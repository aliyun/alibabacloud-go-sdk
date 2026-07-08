// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogFieldConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryType(v string) *DescribeResourceLogFieldConfigRequest
	GetDeliveryType() *string
	SetInstanceId(v string) *DescribeResourceLogFieldConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeResourceLogFieldConfigRequest
	GetRegionId() *string
	SetResource(v string) *DescribeResourceLogFieldConfigRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourceLogFieldConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeResourceLogFieldConfigRequest struct {
	// The log delivery type. Valid values:
	//
	// - **sls**: Simple Log Service.
	//
	// - **kafka**: Kafka.
	//
	// - **syslog**: Syslog.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zpr3*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object whose log field configuration you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.waf.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group to which the WAF instance belongs.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourceLogFieldConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogFieldConfigRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeResourceLogFieldConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceLogFieldConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceLogFieldConfigRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeResourceLogFieldConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourceLogFieldConfigRequest) SetDeliveryType(v string) *DescribeResourceLogFieldConfigRequest {
	s.DeliveryType = &v
	return s
}

func (s *DescribeResourceLogFieldConfigRequest) SetInstanceId(v string) *DescribeResourceLogFieldConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceLogFieldConfigRequest) SetRegionId(v string) *DescribeResourceLogFieldConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceLogFieldConfigRequest) SetResource(v string) *DescribeResourceLogFieldConfigRequest {
	s.Resource = &v
	return s
}

func (s *DescribeResourceLogFieldConfigRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceLogFieldConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceLogFieldConfigRequest) Validate() error {
	return dara.Validate(s)
}
