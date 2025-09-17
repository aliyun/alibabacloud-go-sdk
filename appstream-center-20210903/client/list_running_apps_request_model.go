// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunningAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ListRunningAppsRequest
	GetBizRegionId() *string
	SetClientId(v string) *ListRunningAppsRequest
	GetClientId() *string
	SetClientIp(v string) *ListRunningAppsRequest
	GetClientIp() *string
	SetClientOS(v string) *ListRunningAppsRequest
	GetClientOS() *string
	SetClientVersion(v string) *ListRunningAppsRequest
	GetClientVersion() *string
	SetEndUserId(v string) *ListRunningAppsRequest
	GetEndUserId() *string
	SetLoginRegionId(v string) *ListRunningAppsRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *ListRunningAppsRequest
	GetLoginToken() *string
	SetProductType(v string) *ListRunningAppsRequest
	GetProductType() *string
	SetSessionId(v string) *ListRunningAppsRequest
	GetSessionId() *string
	SetTenantId(v string) *ListRunningAppsRequest
	GetTenantId() *string
	SetUuid(v string) *ListRunningAppsRequest
	GetUuid() *string
}

type ListRunningAppsRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 370b56f8-2812-4b6c-bfa6-2560791cad88
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.32
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_"Windows 10 Enterprise" 10.0 (Build 18363)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 0.1.0-R-20220512.175656
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v1124500957832f30b3e716406562071655aa43b2a723ed2be0837815483d54e025db13ba5469f06f2410d0efc4d302e36
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// a863f4c3-2f1d-4971-8cf7-e2b92ae97764
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1735953493960828
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 470E8C12AB78CE9C3F6627DD0409E51D
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListRunningAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRunningAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRunningAppsRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListRunningAppsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ListRunningAppsRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListRunningAppsRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *ListRunningAppsRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *ListRunningAppsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListRunningAppsRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *ListRunningAppsRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ListRunningAppsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListRunningAppsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListRunningAppsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListRunningAppsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListRunningAppsRequest) SetBizRegionId(v string) *ListRunningAppsRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientId(v string) *ListRunningAppsRequest {
	s.ClientId = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientIp(v string) *ListRunningAppsRequest {
	s.ClientIp = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientOS(v string) *ListRunningAppsRequest {
	s.ClientOS = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientVersion(v string) *ListRunningAppsRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListRunningAppsRequest) SetEndUserId(v string) *ListRunningAppsRequest {
	s.EndUserId = &v
	return s
}

func (s *ListRunningAppsRequest) SetLoginRegionId(v string) *ListRunningAppsRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetLoginToken(v string) *ListRunningAppsRequest {
	s.LoginToken = &v
	return s
}

func (s *ListRunningAppsRequest) SetProductType(v string) *ListRunningAppsRequest {
	s.ProductType = &v
	return s
}

func (s *ListRunningAppsRequest) SetSessionId(v string) *ListRunningAppsRequest {
	s.SessionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetTenantId(v string) *ListRunningAppsRequest {
	s.TenantId = &v
	return s
}

func (s *ListRunningAppsRequest) SetUuid(v string) *ListRunningAppsRequest {
	s.Uuid = &v
	return s
}

func (s *ListRunningAppsRequest) Validate() error {
	return dara.Validate(s)
}
