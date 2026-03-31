// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecModuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyApisecModuleStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyApisecModuleStatusRequest
	GetRegionId() *string
	SetReportStatus(v int64) *ModifyApisecModuleStatusRequest
	GetReportStatus() *int64
	SetResourceGroups(v string) *ModifyApisecModuleStatusRequest
	GetResourceGroups() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecModuleStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetResources(v string) *ModifyApisecModuleStatusRequest
	GetResources() *string
	SetTraceStatus(v int32) *ModifyApisecModuleStatusRequest
	GetTraceStatus() *int32
}

type ModifyApisecModuleStatusRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqtm**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the compliance check feature. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	ReportStatus *int64 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The name of the protected object group to which the protected object belongs.
	//
	// example:
	//
	// group1
	ResourceGroups *string `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The name of the protected object.
	//
	// example:
	//
	// cwaf-***-waf
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The status of the tracing and auditing feature. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	TraceStatus *int32 `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
}

func (s ModifyApisecModuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecModuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecModuleStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecModuleStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecModuleStatusRequest) GetReportStatus() *int64 {
	return s.ReportStatus
}

func (s *ModifyApisecModuleStatusRequest) GetResourceGroups() *string {
	return s.ResourceGroups
}

func (s *ModifyApisecModuleStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecModuleStatusRequest) GetResources() *string {
	return s.Resources
}

func (s *ModifyApisecModuleStatusRequest) GetTraceStatus() *int32 {
	return s.TraceStatus
}

func (s *ModifyApisecModuleStatusRequest) SetInstanceId(v string) *ModifyApisecModuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) SetRegionId(v string) *ModifyApisecModuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) SetReportStatus(v int64) *ModifyApisecModuleStatusRequest {
	s.ReportStatus = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) SetResourceGroups(v string) *ModifyApisecModuleStatusRequest {
	s.ResourceGroups = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecModuleStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) SetResources(v string) *ModifyApisecModuleStatusRequest {
	s.Resources = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) SetTraceStatus(v int32) *ModifyApisecModuleStatusRequest {
	s.TraceStatus = &v
	return s
}

func (s *ModifyApisecModuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
