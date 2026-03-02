// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorArmsAlerts interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorArmsAlerts(v []*MonitorArmsAlert) *MonitorArmsAlerts
	GetMonitorArmsAlerts() []*MonitorArmsAlert
}

type MonitorArmsAlerts struct {
	MonitorArmsAlerts []*MonitorArmsAlert `json:"monitorArmsAlerts,omitempty" xml:"monitorArmsAlerts,omitempty" type:"Repeated"`
}

func (s MonitorArmsAlerts) String() string {
	return dara.Prettify(s)
}

func (s MonitorArmsAlerts) GoString() string {
	return s.String()
}

func (s *MonitorArmsAlerts) GetMonitorArmsAlerts() []*MonitorArmsAlert {
	return s.MonitorArmsAlerts
}

func (s *MonitorArmsAlerts) SetMonitorArmsAlerts(v []*MonitorArmsAlert) *MonitorArmsAlerts {
	s.MonitorArmsAlerts = v
	return s
}

func (s *MonitorArmsAlerts) Validate() error {
	if s.MonitorArmsAlerts != nil {
		for _, item := range s.MonitorArmsAlerts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
