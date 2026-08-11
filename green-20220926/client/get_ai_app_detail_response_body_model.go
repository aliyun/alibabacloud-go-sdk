// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailResponseBody
	GetAppId() *string
	SetAppName(v string) *GetAiAppDetailResponseBody
	GetAppName() *string
	SetChart(v *GetAiAppDetailResponseBodyChart) *GetAiAppDetailResponseBody
	GetChart() *GetAiAppDetailResponseBodyChart
	SetRequestId(v string) *GetAiAppDetailResponseBody
	GetRequestId() *string
	SetRiskEvents(v []*GetAiAppDetailResponseBodyRiskEvents) *GetAiAppDetailResponseBody
	GetRiskEvents() []*GetAiAppDetailResponseBodyRiskEvents
	SetScore(v int32) *GetAiAppDetailResponseBody
	GetScore() *int32
	SetUid(v string) *GetAiAppDetailResponseBody
	GetUid() *string
}

type GetAiAppDetailResponseBody struct {
	// appId。
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// app-xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The chart.
	Chart *GetAiAppDetailResponseBodyChart `json:"Chart,omitempty" xml:"Chart,omitempty" type:"Struct"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The risk events.
	RiskEvents []*GetAiAppDetailResponseBodyRiskEvents `json:"RiskEvents,omitempty" xml:"RiskEvents,omitempty" type:"Repeated"`
	// The score.
	//
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// UID。
	//
	// example:
	//
	// 17726*****370735
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetAiAppDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAiAppDetailResponseBody) GetChart() *GetAiAppDetailResponseBodyChart {
	return s.Chart
}

func (s *GetAiAppDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppDetailResponseBody) GetRiskEvents() []*GetAiAppDetailResponseBodyRiskEvents {
	return s.RiskEvents
}

func (s *GetAiAppDetailResponseBody) GetScore() *int32 {
	return s.Score
}

func (s *GetAiAppDetailResponseBody) GetUid() *string {
	return s.Uid
}

func (s *GetAiAppDetailResponseBody) SetAppId(v string) *GetAiAppDetailResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailResponseBody) SetAppName(v string) *GetAiAppDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAiAppDetailResponseBody) SetChart(v *GetAiAppDetailResponseBodyChart) *GetAiAppDetailResponseBody {
	s.Chart = v
	return s
}

func (s *GetAiAppDetailResponseBody) SetRequestId(v string) *GetAiAppDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppDetailResponseBody) SetRiskEvents(v []*GetAiAppDetailResponseBodyRiskEvents) *GetAiAppDetailResponseBody {
	s.RiskEvents = v
	return s
}

func (s *GetAiAppDetailResponseBody) SetScore(v int32) *GetAiAppDetailResponseBody {
	s.Score = &v
	return s
}

func (s *GetAiAppDetailResponseBody) SetUid(v string) *GetAiAppDetailResponseBody {
	s.Uid = &v
	return s
}

func (s *GetAiAppDetailResponseBody) Validate() error {
	if s.Chart != nil {
		if err := s.Chart.Validate(); err != nil {
			return err
		}
	}
	if s.RiskEvents != nil {
		for _, item := range s.RiskEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppDetailResponseBodyChart struct {
	// The X value of the coordinate point.
	X []*string `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	// The Y value of the coordinate point.
	Y []*GetAiAppDetailResponseBodyChartY `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s GetAiAppDetailResponseBodyChart) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailResponseBodyChart) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailResponseBodyChart) GetX() []*string {
	return s.X
}

func (s *GetAiAppDetailResponseBodyChart) GetY() []*GetAiAppDetailResponseBodyChartY {
	return s.Y
}

func (s *GetAiAppDetailResponseBodyChart) SetX(v []*string) *GetAiAppDetailResponseBodyChart {
	s.X = v
	return s
}

func (s *GetAiAppDetailResponseBodyChart) SetY(v []*GetAiAppDetailResponseBodyChartY) *GetAiAppDetailResponseBodyChart {
	s.Y = v
	return s
}

func (s *GetAiAppDetailResponseBodyChart) Validate() error {
	if s.Y != nil {
		for _, item := range s.Y {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppDetailResponseBodyChartY struct {
	// The returned collection.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// score
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAiAppDetailResponseBodyChartY) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailResponseBodyChartY) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailResponseBodyChartY) GetData() []*int64 {
	return s.Data
}

func (s *GetAiAppDetailResponseBodyChartY) GetName() *string {
	return s.Name
}

func (s *GetAiAppDetailResponseBodyChartY) SetData(v []*int64) *GetAiAppDetailResponseBodyChartY {
	s.Data = v
	return s
}

func (s *GetAiAppDetailResponseBodyChartY) SetName(v string) *GetAiAppDetailResponseBodyChartY {
	s.Name = &v
	return s
}

func (s *GetAiAppDetailResponseBodyChartY) Validate() error {
	return dara.Validate(s)
}

type GetAiAppDetailResponseBodyRiskEvents struct {
	// The risk event code.
	//
	// example:
	//
	// hit_xxx
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The number of events.
	//
	// example:
	//
	// 10
	EventCount *int64 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The list of risk event IDs.
	EventIds []*string `json:"EventIds,omitempty" xml:"EventIds,omitempty" type:"Repeated"`
	// The risk event name.
	//
	// example:
	//
	// xxx
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The event status. Valid values:
	//
	// - **unhandled**: Not handled.
	//
	// - **resolved**: Handled.
	//
	// example:
	//
	// resolved
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The labels.
	Labels []*GetAiAppDetailResponseBodyRiskEventsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s GetAiAppDetailResponseBodyRiskEvents) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailResponseBodyRiskEvents) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailResponseBodyRiskEvents) GetEventCode() *string {
	return s.EventCode
}

func (s *GetAiAppDetailResponseBodyRiskEvents) GetEventCount() *int64 {
	return s.EventCount
}

func (s *GetAiAppDetailResponseBodyRiskEvents) GetEventIds() []*string {
	return s.EventIds
}

func (s *GetAiAppDetailResponseBodyRiskEvents) GetEventName() *string {
	return s.EventName
}

func (s *GetAiAppDetailResponseBodyRiskEvents) GetEventStatus() *string {
	return s.EventStatus
}

func (s *GetAiAppDetailResponseBodyRiskEvents) GetLabels() []*GetAiAppDetailResponseBodyRiskEventsLabels {
	return s.Labels
}

func (s *GetAiAppDetailResponseBodyRiskEvents) SetEventCode(v string) *GetAiAppDetailResponseBodyRiskEvents {
	s.EventCode = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEvents) SetEventCount(v int64) *GetAiAppDetailResponseBodyRiskEvents {
	s.EventCount = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEvents) SetEventIds(v []*string) *GetAiAppDetailResponseBodyRiskEvents {
	s.EventIds = v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEvents) SetEventName(v string) *GetAiAppDetailResponseBodyRiskEvents {
	s.EventName = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEvents) SetEventStatus(v string) *GetAiAppDetailResponseBodyRiskEvents {
	s.EventStatus = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEvents) SetLabels(v []*GetAiAppDetailResponseBodyRiskEventsLabels) *GetAiAppDetailResponseBodyRiskEvents {
	s.Labels = v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEvents) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppDetailResponseBodyRiskEventsLabels struct {
	// The labels.
	//
	// example:
	//
	// inappropriate_profanity
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The label description.
	//
	// example:
	//
	// xxx
	LabelDesc *string `json:"LabelDesc,omitempty" xml:"LabelDesc,omitempty"`
	// The type.
	//
	// example:
	//
	// contentModeration
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiAppDetailResponseBodyRiskEventsLabels) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailResponseBodyRiskEventsLabels) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) GetLabel() *string {
	return s.Label
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) GetType() *string {
	return s.Type
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) SetLabel(v string) *GetAiAppDetailResponseBodyRiskEventsLabels {
	s.Label = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) SetLabelDesc(v string) *GetAiAppDetailResponseBodyRiskEventsLabels {
	s.LabelDesc = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) SetType(v string) *GetAiAppDetailResponseBodyRiskEventsLabels {
	s.Type = &v
	return s
}

func (s *GetAiAppDetailResponseBodyRiskEventsLabels) Validate() error {
	return dara.Validate(s)
}
