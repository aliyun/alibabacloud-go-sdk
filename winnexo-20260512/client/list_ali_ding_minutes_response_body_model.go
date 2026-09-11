// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingMinutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAliDingMinutesResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListAliDingMinutesResponseBody
	GetHasMore() *bool
	SetItems(v []*ListAliDingMinutesResponseBodyItems) *ListAliDingMinutesResponseBody
	GetItems() []*ListAliDingMinutesResponseBodyItems
	SetMessage(v string) *ListAliDingMinutesResponseBody
	GetMessage() *string
	SetNextCursor(v string) *ListAliDingMinutesResponseBody
	GetNextCursor() *string
	SetRequestId(v string) *ListAliDingMinutesResponseBody
	GetRequestId() *string
}

type ListAliDingMinutesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more pages are available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The location clusters.
	Items []*ListAliDingMinutesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next retrieval.
	//
	// example:
	//
	// opaque-next-cursor
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAliDingMinutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliDingMinutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAliDingMinutesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListAliDingMinutesResponseBody) GetItems() []*ListAliDingMinutesResponseBodyItems {
	return s.Items
}

func (s *ListAliDingMinutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAliDingMinutesResponseBody) GetNextCursor() *string {
	return s.NextCursor
}

func (s *ListAliDingMinutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliDingMinutesResponseBody) SetCode(v string) *ListAliDingMinutesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAliDingMinutesResponseBody) SetHasMore(v bool) *ListAliDingMinutesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListAliDingMinutesResponseBody) SetItems(v []*ListAliDingMinutesResponseBodyItems) *ListAliDingMinutesResponseBody {
	s.Items = v
	return s
}

func (s *ListAliDingMinutesResponseBody) SetMessage(v string) *ListAliDingMinutesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAliDingMinutesResponseBody) SetNextCursor(v string) *ListAliDingMinutesResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListAliDingMinutesResponseBody) SetRequestId(v string) *ListAliDingMinutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliDingMinutesResponseBody) Validate() error {
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

type ListAliDingMinutesResponseBodyItems struct {
	// The name of the creator.
	//
	// example:
	//
	// John Doe
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The execution duration of the asynchronous task.
	//
	// example:
	//
	// 3600000
	DurationMs *int64 `json:"durationMs,omitempty" xml:"durationMs,omitempty"`
	// The end timestamp, in milliseconds.
	//
	// example:
	//
	// 2026-09-08T10:00:00+08:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The DingTalk meeting minutes ID.
	//
	// example:
	//
	// 76327569643231383535353939365f3436383537393431335f32
	MinutesId *string `json:"minutesId,omitempty" xml:"minutesId,omitempty"`
	// The start timestamp, in milliseconds.
	//
	// example:
	//
	// 2026-09-08T09:00:00+08:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The task status. Running is returned upon submission.
	//
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The title of the scheduled meeting.
	//
	// example:
	//
	// Weekly Project Meeting
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListAliDingMinutesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingMinutesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAliDingMinutesResponseBodyItems) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListAliDingMinutesResponseBodyItems) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *ListAliDingMinutesResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAliDingMinutesResponseBodyItems) GetMinutesId() *string {
	return s.MinutesId
}

func (s *ListAliDingMinutesResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAliDingMinutesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListAliDingMinutesResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListAliDingMinutesResponseBodyItems) SetCreatorName(v string) *ListAliDingMinutesResponseBodyItems {
	s.CreatorName = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) SetDurationMs(v int64) *ListAliDingMinutesResponseBodyItems {
	s.DurationMs = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) SetEndTime(v string) *ListAliDingMinutesResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) SetMinutesId(v string) *ListAliDingMinutesResponseBodyItems {
	s.MinutesId = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) SetStartTime(v string) *ListAliDingMinutesResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) SetStatus(v string) *ListAliDingMinutesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) SetTitle(v string) *ListAliDingMinutesResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListAliDingMinutesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
