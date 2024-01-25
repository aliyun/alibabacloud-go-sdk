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

type CancelOrderRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) SetInstanceId(v string) *CancelOrderRequest {
	s.InstanceId = &v
	return s
}

type CancelOrderResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) SetData(v bool) *CancelOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CancelOrderResponseBody) SetErrCode(v string) *CancelOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CancelOrderResponseBody) SetErrMessage(v string) *CancelOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CancelOrderResponseBody) SetHttpStatusCode(v int32) *CancelOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelOrderResponseBody) SetRequestId(v string) *CancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderResponseBody) SetSuccess(v bool) *CancelOrderResponseBody {
	s.Success = &v
	return s
}

type CancelOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderResponse) SetHeaders(v map[string]*string) *CancelOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderResponse) SetStatusCode(v int32) *CancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderResponse) SetBody(v *CancelOrderResponseBody) *CancelOrderResponse {
	s.Body = v
	return s
}

type CheckClusterNameRequest struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s CheckClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckClusterNameRequest) GoString() string {
	return s.String()
}

func (s *CheckClusterNameRequest) SetClusterName(v string) *CheckClusterNameRequest {
	s.ClusterName = &v
	return s
}

type CheckClusterNameResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckClusterNameResponseBody) SetData(v bool) *CheckClusterNameResponseBody {
	s.Data = &v
	return s
}

func (s *CheckClusterNameResponseBody) SetErrCode(v string) *CheckClusterNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CheckClusterNameResponseBody) SetErrMessage(v string) *CheckClusterNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CheckClusterNameResponseBody) SetHttpStatusCode(v int32) *CheckClusterNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckClusterNameResponseBody) SetRequestId(v string) *CheckClusterNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckClusterNameResponseBody) SetSuccess(v bool) *CheckClusterNameResponseBody {
	s.Success = &v
	return s
}

type CheckClusterNameResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckClusterNameResponse) GoString() string {
	return s.String()
}

func (s *CheckClusterNameResponse) SetHeaders(v map[string]*string) *CheckClusterNameResponse {
	s.Headers = v
	return s
}

func (s *CheckClusterNameResponse) SetStatusCode(v int32) *CheckClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckClusterNameResponse) SetBody(v *CheckClusterNameResponseBody) *CheckClusterNameResponse {
	s.Body = v
	return s
}

type ConfirmNoticeRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s ConfirmNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmNoticeRequest) GoString() string {
	return s.String()
}

func (s *ConfirmNoticeRequest) SetClusterBizId(v string) *ConfirmNoticeRequest {
	s.ClusterBizId = &v
	return s
}

type ConfirmNoticeResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmNoticeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmNoticeResponseBody) SetData(v bool) *ConfirmNoticeResponseBody {
	s.Data = &v
	return s
}

func (s *ConfirmNoticeResponseBody) SetErrCode(v string) *ConfirmNoticeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfirmNoticeResponseBody) SetErrMessage(v string) *ConfirmNoticeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfirmNoticeResponseBody) SetHttpStatusCode(v int32) *ConfirmNoticeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfirmNoticeResponseBody) SetRequestId(v string) *ConfirmNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmNoticeResponseBody) SetSuccess(v bool) *ConfirmNoticeResponseBody {
	s.Success = &v
	return s
}

type ConfirmNoticeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmNoticeResponse) GoString() string {
	return s.String()
}

func (s *ConfirmNoticeResponse) SetHeaders(v map[string]*string) *ConfirmNoticeResponse {
	s.Headers = v
	return s
}

func (s *ConfirmNoticeResponse) SetStatusCode(v int32) *ConfirmNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmNoticeResponse) SetBody(v *ConfirmNoticeResponseBody) *ConfirmNoticeResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClusterInfo *string `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetClusterInfo(v string) *CreateClusterRequest {
	s.ClusterInfo = &v
	return s
}

type CreateClusterResponseBody struct {
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetData(v string) *CreateClusterResponseBody {
	s.Data = &v
	return s
}

func (s *CreateClusterResponseBody) SetErrCode(v string) *CreateClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateClusterResponseBody) SetErrMessage(v string) *CreateClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateClusterResponseBody) SetHttpStatusCode(v int32) *CreateClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type GetClusterDetailRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *GetClusterDetailRequest) SetClusterBizId(v string) *GetClusterDetailRequest {
	s.ClusterBizId = &v
	return s
}

func (s *GetClusterDetailRequest) SetInstanceId(v string) *GetClusterDetailRequest {
	s.InstanceId = &v
	return s
}

type GetClusterDetailResponseBody struct {
	Data           *GetClusterDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode        *string                           `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                           `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetClusterDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterDetailResponseBody) SetData(v *GetClusterDetailResponseBodyData) *GetClusterDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterDetailResponseBody) SetErrCode(v string) *GetClusterDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetClusterDetailResponseBody) SetErrMessage(v string) *GetClusterDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetClusterDetailResponseBody) SetHttpStatusCode(v int32) *GetClusterDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetClusterDetailResponseBody) SetRequestId(v string) *GetClusterDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterDetailResponseBody) SetSuccess(v bool) *GetClusterDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetClusterDetailResponseBody) SetTotal(v int32) *GetClusterDetailResponseBody {
	s.Total = &v
	return s
}

type GetClusterDetailResponseBodyData struct {
	AutoRenew          *bool                  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BeginTime          *int64                 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ClusterBizId       *string                `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	ClusterId          *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName        *string                `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterStatus      *string                `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	ClusterStatusValue *int32                 `json:"ClusterStatusValue,omitempty" xml:"ClusterStatusValue,omitempty"`
	ClusterType        *string                `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ControlCenterUrl   *string                `json:"ControlCenterUrl,omitempty" xml:"ControlCenterUrl,omitempty"`
	DeployMode         *string                `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Duration           *int32                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpireTime         *int64                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate          *int64                 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *int64                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceConf       map[string]interface{} `json:"InstanceConf,omitempty" xml:"InstanceConf,omitempty"`
	NoticeConfirmed    *bool                  `json:"NoticeConfirmed,omitempty" xml:"NoticeConfirmed,omitempty"`
	OrderBizId         *string                `json:"OrderBizId,omitempty" xml:"OrderBizId,omitempty"`
	PackageType        *string                `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PricingCycle       *string                `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId           *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RunningTime        *int64                 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	Version            *string                `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId             *string                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetClusterDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterDetailResponseBodyData) SetAutoRenew(v bool) *GetClusterDetailResponseBodyData {
	s.AutoRenew = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetBeginTime(v int64) *GetClusterDetailResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetClusterBizId(v string) *GetClusterDetailResponseBodyData {
	s.ClusterBizId = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetClusterId(v string) *GetClusterDetailResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetClusterName(v string) *GetClusterDetailResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetClusterStatus(v string) *GetClusterDetailResponseBodyData {
	s.ClusterStatus = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetClusterStatusValue(v int32) *GetClusterDetailResponseBodyData {
	s.ClusterStatusValue = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetClusterType(v string) *GetClusterDetailResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetControlCenterUrl(v string) *GetClusterDetailResponseBodyData {
	s.ControlCenterUrl = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetDeployMode(v string) *GetClusterDetailResponseBodyData {
	s.DeployMode = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetDuration(v int32) *GetClusterDetailResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetExpireTime(v int64) *GetClusterDetailResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetGmtCreate(v int64) *GetClusterDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetGmtModified(v int64) *GetClusterDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetInstanceConf(v map[string]interface{}) *GetClusterDetailResponseBodyData {
	s.InstanceConf = v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetNoticeConfirmed(v bool) *GetClusterDetailResponseBodyData {
	s.NoticeConfirmed = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetOrderBizId(v string) *GetClusterDetailResponseBodyData {
	s.OrderBizId = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetPackageType(v string) *GetClusterDetailResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetPricingCycle(v string) *GetClusterDetailResponseBodyData {
	s.PricingCycle = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetRegionId(v string) *GetClusterDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetRunningTime(v int64) *GetClusterDetailResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetVersion(v string) *GetClusterDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetClusterDetailResponseBodyData) SetZoneId(v string) *GetClusterDetailResponseBodyData {
	s.ZoneId = &v
	return s
}

type GetClusterDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *GetClusterDetailResponse) SetHeaders(v map[string]*string) *GetClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *GetClusterDetailResponse) SetStatusCode(v int32) *GetClusterDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterDetailResponse) SetBody(v *GetClusterDetailResponseBody) *GetClusterDetailResponse {
	s.Body = v
	return s
}

type HasDefaultRoleResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HasDefaultRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HasDefaultRoleResponseBody) GoString() string {
	return s.String()
}

func (s *HasDefaultRoleResponseBody) SetData(v bool) *HasDefaultRoleResponseBody {
	s.Data = &v
	return s
}

func (s *HasDefaultRoleResponseBody) SetErrCode(v string) *HasDefaultRoleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *HasDefaultRoleResponseBody) SetErrMessage(v string) *HasDefaultRoleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *HasDefaultRoleResponseBody) SetHttpStatusCode(v int32) *HasDefaultRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HasDefaultRoleResponseBody) SetRequestId(v string) *HasDefaultRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *HasDefaultRoleResponseBody) SetSuccess(v bool) *HasDefaultRoleResponseBody {
	s.Success = &v
	return s
}

type HasDefaultRoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HasDefaultRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HasDefaultRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s HasDefaultRoleResponse) GoString() string {
	return s.String()
}

func (s *HasDefaultRoleResponse) SetHeaders(v map[string]*string) *HasDefaultRoleResponse {
	s.Headers = v
	return s
}

func (s *HasDefaultRoleResponse) SetStatusCode(v int32) *HasDefaultRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *HasDefaultRoleResponse) SetBody(v *HasDefaultRoleResponseBody) *HasDefaultRoleResponse {
	s.Body = v
	return s
}

type InitializeClouderaDataPlatformResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitializeClouderaDataPlatformResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeClouderaDataPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeClouderaDataPlatformResponseBody) SetData(v bool) *InitializeClouderaDataPlatformResponseBody {
	s.Data = &v
	return s
}

func (s *InitializeClouderaDataPlatformResponseBody) SetErrCode(v string) *InitializeClouderaDataPlatformResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InitializeClouderaDataPlatformResponseBody) SetErrMessage(v string) *InitializeClouderaDataPlatformResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *InitializeClouderaDataPlatformResponseBody) SetHttpStatusCode(v int32) *InitializeClouderaDataPlatformResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitializeClouderaDataPlatformResponseBody) SetRequestId(v string) *InitializeClouderaDataPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeClouderaDataPlatformResponseBody) SetSuccess(v bool) *InitializeClouderaDataPlatformResponseBody {
	s.Success = &v
	return s
}

type InitializeClouderaDataPlatformResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeClouderaDataPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeClouderaDataPlatformResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeClouderaDataPlatformResponse) GoString() string {
	return s.String()
}

func (s *InitializeClouderaDataPlatformResponse) SetHeaders(v map[string]*string) *InitializeClouderaDataPlatformResponse {
	s.Headers = v
	return s
}

func (s *InitializeClouderaDataPlatformResponse) SetStatusCode(v int32) *InitializeClouderaDataPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeClouderaDataPlatformResponse) SetBody(v *InitializeClouderaDataPlatformResponseBody) *InitializeClouderaDataPlatformResponse {
	s.Body = v
	return s
}

type ListDefaultComponentsRequest struct {
	ClusterType  *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	SecurityMode *string `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
}

