// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFederatedCredentialProviderDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFederatedCredentialProviderDescriptionResponseBody
	GetRequestId() *string
}

type UpdateFederatedCredentialProviderDescriptionResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFederatedCredentialProviderDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFederatedCredentialProviderDescriptionResponseBody) SetRequestId(v string) *UpdateFederatedCredentialProviderDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
