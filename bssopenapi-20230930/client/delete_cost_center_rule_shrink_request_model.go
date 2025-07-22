// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *DeleteCostCenterRuleShrinkRequest
	GetCostCenterId() *int64
	SetFilterExpressionShrink(v string) *DeleteCostCenterRuleShrinkRequest
	GetFilterExpressionShrink() *string
	SetNbid(v string) *DeleteCostCenterRuleShrinkRequest
	GetNbid() *string
}

type DeleteCostCenterRuleShrinkRequest struct {
	// example:
	//
	// 637127
	CostCenterId           *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	FilterExpressionShrink *string `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s DeleteCostCenterRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRuleShrinkRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *DeleteCostCenterRuleShrinkRequest) GetFilterExpressionShrink() *string {
	return s.FilterExpressionShrink
}

func (s *DeleteCostCenterRuleShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteCostCenterRuleShrinkRequest) SetCostCenterId(v int64) *DeleteCostCenterRuleShrinkRequest {
	s.CostCenterId = &v
	return s
}

func (s *DeleteCostCenterRuleShrinkRequest) SetFilterExpressionShrink(v string) *DeleteCostCenterRuleShrinkRequest {
	s.FilterExpressionShrink = &v
	return s
}

func (s *DeleteCostCenterRuleShrinkRequest) SetNbid(v string) *DeleteCostCenterRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCostCenterRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
