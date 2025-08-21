// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateLogstashChargeTypeRequest
	GetBody() *string
	SetClientToken(v string) *UpdateLogstashChargeTypeRequest
	GetClientToken() *string
}

type UpdateLogstashChargeTypeRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateLogstashChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogstashChargeTypeRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateLogstashChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLogstashChargeTypeRequest) SetBody(v string) *UpdateLogstashChargeTypeRequest {
	s.Body = &v
	return s
}

func (s *UpdateLogstashChargeTypeRequest) SetClientToken(v string) *UpdateLogstashChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLogstashChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
