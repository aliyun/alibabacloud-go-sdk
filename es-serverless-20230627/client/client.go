// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAppRequest struct {
	// 应用名
	//
	// This parameter is required.
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// This parameter is required.
	Authentication *CreateAppRequestAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// This parameter is required.
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 应用备注
	Description    *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Network        []*CreateAppRequestNetwork        `json:"network,omitempty" xml:"network,omitempty" type:"Repeated"`
	PrivateNetwork []*CreateAppRequestPrivateNetwork `json:"privateNetwork,omitempty" xml:"privateNetwork,omitempty" type:"Repeated"`
	QuotaInfo      *CreateAppRequestQuotaInfo        `json:"quotaInfo,omitempty" xml:"quotaInfo,omitempty" type:"Struct"`
	RegionId       *string                           `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Scenario       *string                           `json:"scenario,omitempty" xml:"scenario,omitempty"`
	Version        *string                           `json:"version,omitempty" xml:"version,omitempty"`
	DryRun         *bool                             `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAuthentication(v *CreateAppRequestAuthentication) *CreateAppRequest {
	s.Authentication = v
	return s
}

func (s *CreateAppRequest) SetChargeType(v string) *CreateAppRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetNetwork(v []*CreateAppRequestNetwork) *CreateAppRequest {
	s.Network = v
	return s
}

func (s *CreateAppRequest) SetPrivateNetwork(v []*CreateAppRequestPrivateNetwork) *CreateAppRequest {
	s.PrivateNetwork = v
	return s
}

func (s *CreateAppRequest) SetQuotaInfo(v *CreateAppRequestQuotaInfo) *CreateAppRequest {
	s.QuotaInfo = v
	return s
}

func (s *CreateAppRequest) SetRegionId(v string) *CreateAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppRequest) SetScenario(v string) *CreateAppRequest {
	s.Scenario = &v
	return s
}

func (s *CreateAppRequest) SetVersion(v string) *CreateAppRequest {
	s.Version = &v
	return s
}

func (s *CreateAppRequest) SetDryRun(v bool) *CreateAppRequest {
	s.DryRun = &v
	return s
}

type CreateAppRequestAuthentication struct {
	BasicAuth []*CreateAppRequestAuthenticationBasicAuth `json:"basicAuth,omitempty" xml:"basicAuth,omitempty" type:"Repeated"`
}

func (s CreateAppRequestAuthentication) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestAuthentication) GoString() string {
	return s.String()
}

func (s *CreateAppRequestAuthentication) SetBasicAuth(v []*CreateAppRequestAuthenticationBasicAuth) *CreateAppRequestAuthentication {
	s.BasicAuth = v
	return s
}

type CreateAppRequestAuthenticationBasicAuth struct {
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateAppRequestAuthenticationBasicAuth) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestAuthenticationBasicAuth) GoString() string {
	return s.String()
}

func (s *CreateAppRequestAuthenticationBasicAuth) SetPassword(v string) *CreateAppRequestAuthenticationBasicAuth {
	s.Password = &v
	return s
}

func (s *CreateAppRequestAuthenticationBasicAuth) SetUsername(v string) *CreateAppRequestAuthenticationBasicAuth {
	s.Username = &v
	return s
}

type CreateAppRequestNetwork struct {
	Domain       *string                                `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled      *bool                                  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port         *int32                                 `json:"port,omitempty" xml:"port,omitempty"`
	Type         *string                                `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*CreateAppRequestNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s CreateAppRequestNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppRequestNetwork) SetDomain(v string) *CreateAppRequestNetwork {
	s.Domain = &v
	return s
}

func (s *CreateAppRequestNetwork) SetEnabled(v bool) *CreateAppRequestNetwork {
	s.Enabled = &v
	return s
}

func (s *CreateAppRequestNetwork) SetPort(v int32) *CreateAppRequestNetwork {
	s.Port = &v
	return s
}

func (s *CreateAppRequestNetwork) SetType(v string) *CreateAppRequestNetwork {
	s.Type = &v
	return s
}

func (s *CreateAppRequestNetwork) SetWhiteIpGroup(v []*CreateAppRequestNetworkWhiteIpGroup) *CreateAppRequestNetwork {
	s.WhiteIpGroup = v
	return s
}

type CreateAppRequestNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s CreateAppRequestNetworkWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *CreateAppRequestNetworkWhiteIpGroup) SetGroupName(v string) *CreateAppRequestNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *CreateAppRequestNetworkWhiteIpGroup) SetIps(v []*string) *CreateAppRequestNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

type CreateAppRequestPrivateNetwork struct {
	Enabled       *bool                                         `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PvlEndpointId *string                                       `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	Type          *string                                       `json:"type,omitempty" xml:"type,omitempty"`
	VpcId         *string                                       `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup  []*CreateAppRequestPrivateNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s CreateAppRequestPrivateNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestPrivateNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppRequestPrivateNetwork) SetEnabled(v bool) *CreateAppRequestPrivateNetwork {
	s.Enabled = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetPvlEndpointId(v string) *CreateAppRequestPrivateNetwork {
	s.PvlEndpointId = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetType(v string) *CreateAppRequestPrivateNetwork {
	s.Type = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetVpcId(v string) *CreateAppRequestPrivateNetwork {
	s.VpcId = &v
	return s
}

func (s *CreateAppRequestPrivateNetwork) SetWhiteIpGroup(v []*CreateAppRequestPrivateNetworkWhiteIpGroup) *CreateAppRequestPrivateNetwork {
	s.WhiteIpGroup = v
	return s
}

type CreateAppRequestPrivateNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s CreateAppRequestPrivateNetworkWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestPrivateNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) SetGroupName(v string) *CreateAppRequestPrivateNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *CreateAppRequestPrivateNetworkWhiteIpGroup) SetIps(v []*string) *CreateAppRequestPrivateNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

type CreateAppRequestQuotaInfo struct {
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	Cu      *int32  `json:"cu,omitempty" xml:"cu,omitempty"`
	Storage *int32  `json:"storage,omitempty" xml:"storage,omitempty"`
}

func (s CreateAppRequestQuotaInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestQuotaInfo) GoString() string {
	return s.String()
}

func (s *CreateAppRequestQuotaInfo) SetAppType(v string) *CreateAppRequestQuotaInfo {
	s.AppType = &v
	return s
}

func (s *CreateAppRequestQuotaInfo) SetCu(v int32) *CreateAppRequestQuotaInfo {
	s.Cu = &v
	return s
}

func (s *CreateAppRequestQuotaInfo) SetStorage(v int32) *CreateAppRequestQuotaInfo {
	s.Storage = &v
	return s
}

type CreateAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponseBodyResult struct {
	// example:
	//
	// es-serverless-cn-xxx
	InstaneId *string `json:"instaneId,omitempty" xml:"instaneId,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) SetInstaneId(v string) *CreateAppResponseBodyResult {
	s.InstaneId = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateEndpointRequest struct {
	// This parameter is required.
	EndpointZones []*CreateEndpointRequestEndpointZones `json:"endpointZones,omitempty" xml:"endpointZones,omitempty" type:"Repeated"`
	// example:
	//
	// testendpoint
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf664nyle5khp5***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequest) SetEndpointZones(v []*CreateEndpointRequestEndpointZones) *CreateEndpointRequest {
	s.EndpointZones = v
	return s
}

func (s *CreateEndpointRequest) SetName(v string) *CreateEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointRequest) SetVpcId(v string) *CreateEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointRequest) SetType(v string) *CreateEndpointRequest {
	s.Type = &v
	return s
}

type CreateEndpointRequestEndpointZones struct {
	// example:
	//
	// vsw-uf6qmfkqdcw*****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateEndpointRequestEndpointZones) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointRequestEndpointZones) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequestEndpointZones) SetVswitchId(v string) *CreateEndpointRequestEndpointZones {
	s.VswitchId = &v
	return s
}

func (s *CreateEndpointRequestEndpointZones) SetZoneId(v string) *CreateEndpointRequestEndpointZones {
	s.ZoneId = &v
	return s
}

type CreateEndpointResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateEndpointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBody) SetRequestId(v string) *CreateEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEndpointResponseBody) SetResult(v *CreateEndpointResponseBodyResult) *CreateEndpointResponseBody {
	s.Result = v
	return s
}

type CreateEndpointResponseBodyResult struct {
	// example:
	//
	// essep-abd***dks
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
}

func (s CreateEndpointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponseBodyResult) SetEndpointId(v string) *CreateEndpointResponseBodyResult {
	s.EndpointId = &v
	return s
}

type CreateEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointResponse) SetHeaders(v map[string]*string) *CreateEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointResponse) SetStatusCode(v int32) *CreateEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointResponse) SetBody(v *CreateEndpointResponseBody) *CreateEndpointResponse {
	s.Body = v
	return s
}

type DeleteAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetResult(v *DeleteAppResponseBodyResult) *DeleteAppResponseBody {
	s.Result = v
	return s
}

type DeleteAppResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBodyResult) SetInstanceId(v string) *DeleteAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	// example:
	//
	// false
	Detailed *bool `json:"detailed,omitempty" xml:"detailed,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetDetailed(v bool) *GetAppRequest {
	s.Detailed = &v
	return s
}

type GetAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetResult(v *GetAppResponseBodyResult) *GetAppResponseBody {
	s.Result = v
	return s
}

type GetAppResponseBodyResult struct {
	// example:
	//
	// test-app-abc
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// es-severless-test-app
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// 2022-08-15T11:20:52.370Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 2022-08-15T11:21:50.000Z
	ModifiedTime *string                            `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Network      []*GetAppResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Repeated"`
	// example:
	//
	// *******7595
	OwnerId        *string                                   `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	PrivateNetwork []*GetAppResponseBodyResultPrivateNetwork `json:"privateNetwork,omitempty" xml:"privateNetwork,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 7.10
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResult) SetAppId(v string) *GetAppResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppName(v string) *GetAppResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyResult) SetCreateTime(v string) *GetAppResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyResult) SetDescription(v string) *GetAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetAppResponseBodyResult) SetInstanceId(v string) *GetAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetModifiedTime(v string) *GetAppResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *GetAppResponseBodyResult) SetNetwork(v []*GetAppResponseBodyResultNetwork) *GetAppResponseBodyResult {
	s.Network = v
	return s
}

func (s *GetAppResponseBodyResult) SetOwnerId(v string) *GetAppResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetPrivateNetwork(v []*GetAppResponseBodyResultPrivateNetwork) *GetAppResponseBodyResult {
	s.PrivateNetwork = v
	return s
}

func (s *GetAppResponseBodyResult) SetRegionId(v string) *GetAppResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetStatus(v string) *GetAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAppResponseBodyResult) SetVersion(v string) *GetAppResponseBodyResult {
	s.Version = &v
	return s
}

