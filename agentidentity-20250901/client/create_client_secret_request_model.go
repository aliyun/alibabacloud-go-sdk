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
	// The ID of the client. If you specify ClientId, you do not need to specify UserPoolName. Specify either ClientName or ClientId. If both are specified, ClientId takes precedence.
	//
	// example:
	//
	// client_xxxxxxxxxxxxxxxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The name of the client. Specify either ClientName or ClientId.
	//
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// The name of the user pool.
	//
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
