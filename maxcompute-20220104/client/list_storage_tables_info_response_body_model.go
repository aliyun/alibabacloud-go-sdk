// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageTablesInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListStorageTablesInfoResponseBodyData) *ListStorageTablesInfoResponseBody
	GetData() *ListStorageTablesInfoResponseBodyData
	SetErrorCode(v string) *ListStorageTablesInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListStorageTablesInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListStorageTablesInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListStorageTablesInfoResponseBody
	GetRequestId() *string
}

type ListStorageTablesInfoResponseBody struct {
	Data      *ListStorageTablesInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                                 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStorageTablesInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBody) GetData() *ListStorageTablesInfoResponseBodyData {
	return s.Data
}

func (s *ListStorageTablesInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStorageTablesInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListStorageTablesInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListStorageTablesInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStorageTablesInfoResponseBody) SetData(v *ListStorageTablesInfoResponseBodyData) *ListStorageTablesInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetErrorCode(v string) *ListStorageTablesInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetErrorMsg(v string) *ListStorageTablesInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetHttpCode(v int32) *ListStorageTablesInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetRequestId(v string) *ListStorageTablesInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStorageTablesInfoResponseBodyData struct {
	Date                 *string                                                      `json:"date,omitempty" xml:"date,omitempty"`
	PageNumber           *int64                                                       `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize             *int64                                                       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StorageTableInfoList []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList `json:"storageTableInfoList,omitempty" xml:"storageTableInfoList,omitempty" type:"Repeated"`
	TotalCount           *int64                                                       `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStorageTablesInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *ListStorageTablesInfoResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStorageTablesInfoResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStorageTablesInfoResponseBodyData) GetStorageTableInfoList() []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	return s.StorageTableInfoList
}

func (s *ListStorageTablesInfoResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStorageTablesInfoResponseBodyData) SetDate(v string) *ListStorageTablesInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetPageNumber(v int64) *ListStorageTablesInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetPageSize(v int64) *ListStorageTablesInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetStorageTableInfoList(v []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList) *ListStorageTablesInfoResponseBodyData {
	s.StorageTableInfoList = v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetTotalCount(v int64) *ListStorageTablesInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) Validate() error {
	if s.StorageTableInfoList != nil {
		for _, item := range s.StorageTableInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStorageTablesInfoResponseBodyDataStorageTableInfoList struct {
	Date                     *string  `json:"date,omitempty" xml:"date,omitempty"`
	IsPartitioned            *bool    `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	LastAccessTime           *int64   `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	LongTermStorage          *float64 `json:"longTermStorage,omitempty" xml:"longTermStorage,omitempty"`
	LongTermStorageFileCount *int64   `json:"longTermStorageFileCount,omitempty" xml:"longTermStorageFileCount,omitempty"`
	LongTermStorageUnit      *string  `json:"longTermStorageUnit,omitempty" xml:"longTermStorageUnit,omitempty"`
	LowFreqStorage           *float64 `json:"lowFreqStorage,omitempty" xml:"lowFreqStorage,omitempty"`
	LowFreqStorageFileCount  *int64   `json:"lowFreqStorageFileCount,omitempty" xml:"lowFreqStorageFileCount,omitempty"`
	LowFreqStorageUnit       *string  `json:"lowFreqStorageUnit,omitempty" xml:"lowFreqStorageUnit,omitempty"`
	ProjectName              *string  `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Rate                     *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	SchemaName               *string  `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	StandardStorage          *float64 `json:"standardStorage,omitempty" xml:"standardStorage,omitempty"`
	StandardStorageFileCount *int64   `json:"standardStorageFileCount,omitempty" xml:"standardStorageFileCount,omitempty"`
	StandardStorageUnit      *string  `json:"standardStorageUnit,omitempty" xml:"standardStorageUnit,omitempty"`
	StorageType              *string  `json:"storageType,omitempty" xml:"storageType,omitempty"`
	TableName                *string  `json:"tableName,omitempty" xml:"tableName,omitempty"`
	TotalFrequency           *int64   `json:"totalFrequency,omitempty" xml:"totalFrequency,omitempty"`
	TotalInputAmount         *float64 `json:"totalInputAmount,omitempty" xml:"totalInputAmount,omitempty"`
	TotalInputAmountUnit     *string  `json:"totalInputAmountUnit,omitempty" xml:"totalInputAmountUnit,omitempty"`
	TotalStorage             *float64 `json:"totalStorage,omitempty" xml:"totalStorage,omitempty"`
	TotalStorageFileCount    *int64   `json:"totalStorageFileCount,omitempty" xml:"totalStorageFileCount,omitempty"`
	TotalStorageUnit         *string  `json:"totalStorageUnit,omitempty" xml:"totalStorageUnit,omitempty"`
}

func (s ListStorageTablesInfoResponseBodyDataStorageTableInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetDate() *string {
	return s.Date
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetIsPartitioned() *bool {
	return s.IsPartitioned
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLongTermStorage() *float64 {
	return s.LongTermStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLongTermStorageFileCount() *int64 {
	return s.LongTermStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLongTermStorageUnit() *string {
	return s.LongTermStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLowFreqStorage() *float64 {
	return s.LowFreqStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLowFreqStorageFileCount() *int64 {
	return s.LowFreqStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLowFreqStorageUnit() *string {
	return s.LowFreqStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetRate() *float64 {
	return s.Rate
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStandardStorage() *float64 {
	return s.StandardStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStandardStorageFileCount() *int64 {
	return s.StandardStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStandardStorageUnit() *string {
	return s.StandardStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTableName() *string {
	return s.TableName
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalFrequency() *int64 {
	return s.TotalFrequency
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalInputAmount() *float64 {
	return s.TotalInputAmount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalInputAmountUnit() *string {
	return s.TotalInputAmountUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalStorage() *float64 {
	return s.TotalStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalStorageFileCount() *int64 {
	return s.TotalStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalStorageUnit() *string {
	return s.TotalStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetDate(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetIsPartitioned(v bool) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.IsPartitioned = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLastAccessTime(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LastAccessTime = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetProjectName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetRate(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.Rate = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetSchemaName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.SchemaName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStorageType(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StorageType = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTableName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TableName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalFrequency(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalFrequency = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalInputAmount(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalInputAmount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalInputAmountUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalInputAmountUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) Validate() error {
	return dara.Validate(s)
}
