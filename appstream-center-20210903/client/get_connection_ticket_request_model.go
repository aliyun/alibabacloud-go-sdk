// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *GetConnectionTicketRequest
	GetAccessType() *string
	SetAppId(v string) *GetConnectionTicketRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *GetConnectionTicketRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceId(v string) *GetConnectionTicketRequest
	GetAppInstanceId() *string
	SetAppPolicyId(v string) *GetConnectionTicketRequest
	GetAppPolicyId() *string
	SetAppVersion(v string) *GetConnectionTicketRequest
	GetAppVersion() *string
	SetAutoConnectInQueue(v bool) *GetConnectionTicketRequest
	GetAutoConnectInQueue() *bool
	SetBizRegionId(v string) *GetConnectionTicketRequest
	GetBizRegionId() *string
	SetClientId(v string) *GetConnectionTicketRequest
	GetClientId() *string
	SetClientIp(v string) *GetConnectionTicketRequest
	GetClientIp() *string
	SetClientOS(v string) *GetConnectionTicketRequest
	GetClientOS() *string
	SetClientType(v string) *GetConnectionTicketRequest
	GetClientType() *string
	SetClientVersion(v string) *GetConnectionTicketRequest
	GetClientVersion() *string
	SetConnectionProperties(v string) *GetConnectionTicketRequest
	GetConnectionProperties() *string
	SetEndUserId(v string) *GetConnectionTicketRequest
	GetEndUserId() *string
	SetEnvironmentConfig(v string) *GetConnectionTicketRequest
	GetEnvironmentConfig() *string
	SetLoginRegionId(v string) *GetConnectionTicketRequest
	GetLoginRegionId() *string
	SetLoginToken(v string) *GetConnectionTicketRequest
	GetLoginToken() *string
	SetParam(v string) *GetConnectionTicketRequest
	GetParam() *string
	SetProductType(v string) *GetConnectionTicketRequest
	GetProductType() *string
	SetResourceId(v string) *GetConnectionTicketRequest
	GetResourceId() *string
	SetSessionId(v string) *GetConnectionTicketRequest
	GetSessionId() *string
	SetTaskId(v string) *GetConnectionTicketRequest
	GetTaskId() *string
	SetTenantId(v string) *GetConnectionTicketRequest
	GetTenantId() *string
	SetUuid(v string) *GetConnectionTicketRequest
	GetUuid() *string
}

type GetConnectionTicketRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// ca-etn4zizgaezo9gis9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-bw1o1gcwvd3e1ipeu
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppPolicyId        *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// example:
	//
	// 1.0.0.1
	AppVersion         *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AutoConnectInQueue *bool   `json:"AutoConnectInQueue,omitempty" xml:"AutoConnectInQueue,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// f2463208-ec89-4309-8e8c-8b17acdcab93
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.21
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_"Windows 10 Enterprise LTSC 2019" 10.0 (Build 17763)
	ClientOS   *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion        *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// test.test
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EnvironmentConfig *string `json:"EnvironmentConfig,omitempty" xml:"EnvironmentConfig,omitempty"`
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// example:
	//
	// v1c4e2ef03d620f0f6cb41634196161219054e12d8aa5a13deb9ed14eebb76d674559115ad2e27a57f6820c1fd33e0ca36
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// /home/test/test.jpg
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 09e2b2e6-3181-4c84-9539-6fc9f1c3199e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 6f41731b-7091-4954-80c8-1d1e0b3ebb48
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1126819517152528
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// A8B35215993FBF283F28D617975204C4
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *GetConnectionTicketRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetConnectionTicketRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetConnectionTicketRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *GetConnectionTicketRequest) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *GetConnectionTicketRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetConnectionTicketRequest) GetAutoConnectInQueue() *bool {
	return s.AutoConnectInQueue
}

func (s *GetConnectionTicketRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetConnectionTicketRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetConnectionTicketRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetConnectionTicketRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetConnectionTicketRequest) GetClientType() *string {
	return s.ClientType
}

func (s *GetConnectionTicketRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetConnectionTicketRequest) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *GetConnectionTicketRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *GetConnectionTicketRequest) GetEnvironmentConfig() *string {
	return s.EnvironmentConfig
}

func (s *GetConnectionTicketRequest) GetLoginRegionId() *string {
	return s.LoginRegionId
}

func (s *GetConnectionTicketRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *GetConnectionTicketRequest) GetParam() *string {
	return s.Param
}

func (s *GetConnectionTicketRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetConnectionTicketRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetConnectionTicketRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetConnectionTicketRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetConnectionTicketRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetConnectionTicketRequest) SetAccessType(v string) *GetConnectionTicketRequest {
	s.AccessType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppId(v string) *GetConnectionTicketRequest {
	s.AppId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupId(v string) *GetConnectionTicketRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceId(v string) *GetConnectionTicketRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppPolicyId(v string) *GetConnectionTicketRequest {
	s.AppPolicyId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppVersion(v string) *GetConnectionTicketRequest {
	s.AppVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAutoConnectInQueue(v bool) *GetConnectionTicketRequest {
	s.AutoConnectInQueue = &v
	return s
}

func (s *GetConnectionTicketRequest) SetBizRegionId(v string) *GetConnectionTicketRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientId(v string) *GetConnectionTicketRequest {
	s.ClientId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientIp(v string) *GetConnectionTicketRequest {
	s.ClientIp = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientType(v string) *GetConnectionTicketRequest {
	s.ClientType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetConnectionProperties(v string) *GetConnectionTicketRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEnvironmentConfig(v string) *GetConnectionTicketRequest {
	s.EnvironmentConfig = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginRegionId(v string) *GetConnectionTicketRequest {
	s.LoginRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginToken(v string) *GetConnectionTicketRequest {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketRequest) SetParam(v string) *GetConnectionTicketRequest {
	s.Param = &v
	return s
}

func (s *GetConnectionTicketRequest) SetProductType(v string) *GetConnectionTicketRequest {
	s.ProductType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceId(v string) *GetConnectionTicketRequest {
	s.ResourceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetSessionId(v string) *GetConnectionTicketRequest {
	s.SessionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTenantId(v string) *GetConnectionTicketRequest {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

func (s *GetConnectionTicketRequest) Validate() error {
	return dara.Validate(s)
}
