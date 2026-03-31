// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogDeliveryConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryNameLike(v string) *DescribeLogDeliveryConfigsRequest
	GetDeliveryNameLike() *string
	SetDeliveryType(v string) *DescribeLogDeliveryConfigsRequest
	GetDeliveryType() *string
	SetInstanceId(v string) *DescribeLogDeliveryConfigsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeLogDeliveryConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLogDeliveryConfigsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeLogDeliveryConfigsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeLogDeliveryConfigsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeLogDeliveryConfigsRequest struct {
	// The name of the log delivery configuration. Fuzzy queries are supported.
	//
	// example:
	//
	// test
	DeliveryNameLike *string `json:"DeliveryNameLike,omitempty" xml:"DeliveryNameLike,omitempty"`
	// The type of the log delivery configuration that you want to query. Valid values:
	//
	// 	- **syslog**: Logs are delivered to a syslog service.
	//
	// 	- **kafka**: Logs are delivered to a Kafka service.
	//
	// example:
	//
	// kafka
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeLogDeliveryConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogDeliveryConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogDeliveryConfigsRequest) GetDeliveryNameLike() *string {
	return s.DeliveryNameLike
}

func (s *DescribeLogDeliveryConfigsRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeLogDeliveryConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLogDeliveryConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLogDeliveryConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLogDeliveryConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogDeliveryConfigsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeLogDeliveryConfigsRequest) SetDeliveryNameLike(v string) *DescribeLogDeliveryConfigsRequest {
	s.DeliveryNameLike = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) SetDeliveryType(v string) *DescribeLogDeliveryConfigsRequest {
	s.DeliveryType = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) SetInstanceId(v string) *DescribeLogDeliveryConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) SetMaxResults(v int32) *DescribeLogDeliveryConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) SetNextToken(v string) *DescribeLogDeliveryConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) SetRegionId(v string) *DescribeLogDeliveryConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) SetResourceManagerResourceGroupId(v string) *DescribeLogDeliveryConfigsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeLogDeliveryConfigsRequest) Validate() error {
	return dara.Validate(s)
}
