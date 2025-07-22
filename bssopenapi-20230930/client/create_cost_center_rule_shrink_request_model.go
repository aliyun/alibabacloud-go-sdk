// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *CreateCostCenterRuleShrinkRequest
	GetCostCenterId() *int64
	SetFilterExpressionShrink(v string) *CreateCostCenterRuleShrinkRequest
	GetFilterExpressionShrink() *string
	SetNbid(v string) *CreateCostCenterRuleShrinkRequest
	GetNbid() *string
}

type CreateCostCenterRuleShrinkRequest struct {
	// example:
	//
	// 485938
	CostCenterId           *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	FilterExpressionShrink *string `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty"`
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateCostCenterRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRuleShrinkRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *CreateCostCenterRuleShrinkRequest) GetFilterExpressionShrink() *string {
	return s.FilterExpressionShrink
}

func (s *CreateCostCenterRuleShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateCostCenterRuleShrinkRequest) SetCostCenterId(v int64) *CreateCostCenterRuleShrinkRequest {
	s.CostCenterId = &v
	return s
}

func (s *CreateCostCenterRuleShrinkRequest) SetFilterExpressionShrink(v string) *CreateCostCenterRuleShrinkRequest {
	s.FilterExpressionShrink = &v
	return s
}

func (s *CreateCostCenterRuleShrinkRequest) SetNbid(v string) *CreateCostCenterRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *CreateCostCenterRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