type GetAppResponseBodyResultNetwork struct {
	Domain       *string                                        `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled      *bool                                          `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port         *int32                                         `json:"port,omitempty" xml:"port,omitempty"`
	Type         *string                                        `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*GetAppResponseBodyResultNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultNetwork) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultNetwork) SetDomain(v string) *GetAppResponseBodyResultNetwork {
	s.Domain = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetEnabled(v bool) *GetAppResponseBodyResultNetwork {
	s.Enabled = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetPort(v int32) *GetAppResponseBodyResultNetwork {
	s.Port = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetType(v string) *GetAppResponseBodyResultNetwork {
	s.Type = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetWhiteIpGroup(v []*GetAppResponseBodyResultNetworkWhiteIpGroup) *GetAppResponseBodyResultNetwork {
	s.WhiteIpGroup = v
	return s
}

type GetAppResponseBodyResultNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultNetworkWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResultNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) SetGroupName(v string) *GetAppResponseBodyResultNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) SetIps(v []*string) *GetAppResponseBodyResultNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

type GetAppResponseBodyResultPrivateNetwork struct {
	Domain        *string                                               `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled       *bool                                                 `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port          *int32                                                `json:"port,omitempty" xml:"port,omitempty"`
	PvlEndpointId *string                                               `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	Type          *string                                               `json:"type,omitempty" xml:"type,omitempty"`
	VpcId         *string                                               `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup  []*GetAppResponseBodyResultPrivateNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultPrivateNetwork) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResultPrivateNetwork) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetDomain(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.Domain = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetEnabled(v bool) *GetAppResponseBodyResultPrivateNetwork {
	s.Enabled = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetPort(v int32) *GetAppResponseBodyResultPrivateNetwork {
	s.Port = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetPvlEndpointId(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.PvlEndpointId = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetType(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.Type = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetVpcId(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.VpcId = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetWhiteIpGroup(v []*GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) *GetAppResponseBodyResultPrivateNetwork {
	s.WhiteIpGroup = v
	return s
}

type GetAppResponseBodyResultPrivateNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) SetGroupName(v string) *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) SetIps(v []*string) *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetAppQuotaResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAppQuotaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAppQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBody) SetRequestId(v string) *GetAppQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppQuotaResponseBody) SetResult(v *GetAppQuotaResponseBodyResult) *GetAppQuotaResponseBody {
	s.Result = v
	return s
}

type GetAppQuotaResponseBodyResult struct {
	LimiterInfo *GetAppQuotaResponseBodyResultLimiterInfo `json:"limiterInfo,omitempty" xml:"limiterInfo,omitempty" type:"Struct"`
	QuotaInfo   map[string]interface{}                    `json:"quotaInfo,omitempty" xml:"quotaInfo,omitempty"`
}

func (s GetAppQuotaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAppQuotaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBodyResult) SetLimiterInfo(v *GetAppQuotaResponseBodyResultLimiterInfo) *GetAppQuotaResponseBodyResult {
	s.LimiterInfo = v
	return s
}

func (s *GetAppQuotaResponseBodyResult) SetQuotaInfo(v map[string]interface{}) *GetAppQuotaResponseBodyResult {
	s.QuotaInfo = v
	return s
}

type GetAppQuotaResponseBodyResultLimiterInfo struct {
	Limiters []*GetAppQuotaResponseBodyResultLimiterInfoLimiters `json:"limiters,omitempty" xml:"limiters,omitempty" type:"Repeated"`
}

func (s GetAppQuotaResponseBodyResultLimiterInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAppQuotaResponseBodyResultLimiterInfo) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBodyResultLimiterInfo) SetLimiters(v []*GetAppQuotaResponseBodyResultLimiterInfoLimiters) *GetAppQuotaResponseBodyResultLimiterInfo {
	s.Limiters = v
	return s
}

type GetAppQuotaResponseBodyResultLimiterInfoLimiters struct {
	// example:
	//
	// true
	Immutable *bool `json:"immutable,omitempty" xml:"immutable,omitempty"`
	// example:
	//
	// 10
	MaxValue *int64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	// example:
	//
	// 1
	MinValue *int64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
	// example:
	//
	// INDEX_NUMBER_OF_SHARDS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAppQuotaResponseBodyResultLimiterInfoLimiters) String() string {
	return tea.Prettify(s)
}

func (s GetAppQuotaResponseBodyResultLimiterInfoLimiters) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetImmutable(v bool) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.Immutable = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetMaxValue(v int64) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.MaxValue = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetMinValue(v int64) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.MinValue = &v
	return s
}

func (s *GetAppQuotaResponseBodyResultLimiterInfoLimiters) SetType(v string) *GetAppQuotaResponseBodyResultLimiterInfoLimiters {
	s.Type = &v
	return s
}

type GetAppQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetAppQuotaResponse) SetHeaders(v map[string]*string) *GetAppQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetAppQuotaResponse) SetStatusCode(v int32) *GetAppQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppQuotaResponse) SetBody(v *GetAppQuotaResponseBody) *GetAppQuotaResponse {
	s.Body = v
	return s
}

type GetMonitorDataRequest struct {
	// example:
	//
	// {"start":1689245180581,"end":1689246950582,"queries":[{"metric":"aliyunes.elasticsearch.index.docs.count","aggregator":"sum","downsample":"avg","tags":{"resource":"{appName}"},"filters":[],"granularity":"auto"}]}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *GetMonitorDataRequest) SetBody(v string) *GetMonitorDataRequest {
	s.Body = &v
	return s
}

type GetMonitorDataResponseBody struct {
	// example:
	//
	// InternalServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// internal server error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetMonitorDataResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonitorDataResponseBody) SetCode(v string) *GetMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetMonitorDataResponseBody) SetMessage(v string) *GetMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonitorDataResponseBody) SetRequestId(v string) *GetMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonitorDataResponseBody) SetResult(v []*GetMonitorDataResponseBodyResult) *GetMonitorDataResponseBody {
	s.Result = v
	return s
}

func (s *GetMonitorDataResponseBody) SetSuccess(v bool) *GetMonitorDataResponseBody {
	s.Success = &v
	return s
}

type GetMonitorDataResponseBodyResult struct {
	// example:
	//
	// {
	//
	//     "1689480600":28676235.104761902
	//
	// }
	Dps map[string]interface{} `json:"dps,omitempty" xml:"dps,omitempty"`
	// example:
	//
	// 1
	Integrity *float32 `json:"integrity,omitempty" xml:"integrity,omitempty"`
	// example:
	//
	// 1689566839447
	MessageWatermark *int64 `json:"messageWatermark,omitempty" xml:"messageWatermark,omitempty"`
	// example:
	//
	// elasticsearch-server.logic_cpu.cpu
	Metric *string `json:"metric,omitempty" xml:"metric,omitempty"`
	// example:
	//
	// 172455913.18935508
	Summary *float32 `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// {
	//
	//                 "indexName":"test"
	//
	//             }
	Tags map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetMonitorDataResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMonitorDataResponseBodyResult) SetDps(v map[string]interface{}) *GetMonitorDataResponseBodyResult {
	s.Dps = v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetIntegrity(v float32) *GetMonitorDataResponseBodyResult {
	s.Integrity = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetMessageWatermark(v int64) *GetMonitorDataResponseBodyResult {
	s.MessageWatermark = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetMetric(v string) *GetMonitorDataResponseBodyResult {
	s.Metric = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetSummary(v float32) *GetMonitorDataResponseBodyResult {
	s.Summary = &v
	return s
}

func (s *GetMonitorDataResponseBodyResult) SetTags(v map[string]interface{}) *GetMonitorDataResponseBodyResult {
	s.Tags = v
	return s
}

type GetMonitorDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorDataResponse) SetHeaders(v map[string]*string) *GetMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorDataResponse) SetStatusCode(v int32) *GetMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitorDataResponse) SetBody(v *GetMonitorDataResponseBody) *GetMonitorDataResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// example:
	//
	// es-severless-test-app
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// 2023-08-29T02:37:22Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// metrics-logs-online
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// desc
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetAppName(v string) *ListAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListAppsRequest) SetCreateTime(v string) *ListAppsRequest {
	s.CreateTime = &v
	return s
}

