// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAapName(v string) *CreateClientKeyRequest
	GetAapName() *string
	SetNotAfter(v string) *CreateClientKeyRequest
	GetNotAfter() *string
	SetNotBefore(v string) *CreateClientKeyRequest
	GetNotBefore() *string
	SetPassword(v string) *CreateClientKeyRequest
	GetPassword() *string
}

type CreateClientKeyRequest struct {
	// The operation that you want to perform. Set the value to **CreateClientKey**.
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
	// The encryption password of the client key.
	//
	// The password must be 8 to 64 characters in length and must contain at least two of the following types: digits, letters, and special characters. Special characters include `~ ! @ # $ % ^ & 	- ? _ -`.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The end of the validity period of the client key.
	//
	// Specify the time in the ISO 8601 standard. The time must be in UTC. The time must be in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >
	//
	// 	- If you do not configure NotAfter, the default value is the time when the client key was created plus five years.
	//
	// 	- If you configure NotAfter, you must configure NotBefore.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The name of the AAP.
	//
	// This parameter is required.
	//
	// example:
	//
	// bcfefe15-46f0****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateClientKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateClientKeyRequest) GetAapName() *string {
	return s.AapName
}

func (s *CreateClientKeyRequest) GetNotAfter() *string {
	return s.NotAfter
}

func (s *CreateClientKeyRequest) GetNotBefore() *string {
	return s.NotBefore
}

func (s *CreateClientKeyRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateClientKeyRequest) SetAapName(v string) *CreateClientKeyRequest {
	s.AapName = &v
	return s
}

func (s *CreateClientKeyRequest) SetNotAfter(v string) *CreateClientKeyRequest {
	s.NotAfter = &v
	return s
}

func (s *CreateClientKeyRequest) SetNotBefore(v string) *CreateClientKeyRequest {
	s.NotBefore = &v
	return s
}

func (s *CreateClientKeyRequest) SetPassword(v string) *CreateClientKeyRequest {
	s.Password = &v
	return s
}

func (s *CreateClientKeyRequest) Validate() error {
	return dara.Validate(s)
}
