// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationFederatedCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationFederatedCredentialResponseBody
	GetRequestId() *string
}

type DisableApplicationFederatedCredentialResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationFederatedCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationFederatedCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationFederatedCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationFederatedCredentialResponseBody) SetRequestId(v string) *DisableApplicationFederatedCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationFederatedCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
