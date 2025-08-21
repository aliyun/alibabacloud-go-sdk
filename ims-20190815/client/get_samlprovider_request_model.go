// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSAMLProviderName(v string) *GetSAMLProviderRequest
	GetSAMLProviderName() *string
}

type GetSAMLProviderRequest struct {
	// The name of the IdP.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s GetSAMLProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderRequest) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *GetSAMLProviderRequest) SetSAMLProviderName(v string) *GetSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

func (s *GetSAMLProviderRequest) Validate() error {
	return dara.Validate(s)
}
