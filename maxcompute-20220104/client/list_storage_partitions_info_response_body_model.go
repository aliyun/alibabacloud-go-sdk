// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStoragePartitionsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListStoragePartitionsInfoResponseBodyData) *ListStoragePartitionsInfoResponseBody
	GetData() *ListStoragePartitionsInfoResponseBodyData
	SetErrorCode(v string) *ListStoragePartitionsInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListStoragePartitionsInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListStoragePartitionsInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListStoragePartitionsInfoResponseBody
	GetRequestId() *string
}

type ListStoragePartitionsInfoResponseBody struct {
	Data      *ListStoragePartitionsInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                                    `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                    `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                                     `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBody) GetData() *ListStoragePartitionsInfoResponseBodyData {
	return s.Data
}

func (s *ListStoragePartitionsInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStoragePartitionsInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListStoragePartitionsInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListStoragePartitionsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStoragePartitionsInfoResponseBody) SetData(v *ListStoragePartitionsInfoResponseBodyData) *ListStoragePartitionsInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetErrorCode(v string) *ListStoragePartitionsInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetErrorMsg(v string) *ListStoragePartitionsInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetHttpCode(v int32) *ListStoragePartitionsInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetRequestId(v string) *ListStoragePartitionsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStoragePartitionsInfoResponseBodyData struct {
	Date                     *string                                                              `json:"date,omitempty" xml:"date,omitempty"`
	PageNumber               *int64                                                               `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize                 *int64                                                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StoragePartitionInfoList []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList `json:"storagePartitionInfoList,omitempty" xml:"storagePartitionInfoList,omitempty" type:"Repeated"`
	TotalCount               *int64                                                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetStoragePartitionInfoList() []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	return s.StoragePartitionInfoList
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetDate(v string) *ListStoragePartitionsInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetPageNumber(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetPageSize(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetStoragePartitionInfoList(v []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) *ListStoragePartitionsInfoResponseBodyData {
	s.StoragePartitionInfoList = v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetTotalCount(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) Validate() error {
	if s.StoragePartitionInfoList != nil {
		for _, item := range s.StoragePartitionInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList struct {
	FileCount            *int64   `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	FileSize             *float64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	FileSizeUnit         *string  `json:"fileSizeUnit,omitempty" xml:"fileSizeUnit,omitempty"`
	IsPartitioned        *bool    `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	LastAccessTime       *int64   `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	Partition            *string  `json:"partition,omitempty" xml:"partition,omitempty"`
	ProjectName          *string  `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Rate                 *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	SchemaName           *string  `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	StorageType          *string  `json:"storageType,omitempty" xml:"storageType,omitempty"`
	TableName            *string  `json:"tableName,omitempty" xml:"tableName,omitempty"`
	TotalFrequency       *int64   `json:"totalFrequency,omitempty" xml:"totalFrequency,omitempty"`
	TotalInputAmount     *float64 `json:"totalInputAmount,omitempty" xml:"totalInputAmount,omitempty"`
	TotalInputAmountUnit *string  `json:"totalInputAmountUnit,omitempty" xml:"totalInputAmountUnit,omitempty"`
	Type                 *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetFileCount() *int64 {
	return s.FileCount
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetFileSize() *float64 {
	return s.FileSize
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetFileSizeUnit() *string {
	return s.FileSizeUnit
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetIsPartitioned() *bool {
	return s.IsPartitioned
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetPartition() *string {
	return s.Partition
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetRate() *float64 {
	return s.Rate
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTableName() *string {
	return s.TableName
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTotalFrequency() *int64 {
	return s.TotalFrequency
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTotalInputAmount() *float64 {
	return s.TotalInputAmount
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTotalInputAmountUnit() *string {
	return s.TotalInputAmountUnit
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetType() *string {
	return s.Type
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileCount(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileCount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileSize(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileSize = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileSizeUnit(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileSizeUnit = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetIsPartitioned(v bool) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.IsPartitioned = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetLastAccessTime(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.LastAccessTime = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetPartition(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Partition = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetProjectName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetRate(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Rate = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetSchemaName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.SchemaName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetStorageType(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.StorageType = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTableName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TableName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalFrequency(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalFrequency = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalInputAmount(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalInputAmount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalInputAmountUnit(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalInputAmountUnit = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetType(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Type = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) Validate() error {
	return dara.Validate(s)
}
