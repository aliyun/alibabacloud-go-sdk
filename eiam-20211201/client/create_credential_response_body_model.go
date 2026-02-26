// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *CreateCredentialResponseBody
	GetCredentialId() *string
	SetRequestId(v string) *CreateCredentialResponseBody
	GetRequestId() *string
}

type CreateCredentialResponseBody struct {
	// example:
	//
	// cred_mkv7rgt4d7i4u7zqtzev2mxxxx
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBody) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCredentialResponseBody) SetCredentialId(v string) *CreateCredentialResponseBody {
	s.CredentialId = &v
	return s
}

func (s *CreateCredentialResponseBody) SetRequestId(v string) *CreateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
