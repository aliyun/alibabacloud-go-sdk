// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceInfoDetailResult interface {
	dara.Model
	String() string
	GoString() string
	SetTraceDetails(v []*TraceInfoDetailResultTraceDetails) *TraceInfoDetailResult
	GetTraceDetails() []*TraceInfoDetailResultTraceDetails
}

type TraceInfoDetailResult struct {
	TraceDetails []*TraceInfoDetailResultTraceDetails `json:"traceDetails,omitempty" xml:"traceDetails,omitempty" type:"Repeated"`
}

func (s TraceInfoDetailResult) String() string {
	return dara.Prettify(s)
}

func (s TraceInfoDetailResult) GoString() string {
	return s.String()
}

func (s *TraceInfoDetailResult) GetTraceDetails() []*TraceInfoDetailResultTraceDetails {
	return s.TraceDetails
}

func (s *TraceInfoDetailResult) SetTraceDetails(v []*TraceInfoDetailResultTraceDetails) *TraceInfoDetailResult {
	s.TraceDetails = v
	return s
}

func (s *TraceInfoDetailResult) Validate() error {
	if s.TraceDetails != nil {
		for _, item := range s.TraceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TraceInfoDetailResultTraceDetails struct {
	Children      []map[string]interface{}  `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	Duration      *int64                    `json:"duration,omitempty" xml:"duration,omitempty"`
	HaveStack     *bool                     `json:"haveStack,omitempty" xml:"haveStack,omitempty"`
	LogEventList  []*TraceSpansLogEventList `json:"logEventList,omitempty" xml:"logEventList,omitempty" type:"Repeated"`
	OperationName *string                   `json:"operationName,omitempty" xml:"operationName,omitempty"`
	ParentSpanId  *string                   `json:"parentSpanId,omitempty" xml:"parentSpanId,omitempty"`
	ResultCode    *string                   `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
	RpcId         *string                   `json:"rpcId,omitempty" xml:"rpcId,omitempty"`
	RpcType       *int64                    `json:"rpcType,omitempty" xml:"rpcType,omitempty"`
	RpcTypeName   *string                   `json:"rpcTypeName,omitempty" xml:"rpcTypeName,omitempty"`
	ServiceIp     *string                   `json:"serviceIp,omitempty" xml:"serviceIp,omitempty"`
	ServiceName   *string                   `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	SpanId        *string                   `json:"spanId,omitempty" xml:"spanId,omitempty"`
	TagEntryList  []*TagEntry               `json:"tagEntryList,omitempty" xml:"tagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                    `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	TraceId       *string                   `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TraceInfoDetailResultTraceDetails) String() string {
	return dara.Prettify(s)
}

func (s TraceInfoDetailResultTraceDetails) GoString() string {
	return s.String()
}

func (s *TraceInfoDetailResultTraceDetails) GetChildren() []map[string]interface{} {
	return s.Children
}

func (s *TraceInfoDetailResultTraceDetails) GetDuration() *int64 {
	return s.Duration
}

func (s *TraceInfoDetailResultTraceDetails) GetHaveStack() *bool {
	return s.HaveStack
}

func (s *TraceInfoDetailResultTraceDetails) GetLogEventList() []*TraceSpansLogEventList {
	return s.LogEventList
}

func (s *TraceInfoDetailResultTraceDetails) GetOperationName() *string {
	return s.OperationName
}

func (s *TraceInfoDetailResultTraceDetails) GetParentSpanId() *string {
	return s.ParentSpanId
}

func (s *TraceInfoDetailResultTraceDetails) GetResultCode() *string {
	return s.ResultCode
}

func (s *TraceInfoDetailResultTraceDetails) GetRpcId() *string {
	return s.RpcId
}

func (s *TraceInfoDetailResultTraceDetails) GetRpcType() *int64 {
	return s.RpcType
}

func (s *TraceInfoDetailResultTraceDetails) GetRpcTypeName() *string {
	return s.RpcTypeName
}

func (s *TraceInfoDetailResultTraceDetails) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *TraceInfoDetailResultTraceDetails) GetServiceName() *string {
	return s.ServiceName
}

func (s *TraceInfoDetailResultTraceDetails) GetSpanId() *string {
	return s.SpanId
}

func (s *TraceInfoDetailResultTraceDetails) GetTagEntryList() []*TagEntry {
	return s.TagEntryList
}

func (s *TraceInfoDetailResultTraceDetails) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TraceInfoDetailResultTraceDetails) GetTraceId() *string {
	return s.TraceId
}

func (s *TraceInfoDetailResultTraceDetails) SetChildren(v []map[string]interface{}) *TraceInfoDetailResultTraceDetails {
	s.Children = v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetDuration(v int64) *TraceInfoDetailResultTraceDetails {
	s.Duration = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetHaveStack(v bool) *TraceInfoDetailResultTraceDetails {
	s.HaveStack = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetLogEventList(v []*TraceSpansLogEventList) *TraceInfoDetailResultTraceDetails {
	s.LogEventList = v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetOperationName(v string) *TraceInfoDetailResultTraceDetails {
	s.OperationName = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetParentSpanId(v string) *TraceInfoDetailResultTraceDetails {
	s.ParentSpanId = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetResultCode(v string) *TraceInfoDetailResultTraceDetails {
	s.ResultCode = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetRpcId(v string) *TraceInfoDetailResultTraceDetails {
	s.RpcId = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetRpcType(v int64) *TraceInfoDetailResultTraceDetails {
	s.RpcType = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetRpcTypeName(v string) *TraceInfoDetailResultTraceDetails {
	s.RpcTypeName = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetServiceIp(v string) *TraceInfoDetailResultTraceDetails {
	s.ServiceIp = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetServiceName(v string) *TraceInfoDetailResultTraceDetails {
	s.ServiceName = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetSpanId(v string) *TraceInfoDetailResultTraceDetails {
	s.SpanId = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetTagEntryList(v []*TagEntry) *TraceInfoDetailResultTraceDetails {
	s.TagEntryList = v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetTimestamp(v int64) *TraceInfoDetailResultTraceDetails {
	s.Timestamp = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) SetTraceId(v string) *TraceInfoDetailResultTraceDetails {
	s.TraceId = &v
	return s
}

func (s *TraceInfoDetailResultTraceDetails) Validate() error {
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
