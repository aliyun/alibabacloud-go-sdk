// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientSecret(v *CreateClientSecretResponseBodyClientSecret) *CreateClientSecretResponseBody
	GetClientSecret() *CreateClientSecretResponseBodyClientSecret
	SetRequestId(v string) *CreateClientSecretResponseBody
	GetRequestId() *string
}

type CreateClientSecretResponseBody struct {
	ClientSecret *CreateClientSecretResponseBodyClientSecret `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty" type:"Struct"`
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClientSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientSecretResponseBody) GetClientSecret() *CreateClientSecretResponseBodyClientSecret {
	return s.ClientSecret
}

func (s *CreateClientSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientSecretResponseBody) SetClientSecret(v *CreateClientSecretResponseBodyClientSecret) *CreateClientSecretResponseBody {
	s.ClientSecret = v
	return s
}

func (s *CreateClientSecretResponseBody) SetRequestId(v string) *CreateClientSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientSecretResponseBody) Validate() error {
	if s.ClientSecret != nil {
		if err := s.ClientSecret.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClientSecretResponseBodyClientSecret struct {
	// example:
	//
	// client_xxxxxxxxxxxxxxxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// example:
	//
	// secret_xxxxxxxxxxxxxxxxxxxx
	ClientSecretId *string `json:"ClientSecretId,omitempty" xml:"ClientSecretId,omitempty"`
	// example:
	//
	// xxxxxxxxxxxxxxxx
	ClientSecretValue *string `json:"ClientSecretValue,omitempty" xml:"ClientSecretValue,omitempty"`
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateClientSecretResponseBodyClientSecret) String() string {
	return dara.Prettify(s)
}

func (s CreateClientSecretResponseBodyClientSecret) GoString() string {
	return s.String()
}

func (s *CreateClientSecretResponseBodyClientSecret) GetClientId() *string {
	return s.ClientId
}

func (s *CreateClientSecretResponseBodyClientSecret) GetClientName() *string {
	return s.ClientName
}

func (s *CreateClientSecretResponseBodyClientSecret) GetClientSecretId() *string {
	return s.ClientSecretId
}

func (s *CreateClientSecretResponseBodyClientSecret) GetClientSecretValue() *string {
	return s.ClientSecretValue
}

func (s *CreateClientSecretResponseBodyClientSecret) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateClientSecretResponseBodyClientSecret) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateClientSecretResponseBodyClientSecret) SetClientId(v string) *CreateClientSecretResponseBodyClientSecret {
	s.ClientId = &v
	return s
}

func (s *CreateClientSecretResponseBodyClientSecret) SetClientName(v string) *CreateClientSecretResponseBodyClientSecret {
	s.ClientName = &v
	return s
}

func (s *CreateClientSecretResponseBodyClientSecret) SetClientSecretId(v string) *CreateClientSecretResponseBodyClientSecret {
	s.ClientSecretId = &v
	return s
}

func (s *CreateClientSecretResponseBodyClientSecret) SetClientSecretValue(v string) *CreateClientSecretResponseBodyClientSecret {
	s.ClientSecretValue = &v
	return s
}

func (s *CreateClientSecretResponseBodyClientSecret) SetCreateTime(v string) *CreateClientSecretResponseBodyClientSecret {
	s.CreateTime = &v
	return s
}

func (s *CreateClientSecretResponseBodyClientSecret) SetUserPoolName(v string) *CreateClientSecretResponseBodyClientSecret {
	s.UserPoolName = &v
	return s
}

func (s *CreateClientSecretResponseBodyClientSecret) Validate() error {
	return dara.Validate(s)
}
