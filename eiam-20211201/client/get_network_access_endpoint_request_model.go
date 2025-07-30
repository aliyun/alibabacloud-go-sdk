// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkAccessEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNetworkAccessEndpointRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *GetNetworkAccessEndpointRequest
	GetNetworkAccessEndpointId() *string
}

type GetNetworkAccessEndpointRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unique identifier of the network access endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// nae-examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
}

func (s GetNetworkAccessEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNetworkAccessEndpointRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *GetNetworkAccessEndpointRequest) SetInstanceId(v string) *GetNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkAccessEndpointRequest) SetNetworkAccessEndpointId(v string) *GetNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *GetNetworkAccessEndpointRequest) Validate() error {
	return dara.Validate(s)
}
