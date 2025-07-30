// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAccessEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAccessEndpointId(v string) *CreateNetworkAccessEndpointResponseBody
	GetNetworkAccessEndpointId() *string
	SetRequestId(v string) *CreateNetworkAccessEndpointResponseBody
	GetRequestId() *string
}

type CreateNetworkAccessEndpointResponseBody struct {
	// The unique identifier of the network access endpoint.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkAccessEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkAccessEndpointResponseBody) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *CreateNetworkAccessEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkAccessEndpointResponseBody) SetNetworkAccessEndpointId(v string) *CreateNetworkAccessEndpointResponseBody {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *CreateNetworkAccessEndpointResponseBody) SetRequestId(v string) *CreateNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkAccessEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
