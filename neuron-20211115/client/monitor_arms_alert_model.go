// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorArmsAlert interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *MonitorArmsAlert
	GetCondition() *string
	SetEnv(v string) *MonitorArmsAlert
	GetEnv() *string
	SetGmtCreate(v string) *MonitorArmsAlert
	GetGmtCreate() *string
	SetGmtModified(v string) *MonitorArmsAlert
	GetGmtModified() *string
	SetId(v int64) *MonitorArmsAlert
	GetId() *int64
	SetMonitorArmsAlertRules(v []*MonitorArmsAlertRule) *MonitorArmsAlert
	GetMonitorArmsAlertRules() []*MonitorArmsAlertRule
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorArmsAlert
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
	SetName(v string) *MonitorArmsAlert
	GetName() *string
	SetPbcId(v string) *MonitorArmsAlert
	GetPbcId() *string
	SetRemark(v string) *MonitorArmsAlert
	GetRemark() *string
	SetServiceGroupId(v string) *MonitorArmsAlert
	GetServiceGroupId() *string
	SetType(v string) *MonitorArmsAlert
	GetType() *string
}

type MonitorArmsAlert struct {
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
	Id                    *int64                  `json:"id,omitempty" xml:"id,omitempty"`
	MonitorArmsAlertRules []*MonitorArmsAlertRule `json:"monitorArmsAlertRules,omitempty" xml:"monitorArmsAlertRules,omitempty" type:"Repeated"`
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

func (s MonitorArmsAlert) String() string {
	return dara.Prettify(s)
}

func (s MonitorArmsAlert) GoString() string {
	return s.String()
}

func (s *MonitorArmsAlert) GetCondition() *string {
	return s.Condition
}

func (s *MonitorArmsAlert) GetEnv() *string {
	return s.Env
}

func (s *MonitorArmsAlert) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MonitorArmsAlert) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MonitorArmsAlert) GetId() *int64 {
	return s.Id
}

func (s *MonitorArmsAlert) GetMonitorArmsAlertRules() []*MonitorArmsAlertRule {
	return s.MonitorArmsAlertRules
}

func (s *MonitorArmsAlert) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *MonitorArmsAlert) GetName() *string {
	return s.Name
}

func (s *MonitorArmsAlert) GetPbcId() *string {
	return s.PbcId
}

func (s *MonitorArmsAlert) GetRemark() *string {
	return s.Remark
}

func (s *MonitorArmsAlert) GetServiceGroupId() *string {
	return s.ServiceGroupId
}

func (s *MonitorArmsAlert) GetType() *string {
	return s.Type
}

func (s *MonitorArmsAlert) SetCondition(v string) *MonitorArmsAlert {
	s.Condition = &v
	return s
}

func (s *MonitorArmsAlert) SetEnv(v string) *MonitorArmsAlert {
	s.Env = &v
	return s
}

func (s *MonitorArmsAlert) SetGmtCreate(v string) *MonitorArmsAlert {
	s.GmtCreate = &v
	return s
}

func (s *MonitorArmsAlert) SetGmtModified(v string) *MonitorArmsAlert {
	s.GmtModified = &v
	return s
}

func (s *MonitorArmsAlert) SetId(v int64) *MonitorArmsAlert {
	s.Id = &v
	return s
}

func (s *MonitorArmsAlert) SetMonitorArmsAlertRules(v []*MonitorArmsAlertRule) *MonitorArmsAlert {
	s.MonitorArmsAlertRules = v
	return s
}

func (s *MonitorArmsAlert) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorArmsAlert {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *MonitorArmsAlert) SetName(v string) *MonitorArmsAlert {
	s.Name = &v
	return s
}

func (s *MonitorArmsAlert) SetPbcId(v string) *MonitorArmsAlert {
	s.PbcId = &v
	return s
}

func (s *MonitorArmsAlert) SetRemark(v string) *MonitorArmsAlert {
	s.Remark = &v
	return s
}

func (s *MonitorArmsAlert) SetServiceGroupId(v string) *MonitorArmsAlert {
	s.ServiceGroupId = &v
	return s
}

func (s *MonitorArmsAlert) SetType(v string) *MonitorArmsAlert {
	s.Type = &v
	return s
}

func (s *MonitorArmsAlert) Validate() error {
	if s.MonitorArmsAlertRules != nil {
		for _, item := range s.MonitorArmsAlertRules {
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
