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
	SetEvaluationDomain(v string) *RunEvaluationShrinkRequest
	GetEvaluationDomain() *string
	SetMetricIdsShrink(v string) *RunEvaluationShrinkRequest
	GetMetricIdsShrink() *string
	SetRegionId(v string) *RunEvaluationShrinkRequest
	GetRegionId() *string
	SetScope(v string) *RunEvaluationShrinkRequest
	GetScope() *string
}

type RunEvaluationShrinkRequest struct {
	// The ID of the member account. This parameter is applicable only to the multi-account check pattern.
	//
	// example:
	//
	// 176618589410****
	AccountId        *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	EvaluationDomain *string `json:"EvaluationDomain,omitempty" xml:"EvaluationDomain,omitempty"`
	// The list of check item IDs to check.
	MetricIdsShrink *string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of the governance maturity check. Valid values:
	//
	// - Account (default): runs a single-account governance maturity check that checks only the current account.
	//
	// - ResourceDirectory: runs a multi-account governance maturity check that checks all member accounts in the resource directory. Before you perform this operation, upgrade to the multi-account governance maturity check.
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

func (s *RunEvaluationShrinkRequest) GetEvaluationDomain() *string {
	return s.EvaluationDomain
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

func (s *RunEvaluationShrinkRequest) SetEvaluationDomain(v string) *RunEvaluationShrinkRequest {
	s.EvaluationDomain = &v
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
