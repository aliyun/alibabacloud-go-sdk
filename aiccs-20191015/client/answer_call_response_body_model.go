// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AnswerCallResponseBody
	GetCode() *string
	SetMessage(v string) *AnswerCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *AnswerCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnswerCallResponseBody
	GetSuccess() *bool
}

type AnswerCallResponseBody struct {
	// Status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnswerCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *AnswerCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AnswerCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnswerCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnswerCallResponseBody) SetCode(v string) *AnswerCallResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerCallResponseBody) SetMessage(v string) *AnswerCallResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerCallResponseBody) SetRequestId(v string) *AnswerCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerCallResponseBody) SetSuccess(v bool) *AnswerCallResponseBody {
	s.Success = &v
	return s
}

func (s *AnswerCallResponseBody) Validate() error {
	return dara.Validate(s)
}
