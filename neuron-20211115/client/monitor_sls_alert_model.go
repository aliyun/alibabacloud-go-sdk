// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorSlsAlert interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *MonitorSlsAlert
	GetCondition() *string
	SetEnv(v string) *MonitorSlsAlert
	GetEnv() *string
	SetGmtCreate(v string) *MonitorSlsAlert
	GetGmtCreate() *string
	SetGmtModified(v string) *MonitorSlsAlert
	GetGmtModified() *string
	SetId(v int64) *MonitorSlsAlert
	GetId() *int64
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorSlsAlert
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
	SetMonitorSlsAlertRules(v []*MonitorSlsAlertRule) *MonitorSlsAlert
	GetMonitorSlsAlertRules() []*MonitorSlsAlertRule
	SetName(v string) *MonitorSlsAlert
	GetName() *string
	SetPbcId(v string) *MonitorSlsAlert
	GetPbcId() *string
	SetRemark(v string) *MonitorSlsAlert
	GetRemark() *string
	SetServiceGroupId(v string) *MonitorSlsAlert
	GetServiceGroupId() *string
	SetType(v string) *MonitorSlsAlert
	GetType() *string
}

type MonitorSlsAlert struct {
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
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	MonitorNotifyStrategy *MonitorNotifyStrategy `json:"monitorNotifyStrategy,omitempty" xml:"monitorNotifyStrategy,omitempty"`
	MonitorSlsAlertRules  []*MonitorSlsAlertRule `json:"monitorSlsAlertRules,omitempty" xml:"monitorSlsAlertRules,omitempty" type:"Repeated"`
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

func (s MonitorSlsAlert) String() string {
	return dara.Prettify(s)
}

func (s MonitorSlsAlert) GoString() string {
	return s.String()
}

func (s *MonitorSlsAlert) GetCondition() *string {
	return s.Condition
}

func (s *MonitorSlsAlert) GetEnv() *string {
	return s.Env
}

func (s *MonitorSlsAlert) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MonitorSlsAlert) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MonitorSlsAlert) GetId() *int64 {
	return s.Id
}

func (s *MonitorSlsAlert) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *MonitorSlsAlert) GetMonitorSlsAlertRules() []*MonitorSlsAlertRule {
	return s.MonitorSlsAlertRules
}

func (s *MonitorSlsAlert) GetName() *string {
	return s.Name
}

func (s *MonitorSlsAlert) GetPbcId() *string {
	return s.PbcId
}

func (s *MonitorSlsAlert) GetRemark() *string {
	return s.Remark
}

func (s *MonitorSlsAlert) GetServiceGroupId() *string {
	return s.ServiceGroupId
}

func (s *MonitorSlsAlert) GetType() *string {
	return s.Type
}

func (s *MonitorSlsAlert) SetCondition(v string) *MonitorSlsAlert {
	s.Condition = &v
	return s
}

func (s *MonitorSlsAlert) SetEnv(v string) *MonitorSlsAlert {
	s.Env = &v
	return s
}

func (s *MonitorSlsAlert) SetGmtCreate(v string) *MonitorSlsAlert {
	s.GmtCreate = &v
	return s
}

func (s *MonitorSlsAlert) SetGmtModified(v string) *MonitorSlsAlert {
	s.GmtModified = &v
	return s
}

func (s *MonitorSlsAlert) SetId(v int64) *MonitorSlsAlert {
	s.Id = &v
	return s
}

func (s *MonitorSlsAlert) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorSlsAlert {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *MonitorSlsAlert) SetMonitorSlsAlertRules(v []*MonitorSlsAlertRule) *MonitorSlsAlert {
	s.MonitorSlsAlertRules = v
	return s
}

func (s *MonitorSlsAlert) SetName(v string) *MonitorSlsAlert {
	s.Name = &v
	return s
}

func (s *MonitorSlsAlert) SetPbcId(v string) *MonitorSlsAlert {
	s.PbcId = &v
	return s
}

func (s *MonitorSlsAlert) SetRemark(v string) *MonitorSlsAlert {
	s.Remark = &v
	return s
}

func (s *MonitorSlsAlert) SetServiceGroupId(v string) *MonitorSlsAlert {
	s.ServiceGroupId = &v
	return s
}

func (s *MonitorSlsAlert) SetType(v string) *MonitorSlsAlert {
	s.Type = &v
	return s
}

func (s *MonitorSlsAlert) Validate() error {
	if s.MonitorNotifyStrategy != nil {
		if err := s.MonitorNotifyStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.MonitorSlsAlertRules != nil {
		for _, item := range s.MonitorSlsAlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
