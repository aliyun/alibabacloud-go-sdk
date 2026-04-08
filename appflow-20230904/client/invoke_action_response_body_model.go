// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InvokeActionResponseBody
	GetRequestId() *string
	SetResult(v *InvokeActionResponseBodyResult) *InvokeActionResponseBody
	GetResult() *InvokeActionResponseBodyResult
}

type InvokeActionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 346E5EE9-D5FE-56A0-B3E2-A43E0F67302A
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InvokeActionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InvokeActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeActionResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeActionResponseBody) GetResult() *InvokeActionResponseBodyResult {
	return s.Result
}

func (s *InvokeActionResponseBody) SetRequestId(v string) *InvokeActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeActionResponseBody) SetResult(v *InvokeActionResponseBodyResult) *InvokeActionResponseBody {
	s.Result = v
	return s
}

func (s *InvokeActionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeActionResponseBodyResult struct {
	// example:
	//
	// The provided parameter xxx is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// {
	//
	//   "output": {
	//
	//     "finishReason": "stop",
	//
	//     "text": "hello"
	//
	//   },
	//
	//   "requestId": "433944e3-2163-9c7a-9664-2a27c0dda0ec",
	//
	//   "usage": {
	//
	//     "inputTokens": 13,
	//
	//     "outputTokens": 16,
	//
	//     "totalTokens": 29
	//
	//   }
	//
	// }
	Output interface{} `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// RUNNING、COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s InvokeActionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s InvokeActionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InvokeActionResponseBodyResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *InvokeActionResponseBodyResult) GetOutput() interface{} {
	return s.Output
}

func (s *InvokeActionResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *InvokeActionResponseBodyResult) SetErrorMessage(v string) *InvokeActionResponseBodyResult {
	s.ErrorMessage = &v
	return s
}

func (s *InvokeActionResponseBodyResult) SetOutput(v interface{}) *InvokeActionResponseBodyResult {
	s.Output = v
	return s
}

func (s *InvokeActionResponseBodyResult) SetStatus(v string) *InvokeActionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *InvokeActionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
