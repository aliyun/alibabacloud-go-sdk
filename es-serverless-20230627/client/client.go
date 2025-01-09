// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelSpecReviewTaskResponseBody struct {
	// example:
	//
	// 1B64F3E0-25D5-5043-B5C8-4FF22BB12CCD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CancelSpecReviewTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelSpecReviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSpecReviewTaskResponseBody) SetRequestId(v string) *CancelSpecReviewTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSpecReviewTaskResponseBody) SetResult(v bool) *CancelSpecReviewTaskResponseBody {
	s.Result = &v
	return s
}

type CancelSpecReviewTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSpecReviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSpecReviewTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelSpecReviewTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelSpecReviewTaskResponse) SetHeaders(v map[string]*string) *CancelSpecReviewTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelSpecReviewTaskResponse) SetStatusCode(v int32) *CancelSpecReviewTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSpecReviewTaskResponse) SetBody(v *CancelSpecReviewTaskResponseBody) *CancelSpecReviewTaskResponse {
	s.Body = v
	return s
}

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

type CreateSnapshotRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// product_info
	Indices *string `json:"indices,omitempty" xml:"indices,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qingning
	Snapshot *string `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetIndices(v string) *CreateSnapshotRequest {
	s.Indices = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshot(v string) *CreateSnapshotRequest {
	s.Snapshot = &v
	return s
}

func (s *CreateSnapshotRequest) SetDryRun(v bool) *CreateSnapshotRequest {
	s.DryRun = &v
	return s
}

type CreateSnapshotResponseBody struct {
	// example:
	//
	// 03761BE5-D12F-55B4-9C93-0255C11DE44A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetResult(v bool) *CreateSnapshotResponseBody {
	s.Result = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
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

type DeleteDictRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeleteDictRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDictRequest) GoString() string {
	return s.String()
}

func (s *DeleteDictRequest) SetName(v string) *DeleteDictRequest {
	s.Name = &v
	return s
}

func (s *DeleteDictRequest) SetType(v string) *DeleteDictRequest {
	s.Type = &v
	return s
}

type DeleteDictResponseBody struct {
	// example:
	//
	// 2BF6B57E-5AAD-5389-80CD-E200BBF91FF9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteDictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDictResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDictResponseBody) SetRequestId(v string) *DeleteDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDictResponseBody) SetResult(v bool) *DeleteDictResponseBody {
	s.Result = &v
	return s
}

type DeleteDictResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDictResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDictResponse) GoString() string {
	return s.String()
}

func (s *DeleteDictResponse) SetHeaders(v map[string]*string) *DeleteDictResponse {
	s.Headers = v
	return s
}

func (s *DeleteDictResponse) SetStatusCode(v int32) *DeleteDictResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDictResponse) SetBody(v *DeleteDictResponseBody) *DeleteDictResponse {
	s.Body = v
	return s
}

type DeleteEndpointResponseBody struct {
	// example:
	//
	// 1305A3E0-A291-54BA-A3B2-7D1C12EC4112
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointResponseBody) SetRequestId(v string) *DeleteEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEndpointResponseBody) SetResult(v bool) *DeleteEndpointResponseBody {
	s.Result = &v
	return s
}

type DeleteEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteEndpointResponse) SetHeaders(v map[string]*string) *DeleteEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteEndpointResponse) SetStatusCode(v int32) *DeleteEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEndpointResponse) SetBody(v *DeleteEndpointResponseBody) *DeleteEndpointResponse {
	s.Body = v
	return s
}

type DeleteSnapshotResponseBody struct {
	// example:
	//
	// 16484F36-A2A3-5A05-B242-0BF2BF6AA326
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetResult(v bool) *DeleteSnapshotResponseBody {
	s.Result = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
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
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
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

func (s *GetAppResponseBodyResult) SetAppType(v string) *GetAppResponseBodyResult {
	s.AppType = &v
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

type GetSnapshotSettingResponseBody struct {
	// example:
	//
	// 7B6CE6E1-5BA0-56DA-BFFD-8D90692F1EFC
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSnapshotSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSnapshotSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotSettingResponseBody) SetRequestId(v string) *GetSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSnapshotSettingResponseBody) SetResult(v *GetSnapshotSettingResponseBodyResult) *GetSnapshotSettingResponseBody {
	s.Result = v
	return s
}

type GetSnapshotSettingResponseBodyResult struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 0 0 01 ? 	- 	- *
	QuartzRegex *string `json:"quartzRegex,omitempty" xml:"quartzRegex,omitempty"`
}

func (s GetSnapshotSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSnapshotSettingResponseBodyResult) SetEnable(v bool) *GetSnapshotSettingResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *GetSnapshotSettingResponseBodyResult) SetQuartzRegex(v string) *GetSnapshotSettingResponseBodyResult {
	s.QuartzRegex = &v
	return s
}

type GetSnapshotSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSnapshotSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotSettingResponse) SetHeaders(v map[string]*string) *GetSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotSettingResponse) SetStatusCode(v int32) *GetSnapshotSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSnapshotSettingResponse) SetBody(v *GetSnapshotSettingResponseBody) *GetSnapshotSettingResponse {
	s.Body = v
	return s
}

type GetSpecReviewTaskResponseBody struct {
	// example:
	//
	// E310AC54-957A-5FD5-B85B-E972B2BDA8DE
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSpecReviewTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSpecReviewTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpecReviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpecReviewTaskResponseBody) SetRequestId(v string) *GetSpecReviewTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpecReviewTaskResponseBody) SetResult(v *GetSpecReviewTaskResponseBodyResult) *GetSpecReviewTaskResponseBody {
	s.Result = v
	return s
}

type GetSpecReviewTaskResponseBodyResult struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 339
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// example:
	//
	// {
	//
	//                "limiters": [
	//
	//                     {
	//
	//                          "type": "INDEX_QUOTA",
	//
	//                          "maxValue": 500,
	//
	//                          "immutable": false
	//
	//                     }
	//
	//                ]
	//
	//           }
	ApplyLimiter map[string]interface{} `json:"applyLimiter,omitempty" xml:"applyLimiter,omitempty"`
	// example:
	//
	// {
	//
	//                "appType": "TRIAL",
	//
	//                "cu": 4,
	//
	//                "storage": 100
	//
	//           }
	ApplyQuota  map[string]interface{} `json:"applyQuota,omitempty" xml:"applyQuota,omitempty"`
	ApplyReason *string                `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// example:
	//
	// {
	//
	//                "limiters": [
	//
	//                     {
	//
	//                          "type": "INDEX_QUOTA",
	//
	//                          "maxValue": 500,
	//
	//                          "immutable": false
	//
	//                     }
	//
	//                ]
	//
	//           }
	EffectiveLimiter map[string]interface{} `json:"effectiveLimiter,omitempty" xml:"effectiveLimiter,omitempty"`
	// example:
	//
	// {
	//
	//                "appType": "TRIAL",
	//
	//                "cu": 4,
	//
	//                "storage": 100
	//
	//           }
	EffectiveQuota map[string]interface{} `json:"effectiveQuota,omitempty" xml:"effectiveQuota,omitempty"`
	// example:
	//
	// 2024-05-30T06:28:07.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-05-30T06:28:07.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// {
	//
	//                "limiters": [
	//
	//                     {
	//
	//                          "type": "INDEX_QUOTA",
	//
	//                          "maxValue": 500,
	//
	//                          "immutable": false
	//
	//                     }
	//
	//                ]
	//
	//           }
	OldLimiter map[string]interface{} `json:"oldLimiter,omitempty" xml:"oldLimiter,omitempty"`
	// example:
	//
	// {
	//
	//                "appType": "TRIAL",
	//
	//                "cu": 2,
	//
	//                "storage": 1
	//
	//           }
	OldQuota map[string]interface{} `json:"oldQuota,omitempty" xml:"oldQuota,omitempty"`
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// QUOTA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSpecReviewTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSpecReviewTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSpecReviewTaskResponseBodyResult) SetId(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetAppName(v string) *GetSpecReviewTaskResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetApplyLimiter(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.ApplyLimiter = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetApplyQuota(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.ApplyQuota = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetApplyReason(v string) *GetSpecReviewTaskResponseBodyResult {
	s.ApplyReason = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetEffectiveLimiter(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.EffectiveLimiter = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetEffectiveQuota(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.EffectiveQuota = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetGmtCreate(v string) *GetSpecReviewTaskResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetGmtModified(v string) *GetSpecReviewTaskResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetOldLimiter(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.OldLimiter = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetOldQuota(v map[string]interface{}) *GetSpecReviewTaskResponseBodyResult {
	s.OldQuota = v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetSource(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetStatus(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSpecReviewTaskResponseBodyResult) SetType(v string) *GetSpecReviewTaskResponseBodyResult {
	s.Type = &v
	return s
}

type GetSpecReviewTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpecReviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpecReviewTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpecReviewTaskResponse) GoString() string {
	return s.String()
}

func (s *GetSpecReviewTaskResponse) SetHeaders(v map[string]*string) *GetSpecReviewTaskResponse {
	s.Headers = v
	return s
}

func (s *GetSpecReviewTaskResponse) SetStatusCode(v int32) *GetSpecReviewTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpecReviewTaskResponse) SetBody(v *GetSpecReviewTaskResponseBody) *GetSpecReviewTaskResponse {
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
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
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

func (s *ListAppsResponseBodyResult) SetAppType(v string) *ListAppsResponseBodyResult {
	s.AppType = &v
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

type ListDictsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDictsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDictsRequest) GoString() string {
	return s.String()
}

func (s *ListDictsRequest) SetPageNumber(v int32) *ListDictsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDictsRequest) SetPageSize(v int32) *ListDictsRequest {
	s.PageSize = &v
	return s
}

type ListDictsResponseBody struct {
	// example:
	//
	// E92BCBB9-3CFE-58DD-8D8C-56DF46AB3BF3
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListDictsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDictsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBody) SetRequestId(v string) *ListDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDictsResponseBody) SetResult(v []*ListDictsResponseBodyResult) *ListDictsResponseBody {
	s.Result = v
	return s
}

func (s *ListDictsResponseBody) SetTotalCount(v int32) *ListDictsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDictsResponseBodyResult struct {
	// example:
	//
	// http://es-serverless-****.oss-cn-hangzhou.aliyuncs.com/app/es7**190/0/config/analysis-ik/stopword.dic?Expires=1705923089&OSSAccessKeyId=STS.NV18q****UkVp6LNj&Signat
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// example:
	//
	// a.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDictsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBodyResult) SetDownloadUrl(v string) *ListDictsResponseBodyResult {
	s.DownloadUrl = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetName(v string) *ListDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetSourceType(v string) *ListDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetType(v string) *ListDictsResponseBodyResult {
	s.Type = &v
	return s
}

type ListDictsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDictsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDictsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDictsResponse) GoString() string {
	return s.String()
}

func (s *ListDictsResponse) SetHeaders(v map[string]*string) *ListDictsResponse {
	s.Headers = v
	return s
}

func (s *ListDictsResponse) SetStatusCode(v int32) *ListDictsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDictsResponse) SetBody(v *ListDictsResponseBody) *ListDictsResponse {
	s.Body = v
	return s
}

type ListEndpointsRequest struct {
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
	// ep-bp1id41dd116e52e****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-bp1212sb7cj2j4e6x****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointsRequest) SetPageNumber(v int32) *ListEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointsRequest) SetPageSize(v int32) *ListEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEndpointsRequest) SetResourceId(v string) *ListEndpointsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListEndpointsRequest) SetType(v string) *ListEndpointsRequest {
	s.Type = &v
	return s
}

func (s *ListEndpointsRequest) SetVpcId(v string) *ListEndpointsRequest {
	s.VpcId = &v
	return s
}

type ListEndpointsResponseBody struct {
	// example:
	//
	// D6030CE6-9FEB-5B2F-84AC-3ADE3CBA89E5
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListEndpointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBody) SetRequestId(v string) *ListEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointsResponseBody) SetResult(v []*ListEndpointsResponseBodyResult) *ListEndpointsResponseBody {
	s.Result = v
	return s
}

func (s *ListEndpointsResponseBody) SetTotalCount(v int32) *ListEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEndpointsResponseBodyResult struct {
	// example:
	//
	// Pending
	ConnectionStatus *string `json:"connectionStatus,omitempty" xml:"connectionStatus,omitempty"`
	// example:
	//
	// 1701259721
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// ep-bp1i522d****a3.epsrv-bp1f****gei.cn-hangzhou.privatelink.aliyuncs.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// essep-2f46b743f60****
	EndpointId    *string                                         `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	EndpointZones []*ListEndpointsResponseBodyResultEndpointZones `json:"endpointZones,omitempty" xml:"endpointZones,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ep-bp1id41dd116e52e****
	ResourceId       *string   `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1701259721
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// example:
	//
	// vpc-uf6gykvwcirp886ef****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEndpointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyResult) SetConnectionStatus(v string) *ListEndpointsResponseBodyResult {
	s.ConnectionStatus = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetCreated(v int32) *ListEndpointsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetDomain(v string) *ListEndpointsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetEndpointId(v string) *ListEndpointsResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetEndpointZones(v []*ListEndpointsResponseBodyResultEndpointZones) *ListEndpointsResponseBodyResult {
	s.EndpointZones = v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetName(v string) *ListEndpointsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetResourceId(v string) *ListEndpointsResponseBodyResult {
	s.ResourceId = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetSecurityGroupIds(v []*string) *ListEndpointsResponseBodyResult {
	s.SecurityGroupIds = v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetStatus(v string) *ListEndpointsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetType(v string) *ListEndpointsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetUpdated(v int32) *ListEndpointsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListEndpointsResponseBodyResult) SetVpcId(v string) *ListEndpointsResponseBodyResult {
	s.VpcId = &v
	return s
}

type ListEndpointsResponseBodyResultEndpointZones struct {
	// example:
	//
	// vsw-bp194pz9iez****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListEndpointsResponseBodyResultEndpointZones) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponseBodyResultEndpointZones) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponseBodyResultEndpointZones) SetVSwitchId(v string) *ListEndpointsResponseBodyResultEndpointZones {
	s.VSwitchId = &v
	return s
}

func (s *ListEndpointsResponseBodyResultEndpointZones) SetZoneId(v string) *ListEndpointsResponseBodyResultEndpointZones {
	s.ZoneId = &v
	return s
}

type ListEndpointsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListEndpointsResponse) SetHeaders(v map[string]*string) *ListEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListEndpointsResponse) SetStatusCode(v int32) *ListEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEndpointsResponse) SetBody(v *ListEndpointsResponseBody) *ListEndpointsResponse {
	s.Body = v
	return s
}

type ListIndicesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8C85CCB3-C0C9-521C-B599-F903E14A8793
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListIndicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBody) SetRequestId(v string) *ListIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndicesResponseBody) SetResult(v []interface{}) *ListIndicesResponseBody {
	s.Result = v
	return s
}

type ListIndicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListIndicesResponse) SetHeaders(v map[string]*string) *ListIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListIndicesResponse) SetStatusCode(v int32) *ListIndicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndicesResponse) SetBody(v *ListIndicesResponseBody) *ListIndicesResponse {
	s.Body = v
	return s
}

type ListSnapshotRepositoriesResponseBody struct {
	// example:
	//
	// 56E0591D-7D62-56A2-993E-952FB2026C69
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSnapshotRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotRepositoriesResponseBody) SetRequestId(v string) *ListSnapshotRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotRepositoriesResponseBody) SetResult(v []map[string]interface{}) *ListSnapshotRepositoriesResponseBody {
	s.Result = v
	return s
}

type ListSnapshotRepositoriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotRepositoriesResponse) SetHeaders(v map[string]*string) *ListSnapshotRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotRepositoriesResponse) SetStatusCode(v int32) *ListSnapshotRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotRepositoriesResponse) SetBody(v *ListSnapshotRepositoriesResponseBody) *ListSnapshotRepositoriesResponse {
	s.Body = v
	return s
}

