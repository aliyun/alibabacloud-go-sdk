// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRouteItem interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointType(v string) *RouteItem
	GetEndpointType() *string
	SetInstanceDomain(v string) *RouteItem
	GetInstanceDomain() *string
	SetStorageDomain(v string) *RouteItem
	GetStorageDomain() *string
}

type RouteItem struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// This parameter is required.
	InstanceDomain *string `json:"InstanceDomain,omitempty" xml:"InstanceDomain,omitempty"`
	// This parameter is required.
	StorageDomain *string `json:"StorageDomain,omitempty" xml:"StorageDomain,omitempty"`
}

func (s RouteItem) String() string {
	return dara.Prettify(s)
}

func (s RouteItem) GoString() string {
	return s.String()
}

func (s *RouteItem) GetEndpointType() *string {
	return s.EndpointType
}

func (s *RouteItem) GetInstanceDomain() *string {
	return s.InstanceDomain
}

func (s *RouteItem) GetStorageDomain() *string {
	return s.StorageDomain
}

func (s *RouteItem) SetEndpointType(v string) *RouteItem {
	s.EndpointType = &v
	return s
}

func (s *RouteItem) SetInstanceDomain(v string) *RouteItem {
	s.InstanceDomain = &v
	return s
}

func (s *RouteItem) SetStorageDomain(v string) *RouteItem {
	s.StorageDomain = &v
	return s
}

func (s *RouteItem) Validate() error {
	return dara.Validate(s)
}
