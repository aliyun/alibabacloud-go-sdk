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
	// The name of the attached application access point (AAP).
	//
	// This parameter is required.
	//
	// example:
	//
	// aap_test
	AapName *string `json:"AapName,omitempty" xml:"AapName,omitempty"`
	// The end time of the validity period of the ClientKey.
	//
	// The time is specified in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	//  >
	//
	//  >- If you do not specify NotAfter, the default value is the creation time of the ClientKey plus 5 years.
	//
	//  >- If you specify NotAfter, you must also specify NotBefore.
	//
	// example:
	//
	// 2028-08-31T17:14:33Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The start time of the validity period of the ClientKey.
	//
	// The time is specified in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	//  >
	//
	//  >- If you do not specify NotBefore, the default value is the creation time of the ClientKey.
	//
	//  >- If you specify NotBefore, you must also specify NotAfter.
	//
	// example:
	//
	// 2023-08-31T17:14:33Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The passphrase of the ClientKey.
	//
	// The passphrase must be 8 to 64 characters in length and can contain digits, uppercase and lowercase letters, and special characters such as `~!@#$%^&*?_-`. It must contain at least two of these character types.
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
