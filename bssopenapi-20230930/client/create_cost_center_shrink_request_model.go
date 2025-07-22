// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterEntityListShrink(v string) *CreateCostCenterShrinkRequest
	GetCostCenterEntityListShrink() *string
	SetNbid(v string) *CreateCostCenterShrinkRequest
	GetNbid() *string
}

type CreateCostCenterShrinkRequest struct {
	// This parameter is required.
	CostCenterEntityListShrink *string `json:"CostCenterEntityList,omitempty" xml:"CostCenterEntityList,omitempty"`
	// example:
	//
	// 2084210001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateCostCenterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCostCenterShrinkRequest) GetCostCenterEntityListShrink() *string {
	return s.CostCenterEntityListShrink
}

func (s *CreateCostCenterShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateCostCenterShrinkRequest) SetCostCenterEntityListShrink(v string) *CreateCostCenterShrinkRequest {
	s.CostCenterEntityListShrink = &v
	return s
}

func (s *CreateCostCenterShrinkRequest) SetNbid(v string) *CreateCostCenterShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *CreateCostCenterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
