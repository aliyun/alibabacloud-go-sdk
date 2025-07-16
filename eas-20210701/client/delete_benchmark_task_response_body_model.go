// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBenchmarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteBenchmarkTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBenchmarkTaskResponseBody
	GetRequestId() *string
}

type DeleteBenchmarkTaskResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Benchmark task [benchmark-test-service-234c] is Deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBenchmarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBenchmarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBenchmarkTaskResponseBody) SetMessage(v string) *DeleteBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBenchmarkTaskResponseBody) SetRequestId(v string) *DeleteBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBenchmarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
