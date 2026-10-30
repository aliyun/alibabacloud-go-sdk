// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterCapacityReservationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTransitRouterId(v string) *DescribeTransitRouterCapacityReservationsRequest
	GetTransitRouterId() *string
	SetVersion(v string) *DescribeTransitRouterCapacityReservationsRequest
	GetVersion() *string
}

type DescribeTransitRouterCapacityReservationsRequest struct {
	// The ID of the Enterprise Edition transit router instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-uf6llz2286805i44g****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeTransitRouterCapacityReservationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterCapacityReservationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterCapacityReservationsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTransitRouterCapacityReservationsRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeTransitRouterCapacityReservationsRequest) SetTransitRouterId(v string) *DescribeTransitRouterCapacityReservationsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsRequest) SetVersion(v string) *DescribeTransitRouterCapacityReservationsRequest {
	s.Version = &v
	return s
}

func (s *DescribeTransitRouterCapacityReservationsRequest) Validate() error {
	return dara.Validate(s)
}
