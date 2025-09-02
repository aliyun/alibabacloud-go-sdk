// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableListByCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableListByCategoryResponseBodyData) *GetMetaTableListByCategoryResponseBody
	GetData() *GetMetaTableListByCategoryResponseBodyData
	SetErrorCode(v string) *GetMetaTableListByCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableListByCategoryResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableListByCategoryResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableListByCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableListByCategoryResponseBody
	GetSuccess() *bool
}

type GetMetaTableListByCategoryResponseBody struct {
	// The business data returned.
	Data *GetMetaTableListByCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The ID of the request.
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

func (s GetMetaTableListByCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableListByCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableListByCategoryResponseBody) GetData() *GetMetaTableListByCategoryResponseBodyData {
	return s.Data
}

func (s *GetMetaTableListByCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableListByCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableListByCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableListByCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableListByCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableListByCategoryResponseBody) SetData(v *GetMetaTableListByCategoryResponseBodyData) *GetMetaTableListByCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableListByCategoryResponseBody) SetErrorCode(v string) *GetMetaTableListByCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBody) SetErrorMessage(v string) *GetMetaTableListByCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBody) SetHttpStatusCode(v int32) *GetMetaTableListByCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBody) SetRequestId(v string) *GetMetaTableListByCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBody) SetSuccess(v bool) *GetMetaTableListByCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableListByCategoryResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned result.
	TableGuidList []*string `json:"TableGuidList,omitempty" xml:"TableGuidList,omitempty" type:"Repeated"`
	// The total number of metatables.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMetaTableListByCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableListByCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableListByCategoryResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTableListByCategoryResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableListByCategoryResponseBodyData) GetTableGuidList() []*string {
	return s.TableGuidList
}

func (s *GetMetaTableListByCategoryResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMetaTableListByCategoryResponseBodyData) SetPageNumber(v int32) *GetMetaTableListByCategoryResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBodyData) SetPageSize(v int32) *GetMetaTableListByCategoryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBodyData) SetTableGuidList(v []*string) *GetMetaTableListByCategoryResponseBodyData {
	s.TableGuidList = v
	return s
}

func (s *GetMetaTableListByCategoryResponseBodyData) SetTotalCount(v int64) *GetMetaTableListByCategoryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetMetaTableListByCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
