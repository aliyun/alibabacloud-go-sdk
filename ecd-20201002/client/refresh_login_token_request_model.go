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
	SetDirectoryId(v string) *RefreshLoginTokenRequest
	GetDirectoryId() *string
	SetEndUserId(v string) *RefreshLoginTokenRequest
	GetEndUserId() *string
	SetLoginToken(v string) *RefreshLoginTokenRequest
	GetLoginToken() *string
	SetOfficeSiteId(v string) *RefreshLoginTokenRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *RefreshLoginTokenRequest
	GetRegionId() *string
	SetSessionId(v string) *RefreshLoginTokenRequest
	GetSessionId() *string
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
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-shanghai+dir-238191****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cd45e873-650d-4d70-acb9-f996187a****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
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

func (s *RefreshLoginTokenRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RefreshLoginTokenRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RefreshLoginTokenRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *RefreshLoginTokenRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *RefreshLoginTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RefreshLoginTokenRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RefreshLoginTokenRequest) SetClientId(v string) *RefreshLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetDirectoryId(v string) *RefreshLoginTokenRequest {
	s.DirectoryId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetEndUserId(v string) *RefreshLoginTokenRequest {
	s.EndUserId = &v
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

func (s *RefreshLoginTokenRequest) SetRegionId(v string) *RefreshLoginTokenRequest {
	s.RegionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetSessionId(v string) *RefreshLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) Validate() error {
	return dara.Validate(s)
}
