// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTokenPlanInviteLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeTokenPlanInviteLinkResponseBody
	GetCode() *string
	SetMessage(v string) *RevokeTokenPlanInviteLinkResponseBody
	GetMessage() *string
	SetSuccess(v bool) *RevokeTokenPlanInviteLinkResponseBody
	GetSuccess() *bool
}

type RevokeTokenPlanInviteLinkResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// False
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeTokenPlanInviteLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeTokenPlanInviteLinkResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTokenPlanInviteLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeTokenPlanInviteLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeTokenPlanInviteLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeTokenPlanInviteLinkResponseBody) SetCode(v string) *RevokeTokenPlanInviteLinkResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeTokenPlanInviteLinkResponseBody) SetMessage(v string) *RevokeTokenPlanInviteLinkResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeTokenPlanInviteLinkResponseBody) SetSuccess(v bool) *RevokeTokenPlanInviteLinkResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeTokenPlanInviteLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
