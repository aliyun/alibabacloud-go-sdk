// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChangeAccountEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetryChangeAccountEmailResponseBody
	GetRequestId() *string
}

type RetryChangeAccountEmailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryChangeAccountEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryChangeAccountEmailResponseBody) GoString() string {
	return s.String()
}

func (s *RetryChangeAccountEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryChangeAccountEmailResponseBody) SetRequestId(v string) *RetryChangeAccountEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryChangeAccountEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
