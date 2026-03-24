// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEscalationStageStruct interface {
	dara.Model
	String() string
	GoString() string
	SetContact(v []*IncidentContactStruct) *IncidentEscalationStageStruct
	GetContact() []*IncidentContactStruct
	SetCycleNotifyCount(v int32) *IncidentEscalationStageStruct
	GetCycleNotifyCount() *int32
	SetCycleNotifyTime(v int32) *IncidentEscalationStageStruct
	GetCycleNotifyTime() *int32
	SetDescription(v string) *IncidentEscalationStageStruct
	GetDescription() *string
	SetEffectTime(v string) *IncidentEscalationStageStruct
	GetEffectTime() *string
	SetName(v string) *IncidentEscalationStageStruct
	GetName() *string
	SetStageIndex(v int32) *IncidentEscalationStageStruct
	GetStageIndex() *int32
	SetTimeZone(v string) *IncidentEscalationStageStruct
	GetTimeZone() *string
	SetWaitToNextStageTime(v int32) *IncidentEscalationStageStruct
	GetWaitToNextStageTime() *int32
}

type IncidentEscalationStageStruct struct {
	// 联系人列表
	Contact []*IncidentContactStruct `json:"contact,omitempty" xml:"contact,omitempty" type:"Repeated"`
	// 周期通知统计次数
	//
	// example:
	//
	// 3
	CycleNotifyCount *int32 `json:"cycleNotifyCount,omitempty" xml:"cycleNotifyCount,omitempty"`
	// 周期通知时间
	//
	// example:
	//
	// 300
	CycleNotifyTime *int32 `json:"cycleNotifyTime,omitempty" xml:"cycleNotifyTime,omitempty"`
	// 描述
	//
	// example:
	//
	// 负责处理生产环境告警的主负责人
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 有效时间
	//
	// example:
	//
	// 2025-04-05T00:00:00Z
	EffectTime *string `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	// 名称
	//
	// example:
	//
	// 运维团队
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 阶段指数
	//
	// example:
	//
	// 1
	StageIndex *int32 `json:"stageIndex,omitempty" xml:"stageIndex,omitempty"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// 等待进入下一阶段的时间
	//
	// example:
	//
	// 600
	WaitToNextStageTime *int32 `json:"waitToNextStageTime,omitempty" xml:"waitToNextStageTime,omitempty"`
}

func (s IncidentEscalationStageStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentEscalationStageStruct) GoString() string {
	return s.String()
}

func (s *IncidentEscalationStageStruct) GetContact() []*IncidentContactStruct {
	return s.Contact
}

func (s *IncidentEscalationStageStruct) GetCycleNotifyCount() *int32 {
	return s.CycleNotifyCount
}

func (s *IncidentEscalationStageStruct) GetCycleNotifyTime() *int32 {
	return s.CycleNotifyTime
}

func (s *IncidentEscalationStageStruct) GetDescription() *string {
	return s.Description
}

func (s *IncidentEscalationStageStruct) GetEffectTime() *string {
	return s.EffectTime
}

func (s *IncidentEscalationStageStruct) GetName() *string {
	return s.Name
}

func (s *IncidentEscalationStageStruct) GetStageIndex() *int32 {
	return s.StageIndex
}

func (s *IncidentEscalationStageStruct) GetTimeZone() *string {
	return s.TimeZone
}

func (s *IncidentEscalationStageStruct) GetWaitToNextStageTime() *int32 {
	return s.WaitToNextStageTime
}

func (s *IncidentEscalationStageStruct) SetContact(v []*IncidentContactStruct) *IncidentEscalationStageStruct {
	s.Contact = v
	return s
}

func (s *IncidentEscalationStageStruct) SetCycleNotifyCount(v int32) *IncidentEscalationStageStruct {
	s.CycleNotifyCount = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetCycleNotifyTime(v int32) *IncidentEscalationStageStruct {
	s.CycleNotifyTime = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetDescription(v string) *IncidentEscalationStageStruct {
	s.Description = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetEffectTime(v string) *IncidentEscalationStageStruct {
	s.EffectTime = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetName(v string) *IncidentEscalationStageStruct {
	s.Name = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetStageIndex(v int32) *IncidentEscalationStageStruct {
	s.StageIndex = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetTimeZone(v string) *IncidentEscalationStageStruct {
	s.TimeZone = &v
	return s
}

func (s *IncidentEscalationStageStruct) SetWaitToNextStageTime(v int32) *IncidentEscalationStageStruct {
	s.WaitToNextStageTime = &v
	return s
}

func (s *IncidentEscalationStageStruct) Validate() error {
	if s.Contact != nil {
		for _, item := range s.Contact {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
