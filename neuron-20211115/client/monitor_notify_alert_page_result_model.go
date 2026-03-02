// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorNotifyAlertPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetIntTotal(v int32) *MonitorNotifyAlertPageResult
	GetIntTotal() *int32
	SetNotifyAlerts(v []*MonitorNotifyAlert) *MonitorNotifyAlertPageResult
	GetNotifyAlerts() []*MonitorNotifyAlert
	SetRequestId(v string) *MonitorNotifyAlertPageResult
	GetRequestId() *string
	SetTotal(v int64) *MonitorNotifyAlertPageResult
	GetTotal() *int64
}

type MonitorNotifyAlertPageResult struct {
	IntTotal     *int32                `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	NotifyAlerts []*MonitorNotifyAlert `json:"notifyAlerts,omitempty" xml:"notifyAlerts,omitempty" type:"Repeated"`
	RequestId    *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MonitorNotifyAlertPageResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorNotifyAlertPageResult) GoString() string {
	return s.String()
}

func (s *MonitorNotifyAlertPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MonitorNotifyAlertPageResult) GetNotifyAlerts() []*MonitorNotifyAlert {
	return s.NotifyAlerts
}

func (s *MonitorNotifyAlertPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorNotifyAlertPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MonitorNotifyAlertPageResult) SetIntTotal(v int32) *MonitorNotifyAlertPageResult {
	s.IntTotal = &v
	return s
}

func (s *MonitorNotifyAlertPageResult) SetNotifyAlerts(v []*MonitorNotifyAlert) *MonitorNotifyAlertPageResult {
	s.NotifyAlerts = v
	return s
}

func (s *MonitorNotifyAlertPageResult) SetRequestId(v string) *MonitorNotifyAlertPageResult {
	s.RequestId = &v
	return s
}

func (s *MonitorNotifyAlertPageResult) SetTotal(v int64) *MonitorNotifyAlertPageResult {
	s.Total = &v
	return s
}

func (s *MonitorNotifyAlertPageResult) Validate() error {
	if s.NotifyAlerts != nil {
		for _, item := range s.NotifyAlerts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
