// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *DeleteElasticPlanRequest
	GetElasticPlanName() *string
}

type DeleteElasticPlanRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
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

func (s DeleteElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DeleteElasticPlanRequest) SetDBClusterId(v string) *DeleteElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetElasticPlanName(v string) *DeleteElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DeleteElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