func (s *ListAppsRequest) SetDescription(v string) *ListAppsRequest {
	s.Description = &v
	return s
}

func (s *ListAppsRequest) SetOrderType(v string) *ListAppsRequest {
	s.OrderType = &v
	return s
}

func (s *ListAppsRequest) SetPageNumber(v int32) *ListAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetStatus(v string) *ListAppsRequest {
	s.Status = &v
	return s
}

type ListAppsResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListAppsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetResult(v []*ListAppsResponseBodyResult) *ListAppsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppsResponseBody) SetTotalCount(v int32) *ListAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppsResponseBodyResult struct {
	// example:
	//
	// test-abc
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// es-severless-test-app
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// 2022-12-27T07:09:11.000Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 应用备注
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 2022-12-27T07:09:11.000Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// OwnerID账号ID
	//
	// example:
	//
	// *********7595
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 代表region的资源属性字段
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 代表资源状态的资源属性字段
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 7.10
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyResult) SetAppId(v string) *ListAppsResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetAppName(v string) *ListAppsResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetCreateTime(v string) *ListAppsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetDescription(v string) *ListAppsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetInstanceId(v string) *ListAppsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetModifiedTime(v string) *ListAppsResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetOwnerId(v string) *ListAppsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetRegionId(v string) *ListAppsResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetStatus(v string) *ListAppsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAppsResponseBodyResult) SetVersion(v string) *ListAppsResponseBodyResult {
	s.Version = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	ApplyReason    *string                         `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	Authentication *UpdateAppRequestAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	ContactInfo    *string                         `json:"contactInfo,omitempty" xml:"contactInfo,omitempty"`
	// 应用备注
	Description    *string                           `json:"description,omitempty" xml:"description,omitempty"`
	LimiterInfo    *UpdateAppRequestLimiterInfo      `json:"limiterInfo,omitempty" xml:"limiterInfo,omitempty" type:"Struct"`
	Network        []*UpdateAppRequestNetwork        `json:"network,omitempty" xml:"network,omitempty" type:"Repeated"`
	PrivateNetwork []*UpdateAppRequestPrivateNetwork `json:"privateNetwork,omitempty" xml:"privateNetwork,omitempty" type:"Repeated"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetApplyReason(v string) *UpdateAppRequest {
	s.ApplyReason = &v
	return s
}

func (s *UpdateAppRequest) SetAuthentication(v *UpdateAppRequestAuthentication) *UpdateAppRequest {
	s.Authentication = v
	return s
}

func (s *UpdateAppRequest) SetContactInfo(v string) *UpdateAppRequest {
	s.ContactInfo = &v
	return s
}

func (s *UpdateAppRequest) SetDescription(v string) *UpdateAppRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppRequest) SetLimiterInfo(v *UpdateAppRequestLimiterInfo) *UpdateAppRequest {
	s.LimiterInfo = v
	return s
}

