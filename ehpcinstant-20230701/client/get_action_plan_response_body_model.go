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
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// example:
	//
	// TestActionPlan
	ActionPlanName *string `json:"ActionPlanName,omitempty" xml:"ActionPlanName,omitempty"`
	// example:
	//
	// Standard
	AllocationSpec *string `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	// example:
	//
	// ci-vm-rYfypJKwlN9Y
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2025-08-10 18:28:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1000
	DesiredCapacity *float32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// bHMgLWFsCmxzIC1hbGggfCB3YyAtbA==
	PrologScript *string                             `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	Regions      []*GetActionPlanResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// VCpuCapacity
	ResourceType *string                               `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Resources    []*GetActionPlanResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1000
	TotalCapacity *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
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
	return dara.Validate(s)
}

type GetActionPlanResponseBodyRegions struct {
	// example:
	//
	// cn-hangzhou
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	VSwitchIds       []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
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
	// example:
	//
	// 64
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
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
