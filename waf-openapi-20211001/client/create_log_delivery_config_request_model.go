// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogDeliveryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryDetail(v string) *CreateLogDeliveryConfigRequest
	GetDeliveryDetail() *string
	SetDeliveryName(v string) *CreateLogDeliveryConfigRequest
	GetDeliveryName() *string
	SetDeliveryType(v string) *CreateLogDeliveryConfigRequest
	GetDeliveryType() *string
	SetInstanceId(v string) *CreateLogDeliveryConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateLogDeliveryConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateLogDeliveryConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type CreateLogDeliveryConfigRequest struct {
	// The content of the log delivery configuration. Set the value to a JSON string that contains multiple parameters.
	//
	// >  The parameters vary based on the type of the **log delivery configuration*	- specified by **DeliveryType**. For more information, see **Parameter description for log delivery configuration**.
	//
	// This parameter is required.
	DeliveryDetail *string `json:"DeliveryDetail,omitempty" xml:"DeliveryDetail,omitempty"`
	// The name of the log delivery configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The type of the log delivery configuration. Valid values:
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
	// waf_elasticity-cn-0xldbqt****
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

func (s CreateLogDeliveryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogDeliveryConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLogDeliveryConfigRequest) GetDeliveryDetail() *string {
	return s.DeliveryDetail
}

func (s *CreateLogDeliveryConfigRequest) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *CreateLogDeliveryConfigRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *CreateLogDeliveryConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLogDeliveryConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLogDeliveryConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateLogDeliveryConfigRequest) SetDeliveryDetail(v string) *CreateLogDeliveryConfigRequest {
	s.DeliveryDetail = &v
	return s
}

func (s *CreateLogDeliveryConfigRequest) SetDeliveryName(v string) *CreateLogDeliveryConfigRequest {
	s.DeliveryName = &v
	return s
}

func (s *CreateLogDeliveryConfigRequest) SetDeliveryType(v string) *CreateLogDeliveryConfigRequest {
	s.DeliveryType = &v
	return s
}

func (s *CreateLogDeliveryConfigRequest) SetInstanceId(v string) *CreateLogDeliveryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLogDeliveryConfigRequest) SetRegionId(v string) *CreateLogDeliveryConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLogDeliveryConfigRequest) SetResourceManagerResourceGroupId(v string) *CreateLogDeliveryConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateLogDeliveryConfigRequest) Validate() error {
	return dara.Validate(s)
}
