// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAnycastEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *ReleaseAnycastEipAddressRequest
	GetAnycastId() *string
	SetClientToken(v string) *ReleaseAnycastEipAddressRequest
	GetClientToken() *string
}

type ReleaseAnycastEipAddressRequest struct {
	// The ID of the Anycast EIP to be released.
	//
	// This parameter is required.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ReleaseAnycastEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseAnycastEipAddressRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *ReleaseAnycastEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseAnycastEipAddressRequest) SetAnycastId(v string) *ReleaseAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *ReleaseAnycastEipAddressRequest) SetClientToken(v string) *ReleaseAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseAnycastEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
