// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCredentialProviderResponseBody
	GetRequestId() *string
}

type DisableCredentialProviderResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCredentialProviderResponseBody) SetRequestId(v string) *DisableCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
