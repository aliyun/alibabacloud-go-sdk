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
	// The number of entries per page.
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
	Logs []*ListRtcMPUEventSubRecordResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The ID of the request.
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
	return dara.Validate(s)
}

type ListRtcMPUEventSubRecordResponseBodyLogs struct {
	// The ID of the application.
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
	// For more information about the callback, see [CreateRtcMPUEventSub](https://help.aliyun.com/document_detail/2804583.html).
	//
	// example:
	//
	// {\\"EventType\\":1,\\"MsgId\\":\\"42bba8b5-94ab-468c-9dae-9b501dd6c***\\",\\"AppId\\":\\"rtcdev\\",\\"SubId\\":\\"Sub-9799B2C45009799B2C4***\\",\\"TaskId\\":\\"mpucallbacktest\\",\\"CallbackTs\\":1712656430***,\\"Payload\\":{\\"DstUrl\\":\\"rtmp://domain/app/stream?auth\\",\\"EventTs\\":1712656430***,\\"EventCode\\":1,\\"ErrorCode\\":0,\\"ErrorMessage\\":\\"\\"}}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code. 200 indicates that the callback is successful.
	//
	// example:
	//
	// 200
	HTTPCode *string `json:"HTTPCode,omitempty" xml:"HTTPCode,omitempty"`
	// The ID of the callback record.
	//
	// example:
	//
	// 42bba8b5-********-9b501dd6cb6e
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The ID of the subscription.
	//
	// example:
	//
	// Sub-******9799B2C4500******
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The time when the callback was invoked. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
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