type ListSnapshotsRequest struct {
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
	// aliyun_auto_snapshot
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// example:
	//
	// qingning
	Snapshot *string `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) SetPageNumber(v int32) *ListSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v int32) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetRepository(v string) *ListSnapshotsRequest {
	s.Repository = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshot(v string) *ListSnapshotsRequest {
	s.Snapshot = &v
	return s
}

type ListSnapshotsResponseBody struct {
	// example:
	//
	// ODgyObrnP3
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 22E9EE78-F567-550A-8F7C-20E9CD3DE489
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) SetNextToken(v string) *ListSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetResult(v []map[string]interface{}) *ListSnapshotsResponseBody {
	s.Result = v
	return s
}

func (s *ListSnapshotsResponseBody) SetTotalCount(v int32) *ListSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSnapshotsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponse) SetHeaders(v map[string]*string) *ListSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotsResponse) SetStatusCode(v int32) *ListSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotsResponse) SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse {
	s.Body = v
	return s
}

type ListSpecReviewTasksRequest struct {
	// example:
	//
	// 1
	Page       *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// QUOTA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSpecReviewTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpecReviewTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksRequest) SetPage(v int32) *ListSpecReviewTasksRequest {
	s.Page = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetPageNumber(v int32) *ListSpecReviewTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetPageSize(v int32) *ListSpecReviewTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetSize(v int32) *ListSpecReviewTasksRequest {
	s.Size = &v
	return s
}

func (s *ListSpecReviewTasksRequest) SetType(v string) *ListSpecReviewTasksRequest {
	s.Type = &v
	return s
}

type ListSpecReviewTasksResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 55F7B3FE-05D8-5F0F-BD55-A18967D447DC
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListSpecReviewTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSpecReviewTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpecReviewTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksResponseBody) SetRequestId(v string) *ListSpecReviewTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpecReviewTasksResponseBody) SetResult(v []*ListSpecReviewTasksResponseBodyResult) *ListSpecReviewTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListSpecReviewTasksResponseBody) SetTotalCount(v int32) *ListSpecReviewTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListSpecReviewTasksResponseBodyResult struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 339
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	AppName     *string `json:"appName,omitempty" xml:"appName,omitempty"`
	ApplyReason *string `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// example:
	//
	// 2024-05-27T10:13:22.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// QUOTA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSpecReviewTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSpecReviewTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksResponseBodyResult) SetId(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetAppName(v string) *ListSpecReviewTasksResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetApplyReason(v string) *ListSpecReviewTasksResponseBodyResult {
	s.ApplyReason = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetGmtCreate(v string) *ListSpecReviewTasksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetSource(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetStatus(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSpecReviewTasksResponseBodyResult) SetType(v string) *ListSpecReviewTasksResponseBodyResult {
	s.Type = &v
	return s
}

type ListSpecReviewTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSpecReviewTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSpecReviewTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpecReviewTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksResponse) SetHeaders(v map[string]*string) *ListSpecReviewTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSpecReviewTasksResponse) SetStatusCode(v int32) *ListSpecReviewTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpecReviewTasksResponse) SetBody(v *ListSpecReviewTasksResponseBody) *ListSpecReviewTasksResponse {
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

type UpdateDictRequest struct {
	// example:
	//
	// true
	AllowCover *bool `json:"allowCover,omitempty" xml:"allowCover,omitempty"`
	// This parameter is required.
	Files []*UpdateDictRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateDictRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictRequest) GoString() string {
	return s.String()
}

func (s *UpdateDictRequest) SetAllowCover(v bool) *UpdateDictRequest {
	s.AllowCover = &v
	return s
}

func (s *UpdateDictRequest) SetFiles(v []*UpdateDictRequestFiles) *UpdateDictRequest {
	s.Files = v
	return s
}

func (s *UpdateDictRequest) SetSourceType(v string) *UpdateDictRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateDictRequest) SetType(v string) *UpdateDictRequest {
	s.Type = &v
	return s
}

