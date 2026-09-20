// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AllocatePublicNetworkAddressRequest
	GetClientToken() *string
	SetClusterId(v string) *AllocatePublicNetworkAddressRequest
	GetClusterId() *string
}

type AllocatePublicNetworkAddressRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the value of RequestId as the value of ClientToken. The value of RequestId may be different for each API request.
	//
	// example:
	//
	// 83b2b5e117a5b8bce0fae88d90576a84_6452320_82718582
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s AllocatePublicNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocatePublicNetworkAddressRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AllocatePublicNetworkAddressRequest) SetClientToken(v string) *AllocatePublicNetworkAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetClusterId(v string) *AllocatePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