func (s ListDefaultComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListDefaultComponentsRequest) SetClusterType(v string) *ListDefaultComponentsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDefaultComponentsRequest) SetSecurityMode(v string) *ListDefaultComponentsRequest {
	s.SecurityMode = &v
	return s
}

type ListDefaultComponentsResponseBody struct {
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *string   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDefaultComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDefaultComponentsResponseBody) SetData(v []*string) *ListDefaultComponentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDefaultComponentsResponseBody) SetErrCode(v string) *ListDefaultComponentsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDefaultComponentsResponseBody) SetErrMessage(v string) *ListDefaultComponentsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDefaultComponentsResponseBody) SetHttpStatusCode(v string) *ListDefaultComponentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDefaultComponentsResponseBody) SetRequestId(v string) *ListDefaultComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDefaultComponentsResponseBody) SetSuccess(v bool) *ListDefaultComponentsResponseBody {
	s.Success = &v
	return s
}

type ListDefaultComponentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDefaultComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDefaultComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListDefaultComponentsResponse) SetHeaders(v map[string]*string) *ListDefaultComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListDefaultComponentsResponse) SetStatusCode(v int32) *ListDefaultComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDefaultComponentsResponse) SetBody(v *ListDefaultComponentsResponseBody) *ListDefaultComponentsResponse {
	s.Body = v
	return s
}

type ListNodeGroupConstraintsRequest struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s ListNodeGroupConstraintsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupConstraintsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupConstraintsRequest) SetClusterType(v string) *ListNodeGroupConstraintsRequest {
	s.ClusterType = &v
	return s
}

type ListNodeGroupConstraintsResponseBody struct {
	Data           []*ListNodeGroupConstraintsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string                                     `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                                     `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodeGroupConstraintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupConstraintsResponseBody) SetData(v []*ListNodeGroupConstraintsResponseBodyData) *ListNodeGroupConstraintsResponseBody {
	s.Data = v
	return s
}

func (s *ListNodeGroupConstraintsResponseBody) SetErrCode(v string) *ListNodeGroupConstraintsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBody) SetErrMessage(v string) *ListNodeGroupConstraintsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBody) SetHttpStatusCode(v int32) *ListNodeGroupConstraintsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBody) SetRequestId(v string) *ListNodeGroupConstraintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBody) SetSuccess(v bool) *ListNodeGroupConstraintsResponseBody {
	s.Success = &v
	return s
}

type ListNodeGroupConstraintsResponseBodyData struct {
	HostGroupType                 *string   `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	RecommendedInstanceTypes      []*string `json:"RecommendedInstanceTypes,omitempty" xml:"RecommendedInstanceTypes,omitempty" type:"Repeated"`
	AvailableDataDiskCategories   []*string `json:"availableDataDiskCategories,omitempty" xml:"availableDataDiskCategories,omitempty" type:"Repeated"`
	AvailableInstanceTypes        []*string `json:"availableInstanceTypes,omitempty" xml:"availableInstanceTypes,omitempty" type:"Repeated"`
	AvailableSystemDiskCategories []*string `json:"availableSystemDiskCategories,omitempty" xml:"availableSystemDiskCategories,omitempty" type:"Repeated"`
	DefaultDataDiskCount          *int32    `json:"defaultDataDiskCount,omitempty" xml:"defaultDataDiskCount,omitempty"`
	DefaultDataDiskSize           *int32    `json:"defaultDataDiskSize,omitempty" xml:"defaultDataDiskSize,omitempty"`
	DefaultNodeCount              *int32    `json:"defaultNodeCount,omitempty" xml:"defaultNodeCount,omitempty"`
	DefaultSystemDiskSize         *int32    `json:"defaultSystemDiskSize,omitempty" xml:"defaultSystemDiskSize,omitempty"`
	MaxDataDiskCount              *int32    `json:"maxDataDiskCount,omitempty" xml:"maxDataDiskCount,omitempty"`
	MaxDataDiskSize               *int32    `json:"maxDataDiskSize,omitempty" xml:"maxDataDiskSize,omitempty"`
	MaxNodeCount                  *int32    `json:"maxNodeCount,omitempty" xml:"maxNodeCount,omitempty"`
	MaxSystemDiskSize             *int32    `json:"maxSystemDiskSize,omitempty" xml:"maxSystemDiskSize,omitempty"`
	MinDataDiskCount              *int32    `json:"minDataDiskCount,omitempty" xml:"minDataDiskCount,omitempty"`
	MinDataDiskSize               *int32    `json:"minDataDiskSize,omitempty" xml:"minDataDiskSize,omitempty"`
	MinNodeCount                  *int32    `json:"minNodeCount,omitempty" xml:"minNodeCount,omitempty"`
	MinSystemDiskSize             *int32    `json:"minSystemDiskSize,omitempty" xml:"minSystemDiskSize,omitempty"`
	NodeGroupName                 *string   `json:"nodeGroupName,omitempty" xml:"nodeGroupName,omitempty"`
}

func (s ListNodeGroupConstraintsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupConstraintsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetHostGroupType(v string) *ListNodeGroupConstraintsResponseBodyData {
	s.HostGroupType = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetRecommendedInstanceTypes(v []*string) *ListNodeGroupConstraintsResponseBodyData {
	s.RecommendedInstanceTypes = v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetAvailableDataDiskCategories(v []*string) *ListNodeGroupConstraintsResponseBodyData {
	s.AvailableDataDiskCategories = v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetAvailableInstanceTypes(v []*string) *ListNodeGroupConstraintsResponseBodyData {
	s.AvailableInstanceTypes = v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetAvailableSystemDiskCategories(v []*string) *ListNodeGroupConstraintsResponseBodyData {
	s.AvailableSystemDiskCategories = v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetDefaultDataDiskCount(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.DefaultDataDiskCount = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetDefaultDataDiskSize(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.DefaultDataDiskSize = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetDefaultNodeCount(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.DefaultNodeCount = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetDefaultSystemDiskSize(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.DefaultSystemDiskSize = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMaxDataDiskCount(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MaxDataDiskCount = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMaxDataDiskSize(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MaxDataDiskSize = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMaxNodeCount(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MaxNodeCount = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMaxSystemDiskSize(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MaxSystemDiskSize = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMinDataDiskCount(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MinDataDiskCount = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMinDataDiskSize(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MinDataDiskSize = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMinNodeCount(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MinNodeCount = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetMinSystemDiskSize(v int32) *ListNodeGroupConstraintsResponseBodyData {
	s.MinSystemDiskSize = &v
	return s
}

func (s *ListNodeGroupConstraintsResponseBodyData) SetNodeGroupName(v string) *ListNodeGroupConstraintsResponseBodyData {
	s.NodeGroupName = &v
	return s
}

type ListNodeGroupConstraintsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupConstraintsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupConstraintsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupConstraintsResponse) SetHeaders(v map[string]*string) *ListNodeGroupConstraintsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupConstraintsResponse) SetStatusCode(v int32) *ListNodeGroupConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupConstraintsResponse) SetBody(v *ListNodeGroupConstraintsResponseBody) *ListNodeGroupConstraintsResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetClusterBizId(v string) *ListNodesRequest {
	s.ClusterBizId = &v
	return s
}

type ListNodesResponseBody struct {
	Data           []*ListNodesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string                      `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                      `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetData(v []*ListNodesResponseBodyData) *ListNodesResponseBody {
	s.Data = v
	return s
}

func (s *ListNodesResponseBody) SetErrCode(v string) *ListNodesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListNodesResponseBody) SetErrMessage(v string) *ListNodesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListNodesResponseBody) SetHttpStatusCode(v int32) *ListNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetSuccess(v bool) *ListNodesResponseBody {
	s.Success = &v
	return s
}

type ListNodesResponseBodyData struct {
	CreateTime     *int64                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EcsNodeDtoList []*ListNodesResponseBodyDataEcsNodeDtoList `json:"EcsNodeDtoList,omitempty" xml:"EcsNodeDtoList,omitempty" type:"Repeated"`
	ExpireTime     *int64                                     `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceConf   map[string]interface{}                     `json:"InstanceConf,omitempty" xml:"InstanceConf,omitempty"`
	InstanceId     *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName   *string                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ListNodesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyData) SetCreateTime(v int64) *ListNodesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyData) SetEcsNodeDtoList(v []*ListNodesResponseBodyDataEcsNodeDtoList) *ListNodesResponseBodyData {
	s.EcsNodeDtoList = v
	return s
}

