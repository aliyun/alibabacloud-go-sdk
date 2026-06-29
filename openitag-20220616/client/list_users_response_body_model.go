// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListUsersResponseBody
	GetCode() *int32
	SetDetails(v string) *ListUsersResponseBody
	GetDetails() *string
	SetErrorCode(v string) *ListUsersResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListUsersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUsersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListUsersResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListUsersResponseBody
	GetTotalPage() *int32
	SetUsers(v []*SimpleUser) *ListUsersResponseBody
	GetUsers() []*SimpleUser
}

type ListUsersResponseBody struct {
	// Return encoding. The default value is 0, indicating Normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// - When Success is false, returns a business error code.
	//
	// - When Success is true, returns an empty value.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number of the queried annotate member list returned in the response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of annotate members displayed per page in the response.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of annotate members.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// User List.
	Users []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListUsersResponseBody) GetDetails() *string {
	return s.Details
}

func (s *ListUsersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUsersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListUsersResponseBody) GetUsers() []*SimpleUser {
	return s.Users
}

func (s *ListUsersResponseBody) SetCode(v int32) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetDetails(v string) *ListUsersResponseBody {
	s.Details = &v
	return s
}

func (s *ListUsersResponseBody) SetErrorCode(v string) *ListUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalPage(v int32) *ListUsersResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*SimpleUser) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
