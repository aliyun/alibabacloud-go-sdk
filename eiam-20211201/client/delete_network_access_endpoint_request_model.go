// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAccessEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNetworkAccessEndpointRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *DeleteNetworkAccessEndpointRequest
	GetNetworkAccessEndpointId() *string
}

type DeleteNetworkAccessEndpointRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Network Access Endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
}

func (s DeleteNetworkAccessEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAccessEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNetworkAccessEndpointRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *DeleteNetworkAccessEndpointRequest) SetInstanceId(v string) *DeleteNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNetworkAccessEndpointRequest) SetNetworkAccessEndpointId(v string) *DeleteNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *DeleteNetworkAccessEndpointRequest) Validate() error {
	return dara.Validate(s)
}
