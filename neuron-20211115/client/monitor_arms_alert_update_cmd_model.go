// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorArmsAlertUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *MonitorArmsAlertUpdateCmd
	GetCondition() *string
	SetId(v int64) *MonitorArmsAlertUpdateCmd
	GetId() *int64
	SetMonitorArmsAlertRules(v []*MonitorArmsAlertRule) *MonitorArmsAlertUpdateCmd
	GetMonitorArmsAlertRules() []*MonitorArmsAlertRule
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorArmsAlertUpdateCmd
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
}

type MonitorArmsAlertUpdateCmd struct {
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id                    *int64                  `json:"id,omitempty" xml:"id,omitempty"`
	MonitorArmsAlertRules []*MonitorArmsAlertRule `json:"monitorArmsAlertRules,omitempty" xml:"monitorArmsAlertRules,omitempty" type:"Repeated"`
	MonitorNotifyStrategy *MonitorNotifyStrategy  `json:"monitorNotifyStrategy,omitempty" xml:"monitorNotifyStrategy,omitempty"`
}

func (s MonitorArmsAlertUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorArmsAlertUpdateCmd) GoString() string {
	return s.String()
}

func (s *MonitorArmsAlertUpdateCmd) GetCondition() *string {
	return s.Condition
}

func (s *MonitorArmsAlertUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *MonitorArmsAlertUpdateCmd) GetMonitorArmsAlertRules() []*MonitorArmsAlertRule {
	return s.MonitorArmsAlertRules
}

func (s *MonitorArmsAlertUpdateCmd) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *MonitorArmsAlertUpdateCmd) SetCondition(v string) *MonitorArmsAlertUpdateCmd {
	s.Condition = &v
	return s
}

func (s *MonitorArmsAlertUpdateCmd) SetId(v int64) *MonitorArmsAlertUpdateCmd {
	s.Id = &v
	return s
}

func (s *MonitorArmsAlertUpdateCmd) SetMonitorArmsAlertRules(v []*MonitorArmsAlertRule) *MonitorArmsAlertUpdateCmd {
	s.MonitorArmsAlertRules = v
	return s
}

func (s *MonitorArmsAlertUpdateCmd) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorArmsAlertUpdateCmd {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *MonitorArmsAlertUpdateCmd) Validate() error {
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
