// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RunFunctionRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Body    *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s RunFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionRequest) GoString() string {
	return s.String()
}

func (s *RunFunctionRequest) SetSpaceId(v string) *RunFunctionRequest {
	s.SpaceId = &v
	return s
}

func (s *RunFunctionRequest) SetBody(v string) *RunFunctionRequest {
	s.Body = &v
	return s
}

type RunFunctionResponseBody struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result      *string                             `json:"Result,omitempty" xml:"Result,omitempty"`
	RuntimeMeta *RunFunctionResponseBodyRuntimeMeta `json:"RuntimeMeta,omitempty" xml:"RuntimeMeta,omitempty" type:"Struct"`
}

func (s RunFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RunFunctionResponseBody) SetRequestId(v string) *RunFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunFunctionResponseBody) SetResult(v string) *RunFunctionResponseBody {
	s.Result = &v
	return s
}

func (s *RunFunctionResponseBody) SetRuntimeMeta(v *RunFunctionResponseBodyRuntimeMeta) *RunFunctionResponseBody {
	s.RuntimeMeta = v
	return s
}

type RunFunctionResponseBodyRuntimeMeta struct {
	InvocationDuration *int32  `json:"InvocationDuration,omitempty" xml:"InvocationDuration,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BillingDuration    *int32  `json:"BillingDuration,omitempty" xml:"BillingDuration,omitempty"`
	MaxMemoryUsage     *int32  `json:"MaxMemoryUsage,omitempty" xml:"MaxMemoryUsage,omitempty"`
}

func (s RunFunctionResponseBodyRuntimeMeta) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionResponseBodyRuntimeMeta) GoString() string {
	return s.String()
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetInvocationDuration(v int32) *RunFunctionResponseBodyRuntimeMeta {
	s.InvocationDuration = &v
	return s
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetRequestId(v string) *RunFunctionResponseBodyRuntimeMeta {
	s.RequestId = &v
	return s
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetBillingDuration(v int32) *RunFunctionResponseBodyRuntimeMeta {
	s.BillingDuration = &v
	return s
}

func (s *RunFunctionResponseBodyRuntimeMeta) SetMaxMemoryUsage(v int32) *RunFunctionResponseBodyRuntimeMeta {
	s.MaxMemoryUsage = &v
	return s
}

type RunFunctionResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s RunFunctionResponse) GoString() string {
	return s.String()
}

func (s *RunFunctionResponse) SetHeaders(v map[string]*string) *RunFunctionResponse {
	s.Headers = v
	return s
}

func (s *RunFunctionResponse) SetBody(v *RunFunctionResponseBody) *RunFunctionResponse {
	s.Body = v
	return s
}

type ListFunctionRequest struct {
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	FilterBy *string `json:"FilterBy,omitempty" xml:"FilterBy,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionRequest) SetPageNum(v int32) *ListFunctionRequest {
	s.PageNum = &v
	return s
}

func (s *ListFunctionRequest) SetPageSize(v int32) *ListFunctionRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionRequest) SetFilterBy(v string) *ListFunctionRequest {
	s.FilterBy = &v
	return s
}

func (s *ListFunctionRequest) SetSpaceId(v string) *ListFunctionRequest {
	s.SpaceId = &v
	return s
}

type ListFunctionResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Paginator *ListFunctionResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	DataList  []*ListFunctionResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
}

func (s ListFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBody) SetRequestId(v string) *ListFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionResponseBody) SetPaginator(v *ListFunctionResponseBodyPaginator) *ListFunctionResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFunctionResponseBody) SetDataList(v []*ListFunctionResponseBodyDataList) *ListFunctionResponseBody {
	s.DataList = v
	return s
}

type ListFunctionResponseBodyPaginator struct {
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
}

func (s ListFunctionResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBodyPaginator) SetPageNum(v int32) *ListFunctionResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFunctionResponseBodyPaginator) SetPageSize(v int32) *ListFunctionResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFunctionResponseBodyPaginator) SetTotal(v int32) *ListFunctionResponseBodyPaginator {
	s.Total = &v
	return s
}

func (s *ListFunctionResponseBodyPaginator) SetPageCount(v int32) *ListFunctionResponseBodyPaginator {
	s.PageCount = &v
	return s
}

type ListFunctionResponseBodyDataList struct {
	TimingTriggerConfig *string                               `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
	HttpTriggerPath     *string                               `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	CreatedAt           *string                               `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ModifiedAt          *string                               `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Name                *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc                *string                               `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Spec                *ListFunctionResponseBodyDataListSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s ListFunctionResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBodyDataList) SetTimingTriggerConfig(v string) *ListFunctionResponseBodyDataList {
	s.TimingTriggerConfig = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetHttpTriggerPath(v string) *ListFunctionResponseBodyDataList {
	s.HttpTriggerPath = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetCreatedAt(v string) *ListFunctionResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetModifiedAt(v string) *ListFunctionResponseBodyDataList {
	s.ModifiedAt = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetName(v string) *ListFunctionResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetDesc(v string) *ListFunctionResponseBodyDataList {
	s.Desc = &v
	return s
}

func (s *ListFunctionResponseBodyDataList) SetSpec(v *ListFunctionResponseBodyDataListSpec) *ListFunctionResponseBodyDataList {
	s.Spec = v
	return s
}

type ListFunctionResponseBodyDataListSpec struct {
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListFunctionResponseBodyDataListSpec) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponseBodyDataListSpec) GoString() string {
	return s.String()
}

func (s *ListFunctionResponseBodyDataListSpec) SetTimeout(v string) *ListFunctionResponseBodyDataListSpec {
	s.Timeout = &v
	return s
}

func (s *ListFunctionResponseBodyDataListSpec) SetRuntime(v string) *ListFunctionResponseBodyDataListSpec {
	s.Runtime = &v
	return s
}

func (s *ListFunctionResponseBodyDataListSpec) SetInstanceConcurrency(v int32) *ListFunctionResponseBodyDataListSpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *ListFunctionResponseBodyDataListSpec) SetMemory(v string) *ListFunctionResponseBodyDataListSpec {
	s.Memory = &v
	return s
}

type ListFunctionResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionResponse) SetHeaders(v map[string]*string) *ListFunctionResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionResponse) SetBody(v *ListFunctionResponseBody) *ListFunctionResponse {
	s.Body = v
	return s
}

type GetWebHostingCertificateDetailRequest struct {
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	CustomDomain *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s GetWebHostingCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailRequest) SetSpaceId(v string) *GetWebHostingCertificateDetailRequest {
	s.SpaceId = &v
	return s
}

func (s *GetWebHostingCertificateDetailRequest) SetCustomDomain(v string) *GetWebHostingCertificateDetailRequest {
	s.CustomDomain = &v
	return s
}

type GetWebHostingCertificateDetailResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetWebHostingCertificateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetWebHostingCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailResponseBody) SetRequestId(v string) *GetWebHostingCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBody) SetData(v *GetWebHostingCertificateDetailResponseBodyData) *GetWebHostingCertificateDetailResponseBody {
	s.Data = v
	return s
}

type GetWebHostingCertificateDetailResponseBodyData struct {
	CertLife                *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertDomainName          *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertExpiredTime         *int64  `json:"CertExpiredTime,omitempty" xml:"CertExpiredTime,omitempty"`
}

func (s GetWebHostingCertificateDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertLife(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertLife = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertType(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertType = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertDomainName(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertDomainName = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetServerCertificateStatus(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.ServerCertificateStatus = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertName(v string) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertName = &v
	return s
}

func (s *GetWebHostingCertificateDetailResponseBodyData) SetCertExpiredTime(v int64) *GetWebHostingCertificateDetailResponseBodyData {
	s.CertExpiredTime = &v
	return s
}

type GetWebHostingCertificateDetailResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingCertificateDetailResponse) SetHeaders(v map[string]*string) *GetWebHostingCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingCertificateDetailResponse) SetBody(v *GetWebHostingCertificateDetailResponseBody) *GetWebHostingCertificateDetailResponse {
	s.Body = v
	return s
}

type UpdateSpaceRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateSpaceRequest) SetSpaceId(v string) *UpdateSpaceRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSpaceRequest) SetDesc(v string) *UpdateSpaceRequest {
	s.Desc = &v
	return s
}

func (s *UpdateSpaceRequest) SetStatus(v string) *UpdateSpaceRequest {
	s.Status = &v
	return s
}

type UpdateSpaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSpaceResponseBody) SetRequestId(v string) *UpdateSpaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSpaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateSpaceResponse) SetHeaders(v map[string]*string) *UpdateSpaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateSpaceResponse) SetBody(v *UpdateSpaceResponseBody) *UpdateSpaceResponse {
	s.Body = v
	return s
}

type SaveWebHostingCustomDomainConfigRequest struct {
	SpaceId           *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	ForceRedirectType *string `json:"ForceRedirectType,omitempty" xml:"ForceRedirectType,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveWebHostingCustomDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainConfigRequest) SetSpaceId(v string) *SaveWebHostingCustomDomainConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *SaveWebHostingCustomDomainConfigRequest) SetForceRedirectType(v string) *SaveWebHostingCustomDomainConfigRequest {
	s.ForceRedirectType = &v
	return s
}

func (s *SaveWebHostingCustomDomainConfigRequest) SetDomainName(v string) *SaveWebHostingCustomDomainConfigRequest {
	s.DomainName = &v
	return s
}

type SaveWebHostingCustomDomainConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveWebHostingCustomDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainConfigResponseBody) SetRequestId(v string) *SaveWebHostingCustomDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveWebHostingCustomDomainConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWebHostingCustomDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebHostingCustomDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainConfigResponse) SetHeaders(v map[string]*string) *SaveWebHostingCustomDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveWebHostingCustomDomainConfigResponse) SetBody(v *SaveWebHostingCustomDomainConfigResponseBody) *SaveWebHostingCustomDomainConfigResponse {
	s.Body = v
	return s
}

type ListFunctionSpecResponseBody struct {
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MemoryList  []*ListFunctionSpecResponseBodyMemoryList  `json:"MemoryList,omitempty" xml:"MemoryList,omitempty" type:"Repeated"`
	RuntimeList []*ListFunctionSpecResponseBodyRuntimeList `json:"RuntimeList,omitempty" xml:"RuntimeList,omitempty" type:"Repeated"`
}

func (s ListFunctionSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionSpecResponseBody) SetRequestId(v string) *ListFunctionSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionSpecResponseBody) SetMemoryList(v []*ListFunctionSpecResponseBodyMemoryList) *ListFunctionSpecResponseBody {
	s.MemoryList = v
	return s
}

func (s *ListFunctionSpecResponseBody) SetRuntimeList(v []*ListFunctionSpecResponseBodyRuntimeList) *ListFunctionSpecResponseBody {
	s.RuntimeList = v
	return s
}

type ListFunctionSpecResponseBodyMemoryList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFunctionSpecResponseBodyMemoryList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionSpecResponseBodyMemoryList) GoString() string {
	return s.String()
}

func (s *ListFunctionSpecResponseBodyMemoryList) SetName(v string) *ListFunctionSpecResponseBodyMemoryList {
	s.Name = &v
	return s
}

type ListFunctionSpecResponseBodyRuntimeList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFunctionSpecResponseBodyRuntimeList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionSpecResponseBodyRuntimeList) GoString() string {
	return s.String()
}

func (s *ListFunctionSpecResponseBodyRuntimeList) SetName(v string) *ListFunctionSpecResponseBodyRuntimeList {
	s.Name = &v
	return s
}

type ListFunctionSpecResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionSpecResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionSpecResponse) SetHeaders(v map[string]*string) *ListFunctionSpecResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionSpecResponse) SetBody(v *ListFunctionSpecResponseBody) *ListFunctionSpecResponse {
	s.Body = v
	return s
}

type DeleteWechatOpenPlatformConfigRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteWechatOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWechatOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteWechatOpenPlatformConfigRequest) SetSpaceId(v string) *DeleteWechatOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteWechatOpenPlatformConfigRequest) SetAppId(v string) *DeleteWechatOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

type DeleteWechatOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWechatOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWechatOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWechatOpenPlatformConfigResponseBody) SetRequestId(v string) *DeleteWechatOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWechatOpenPlatformConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWechatOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWechatOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWechatOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWechatOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *DeleteWechatOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWechatOpenPlatformConfigResponse) SetBody(v *DeleteWechatOpenPlatformConfigResponseBody) *DeleteWechatOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type CreateSpaceRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc        *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	WorkspaceId *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateSpaceRequest) SetName(v string) *CreateSpaceRequest {
	s.Name = &v
	return s
}

func (s *CreateSpaceRequest) SetDesc(v string) *CreateSpaceRequest {
	s.Desc = &v
	return s
}

func (s *CreateSpaceRequest) SetWorkspaceId(v int64) *CreateSpaceRequest {
	s.WorkspaceId = &v
	return s
}

type CreateSpaceResponseBody struct {
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponseBody) SetSpaceId(v string) *CreateSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *CreateSpaceResponseBody) SetRequestId(v string) *CreateSpaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateSpaceResponse) SetHeaders(v map[string]*string) *CreateSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateSpaceResponse) SetBody(v *CreateSpaceResponseBody) *CreateSpaceResponse {
	s.Body = v
	return s
}

type OpenProductRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OpenProductRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenProductRequest) GoString() string {
	return s.String()
}

func (s *OpenProductRequest) SetData(v string) *OpenProductRequest {
	s.Data = &v
	return s
}

type OpenProductResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Synchro   *string `json:"synchro,omitempty" xml:"synchro,omitempty"`
}

func (s OpenProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenProductResponseBody) GoString() string {
	return s.String()
}

func (s *OpenProductResponseBody) SetRequestId(v string) *OpenProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenProductResponseBody) SetData(v string) *OpenProductResponseBody {
	s.Data = &v
	return s
}

func (s *OpenProductResponseBody) SetSynchro(v string) *OpenProductResponseBody {
	s.Synchro = &v
	return s
}

type OpenProductResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenProductResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenProductResponse) GoString() string {
	return s.String()
}

func (s *OpenProductResponse) SetHeaders(v map[string]*string) *OpenProductResponse {
	s.Headers = v
	return s
}

func (s *OpenProductResponse) SetBody(v *OpenProductResponseBody) *OpenProductResponse {
	s.Body = v
	return s
}

type OpenServiceRequest struct {
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s OpenServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenServiceRequest) SetSpaceId(v string) *OpenServiceRequest {
	s.SpaceId = &v
	return s
}

func (s *OpenServiceRequest) SetServiceName(v string) *OpenServiceRequest {
	s.ServiceName = &v
	return s
}

type OpenServiceResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	Count         *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s OpenServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBody) SetRequestId(v string) *OpenServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenServiceResponseBody) SetServiceStatus(v string) *OpenServiceResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *OpenServiceResponseBody) SetCount(v int32) *OpenServiceResponseBody {
	s.Count = &v
	return s
}

type OpenServiceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenServiceResponse) SetHeaders(v map[string]*string) *OpenServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenServiceResponse) SetBody(v *OpenServiceResponseBody) *OpenServiceResponse {
	s.Body = v
	return s
}

type DeleteSpaceRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpaceRequest) SetSpaceId(v string) *DeleteSpaceRequest {
	s.SpaceId = &v
	return s
}

type DeleteSpaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponseBody) SetRequestId(v string) *DeleteSpaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpaceResponse) SetHeaders(v map[string]*string) *DeleteSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpaceResponse) SetBody(v *DeleteSpaceResponseBody) *DeleteSpaceResponse {
	s.Body = v
	return s
}

type DeleteAntOpenPlatformConfigRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAntOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAntOpenPlatformConfigRequest) SetSpaceId(v string) *DeleteAntOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteAntOpenPlatformConfigRequest) SetAppId(v string) *DeleteAntOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

type DeleteAntOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAntOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAntOpenPlatformConfigResponseBody) SetRequestId(v string) *DeleteAntOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAntOpenPlatformConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAntOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAntOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAntOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *DeleteAntOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntOpenPlatformConfigResponse) SetBody(v *DeleteAntOpenPlatformConfigResponseBody) *DeleteAntOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type DescribeFCOpenStatusResponseBody struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsAuthorized *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
}

func (s DescribeFCOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFCOpenStatusResponseBody) SetStatus(v string) *DescribeFCOpenStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFCOpenStatusResponseBody) SetRequestId(v string) *DescribeFCOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFCOpenStatusResponseBody) SetIsAuthorized(v bool) *DescribeFCOpenStatusResponseBody {
	s.IsAuthorized = &v
	return s
}

type DescribeFCOpenStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFCOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFCOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeFCOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeFCOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeFCOpenStatusResponse) SetBody(v *DescribeFCOpenStatusResponseBody) *DescribeFCOpenStatusResponse {
	s.Body = v
	return s
}

type DescribeFileUploadSignedUrlRequest struct {
	Filename    *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s DescribeFileUploadSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileUploadSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignedUrlRequest) SetFilename(v string) *DescribeFileUploadSignedUrlRequest {
	s.Filename = &v
	return s
}

func (s *DescribeFileUploadSignedUrlRequest) SetSize(v int64) *DescribeFileUploadSignedUrlRequest {
	s.Size = &v
	return s
}

func (s *DescribeFileUploadSignedUrlRequest) SetSpaceId(v string) *DescribeFileUploadSignedUrlRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeFileUploadSignedUrlRequest) SetContentType(v string) *DescribeFileUploadSignedUrlRequest {
	s.ContentType = &v
	return s
}

type DescribeFileUploadSignedUrlResponseBody struct {
	SignUrl   *string `json:"SignUrl,omitempty" xml:"SignUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeFileUploadSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileUploadSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignedUrlResponseBody) SetSignUrl(v string) *DescribeFileUploadSignedUrlResponseBody {
	s.SignUrl = &v
	return s
}

func (s *DescribeFileUploadSignedUrlResponseBody) SetRequestId(v string) *DescribeFileUploadSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileUploadSignedUrlResponseBody) SetId(v string) *DescribeFileUploadSignedUrlResponseBody {
	s.Id = &v
	return s
}

type DescribeFileUploadSignedUrlResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFileUploadSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileUploadSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileUploadSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignedUrlResponse) SetHeaders(v map[string]*string) *DescribeFileUploadSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileUploadSignedUrlResponse) SetBody(v *DescribeFileUploadSignedUrlResponseBody) *DescribeFileUploadSignedUrlResponse {
	s.Body = v
	return s
}

type DeleteFileRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetId(v string) *DeleteFileRequest {
	s.Id = &v
	return s
}

func (s *DeleteFileRequest) SetSpaceId(v string) *DeleteFileRequest {
	s.SpaceId = &v
	return s
}

type DeleteFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFileResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type QueryDBImportTaskStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryDBImportTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBImportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDBImportTaskStatusRequest) SetSpaceId(v string) *QueryDBImportTaskStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBImportTaskStatusRequest) SetTaskId(v string) *QueryDBImportTaskStatusRequest {
	s.TaskId = &v
	return s
}

type QueryDBImportTaskStatusResponseBody struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FailedCount   *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
	SuccessCount  *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s QueryDBImportTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBImportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBImportTaskStatusResponseBody) SetStatus(v string) *QueryDBImportTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetFailedCount(v string) *QueryDBImportTaskStatusResponseBody {
	s.FailedCount = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetRequestId(v string) *QueryDBImportTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetDetailMessage(v string) *QueryDBImportTaskStatusResponseBody {
	s.DetailMessage = &v
	return s
}

func (s *QueryDBImportTaskStatusResponseBody) SetSuccessCount(v string) *QueryDBImportTaskStatusResponseBody {
	s.SuccessCount = &v
	return s
}

type QueryDBImportTaskStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBImportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBImportTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBImportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDBImportTaskStatusResponse) SetHeaders(v map[string]*string) *QueryDBImportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDBImportTaskStatusResponse) SetBody(v *QueryDBImportTaskStatusResponseBody) *QueryDBImportTaskStatusResponse {
	s.Body = v
	return s
}

type RegisterFileRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s RegisterFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterFileRequest) GoString() string {
	return s.String()
}

func (s *RegisterFileRequest) SetId(v string) *RegisterFileRequest {
	s.Id = &v
	return s
}

func (s *RegisterFileRequest) SetSpaceId(v string) *RegisterFileRequest {
	s.SpaceId = &v
	return s
}

type RegisterFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterFileResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterFileResponseBody) SetRequestId(v string) *RegisterFileResponseBody {
	s.RequestId = &v
	return s
}

type RegisterFileResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterFileResponse) GoString() string {
	return s.String()
}

func (s *RegisterFileResponse) SetHeaders(v map[string]*string) *RegisterFileResponse {
	s.Headers = v
	return s
}

func (s *RegisterFileResponse) SetBody(v *RegisterFileResponseBody) *RegisterFileResponse {
	s.Body = v
	return s
}

type SaveAntOpenPlatformConfigRequest struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PublicKey  *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	SignMode   *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
	AppCert    *string `json:"AppCert,omitempty" xml:"AppCert,omitempty"`
	PublicCert *string `json:"PublicCert,omitempty" xml:"PublicCert,omitempty"`
	RootCert   *string `json:"RootCert,omitempty" xml:"RootCert,omitempty"`
}

func (s SaveAntOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAntOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveAntOpenPlatformConfigRequest) SetSpaceId(v string) *SaveAntOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetAppId(v string) *SaveAntOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetPublicKey(v string) *SaveAntOpenPlatformConfigRequest {
	s.PublicKey = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetPrivateKey(v string) *SaveAntOpenPlatformConfigRequest {
	s.PrivateKey = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetSignMode(v string) *SaveAntOpenPlatformConfigRequest {
	s.SignMode = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetAppCert(v string) *SaveAntOpenPlatformConfigRequest {
	s.AppCert = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetPublicCert(v string) *SaveAntOpenPlatformConfigRequest {
	s.PublicCert = &v
	return s
}

func (s *SaveAntOpenPlatformConfigRequest) SetRootCert(v string) *SaveAntOpenPlatformConfigRequest {
	s.RootCert = &v
	return s
}

type SaveAntOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveAntOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAntOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAntOpenPlatformConfigResponseBody) SetRequestId(v string) *SaveAntOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveAntOpenPlatformConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveAntOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAntOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAntOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveAntOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *SaveAntOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveAntOpenPlatformConfigResponse) SetBody(v *SaveAntOpenPlatformConfigResponseBody) *SaveAntOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type DescribeFunctionRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionRequest) GoString() string {
	return s.String()
}

func (s *DescribeFunctionRequest) SetName(v string) *DescribeFunctionRequest {
	s.Name = &v
	return s
}

func (s *DescribeFunctionRequest) SetSpaceId(v string) *DescribeFunctionRequest {
	s.SpaceId = &v
	return s
}

type DescribeFunctionResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Function   *DescribeFunctionResponseBodyFunction   `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	Deployment *DescribeFunctionResponseBodyDeployment `json:"Deployment,omitempty" xml:"Deployment,omitempty" type:"Struct"`
}

