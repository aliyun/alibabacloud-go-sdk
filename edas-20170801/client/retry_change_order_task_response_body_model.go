// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChangeOrderTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RetryChangeOrderTaskResponseBody
	GetCode() *int32
	SetData(v string) *RetryChangeOrderTaskResponseBody
	GetData() *string
	SetMessage(v string) *RetryChangeOrderTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryChangeOrderTaskResponseBody
	GetRequestId() *string
}

type RetryChangeOrderTaskResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The retry information.
	//
	// example:
	//
	// success retry task
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4823-bhjf-23u4-eiufh
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryChangeOrderTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryChangeOrderTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryChangeOrderTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RetryChangeOrderTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *RetryChangeOrderTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryChangeOrderTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryChangeOrderTaskResponseBody) SetCode(v int32) *RetryChangeOrderTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RetryChangeOrderTaskResponseBody) SetData(v string) *RetryChangeOrderTaskResponseBody {
	s.Data = &v
	return s
}

func (s *RetryChangeOrderTaskResponseBody) SetMessage(v string) *RetryChangeOrderTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RetryChangeOrderTaskResponseBody) SetRequestId(v string) *RetryChangeOrderTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryChangeOrderTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
