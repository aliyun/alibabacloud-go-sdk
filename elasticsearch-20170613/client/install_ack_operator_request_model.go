// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAckOperatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *InstallAckOperatorRequest
	GetClientToken() *string
	SetBody(v string) *InstallAckOperatorRequest
	GetBody() *string
}

type InstallAckOperatorRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAckOperatorRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAckOperatorRequest) GoString() string {
	return s.String()
}

func (s *InstallAckOperatorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallAckOperatorRequest) GetBody() *string {
	return s.Body
}

func (s *InstallAckOperatorRequest) SetClientToken(v string) *InstallAckOperatorRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallAckOperatorRequest) SetBody(v string) *InstallAckOperatorRequest {
	s.Body = &v
	return s
}

func (s *InstallAckOperatorRequest) Validate() error {
	return dara.Validate(s)
}
