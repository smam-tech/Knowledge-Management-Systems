package kms

import (
	"dependency"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Article_API struct {
	ArticleID            int `json:"ArticleID" query:"ArticleID"`
	OwnerID              int
	OwnerUsername        string
	OwnerName            string
	LastEditedByID       int
	LastEditedByUsername string
	LastEditedByName     string
	LastEditedTime       time.Time
	Tag                  []string
	Title                string
	CategoryID           int
	CategoryName         string
	CategoryParent       string
	CategoryDescription  string
	Article              string
	FileID               []int
	DocID                []int
	IsActive             bool
}

type Article_GrapesJS struct {
	StatusCode int
	Id         int                    `json:"id" query:"id"`
	Data       string                 `json:"data"`
	Dumpvalue  map[string]interface{} `json:"dump"`
}

type GrapesJS_Data struct {
	Data      map[string]interface{}   `json:"data"`
	PagesHtml []map[string]interface{} `json:"pagesHtml"`
}

func ListArticle(c echo.Context) error {
	var LimitQuery string
	var ValuesQuery []interface{}

	permission, _, _ := Check_Admin_Permission_API(c)
	res := ResponseList{}
	limit := new(dependency.QueryType)
	err := c.Bind(limit)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	if permission {
		LimitQuery, ValuesQuery, res.Info, err = limit.QueryMaker(Database, "kms_article")
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		ArticleList, _ := ReadArticle(LimitQuery, ValuesQuery)
		var ArticleListAPI []Article_API
		for _, x := range ArticleList {
			tmp, err := x.ToAPI(c)
			if err != nil {
				Logger.Error(err.Error())
				res.StatusCode = http.StatusInternalServerError
				res.Data = err
				return c.JSON(http.StatusInternalServerError, res)
			}
			ArticleListAPI = append(ArticleListAPI, tmp)
		}
		res.StatusCode = http.StatusOK
		res.Data = ArticleListAPI
		return c.JSON(http.StatusOK, res)
	} else {
		AllowedCategoryList, err := GetCurrentUserReadCategoryList(c)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		var wherequery []dependency.WhereType
		if limit.Query != "" {
			err = json.Unmarshal([]byte(limit.Query), &wherequery)
			if err != nil {
				err = errors.New("query field json read error : " + err.Error())
				Logger.Error(err.Error())
				res.StatusCode = http.StatusInternalServerError
				res.Data = err
				return c.JSON(http.StatusInternalServerError, res)
			}
		}
		var convertedAllowedCategoryList []interface{}
		for _, v := range AllowedCategoryList {
			convertedAllowedCategoryList = append(convertedAllowedCategoryList, v)
		}
		singlewherequery := dependency.WhereType{
			Field:    "CategoryID",
			Operator: "IN",
			Logic:    "AND",
			Values:   convertedAllowedCategoryList,
		}
		wherequery = append(wherequery, singlewherequery)
		a, err := json.Marshal(wherequery)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		limit.Query = string(a)
		LimitQuery, ValuesQuery, res.Info, err = limit.QueryMaker(Database, "kms_article")
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		ArticleList, _ := ReadArticle(LimitQuery, ValuesQuery)
		var ArticleListAPI []Article_API
		for _, x := range ArticleList {
			tmp, err := x.ToAPI(c)
			if err != nil {
				Logger.Error(err.Error())
				res.StatusCode = http.StatusInternalServerError
				res.Data = err
				return c.JSON(http.StatusInternalServerError, res)
			}
			ArticleListAPI = append(ArticleListAPI, tmp)
		}
		res.StatusCode = http.StatusOK
		res.Data = ArticleListAPI
		return c.JSON(http.StatusOK, res)
	}
}

func ListArticleID(c echo.Context) error {
	var LimitQuery string
	var ValuesQuery []interface{}

	permission, _, _ := Check_Admin_Permission_API(c)
	res := ResponseList{}
	limit := new(dependency.QueryType)
	err := c.Bind(limit)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	if permission {
		LimitQuery, ValuesQuery, res.Info, err = limit.QueryMaker(Database, "kms_article")
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		ArticleList, _ := ReadArticleID(LimitQuery, ValuesQuery)
		res.StatusCode = http.StatusOK
		res.Data = ArticleList
		return c.JSON(http.StatusOK, res)
	} else {
		AllowedCategoryList, err := GetCurrentUserReadCategoryList(c)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		var wherequery []dependency.WhereType
		if limit.Query != "" {
			err = json.Unmarshal([]byte(limit.Query), &wherequery)
			if err != nil {
				err = errors.New("query field json read error : " + err.Error())
				Logger.Error(err.Error())
				res.StatusCode = http.StatusInternalServerError
				res.Data = err
				return c.JSON(http.StatusInternalServerError, res)
			}
		}
		var convertedAllowedCategoryList []interface{}
		for _, v := range AllowedCategoryList {
			convertedAllowedCategoryList = append(convertedAllowedCategoryList, v)
		}
		singlewherequery := dependency.WhereType{
			Field:    "CategoryID",
			Operator: "IN",
			Logic:    "AND",
			Values:   convertedAllowedCategoryList,
		}
		wherequery = append(wherequery, singlewherequery)
		a, err := json.Marshal(wherequery)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		limit.Query = string(a)
		LimitQuery, ValuesQuery, res.Info, err = limit.QueryMaker(Database, "kms_article")
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		ArticleList, _ := ReadArticleID(LimitQuery, ValuesQuery)
		res.StatusCode = http.StatusOK
		res.Data = ArticleList
		return c.JSON(http.StatusOK, res)
	}
}

func ShowArticle(c echo.Context) error {
	var err error
	var res Response
	u := new(Article_API)
	err = c.Bind(u)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	uOri, err := u.ToTable()
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	err = uOri.Read()
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "ARTICLE NOT FOUND"
		return c.JSON(http.StatusBadRequest, res)
	}
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	_, TrueRead, TrueUpdate, _, err := GetTruePermission(c, uOri.CategoryID, role_id)
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	permission, _, _ := Check_Admin_Permission_API(c)
	if TrueRead || TrueUpdate || permission {
		*u, err = uOri.ToAPI(c)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.StatusCode = http.StatusOK
		res.Data = *u
		return c.JSON(http.StatusOK, res)
	} else {
		res.StatusCode = http.StatusForbidden
		res.Data = "YOU DONT HAVE PERMISSION TO READ THIS ARTICLE"
		return c.JSON(http.StatusForbidden, res)
	}
}

