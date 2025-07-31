// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTenantMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateTenantMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTenantMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTenantMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTenantMemberResponseBody
	GetSuccess() *bool
}

type UpdateTenantMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTenantMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTenantMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTenantMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTenantMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTenantMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTenantMemberResponseBody) SetCode(v string) *UpdateTenantMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetHttpStatusCode(v int32) *UpdateTenantMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetMessage(v string) *UpdateTenantMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetRequestId(v string) *UpdateTenantMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) SetSuccess(v bool) *UpdateTenantMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTenantMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
