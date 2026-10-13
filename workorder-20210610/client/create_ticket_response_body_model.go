// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateTicketResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateTicketResponseBody
	GetCode() *int32
	SetData(v string) *CreateTicketResponseBody
	GetData() *string
	SetMessage(v string) *CreateTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTicketResponseBody
	GetSuccess() *bool
}

type CreateTicketResponseBody struct {
	// The details of the permission error. An empty string is returned if the permissions are sufficient.
	//
	// example:
	//
	// {\\"AuthAction\\":\\"ram:CreateTicket\\",\\"AuthPrincipalDisplayName\\":\\"2146216584788xxxxx\\",\\"AuthPrincipalOwnerId\\":\\"1135850448xxxxx\\",\\"AuthPrincipalType\\":\\"SubUser\\",\\"NoPermissionType\\":\\"ImplicitDeny\\",\\"PolicyType\\":\\"AccountLevelIdentityBasedPolicy\\"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the API request.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value, which is the ticket ID.
	//
	// example:
	//
	// 0005PYGCW
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message. This parameter is returned when the value of Success is false.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Each request has a unique ID.
	//
	// example:
	//
	// 0254B7DE-7365-57B4-8E38-14FE927E3FEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful. A value of true indicates that the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateTicketResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateTicketResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTicketResponseBody) SetAccessDeniedDetail(v string) *CreateTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateTicketResponseBody) SetCode(v int32) *CreateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketResponseBody) SetData(v string) *CreateTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketResponseBody) SetMessage(v string) *CreateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
