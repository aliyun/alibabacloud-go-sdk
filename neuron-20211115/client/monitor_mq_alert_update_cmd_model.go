// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMqAlertUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *MonitorMqAlertUpdateCmd
	GetCondition() *string
	SetId(v int64) *MonitorMqAlertUpdateCmd
	GetId() *int64
	SetMonitorMqAlertRules(v []*MonitorMqAlertRule) *MonitorMqAlertUpdateCmd
	GetMonitorMqAlertRules() []*MonitorMqAlertRule
	SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorMqAlertUpdateCmd
	GetMonitorNotifyStrategy() *MonitorNotifyStrategy
}

type MonitorMqAlertUpdateCmd struct {
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
	MonitorMqAlertRules   []*MonitorMqAlertRule  `json:"monitorMqAlertRules,omitempty" xml:"monitorMqAlertRules,omitempty" type:"Repeated"`
	MonitorNotifyStrategy *MonitorNotifyStrategy `json:"monitorNotifyStrategy,omitempty" xml:"monitorNotifyStrategy,omitempty"`
}

func (s MonitorMqAlertUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorMqAlertUpdateCmd) GoString() string {
	return s.String()
}

func (s *MonitorMqAlertUpdateCmd) GetCondition() *string {
	return s.Condition
}

func (s *MonitorMqAlertUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *MonitorMqAlertUpdateCmd) GetMonitorMqAlertRules() []*MonitorMqAlertRule {
	return s.MonitorMqAlertRules
}

func (s *MonitorMqAlertUpdateCmd) GetMonitorNotifyStrategy() *MonitorNotifyStrategy {
	return s.MonitorNotifyStrategy
}

func (s *MonitorMqAlertUpdateCmd) SetCondition(v string) *MonitorMqAlertUpdateCmd {
	s.Condition = &v
	return s
}

func (s *MonitorMqAlertUpdateCmd) SetId(v int64) *MonitorMqAlertUpdateCmd {
	s.Id = &v
	return s
}

func (s *MonitorMqAlertUpdateCmd) SetMonitorMqAlertRules(v []*MonitorMqAlertRule) *MonitorMqAlertUpdateCmd {
	s.MonitorMqAlertRules = v
	return s
}

func (s *MonitorMqAlertUpdateCmd) SetMonitorNotifyStrategy(v *MonitorNotifyStrategy) *MonitorMqAlertUpdateCmd {
	s.MonitorNotifyStrategy = v
	return s
}

func (s *MonitorMqAlertUpdateCmd) Validate() error {
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
