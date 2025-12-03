// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceAuthType(v string) *CreateServiceAuthRequest
	GetServiceAuthType() *string
}

type CreateServiceAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RAM
	ServiceAuthType *string `json:"serviceAuthType,omitempty" xml:"serviceAuthType,omitempty"`
}

func (s CreateServiceAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAuthRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceAuthRequest) GetServiceAuthType() *string {
	return s.ServiceAuthType
}

func (s *CreateServiceAuthRequest) SetServiceAuthType(v string) *CreateServiceAuthRequest {
	s.ServiceAuthType = &v
	return s
}

func (s *CreateServiceAuthRequest) Validate() error {
	return dara.Validate(s)
}
