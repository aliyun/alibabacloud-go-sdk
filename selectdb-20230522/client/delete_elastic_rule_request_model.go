// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteElasticRuleRequest
	GetClusterId() *string
	SetDbInstanceId(v string) *DeleteElasticRuleRequest
	GetDbInstanceId() *string
	SetProduct(v string) *DeleteElasticRuleRequest
	GetProduct() *string
	SetRegionId(v string) *DeleteElasticRuleRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DeleteElasticRuleRequest
	GetResourceOwnerId() *int64
	SetRuleId(v int64) *DeleteElasticRuleRequest
	GetRuleId() *int64
}

type DeleteElasticRuleRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxx302i5-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
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
	// cn-hanghzou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100458
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteElasticRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteElasticRuleRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DeleteElasticRuleRequest) GetProduct() *string {
	return s.Product
}

func (s *DeleteElasticRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteElasticRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteElasticRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteElasticRuleRequest) SetClusterId(v string) *DeleteElasticRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteElasticRuleRequest) SetDbInstanceId(v string) *DeleteElasticRuleRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DeleteElasticRuleRequest) SetProduct(v string) *DeleteElasticRuleRequest {
	s.Product = &v
	return s
}

func (s *DeleteElasticRuleRequest) SetRegionId(v string) *DeleteElasticRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteElasticRuleRequest) SetResourceOwnerId(v int64) *DeleteElasticRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteElasticRuleRequest) SetRuleId(v int64) *DeleteElasticRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteElasticRuleRequest) Validate() error {
	return dara.Validate(s)
}
