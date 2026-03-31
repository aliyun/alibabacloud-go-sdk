// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMmsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *RetryMmsJobResponseBody
	GetData() *int64
	SetRequestId(v string) *RetryMmsJobResponseBody
	GetRequestId() *string
}

type RetryMmsJobResponseBody struct {
	// example:
	//
	// 78
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 7F5DAD1C-9EC2-5FE5-97CF-BCE21B4ABA29
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RetryMmsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *RetryMmsJobResponseBody) GetData() *int64 {
	return s.Data
}

func (s *RetryMmsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryMmsJobResponseBody) SetData(v int64) *RetryMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *RetryMmsJobResponseBody) SetRequestId(v string) *RetryMmsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryMmsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
