// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficControlTargetTrafficHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTrafficControlTargetTrafficHistoryResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListTrafficControlTargetTrafficHistoryResponseBody
	GetTotalCount() *string
	SetTrafficControlTaskTrafficHistories(v []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) *ListTrafficControlTargetTrafficHistoryResponseBody
	GetTrafficControlTaskTrafficHistories() []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories
}

type ListTrafficControlTargetTrafficHistoryResponseBody struct {
	RequestId                          *string                                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                         *string                                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficControlTaskTrafficHistories []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories `json:"TrafficControlTaskTrafficHistories,omitempty" xml:"TrafficControlTaskTrafficHistories,omitempty" type:"Repeated"`
}

func (s ListTrafficControlTargetTrafficHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) GetTrafficControlTaskTrafficHistories() []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	return s.TrafficControlTaskTrafficHistories
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) SetRequestId(v string) *ListTrafficControlTargetTrafficHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) SetTotalCount(v string) *ListTrafficControlTargetTrafficHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) SetTrafficControlTaskTrafficHistories(v []*ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) *ListTrafficControlTargetTrafficHistoryResponseBody {
	s.TrafficControlTaskTrafficHistories = v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBody) Validate() error {
	if s.TrafficControlTaskTrafficHistories != nil {
		for _, item := range s.TrafficControlTaskTrafficHistories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories struct {
	ExperimentId                   *string  `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	ItemId                         *string  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	RecordTime                     *string  `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
	TrafficControlTargetAimTraffic *float64 `json:"TrafficControlTargetAimTraffic,omitempty" xml:"TrafficControlTargetAimTraffic,omitempty"`
	TrafficControlTargetTraffic    *float64 `json:"TrafficControlTargetTraffic,omitempty" xml:"TrafficControlTargetTraffic,omitempty"`
	TrafficControlTaskTraffic      *float64 `json:"TrafficControlTaskTraffic,omitempty" xml:"TrafficControlTaskTraffic,omitempty"`
}

func (s ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GetItemId() *string {
	return s.ItemId
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GetRecordTime() *string {
	return s.RecordTime
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GetTrafficControlTargetAimTraffic() *float64 {
	return s.TrafficControlTargetAimTraffic
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GetTrafficControlTargetTraffic() *float64 {
	return s.TrafficControlTargetTraffic
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) GetTrafficControlTaskTraffic() *float64 {
	return s.TrafficControlTaskTraffic
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetExperimentId(v string) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.ExperimentId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetItemId(v string) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.ItemId = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetRecordTime(v string) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.RecordTime = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetTrafficControlTargetAimTraffic(v float64) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.TrafficControlTargetAimTraffic = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetTrafficControlTargetTraffic(v float64) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.TrafficControlTargetTraffic = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) SetTrafficControlTaskTraffic(v float64) *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories {
	s.TrafficControlTaskTraffic = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponseBodyTrafficControlTaskTrafficHistories) Validate() error {
	return dara.Validate(s)
}
