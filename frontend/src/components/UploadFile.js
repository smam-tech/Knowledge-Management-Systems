import React, { useState, useRef } from 'react';
import { Button } from '@/components/ui/button';
import { KmsAPIBlob } from '@/dep/kms/kmsHandler';
import { alertAdd } from './Feature';

function UploadFile({ categoryID, FileAdd }) {
  const [selectedFile, setSelectedFile] = useState(null);
  const [currentState, setCurrentState] = useState(1);
  const fileInputRef = useRef(null);
  const handleFileChange = (event) => {
    const file = event.target.files[0];
    console.log(file);
    console.log(typeof (file));
    setSelectedFile(file);
    setCurrentState(2);
  };

  const categoryIdValue = categoryID;

  const handleUpload = async (e) => {
    e.preventDefault();
    console.log('Upload button clicked in UploadFile');
    try {
      const FileSent = selectedFile;
      console.log('before sending API');
      const formData = new FormData();
      formData.append('CategoryID', categoryIdValue);
      formData.append('File', FileSent);
      const response = await KmsAPIBlob('POST', 'file', formData);
      alertAdd(response);
      if (response.status === 200) {
        setSelectedFile(null);
        FileAdd(response.body.Data.FileID);
        setCurrentState(3);
        fileInputRef.current.value = '';
      } else {
        setCurrentState(4);
      }
    } catch (error) {
      console.log(error);
      console.log('An error occurred');
    }
  };

  return (
    <div className="flex flex-col">
      <div className="grid grid-cols-5 grid-rows-2 gap-4">
        <div className="col-span-3">
          <input
            ref={fileInputRef}
            className="block px-2 py-2 w-full md:w-3/4 text-sm text-gray-900 border border-gray-300 rounded cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
            accept="*/*"
            type="file"
            onChange={handleFileChange}
          />
        </div>
        <div className="col-start-4 col-span-2">
          <Button
            type="button"
            onClick={handleUpload}
            className="rounded bg-blue-500 text-white w-full md:w-36"
          >
            Upload File
          </Button>
        </div>
        <div className="col-span-5 row-start-2">
          <p className="text-xs mt-1 mb-4">
            Upload a File from your device. Image should be a .jpg, .jpeg, .png file. Notes: to use the uploaded File on this article, you need to click the upload button.
          </p>
          <p>
            Current State :
            {' '}
            {currentState}
          </p>
        </div>
      </div>
    </div>
  );
}

export default UploadFile;
