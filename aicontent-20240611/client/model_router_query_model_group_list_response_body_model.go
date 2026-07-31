// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryModelGroupListResponseBodyData) *ModelRouterQueryModelGroupListResponseBody
	GetData() *ModelRouterQueryModelGroupListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryModelGroupListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelGroupListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryModelGroupListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryModelGroupListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelGroupListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelGroupListResponseBody struct {
	// The returned data.
	Data *ModelRouterQueryModelGroupListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// An unused parameter.
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryModelGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetData() *ModelRouterQueryModelGroupListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelGroupListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetData(v *ModelRouterQueryModelGroupListResponseBodyData) *ModelRouterQueryModelGroupListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetErrCode(v string) *ModelRouterQueryModelGroupListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetErrMessage(v string) *ModelRouterQueryModelGroupListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetMaxResults(v int32) *ModelRouterQueryModelGroupListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetNextToken(v string) *ModelRouterQueryModelGroupListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetRequestId(v string) *ModelRouterQueryModelGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) SetSuccess(v bool) *ModelRouterQueryModelGroupListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryModelGroupListResponseBodyData struct {
	// The list of elements.
	List []*ModelGroupDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The requested page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryModelGroupListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) GetList() []*ModelGroupDTO {
	return s.List
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) SetList(v []*ModelGroupDTO) *ModelRouterQueryModelGroupListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) SetPage(v int32) *ModelRouterQueryModelGroupListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryModelGroupListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) SetTotal(v int32) *ModelRouterQueryModelGroupListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryModelGroupListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
