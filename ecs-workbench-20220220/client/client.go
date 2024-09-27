// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetInstanceRecordConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceRecordConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigRequest) SetInstanceId(v string) *GetInstanceRecordConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRecordConfigRequest) SetRegionId(v string) *GetInstanceRecordConfigRequest {
	s.RegionId = &v
	return s
}

type GetInstanceRecordConfigResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *GetInstanceRecordConfigResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceRecordConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigResponseBody) SetCode(v string) *GetInstanceRecordConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetMessage(v string) *GetInstanceRecordConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetRequestId(v string) *GetInstanceRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetRoot(v *GetInstanceRecordConfigResponseBodyRoot) *GetInstanceRecordConfigResponseBody {
	s.Root = v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetSuccess(v bool) *GetInstanceRecordConfigResponseBody {
	s.Success = &v
	return s
}

type GetInstanceRecordConfigResponseBodyRoot struct {
	// example:
	//
	// 7
	ExpirationDays *int32 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// example:
	//
	// i-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// acs:oss:cn-shanghai:123:workbench-record-123-1/record
	RecordStorageTarget *string `json:"RecordStorageTarget,omitempty" xml:"RecordStorageTarget,omitempty"`
}

func (s GetInstanceRecordConfigResponseBodyRoot) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRecordConfigResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetExpirationDays(v int32) *GetInstanceRecordConfigResponseBodyRoot {
	s.ExpirationDays = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetInstanceId(v string) *GetInstanceRecordConfigResponseBodyRoot {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetParentId(v string) *GetInstanceRecordConfigResponseBodyRoot {
	s.ParentId = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetRecordStorageTarget(v string) *GetInstanceRecordConfigResponseBodyRoot {
	s.RecordStorageTarget = &v
	return s
}

type GetInstanceRecordConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceRecordConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigResponse) SetHeaders(v map[string]*string) *GetInstanceRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceRecordConfigResponse) SetStatusCode(v int32) *GetInstanceRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceRecordConfigResponse) SetBody(v *GetInstanceRecordConfigResponseBody) *GetInstanceRecordConfigResponse {
	s.Body = v
	return s
}

type ListInstanceRecordsRequest struct {
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsRequest) SetInstanceId(v string) *ListInstanceRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRecordsRequest) SetPageNumber(v int32) *ListInstanceRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceRecordsRequest) SetPageSize(v int32) *ListInstanceRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRecordsRequest) SetRegionId(v string) *ListInstanceRecordsRequest {
	s.RegionId = &v
	return s
}

type ListInstanceRecordsResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *ListInstanceRecordsResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstanceRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponseBody) SetCode(v string) *ListInstanceRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetMessage(v string) *ListInstanceRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetRequestId(v string) *ListInstanceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetRoot(v *ListInstanceRecordsResponseBodyRoot) *ListInstanceRecordsResponseBody {
	s.Root = v
	return s
}

func (s *ListInstanceRecordsResponseBody) SetSuccess(v bool) *ListInstanceRecordsResponseBody {
	s.Success = &v
	return s
}

type ListInstanceRecordsResponseBodyRoot struct {
	RecordList []*ListInstanceRecordsResponseBodyRootRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceRecordsResponseBodyRoot) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRecordsResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponseBodyRoot) SetRecordList(v []*ListInstanceRecordsResponseBodyRootRecordList) *ListInstanceRecordsResponseBodyRoot {
	s.RecordList = v
	return s
}

func (s *ListInstanceRecordsResponseBodyRoot) SetTotalCount(v int32) *ListInstanceRecordsResponseBodyRoot {
	s.TotalCount = &v
	return s
}

