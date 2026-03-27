// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEscalationPolicyForView interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *IncidentEscalationPolicyForView
	GetCreateTime() *string
	SetDescription(v string) *IncidentEscalationPolicyForView
	GetDescription() *string
	SetEnable(v bool) *IncidentEscalationPolicyForView
	GetEnable() *bool
	SetEscalationStageList(v []*IncidentEscalationStageForView) *IncidentEscalationPolicyForView
	GetEscalationStageList() []*IncidentEscalationStageForView
	SetName(v string) *IncidentEscalationPolicyForView
	GetName() *string
	SetOwnerType(v string) *IncidentEscalationPolicyForView
	GetOwnerType() *string
	SetRegionId(v string) *IncidentEscalationPolicyForView
	GetRegionId() *string
	SetSource(v string) *IncidentEscalationPolicyForView
	GetSource() *string
	SetSyncFromType(v string) *IncidentEscalationPolicyForView
	GetSyncFromType() *string
	SetUpdateTime(v string) *IncidentEscalationPolicyForView
	GetUpdateTime() *string
	SetUserId(v string) *IncidentEscalationPolicyForView
	GetUserId() *string
	SetUuid(v string) *IncidentEscalationPolicyForView
	GetUuid() *string
	SetWorkspace(v string) *IncidentEscalationPolicyForView
	GetWorkspace() *string
}

type IncidentEscalationPolicyForView struct {
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// workspace api monitor update test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Enable              *bool                             `json:"enable,omitempty" xml:"enable,omitempty"`
	EscalationStageList []*IncidentEscalationStageForView `json:"escalationStageList,omitempty" xml:"escalationStageList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// "Critical-Alert-Escalation"
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// "USER"
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// example:
	//
	// "cn-hangzhou"
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// "MANUAL"
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// "ARMS"
	SyncFromType *string `json:"syncFromType,omitempty" xml:"syncFromType,omitempty"`
	// example:
	//
	// 2025-04-15T02:02:50Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// "user-abc123"
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// "a1b2c3d4-e5f6-7890-1234-567890abcdef"
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// example:
	//
	// "ws-xyz789"
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentEscalationPolicyForView) String() string {
	return dara.Prettify(s)
}

func (s IncidentEscalationPolicyForView) GoString() string {
	return s.String()
}

func (s *IncidentEscalationPolicyForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *IncidentEscalationPolicyForView) GetDescription() *string {
	return s.Description
}

func (s *IncidentEscalationPolicyForView) GetEnable() *bool {
	return s.Enable
}

func (s *IncidentEscalationPolicyForView) GetEscalationStageList() []*IncidentEscalationStageForView {
	return s.EscalationStageList
}

func (s *IncidentEscalationPolicyForView) GetName() *string {
	return s.Name
}

func (s *IncidentEscalationPolicyForView) GetOwnerType() *string {
	return s.OwnerType
}

func (s *IncidentEscalationPolicyForView) GetRegionId() *string {
	return s.RegionId
}

func (s *IncidentEscalationPolicyForView) GetSource() *string {
	return s.Source
}

func (s *IncidentEscalationPolicyForView) GetSyncFromType() *string {
	return s.SyncFromType
}

func (s *IncidentEscalationPolicyForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *IncidentEscalationPolicyForView) GetUserId() *string {
	return s.UserId
}

func (s *IncidentEscalationPolicyForView) GetUuid() *string {
	return s.Uuid
}

func (s *IncidentEscalationPolicyForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentEscalationPolicyForView) SetCreateTime(v string) *IncidentEscalationPolicyForView {
	s.CreateTime = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetDescription(v string) *IncidentEscalationPolicyForView {
	s.Description = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetEnable(v bool) *IncidentEscalationPolicyForView {
	s.Enable = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetEscalationStageList(v []*IncidentEscalationStageForView) *IncidentEscalationPolicyForView {
	s.EscalationStageList = v
	return s
}

func (s *IncidentEscalationPolicyForView) SetName(v string) *IncidentEscalationPolicyForView {
	s.Name = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetOwnerType(v string) *IncidentEscalationPolicyForView {
	s.OwnerType = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetRegionId(v string) *IncidentEscalationPolicyForView {
	s.RegionId = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetSource(v string) *IncidentEscalationPolicyForView {
	s.Source = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetSyncFromType(v string) *IncidentEscalationPolicyForView {
	s.SyncFromType = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetUpdateTime(v string) *IncidentEscalationPolicyForView {
	s.UpdateTime = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetUserId(v string) *IncidentEscalationPolicyForView {
	s.UserId = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetUuid(v string) *IncidentEscalationPolicyForView {
	s.Uuid = &v
	return s
}

func (s *IncidentEscalationPolicyForView) SetWorkspace(v string) *IncidentEscalationPolicyForView {
	s.Workspace = &v
	return s
}

func (s *IncidentEscalationPolicyForView) Validate() error {
	if s.EscalationStageList != nil {
		for _, item := range s.EscalationStageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
