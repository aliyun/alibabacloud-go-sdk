// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterEntityListShrink(v string) *ModifyCostCenterShrinkRequest
	GetCostCenterEntityListShrink() *string
	SetNbid(v string) *ModifyCostCenterShrinkRequest
	GetNbid() *string
}

type ModifyCostCenterShrinkRequest struct {
	// This parameter is required.
	CostCenterEntityListShrink *string `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ModifyCostCenterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterShrinkRequest) GetCostCenterEntityListShrink() *string {
	return s.CostCenterEntityListShrink
}

func (s *ModifyCostCenterShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ModifyCostCenterShrinkRequest) SetCostCenterEntityListShrink(v string) *ModifyCostCenterShrinkRequest {
	s.CostCenterEntityListShrink = &v
	return s
}

func (s *ModifyCostCenterShrinkRequest) SetNbid(v string) *ModifyCostCenterShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *ModifyCostCenterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
