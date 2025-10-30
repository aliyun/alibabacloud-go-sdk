// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogDeliveryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryDetail(v string) *ModifyLogDeliveryConfigRequest
	GetDeliveryDetail() *string
	SetDeliveryName(v string) *ModifyLogDeliveryConfigRequest
	GetDeliveryName() *string
	SetDeliveryType(v string) *ModifyLogDeliveryConfigRequest
	GetDeliveryType() *string
	SetInstanceId(v string) *ModifyLogDeliveryConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyLogDeliveryConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyLogDeliveryConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyLogDeliveryConfigRequest struct {
	// The content of the log delivery configuration. Set the value to a JSON string that contains multiple parameters.
	//
	// >  This parameter is the same as the **DeliveryDetail*	- parameter of the **CreateLogDeliveryConfig*	- operation. For more information, see **Parameter description for log delivery configuration*	- of the [CreateLogDeliveryConfig](~~CreateLogDeliveryConfig~~) operation.
	//
	// This parameter is required.
	DeliveryDetail *string `json:"DeliveryDetail,omitempty" xml:"DeliveryDetail,omitempty"`
	// The name of the log delivery configuration that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The type of the log delivery configuration that you want to modify. Valid values:
	//
	// 	- **syslog**: Logs are delivered to a syslog service.
	//
	// 	- **kafka**: Logs are delivered to a Kafka service.
	//
	// This parameter is required.
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
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s ModifyLogDeliveryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogDeliveryConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogDeliveryConfigRequest) GetDeliveryDetail() *string {
	return s.DeliveryDetail
}

func (s *ModifyLogDeliveryConfigRequest) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *ModifyLogDeliveryConfigRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *ModifyLogDeliveryConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyLogDeliveryConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLogDeliveryConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyLogDeliveryConfigRequest) SetDeliveryDetail(v string) *ModifyLogDeliveryConfigRequest {
	s.DeliveryDetail = &v
	return s
}

func (s *ModifyLogDeliveryConfigRequest) SetDeliveryName(v string) *ModifyLogDeliveryConfigRequest {
	s.DeliveryName = &v
	return s
}

func (s *ModifyLogDeliveryConfigRequest) SetDeliveryType(v string) *ModifyLogDeliveryConfigRequest {
	s.DeliveryType = &v
	return s
}

func (s *ModifyLogDeliveryConfigRequest) SetInstanceId(v string) *ModifyLogDeliveryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLogDeliveryConfigRequest) SetRegionId(v string) *ModifyLogDeliveryConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLogDeliveryConfigRequest) SetResourceManagerResourceGroupId(v string) *ModifyLogDeliveryConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyLogDeliveryConfigRequest) Validate() error {
	return dara.Validate(s)
}
