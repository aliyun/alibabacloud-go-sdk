// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteClientSecretRequest
	GetClientId() *string
	SetClientName(v string) *DeleteClientSecretRequest
	GetClientName() *string
	SetClientSecretId(v string) *DeleteClientSecretRequest
	GetClientSecretId() *string
	SetUserPoolName(v string) *DeleteClientSecretRequest
	GetUserPoolName() *string
}

type DeleteClientSecretRequest struct {
	// example:
	//
	// client-xxxxxxxxxxxxxxxxxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// example:
	//
	// secret-xxxxxxxxxxxxxxxxxxxx
	ClientSecretId *string `json:"ClientSecretId,omitempty" xml:"ClientSecretId,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s DeleteClientSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientSecretRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteClientSecretRequest) GetClientName() *string {
	return s.ClientName
}

func (s *DeleteClientSecretRequest) GetClientSecretId() *string {
	return s.ClientSecretId
}

func (s *DeleteClientSecretRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteClientSecretRequest) SetClientId(v string) *DeleteClientSecretRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteClientSecretRequest) SetClientName(v string) *DeleteClientSecretRequest {
	s.ClientName = &v
	return s
}

func (s *DeleteClientSecretRequest) SetClientSecretId(v string) *DeleteClientSecretRequest {
	s.ClientSecretId = &v
	return s
}

func (s *DeleteClientSecretRequest) SetUserPoolName(v string) *DeleteClientSecretRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteClientSecretRequest) Validate() error {
	return dara.Validate(s)
}
