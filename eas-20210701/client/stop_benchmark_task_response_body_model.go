// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *StopBenchmarkTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopBenchmarkTaskResponseBody
	GetRequestId() *string
}

type StopBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Stopping
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopBenchmarkTaskResponseBody) SetMessage(v string) *StopBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopBenchmarkTaskResponseBody) SetRequestId(v string) *StopBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
