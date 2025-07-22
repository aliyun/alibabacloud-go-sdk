// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *ModifyCostCenterRuleShrinkRequest
	GetCostCenterId() *int64
	SetFilterExpressionShrink(v string) *ModifyCostCenterRuleShrinkRequest
	GetFilterExpressionShrink() *string
	SetNbid(v string) *ModifyCostCenterRuleShrinkRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *ModifyCostCenterRuleShrinkRequest
	GetOwnerAccountId() *int64
}

type ModifyCostCenterRuleShrinkRequest struct {
	// example:
	//
	// 485938
	CostCenterId           *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	FilterExpressionShrink *string `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 1234567812345678
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s ModifyCostCenterRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRuleShrinkRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ModifyCostCenterRuleShrinkRequest) GetFilterExpressionShrink() *string {
	return s.FilterExpressionShrink
}

func (s *ModifyCostCenterRuleShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ModifyCostCenterRuleShrinkRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *ModifyCostCenterRuleShrinkRequest) SetCostCenterId(v int64) *ModifyCostCenterRuleShrinkRequest {
	s.CostCenterId = &v
	return s
}

func (s *ModifyCostCenterRuleShrinkRequest) SetFilterExpressionShrink(v string) *ModifyCostCenterRuleShrinkRequest {
	s.FilterExpressionShrink = &v
	return s
}

func (s *ModifyCostCenterRuleShrinkRequest) SetNbid(v string) *ModifyCostCenterRuleShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *ModifyCostCenterRuleShrinkRequest) SetOwnerAccountId(v int64) *ModifyCostCenterRuleShrinkRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *ModifyCostCenterRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