func (s *ListNodesResponseBodyData) SetExpireTime(v int64) *ListNodesResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListNodesResponseBodyData) SetInstanceConf(v map[string]interface{}) *ListNodesResponseBodyData {
	s.InstanceConf = v
	return s
}

func (s *ListNodesResponseBodyData) SetInstanceId(v string) *ListNodesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListNodesResponseBodyData) SetInstanceName(v string) *ListNodesResponseBodyData {
	s.InstanceName = &v
	return s
}

type ListNodesResponseBodyDataEcsNodeDtoList struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CpuCount           *int32  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	ExpireTime         *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Index              *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MemorySize         *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	NodeGroupId        *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName      *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeGroupType      *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	NodeId             *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeName           *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeResourceType   *string `json:"NodeResourceType,omitempty" xml:"NodeResourceType,omitempty"`
	NodeStatus         *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	PrivateIp          *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp           *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	RunningTime        *int64  `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SystemDiskCapacity *int32  `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
	SystemDiskCount    *int32  `json:"SystemDiskCount,omitempty" xml:"SystemDiskCount,omitempty"`
	SystemDiskType     *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
}

func (s ListNodesResponseBodyDataEcsNodeDtoList) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyDataEcsNodeDtoList) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetBeginTime(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.BeginTime = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetCpuCount(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.CpuCount = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetDiskCapacity(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.DiskCapacity = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetDiskCount(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.DiskCount = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetDiskType(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.DiskType = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetExpireTime(v int64) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.ExpireTime = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetGmtCreate(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.GmtCreate = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetGmtModified(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.GmtModified = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetIndex(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.Index = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetInstanceType(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.InstanceType = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetMemorySize(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.MemorySize = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeGroupId(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeGroupName(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeGroupName = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeGroupType(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeGroupType = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeId(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeId = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeName(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeName = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeResourceType(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeResourceType = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetNodeStatus(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.NodeStatus = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetPrivateIp(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.PrivateIp = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetPublicIp(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.PublicIp = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetRunningTime(v int64) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.RunningTime = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetSerialNumber(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.SerialNumber = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetSystemDiskCapacity(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.SystemDiskCapacity = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetSystemDiskCount(v int32) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.SystemDiskCount = &v
	return s
}

func (s *ListNodesResponseBodyDataEcsNodeDtoList) SetSystemDiskType(v string) *ListNodesResponseBodyDataEcsNodeDtoList {
	s.SystemDiskType = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListOperationsRequest struct {
	ClusterBizId          *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	ParentOperationNodeId *string `json:"ParentOperationNodeId,omitempty" xml:"ParentOperationNodeId,omitempty"`
}

func (s ListOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationsRequest) SetClusterBizId(v string) *ListOperationsRequest {
	s.ClusterBizId = &v
	return s
}

func (s *ListOperationsRequest) SetParentOperationNodeId(v string) *ListOperationsRequest {
	s.ParentOperationNodeId = &v
	return s
}

type ListOperationsResponseBody struct {
	Data           []*ListOperationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string                           `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                           `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationsResponseBody) SetData(v []*ListOperationsResponseBodyData) *ListOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationsResponseBody) SetErrCode(v string) *ListOperationsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListOperationsResponseBody) SetErrMessage(v string) *ListOperationsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListOperationsResponseBody) SetHttpStatusCode(v int32) *ListOperationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOperationsResponseBody) SetRequestId(v string) *ListOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationsResponseBody) SetSuccess(v bool) *ListOperationsResponseBody {
	s.Success = &v
	return s
}

type ListOperationsResponseBodyData struct {
	EndTime                *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HasChildOperationNodes *bool   `json:"HasChildOperationNodes,omitempty" xml:"HasChildOperationNodes,omitempty"`
	HasOperationTask       *bool   `json:"HasOperationTask,omitempty" xml:"HasOperationTask,omitempty"`
	OperationId            *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	OperationNodeId        *string `json:"OperationNodeId,omitempty" xml:"OperationNodeId,omitempty"`
	OperationNodeName      *int32  `json:"OperationNodeName,omitempty" xml:"OperationNodeName,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOperationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationsResponseBodyData) SetEndTime(v int64) *ListOperationsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetHasChildOperationNodes(v bool) *ListOperationsResponseBodyData {
	s.HasChildOperationNodes = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetHasOperationTask(v bool) *ListOperationsResponseBodyData {
	s.HasOperationTask = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetOperationId(v string) *ListOperationsResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetOperationNodeId(v string) *ListOperationsResponseBodyData {
	s.OperationNodeId = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetOperationNodeName(v int32) *ListOperationsResponseBodyData {
	s.OperationNodeName = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetStartTime(v int64) *ListOperationsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListOperationsResponseBodyData) SetStatus(v string) *ListOperationsResponseBodyData {
	s.Status = &v
	return s
}

type ListOperationsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationsResponse) SetHeaders(v map[string]*string) *ListOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationsResponse) SetStatusCode(v int32) *ListOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationsResponse) SetBody(v *ListOperationsResponseBody) *ListOperationsResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	Data           []*ListRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                        `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetData(v []*ListRegionsResponseBodyData) *ListRegionsResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionsResponseBody) SetErrCode(v string) *ListRegionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListRegionsResponseBody) SetErrMessage(v string) *ListRegionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListRegionsResponseBody) SetHttpStatusCode(v int32) *ListRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

type ListRegionsResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ListRegionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyData) SetDescription(v string) *ListRegionsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetRegionId(v string) *ListRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetRegionName(v string) *ListRegionsResponseBodyData {
	s.RegionName = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListZonesRequest) GoString() string {
	return s.String()
}

func (s *ListZonesRequest) SetRegionId(v string) *ListZonesRequest {
	s.RegionId = &v
	return s
}

type ListZonesResponseBody struct {
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrMessage     *string   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ErrorCode      *string   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBody) SetData(v []*string) *ListZonesResponseBody {
	s.Data = v
	return s
}

func (s *ListZonesResponseBody) SetErrMessage(v string) *ListZonesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListZonesResponseBody) SetErrorCode(v string) *ListZonesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListZonesResponseBody) SetHttpStatusCode(v int32) *ListZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListZonesResponseBody) SetRequestId(v string) *ListZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZonesResponseBody) SetSuccess(v bool) *ListZonesResponseBody {
	s.Success = &v
	return s
}

type ListZonesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponse) GoString() string {
	return s.String()
}

func (s *ListZonesResponse) SetHeaders(v map[string]*string) *ListZonesResponse {
	s.Headers = v
	return s
}

func (s *ListZonesResponse) SetStatusCode(v int32) *ListZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZonesResponse) SetBody(v *ListZonesResponseBody) *ListZonesResponse {
	s.Body = v
	return s
}

type QueryOrderRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s QueryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderRequest) SetClusterBizId(v string) *QueryOrderRequest {
	s.ClusterBizId = &v
	return s
}

type QueryOrderResponseBody struct {
	Data           []*QueryOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string                       `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                       `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBody) SetData(v []*QueryOrderResponseBodyData) *QueryOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrderResponseBody) SetErrCode(v string) *QueryOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryOrderResponseBody) SetErrMessage(v string) *QueryOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryOrderResponseBody) SetHttpStatusCode(v int32) *QueryOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryOrderResponseBody) SetRequestId(v string) *QueryOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderResponseBody) SetSuccess(v bool) *QueryOrderResponseBody {
	s.Success = &v
	return s
}

type QueryOrderResponseBodyData struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	OrderId     *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus *string   `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType   *string   `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s QueryOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyData) SetInstanceIds(v []*string) *QueryOrderResponseBodyData {
	s.InstanceIds = v
	return s
}

