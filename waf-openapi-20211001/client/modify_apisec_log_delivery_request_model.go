// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssertKey(v string) *ModifyApisecLogDeliveryRequest
	GetAssertKey() *string
	SetInstanceId(v string) *ModifyApisecLogDeliveryRequest
	GetInstanceId() *string
	SetLogRegionId(v string) *ModifyApisecLogDeliveryRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *ModifyApisecLogDeliveryRequest
	GetLogStoreName() *string
	SetProjectName(v string) *ModifyApisecLogDeliveryRequest
	GetProjectName() *string
	SetRegionId(v string) *ModifyApisecLogDeliveryRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecLogDeliveryRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyApisecLogDeliveryRequest struct {
	// The type of the log subscription. Valid values:
	//
	// 	- **risk**: risk information.
	//
	// 	- **event**: attack event information.
	//
	// 	- **asset**: asset information.
	//
	// This parameter is required.
	//
	// example:
	//
	// risk
	AssertKey *string `json:"AssertKey,omitempty" xml:"AssertKey,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3_public_cn-uqm2z****0a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where logs are stored.
	//
	// >  You can call the [DescribeUserSlsLogRegions](https://help.aliyun.com/document_detail/2712598.html) operation to query available log storage regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The name of the Logstore in Simple Log Service.
	//
	// >  API security logs can be delivered only to Logstores whose names start with apisec-.
	//
	// This parameter is required.
	//
	// example:
	//
	// apisec-logstore***
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// >  API security logs can be delivered only to projects whose names start with apisec-.
	//
	// This parameter is required.
	//
	// example:
	//
	// apisec-project-14316572********
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
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

func (s ModifyApisecLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecLogDeliveryRequest) GetAssertKey() *string {
	return s.AssertKey
}

func (s *ModifyApisecLogDeliveryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecLogDeliveryRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ModifyApisecLogDeliveryRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ModifyApisecLogDeliveryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ModifyApisecLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecLogDeliveryRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecLogDeliveryRequest) SetAssertKey(v string) *ModifyApisecLogDeliveryRequest {
	s.AssertKey = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) SetInstanceId(v string) *ModifyApisecLogDeliveryRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) SetLogRegionId(v string) *ModifyApisecLogDeliveryRequest {
	s.LogRegionId = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) SetLogStoreName(v string) *ModifyApisecLogDeliveryRequest {
	s.LogStoreName = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) SetProjectName(v string) *ModifyApisecLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) SetRegionId(v string) *ModifyApisecLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecLogDeliveryRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
