// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVirusScanTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListVirusScanTasksResponseBodyTasks) *ListVirusScanTasksResponseBody
	GetTasks() []*ListVirusScanTasksResponseBodyTasks
	SetTotalNum(v int32) *ListVirusScanTasksResponseBody
	GetTotalNum() *int32
}

type ListVirusScanTasksResponseBody struct {
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListVirusScanTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListVirusScanTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanTasksResponseBody) GetTasks() []*ListVirusScanTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListVirusScanTasksResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListVirusScanTasksResponseBody) SetRequestId(v string) *ListVirusScanTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanTasksResponseBody) SetTasks(v []*ListVirusScanTasksResponseBodyTasks) *ListVirusScanTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListVirusScanTasksResponseBody) SetTotalNum(v int32) *ListVirusScanTasksResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListVirusScanTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanTasksResponseBodyTasks struct {
	// example:
	//
	// 2026-08-21 10:24:31
	CreateTime       *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomMatchGroup []*ListVirusScanTasksResponseBodyTasksCustomMatchGroup `json:"CustomMatchGroup,omitempty" xml:"CustomMatchGroup,omitempty" type:"Repeated"`
	// example:
	//
	// 1786377600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Quarantine
	HighRiskOperation *string `json:"HighRiskOperation,omitempty" xml:"HighRiskOperation,omitempty"`
	// example:
	//
	// None
	LowRiskOperation *string `json:"LowRiskOperation,omitempty" xml:"LowRiskOperation,omitempty"`
	// example:
	//
	// UserGroupNormal
	MatchMode        *string                                                `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	MatchTargetInfos []*ListVirusScanTasksResponseBodyTasksMatchTargetInfos `json:"MatchTargetInfos,omitempty" xml:"MatchTargetInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	MaxCpuUsage *int64 `json:"MaxCpuUsage,omitempty" xml:"MaxCpuUsage,omitempty"`
	// example:
	//
	// Notify
	MidRiskOperation *string `json:"MidRiskOperation,omitempty" xml:"MidRiskOperation,omitempty"`
	// example:
	//
	// Balance
	PerformanceMode *string `json:"PerformanceMode,omitempty" xml:"PerformanceMode,omitempty"`
	// example:
	//
	// Quick
	ScanMode    *string   `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	ScanPath    []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	ScanTargets []*string `json:"ScanTargets,omitempty" xml:"ScanTargets,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 研发部门全盘扫描
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// example:
	//
	// v1:1024772
	TaskId    *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListVirusScanTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListVirusScanTasksResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVirusScanTasksResponseBodyTasks) GetCustomMatchGroup() []*ListVirusScanTasksResponseBodyTasksCustomMatchGroup {
	return s.CustomMatchGroup
}

func (s *ListVirusScanTasksResponseBodyTasks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListVirusScanTasksResponseBodyTasks) GetHighRiskOperation() *string {
	return s.HighRiskOperation
}

func (s *ListVirusScanTasksResponseBodyTasks) GetLowRiskOperation() *string {
	return s.LowRiskOperation
}

func (s *ListVirusScanTasksResponseBodyTasks) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListVirusScanTasksResponseBodyTasks) GetMatchTargetInfos() []*ListVirusScanTasksResponseBodyTasksMatchTargetInfos {
	return s.MatchTargetInfos
}

func (s *ListVirusScanTasksResponseBodyTasks) GetMaxCpuUsage() *int64 {
	return s.MaxCpuUsage
}

func (s *ListVirusScanTasksResponseBodyTasks) GetMidRiskOperation() *string {
	return s.MidRiskOperation
}

func (s *ListVirusScanTasksResponseBodyTasks) GetPerformanceMode() *string {
	return s.PerformanceMode
}

func (s *ListVirusScanTasksResponseBodyTasks) GetScanMode() *string {
	return s.ScanMode
}

func (s *ListVirusScanTasksResponseBodyTasks) GetScanPath() []*string {
	return s.ScanPath
}

func (s *ListVirusScanTasksResponseBodyTasks) GetScanTargets() []*string {
	return s.ScanTargets
}

func (s *ListVirusScanTasksResponseBodyTasks) GetStatus() *int32 {
	return s.Status
}

