// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsOptimizationStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListApsOptimizationStrategyRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ListApsOptimizationStrategyRequest
	GetRegionId() *string
}

type ListApsOptimizationStrategyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApsOptimizationStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApsOptimizationStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListApsOptimizationStrategyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsOptimizationStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApsOptimizationStrategyRequest) SetDBClusterId(v string) *ListApsOptimizationStrategyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListApsOptimizationStrategyRequest) SetRegionId(v string) *ListApsOptimizationStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ListApsOptimizationStrategyRequest) Validate() error {
	return dara.Validate(s)
}
