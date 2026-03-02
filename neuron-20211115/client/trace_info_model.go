// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int64) *TraceInfo
	GetDuration() *int64
	SetOperationName(v string) *TraceInfo
	GetOperationName() *string
	SetServiceIp(v string) *TraceInfo
	GetServiceIp() *string
	SetServiceName(v string) *TraceInfo
	GetServiceName() *string
	SetTimestamp(v int64) *TraceInfo
	GetTimestamp() *int64
	SetTraceId(v string) *TraceInfo
	GetTraceId() *string
}

type TraceInfo struct {
	Duration      *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	OperationName *string `json:"operationName,omitempty" xml:"operationName,omitempty"`
	ServiceIp     *string `json:"serviceIp,omitempty" xml:"serviceIp,omitempty"`
	ServiceName   *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	Timestamp     *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	TraceId       *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TraceInfo) String() string {
	return dara.Prettify(s)
}

func (s TraceInfo) GoString() string {
	return s.String()
}

func (s *TraceInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *TraceInfo) GetOperationName() *string {
	return s.OperationName
}

func (s *TraceInfo) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *TraceInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *TraceInfo) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TraceInfo) GetTraceId() *string {
	return s.TraceId
}

func (s *TraceInfo) SetDuration(v int64) *TraceInfo {
	s.Duration = &v
	return s
}

func (s *TraceInfo) SetOperationName(v string) *TraceInfo {
	s.OperationName = &v
	return s
}

func (s *TraceInfo) SetServiceIp(v string) *TraceInfo {
	s.ServiceIp = &v
	return s
}

func (s *TraceInfo) SetServiceName(v string) *TraceInfo {
	s.ServiceName = &v
	return s
}

func (s *TraceInfo) SetTimestamp(v int64) *TraceInfo {
	s.Timestamp = &v
	return s
}

func (s *TraceInfo) SetTraceId(v string) *TraceInfo {
	s.TraceId = &v
	return s
}

func (s *TraceInfo) Validate() error {
	return dara.Validate(s)
}
