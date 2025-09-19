// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryAgentlessTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetryAgentlessTaskResponseBody
	GetRequestId() *string
}

type RetryAgentlessTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryAgentlessTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryAgentlessTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryAgentlessTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryAgentlessTaskResponseBody) SetRequestId(v string) *RetryAgentlessTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryAgentlessTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
