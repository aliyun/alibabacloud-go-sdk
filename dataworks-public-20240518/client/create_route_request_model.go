// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidr(v string) *CreateRouteRequest
	GetDestinationCidr() *string
	SetNetworkId(v int64) *CreateRouteRequest
	GetNetworkId() *int64
	SetResourceGroupId(v string) *CreateRouteRequest
	GetResourceGroupId() *string
}

type CreateRouteRequest struct {
	// The CIDR blocks of the destination-based route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// Unique identifier of the serverless resource group.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteRequest) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *CreateRouteRequest) GetNetworkId() *int64 {
	return s.NetworkId
}

func (s *CreateRouteRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRouteRequest) SetDestinationCidr(v string) *CreateRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *CreateRouteRequest) SetNetworkId(v int64) *CreateRouteRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateRouteRequest) SetResourceGroupId(v string) *CreateRouteRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRouteRequest) Validate() error {
	return dara.Validate(s)
}
