// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientName(v string) *ListClientSecretsRequest
	GetClientName() *string
	SetUserPoolName(v string) *ListClientSecretsRequest
	GetUserPoolName() *string
}

type ListClientSecretsRequest struct {
	// example:
	//
	// my-web-app
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListClientSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListClientSecretsRequest) GetClientName() *string {
	return s.ClientName
}

func (s *ListClientSecretsRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListClientSecretsRequest) SetClientName(v string) *ListClientSecretsRequest {
	s.ClientName = &v
	return s
}

func (s *ListClientSecretsRequest) SetUserPoolName(v string) *ListClientSecretsRequest {
	s.UserPoolName = &v
	return s
}

func (s *ListClientSecretsRequest) Validate() error {
	return dara.Validate(s)
}
