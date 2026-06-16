// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthnSourceMappingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *ListUserAuthnSourceMappingsRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *ListUserAuthnSourceMappingsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListUserAuthnSourceMappingsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserAuthnSourceMappingsRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListUserAuthnSourceMappingsRequest
	GetPreviousToken() *string
	SetUserExternalId(v string) *ListUserAuthnSourceMappingsRequest
	GetUserExternalId() *string
	SetUserId(v string) *ListUserAuthnSourceMappingsRequest
	GetUserId() *string
}

type ListUserAuthnSourceMappingsRequest struct {
	// The ID of the source identity provider (IdP).
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token for the previous page of results.
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The external ID.
	//
	// example:
	//
	// xxxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_ue2jvisn35exxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserAuthnSourceMappingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthnSourceMappingsRequest) GoString() string {
	return s.String()
}

func (s *ListUserAuthnSourceMappingsRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *ListUserAuthnSourceMappingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserAuthnSourceMappingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserAuthnSourceMappingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAuthnSourceMappingsRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListUserAuthnSourceMappingsRequest) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *ListUserAuthnSourceMappingsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListUserAuthnSourceMappingsRequest) SetIdentityProviderId(v string) *ListUserAuthnSourceMappingsRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) SetInstanceId(v string) *ListUserAuthnSourceMappingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) SetMaxResults(v int32) *ListUserAuthnSourceMappingsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) SetNextToken(v string) *ListUserAuthnSourceMappingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) SetPreviousToken(v string) *ListUserAuthnSourceMappingsRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) SetUserExternalId(v string) *ListUserAuthnSourceMappingsRequest {
	s.UserExternalId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) SetUserId(v string) *ListUserAuthnSourceMappingsRequest {
	s.UserId = &v
	return s
}

func (s *ListUserAuthnSourceMappingsRequest) Validate() error {
	return dara.Validate(s)
}
