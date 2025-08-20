// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeElasticPlanAttributeRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *DescribeElasticPlanAttributeRequest
	GetElasticPlanName() *string
}

type DescribeElasticPlanAttributeRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >  You can call the [DescribeElasticPlans](https://help.aliyun.com/document_detail/601334.html) operation to query the names of scaling plans.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
}

func (s DescribeElasticPlanAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeElasticPlanAttributeRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlanAttributeRequest) SetDBClusterId(v string) *DescribeElasticPlanAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanAttributeRequest) SetElasticPlanName(v string) *DescribeElasticPlanAttributeRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanAttributeRequest) Validate() error {
	return dara.Validate(s)
}
