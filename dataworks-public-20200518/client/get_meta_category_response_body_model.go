// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaCategoryResponseBodyData) *GetMetaCategoryResponseBody
	GetData() *GetMetaCategoryResponseBodyData
	SetErrorCode(v string) *GetMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaCategoryResponseBody
	GetSuccess() *bool
}

type GetMetaCategoryResponseBody struct {
	// The business data.
	Data *GetMetaCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaCategoryResponseBody) GetData() *GetMetaCategoryResponseBodyData {
	return s.Data
}

func (s *GetMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaCategoryResponseBody) SetData(v *GetMetaCategoryResponseBodyData) *GetMetaCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaCategoryResponseBody) SetErrorCode(v string) *GetMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaCategoryResponseBody) SetErrorMessage(v string) *GetMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaCategoryResponseBody) SetHttpStatusCode(v int32) *GetMetaCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaCategoryResponseBody) SetRequestId(v string) *GetMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaCategoryResponseBody) SetSuccess(v bool) *GetMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaCategoryResponseBodyData struct {
	// The information about the category tree.
	DataEntityList []*GetMetaCategoryResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of categories returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaCategoryResponseBodyData) GetDataEntityList() []*GetMetaCategoryResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *GetMetaCategoryResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaCategoryResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaCategoryResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaCategoryResponseBodyData) SetDataEntityList(v []*GetMetaCategoryResponseBodyDataDataEntityList) *GetMetaCategoryResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *GetMetaCategoryResponseBodyData) SetPageNum(v int32) *GetMetaCategoryResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetMetaCategoryResponseBodyData) SetPageSize(v int32) *GetMetaCategoryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaCategoryResponseBodyData) SetTotalCount(v int64) *GetMetaCategoryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaCategoryResponseBodyDataDataEntityList struct {
	// The category ID.
	//
	// example:
	//
	// 133
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The remarks of the category.
	//
	// example:
	//
	// category 1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the category was created.
	//
	// example:
	//
	// 1541576644000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of category levels.
	//
	// example:
	//
	// 1
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The ID of the user that performed the last operation.
	//
	// example:
	//
	// 12345
	LastOperatorId *string `json:"LastOperatorId,omitempty" xml:"LastOperatorId,omitempty"`
	// The time when the category was last modified.
	//
	// example:
	//
	// 1541576644000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// category 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The category owner ID.
	//
	// example:
	//
	// 123
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 12
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s GetMetaCategoryResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCategoryResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetComment() *string {
	return s.Comment
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetDepth() *int32 {
	return s.Depth
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetLastOperatorId() *string {
	return s.LastOperatorId
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetName() *string {
	return s.Name
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetCategoryId(v int64) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.CategoryId = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetComment(v string) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.Comment = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetCreateTime(v int64) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.CreateTime = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetDepth(v int32) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.Depth = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetLastOperatorId(v string) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.LastOperatorId = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetModifiedTime(v int64) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.ModifiedTime = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetName(v string) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.Name = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetOwnerId(v string) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.OwnerId = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) SetParentCategoryId(v int64) *GetMetaCategoryResponseBodyDataDataEntityList {
	s.ParentCategoryId = &v
	return s
}

func (s *GetMetaCategoryResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
