// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFederatedCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableFederatedCredentialProviderResponseBody
	GetRequestId() *string
}

type DisableFederatedCredentialProviderResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableFederatedCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableFederatedCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFederatedCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableFederatedCredentialProviderResponseBody) SetRequestId(v string) *DisableFederatedCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableFederatedCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
