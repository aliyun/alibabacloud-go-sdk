// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAPIKeyCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKeyCredentialProviderName(v string) *DeleteAPIKeyCredentialProviderRequest
	GetAPIKeyCredentialProviderName() *string
}

type DeleteAPIKeyCredentialProviderRequest struct {
	// example:
	//
	// api-key-dash-scope
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
}

func (s DeleteAPIKeyCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAPIKeyCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteAPIKeyCredentialProviderRequest) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *DeleteAPIKeyCredentialProviderRequest) SetAPIKeyCredentialProviderName(v string) *DeleteAPIKeyCredentialProviderRequest {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *DeleteAPIKeyCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