func (s *QueryOrderResponseBodyData) SetOrderId(v string) *QueryOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetOrderStatus(v string) *QueryOrderResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetOrderType(v string) *QueryOrderResponseBodyData {
	s.OrderType = &v
	return s
}

type QueryOrderResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderResponse) SetHeaders(v map[string]*string) *QueryOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderResponse) SetStatusCode(v int32) *QueryOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderResponse) SetBody(v *QueryOrderResponseBody) *QueryOrderResponse {
	s.Body = v
	return s
}

type QueryPriceRequest struct {
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeGroupSpecs *string `json:"NodeGroupSpecs,omitempty" xml:"NodeGroupSpecs,omitempty"`
	PricingCycle   *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryPriceRequest) SetDuration(v int32) *QueryPriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryPriceRequest) SetInstanceId(v string) *QueryPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryPriceRequest) SetNodeGroupSpecs(v string) *QueryPriceRequest {
	s.NodeGroupSpecs = &v
	return s
}

func (s *QueryPriceRequest) SetPricingCycle(v string) *QueryPriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryPriceRequest) SetRegionId(v string) *QueryPriceRequest {
	s.RegionId = &v
	return s
}

type QueryPriceResponseBody struct {
	Data           *QueryPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode        *string                     `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                     `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPriceResponseBody) SetData(v *QueryPriceResponseBodyData) *QueryPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryPriceResponseBody) SetErrCode(v string) *QueryPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryPriceResponseBody) SetErrMessage(v string) *QueryPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryPriceResponseBody) SetHttpStatusCode(v int32) *QueryPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryPriceResponseBody) SetRequestId(v string) *QueryPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPriceResponseBody) SetSuccess(v bool) *QueryPriceResponseBody {
	s.Success = &v
	return s
}

type QueryPriceResponseBodyData struct {
	DiscountPrice *float64                                 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	EcsPriceInfo  *QueryPriceResponseBodyDataEcsPriceInfo  `json:"EcsPriceInfo,omitempty" xml:"EcsPriceInfo,omitempty" type:"Struct"`
	SoftPriceInfo *QueryPriceResponseBodyDataSoftPriceInfo `json:"SoftPriceInfo,omitempty" xml:"SoftPriceInfo,omitempty" type:"Struct"`
	SumPrice      *float64                                 `json:"SumPrice,omitempty" xml:"SumPrice,omitempty"`
}

func (s QueryPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPriceResponseBodyData) SetDiscountPrice(v float64) *QueryPriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *QueryPriceResponseBodyData) SetEcsPriceInfo(v *QueryPriceResponseBodyDataEcsPriceInfo) *QueryPriceResponseBodyData {
	s.EcsPriceInfo = v
	return s
}

func (s *QueryPriceResponseBodyData) SetSoftPriceInfo(v *QueryPriceResponseBodyDataSoftPriceInfo) *QueryPriceResponseBodyData {
	s.SoftPriceInfo = v
	return s
}

func (s *QueryPriceResponseBodyData) SetSumPrice(v float64) *QueryPriceResponseBodyData {
	s.SumPrice = &v
	return s
}

type QueryPriceResponseBodyDataEcsPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryPriceResponseBodyDataEcsPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryPriceResponseBodyDataEcsPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryPriceResponseBodyDataEcsPriceInfo) SetCurrency(v string) *QueryPriceResponseBodyDataEcsPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryPriceResponseBodyDataEcsPriceInfo) SetDiscountPrice(v float64) *QueryPriceResponseBodyDataEcsPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *QueryPriceResponseBodyDataEcsPriceInfo) SetOriginalPrice(v float64) *QueryPriceResponseBodyDataEcsPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *QueryPriceResponseBodyDataEcsPriceInfo) SetTradePrice(v float64) *QueryPriceResponseBodyDataEcsPriceInfo {
	s.TradePrice = &v
	return s
}

type QueryPriceResponseBodyDataSoftPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryPriceResponseBodyDataSoftPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryPriceResponseBodyDataSoftPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryPriceResponseBodyDataSoftPriceInfo) SetCurrency(v string) *QueryPriceResponseBodyDataSoftPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryPriceResponseBodyDataSoftPriceInfo) SetDiscountPrice(v float64) *QueryPriceResponseBodyDataSoftPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *QueryPriceResponseBodyDataSoftPriceInfo) SetOriginalPrice(v float64) *QueryPriceResponseBodyDataSoftPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *QueryPriceResponseBodyDataSoftPriceInfo) SetTradePrice(v float64) *QueryPriceResponseBodyDataSoftPriceInfo {
	s.TradePrice = &v
	return s
}

type QueryPriceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryPriceResponse) SetHeaders(v map[string]*string) *QueryPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryPriceResponse) SetStatusCode(v int32) *QueryPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPriceResponse) SetBody(v *QueryPriceResponseBody) *QueryPriceResponse {
	s.Body = v
	return s
}

type QueryRenewOrderRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s QueryRenewOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewOrderRequest) SetClusterBizId(v string) *QueryRenewOrderRequest {
	s.ClusterBizId = &v
	return s
}

type QueryRenewOrderResponseBody struct {
	Data           []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrMessage     *string  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ErrorCode      *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRenewOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRenewOrderResponseBody) SetData(v []*int64) *QueryRenewOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryRenewOrderResponseBody) SetErrMessage(v string) *QueryRenewOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryRenewOrderResponseBody) SetErrorCode(v string) *QueryRenewOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryRenewOrderResponseBody) SetHttpStatusCode(v int32) *QueryRenewOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRenewOrderResponseBody) SetRequestId(v string) *QueryRenewOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRenewOrderResponseBody) SetSuccess(v bool) *QueryRenewOrderResponseBody {
	s.Success = &v
	return s
}

type QueryRenewOrderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRenewOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRenewOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryRenewOrderResponse) SetHeaders(v map[string]*string) *QueryRenewOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryRenewOrderResponse) SetStatusCode(v int32) *QueryRenewOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRenewOrderResponse) SetBody(v *QueryRenewOrderResponseBody) *QueryRenewOrderResponse {
	s.Body = v
	return s
}

type QueryRenewPriceRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	Instances    *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
}

func (s QueryRenewPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceRequest) SetClusterBizId(v string) *QueryRenewPriceRequest {
	s.ClusterBizId = &v
	return s
}

func (s *QueryRenewPriceRequest) SetInstances(v string) *QueryRenewPriceRequest {
	s.Instances = &v
	return s
}

type QueryRenewPriceResponseBody struct {
	Data           *QueryRenewPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode        *string                          `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                          `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRenewPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBody) SetData(v *QueryRenewPriceResponseBodyData) *QueryRenewPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryRenewPriceResponseBody) SetErrCode(v string) *QueryRenewPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetErrMessage(v string) *QueryRenewPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetHttpStatusCode(v int32) *QueryRenewPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetRequestId(v string) *QueryRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRenewPriceResponseBody) SetSuccess(v bool) *QueryRenewPriceResponseBody {
	s.Success = &v
	return s
}

type QueryRenewPriceResponseBodyData struct {
	CdpSoftPriceInfo *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo `json:"CdpSoftPriceInfo,omitempty" xml:"CdpSoftPriceInfo,omitempty" type:"Struct"`
	DiscountPrice    *float32                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	EcsPriceInfo     *QueryRenewPriceResponseBodyDataEcsPriceInfo     `json:"EcsPriceInfo,omitempty" xml:"EcsPriceInfo,omitempty" type:"Struct"`
	SumPrice         *float32                                         `json:"SumPrice,omitempty" xml:"SumPrice,omitempty"`
}

