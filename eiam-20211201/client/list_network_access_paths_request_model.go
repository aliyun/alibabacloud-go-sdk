// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessPathsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNetworkAccessPathsRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *ListNetworkAccessPathsRequest
	GetNetworkAccessEndpointId() *string
}

type ListNetworkAccessPathsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Network access endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
}

func (s ListNetworkAccessPathsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessPathsRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkAccessPathsRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListNetworkAccessPathsRequest) SetInstanceId(v string) *ListNetworkAccessPathsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessPathsRequest) SetNetworkAccessEndpointId(v string) *ListNetworkAccessPathsRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListNetworkAccessPathsRequest) Validate() error {
	return dara.Validate(s)
}