func (s *UpdateDictRequest) SetDryRun(v bool) *UpdateDictRequest {
	s.DryRun = &v
	return s
}

type UpdateDictRequestFiles struct {
	// example:
	//
	// dic_0.dic
	Name      *string                          `json:"name,omitempty" xml:"name,omitempty"`
	OssObject *UpdateDictRequestFilesOssObject `json:"ossObject,omitempty" xml:"ossObject,omitempty" type:"Struct"`
}

func (s UpdateDictRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictRequestFiles) GoString() string {
	return s.String()
}

func (s *UpdateDictRequestFiles) SetName(v string) *UpdateDictRequestFiles {
	s.Name = &v
	return s
}

func (s *UpdateDictRequestFiles) SetOssObject(v *UpdateDictRequestFilesOssObject) *UpdateDictRequestFiles {
	s.OssObject = v
	return s
}

type UpdateDictRequestFilesOssObject struct {
	// example:
	//
	// bucket1
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// oss/dic_0.dic
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateDictRequestFilesOssObject) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictRequestFilesOssObject) GoString() string {
	return s.String()
}

func (s *UpdateDictRequestFilesOssObject) SetBucketName(v string) *UpdateDictRequestFilesOssObject {
	s.BucketName = &v
	return s
}

func (s *UpdateDictRequestFilesOssObject) SetKey(v string) *UpdateDictRequestFilesOssObject {
	s.Key = &v
	return s
}

