// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicySetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePolicySetResponseBody
	GetRequestId() *string
}

type UpdatePolicySetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolicySetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicySetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicySetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicySetResponseBody) SetRequestId(v string) *UpdatePolicySetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicySetResponseBody) Validate() error {
	return dara.Validate(s)
}
