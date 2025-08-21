// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClassicNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkId(v string) *CreateClassicNetworkResponseBody
	GetNetworkId() *string
	SetRequestId(v string) *CreateClassicNetworkResponseBody
	GetRequestId() *string
}

type CreateClassicNetworkResponseBody struct {
	// The ID of the network.
	//
	// example:
	//
	// n-5s9ayrxsd9hszrlt5fgv2****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A004E06-AC1B-5806-BA5E-41AB6B02DE83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClassicNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClassicNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClassicNetworkResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateClassicNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClassicNetworkResponseBody) SetNetworkId(v string) *CreateClassicNetworkResponseBody {
	s.NetworkId = &v
	return s
}

func (s *CreateClassicNetworkResponseBody) SetRequestId(v string) *CreateClassicNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClassicNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
