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
	SetRequireUuidValidation(v bool) *GetConnectionTicketRequest
	GetRequireUuidValidation() *bool
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
	// The access type. If this parameter is not passed, both types are displayed. Valid values: INTERNET (Internet) and VPC (Express Connect circuit).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The application ID.
	//
	// > Required for the first call, and optional for subsequent calls.
	//
	// example:
	//
	// ca-etn4zizga******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID.
	//
	// >
	//
	// > - When this parameter is passed, an application instance is assigned only from the specified and authorized delivery groups.
	//
	// > - When the AppInstanceId parameter is passed, this parameter is required.
	//
	// example:
	//
	// aig-bw1o1gcwv******
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The application instance ID.
	//
	// >
	//
	// > - When this parameter is passed, the system only attempts to assign the specified application instance.
	//
	// > - When this parameter is passed, the AppInstanceGroupId parameter must also be passed.
	//
	// example:
	//
	// ai-ij4a6kd4bn2******
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-0clfzcy0adpcf****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// The application version number.
	//
	// example:
	//
	// 1.0.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// Specifies whether to automatically connect when the turn is reached.
	//
	// > This parameter is not available.
	//
	// example:
	//
	// true
	AutoConnectInQueue *bool `json:"AutoConnectInQueue,omitempty" xml:"AutoConnectInQueue,omitempty"`
	// The region ID.
	//
	// > When this parameter is passed, an application instance is assigned only from the delivery groups in the specified region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The client ID.
	//
	// example:
	//
	// f2463208-ec89-4309-8e8c-8b1******
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 22.21.*.*
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The client operating system.
	//
	// example:
	//
	// windows_"Windows 10 Enterprise LTSC 2019" 10.0 (Build 17763)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client type.
	//
	// example:
	//
	// html5
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The client version.
	//
	// example:
	//
	// 2.0.1-D-20211008.101607
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The connection properties in JSON format.
	//
	// The following parameters are supported:
	//
	// - authMode: the authorization mode, which is determined based on the authorization mode of the delivery group. Valid values: App (application authorization), AppInstanceGroup (delivery group authorization), and Session (persistent session authorization).
	//
	// - appInstanceGroupScheduleTags: the delivery group scheduling tags. Multiple tag key-value pairs are supported for filtering. Multiple tag pairs have an AND relationship. The key must have a value, and the value can be empty. Format example: [{"key":"Resolution","value":"720p"},{"key":"Name"}].
	//
	// - appInstanceGroupIds: the list of specified delivery group IDs.
	//
	// - appInstancePersistentId: the specified persistent session ID.
	//
	// - avatarId: the user avatar ID. Default value: default. This parameter is used to support assigning multiple session instances to a single user.
	//
	// > Used when connecting to a cloud phone. Pass the fixed value `{"authMode":"Session"}`.
	//
	// example:
	//
	// {"authMode":"Session"}
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The username of the convenience account.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The environment configuration.
	//
	// example:
	//
	// {"screenWidth":2560,"screenHeight":1564}
	EnvironmentConfig *string `json:"EnvironmentConfig,omitempty" xml:"EnvironmentConfig,omitempty"`
	// The logon region ID. When a user logs on to a workspace, this parameter specifies the ID of the region to which the workspace belongs.
	//
	// example:
	//
	// cn-hangzhou
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	// The logon token.
	//
	// Obtained by calling the GetLoginToken or GetStsToken API operation.
	//
	// The GetLoginToken and GetStsToken API operations belong to the client logon authentication APIs of the 2021-02-20 version of WUYING Cloud App (product code: appstream-center). These two API operations are not provided in the 2021-09-03 version. OpenAPI Explorer links: https://api.aliyun.com/api/appstream-center/2021-02-20/GetLoginToken and https://api.aliyun.com/api/appstream-center/2021-02-20/GetStsToken.
	//
	// Note: This parameter is required. If it is not provided, the API returns InvalidParameter.loginToken(400).
	//
	// example:
	//
	// v2c4e2ef03d62******
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The application startup parameters.
	//
	// example:
	//
	// /q /n
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Specifies whether UUID validation is required.
	//
	// example:
	//
	// true
	RequireUuidValidation *bool `json:"RequireUuidValidation,omitempty" xml:"RequireUuidValidation,omitempty"`
	// The resource ID.
	//
	// > Used when connecting to a cloud phone. Pass the persistent session ID that corresponds to the cloud phone instance.
	//
	// > When connecting to a cloud browser, pass the cloud browser ID.
	//
	// > When connecting to a cloud application, pass the application ID for the application authorization scenario, the delivery group ID for the delivery group authorization scenario, or the persistent session ID for the persistent session authorization scenario.
	//
	// example:
	//
	// p-0cdr5c9kl******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The logon session ID. Obtained by calling the GetLoginToken or GetStsToken API operation. This parameter is used together with LoginToken.
	//
	// The GetLoginToken and GetStsToken API operations belong to the client logon authentication APIs of the 2021-02-20 version of WUYING Cloud App (product code: appstream-center). These two API operations are not provided in the 2021-09-03 version. OpenAPI Explorer links: https://api.aliyun.com/api/appstream-center/2021-02-20/GetLoginToken and https://api.aliyun.com/api/appstream-center/2021-02-20/GetStsToken.
	//
	// Note: This parameter is required. If it is not provided, the API returns InvalidParameter.sessionId(400). This parameter must be used together with LoginToken.
	//
	// example:
	//
	// 09e2b2e6-3181******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The node ID.
	//
	// > Required when you invoke the API not for the first time. Use this parameter query to check the node status and connection credentials.
	//
	// example:
	//
	// 6f41731b-7091-4******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The tenant ID, which is the Alibaba Cloud account UID.
	//
	// example:
	//
	// 1126819517******
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The unique identifier of the logon device.
	//
	// example:
	//
	// A8B35215993FBF283F28D61******
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

func (s *GetConnectionTicketRequest) GetRequireUuidValidation() *bool {
	return s.RequireUuidValidation
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

func (s *GetConnectionTicketRequest) SetRequireUuidValidation(v bool) *GetConnectionTicketRequest {
	s.RequireUuidValidation = &v
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