func (s DescribeFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBody) SetRequestId(v string) *DescribeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFunctionResponseBody) SetFunction(v *DescribeFunctionResponseBodyFunction) *DescribeFunctionResponseBody {
	s.Function = v
	return s
}

func (s *DescribeFunctionResponseBody) SetDeployment(v *DescribeFunctionResponseBodyDeployment) *DescribeFunctionResponseBody {
	s.Deployment = v
	return s
}

type DescribeFunctionResponseBodyFunction struct {
	TimingTriggerConfig *string                                   `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
	HttpTriggerPath     *string                                   `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	CreatedAt           *string                                   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name                *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ModifiedAt          *string                                   `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Desc                *string                                   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Spec                *DescribeFunctionResponseBodyFunctionSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s DescribeFunctionResponseBodyFunction) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBodyFunction) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBodyFunction) SetTimingTriggerConfig(v string) *DescribeFunctionResponseBodyFunction {
	s.TimingTriggerConfig = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetHttpTriggerPath(v string) *DescribeFunctionResponseBodyFunction {
	s.HttpTriggerPath = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetCreatedAt(v string) *DescribeFunctionResponseBodyFunction {
	s.CreatedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetName(v string) *DescribeFunctionResponseBodyFunction {
	s.Name = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetModifiedAt(v string) *DescribeFunctionResponseBodyFunction {
	s.ModifiedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetDesc(v string) *DescribeFunctionResponseBodyFunction {
	s.Desc = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunction) SetSpec(v *DescribeFunctionResponseBodyFunctionSpec) *DescribeFunctionResponseBodyFunction {
	s.Spec = v
	return s
}

type DescribeFunctionResponseBodyFunctionSpec struct {
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeFunctionResponseBodyFunctionSpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBodyFunctionSpec) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetTimeout(v string) *DescribeFunctionResponseBodyFunctionSpec {
	s.Timeout = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetRuntime(v string) *DescribeFunctionResponseBodyFunctionSpec {
	s.Runtime = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetInstanceConcurrency(v int32) *DescribeFunctionResponseBodyFunctionSpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *DescribeFunctionResponseBodyFunctionSpec) SetMemory(v string) *DescribeFunctionResponseBodyFunctionSpec {
	s.Memory = &v
	return s
}

type DescribeFunctionResponseBodyDeployment struct {
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DeploymentId      *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	DownloadSignedUrl *string `json:"DownloadSignedUrl,omitempty" xml:"DownloadSignedUrl,omitempty"`
	VersionNo         *string `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
	ModifiedAt        *string `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
}

func (s DescribeFunctionResponseBodyDeployment) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponseBodyDeployment) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponseBodyDeployment) SetCreatedAt(v string) *DescribeFunctionResponseBodyDeployment {
	s.CreatedAt = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetDeploymentId(v string) *DescribeFunctionResponseBodyDeployment {
	s.DeploymentId = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetDownloadSignedUrl(v string) *DescribeFunctionResponseBodyDeployment {
	s.DownloadSignedUrl = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetVersionNo(v string) *DescribeFunctionResponseBodyDeployment {
	s.VersionNo = &v
	return s
}

func (s *DescribeFunctionResponseBodyDeployment) SetModifiedAt(v string) *DescribeFunctionResponseBodyDeployment {
	s.ModifiedAt = &v
	return s
}

type DescribeFunctionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFunctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeFunctionResponse) SetHeaders(v map[string]*string) *DescribeFunctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeFunctionResponse) SetBody(v *DescribeFunctionResponseBody) *DescribeFunctionResponse {
	s.Body = v
	return s
}

type OpenWebHostingServiceRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s OpenWebHostingServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenWebHostingServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenWebHostingServiceRequest) SetSpaceId(v string) *OpenWebHostingServiceRequest {
	s.SpaceId = &v
	return s
}

type OpenWebHostingServiceResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenWebHostingServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenWebHostingServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenWebHostingServiceResponseBody) SetData(v bool) *OpenWebHostingServiceResponseBody {
	s.Data = &v
	return s
}

func (s *OpenWebHostingServiceResponseBody) SetRequestId(v string) *OpenWebHostingServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenWebHostingServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenWebHostingServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenWebHostingServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenWebHostingServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenWebHostingServiceResponse) SetHeaders(v map[string]*string) *OpenWebHostingServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenWebHostingServiceResponse) SetBody(v *OpenWebHostingServiceResponseBody) *OpenWebHostingServiceResponse {
	s.Body = v
	return s
}

type DescribeSmsSignRequest struct {
	SignId  *string `json:"SignId,omitempty" xml:"SignId,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignRequest) SetSignId(v string) *DescribeSmsSignRequest {
	s.SignId = &v
	return s
}

func (s *DescribeSmsSignRequest) SetSpaceId(v string) *DescribeSmsSignRequest {
	s.SpaceId = &v
	return s
}

type DescribeSmsSignResponseBody struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SignName   *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DescribeSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignResponseBody) SetSpaceId(v string) *DescribeSmsSignResponseBody {
	s.SpaceId = &v
	return s
}

func (s *DescribeSmsSignResponseBody) SetUpdateTime(v string) *DescribeSmsSignResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSmsSignResponseBody) SetRequestId(v string) *DescribeSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmsSignResponseBody) SetRemark(v string) *DescribeSmsSignResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeSmsSignResponseBody) SetCreateTime(v string) *DescribeSmsSignResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSmsSignResponseBody) SetSignName(v string) *DescribeSmsSignResponseBody {
	s.SignName = &v
	return s
}

type DescribeSmsSignResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignResponse) SetHeaders(v map[string]*string) *DescribeSmsSignResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmsSignResponse) SetBody(v *DescribeSmsSignResponseBody) *DescribeSmsSignResponse {
	s.Body = v
	return s
}

type ListAvailableCertificatesRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s ListAvailableCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesRequest) SetSpaceId(v string) *ListAvailableCertificatesRequest {
	s.SpaceId = &v
	return s
}

func (s *ListAvailableCertificatesRequest) SetDomain(v string) *ListAvailableCertificatesRequest {
	s.Domain = &v
	return s
}

type ListAvailableCertificatesResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListAvailableCertificatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListAvailableCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesResponseBody) SetRequestId(v string) *ListAvailableCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableCertificatesResponseBody) SetData(v []*ListAvailableCertificatesResponseBodyData) *ListAvailableCertificatesResponseBody {
	s.Data = v
	return s
}

type ListAvailableCertificatesResponseBodyData struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListAvailableCertificatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesResponseBodyData) SetName(v string) *ListAvailableCertificatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAvailableCertificatesResponseBodyData) SetId(v string) *ListAvailableCertificatesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAvailableCertificatesResponseBodyData) SetStatusCode(v string) *ListAvailableCertificatesResponseBodyData {
	s.StatusCode = &v
	return s
}

type ListAvailableCertificatesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableCertificatesResponse) SetHeaders(v map[string]*string) *ListAvailableCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableCertificatesResponse) SetBody(v *ListAvailableCertificatesResponseBody) *ListAvailableCertificatesResponse {
	s.Body = v
	return s
}

type ListOpenPlatformConfigRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigRequest) SetSpaceId(v string) *ListOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *ListOpenPlatformConfigRequest) SetPlatform(v string) *ListOpenPlatformConfigRequest {
	s.Platform = &v
	return s
}

type ListOpenPlatformConfigResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretList []*ListOpenPlatformConfigResponseBodySecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Repeated"`
}

func (s ListOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigResponseBody) SetRequestId(v string) *ListOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBody) SetSecretList(v []*ListOpenPlatformConfigResponseBodySecretList) *ListOpenPlatformConfigResponseBody {
	s.SecretList = v
	return s
}

type ListOpenPlatformConfigResponseBodySecretList struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	AppSecret  *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	PublicCert *string `json:"PublicCert,omitempty" xml:"PublicCert,omitempty"`
	AppCert    *string `json:"AppCert,omitempty" xml:"AppCert,omitempty"`
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RootCert   *string `json:"RootCert,omitempty" xml:"RootCert,omitempty"`
	PublicKey  *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SignMode   *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
}

func (s ListOpenPlatformConfigResponseBodySecretList) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigResponseBodySecretList) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetSpaceId(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.SpaceId = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetAppSecret(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.AppSecret = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPublicCert(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.PublicCert = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetAppCert(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.AppCert = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPrivateKey(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.PrivateKey = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetAppId(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.AppId = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetRootCert(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.RootCert = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPublicKey(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.PublicKey = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetPlatform(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.Platform = &v
	return s
}

func (s *ListOpenPlatformConfigResponseBodySecretList) SetSignMode(v string) *ListOpenPlatformConfigResponseBodySecretList {
	s.SignMode = &v
	return s
}

type ListOpenPlatformConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *ListOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *ListOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *ListOpenPlatformConfigResponse) SetBody(v *ListOpenPlatformConfigResponseBody) *ListOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type ModifyWebHostingConfigRequest struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	IndexPath  *string `json:"IndexPath,omitempty" xml:"IndexPath,omitempty"`
	ErrorPath  *string `json:"ErrorPath,omitempty" xml:"ErrorPath,omitempty"`
	AllowedIps *string `json:"AllowedIps,omitempty" xml:"AllowedIps,omitempty"`
}

func (s ModifyWebHostingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebHostingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebHostingConfigRequest) SetSpaceId(v string) *ModifyWebHostingConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetIndexPath(v string) *ModifyWebHostingConfigRequest {
	s.IndexPath = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetErrorPath(v string) *ModifyWebHostingConfigRequest {
	s.ErrorPath = &v
	return s
}

func (s *ModifyWebHostingConfigRequest) SetAllowedIps(v string) *ModifyWebHostingConfigRequest {
	s.AllowedIps = &v
	return s
}

type ModifyWebHostingConfigResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebHostingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebHostingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebHostingConfigResponseBody) SetData(v bool) *ModifyWebHostingConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyWebHostingConfigResponseBody) SetRequestId(v string) *ModifyWebHostingConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebHostingConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebHostingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebHostingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebHostingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebHostingConfigResponse) SetHeaders(v map[string]*string) *ModifyWebHostingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebHostingConfigResponse) SetBody(v *ModifyWebHostingConfigResponseBody) *ModifyWebHostingConfigResponse {
	s.Body = v
	return s
}

type DeleteSmsSignRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	SignId  *string `json:"SignId,omitempty" xml:"SignId,omitempty"`
}

func (s DeleteSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsSignRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignRequest) SetSpaceId(v string) *DeleteSmsSignRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteSmsSignRequest) SetSignId(v string) *DeleteSmsSignRequest {
	s.SignId = &v
	return s
}

type DeleteSmsSignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignResponseBody) SetRequestId(v string) *DeleteSmsSignResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSmsSignResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsSignResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsSignResponse) SetHeaders(v map[string]*string) *DeleteSmsSignResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsSignResponse) SetBody(v *DeleteSmsSignResponseBody) *DeleteSmsSignResponse {
	s.Body = v
	return s
}

type DescribeSmsOpenStatusResponseBody struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsAuthorized *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
}

func (s DescribeSmsOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmsOpenStatusResponseBody) SetStatus(v string) *DescribeSmsOpenStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSmsOpenStatusResponseBody) SetRequestId(v string) *DescribeSmsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmsOpenStatusResponseBody) SetIsAuthorized(v bool) *DescribeSmsOpenStatusResponseBody {
	s.IsAuthorized = &v
	return s
}

type DescribeSmsOpenStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmsOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeSmsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmsOpenStatusResponse) SetBody(v *DescribeSmsOpenStatusResponseBody) *DescribeSmsOpenStatusResponse {
	s.Body = v
	return s
}

type ListSpaceRequest struct {
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceRequest) GoString() string {
	return s.String()
}

func (s *ListSpaceRequest) SetPageNum(v int32) *ListSpaceRequest {
	s.PageNum = &v
	return s
}

func (s *ListSpaceRequest) SetPageSize(v int32) *ListSpaceRequest {
	s.PageSize = &v
	return s
}

type ListSpaceResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GmtCreate *string                        `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Count     *int32                         `json:"Count,omitempty" xml:"Count,omitempty"`
	Spaces    []*ListSpaceResponseBodySpaces `json:"Spaces,omitempty" xml:"Spaces,omitempty" type:"Repeated"`
}

func (s ListSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpaceResponseBody) SetRequestId(v string) *ListSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpaceResponseBody) SetGmtCreate(v string) *ListSpaceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *ListSpaceResponseBody) SetCount(v int32) *ListSpaceResponseBody {
	s.Count = &v
	return s
}

func (s *ListSpaceResponseBody) SetSpaces(v []*ListSpaceResponseBodySpaces) *ListSpaceResponseBody {
	s.Spaces = v
	return s
}

type ListSpaceResponseBodySpaces struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s ListSpaceResponseBodySpaces) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *ListSpaceResponseBodySpaces) SetStatus(v string) *ListSpaceResponseBodySpaces {
	s.Status = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetGmtCreate(v int64) *ListSpaceResponseBodySpaces {
	s.GmtCreate = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetSpaceId(v string) *ListSpaceResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetName(v string) *ListSpaceResponseBodySpaces {
	s.Name = &v
	return s
}

func (s *ListSpaceResponseBodySpaces) SetDesc(v string) *ListSpaceResponseBodySpaces {
	s.Desc = &v
	return s
}

type ListSpaceResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSpaceResponse) GoString() string {
	return s.String()
}

func (s *ListSpaceResponse) SetHeaders(v map[string]*string) *ListSpaceResponse {
	s.Headers = v
	return s
}

func (s *ListSpaceResponse) SetBody(v *ListSpaceResponseBody) *ListSpaceResponse {
	s.Body = v
	return s
}

type DeleteDBCollectionRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Body    *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s DeleteDBCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBCollectionRequest) SetSpaceId(v string) *DeleteDBCollectionRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteDBCollectionRequest) SetBody(v string) *DeleteDBCollectionRequest {
	s.Body = &v
	return s
}

type DeleteDBCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBCollectionResponseBody) SetRequestId(v string) *DeleteDBCollectionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBCollectionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBCollectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBCollectionResponse) SetHeaders(v map[string]*string) *DeleteDBCollectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBCollectionResponse) SetBody(v *DeleteDBCollectionResponseBody) *DeleteDBCollectionResponse {
	s.Body = v
	return s
}

type CreateFunctionDeploymentRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateFunctionDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionDeploymentRequest) SetName(v string) *CreateFunctionDeploymentRequest {
	s.Name = &v
	return s
}

func (s *CreateFunctionDeploymentRequest) SetSpaceId(v string) *CreateFunctionDeploymentRequest {
	s.SpaceId = &v
	return s
}

type CreateFunctionDeploymentResponseBody struct {
	UploadSignedUrl *string `json:"UploadSignedUrl,omitempty" xml:"UploadSignedUrl,omitempty"`
	DeploymentId    *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFunctionDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionDeploymentResponseBody) SetUploadSignedUrl(v string) *CreateFunctionDeploymentResponseBody {
	s.UploadSignedUrl = &v
	return s
}

func (s *CreateFunctionDeploymentResponseBody) SetDeploymentId(v string) *CreateFunctionDeploymentResponseBody {
	s.DeploymentId = &v
	return s
}

func (s *CreateFunctionDeploymentResponseBody) SetRequestId(v string) *CreateFunctionDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type CreateFunctionDeploymentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionDeploymentResponse) SetHeaders(v map[string]*string) *CreateFunctionDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionDeploymentResponse) SetBody(v *CreateFunctionDeploymentResponseBody) *CreateFunctionDeploymentResponse {
	s.Body = v
	return s
}

type GetWebHostingUploadCredentialRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s GetWebHostingUploadCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialRequest) SetSpaceId(v string) *GetWebHostingUploadCredentialRequest {
	s.SpaceId = &v
	return s
}

func (s *GetWebHostingUploadCredentialRequest) SetFilePath(v string) *GetWebHostingUploadCredentialRequest {
	s.FilePath = &v
	return s
}

type GetWebHostingUploadCredentialResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetWebHostingUploadCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetWebHostingUploadCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialResponseBody) SetRequestId(v string) *GetWebHostingUploadCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBody) SetData(v *GetWebHostingUploadCredentialResponseBodyData) *GetWebHostingUploadCredentialResponseBody {
	s.Data = v
	return s
}

type GetWebHostingUploadCredentialResponseBodyData struct {
	FilePath      *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ExpiredTime   *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Endpoint      *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
}

func (s GetWebHostingUploadCredentialResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetFilePath(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetSignature(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetPolicy(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetSecurityToken(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetExpiredTime(v int64) *GetWebHostingUploadCredentialResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetEndpoint(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetWebHostingUploadCredentialResponseBodyData) SetAccessKeyId(v string) *GetWebHostingUploadCredentialResponseBodyData {
	s.AccessKeyId = &v
	return s
}

type GetWebHostingUploadCredentialResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingUploadCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingUploadCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingUploadCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingUploadCredentialResponse) SetHeaders(v map[string]*string) *GetWebHostingUploadCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingUploadCredentialResponse) SetBody(v *GetWebHostingUploadCredentialResponseBody) *GetWebHostingUploadCredentialResponse {
	s.Body = v
	return s
}

type ListFunctionDeploymentRequest struct {
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentRequest) SetPageNum(v int32) *ListFunctionDeploymentRequest {
	s.PageNum = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetPageSize(v int32) *ListFunctionDeploymentRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetName(v string) *ListFunctionDeploymentRequest {
	s.Name = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetSpaceId(v string) *ListFunctionDeploymentRequest {
	s.SpaceId = &v
	return s
}

func (s *ListFunctionDeploymentRequest) SetStatus(v string) *ListFunctionDeploymentRequest {
	s.Status = &v
	return s
}

type ListFunctionDeploymentResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Paginator *ListFunctionDeploymentResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	DataList  []*ListFunctionDeploymentResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
}

func (s ListFunctionDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBody) SetRequestId(v string) *ListFunctionDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionDeploymentResponseBody) SetPaginator(v *ListFunctionDeploymentResponseBodyPaginator) *ListFunctionDeploymentResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFunctionDeploymentResponseBody) SetDataList(v []*ListFunctionDeploymentResponseBodyDataList) *ListFunctionDeploymentResponseBody {
	s.DataList = v
	return s
}

type ListFunctionDeploymentResponseBodyPaginator struct {
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
}

func (s ListFunctionDeploymentResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetPageNum(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetPageSize(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetTotal(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.Total = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyPaginator) SetPageCount(v int32) *ListFunctionDeploymentResponseBodyPaginator {
	s.PageCount = &v
	return s
}

type ListFunctionDeploymentResponseBodyDataList struct {
	CreatedAt         *string                                           `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DeploymentId      *string                                           `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	DownloadSignedUrl *string                                           `json:"DownloadSignedUrl,omitempty" xml:"DownloadSignedUrl,omitempty"`
	VersionNo         *string                                           `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
	ModifiedAt        *string                                           `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Status            *ListFunctionDeploymentResponseBodyDataListStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s ListFunctionDeploymentResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetCreatedAt(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetDeploymentId(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.DeploymentId = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetDownloadSignedUrl(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.DownloadSignedUrl = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetVersionNo(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.VersionNo = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetModifiedAt(v string) *ListFunctionDeploymentResponseBodyDataList {
	s.ModifiedAt = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataList) SetStatus(v *ListFunctionDeploymentResponseBodyDataListStatus) *ListFunctionDeploymentResponseBodyDataList {
	s.Status = v
	return s
}

type ListFunctionDeploymentResponseBodyDataListStatus struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ListFunctionDeploymentResponseBodyDataListStatus) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponseBodyDataListStatus) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponseBodyDataListStatus) SetStatus(v string) *ListFunctionDeploymentResponseBodyDataListStatus {
	s.Status = &v
	return s
}

func (s *ListFunctionDeploymentResponseBodyDataListStatus) SetLabel(v string) *ListFunctionDeploymentResponseBodyDataListStatus {
	s.Label = &v
	return s
}

type ListFunctionDeploymentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionDeploymentResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionDeploymentResponse) SetHeaders(v map[string]*string) *ListFunctionDeploymentResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionDeploymentResponse) SetBody(v *ListFunctionDeploymentResponseBody) *ListFunctionDeploymentResponse {
	s.Body = v
	return s
}

type AddDingtalkOpenPlatformConfigRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s AddDingtalkOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDingtalkOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *AddDingtalkOpenPlatformConfigRequest) SetAppId(v string) *AddDingtalkOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *AddDingtalkOpenPlatformConfigRequest) SetAppSecret(v string) *AddDingtalkOpenPlatformConfigRequest {
	s.AppSecret = &v
	return s
}

func (s *AddDingtalkOpenPlatformConfigRequest) SetSpaceId(v string) *AddDingtalkOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type AddDingtalkOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDingtalkOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDingtalkOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddDingtalkOpenPlatformConfigResponseBody) SetRequestId(v string) *AddDingtalkOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type AddDingtalkOpenPlatformConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddDingtalkOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDingtalkOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDingtalkOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *AddDingtalkOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *AddDingtalkOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *AddDingtalkOpenPlatformConfigResponse) SetBody(v *AddDingtalkOpenPlatformConfigResponseBody) *AddDingtalkOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type CreateDBRestoreTaskRequest struct {
	SpaceId           *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	BackupId          *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	OriginCollections *string `json:"OriginCollections,omitempty" xml:"OriginCollections,omitempty"`
	NewCollections    *string `json:"NewCollections,omitempty" xml:"NewCollections,omitempty"`
}

func (s CreateDBRestoreTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRestoreTaskRequest) SetSpaceId(v string) *CreateDBRestoreTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateDBRestoreTaskRequest) SetBackupId(v string) *CreateDBRestoreTaskRequest {
	s.BackupId = &v
	return s
}

func (s *CreateDBRestoreTaskRequest) SetOriginCollections(v string) *CreateDBRestoreTaskRequest {
	s.OriginCollections = &v
	return s
}

func (s *CreateDBRestoreTaskRequest) SetNewCollections(v string) *CreateDBRestoreTaskRequest {
	s.NewCollections = &v
	return s
}

type CreateDBRestoreTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBRestoreTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBRestoreTaskResponseBody) SetRequestId(v string) *CreateDBRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBRestoreTaskResponseBody) SetTaskId(v string) *CreateDBRestoreTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDBRestoreTaskResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBRestoreTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBRestoreTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBRestoreTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDBRestoreTaskResponse) SetHeaders(v map[string]*string) *CreateDBRestoreTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDBRestoreTaskResponse) SetBody(v *CreateDBRestoreTaskResponseBody) *CreateDBRestoreTaskResponse {
	s.Body = v
	return s
}

type AttachWebHostingCertificateRequest struct {
	SpaceId           *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Domain            *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CertType          *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName          *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	PrivateKey        *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s AttachWebHostingCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachWebHostingCertificateRequest) GoString() string {
	return s.String()
}

func (s *AttachWebHostingCertificateRequest) SetSpaceId(v string) *AttachWebHostingCertificateRequest {
	s.SpaceId = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetDomain(v string) *AttachWebHostingCertificateRequest {
	s.Domain = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetCertType(v string) *AttachWebHostingCertificateRequest {
	s.CertType = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetCertName(v string) *AttachWebHostingCertificateRequest {
	s.CertName = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetServerCertificate(v string) *AttachWebHostingCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *AttachWebHostingCertificateRequest) SetPrivateKey(v string) *AttachWebHostingCertificateRequest {
	s.PrivateKey = &v
	return s
}

type AttachWebHostingCertificateResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachWebHostingCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachWebHostingCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *AttachWebHostingCertificateResponseBody) SetData(v bool) *AttachWebHostingCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *AttachWebHostingCertificateResponseBody) SetRequestId(v string) *AttachWebHostingCertificateResponseBody {
	s.RequestId = &v
	return s
}

type AttachWebHostingCertificateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachWebHostingCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachWebHostingCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachWebHostingCertificateResponse) GoString() string {
	return s.String()
}

func (s *AttachWebHostingCertificateResponse) SetHeaders(v map[string]*string) *AttachWebHostingCertificateResponse {
	s.Headers = v
	return s
}

func (s *AttachWebHostingCertificateResponse) SetBody(v *AttachWebHostingCertificateResponseBody) *AttachWebHostingCertificateResponse {
	s.Body = v
	return s
}

type ListFileRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileRequest) GoString() string {
	return s.String()
}

func (s *ListFileRequest) SetSpaceId(v string) *ListFileRequest {
	s.SpaceId = &v
	return s
}

func (s *ListFileRequest) SetPageNum(v int32) *ListFileRequest {
	s.PageNum = &v
	return s
}

func (s *ListFileRequest) SetPageSize(v int32) *ListFileRequest {
	s.PageSize = &v
	return s
}

func (s *ListFileRequest) SetKeyword(v string) *ListFileRequest {
	s.Keyword = &v
	return s
}

type ListFileResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Paginator *ListFileResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	DataList  []*ListFileResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
}

func (s ListFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileResponseBody) SetRequestId(v string) *ListFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileResponseBody) SetPaginator(v *ListFileResponseBodyPaginator) *ListFileResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFileResponseBody) SetDataList(v []*ListFileResponseBodyDataList) *ListFileResponseBody {
	s.DataList = v
	return s
}

type ListFileResponseBodyPaginator struct {
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
}

func (s ListFileResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyPaginator) SetPageNum(v int32) *ListFileResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFileResponseBodyPaginator) SetPageSize(v int32) *ListFileResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFileResponseBodyPaginator) SetTotal(v int32) *ListFileResponseBodyPaginator {
	s.Total = &v
	return s
}

func (s *ListFileResponseBodyPaginator) SetPageCount(v int32) *ListFileResponseBodyPaginator {
	s.PageCount = &v
	return s
}

type ListFileResponseBodyDataList struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Size        *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFileResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyDataList) SetType(v string) *ListFileResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetSize(v int32) *ListFileResponseBodyDataList {
	s.Size = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetGmtCreate(v string) *ListFileResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetUrl(v string) *ListFileResponseBodyDataList {
	s.Url = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetGmtModified(v string) *ListFileResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetName(v string) *ListFileResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListFileResponseBodyDataList) SetId(v string) *ListFileResponseBodyDataList {
	s.Id = &v
	return s
}

type ListFileResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponse) GoString() string {
	return s.String()
}

func (s *ListFileResponse) SetHeaders(v map[string]*string) *ListFileResponse {
	s.Headers = v
	return s
}

func (s *ListFileResponse) SetBody(v *ListFileResponseBody) *ListFileResponse {
	s.Body = v
	return s
}

type QueryDBRestoreTaskStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryDBRestoreTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBRestoreTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDBRestoreTaskStatusRequest) SetSpaceId(v string) *QueryDBRestoreTaskStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBRestoreTaskStatusRequest) SetTaskId(v string) *QueryDBRestoreTaskStatusRequest {
	s.TaskId = &v
	return s
}

type QueryDBRestoreTaskStatusResponseBody struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FailedCount   *int64  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
	SuccessCount  *int64  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s QueryDBRestoreTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBRestoreTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetStatus(v string) *QueryDBRestoreTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetFailedCount(v int64) *QueryDBRestoreTaskStatusResponseBody {
	s.FailedCount = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetRequestId(v string) *QueryDBRestoreTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetDetailMessage(v string) *QueryDBRestoreTaskStatusResponseBody {
	s.DetailMessage = &v
	return s
}

func (s *QueryDBRestoreTaskStatusResponseBody) SetSuccessCount(v int64) *QueryDBRestoreTaskStatusResponseBody {
	s.SuccessCount = &v
	return s
}

type QueryDBRestoreTaskStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBRestoreTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBRestoreTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBRestoreTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDBRestoreTaskStatusResponse) SetHeaders(v map[string]*string) *QueryDBRestoreTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDBRestoreTaskStatusResponse) SetBody(v *QueryDBRestoreTaskStatusResponseBody) *QueryDBRestoreTaskStatusResponse {
	s.Body = v
	return s
}

type VerifyWebHostingDomainOwnerRequest struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyWebHostingDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebHostingDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyWebHostingDomainOwnerRequest) SetSpaceId(v string) *VerifyWebHostingDomainOwnerRequest {
	s.SpaceId = &v
	return s
}

func (s *VerifyWebHostingDomainOwnerRequest) SetDomain(v string) *VerifyWebHostingDomainOwnerRequest {
	s.Domain = &v
	return s
}

func (s *VerifyWebHostingDomainOwnerRequest) SetVerifyType(v string) *VerifyWebHostingDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

type VerifyWebHostingDomainOwnerResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyWebHostingDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebHostingDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWebHostingDomainOwnerResponseBody) SetData(v string) *VerifyWebHostingDomainOwnerResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyWebHostingDomainOwnerResponseBody) SetRequestId(v string) *VerifyWebHostingDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyWebHostingDomainOwnerResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyWebHostingDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyWebHostingDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebHostingDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyWebHostingDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyWebHostingDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyWebHostingDomainOwnerResponse) SetBody(v *VerifyWebHostingDomainOwnerResponseBody) *VerifyWebHostingDomainOwnerResponse {
	s.Body = v
	return s
}

type DeleteSmsTemplateRequest struct {
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateRequest) SetTemplateCode(v string) *DeleteSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetSpaceId(v string) *DeleteSmsTemplateRequest {
	s.SpaceId = &v
	return s
}

type DeleteSmsTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateResponseBody) SetRequestId(v string) *DeleteSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSmsTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateResponse) SetHeaders(v map[string]*string) *DeleteSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmsTemplateResponse) SetBody(v *DeleteSmsTemplateResponseBody) *DeleteSmsTemplateResponse {
	s.Body = v
	return s
}

type QueryDBExportTaskStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryDBExportTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBExportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDBExportTaskStatusRequest) SetSpaceId(v string) *QueryDBExportTaskStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBExportTaskStatusRequest) SetTaskId(v string) *QueryDBExportTaskStatusRequest {
	s.TaskId = &v
	return s
}

type QueryDBExportTaskStatusResponseBody struct {
	ExportedCount *string `json:"ExportedCount,omitempty" xml:"ExportedCount,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DownloadUrl   *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	DetailMessage *string `json:"DetailMessage,omitempty" xml:"DetailMessage,omitempty"`
}

