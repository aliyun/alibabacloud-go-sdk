// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogFieldConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryType(v string) *ModifyResourceLogFieldConfigRequest
	GetDeliveryType() *string
	SetExtendConfig(v string) *ModifyResourceLogFieldConfigRequest
	GetExtendConfig() *string
	SetFieldList(v string) *ModifyResourceLogFieldConfigRequest
	GetFieldList() *string
	SetInstanceId(v string) *ModifyResourceLogFieldConfigRequest
	GetInstanceId() *string
	SetLogDeliveryStrategy(v string) *ModifyResourceLogFieldConfigRequest
	GetLogDeliveryStrategy() *string
	SetRegionId(v string) *ModifyResourceLogFieldConfigRequest
	GetRegionId() *string
	SetResource(v string) *ModifyResourceLogFieldConfigRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyResourceLogFieldConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyResourceLogFieldConfigRequest struct {
	// The log delivery type. Valid values:
	//
	// - **sls**: Simple Log Service.
	//
	// - **kafka**: external delivery to Kafka logs.
	//
	// - **syslog**: external delivery to Syslog logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The log delivery extension configuration. The value is a string converted from a JSON object constructed with a series of parameters.
	//
	// > For more information, see the **Log delivery extension configuration*	- parameter description.
	//
	// example:
	//
	// {\\"request_header\\":\\"Ali-Cdn-Real-Ip\\"}
	ExtendConfig *string `json:"ExtendConfig,omitempty" xml:"ExtendConfig,omitempty"`
	// The list of log fields to deliver. Specify the fields in the "a,b,c,..." format.
	//
	// >   - All required log fields must be included. You can invoke the [DescribeCommonLogFields](~~DescribeCommonLogFields~~) operation to view the log fields supported by Simple Log Service for WAF.
	//
	// > - If the log fields include **request_header**, use the **delivery extension configuration*	- (**ExtendConfig**) parameter to specify the request headers to deliver.
	//
	// This parameter is required.
	//
	// example:
	//
	// account,acl_action,acl_rule_id,acl_rule_type,acl_test,antiscan_action,antiscan_rule_id,antiscan_rule_type,antiscan_test,body_bytes_sent,bypass_matched_ids
	FieldList *string `json:"FieldList,omitempty" xml:"FieldList,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-uax****3k0e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The log delivery strategy. Multiple strategies are supported. The value is a string converted from a JSON array constructed with a series of parameters.
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
	// The protected object to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// cwaf-***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyResourceLogFieldConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogFieldConfigRequest) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *ModifyResourceLogFieldConfigRequest) GetExtendConfig() *string {
	return s.ExtendConfig
}

func (s *ModifyResourceLogFieldConfigRequest) GetFieldList() *string {
	return s.FieldList
}

func (s *ModifyResourceLogFieldConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyResourceLogFieldConfigRequest) GetLogDeliveryStrategy() *string {
	return s.LogDeliveryStrategy
}

func (s *ModifyResourceLogFieldConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceLogFieldConfigRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyResourceLogFieldConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyResourceLogFieldConfigRequest) SetDeliveryType(v string) *ModifyResourceLogFieldConfigRequest {
	s.DeliveryType = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetExtendConfig(v string) *ModifyResourceLogFieldConfigRequest {
	s.ExtendConfig = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetFieldList(v string) *ModifyResourceLogFieldConfigRequest {
	s.FieldList = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetInstanceId(v string) *ModifyResourceLogFieldConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetLogDeliveryStrategy(v string) *ModifyResourceLogFieldConfigRequest {
	s.LogDeliveryStrategy = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetRegionId(v string) *ModifyResourceLogFieldConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetResource(v string) *ModifyResourceLogFieldConfigRequest {
	s.Resource = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) SetResourceManagerResourceGroupId(v string) *ModifyResourceLogFieldConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyResourceLogFieldConfigRequest) Validate() error {
	return dara.Validate(s)
}
