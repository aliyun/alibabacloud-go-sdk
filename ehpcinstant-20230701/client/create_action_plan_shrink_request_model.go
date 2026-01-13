// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActionPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanName(v string) *CreateActionPlanShrinkRequest
	GetActionPlanName() *string
	SetAllocationSpec(v string) *CreateActionPlanShrinkRequest
	GetAllocationSpec() *string
	SetAppId(v string) *CreateActionPlanShrinkRequest
	GetAppId() *string
	SetDesiredCapacity(v float64) *CreateActionPlanShrinkRequest
	GetDesiredCapacity() *float64
	SetIntervalMinutes(v int32) *CreateActionPlanShrinkRequest
	GetIntervalMinutes() *int32
	SetLevel(v string) *CreateActionPlanShrinkRequest
	GetLevel() *string
	SetPrologScript(v string) *CreateActionPlanShrinkRequest
	GetPrologScript() *string
	SetRegionsShrink(v string) *CreateActionPlanShrinkRequest
	GetRegionsShrink() *string
	SetResourceType(v string) *CreateActionPlanShrinkRequest
	GetResourceType() *string
	SetResourcesShrink(v string) *CreateActionPlanShrinkRequest
	GetResourcesShrink() *string
	SetScript(v string) *CreateActionPlanShrinkRequest
	GetScript() *string
}

type CreateActionPlanShrinkRequest struct {
	// The name of the execution plan.
	//
	// example:
	//
	// TestActionPlan
	ActionPlanName *string `json:"ActionPlanName,omitempty" xml:"ActionPlanName,omitempty"`
	// The type of the resource.
	//
	// 	- Standard
	//
	// 	- Dedicated: You must enable a whitelist for use.
	//
	// 	- Economic: You must enable a whitelist for use.
	//
	// example:
	//
	// Standard
	AllocationSpec *string `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// ci-vm-rYfypJKwlN9Y
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The expected scale of resources for the execution plan. If the ResourceType parameter is set to VcpuCapacity, the execution plan is expected to have 10000 vCPUs.
	//
	// example:
	//
	// 1000
	DesiredCapacity *float64 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// example:
	//
	// 60
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
	// The computing power level. This value is valid only when the resource type is Economic. The following disk categories are supported:
	//
	// 	- General
	//
	// 	- Performance
	//
	// Default value: General
	//
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The pre-processing script. Base64 encoding is required.
	//
	// example:
	//
	// bHMgLWFsCmxzIC1hbGggfCB3YyAtbA==
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// The list of resource configurations in the region where the execution plan runs.
	RegionsShrink *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// Target resource type: the capacity of vCPUs or the number of execution nodes. Valid values:
	//
	// 	- VCpuCapacity
	//
	// 	- ExecutorCapacity
	//
	// example:
	//
	// VCpuCapacity
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of resource configurations of the execution plan runtime environment. You can configure 1 to 10 resources.
	//
	// example:
	//
	// 1000
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The running-job script. Base64 encoding is required.
	//
	// example:
	//
	// bHMgLWFsCmxzIC1hbGggfCB3YyAtbA==
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateActionPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateActionPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateActionPlanShrinkRequest) GetActionPlanName() *string {
	return s.ActionPlanName
}

func (s *CreateActionPlanShrinkRequest) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *CreateActionPlanShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateActionPlanShrinkRequest) GetDesiredCapacity() *float64 {
	return s.DesiredCapacity
}

func (s *CreateActionPlanShrinkRequest) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *CreateActionPlanShrinkRequest) GetLevel() *string {
	return s.Level
}

func (s *CreateActionPlanShrinkRequest) GetPrologScript() *string {
	return s.PrologScript
}

func (s *CreateActionPlanShrinkRequest) GetRegionsShrink() *string {
	return s.RegionsShrink
}

func (s *CreateActionPlanShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateActionPlanShrinkRequest) GetResourcesShrink() *string {
	return s.ResourcesShrink
}

func (s *CreateActionPlanShrinkRequest) GetScript() *string {
	return s.Script
}

func (s *CreateActionPlanShrinkRequest) SetActionPlanName(v string) *CreateActionPlanShrinkRequest {
	s.ActionPlanName = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetAllocationSpec(v string) *CreateActionPlanShrinkRequest {
	s.AllocationSpec = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetAppId(v string) *CreateActionPlanShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetDesiredCapacity(v float64) *CreateActionPlanShrinkRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetIntervalMinutes(v int32) *CreateActionPlanShrinkRequest {
	s.IntervalMinutes = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetLevel(v string) *CreateActionPlanShrinkRequest {
	s.Level = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetPrologScript(v string) *CreateActionPlanShrinkRequest {
	s.PrologScript = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetRegionsShrink(v string) *CreateActionPlanShrinkRequest {
	s.RegionsShrink = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetResourceType(v string) *CreateActionPlanShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetResourcesShrink(v string) *CreateActionPlanShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) SetScript(v string) *CreateActionPlanShrinkRequest {
	s.Script = &v
	return s
}

func (s *CreateActionPlanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