func (s QueryDBExportTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBExportTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBExportTaskStatusResponseBody) SetExportedCount(v string) *QueryDBExportTaskStatusResponseBody {
	s.ExportedCount = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetStatus(v string) *QueryDBExportTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetRequestId(v string) *QueryDBExportTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetDownloadUrl(v string) *QueryDBExportTaskStatusResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *QueryDBExportTaskStatusResponseBody) SetDetailMessage(v string) *QueryDBExportTaskStatusResponseBody {
	s.DetailMessage = &v
	return s
}

type QueryDBExportTaskStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBExportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBExportTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBExportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDBExportTaskStatusResponse) SetHeaders(v map[string]*string) *QueryDBExportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDBExportTaskStatusResponse) SetBody(v *QueryDBExportTaskStatusResponseBody) *QueryDBExportTaskStatusResponse {
	s.Body = v
	return s
}

type CreateDBImportTaskRequest struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	FileType   *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateDBImportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDBImportTaskRequest) SetSpaceId(v string) *CreateDBImportTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateDBImportTaskRequest) SetCollection(v string) *CreateDBImportTaskRequest {
	s.Collection = &v
	return s
}

func (s *CreateDBImportTaskRequest) SetFileType(v string) *CreateDBImportTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateDBImportTaskRequest) SetMode(v string) *CreateDBImportTaskRequest {
	s.Mode = &v
	return s
}

type CreateDBImportTaskResponseBody struct {
	Host        *string `json:"Host,omitempty" xml:"Host,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileKey     *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Signature   *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBImportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBImportTaskResponseBody) SetHost(v string) *CreateDBImportTaskResponseBody {
	s.Host = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetExpireTime(v string) *CreateDBImportTaskResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetFileKey(v string) *CreateDBImportTaskResponseBody {
	s.FileKey = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetAccessKeyId(v string) *CreateDBImportTaskResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetSignature(v string) *CreateDBImportTaskResponseBody {
	s.Signature = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetRequestId(v string) *CreateDBImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetPolicy(v string) *CreateDBImportTaskResponseBody {
	s.Policy = &v
	return s
}

func (s *CreateDBImportTaskResponseBody) SetTaskId(v string) *CreateDBImportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDBImportTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBImportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBImportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDBImportTaskResponse) SetHeaders(v map[string]*string) *CreateDBImportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDBImportTaskResponse) SetBody(v *CreateDBImportTaskResponseBody) *CreateDBImportTaskResponse {
	s.Body = v
	return s
}

type CheckSmsHasAuthorizedToMPSRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CheckSmsHasAuthorizedToMPSRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsHasAuthorizedToMPSRequest) GoString() string {
	return s.String()
}

func (s *CheckSmsHasAuthorizedToMPSRequest) SetSpaceId(v string) *CheckSmsHasAuthorizedToMPSRequest {
	s.SpaceId = &v
	return s
}

type CheckSmsHasAuthorizedToMPSResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsAuthorized *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
}

func (s CheckSmsHasAuthorizedToMPSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsHasAuthorizedToMPSResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSmsHasAuthorizedToMPSResponseBody) SetRequestId(v string) *CheckSmsHasAuthorizedToMPSResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSmsHasAuthorizedToMPSResponseBody) SetIsAuthorized(v bool) *CheckSmsHasAuthorizedToMPSResponseBody {
	s.IsAuthorized = &v
	return s
}

type CheckSmsHasAuthorizedToMPSResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckSmsHasAuthorizedToMPSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckSmsHasAuthorizedToMPSResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsHasAuthorizedToMPSResponse) GoString() string {
	return s.String()
}

func (s *CheckSmsHasAuthorizedToMPSResponse) SetHeaders(v map[string]*string) *CheckSmsHasAuthorizedToMPSResponse {
	s.Headers = v
	return s
}

func (s *CheckSmsHasAuthorizedToMPSResponse) SetBody(v *CheckSmsHasAuthorizedToMPSResponseBody) *CheckSmsHasAuthorizedToMPSResponse {
	s.Body = v
	return s
}

type DescribeServicePolicyRequest struct {
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
}

func (s DescribeServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeServicePolicyRequest) SetSpaceId(v string) *DescribeServicePolicyRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeServicePolicyRequest) SetServiceName(v string) *DescribeServicePolicyRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeServicePolicyRequest) SetCollectionName(v string) *DescribeServicePolicyRequest {
	s.CollectionName = &v
	return s
}

type DescribeServicePolicyResponseBody struct {
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
}

func (s DescribeServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServicePolicyResponseBody) SetSpaceId(v string) *DescribeServicePolicyResponseBody {
	s.SpaceId = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetRequestId(v string) *DescribeServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetPolicy(v string) *DescribeServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetPolicyName(v string) *DescribeServicePolicyResponseBody {
	s.PolicyName = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetServiceName(v string) *DescribeServicePolicyResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServicePolicyResponseBody) SetCollectionName(v string) *DescribeServicePolicyResponseBody {
	s.CollectionName = &v
	return s
}

type DescribeServicePolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeServicePolicyResponse) SetHeaders(v map[string]*string) *DescribeServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeServicePolicyResponse) SetBody(v *DescribeServicePolicyResponseBody) *DescribeServicePolicyResponse {
	s.Body = v
	return s
}

type ListSmsTemplatesRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListSmsTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSmsTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListSmsTemplatesRequest) SetPageNumber(v int32) *ListSmsTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSmsTemplatesRequest) SetPageSize(v int32) *ListSmsTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmsTemplatesRequest) SetSpaceId(v string) *ListSmsTemplatesRequest {
	s.SpaceId = &v
	return s
}

type ListSmsTemplatesResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber   *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SmsTemplates []*ListSmsTemplatesResponseBodySmsTemplates `json:"SmsTemplates,omitempty" xml:"SmsTemplates,omitempty" type:"Repeated"`
}

func (s ListSmsTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSmsTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmsTemplatesResponseBody) SetRequestId(v string) *ListSmsTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmsTemplatesResponseBody) SetPageNumber(v int32) *ListSmsTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSmsTemplatesResponseBody) SetPageSize(v int32) *ListSmsTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSmsTemplatesResponseBody) SetTotalCount(v int32) *ListSmsTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSmsTemplatesResponseBody) SetSmsTemplates(v []*ListSmsTemplatesResponseBodySmsTemplates) *ListSmsTemplatesResponseBody {
	s.SmsTemplates = v
	return s
}

type ListSmsTemplatesResponseBodySmsTemplates struct {
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TemplateCode    *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateType    *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListSmsTemplatesResponseBodySmsTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListSmsTemplatesResponseBodySmsTemplates) GoString() string {
	return s.String()
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetUpdateTime(v string) *ListSmsTemplatesResponseBodySmsTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetTemplateContent(v string) *ListSmsTemplatesResponseBodySmsTemplates {
	s.TemplateContent = &v
	return s
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetRemark(v string) *ListSmsTemplatesResponseBodySmsTemplates {
	s.Remark = &v
	return s
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetTemplateCode(v string) *ListSmsTemplatesResponseBodySmsTemplates {
	s.TemplateCode = &v
	return s
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetCreateTime(v string) *ListSmsTemplatesResponseBodySmsTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetTemplateType(v int32) *ListSmsTemplatesResponseBodySmsTemplates {
	s.TemplateType = &v
	return s
}

func (s *ListSmsTemplatesResponseBodySmsTemplates) SetTemplateName(v string) *ListSmsTemplatesResponseBodySmsTemplates {
	s.TemplateName = &v
	return s
}

type ListSmsTemplatesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSmsTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSmsTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSmsTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListSmsTemplatesResponse) SetHeaders(v map[string]*string) *ListSmsTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListSmsTemplatesResponse) SetBody(v *ListSmsTemplatesResponseBody) *ListSmsTemplatesResponse {
	s.Body = v
	return s
}

type QueryDBBackupCollectionsRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s QueryDBBackupCollectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupCollectionsRequest) GoString() string {
	return s.String()
}

func (s *QueryDBBackupCollectionsRequest) SetSpaceId(v string) *QueryDBBackupCollectionsRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDBBackupCollectionsRequest) SetBackupId(v string) *QueryDBBackupCollectionsRequest {
	s.BackupId = &v
	return s
}

type QueryDBBackupCollectionsResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Collections []*string `json:"Collections,omitempty" xml:"Collections,omitempty" type:"Repeated"`
}

func (s QueryDBBackupCollectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBBackupCollectionsResponseBody) SetRequestId(v string) *QueryDBBackupCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBBackupCollectionsResponseBody) SetCollections(v []*string) *QueryDBBackupCollectionsResponseBody {
	s.Collections = v
	return s
}

type QueryDBBackupCollectionsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBBackupCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBBackupCollectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupCollectionsResponse) GoString() string {
	return s.String()
}

func (s *QueryDBBackupCollectionsResponse) SetHeaders(v map[string]*string) *QueryDBBackupCollectionsResponse {
	s.Headers = v
	return s
}

func (s *QueryDBBackupCollectionsResponse) SetBody(v *QueryDBBackupCollectionsResponseBody) *QueryDBBackupCollectionsResponse {
	s.Body = v
	return s
}

type QueryServiceStatusRequest struct {
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s QueryServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceStatusRequest) SetSpaceId(v string) *QueryServiceStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryServiceStatusRequest) SetServiceName(v string) *QueryServiceStatusRequest {
	s.ServiceName = &v
	return s
}

type QueryServiceStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	Count         *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s QueryServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceStatusResponseBody) SetRequestId(v string) *QueryServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServiceStatusResponseBody) SetServiceStatus(v string) *QueryServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *QueryServiceStatusResponseBody) SetCount(v int32) *QueryServiceStatusResponseBody {
	s.Count = &v
	return s
}

type QueryServiceStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceStatusResponse) SetHeaders(v map[string]*string) *QueryServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceStatusResponse) SetBody(v *QueryServiceStatusResponseBody) *QueryServiceStatusResponse {
	s.Body = v
	return s
}

type DescribeSpaceClientConfigRequest struct {
	SpaceId     *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Detail      *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	WorkspaceId *int64  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeSpaceClientConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceClientConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpaceClientConfigRequest) SetSpaceId(v string) *DescribeSpaceClientConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeSpaceClientConfigRequest) SetDetail(v string) *DescribeSpaceClientConfigRequest {
	s.Detail = &v
	return s
}

func (s *DescribeSpaceClientConfigRequest) SetWorkspaceId(v int64) *DescribeSpaceClientConfigRequest {
	s.WorkspaceId = &v
	return s
}

type DescribeSpaceClientConfigResponseBody struct {
	ApiKey             *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	SpaceId            *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PrivateKey         *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	Endpoint           *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FileUploadEndpoint *string `json:"FileUploadEndpoint,omitempty" xml:"FileUploadEndpoint,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeSpaceClientConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceClientConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpaceClientConfigResponseBody) SetApiKey(v string) *DescribeSpaceClientConfigResponseBody {
	s.ApiKey = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetSpaceId(v string) *DescribeSpaceClientConfigResponseBody {
	s.SpaceId = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetRequestId(v string) *DescribeSpaceClientConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetPrivateKey(v string) *DescribeSpaceClientConfigResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetEndpoint(v string) *DescribeSpaceClientConfigResponseBody {
	s.Endpoint = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetFileUploadEndpoint(v string) *DescribeSpaceClientConfigResponseBody {
	s.FileUploadEndpoint = &v
	return s
}

func (s *DescribeSpaceClientConfigResponseBody) SetName(v string) *DescribeSpaceClientConfigResponseBody {
	s.Name = &v
	return s
}

type DescribeSpaceClientConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSpaceClientConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSpaceClientConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceClientConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpaceClientConfigResponse) SetHeaders(v map[string]*string) *DescribeSpaceClientConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpaceClientConfigResponse) SetBody(v *DescribeSpaceClientConfigResponseBody) *DescribeSpaceClientConfigResponse {
	s.Body = v
	return s
}

type SaveBuiltinFunctionTemplateRequest struct {
	BizId                             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BuiltinFunctionTemplateCategoryId *string `json:"BuiltinFunctionTemplateCategoryId,omitempty" xml:"BuiltinFunctionTemplateCategoryId,omitempty"`
	BuiltinFunctionTemplateProfile    *string `json:"BuiltinFunctionTemplateProfile,omitempty" xml:"BuiltinFunctionTemplateProfile,omitempty"`
}

func (s SaveBuiltinFunctionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBuiltinFunctionTemplateRequest) GoString() string {
	return s.String()
}

func (s *SaveBuiltinFunctionTemplateRequest) SetBizId(v string) *SaveBuiltinFunctionTemplateRequest {
	s.BizId = &v
	return s
}

func (s *SaveBuiltinFunctionTemplateRequest) SetBuiltinFunctionTemplateCategoryId(v string) *SaveBuiltinFunctionTemplateRequest {
	s.BuiltinFunctionTemplateCategoryId = &v
	return s
}

func (s *SaveBuiltinFunctionTemplateRequest) SetBuiltinFunctionTemplateProfile(v string) *SaveBuiltinFunctionTemplateRequest {
	s.BuiltinFunctionTemplateProfile = &v
	return s
}

type SaveBuiltinFunctionTemplateResponseBody struct {
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BuiltinFunctionTemplateId *string `json:"BuiltinFunctionTemplateId,omitempty" xml:"BuiltinFunctionTemplateId,omitempty"`
}

func (s SaveBuiltinFunctionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveBuiltinFunctionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBuiltinFunctionTemplateResponseBody) SetRequestId(v string) *SaveBuiltinFunctionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBuiltinFunctionTemplateResponseBody) SetBuiltinFunctionTemplateId(v string) *SaveBuiltinFunctionTemplateResponseBody {
	s.BuiltinFunctionTemplateId = &v
	return s
}

type SaveBuiltinFunctionTemplateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveBuiltinFunctionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveBuiltinFunctionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBuiltinFunctionTemplateResponse) GoString() string {
	return s.String()
}

func (s *SaveBuiltinFunctionTemplateResponse) SetHeaders(v map[string]*string) *SaveBuiltinFunctionTemplateResponse {
	s.Headers = v
	return s
}

