// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEvaluationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *RunEvaluationRequest
	GetAccountId() *int64
	SetEvaluationDomain(v string) *RunEvaluationRequest
	GetEvaluationDomain() *string
	SetMetricIds(v []*string) *RunEvaluationRequest
	GetMetricIds() []*string
	SetRegionId(v string) *RunEvaluationRequest
	GetRegionId() *string
	SetScope(v string) *RunEvaluationRequest
	GetScope() *string
}

type RunEvaluationRequest struct {
	// The ID of the member account. This parameter is applicable only to the multi-account check pattern.
	//
	// example:
	//
	// 176618589410****
	AccountId        *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	EvaluationDomain *string `json:"EvaluationDomain,omitempty" xml:"EvaluationDomain,omitempty"`
	// The list of check item IDs to check.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
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

func (s RunEvaluationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEvaluationRequest) GoString() string {
	return s.String()
}

func (s *RunEvaluationRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *RunEvaluationRequest) GetEvaluationDomain() *string {
	return s.EvaluationDomain
}

func (s *RunEvaluationRequest) GetMetricIds() []*string {
	return s.MetricIds
}

func (s *RunEvaluationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunEvaluationRequest) GetScope() *string {
	return s.Scope
}

func (s *RunEvaluationRequest) SetAccountId(v int64) *RunEvaluationRequest {
	s.AccountId = &v
	return s
}

func (s *RunEvaluationRequest) SetEvaluationDomain(v string) *RunEvaluationRequest {
	s.EvaluationDomain = &v
	return s
}

func (s *RunEvaluationRequest) SetMetricIds(v []*string) *RunEvaluationRequest {
	s.MetricIds = v
	return s
}

func (s *RunEvaluationRequest) SetRegionId(v string) *RunEvaluationRequest {
	s.RegionId = &v
	return s
}

func (s *RunEvaluationRequest) SetScope(v string) *RunEvaluationRequest {
	s.Scope = &v
	return s
}

func (s *RunEvaluationRequest) Validate() error {
	return dara.Validate(s)
}
