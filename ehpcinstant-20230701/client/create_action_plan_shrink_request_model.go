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
	// The resource type.
	//
	// - Standard: Standard.
	//
	// - Dedicated: Dedicated. This type is available only to users in the whitelist.
	//
	// - Economic: Economy. This type is available only to users in the whitelist.
	//
	// example:
	//
	// Standard
	AllocationSpec *string `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	// The application ID.
	//
	// example:
	//
	// ci-vm-rYfypJKwlN9Y
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The desired size of the resource for the execution plan. For example, if you set ResourceType to VcpuCapacity, this parameter specifies the number of vCPUs that you want to maintain for the execution plan.
	//
	// example:
	//
	// 1000
	DesiredCapacity *float64 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// example:
	//
	// 60
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
	// The computing power level. This parameter is valid only when you set AllocationSpec to Economic. The following types are supported:
	//
	// - General: General-purpose.
	//
	// - Performance: Compute-optimized.
	//
	// Default value: General
	//
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The pre-execution script. The script must be Base64-encoded.
	//
	// example:
	//
	// bHMgLWFsCmxzIC1hbGggfCB3YyAtbA==
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// A list of regional resource configurations for the runtime environment of the execution plan.
	RegionsShrink *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The type of resource for the execution target. The value can be the vCPU capacity or the number of executor nodes. Valid values:
	//
	// - VCpuCapacity: vCPU capacity
	//
	// - ExecutorCapacity: number of executor nodes
	//
	// example:
	//
	// VCpuCapacity
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// A list of resource configurations for the runtime environment of the execution plan. You can specify 1 to 10 resource configurations.
	//
	// example:
	//
	// 1000
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The script to run the job. The script must be Base64-encoded.
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