type ListInstanceRecordsResponseBodyRootRecordList struct {
	// example:
	//
	// 1234
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 2023-11-16T02:59:39Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2023-04-10T12:41:28Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// i-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// https://ecs-workbench.aliyun.com/view/instance/record/replay/abc
	InstanceRecordUrl *string `json:"InstanceRecordUrl,omitempty" xml:"InstanceRecordUrl,omitempty"`
	// example:
	//
	// 123
	RecordDurationMillis *int64 `json:"RecordDurationMillis,omitempty" xml:"RecordDurationMillis,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// abc
	TerminalSessionToken *string `json:"TerminalSessionToken,omitempty" xml:"TerminalSessionToken,omitempty"`
}

func (s ListInstanceRecordsResponseBodyRootRecordList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRecordsResponseBodyRootRecordList) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetAccountId(v int64) *ListInstanceRecordsResponseBodyRootRecordList {
	s.AccountId = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetExpireTime(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.ExpireTime = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetGmtCreate(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.GmtCreate = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetInstanceId(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetInstanceRecordUrl(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.InstanceRecordUrl = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetRecordDurationMillis(v int64) *ListInstanceRecordsResponseBodyRootRecordList {
	s.RecordDurationMillis = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetStatus(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.Status = &v
	return s
}

func (s *ListInstanceRecordsResponseBodyRootRecordList) SetTerminalSessionToken(v string) *ListInstanceRecordsResponseBodyRootRecordList {
	s.TerminalSessionToken = &v
	return s
}

type ListInstanceRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponse) SetHeaders(v map[string]*string) *ListInstanceRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRecordsResponse) SetStatusCode(v int32) *ListInstanceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRecordsResponse) SetBody(v *ListInstanceRecordsResponseBody) *ListInstanceRecordsResponse {
	s.Body = v
	return s
}

type ListTerminalCommandsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	TerminalSessionToken *string `json:"TerminalSessionToken,omitempty" xml:"TerminalSessionToken,omitempty"`
}

func (s ListTerminalCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsRequest) SetPageNumber(v int32) *ListTerminalCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTerminalCommandsRequest) SetPageSize(v int32) *ListTerminalCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTerminalCommandsRequest) SetRegionId(v string) *ListTerminalCommandsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTerminalCommandsRequest) SetTerminalSessionToken(v string) *ListTerminalCommandsRequest {
	s.TerminalSessionToken = &v
	return s
}

type ListTerminalCommandsResponseBody struct {
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TerminalCommandList []*ListTerminalCommandsResponseBodyTerminalCommandList `json:"TerminalCommandList,omitempty" xml:"TerminalCommandList,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTerminalCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsResponseBody) SetPageNumber(v int32) *ListTerminalCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetPageSize(v int32) *ListTerminalCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetRequestId(v string) *ListTerminalCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetTerminalCommandList(v []*ListTerminalCommandsResponseBodyTerminalCommandList) *ListTerminalCommandsResponseBody {
	s.TerminalCommandList = v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetTotalCount(v int32) *ListTerminalCommandsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTerminalCommandsResponseBodyTerminalCommandList struct {
	// example:
	//
	// ls
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// example:
	//
	// 2024-04-16T03:53:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// /root
	ExecutePath *string `json:"ExecutePath,omitempty" xml:"ExecutePath,omitempty"`
	// example:
	//
	// root
	LoginUser *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
}

func (s ListTerminalCommandsResponseBodyTerminalCommandList) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalCommandsResponseBodyTerminalCommandList) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetCommandLine(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.CommandLine = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetCreateTime(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.CreateTime = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetExecutePath(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.ExecutePath = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetLoginUser(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.LoginUser = &v
	return s
}

type ListTerminalCommandsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerminalCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerminalCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTerminalCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsResponse) SetHeaders(v map[string]*string) *ListTerminalCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListTerminalCommandsResponse) SetStatusCode(v int32) *ListTerminalCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerminalCommandsResponse) SetBody(v *ListTerminalCommandsResponseBody) *ListTerminalCommandsResponse {
	s.Body = v
	return s
}

type LoginInstanceRequest struct {
	InstanceLoginInfo *LoginInstanceRequestInstanceLoginInfo `json:"InstanceLoginInfo,omitempty" xml:"InstanceLoginInfo,omitempty" type:"Struct"`
	PartnerInfo       *LoginInstanceRequestPartnerInfo       `json:"PartnerInfo,omitempty" xml:"PartnerInfo,omitempty" type:"Struct"`
	UserAccount       *LoginInstanceRequestUserAccount       `json:"UserAccount,omitempty" xml:"UserAccount,omitempty" type:"Struct"`
}

func (s LoginInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequest) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequest) SetInstanceLoginInfo(v *LoginInstanceRequestInstanceLoginInfo) *LoginInstanceRequest {
	s.InstanceLoginInfo = v
	return s
}

func (s *LoginInstanceRequest) SetPartnerInfo(v *LoginInstanceRequestPartnerInfo) *LoginInstanceRequest {
	s.PartnerInfo = v
	return s
}

