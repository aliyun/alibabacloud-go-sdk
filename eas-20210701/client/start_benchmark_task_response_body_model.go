// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *StartBenchmarkTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartBenchmarkTaskResponseBody
	GetRequestId() *string
}

type StartBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Starting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBenchmarkTaskResponseBody) SetMessage(v string) *StartBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartBenchmarkTaskResponseBody) SetRequestId(v string) *StartBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
