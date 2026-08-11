// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppRiskEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAiAppRiskEventResponseBodyData) *ListAiAppRiskEventResponseBody
	GetData() []*ListAiAppRiskEventResponseBodyData
	SetRequestId(v string) *ListAiAppRiskEventResponseBody
	GetRequestId() *string
}

type ListAiAppRiskEventResponseBody struct {
	// The returned data.
	Data []*ListAiAppRiskEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID assigned by the backend to uniquely identify a request. This ID can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAiAppRiskEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventResponseBody) GetData() []*ListAiAppRiskEventResponseBodyData {
	return s.Data
}

func (s *ListAiAppRiskEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAiAppRiskEventResponseBody) SetData(v []*ListAiAppRiskEventResponseBodyData) *ListAiAppRiskEventResponseBody {
	s.Data = v
	return s
}

func (s *ListAiAppRiskEventResponseBody) SetRequestId(v string) *ListAiAppRiskEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAiAppRiskEventResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAiAppRiskEventResponseBodyData struct {
	// The end time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2026-03-25 10:22:02
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event code that identifies the type or category of the event.
	//
	// example:
	//
	// hit-xxxx
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The event description that provides details about the risk event.
	//
	// example:
	//
	// desc-xxx
	EventDesc *string `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	// The event description in English.
	//
	// example:
	//
	// desc-xxx
	EventDescEn *string `json:"EventDescEn,omitempty" xml:"EventDescEn,omitempty"`
	// The event ID that uniquely identifies a risk event.
	//
	// example:
	//
	// id-xxx
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event name that briefly describes the risk event.
	//
	// example:
	//
	// name-xxx
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The label used to mark or categorize the event.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The label description that provides details about the label.
	//
	// example:
	//
	// desc-xxx
	LabelDesc *string `json:"LabelDesc,omitempty" xml:"LabelDesc,omitempty"`
	// The risk level that indicates the severity of the event, such as high, medium, or low.
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The effective period. Format: YYYY-MM-DD HH:mm:ss (default time zone: UTC+08:00).
	//
	// example:
	//
	// 2025-07-22 16:41:15
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The event status that indicates the current processing state of the event, such as pending or resolved.
	//
	// example:
	//
	// resovled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The event type that indicates the category of the risk event, such as security or performance.
	//
	// example:
	//
	// sensitiveData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAiAppRiskEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAiAppRiskEventResponseBodyData) GetEventCode() *string {
	return s.EventCode
}

func (s *ListAiAppRiskEventResponseBodyData) GetEventDesc() *string {
	return s.EventDesc
}

func (s *ListAiAppRiskEventResponseBodyData) GetEventDescEn() *string {
	return s.EventDescEn
}

func (s *ListAiAppRiskEventResponseBodyData) GetEventId() *string {
	return s.EventId
}

func (s *ListAiAppRiskEventResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *ListAiAppRiskEventResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *ListAiAppRiskEventResponseBodyData) GetLabelDesc() *string {
	return s.LabelDesc
}

func (s *ListAiAppRiskEventResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *ListAiAppRiskEventResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAiAppRiskEventResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAiAppRiskEventResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListAiAppRiskEventResponseBodyData) SetEndTime(v string) *ListAiAppRiskEventResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetEventCode(v string) *ListAiAppRiskEventResponseBodyData {
	s.EventCode = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetEventDesc(v string) *ListAiAppRiskEventResponseBodyData {
	s.EventDesc = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetEventDescEn(v string) *ListAiAppRiskEventResponseBodyData {
	s.EventDescEn = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetEventId(v string) *ListAiAppRiskEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetEventName(v string) *ListAiAppRiskEventResponseBodyData {
	s.EventName = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetLabel(v string) *ListAiAppRiskEventResponseBodyData {
	s.Label = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetLabelDesc(v string) *ListAiAppRiskEventResponseBodyData {
	s.LabelDesc = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetLevel(v string) *ListAiAppRiskEventResponseBodyData {
	s.Level = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetStartTime(v string) *ListAiAppRiskEventResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetStatus(v string) *ListAiAppRiskEventResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) SetType(v string) *ListAiAppRiskEventResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAiAppRiskEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
