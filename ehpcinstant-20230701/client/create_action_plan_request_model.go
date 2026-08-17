// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActionPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanName(v string) *CreateActionPlanRequest
	GetActionPlanName() *string
	SetAllocationSpec(v string) *CreateActionPlanRequest
	GetAllocationSpec() *string
	SetAppId(v string) *CreateActionPlanRequest
	GetAppId() *string
	SetDesiredCapacity(v float64) *CreateActionPlanRequest
	GetDesiredCapacity() *float64
	SetIntervalMinutes(v int32) *CreateActionPlanRequest
	GetIntervalMinutes() *int32
	SetLevel(v string) *CreateActionPlanRequest
	GetLevel() *string
	SetPrologScript(v string) *CreateActionPlanRequest
	GetPrologScript() *string
	SetRegions(v []*CreateActionPlanRequestRegions) *CreateActionPlanRequest
	GetRegions() []*CreateActionPlanRequestRegions
	SetResourceType(v string) *CreateActionPlanRequest
	GetResourceType() *string
	SetResources(v []*CreateActionPlanRequestResources) *CreateActionPlanRequest
	GetResources() []*CreateActionPlanRequestResources
	SetScript(v string) *CreateActionPlanRequest
	GetScript() *string
}

type CreateActionPlanRequest struct {
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
	Regions []*CreateActionPlanRequestRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
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
	Resources []*CreateActionPlanRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The script to run the job. The script must be Base64-encoded.
	//
	// example:
	//
	// bHMgLWFsCmxzIC1hbGggfCB3YyAtbA==
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateActionPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateActionPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateActionPlanRequest) GetActionPlanName() *string {
	return s.ActionPlanName
}

func (s *CreateActionPlanRequest) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *CreateActionPlanRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateActionPlanRequest) GetDesiredCapacity() *float64 {
	return s.DesiredCapacity
}

func (s *CreateActionPlanRequest) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *CreateActionPlanRequest) GetLevel() *string {
	return s.Level
}

func (s *CreateActionPlanRequest) GetPrologScript() *string {
	return s.PrologScript
}

func (s *CreateActionPlanRequest) GetRegions() []*CreateActionPlanRequestRegions {
	return s.Regions
}

func (s *CreateActionPlanRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateActionPlanRequest) GetResources() []*CreateActionPlanRequestResources {
	return s.Resources
}

func (s *CreateActionPlanRequest) GetScript() *string {
	return s.Script
}

func (s *CreateActionPlanRequest) SetActionPlanName(v string) *CreateActionPlanRequest {
	s.ActionPlanName = &v
	return s
}

func (s *CreateActionPlanRequest) SetAllocationSpec(v string) *CreateActionPlanRequest {
	s.AllocationSpec = &v
	return s
}

func (s *CreateActionPlanRequest) SetAppId(v string) *CreateActionPlanRequest {
	s.AppId = &v
	return s
}

func (s *CreateActionPlanRequest) SetDesiredCapacity(v float64) *CreateActionPlanRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateActionPlanRequest) SetIntervalMinutes(v int32) *CreateActionPlanRequest {
	s.IntervalMinutes = &v
	return s
}

func (s *CreateActionPlanRequest) SetLevel(v string) *CreateActionPlanRequest {
	s.Level = &v
	return s
}

func (s *CreateActionPlanRequest) SetPrologScript(v string) *CreateActionPlanRequest {
	s.PrologScript = &v
	return s
}

func (s *CreateActionPlanRequest) SetRegions(v []*CreateActionPlanRequestRegions) *CreateActionPlanRequest {
	s.Regions = v
	return s
}

func (s *CreateActionPlanRequest) SetResourceType(v string) *CreateActionPlanRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateActionPlanRequest) SetResources(v []*CreateActionPlanRequestResources) *CreateActionPlanRequest {
	s.Resources = v
	return s
}

func (s *CreateActionPlanRequest) SetScript(v string) *CreateActionPlanRequest {
	s.Script = &v
	return s
}

func (s *CreateActionPlanRequest) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateActionPlanRequestRegions struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// A list of security groups that are available for the execution plan in the region. You can specify 0 to 5 security groups.
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// A list of security group IDs. You can call the [DescribeSecurityGroups](https://api.aliyun.com/document/Ecs/2014-05-26/DescribeSecurityGroups) operation to query information about available security groups.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// A list of vSwitches that are available for the execution plan in the region. You can specify 0 to 5 vSwitches.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateActionPlanRequestRegions) String() string {
	return dara.Prettify(s)
}

func (s CreateActionPlanRequestRegions) GoString() string {
	return s.String()
}

func (s *CreateActionPlanRequestRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateActionPlanRequestRegions) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *CreateActionPlanRequestRegions) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateActionPlanRequestRegions) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateActionPlanRequestRegions) SetRegionId(v string) *CreateActionPlanRequestRegions {
	s.RegionId = &v
	return s
}

func (s *CreateActionPlanRequestRegions) SetSecurityGroupId(v []*string) *CreateActionPlanRequestRegions {
	s.SecurityGroupId = v
	return s
}

func (s *CreateActionPlanRequestRegions) SetSecurityGroupIds(v []*string) *CreateActionPlanRequestRegions {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateActionPlanRequestRegions) SetVSwitchIds(v []*string) *CreateActionPlanRequestRegions {
	s.VSwitchIds = v
	return s
}

func (s *CreateActionPlanRequestRegions) Validate() error {
	return dara.Validate(s)
}

type CreateActionPlanRequestResources struct {
	// The number of vCPUs for the runtime environment.
	//
	// example:
	//
	// 2
	Cores *float64 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The memory size of the runtime environment. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateActionPlanRequestResources) String() string {
	return dara.Prettify(s)
}

func (s CreateActionPlanRequestResources) GoString() string {
	return s.String()
}

func (s *CreateActionPlanRequestResources) GetCores() *float64 {
	return s.Cores
}

func (s *CreateActionPlanRequestResources) GetMemory() *float64 {
	return s.Memory
}

func (s *CreateActionPlanRequestResources) SetCores(v float64) *CreateActionPlanRequestResources {
	s.Cores = &v
	return s
}

func (s *CreateActionPlanRequestResources) SetMemory(v float64) *CreateActionPlanRequestResources {
	s.Memory = &v
	return s
}

func (s *CreateActionPlanRequestResources) Validate() error {
	return dara.Validate(s)
}
