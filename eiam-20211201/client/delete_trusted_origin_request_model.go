// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrustedOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTrustedOriginRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteTrustedOriginRequest
	GetInstanceId() *string
	SetTrustedOriginId(v string) *DeleteTrustedOriginRequest
	GetTrustedOriginId() *string
}

type DeleteTrustedOriginRequest struct {
	// A client token that is used to ensure the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique among different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the trusted origin.
	//
	// This parameter is required.
	//
	// example:
	//
	// to_example
	TrustedOriginId *string `json:"TrustedOriginId,omitempty" xml:"TrustedOriginId,omitempty"`
}

func (s DeleteTrustedOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrustedOriginRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrustedOriginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTrustedOriginRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTrustedOriginRequest) GetTrustedOriginId() *string {
	return s.TrustedOriginId
}

func (s *DeleteTrustedOriginRequest) SetClientToken(v string) *DeleteTrustedOriginRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTrustedOriginRequest) SetInstanceId(v string) *DeleteTrustedOriginRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTrustedOriginRequest) SetTrustedOriginId(v string) *DeleteTrustedOriginRequest {
	s.TrustedOriginId = &v
	return s
}

func (s *DeleteTrustedOriginRequest) Validate() error {
	return dara.Validate(s)
}
