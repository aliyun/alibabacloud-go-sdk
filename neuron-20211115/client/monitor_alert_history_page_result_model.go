// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorAlertHistoryPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetAlertHistories(v []*MonitorAlertHistory) *MonitorAlertHistoryPageResult
	GetAlertHistories() []*MonitorAlertHistory
	SetIntTotal(v int32) *MonitorAlertHistoryPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MonitorAlertHistoryPageResult
	GetRequestId() *string
	SetTotal(v int64) *MonitorAlertHistoryPageResult
	GetTotal() *int64
}

type MonitorAlertHistoryPageResult struct {
	AlertHistories []*MonitorAlertHistory `json:"alertHistories,omitempty" xml:"alertHistories,omitempty" type:"Repeated"`
	IntTotal       *int32                 `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId      *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MonitorAlertHistoryPageResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorAlertHistoryPageResult) GoString() string {
	return s.String()
}

func (s *MonitorAlertHistoryPageResult) GetAlertHistories() []*MonitorAlertHistory {
	return s.AlertHistories
}

func (s *MonitorAlertHistoryPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MonitorAlertHistoryPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorAlertHistoryPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MonitorAlertHistoryPageResult) SetAlertHistories(v []*MonitorAlertHistory) *MonitorAlertHistoryPageResult {
	s.AlertHistories = v
	return s
}

func (s *MonitorAlertHistoryPageResult) SetIntTotal(v int32) *MonitorAlertHistoryPageResult {
	s.IntTotal = &v
	return s
}

func (s *MonitorAlertHistoryPageResult) SetRequestId(v string) *MonitorAlertHistoryPageResult {
	s.RequestId = &v
	return s
}

func (s *MonitorAlertHistoryPageResult) SetTotal(v int64) *MonitorAlertHistoryPageResult {
	s.Total = &v
	return s
}

func (s *MonitorAlertHistoryPageResult) Validate() error {
	if s.AlertHistories != nil {
		for _, item := range s.AlertHistories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
