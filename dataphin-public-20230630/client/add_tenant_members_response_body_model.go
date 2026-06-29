// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTenantMembersResponseBody
	GetCode() *string
	SetData(v bool) *AddTenantMembersResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddTenantMembersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddTenantMembersResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddTenantMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTenantMembersResponseBody
	GetSuccess() *bool
}

type AddTenantMembersResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of adding members to the tenant.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTenantMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddTenantMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTenantMembersResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddTenantMembersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddTenantMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTenantMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTenantMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTenantMembersResponseBody) SetCode(v string) *AddTenantMembersResponseBody {
	s.Code = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetData(v bool) *AddTenantMembersResponseBody {
	s.Data = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetHttpStatusCode(v int32) *AddTenantMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetMessage(v string) *AddTenantMembersResponseBody {
	s.Message = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetRequestId(v string) *AddTenantMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTenantMembersResponseBody) SetSuccess(v bool) *AddTenantMembersResponseBody {
	s.Success = &v
	return s
}

func (s *AddTenantMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
