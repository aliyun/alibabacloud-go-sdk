// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v []*string) *DeleteMetricRulesRequest
	GetId() []*string
	SetRegionId(v string) *DeleteMetricRulesRequest
	GetRegionId() *string
}

type DeleteMetricRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ab05733c97b7ce239fb1b53393dc1697c7e12****
	Id       []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesRequest) GetId() []*string {
	return s.Id
}

func (s *DeleteMetricRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMetricRulesRequest) SetId(v []*string) *DeleteMetricRulesRequest {
	s.Id = v
	return s
}

func (s *DeleteMetricRulesRequest) SetRegionId(v string) *DeleteMetricRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMetricRulesRequest) Validate() error {
	return dara.Validate(s)
}