func (s *UpdateAppRequest) SetNetwork(v []*UpdateAppRequestNetwork) *UpdateAppRequest {
	s.Network = v
	return s
}

func (s *UpdateAppRequest) SetPrivateNetwork(v []*UpdateAppRequestPrivateNetwork) *UpdateAppRequest {
	s.PrivateNetwork = v
	return s
}

type UpdateAppRequestAuthentication struct {
	BasicAuth []*UpdateAppRequestAuthenticationBasicAuth `json:"basicAuth,omitempty" xml:"basicAuth,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestAuthentication) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestAuthentication) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestAuthentication) SetBasicAuth(v []*UpdateAppRequestAuthenticationBasicAuth) *UpdateAppRequestAuthentication {
	s.BasicAuth = v
	return s
}

type UpdateAppRequestAuthenticationBasicAuth struct {
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s UpdateAppRequestAuthenticationBasicAuth) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestAuthenticationBasicAuth) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestAuthenticationBasicAuth) SetPassword(v string) *UpdateAppRequestAuthenticationBasicAuth {
	s.Password = &v
	return s
}

func (s *UpdateAppRequestAuthenticationBasicAuth) SetUsername(v string) *UpdateAppRequestAuthenticationBasicAuth {
	s.Username = &v
	return s
}

type UpdateAppRequestLimiterInfo struct {
	Limiters []*UpdateAppRequestLimiterInfoLimiters `json:"limiters,omitempty" xml:"limiters,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestLimiterInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestLimiterInfo) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestLimiterInfo) SetLimiters(v []*UpdateAppRequestLimiterInfoLimiters) *UpdateAppRequestLimiterInfo {
	s.Limiters = v
	return s
}

type UpdateAppRequestLimiterInfoLimiters struct {
	MaxValue *int32    `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *int32    `json:"minValue,omitempty" xml:"minValue,omitempty"`
	Type     *string   `json:"type,omitempty" xml:"type,omitempty"`
	Values   []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestLimiterInfoLimiters) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestLimiterInfoLimiters) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetMaxValue(v int32) *UpdateAppRequestLimiterInfoLimiters {
	s.MaxValue = &v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetMinValue(v int32) *UpdateAppRequestLimiterInfoLimiters {
	s.MinValue = &v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetType(v string) *UpdateAppRequestLimiterInfoLimiters {
	s.Type = &v
	return s
}

func (s *UpdateAppRequestLimiterInfoLimiters) SetValues(v []*string) *UpdateAppRequestLimiterInfoLimiters {
	s.Values = v
	return s
}

type UpdateAppRequestNetwork struct {
	Domain       *string                                `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled      *bool                                  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port         *int32                                 `json:"port,omitempty" xml:"port,omitempty"`
	Type         *string                                `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*UpdateAppRequestNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestNetwork) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestNetwork) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestNetwork) SetDomain(v string) *UpdateAppRequestNetwork {
	s.Domain = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetEnabled(v bool) *UpdateAppRequestNetwork {
	s.Enabled = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetPort(v int32) *UpdateAppRequestNetwork {
	s.Port = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetType(v string) *UpdateAppRequestNetwork {
	s.Type = &v
	return s
}

func (s *UpdateAppRequestNetwork) SetWhiteIpGroup(v []*UpdateAppRequestNetworkWhiteIpGroup) *UpdateAppRequestNetwork {
	s.WhiteIpGroup = v
	return s
}

type UpdateAppRequestNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestNetworkWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) SetGroupName(v string) *UpdateAppRequestNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateAppRequestNetworkWhiteIpGroup) SetIps(v []*string) *UpdateAppRequestNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

type UpdateAppRequestPrivateNetwork struct {
	Enabled       *bool                                         `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PvlEndpointId *string                                       `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	Type          *string                                       `json:"type,omitempty" xml:"type,omitempty"`
	VpcId         *string                                       `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup  []*UpdateAppRequestPrivateNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestPrivateNetwork) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestPrivateNetwork) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestPrivateNetwork) SetEnabled(v bool) *UpdateAppRequestPrivateNetwork {
	s.Enabled = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetPvlEndpointId(v string) *UpdateAppRequestPrivateNetwork {
	s.PvlEndpointId = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetType(v string) *UpdateAppRequestPrivateNetwork {
	s.Type = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetVpcId(v string) *UpdateAppRequestPrivateNetwork {
	s.VpcId = &v
	return s
}

func (s *UpdateAppRequestPrivateNetwork) SetWhiteIpGroup(v []*UpdateAppRequestPrivateNetworkWhiteIpGroup) *UpdateAppRequestPrivateNetwork {
	s.WhiteIpGroup = v
	return s
}

type UpdateAppRequestPrivateNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s UpdateAppRequestPrivateNetworkWhiteIpGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequestPrivateNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) SetGroupName(v string) *UpdateAppRequestPrivateNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateAppRequestPrivateNetworkWhiteIpGroup) SetIps(v []*string) *UpdateAppRequestPrivateNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

type UpdateAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetResult(v *UpdateAppResponseBodyResult) *UpdateAppResponseBody {
	s.Result = v
	return s
}

type UpdateAppResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBodyResult) SetInstanceId(v string) *UpdateAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

type UpdateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetStatusCode(v int32) *UpdateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("es-serverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建Serverless应用
//
// @param request - CreateAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Authentication)) {
		body["authentication"] = request.Authentication
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Network)) {
		body["network"] = request.Network
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateNetwork)) {
		body["privateNetwork"] = request.PrivateNetwork
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaInfo)) {
		body["quotaInfo"] = request.QuotaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Serverless应用
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建端点
//
// @param request - CreateEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpointWithOptions(request *CreateEndpointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointZones)) {
		body["endpointZones"] = request.EndpointZones
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["vpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEndpoint"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/endpoints"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建端点
//
// @param request - CreateEndpointRequest
//
// @return CreateEndpointResponse
func (client *Client) CreateEndpoint(request *CreateEndpointRequest) (_result *CreateEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEndpointResponse{}
	_body, _err := client.CreateEndpointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Serverless应用。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Serverless应用。
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(appName *string) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Serverless应用详情
//
// @param request - GetAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppResponse
func (client *Client) GetAppWithOptions(appName *string, request *GetAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detailed)) {
		query["detailed"] = request.Detailed
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Serverless应用详情
//
// @param request - GetAppRequest
//
// @return GetAppResponse
func (client *Client) GetApp(appName *string, request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Serverless应用配额详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppQuotaResponse
func (client *Client) GetAppQuotaWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppQuota"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/quota"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Serverless应用配额详情
//
// @return GetAppQuotaResponse
func (client *Client) GetAppQuota(appName *string) (_result *GetAppQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppQuotaResponse{}
	_body, _err := client.GetAppQuotaWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取监控数据
//
// @param request - GetMonitorDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorDataResponse
func (client *Client) GetMonitorDataWithOptions(request *GetMonitorDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMonitorData"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/emon/metrics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取监控数据
//
// @param request - GetMonitorDataRequest
//
// @return GetMonitorDataResponse
func (client *Client) GetMonitorData(request *GetMonitorDataRequest) (_result *GetMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMonitorDataResponse{}
	_body, _err := client.GetMonitorDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看Serverless应用列表
//
// @param request - ListAppsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsResponse
func (client *Client) ListAppsWithOptions(request *ListAppsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["orderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看Serverless应用列表
//
// @param request - ListAppsRequest
//
// @return ListAppsResponse
func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑Serverless应用
//
// @param request - UpdateAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppResponse
func (client *Client) UpdateAppWithOptions(appName *string, request *UpdateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyReason)) {
		body["applyReason"] = request.ApplyReason
	}

	if !tea.BoolValue(util.IsUnset(request.Authentication)) {
		body["authentication"] = request.Authentication
	}

	if !tea.BoolValue(util.IsUnset(request.ContactInfo)) {
		body["contactInfo"] = request.ContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.LimiterInfo)) {
		body["limiterInfo"] = request.LimiterInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Network)) {
		body["network"] = request.Network
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateNetwork)) {
		body["privateNetwork"] = request.PrivateNetwork
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApp"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑Serverless应用
//
// @param request - UpdateAppRequest
//
// @return UpdateAppResponse
func (client *Client) UpdateApp(appName *string, request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
