// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIds(v string) *UpdateOIDCProviderRequest
	GetClientIds() *string
	SetIssuanceLimitTime(v int64) *UpdateOIDCProviderRequest
	GetIssuanceLimitTime() *int64
	SetNewDescription(v string) *UpdateOIDCProviderRequest
	GetNewDescription() *string
	SetOIDCProviderName(v string) *UpdateOIDCProviderRequest
	GetOIDCProviderName() *string
}

type UpdateOIDCProviderRequest struct {
	ClientIds         *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	IssuanceLimitTime *int64  `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	NewDescription    *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	OIDCProviderName  *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s UpdateOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *UpdateOIDCProviderRequest) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *UpdateOIDCProviderRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *UpdateOIDCProviderRequest) SetClientIds(v string) *UpdateOIDCProviderRequest {
	s.ClientIds = &v
	return s
}

func (s *UpdateOIDCProviderRequest) SetIssuanceLimitTime(v int64) *UpdateOIDCProviderRequest {
	s.IssuanceLimitTime = &v
	return s
}

func (s *UpdateOIDCProviderRequest) SetNewDescription(v string) *UpdateOIDCProviderRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateOIDCProviderRequest) SetOIDCProviderName(v string) *UpdateOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *UpdateOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
