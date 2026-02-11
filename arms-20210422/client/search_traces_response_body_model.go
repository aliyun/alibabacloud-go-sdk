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
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
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
