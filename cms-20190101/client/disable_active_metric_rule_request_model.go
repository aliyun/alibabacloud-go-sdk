// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableActiveMetricRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *DisableActiveMetricRuleRequest
	GetProduct() *string
	SetRegionId(v string) *DisableActiveMetricRuleRequest
	GetRegionId() *string
}

type DisableActiveMetricRuleRequest struct {
	// The cloud service for which you want to disable proactive alerting. Valid values:
	//
	// 	- ECS: Elastic Compute Service (ECS)
	//
	// 	- rds: ApsaraDB RDS
	//
	// 	- slb: Server Load Balancer (SLB)
	//
	// 	- redis_standard: Redis Open-Source Edition (standard architecture)
	//
	// 	- redis_sharding: Redis Open-Source Edition (cluster architecture)
	//
	// 	- redis_splitrw: Redis Open-Source Edition (read/write splitting architecture)
	//
	// 	- mongodb: ApsaraDB for MongoDB of the replica set architecture
	//
	// 	- mongodb_sharding: ApsaraDB for MongoDB of the sharded cluster architecture
	//
	// 	- hbase: ApsaraDB for HBase
	//
	// 	- elasticsearch: Elasticsearch
	//
	// 	- opensearch: OpenSearch
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableActiveMetricRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableActiveMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableActiveMetricRuleRequest) GetProduct() *string {
	return s.Product
}

func (s *DisableActiveMetricRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableActiveMetricRuleRequest) SetProduct(v string) *DisableActiveMetricRuleRequest {
	s.Product = &v
	return s
}

func (s *DisableActiveMetricRuleRequest) SetRegionId(v string) *DisableActiveMetricRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DisableActiveMetricRuleRequest) Validate() error {
	return dara.Validate(s)
}
