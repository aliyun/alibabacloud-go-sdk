// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorSlsAlerts interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorSlsAlerts(v []*MonitorSlsAlert) *MonitorSlsAlerts
	GetMonitorSlsAlerts() []*MonitorSlsAlert
}

type MonitorSlsAlerts struct {
	MonitorSlsAlerts []*MonitorSlsAlert `json:"monitorSlsAlerts,omitempty" xml:"monitorSlsAlerts,omitempty" type:"Repeated"`
}

func (s MonitorSlsAlerts) String() string {
	return dara.Prettify(s)
}

func (s MonitorSlsAlerts) GoString() string {
	return s.String()
}

func (s *MonitorSlsAlerts) GetMonitorSlsAlerts() []*MonitorSlsAlert {
	return s.MonitorSlsAlerts
}

func (s *MonitorSlsAlerts) SetMonitorSlsAlerts(v []*MonitorSlsAlert) *MonitorSlsAlerts {
	s.MonitorSlsAlerts = v
	return s
}

func (s *MonitorSlsAlerts) Validate() error {
	if s.MonitorSlsAlerts != nil {
		for _, item := range s.MonitorSlsAlerts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
