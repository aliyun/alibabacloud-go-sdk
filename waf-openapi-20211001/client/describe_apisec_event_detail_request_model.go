// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecEventDetailRequest
	GetClusterId() *string
	SetDetailType(v string) *DescribeApisecEventDetailRequest
	GetDetailType() *string
	SetEventId(v string) *DescribeApisecEventDetailRequest
	GetEventId() *string
	SetEventScope(v string) *DescribeApisecEventDetailRequest
	GetEventScope() *string
	SetInstanceId(v string) *DescribeApisecEventDetailRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeApisecEventDetailRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecEventDetailRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecEventDetailRequest struct {
	// The ID of the Hybrid Cloud WAF cluster.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query information about hybrid cloud WAF clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the detailed information about the security event. Valid values:
	//
	// - **event_info*	- (default): attack details.
	//
	// - **api_info**: API information.
	//
	// - **cnt_info**: attack trend.
	//
	// - **ip_info**: attacker IP information.
	//
	// - **sensitive_info**: information about access to sensitive data.
	//
	// - **request_data**: request information.
	//
	// - **response_data**: response information.
	//
	// example:
	//
	// event_info
	DetailType *string `json:"DetailType,omitempty" xml:"DetailType,omitempty"`
	// The ID of the API security event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18ba94fea9***e66ba0557b7b91
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The dimension of the security event. Valid values:
	//
	// - **ip*	- (default): IP security event.
	//
	// - **account**: account security event.
	//
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqtm005
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeApisecEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventDetailRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecEventDetailRequest) GetDetailType() *string {
	return s.DetailType
}

func (s *DescribeApisecEventDetailRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeApisecEventDetailRequest) GetEventScope() *string {
	return s.EventScope
}

func (s *DescribeApisecEventDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecEventDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecEventDetailRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecEventDetailRequest) SetClusterId(v string) *DescribeApisecEventDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) SetDetailType(v string) *DescribeApisecEventDetailRequest {
	s.DetailType = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) SetEventId(v string) *DescribeApisecEventDetailRequest {
	s.EventId = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) SetEventScope(v string) *DescribeApisecEventDetailRequest {
	s.EventScope = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) SetInstanceId(v string) *DescribeApisecEventDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) SetRegionId(v string) *DescribeApisecEventDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecEventDetailRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
