// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkDomainId(v string) *CreateNetworkDomainResponseBody
	GetNetworkDomainId() *string
	SetRequestId(v string) *CreateNetworkDomainResponseBody
	GetRequestId() *string
}

type CreateNetworkDomainResponseBody struct {
	// The ID of the network domain.
	//
	// example:
	//
	// 31
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A2873E9C-A7EA-5735-845C-65D3792623D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkDomainResponseBody) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *CreateNetworkDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkDomainResponseBody) SetNetworkDomainId(v string) *CreateNetworkDomainResponseBody {
	s.NetworkDomainId = &v
	return s
}

func (s *CreateNetworkDomainResponseBody) SetRequestId(v string) *CreateNetworkDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
