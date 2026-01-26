// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchTracesResponseBody
	GetRequestId() *string
	SetTraceInfos(v []*SearchTracesResponseBodyTraceInfos) *SearchTracesResponseBody
	GetTraceInfos() []*SearchTracesResponseBodyTraceInfos
}

type SearchTracesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the returned traces.
	TraceInfos []*SearchTracesResponseBodyTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTracesResponseBody) GetTraceInfos() []*SearchTracesResponseBodyTraceInfos {
	return s.TraceInfos
}

func (s *SearchTracesResponseBody) SetRequestId(v string) *SearchTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTracesResponseBody) SetTraceInfos(v []*SearchTracesResponseBodyTraceInfos) *SearchTracesResponseBody {
	s.TraceInfos = v
	return s
}

func (s *SearchTracesResponseBody) Validate() error {
	if s.TraceInfos != nil {
		for _, item := range s.TraceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTracesResponseBodyTraceInfos struct {
	// The amount of time consumed by the trace. Unit: milliseconds.
	//
	// example:
	//
	// 6
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the traced span.
	//
	// example:
	//
	// get***
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The IP address of the host where the application resides.
	//
	// example:
	//
	// ``172.20.**.**``
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// arms-k8s-demo-subcomponent
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Span ID. You can get it from the **Trace Explorer*	- page of the ARMS console.
	//
	// example:
	//
	// be3d6dcf5750e***
	SpanID *string `json:"SpanID,omitempty" xml:"SpanID,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1595174436993
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// ac1400a115951744369937024d****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyTraceInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesResponseBodyTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyTraceInfos) GetDuration() *int64 {
	return s.Duration
}

func (s *SearchTracesResponseBodyTraceInfos) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesResponseBodyTraceInfos) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesResponseBodyTraceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesResponseBodyTraceInfos) GetSpanID() *string {
	return s.SpanID
}

func (s *SearchTracesResponseBodyTraceInfos) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SearchTracesResponseBodyTraceInfos) GetTraceID() *string {
	return s.TraceID
}

func (s *SearchTracesResponseBodyTraceInfos) SetDuration(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetOperationName(v string) *SearchTracesResponseBodyTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceIp(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceName(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetSpanID(v string) *SearchTracesResponseBodyTraceInfos {
	s.SpanID = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTimestamp(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTraceID(v string) *SearchTracesResponseBodyTraceInfos {
	s.TraceID = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) Validate() error {
	return dara.Validate(s)
}
