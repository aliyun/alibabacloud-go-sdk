// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEscalationStruct interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *IncidentEscalationStruct
	GetCreateTime() *int64
	SetDescription(v string) *IncidentEscalationStruct
	GetDescription() *string
	SetIncidentEscalationId(v string) *IncidentEscalationStruct
	GetIncidentEscalationId() *string
	SetModifyTime(v int64) *IncidentEscalationStruct
	GetModifyTime() *int64
	SetName(v string) *IncidentEscalationStruct
	GetName() *string
	SetRegionId(v string) *IncidentEscalationStruct
	GetRegionId() *string
	SetStage(v []*IncidentEscalationStageStruct) *IncidentEscalationStruct
	GetStage() []*IncidentEscalationStageStruct
	SetWorkspace(v string) *IncidentEscalationStruct
	GetWorkspace() *string
}

type IncidentEscalationStruct struct {
	// Creation time.
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Description.
	//
	// example:
	//
	// When an alert is not acknowledged, notify the operations team, on-duty manager, and CTO sequentially.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event Escalation ID.
	//
	// example:
	//
	// 52631388567
	IncidentEscalationId *string `json:"incidentEscalationId,omitempty" xml:"incidentEscalationId,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	ModifyTime *int64 `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// Name.
	//
	// example:
	//
	// Production Environment Alert Escalation Policy.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Phase.
	Stage []*IncidentEscalationStageStruct `json:"stage,omitempty" xml:"stage,omitempty" type:"Repeated"`
	// Workspace.
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentEscalationStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentEscalationStruct) GoString() string {
	return s.String()
}

func (s *IncidentEscalationStruct) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *IncidentEscalationStruct) GetDescription() *string {
	return s.Description
}

func (s *IncidentEscalationStruct) GetIncidentEscalationId() *string {
	return s.IncidentEscalationId
}

func (s *IncidentEscalationStruct) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *IncidentEscalationStruct) GetName() *string {
	return s.Name
}

func (s *IncidentEscalationStruct) GetRegionId() *string {
	return s.RegionId
}

func (s *IncidentEscalationStruct) GetStage() []*IncidentEscalationStageStruct {
	return s.Stage
}

func (s *IncidentEscalationStruct) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentEscalationStruct) SetCreateTime(v int64) *IncidentEscalationStruct {
	s.CreateTime = &v
	return s
}

func (s *IncidentEscalationStruct) SetDescription(v string) *IncidentEscalationStruct {
	s.Description = &v
	return s
}

func (s *IncidentEscalationStruct) SetIncidentEscalationId(v string) *IncidentEscalationStruct {
	s.IncidentEscalationId = &v
	return s
}

func (s *IncidentEscalationStruct) SetModifyTime(v int64) *IncidentEscalationStruct {
	s.ModifyTime = &v
	return s
}

func (s *IncidentEscalationStruct) SetName(v string) *IncidentEscalationStruct {
	s.Name = &v
	return s
}

func (s *IncidentEscalationStruct) SetRegionId(v string) *IncidentEscalationStruct {
	s.RegionId = &v
	return s
}

func (s *IncidentEscalationStruct) SetStage(v []*IncidentEscalationStageStruct) *IncidentEscalationStruct {
	s.Stage = v
	return s
}

func (s *IncidentEscalationStruct) SetWorkspace(v string) *IncidentEscalationStruct {
	s.Workspace = &v
	return s
}

func (s *IncidentEscalationStruct) Validate() error {
	if s.Stage != nil {
		for _, item := range s.Stage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
