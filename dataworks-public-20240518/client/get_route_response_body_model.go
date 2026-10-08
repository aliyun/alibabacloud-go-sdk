// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRouteResponseBody
	GetRequestId() *string
	SetRoute(v *GetRouteResponseBodyRoute) *GetRouteResponseBody
	GetRoute() *GetRouteResponseBodyRoute
	SetSuccess(v bool) *GetRouteResponseBody
	GetSuccess() *bool
}

type GetRouteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the route.
	Route *GetRouteResponseBodyRoute `json:"Route,omitempty" xml:"Route,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRouteResponseBody) GetRoute() *GetRouteResponseBodyRoute {
	return s.Route
}

func (s *GetRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRouteResponseBody) SetRequestId(v string) *GetRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRouteResponseBody) SetRoute(v *GetRouteResponseBodyRoute) *GetRouteResponseBody {
	s.Route = v
	return s
}

func (s *GetRouteResponseBody) SetSuccess(v bool) *GetRouteResponseBody {
	s.Success = &v
	return s
}

func (s *GetRouteResponseBody) Validate() error {
	if s.Route != nil {
		if err := s.Route.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRouteResponseBodyRoute struct {
	// The time when the route was created. The value is a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The CIDR block of the destination-based route.
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The route ID.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The network ID.
	//
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The network resource ID.
	//
	// example:
	//
	// ns-679XXXXX
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetRouteResponseBodyRoute) String() string {
	return dara.Prettify(s)
}

func (s GetRouteResponseBodyRoute) GoString() string {
	return s.String()
}

func (s *GetRouteResponseBodyRoute) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRouteResponseBodyRoute) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *GetRouteResponseBodyRoute) GetId() *int64 {
	return s.Id
}

func (s *GetRouteResponseBodyRoute) GetNetworkId() *int64 {
	return s.NetworkId
}

func (s *GetRouteResponseBodyRoute) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRouteResponseBodyRoute) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetRouteResponseBodyRoute) SetCreateTime(v int64) *GetRouteResponseBodyRoute {
	s.CreateTime = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetDestinationCidr(v string) *GetRouteResponseBodyRoute {
	s.DestinationCidr = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetId(v int64) *GetRouteResponseBodyRoute {
	s.Id = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetNetworkId(v int64) *GetRouteResponseBodyRoute {
	s.NetworkId = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetResourceGroupId(v string) *GetRouteResponseBodyRoute {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetResourceId(v string) *GetRouteResponseBodyRoute {
	s.ResourceId = &v
	return s
}

func (s *GetRouteResponseBodyRoute) Validate() error {
	return dara.Validate(s)
}
