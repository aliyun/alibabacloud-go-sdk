// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAiAppByPageResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListAiAppByPageResponseBodyItems) *ListAiAppByPageResponseBody
	GetItems() []*ListAiAppByPageResponseBodyItems
	SetPageSize(v int32) *ListAiAppByPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAiAppByPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAiAppByPageResponseBody
	GetTotalCount() *int64
}

type ListAiAppByPageResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data on the current page.
	Items []*ListAiAppByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID assigned by the backend to uniquely identify the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAiAppByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiAppByPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAiAppByPageResponseBody) GetItems() []*ListAiAppByPageResponseBodyItems {
	return s.Items
}

func (s *ListAiAppByPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiAppByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiAppByPageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAiAppByPageResponseBody) SetCurrentPage(v int32) *ListAiAppByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListAiAppByPageResponseBody) SetItems(v []*ListAiAppByPageResponseBodyItems) *ListAiAppByPageResponseBody {
	s.Items = v
	return s
}

func (s *ListAiAppByPageResponseBody) SetPageSize(v int32) *ListAiAppByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAiAppByPageResponseBody) SetRequestId(v string) *ListAiAppByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiAppByPageResponseBody) SetTotalCount(v int64) *ListAiAppByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAiAppByPageResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAiAppByPageResponseBodyItems struct {
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
	// name-xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application status.
	//
	// example:
	//
	// online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// The channel.
	//
	// example:
	//
	// bailian
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The last active time.
	//
	// example:
	//
	// 2026-01-01 00:00:00
	LastTraceTime *string `json:"LastTraceTime,omitempty" xml:"LastTraceTime,omitempty"`
	// The risk events.
	RiskEvents []*ListAiAppByPageResponseBodyItemsRiskEvents `json:"RiskEvents,omitempty" xml:"RiskEvents,omitempty" type:"Repeated"`
	// The risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The Tracing Analysis status.
	//
	// example:
	//
	// enable
	TraceStatus *string `json:"TraceStatus,omitempty" xml:"TraceStatus,omitempty"`
	// UID。
	//
	// example:
	//
	// 104813*****2399
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The number of alerts.
	//
	// example:
	//
	// 10
	WarningCount *int32 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s ListAiAppByPageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAiAppByPageResponseBodyItems) GetAppId() *string {
	return s.AppId
}

func (s *ListAiAppByPageResponseBodyItems) GetAppName() *string {
	return s.AppName
}

func (s *ListAiAppByPageResponseBodyItems) GetAppStatus() *string {
	return s.AppStatus
}

func (s *ListAiAppByPageResponseBodyItems) GetChannel() *string {
	return s.Channel
}

func (s *ListAiAppByPageResponseBodyItems) GetLastTraceTime() *string {
	return s.LastTraceTime
}

func (s *ListAiAppByPageResponseBodyItems) GetRiskEvents() []*ListAiAppByPageResponseBodyItemsRiskEvents {
	return s.RiskEvents
}

func (s *ListAiAppByPageResponseBodyItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListAiAppByPageResponseBodyItems) GetTraceStatus() *string {
	return s.TraceStatus
}

func (s *ListAiAppByPageResponseBodyItems) GetUid() *string {
	return s.Uid
}

func (s *ListAiAppByPageResponseBodyItems) GetWarningCount() *int32 {
	return s.WarningCount
}

func (s *ListAiAppByPageResponseBodyItems) SetAppId(v string) *ListAiAppByPageResponseBodyItems {
	s.AppId = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetAppName(v string) *ListAiAppByPageResponseBodyItems {
	s.AppName = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetAppStatus(v string) *ListAiAppByPageResponseBodyItems {
	s.AppStatus = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetChannel(v string) *ListAiAppByPageResponseBodyItems {
	s.Channel = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetLastTraceTime(v string) *ListAiAppByPageResponseBodyItems {
	s.LastTraceTime = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetRiskEvents(v []*ListAiAppByPageResponseBodyItemsRiskEvents) *ListAiAppByPageResponseBodyItems {
	s.RiskEvents = v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetRiskLevel(v string) *ListAiAppByPageResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetTraceStatus(v string) *ListAiAppByPageResponseBodyItems {
	s.TraceStatus = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetUid(v string) *ListAiAppByPageResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) SetWarningCount(v int32) *ListAiAppByPageResponseBodyItems {
	s.WarningCount = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItems) Validate() error {
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

type ListAiAppByPageResponseBodyItemsRiskEvents struct {
	// The risk event code.
	//
	// example:
	//
	// hit-xxx
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The number of events.
	//
	// example:
	//
	// 10
	EventCount *int64 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The event descriptions.
	EventDescs []*string `json:"EventDescs,omitempty" xml:"EventDescs,omitempty" type:"Repeated"`
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
	// The list of label items.
	Labels []*ListAiAppByPageResponseBodyItemsRiskEventsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s ListAiAppByPageResponseBodyItemsRiskEvents) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppByPageResponseBodyItemsRiskEvents) GoString() string {
	return s.String()
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetEventCode() *string {
	return s.EventCode
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetEventCount() *int64 {
	return s.EventCount
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetEventDescs() []*string {
	return s.EventDescs
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetEventIds() []*string {
	return s.EventIds
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetEventName() *string {
	return s.EventName
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetEventStatus() *string {
	return s.EventStatus
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) GetLabels() []*ListAiAppByPageResponseBodyItemsRiskEventsLabels {
	return s.Labels
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetEventCode(v string) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.EventCode = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetEventCount(v int64) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.EventCount = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetEventDescs(v []*string) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.EventDescs = v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetEventIds(v []*string) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.EventIds = v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetEventName(v string) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.EventName = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetEventStatus(v string) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.EventStatus = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) SetLabels(v []*ListAiAppByPageResponseBodyItemsRiskEventsLabels) *ListAiAppByPageResponseBodyItemsRiskEvents {
	s.Labels = v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEvents) Validate() error {
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

type ListAiAppByPageResponseBodyItemsRiskEventsLabels struct {
	// The label name.
	//
	// example:
	//
	// porn
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
	// sensitiveData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAiAppByPageResponseBodyItemsRiskEventsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppByPageResponseBodyItemsRiskEventsLabels) GoString() string {
	return s.String()
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) GetLabel() *string {
	return s.Label
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) GetType() *string {
	return s.Type
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) SetLabel(v string) *ListAiAppByPageResponseBodyItemsRiskEventsLabels {
	s.Label = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) SetLabelDesc(v string) *ListAiAppByPageResponseBodyItemsRiskEventsLabels {
	s.LabelDesc = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) SetType(v string) *ListAiAppByPageResponseBodyItemsRiskEventsLabels {
	s.Type = &v
	return s
}

func (s *ListAiAppByPageResponseBodyItemsRiskEventsLabels) Validate() error {
	return dara.Validate(s)
}
