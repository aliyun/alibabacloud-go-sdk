// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableChangeLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableChangeLogResponseBodyData) *GetMetaTableChangeLogResponseBody
	GetData() *GetMetaTableChangeLogResponseBodyData
	SetErrorCode(v string) *GetMetaTableChangeLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableChangeLogResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableChangeLogResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableChangeLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableChangeLogResponseBody
	GetSuccess() *bool
}

type GetMetaTableChangeLogResponseBody struct {
	// The business data.
	Data *GetMetaTableChangeLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
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

func (s GetMetaTableChangeLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableChangeLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableChangeLogResponseBody) GetData() *GetMetaTableChangeLogResponseBodyData {
	return s.Data
}

func (s *GetMetaTableChangeLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableChangeLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableChangeLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableChangeLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableChangeLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableChangeLogResponseBody) SetData(v *GetMetaTableChangeLogResponseBodyData) *GetMetaTableChangeLogResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableChangeLogResponseBody) SetErrorCode(v string) *GetMetaTableChangeLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBody) SetErrorMessage(v string) *GetMetaTableChangeLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBody) SetHttpStatusCode(v int32) *GetMetaTableChangeLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBody) SetRequestId(v string) *GetMetaTableChangeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBody) SetSuccess(v bool) *GetMetaTableChangeLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableChangeLogResponseBodyData struct {
	// The list of instances.
	DataEntityList []*GetMetaTableChangeLogResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of metatables.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaTableChangeLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableChangeLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableChangeLogResponseBodyData) GetDataEntityList() []*GetMetaTableChangeLogResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *GetMetaTableChangeLogResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTableChangeLogResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableChangeLogResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaTableChangeLogResponseBodyData) SetDataEntityList(v []*GetMetaTableChangeLogResponseBodyDataDataEntityList) *GetMetaTableChangeLogResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyData) SetPageNumber(v int32) *GetMetaTableChangeLogResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyData) SetPageSize(v int32) *GetMetaTableChangeLogResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyData) SetTotalCount(v int64) *GetMetaTableChangeLogResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableChangeLogResponseBodyDataDataEntityList struct {
	// The content of the change.
	//
	// example:
	//
	// "[{\\"action\\":\\"ADD_COLUMN\\",\\"value\\":[{\\"originName\\":\\"\\",\\"originType\\":\\"\\",\\"originComment\\":\\"\\",\\"name\\":\\"id\\",\\"type\\":\\"struct<name:string>\\",\\"comment\\":\\"\\"}]}]",
	ChangeContent *string `json:"ChangeContent,omitempty" xml:"ChangeContent,omitempty"`
	// The type of the change.
	//
	// example:
	//
	// CREATE_TABLE
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The time when the metatable was created.
	//
	// example:
	//
	// 1590722845000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the metatable was modified.
	//
	// example:
	//
	// 1590722845000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The entity on which the change was made. Valid values: TABLE and PARTITION.
	//
	// example:
	//
	// TABLE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// abc
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GetMetaTableChangeLogResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableChangeLogResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) GetChangeContent() *string {
	return s.ChangeContent
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) GetChangeType() *string {
	return s.ChangeType
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) GetOperator() *string {
	return s.Operator
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) SetChangeContent(v string) *GetMetaTableChangeLogResponseBodyDataDataEntityList {
	s.ChangeContent = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) SetChangeType(v string) *GetMetaTableChangeLogResponseBodyDataDataEntityList {
	s.ChangeType = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) SetCreateTime(v int64) *GetMetaTableChangeLogResponseBodyDataDataEntityList {
	s.CreateTime = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) SetModifiedTime(v int64) *GetMetaTableChangeLogResponseBodyDataDataEntityList {
	s.ModifiedTime = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) SetObjectType(v string) *GetMetaTableChangeLogResponseBodyDataDataEntityList {
	s.ObjectType = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) SetOperator(v string) *GetMetaTableChangeLogResponseBodyDataDataEntityList {
	s.Operator = &v
	return s
}

func (s *GetMetaTableChangeLogResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
