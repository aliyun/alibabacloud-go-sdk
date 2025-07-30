// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeElasticRulesRequest
	GetClusterId() *string
	SetDbInstanceId(v string) *DescribeElasticRulesRequest
	GetDbInstanceId() *string
	SetProduct(v string) *DescribeElasticRulesRequest
	GetProduct() *string
	SetRegionId(v string) *DescribeElasticRulesRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeElasticRulesRequest
	GetResourceOwnerId() *int64
}

type DescribeElasticRulesRequest struct {
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
}

func (s DescribeElasticRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeElasticRulesRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeElasticRulesRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeElasticRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticRulesRequest) SetClusterId(v string) *DescribeElasticRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeElasticRulesRequest) SetDbInstanceId(v string) *DescribeElasticRulesRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeElasticRulesRequest) SetProduct(v string) *DescribeElasticRulesRequest {
	s.Product = &v
	return s
}

func (s *DescribeElasticRulesRequest) SetRegionId(v string) *DescribeElasticRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticRulesRequest) SetResourceOwnerId(v int64) *DescribeElasticRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticRulesRequest) Validate() error {
	return dara.Validate(s)
}