type UpdateDictResponseBody struct {
	// example:
	//
	// 12797BCC-E0B5-5A47-B4B9-A14DDF0B0200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateDictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDictResponseBody) SetRequestId(v string) *UpdateDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDictResponseBody) SetResult(v bool) *UpdateDictResponseBody {
	s.Result = &v
	return s
}

type UpdateDictResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDictResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDictResponse) GoString() string {
	return s.String()
}

func (s *UpdateDictResponse) SetHeaders(v map[string]*string) *UpdateDictResponse {
	s.Headers = v
	return s
}

func (s *UpdateDictResponse) SetStatusCode(v int32) *UpdateDictResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDictResponse) SetBody(v *UpdateDictResponseBody) *UpdateDictResponse {
	s.Body = v
	return s
}

type UpdateEndpointRequest struct {
	// This parameter is required.
	EndpointZones []*UpdateEndpointRequestEndpointZones `json:"endpointZones,omitempty" xml:"endpointZones,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointRequest) SetEndpointZones(v []*UpdateEndpointRequestEndpointZones) *UpdateEndpointRequest {
	s.EndpointZones = v
	return s
}

func (s *UpdateEndpointRequest) SetName(v string) *UpdateEndpointRequest {
	s.Name = &v
	return s
}

type UpdateEndpointRequestEndpointZones struct {
	// example:
	//
	// vsw-bp18r8uwnukv3rvi9****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s UpdateEndpointRequestEndpointZones) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointRequestEndpointZones) GoString() string {
	return s.String()
}

func (s *UpdateEndpointRequestEndpointZones) SetVSwitchId(v string) *UpdateEndpointRequestEndpointZones {
	s.VSwitchId = &v
	return s
}

func (s *UpdateEndpointRequestEndpointZones) SetZoneId(v string) *UpdateEndpointRequestEndpointZones {
	s.ZoneId = &v
	return s
}

type UpdateEndpointResponseBody struct {
	// example:
	//
	// FBAD8493-87FA-583E-8A4C-D487F2DE90FC
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateEndpointResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointResponseBody) SetRequestId(v string) *UpdateEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEndpointResponseBody) SetResult(v *UpdateEndpointResponseBodyResult) *UpdateEndpointResponseBody {
	s.Result = v
	return s
}

type UpdateEndpointResponseBodyResult struct {
	// example:
	//
	// ep-bp1i98bcbb1540d0****
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
}

func (s UpdateEndpointResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateEndpointResponseBodyResult) SetEndpointId(v string) *UpdateEndpointResponseBodyResult {
	s.EndpointId = &v
	return s
}

type UpdateEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointResponse) SetHeaders(v map[string]*string) *UpdateEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointResponse) SetStatusCode(v int32) *UpdateEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointResponse) SetBody(v *UpdateEndpointResponseBody) *UpdateEndpointResponse {
	s.Body = v
	return s
}

type UpdateSnapshotSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0 0 01 ? 	- 	- *
	QuartzRegex *string `json:"quartzRegex,omitempty" xml:"quartzRegex,omitempty"`
}

func (s UpdateSnapshotSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingRequest) SetEnable(v bool) *UpdateSnapshotSettingRequest {
	s.Enable = &v
	return s
}

func (s *UpdateSnapshotSettingRequest) SetQuartzRegex(v string) *UpdateSnapshotSettingRequest {
	s.QuartzRegex = &v
	return s
}

type UpdateSnapshotSettingResponseBody struct {
	// example:
	//
	// A7B03723-AA73-5A5F-B71C-270792614DD8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {
	//
	//     "quartzRegex": "0 0 01 ? 	- 	- *",
	//
	//     "enable": true
	//
	//   }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSnapshotSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponseBody) SetRequestId(v string) *UpdateSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBody) SetResult(v map[string]interface{}) *UpdateSnapshotSettingResponseBody {
	s.Result = v
	return s
}

type UpdateSnapshotSettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSnapshotSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSnapshotSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponse) SetHeaders(v map[string]*string) *UpdateSnapshotSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotSettingResponse) SetStatusCode(v int32) *UpdateSnapshotSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSnapshotSettingResponse) SetBody(v *UpdateSnapshotSettingResponseBody) *UpdateSnapshotSettingResponse {
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
// 撤销规格审批
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSpecReviewTaskResponse
func (client *Client) CancelSpecReviewTaskWithOptions(appName *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelSpecReviewTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelSpecReviewTask"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/spec-review-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelSpecReviewTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销规格审批
//
// @return CancelSpecReviewTaskResponse
func (client *Client) CancelSpecReviewTask(appName *string, taskId *string) (_result *CancelSpecReviewTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelSpecReviewTaskResponse{}
	_body, _err := client.CancelSpecReviewTaskWithOptions(appName, taskId, headers, runtime)
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
// 创建快照
//
// @param request - CreateSnapshotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithOptions(appName *string, repository *string, request *CreateSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Indices)) {
		body["indices"] = request.Indices
	}

	if !tea.BoolValue(util.IsUnset(request.Snapshot)) {
		body["snapshot"] = request.Snapshot
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/snapshot-repositories/" + tea.StringValue(openapiutil.GetEncodeParam(repository)) + "/snapshots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建快照
//
// @param request - CreateSnapshotRequest
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshot(appName *string, repository *string, request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(appName, repository, request, headers, runtime)
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
// 删除词典
//
// @param request - DeleteDictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDictResponse
func (client *Client) DeleteDictWithOptions(appName *string, request *DeleteDictRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDict"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/dicts/actions/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除词典
//
// @param request - DeleteDictRequest
//
// @return DeleteDictResponse
func (client *Client) DeleteDict(appName *string, request *DeleteDictRequest) (_result *DeleteDictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDictResponse{}
	_body, _err := client.DeleteDictWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除端点
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpointWithOptions(endpointId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEndpointResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEndpoint"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/endpoints/" + tea.StringValue(openapiutil.GetEncodeParam(endpointId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除端点
//
// @return DeleteEndpointResponse
func (client *Client) DeleteEndpoint(endpointId *string) (_result *DeleteEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEndpointResponse{}
	_body, _err := client.DeleteEndpointWithOptions(endpointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除快照
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithOptions(appName *string, repository *string, snapshot *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/snapshot-repositories/" + tea.StringValue(openapiutil.GetEncodeParam(repository)) + "/snapshots/" + tea.StringValue(openapiutil.GetEncodeParam(snapshot))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除快照
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshot(appName *string, repository *string, snapshot *string) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(appName, repository, snapshot, headers, runtime)
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
// 获取自动备份配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSnapshotSettingResponse
func (client *Client) GetSnapshotSettingWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSnapshotSettingResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSnapshotSetting"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/auto-snapshot-setting"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自动备份配置
//
// @return GetSnapshotSettingResponse
func (client *Client) GetSnapshotSetting(appName *string) (_result *GetSnapshotSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSnapshotSettingResponse{}
	_body, _err := client.GetSnapshotSettingWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取配额审批详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpecReviewTaskResponse
func (client *Client) GetSpecReviewTaskWithOptions(appName *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSpecReviewTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSpecReviewTask"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/spec-review-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpecReviewTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取配额审批详情
//
// @return GetSpecReviewTaskResponse
func (client *Client) GetSpecReviewTask(appName *string, taskId *string) (_result *GetSpecReviewTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSpecReviewTaskResponse{}
	_body, _err := client.GetSpecReviewTaskWithOptions(appName, taskId, headers, runtime)
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
// 获取词典列表
//
// @param request - ListDictsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDictsResponse
func (client *Client) ListDictsWithOptions(appName *string, request *ListDictsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDictsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDicts"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/dicts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDictsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取词典列表
//
// @param request - ListDictsRequest
//
// @return ListDictsResponse
func (client *Client) ListDicts(appName *string, request *ListDictsRequest) (_result *ListDictsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDictsResponse{}
	_body, _err := client.ListDictsWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取端点信息列表
//
// @param request - ListEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEndpointsResponse
func (client *Client) ListEndpointsWithOptions(request *ListEndpointsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["vpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEndpoints"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/endpoints"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取端点信息列表
//
// @param request - ListEndpointsRequest
//
// @return ListEndpointsResponse
func (client *Client) ListEndpoints(request *ListEndpointsRequest) (_result *ListEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEndpointsResponse{}
	_body, _err := client.ListEndpointsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看索引列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndicesResponse
func (client *Client) ListIndicesWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndicesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndices"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/indices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看索引列表
//
// @return ListIndicesResponse
func (client *Client) ListIndices(appName *string) (_result *ListIndicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndicesResponse{}
	_body, _err := client.ListIndicesWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取快照仓库列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotRepositoriesResponse
func (client *Client) ListSnapshotRepositoriesWithOptions(appName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSnapshotRepositoriesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshotRepositories"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/snapshot-repositories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取快照仓库列表
//
// @return ListSnapshotRepositoriesResponse
func (client *Client) ListSnapshotRepositories(appName *string) (_result *ListSnapshotRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotRepositoriesResponse{}
	_body, _err := client.ListSnapshotRepositoriesWithOptions(appName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取仓库的快照列表
//
// @param request - ListSnapshotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithOptions(appName *string, request *ListSnapshotsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Repository)) {
		query["repository"] = request.Repository
	}

	if !tea.BoolValue(util.IsUnset(request.Snapshot)) {
		query["snapshot"] = request.Snapshot
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshots"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/snapshots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取仓库的快照列表
//
// @param request - ListSnapshotsRequest
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshots(appName *string, request *ListSnapshotsRequest) (_result *ListSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotsResponse{}
	_body, _err := client.ListSnapshotsWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取规格审批列表
//
// @param request - ListSpecReviewTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSpecReviewTasksResponse
func (client *Client) ListSpecReviewTasksWithOptions(appName *string, request *ListSpecReviewTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSpecReviewTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSpecReviewTasks"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/spec-review-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSpecReviewTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取规格审批列表
//
// @param request - ListSpecReviewTasksRequest
//
// @return ListSpecReviewTasksResponse
func (client *Client) ListSpecReviewTasks(appName *string, request *ListSpecReviewTasksRequest) (_result *ListSpecReviewTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSpecReviewTasksResponse{}
	_body, _err := client.ListSpecReviewTasksWithOptions(appName, request, headers, runtime)
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

// Summary:
//
// 创建或更新词典
//
// @param request - UpdateDictRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDictResponse
func (client *Client) UpdateDictWithOptions(appName *string, request *UpdateDictRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowCover)) {
		query["allowCover"] = request.AllowCover
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDict"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/dicts"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDictResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新词典
//
// @param request - UpdateDictRequest
//
// @return UpdateDictResponse
func (client *Client) UpdateDict(appName *string, request *UpdateDictRequest) (_result *UpdateDictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDictResponse{}
	_body, _err := client.UpdateDictWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改端点信息
//
// @param request - UpdateEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEndpointResponse
func (client *Client) UpdateEndpointWithOptions(endpointId *string, request *UpdateEndpointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointZones)) {
		body["endpointZones"] = request.EndpointZones
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEndpoint"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/endpoints/" + tea.StringValue(openapiutil.GetEncodeParam(endpointId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改端点信息
//
// @param request - UpdateEndpointRequest
//
// @return UpdateEndpointResponse
func (client *Client) UpdateEndpoint(endpointId *string, request *UpdateEndpointRequest) (_result *UpdateEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEndpointResponse{}
	_body, _err := client.UpdateEndpointWithOptions(endpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自动备份配置
//
// @param request - UpdateSnapshotSettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSettingWithOptions(appName *string, request *UpdateSnapshotSettingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSnapshotSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.QuartzRegex)) {
		body["quartzRegex"] = request.QuartzRegex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSnapshotSetting"),
		Version:     tea.String("2023-06-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/es-serverless/instances/" + tea.StringValue(openapiutil.GetEncodeParam(appName)) + "/auto-snapshot-setting"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自动备份配置
//
// @param request - UpdateSnapshotSettingRequest
//
// @return UpdateSnapshotSettingResponse
func (client *Client) UpdateSnapshotSetting(appName *string, request *UpdateSnapshotSettingRequest) (_result *UpdateSnapshotSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotSettingResponse{}
	_body, _err := client.UpdateSnapshotSettingWithOptions(appName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
