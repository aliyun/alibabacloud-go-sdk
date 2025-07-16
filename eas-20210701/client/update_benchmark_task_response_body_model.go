// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateBenchmarkTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBenchmarkTaskResponseBody
	GetRequestId() *string
}

type UpdateBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-larec-test-1076] is Updating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBenchmarkTaskResponseBody) SetMessage(v string) *UpdateBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBenchmarkTaskResponseBody) SetRequestId(v string) *UpdateBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
