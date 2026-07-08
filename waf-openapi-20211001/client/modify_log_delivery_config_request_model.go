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
	// The details of the log delivery configuration, in JSON format.
	//
	// > The value of this parameter is the same as the **DeliveryDetail*	- parameter of the **CreateLogDeliveryConfig*	- operation. For more information, see [CreateLogDeliveryConfig]().
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "rfcVersion": "rfc3164",
	//
	//   "protocol": "tcp",
	//
	//   "servers": [
	//
	//     {
	//
	//       "address": "1.1.1.1",
	//
	//       "port": 20
	//
	//     }
	//
	//   ]
	//
	// }
	DeliveryDetail *string `json:"DeliveryDetail,omitempty" xml:"DeliveryDetail,omitempty"`
	// The name of the log delivery configuration that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The type of the log delivery destination. Valid values:
	//
	// - **syslog**: delivers logs to a syslog server.
	//
	// - **kafka**: delivers logs to a Kafka cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// syslog
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
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
