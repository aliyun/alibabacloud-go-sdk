// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableActiveMetricRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableActiveMetricRuleRequest
  GetClientToken() *string 
  SetDryRun(v bool) *EnableActiveMetricRuleRequest
  GetDryRun() *bool 
  SetProduct(v string) *EnableActiveMetricRuleRequest
  GetProduct() *string 
  SetRegionId(v string) *EnableActiveMetricRuleRequest
  GetRegionId() *string 
}

type EnableActiveMetricRuleRequest struct {
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can generate the value of this parameter on the client side. Make sure that the value is unique across different requests. The value can be up to 64 ASCII characters in length.
  // 
  // example:
  // 
  // 00000-xxxx-xxxxx--xxxx
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // Specifies whether to check the availability of the webhook.
  // 
  // example:
  // 
  // false
  DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
  // The name of the Alibaba Cloud service that supports one-click alert rules. Valid values:
  // 
  // 	- ecs: Elastic Compute Service (ECS).
  // 
  // 	- rds: ApsaraDB RDS.
  // 
  // 	- slb: Server Load Balancer (SLB).
  // 
  // 	- redis_standard: Standard architecture of open-source Redis.
  // 
  // 	- redis_sharding: Cluster architecture of open-source Redis.
  // 
  // 	- redis_splitrw: Read/write splitting architecture of open-source Redis.
  // 
  // 	- mongodb: ApsaraDB for MongoDB (replica set).
  // 
  // 	- mongodb_sharding: ApsaraDB for MongoDB (sharded cluster).
  // 
  // 	- hbase: ApsaraDB for HBase.
  // 
  // 	- elasticsearch: Elasticsearch.
  // 
  // 	- opensearch: OpenSearch.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ecs
  Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableActiveMetricRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableActiveMetricRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableActiveMetricRuleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableActiveMetricRuleRequest) GetDryRun() *bool  {
  return s.DryRun
}

func (s *EnableActiveMetricRuleRequest) GetProduct() *string  {
  return s.Product
}

func (s *EnableActiveMetricRuleRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableActiveMetricRuleRequest) SetClientToken(v string) *EnableActiveMetricRuleRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableActiveMetricRuleRequest) SetDryRun(v bool) *EnableActiveMetricRuleRequest {
  s.DryRun = &v
  return s
}

func (s *EnableActiveMetricRuleRequest) SetProduct(v string) *EnableActiveMetricRuleRequest {
  s.Product = &v
  return s
}

func (s *EnableActiveMetricRuleRequest) SetRegionId(v string) *EnableActiveMetricRuleRequest {
  s.RegionId = &v
  return s
}

func (s *EnableActiveMetricRuleRequest) Validate() error {
  return dara.Validate(s)
}

