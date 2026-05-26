// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *CreateClientSecretRequest
	GetClientId() *string
	SetClientName(v string) *CreateClientSecretRequest
	GetClientName() *string
	SetUserPoolName(v string) *CreateClientSecretRequest
	GetUserPoolName() *string
}

type CreateClientSecretRequest struct {
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
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateClientSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateClientSecretRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CreateClientSecretRequest) GetClientName() *string {
	return s.ClientName
}

func (s *CreateClientSecretRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateClientSecretRequest) SetClientId(v string) *CreateClientSecretRequest {
	s.ClientId = &v
	return s
}

func (s *CreateClientSecretRequest) SetClientName(v string) *CreateClientSecretRequest {
	s.ClientName = &v
	return s
}

func (s *CreateClientSecretRequest) SetUserPoolName(v string) *CreateClientSecretRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