func (s *ListVirusScanTasksResponseBodyTasks) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *ListVirusScanTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVirusScanTasksResponseBodyTasks) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListVirusScanTasksResponseBodyTasks) SetCreateTime(v string) *ListVirusScanTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetCustomMatchGroup(v []*ListVirusScanTasksResponseBodyTasksCustomMatchGroup) *ListVirusScanTasksResponseBodyTasks {
	s.CustomMatchGroup = v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetEndTime(v int64) *ListVirusScanTasksResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetHighRiskOperation(v string) *ListVirusScanTasksResponseBodyTasks {
	s.HighRiskOperation = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetLowRiskOperation(v string) *ListVirusScanTasksResponseBodyTasks {
	s.LowRiskOperation = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetMatchMode(v string) *ListVirusScanTasksResponseBodyTasks {
	s.MatchMode = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetMatchTargetInfos(v []*ListVirusScanTasksResponseBodyTasksMatchTargetInfos) *ListVirusScanTasksResponseBodyTasks {
	s.MatchTargetInfos = v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetMaxCpuUsage(v int64) *ListVirusScanTasksResponseBodyTasks {
	s.MaxCpuUsage = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetMidRiskOperation(v string) *ListVirusScanTasksResponseBodyTasks {
	s.MidRiskOperation = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetPerformanceMode(v string) *ListVirusScanTasksResponseBodyTasks {
	s.PerformanceMode = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetScanMode(v string) *ListVirusScanTasksResponseBodyTasks {
	s.ScanMode = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetScanPath(v []*string) *ListVirusScanTasksResponseBodyTasks {
	s.ScanPath = v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetScanTargets(v []*string) *ListVirusScanTasksResponseBodyTasks {
	s.ScanTargets = v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetStatus(v int32) *ListVirusScanTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetTaskDescription(v string) *ListVirusScanTasksResponseBodyTasks {
	s.TaskDescription = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetTaskId(v string) *ListVirusScanTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) SetWhitelist(v []*string) *ListVirusScanTasksResponseBodyTasks {
	s.Whitelist = v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasks) Validate() error {
	if s.CustomMatchGroup != nil {
		for _, item := range s.CustomMatchGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MatchTargetInfos != nil {
		for _, item := range s.MatchTargetInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanTasksResponseBodyTasksCustomMatchGroup struct {
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
	// example:
	//
	// idp-7c3f9a2e5b18****
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s ListVirusScanTasksResponseBodyTasksCustomMatchGroup) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTasksResponseBodyTasksCustomMatchGroup) GoString() string {
	return s.String()
}

func (s *ListVirusScanTasksResponseBodyTasksCustomMatchGroup) GetGroup() []*string {
	return s.Group
}

func (s *ListVirusScanTasksResponseBodyTasksCustomMatchGroup) GetIdpId() *string {
	return s.IdpId
}

func (s *ListVirusScanTasksResponseBodyTasksCustomMatchGroup) SetGroup(v []*string) *ListVirusScanTasksResponseBodyTasksCustomMatchGroup {
	s.Group = v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasksCustomMatchGroup) SetIdpId(v string) *ListVirusScanTasksResponseBodyTasksCustomMatchGroup {
	s.IdpId = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasksCustomMatchGroup) Validate() error {
	return dara.Validate(s)
}

type ListVirusScanTasksResponseBodyTasksMatchTargetInfos struct {
	// example:
	//
	// usergroup-9d4f2a7b3c1e****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// example:
	//
	// 研发部门
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
}

func (s ListVirusScanTasksResponseBodyTasksMatchTargetInfos) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTasksResponseBodyTasksMatchTargetInfos) GoString() string {
	return s.String()
}

func (s *ListVirusScanTasksResponseBodyTasksMatchTargetInfos) GetTargetId() *string {
	return s.TargetId
}

func (s *ListVirusScanTasksResponseBodyTasksMatchTargetInfos) GetTargetName() *string {
	return s.TargetName
}

func (s *ListVirusScanTasksResponseBodyTasksMatchTargetInfos) SetTargetId(v string) *ListVirusScanTasksResponseBodyTasksMatchTargetInfos {
	s.TargetId = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasksMatchTargetInfos) SetTargetName(v string) *ListVirusScanTasksResponseBodyTasksMatchTargetInfos {
	s.TargetName = &v
	return s
}

func (s *ListVirusScanTasksResponseBodyTasksMatchTargetInfos) Validate() error {
	return dara.Validate(s)
}
