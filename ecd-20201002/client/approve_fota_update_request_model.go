// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveFotaUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *ApproveFotaUpdateRequest
	GetAppVersion() *string
	SetClientId(v string) *ApproveFotaUpdateRequest
	GetClientId() *string
	SetDesktopId(v string) *ApproveFotaUpdateRequest
	GetDesktopId() *string
	SetLoginToken(v string) *ApproveFotaUpdateRequest
	GetLoginToken() *string
	SetRegionId(v string) *ApproveFotaUpdateRequest
	GetRegionId() *string
	SetSessionId(v string) *ApproveFotaUpdateRequest
	GetSessionId() *string
	SetTargetStatus(v string) *ApproveFotaUpdateRequest
	GetTargetStatus() *string
	SetUuid(v string) *ApproveFotaUpdateRequest
	GetUuid() *string
}

type ApproveFotaUpdateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1-D-20220513.14****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d4452bd7-88df-4b90-a409-806da674****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// ecd-138dsptkrt00u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v18390c954ce59e2915ef16663205371721e0db9a46179ac89249a203320459523cb54ad08efe324784dd0eba25950****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 4771b873-c757-4893-973c-7f5beejh****
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TargetStatus *string `json:"TargetStatus,omitempty" xml:"TargetStatus,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApproveFotaUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveFotaUpdateRequest) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ApproveFotaUpdateRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ApproveFotaUpdateRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ApproveFotaUpdateRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ApproveFotaUpdateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApproveFotaUpdateRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ApproveFotaUpdateRequest) GetTargetStatus() *string {
	return s.TargetStatus
}

func (s *ApproveFotaUpdateRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ApproveFotaUpdateRequest) SetAppVersion(v string) *ApproveFotaUpdateRequest {
	s.AppVersion = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetClientId(v string) *ApproveFotaUpdateRequest {
	s.ClientId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetDesktopId(v string) *ApproveFotaUpdateRequest {
	s.DesktopId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetLoginToken(v string) *ApproveFotaUpdateRequest {
	s.LoginToken = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetRegionId(v string) *ApproveFotaUpdateRequest {
	s.RegionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetSessionId(v string) *ApproveFotaUpdateRequest {
	s.SessionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetTargetStatus(v string) *ApproveFotaUpdateRequest {
	s.TargetStatus = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetUuid(v string) *ApproveFotaUpdateRequest {
	s.Uuid = &v
	return s
}

func (s *ApproveFotaUpdateRequest) Validate() error {
	return dara.Validate(s)
}
