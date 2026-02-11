// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipleTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody
	GetMultiCallChainInfos() []*GetMultipleTraceResponseBodyMultiCallChainInfos
	SetRequestId(v string) *GetMultipleTraceResponseBody
	GetRequestId() *string
}

type GetMultipleTraceResponseBody struct {
	MultiCallChainInfos []*GetMultipleTraceResponseBodyMultiCallChainInfos `json:"MultiCallChainInfos,omitempty" xml:"MultiCallChainInfos,omitempty" type:"Repeated"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultipleTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBody) GetMultiCallChainInfos() []*GetMultipleTraceResponseBodyMultiCallChainInfos {
	return s.MultiCallChainInfos
}

func (s *GetMultipleTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultipleTraceResponseBody) SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody {
	s.MultiCallChainInfos = v
	return s
}

func (s *GetMultipleTraceResponseBody) SetRequestId(v string) *GetMultipleTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultipleTraceResponseBody) Validate() error {
	if s.MultiCallChainInfos != nil {
		for _, item := range s.MultiCallChainInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMultipleTraceResponseBodyMultiCallChainInfos struct {
	Spans   []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
	TraceID *string                                                 `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) GetSpans() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	return s.Spans
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) GetTraceID() *string {
	return s.TraceID
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetSpans(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.Spans = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.TraceID = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) Validate() error {
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

type GetMultipleTraceResponseBodyMultiCallChainInfosSpans struct {
	Duration      *int64                                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                                               `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	OperationName *string                                                             `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                                             `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                                             `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType       *int32                                                              `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServiceIp     *string                                                             `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                                             `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                                             `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                                                              `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                                             `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetDuration() *int64 {
	return s.Duration
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetHaveStack() *bool {
	return s.HaveStack
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetLogEventList() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	return s.LogEventList
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetOperationName() *string {
	return s.OperationName
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetParentSpanId() *string {
	return s.ParentSpanId
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetRpcId() *string {
	return s.RpcId
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetRpcType() *int32 {
	return s.RpcType
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetSpanId() *string {
	return s.SpanId
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetTagEntryList() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	return s.TagEntryList
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GetTraceID() *string {
	return s.TraceID
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetDuration(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Duration = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetHaveStack(v bool) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.HaveStack = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetLogEventList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.LogEventList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetOperationName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.OperationName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetParentSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetResultCode(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ResultCode = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcType(v int32) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcType = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceIp(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceIp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.SpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TraceID = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) Validate() error {
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

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList struct {
	TagEntryList []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp    *int64                                                                          `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GetTagEntryList() []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	return s.TagEntryList
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) Validate() error {
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

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) Validate() error {
	return dara.Validate(s)
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) String() string {
	return dara.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GetKey() *string {
	return s.Key
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GetValue() *string {
	return s.Value
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Value = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) Validate() error {
	return dara.Validate(s)
}