func (s *SaveBuiltinFunctionTemplateResponse) SetBody(v *SaveBuiltinFunctionTemplateResponseBody) *SaveBuiltinFunctionTemplateResponse {
	s.Body = v
	return s
}

type DescribeISVFileUploadSignedUrlRequest struct {
	Filename   *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeISVFileUploadSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeISVFileUploadSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeISVFileUploadSignedUrlRequest) SetFilename(v string) *DescribeISVFileUploadSignedUrlRequest {
	s.Filename = &v
	return s
}

func (s *DescribeISVFileUploadSignedUrlRequest) SetBucketName(v string) *DescribeISVFileUploadSignedUrlRequest {
	s.BucketName = &v
	return s
}

func (s *DescribeISVFileUploadSignedUrlRequest) SetTenantId(v string) *DescribeISVFileUploadSignedUrlRequest {
	s.TenantId = &v
	return s
}

type DescribeISVFileUploadSignedUrlResponseBody struct {
	SignUrl   *string `json:"SignUrl,omitempty" xml:"SignUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeISVFileUploadSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeISVFileUploadSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeISVFileUploadSignedUrlResponseBody) SetSignUrl(v string) *DescribeISVFileUploadSignedUrlResponseBody {
	s.SignUrl = &v
	return s
}

func (s *DescribeISVFileUploadSignedUrlResponseBody) SetRequestId(v string) *DescribeISVFileUploadSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeISVFileUploadSignedUrlResponseBody) SetId(v string) *DescribeISVFileUploadSignedUrlResponseBody {
	s.Id = &v
	return s
}

type DescribeISVFileUploadSignedUrlResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeISVFileUploadSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeISVFileUploadSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeISVFileUploadSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeISVFileUploadSignedUrlResponse) SetHeaders(v map[string]*string) *DescribeISVFileUploadSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeISVFileUploadSignedUrlResponse) SetBody(v *DescribeISVFileUploadSignedUrlResponseBody) *DescribeISVFileUploadSignedUrlResponse {
	s.Body = v
	return s
}

type CreateBuiltinFunctionTemplateRequest struct {
	BuiltinFunctionTemplateCategoryId *string `json:"BuiltinFunctionTemplateCategoryId,omitempty" xml:"BuiltinFunctionTemplateCategoryId,omitempty"`
}

func (s CreateBuiltinFunctionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBuiltinFunctionTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateBuiltinFunctionTemplateRequest) SetBuiltinFunctionTemplateCategoryId(v string) *CreateBuiltinFunctionTemplateRequest {
	s.BuiltinFunctionTemplateCategoryId = &v
	return s
}

type CreateBuiltinFunctionTemplateResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BizId             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ArtifactUploadUrl *string `json:"ArtifactUploadUrl,omitempty" xml:"ArtifactUploadUrl,omitempty"`
}

func (s CreateBuiltinFunctionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBuiltinFunctionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBuiltinFunctionTemplateResponseBody) SetRequestId(v string) *CreateBuiltinFunctionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBuiltinFunctionTemplateResponseBody) SetBizId(v string) *CreateBuiltinFunctionTemplateResponseBody {
	s.BizId = &v
	return s
}

func (s *CreateBuiltinFunctionTemplateResponseBody) SetArtifactUploadUrl(v string) *CreateBuiltinFunctionTemplateResponseBody {
	s.ArtifactUploadUrl = &v
	return s
}

type CreateBuiltinFunctionTemplateResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBuiltinFunctionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBuiltinFunctionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBuiltinFunctionTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateBuiltinFunctionTemplateResponse) SetHeaders(v map[string]*string) *CreateBuiltinFunctionTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateBuiltinFunctionTemplateResponse) SetBody(v *CreateBuiltinFunctionTemplateResponseBody) *CreateBuiltinFunctionTemplateResponse {
	s.Body = v
	return s
}

type GetWebHostingStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusRequest) SetSpaceId(v string) *GetWebHostingStatusRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingStatusResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetWebHostingStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetWebHostingStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusResponseBody) SetRequestId(v string) *GetWebHostingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebHostingStatusResponseBody) SetData(v *GetWebHostingStatusResponseBodyData) *GetWebHostingStatusResponseBody {
	s.Data = v
	return s
}

type GetWebHostingStatusResponseBodyData struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusResponseBodyData) SetStatus(v string) *GetWebHostingStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWebHostingStatusResponseBodyData) SetSpaceId(v string) *GetWebHostingStatusResponseBodyData {
	s.SpaceId = &v
	return s
}

type GetWebHostingStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingStatusResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingStatusResponse) SetHeaders(v map[string]*string) *GetWebHostingStatusResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingStatusResponse) SetBody(v *GetWebHostingStatusResponseBody) *GetWebHostingStatusResponse {
	s.Body = v
	return s
}

type ListFunctionLogRequest struct {
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	FromDate     *int64  `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	ToDate       *int64  `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionLogRequest) SetPageNum(v int32) *ListFunctionLogRequest {
	s.PageNum = &v
	return s
}

func (s *ListFunctionLogRequest) SetPageSize(v int32) *ListFunctionLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionLogRequest) SetName(v string) *ListFunctionLogRequest {
	s.Name = &v
	return s
}

func (s *ListFunctionLogRequest) SetSpaceId(v string) *ListFunctionLogRequest {
	s.SpaceId = &v
	return s
}

func (s *ListFunctionLogRequest) SetLogRequestId(v string) *ListFunctionLogRequest {
	s.LogRequestId = &v
	return s
}

func (s *ListFunctionLogRequest) SetFromDate(v int64) *ListFunctionLogRequest {
	s.FromDate = &v
	return s
}

func (s *ListFunctionLogRequest) SetToDate(v int64) *ListFunctionLogRequest {
	s.ToDate = &v
	return s
}

func (s *ListFunctionLogRequest) SetStatus(v string) *ListFunctionLogRequest {
	s.Status = &v
	return s
}

type ListFunctionLogResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Paginator *ListFunctionLogResponseBodyPaginator  `json:"Paginator,omitempty" xml:"Paginator,omitempty" type:"Struct"`
	DataList  []*ListFunctionLogResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
}

func (s ListFunctionLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponseBody) SetRequestId(v string) *ListFunctionLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionLogResponseBody) SetPaginator(v *ListFunctionLogResponseBodyPaginator) *ListFunctionLogResponseBody {
	s.Paginator = v
	return s
}

func (s *ListFunctionLogResponseBody) SetDataList(v []*ListFunctionLogResponseBodyDataList) *ListFunctionLogResponseBody {
	s.DataList = v
	return s
}

type ListFunctionLogResponseBodyPaginator struct {
	PageNum   *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
}

func (s ListFunctionLogResponseBodyPaginator) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponseBodyPaginator) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponseBodyPaginator) SetPageNum(v int32) *ListFunctionLogResponseBodyPaginator {
	s.PageNum = &v
	return s
}

func (s *ListFunctionLogResponseBodyPaginator) SetPageSize(v int32) *ListFunctionLogResponseBodyPaginator {
	s.PageSize = &v
	return s
}

func (s *ListFunctionLogResponseBodyPaginator) SetTotal(v int32) *ListFunctionLogResponseBodyPaginator {
	s.Total = &v
	return s
}

func (s *ListFunctionLogResponseBodyPaginator) SetPageCount(v int32) *ListFunctionLogResponseBodyPaginator {
	s.PageCount = &v
	return s
}

type ListFunctionLogResponseBodyDataList struct {
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SpaceId      *string   `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FunctionName *string   `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Timestamps   []*string `json:"Timestamps,omitempty" xml:"Timestamps,omitempty" type:"Repeated"`
	Contents     []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	Levels       []*string `json:"Levels,omitempty" xml:"Levels,omitempty" type:"Repeated"`
}

func (s ListFunctionLogResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponseBodyDataList) SetStatus(v string) *ListFunctionLogResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetSpaceId(v string) *ListFunctionLogResponseBodyDataList {
	s.SpaceId = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetRequestId(v string) *ListFunctionLogResponseBodyDataList {
	s.RequestId = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetFunctionName(v string) *ListFunctionLogResponseBodyDataList {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetTimestamps(v []*string) *ListFunctionLogResponseBodyDataList {
	s.Timestamps = v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetContents(v []*string) *ListFunctionLogResponseBodyDataList {
	s.Contents = v
	return s
}

func (s *ListFunctionLogResponseBodyDataList) SetLevels(v []*string) *ListFunctionLogResponseBodyDataList {
	s.Levels = v
	return s
}

type ListFunctionLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFunctionLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionLogResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionLogResponse) SetHeaders(v map[string]*string) *ListFunctionLogResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionLogResponse) SetBody(v *ListFunctionLogResponseBody) *ListFunctionLogResponse {
	s.Body = v
	return s
}

type ListWebHostingFilesRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Prefix   *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListWebHostingFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesRequest) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesRequest) SetSpaceId(v string) *ListWebHostingFilesRequest {
	s.SpaceId = &v
	return s
}

func (s *ListWebHostingFilesRequest) SetPrefix(v string) *ListWebHostingFilesRequest {
	s.Prefix = &v
	return s
}

func (s *ListWebHostingFilesRequest) SetMarker(v string) *ListWebHostingFilesRequest {
	s.Marker = &v
	return s
}

func (s *ListWebHostingFilesRequest) SetPageSize(v int32) *ListWebHostingFilesRequest {
	s.PageSize = &v
	return s
}

type ListWebHostingFilesResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListWebHostingFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListWebHostingFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponseBody) SetRequestId(v string) *ListWebHostingFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebHostingFilesResponseBody) SetData(v *ListWebHostingFilesResponseBodyData) *ListWebHostingFilesResponseBody {
	s.Data = v
	return s
}

type ListWebHostingFilesResponseBodyData struct {
	NextMarker      *string                                               `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Count           *int32                                                `json:"Count,omitempty" xml:"Count,omitempty"`
	WebHostingFiles []*ListWebHostingFilesResponseBodyDataWebHostingFiles `json:"WebHostingFiles,omitempty" xml:"WebHostingFiles,omitempty" type:"Repeated"`
}

func (s ListWebHostingFilesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponseBodyData) SetNextMarker(v string) *ListWebHostingFilesResponseBodyData {
	s.NextMarker = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyData) SetCount(v int32) *ListWebHostingFilesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyData) SetWebHostingFiles(v []*ListWebHostingFilesResponseBodyDataWebHostingFiles) *ListWebHostingFilesResponseBodyData {
	s.WebHostingFiles = v
	return s
}

type ListWebHostingFilesResponseBodyDataWebHostingFiles struct {
	FilePath         *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	ContentType      *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ETag             *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	LastModifiedTime *int64  `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	SignedUrl        *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s ListWebHostingFilesResponseBodyDataWebHostingFiles) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponseBodyDataWebHostingFiles) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetFilePath(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.FilePath = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetContentType(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.ContentType = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetETag(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.ETag = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetSize(v int64) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.Size = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetLastModifiedTime(v int64) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.LastModifiedTime = &v
	return s
}

func (s *ListWebHostingFilesResponseBodyDataWebHostingFiles) SetSignedUrl(v string) *ListWebHostingFilesResponseBodyDataWebHostingFiles {
	s.SignedUrl = &v
	return s
}

type ListWebHostingFilesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWebHostingFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWebHostingFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingFilesResponse) GoString() string {
	return s.String()
}

func (s *ListWebHostingFilesResponse) SetHeaders(v map[string]*string) *ListWebHostingFilesResponse {
	s.Headers = v
	return s
}

func (s *ListWebHostingFilesResponse) SetBody(v *ListWebHostingFilesResponseBody) *ListWebHostingFilesResponse {
	s.Body = v
	return s
}

type DescribeFileRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileRequest) SetId(v string) *DescribeFileRequest {
	s.Id = &v
	return s
}

func (s *DescribeFileRequest) SetSpaceId(v string) *DescribeFileRequest {
	s.SpaceId = &v
	return s
}

type DescribeFileResponseBody struct {
	Type        *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	GmtModified *string  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId   *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size        *float32 `json:"Size,omitempty" xml:"Size,omitempty"`
	GmtCreate   *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBody) SetType(v string) *DescribeFileResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeFileResponseBody) SetUrl(v string) *DescribeFileResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeFileResponseBody) SetGmtModified(v string) *DescribeFileResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFileResponseBody) SetRequestId(v string) *DescribeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileResponseBody) SetSize(v float32) *DescribeFileResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeFileResponseBody) SetGmtCreate(v string) *DescribeFileResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFileResponseBody) SetName(v string) *DescribeFileResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFileResponseBody) SetId(v string) *DescribeFileResponseBody {
	s.Id = &v
	return s
}

type DescribeFileResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileResponse) SetHeaders(v map[string]*string) *DescribeFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileResponse) SetBody(v *DescribeFileResponseBody) *DescribeFileResponse {
	s.Body = v
	return s
}

type MoveWebHostingFileRequest struct {
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	SourceFilePath *string `json:"SourceFilePath,omitempty" xml:"SourceFilePath,omitempty"`
	TargetFilePath *string `json:"TargetFilePath,omitempty" xml:"TargetFilePath,omitempty"`
}

func (s MoveWebHostingFileRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveWebHostingFileRequest) GoString() string {
	return s.String()
}

func (s *MoveWebHostingFileRequest) SetSpaceId(v string) *MoveWebHostingFileRequest {
	s.SpaceId = &v
	return s
}

func (s *MoveWebHostingFileRequest) SetSourceFilePath(v string) *MoveWebHostingFileRequest {
	s.SourceFilePath = &v
	return s
}

func (s *MoveWebHostingFileRequest) SetTargetFilePath(v string) *MoveWebHostingFileRequest {
	s.TargetFilePath = &v
	return s
}

type MoveWebHostingFileResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveWebHostingFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveWebHostingFileResponseBody) GoString() string {
	return s.String()
}

func (s *MoveWebHostingFileResponseBody) SetData(v bool) *MoveWebHostingFileResponseBody {
	s.Data = &v
	return s
}

func (s *MoveWebHostingFileResponseBody) SetRequestId(v string) *MoveWebHostingFileResponseBody {
	s.RequestId = &v
	return s
}

type MoveWebHostingFileResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveWebHostingFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveWebHostingFileResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveWebHostingFileResponse) GoString() string {
	return s.String()
}

func (s *MoveWebHostingFileResponse) SetHeaders(v map[string]*string) *MoveWebHostingFileResponse {
	s.Headers = v
	return s
}

func (s *MoveWebHostingFileResponse) SetBody(v *MoveWebHostingFileResponseBody) *MoveWebHostingFileResponse {
	s.Body = v
	return s
}

type CreateSmsTemplateRequest struct {
	SpaceId         *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TemplateType    *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateRequest) SetSpaceId(v string) *CreateSmsTemplateRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateType(v int32) *CreateSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateName(v string) *CreateSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateContent(v string) *CreateSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetRemark(v string) *CreateSmsTemplateRequest {
	s.Remark = &v
	return s
}

type CreateSmsTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponseBody) SetRequestId(v string) *CreateSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateSmsTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateResponse) SetHeaders(v map[string]*string) *CreateSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsTemplateResponse) SetBody(v *CreateSmsTemplateResponseBody) *CreateSmsTemplateResponse {
	s.Body = v
	return s
}

type DescribeSmsTemplateStatusRequest struct {
	TemplateCodes *string `json:"TemplateCodes,omitempty" xml:"TemplateCodes,omitempty"`
	SpaceId       *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeSmsTemplateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateStatusRequest) SetTemplateCodes(v string) *DescribeSmsTemplateStatusRequest {
	s.TemplateCodes = &v
	return s
}

func (s *DescribeSmsTemplateStatusRequest) SetSpaceId(v string) *DescribeSmsTemplateStatusRequest {
	s.SpaceId = &v
	return s
}

type DescribeSmsTemplateStatusResponseBody struct {
	RequestId        *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateStatuses []*DescribeSmsTemplateStatusResponseBodyTemplateStatuses `json:"TemplateStatuses,omitempty" xml:"TemplateStatuses,omitempty" type:"Repeated"`
}

func (s DescribeSmsTemplateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateStatusResponseBody) SetRequestId(v string) *DescribeSmsTemplateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmsTemplateStatusResponseBody) SetTemplateStatuses(v []*DescribeSmsTemplateStatusResponseBodyTemplateStatuses) *DescribeSmsTemplateStatusResponseBody {
	s.TemplateStatuses = v
	return s
}

type DescribeSmsTemplateStatusResponseBodyTemplateStatuses struct {
	TemplateCode   *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
}

func (s DescribeSmsTemplateStatusResponseBodyTemplateStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateStatusResponseBodyTemplateStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateStatusResponseBodyTemplateStatuses) SetTemplateCode(v string) *DescribeSmsTemplateStatusResponseBodyTemplateStatuses {
	s.TemplateCode = &v
	return s
}

func (s *DescribeSmsTemplateStatusResponseBodyTemplateStatuses) SetReason(v string) *DescribeSmsTemplateStatusResponseBodyTemplateStatuses {
	s.Reason = &v
	return s
}

func (s *DescribeSmsTemplateStatusResponseBodyTemplateStatuses) SetTemplateStatus(v string) *DescribeSmsTemplateStatusResponseBodyTemplateStatuses {
	s.TemplateStatus = &v
	return s
}

type DescribeSmsTemplateStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmsTemplateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmsTemplateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateStatusResponse) SetHeaders(v map[string]*string) *DescribeSmsTemplateStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmsTemplateStatusResponse) SetBody(v *DescribeSmsTemplateStatusResponseBody) *DescribeSmsTemplateStatusResponse {
	s.Body = v
	return s
}

type BindWebHostingCustomDomainRequest struct {
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	CustomDomain *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s BindWebHostingCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BindWebHostingCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *BindWebHostingCustomDomainRequest) SetSpaceId(v string) *BindWebHostingCustomDomainRequest {
	s.SpaceId = &v
	return s
}

func (s *BindWebHostingCustomDomainRequest) SetCustomDomain(v string) *BindWebHostingCustomDomainRequest {
	s.CustomDomain = &v
	return s
}

type BindWebHostingCustomDomainResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindWebHostingCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindWebHostingCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BindWebHostingCustomDomainResponseBody) SetData(v bool) *BindWebHostingCustomDomainResponseBody {
	s.Data = &v
	return s
}

func (s *BindWebHostingCustomDomainResponseBody) SetRequestId(v string) *BindWebHostingCustomDomainResponseBody {
	s.RequestId = &v
	return s
}

type BindWebHostingCustomDomainResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindWebHostingCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindWebHostingCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BindWebHostingCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *BindWebHostingCustomDomainResponse) SetHeaders(v map[string]*string) *BindWebHostingCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *BindWebHostingCustomDomainResponse) SetBody(v *BindWebHostingCustomDomainResponseBody) *BindWebHostingCustomDomainResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Runtime *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetName(v string) *CreateFunctionRequest {
	s.Name = &v
	return s
}

func (s *CreateFunctionRequest) SetDesc(v string) *CreateFunctionRequest {
	s.Desc = &v
	return s
}

func (s *CreateFunctionRequest) SetSpaceId(v string) *CreateFunctionRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateFunctionRequest) SetRuntime(v string) *CreateFunctionRequest {
	s.Runtime = &v
	return s
}

type CreateFunctionResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreatedAt  *string                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name       *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ModifiedAt *string                         `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Desc       *string                         `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Spec       *CreateFunctionResponseBodySpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionResponseBody) SetCreatedAt(v string) *CreateFunctionResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateFunctionResponseBody) SetName(v string) *CreateFunctionResponseBody {
	s.Name = &v
	return s
}

func (s *CreateFunctionResponseBody) SetModifiedAt(v string) *CreateFunctionResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *CreateFunctionResponseBody) SetDesc(v string) *CreateFunctionResponseBody {
	s.Desc = &v
	return s
}

func (s *CreateFunctionResponseBody) SetSpec(v *CreateFunctionResponseBodySpec) *CreateFunctionResponseBody {
	s.Spec = v
	return s
}

type CreateFunctionResponseBodySpec struct {
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	InstanceConcurrency *string `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateFunctionResponseBodySpec) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBodySpec) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBodySpec) SetTimeout(v string) *CreateFunctionResponseBodySpec {
	s.Timeout = &v
	return s
}

func (s *CreateFunctionResponseBodySpec) SetRuntime(v string) *CreateFunctionResponseBodySpec {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionResponseBodySpec) SetInstanceConcurrency(v string) *CreateFunctionResponseBodySpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateFunctionResponseBodySpec) SetMemory(v string) *CreateFunctionResponseBodySpec {
	s.Memory = &v
	return s
}

type CreateFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type DeleteDingtalkOpenPlatformConfigRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteDingtalkOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDingtalkOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDingtalkOpenPlatformConfigRequest) SetAppId(v string) *DeleteDingtalkOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *DeleteDingtalkOpenPlatformConfigRequest) SetSpaceId(v string) *DeleteDingtalkOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type DeleteDingtalkOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDingtalkOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDingtalkOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDingtalkOpenPlatformConfigResponseBody) SetRequestId(v string) *DeleteDingtalkOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDingtalkOpenPlatformConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDingtalkOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDingtalkOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDingtalkOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDingtalkOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *DeleteDingtalkOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDingtalkOpenPlatformConfigResponse) SetBody(v *DeleteDingtalkOpenPlatformConfigResponseBody) *DeleteDingtalkOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type ListExtensionsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExtensionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsRequest) GoString() string {
	return s.String()
}

func (s *ListExtensionsRequest) SetPageNumber(v int32) *ListExtensionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExtensionsRequest) SetPageSize(v int32) *ListExtensionsRequest {
	s.PageSize = &v
	return s
}

type ListExtensionsResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Extensions []*ListExtensionsResponseBodyExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
}

func (s ListExtensionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBody) SetRequestId(v string) *ListExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExtensionsResponseBody) SetPageNumber(v int32) *ListExtensionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExtensionsResponseBody) SetPageSize(v int32) *ListExtensionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExtensionsResponseBody) SetTotalCount(v int32) *ListExtensionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExtensionsResponseBody) SetExtensions(v []*ListExtensionsResponseBodyExtensions) *ListExtensionsResponseBody {
	s.Extensions = v
	return s
}

