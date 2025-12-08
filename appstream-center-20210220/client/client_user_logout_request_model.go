// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientUserLogoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *ClientUserLogoutRequest
	GetClientId() *string
	SetLoginToken(v string) *ClientUserLogoutRequest
	GetLoginToken() *string
	SetOfficeSiteId(v string) *ClientUserLogoutRequest
	GetOfficeSiteId() *string
	SetProfileRegion(v string) *ClientUserLogoutRequest
	GetProfileRegion() *string
	SetSessionId(v string) *ClientUserLogoutRequest
	GetSessionId() *string
}

type ClientUserLogoutRequest struct {
	// example:
	//
	// eac19bef-1e45-4190-a03a-4ea74b****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// v22369636c721ba6b3ddb1683341016775c3f63e4d0e78f120f9a0544ed826b7af7daf747c402f0d0730b52f451b70****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-hongkong+dir-643067****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn_hangzhou
	ProfileRegion *string `json:"ProfileRegion,omitempty" xml:"ProfileRegion,omitempty"`
	// example:
	//
	// 597e869d-ea14-4b83-9490-714f68****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ClientUserLogoutRequest) String() string {
	return dara.Prettify(s)
}

func (s ClientUserLogoutRequest) GoString() string {
	return s.String()
}

func (s *ClientUserLogoutRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ClientUserLogoutRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ClientUserLogoutRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ClientUserLogoutRequest) GetProfileRegion() *string {
	return s.ProfileRegion
}

func (s *ClientUserLogoutRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ClientUserLogoutRequest) SetClientId(v string) *ClientUserLogoutRequest {
	s.ClientId = &v
	return s
}

func (s *ClientUserLogoutRequest) SetLoginToken(v string) *ClientUserLogoutRequest {
	s.LoginToken = &v
	return s
}

func (s *ClientUserLogoutRequest) SetOfficeSiteId(v string) *ClientUserLogoutRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ClientUserLogoutRequest) SetProfileRegion(v string) *ClientUserLogoutRequest {
	s.ProfileRegion = &v
	return s
}

func (s *ClientUserLogoutRequest) SetSessionId(v string) *ClientUserLogoutRequest {
	s.SessionId = &v
	return s
}

func (s *ClientUserLogoutRequest) Validate() error {
	return dara.Validate(s)
}
