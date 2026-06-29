// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanInviteLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenPlanInviteLinkResponseBody
	GetCode() *string
	SetData(v *GetTokenPlanInviteLinkResponseBodyData) *GetTokenPlanInviteLinkResponseBody
	GetData() *GetTokenPlanInviteLinkResponseBodyData
	SetMessage(v string) *GetTokenPlanInviteLinkResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetTokenPlanInviteLinkResponseBody
	GetSuccess() *bool
}

type GetTokenPlanInviteLinkResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetTokenPlanInviteLinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the call was successful. Valid values:
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

func (s GetTokenPlanInviteLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanInviteLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenPlanInviteLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenPlanInviteLinkResponseBody) GetData() *GetTokenPlanInviteLinkResponseBodyData {
	return s.Data
}

func (s *GetTokenPlanInviteLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenPlanInviteLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTokenPlanInviteLinkResponseBody) SetCode(v string) *GetTokenPlanInviteLinkResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenPlanInviteLinkResponseBody) SetData(v *GetTokenPlanInviteLinkResponseBodyData) *GetTokenPlanInviteLinkResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenPlanInviteLinkResponseBody) SetMessage(v string) *GetTokenPlanInviteLinkResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenPlanInviteLinkResponseBody) SetSuccess(v bool) *GetTokenPlanInviteLinkResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenPlanInviteLinkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenPlanInviteLinkResponseBodyData struct {
	// The expiration time. Unit: milliseconds.
	//
	// example:
	//
	// 1778379206
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The generated token.
	//
	// example:
	//
	// sk-ws-D.****.*******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenPlanInviteLinkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanInviteLinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenPlanInviteLinkResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetTokenPlanInviteLinkResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetTokenPlanInviteLinkResponseBodyData) SetExpireTime(v int64) *GetTokenPlanInviteLinkResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenPlanInviteLinkResponseBodyData) SetToken(v string) *GetTokenPlanInviteLinkResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetTokenPlanInviteLinkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
