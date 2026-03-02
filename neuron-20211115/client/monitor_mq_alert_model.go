// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMqAlert interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *MonitorMqAlert
	GetCondition() *string
	SetEnv(v string) *MonitorMqAlert
	GetEnv() *string
	SetGmtCreate(v string) *MonitorMqAlert
	GetGmtCreate() *string
	SetGmtModified(v string) *MonitorMqAlert
	GetGmtModified() *string
	SetId(v int64) *MonitorMqAlert
	GetId() *int64
	SetMonitorMqAlertRules(v []*MonitorMqAlertRule) *MonitorMqAlert
	GetMonitorMqAlertRules() []*MonitorMqAlertRule
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorMqAlert
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
	SetName(v string) *MonitorMqAlert
	GetName() *string
	SetPbcId(v string) *MonitorMqAlert
	GetPbcId() *string
	SetRemark(v string) *MonitorMqAlert
	GetRemark() *string
	SetServiceGroupId(v string) *MonitorMqAlert
	GetServiceGroupId() *string
	SetType(v string) *MonitorMqAlert
	GetType() *string
}

type MonitorMqAlert struct {
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 2022-08-24T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-08-24T00:00:00.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id                  *int64                `json:"id,omitempty" xml:"id,omitempty"`
	MonitorMqAlertRules []*MonitorMqAlertRule `json:"monitorMqAlertRules,omitempty" xml:"monitorMqAlertRules,omitempty" type:"Repeated"`
	// This parameter is required.
	MonitorNotifyStrategy *MonitorNotifyStrategy `json:"monitorNotifyStrategy,omitempty" xml:"monitorNotifyStrategy,omitempty"`
	// example:
	//
	// 报警任务1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1001
	PbcId *string `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// 测试任务
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1001
	ServiceGroupId *string `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	// example:
	//
	// ARMS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MonitorMqAlert) String() string {
	return dara.Prettify(s)
}

func (s MonitorMqAlert) GoString() string {
	return s.String()
}

func (s *MonitorMqAlert) GetCondition() *string {
	return s.Condition
}

func (s *MonitorMqAlert) GetEnv() *string {
	return s.Env
}

func (s *MonitorMqAlert) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MonitorMqAlert) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MonitorMqAlert) GetId() *int64 {
	return s.Id
}

func (s *MonitorMqAlert) GetMonitorMqAlertRules() []*MonitorMqAlertRule {
	return s.MonitorMqAlertRules
}

func (s *MonitorMqAlert) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *MonitorMqAlert) GetName() *string {
	return s.Name
}

func (s *MonitorMqAlert) GetPbcId() *string {
	return s.PbcId
}

func (s *MonitorMqAlert) GetRemark() *string {
	return s.Remark
}

func (s *MonitorMqAlert) GetServiceGroupId() *string {
	return s.ServiceGroupId
}

func (s *MonitorMqAlert) GetType() *string {
	return s.Type
}

func (s *MonitorMqAlert) SetCondition(v string) *MonitorMqAlert {
	s.Condition = &v
	return s
}

func (s *MonitorMqAlert) SetEnv(v string) *MonitorMqAlert {
	s.Env = &v
	return s
}

func (s *MonitorMqAlert) SetGmtCreate(v string) *MonitorMqAlert {
	s.GmtCreate = &v
	return s
}

func (s *MonitorMqAlert) SetGmtModified(v string) *MonitorMqAlert {
	s.GmtModified = &v
	return s
}

func (s *MonitorMqAlert) SetId(v int64) *MonitorMqAlert {
	s.Id = &v
	return s
}

func (s *MonitorMqAlert) SetMonitorMqAlertRules(v []*MonitorMqAlertRule) *MonitorMqAlert {
	s.MonitorMqAlertRules = v
	return s
}

func (s *MonitorMqAlert) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorMqAlert {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *MonitorMqAlert) SetName(v string) *MonitorMqAlert {
	s.Name = &v
	return s
}

func (s *MonitorMqAlert) SetPbcId(v string) *MonitorMqAlert {
	s.PbcId = &v
	return s
}

func (s *MonitorMqAlert) SetRemark(v string) *MonitorMqAlert {
	s.Remark = &v
	return s
}

func (s *MonitorMqAlert) SetServiceGroupId(v string) *MonitorMqAlert {
	s.ServiceGroupId = &v
	return s
}

func (s *MonitorMqAlert) SetType(v string) *MonitorMqAlert {
	s.Type = &v
	return s
}

func (s *MonitorMqAlert) Validate() error {
	if s.MonitorMqAlertRules != nil {
		for _, item := range s.MonitorMqAlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MonitorNotifyStrategy != nil {
		if err := s.MonitorNotifyStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
