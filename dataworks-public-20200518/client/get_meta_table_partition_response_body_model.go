// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTablePartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTablePartitionResponseBodyData) *GetMetaTablePartitionResponseBody
	GetData() *GetMetaTablePartitionResponseBodyData
	SetErrorCode(v string) *GetMetaTablePartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTablePartitionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTablePartitionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTablePartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTablePartitionResponseBody
	GetSuccess() *bool
}

type GetMetaTablePartitionResponseBody struct {
	// The returned result.
	Data *GetMetaTablePartitionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTablePartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionResponseBody) GetData() *GetMetaTablePartitionResponseBodyData {
	return s.Data
}

func (s *GetMetaTablePartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTablePartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTablePartitionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTablePartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTablePartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTablePartitionResponseBody) SetData(v *GetMetaTablePartitionResponseBodyData) *GetMetaTablePartitionResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTablePartitionResponseBody) SetErrorCode(v string) *GetMetaTablePartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTablePartitionResponseBody) SetErrorMessage(v string) *GetMetaTablePartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTablePartitionResponseBody) SetHttpStatusCode(v int32) *GetMetaTablePartitionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTablePartitionResponseBody) SetRequestId(v string) *GetMetaTablePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTablePartitionResponseBody) SetSuccess(v bool) *GetMetaTablePartitionResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTablePartitionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTablePartitionResponseBodyData struct {
	// The list of partitions.
	DataEntityList []*GetMetaTablePartitionResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of partitions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaTablePartitionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionResponseBodyData) GetDataEntityList() []*GetMetaTablePartitionResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *GetMetaTablePartitionResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTablePartitionResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTablePartitionResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaTablePartitionResponseBodyData) SetDataEntityList(v []*GetMetaTablePartitionResponseBodyDataDataEntityList) *GetMetaTablePartitionResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *GetMetaTablePartitionResponseBodyData) SetPageNumber(v int32) *GetMetaTablePartitionResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyData) SetPageSize(v int32) *GetMetaTablePartitionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyData) SetTotalCount(v int64) *GetMetaTablePartitionResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaTablePartitionResponseBodyDataDataEntityList struct {
	// The comment.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the partition was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1590032868000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The size of the partition. Unit: bytes.
	//
	// example:
	//
	// 19
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The time when the partition was modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1590032868000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The GUID of the partition.
	//
	// example:
	//
	// odps.engine_name.table_name.pt=20170614
	PartitionGuid *string `json:"PartitionGuid,omitempty" xml:"PartitionGuid,omitempty"`
	// The location of the Hive partition.
	//
	// example:
	//
	// abc
	PartitionLocation *string `json:"PartitionLocation,omitempty" xml:"PartitionLocation,omitempty"`
	// The name of the partition.
	//
	// example:
	//
	// pt=20170614
	PartitionName *string `json:"PartitionName,omitempty" xml:"PartitionName,omitempty"`
	// The path of the partition.
	//
	// example:
	//
	// abc
	PartitionPath *string `json:"PartitionPath,omitempty" xml:"PartitionPath,omitempty"`
	// The type of the partition.
	//
	// example:
	//
	// abc
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// The number of entries in the partition.
	//
	// example:
	//
	// 233
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The unique identifier of the metatable.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetMetaTablePartitionResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetComment() *string {
	return s.Comment
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetDataSize() *int64 {
	return s.DataSize
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetPartitionGuid() *string {
	return s.PartitionGuid
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetPartitionLocation() *string {
	return s.PartitionLocation
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetPartitionName() *string {
	return s.PartitionName
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetPartitionPath() *string {
	return s.PartitionPath
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetPartitionType() *string {
	return s.PartitionType
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetComment(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.Comment = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetCreateTime(v int64) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.CreateTime = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetDataSize(v int64) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.DataSize = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetModifiedTime(v int64) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.ModifiedTime = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetPartitionGuid(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.PartitionGuid = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetPartitionLocation(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.PartitionLocation = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetPartitionName(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.PartitionName = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetPartitionPath(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.PartitionPath = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetPartitionType(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.PartitionType = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetRecordCount(v int64) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.RecordCount = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) SetTableGuid(v string) *GetMetaTablePartitionResponseBodyDataDataEntityList {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTablePartitionResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
