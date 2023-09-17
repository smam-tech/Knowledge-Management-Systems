'use client';

import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import { CoreAPIGET } from '@/dep/core/coreHandler';
import { ItmsPerPageComp, PaginationComp } from '@/components/PaginationControls';

function HistoryTable() {
  const [data, setData] = useState([]);
  const [usNames, setUsNames] = useState({});
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(20);
  const [pageInfo, setPageInfo] = useState({ TotalPage: 0 });

  const updateUsNames = async (historyData) => {
    const fetchUsName = async (userID) => {
      try {
        const responseUs = await CoreAPIGET(`user?UserID=${userID}`);
        return responseUs.body.Data.Name;
      } catch (error) {
        console.error('Error fetching user name:', error);
        return null;
      }
    };

    const usNamesMap = {};
    for (const history of historyData) {
      if (!usNamesMap[history.UserID]) {
        usNamesMap[history.UserID] = await fetchUsName(history.UserID);
      }
    }
    setUsNames(usNamesMap);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await CoreAPIGET(`listhistory?page=${currentPage}&num=${itemsPerPage}`);
        const reversedData = response.body.Data.reverse();
        const pageInfo = response.body.Info;
        setData(reversedData);
        setPageInfo(pageInfo);
        if (reversedData.length > 0) {
          updateUsNames(reversedData);
        }
      } catch (error) {
        console.error('Error fetching history data:', error);
      }
    };

    fetchData();
  }, [itemsPerPage, currentPage]);

  const handlePageChange = (pageNumber) => {
    setCurrentPage(pageNumber);
  };
  return (
    <section className="h-screen flex flex-col flex-auto">
      <div className="flex flex-col">
        <div className="px-4 py-2 bg-gray-300 tracking-wide font-medium shadow rounded mb-4">
          <Link href="/activity-log" className="text-blue-500 hover:text-blue-800">
            Activity Log
          </Link>
        </div>
        <div className="max-w-3xl mx-auto p-4">

          <table className="w-full border">
            <thead>
              <tr className="bg-gray-100">
                <th className="px-4 py-2">Type</th>
                <th className="px-4 py-2">Changes</th>
                <th className="px-4 py-2">User</th>
                <th className="px-4 py-2">Time</th>
              </tr>
            </thead>
            <tbody>
              {data.map((history) => (
                <tr key={history.HistoryID} className="border-b">
                  <td className="px-4 py-2">{history.ActivityType}</td>
                  <td className="px-4 py-2">{history.Changes}</td>
                  <td className="px-4 py-2">
                    {' '}
                    {usNames[history.UserID] || '-'}
                  </td>
                  <td className="px-4 py-2">{history.Time}</td>
                </tr>
              ))}
            </tbody>
          </table>
          <PaginationComp
            currentPage={currentPage}
            totalPages={pageInfo.TotalPage}
            totalRow={pageInfo.TotalRow}
            itemsPerPage={itemsPerPage}
            handlePageChange={handlePageChange}
            upperLimit={pageInfo.UpperLimit}
            lowerLimit={pageInfo.LowerLimit}
          />

          <ItmsPerPageComp
            itemsPerPage={itemsPerPage}
            setItemsPerPage={setItemsPerPage}
            setCurrentPage={setCurrentPage}
          />
        </div>
      </div>
    </section>
  );
}

export default HistoryTable;