func AddArticle(c echo.Context) error {
	var err error
	var res Response
	u := new(Article_API)
	err = c.Bind(u)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	TrueCreate, _, _, _, err := GetTruePermission(c, u.CategoryID, role_id)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusForbidden
		res.Data = err
		return c.JSON(http.StatusForbidden, res)
	}
	permission, _, _ := Check_Admin_Permission_API(c)
	if TrueCreate || permission {
		uOri, err := u.ToTable()
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		_, err = uOri.Create()
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusConflict
			res.Data = "CREATE ERROR : " + err.Error()
			return c.JSON(http.StatusConflict, res)
		}
		// resulta, err := u.ConvForSolr()
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "UPDATE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		// err = resulta.PrepareSolrData(c)
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "UPDATE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		// _, _, err = SolrCallUpdate("POST", resulta)
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "UPDATE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		*u, err = uOri.ToAPI(c)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.StatusCode = http.StatusOK
		res.Data = *u
		err = RecordHistory(c, "Article", "Added Article : "+u.Title+"("+strconv.Itoa(u.ArticleID)+")")
		if err != nil {
			Logger.Error("failed to record article change history " + err.Error())
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res.StatusCode = http.StatusForbidden
		res.Data = "YOU DONT HAVE PERMISSION TO CREATE ARTICLE IN THIS CATEGORY"
		return c.JSON(http.StatusForbidden, res)
	}
}

func EditArticle(c echo.Context) error {
	var err error
	var res Response
	u := new(Article_API)
	err = c.Bind(u)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	oriu, err := u.ToTable()
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	err = oriu.Read()
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "ORIGINAL NOT FOUND"
		return c.JSON(http.StatusBadRequest, res)
	}
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	_, _, TrueUpdate, _, err := GetTruePermission(c, oriu.CategoryID, role_id)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusForbidden
		res.Data = err
		return c.JSON(http.StatusForbidden, res)
	}
	permission, _, _ := Check_Admin_Permission_API(c)
	if TrueUpdate || permission {
		// resulta, err := u.ConvForSolr()
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "UPDATE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		// err = resulta.PrepareSolrData(c)
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "UPDATE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		// _, _, err = SolrCallUpdate("POST", resulta)
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "UPDATE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		oriu, err = u.ToTable()
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		err = oriu.Update()
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = "UPDATE ERROR : " + err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		*u, err = oriu.ToAPI(c)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.StatusCode = http.StatusOK
		res.Data = *u
		err = RecordHistory(c, "Article", "Edited Article : "+u.Title+"("+strconv.Itoa(u.ArticleID)+")")
		if err != nil {
			Logger.Error("failed to record article change history " + err.Error())
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res.StatusCode = http.StatusForbidden
		res.Data = "YOU DONT HAVE PERMISSION TO EDIT THIS ARTICLE"
		return c.JSON(http.StatusForbidden, res)
	}
}

