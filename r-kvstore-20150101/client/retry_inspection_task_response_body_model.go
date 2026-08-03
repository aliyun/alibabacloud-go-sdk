// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetryInspectionTaskResponseBody
	GetRequestId() *string
}

type RetryInspectionTaskResponseBody struct {
	// example:
	//
	// 794120D1-E0CF-4713-BAE4-EBAEA04506AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryInspectionTaskResponseBody) SetRequestId(v string) *RetryInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
