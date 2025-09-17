// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UnbindRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *UnbindRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *UnbindRequest
	GetAppInstanceId() *string
	SetClientId(v string) *UnbindRequest
	GetClientId() *string
	SetClientIp(v string) *UnbindRequest
	GetClientIp() *string
	SetClientOS(v string) *UnbindRequest
	GetClientOS() *string
	SetClientVersion(v string) *UnbindRequest
	GetClientVersion() *string
	SetEndUserId(v string) *UnbindRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *UnbindRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *UnbindRequest
	GetLoginToken() *string
	SetProductType(v string) *UnbindRequest
	GetProductType() *string
	SetSessionId(v string) *UnbindRequest
	GetSessionId() *string
	SetTenantId(v int64) *UnbindRequest
	GetTenantId() *int64
}

type UnbindRequest struct {
	// example:
	//
	// ca-fxwp4koxs8hopi94e
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aig-e1l4kqqykxt4uzdx9
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// eac19bef-1e45-4190-a03a-4ea74b699ca7
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 139.129.223.122
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Pro\\" 10.0 (Build 19041)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 2.0.1-D-20220303.171122
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-shanghai
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1e9c8e83d83ea11270871640059145702bde8c5be8c6b9a854ffb6a43bd2673c19a5551c83800724e024f488dbfb0b247
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11040139-4fb4-4b35-9b44-6c07c746a43e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1569416393841402
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UnbindRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) GetAppId() *string {
	return s.AppId
}

func (s *UnbindRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *UnbindRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *UnbindRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UnbindRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *UnbindRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *UnbindRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *UnbindRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *UnbindRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *UnbindRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *UnbindRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UnbindRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UnbindRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *UnbindRequest) SetAppId(v string) *UnbindRequest {
	s.AppId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceGroupId(v string) *UnbindRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceId(v string) *UnbindRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindRequest) SetClientId(v string) *UnbindRequest {
	s.ClientId = &v
	return s
}

func (s *UnbindRequest) SetClientIp(v string) *UnbindRequest {
	s.ClientIp = &v
	return s
}

func (s *UnbindRequest) SetClientOS(v string) *UnbindRequest {
	s.ClientOS = &v
	return s
}

func (s *UnbindRequest) SetClientVersion(v string) *UnbindRequest {
	s.ClientVersion = &v
	return s
}

func (s *UnbindRequest) SetEndUserId(v string) *UnbindRequest {
	s.EndUserId = &v
	return s
}

func (s *UnbindRequest) SetLoginRegionId(v string) *UnbindRequest {
	s.LoginRegionId = &v
	return s
}

func (s *UnbindRequest) SetLoginToken(v string) *UnbindRequest {
	s.LoginToken = &v
	return s
}

func (s *UnbindRequest) SetProductType(v string) *UnbindRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindRequest) SetSessionId(v string) *UnbindRequest {
	s.SessionId = &v
	return s
}

func (s *UnbindRequest) SetTenantId(v int64) *UnbindRequest {
	s.TenantId = &v
	return s
}

func (s *UnbindRequest) Validate() error {
	return dara.Validate(s)
}
