// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorMqAlerts interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorMqAlerts(v []*MonitorMqAlert) *MonitorMqAlerts
	GetMonitorMqAlerts() []*MonitorMqAlert
}

type MonitorMqAlerts struct {
	MonitorMqAlerts []*MonitorMqAlert `json:"monitorMqAlerts,omitempty" xml:"monitorMqAlerts,omitempty" type:"Repeated"`
}

func (s MonitorMqAlerts) String() string {
	return dara.Prettify(s)
}

func (s MonitorMqAlerts) GoString() string {
	return s.String()
}

func (s *MonitorMqAlerts) GetMonitorMqAlerts() []*MonitorMqAlert {
	return s.MonitorMqAlerts
}

func (s *MonitorMqAlerts) SetMonitorMqAlerts(v []*MonitorMqAlert) *MonitorMqAlerts {
	s.MonitorMqAlerts = v
	return s
}

func (s *MonitorMqAlerts) Validate() error {
	if s.MonitorMqAlerts != nil {
		for _, item := range s.MonitorMqAlerts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
