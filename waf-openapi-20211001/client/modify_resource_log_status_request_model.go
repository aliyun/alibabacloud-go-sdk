// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyResourceLogStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyResourceLogStatusRequest
	GetRegionId() *string
	SetResource(v string) *ModifyResourceLogStatusRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyResourceLogStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetStatus(v bool) *ModifyResourceLogStatusRequest
	GetStatus() *bool
	SetTraceConfig(v *ModifyResourceLogStatusRequestTraceConfig) *ModifyResourceLogStatusRequest
	GetTraceConfig() *ModifyResourceLogStatusRequestTraceConfig
	SetTraceStatus(v bool) *ModifyResourceLogStatusRequest
	GetTraceStatus() *bool
}

type ModifyResourceLogStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11zcl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object on which you want to manage the log collection feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-wewbb23dfsetetcic1242-0****
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Specifies whether to enable the log collection feature for the protected object. Valid values:
	//
	// 	- **true:*	- enables the log collection feature.
	//
	// 	- **false:*	- disables the log collection feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status      *bool                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	TraceConfig *ModifyResourceLogStatusRequestTraceConfig `json:"TraceConfig,omitempty" xml:"TraceConfig,omitempty" type:"Struct"`
	TraceStatus *bool                                      `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
}

func (s ModifyResourceLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyResourceLogStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceLogStatusRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyResourceLogStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyResourceLogStatusRequest) GetStatus() *bool {
	return s.Status
}

func (s *ModifyResourceLogStatusRequest) GetTraceConfig() *ModifyResourceLogStatusRequestTraceConfig {
	return s.TraceConfig
}

func (s *ModifyResourceLogStatusRequest) GetTraceStatus() *bool {
	return s.TraceStatus
}

func (s *ModifyResourceLogStatusRequest) SetInstanceId(v string) *ModifyResourceLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetRegionId(v string) *ModifyResourceLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetResource(v string) *ModifyResourceLogStatusRequest {
	s.Resource = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyResourceLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetStatus(v bool) *ModifyResourceLogStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetTraceConfig(v *ModifyResourceLogStatusRequestTraceConfig) *ModifyResourceLogStatusRequest {
	s.TraceConfig = v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetTraceStatus(v bool) *ModifyResourceLogStatusRequest {
	s.TraceStatus = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) Validate() error {
	if s.TraceConfig != nil {
		if err := s.TraceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyResourceLogStatusRequestTraceConfig struct {
	// example:
	//
	// 0
	RatePerMille *int32 `json:"RatePerMille,omitempty" xml:"RatePerMille,omitempty"`
	// example:
	//
	// cms-test
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s ModifyResourceLogStatusRequestTraceConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogStatusRequestTraceConfig) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusRequestTraceConfig) GetRatePerMille() *int32 {
	return s.RatePerMille
}

func (s *ModifyResourceLogStatusRequestTraceConfig) GetWorkspace() *string {
	return s.Workspace
}

func (s *ModifyResourceLogStatusRequestTraceConfig) SetRatePerMille(v int32) *ModifyResourceLogStatusRequestTraceConfig {
	s.RatePerMille = &v
	return s
}

func (s *ModifyResourceLogStatusRequestTraceConfig) SetWorkspace(v string) *ModifyResourceLogStatusRequestTraceConfig {
	s.Workspace = &v
	return s
}

func (s *ModifyResourceLogStatusRequestTraceConfig) Validate() error {
	return dara.Validate(s)
}
