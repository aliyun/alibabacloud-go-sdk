// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationFederatedCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationFederatedCredentialResponseBody
	GetRequestId() *string
}

type DeleteApplicationFederatedCredentialResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationFederatedCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationFederatedCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationFederatedCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationFederatedCredentialResponseBody) SetRequestId(v string) *DeleteApplicationFederatedCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationFederatedCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
