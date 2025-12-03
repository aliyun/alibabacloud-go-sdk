// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCredentialType(v string) *ListServiceCredentialsRequest
	GetServiceCredentialType() *string
}

type ListServiceCredentialsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// username_password
	ServiceCredentialType *string `json:"serviceCredentialType,omitempty" xml:"serviceCredentialType,omitempty"`
}

func (s ListServiceCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceCredentialsRequest) GetServiceCredentialType() *string {
	return s.ServiceCredentialType
}

func (s *ListServiceCredentialsRequest) SetServiceCredentialType(v string) *ListServiceCredentialsRequest {
	s.ServiceCredentialType = &v
	return s
}

func (s *ListServiceCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
