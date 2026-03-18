// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCredentialProviderResponseBody
	GetRequestId() *string
}

type UpdateCredentialProviderResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialProviderResponseBody) SetRequestId(v string) *UpdateCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
