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
	SetMetricIds(v []*string) *RunEvaluationRequest
	GetMetricIds() []*string
	SetRegionId(v string) *RunEvaluationRequest
	GetRegionId() *string
	SetScope(v string) *RunEvaluationRequest
	GetScope() *string
}

type RunEvaluationRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The IDs of the check items to be checked.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
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

func (s RunEvaluationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunEvaluationRequest) GoString() string {
	return s.String()
}

func (s *RunEvaluationRequest) GetAccountId() *int64 {
	return s.AccountId
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
