// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	AuthenticationType        *string                                       `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	Certificate               *string                                       `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CredentialToken           *string                                       `json:"CredentialToken,omitempty" xml:"CredentialToken,omitempty"`
	DockerContainerName       *string                                       `json:"DockerContainerName,omitempty" xml:"DockerContainerName,omitempty"`
	DockerExec                *string                                       `json:"DockerExec,omitempty" xml:"DockerExec,omitempty"`
	DurationSeconds           *int64                                        `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	ExpireTime                *string                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Host                      *string                                       `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceId                *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType              *string                                       `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LoginByInstanceCredential *bool                                         `json:"LoginByInstanceCredential,omitempty" xml:"LoginByInstanceCredential,omitempty"`
	LoginByInstanceShortcut   *bool                                         `json:"LoginByInstanceShortcut,omitempty" xml:"LoginByInstanceShortcut,omitempty"`
	NetworkAccessMode         *string                                       `json:"NetworkAccessMode,omitempty" xml:"NetworkAccessMode,omitempty"`
	Options                   *LoginInstanceRequestInstanceLoginInfoOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	PassPhrase                *string                                       `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	Password                  *string                                       `json:"Password,omitempty" xml:"Password,omitempty"`
	Port                      *int32                                        `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol                  *string                                       `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId                  *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId           *string                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ShortcutToken             *string                                       `json:"ShortcutToken,omitempty" xml:"ShortcutToken,omitempty"`
	Username                  *string                                       `json:"Username,omitempty" xml:"Username,omitempty"`
	VpcId                     *string                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	AudioMuteSeconds                 *int32                                                     `json:"AudioMuteSeconds,omitempty" xml:"AudioMuteSeconds,omitempty"`
	ContainerInfo                    *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo `json:"ContainerInfo,omitempty" xml:"ContainerInfo,omitempty" type:"Struct"`
	FixedHeight                      *int32                                                     `json:"FixedHeight,omitempty" xml:"FixedHeight,omitempty"`
	FixedWidth                       *int32                                                     `json:"FixedWidth,omitempty" xml:"FixedWidth,omitempty"`
	NotificationEventTypes           *string                                                    `json:"NotificationEventTypes,omitempty" xml:"NotificationEventTypes,omitempty"`
	NotificationRecipientUrl         *string                                                    `json:"NotificationRecipientUrl,omitempty" xml:"NotificationRecipientUrl,omitempty"`
	NotificationRetryIntervalSeconds *int32                                                     `json:"NotificationRetryIntervalSeconds,omitempty" xml:"NotificationRetryIntervalSeconds,omitempty"`
	NotificationRetryLimit           *int32                                                     `json:"NotificationRetryLimit,omitempty" xml:"NotificationRetryLimit,omitempty"`
	OperationDisableSeconds          *int32                                                     `json:"OperationDisableSeconds,omitempty" xml:"OperationDisableSeconds,omitempty"`
	SessionControl                   *string                                                    `json:"SessionControl,omitempty" xml:"SessionControl,omitempty"`
	VideoFreezeSeconds               *int32                                                     `json:"VideoFreezeSeconds,omitempty" xml:"VideoFreezeSeconds,omitempty"`
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
	ClusterId     *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName   *string                `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ContainerName *string                `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Deployment    *string                `json:"Deployment,omitempty" xml:"Deployment,omitempty"`
	Endpoint      *string                `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Headers       map[string]interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Namespace     *string                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PodName       *string                `json:"PodName,omitempty" xml:"PodName,omitempty"`
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
	PartnerId   *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
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
	AccountId        *int64                                  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountPlatform  *string                                 `json:"AccountPlatform,omitempty" xml:"AccountPlatform,omitempty"`
	AccountStructure *string                                 `json:"AccountStructure,omitempty" xml:"AccountStructure,omitempty"`
	DurationSeconds  *int64                                  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	EmpId            *string                                 `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	ExpireTime       *string                                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LoginName        *string                                 `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Options          *LoginInstanceRequestUserAccountOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	ParentId         *int64                                  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *LoginInstanceResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	Success   *string                        `json:"Success,omitempty" xml:"Success,omitempty"`
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
	LoginFormActionUrl *string `json:"LoginFormActionUrl,omitempty" xml:"LoginFormActionUrl,omitempty"`
	LoginUrl           *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
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
	InstanceId         *string                                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceLoginToken *string                                                              `json:"InstanceLoginToken,omitempty" xml:"InstanceLoginToken,omitempty"`
	InstanceLoginView  *LoginInstanceResponseBodyRootInstanceLoginInfoListInstanceLoginView `json:"InstanceLoginView,omitempty" xml:"InstanceLoginView,omitempty" type:"Struct"`
	LoginSuccess       *bool                                                                `json:"LoginSuccess,omitempty" xml:"LoginSuccess,omitempty"`
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
