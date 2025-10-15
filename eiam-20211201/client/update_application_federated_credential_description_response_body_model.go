// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationFederatedCredentialDescriptionResponseBody
	GetRequestId() *string
}

type UpdateApplicationFederatedCredentialDescriptionResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationFederatedCredentialDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponseBody) SetRequestId(v string) *UpdateApplicationFederatedCredentialDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
