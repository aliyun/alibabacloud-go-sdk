// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *AssociateRouteTableRequest
	GetRegionId() *string
	SetRouteTableId(v string) *AssociateRouteTableRequest
	GetRouteTableId() *string
	SetVSwitchId(v string) *AssociateRouteTableRequest
	GetVSwitchId() *string
}

type AssociateRouteTableRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AssociateRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableRequest) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateRouteTableRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *AssociateRouteTableRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AssociateRouteTableRequest) SetRegionId(v string) *AssociateRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateRouteTableRequest) SetRouteTableId(v string) *AssociateRouteTableRequest {
	s.RouteTableId = &v
	return s
}

func (s *AssociateRouteTableRequest) SetVSwitchId(v string) *AssociateRouteTableRequest {
	s.VSwitchId = &v
	return s
}

func (s *AssociateRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