func (s QueryRenewPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyData) SetCdpSoftPriceInfo(v *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) *QueryRenewPriceResponseBodyData {
	s.CdpSoftPriceInfo = v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetDiscountPrice(v float32) *QueryRenewPriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetEcsPriceInfo(v *QueryRenewPriceResponseBodyDataEcsPriceInfo) *QueryRenewPriceResponseBodyData {
	s.EcsPriceInfo = v
	return s
}

func (s *QueryRenewPriceResponseBodyData) SetSumPrice(v float32) *QueryRenewPriceResponseBodyData {
	s.SumPrice = &v
	return s
}

type QueryRenewPriceResponseBodyDataCdpSoftPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) SetCurrency(v string) *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) SetDiscountPrice(v float32) *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) SetOriginalPrice(v float32) *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo) SetTradePrice(v float32) *QueryRenewPriceResponseBodyDataCdpSoftPriceInfo {
	s.TradePrice = &v
	return s
}

type QueryRenewPriceResponseBodyDataEcsPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryRenewPriceResponseBodyDataEcsPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewPriceResponseBodyDataEcsPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponseBodyDataEcsPriceInfo) SetCurrency(v string) *QueryRenewPriceResponseBodyDataEcsPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataEcsPriceInfo) SetDiscountPrice(v float32) *QueryRenewPriceResponseBodyDataEcsPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataEcsPriceInfo) SetOriginalPrice(v float32) *QueryRenewPriceResponseBodyDataEcsPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *QueryRenewPriceResponseBodyDataEcsPriceInfo) SetTradePrice(v float32) *QueryRenewPriceResponseBodyDataEcsPriceInfo {
	s.TradePrice = &v
	return s
}

type QueryRenewPriceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRenewPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceResponse) SetHeaders(v map[string]*string) *QueryRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryRenewPriceResponse) SetStatusCode(v int32) *QueryRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRenewPriceResponse) SetBody(v *QueryRenewPriceResponseBody) *QueryRenewPriceResponse {
	s.Body = v
	return s
}

type QueryScaleUpOrderRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryScaleUpOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryScaleUpOrderRequest) SetClusterBizId(v string) *QueryScaleUpOrderRequest {
	s.ClusterBizId = &v
	return s
}

func (s *QueryScaleUpOrderRequest) SetInstanceId(v string) *QueryScaleUpOrderRequest {
	s.InstanceId = &v
	return s
}

type QueryScaleUpOrderResponseBody struct {
	Data           []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int64   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryScaleUpOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScaleUpOrderResponseBody) SetData(v []*int64) *QueryScaleUpOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryScaleUpOrderResponseBody) SetErrCode(v string) *QueryScaleUpOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryScaleUpOrderResponseBody) SetErrMessage(v string) *QueryScaleUpOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryScaleUpOrderResponseBody) SetHttpStatusCode(v int64) *QueryScaleUpOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryScaleUpOrderResponseBody) SetRequestId(v string) *QueryScaleUpOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScaleUpOrderResponseBody) SetSuccess(v bool) *QueryScaleUpOrderResponseBody {
	s.Success = &v
	return s
}

type QueryScaleUpOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScaleUpOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScaleUpOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryScaleUpOrderResponse) SetHeaders(v map[string]*string) *QueryScaleUpOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryScaleUpOrderResponse) SetStatusCode(v int32) *QueryScaleUpOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScaleUpOrderResponse) SetBody(v *QueryScaleUpOrderResponseBody) *QueryScaleUpOrderResponse {
	s.Body = v
	return s
}

type QueryScaleUpPriceRequest struct {
	ClusterBizId  *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	CoreNumber    *int64  `json:"CoreNumber,omitempty" xml:"CoreNumber,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s QueryScaleUpPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryScaleUpPriceRequest) SetClusterBizId(v string) *QueryScaleUpPriceRequest {
	s.ClusterBizId = &v
	return s
}

func (s *QueryScaleUpPriceRequest) SetCoreNumber(v int64) *QueryScaleUpPriceRequest {
	s.CoreNumber = &v
	return s
}

func (s *QueryScaleUpPriceRequest) SetDuration(v int64) *QueryScaleUpPriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryScaleUpPriceRequest) SetInstanceId(v string) *QueryScaleUpPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryScaleUpPriceRequest) SetInstanceType(v string) *QueryScaleUpPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *QueryScaleUpPriceRequest) SetNodeGroupType(v string) *QueryScaleUpPriceRequest {
	s.NodeGroupType = &v
	return s
}

func (s *QueryScaleUpPriceRequest) SetPricingCycle(v string) *QueryScaleUpPriceRequest {
	s.PricingCycle = &v
	return s
}

type QueryScaleUpPriceResponseBody struct {
	Data           *QueryScaleUpPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode        *string                            `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                            `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int64                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryScaleUpPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScaleUpPriceResponseBody) SetData(v *QueryScaleUpPriceResponseBodyData) *QueryScaleUpPriceResponseBody {
	s.Data = v
	return s
}

func (s *QueryScaleUpPriceResponseBody) SetErrCode(v string) *QueryScaleUpPriceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryScaleUpPriceResponseBody) SetErrMessage(v string) *QueryScaleUpPriceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryScaleUpPriceResponseBody) SetHttpStatusCode(v int64) *QueryScaleUpPriceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryScaleUpPriceResponseBody) SetRequestId(v string) *QueryScaleUpPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScaleUpPriceResponseBody) SetSuccess(v bool) *QueryScaleUpPriceResponseBody {
	s.Success = &v
	return s
}

type QueryScaleUpPriceResponseBodyData struct {
	DiscountPrice *float32                                        `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	EcsPriceInfo  *QueryScaleUpPriceResponseBodyDataEcsPriceInfo  `json:"EcsPriceInfo,omitempty" xml:"EcsPriceInfo,omitempty" type:"Struct"`
	SoftPriceInfo *QueryScaleUpPriceResponseBodyDataSoftPriceInfo `json:"SoftPriceInfo,omitempty" xml:"SoftPriceInfo,omitempty" type:"Struct"`
	SumPrice      *float32                                        `json:"SumPrice,omitempty" xml:"SumPrice,omitempty"`
}

func (s QueryScaleUpPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryScaleUpPriceResponseBodyData) SetDiscountPrice(v float32) *QueryScaleUpPriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyData) SetEcsPriceInfo(v *QueryScaleUpPriceResponseBodyDataEcsPriceInfo) *QueryScaleUpPriceResponseBodyData {
	s.EcsPriceInfo = v
	return s
}

func (s *QueryScaleUpPriceResponseBodyData) SetSoftPriceInfo(v *QueryScaleUpPriceResponseBodyDataSoftPriceInfo) *QueryScaleUpPriceResponseBodyData {
	s.SoftPriceInfo = v
	return s
}

func (s *QueryScaleUpPriceResponseBodyData) SetSumPrice(v float32) *QueryScaleUpPriceResponseBodyData {
	s.SumPrice = &v
	return s
}

type QueryScaleUpPriceResponseBodyDataEcsPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryScaleUpPriceResponseBodyDataEcsPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpPriceResponseBodyDataEcsPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryScaleUpPriceResponseBodyDataEcsPriceInfo) SetCurrency(v string) *QueryScaleUpPriceResponseBodyDataEcsPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyDataEcsPriceInfo) SetDiscountPrice(v float32) *QueryScaleUpPriceResponseBodyDataEcsPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyDataEcsPriceInfo) SetOriginalPrice(v float32) *QueryScaleUpPriceResponseBodyDataEcsPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyDataEcsPriceInfo) SetTradePrice(v float32) *QueryScaleUpPriceResponseBodyDataEcsPriceInfo {
	s.TradePrice = &v
	return s
}

type QueryScaleUpPriceResponseBodyDataSoftPriceInfo struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s QueryScaleUpPriceResponseBodyDataSoftPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpPriceResponseBodyDataSoftPriceInfo) GoString() string {
	return s.String()
}

func (s *QueryScaleUpPriceResponseBodyDataSoftPriceInfo) SetCurrency(v string) *QueryScaleUpPriceResponseBodyDataSoftPriceInfo {
	s.Currency = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyDataSoftPriceInfo) SetDiscountPrice(v float32) *QueryScaleUpPriceResponseBodyDataSoftPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyDataSoftPriceInfo) SetOriginalPrice(v float32) *QueryScaleUpPriceResponseBodyDataSoftPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *QueryScaleUpPriceResponseBodyDataSoftPriceInfo) SetTradePrice(v float32) *QueryScaleUpPriceResponseBodyDataSoftPriceInfo {
	s.TradePrice = &v
	return s
}

type QueryScaleUpPriceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScaleUpPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScaleUpPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryScaleUpPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryScaleUpPriceResponse) SetHeaders(v map[string]*string) *QueryScaleUpPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryScaleUpPriceResponse) SetStatusCode(v int32) *QueryScaleUpPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScaleUpPriceResponse) SetBody(v *QueryScaleUpPriceResponseBody) *QueryScaleUpPriceResponse {
	s.Body = v
	return s
}

type ReleaseClusterRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterRequest) SetInstanceId(v string) *ReleaseClusterRequest {
	s.InstanceId = &v
	return s
}

type ReleaseClusterResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterResponseBody) SetData(v bool) *ReleaseClusterResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseClusterResponseBody) SetErrCode(v string) *ReleaseClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReleaseClusterResponseBody) SetErrMessage(v string) *ReleaseClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ReleaseClusterResponseBody) SetHttpStatusCode(v int32) *ReleaseClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseClusterResponseBody) SetRequestId(v string) *ReleaseClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseClusterResponseBody) SetSuccess(v bool) *ReleaseClusterResponseBody {
	s.Success = &v
	return s
}

type ReleaseClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterResponse) SetHeaders(v map[string]*string) *ReleaseClusterResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterResponse) SetStatusCode(v int32) *ReleaseClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseClusterResponse) SetBody(v *ReleaseClusterResponseBody) *ReleaseClusterResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	Instances    *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClusterBizId(v string) *RenewInstanceRequest {
	s.ClusterBizId = &v
	return s
}

func (s *RenewInstanceRequest) SetInstances(v string) *RenewInstanceRequest {
	s.Instances = &v
	return s
}

type RenewInstanceResponseBody struct {
	Data           *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode        *string                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                        `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RenewInstanceResponseBody) SetErrCode(v string) *RenewInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrMessage(v string) *RenewInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetHttpStatusCode(v int32) *RenewInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

type RenewInstanceResponseBodyData struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s RenewInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBodyData) SetOrderIds(v []*string) *RenewInstanceResponseBodyData {
	s.OrderIds = v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ScaleUpClusterRequest struct {
	ClusterBizId  *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	CoreNumber    *int64  `json:"CoreNumber,omitempty" xml:"CoreNumber,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s ScaleUpClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleUpClusterRequest) GoString() string {
	return s.String()
}

func (s *ScaleUpClusterRequest) SetClusterBizId(v string) *ScaleUpClusterRequest {
	s.ClusterBizId = &v
	return s
}

func (s *ScaleUpClusterRequest) SetCoreNumber(v int64) *ScaleUpClusterRequest {
	s.CoreNumber = &v
	return s
}

func (s *ScaleUpClusterRequest) SetDuration(v int64) *ScaleUpClusterRequest {
	s.Duration = &v
	return s
}

func (s *ScaleUpClusterRequest) SetInstanceId(v string) *ScaleUpClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *ScaleUpClusterRequest) SetInstanceType(v string) *ScaleUpClusterRequest {
	s.InstanceType = &v
	return s
}

func (s *ScaleUpClusterRequest) SetNodeGroupType(v string) *ScaleUpClusterRequest {
	s.NodeGroupType = &v
	return s
}

func (s *ScaleUpClusterRequest) SetPricingCycle(v string) *ScaleUpClusterRequest {
	s.PricingCycle = &v
	return s
}

type ScaleUpClusterResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScaleUpClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleUpClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleUpClusterResponseBody) SetData(v bool) *ScaleUpClusterResponseBody {
	s.Data = &v
	return s
}

func (s *ScaleUpClusterResponseBody) SetErrCode(v string) *ScaleUpClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ScaleUpClusterResponseBody) SetErrMessage(v string) *ScaleUpClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ScaleUpClusterResponseBody) SetHttpStatusCode(v int64) *ScaleUpClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ScaleUpClusterResponseBody) SetRequestId(v string) *ScaleUpClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleUpClusterResponseBody) SetSuccess(v bool) *ScaleUpClusterResponseBody {
	s.Success = &v
	return s
}

type ScaleUpClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleUpClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleUpClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleUpClusterResponse) GoString() string {
	return s.String()
}

func (s *ScaleUpClusterResponse) SetHeaders(v map[string]*string) *ScaleUpClusterResponse {
	s.Headers = v
	return s
}

func (s *ScaleUpClusterResponse) SetStatusCode(v int32) *ScaleUpClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleUpClusterResponse) SetBody(v *ScaleUpClusterResponseBody) *ScaleUpClusterResponse {
	s.Body = v
	return s
}

type SearchClusterInstancesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchClusterInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchClusterInstancesRequest) GoString() string {
	return s.String()
}

func (s *SearchClusterInstancesRequest) SetClusterId(v string) *SearchClusterInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *SearchClusterInstancesRequest) SetClusterName(v string) *SearchClusterInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *SearchClusterInstancesRequest) SetPageNumber(v int32) *SearchClusterInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchClusterInstancesRequest) SetPageSize(v int32) *SearchClusterInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchClusterInstancesRequest) SetRegionId(v string) *SearchClusterInstancesRequest {
	s.RegionId = &v
	return s
}

type SearchClusterInstancesResponseBody struct {
	Data           []*SearchClusterInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrCode        *string                                   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchClusterInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchClusterInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchClusterInstancesResponseBody) SetData(v []*SearchClusterInstancesResponseBodyData) *SearchClusterInstancesResponseBody {
	s.Data = v
	return s
}

func (s *SearchClusterInstancesResponseBody) SetErrCode(v string) *SearchClusterInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SearchClusterInstancesResponseBody) SetErrMessage(v string) *SearchClusterInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SearchClusterInstancesResponseBody) SetHttpStatusCode(v int32) *SearchClusterInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchClusterInstancesResponseBody) SetRequestId(v string) *SearchClusterInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchClusterInstancesResponseBody) SetSuccess(v bool) *SearchClusterInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *SearchClusterInstancesResponseBody) SetTotal(v int32) *SearchClusterInstancesResponseBody {
	s.Total = &v
	return s
}

