// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTokenPlanOrgInviteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetTokenPlanOrgInviteConfigResponseBody
	GetCode() *string
	SetMessage(v string) *SetTokenPlanOrgInviteConfigResponseBody
	GetMessage() *string
	SetSuccess(v bool) *SetTokenPlanOrgInviteConfigResponseBody
	GetSuccess() *bool
}

type SetTokenPlanOrgInviteConfigResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetTokenPlanOrgInviteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetTokenPlanOrgInviteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) SetCode(v string) *SetTokenPlanOrgInviteConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) SetMessage(v string) *SetTokenPlanOrgInviteConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) SetSuccess(v bool) *SetTokenPlanOrgInviteConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetTokenPlanOrgInviteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
