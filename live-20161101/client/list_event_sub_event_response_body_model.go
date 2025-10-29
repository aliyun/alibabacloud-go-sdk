// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventSubEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListEventSubEventResponseBody
	GetCount() *int64
	SetHasMore(v bool) *ListEventSubEventResponseBody
	GetHasMore() *bool
	SetLogs(v []*ListEventSubEventResponseBodyLogs) *ListEventSubEventResponseBody
	GetLogs() []*ListEventSubEventResponseBodyLogs
	SetRequestId(v string) *ListEventSubEventResponseBody
	GetRequestId() *string
}

type ListEventSubEventResponseBody struct {
	// The total number of callback records returned on the current page.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The callback records.
	Logs []*ListEventSubEventResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CC8CB656-A7BA-1811-9D6B-4CC187E988BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEventSubEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventSubEventResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListEventSubEventResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListEventSubEventResponseBody) GetLogs() []*ListEventSubEventResponseBodyLogs {
	return s.Logs
}

func (s *ListEventSubEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventSubEventResponseBody) SetCount(v int64) *ListEventSubEventResponseBody {
	s.Count = &v
	return s
}

func (s *ListEventSubEventResponseBody) SetHasMore(v bool) *ListEventSubEventResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListEventSubEventResponseBody) SetLogs(v []*ListEventSubEventResponseBodyLogs) *ListEventSubEventResponseBody {
	s.Logs = v
	return s
}

func (s *ListEventSubEventResponseBody) SetRequestId(v string) *ListEventSubEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventSubEventResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventSubEventResponseBodyLogs struct {
	// The application ID.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The HTTP status code. A value of 200 indicates that the callback was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The callback duration. Unit: milliseconds.
	//
	// example:
	//
	// 22
	Cost *int32 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The details about the callback.
	//
	// example:
	//
	// {\\"MsgId\\":\\"875d5266cbabb1834cc84a105cf68454\\",\\"MsgTimestamp\\":1697545591,\\"SubscribeId\\":\\"09be0d2254cb5a89f4cbd86403ec5343\\",\\"AppId\\":\\"xxx\\",\\"ChannelId\\":\\"9099\\",\\"Contents\\":[{\\"Event\\":\\"UserEvent\\",\\"UserEvent\\":{\\"UserId\\":\\"linux_test\\",\\"EventTag\\":\\"Leave\\",\\"SessionId\\":\\"je7y2sBZJZQ0VBJZrh4LnBkxvGH2WyVs\\",\\"Timestamp\\":1697545591,\\"ChannelProfile\\":\\"interactive_live\\",\\"US\\":5068748604047364,\\"Reason\\":1,\\"Role\\":1,\\"TerminalType\\":6,\\"UserType\\":2}}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the callback record.
	//
	// example:
	//
	// 875d5266cbabb1834cc84a105cf6****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The subscription ID.
	//
	// example:
	//
	// ad53276431c****
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The time when the callback was generated.
	//
	// example:
	//
	// 2023-10-17 20:26:31.988
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://****.com/callback
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListEventSubEventResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubEventResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListEventSubEventResponseBodyLogs) GetAppId() *string {
	return s.AppId
}

func (s *ListEventSubEventResponseBodyLogs) GetCode() *int32 {
	return s.Code
}

func (s *ListEventSubEventResponseBodyLogs) GetCost() *int32 {
	return s.Cost
}

func (s *ListEventSubEventResponseBodyLogs) GetData() *string {
	return s.Data
}

func (s *ListEventSubEventResponseBodyLogs) GetMessageId() *string {
	return s.MessageId
}

func (s *ListEventSubEventResponseBodyLogs) GetSubId() *string {
	return s.SubId
}

func (s *ListEventSubEventResponseBodyLogs) GetTime() *string {
	return s.Time
}

func (s *ListEventSubEventResponseBodyLogs) GetUrl() *string {
	return s.Url
}

func (s *ListEventSubEventResponseBodyLogs) SetAppId(v string) *ListEventSubEventResponseBodyLogs {
	s.AppId = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetCode(v int32) *ListEventSubEventResponseBodyLogs {
	s.Code = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetCost(v int32) *ListEventSubEventResponseBodyLogs {
	s.Cost = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetData(v string) *ListEventSubEventResponseBodyLogs {
	s.Data = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetMessageId(v string) *ListEventSubEventResponseBodyLogs {
	s.MessageId = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetSubId(v string) *ListEventSubEventResponseBodyLogs {
	s.SubId = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetTime(v string) *ListEventSubEventResponseBodyLogs {
	s.Time = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) SetUrl(v string) *ListEventSubEventResponseBodyLogs {
	s.Url = &v
	return s
}

func (s *ListEventSubEventResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
