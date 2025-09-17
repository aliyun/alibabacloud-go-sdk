// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *StopAppRequest
	GetAliUid() *int64
	SetApiType(v string) *StopAppRequest
	GetApiType() *string
	SetAppId(v string) *StopAppRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *StopAppRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *StopAppRequest
	GetAppInstanceId() *string
	SetBizRegionId(v string) *StopAppRequest
	GetBizRegionId() *string
	SetClientChannel(v string) *StopAppRequest
	GetClientChannel() *string
	SetClientId(v string) *StopAppRequest
	GetClientId() *string
	SetClientIp(v string) *StopAppRequest
	GetClientIp() *string
	SetClientOS(v string) *StopAppRequest
	GetClientOS() *string
	SetClientVersion(v string) *StopAppRequest
	GetClientVersion() *string
	SetEndUserId(v string) *StopAppRequest
	GetEndUserId() *string
	SetForceStop(v bool) *StopAppRequest
	GetForceStop() *bool
	SetIdpId(v string) *StopAppRequest
	GetIdpId() *string
	SetLoginRegionId(v string) *StopAppRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *StopAppRequest
	GetLoginToken() *string
	SetProductType(v string) *StopAppRequest
	GetProductType() *string
	SetRegionId(v string) *StopAppRequest
	GetRegionId() *string
	SetSessionId(v string) *StopAppRequest
	GetSessionId() *string
	SetUuid(v string) *StopAppRequest
	GetUuid() *string
	SetWyId(v string) *StopAppRequest
	GetWyId() *string
}

type StopAppRequest struct {
	// example:
	//
	// 1924794279035094
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// AnonymousUserAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// example:
	//
	// ca-fxwp4koywsglzvvex
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-89ibriac2wudyph38
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// ai-d297eyf83g5niwnjl
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// pc
	ClientChannel *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	// example:
	//
	// 91b79184-51d0-42ad-8475-78cae95b0aa6
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.79
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_"Windows 10 Enterprise" 10.0 (Build 19042)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 3.1.1-R-20211022.144255
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// test.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// example:
	//
	// idp-9ie5smicnct2xodn2
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v185fdd7f6d39fa7861981639366085772e150a390a5bb7b43c4e62440d94fc392b945770e1596cebe90085ce0af4d330e
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 597e869d-ea14-4b83-9490-714f68bfe935
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B26903
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// ac3cb49059261898
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s StopAppRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAppRequest) GoString() string {
	return s.String()
}

func (s *StopAppRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *StopAppRequest) GetApiType() *string {
	return s.ApiType
}

func (s *StopAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopAppRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *StopAppRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *StopAppRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *StopAppRequest) GetClientChannel() *string {
	return s.ClientChannel
}

func (s *StopAppRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StopAppRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *StopAppRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *StopAppRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *StopAppRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *StopAppRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopAppRequest) GetIdpId() *string {
	return s.IdpId
}

func (s *StopAppRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *StopAppRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *StopAppRequest) GetProductType() *string {
	return s.ProductType
}

func (s *StopAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopAppRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopAppRequest) GetUuid() *string {
	return s.Uuid
}

func (s *StopAppRequest) GetWyId() *string {
	return s.WyId
}

func (s *StopAppRequest) SetAliUid(v int64) *StopAppRequest {
	s.AliUid = &v
	return s
}

func (s *StopAppRequest) SetApiType(v string) *StopAppRequest {
	s.ApiType = &v
	return s
}

func (s *StopAppRequest) SetAppId(v string) *StopAppRequest {
	s.AppId = &v
	return s
}

func (s *StopAppRequest) SetAppInstanceGroupId(v string) *StopAppRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *StopAppRequest) SetAppInstanceId(v string) *StopAppRequest {
	s.AppInstanceId = &v
	return s
}

func (s *StopAppRequest) SetBizRegionId(v string) *StopAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *StopAppRequest) SetClientChannel(v string) *StopAppRequest {
	s.ClientChannel = &v
	return s
}

func (s *StopAppRequest) SetClientId(v string) *StopAppRequest {
	s.ClientId = &v
	return s
}

func (s *StopAppRequest) SetClientIp(v string) *StopAppRequest {
	s.ClientIp = &v
	return s
}

func (s *StopAppRequest) SetClientOS(v string) *StopAppRequest {
	s.ClientOS = &v
	return s
}

func (s *StopAppRequest) SetClientVersion(v string) *StopAppRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopAppRequest) SetEndUserId(v string) *StopAppRequest {
	s.EndUserId = &v
	return s
}

func (s *StopAppRequest) SetForceStop(v bool) *StopAppRequest {
	s.ForceStop = &v
	return s
}

func (s *StopAppRequest) SetIdpId(v string) *StopAppRequest {
	s.IdpId = &v
	return s
}

func (s *StopAppRequest) SetLoginRegionId(v string) *StopAppRequest {
	s.LoginRegionId = &v
	return s
}

func (s *StopAppRequest) SetLoginToken(v string) *StopAppRequest {
	s.LoginToken = &v
	return s
}

func (s *StopAppRequest) SetProductType(v string) *StopAppRequest {
	s.ProductType = &v
	return s
}

func (s *StopAppRequest) SetRegionId(v string) *StopAppRequest {
	s.RegionId = &v
	return s
}

func (s *StopAppRequest) SetSessionId(v string) *StopAppRequest {
	s.SessionId = &v
	return s
}

func (s *StopAppRequest) SetUuid(v string) *StopAppRequest {
	s.Uuid = &v
	return s
}

func (s *StopAppRequest) SetWyId(v string) *StopAppRequest {
	s.WyId = &v
	return s
}

func (s *StopAppRequest) Validate() error {
	return dara.Validate(s)
}
