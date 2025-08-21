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
	// The ID of the client. If you want to specify multiple client IDs, separate the client IDs with commas (,).
	//
	// A client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// A client ID can be up to 128 characters in length.
	//
	// >  If you specify this parameter, all the client IDs of the OIDC IdP are replaced. If you need to only add or remove a client ID, call the AddClientIdToOIDCProvider or RemoveClientIdFromOIDCProvider operation. For more information, see [AddClientIdToOIDCProvider](https://help.aliyun.com/document_detail/332057.html) or [RemoveClientIdFromOIDCProvider](https://help.aliyun.com/document_detail/332058.html).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The description of the OIDC IdP.
	//
	// The description can be up to 256 characters in length.
	//
	// example:
	//
	// This is a new OIDC Provider.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
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
