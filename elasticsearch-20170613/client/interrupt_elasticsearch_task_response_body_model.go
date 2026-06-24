// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptElasticsearchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InterruptElasticsearchTaskResponseBody
	GetCode() *string
	SetMessage(v string) *InterruptElasticsearchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *InterruptElasticsearchTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *InterruptElasticsearchTaskResponseBody
	GetResult() *bool
}

type InterruptElasticsearchTaskResponseBody struct {
	// The error code. This parameter is returned only when an exception occurs.
	//
	// example:
	//
	// InstanceStatusNotSupportCurrentAction
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an exception occurs.
	//
	// example:
	//
	// The cluster is running tasks or in an error status. Try again later.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The change is interrupted.
	//
	// - false: The change failed to be interrupted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InterruptElasticsearchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InterruptElasticsearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *InterruptElasticsearchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *InterruptElasticsearchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InterruptElasticsearchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InterruptElasticsearchTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InterruptElasticsearchTaskResponseBody) SetCode(v string) *InterruptElasticsearchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) SetMessage(v string) *InterruptElasticsearchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) SetRequestId(v string) *InterruptElasticsearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) SetResult(v bool) *InterruptElasticsearchTaskResponseBody {
	s.Result = &v
	return s
}

func (s *InterruptElasticsearchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
