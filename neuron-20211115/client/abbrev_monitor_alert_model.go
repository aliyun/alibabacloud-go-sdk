// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbbrevMonitorAlert interface {
	dara.Model
	String() string
	GoString() string
	SetAbbrevMonitorAlertRules(v []*AbbrevMonitorAlertRule) *AbbrevMonitorAlert
	GetAbbrevMonitorAlertRules() []*AbbrevMonitorAlertRule
	SetCondition(v string) *AbbrevMonitorAlert
	GetCondition() *string
	SetEnv(v string) *AbbrevMonitorAlert
	GetEnv() *string
	SetGmtCreate(v string) *AbbrevMonitorAlert
	GetGmtCreate() *string
	SetGmtModified(v string) *AbbrevMonitorAlert
	GetGmtModified() *string
	SetId(v int64) *AbbrevMonitorAlert
	GetId() *int64
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *AbbrevMonitorAlert
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
	SetName(v string) *AbbrevMonitorAlert
	GetName() *string
	SetPbcId(v string) *AbbrevMonitorAlert
	GetPbcId() *string
	SetRemark(v string) *AbbrevMonitorAlert
	GetRemark() *string
	SetServiceGroupId(v string) *AbbrevMonitorAlert
	GetServiceGroupId() *string
	SetType(v string) *AbbrevMonitorAlert
	GetType() *string
}

type AbbrevMonitorAlert struct {
	AbbrevMonitorAlertRules []*AbbrevMonitorAlertRule `json:"abbrevMonitorAlertRules,omitempty" xml:"abbrevMonitorAlertRules,omitempty" type:"Repeated"`
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

func (s AbbrevMonitorAlert) String() string {
	return dara.Prettify(s)
}

func (s AbbrevMonitorAlert) GoString() string {
	return s.String()
}

func (s *AbbrevMonitorAlert) GetAbbrevMonitorAlertRules() []*AbbrevMonitorAlertRule {
	return s.AbbrevMonitorAlertRules
}

func (s *AbbrevMonitorAlert) GetCondition() *string {
	return s.Condition
}

func (s *AbbrevMonitorAlert) GetEnv() *string {
	return s.Env
}

func (s *AbbrevMonitorAlert) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AbbrevMonitorAlert) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AbbrevMonitorAlert) GetId() *int64 {
	return s.Id
}

func (s *AbbrevMonitorAlert) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *AbbrevMonitorAlert) GetName() *string {
	return s.Name
}

func (s *AbbrevMonitorAlert) GetPbcId() *string {
	return s.PbcId
}

func (s *AbbrevMonitorAlert) GetRemark() *string {
	return s.Remark
}

func (s *AbbrevMonitorAlert) GetServiceGroupId() *string {
	return s.ServiceGroupId
}

func (s *AbbrevMonitorAlert) GetType() *string {
	return s.Type
}

func (s *AbbrevMonitorAlert) SetAbbrevMonitorAlertRules(v []*AbbrevMonitorAlertRule) *AbbrevMonitorAlert {
	s.AbbrevMonitorAlertRules = v
	return s
}

func (s *AbbrevMonitorAlert) SetCondition(v string) *AbbrevMonitorAlert {
	s.Condition = &v
	return s
}

func (s *AbbrevMonitorAlert) SetEnv(v string) *AbbrevMonitorAlert {
	s.Env = &v
	return s
}

func (s *AbbrevMonitorAlert) SetGmtCreate(v string) *AbbrevMonitorAlert {
	s.GmtCreate = &v
	return s
}

func (s *AbbrevMonitorAlert) SetGmtModified(v string) *AbbrevMonitorAlert {
	s.GmtModified = &v
	return s
}

func (s *AbbrevMonitorAlert) SetId(v int64) *AbbrevMonitorAlert {
	s.Id = &v
	return s
}

func (s *AbbrevMonitorAlert) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *AbbrevMonitorAlert {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *AbbrevMonitorAlert) SetName(v string) *AbbrevMonitorAlert {
	s.Name = &v
	return s
}

func (s *AbbrevMonitorAlert) SetPbcId(v string) *AbbrevMonitorAlert {
	s.PbcId = &v
	return s
}

func (s *AbbrevMonitorAlert) SetRemark(v string) *AbbrevMonitorAlert {
	s.Remark = &v
	return s
}

func (s *AbbrevMonitorAlert) SetServiceGroupId(v string) *AbbrevMonitorAlert {
	s.ServiceGroupId = &v
	return s
}

func (s *AbbrevMonitorAlert) SetType(v string) *AbbrevMonitorAlert {
	s.Type = &v
	return s
}

func (s *AbbrevMonitorAlert) Validate() error {
	if s.AbbrevMonitorAlertRules != nil {
		for _, item := range s.AbbrevMonitorAlertRules {
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
