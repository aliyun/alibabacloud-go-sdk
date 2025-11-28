// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteRouteRuleRequest
	GetRegionId() *string
	SetRouteRuleId(v string) *DeleteRouteRuleRequest
	GetRouteRuleId() *string
}

type DeleteRouteRuleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RouteRuleId *string `json:"RouteRuleId,omitempty" xml:"RouteRuleId,omitempty"`
}

func (s DeleteRouteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteRuleRequest) GetRouteRuleId() *string {
	return s.RouteRuleId
}

func (s *DeleteRouteRuleRequest) SetRegionId(v string) *DeleteRouteRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteRuleRequest) SetRouteRuleId(v string) *DeleteRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

func (s *DeleteRouteRuleRequest) Validate() error {
	return dara.Validate(s)
}
