// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudDriveServiceMountTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *GetCloudDriveServiceMountTokenRequest
	GetClientId() *string
	SetLoginToken(v string) *GetCloudDriveServiceMountTokenRequest
	GetLoginToken() *string
	SetOfficeSiteId(v string) *GetCloudDriveServiceMountTokenRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *GetCloudDriveServiceMountTokenRequest
	GetRegionId() *string
	SetSessionId(v string) *GetCloudDriveServiceMountTokenRequest
	GetSessionId() *string
}

type GetCloudDriveServiceMountTokenRequest struct {
	// example:
	//
	// 00e122c3-13fb-4fc3-bc7a-5d9acb89****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// v1972cd3446f0e523598916520951742474e6624fcdea6652994d47bc6157d27f7cc900c339db67882j3no4nh5bk3b4****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-7186763****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 14e1fe41-ce9b-491d-aa8c-345jk2n4bk****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetCloudDriveServiceMountTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenRequest) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetCloudDriveServiceMountTokenRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GetCloudDriveServiceMountTokenRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetCloudDriveServiceMountTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCloudDriveServiceMountTokenRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCloudDriveServiceMountTokenRequest) SetClientId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetLoginToken(v string) *GetCloudDriveServiceMountTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetOfficeSiteId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetRegionId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetSessionId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.SessionId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) Validate() error {
	return dara.Validate(s)
}