func DeleteArticle(c echo.Context) error {
	var err error
	var res Response
	u := new(Article_API)
	err = c.Bind(u)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	oriu, err := u.ToTable()
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	err = oriu.Read()
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "ORIGINAL ARTICLE NOT FOUND"
		return c.JSON(http.StatusBadRequest, res)
	}
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	_, _, _, TrueDelete, err := GetTruePermission(c, oriu.CategoryID, role_id)
	// err = permission_kms.Read()
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusForbidden
		res.Data = err
		return c.JSON(http.StatusForbidden, res)
	}
	permission, _, _ := Check_Admin_Permission_API(c)
	if TrueDelete || permission {
		err = oriu.Delete()
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = "DELETE ERROR : " + err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		// err = DeleteSolrDocument(strconv.Itoa(u.ArticleID))
		// if err != nil {
		// 	res.StatusCode = http.StatusBadRequest
		// 	res.Data = "DELETE ERROR : " + err.Error()
		// 	return c.JSON(http.StatusBadRequest, res)
		// }
		res.StatusCode = http.StatusOK
		res.Data = u
		err = RecordHistory(c, "Article", "Deleted Article : "+u.Title+"("+strconv.Itoa(u.ArticleID)+")")
		if err != nil {
			Logger.Error("failed to record article change history " + err.Error())
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res.StatusCode = http.StatusForbidden
		res.Data = "YOU DONT HAVE PERMISSION TO DELETE THIS ARTICLE"
		return c.JSON(http.StatusForbidden, res)
	}
}

func QueryArticle(c echo.Context) error {
	var err error
	var res Response
	query := c.QueryParam("query")
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	q := c.QueryParam("q")
	search := c.QueryParam("search")

	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	CategoryIDList, err := GetReadCategoryList(c, role_id)
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	response, _, err := SolrCallQuery(CategoryIDList, q, query, search)
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.StatusCode = http.StatusOK
	res.Data = string(response)
	return c.JSON(http.StatusOK, res)
}

func UpdateArticleGrapesjs(c echo.Context) error {
	var err error
	var res Article_GrapesJS
	u := new(Article_GrapesJS)
	err = c.Bind(u)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	u.Data, err = dependency.MapToJson(u.Dumpvalue)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}
	// fmt.Println("BEGIN PRINTING")
	// fmt.Println("ID : " + strconv.Itoa(u.Id))
	// fmt.Println("Data : " + u.Data)
	// fmt.Println(u.Dumpvalue)
	// fmt.Println(user)
	checkArticle := Article_Table{ArticleID: u.Id}
	err = checkArticle.ReadShort()
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusNotFound
		res.Data = "Article Check Error : " + err.Error()
		return c.JSON(http.StatusNotFound, res)
	}

	_, _, TrueUpdate, _, err := GetTruePermission(c, checkArticle.CategoryID, role_id)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusForbidden
		res.Data = err.Error()
		return c.JSON(http.StatusForbidden, res)
	}
	permission, _, _ := Check_Admin_Permission_API(c)
	if TrueUpdate || permission {
		updateArticle := Article_Table{
			ArticleID:      u.Id,
			OwnerID:        0,
			LastEditedByID: 0,
			LastEditedTime: time.Time{},
			Tag:            "",
			Title:          "",
			CategoryID:     0,
			Article:        u.Data,
			FileID:         "",
			DocID:          "",
			IsActive:       0,
		}
		err = updateArticle.UpdateArticleOnly()
		if err != nil {
			Logger.Warn(err.Error())
			res.StatusCode = http.StatusBadRequest
			res.Data = "UPDATE ERROR : " + err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		res.StatusCode = http.StatusOK
		return c.JSON(http.StatusOK, res)
	} else {
		res.StatusCode = http.StatusForbidden
		res.Data = "YOU DONT HAVE PERMISSION TO EDIT THIS ARTICLE"
		return c.JSON(http.StatusForbidden, res)
	}
}

