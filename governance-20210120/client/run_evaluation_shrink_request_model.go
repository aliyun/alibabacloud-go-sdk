// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEvaluationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *RunEvaluationShrinkRequest
	GetAccountId() *int64
	SetMetricIdsShrink(v string) *RunEvaluationShrinkRequest
	GetMetricIdsShrink() *string
	SetRegionId(v string) *RunEvaluationShrinkRequest
	GetRegionId() *string
	SetScope(v string) *RunEvaluationShrinkRequest
	GetScope() *string
}

type RunEvaluationShrinkRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The IDs of the check items to be checked.
	MetricIdsShrink *string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The check range of the governance maturity check. Valid values:
	//
	// 	- Account (default): A single-account governance maturity check is performed to check only the Alibaba Cloud account that you use to access Cloud Governance Center.
	//
	// 	- ResourceDirectory: A multi-account governance maturity check is performed to check all members within a resource directory. Before you perform a multi-account governance maturity check, you must enable the multi-account governance maturity check feature.
	//
	// example:
	//
	// ResourceDirectory
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s RunEvaluationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEvaluationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunEvaluationShrinkRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *RunEvaluationShrinkRequest) GetMetricIdsShrink() *string {
	return s.MetricIdsShrink
}

func (s *RunEvaluationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunEvaluationShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *RunEvaluationShrinkRequest) SetAccountId(v int64) *RunEvaluationShrinkRequest {
	s.AccountId = &v
	return s
}

func (s *RunEvaluationShrinkRequest) SetMetricIdsShrink(v string) *RunEvaluationShrinkRequest {
	s.MetricIdsShrink = &v
	return s
}

func (s *RunEvaluationShrinkRequest) SetRegionId(v string) *RunEvaluationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunEvaluationShrinkRequest) SetScope(v string) *RunEvaluationShrinkRequest {
	s.Scope = &v
	return s
}

func (s *RunEvaluationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
