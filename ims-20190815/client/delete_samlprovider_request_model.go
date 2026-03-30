// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSAMLProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSAMLProviderName(v string) *DeleteSAMLProviderRequest
	GetSAMLProviderName() *string
}

type DeleteSAMLProviderRequest struct {
	// The name of the IdP that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s DeleteSAMLProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderRequest) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *DeleteSAMLProviderRequest) SetSAMLProviderName(v string) *DeleteSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

func (s *DeleteSAMLProviderRequest) Validate() error {
	return dara.Validate(s)
}
