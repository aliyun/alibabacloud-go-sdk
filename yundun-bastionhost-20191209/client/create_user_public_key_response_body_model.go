// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicKeyId(v string) *CreateUserPublicKeyResponseBody
	GetPublicKeyId() *string
	SetRequestId(v string) *CreateUserPublicKeyResponseBody
	GetRequestId() *string
}

type CreateUserPublicKeyResponseBody struct {
	// The ID of the public key.
	//
	// example:
	//
	// 1
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5EAB922E-F476-5DFA-9290-313C608E724B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserPublicKeyResponseBody) GetPublicKeyId() *string {
	return s.PublicKeyId
}

func (s *CreateUserPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserPublicKeyResponseBody) SetPublicKeyId(v string) *CreateUserPublicKeyResponseBody {
	s.PublicKeyId = &v
	return s
}

func (s *CreateUserPublicKeyResponseBody) SetRequestId(v string) *CreateUserPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