func (s *LoginInstanceRequest) SetUserAccount(v *LoginInstanceRequestUserAccount) *LoginInstanceRequest {
	s.UserAccount = v
	return s
}

type LoginInstanceRequestInstanceLoginInfo struct {
	// example:
	//
	// password/certificate
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// example:
	//
	// ----begin----
	//
	// ----end----
	Certificate         *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CredentialToken     *string `json:"CredentialToken,omitempty" xml:"CredentialToken,omitempty"`
	DockerContainerName *string `json:"DockerContainerName,omitempty" xml:"DockerContainerName,omitempty"`
	DockerExec          *string `json:"DockerExec,omitempty" xml:"DockerExec,omitempty"`
	// example:
	//
	// 123
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// example:
	//
	// 2022-11-30 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 127.0.0.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// i-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ecs/eci/ack
	InstanceType              *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LoginByInstanceCredential *bool   `json:"LoginByInstanceCredential,omitempty" xml:"LoginByInstanceCredential,omitempty"`
	LoginByInstanceShortcut   *bool   `json:"LoginByInstanceShortcut,omitempty" xml:"LoginByInstanceShortcut,omitempty"`
	// example:
	//
	// vpc/classic
	NetworkAccessMode *string                                       `json:"NetworkAccessMode,omitempty" xml:"NetworkAccessMode,omitempty"`
	Options           *LoginInstanceRequestInstanceLoginInfoOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// example:
	//
	// xxxx
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// example:
	//
	// xxxxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 22/3389
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// ssh/rdp/ack
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// cn-hangzhou/cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ShortcutToken   *string `json:"ShortcutToken,omitempty" xml:"ShortcutToken,omitempty"`
	// example:
	//
	// root/Administrator
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// vpc-abc
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfo) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfo) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetAuthenticationType(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.AuthenticationType = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetCertificate(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Certificate = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetCredentialToken(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.CredentialToken = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetDockerContainerName(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.DockerContainerName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetDockerExec(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.DockerExec = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetDurationSeconds(v int64) *LoginInstanceRequestInstanceLoginInfo {
	s.DurationSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetExpireTime(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.ExpireTime = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetHost(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Host = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetInstanceId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetInstanceType(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.InstanceType = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetLoginByInstanceCredential(v bool) *LoginInstanceRequestInstanceLoginInfo {
	s.LoginByInstanceCredential = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetLoginByInstanceShortcut(v bool) *LoginInstanceRequestInstanceLoginInfo {
	s.LoginByInstanceShortcut = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetNetworkAccessMode(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.NetworkAccessMode = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetOptions(v *LoginInstanceRequestInstanceLoginInfoOptions) *LoginInstanceRequestInstanceLoginInfo {
	s.Options = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetPassPhrase(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.PassPhrase = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetPassword(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Password = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetPort(v int32) *LoginInstanceRequestInstanceLoginInfo {
	s.Port = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetProtocol(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Protocol = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetRegionId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.RegionId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetResourceGroupId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetShortcutToken(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.ShortcutToken = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetUsername(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Username = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetVpcId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.VpcId = &v
	return s
}

type LoginInstanceRequestInstanceLoginInfoOptions struct {
	AudioMuteSeconds *int32                                                     `json:"AudioMuteSeconds,omitempty" xml:"AudioMuteSeconds,omitempty"`
	ContainerInfo    *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo `json:"ContainerInfo,omitempty" xml:"ContainerInfo,omitempty" type:"Struct"`
	FixedHeight      *int32                                                     `json:"FixedHeight,omitempty" xml:"FixedHeight,omitempty"`
	FixedWidth       *int32                                                     `json:"FixedWidth,omitempty" xml:"FixedWidth,omitempty"`
	// example:
	//
	// abc
	NotificationEventTypes *string `json:"NotificationEventTypes,omitempty" xml:"NotificationEventTypes,omitempty"`
	// example:
	//
	// abc
	NotificationRecipientUrl *string `json:"NotificationRecipientUrl,omitempty" xml:"NotificationRecipientUrl,omitempty"`
	// example:
	//
	// 10
	NotificationRetryIntervalSeconds *int32 `json:"NotificationRetryIntervalSeconds,omitempty" xml:"NotificationRetryIntervalSeconds,omitempty"`
	// example:
	//
	// 3
	NotificationRetryLimit  *int32 `json:"NotificationRetryLimit,omitempty" xml:"NotificationRetryLimit,omitempty"`
	OperationDisableSeconds *int32 `json:"OperationDisableSeconds,omitempty" xml:"OperationDisableSeconds,omitempty"`
	// example:
	//
	// abc
	SessionControl     *string `json:"SessionControl,omitempty" xml:"SessionControl,omitempty"`
	VideoFreezeSeconds *int32  `json:"VideoFreezeSeconds,omitempty" xml:"VideoFreezeSeconds,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfoOptions) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfoOptions) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetAudioMuteSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.AudioMuteSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetContainerInfo(v *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.ContainerInfo = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetFixedHeight(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.FixedHeight = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetFixedWidth(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.FixedWidth = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationEventTypes(v string) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationEventTypes = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationRecipientUrl(v string) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationRecipientUrl = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationRetryIntervalSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationRetryIntervalSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationRetryLimit(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationRetryLimit = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetOperationDisableSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.OperationDisableSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetSessionControl(v string) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.SessionControl = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetVideoFreezeSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.VideoFreezeSeconds = &v
	return s
}

type LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo struct {
	// example:
	//
	// abcdef
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// abc
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// abc
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// example:
	//
	// abc
	Deployment *string `json:"Deployment,omitempty" xml:"Deployment,omitempty"`
	// example:
	//
	// abc
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// {"abc":"def"}
	Headers map[string]interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// abc
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// abc
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetClusterId(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.ClusterId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetClusterName(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.ClusterName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetContainerName(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.ContainerName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetDeployment(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Deployment = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetEndpoint(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Endpoint = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetHeaders(v map[string]interface{}) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Headers = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetNamespace(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Namespace = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetPodName(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.PodName = &v
	return s
}

type LoginInstanceRequestPartnerInfo struct {
	// example:
	//
	// abc
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// example:
	//
	// abc
	PartnerName *string `json:"PartnerName,omitempty" xml:"PartnerName,omitempty"`
}

func (s LoginInstanceRequestPartnerInfo) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequestPartnerInfo) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestPartnerInfo) SetPartnerId(v string) *LoginInstanceRequestPartnerInfo {
	s.PartnerId = &v
	return s
}

func (s *LoginInstanceRequestPartnerInfo) SetPartnerName(v string) *LoginInstanceRequestPartnerInfo {
	s.PartnerName = &v
	return s
}

type LoginInstanceRequestUserAccount struct {
	// example:
	//
	// 1234
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// aas
	AccountPlatform *string `json:"AccountPlatform,omitempty" xml:"AccountPlatform,omitempty"`
	// example:
	//
	// 2/3/4
	AccountStructure *string `json:"AccountStructure,omitempty" xml:"AccountStructure,omitempty"`
	// example:
	//
	// 100
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// example:
	//
	// 123abc
	EmpId *string `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	// example:
	//
	// 2022-11-30 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// abc
	LoginName *string                                 `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Options   *LoginInstanceRequestUserAccountOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s LoginInstanceRequestUserAccount) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequestUserAccount) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestUserAccount) SetAccountId(v int64) *LoginInstanceRequestUserAccount {
	s.AccountId = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetAccountPlatform(v string) *LoginInstanceRequestUserAccount {
	s.AccountPlatform = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetAccountStructure(v string) *LoginInstanceRequestUserAccount {
	s.AccountStructure = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetDurationSeconds(v int64) *LoginInstanceRequestUserAccount {
	s.DurationSeconds = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetEmpId(v string) *LoginInstanceRequestUserAccount {
	s.EmpId = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetExpireTime(v string) *LoginInstanceRequestUserAccount {
	s.ExpireTime = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetLoginName(v string) *LoginInstanceRequestUserAccount {
	s.LoginName = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetOptions(v *LoginInstanceRequestUserAccountOptions) *LoginInstanceRequestUserAccount {
	s.Options = v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetParentId(v int64) *LoginInstanceRequestUserAccount {
	s.ParentId = &v
	return s
}

type LoginInstanceRequestUserAccountOptions struct {
	// example:
	//
	// 3
	LoginLimit *int64 `json:"LoginLimit,omitempty" xml:"LoginLimit,omitempty"`
}

func (s LoginInstanceRequestUserAccountOptions) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequestUserAccountOptions) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestUserAccountOptions) SetLoginLimit(v int64) *LoginInstanceRequestUserAccountOptions {
	s.LoginLimit = &v
	return s
}

type LoginInstanceResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// abc-123
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *LoginInstanceResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true/false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LoginInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBody) SetCode(v string) *LoginInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *LoginInstanceResponseBody) SetMessage(v string) *LoginInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRequestId(v string) *LoginInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRoot(v *LoginInstanceResponseBodyRoot) *LoginInstanceResponseBody {
	s.Root = v
	return s
}

func (s *LoginInstanceResponseBody) SetSuccess(v string) *LoginInstanceResponseBody {
	s.Success = &v
	return s
}

type LoginInstanceResponseBodyRoot struct {
	DisposableAccount     *LoginInstanceResponseBodyRootDisposableAccount       `json:"DisposableAccount,omitempty" xml:"DisposableAccount,omitempty" type:"Struct"`
	InstanceLoginInfoList []*LoginInstanceResponseBodyRootInstanceLoginInfoList `json:"InstanceLoginInfoList,omitempty" xml:"InstanceLoginInfoList,omitempty" type:"Repeated"`
	SessionControl        *LoginInstanceResponseBodyRootSessionControl          `json:"SessionControl,omitempty" xml:"SessionControl,omitempty" type:"Struct"`
}

func (s LoginInstanceResponseBodyRoot) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRoot) SetDisposableAccount(v *LoginInstanceResponseBodyRootDisposableAccount) *LoginInstanceResponseBodyRoot {
	s.DisposableAccount = v
	return s
}

func (s *LoginInstanceResponseBodyRoot) SetInstanceLoginInfoList(v []*LoginInstanceResponseBodyRootInstanceLoginInfoList) *LoginInstanceResponseBodyRoot {
	s.InstanceLoginInfoList = v
	return s
}

func (s *LoginInstanceResponseBodyRoot) SetSessionControl(v *LoginInstanceResponseBodyRootSessionControl) *LoginInstanceResponseBodyRoot {
	s.SessionControl = v
	return s
}

type LoginInstanceResponseBodyRootDisposableAccount struct {
	// example:
	//
	// abc
	LoginFormActionUrl *string `json:"LoginFormActionUrl,omitempty" xml:"LoginFormActionUrl,omitempty"`
	// example:
	//
	// abc
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
}

func (s LoginInstanceResponseBodyRootDisposableAccount) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBodyRootDisposableAccount) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) SetLoginFormActionUrl(v string) *LoginInstanceResponseBodyRootDisposableAccount {
	s.LoginFormActionUrl = &v
	return s
}

func (s *LoginInstanceResponseBodyRootDisposableAccount) SetLoginUrl(v string) *LoginInstanceResponseBodyRootDisposableAccount {
	s.LoginUrl = &v
	return s
}

type LoginInstanceResponseBodyRootInstanceLoginInfoList struct {
	// example:
	//
	// i-abc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 134
	InstanceLoginToken *string                                                              `json:"InstanceLoginToken,omitempty" xml:"InstanceLoginToken,omitempty"`
	InstanceLoginView  *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView `json:"InstanceLoginView,omitempty" xml:"InstanceLoginView,omitempty" type:"Struct"`
	// example:
	//
	// true
	LoginSuccess *bool `json:"LoginSuccess,omitempty" xml:"LoginSuccess,omitempty"`
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoList) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoList) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetInstanceId(v string) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetInstanceLoginToken(v string) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.InstanceLoginToken = &v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetInstanceLoginView(v *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.InstanceLoginView = v
	return s
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoList) SetLoginSuccess(v bool) *LoginInstanceResponseBodyRootInstanceLoginInfoList {
	s.LoginSuccess = &v
	return s
}

type LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView struct {
	// example:
	//
	// abc
	DefaultViewUrl *string `json:"DefaultViewUrl,omitempty" xml:"DefaultViewUrl,omitempty"`
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView) SetDefaultViewUrl(v string) *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView {
	s.DefaultViewUrl = &v
	return s
}

type LoginInstanceResponseBodyRootSessionControl struct {
	// example:
	//
	// abc
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
}

func (s LoginInstanceResponseBodyRootSessionControl) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBodyRootSessionControl) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBodyRootSessionControl) SetBaseUrl(v string) *LoginInstanceResponseBodyRootSessionControl {
	s.BaseUrl = &v
	return s
}

type LoginInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponse) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponse) SetHeaders(v map[string]*string) *LoginInstanceResponse {
	s.Headers = v
	return s
}

func (s *LoginInstanceResponse) SetStatusCode(v int32) *LoginInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginInstanceResponse) SetBody(v *LoginInstanceResponseBody) *LoginInstanceResponse {
	s.Body = v
	return s
}

type SetInstanceRecordConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 7
	ExpirationDays *int32 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// acs:oss:cn-shanghai:123:workbench-record-123-1/record
	RecordStorageTarget *string `json:"RecordStorageTarget,omitempty" xml:"RecordStorageTarget,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetInstanceRecordConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceRecordConfigRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceRecordConfigRequest) SetEnabled(v bool) *SetInstanceRecordConfigRequest {
	s.Enabled = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetExpirationDays(v int32) *SetInstanceRecordConfigRequest {
	s.ExpirationDays = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetInstanceId(v string) *SetInstanceRecordConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetRecordStorageTarget(v string) *SetInstanceRecordConfigRequest {
	s.RecordStorageTarget = &v
	return s
}

func (s *SetInstanceRecordConfigRequest) SetRegionId(v string) *SetInstanceRecordConfigRequest {
	s.RegionId = &v
	return s
}

type SetInstanceRecordConfigResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Root *bool `json:"Root,omitempty" xml:"Root,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetInstanceRecordConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceRecordConfigResponseBody) SetCode(v string) *SetInstanceRecordConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetMessage(v string) *SetInstanceRecordConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetRequestId(v string) *SetInstanceRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetRoot(v bool) *SetInstanceRecordConfigResponseBody {
	s.Root = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetSuccess(v bool) *SetInstanceRecordConfigResponseBody {
	s.Success = &v
	return s
}

type SetInstanceRecordConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstanceRecordConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstanceRecordConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceRecordConfigResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceRecordConfigResponse) SetHeaders(v map[string]*string) *SetInstanceRecordConfigResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceRecordConfigResponse) SetStatusCode(v int32) *SetInstanceRecordConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceRecordConfigResponse) SetBody(v *SetInstanceRecordConfigResponseBody) *SetInstanceRecordConfigResponse {
	s.Body = v
	return s
}

type ViewInstanceRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	TerminalSessionToken *string `json:"TerminalSessionToken,omitempty" xml:"TerminalSessionToken,omitempty"`
}

func (s ViewInstanceRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ViewInstanceRecordsRequest) GoString() string {
	return s.String()
}

func (s *ViewInstanceRecordsRequest) SetInstanceId(v string) *ViewInstanceRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ViewInstanceRecordsRequest) SetRegionId(v string) *ViewInstanceRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *ViewInstanceRecordsRequest) SetTerminalSessionToken(v string) *ViewInstanceRecordsRequest {
	s.TerminalSessionToken = &v
	return s
}

type ViewInstanceRecordsResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Root *bool `json:"Root,omitempty" xml:"Root,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ViewInstanceRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ViewInstanceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ViewInstanceRecordsResponseBody) SetCode(v string) *ViewInstanceRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetMessage(v string) *ViewInstanceRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetRequestId(v string) *ViewInstanceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetRoot(v bool) *ViewInstanceRecordsResponseBody {
	s.Root = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetSuccess(v bool) *ViewInstanceRecordsResponseBody {
	s.Success = &v
	return s
}

type ViewInstanceRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewInstanceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewInstanceRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ViewInstanceRecordsResponse) GoString() string {
	return s.String()
}

func (s *ViewInstanceRecordsResponse) SetHeaders(v map[string]*string) *ViewInstanceRecordsResponse {
	s.Headers = v
	return s
}

func (s *ViewInstanceRecordsResponse) SetStatusCode(v int32) *ViewInstanceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewInstanceRecordsResponse) SetBody(v *ViewInstanceRecordsResponseBody) *ViewInstanceRecordsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecs-workbench"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例录屏配置
//
// @param request - GetInstanceRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceRecordConfigResponse
func (client *Client) GetInstanceRecordConfigWithOptions(request *GetInstanceRecordConfigRequest, runtime *util.RuntimeOptions) (_result *GetInstanceRecordConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceRecordConfig"),
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceRecordConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例录屏配置
//
// @param request - GetInstanceRecordConfigRequest
//
// @return GetInstanceRecordConfigResponse
func (client *Client) GetInstanceRecordConfig(request *GetInstanceRecordConfigRequest) (_result *GetInstanceRecordConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceRecordConfigResponse{}
	_body, _err := client.GetInstanceRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例录屏记录列表
//
// @param request - ListInstanceRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceRecordsResponse
func (client *Client) ListInstanceRecordsWithOptions(request *ListInstanceRecordsRequest, runtime *util.RuntimeOptions) (_result *ListInstanceRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceRecords"),
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例录屏记录列表
//
// @param request - ListInstanceRecordsRequest
//
// @return ListInstanceRecordsResponse
func (client *Client) ListInstanceRecords(request *ListInstanceRecordsRequest) (_result *ListInstanceRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceRecordsResponse{}
	_body, _err := client.ListInstanceRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看实例Workbench登录后执行命令的历史列表。
//
// @param request - ListTerminalCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalCommandsResponse
func (client *Client) ListTerminalCommandsWithOptions(request *ListTerminalCommandsRequest, runtime *util.RuntimeOptions) (_result *ListTerminalCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalSessionToken)) {
		body["TerminalSessionToken"] = request.TerminalSessionToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTerminalCommands"),
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTerminalCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看实例Workbench登录后执行命令的历史列表。
//
// @param request - ListTerminalCommandsRequest
//
// @return ListTerminalCommandsResponse
func (client *Client) ListTerminalCommands(request *ListTerminalCommandsRequest) (_result *ListTerminalCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTerminalCommandsResponse{}
	_body, _err := client.ListTerminalCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 登录实例
//
// @param request - LoginInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginInstanceResponse
func (client *Client) LoginInstanceWithOptions(request *LoginInstanceRequest, runtime *util.RuntimeOptions) (_result *LoginInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceLoginInfo)) {
		query["InstanceLoginInfo"] = request.InstanceLoginInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerInfo)) {
		query["PartnerInfo"] = request.PartnerInfo
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccount)) {
		query["UserAccount"] = request.UserAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LoginInstance"),
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoginInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 登录实例
//
// @param request - LoginInstanceRequest
//
// @return LoginInstanceResponse
func (client *Client) LoginInstance(request *LoginInstanceRequest) (_result *LoginInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoginInstanceResponse{}
	_body, _err := client.LoginInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置实例录屏配置
//
// @param request - SetInstanceRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstanceRecordConfigResponse
func (client *Client) SetInstanceRecordConfigWithOptions(request *SetInstanceRecordConfigRequest, runtime *util.RuntimeOptions) (_result *SetInstanceRecordConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		body["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.ExpirationDays)) {
		body["ExpirationDays"] = request.ExpirationDays
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RecordStorageTarget)) {
		body["RecordStorageTarget"] = request.RecordStorageTarget
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetInstanceRecordConfig"),
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInstanceRecordConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置实例录屏配置
//
// @param request - SetInstanceRecordConfigRequest
//
// @return SetInstanceRecordConfigResponse
func (client *Client) SetInstanceRecordConfig(request *SetInstanceRecordConfigRequest) (_result *SetInstanceRecordConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceRecordConfigResponse{}
	_body, _err := client.SetInstanceRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看实例录屏内容
//
// @param request - ViewInstanceRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewInstanceRecordsResponse
func (client *Client) ViewInstanceRecordsWithOptions(request *ViewInstanceRecordsRequest, runtime *util.RuntimeOptions) (_result *ViewInstanceRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalSessionToken)) {
		body["TerminalSessionToken"] = request.TerminalSessionToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ViewInstanceRecords"),
		Version:     tea.String("2022-02-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ViewInstanceRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看实例录屏内容
//
// @param request - ViewInstanceRecordsRequest
//
// @return ViewInstanceRecordsResponse
func (client *Client) ViewInstanceRecords(request *ViewInstanceRecordsRequest) (_result *ViewInstanceRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ViewInstanceRecordsResponse{}
	_body, _err := client.ViewInstanceRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
