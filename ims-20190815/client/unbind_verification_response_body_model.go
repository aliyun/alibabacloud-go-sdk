// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindVerificationResponseBody
	GetRequestId() *string
}

type UnbindVerificationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindVerificationResponseBody) SetRequestId(v string) *UnbindVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindVerificationResponseBody) Validate() error {
	return dara.Validate(s)
}
