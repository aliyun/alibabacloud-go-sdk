// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshLoginTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *RefreshLoginTokenRequest
	GetClientId() *string
	SetClientType(v string) *RefreshLoginTokenRequest
	GetClientType() *string
	SetEndUserId(v string) *RefreshLoginTokenRequest
	GetEndUserId() *string
	SetLoginIdentifier(v string) *RefreshLoginTokenRequest
	GetLoginIdentifier() *string
	SetLoginToken(v string) *RefreshLoginTokenRequest
	GetLoginToken() *string
	SetOfficeSiteId(v string) *RefreshLoginTokenRequest
	GetOfficeSiteId() *string
	SetProfileRegion(v string) *RefreshLoginTokenRequest
	GetProfileRegion() *string
	SetSessionId(v string) *RefreshLoginTokenRequest
	GetSessionId() *string
	SetUuid(v string) *RefreshLoginTokenRequest
	GetUuid() *string
}

type RefreshLoginTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f4a0dc8e-1702-4728-9a60-95b27a35****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// windows
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// Alibaba****
	LoginIdentifier *string `json:"LoginIdentifier,omitempty" xml:"LoginIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1046341d8d4e2f05c4aa168196009613594aaf451499bfc75e54699efa7230bc968e1debb1fa4063b01e5d327b467****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-shenzhen+dir-436909****
	OfficeSiteId  *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	ProfileRegion *string `json:"ProfileRegion,omitempty" xml:"ProfileRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6442b2fd-ed3e-423a-8e6e-352d26a4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RefreshLoginTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *RefreshLoginTokenRequest) GetClientType() *string {
	return s.ClientType
}

func (s *RefreshLoginTokenRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RefreshLoginTokenRequest) GetLoginIdentifier() *string {
	return s.LoginIdentifier
}

func (s *RefreshLoginTokenRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *RefreshLoginTokenRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *RefreshLoginTokenRequest) GetProfileRegion() *string {
	return s.ProfileRegion
}

func (s *RefreshLoginTokenRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RefreshLoginTokenRequest) GetUuid() *string {
	return s.Uuid
}

func (s *RefreshLoginTokenRequest) SetClientId(v string) *RefreshLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetClientType(v string) *RefreshLoginTokenRequest {
	s.ClientType = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetEndUserId(v string) *RefreshLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginIdentifier(v string) *RefreshLoginTokenRequest {
	s.LoginIdentifier = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginToken(v string) *RefreshLoginTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetOfficeSiteId(v string) *RefreshLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetProfileRegion(v string) *RefreshLoginTokenRequest {
	s.ProfileRegion = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetSessionId(v string) *RefreshLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetUuid(v string) *RefreshLoginTokenRequest {
	s.Uuid = &v
	return s
}

func (s *RefreshLoginTokenRequest) Validate() error {
	return dara.Validate(s)
}
