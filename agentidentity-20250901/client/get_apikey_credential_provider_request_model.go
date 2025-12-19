// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAPIKeyCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKeyCredentialProviderName(v string) *GetAPIKeyCredentialProviderRequest
	GetAPIKeyCredentialProviderName() *string
}

type GetAPIKeyCredentialProviderRequest struct {
	// example:
	//
	// api-key-dash-scope
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
}

func (s GetAPIKeyCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAPIKeyCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *GetAPIKeyCredentialProviderRequest) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *GetAPIKeyCredentialProviderRequest) SetAPIKeyCredentialProviderName(v string) *GetAPIKeyCredentialProviderRequest {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *GetAPIKeyCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
