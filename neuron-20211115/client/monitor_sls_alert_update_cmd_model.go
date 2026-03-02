// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorSlsAlertUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *MonitorSlsAlertUpdateCmd
	GetCondition() *string
	SetId(v int64) *MonitorSlsAlertUpdateCmd
	GetId() *int64
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorSlsAlertUpdateCmd
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
	SetMonitorSlsAlertRules(v []*MonitorSlsAlertRule) *MonitorSlsAlertUpdateCmd
	GetMonitorSlsAlertRules() []*MonitorSlsAlertRule
}

type MonitorSlsAlertUpdateCmd struct {
	// example:
	//
	// OR
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id                    *int64                 `json:"id,omitempty" xml:"id,omitempty"`
	MonitorNotifyStrategy *MonitorNotifyStrategy `json:"monitorNotifyStrategy,omitempty" xml:"monitorNotifyStrategy,omitempty"`
	MonitorSlsAlertRules  []*MonitorSlsAlertRule `json:"monitorSlsAlertRules,omitempty" xml:"monitorSlsAlertRules,omitempty" type:"Repeated"`
}

func (s MonitorSlsAlertUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorSlsAlertUpdateCmd) GoString() string {
	return s.String()
}

func (s *MonitorSlsAlertUpdateCmd) GetCondition() *string {
	return s.Condition
}

func (s *MonitorSlsAlertUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *MonitorSlsAlertUpdateCmd) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *MonitorSlsAlertUpdateCmd) GetMonitorSlsAlertRules() []*MonitorSlsAlertRule {
	return s.MonitorSlsAlertRules
}

func (s *MonitorSlsAlertUpdateCmd) SetCondition(v string) *MonitorSlsAlertUpdateCmd {
	s.Condition = &v
	return s
}

func (s *MonitorSlsAlertUpdateCmd) SetId(v int64) *MonitorSlsAlertUpdateCmd {
	s.Id = &v
	return s
}

func (s *MonitorSlsAlertUpdateCmd) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorSlsAlertUpdateCmd {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *MonitorSlsAlertUpdateCmd) SetMonitorSlsAlertRules(v []*MonitorSlsAlertRule) *MonitorSlsAlertUpdateCmd {
	s.MonitorSlsAlertRules = v
	return s
}

func (s *MonitorSlsAlertUpdateCmd) Validate() error {
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
