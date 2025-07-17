// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidr(v string) *UpdateRouteRequest
	GetDestinationCidr() *string
	SetId(v int64) *UpdateRouteRequest
	GetId() *int64
}

type UpdateRouteRequest struct {
	// The destination CIDR block of the route that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The route ID of the network resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteRequest) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *UpdateRouteRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateRouteRequest) SetDestinationCidr(v string) *UpdateRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *UpdateRouteRequest) SetId(v int64) *UpdateRouteRequest {
	s.Id = &v
	return s
}

func (s *UpdateRouteRequest) Validate() error {
	return dara.Validate(s)
}