type ListExtensionsResponseBodyExtensions struct {
	ExtensionDocumentationLink *string `json:"ExtensionDocumentationLink,omitempty" xml:"ExtensionDocumentationLink,omitempty"`
	ExtensionId                *string `json:"ExtensionId,omitempty" xml:"ExtensionId,omitempty"`
	ExtensionDesc              *string `json:"ExtensionDesc,omitempty" xml:"ExtensionDesc,omitempty"`
	ExtensionName              *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	Enabled                    *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ListExtensionsResponseBodyExtensions) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsResponseBodyExtensions) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionDocumentationLink(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionDocumentationLink = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionId(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionId = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionDesc(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionDesc = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetExtensionName(v string) *ListExtensionsResponseBodyExtensions {
	s.ExtensionName = &v
	return s
}

func (s *ListExtensionsResponseBodyExtensions) SetEnabled(v string) *ListExtensionsResponseBodyExtensions {
	s.Enabled = &v
	return s
}

type ListExtensionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExtensionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExtensionsResponse) GoString() string {
	return s.String()
}

func (s *ListExtensionsResponse) SetHeaders(v map[string]*string) *ListExtensionsResponse {
	s.Headers = v
	return s
}

func (s *ListExtensionsResponse) SetBody(v *ListExtensionsResponseBody) *ListExtensionsResponse {
	s.Body = v
	return s
}

type EnableSmsServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSmsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSmsServiceResponseBody) SetRequestId(v string) *EnableSmsServiceResponseBody {
	s.RequestId = &v
	return s
}

type EnableSmsServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSmsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSmsServiceResponse) GoString() string {
	return s.String()
}

func (s *EnableSmsServiceResponse) SetHeaders(v map[string]*string) *EnableSmsServiceResponse {
	s.Headers = v
	return s
}

func (s *EnableSmsServiceResponse) SetBody(v *EnableSmsServiceResponseBody) *EnableSmsServiceResponse {
	s.Body = v
	return s
}

type ReleaseBuiltinFunctionTemplateRequest struct {
	BuiltinFunctionTemplateId *string `json:"BuiltinFunctionTemplateId,omitempty" xml:"BuiltinFunctionTemplateId,omitempty"`
}

func (s ReleaseBuiltinFunctionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseBuiltinFunctionTemplateRequest) GoString() string {
	return s.String()
}

func (s *ReleaseBuiltinFunctionTemplateRequest) SetBuiltinFunctionTemplateId(v string) *ReleaseBuiltinFunctionTemplateRequest {
	s.BuiltinFunctionTemplateId = &v
	return s
}

type ReleaseBuiltinFunctionTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseBuiltinFunctionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseBuiltinFunctionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseBuiltinFunctionTemplateResponseBody) SetRequestId(v string) *ReleaseBuiltinFunctionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseBuiltinFunctionTemplateResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseBuiltinFunctionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseBuiltinFunctionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseBuiltinFunctionTemplateResponse) GoString() string {
	return s.String()
}

func (s *ReleaseBuiltinFunctionTemplateResponse) SetHeaders(v map[string]*string) *ReleaseBuiltinFunctionTemplateResponse {
	s.Headers = v
	return s
}

func (s *ReleaseBuiltinFunctionTemplateResponse) SetBody(v *ReleaseBuiltinFunctionTemplateResponseBody) *ReleaseBuiltinFunctionTemplateResponse {
	s.Body = v
	return s
}

type CreateSmsSignRequest struct {
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s CreateSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignRequest) SetSignName(v string) *CreateSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *CreateSmsSignRequest) SetRemark(v string) *CreateSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsSignRequest) SetSpaceId(v string) *CreateSmsSignRequest {
	s.SpaceId = &v
	return s
}

type CreateSmsSignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBody) SetRequestId(v string) *CreateSmsSignResponseBody {
	s.RequestId = &v
	return s
}

type CreateSmsSignResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponse) SetHeaders(v map[string]*string) *CreateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsSignResponse) SetBody(v *CreateSmsSignResponseBody) *CreateSmsSignResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	Desc                *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Memory              *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Timeout             *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	HttpTriggerPath     *string `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	TimingTriggerConfig *string `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetDesc(v string) *UpdateFunctionRequest {
	s.Desc = &v
	return s
}

func (s *UpdateFunctionRequest) SetName(v string) *UpdateFunctionRequest {
	s.Name = &v
	return s
}

func (s *UpdateFunctionRequest) SetSpaceId(v string) *UpdateFunctionRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateFunctionRequest) SetMemory(v int32) *UpdateFunctionRequest {
	s.Memory = &v
	return s
}

func (s *UpdateFunctionRequest) SetTimeout(v int32) *UpdateFunctionRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateFunctionRequest) SetHttpTriggerPath(v string) *UpdateFunctionRequest {
	s.HttpTriggerPath = &v
	return s
}

func (s *UpdateFunctionRequest) SetTimingTriggerConfig(v string) *UpdateFunctionRequest {
	s.TimingTriggerConfig = &v
	return s
}

func (s *UpdateFunctionRequest) SetInstanceConcurrency(v int32) *UpdateFunctionRequest {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionRequest) SetRuntime(v string) *UpdateFunctionRequest {
	s.Runtime = &v
	return s
}

type UpdateFunctionResponseBody struct {
	RequestId           *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimingTriggerConfig *string                         `json:"TimingTriggerConfig,omitempty" xml:"TimingTriggerConfig,omitempty"`
	HttpTriggerPath     *string                         `json:"HttpTriggerPath,omitempty" xml:"HttpTriggerPath,omitempty"`
	CreatedAt           *string                         `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name                *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ModifiedAt          *string                         `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	Desc                *string                         `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Spec                *UpdateFunctionResponseBodySpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetTimingTriggerConfig(v string) *UpdateFunctionResponseBody {
	s.TimingTriggerConfig = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetHttpTriggerPath(v string) *UpdateFunctionResponseBody {
	s.HttpTriggerPath = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetCreatedAt(v string) *UpdateFunctionResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetName(v string) *UpdateFunctionResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetModifiedAt(v string) *UpdateFunctionResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetDesc(v string) *UpdateFunctionResponseBody {
	s.Desc = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSpec(v *UpdateFunctionResponseBodySpec) *UpdateFunctionResponseBody {
	s.Spec = v
	return s
}

type UpdateFunctionResponseBodySpec struct {
	Timeout             *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Runtime             *string `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	InstanceConcurrency *int32  `json:"InstanceConcurrency,omitempty" xml:"InstanceConcurrency,omitempty"`
	Memory              *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s UpdateFunctionResponseBodySpec) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBodySpec) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBodySpec) SetTimeout(v string) *UpdateFunctionResponseBodySpec {
	s.Timeout = &v
	return s
}

func (s *UpdateFunctionResponseBodySpec) SetRuntime(v string) *UpdateFunctionResponseBodySpec {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionResponseBodySpec) SetInstanceConcurrency(v int32) *UpdateFunctionResponseBodySpec {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionResponseBodySpec) SetMemory(v string) *UpdateFunctionResponseBodySpec {
	s.Memory = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateHttpTriggerConfigRequest struct {
	EnableService           *bool   `json:"EnableService,omitempty" xml:"EnableService,omitempty"`
	SpaceId                 *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	CustomDomain            *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
	CustomDomainCertificate *string `json:"CustomDomainCertificate,omitempty" xml:"CustomDomainCertificate,omitempty"`
	CustomDomainPrivateKey  *string `json:"CustomDomainPrivateKey,omitempty" xml:"CustomDomainPrivateKey,omitempty"`
}

func (s UpdateHttpTriggerConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpTriggerConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpTriggerConfigRequest) SetEnableService(v bool) *UpdateHttpTriggerConfigRequest {
	s.EnableService = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetSpaceId(v string) *UpdateHttpTriggerConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetCustomDomain(v string) *UpdateHttpTriggerConfigRequest {
	s.CustomDomain = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetCustomDomainCertificate(v string) *UpdateHttpTriggerConfigRequest {
	s.CustomDomainCertificate = &v
	return s
}

func (s *UpdateHttpTriggerConfigRequest) SetCustomDomainPrivateKey(v string) *UpdateHttpTriggerConfigRequest {
	s.CustomDomainPrivateKey = &v
	return s
}

type UpdateHttpTriggerConfigResponseBody struct {
	EnableService               *bool   `json:"EnableService,omitempty" xml:"EnableService,omitempty"`
	CustomDomainCname           *string `json:"CustomDomainCname,omitempty" xml:"CustomDomainCname,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DefaultEndpoint             *string `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty"`
	CustomDomainCertificateInfo *string `json:"CustomDomainCertificateInfo,omitempty" xml:"CustomDomainCertificateInfo,omitempty"`
	CustomDomain                *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s UpdateHttpTriggerConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpTriggerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpTriggerConfigResponseBody) SetEnableService(v bool) *UpdateHttpTriggerConfigResponseBody {
	s.EnableService = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetCustomDomainCname(v string) *UpdateHttpTriggerConfigResponseBody {
	s.CustomDomainCname = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetRequestId(v string) *UpdateHttpTriggerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetDefaultEndpoint(v string) *UpdateHttpTriggerConfigResponseBody {
	s.DefaultEndpoint = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetCustomDomainCertificateInfo(v string) *UpdateHttpTriggerConfigResponseBody {
	s.CustomDomainCertificateInfo = &v
	return s
}

func (s *UpdateHttpTriggerConfigResponseBody) SetCustomDomain(v string) *UpdateHttpTriggerConfigResponseBody {
	s.CustomDomain = &v
	return s
}

type UpdateHttpTriggerConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateHttpTriggerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHttpTriggerConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHttpTriggerConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpTriggerConfigResponse) SetHeaders(v map[string]*string) *UpdateHttpTriggerConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpTriggerConfigResponse) SetBody(v *UpdateHttpTriggerConfigResponseBody) *UpdateHttpTriggerConfigResponse {
	s.Body = v
	return s
}

type ResetServerSecretRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ResetServerSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetServerSecretRequest) GoString() string {
	return s.String()
}

func (s *ResetServerSecretRequest) SetSpaceId(v string) *ResetServerSecretRequest {
	s.SpaceId = &v
	return s
}

type ResetServerSecretResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetServerSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetServerSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ResetServerSecretResponseBody) SetRequestId(v string) *ResetServerSecretResponseBody {
	s.RequestId = &v
	return s
}

type ResetServerSecretResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetServerSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetServerSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetServerSecretResponse) GoString() string {
	return s.String()
}

func (s *ResetServerSecretResponse) SetHeaders(v map[string]*string) *ResetServerSecretResponse {
	s.Headers = v
	return s
}

func (s *ResetServerSecretResponse) SetBody(v *ResetServerSecretResponseBody) *ResetServerSecretResponse {
	s.Body = v
	return s
}

type GetWebHostingDomainVerificationContentRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetWebHostingDomainVerificationContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentRequest) SetSpaceId(v string) *GetWebHostingDomainVerificationContentRequest {
	s.SpaceId = &v
	return s
}

func (s *GetWebHostingDomainVerificationContentRequest) SetDomain(v string) *GetWebHostingDomainVerificationContentRequest {
	s.Domain = &v
	return s
}

type GetWebHostingDomainVerificationContentResponseBody struct {
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetWebHostingDomainVerificationContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetWebHostingDomainVerificationContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentResponseBody) SetRequestId(v string) *GetWebHostingDomainVerificationContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebHostingDomainVerificationContentResponseBody) SetData(v *GetWebHostingDomainVerificationContentResponseBodyData) *GetWebHostingDomainVerificationContentResponseBody {
	s.Data = v
	return s
}

type GetWebHostingDomainVerificationContentResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetWebHostingDomainVerificationContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentResponseBodyData) SetContent(v string) *GetWebHostingDomainVerificationContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetWebHostingDomainVerificationContentResponseBodyData) SetDomain(v string) *GetWebHostingDomainVerificationContentResponseBodyData {
	s.Domain = &v
	return s
}

type GetWebHostingDomainVerificationContentResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingDomainVerificationContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingDomainVerificationContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingDomainVerificationContentResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingDomainVerificationContentResponse) SetHeaders(v map[string]*string) *GetWebHostingDomainVerificationContentResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingDomainVerificationContentResponse) SetBody(v *GetWebHostingDomainVerificationContentResponseBody) *GetWebHostingDomainVerificationContentResponse {
	s.Body = v
	return s
}

type UpdateDingtalkOpenPlatformConfigRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s UpdateDingtalkOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingtalkOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDingtalkOpenPlatformConfigRequest) SetAppId(v string) *UpdateDingtalkOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *UpdateDingtalkOpenPlatformConfigRequest) SetAppSecret(v string) *UpdateDingtalkOpenPlatformConfigRequest {
	s.AppSecret = &v
	return s
}

func (s *UpdateDingtalkOpenPlatformConfigRequest) SetSpaceId(v string) *UpdateDingtalkOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

type UpdateDingtalkOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDingtalkOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingtalkOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDingtalkOpenPlatformConfigResponseBody) SetRequestId(v string) *UpdateDingtalkOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDingtalkOpenPlatformConfigResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDingtalkOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDingtalkOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDingtalkOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateDingtalkOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *UpdateDingtalkOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateDingtalkOpenPlatformConfigResponse) SetBody(v *UpdateDingtalkOpenPlatformConfigResponseBody) *UpdateDingtalkOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type CheckMpServerlessRoleExistsRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CheckMpServerlessRoleExistsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMpServerlessRoleExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckMpServerlessRoleExistsRequest) SetRoleName(v string) *CheckMpServerlessRoleExistsRequest {
	s.RoleName = &v
	return s
}

type CheckMpServerlessRoleExistsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Exists    *bool   `json:"Exists,omitempty" xml:"Exists,omitempty"`
}

func (s CheckMpServerlessRoleExistsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMpServerlessRoleExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMpServerlessRoleExistsResponseBody) SetRequestId(v string) *CheckMpServerlessRoleExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMpServerlessRoleExistsResponseBody) SetExists(v bool) *CheckMpServerlessRoleExistsResponseBody {
	s.Exists = &v
	return s
}

type CheckMpServerlessRoleExistsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckMpServerlessRoleExistsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMpServerlessRoleExistsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMpServerlessRoleExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckMpServerlessRoleExistsResponse) SetHeaders(v map[string]*string) *CheckMpServerlessRoleExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckMpServerlessRoleExistsResponse) SetBody(v *CheckMpServerlessRoleExistsResponseBody) *CheckMpServerlessRoleExistsResponse {
	s.Body = v
	return s
}

type EnableExtensionRequest struct {
	ExtensionId *string `json:"ExtensionId,omitempty" xml:"ExtensionId,omitempty"`
}

func (s EnableExtensionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableExtensionRequest) GoString() string {
	return s.String()
}

func (s *EnableExtensionRequest) SetExtensionId(v string) *EnableExtensionRequest {
	s.ExtensionId = &v
	return s
}

type EnableExtensionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableExtensionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableExtensionResponseBody) SetRequestId(v string) *EnableExtensionResponseBody {
	s.RequestId = &v
	return s
}

type EnableExtensionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableExtensionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableExtensionResponse) GoString() string {
	return s.String()
}

func (s *EnableExtensionResponse) SetHeaders(v map[string]*string) *EnableExtensionResponse {
	s.Headers = v
	return s
}

func (s *EnableExtensionResponse) SetBody(v *EnableExtensionResponseBody) *EnableExtensionResponse {
	s.Body = v
	return s
}

type ListSmsSignsForAccountRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListSmsSignsForAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListSmsSignsForAccountRequest) SetPageNumber(v int32) *ListSmsSignsForAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSmsSignsForAccountRequest) SetPageSize(v int32) *ListSmsSignsForAccountRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmsSignsForAccountRequest) SetSpaceId(v string) *ListSmsSignsForAccountRequest {
	s.SpaceId = &v
	return s
}

type ListSmsSignsForAccountResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SmsSigns   []*ListSmsSignsForAccountResponseBodySmsSigns `json:"SmsSigns,omitempty" xml:"SmsSigns,omitempty" type:"Repeated"`
}

func (s ListSmsSignsForAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmsSignsForAccountResponseBody) SetRequestId(v string) *ListSmsSignsForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmsSignsForAccountResponseBody) SetPageNumber(v int32) *ListSmsSignsForAccountResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSmsSignsForAccountResponseBody) SetPageSize(v int32) *ListSmsSignsForAccountResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSmsSignsForAccountResponseBody) SetTotalCount(v int32) *ListSmsSignsForAccountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSmsSignsForAccountResponseBody) SetSmsSigns(v []*ListSmsSignsForAccountResponseBodySmsSigns) *ListSmsSignsForAccountResponseBody {
	s.SmsSigns = v
	return s
}

type ListSmsSignsForAccountResponseBodySmsSigns struct {
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s ListSmsSignsForAccountResponseBodySmsSigns) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsForAccountResponseBodySmsSigns) GoString() string {
	return s.String()
}

func (s *ListSmsSignsForAccountResponseBodySmsSigns) SetSignName(v string) *ListSmsSignsForAccountResponseBodySmsSigns {
	s.SignName = &v
	return s
}

type ListSmsSignsForAccountResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSmsSignsForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSmsSignsForAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListSmsSignsForAccountResponse) SetHeaders(v map[string]*string) *ListSmsSignsForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListSmsSignsForAccountResponse) SetBody(v *ListSmsSignsForAccountResponseBody) *ListSmsSignsForAccountResponse {
	s.Body = v
	return s
}

type ListCorsDomainsRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListCorsDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsRequest) SetSpaceId(v string) *ListCorsDomainsRequest {
	s.SpaceId = &v
	return s
}

type ListCorsDomainsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Domains   []*ListCorsDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s ListCorsDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsResponseBody) SetRequestId(v string) *ListCorsDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCorsDomainsResponseBody) SetDomains(v []*ListCorsDomainsResponseBodyDomains) *ListCorsDomainsResponseBody {
	s.Domains = v
	return s
}

type ListCorsDomainsResponseBodyDomains struct {
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s ListCorsDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsResponseBodyDomains) SetDomain(v string) *ListCorsDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *ListCorsDomainsResponseBodyDomains) SetDomainId(v string) *ListCorsDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

type ListCorsDomainsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCorsDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCorsDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCorsDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListCorsDomainsResponse) SetHeaders(v map[string]*string) *ListCorsDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListCorsDomainsResponse) SetBody(v *ListCorsDomainsResponseBody) *ListCorsDomainsResponse {
	s.Body = v
	return s
}

type ListDingtalkOpenPlatformConfigsRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListDingtalkOpenPlatformConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsRequest) SetSpaceId(v string) *ListDingtalkOpenPlatformConfigsRequest {
	s.SpaceId = &v
	return s
}

type ListDingtalkOpenPlatformConfigsResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Configs   []*ListDingtalkOpenPlatformConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s ListDingtalkOpenPlatformConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsResponseBody) SetRequestId(v string) *ListDingtalkOpenPlatformConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBody) SetConfigs(v []*ListDingtalkOpenPlatformConfigsResponseBodyConfigs) *ListDingtalkOpenPlatformConfigsResponseBody {
	s.Configs = v
	return s
}

type ListDingtalkOpenPlatformConfigsResponseBodyConfigs struct {
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	AppSecret  *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListDingtalkOpenPlatformConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetUpdateTime(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.UpdateTime = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetAppSecret(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.AppSecret = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetAppId(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.AppId = &v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponseBodyConfigs) SetCreateTime(v string) *ListDingtalkOpenPlatformConfigsResponseBodyConfigs {
	s.CreateTime = &v
	return s
}

type ListDingtalkOpenPlatformConfigsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDingtalkOpenPlatformConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDingtalkOpenPlatformConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDingtalkOpenPlatformConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListDingtalkOpenPlatformConfigsResponse) SetHeaders(v map[string]*string) *ListDingtalkOpenPlatformConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListDingtalkOpenPlatformConfigsResponse) SetBody(v *ListDingtalkOpenPlatformConfigsResponseBody) *ListDingtalkOpenPlatformConfigsResponse {
	s.Body = v
	return s
}

type CreateDBExportTaskRequest struct {
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	FileType   *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	Fields     *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
}

func (s CreateDBExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDBExportTaskRequest) SetSpaceId(v string) *CreateDBExportTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateDBExportTaskRequest) SetCollection(v string) *CreateDBExportTaskRequest {
	s.Collection = &v
	return s
}

func (s *CreateDBExportTaskRequest) SetFileType(v string) *CreateDBExportTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateDBExportTaskRequest) SetFields(v string) *CreateDBExportTaskRequest {
	s.Fields = &v
	return s
}

type CreateDBExportTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBExportTaskResponseBody) SetRequestId(v string) *CreateDBExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBExportTaskResponseBody) SetTaskId(v string) *CreateDBExportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDBExportTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDBExportTaskResponse) SetHeaders(v map[string]*string) *CreateDBExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDBExportTaskResponse) SetBody(v *CreateDBExportTaskResponseBody) *CreateDBExportTaskResponse {
	s.Body = v
	return s
}

type GetWebHostingConfigRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s GetWebHostingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigRequest) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigRequest) SetSpaceId(v string) *GetWebHostingConfigRequest {
	s.SpaceId = &v
	return s
}

type GetWebHostingConfigResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetWebHostingConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetWebHostingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigResponseBody) SetRequestId(v string) *GetWebHostingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebHostingConfigResponseBody) SetData(v *GetWebHostingConfigResponseBodyData) *GetWebHostingConfigResponseBody {
	s.Data = v
	return s
}

type GetWebHostingConfigResponseBodyData struct {
	SpaceId       *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	AllowedIps    *string `json:"AllowedIps,omitempty" xml:"AllowedIps,omitempty"`
	ErrorPath     *string `json:"ErrorPath,omitempty" xml:"ErrorPath,omitempty"`
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	IndexPath     *string `json:"IndexPath,omitempty" xml:"IndexPath,omitempty"`
}

