// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterClass(v string) *ModifyElasticRuleRequest
	GetClusterClass() *string
	SetClusterId(v string) *ModifyElasticRuleRequest
	GetClusterId() *string
	SetDbInstanceId(v string) *ModifyElasticRuleRequest
	GetDbInstanceId() *string
	SetElasticRuleStartTime(v string) *ModifyElasticRuleRequest
	GetElasticRuleStartTime() *string
	SetExecutionPeriod(v string) *ModifyElasticRuleRequest
	GetExecutionPeriod() *string
	SetProduct(v string) *ModifyElasticRuleRequest
	GetProduct() *string
	SetRegionId(v string) *ModifyElasticRuleRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ModifyElasticRuleRequest
	GetResourceOwnerId() *int64
	SetRuleId(v int64) *ModifyElasticRuleRequest
	GetRuleId() *int64
}

type ModifyElasticRuleRequest struct {
	// The rule for computing resources of the required cluster.
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
	// selectdb-cn-nwy3jv1oa02-be
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
	// example:
	//
	// Day
	ExecutionPeriod *string `json:"ExecutionPeriod,omitempty" xml:"ExecutionPeriod,omitempty"`
	// The cloud service.
	//
	// example:
	//
	// selectdb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5467
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyElasticRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticRuleRequest) GetClusterClass() *string {
	return s.ClusterClass
}

func (s *ModifyElasticRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyElasticRuleRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyElasticRuleRequest) GetElasticRuleStartTime() *string {
	return s.ElasticRuleStartTime
}

func (s *ModifyElasticRuleRequest) GetExecutionPeriod() *string {
	return s.ExecutionPeriod
}

func (s *ModifyElasticRuleRequest) GetProduct() *string {
	return s.Product
}

func (s *ModifyElasticRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyElasticRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyElasticRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyElasticRuleRequest) SetClusterClass(v string) *ModifyElasticRuleRequest {
	s.ClusterClass = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetClusterId(v string) *ModifyElasticRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetDbInstanceId(v string) *ModifyElasticRuleRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetElasticRuleStartTime(v string) *ModifyElasticRuleRequest {
	s.ElasticRuleStartTime = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetExecutionPeriod(v string) *ModifyElasticRuleRequest {
	s.ExecutionPeriod = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetProduct(v string) *ModifyElasticRuleRequest {
	s.Product = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetRegionId(v string) *ModifyElasticRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetResourceOwnerId(v int64) *ModifyElasticRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyElasticRuleRequest) SetRuleId(v int64) *ModifyElasticRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyElasticRuleRequest) Validate() error {
	return dara.Validate(s)
}
