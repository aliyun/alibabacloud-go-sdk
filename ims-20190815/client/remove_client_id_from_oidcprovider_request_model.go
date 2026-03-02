// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientIdFromOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *RemoveClientIdFromOIDCProviderRequest
	GetClientId() *string
	SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderRequest
	GetOIDCProviderName() *string
}

type RemoveClientIdFromOIDCProviderRequest struct {
	ClientId         *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RemoveClientIdFromOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *RemoveClientIdFromOIDCProviderRequest) SetClientId(v string) *RemoveClientIdFromOIDCProviderRequest {
	s.ClientId = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderRequest) SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
