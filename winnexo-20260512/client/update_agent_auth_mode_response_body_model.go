// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentAuthModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentAuthModeResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAgentAuthModeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentAuthModeResponseBody
	GetRequestId() *string
	SetUpdated(v bool) *UpdateAgentAuthModeResponseBody
	GetUpdated() *bool
}

type UpdateAgentAuthModeResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The updated item.
	//
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateAgentAuthModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentAuthModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentAuthModeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentAuthModeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentAuthModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentAuthModeResponseBody) GetUpdated() *bool {
	return s.Updated
}

func (s *UpdateAgentAuthModeResponseBody) SetCode(v string) *UpdateAgentAuthModeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentAuthModeResponseBody) SetMessage(v string) *UpdateAgentAuthModeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentAuthModeResponseBody) SetRequestId(v string) *UpdateAgentAuthModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentAuthModeResponseBody) SetUpdated(v bool) *UpdateAgentAuthModeResponseBody {
	s.Updated = &v
	return s
}

func (s *UpdateAgentAuthModeResponseBody) Validate() error {
	return dara.Validate(s)
}
