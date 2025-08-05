// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageProjectsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListStorageProjectsInfoResponseBodyData) *ListStorageProjectsInfoResponseBody
	GetData() *ListStorageProjectsInfoResponseBodyData
	SetErrorCode(v string) *ListStorageProjectsInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListStorageProjectsInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListStorageProjectsInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListStorageProjectsInfoResponseBody
	GetRequestId() *string
}

type ListStorageProjectsInfoResponseBody struct {
	Data *ListStorageProjectsInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0bc3b4b016674434996033675e71ee
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStorageProjectsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStorageProjectsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStorageProjectsInfoResponseBody) GetData() *ListStorageProjectsInfoResponseBodyData {
	return s.Data
}

func (s *ListStorageProjectsInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStorageProjectsInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListStorageProjectsInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListStorageProjectsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStorageProjectsInfoResponseBody) SetData(v *ListStorageProjectsInfoResponseBodyData) *ListStorageProjectsInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStorageProjectsInfoResponseBody) SetErrorCode(v string) *ListStorageProjectsInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBody) SetErrorMsg(v string) *ListStorageProjectsInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBody) SetHttpCode(v int32) *ListStorageProjectsInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBody) SetRequestId(v string) *ListStorageProjectsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStorageProjectsInfoResponseBodyData struct {
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize               *int64                                                           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StorageProjectInfoList []*ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList `json:"storageProjectInfoList,omitempty" xml:"storageProjectInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStorageProjectsInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStorageProjectsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStorageProjectsInfoResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *ListStorageProjectsInfoResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStorageProjectsInfoResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStorageProjectsInfoResponseBodyData) GetStorageProjectInfoList() []*ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	return s.StorageProjectInfoList
}

func (s *ListStorageProjectsInfoResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStorageProjectsInfoResponseBodyData) SetDate(v string) *ListStorageProjectsInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyData) SetPageNumber(v int64) *ListStorageProjectsInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyData) SetPageSize(v int64) *ListStorageProjectsInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyData) SetStorageProjectInfoList(v []*ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) *ListStorageProjectsInfoResponseBodyData {
	s.StorageProjectInfoList = v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyData) SetTotalCount(v int64) *ListStorageProjectsInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList struct {
	// example:
	//
	// 20250528
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 300
	LongTermStorage *float64 `json:"longTermStorage,omitempty" xml:"longTermStorage,omitempty"`
	// example:
	//
	// GB
	LongTermStorageUnit *string `json:"longTermStorageUnit,omitempty" xml:"longTermStorageUnit,omitempty"`
	// example:
	//
	// 200
	LowFreqStorage *float64 `json:"lowFreqStorage,omitempty" xml:"lowFreqStorage,omitempty"`
	// example:
	//
	// GB
	LowFreqStorageUnit *string `json:"lowFreqStorageUnit,omitempty" xml:"lowFreqStorageUnit,omitempty"`
	// example:
	//
	// max_testproject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// 0.011872406445069006
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// example:
	//
	// 0
	RecycleBinStorage *float64 `json:"recycleBinStorage,omitempty" xml:"recycleBinStorage,omitempty"`
	// example:
	//
	// B
	RecycleBinStorageUnit *string `json:"recycleBinStorageUnit,omitempty" xml:"recycleBinStorageUnit,omitempty"`
	// example:
	//
	// 500
	StandardStorage *float64 `json:"standardStorage,omitempty" xml:"standardStorage,omitempty"`
	// example:
	//
	// GB
	StandardStorageUnit *string `json:"standardStorageUnit,omitempty" xml:"standardStorageUnit,omitempty"`
	// example:
	//
	// 1749105045512
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 1
	TotalStorage *float64 `json:"totalStorage,omitempty" xml:"totalStorage,omitempty"`
	// example:
	//
	// TB
	TotalStorageUnit *string `json:"totalStorageUnit,omitempty" xml:"totalStorageUnit,omitempty"`
}

func (s ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GoString() string {
	return s.String()
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetDate() *string {
	return s.Date
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetLongTermStorage() *float64 {
	return s.LongTermStorage
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetLongTermStorageUnit() *string {
	return s.LongTermStorageUnit
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetLowFreqStorage() *float64 {
	return s.LowFreqStorage
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetLowFreqStorageUnit() *string {
	return s.LowFreqStorageUnit
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetRate() *float64 {
	return s.Rate
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetRecycleBinStorage() *float64 {
	return s.RecycleBinStorage
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetRecycleBinStorageUnit() *string {
	return s.RecycleBinStorageUnit
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetStandardStorage() *float64 {
	return s.StandardStorage
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetStandardStorageUnit() *string {
	return s.StandardStorageUnit
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetTotalStorage() *float64 {
	return s.TotalStorage
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) GetTotalStorageUnit() *string {
	return s.TotalStorageUnit
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetDate(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.Date = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetLongTermStorage(v float64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.LongTermStorage = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetLongTermStorageUnit(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.LongTermStorageUnit = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetLowFreqStorage(v float64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.LowFreqStorage = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetLowFreqStorageUnit(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.LowFreqStorageUnit = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetProjectName(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetRate(v float64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.Rate = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetRecycleBinStorage(v float64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.RecycleBinStorage = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetRecycleBinStorageUnit(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.RecycleBinStorageUnit = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetStandardStorage(v float64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.StandardStorage = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetStandardStorageUnit(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.StandardStorageUnit = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetTimestamp(v int64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.Timestamp = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetTotalStorage(v float64) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.TotalStorage = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) SetTotalStorageUnit(v string) *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList {
	s.TotalStorageUnit = &v
	return s
}

func (s *ListStorageProjectsInfoResponseBodyDataStorageProjectInfoList) Validate() error {
	return dara.Validate(s)
}
