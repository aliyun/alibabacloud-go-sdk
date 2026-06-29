// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenPlanInviteLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTokenPlanInviteLinkResponseBody
	GetCode() *string
	SetData(v *CreateTokenPlanInviteLinkResponseBodyData) *CreateTokenPlanInviteLinkResponseBody
	GetData() *CreateTokenPlanInviteLinkResponseBodyData
	SetMessage(v string) *CreateTokenPlanInviteLinkResponseBody
	GetMessage() *string
	SetSuccess(v bool) *CreateTokenPlanInviteLinkResponseBody
	GetSuccess() *bool
}

type CreateTokenPlanInviteLinkResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business parameters.
	Data *CreateTokenPlanInviteLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTokenPlanInviteLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenPlanInviteLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenPlanInviteLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTokenPlanInviteLinkResponseBody) GetData() *CreateTokenPlanInviteLinkResponseBodyData {
	return s.Data
}

func (s *CreateTokenPlanInviteLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTokenPlanInviteLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTokenPlanInviteLinkResponseBody) SetCode(v string) *CreateTokenPlanInviteLinkResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTokenPlanInviteLinkResponseBody) SetData(v *CreateTokenPlanInviteLinkResponseBodyData) *CreateTokenPlanInviteLinkResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenPlanInviteLinkResponseBody) SetMessage(v string) *CreateTokenPlanInviteLinkResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTokenPlanInviteLinkResponseBody) SetSuccess(v bool) *CreateTokenPlanInviteLinkResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTokenPlanInviteLinkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTokenPlanInviteLinkResponseBodyData struct {
	// The generated token.
	//
	// example:
	//
	// sk-ws-D.****.*******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenPlanInviteLinkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenPlanInviteLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenPlanInviteLinkResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CreateTokenPlanInviteLinkResponseBodyData) SetToken(v string) *CreateTokenPlanInviteLinkResponseBodyData {
	s.Token = &v
	return s
}

func (s *CreateTokenPlanInviteLinkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
