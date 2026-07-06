// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceCredentialResponseBody
	GetRequestId() *string
}

type UpdateServiceCredentialResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 53C8FBFD-F2E9-50F2-AD63-B6566B3D4D7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceCredentialResponseBody) SetRequestId(v string) *UpdateServiceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
