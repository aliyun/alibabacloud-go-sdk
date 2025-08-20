// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableElasticPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableElasticPlanRequest
  GetDBClusterId() *string 
  SetElasticPlanName(v string) *EnableElasticPlanRequest
  GetElasticPlanName() *string 
}

type EnableElasticPlanRequest struct {
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

func (s EnableElasticPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableElasticPlanRequest) GoString() string {
  return s.String()
}

func (s *EnableElasticPlanRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableElasticPlanRequest) GetElasticPlanName() *string  {
  return s.ElasticPlanName
}

func (s *EnableElasticPlanRequest) SetDBClusterId(v string) *EnableElasticPlanRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableElasticPlanRequest) SetElasticPlanName(v string) *EnableElasticPlanRequest {
  s.ElasticPlanName = &v
  return s
}

func (s *EnableElasticPlanRequest) Validate() error {
  return dara.Validate(s)
}

