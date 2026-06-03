// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyEmailResponseBody
	GetRequestId() *string
}

type VerifyEmailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyEmailResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyEmailResponseBody) SetRequestId(v string) *VerifyEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
