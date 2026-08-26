// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRtcMPUEventSubRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListRtcMPUEventSubRecordResponseBody
	GetCount() *int64
	SetHasMore(v bool) *ListRtcMPUEventSubRecordResponseBody
	GetHasMore() *bool
	SetLogs(v []*ListRtcMPUEventSubRecordResponseBodyLogs) *ListRtcMPUEventSubRecordResponseBody
	GetLogs() []*ListRtcMPUEventSubRecordResponseBodyLogs
	SetRequestId(v string) *ListRtcMPUEventSubRecordResponseBody
	GetRequestId() *string
}

type ListRtcMPUEventSubRecordResponseBody struct {
	// The total number of callback records returned on the current page.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Indicates whether there is a next page.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The callback records.
	Logs []*ListRtcMPUEventSubRecordResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRtcMPUEventSubRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUEventSubRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListRtcMPUEventSubRecordResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListRtcMPUEventSubRecordResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListRtcMPUEventSubRecordResponseBody) GetLogs() []*ListRtcMPUEventSubRecordResponseBodyLogs {
	return s.Logs
}

func (s *ListRtcMPUEventSubRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRtcMPUEventSubRecordResponseBody) SetCount(v int64) *ListRtcMPUEventSubRecordResponseBody {
	s.Count = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBody) SetHasMore(v bool) *ListRtcMPUEventSubRecordResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBody) SetLogs(v []*ListRtcMPUEventSubRecordResponseBodyLogs) *ListRtcMPUEventSubRecordResponseBody {
	s.Logs = v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBody) SetRequestId(v string) *ListRtcMPUEventSubRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBody) Validate() error {
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

type ListRtcMPUEventSubRecordResponseBodyLogs struct {
	// The ID of the subscribed application.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://testcallback***.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The callback duration. Unit: milliseconds.
	//
	// example:
	//
	// 22
	Cost *int64 `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The callback content. For more information, see [Create a stream mixing and forwarding event callback](https://help.aliyun.com/document_detail/2804583.html).
	//
	// example:
	//
	// {\\"EventType\\":1,\\"MsgId\\":\\"42bba8b5-94ab-468c-9dae-9b501dd6c***\\",\\"AppId\\":\\"rtcdev\\",\\"SubId\\":\\"Sub-9799B2C45009799B2C4***\\",\\"TaskId\\":\\"mpucallbacktest\\",\\"CallbackTs\\":1712656430***,\\"Payload\\":{\\"DstUrl\\":\\"rtmp://domain/app/stream?auth\\",\\"EventTs\\":1712656430***,\\"EventCode\\":1,\\"ErrorCode\\":0,\\"ErrorMessage\\":\\"\\"}}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code. A value of 200 indicates that the callback was successful.
	//
	// example:
	//
	// 200
	HTTPCode *string `json:"HTTPCode,omitempty" xml:"HTTPCode,omitempty"`
	// The callback record ID.
	//
	// example:
	//
	// 42bba8b5-********-9b501dd6cb6e
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The event callback ID.
	//
	// example:
	//
	// Sub-******9799B2C4500******
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The time when the callback was invoked.
	//
	// Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 1970-01-01T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListRtcMPUEventSubRecordResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUEventSubRecordResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetAppId() *string {
	return s.AppId
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetCost() *int64 {
	return s.Cost
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetData() *string {
	return s.Data
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetHTTPCode() *string {
	return s.HTTPCode
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetMsgId() *string {
	return s.MsgId
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetSubId() *string {
	return s.SubId
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) GetTime() *string {
	return s.Time
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetAppId(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.AppId = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetCallbackUrl(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.CallbackUrl = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetCost(v int64) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.Cost = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetData(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.Data = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetHTTPCode(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.HTTPCode = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetMsgId(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.MsgId = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetSubId(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.SubId = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) SetTime(v string) *ListRtcMPUEventSubRecordResponseBodyLogs {
	s.Time = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
