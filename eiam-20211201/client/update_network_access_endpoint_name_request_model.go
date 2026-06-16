// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessEndpointNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateNetworkAccessEndpointNameRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *UpdateNetworkAccessEndpointNameRequest
	GetNetworkAccessEndpointId() *string
	SetNetworkAccessEndpointName(v string) *UpdateNetworkAccessEndpointNameRequest
	GetNetworkAccessEndpointName() *string
}

type UpdateNetworkAccessEndpointNameRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the private network access endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The name of the private network access endpoint. The endpoint type must be private.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC access endpoint for xx service
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNetworkAccessEndpointNameRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *UpdateNetworkAccessEndpointNameRequest) GetNetworkAccessEndpointName() *string {
	return s.NetworkAccessEndpointName
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetInstanceId(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetNetworkAccessEndpointId(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetNetworkAccessEndpointName(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) Validate() error {
	return dara.Validate(s)
}
