// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationFederatedCredentialResponseBody
	GetRequestId() *string
}

type UpdateApplicationFederatedCredentialResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationFederatedCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationFederatedCredentialResponseBody) SetRequestId(v string) *UpdateApplicationFederatedCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
