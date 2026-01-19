// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientPublicKeyId(v string) *CreateClientPublicKeyResponseBody
	GetClientPublicKeyId() *string
	SetRequestId(v string) *CreateClientPublicKeyResponseBody
	GetRequestId() *string
}

type CreateClientPublicKeyResponseBody struct {
	// example:
	//
	// KEYCKmEYW9byWTdjuRbmCjd2Bhg6VpkAxxxx
	ClientPublicKeyId *string `json:"ClientPublicKeyId,omitempty" xml:"ClientPublicKeyId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClientPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientPublicKeyResponseBody) GetClientPublicKeyId() *string {
	return s.ClientPublicKeyId
}

func (s *CreateClientPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientPublicKeyResponseBody) SetClientPublicKeyId(v string) *CreateClientPublicKeyResponseBody {
	s.ClientPublicKeyId = &v
	return s
}

func (s *CreateClientPublicKeyResponseBody) SetRequestId(v string) *CreateClientPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
