// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTraceResponseBody
	GetRequestId() *string
	SetSpans(v []*GetTraceResponseBodySpans) *GetTraceResponseBody
	GetSpans() []*GetTraceResponseBodySpans
}

type GetTraceResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spans     []*GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceResponseBody) GetSpans() []*GetTraceResponseBodySpans {
	return s.Spans
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSpans(v []*GetTraceResponseBodySpans) *GetTraceResponseBody {
	s.Spans = v
	return s
}

func (s *GetTraceResponseBody) Validate() error {
	if s.Spans != nil {
		for _, item := range s.Spans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodySpans struct {
	Duration      *int64                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                    `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  []*GetTraceResponseBodySpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	OperationName *string                                  `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                  `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                  `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType       *int32                                   `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServiceIp     *string                                  `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                  `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                  `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  []*GetTraceResponseBodySpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                  `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpans) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) GetDuration() *int64 {
	return s.Duration
}

func (s *GetTraceResponseBodySpans) GetHaveStack() *bool {
	return s.HaveStack
}

func (s *GetTraceResponseBodySpans) GetLogEventList() []*GetTraceResponseBodySpansLogEventList {
	return s.LogEventList
}

func (s *GetTraceResponseBodySpans) GetOperationName() *string {
	return s.OperationName
}

func (s *GetTraceResponseBodySpans) GetParentSpanId() *string {
	return s.ParentSpanId
}

func (s *GetTraceResponseBodySpans) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetTraceResponseBodySpans) GetRpcId() *string {
	return s.RpcId
}

func (s *GetTraceResponseBodySpans) GetRpcType() *int32 {
	return s.RpcType
}

func (s *GetTraceResponseBodySpans) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *GetTraceResponseBodySpans) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetTraceResponseBodySpans) GetSpanId() *string {
	return s.SpanId
}

func (s *GetTraceResponseBodySpans) GetTagEntryList() []*GetTraceResponseBodySpansTagEntryList {
	return s.TagEntryList
}

func (s *GetTraceResponseBodySpans) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetTraceResponseBodySpans) GetTraceID() *string {
	return s.TraceID
}

func (s *GetTraceResponseBodySpans) SetDuration(v int64) *GetTraceResponseBodySpans {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetHaveStack(v bool) *GetTraceResponseBodySpans {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetLogEventList(v []*GetTraceResponseBodySpansLogEventList) *GetTraceResponseBodySpans {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetOperationName(v string) *GetTraceResponseBodySpans {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetParentSpanId(v string) *GetTraceResponseBodySpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetResultCode(v string) *GetTraceResponseBodySpans {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcId(v string) *GetTraceResponseBodySpans {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcType(v int32) *GetTraceResponseBodySpans {
	s.RpcType = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceIp(v string) *GetTraceResponseBodySpans {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceName(v string) *GetTraceResponseBodySpans {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetSpanId(v string) *GetTraceResponseBodySpans {
	s.SpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTagEntryList(v []*GetTraceResponseBodySpansTagEntryList) *GetTraceResponseBodySpans {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetTimestamp(v int64) *GetTraceResponseBodySpans {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTraceID(v string) *GetTraceResponseBodySpans {
	s.TraceID = &v
	return s
}

func (s *GetTraceResponseBodySpans) Validate() error {
	if s.LogEventList != nil {
		for _, item := range s.LogEventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagEntryList != nil {
		for _, item := range s.TagEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodySpansLogEventList struct {
	TagEntryList []*GetTraceResponseBodySpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp    *int64                                               `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventList) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventList) GetTagEntryList() []*GetTraceResponseBodySpansLogEventListTagEntryList {
	return s.TagEntryList
}

func (s *GetTraceResponseBodySpansLogEventList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetTraceResponseBodySpansLogEventList) SetTagEntryList(v []*GetTraceResponseBodySpansLogEventListTagEntryList) *GetTraceResponseBodySpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansLogEventList) SetTimestamp(v int64) *GetTraceResponseBodySpansLogEventList {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventList) Validate() error {
	if s.TagEntryList != nil {
		for _, item := range s.TagEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodySpansLogEventListTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetKey(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetValue(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodySpansTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodySpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetTraceResponseBodySpansTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetTraceResponseBodySpansTagEntryList) SetKey(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansTagEntryList) SetValue(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Value = &v
	return s
}

func (s *GetTraceResponseBodySpansTagEntryList) Validate() error {
	return dara.Validate(s)
}
