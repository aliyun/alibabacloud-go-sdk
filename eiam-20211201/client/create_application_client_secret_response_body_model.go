// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationClientSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationClientSecret(v *CreateApplicationClientSecretResponseBodyApplicationClientSecret) *CreateApplicationClientSecretResponseBody
	GetApplicationClientSecret() *CreateApplicationClientSecretResponseBodyApplicationClientSecret
	SetRequestId(v string) *CreateApplicationClientSecretResponseBody
	GetRequestId() *string
}

type CreateApplicationClientSecretResponseBody struct {
	// The information about the client key.
	ApplicationClientSecret *CreateApplicationClientSecretResponseBodyApplicationClientSecret `json:"ApplicationClientSecret,omitempty" xml:"ApplicationClientSecret,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationClientSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretResponseBody) GetApplicationClientSecret() *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	return s.ApplicationClientSecret
}

func (s *CreateApplicationClientSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationClientSecretResponseBody) SetApplicationClientSecret(v *CreateApplicationClientSecretResponseBodyApplicationClientSecret) *CreateApplicationClientSecretResponseBody {
	s.ApplicationClientSecret = v
	return s
}

func (s *CreateApplicationClientSecretResponseBody) SetRequestId(v string) *CreateApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationClientSecretResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationClientSecretResponseBodyApplicationClientSecret struct {
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client key secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The client key ID of the application.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s CreateApplicationClientSecretResponseBodyApplicationClientSecret) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationClientSecretResponseBodyApplicationClientSecret) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) GetClientId() *string {
	return s.ClientId
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) GetSecretId() *string {
	return s.SecretId
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) SetClientId(v string) *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientId = &v
	return s
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) SetClientSecret(v string) *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientSecret = &v
	return s
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) SetSecretId(v string) *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	s.SecretId = &v
	return s
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) Validate() error {
	return dara.Validate(s)
}
