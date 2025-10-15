// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFederatedCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFederatedCredentialProviderResponseBody
	GetRequestId() *string
}

type DeleteFederatedCredentialProviderResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFederatedCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFederatedCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFederatedCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFederatedCredentialProviderResponseBody) SetRequestId(v string) *DeleteFederatedCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFederatedCredentialProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
