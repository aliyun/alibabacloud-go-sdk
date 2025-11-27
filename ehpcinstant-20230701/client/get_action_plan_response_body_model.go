// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActionPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanId(v string) *GetActionPlanResponseBody
	GetActionPlanId() *string
	SetActionPlanName(v string) *GetActionPlanResponseBody
	GetActionPlanName() *string
	SetAllocationSpec(v string) *GetActionPlanResponseBody
	GetAllocationSpec() *string
	SetAppId(v string) *GetActionPlanResponseBody
	GetAppId() *string
	SetCreateTime(v string) *GetActionPlanResponseBody
	GetCreateTime() *string
	SetDesiredCapacity(v float32) *GetActionPlanResponseBody
	GetDesiredCapacity() *float32
	SetLevel(v string) *GetActionPlanResponseBody
	GetLevel() *string
	SetPrologScript(v string) *GetActionPlanResponseBody
	GetPrologScript() *string
	SetRegions(v []*GetActionPlanResponseBodyRegions) *GetActionPlanResponseBody
	GetRegions() []*GetActionPlanResponseBodyRegions
	SetRequestId(v string) *GetActionPlanResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetActionPlanResponseBody
	GetResourceType() *string
	SetResources(v []*GetActionPlanResponseBodyResources) *GetActionPlanResponseBody
	GetResources() []*GetActionPlanResponseBodyResources
	SetStatus(v string) *GetActionPlanResponseBody
	GetStatus() *string
	SetTotalCapacity(v float32) *GetActionPlanResponseBody
	GetTotalCapacity() *float32
	SetUpdateTime(v string) *GetActionPlanResponseBody
	GetUpdateTime() *string
}

type GetActionPlanResponseBody struct {
	// The ID of the execution plan.
	//
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// The name of the execution plan.
	//
	// example:
	//
	// TestActionPlan
	ActionPlanName *string `json:"ActionPlanName,omitempty" xml:"ActionPlanName,omitempty"`
	// The type of the resource.
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
	// The time when the execution plan was created.
	//
	// example:
	//
	// 2025-08-10 18:28:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expected scale of resources for the execution plan. If the ResourceType parameter is set to VcpuCapacity, the execution plan is expected to have 10000 vCPUs.
	//
	// example:
	//
	// 1000
	DesiredCapacity *float32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// The computing power level.
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
	Regions []*GetActionPlanResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The list of resource configurations of the execution plan runtime environment.
	Resources []*GetActionPlanResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The status of the execution plan. The possible values are as follows:
	//
	// 	- Active Instant tasks are dynamically managed only when the execution plan is in the Active state.
	//
	// 	- Inactive Instant tasks are no longer managed by execution plans in the Inactive state.
	//
	// 	- Deleting You cannot modify the parameters of an execution plan in this state.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The size of the resources currently managed by the execution plan.
	//
	// example:
	//
	// 1000
	TotalCapacity *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The time when the execution plan was last modified.
	//
	// example:
	//
	// 2025-08-10 18:28:05
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetActionPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetActionPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetActionPlanResponseBody) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *GetActionPlanResponseBody) GetActionPlanName() *string {
	return s.ActionPlanName
}

func (s *GetActionPlanResponseBody) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *GetActionPlanResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetActionPlanResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetActionPlanResponseBody) GetDesiredCapacity() *float32 {
	return s.DesiredCapacity
}

func (s *GetActionPlanResponseBody) GetLevel() *string {
	return s.Level
}

func (s *GetActionPlanResponseBody) GetPrologScript() *string {
	return s.PrologScript
}

func (s *GetActionPlanResponseBody) GetRegions() []*GetActionPlanResponseBodyRegions {
	return s.Regions
}

func (s *GetActionPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetActionPlanResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetActionPlanResponseBody) GetResources() []*GetActionPlanResponseBodyResources {
	return s.Resources
}

func (s *GetActionPlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetActionPlanResponseBody) GetTotalCapacity() *float32 {
	return s.TotalCapacity
}