func (s GetWebHostingConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigResponseBodyData) SetSpaceId(v string) *GetWebHostingConfigResponseBodyData {
	s.SpaceId = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetAllowedIps(v string) *GetWebHostingConfigResponseBodyData {
	s.AllowedIps = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetErrorPath(v string) *GetWebHostingConfigResponseBodyData {
	s.ErrorPath = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetDefaultDomain(v string) *GetWebHostingConfigResponseBodyData {
	s.DefaultDomain = &v
	return s
}

func (s *GetWebHostingConfigResponseBodyData) SetIndexPath(v string) *GetWebHostingConfigResponseBodyData {
	s.IndexPath = &v
	return s
}

type GetWebHostingConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWebHostingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWebHostingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebHostingConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWebHostingConfigResponse) SetHeaders(v map[string]*string) *GetWebHostingConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWebHostingConfigResponse) SetBody(v *GetWebHostingConfigResponseBody) *GetWebHostingConfigResponse {
	s.Body = v
	return s
}

type UnbindWebHostingCustomDomainRequest struct {
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	CustomDomain *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s UnbindWebHostingCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindWebHostingCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *UnbindWebHostingCustomDomainRequest) SetSpaceId(v string) *UnbindWebHostingCustomDomainRequest {
	s.SpaceId = &v
	return s
}

func (s *UnbindWebHostingCustomDomainRequest) SetCustomDomain(v string) *UnbindWebHostingCustomDomainRequest {
	s.CustomDomain = &v
	return s
}

type UnbindWebHostingCustomDomainResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindWebHostingCustomDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindWebHostingCustomDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindWebHostingCustomDomainResponseBody) SetData(v bool) *UnbindWebHostingCustomDomainResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindWebHostingCustomDomainResponseBody) SetRequestId(v string) *UnbindWebHostingCustomDomainResponseBody {
	s.RequestId = &v
	return s
}

type UnbindWebHostingCustomDomainResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindWebHostingCustomDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindWebHostingCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindWebHostingCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *UnbindWebHostingCustomDomainResponse) SetHeaders(v map[string]*string) *UnbindWebHostingCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *UnbindWebHostingCustomDomainResponse) SetBody(v *UnbindWebHostingCustomDomainResponseBody) *UnbindWebHostingCustomDomainResponse {
	s.Body = v
	return s
}

type DescribeSmsTemplateRequest struct {
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateRequest) SetTemplateCode(v string) *DescribeSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *DescribeSmsTemplateRequest) SetSpaceId(v string) *DescribeSmsTemplateRequest {
	s.SpaceId = &v
	return s
}

type DescribeSmsTemplateResponseBody struct {
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateType    *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateResponseBody) SetUpdateTime(v string) *DescribeSmsTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSmsTemplateResponseBody) SetTemplateContent(v string) *DescribeSmsTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *DescribeSmsTemplateResponseBody) SetRequestId(v string) *DescribeSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmsTemplateResponseBody) SetCreateTime(v string) *DescribeSmsTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSmsTemplateResponseBody) SetTemplateType(v string) *DescribeSmsTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *DescribeSmsTemplateResponseBody) SetTemplateName(v string) *DescribeSmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

type DescribeSmsTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmsTemplateResponse) SetHeaders(v map[string]*string) *DescribeSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmsTemplateResponse) SetBody(v *DescribeSmsTemplateResponseBody) *DescribeSmsTemplateResponse {
	s.Body = v
	return s
}

type SaveWebHostingCustomDomainCorsConfigRequest struct {
	SpaceId                  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	DomainName               *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableCors               *bool   `json:"EnableCors,omitempty" xml:"EnableCors,omitempty"`
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
}

func (s SaveWebHostingCustomDomainCorsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainCorsConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetSpaceId(v string) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetDomainName(v string) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetEnableCors(v bool) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.EnableCors = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigRequest) SetAccessControlAllowOrigin(v string) *SaveWebHostingCustomDomainCorsConfigRequest {
	s.AccessControlAllowOrigin = &v
	return s
}

type SaveWebHostingCustomDomainCorsConfigResponseBody struct {
	// Id of the request
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s SaveWebHostingCustomDomainCorsConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainCorsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetCode(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetMessage(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetHttpStatusCode(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetRequestId(v string) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetSuccess(v bool) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponseBody) SetData(v bool) *SaveWebHostingCustomDomainCorsConfigResponseBody {
	s.Data = &v
	return s
}

type SaveWebHostingCustomDomainCorsConfigResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWebHostingCustomDomainCorsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebHostingCustomDomainCorsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebHostingCustomDomainCorsConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveWebHostingCustomDomainCorsConfigResponse) SetHeaders(v map[string]*string) *SaveWebHostingCustomDomainCorsConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveWebHostingCustomDomainCorsConfigResponse) SetBody(v *SaveWebHostingCustomDomainCorsConfigResponseBody) *SaveWebHostingCustomDomainCorsConfigResponse {
	s.Body = v
	return s
}

type BatchDeleteWebHostingFilesRequest struct {
	SpaceId   *string   `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
}

func (s BatchDeleteWebHostingFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteWebHostingFilesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteWebHostingFilesRequest) SetSpaceId(v string) *BatchDeleteWebHostingFilesRequest {
	s.SpaceId = &v
	return s
}

func (s *BatchDeleteWebHostingFilesRequest) SetFilePaths(v []*string) *BatchDeleteWebHostingFilesRequest {
	s.FilePaths = v
	return s
}

type BatchDeleteWebHostingFilesResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteWebHostingFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteWebHostingFilesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteWebHostingFilesResponseBody) SetData(v bool) *BatchDeleteWebHostingFilesResponseBody {
	s.Data = &v
	return s
}

func (s *BatchDeleteWebHostingFilesResponseBody) SetRequestId(v string) *BatchDeleteWebHostingFilesResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteWebHostingFilesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeleteWebHostingFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteWebHostingFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteWebHostingFilesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteWebHostingFilesResponse) SetHeaders(v map[string]*string) *BatchDeleteWebHostingFilesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteWebHostingFilesResponse) SetBody(v *BatchDeleteWebHostingFilesResponseBody) *BatchDeleteWebHostingFilesResponse {
	s.Body = v
	return s
}

type DeleteCorsDomainRequest struct {
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteCorsDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorsDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteCorsDomainRequest) SetDomainId(v string) *DeleteCorsDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteCorsDomainRequest) SetSpaceId(v string) *DeleteCorsDomainRequest {
	s.SpaceId = &v
	return s
}

type DeleteCorsDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCorsDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorsDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCorsDomainResponseBody) SetRequestId(v string) *DeleteCorsDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCorsDomainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCorsDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCorsDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCorsDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCorsDomainResponse) SetHeaders(v map[string]*string) *DeleteCorsDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteCorsDomainResponse) SetBody(v *DeleteCorsDomainResponseBody) *DeleteCorsDomainResponse {
	s.Body = v
	return s
}

type DescribeHttpTriggerConfigRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeHttpTriggerConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpTriggerConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpTriggerConfigRequest) SetSpaceId(v string) *DescribeHttpTriggerConfigRequest {
	s.SpaceId = &v
	return s
}

type DescribeHttpTriggerConfigResponseBody struct {
	EnableService               *bool   `json:"EnableService,omitempty" xml:"EnableService,omitempty"`
	CustomDomainCname           *string `json:"CustomDomainCname,omitempty" xml:"CustomDomainCname,omitempty"`
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DefaultEndpoint             *string `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty"`
	CustomDomainCertificateInfo *string `json:"CustomDomainCertificateInfo,omitempty" xml:"CustomDomainCertificateInfo,omitempty"`
	CustomDomain                *string `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty"`
}

func (s DescribeHttpTriggerConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpTriggerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpTriggerConfigResponseBody) SetEnableService(v bool) *DescribeHttpTriggerConfigResponseBody {
	s.EnableService = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetCustomDomainCname(v string) *DescribeHttpTriggerConfigResponseBody {
	s.CustomDomainCname = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetRequestId(v string) *DescribeHttpTriggerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetDefaultEndpoint(v string) *DescribeHttpTriggerConfigResponseBody {
	s.DefaultEndpoint = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetCustomDomainCertificateInfo(v string) *DescribeHttpTriggerConfigResponseBody {
	s.CustomDomainCertificateInfo = &v
	return s
}

func (s *DescribeHttpTriggerConfigResponseBody) SetCustomDomain(v string) *DescribeHttpTriggerConfigResponseBody {
	s.CustomDomain = &v
	return s
}

type DescribeHttpTriggerConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHttpTriggerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHttpTriggerConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpTriggerConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpTriggerConfigResponse) SetHeaders(v map[string]*string) *DescribeHttpTriggerConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpTriggerConfigResponse) SetBody(v *DescribeHttpTriggerConfigResponseBody) *DescribeHttpTriggerConfigResponse {
	s.Body = v
	return s
}

type SaveAppAuthTokenRequest struct {
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	IsvAppId     *string `json:"IsvAppId,omitempty" xml:"IsvAppId,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppAuthToken *string `json:"AppAuthToken,omitempty" xml:"AppAuthToken,omitempty"`
}

func (s SaveAppAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAppAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *SaveAppAuthTokenRequest) SetSpaceId(v string) *SaveAppAuthTokenRequest {
	s.SpaceId = &v
	return s
}

func (s *SaveAppAuthTokenRequest) SetIsvAppId(v string) *SaveAppAuthTokenRequest {
	s.IsvAppId = &v
	return s
}

func (s *SaveAppAuthTokenRequest) SetAppId(v string) *SaveAppAuthTokenRequest {
	s.AppId = &v
	return s
}

func (s *SaveAppAuthTokenRequest) SetAppAuthToken(v string) *SaveAppAuthTokenRequest {
	s.AppAuthToken = &v
	return s
}

type SaveAppAuthTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveAppAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAppAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAppAuthTokenResponseBody) SetRequestId(v string) *SaveAppAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

type SaveAppAuthTokenResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveAppAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAppAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAppAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *SaveAppAuthTokenResponse) SetHeaders(v map[string]*string) *SaveAppAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *SaveAppAuthTokenResponse) SetBody(v *SaveAppAuthTokenResponseBody) *SaveAppAuthTokenResponse {
	s.Body = v
	return s
}

type DescribeSmsSignStatusRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	SignIds *string `json:"SignIds,omitempty" xml:"SignIds,omitempty"`
}

func (s DescribeSmsSignStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignStatusRequest) SetSpaceId(v string) *DescribeSmsSignStatusRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeSmsSignStatusRequest) SetSignIds(v string) *DescribeSmsSignStatusRequest {
	s.SignIds = &v
	return s
}

type DescribeSmsSignStatusResponseBody struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignStatuses []*DescribeSmsSignStatusResponseBodySignStatuses `json:"SignStatuses,omitempty" xml:"SignStatuses,omitempty" type:"Repeated"`
}

func (s DescribeSmsSignStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignStatusResponseBody) SetRequestId(v string) *DescribeSmsSignStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmsSignStatusResponseBody) SetSignStatuses(v []*DescribeSmsSignStatusResponseBodySignStatuses) *DescribeSmsSignStatusResponseBody {
	s.SignStatuses = v
	return s
}

type DescribeSmsSignStatusResponseBodySignStatuses struct {
	SignId     *string `json:"SignId,omitempty" xml:"SignId,omitempty"`
	SignStatus *int32  `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
	Reason     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	SignName   *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s DescribeSmsSignStatusResponseBodySignStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignStatusResponseBodySignStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignStatusResponseBodySignStatuses) SetSignId(v string) *DescribeSmsSignStatusResponseBodySignStatuses {
	s.SignId = &v
	return s
}

func (s *DescribeSmsSignStatusResponseBodySignStatuses) SetSignStatus(v int32) *DescribeSmsSignStatusResponseBodySignStatuses {
	s.SignStatus = &v
	return s
}

func (s *DescribeSmsSignStatusResponseBodySignStatuses) SetReason(v string) *DescribeSmsSignStatusResponseBodySignStatuses {
	s.Reason = &v
	return s
}

func (s *DescribeSmsSignStatusResponseBodySignStatuses) SetSignName(v string) *DescribeSmsSignStatusResponseBodySignStatuses {
	s.SignName = &v
	return s
}

type DescribeSmsSignStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSmsSignStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSmsSignStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSmsSignStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmsSignStatusResponse) SetHeaders(v map[string]*string) *DescribeSmsSignStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmsSignStatusResponse) SetBody(v *DescribeSmsSignStatusResponseBody) *DescribeSmsSignStatusResponse {
	s.Body = v
	return s
}

type SaveWechatOpenPlatformConfigRequest struct {
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
}

func (s SaveWechatOpenPlatformConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWechatOpenPlatformConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveWechatOpenPlatformConfigRequest) SetSpaceId(v string) *SaveWechatOpenPlatformConfigRequest {
	s.SpaceId = &v
	return s
}

func (s *SaveWechatOpenPlatformConfigRequest) SetAppId(v string) *SaveWechatOpenPlatformConfigRequest {
	s.AppId = &v
	return s
}

func (s *SaveWechatOpenPlatformConfigRequest) SetAppSecret(v string) *SaveWechatOpenPlatformConfigRequest {
	s.AppSecret = &v
	return s
}

type SaveWechatOpenPlatformConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveWechatOpenPlatformConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWechatOpenPlatformConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWechatOpenPlatformConfigResponseBody) SetRequestId(v string) *SaveWechatOpenPlatformConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveWechatOpenPlatformConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWechatOpenPlatformConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWechatOpenPlatformConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWechatOpenPlatformConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveWechatOpenPlatformConfigResponse) SetHeaders(v map[string]*string) *SaveWechatOpenPlatformConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveWechatOpenPlatformConfigResponse) SetBody(v *SaveWechatOpenPlatformConfigResponseBody) *SaveWechatOpenPlatformConfigResponse {
	s.Body = v
	return s
}

type DescribeSpaceRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DescribeSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpaceRequest) SetSpaceId(v string) *DescribeSpaceRequest {
	s.SpaceId = &v
	return s
}

type DescribeSpaceResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s DescribeSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpaceResponseBody) SetStatus(v string) *DescribeSpaceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSpaceResponseBody) SetSpaceId(v string) *DescribeSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *DescribeSpaceResponseBody) SetRequestId(v string) *DescribeSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpaceResponseBody) SetGmtCreate(v string) *DescribeSpaceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSpaceResponseBody) SetName(v string) *DescribeSpaceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeSpaceResponseBody) SetDesc(v string) *DescribeSpaceResponseBody {
	s.Desc = &v
	return s
}

type DescribeSpaceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpaceResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpaceResponse) SetHeaders(v map[string]*string) *DescribeSpaceResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpaceResponse) SetBody(v *DescribeSpaceResponseBody) *DescribeSpaceResponse {
	s.Body = v
	return s
}

type RenameDBCollectionRequest struct {
	SpaceId          *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	OriginCollection *string `json:"OriginCollection,omitempty" xml:"OriginCollection,omitempty"`
	NewCollection    *string `json:"NewCollection,omitempty" xml:"NewCollection,omitempty"`
}

func (s RenameDBCollectionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDBCollectionRequest) GoString() string {
	return s.String()
}

func (s *RenameDBCollectionRequest) SetSpaceId(v string) *RenameDBCollectionRequest {
	s.SpaceId = &v
	return s
}

func (s *RenameDBCollectionRequest) SetOriginCollection(v string) *RenameDBCollectionRequest {
	s.OriginCollection = &v
	return s
}

func (s *RenameDBCollectionRequest) SetNewCollection(v string) *RenameDBCollectionRequest {
	s.NewCollection = &v
	return s
}

type RenameDBCollectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameDBCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDBCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDBCollectionResponseBody) SetRequestId(v string) *RenameDBCollectionResponseBody {
	s.RequestId = &v
	return s
}

type RenameDBCollectionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameDBCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameDBCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDBCollectionResponse) GoString() string {
	return s.String()
}

func (s *RenameDBCollectionResponse) SetHeaders(v map[string]*string) *RenameDBCollectionResponse {
	s.Headers = v
	return s
}

func (s *RenameDBCollectionResponse) SetBody(v *RenameDBCollectionResponseBody) *RenameDBCollectionResponse {
	s.Body = v
	return s
}

type ListSmsSignsRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SpaceId    *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListSmsSignsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsRequest) GoString() string {
	return s.String()
}

func (s *ListSmsSignsRequest) SetPageNumber(v int32) *ListSmsSignsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSmsSignsRequest) SetPageSize(v int32) *ListSmsSignsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmsSignsRequest) SetSpaceId(v string) *ListSmsSignsRequest {
	s.SpaceId = &v
	return s
}

type ListSmsSignsResponseBody struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SmsSigns   []*ListSmsSignsResponseBodySmsSigns `json:"SmsSigns,omitempty" xml:"SmsSigns,omitempty" type:"Repeated"`
}

func (s ListSmsSignsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmsSignsResponseBody) SetRequestId(v string) *ListSmsSignsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmsSignsResponseBody) SetPageNumber(v int32) *ListSmsSignsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSmsSignsResponseBody) SetPageSize(v int32) *ListSmsSignsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSmsSignsResponseBody) SetTotalCount(v int32) *ListSmsSignsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSmsSignsResponseBody) SetSmsSigns(v []*ListSmsSignsResponseBodySmsSigns) *ListSmsSignsResponseBody {
	s.SmsSigns = v
	return s
}

type ListSmsSignsResponseBodySmsSigns struct {
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	SignId     *string `json:"SignId,omitempty" xml:"SignId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SignName   *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListSmsSignsResponseBodySmsSigns) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsResponseBodySmsSigns) GoString() string {
	return s.String()
}

func (s *ListSmsSignsResponseBodySmsSigns) SetUpdateTime(v string) *ListSmsSignsResponseBodySmsSigns {
	s.UpdateTime = &v
	return s
}

func (s *ListSmsSignsResponseBodySmsSigns) SetSignId(v string) *ListSmsSignsResponseBodySmsSigns {
	s.SignId = &v
	return s
}

func (s *ListSmsSignsResponseBodySmsSigns) SetRemark(v string) *ListSmsSignsResponseBodySmsSigns {
	s.Remark = &v
	return s
}

func (s *ListSmsSignsResponseBodySmsSigns) SetSignName(v string) *ListSmsSignsResponseBodySmsSigns {
	s.SignName = &v
	return s
}

func (s *ListSmsSignsResponseBodySmsSigns) SetCreateTime(v string) *ListSmsSignsResponseBodySmsSigns {
	s.CreateTime = &v
	return s
}

type ListSmsSignsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSmsSignsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSmsSignsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSmsSignsResponse) GoString() string {
	return s.String()
}

func (s *ListSmsSignsResponse) SetHeaders(v map[string]*string) *ListSmsSignsResponse {
	s.Headers = v
	return s
}

func (s *ListSmsSignsResponse) SetBody(v *ListSmsSignsResponseBody) *ListSmsSignsResponse {
	s.Body = v
	return s
}

type DescribeProductOpenStatusRequest struct {
	Name   *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc   *string                                   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Labels []*DescribeProductOpenStatusRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DescribeProductOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductOpenStatusRequest) SetName(v string) *DescribeProductOpenStatusRequest {
	s.Name = &v
	return s
}

func (s *DescribeProductOpenStatusRequest) SetDesc(v string) *DescribeProductOpenStatusRequest {
	s.Desc = &v
	return s
}

func (s *DescribeProductOpenStatusRequest) SetLabels(v []*DescribeProductOpenStatusRequestLabels) *DescribeProductOpenStatusRequest {
	s.Labels = v
	return s
}

type DescribeProductOpenStatusRequestLabels struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeProductOpenStatusRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductOpenStatusRequestLabels) GoString() string {
	return s.String()
}

func (s *DescribeProductOpenStatusRequestLabels) SetValue(v string) *DescribeProductOpenStatusRequestLabels {
	s.Value = &v
	return s
}

func (s *DescribeProductOpenStatusRequestLabels) SetName(v string) *DescribeProductOpenStatusRequestLabels {
	s.Name = &v
	return s
}

type DescribeProductOpenStatusResponseBody struct {
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProductOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductOpenStatusResponseBody) SetSpaceId(v string) *DescribeProductOpenStatusResponseBody {
	s.SpaceId = &v
	return s
}

func (s *DescribeProductOpenStatusResponseBody) SetRequestId(v string) *DescribeProductOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProductOpenStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProductOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProductOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeProductOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductOpenStatusResponse) SetBody(v *DescribeProductOpenStatusResponseBody) *DescribeProductOpenStatusResponse {
	s.Body = v
	return s
}

type UpdateSmsSignRequest struct {
	SignId   *string `json:"SignId,omitempty" xml:"SignId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s UpdateSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignRequest) SetSignId(v string) *UpdateSmsSignRequest {
	s.SignId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetRemark(v string) *UpdateSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSpaceId(v string) *UpdateSmsSignRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignName(v string) *UpdateSmsSignRequest {
	s.SignName = &v
	return s
}

type UpdateSmsSignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignResponseBody) SetRequestId(v string) *UpdateSmsSignResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSmsSignResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsSignResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignResponse) SetHeaders(v map[string]*string) *UpdateSmsSignResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsSignResponse) SetBody(v *UpdateSmsSignResponseBody) *UpdateSmsSignResponse {
	s.Body = v
	return s
}

type DeleteWebHostingCertificateRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteWebHostingCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingCertificateRequest) SetSpaceId(v string) *DeleteWebHostingCertificateRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteWebHostingCertificateRequest) SetDomain(v string) *DeleteWebHostingCertificateRequest {
	s.Domain = &v
	return s
}

type DeleteWebHostingCertificateResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebHostingCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingCertificateResponseBody) SetData(v bool) *DeleteWebHostingCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWebHostingCertificateResponseBody) SetRequestId(v string) *DeleteWebHostingCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebHostingCertificateResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebHostingCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebHostingCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingCertificateResponse) SetHeaders(v map[string]*string) *DeleteWebHostingCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebHostingCertificateResponse) SetBody(v *DeleteWebHostingCertificateResponseBody) *DeleteWebHostingCertificateResponse {
	s.Body = v
	return s
}

type QueryDBBackupDumpTimesRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s QueryDBBackupDumpTimesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesRequest) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesRequest) SetSpaceId(v string) *QueryDBBackupDumpTimesRequest {
	s.SpaceId = &v
	return s
}

type QueryDBBackupDumpTimesResponseBody struct {
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupDumpTimes []*QueryDBBackupDumpTimesResponseBodyBackupDumpTimes `json:"BackupDumpTimes,omitempty" xml:"BackupDumpTimes,omitempty" type:"Repeated"`
}

func (s QueryDBBackupDumpTimesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesResponseBody) SetRequestId(v string) *QueryDBBackupDumpTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDBBackupDumpTimesResponseBody) SetBackupDumpTimes(v []*QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) *QueryDBBackupDumpTimesResponseBody {
	s.BackupDumpTimes = v
	return s
}

type QueryDBBackupDumpTimesResponseBodyBackupDumpTimes struct {
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DumpTime *string `json:"DumpTime,omitempty" xml:"DumpTime,omitempty"`
}

func (s QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) SetBackupId(v string) *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes {
	s.BackupId = &v
	return s
}