type SearchClusterInstancesResponseBodyData struct {
	BeginTime           *int64                                                     `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ClusterBizId        *string                                                    `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	ClusterId           *string                                                    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterInstanceInfo *SearchClusterInstancesResponseBodyDataClusterInstanceInfo `json:"ClusterInstanceInfo,omitempty" xml:"ClusterInstanceInfo,omitempty" type:"Struct"`
	ClusterName         *string                                                    `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterStatus       *string                                                    `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	ClusterStatusValue  *int32                                                     `json:"ClusterStatusValue,omitempty" xml:"ClusterStatusValue,omitempty"`
	ClusterType         *string                                                    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ControlCenterUrl    *string                                                    `json:"ControlCenterUrl,omitempty" xml:"ControlCenterUrl,omitempty"`
	Duration            *int32                                                     `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EcsGroupList        []*SearchClusterInstancesResponseBodyDataEcsGroupList      `json:"EcsGroupList,omitempty" xml:"EcsGroupList,omitempty" type:"Repeated"`
	ExpireTime          *int64                                                     `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FailReason          *string                                                    `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	GmtCreate           *int64                                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *int64                                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceConf        map[string]interface{}                                     `json:"InstanceConf,omitempty" xml:"InstanceConf,omitempty"`
	NoticeConfirmed     *bool                                                      `json:"NoticeConfirmed,omitempty" xml:"NoticeConfirmed,omitempty"`
	OrderBizId          *string                                                    `json:"OrderBizId,omitempty" xml:"OrderBizId,omitempty"`
	PackageType         *string                                                    `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PricingCycle        *string                                                    `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	RegionId            *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RunningTime         *int64                                                     `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	ZoneId              *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchClusterInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchClusterInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchClusterInstancesResponseBodyData) SetBeginTime(v int64) *SearchClusterInstancesResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterBizId(v string) *SearchClusterInstancesResponseBodyData {
	s.ClusterBizId = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterId(v string) *SearchClusterInstancesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterInstanceInfo(v *SearchClusterInstancesResponseBodyDataClusterInstanceInfo) *SearchClusterInstancesResponseBodyData {
	s.ClusterInstanceInfo = v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterName(v string) *SearchClusterInstancesResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterStatus(v string) *SearchClusterInstancesResponseBodyData {
	s.ClusterStatus = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterStatusValue(v int32) *SearchClusterInstancesResponseBodyData {
	s.ClusterStatusValue = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetClusterType(v string) *SearchClusterInstancesResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetControlCenterUrl(v string) *SearchClusterInstancesResponseBodyData {
	s.ControlCenterUrl = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetDuration(v int32) *SearchClusterInstancesResponseBodyData {
	s.Duration = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetEcsGroupList(v []*SearchClusterInstancesResponseBodyDataEcsGroupList) *SearchClusterInstancesResponseBodyData {
	s.EcsGroupList = v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetExpireTime(v int64) *SearchClusterInstancesResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetFailReason(v string) *SearchClusterInstancesResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetGmtCreate(v int64) *SearchClusterInstancesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetGmtModified(v int64) *SearchClusterInstancesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetInstanceConf(v map[string]interface{}) *SearchClusterInstancesResponseBodyData {
	s.InstanceConf = v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetNoticeConfirmed(v bool) *SearchClusterInstancesResponseBodyData {
	s.NoticeConfirmed = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetOrderBizId(v string) *SearchClusterInstancesResponseBodyData {
	s.OrderBizId = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetPackageType(v string) *SearchClusterInstancesResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetPricingCycle(v string) *SearchClusterInstancesResponseBodyData {
	s.PricingCycle = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetRegionId(v string) *SearchClusterInstancesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetRunningTime(v int64) *SearchClusterInstancesResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyData) SetZoneId(v string) *SearchClusterInstancesResponseBodyData {
	s.ZoneId = &v
	return s
}

type SearchClusterInstancesResponseBodyDataClusterInstanceInfo struct {
	ControlCenterLoginName *string `json:"ControlCenterLoginName,omitempty" xml:"ControlCenterLoginName,omitempty"`
	ControlCenterUrl       *string `json:"ControlCenterUrl,omitempty" xml:"ControlCenterUrl,omitempty"`
	SgId                   *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	VpcId                  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId                  *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s SearchClusterInstancesResponseBodyDataClusterInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchClusterInstancesResponseBodyDataClusterInstanceInfo) GoString() string {
	return s.String()
}

func (s *SearchClusterInstancesResponseBodyDataClusterInstanceInfo) SetControlCenterLoginName(v string) *SearchClusterInstancesResponseBodyDataClusterInstanceInfo {
	s.ControlCenterLoginName = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataClusterInstanceInfo) SetControlCenterUrl(v string) *SearchClusterInstancesResponseBodyDataClusterInstanceInfo {
	s.ControlCenterUrl = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataClusterInstanceInfo) SetSgId(v string) *SearchClusterInstancesResponseBodyDataClusterInstanceInfo {
	s.SgId = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataClusterInstanceInfo) SetVpcId(v string) *SearchClusterInstancesResponseBodyDataClusterInstanceInfo {
	s.VpcId = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataClusterInstanceInfo) SetVswId(v string) *SearchClusterInstancesResponseBodyDataClusterInstanceInfo {
	s.VswId = &v
	return s
}

type SearchClusterInstancesResponseBodyDataEcsGroupList struct {
	CpuCount           *int32  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupName      *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	HostGroupType      *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MemorySize         *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	SystemDiskCapacity *string `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
	SystemDiskCount    *int32  `json:"SystemDiskCount,omitempty" xml:"SystemDiskCount,omitempty"`
	SystemDiskType     *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
}

func (s SearchClusterInstancesResponseBodyDataEcsGroupList) String() string {
	return tea.Prettify(s)
}

func (s SearchClusterInstancesResponseBodyDataEcsGroupList) GoString() string {
	return s.String()
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetCpuCount(v int32) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.CpuCount = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetDiskCapacity(v int32) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.DiskCapacity = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetDiskCount(v int32) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.DiskCount = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetDiskType(v string) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.DiskType = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetHostGroupName(v string) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.HostGroupName = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetHostGroupType(v string) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.HostGroupType = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetInstanceType(v string) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.InstanceType = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetMemorySize(v int32) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.MemorySize = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetNodeCount(v int32) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.NodeCount = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetSystemDiskCapacity(v string) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.SystemDiskCapacity = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetSystemDiskCount(v int32) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.SystemDiskCount = &v
	return s
}

func (s *SearchClusterInstancesResponseBodyDataEcsGroupList) SetSystemDiskType(v string) *SearchClusterInstancesResponseBodyDataEcsGroupList {
	s.SystemDiskType = &v
	return s
}

type SearchClusterInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchClusterInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchClusterInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchClusterInstancesResponse) GoString() string {
	return s.String()
}

func (s *SearchClusterInstancesResponse) SetHeaders(v map[string]*string) *SearchClusterInstancesResponse {
	s.Headers = v
	return s
}

func (s *SearchClusterInstancesResponse) SetStatusCode(v int32) *SearchClusterInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchClusterInstancesResponse) SetBody(v *SearchClusterInstancesResponseBody) *SearchClusterInstancesResponse {
	s.Body = v
	return s
}

type SingleOrderRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SingleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleOrderRequest) GoString() string {
	return s.String()
}

func (s *SingleOrderRequest) SetInstanceId(v string) *SingleOrderRequest {
	s.InstanceId = &v
	return s
}

type SingleOrderResponseBody struct {
	Data           *SingleOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode        *string                      `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string                      `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SingleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SingleOrderResponseBody) SetData(v *SingleOrderResponseBodyData) *SingleOrderResponseBody {
	s.Data = v
	return s
}

func (s *SingleOrderResponseBody) SetErrCode(v string) *SingleOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SingleOrderResponseBody) SetErrMessage(v string) *SingleOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SingleOrderResponseBody) SetHttpStatusCode(v int32) *SingleOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SingleOrderResponseBody) SetRequestId(v string) *SingleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleOrderResponseBody) SetSuccess(v bool) *SingleOrderResponseBody {
	s.Success = &v
	return s
}

type SingleOrderResponseBodyData struct {
	ClusterId     *string                                    `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterSize   *int32                                     `json:"ClusterSize,omitempty" xml:"ClusterSize,omitempty"`
	ClusterStatus *int32                                     `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	DeployMode    *string                                    `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	Duration      *int32                                     `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EcsGroupList  []*SingleOrderResponseBodyDataEcsGroupList `json:"EcsGroupList,omitempty" xml:"EcsGroupList,omitempty" type:"Repeated"`
	InstanceId    *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId       *string                                    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PackageType   *string                                    `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PricingCycle  *string                                    `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	StorageSize   *int32                                     `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s SingleOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SingleOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *SingleOrderResponseBodyData) SetClusterId(v string) *SingleOrderResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetClusterSize(v int32) *SingleOrderResponseBodyData {
	s.ClusterSize = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetClusterStatus(v int32) *SingleOrderResponseBodyData {
	s.ClusterStatus = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetDeployMode(v string) *SingleOrderResponseBodyData {
	s.DeployMode = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetDuration(v int32) *SingleOrderResponseBodyData {
	s.Duration = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetEcsGroupList(v []*SingleOrderResponseBodyDataEcsGroupList) *SingleOrderResponseBodyData {
	s.EcsGroupList = v
	return s
}

func (s *SingleOrderResponseBodyData) SetInstanceId(v string) *SingleOrderResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetOrderId(v string) *SingleOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetPackageType(v string) *SingleOrderResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetPricingCycle(v string) *SingleOrderResponseBodyData {
	s.PricingCycle = &v
	return s
}

func (s *SingleOrderResponseBodyData) SetStorageSize(v int32) *SingleOrderResponseBodyData {
	s.StorageSize = &v
	return s
}

type SingleOrderResponseBodyDataEcsGroupList struct {
	CpuCount           *int32  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupName      *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	HostGroupType      *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MemorySize         *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	SystemDiskCapacity *int32  `json:"SystemDiskCapacity,omitempty" xml:"SystemDiskCapacity,omitempty"`
	SystemDiskCount    *int32  `json:"SystemDiskCount,omitempty" xml:"SystemDiskCount,omitempty"`
	SystemDiskType     *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
}

func (s SingleOrderResponseBodyDataEcsGroupList) String() string {
	return tea.Prettify(s)
}

func (s SingleOrderResponseBodyDataEcsGroupList) GoString() string {
	return s.String()
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetCpuCount(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.CpuCount = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetDiskCapacity(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.DiskCapacity = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetDiskCount(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.DiskCount = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetDiskType(v string) *SingleOrderResponseBodyDataEcsGroupList {
	s.DiskType = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetHostGroupName(v string) *SingleOrderResponseBodyDataEcsGroupList {
	s.HostGroupName = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetHostGroupType(v string) *SingleOrderResponseBodyDataEcsGroupList {
	s.HostGroupType = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetInstanceType(v string) *SingleOrderResponseBodyDataEcsGroupList {
	s.InstanceType = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetMemorySize(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.MemorySize = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetNodeCount(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.NodeCount = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetSystemDiskCapacity(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.SystemDiskCapacity = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetSystemDiskCount(v int32) *SingleOrderResponseBodyDataEcsGroupList {
	s.SystemDiskCount = &v
	return s
}

func (s *SingleOrderResponseBodyDataEcsGroupList) SetSystemDiskType(v string) *SingleOrderResponseBodyDataEcsGroupList {
	s.SystemDiskType = &v
	return s
}

type SingleOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SingleOrderResponse) GoString() string {
	return s.String()
}

func (s *SingleOrderResponse) SetHeaders(v map[string]*string) *SingleOrderResponse {
	s.Headers = v
	return s
}

func (s *SingleOrderResponse) SetStatusCode(v int32) *SingleOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleOrderResponse) SetBody(v *SingleOrderResponseBody) *SingleOrderResponse {
	s.Body = v
	return s
}

type UpdateClusterNameRequest struct {
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	ClusterName  *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s UpdateClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterNameRequest) SetClusterBizId(v string) *UpdateClusterNameRequest {
	s.ClusterBizId = &v
	return s
}

func (s *UpdateClusterNameRequest) SetClusterName(v string) *UpdateClusterNameRequest {
	s.ClusterName = &v
	return s
}

type UpdateClusterNameResponseBody struct {
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterNameResponseBody) SetData(v bool) *UpdateClusterNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateClusterNameResponseBody) SetErrCode(v string) *UpdateClusterNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateClusterNameResponseBody) SetErrMessage(v string) *UpdateClusterNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateClusterNameResponseBody) SetHttpStatusCode(v int32) *UpdateClusterNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateClusterNameResponseBody) SetRequestId(v string) *UpdateClusterNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterNameResponseBody) SetSuccess(v bool) *UpdateClusterNameResponseBody {
	s.Success = &v
	return s
}

type UpdateClusterNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterNameResponse) SetHeaders(v map[string]*string) *UpdateClusterNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterNameResponse) SetStatusCode(v int32) *UpdateClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterNameResponse) SetBody(v *UpdateClusterNameResponseBody) *UpdateClusterNameResponse {
	s.Body = v
	return s
}

type UploadLicenseResponseBody struct {
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s UploadLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *UploadLicenseResponseBody) SetData(v string) *UploadLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *UploadLicenseResponseBody) SetErrCode(v string) *UploadLicenseResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UploadLicenseResponseBody) SetErrMessage(v string) *UploadLicenseResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UploadLicenseResponseBody) SetHttpStatusCode(v int32) *UploadLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadLicenseResponseBody) SetRequestId(v string) *UploadLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadLicenseResponseBody) SetSuccess(v bool) *UploadLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *UploadLicenseResponseBody) SetTotal(v int32) *UploadLicenseResponseBody {
	s.Total = &v
	return s
}

type UploadLicenseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadLicenseResponse) GoString() string {
	return s.String()
}

func (s *UploadLicenseResponse) SetHeaders(v map[string]*string) *UploadLicenseResponse {
	s.Headers = v
	return s
}

func (s *UploadLicenseResponse) SetStatusCode(v int32) *UploadLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadLicenseResponse) SetBody(v *UploadLicenseResponseBody) *UploadLicenseResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cdp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelOrderWithOptions(request *CancelOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOrder"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOrder(request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelOrderResponse{}
	_body, _err := client.CancelOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckClusterNameWithOptions(request *CheckClusterNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckClusterName"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/check/cluster_name"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckClusterNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckClusterName(request *CheckClusterNameRequest) (_result *CheckClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckClusterNameResponse{}
	_body, _err := client.CheckClusterNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmNoticeWithOptions(request *ConfirmNoticeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConfirmNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmNotice"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/confirm_notice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmNotice(request *ConfirmNoticeRequest) (_result *ConfirmNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmNoticeResponse{}
	_body, _err := client.ConfirmNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterInfo)) {
		query["ClusterInfo"] = request.ClusterInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterDetailWithOptions(request *GetClusterDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterDetail"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/detail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterDetail(request *GetClusterDetailRequest) (_result *GetClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterDetailResponse{}
	_body, _err := client.GetClusterDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HasDefaultRoleWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *HasDefaultRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("HasDefaultRole"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/user/has_default_role"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &HasDefaultRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HasDefaultRole() (_result *HasDefaultRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HasDefaultRoleResponse{}
	_body, _err := client.HasDefaultRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeClouderaDataPlatformWithOptions(ClientToken *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitializeClouderaDataPlatformResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("InitializeClouderaDataPlatform"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/user/create_default_role"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeClouderaDataPlatformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeClouderaDataPlatform(ClientToken *string) (_result *InitializeClouderaDataPlatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitializeClouderaDataPlatformResponse{}
	_body, _err := client.InitializeClouderaDataPlatformWithOptions(ClientToken, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDefaultComponentsWithOptions(request *ListDefaultComponentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDefaultComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityMode)) {
		query["SecurityMode"] = request.SecurityMode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDefaultComponents"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cdp/defaultComponents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDefaultComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDefaultComponents(request *ListDefaultComponentsRequest) (_result *ListDefaultComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDefaultComponentsResponse{}
	_body, _err := client.ListDefaultComponentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeGroupConstraintsWithOptions(request *ListNodeGroupConstraintsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNodeGroupConstraintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeGroupConstraints"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cdp/nodeGroupConstraints"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeGroupConstraints(request *ListNodeGroupConstraintsRequest) (_result *ListNodeGroupConstraintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodeGroupConstraintsResponse{}
	_body, _err := client.ListNodeGroupConstraintsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(request *ListNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/nodes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOperationsWithOptions(request *ListOperationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentOperationNodeId)) {
		query["ParentOperationNodeId"] = request.ParentOperationNodeId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOperations"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/operations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOperations(request *ListOperationsRequest) (_result *ListOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationsResponse{}
	_body, _err := client.ListOperationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/region/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListZonesWithOptions(request *ListZonesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListZones"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/user/zones"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListZones(request *ListZonesRequest) (_result *ListZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListZonesResponse{}
	_body, _err := client.ListZonesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderWithOptions(request *QueryOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrder"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrder(request *QueryOrderRequest) (_result *QueryOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryOrderResponse{}
	_body, _err := client.QueryOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPriceWithOptions(request *QueryPriceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupSpecs)) {
		query["NodeGroupSpecs"] = request.NodeGroupSpecs
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPrice"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/buy/query_price"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPrice(request *QueryPriceRequest) (_result *QueryPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPriceResponse{}
	_body, _err := client.QueryPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRenewOrderWithOptions(request *QueryRenewOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryRenewOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRenewOrder"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/query_renew_order"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRenewOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRenewOrder(request *QueryRenewOrderRequest) (_result *QueryRenewOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRenewOrderResponse{}
	_body, _err := client.QueryRenewOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRenewPriceWithOptions(request *QueryRenewPriceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryRenewPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		query["Instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRenewPrice"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/buy/query_renew_price"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRenewPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRenewPrice(request *QueryRenewPriceRequest) (_result *QueryRenewPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRenewPriceResponse{}
	_body, _err := client.QueryRenewPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryScaleUpOrderWithOptions(request *QueryScaleUpOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryScaleUpOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryScaleUpOrder"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/query_scale_up_order"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScaleUpOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryScaleUpOrder(request *QueryScaleUpOrderRequest) (_result *QueryScaleUpOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryScaleUpOrderResponse{}
	_body, _err := client.QueryScaleUpOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryScaleUpPriceWithOptions(request *QueryScaleUpPriceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryScaleUpPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreNumber)) {
		query["CoreNumber"] = request.CoreNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupType)) {
		query["NodeGroupType"] = request.NodeGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryScaleUpPrice"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/buy/query_scale_up_price"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryScaleUpPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryScaleUpPrice(request *QueryScaleUpPriceRequest) (_result *QueryScaleUpPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryScaleUpPriceResponse{}
	_body, _err := client.QueryScaleUpPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseClusterWithOptions(request *ReleaseClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseCluster"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseCluster(request *ReleaseClusterRequest) (_result *ReleaseClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseClusterResponse{}
	_body, _err := client.ReleaseClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		query["Instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/renew_instance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleUpClusterWithOptions(request *ScaleUpClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScaleUpClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CoreNumber)) {
		query["CoreNumber"] = request.CoreNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupType)) {
		query["NodeGroupType"] = request.NodeGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ScaleUpCluster"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/scale_up"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ScaleUpClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleUpCluster(request *ScaleUpClusterRequest) (_result *ScaleUpClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleUpClusterResponse{}
	_body, _err := client.ScaleUpClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchClusterInstancesWithOptions(request *SearchClusterInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchClusterInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchClusterInstances"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchClusterInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchClusterInstances(request *SearchClusterInstancesRequest) (_result *SearchClusterInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchClusterInstancesResponse{}
	_body, _err := client.SearchClusterInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SingleOrderWithOptions(request *SingleOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SingleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleOrder"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/order/single"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SingleOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SingleOrder(request *SingleOrderRequest) (_result *SingleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SingleOrderResponse{}
	_body, _err := client.SingleOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterNameWithOptions(request *UpdateClusterNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterBizId)) {
		query["ClusterBizId"] = request.ClusterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClusterName"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/cluster/update_name"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClusterNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClusterName(request *UpdateClusterNameRequest) (_result *UpdateClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateClusterNameResponse{}
	_body, _err := client.UpdateClusterNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadLicenseWithOptions(RegionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadLicenseResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UploadLicense"),
		Version:     tea.String("2021-04-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/user/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadLicense(RegionId *string) (_result *UploadLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadLicenseResponse{}
	_body, _err := client.UploadLicenseWithOptions(RegionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