func (s *GetActionPlanResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetActionPlanResponseBody) SetActionPlanId(v string) *GetActionPlanResponseBody {
	s.ActionPlanId = &v
	return s
}

func (s *GetActionPlanResponseBody) SetActionPlanName(v string) *GetActionPlanResponseBody {
	s.ActionPlanName = &v
	return s
}

func (s *GetActionPlanResponseBody) SetAllocationSpec(v string) *GetActionPlanResponseBody {
	s.AllocationSpec = &v
	return s
}

func (s *GetActionPlanResponseBody) SetAppId(v string) *GetActionPlanResponseBody {
	s.AppId = &v
	return s
}

func (s *GetActionPlanResponseBody) SetCreateTime(v string) *GetActionPlanResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetActionPlanResponseBody) SetDesiredCapacity(v float32) *GetActionPlanResponseBody {
	s.DesiredCapacity = &v
	return s
}

func (s *GetActionPlanResponseBody) SetLevel(v string) *GetActionPlanResponseBody {
	s.Level = &v
	return s
}

func (s *GetActionPlanResponseBody) SetPrologScript(v string) *GetActionPlanResponseBody {
	s.PrologScript = &v
	return s
}

func (s *GetActionPlanResponseBody) SetRegions(v []*GetActionPlanResponseBodyRegions) *GetActionPlanResponseBody {
	s.Regions = v
	return s
}

func (s *GetActionPlanResponseBody) SetRequestId(v string) *GetActionPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetActionPlanResponseBody) SetResourceType(v string) *GetActionPlanResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetActionPlanResponseBody) SetResources(v []*GetActionPlanResponseBodyResources) *GetActionPlanResponseBody {
	s.Resources = v
	return s
}

func (s *GetActionPlanResponseBody) SetStatus(v string) *GetActionPlanResponseBody {
	s.Status = &v
	return s
}

func (s *GetActionPlanResponseBody) SetTotalCapacity(v float32) *GetActionPlanResponseBody {
	s.TotalCapacity = &v
	return s
}

func (s *GetActionPlanResponseBody) SetUpdateTime(v string) *GetActionPlanResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetActionPlanResponseBody) Validate() error {
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

type GetActionPlanResponseBodyRegions struct {
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of security groups available for the execution plan in the region.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The list of VSwitches available for the execution plan in the region.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetActionPlanResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s GetActionPlanResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *GetActionPlanResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *GetActionPlanResponseBodyRegions) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *GetActionPlanResponseBodyRegions) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetActionPlanResponseBodyRegions) SetRegionId(v string) *GetActionPlanResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *GetActionPlanResponseBodyRegions) SetSecurityGroupIds(v []*string) *GetActionPlanResponseBodyRegions {
	s.SecurityGroupIds = v
	return s
}

func (s *GetActionPlanResponseBodyRegions) SetVSwitchIds(v []*string) *GetActionPlanResponseBodyRegions {
	s.VSwitchIds = v
	return s
}

func (s *GetActionPlanResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type GetActionPlanResponseBodyResources struct {
	// The number of CPUs in the running environment.
	//
	// example:
	//
	// 64
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The memory size of the running environment. Unit: GiB.
	//
	// example:
	//
	// 128
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s GetActionPlanResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s GetActionPlanResponseBodyResources) GoString() string {
	return s.String()
}

func (s *GetActionPlanResponseBodyResources) GetCores() *float32 {
	return s.Cores
}

func (s *GetActionPlanResponseBodyResources) GetMemory() *float32 {
	return s.Memory
}

func (s *GetActionPlanResponseBodyResources) SetCores(v float32) *GetActionPlanResponseBodyResources {
	s.Cores = &v
	return s
}

func (s *GetActionPlanResponseBodyResources) SetMemory(v float32) *GetActionPlanResponseBodyResources {
	s.Memory = &v
	return s
}

func (s *GetActionPlanResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