func (s *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes) SetDumpTime(v string) *QueryDBBackupDumpTimesResponseBodyBackupDumpTimes {
	s.DumpTime = &v
	return s
}

type QueryDBBackupDumpTimesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryDBBackupDumpTimesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDBBackupDumpTimesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDBBackupDumpTimesResponse) GoString() string {
	return s.String()
}

func (s *QueryDBBackupDumpTimesResponse) SetHeaders(v map[string]*string) *QueryDBBackupDumpTimesResponse {
	s.Headers = v
	return s
}

func (s *QueryDBBackupDumpTimesResponse) SetBody(v *QueryDBBackupDumpTimesResponseBody) *QueryDBBackupDumpTimesResponse {
	s.Body = v
	return s
}

type DeployFunctionRequest struct {
	DeploymentId *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	SpaceId      *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeployFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeployFunctionRequest) SetDeploymentId(v string) *DeployFunctionRequest {
	s.DeploymentId = &v
	return s
}

func (s *DeployFunctionRequest) SetSpaceId(v string) *DeployFunctionRequest {
	s.SpaceId = &v
	return s
}

type DeployFunctionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeployFunctionResponseBody) SetRequestId(v string) *DeployFunctionResponseBody {
	s.RequestId = &v
	return s
}

type DeployFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeployFunctionResponse) SetHeaders(v map[string]*string) *DeployFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeployFunctionResponse) SetBody(v *DeployFunctionResponseBody) *DeployFunctionResponse {
	s.Body = v
	return s
}

type AttachSmsSignRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s AttachSmsSignRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachSmsSignRequest) GoString() string {
	return s.String()
}

func (s *AttachSmsSignRequest) SetSpaceId(v string) *AttachSmsSignRequest {
	s.SpaceId = &v
	return s
}

func (s *AttachSmsSignRequest) SetSignName(v string) *AttachSmsSignRequest {
	s.SignName = &v
	return s
}

type AttachSmsSignResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachSmsSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSmsSignResponseBody) SetRequestId(v string) *AttachSmsSignResponseBody {
	s.RequestId = &v
	return s
}

type AttachSmsSignResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachSmsSignResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachSmsSignResponse) GoString() string {
	return s.String()
}

func (s *AttachSmsSignResponse) SetHeaders(v map[string]*string) *AttachSmsSignResponse {
	s.Headers = v
	return s
}

func (s *AttachSmsSignResponse) SetBody(v *AttachSmsSignResponseBody) *AttachSmsSignResponse {
	s.Body = v
	return s
}

type UpdateServicePolicyRequest struct {
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s UpdateServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateServicePolicyRequest) SetSpaceId(v string) *UpdateServicePolicyRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetServiceName(v string) *UpdateServicePolicyRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetPolicy(v string) *UpdateServicePolicyRequest {
	s.Policy = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetCollectionName(v string) *UpdateServicePolicyRequest {
	s.CollectionName = &v
	return s
}

func (s *UpdateServicePolicyRequest) SetPolicyName(v string) *UpdateServicePolicyRequest {
	s.PolicyName = &v
	return s
}

type UpdateServicePolicyResponseBody struct {
	SpaceId        *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Policy         *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
}

func (s UpdateServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServicePolicyResponseBody) SetSpaceId(v string) *UpdateServicePolicyResponseBody {
	s.SpaceId = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetRequestId(v string) *UpdateServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetPolicy(v string) *UpdateServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetPolicyName(v string) *UpdateServicePolicyResponseBody {
	s.PolicyName = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetServiceName(v string) *UpdateServicePolicyResponseBody {
	s.ServiceName = &v
	return s
}

func (s *UpdateServicePolicyResponseBody) SetCollectionName(v string) *UpdateServicePolicyResponseBody {
	s.CollectionName = &v
	return s
}

type UpdateServicePolicyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateServicePolicyResponse) SetHeaders(v map[string]*string) *UpdateServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateServicePolicyResponse) SetBody(v *UpdateServicePolicyResponseBody) *UpdateServicePolicyResponse {
	s.Body = v
	return s
}

type AddCorsDomainRequest struct {
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s AddCorsDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCorsDomainRequest) GoString() string {
	return s.String()
}

func (s *AddCorsDomainRequest) SetDomain(v string) *AddCorsDomainRequest {
	s.Domain = &v
	return s
}

func (s *AddCorsDomainRequest) SetSpaceId(v string) *AddCorsDomainRequest {
	s.SpaceId = &v
	return s
}

type AddCorsDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainId  *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s AddCorsDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCorsDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddCorsDomainResponseBody) SetRequestId(v string) *AddCorsDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCorsDomainResponseBody) SetDomainId(v string) *AddCorsDomainResponseBody {
	s.DomainId = &v
	return s
}

type AddCorsDomainResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCorsDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCorsDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCorsDomainResponse) GoString() string {
	return s.String()
}

func (s *AddCorsDomainResponse) SetHeaders(v map[string]*string) *AddCorsDomainResponse {
	s.Headers = v
	return s
}

func (s *AddCorsDomainResponse) SetBody(v *AddCorsDomainResponseBody) *AddCorsDomainResponse {
	s.Body = v
	return s
}

type DescribeWebHostingFileRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s DescribeWebHostingFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileRequest) SetSpaceId(v string) *DescribeWebHostingFileRequest {
	s.SpaceId = &v
	return s
}

func (s *DescribeWebHostingFileRequest) SetFilePath(v string) *DescribeWebHostingFileRequest {
	s.FilePath = &v
	return s
}

type DescribeWebHostingFileResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeWebHostingFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeWebHostingFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileResponseBody) SetRequestId(v string) *DescribeWebHostingFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebHostingFileResponseBody) SetData(v *DescribeWebHostingFileResponseBodyData) *DescribeWebHostingFileResponseBody {
	s.Data = v
	return s
}

type DescribeWebHostingFileResponseBodyData struct {
	FilePath         *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	ContentType      *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ETag             *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Exists           *bool   `json:"Exists,omitempty" xml:"Exists,omitempty"`
	LastModifiedTime *int64  `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	SignedUrl        *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s DescribeWebHostingFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileResponseBodyData) SetFilePath(v string) *DescribeWebHostingFileResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetContentType(v string) *DescribeWebHostingFileResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetETag(v string) *DescribeWebHostingFileResponseBodyData {
	s.ETag = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetSize(v int64) *DescribeWebHostingFileResponseBodyData {
	s.Size = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetExists(v bool) *DescribeWebHostingFileResponseBodyData {
	s.Exists = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetLastModifiedTime(v int64) *DescribeWebHostingFileResponseBodyData {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeWebHostingFileResponseBodyData) SetSignedUrl(v string) *DescribeWebHostingFileResponseBodyData {
	s.SignedUrl = &v
	return s
}

type DescribeWebHostingFileResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebHostingFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebHostingFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebHostingFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebHostingFileResponse) SetHeaders(v map[string]*string) *DescribeWebHostingFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebHostingFileResponse) SetBody(v *DescribeWebHostingFileResponseBody) *DescribeWebHostingFileResponse {
	s.Body = v
	return s
}

type UpdateSmsTemplateRequest struct {
	SpaceId         *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TemplateCode    *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateType    *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateSmsTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateRequest) SetSpaceId(v string) *UpdateSmsTemplateRequest {
	s.SpaceId = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateCode(v string) *UpdateSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateType(v string) *UpdateSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateName(v string) *UpdateSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateContent(v string) *UpdateSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetRemark(v string) *UpdateSmsTemplateRequest {
	s.Remark = &v
	return s
}

type UpdateSmsTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmsTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateResponseBody) SetRequestId(v string) *UpdateSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSmsTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSmsTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSmsTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmsTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateResponse) SetHeaders(v map[string]*string) *UpdateSmsTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsTemplateResponse) SetBody(v *UpdateSmsTemplateResponseBody) *UpdateSmsTemplateResponse {
	s.Body = v
	return s
}

type VerifyBuiltinFunctionTemplateRequest struct {
	BuiltinFunctionTemplateId *string `json:"BuiltinFunctionTemplateId,omitempty" xml:"BuiltinFunctionTemplateId,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VerifyBuiltinFunctionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBuiltinFunctionTemplateRequest) GoString() string {
	return s.String()
}

func (s *VerifyBuiltinFunctionTemplateRequest) SetBuiltinFunctionTemplateId(v string) *VerifyBuiltinFunctionTemplateRequest {
	s.BuiltinFunctionTemplateId = &v
	return s
}

func (s *VerifyBuiltinFunctionTemplateRequest) SetStatus(v string) *VerifyBuiltinFunctionTemplateRequest {
	s.Status = &v
	return s
}

type VerifyBuiltinFunctionTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyBuiltinFunctionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyBuiltinFunctionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyBuiltinFunctionTemplateResponseBody) SetRequestId(v string) *VerifyBuiltinFunctionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type VerifyBuiltinFunctionTemplateResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyBuiltinFunctionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyBuiltinFunctionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBuiltinFunctionTemplateResponse) GoString() string {
	return s.String()
}

func (s *VerifyBuiltinFunctionTemplateResponse) SetHeaders(v map[string]*string) *VerifyBuiltinFunctionTemplateResponse {
	s.Headers = v
	return s
}

func (s *VerifyBuiltinFunctionTemplateResponse) SetBody(v *VerifyBuiltinFunctionTemplateResponseBody) *VerifyBuiltinFunctionTemplateResponse {
	s.Body = v
	return s
}

type DeleteWebHostingFileRequest struct {
	SpaceId  *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s DeleteWebHostingFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingFileRequest) SetSpaceId(v string) *DeleteWebHostingFileRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteWebHostingFileRequest) SetFilePath(v string) *DeleteWebHostingFileRequest {
	s.FilePath = &v
	return s
}

type DeleteWebHostingFileResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebHostingFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingFileResponseBody) SetData(v bool) *DeleteWebHostingFileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWebHostingFileResponseBody) SetRequestId(v string) *DeleteWebHostingFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebHostingFileResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebHostingFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebHostingFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebHostingFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebHostingFileResponse) SetHeaders(v map[string]*string) *DeleteWebHostingFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebHostingFileResponse) SetBody(v *DeleteWebHostingFileResponseBody) *DeleteWebHostingFileResponse {
	s.Body = v
	return s
}

type ListWebHostingCustomDomainsRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s ListWebHostingCustomDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsRequest) SetSpaceId(v string) *ListWebHostingCustomDomainsRequest {
	s.SpaceId = &v
	return s
}

type ListWebHostingCustomDomainsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListWebHostingCustomDomainsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListWebHostingCustomDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsResponseBody) SetRequestId(v string) *ListWebHostingCustomDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBody) SetData(v []*ListWebHostingCustomDomainsResponseBodyData) *ListWebHostingCustomDomainsResponseBody {
	s.Data = v
	return s
}

