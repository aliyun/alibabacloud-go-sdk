// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteRouteTableRequest
	GetRegionId() *string
	SetRouteTableId(v string) *DeleteRouteTableRequest
	GetRouteTableId() *string
}

type DeleteRouteTableRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteTableRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteRouteTableRequest) SetRegionId(v string) *DeleteRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteTableRequest) SetRouteTableId(v string) *DeleteRouteTableRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
