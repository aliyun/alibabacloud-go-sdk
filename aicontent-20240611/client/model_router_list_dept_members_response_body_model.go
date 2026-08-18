// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListDeptMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterListDeptMembersResponseBodyData) *ModelRouterListDeptMembersResponseBody
	GetData() *ModelRouterListDeptMembersResponseBodyData
	SetErrCode(v string) *ModelRouterListDeptMembersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterListDeptMembersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterListDeptMembersResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterListDeptMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterListDeptMembersResponseBody
	GetSuccess() *bool
}

type ModelRouterListDeptMembersResponseBody struct {
	// The response data object.
	//
	// example:
	//
	// {}
	Data *ModelRouterListDeptMembersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
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

func (s ModelRouterListDeptMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListDeptMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterListDeptMembersResponseBody) GetData() *ModelRouterListDeptMembersResponseBodyData {
	return s.Data
}

func (s *ModelRouterListDeptMembersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterListDeptMembersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterListDeptMembersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterListDeptMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterListDeptMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterListDeptMembersResponseBody) SetData(v *ModelRouterListDeptMembersResponseBodyData) *ModelRouterListDeptMembersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterListDeptMembersResponseBody) SetErrCode(v string) *ModelRouterListDeptMembersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBody) SetErrMessage(v string) *ModelRouterListDeptMembersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBody) SetHttpStatusCode(v int32) *ModelRouterListDeptMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBody) SetRequestId(v string) *ModelRouterListDeptMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBody) SetSuccess(v bool) *ModelRouterListDeptMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterListDeptMembersResponseBodyData struct {
	// The list of returned data.
	//
	// example:
	//
	// []
	List []*DeptMemberDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
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
	// 0
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterListDeptMembersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListDeptMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterListDeptMembersResponseBodyData) GetList() []*DeptMemberDTO {
	return s.List
}

func (s *ModelRouterListDeptMembersResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterListDeptMembersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterListDeptMembersResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterListDeptMembersResponseBodyData) SetList(v []*DeptMemberDTO) *ModelRouterListDeptMembersResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterListDeptMembersResponseBodyData) SetPage(v int32) *ModelRouterListDeptMembersResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBodyData) SetPageSize(v int32) *ModelRouterListDeptMembersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBodyData) SetTotal(v int32) *ModelRouterListDeptMembersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterListDeptMembersResponseBodyData) Validate() error {
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
