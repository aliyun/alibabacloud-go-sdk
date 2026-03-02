// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTraceDetailResponseBody
	GetRequestId() *string
	SetResult(v *TraceInfoDetailResult) *GetTraceDetailResponseBody
	GetResult() *TraceInfoDetailResult
}

type GetTraceDetailResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// sdadawqewe
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *TraceInfoDetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetTraceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceDetailResponseBody) GetResult() *TraceInfoDetailResult {
	return s.Result
}

func (s *GetTraceDetailResponseBody) SetRequestId(v string) *GetTraceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceDetailResponseBody) SetResult(v *TraceInfoDetailResult) *GetTraceDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetTraceDetailResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
