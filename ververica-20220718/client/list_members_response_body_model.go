// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Member) *ListMembersResponseBody
	GetData() []*Member
	SetErrorCode(v string) *ListMembersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMembersResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListMembersResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListMembersResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMembersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMembersResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListMembersResponseBody
	GetTotalSize() *int32
}

type ListMembersResponseBody struct {
	// 	- If the value of success was false, a null value was returned.
	//
	// 	- If the value of success was true, the authorization information was returned.
	Data []*Member `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) GetData() []*Member {
	return s.Data
}

func (s *ListMembersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMembersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMembersResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListMembersResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMembersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMembersResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListMembersResponseBody) SetData(v []*Member) *ListMembersResponseBody {
	s.Data = v
	return s
}

func (s *ListMembersResponseBody) SetErrorCode(v string) *ListMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMembersResponseBody) SetErrorMessage(v string) *ListMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMembersResponseBody) SetHttpCode(v int32) *ListMembersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListMembersResponseBody) SetPageIndex(v int32) *ListMembersResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMembersResponseBody) SetPageSize(v int32) *ListMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetSuccess(v bool) *ListMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListMembersResponseBody) SetTotalSize(v int32) *ListMembersResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListMembersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
