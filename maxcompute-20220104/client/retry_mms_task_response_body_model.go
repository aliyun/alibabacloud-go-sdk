// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMmsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *RetryMmsTaskResponseBody
	GetData() *int64
	SetRequestId(v string) *RetryMmsTaskResponseBody
	GetRequestId() *string
}

type RetryMmsTaskResponseBody struct {
	// example:
	//
	// success
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RetryMmsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryMmsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryMmsTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *RetryMmsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryMmsTaskResponseBody) SetData(v int64) *RetryMmsTaskResponseBody {
	s.Data = &v
	return s
}

func (s *RetryMmsTaskResponseBody) SetRequestId(v string) *RetryMmsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryMmsTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
