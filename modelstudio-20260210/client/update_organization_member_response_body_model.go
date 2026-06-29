// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateOrganizationMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateOrganizationMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateOrganizationMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOrganizationMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOrganizationMemberResponseBody
	GetSuccess() *bool
}

type UpdateOrganizationMemberResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// None
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9531C132-DF05-5C7F-8BB0-96EA8C4D00D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOrganizationMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateOrganizationMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateOrganizationMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOrganizationMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrganizationMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOrganizationMemberResponseBody) SetCode(v string) *UpdateOrganizationMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetHttpStatusCode(v int32) *UpdateOrganizationMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetMessage(v string) *UpdateOrganizationMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetRequestId(v string) *UpdateOrganizationMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) SetSuccess(v bool) *UpdateOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOrganizationMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
