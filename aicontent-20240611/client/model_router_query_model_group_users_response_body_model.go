// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryModelGroupUsersResponseBodyData) *ModelRouterQueryModelGroupUsersResponseBody
	GetData() *ModelRouterQueryModelGroupUsersResponseBodyData
	SetErrCode(v string) *ModelRouterQueryModelGroupUsersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelGroupUsersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupUsersResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryModelGroupUsersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryModelGroupUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelGroupUsersResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelGroupUsersResponseBody struct {
	// The response data struct.
	Data *ModelRouterQueryModelGroupUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault code.
	//
	// example:
	//
	// B_USER_NOT_FOUND_EXCEPTION
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
	// nextToken
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

func (s ModelRouterQueryModelGroupUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetData() *ModelRouterQueryModelGroupUsersResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetData(v *ModelRouterQueryModelGroupUsersResponseBodyData) *ModelRouterQueryModelGroupUsersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetErrCode(v string) *ModelRouterQueryModelGroupUsersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetErrMessage(v string) *ModelRouterQueryModelGroupUsersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelGroupUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetMaxResults(v int32) *ModelRouterQueryModelGroupUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetNextToken(v string) *ModelRouterQueryModelGroupUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetRequestId(v string) *ModelRouterQueryModelGroupUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) SetSuccess(v bool) *ModelRouterQueryModelGroupUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryModelGroupUsersResponseBodyData struct {
	// The bound users.
	List []*ModelGroupUserDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The requested page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// None
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryModelGroupUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) GetList() []*ModelGroupUserDTO {
	return s.List
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) SetList(v []*ModelGroupUserDTO) *ModelRouterQueryModelGroupUsersResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) SetPage(v int32) *ModelRouterQueryModelGroupUsersResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) SetPageSize(v int32) *ModelRouterQueryModelGroupUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) SetTotal(v int32) *ModelRouterQueryModelGroupUsersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersResponseBodyData) Validate() error {
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
