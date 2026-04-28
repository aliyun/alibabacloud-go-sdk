// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBudgetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetPolicyId(v string) *DeleteBudgetPolicyRequest
	GetBudgetPolicyId() *string
	SetGwClusterId(v string) *DeleteBudgetPolicyRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DeleteBudgetPolicyRequest
	GetRegionId() *string
}

type DeleteBudgetPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 023aacc1effc4b56bb154bfbec6baxxx
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBudgetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBudgetPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBudgetPolicyRequest) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *DeleteBudgetPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteBudgetPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBudgetPolicyRequest) SetBudgetPolicyId(v string) *DeleteBudgetPolicyRequest {
	s.BudgetPolicyId = &v
	return s
}

func (s *DeleteBudgetPolicyRequest) SetGwClusterId(v string) *DeleteBudgetPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteBudgetPolicyRequest) SetRegionId(v string) *DeleteBudgetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBudgetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
