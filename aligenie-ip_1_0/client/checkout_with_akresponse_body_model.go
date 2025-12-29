// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckoutWithAKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CheckoutWithAKResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckoutWithAKResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckoutWithAKResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *CheckoutWithAKResponseBody
	GetStatusCode() *int32
}

type CheckoutWithAKResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CheckoutWithAKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckoutWithAKResponseBody) GoString() string {
	return s.String()
}

func (s *CheckoutWithAKResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckoutWithAKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckoutWithAKResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckoutWithAKResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckoutWithAKResponseBody) SetMessage(v string) *CheckoutWithAKResponseBody {
	s.Message = &v
	return s
}

func (s *CheckoutWithAKResponseBody) SetRequestId(v string) *CheckoutWithAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckoutWithAKResponseBody) SetResult(v bool) *CheckoutWithAKResponseBody {
	s.Result = &v
	return s
}

func (s *CheckoutWithAKResponseBody) SetStatusCode(v int32) *CheckoutWithAKResponseBody {
	s.StatusCode = &v
	return s
}

func (s *CheckoutWithAKResponseBody) Validate() error {
	return dara.Validate(s)
}
