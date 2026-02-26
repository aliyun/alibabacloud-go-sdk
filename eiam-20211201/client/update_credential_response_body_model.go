// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCredentialResponseBody
	GetRequestId() *string
}

type UpdateCredentialResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialResponseBody) SetRequestId(v string) *UpdateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
