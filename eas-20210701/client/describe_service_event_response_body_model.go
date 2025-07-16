// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*DescribeServiceEventResponseBodyEvents) *DescribeServiceEventResponseBody
	GetEvents() []*DescribeServiceEventResponseBodyEvents
	SetPageNum(v int64) *DescribeServiceEventResponseBody
	GetPageNum() *int64
	SetRequestId(v string) *DescribeServiceEventResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeServiceEventResponseBody
	GetTotalCount() *int64
	SetTotalPageNum(v int64) *DescribeServiceEventResponseBody
	GetTotalPageNum() *int64
}

type DescribeServiceEventResponseBody struct {
	// The events.
	Events []*DescribeServiceEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D491C94-6239-5318-B4B4-799D859***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 29
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 12
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s DescribeServiceEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventResponseBody) GetEvents() []*DescribeServiceEventResponseBodyEvents {
	return s.Events
}

func (s *DescribeServiceEventResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeServiceEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceEventResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeServiceEventResponseBody) GetTotalPageNum() *int64 {
	return s.TotalPageNum
}

func (s *DescribeServiceEventResponseBody) SetEvents(v []*DescribeServiceEventResponseBodyEvents) *DescribeServiceEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeServiceEventResponseBody) SetPageNum(v int64) *DescribeServiceEventResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceEventResponseBody) SetRequestId(v string) *DescribeServiceEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceEventResponseBody) SetTotalCount(v int64) *DescribeServiceEventResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeServiceEventResponseBody) SetTotalPageNum(v int64) *DescribeServiceEventResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *DescribeServiceEventResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceEventResponseBodyEvents struct {
	// The returned message. The message is formatted and returned in the JSON format.
	//
	// example:
	//
	// {\\"versionId\\":1,\\"message\\":\\"Stage scale complete\\",\\"availableInstance\\":1,\\"unavailableInstance\\":0}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The cause of the event. The information about the change in the service status is returned.
	//
	// example:
	//
	// Updating
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The time when the event occurred. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-09 06:30:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The event type. Valid values:
	//
	// 	- Normal
	//
	// 	- Warning
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeServiceEventResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventResponseBodyEvents) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceEventResponseBodyEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeServiceEventResponseBodyEvents) GetTime() *string {
	return s.Time
}

func (s *DescribeServiceEventResponseBodyEvents) GetType() *string {
	return s.Type
}

func (s *DescribeServiceEventResponseBodyEvents) SetMessage(v string) *DescribeServiceEventResponseBodyEvents {
	s.Message = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) SetReason(v string) *DescribeServiceEventResponseBodyEvents {
	s.Reason = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) SetTime(v string) *DescribeServiceEventResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) SetType(v string) *DescribeServiceEventResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *DescribeServiceEventResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