func ReadArticleGrapesjs(c echo.Context) error {
	var err error
	var res Article_GrapesJS
	u := new(Article_GrapesJS)
	err = c.Bind(u)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusBadRequest
		res.Data = "DATA INPUT ERROR : " + err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	_, user, _ := Check_Admin_Permission_API(c)
	role_id, err := dependency.InterfaceToInt(user["RoleID"])
	if err != nil {
		Logger.Error(err.Error())
		res.StatusCode = http.StatusInternalServerError
		res.Data = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}
	fmt.Println(u)
	fmt.Println(user)
	checkArticle := Article_Table{ArticleID: u.Id}
	err = checkArticle.Read()
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusNotFound
		res.Data = "Article Check Error : " + err.Error()
		return c.JSON(http.StatusNotFound, res)
	}
	_, TrueRead, TrueUpdate, _, err := GetTruePermission(c, checkArticle.CategoryID, role_id)
	if err != nil {
		Logger.Warn(err.Error())
		res.StatusCode = http.StatusForbidden
		res.Data = err.Error()
		return c.JSON(http.StatusForbidden, res)
	}
	permission, _, _ := Check_Admin_Permission_API(c)
	if TrueRead || TrueUpdate || permission {
		tmp := checkArticle.Article
		tmp_divided := GrapesJS_Data{}
		err = json.Unmarshal([]byte(tmp), &tmp_divided)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err.Error()
			return c.JSON(http.StatusInternalServerError, res)
		}
		tmp, err = dependency.MapToJson(tmp_divided.Data)
		if err != nil {
			Logger.Error(err.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Data = err.Error()
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.Data = tmp
		res.Id = checkArticle.ArticleID
		res.StatusCode = http.StatusOK
		return c.JSON(http.StatusOK, res)
	} else {
		res.StatusCode = http.StatusForbidden
		res.Data = "YOU DONT HAVE PERMISSION TO EDIT THIS ARTICLE"
		return c.JSON(http.StatusForbidden, res)
	}
}

func (data Article_Table) ToAPI(c echo.Context) (res Article_API, err error) {
	res.Article = data.Article
	res.ArticleID = data.ArticleID
	res.CategoryID = data.CategoryID
	res.DocID, err = dependency.ConvStringToIntArray(data.DocID)
	if err != nil {
		return res, err
	}
	res.FileID, err = dependency.ConvStringToIntArray(data.FileID)
	if err != nil {
		return res, err
	}
	res.IsActive = data.IsActive == 1
	res.LastEditedByID = data.LastEditedByID
	res.LastEditedTime = data.LastEditedTime
	res.OwnerID = data.OwnerID
	res.Tag = dependency.ConvStringToStringArray(data.Tag)
	res.Title = data.Title
	CategoryData := Category{
		CategoryID:          data.CategoryID,
		CategoryName:        "",
		CategoryParentID:    0,
		CategoryDescription: "",
	}
	err = CategoryData.Read()
	if err != nil {
		return res, err
	}
	res.CategoryName = CategoryData.CategoryName
	res.CategoryDescription = CategoryData.CategoryDescription
	res.OwnerName, res.OwnerUsername, err = GetNameUsername(c, data.OwnerID)
	if err != nil {
		return res, err
	}
	res.LastEditedByName, res.LastEditedByUsername, err = GetNameUsername(c, data.LastEditedByID)
	return res, err
}

func (data Article_API) ToTable() (res Article_Table, err error) {
	res.Article = data.Article
	res.ArticleID = data.ArticleID
	res.CategoryID = data.CategoryID
	res.DocID, err = dependency.ConvIntArrayToStringUnique(data.DocID)
	if err != nil {
		return res, err
	}
	res.FileID, err = dependency.ConvIntArrayToStringUnique(data.FileID)
	if err != nil {
		return res, err
	}
	res.IsActive = dependency.BooltoInt(data.IsActive)
	res.LastEditedByID = data.LastEditedByID
	res.LastEditedTime = data.LastEditedTime
	res.OwnerID = data.OwnerID
	res.Tag, err = dependency.ConvStringArrayToString(data.Tag)
	res.Title = data.Title
	return res, err
}
