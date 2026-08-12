// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserLogFieldConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryType(v string) *ModifyUserLogFieldConfigRequest
	GetDeliveryType() *string
	SetExtendConfig(v string) *ModifyUserLogFieldConfigRequest
	GetExtendConfig() *string
	SetFieldList(v string) *ModifyUserLogFieldConfigRequest
	GetFieldList() *string
	SetInstanceId(v string) *ModifyUserLogFieldConfigRequest
	GetInstanceId() *string
	SetLogDeliveryStrategy(v string) *ModifyUserLogFieldConfigRequest
	GetLogDeliveryStrategy() *string
	SetRegionId(v string) *ModifyUserLogFieldConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyUserLogFieldConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyUserLogFieldConfigRequest struct {
	// The delivery type. Valid values:
	//
	// - **sls**: Simple Log Service.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The extended configuration for log delivery. The value is a JSON string constructed from a set of parameters.
	//
	// > For more information, see the **Log delivery extended configuration*	- parameter description.
	//
	// example:
	//
	// {\\"request_header\\":\\"App-Id,channelCode\\"}
	ExtendConfig *string `json:"ExtendConfig,omitempty" xml:"ExtendConfig,omitempty"`
	// The list of log fields to deliver. Specify the fields in the "a,b,c,..." format.
	//
	// >   - All required log fields must be included. You can invoke the [DescribeCommonLogFields](~~DescribeCommonLogFields~~) operation to query the log fields supported by Simple Log Service for WAF.
	//
	// > - If the log fields include **request_header**, use the **delivery extension configuration*	- (**ExtendConfig**) parameter to specify the request headers to deliver.
	//
	// This parameter is required.
	//
	// example:
	//
	// account,acl_action,acl_rule_id,acl_rule_type,acl_test,antiscan_action,antiscan_rule_id,antiscan_rule_type,antiscan_test,body_bytes_sent,bypass_matched_ids
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-fou****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The log delivery strategy. Multiple strategies are supported. The value is a JSON array string constructed from a set of parameters.
	//
	// > For more information, see the **Log delivery strategy*	- parameter description.
	//
	// example:
	//
	// [{\\"logType\\":\\"blockLog\\",\\"rate\\":100},{\\"logType\\":\\"normalRequestLog\\",\\"rate\\":100},{\\"logType\\":\\"checkLog\\",\\"rate\\":100}]
	LogDeliveryStrategy *string `json:"LogDeliveryStrategy,omitempty" xml:"LogDeliveryStrategy,omitempty"`
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

func (s ModifyUserLogFieldConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserLogFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserLogFieldConfigRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *ModifyUserLogFieldConfigRequest) GetExtendConfig() *string {
	return s.ExtendConfig
}

func (s *ModifyUserLogFieldConfigRequest) GetFieldList() *string {
	return s.FieldList
}

func (s *ModifyUserLogFieldConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserLogFieldConfigRequest) GetLogDeliveryStrategy() *string {
	return s.LogDeliveryStrategy
}

func (s *ModifyUserLogFieldConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserLogFieldConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyUserLogFieldConfigRequest) SetDeliveryType(v string) *ModifyUserLogFieldConfigRequest {
	s.DeliveryType = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) SetExtendConfig(v string) *ModifyUserLogFieldConfigRequest {
	s.ExtendConfig = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) SetFieldList(v string) *ModifyUserLogFieldConfigRequest {
	s.FieldList = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) SetInstanceId(v string) *ModifyUserLogFieldConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) SetLogDeliveryStrategy(v string) *ModifyUserLogFieldConfigRequest {
	s.LogDeliveryStrategy = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) SetRegionId(v string) *ModifyUserLogFieldConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) SetResourceManagerResourceGroupId(v string) *ModifyUserLogFieldConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyUserLogFieldConfigRequest) Validate() error {
	return dara.Validate(s)
}
