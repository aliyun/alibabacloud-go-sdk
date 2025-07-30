// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterClass(v string) *CreateElasticRuleRequest
	GetClusterClass() *string
	SetClusterId(v string) *CreateElasticRuleRequest
	GetClusterId() *string
	SetDbInstanceId(v string) *CreateElasticRuleRequest
	GetDbInstanceId() *string
	SetElasticRuleStartTime(v string) *CreateElasticRuleRequest
	GetElasticRuleStartTime() *string
	SetExecutionPeriod(v string) *CreateElasticRuleRequest
	GetExecutionPeriod() *string
	SetRegionId(v string) *CreateElasticRuleRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *CreateElasticRuleRequest
	GetResourceOwnerId() *int64
}

type CreateElasticRuleRequest struct {
	// The rule for computing resources of the required cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb.2xlarge
	ClusterClass *string `json:"ClusterClass,omitempty" xml:"ClusterClass,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The time when you want to execute the scheduled scaling rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00
	ElasticRuleStartTime *string `json:"ElasticRuleStartTime,omitempty" xml:"ElasticRuleStartTime,omitempty"`
	// The execution cycle.
	//
	// Valid value:
	//
	// 	- Day
	//
	// This parameter is required.
	//
	// example:
	//
	// Day
	ExecutionPeriod *string `json:"ExecutionPeriod,omitempty" xml:"ExecutionPeriod,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateElasticRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticRuleRequest) GetClusterClass() *string {
	return s.ClusterClass
}

func (s *CreateElasticRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateElasticRuleRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateElasticRuleRequest) GetElasticRuleStartTime() *string {
	return s.ElasticRuleStartTime
}

func (s *CreateElasticRuleRequest) GetExecutionPeriod() *string {
	return s.ExecutionPeriod
}

func (s *CreateElasticRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateElasticRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateElasticRuleRequest) SetClusterClass(v string) *CreateElasticRuleRequest {
	s.ClusterClass = &v
	return s
}

func (s *CreateElasticRuleRequest) SetClusterId(v string) *CreateElasticRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateElasticRuleRequest) SetDbInstanceId(v string) *CreateElasticRuleRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateElasticRuleRequest) SetElasticRuleStartTime(v string) *CreateElasticRuleRequest {
	s.ElasticRuleStartTime = &v
	return s
}

func (s *CreateElasticRuleRequest) SetExecutionPeriod(v string) *CreateElasticRuleRequest {
	s.ExecutionPeriod = &v
	return s
}

func (s *CreateElasticRuleRequest) SetRegionId(v string) *CreateElasticRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateElasticRuleRequest) SetResourceOwnerId(v int64) *CreateElasticRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateElasticRuleRequest) Validate() error {
	return dara.Validate(s)
}
