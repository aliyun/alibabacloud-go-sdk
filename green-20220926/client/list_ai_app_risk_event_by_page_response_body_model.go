// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppRiskEventByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAiAppRiskEventByPageResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListAiAppRiskEventByPageResponseBodyItems) *ListAiAppRiskEventByPageResponseBody
	GetItems() []*ListAiAppRiskEventByPageResponseBodyItems
	SetMaxResults(v int32) *ListAiAppRiskEventByPageResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAiAppRiskEventByPageResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListAiAppRiskEventByPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAiAppRiskEventByPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAiAppRiskEventByPageResponseBody
	GetTotalCount() *int64
}

type ListAiAppRiskEventByPageResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned data.
	Items []*ListAiAppRiskEventByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of results returned per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more pages exist.
	//
	// example:
	//
	// 1a320d468c75e987f297484532c16e34d0ab6e7e43f8b73d
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAiAppRiskEventByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventByPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAiAppRiskEventByPageResponseBody) GetItems() []*ListAiAppRiskEventByPageResponseBodyItems {
	return s.Items
}

func (s *ListAiAppRiskEventByPageResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAiAppRiskEventByPageResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAiAppRiskEventByPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAiAppRiskEventByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiAppRiskEventByPageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAiAppRiskEventByPageResponseBody) SetCurrentPage(v int32) *ListAiAppRiskEventByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) SetItems(v []*ListAiAppRiskEventByPageResponseBodyItems) *ListAiAppRiskEventByPageResponseBody {
	s.Items = v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) SetMaxResults(v int32) *ListAiAppRiskEventByPageResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) SetNextToken(v string) *ListAiAppRiskEventByPageResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) SetPageSize(v int32) *ListAiAppRiskEventByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) SetRequestId(v string) *ListAiAppRiskEventByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) SetTotalCount(v int64) *ListAiAppRiskEventByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBody) Validate() error {
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

type ListAiAppRiskEventByPageResponseBodyItems struct {
	// The unique ID of the AI application.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the AI application.
	//
	// example:
	//
	// name-xxx
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The channel source.
	//
	// example:
	//
	// bailian
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The end time that indicates when the event was resolved.
	//
	// example:
	//
	// 2026-01-10 11:42:31
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event code that identifies the type or category of the event.
	//
	// example:
	//
	// de_aamexg3015
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The detailed description of the risk event.
	//
	// example:
	//
	// xxx
	EventDesc *string `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	// The detailed description of the risk event in English.
	//
	// example:
	//
	// xxx
	EventDescEn *string `json:"EventDescEn,omitempty" xml:"EventDescEn,omitempty"`
	// The event ID that uniquely identifies a risk event.
	//
	// example:
	//
	// e-a7gvnv3vid536dfxj
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The brief name that describes the risk event.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The time when the event was handled.
	//
	// example:
	//
	// 2026-01-10 11:42:31
	HandleTime *string `json:"HandleTime,omitempty" xml:"HandleTime,omitempty"`
	// The label used to mark or categorize the event.
	//
	// example:
	//
	// label-03
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The detailed description of the label.
	//
	// example:
	//
	// xxxx
	LabelDesc *string `json:"LabelDesc,omitempty" xml:"LabelDesc,omitempty"`
	// The risk level that indicates the severity of the event, such as high, medium, or low.
	//
	// example:
	//
	// normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The start time that indicates when the event occurred.
	//
	// example:
	//
	// 2025-12-21 15:30:19
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status that indicates the current processing state of the event, such as pending or resolved.
	//
	// example:
	//
	// resolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The event type that indicates the category of the risk event, such as security or performance.
	//
	// example:
	//
	// 0
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAiAppRiskEventByPageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetAppId() *string {
	return s.AppId
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetAppName() *string {
	return s.AppName
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetChannel() *string {
	return s.Channel
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetEventCode() *string {
	return s.EventCode
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetEventDesc() *string {
	return s.EventDesc
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetEventDescEn() *string {
	return s.EventDescEn
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetEventId() *string {
	return s.EventId
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetEventName() *string {
	return s.EventName
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetHandleTime() *string {
	return s.HandleTime
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetLabel() *string {
	return s.Label
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetLevel() *string {
	return s.Level
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetAppId(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.AppId = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetAppName(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.AppName = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetChannel(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.Channel = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetEndTime(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetEventCode(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.EventCode = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetEventDesc(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.EventDesc = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetEventDescEn(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.EventDescEn = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetEventId(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.EventId = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetEventName(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.EventName = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetHandleTime(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.HandleTime = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetLabel(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.Label = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetLabelDesc(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.LabelDesc = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetLevel(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.Level = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetStartTime(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetStatus(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) SetType(v string) *ListAiAppRiskEventByPageResponseBodyItems {
	s.Type = &v
	return s
}

func (s *ListAiAppRiskEventByPageResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