type ListWebHostingCustomDomainsResponseBodyData struct {
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Domain                   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	UpdateTime               *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	SslProtocol              *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
	ForceRedirectType        *string `json:"ForceRedirectType,omitempty" xml:"ForceRedirectType,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime               *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Cname                    *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	EnableCors               *bool   `json:"EnableCors,omitempty" xml:"EnableCors,omitempty"`
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
}

func (s ListWebHostingCustomDomainsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetStatus(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetDomain(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Domain = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetUpdateTime(v int64) *ListWebHostingCustomDomainsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetSslProtocol(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.SslProtocol = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetForceRedirectType(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.ForceRedirectType = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetDescription(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetCreateTime(v int64) *ListWebHostingCustomDomainsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetCname(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.Cname = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetEnableCors(v bool) *ListWebHostingCustomDomainsResponseBodyData {
	s.EnableCors = &v
	return s
}

func (s *ListWebHostingCustomDomainsResponseBodyData) SetAccessControlAllowOrigin(v string) *ListWebHostingCustomDomainsResponseBodyData {
	s.AccessControlAllowOrigin = &v
	return s
}

type ListWebHostingCustomDomainsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWebHostingCustomDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWebHostingCustomDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWebHostingCustomDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListWebHostingCustomDomainsResponse) SetHeaders(v map[string]*string) *ListWebHostingCustomDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListWebHostingCustomDomainsResponse) SetBody(v *ListWebHostingCustomDomainsResponseBody) *ListWebHostingCustomDomainsResponse {
	s.Body = v
	return s
}

type RunDBCommandRequest struct {
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	Body    *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s RunDBCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDBCommandRequest) GoString() string {
	return s.String()
}

func (s *RunDBCommandRequest) SetSpaceId(v string) *RunDBCommandRequest {
	s.SpaceId = &v
	return s
}

func (s *RunDBCommandRequest) SetBody(v string) *RunDBCommandRequest {
	s.Body = &v
	return s
}

type RunDBCommandResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	AffectedDocs *int32  `json:"AffectedDocs,omitempty" xml:"AffectedDocs,omitempty"`
}

func (s RunDBCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunDBCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunDBCommandResponseBody) SetRequestId(v string) *RunDBCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDBCommandResponseBody) SetResult(v string) *RunDBCommandResponseBody {
	s.Result = &v
	return s
}

func (s *RunDBCommandResponseBody) SetAffectedDocs(v int32) *RunDBCommandResponseBody {
	s.AffectedDocs = &v
	return s
}

type RunDBCommandResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunDBCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunDBCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunDBCommandResponse) GoString() string {
	return s.String()
}

func (s *RunDBCommandResponse) SetHeaders(v map[string]*string) *RunDBCommandResponse {
	s.Headers = v
	return s
}

func (s *RunDBCommandResponse) SetBody(v *RunDBCommandResponseBody) *RunDBCommandResponse {
	s.Body = v
	return s
}

type DeleteFunctionRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
}

func (s DeleteFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) SetName(v string) *DeleteFunctionRequest {
	s.Name = &v
	return s
}

func (s *DeleteFunctionRequest) SetSpaceId(v string) *DeleteFunctionRequest {
	s.SpaceId = &v
	return s
}

type DeleteFunctionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetBody(v *DeleteFunctionResponseBody) *DeleteFunctionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mpserverless"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) RunFunctionWithOptions(request *RunFunctionRequest, runtime *util.RuntimeOptions) (_result *RunFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunFunction(request *RunFunctionRequest) (_result *RunFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunFunctionResponse{}
	_body, _err := client.RunFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionWithOptions(request *ListFunctionRequest, runtime *util.RuntimeOptions) (_result *ListFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunction(request *ListFunctionRequest) (_result *ListFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionResponse{}
	_body, _err := client.ListFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingCertificateDetailWithOptions(request *GetWebHostingCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebHostingCertificateDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebHostingCertificateDetail"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingCertificateDetail(request *GetWebHostingCertificateDetailRequest) (_result *GetWebHostingCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingCertificateDetailResponse{}
	_body, _err := client.GetWebHostingCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSpaceWithOptions(request *UpdateSpaceRequest, runtime *util.RuntimeOptions) (_result *UpdateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSpace"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSpace(request *UpdateSpaceRequest) (_result *UpdateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSpaceResponse{}
	_body, _err := client.UpdateSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainConfigWithOptions(request *SaveWebHostingCustomDomainConfigRequest, runtime *util.RuntimeOptions) (_result *SaveWebHostingCustomDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveWebHostingCustomDomainConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveWebHostingCustomDomainConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainConfig(request *SaveWebHostingCustomDomainConfigRequest) (_result *SaveWebHostingCustomDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebHostingCustomDomainConfigResponse{}
	_body, _err := client.SaveWebHostingCustomDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionSpecWithOptions(runtime *util.RuntimeOptions) (_result *ListFunctionSpecResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListFunctionSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunctionSpec"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionSpec() (_result *ListFunctionSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionSpecResponse{}
	_body, _err := client.ListFunctionSpecWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWechatOpenPlatformConfigWithOptions(request *DeleteWechatOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteWechatOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWechatOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWechatOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWechatOpenPlatformConfig(request *DeleteWechatOpenPlatformConfigRequest) (_result *DeleteWechatOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWechatOpenPlatformConfigResponse{}
	_body, _err := client.DeleteWechatOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSpaceWithOptions(request *CreateSpaceRequest, runtime *util.RuntimeOptions) (_result *CreateSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSpace"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSpace(request *CreateSpaceRequest) (_result *CreateSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSpaceResponse{}
	_body, _err := client.CreateSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenProductWithOptions(request *OpenProductRequest, runtime *util.RuntimeOptions) (_result *OpenProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenProductResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenProduct"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenProduct(request *OpenProductRequest) (_result *OpenProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenProductResponse{}
	_body, _err := client.OpenProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenServiceWithOptions(request *OpenServiceRequest, runtime *util.RuntimeOptions) (_result *OpenServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenService"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenService(request *OpenServiceRequest) (_result *OpenServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenServiceResponse{}
	_body, _err := client.OpenServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSpaceWithOptions(request *DeleteSpaceRequest, runtime *util.RuntimeOptions) (_result *DeleteSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSpace"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpace(request *DeleteSpaceRequest) (_result *DeleteSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSpaceResponse{}
	_body, _err := client.DeleteSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAntOpenPlatformConfigWithOptions(request *DeleteAntOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteAntOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAntOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAntOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAntOpenPlatformConfig(request *DeleteAntOpenPlatformConfigRequest) (_result *DeleteAntOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAntOpenPlatformConfigResponse{}
	_body, _err := client.DeleteAntOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFCOpenStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeFCOpenStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeFCOpenStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFCOpenStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFCOpenStatus() (_result *DescribeFCOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFCOpenStatusResponse{}
	_body, _err := client.DescribeFCOpenStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileUploadSignedUrlWithOptions(request *DescribeFileUploadSignedUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeFileUploadSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFileUploadSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFileUploadSignedUrl"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileUploadSignedUrl(request *DescribeFileUploadSignedUrlRequest) (_result *DescribeFileUploadSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileUploadSignedUrlResponse{}
	_body, _err := client.DescribeFileUploadSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBImportTaskStatusWithOptions(request *QueryDBImportTaskStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDBImportTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDBImportTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDBImportTaskStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBImportTaskStatus(request *QueryDBImportTaskStatusRequest) (_result *QueryDBImportTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBImportTaskStatusResponse{}
	_body, _err := client.QueryDBImportTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterFileWithOptions(request *RegisterFileRequest, runtime *util.RuntimeOptions) (_result *RegisterFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterFile(request *RegisterFileRequest) (_result *RegisterFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterFileResponse{}
	_body, _err := client.RegisterFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAntOpenPlatformConfigWithOptions(request *SaveAntOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *SaveAntOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveAntOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveAntOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveAntOpenPlatformConfig(request *SaveAntOpenPlatformConfigRequest) (_result *SaveAntOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveAntOpenPlatformConfigResponse{}
	_body, _err := client.SaveAntOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFunctionWithOptions(request *DescribeFunctionRequest, runtime *util.RuntimeOptions) (_result *DescribeFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFunction(request *DescribeFunctionRequest) (_result *DescribeFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFunctionResponse{}
	_body, _err := client.DescribeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenWebHostingServiceWithOptions(request *OpenWebHostingServiceRequest, runtime *util.RuntimeOptions) (_result *OpenWebHostingServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenWebHostingServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenWebHostingService"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenWebHostingService(request *OpenWebHostingServiceRequest) (_result *OpenWebHostingServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenWebHostingServiceResponse{}
	_body, _err := client.OpenWebHostingServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmsSignWithOptions(request *DescribeSmsSignRequest, runtime *util.RuntimeOptions) (_result *DescribeSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmsSign"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmsSign(request *DescribeSmsSignRequest) (_result *DescribeSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmsSignResponse{}
	_body, _err := client.DescribeSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableCertificatesWithOptions(request *ListAvailableCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListAvailableCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAvailableCertificatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAvailableCertificates"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableCertificates(request *ListAvailableCertificatesRequest) (_result *ListAvailableCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableCertificatesResponse{}
	_body, _err := client.ListAvailableCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenPlatformConfigWithOptions(request *ListOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *ListOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenPlatformConfig(request *ListOpenPlatformConfigRequest) (_result *ListOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenPlatformConfigResponse{}
	_body, _err := client.ListOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebHostingConfigWithOptions(request *ModifyWebHostingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyWebHostingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebHostingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebHostingConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebHostingConfig(request *ModifyWebHostingConfigRequest) (_result *ModifyWebHostingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebHostingConfigResponse{}
	_body, _err := client.ModifyWebHostingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSmsSignWithOptions(request *DeleteSmsSignRequest, runtime *util.RuntimeOptions) (_result *DeleteSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSmsSign"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (_result *DeleteSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.DeleteSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmsOpenStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeSmsOpenStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeSmsOpenStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmsOpenStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmsOpenStatus() (_result *DescribeSmsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmsOpenStatusResponse{}
	_body, _err := client.DescribeSmsOpenStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSpaceWithOptions(request *ListSpaceRequest, runtime *util.RuntimeOptions) (_result *ListSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSpace"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSpace(request *ListSpaceRequest) (_result *ListSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSpaceResponse{}
	_body, _err := client.ListSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBCollectionWithOptions(request *DeleteDBCollectionRequest, runtime *util.RuntimeOptions) (_result *DeleteDBCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBCollectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBCollection"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBCollection(request *DeleteDBCollectionRequest) (_result *DeleteDBCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBCollectionResponse{}
	_body, _err := client.DeleteDBCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionDeploymentWithOptions(request *CreateFunctionDeploymentRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFunctionDeploymentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFunctionDeployment"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunctionDeployment(request *CreateFunctionDeploymentRequest) (_result *CreateFunctionDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionDeploymentResponse{}
	_body, _err := client.CreateFunctionDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingUploadCredentialWithOptions(request *GetWebHostingUploadCredentialRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingUploadCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebHostingUploadCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebHostingUploadCredential"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingUploadCredential(request *GetWebHostingUploadCredentialRequest) (_result *GetWebHostingUploadCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingUploadCredentialResponse{}
	_body, _err := client.GetWebHostingUploadCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionDeploymentWithOptions(request *ListFunctionDeploymentRequest, runtime *util.RuntimeOptions) (_result *ListFunctionDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFunctionDeploymentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunctionDeployment"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionDeployment(request *ListFunctionDeploymentRequest) (_result *ListFunctionDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionDeploymentResponse{}
	_body, _err := client.ListFunctionDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDingtalkOpenPlatformConfigWithOptions(request *AddDingtalkOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *AddDingtalkOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddDingtalkOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDingtalkOpenPlatformConfig(request *AddDingtalkOpenPlatformConfigRequest) (_result *AddDingtalkOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.AddDingtalkOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBRestoreTaskWithOptions(request *CreateDBRestoreTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDBRestoreTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBRestoreTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBRestoreTask"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBRestoreTask(request *CreateDBRestoreTaskRequest) (_result *CreateDBRestoreTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBRestoreTaskResponse{}
	_body, _err := client.CreateDBRestoreTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachWebHostingCertificateWithOptions(request *AttachWebHostingCertificateRequest, runtime *util.RuntimeOptions) (_result *AttachWebHostingCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachWebHostingCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachWebHostingCertificate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachWebHostingCertificate(request *AttachWebHostingCertificateRequest) (_result *AttachWebHostingCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachWebHostingCertificateResponse{}
	_body, _err := client.AttachWebHostingCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFileWithOptions(request *ListFileRequest, runtime *util.RuntimeOptions) (_result *ListFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFile(request *ListFileRequest) (_result *ListFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFileResponse{}
	_body, _err := client.ListFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBRestoreTaskStatusWithOptions(request *QueryDBRestoreTaskStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDBRestoreTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDBRestoreTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDBRestoreTaskStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBRestoreTaskStatus(request *QueryDBRestoreTaskStatusRequest) (_result *QueryDBRestoreTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBRestoreTaskStatusResponse{}
	_body, _err := client.QueryDBRestoreTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyWebHostingDomainOwnerWithOptions(request *VerifyWebHostingDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyWebHostingDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyWebHostingDomainOwnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyWebHostingDomainOwner"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyWebHostingDomainOwner(request *VerifyWebHostingDomainOwnerRequest) (_result *VerifyWebHostingDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyWebHostingDomainOwnerResponse{}
	_body, _err := client.VerifyWebHostingDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSmsTemplateWithOptions(request *DeleteSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSmsTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (_result *DeleteSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.DeleteSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBExportTaskStatusWithOptions(request *QueryDBExportTaskStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDBExportTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDBExportTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDBExportTaskStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBExportTaskStatus(request *QueryDBExportTaskStatusRequest) (_result *QueryDBExportTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBExportTaskStatusResponse{}
	_body, _err := client.QueryDBExportTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBImportTaskWithOptions(request *CreateDBImportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDBImportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBImportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBImportTask"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBImportTask(request *CreateDBImportTaskRequest) (_result *CreateDBImportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBImportTaskResponse{}
	_body, _err := client.CreateDBImportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckSmsHasAuthorizedToMPSWithOptions(request *CheckSmsHasAuthorizedToMPSRequest, runtime *util.RuntimeOptions) (_result *CheckSmsHasAuthorizedToMPSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckSmsHasAuthorizedToMPSResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckSmsHasAuthorizedToMPS"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSmsHasAuthorizedToMPS(request *CheckSmsHasAuthorizedToMPSRequest) (_result *CheckSmsHasAuthorizedToMPSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSmsHasAuthorizedToMPSResponse{}
	_body, _err := client.CheckSmsHasAuthorizedToMPSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServicePolicyWithOptions(request *DescribeServicePolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeServicePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeServicePolicy"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServicePolicy(request *DescribeServicePolicyRequest) (_result *DescribeServicePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServicePolicyResponse{}
	_body, _err := client.DescribeServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSmsTemplatesWithOptions(request *ListSmsTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListSmsTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSmsTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSmsTemplates"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSmsTemplates(request *ListSmsTemplatesRequest) (_result *ListSmsTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSmsTemplatesResponse{}
	_body, _err := client.ListSmsTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBBackupCollectionsWithOptions(request *QueryDBBackupCollectionsRequest, runtime *util.RuntimeOptions) (_result *QueryDBBackupCollectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDBBackupCollectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDBBackupCollections"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBBackupCollections(request *QueryDBBackupCollectionsRequest) (_result *QueryDBBackupCollectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBBackupCollectionsResponse{}
	_body, _err := client.QueryDBBackupCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceStatusWithOptions(request *QueryServiceStatusRequest, runtime *util.RuntimeOptions) (_result *QueryServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryServiceStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceStatus(request *QueryServiceStatusRequest) (_result *QueryServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServiceStatusResponse{}
	_body, _err := client.QueryServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSpaceClientConfigWithOptions(request *DescribeSpaceClientConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeSpaceClientConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSpaceClientConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSpaceClientConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSpaceClientConfig(request *DescribeSpaceClientConfigRequest) (_result *DescribeSpaceClientConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpaceClientConfigResponse{}
	_body, _err := client.DescribeSpaceClientConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveBuiltinFunctionTemplateWithOptions(request *SaveBuiltinFunctionTemplateRequest, runtime *util.RuntimeOptions) (_result *SaveBuiltinFunctionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveBuiltinFunctionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveBuiltinFunctionTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveBuiltinFunctionTemplate(request *SaveBuiltinFunctionTemplateRequest) (_result *SaveBuiltinFunctionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveBuiltinFunctionTemplateResponse{}
	_body, _err := client.SaveBuiltinFunctionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeISVFileUploadSignedUrlWithOptions(request *DescribeISVFileUploadSignedUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeISVFileUploadSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeISVFileUploadSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeISVFileUploadSignedUrl"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeISVFileUploadSignedUrl(request *DescribeISVFileUploadSignedUrlRequest) (_result *DescribeISVFileUploadSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeISVFileUploadSignedUrlResponse{}
	_body, _err := client.DescribeISVFileUploadSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBuiltinFunctionTemplateWithOptions(request *CreateBuiltinFunctionTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateBuiltinFunctionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBuiltinFunctionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBuiltinFunctionTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBuiltinFunctionTemplate(request *CreateBuiltinFunctionTemplateRequest) (_result *CreateBuiltinFunctionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBuiltinFunctionTemplateResponse{}
	_body, _err := client.CreateBuiltinFunctionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingStatusWithOptions(request *GetWebHostingStatusRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebHostingStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebHostingStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingStatus(request *GetWebHostingStatusRequest) (_result *GetWebHostingStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingStatusResponse{}
	_body, _err := client.GetWebHostingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionLogWithOptions(request *ListFunctionLogRequest, runtime *util.RuntimeOptions) (_result *ListFunctionLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFunctionLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFunctionLog"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionLog(request *ListFunctionLogRequest) (_result *ListFunctionLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionLogResponse{}
	_body, _err := client.ListFunctionLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWebHostingFilesWithOptions(request *ListWebHostingFilesRequest, runtime *util.RuntimeOptions) (_result *ListWebHostingFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListWebHostingFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWebHostingFiles"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWebHostingFiles(request *ListWebHostingFilesRequest) (_result *ListWebHostingFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWebHostingFilesResponse{}
	_body, _err := client.ListWebHostingFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileWithOptions(request *DescribeFileRequest, runtime *util.RuntimeOptions) (_result *DescribeFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFile(request *DescribeFileRequest) (_result *DescribeFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileResponse{}
	_body, _err := client.DescribeFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveWebHostingFileWithOptions(request *MoveWebHostingFileRequest, runtime *util.RuntimeOptions) (_result *MoveWebHostingFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveWebHostingFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveWebHostingFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveWebHostingFile(request *MoveWebHostingFileRequest) (_result *MoveWebHostingFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveWebHostingFileResponse{}
	_body, _err := client.MoveWebHostingFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSmsTemplateWithOptions(request *CreateSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSmsTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSmsTemplate(request *CreateSmsTemplateRequest) (_result *CreateSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CreateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmsTemplateStatusWithOptions(request *DescribeSmsTemplateStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSmsTemplateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSmsTemplateStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmsTemplateStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmsTemplateStatus(request *DescribeSmsTemplateStatusRequest) (_result *DescribeSmsTemplateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmsTemplateStatusResponse{}
	_body, _err := client.DescribeSmsTemplateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindWebHostingCustomDomainWithOptions(request *BindWebHostingCustomDomainRequest, runtime *util.RuntimeOptions) (_result *BindWebHostingCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindWebHostingCustomDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindWebHostingCustomDomain"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindWebHostingCustomDomain(request *BindWebHostingCustomDomainRequest) (_result *BindWebHostingCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindWebHostingCustomDomainResponse{}
	_body, _err := client.BindWebHostingCustomDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDingtalkOpenPlatformConfigWithOptions(request *DeleteDingtalkOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDingtalkOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDingtalkOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDingtalkOpenPlatformConfig(request *DeleteDingtalkOpenPlatformConfigRequest) (_result *DeleteDingtalkOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.DeleteDingtalkOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExtensionsWithOptions(request *ListExtensionsRequest, runtime *util.RuntimeOptions) (_result *ListExtensionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListExtensionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExtensions"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExtensions(request *ListExtensionsRequest) (_result *ListExtensionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExtensionsResponse{}
	_body, _err := client.ListExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSmsServiceWithOptions(runtime *util.RuntimeOptions) (_result *EnableSmsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &EnableSmsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableSmsService"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSmsService() (_result *EnableSmsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSmsServiceResponse{}
	_body, _err := client.EnableSmsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseBuiltinFunctionTemplateWithOptions(request *ReleaseBuiltinFunctionTemplateRequest, runtime *util.RuntimeOptions) (_result *ReleaseBuiltinFunctionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseBuiltinFunctionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseBuiltinFunctionTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseBuiltinFunctionTemplate(request *ReleaseBuiltinFunctionTemplateRequest) (_result *ReleaseBuiltinFunctionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseBuiltinFunctionTemplateResponse{}
	_body, _err := client.ReleaseBuiltinFunctionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSmsSignWithOptions(request *CreateSmsSignRequest, runtime *util.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSmsSign"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (_result *CreateSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CreateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHttpTriggerConfigWithOptions(request *UpdateHttpTriggerConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateHttpTriggerConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateHttpTriggerConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateHttpTriggerConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHttpTriggerConfig(request *UpdateHttpTriggerConfigRequest) (_result *UpdateHttpTriggerConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHttpTriggerConfigResponse{}
	_body, _err := client.UpdateHttpTriggerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetServerSecretWithOptions(request *ResetServerSecretRequest, runtime *util.RuntimeOptions) (_result *ResetServerSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetServerSecretResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetServerSecret"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetServerSecret(request *ResetServerSecretRequest) (_result *ResetServerSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetServerSecretResponse{}
	_body, _err := client.ResetServerSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingDomainVerificationContentWithOptions(request *GetWebHostingDomainVerificationContentRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingDomainVerificationContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebHostingDomainVerificationContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebHostingDomainVerificationContent"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingDomainVerificationContent(request *GetWebHostingDomainVerificationContentRequest) (_result *GetWebHostingDomainVerificationContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingDomainVerificationContentResponse{}
	_body, _err := client.GetWebHostingDomainVerificationContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDingtalkOpenPlatformConfigWithOptions(request *UpdateDingtalkOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateDingtalkOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDingtalkOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDingtalkOpenPlatformConfig(request *UpdateDingtalkOpenPlatformConfigRequest) (_result *UpdateDingtalkOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDingtalkOpenPlatformConfigResponse{}
	_body, _err := client.UpdateDingtalkOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMpServerlessRoleExistsWithOptions(request *CheckMpServerlessRoleExistsRequest, runtime *util.RuntimeOptions) (_result *CheckMpServerlessRoleExistsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckMpServerlessRoleExistsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckMpServerlessRoleExists"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMpServerlessRoleExists(request *CheckMpServerlessRoleExistsRequest) (_result *CheckMpServerlessRoleExistsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMpServerlessRoleExistsResponse{}
	_body, _err := client.CheckMpServerlessRoleExistsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableExtensionWithOptions(request *EnableExtensionRequest, runtime *util.RuntimeOptions) (_result *EnableExtensionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableExtensionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableExtension"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableExtension(request *EnableExtensionRequest) (_result *EnableExtensionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableExtensionResponse{}
	_body, _err := client.EnableExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSmsSignsForAccountWithOptions(request *ListSmsSignsForAccountRequest, runtime *util.RuntimeOptions) (_result *ListSmsSignsForAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSmsSignsForAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSmsSignsForAccount"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSmsSignsForAccount(request *ListSmsSignsForAccountRequest) (_result *ListSmsSignsForAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSmsSignsForAccountResponse{}
	_body, _err := client.ListSmsSignsForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCorsDomainsWithOptions(request *ListCorsDomainsRequest, runtime *util.RuntimeOptions) (_result *ListCorsDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCorsDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCorsDomains"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCorsDomains(request *ListCorsDomainsRequest) (_result *ListCorsDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCorsDomainsResponse{}
	_body, _err := client.ListCorsDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDingtalkOpenPlatformConfigsWithOptions(request *ListDingtalkOpenPlatformConfigsRequest, runtime *util.RuntimeOptions) (_result *ListDingtalkOpenPlatformConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDingtalkOpenPlatformConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDingtalkOpenPlatformConfigs"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDingtalkOpenPlatformConfigs(request *ListDingtalkOpenPlatformConfigsRequest) (_result *ListDingtalkOpenPlatformConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDingtalkOpenPlatformConfigsResponse{}
	_body, _err := client.ListDingtalkOpenPlatformConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBExportTaskWithOptions(request *CreateDBExportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDBExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBExportTask"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBExportTask(request *CreateDBExportTaskRequest) (_result *CreateDBExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBExportTaskResponse{}
	_body, _err := client.CreateDBExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebHostingConfigWithOptions(request *GetWebHostingConfigRequest, runtime *util.RuntimeOptions) (_result *GetWebHostingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWebHostingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWebHostingConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebHostingConfig(request *GetWebHostingConfigRequest) (_result *GetWebHostingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWebHostingConfigResponse{}
	_body, _err := client.GetWebHostingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindWebHostingCustomDomainWithOptions(request *UnbindWebHostingCustomDomainRequest, runtime *util.RuntimeOptions) (_result *UnbindWebHostingCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindWebHostingCustomDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindWebHostingCustomDomain"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindWebHostingCustomDomain(request *UnbindWebHostingCustomDomainRequest) (_result *UnbindWebHostingCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindWebHostingCustomDomainResponse{}
	_body, _err := client.UnbindWebHostingCustomDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmsTemplateWithOptions(request *DescribeSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmsTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmsTemplate(request *DescribeSmsTemplateRequest) (_result *DescribeSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmsTemplateResponse{}
	_body, _err := client.DescribeSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainCorsConfigWithOptions(request *SaveWebHostingCustomDomainCorsConfigRequest, runtime *util.RuntimeOptions) (_result *SaveWebHostingCustomDomainCorsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveWebHostingCustomDomainCorsConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveWebHostingCustomDomainCorsConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebHostingCustomDomainCorsConfig(request *SaveWebHostingCustomDomainCorsConfigRequest) (_result *SaveWebHostingCustomDomainCorsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebHostingCustomDomainCorsConfigResponse{}
	_body, _err := client.SaveWebHostingCustomDomainCorsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteWebHostingFilesWithOptions(request *BatchDeleteWebHostingFilesRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteWebHostingFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchDeleteWebHostingFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchDeleteWebHostingFiles"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteWebHostingFiles(request *BatchDeleteWebHostingFilesRequest) (_result *BatchDeleteWebHostingFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteWebHostingFilesResponse{}
	_body, _err := client.BatchDeleteWebHostingFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCorsDomainWithOptions(request *DeleteCorsDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteCorsDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCorsDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCorsDomain"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCorsDomain(request *DeleteCorsDomainRequest) (_result *DeleteCorsDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCorsDomainResponse{}
	_body, _err := client.DeleteCorsDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHttpTriggerConfigWithOptions(request *DescribeHttpTriggerConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeHttpTriggerConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHttpTriggerConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHttpTriggerConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHttpTriggerConfig(request *DescribeHttpTriggerConfigRequest) (_result *DescribeHttpTriggerConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHttpTriggerConfigResponse{}
	_body, _err := client.DescribeHttpTriggerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAppAuthTokenWithOptions(request *SaveAppAuthTokenRequest, runtime *util.RuntimeOptions) (_result *SaveAppAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveAppAuthTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveAppAuthToken"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveAppAuthToken(request *SaveAppAuthTokenRequest) (_result *SaveAppAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveAppAuthTokenResponse{}
	_body, _err := client.SaveAppAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSmsSignStatusWithOptions(request *DescribeSmsSignStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSmsSignStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSmsSignStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSmsSignStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSmsSignStatus(request *DescribeSmsSignStatusRequest) (_result *DescribeSmsSignStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSmsSignStatusResponse{}
	_body, _err := client.DescribeSmsSignStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWechatOpenPlatformConfigWithOptions(request *SaveWechatOpenPlatformConfigRequest, runtime *util.RuntimeOptions) (_result *SaveWechatOpenPlatformConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveWechatOpenPlatformConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveWechatOpenPlatformConfig"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWechatOpenPlatformConfig(request *SaveWechatOpenPlatformConfigRequest) (_result *SaveWechatOpenPlatformConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWechatOpenPlatformConfigResponse{}
	_body, _err := client.SaveWechatOpenPlatformConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSpaceWithOptions(request *DescribeSpaceRequest, runtime *util.RuntimeOptions) (_result *DescribeSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSpaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSpace"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSpace(request *DescribeSpaceRequest) (_result *DescribeSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpaceResponse{}
	_body, _err := client.DescribeSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDBCollectionWithOptions(request *RenameDBCollectionRequest, runtime *util.RuntimeOptions) (_result *RenameDBCollectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenameDBCollectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenameDBCollection"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDBCollection(request *RenameDBCollectionRequest) (_result *RenameDBCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameDBCollectionResponse{}
	_body, _err := client.RenameDBCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSmsSignsWithOptions(request *ListSmsSignsRequest, runtime *util.RuntimeOptions) (_result *ListSmsSignsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSmsSignsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSmsSigns"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSmsSigns(request *ListSmsSignsRequest) (_result *ListSmsSignsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSmsSignsResponse{}
	_body, _err := client.ListSmsSignsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProductOpenStatusWithOptions(request *DescribeProductOpenStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeProductOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProductOpenStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProductOpenStatus"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProductOpenStatus(request *DescribeProductOpenStatusRequest) (_result *DescribeProductOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProductOpenStatusResponse{}
	_body, _err := client.DescribeProductOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSmsSignWithOptions(request *UpdateSmsSignRequest, runtime *util.RuntimeOptions) (_result *UpdateSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSmsSign"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSmsSign(request *UpdateSmsSignRequest) (_result *UpdateSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSmsSignResponse{}
	_body, _err := client.UpdateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebHostingCertificateWithOptions(request *DeleteWebHostingCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteWebHostingCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWebHostingCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWebHostingCertificate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebHostingCertificate(request *DeleteWebHostingCertificateRequest) (_result *DeleteWebHostingCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebHostingCertificateResponse{}
	_body, _err := client.DeleteWebHostingCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDBBackupDumpTimesWithOptions(request *QueryDBBackupDumpTimesRequest, runtime *util.RuntimeOptions) (_result *QueryDBBackupDumpTimesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryDBBackupDumpTimesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryDBBackupDumpTimes"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDBBackupDumpTimes(request *QueryDBBackupDumpTimesRequest) (_result *QueryDBBackupDumpTimesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDBBackupDumpTimesResponse{}
	_body, _err := client.QueryDBBackupDumpTimesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployFunctionWithOptions(request *DeployFunctionRequest, runtime *util.RuntimeOptions) (_result *DeployFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployFunction(request *DeployFunctionRequest) (_result *DeployFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployFunctionResponse{}
	_body, _err := client.DeployFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachSmsSignWithOptions(request *AttachSmsSignRequest, runtime *util.RuntimeOptions) (_result *AttachSmsSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachSmsSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachSmsSign"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachSmsSign(request *AttachSmsSignRequest) (_result *AttachSmsSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachSmsSignResponse{}
	_body, _err := client.AttachSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServicePolicyWithOptions(request *UpdateServicePolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServicePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateServicePolicy"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServicePolicy(request *UpdateServicePolicyRequest) (_result *UpdateServicePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServicePolicyResponse{}
	_body, _err := client.UpdateServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCorsDomainWithOptions(request *AddCorsDomainRequest, runtime *util.RuntimeOptions) (_result *AddCorsDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddCorsDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddCorsDomain"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCorsDomain(request *AddCorsDomainRequest) (_result *AddCorsDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCorsDomainResponse{}
	_body, _err := client.AddCorsDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebHostingFileWithOptions(request *DescribeWebHostingFileRequest, runtime *util.RuntimeOptions) (_result *DescribeWebHostingFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebHostingFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebHostingFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebHostingFile(request *DescribeWebHostingFileRequest) (_result *DescribeWebHostingFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebHostingFileResponse{}
	_body, _err := client.DescribeWebHostingFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSmsTemplateWithOptions(request *UpdateSmsTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateSmsTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSmsTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSmsTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSmsTemplate(request *UpdateSmsTemplateRequest) (_result *UpdateSmsTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSmsTemplateResponse{}
	_body, _err := client.UpdateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyBuiltinFunctionTemplateWithOptions(request *VerifyBuiltinFunctionTemplateRequest, runtime *util.RuntimeOptions) (_result *VerifyBuiltinFunctionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyBuiltinFunctionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyBuiltinFunctionTemplate"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyBuiltinFunctionTemplate(request *VerifyBuiltinFunctionTemplateRequest) (_result *VerifyBuiltinFunctionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyBuiltinFunctionTemplateResponse{}
	_body, _err := client.VerifyBuiltinFunctionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebHostingFileWithOptions(request *DeleteWebHostingFileRequest, runtime *util.RuntimeOptions) (_result *DeleteWebHostingFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWebHostingFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWebHostingFile"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebHostingFile(request *DeleteWebHostingFileRequest) (_result *DeleteWebHostingFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebHostingFileResponse{}
	_body, _err := client.DeleteWebHostingFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWebHostingCustomDomainsWithOptions(request *ListWebHostingCustomDomainsRequest, runtime *util.RuntimeOptions) (_result *ListWebHostingCustomDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListWebHostingCustomDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWebHostingCustomDomains"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWebHostingCustomDomains(request *ListWebHostingCustomDomainsRequest) (_result *ListWebHostingCustomDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWebHostingCustomDomainsResponse{}
	_body, _err := client.ListWebHostingCustomDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunDBCommandWithOptions(request *RunDBCommandRequest, runtime *util.RuntimeOptions) (_result *RunDBCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunDBCommandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunDBCommand"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunDBCommand(request *RunDBCommandRequest) (_result *RunDBCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunDBCommandResponse{}
	_body, _err := client.RunDBCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFunction"), tea.String("2019-06-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
