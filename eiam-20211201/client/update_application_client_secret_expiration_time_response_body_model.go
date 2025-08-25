// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationClientSecretExpirationTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationClientSecretExpirationTimeResponseBody
	GetRequestId() *string
}

type UpdateApplicationClientSecretExpirationTimeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationClientSecretExpirationTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationClientSecretExpirationTimeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationClientSecretExpirationTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationClientSecretExpirationTimeResponseBody) SetRequestId(v string) *UpdateApplicationClientSecretExpirationTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationClientSecretExpirationTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
