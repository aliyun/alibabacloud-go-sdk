// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AppsDomain struct {
	// 域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s AppsDomain) String() string {
	return tea.Prettify(s)
}

func (s AppsDomain) GoString() string {
	return s.String()
}

func (s *AppsDomain) SetDomain(v string) *AppsDomain {
	s.Domain = &v
	return s
}

type AppsInteractiveVideoSettings struct {
	Domain      *AppsDomain      `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Severity    *AppsSecurity    `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Transcoding *AppsTranscoding `json:"Transcoding,omitempty" xml:"Transcoding,omitempty"`
	Type        *string          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppsInteractiveVideoSettings) String() string {
	return tea.Prettify(s)
}

func (s AppsInteractiveVideoSettings) GoString() string {
	return s.String()
}

func (s *AppsInteractiveVideoSettings) SetDomain(v *AppsDomain) *AppsInteractiveVideoSettings {
	s.Domain = v
	return s
}

func (s *AppsInteractiveVideoSettings) SetSeverity(v *AppsSecurity) *AppsInteractiveVideoSettings {
	s.Severity = v
	return s
}

func (s *AppsInteractiveVideoSettings) SetTranscoding(v *AppsTranscoding) *AppsInteractiveVideoSettings {
	s.Transcoding = v
	return s
}

func (s *AppsInteractiveVideoSettings) SetType(v string) *AppsInteractiveVideoSettings {
	s.Type = &v
	return s
}

type AppsSecurity struct {
	PlayUrlTtl *int32 `json:"PlayUrlTtl,omitempty" xml:"PlayUrlTtl,omitempty"`
}

func (s AppsSecurity) String() string {
	return tea.Prettify(s)
}

func (s AppsSecurity) GoString() string {
	return s.String()
}

func (s *AppsSecurity) SetPlayUrlTtl(v int32) *AppsSecurity {
	s.PlayUrlTtl = &v
	return s
}

type AppsSettings struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 云帐号Id
	CloudUid *string `json:"CloudUid,omitempty" xml:"CloudUid,omitempty"`
	// 客户非敏感信息
	Customer *CommonCustomer `json:"Customer,omitempty" xml:"Customer,omitempty"`
	// 互动视频配置
	InteractiveVideoSettings *AppsInteractiveVideoSettings `json:"InteractiveVideoSettings,omitempty" xml:"InteractiveVideoSettings,omitempty"`
}

func (s AppsSettings) String() string {
	return tea.Prettify(s)
}

func (s AppsSettings) GoString() string {
	return s.String()
}

func (s *AppsSettings) SetAppId(v string) *AppsSettings {
	s.AppId = &v
	return s
}

func (s *AppsSettings) SetCloudUid(v string) *AppsSettings {
	s.CloudUid = &v
	return s
}

func (s *AppsSettings) SetCustomer(v *CommonCustomer) *AppsSettings {
	s.Customer = v
	return s
}

func (s *AppsSettings) SetInteractiveVideoSettings(v *AppsInteractiveVideoSettings) *AppsSettings {
	s.InteractiveVideoSettings = v
	return s
}

type AppsTranscoding struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppsTranscoding) String() string {
	return tea.Prettify(s)
}

func (s AppsTranscoding) GoString() string {
	return s.String()
}

func (s *AppsTranscoding) SetFlowId(v string) *AppsTranscoding {
	s.FlowId = &v
	return s
}

func (s *AppsTranscoding) SetType(v string) *AppsTranscoding {
	s.Type = &v
	return s
}

type AssetsAuditAssetRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// AssetId
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AssetsAuditAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetsAuditAssetRequest) GoString() string {
	return s.String()
}

func (s *AssetsAuditAssetRequest) SetAppId(v string) *AssetsAuditAssetRequest {
	s.AppId = &v
	return s
}

func (s *AssetsAuditAssetRequest) SetAuditStatus(v string) *AssetsAuditAssetRequest {
	s.AuditStatus = &v
	return s
}

func (s *AssetsAuditAssetRequest) SetId(v string) *AssetsAuditAssetRequest {
	s.Id = &v
	return s
}

type AssetsAuditAssetResponse struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应状态
	Status *RpcStatus `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AssetsAuditAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetsAuditAssetResponse) GoString() string {
	return s.String()
}

func (s *AssetsAuditAssetResponse) SetRequestId(v string) *AssetsAuditAssetResponse {
	s.RequestId = &v
	return s
}

func (s *AssetsAuditAssetResponse) SetStatus(v *RpcStatus) *AssetsAuditAssetResponse {
	s.Status = v
	return s
}

type AssetsCreateAssetRequest struct {
	// Asset请求Item
	Asset *CommonSimpleAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
}

func (s AssetsCreateAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetsCreateAssetRequest) GoString() string {
	return s.String()
}

func (s *AssetsCreateAssetRequest) SetAsset(v *CommonSimpleAsset) *AssetsCreateAssetRequest {
	s.Asset = v
	return s
}

type AssetsCreateAssetResponse struct {
	// 资产信息
	Asset *CommonAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应状态
	Status *RpcStatus `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AssetsCreateAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetsCreateAssetResponse) GoString() string {
	return s.String()
}

func (s *AssetsCreateAssetResponse) SetAsset(v *CommonAsset) *AssetsCreateAssetResponse {
	s.Asset = v
	return s
}

func (s *AssetsCreateAssetResponse) SetRequestId(v string) *AssetsCreateAssetResponse {
	s.RequestId = &v
	return s
}

func (s *AssetsCreateAssetResponse) SetStatus(v *RpcStatus) *AssetsCreateAssetResponse {
	s.Status = v
	return s
}

type AssetsDeleteAssetRequest struct {
	// Asset
	Asset *CommonAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
}

func (s AssetsDeleteAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetsDeleteAssetRequest) GoString() string {
	return s.String()
}

func (s *AssetsDeleteAssetRequest) SetAsset(v *CommonAsset) *AssetsDeleteAssetRequest {
	s.Asset = v
	return s
}

type AssetsDeleteAssetResponse struct {
	// 资产信息
	Asset *CommonAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应状态
	Status *RpcStatus `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AssetsDeleteAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetsDeleteAssetResponse) GoString() string {
	return s.String()
}

func (s *AssetsDeleteAssetResponse) SetAsset(v *CommonAsset) *AssetsDeleteAssetResponse {
	s.Asset = v
	return s
}

func (s *AssetsDeleteAssetResponse) SetRequestId(v string) *AssetsDeleteAssetResponse {
	s.RequestId = &v
	return s
}

func (s *AssetsDeleteAssetResponse) SetStatus(v *RpcStatus) *AssetsDeleteAssetResponse {
	s.Status = v
	return s
}

type AssetsGetAssetRequest struct {
	// AssetId
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AssetsGetAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetsGetAssetRequest) GoString() string {
	return s.String()
}

func (s *AssetsGetAssetRequest) SetId(v string) *AssetsGetAssetRequest {
	s.Id = &v
	return s
}

type AssetsGetAssetResponse struct {
	// Asset
	Asset *CommonAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应状态
	Status *RpcStatus `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AssetsGetAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetsGetAssetResponse) GoString() string {
	return s.String()
}

func (s *AssetsGetAssetResponse) SetAsset(v *CommonAsset) *AssetsGetAssetResponse {
	s.Asset = v
	return s
}

func (s *AssetsGetAssetResponse) SetRequestId(v string) *AssetsGetAssetResponse {
	s.RequestId = &v
	return s
}

func (s *AssetsGetAssetResponse) SetStatus(v *RpcStatus) *AssetsGetAssetResponse {
	s.Status = v
	return s
}

type AssetsListAssetsRequest struct {
	// AppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Optional. Used to specify a subset of fields that should be
	// returned by a get operation or modified by an update operation.
	FieldMask *string `json:"FieldMask,omitempty" xml:"FieldMask,omitempty"`
	// 每页显示个数，最大支持20，参数为空默认显示个数为10。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页Token Optional.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 参数
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 订阅Topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s AssetsListAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetsListAssetsRequest) GoString() string {
	return s.String()
}

func (s *AssetsListAssetsRequest) SetAppId(v string) *AssetsListAssetsRequest {
	s.AppId = &v
	return s
}

func (s *AssetsListAssetsRequest) SetFieldMask(v string) *AssetsListAssetsRequest {
	s.FieldMask = &v
	return s
}

func (s *AssetsListAssetsRequest) SetMaxResults(v int32) *AssetsListAssetsRequest {
	s.MaxResults = &v
	return s
}

func (s *AssetsListAssetsRequest) SetNextToken(v string) *AssetsListAssetsRequest {
	s.NextToken = &v
	return s
}

func (s *AssetsListAssetsRequest) SetParams(v string) *AssetsListAssetsRequest {
	s.Params = &v
	return s
}

func (s *AssetsListAssetsRequest) SetTopic(v string) *AssetsListAssetsRequest {
	s.Topic = &v
	return s
}

type AssetsListAssetsResponse struct {
	// Asset列表
	Assets []*CommonAsset `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// 分页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应状态
	Status *RpcStatus `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AssetsListAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetsListAssetsResponse) GoString() string {
	return s.String()
}

func (s *AssetsListAssetsResponse) SetAssets(v []*CommonAsset) *AssetsListAssetsResponse {
	s.Assets = v
	return s
}

func (s *AssetsListAssetsResponse) SetNextToken(v string) *AssetsListAssetsResponse {
	s.NextToken = &v
	return s
}

func (s *AssetsListAssetsResponse) SetRequestId(v string) *AssetsListAssetsResponse {
	s.RequestId = &v
	return s
}

func (s *AssetsListAssetsResponse) SetStatus(v *RpcStatus) *AssetsListAssetsResponse {
	s.Status = v
	return s
}

type AssetsUpdateAssetRequest struct {
	// Asset
	Asset *CommonAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
}

func (s AssetsUpdateAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s AssetsUpdateAssetRequest) GoString() string {
	return s.String()
}

func (s *AssetsUpdateAssetRequest) SetAsset(v *CommonAsset) *AssetsUpdateAssetRequest {
	s.Asset = v
	return s
}

type AssetsUpdateAssetResponse struct {
	// 资产信息
	Asset *CommonAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 响应状态
	Status *RpcStatus `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AssetsUpdateAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s AssetsUpdateAssetResponse) GoString() string {
	return s.String()
}

func (s *AssetsUpdateAssetResponse) SetAsset(v *CommonAsset) *AssetsUpdateAssetResponse {
	s.Asset = v
	return s
}

func (s *AssetsUpdateAssetResponse) SetRequestId(v string) *AssetsUpdateAssetResponse {
	s.RequestId = &v
	return s
}

func (s *AssetsUpdateAssetResponse) SetStatus(v *RpcStatus) *AssetsUpdateAssetResponse {
	s.Status = v
	return s
}

type CommonAddress struct {
	// 地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 城市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 国家
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// 区域
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// zip
	Zip *string `json:"Zip,omitempty" xml:"Zip,omitempty"`
}

func (s CommonAddress) String() string {
	return tea.Prettify(s)
}

func (s CommonAddress) GoString() string {
	return s.String()
}

func (s *CommonAddress) SetAddress(v string) *CommonAddress {
	s.Address = &v
	return s
}

func (s *CommonAddress) SetCity(v string) *CommonAddress {
	s.City = &v
	return s
}

func (s *CommonAddress) SetCountry(v string) *CommonAddress {
	s.Country = &v
	return s
}

func (s *CommonAddress) SetState(v string) *CommonAddress {
	s.State = &v
	return s
}

func (s *CommonAddress) SetZip(v string) *CommonAddress {
	s.Zip = &v
	return s
}

type CommonAsset struct {
	// 行政区域地址
	Address *CommonAddress `json:"Address,omitempty" xml:"Address,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// 资产描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 扩展字段
	Extends map[string]interface{} `json:"Extends,omitempty" xml:"Extends,omitempty"`
	// 资产ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 图像资源
	Images []*CommonMediaResource `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 定义Label， eg:type:advertise 支持广告类型的label
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 经纬度地理位置
	Location *TypeLatLng `json:"Location,omitempty" xml:"Location,omitempty"`
	// 资产来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 资产状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 概要
	Synopsis *string `json:"Synopsis,omitempty" xml:"Synopsis,omitempty"`
	// 标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// 视频资源
	Videos []*CommonMediaResource `json:"Videos,omitempty" xml:"Videos,omitempty" type:"Repeated"`
}

func (s CommonAsset) String() string {
	return tea.Prettify(s)
}

func (s CommonAsset) GoString() string {
	return s.String()
}

func (s *CommonAsset) SetAddress(v *CommonAddress) *CommonAsset {
	s.Address = v
	return s
}

func (s *CommonAsset) SetAppId(v string) *CommonAsset {
	s.AppId = &v
	return s
}

func (s *CommonAsset) SetAuditStatus(v string) *CommonAsset {
	s.AuditStatus = &v
	return s
}

func (s *CommonAsset) SetAuthor(v string) *CommonAsset {
	s.Author = &v
	return s
}

func (s *CommonAsset) SetCreatedAt(v string) *CommonAsset {
	s.CreatedAt = &v
	return s
}

func (s *CommonAsset) SetDescription(v string) *CommonAsset {
	s.Description = &v
	return s
}

func (s *CommonAsset) SetExtends(v map[string]interface{}) *CommonAsset {
	s.Extends = v
	return s
}

func (s *CommonAsset) SetId(v string) *CommonAsset {
	s.Id = &v
	return s
}

func (s *CommonAsset) SetImages(v []*CommonMediaResource) *CommonAsset {
	s.Images = v
	return s
}

func (s *CommonAsset) SetLabels(v map[string]interface{}) *CommonAsset {
	s.Labels = v
	return s
}

func (s *CommonAsset) SetLocation(v *TypeLatLng) *CommonAsset {
	s.Location = v
	return s
}

func (s *CommonAsset) SetSource(v string) *CommonAsset {
	s.Source = &v
	return s
}

func (s *CommonAsset) SetStatus(v string) *CommonAsset {
	s.Status = &v
	return s
}

func (s *CommonAsset) SetSynopsis(v string) *CommonAsset {
	s.Synopsis = &v
	return s
}

func (s *CommonAsset) SetTags(v []*string) *CommonAsset {
	s.Tags = v
	return s
}

func (s *CommonAsset) SetTitle(v string) *CommonAsset {
	s.Title = &v
	return s
}

func (s *CommonAsset) SetUpdatedAt(v string) *CommonAsset {
	s.UpdatedAt = &v
	return s
}

func (s *CommonAsset) SetVideos(v []*CommonMediaResource) *CommonAsset {
	s.Videos = v
	return s
}

type CommonCustomer struct {
	// 云帐号Id
	CloudUid *string `json:"CloudUid,omitempty" xml:"CloudUid,omitempty"`
}

func (s CommonCustomer) String() string {
	return tea.Prettify(s)
}

func (s CommonCustomer) GoString() string {
	return s.String()
}

func (s *CommonCustomer) SetCloudUid(v string) *CommonCustomer {
	s.CloudUid = &v
	return s
}

type CommonMediaResource struct {
	// media format
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// media id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// media sha1
	Sha1 *string `json:"Sha1,omitempty" xml:"Sha1,omitempty"`
	// media size
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// resource url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CommonMediaResource) String() string {
	return tea.Prettify(s)
}

func (s CommonMediaResource) GoString() string {
	return s.String()
}

func (s *CommonMediaResource) SetFormat(v string) *CommonMediaResource {
	s.Format = &v
	return s
}

func (s *CommonMediaResource) SetId(v string) *CommonMediaResource {
	s.Id = &v
	return s
}

func (s *CommonMediaResource) SetName(v string) *CommonMediaResource {
	s.Name = &v
	return s
}

func (s *CommonMediaResource) SetSha1(v string) *CommonMediaResource {
	s.Sha1 = &v
	return s
}

func (s *CommonMediaResource) SetSize(v int64) *CommonMediaResource {
	s.Size = &v
	return s
}

func (s *CommonMediaResource) SetUrl(v string) *CommonMediaResource {
	s.Url = &v
	return s
}

type CommonSTS struct {
	// AccessKey ID 标识用户
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// AccessKey Secret 验证用户的密钥
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// 临时token
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CommonSTS) String() string {
	return tea.Prettify(s)
}

func (s CommonSTS) GoString() string {
	return s.String()
}

func (s *CommonSTS) SetAccessKeyId(v string) *CommonSTS {
	s.AccessKeyId = &v
	return s
}

func (s *CommonSTS) SetAccessKeySecret(v string) *CommonSTS {
	s.AccessKeySecret = &v
	return s
}

func (s *CommonSTS) SetSecurityToken(v string) *CommonSTS {
	s.SecurityToken = &v
	return s
}

type CommonSimpleAsset struct {
	// 行政区域地址
	Address *CommonAddress `json:"Address,omitempty" xml:"Address,omitempty"`
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 审核状态
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// 资产描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 扩展字段
	Extends map[string]interface{} `json:"Extends,omitempty" xml:"Extends,omitempty"`
	// 资产ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 图像资源
	Image *CommonMediaResource `json:"Image,omitempty" xml:"Image,omitempty"`
	// 定义Label， eg:type:advertise 支持广告类型的label
	Labels map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 经纬度地理位置
	Location *TypeLatLng `json:"Location,omitempty" xml:"Location,omitempty"`
	// 资产来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 资产状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 概要
	Synopsis *string `json:"Synopsis,omitempty" xml:"Synopsis,omitempty"`
	// 标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 视频资源
	Video *CommonMediaResource `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CommonSimpleAsset) String() string {
	return tea.Prettify(s)
}

func (s CommonSimpleAsset) GoString() string {
	return s.String()
}

func (s *CommonSimpleAsset) SetAddress(v *CommonAddress) *CommonSimpleAsset {
	s.Address = v
	return s
}

func (s *CommonSimpleAsset) SetAppId(v string) *CommonSimpleAsset {
	s.AppId = &v
	return s
}

func (s *CommonSimpleAsset) SetAuditStatus(v string) *CommonSimpleAsset {
	s.AuditStatus = &v
	return s
}

func (s *CommonSimpleAsset) SetAuthor(v string) *CommonSimpleAsset {
	s.Author = &v
	return s
}

func (s *CommonSimpleAsset) SetDescription(v string) *CommonSimpleAsset {
	s.Description = &v
	return s
}

func (s *CommonSimpleAsset) SetExtends(v map[string]interface{}) *CommonSimpleAsset {
	s.Extends = v
	return s
}

func (s *CommonSimpleAsset) SetId(v string) *CommonSimpleAsset {
	s.Id = &v
	return s
}

func (s *CommonSimpleAsset) SetImage(v *CommonMediaResource) *CommonSimpleAsset {
	s.Image = v
	return s
}

func (s *CommonSimpleAsset) SetLabels(v map[string]interface{}) *CommonSimpleAsset {
	s.Labels = v
	return s
}

func (s *CommonSimpleAsset) SetLocation(v *TypeLatLng) *CommonSimpleAsset {
	s.Location = v
	return s
}

func (s *CommonSimpleAsset) SetSource(v string) *CommonSimpleAsset {
	s.Source = &v
	return s
}

func (s *CommonSimpleAsset) SetStatus(v string) *CommonSimpleAsset {
	s.Status = &v
	return s
}

func (s *CommonSimpleAsset) SetSynopsis(v string) *CommonSimpleAsset {
	s.Synopsis = &v
	return s
}

func (s *CommonSimpleAsset) SetTags(v []*string) *CommonSimpleAsset {
	s.Tags = v
	return s
}

func (s *CommonSimpleAsset) SetTitle(v string) *CommonSimpleAsset {
	s.Title = &v
	return s
}

func (s *CommonSimpleAsset) SetVideo(v *CommonMediaResource) *CommonSimpleAsset {
	s.Video = v
	return s
}

type CommonStorage struct {
	// OSS bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// OSS object path
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// 阿里云Region，比如 cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// OSS的签名URL
	SignUrl *string `json:"SignUrl,omitempty" xml:"SignUrl,omitempty"`
	// Security Token Service
	Sts *CommonSTS `json:"Sts,omitempty" xml:"Sts,omitempty"`
}

func (s CommonStorage) String() string {
	return tea.Prettify(s)
}

func (s CommonStorage) GoString() string {
	return s.String()
}

func (s *CommonStorage) SetBucket(v string) *CommonStorage {
	s.Bucket = &v
	return s
}

func (s *CommonStorage) SetObject(v string) *CommonStorage {
	s.Object = &v
	return s
}

func (s *CommonStorage) SetRegion(v string) *CommonStorage {
	s.Region = &v
	return s
}

func (s *CommonStorage) SetSignUrl(v string) *CommonStorage {
	s.SignUrl = &v
	return s
}

func (s *CommonStorage) SetSts(v *CommonSTS) *CommonStorage {
	s.Sts = v
	return s
}

type RpcStatus struct {
	// 错误码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误详情
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RpcStatus) String() string {
	return tea.Prettify(s)
}

func (s RpcStatus) GoString() string {
	return s.String()
}

func (s *RpcStatus) SetCode(v int32) *RpcStatus {
	s.Code = &v
	return s
}

func (s *RpcStatus) SetDetail(v string) *RpcStatus {
	s.Detail = &v
	return s
}

func (s *RpcStatus) SetMessage(v string) *RpcStatus {
	s.Message = &v
	return s
}

type TypeLatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s TypeLatLng) String() string {
	return tea.Prettify(s)
}

func (s TypeLatLng) GoString() string {
	return s.String()
}

func (s *TypeLatLng) SetLatitude(v float64) *TypeLatLng {
	s.Latitude = &v
	return s
}

func (s *TypeLatLng) SetLongitude(v float64) *TypeLatLng {
	s.Longitude = &v
	return s
}

type V1MediaSecurityStorageRequest struct {
	AppId        *string              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSettings  *AppsSettings        `json:"AppSettings,omitempty" xml:"AppSettings,omitempty"`
	AssetId      *string              `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	Customer     *CommonCustomer      `json:"Customer,omitempty" xml:"Customer,omitempty"`
	Image        *CommonMediaResource `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageStorage *CommonStorage       `json:"ImageStorage,omitempty" xml:"ImageStorage,omitempty"`
	Video        *CommonMediaResource `json:"Video,omitempty" xml:"Video,omitempty"`
	VideoStorage *CommonStorage       `json:"VideoStorage,omitempty" xml:"VideoStorage,omitempty"`
}

func (s V1MediaSecurityStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s V1MediaSecurityStorageRequest) GoString() string {
	return s.String()
}

func (s *V1MediaSecurityStorageRequest) SetAppId(v string) *V1MediaSecurityStorageRequest {
	s.AppId = &v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetAppSettings(v *AppsSettings) *V1MediaSecurityStorageRequest {
	s.AppSettings = v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetAssetId(v string) *V1MediaSecurityStorageRequest {
	s.AssetId = &v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetCustomer(v *CommonCustomer) *V1MediaSecurityStorageRequest {
	s.Customer = v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetImage(v *CommonMediaResource) *V1MediaSecurityStorageRequest {
	s.Image = v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetImageStorage(v *CommonStorage) *V1MediaSecurityStorageRequest {
	s.ImageStorage = v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetVideo(v *CommonMediaResource) *V1MediaSecurityStorageRequest {
	s.Video = v
	return s
}

func (s *V1MediaSecurityStorageRequest) SetVideoStorage(v *CommonStorage) *V1MediaSecurityStorageRequest {
	s.VideoStorage = v
	return s
}

type V1MediaSecurityStorageResponse struct {
	AssetId      *string              `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	Image        *CommonMediaResource `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageStorage *CommonStorage       `json:"ImageStorage,omitempty" xml:"ImageStorage,omitempty"`
	MediaId      *string              `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Video        *CommonMediaResource `json:"Video,omitempty" xml:"Video,omitempty"`
	VideoStorage *CommonStorage       `json:"VideoStorage,omitempty" xml:"VideoStorage,omitempty"`
}

func (s V1MediaSecurityStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s V1MediaSecurityStorageResponse) GoString() string {
	return s.String()
}

func (s *V1MediaSecurityStorageResponse) SetAssetId(v string) *V1MediaSecurityStorageResponse {
	s.AssetId = &v
	return s
}

func (s *V1MediaSecurityStorageResponse) SetImage(v *CommonMediaResource) *V1MediaSecurityStorageResponse {
	s.Image = v
	return s
}

func (s *V1MediaSecurityStorageResponse) SetImageStorage(v *CommonStorage) *V1MediaSecurityStorageResponse {
	s.ImageStorage = v
	return s
}

func (s *V1MediaSecurityStorageResponse) SetMediaId(v string) *V1MediaSecurityStorageResponse {
	s.MediaId = &v
	return s
}

func (s *V1MediaSecurityStorageResponse) SetVideo(v *CommonMediaResource) *V1MediaSecurityStorageResponse {
	s.Video = v
	return s
}

func (s *V1MediaSecurityStorageResponse) SetVideoStorage(v *CommonStorage) *V1MediaSecurityStorageResponse {
	s.VideoStorage = v
	return s
}

type BanAllCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BanAllCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s BanAllCommentRequest) GoString() string {
	return s.String()
}

func (s *BanAllCommentRequest) SetAppId(v string) *BanAllCommentRequest {
	s.AppId = &v
	return s
}

func (s *BanAllCommentRequest) SetRoomId(v string) *BanAllCommentRequest {
	s.RoomId = &v
	return s
}

func (s *BanAllCommentRequest) SetUserId(v string) *BanAllCommentRequest {
	s.UserId = &v
	return s
}

type BanAllCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作成功标识
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BanAllCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BanAllCommentResponseBody) GoString() string {
	return s.String()
}

func (s *BanAllCommentResponseBody) SetRequestId(v string) *BanAllCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BanAllCommentResponseBody) SetResult(v bool) *BanAllCommentResponseBody {
	s.Result = &v
	return s
}

type BanAllCommentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BanAllCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BanAllCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s BanAllCommentResponse) GoString() string {
	return s.String()
}

func (s *BanAllCommentResponse) SetHeaders(v map[string]*string) *BanAllCommentResponse {
	s.Headers = v
	return s
}

func (s *BanAllCommentResponse) SetStatusCode(v int32) *BanAllCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *BanAllCommentResponse) SetBody(v *BanAllCommentResponseBody) *BanAllCommentResponse {
	s.Body = v
	return s
}

type BanCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 禁言时长（秒）
	BanCommentTime *int64 `json:"BanCommentTime,omitempty" xml:"BanCommentTime,omitempty"`
	// 被禁言的用户在房间内的唯一标识
	BanCommentUser *string `json:"BanCommentUser,omitempty" xml:"BanCommentUser,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BanCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s BanCommentRequest) GoString() string {
	return s.String()
}

func (s *BanCommentRequest) SetAppId(v string) *BanCommentRequest {
	s.AppId = &v
	return s
}

func (s *BanCommentRequest) SetBanCommentTime(v int64) *BanCommentRequest {
	s.BanCommentTime = &v
	return s
}

func (s *BanCommentRequest) SetBanCommentUser(v string) *BanCommentRequest {
	s.BanCommentUser = &v
	return s
}

func (s *BanCommentRequest) SetRoomId(v string) *BanCommentRequest {
	s.RoomId = &v
	return s
}

func (s *BanCommentRequest) SetUserId(v string) *BanCommentRequest {
	s.UserId = &v
	return s
}

type BanCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作是否成功
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BanCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BanCommentResponseBody) GoString() string {
	return s.String()
}

func (s *BanCommentResponseBody) SetRequestId(v string) *BanCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *BanCommentResponseBody) SetResult(v bool) *BanCommentResponseBody {
	s.Result = &v
	return s
}

type BanCommentResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BanCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BanCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s BanCommentResponse) GoString() string {
	return s.String()
}

func (s *BanCommentResponse) SetHeaders(v map[string]*string) *BanCommentResponse {
	s.Headers = v
	return s
}

func (s *BanCommentResponse) SetStatusCode(v int32) *BanCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *BanCommentResponse) SetBody(v *BanCommentResponseBody) *BanCommentResponse {
	s.Body = v
	return s
}

type CancelBanAllCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelBanAllCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelBanAllCommentRequest) GoString() string {
	return s.String()
}

func (s *CancelBanAllCommentRequest) SetAppId(v string) *CancelBanAllCommentRequest {
	s.AppId = &v
	return s
}

func (s *CancelBanAllCommentRequest) SetRoomId(v string) *CancelBanAllCommentRequest {
	s.RoomId = &v
	return s
}

func (s *CancelBanAllCommentRequest) SetUserId(v string) *CancelBanAllCommentRequest {
	s.UserId = &v
	return s
}

type CancelBanAllCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作成功标识
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelBanAllCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelBanAllCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBanAllCommentResponseBody) SetRequestId(v string) *CancelBanAllCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBanAllCommentResponseBody) SetResult(v bool) *CancelBanAllCommentResponseBody {
	s.Result = &v
	return s
}

type CancelBanAllCommentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelBanAllCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelBanAllCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelBanAllCommentResponse) GoString() string {
	return s.String()
}

func (s *CancelBanAllCommentResponse) SetHeaders(v map[string]*string) *CancelBanAllCommentResponse {
	s.Headers = v
	return s
}

func (s *CancelBanAllCommentResponse) SetStatusCode(v int32) *CancelBanAllCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelBanAllCommentResponse) SetBody(v *CancelBanAllCommentResponseBody) *CancelBanAllCommentResponse {
	s.Body = v
	return s
}

type CancelBanCommentRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 取消禁言的用户唯一标识
	BanCommentUser *string `json:"BanCommentUser,omitempty" xml:"BanCommentUser,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户在房间内的唯一标识
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelBanCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelBanCommentRequest) GoString() string {
	return s.String()
}

func (s *CancelBanCommentRequest) SetAppId(v string) *CancelBanCommentRequest {
	s.AppId = &v
	return s
}

func (s *CancelBanCommentRequest) SetBanCommentUser(v string) *CancelBanCommentRequest {
	s.BanCommentUser = &v
	return s
}

func (s *CancelBanCommentRequest) SetRoomId(v string) *CancelBanCommentRequest {
	s.RoomId = &v
	return s
}

func (s *CancelBanCommentRequest) SetUserId(v string) *CancelBanCommentRequest {
	s.UserId = &v
	return s
}

type CancelBanCommentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作成功标识
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CancelBanCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelBanCommentResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBanCommentResponseBody) SetRequestId(v string) *CancelBanCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBanCommentResponseBody) SetResult(v bool) *CancelBanCommentResponseBody {
	s.Result = &v
	return s
}

type CancelBanCommentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelBanCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelBanCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelBanCommentResponse) GoString() string {
	return s.String()
}

func (s *CancelBanCommentResponse) SetHeaders(v map[string]*string) *CancelBanCommentResponse {
	s.Headers = v
	return s
}

func (s *CancelBanCommentResponse) SetStatusCode(v int32) *CancelBanCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelBanCommentResponse) SetBody(v *CancelBanCommentResponseBody) *CancelBanCommentResponse {
	s.Body = v
	return s
}

type CancelUserAdminRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelUserAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelUserAdminRequest) GoString() string {
	return s.String()
}

func (s *CancelUserAdminRequest) SetAppId(v string) *CancelUserAdminRequest {
	s.AppId = &v
	return s
}

func (s *CancelUserAdminRequest) SetRoomId(v string) *CancelUserAdminRequest {
	s.RoomId = &v
	return s
}

func (s *CancelUserAdminRequest) SetUserId(v string) *CancelUserAdminRequest {
	s.UserId = &v
	return s
}

type CancelUserAdminResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUserAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelUserAdminResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUserAdminResponseBody) SetRequestId(v string) *CancelUserAdminResponseBody {
	s.RequestId = &v
	return s
}

type CancelUserAdminResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelUserAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelUserAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelUserAdminResponse) GoString() string {
	return s.String()
}

func (s *CancelUserAdminResponse) SetHeaders(v map[string]*string) *CancelUserAdminResponse {
	s.Headers = v
	return s
}

func (s *CancelUserAdminResponse) SetStatusCode(v int32) *CancelUserAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUserAdminResponse) SetBody(v *CancelUserAdminResponseBody) *CancelUserAdminResponse {
	s.Body = v
	return s
}

type CreateClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建人用户昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人用户ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 课堂标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateClassRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClassRequest) GoString() string {
	return s.String()
}

func (s *CreateClassRequest) SetAppId(v string) *CreateClassRequest {
	s.AppId = &v
	return s
}

func (s *CreateClassRequest) SetCreateNickname(v string) *CreateClassRequest {
	s.CreateNickname = &v
	return s
}

func (s *CreateClassRequest) SetCreateUserId(v string) *CreateClassRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateClassRequest) SetTitle(v string) *CreateClassRequest {
	s.Title = &v
	return s
}

type CreateClassResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *CreateClassResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClassResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClassResponseBody) SetRequestId(v string) *CreateClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClassResponseBody) SetResult(v *CreateClassResponseBodyResult) *CreateClassResponseBody {
	s.Result = v
	return s
}

type CreateClassResponseBodyResult struct {
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 创建人昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 课堂状态，0:未开始 1:上课中 2:已下课。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 课堂标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 白板ID
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
	// 白板录制ID
	WhiteboardRecordId *string `json:"WhiteboardRecordId,omitempty" xml:"WhiteboardRecordId,omitempty"`
}

func (s CreateClassResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateClassResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateClassResponseBodyResult) SetClassId(v string) *CreateClassResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetConfId(v string) *CreateClassResponseBodyResult {
	s.ConfId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetCreateNickname(v string) *CreateClassResponseBodyResult {
	s.CreateNickname = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetCreateUserId(v string) *CreateClassResponseBodyResult {
	s.CreateUserId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetLiveId(v string) *CreateClassResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetRoomId(v string) *CreateClassResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetStatus(v int32) *CreateClassResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetTitle(v string) *CreateClassResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetWhiteboardId(v string) *CreateClassResponseBodyResult {
	s.WhiteboardId = &v
	return s
}

func (s *CreateClassResponseBodyResult) SetWhiteboardRecordId(v string) *CreateClassResponseBodyResult {
	s.WhiteboardRecordId = &v
	return s
}

type CreateClassResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClassResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClassResponse) GoString() string {
	return s.String()
}

func (s *CreateClassResponse) SetHeaders(v map[string]*string) *CreateClassResponse {
	s.Headers = v
	return s
}

func (s *CreateClassResponse) SetStatusCode(v int32) *CreateClassResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClassResponse) SetBody(v *CreateClassResponseBody) *CreateClassResponse {
	s.Body = v
	return s
}

type CreateLiveRequest struct {
	// 主播ID，支持中英文，最大长度128位，缺省时主播为当前创建直播用户。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播推流码率，缺省时默认为3。取值：  -1：流畅lld。 1：标清lsd。 2：高清lhd。 3：超清lud。
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// 直播简介，支持中英文，最大长度2048位。
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 直播资源的唯一标识ID，缺省时系统自动生成36位随机uuid字符串。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID，最大长度36个字符，传空值，则随机生成一个房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播标题，支持中英文，最大长度256位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建直播用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) SetAnchorId(v string) *CreateLiveRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRequest) SetAppId(v string) *CreateLiveRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRequest) SetCodeLevel(v int32) *CreateLiveRequest {
	s.CodeLevel = &v
	return s
}

func (s *CreateLiveRequest) SetIntroduction(v string) *CreateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *CreateLiveRequest) SetLiveId(v string) *CreateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *CreateLiveRequest) SetRoomId(v string) *CreateLiveRequest {
	s.RoomId = &v
	return s
}

func (s *CreateLiveRequest) SetTitle(v string) *CreateLiveRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRequest) SetUserId(v string) *CreateLiveRequest {
	s.UserId = &v
	return s
}

type CreateLiveResponseBody struct {
	// 请求ID。
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateLiveResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBody) SetRequestId(v string) *CreateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveResponseBody) SetResult(v *CreateLiveResponseBodyResult) *CreateLiveResponseBody {
	s.Result = v
	return s
}

type CreateLiveResponseBodyResult struct {
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s CreateLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBodyResult) SetLiveId(v string) *CreateLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

type CreateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveResponse) SetHeaders(v map[string]*string) *CreateLiveResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveResponse) SetStatusCode(v int32) *CreateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveResponse) SetBody(v *CreateLiveResponseBody) *CreateLiveResponse {
	s.Body = v
	return s
}

type CreateLiveRecordSliceFileRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 片段结束时间，时间戳。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 自定义文件名称。
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 片段开始时间，时间戳。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateLiveRecordSliceFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRecordSliceFileRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordSliceFileRequest) SetAppId(v string) *CreateLiveRecordSliceFileRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRecordSliceFileRequest) SetEndTime(v int64) *CreateLiveRecordSliceFileRequest {
	s.EndTime = &v
	return s
}

func (s *CreateLiveRecordSliceFileRequest) SetFileName(v string) *CreateLiveRecordSliceFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateLiveRecordSliceFileRequest) SetLiveId(v string) *CreateLiveRecordSliceFileRequest {
	s.LiveId = &v
	return s
}

func (s *CreateLiveRecordSliceFileRequest) SetStartTime(v int64) *CreateLiveRecordSliceFileRequest {
	s.StartTime = &v
	return s
}

type CreateLiveRecordSliceFileResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *CreateLiveRecordSliceFileResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateLiveRecordSliceFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRecordSliceFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordSliceFileResponseBody) SetRequestId(v string) *CreateLiveRecordSliceFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveRecordSliceFileResponseBody) SetResult(v *CreateLiveRecordSliceFileResponseBodyResult) *CreateLiveRecordSliceFileResponseBody {
	s.Result = v
	return s
}

type CreateLiveRecordSliceFileResponseBodyResult struct {
	// 片段录制文件的地址。
	SliceRecordUrl *string `json:"SliceRecordUrl,omitempty" xml:"SliceRecordUrl,omitempty"`
}

func (s CreateLiveRecordSliceFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRecordSliceFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordSliceFileResponseBodyResult) SetSliceRecordUrl(v string) *CreateLiveRecordSliceFileResponseBodyResult {
	s.SliceRecordUrl = &v
	return s
}

type CreateLiveRecordSliceFileResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLiveRecordSliceFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveRecordSliceFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRecordSliceFileResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordSliceFileResponse) SetHeaders(v map[string]*string) *CreateLiveRecordSliceFileResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveRecordSliceFileResponse) SetStatusCode(v int32) *CreateLiveRecordSliceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveRecordSliceFileResponse) SetBody(v *CreateLiveRecordSliceFileResponseBody) *CreateLiveRecordSliceFileResponse {
	s.Body = v
	return s
}

type CreateLiveRoomRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 是否开启连麦。
	EnableLinkMic *bool `json:"EnableLinkMic,omitempty" xml:"EnableLinkMic,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。最大支持4096个字节。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomRequest) SetAnchorId(v string) *CreateLiveRoomRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRoomRequest) SetAnchorNick(v string) *CreateLiveRoomRequest {
	s.AnchorNick = &v
	return s
}

func (s *CreateLiveRoomRequest) SetAppId(v string) *CreateLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRoomRequest) SetCoverUrl(v string) *CreateLiveRoomRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRoomRequest) SetEnableLinkMic(v bool) *CreateLiveRoomRequest {
	s.EnableLinkMic = &v
	return s
}

func (s *CreateLiveRoomRequest) SetExtension(v map[string]*string) *CreateLiveRoomRequest {
	s.Extension = v
	return s
}

func (s *CreateLiveRoomRequest) SetNotice(v string) *CreateLiveRoomRequest {
	s.Notice = &v
	return s
}

func (s *CreateLiveRoomRequest) SetTitle(v string) *CreateLiveRoomRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRoomRequest) SetUserId(v string) *CreateLiveRoomRequest {
	s.UserId = &v
	return s
}

type CreateLiveRoomShrinkRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 是否开启连麦。
	EnableLinkMic *bool `json:"EnableLinkMic,omitempty" xml:"EnableLinkMic,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。最大支持4096个字节。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateLiveRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomShrinkRequest) SetAnchorId(v string) *CreateLiveRoomShrinkRequest {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetAnchorNick(v string) *CreateLiveRoomShrinkRequest {
	s.AnchorNick = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetAppId(v string) *CreateLiveRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetCoverUrl(v string) *CreateLiveRoomShrinkRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetEnableLinkMic(v bool) *CreateLiveRoomShrinkRequest {
	s.EnableLinkMic = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetExtensionShrink(v string) *CreateLiveRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetNotice(v string) *CreateLiveRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetTitle(v string) *CreateLiveRoomShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRoomShrinkRequest) SetUserId(v string) *CreateLiveRoomShrinkRequest {
	s.UserId = &v
	return s
}

type CreateLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *CreateLiveRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBody) SetRequestId(v string) *CreateLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveRoomResponseBody) SetResult(v *CreateLiveRoomResponseBodyResult) *CreateLiveRoomResponseBody {
	s.Result = v
	return s
}

type CreateLiveRoomResponseBodyResult struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// RTS低延迟播流信息。
	ArtcInfo *CreateLiveRoomResponseBodyResultArtcInfo `json:"ArtcInfo,omitempty" xml:"ArtcInfo,omitempty" type:"Struct"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 原画HLS播放地址。
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址。
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 直播回放地址。
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*CreateLiveRoomResponseBodyResultPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 直播推流地址。
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateLiveRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBodyResult) SetAnchorId(v string) *CreateLiveRoomResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetAnchorNick(v string) *CreateLiveRoomResponseBodyResult {
	s.AnchorNick = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetAppId(v string) *CreateLiveRoomResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetArtcInfo(v *CreateLiveRoomResponseBodyResultArtcInfo) *CreateLiveRoomResponseBodyResult {
	s.ArtcInfo = v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetChatId(v string) *CreateLiveRoomResponseBodyResult {
	s.ChatId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetCoverUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetExtension(v map[string]*string) *CreateLiveRoomResponseBodyResult {
	s.Extension = v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetHlsUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.HlsUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetLiveId(v string) *CreateLiveRoomResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetLiveUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetNotice(v string) *CreateLiveRoomResponseBodyResult {
	s.Notice = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetPlaybackUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetPluginInstanceInfoList(v []*CreateLiveRoomResponseBodyResultPluginInstanceInfoList) *CreateLiveRoomResponseBodyResult {
	s.PluginInstanceInfoList = v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetPushUrl(v string) *CreateLiveRoomResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetRoomId(v string) *CreateLiveRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResult) SetTitle(v string) *CreateLiveRoomResponseBodyResult {
	s.Title = &v
	return s
}

type CreateLiveRoomResponseBodyResultArtcInfo struct {
	// RTS转码流地址，推荐web端使用。
	ArtcH5Url *string `json:"ArtcH5Url,omitempty" xml:"ArtcH5Url,omitempty"`
	// RTS原码流地址，推荐移动端使用。
	ArtcUrl *string `json:"ArtcUrl,omitempty" xml:"ArtcUrl,omitempty"`
}

func (s CreateLiveRoomResponseBodyResultArtcInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBodyResultArtcInfo) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBodyResultArtcInfo) SetArtcH5Url(v string) *CreateLiveRoomResponseBodyResultArtcInfo {
	s.ArtcH5Url = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResultArtcInfo) SetArtcUrl(v string) *CreateLiveRoomResponseBodyResultArtcInfo {
	s.ArtcUrl = &v
	return s
}

type CreateLiveRoomResponseBodyResultPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s CreateLiveRoomResponseBodyResultPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponseBodyResultPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetCreateTime(v int64) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetExtension(v map[string]*string) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginId(v string) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *CreateLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginType(v string) *CreateLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

type CreateLiveRoomResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveRoomResponse) SetHeaders(v map[string]*string) *CreateLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveRoomResponse) SetStatusCode(v int32) *CreateLiveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveRoomResponse) SetBody(v *CreateLiveRoomResponseBody) *CreateLiveRoomResponse {
	s.Body = v
	return s
}

type CreateRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间模板唯一标识，当前支持的取值：default，传空默认为default。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomRequest) SetAppId(v string) *CreateRoomRequest {
	s.AppId = &v
	return s
}

func (s *CreateRoomRequest) SetExtension(v map[string]*string) *CreateRoomRequest {
	s.Extension = v
	return s
}

func (s *CreateRoomRequest) SetNotice(v string) *CreateRoomRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomRequest) SetRoomId(v string) *CreateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomRequest) SetRoomOwnerId(v string) *CreateRoomRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *CreateRoomRequest) SetTemplateId(v string) *CreateRoomRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomRequest) SetTitle(v string) *CreateRoomRequest {
	s.Title = &v
	return s
}

type CreateRoomShrinkRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间模板唯一标识，当前支持的取值：default，传空默认为default。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoomShrinkRequest) SetAppId(v string) *CreateRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetExtensionShrink(v string) *CreateRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetNotice(v string) *CreateRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetRoomId(v string) *CreateRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetRoomOwnerId(v string) *CreateRoomShrinkRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetTemplateId(v string) *CreateRoomShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRoomShrinkRequest) SetTitle(v string) *CreateRoomShrinkRequest {
	s.Title = &v
	return s
}

type CreateRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *CreateRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoomResponseBody) SetRequestId(v string) *CreateRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoomResponseBody) SetResult(v *CreateRoomResponseBodyResult) *CreateRoomResponseBody {
	s.Result = v
	return s
}

type CreateRoomResponseBodyResult struct {
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s CreateRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRoomResponseBodyResult) SetRoomId(v string) *CreateRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

type CreateRoomResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoomResponse) GoString() string {
	return s.String()
}

func (s *CreateRoomResponse) SetHeaders(v map[string]*string) *CreateRoomResponse {
	s.Headers = v
	return s
}

func (s *CreateRoomResponse) SetStatusCode(v int32) *CreateRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoomResponse) SetBody(v *CreateRoomResponseBody) *CreateRoomResponse {
	s.Body = v
	return s
}

type CreateSensitiveWordRequest struct {
	// 用户的应用ID，在控制台创建应用时生成。包含小写字母、数字，长度为6个字符。
	AppId    *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordList []*string `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s CreateSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordRequest) SetAppId(v string) *CreateSensitiveWordRequest {
	s.AppId = &v
	return s
}

func (s *CreateSensitiveWordRequest) SetWordList(v []*string) *CreateSensitiveWordRequest {
	s.WordList = v
	return s
}

type CreateSensitiveWordShrinkRequest struct {
	// 用户的应用ID，在控制台创建应用时生成。包含小写字母、数字，长度为6个字符。
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordListShrink *string `json:"WordList,omitempty" xml:"WordList,omitempty"`
}

func (s CreateSensitiveWordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordShrinkRequest) SetAppId(v string) *CreateSensitiveWordShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateSensitiveWordShrinkRequest) SetWordListShrink(v string) *CreateSensitiveWordShrinkRequest {
	s.WordListShrink = &v
	return s
}

type CreateSensitiveWordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *CreateSensitiveWordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordResponseBody) SetRequestId(v string) *CreateSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSensitiveWordResponseBody) SetResult(v *CreateSensitiveWordResponseBodyResult) *CreateSensitiveWordResponseBody {
	s.Result = v
	return s
}

type CreateSensitiveWordResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSensitiveWordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordResponseBodyResult) SetSuccess(v bool) *CreateSensitiveWordResponseBodyResult {
	s.Success = &v
	return s
}

type CreateSensitiveWordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *CreateSensitiveWordResponse) SetHeaders(v map[string]*string) *CreateSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *CreateSensitiveWordResponse) SetStatusCode(v int32) *CreateSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSensitiveWordResponse) SetBody(v *CreateSensitiveWordResponseBody) *CreateSensitiveWordResponse {
	s.Body = v
	return s
}

type DeleteClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作人用户ID，仅支持中英文数字，下划线，中划线，1~36个字符。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteClassRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClassRequest) GoString() string {
	return s.String()
}

func (s *DeleteClassRequest) SetAppId(v string) *DeleteClassRequest {
	s.AppId = &v
	return s
}

func (s *DeleteClassRequest) SetClassId(v string) *DeleteClassRequest {
	s.ClassId = &v
	return s
}

func (s *DeleteClassRequest) SetUserId(v string) *DeleteClassRequest {
	s.UserId = &v
	return s
}

type DeleteClassResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClassResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClassResponseBody) SetRequestId(v string) *DeleteClassResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClassResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClassResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClassResponse) GoString() string {
	return s.String()
}

func (s *DeleteClassResponse) SetHeaders(v map[string]*string) *DeleteClassResponse {
	s.Headers = v
	return s
}

func (s *DeleteClassResponse) SetStatusCode(v int32) *DeleteClassResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClassResponse) SetBody(v *DeleteClassResponseBody) *DeleteClassResponse {
	s.Body = v
	return s
}

type DeleteCommentRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 需要删除的弹幕id列表
	CommentIdList []*string `json:"CommentIdList,omitempty" xml:"CommentIdList,omitempty" type:"Repeated"`
	// 弹幕的创建者ID。
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 删除的操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommentRequest) SetAppId(v string) *DeleteCommentRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCommentRequest) SetCommentIdList(v []*string) *DeleteCommentRequest {
	s.CommentIdList = v
	return s
}

func (s *DeleteCommentRequest) SetCreatorId(v string) *DeleteCommentRequest {
	s.CreatorId = &v
	return s
}

func (s *DeleteCommentRequest) SetRoomId(v string) *DeleteCommentRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteCommentRequest) SetUserId(v string) *DeleteCommentRequest {
	s.UserId = &v
	return s
}

type DeleteCommentResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用删除直播间弹幕的返回结果。
	Result *DeleteCommentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponseBody) SetRequestId(v string) *DeleteCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommentResponseBody) SetResult(v *DeleteCommentResponseBodyResult) *DeleteCommentResponseBody {
	s.Result = v
	return s
}

type DeleteCommentResponseBodyResult struct {
	// 删除的结果
	DeleteResult *bool `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
}

func (s DeleteCommentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponseBodyResult) SetDeleteResult(v bool) *DeleteCommentResponseBodyResult {
	s.DeleteResult = &v
	return s
}

type DeleteCommentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponse) SetHeaders(v map[string]*string) *DeleteCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommentResponse) SetStatusCode(v int32) *DeleteCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommentResponse) SetBody(v *DeleteCommentResponseBody) *DeleteCommentResponse {
	s.Body = v
	return s
}

type DeleteCommentByCreatorIdRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 需要删除的弹幕id列表
	CommentIdList []*string `json:"CommentIdList,omitempty" xml:"CommentIdList,omitempty" type:"Repeated"`
	// 弹幕的创建者ID。
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 删除的操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteCommentByCreatorIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentByCreatorIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommentByCreatorIdRequest) SetAppId(v string) *DeleteCommentByCreatorIdRequest {
	s.AppId = &v
	return s
}

func (s *DeleteCommentByCreatorIdRequest) SetCommentIdList(v []*string) *DeleteCommentByCreatorIdRequest {
	s.CommentIdList = v
	return s
}

func (s *DeleteCommentByCreatorIdRequest) SetCreatorId(v string) *DeleteCommentByCreatorIdRequest {
	s.CreatorId = &v
	return s
}

func (s *DeleteCommentByCreatorIdRequest) SetRoomId(v string) *DeleteCommentByCreatorIdRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteCommentByCreatorIdRequest) SetUserId(v string) *DeleteCommentByCreatorIdRequest {
	s.UserId = &v
	return s
}

type DeleteCommentByCreatorIdResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用删除直播间弹幕的返回结果。
	Result *DeleteCommentByCreatorIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteCommentByCreatorIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentByCreatorIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommentByCreatorIdResponseBody) SetRequestId(v string) *DeleteCommentByCreatorIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommentByCreatorIdResponseBody) SetResult(v *DeleteCommentByCreatorIdResponseBodyResult) *DeleteCommentByCreatorIdResponseBody {
	s.Result = v
	return s
}

type DeleteCommentByCreatorIdResponseBodyResult struct {
	// 删除的结果
	DeleteResult *bool `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
}

func (s DeleteCommentByCreatorIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentByCreatorIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteCommentByCreatorIdResponseBodyResult) SetDeleteResult(v bool) *DeleteCommentByCreatorIdResponseBodyResult {
	s.DeleteResult = &v
	return s
}

type DeleteCommentByCreatorIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCommentByCreatorIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCommentByCreatorIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentByCreatorIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommentByCreatorIdResponse) SetHeaders(v map[string]*string) *DeleteCommentByCreatorIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommentByCreatorIdResponse) SetStatusCode(v int32) *DeleteCommentByCreatorIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommentByCreatorIdResponse) SetBody(v *DeleteCommentByCreatorIdResponseBody) *DeleteCommentByCreatorIdResponse {
	s.Body = v
	return s
}

type DeleteConferenceRequest struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 会议资源的唯一标识ID
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 房间ID，最大长度36位
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建会议用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceRequest) GoString() string {
	return s.String()
}

func (s *DeleteConferenceRequest) SetAppId(v string) *DeleteConferenceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteConferenceRequest) SetConferenceId(v string) *DeleteConferenceRequest {
	s.ConferenceId = &v
	return s
}

func (s *DeleteConferenceRequest) SetRoomId(v string) *DeleteConferenceRequest {
	s.RoomId = &v
	return s
}

func (s *DeleteConferenceRequest) SetUserId(v string) *DeleteConferenceRequest {
	s.UserId = &v
	return s
}

type DeleteConferenceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConferenceResponseBody) SetRequestId(v string) *DeleteConferenceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConferenceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceResponse) GoString() string {
	return s.String()
}

func (s *DeleteConferenceResponse) SetHeaders(v map[string]*string) *DeleteConferenceResponse {
	s.Headers = v
	return s
}

func (s *DeleteConferenceResponse) SetStatusCode(v int32) *DeleteConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConferenceResponse) SetBody(v *DeleteConferenceResponseBody) *DeleteConferenceResponse {
	s.Body = v
	return s
}

type DeleteLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s DeleteLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRequest) SetLiveId(v string) *DeleteLiveRequest {
	s.LiveId = &v
	return s
}

type DeleteLiveResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponseBody) SetRequestId(v string) *DeleteLiveResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponse) SetHeaders(v map[string]*string) *DeleteLiveResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveResponse) SetStatusCode(v int32) *DeleteLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveResponse) SetBody(v *DeleteLiveResponseBody) *DeleteLiveResponse {
	s.Body = v
	return s
}

type DeleteLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRoomRequest) SetAppId(v string) *DeleteLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *DeleteLiveRoomRequest) SetLiveId(v string) *DeleteLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *DeleteLiveRoomRequest) SetUserId(v string) *DeleteLiveRoomRequest {
	s.UserId = &v
	return s
}

type DeleteLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRoomResponseBody) SetRequestId(v string) *DeleteLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLiveRoomResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRoomResponse) SetHeaders(v map[string]*string) *DeleteLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRoomResponse) SetStatusCode(v int32) *DeleteLiveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRoomResponse) SetBody(v *DeleteLiveRoomResponseBody) *DeleteLiveRoomResponse {
	s.Body = v
	return s
}

type DeleteRecordFileInfoRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DeleteRecordFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordFileInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordFileInfoRequest) SetAppId(v string) *DeleteRecordFileInfoRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRecordFileInfoRequest) SetMaxResults(v int32) *DeleteRecordFileInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *DeleteRecordFileInfoRequest) SetNextToken(v string) *DeleteRecordFileInfoRequest {
	s.NextToken = &v
	return s
}

type DeleteRecordFileInfoResponseBody struct {
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteRecordFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordFileInfoResponseBody) SetMaxResults(v int32) *DeleteRecordFileInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DeleteRecordFileInfoResponseBody) SetNextToken(v string) *DeleteRecordFileInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *DeleteRecordFileInfoResponseBody) SetRequestId(v string) *DeleteRecordFileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordFileInfoResponseBody) SetTotalCount(v int32) *DeleteRecordFileInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DeleteRecordFileInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRecordFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRecordFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordFileInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordFileInfoResponse) SetHeaders(v map[string]*string) *DeleteRecordFileInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordFileInfoResponse) SetStatusCode(v int32) *DeleteRecordFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordFileInfoResponse) SetBody(v *DeleteRecordFileInfoResponseBody) *DeleteRecordFileInfoResponse {
	s.Body = v
	return s
}

type DeleteRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s DeleteRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoomRequest) SetAppId(v string) *DeleteRoomRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRoomRequest) SetRoomId(v string) *DeleteRoomRequest {
	s.RoomId = &v
	return s
}

type DeleteRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoomResponseBody) SetRequestId(v string) *DeleteRoomResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoomResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoomResponse) SetHeaders(v map[string]*string) *DeleteRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoomResponse) SetStatusCode(v int32) *DeleteRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoomResponse) SetBody(v *DeleteRoomResponseBody) *DeleteRoomResponse {
	s.Body = v
	return s
}

type DeleteSensitiveWordRequest struct {
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	AppId    *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordList []*string `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s DeleteSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordRequest) SetAppId(v string) *DeleteSensitiveWordRequest {
	s.AppId = &v
	return s
}

func (s *DeleteSensitiveWordRequest) SetWordList(v []*string) *DeleteSensitiveWordRequest {
	s.WordList = v
	return s
}

type DeleteSensitiveWordShrinkRequest struct {
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	WordListShrink *string `json:"WordList,omitempty" xml:"WordList,omitempty"`
}

func (s DeleteSensitiveWordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordShrinkRequest) SetAppId(v string) *DeleteSensitiveWordShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteSensitiveWordShrinkRequest) SetWordListShrink(v string) *DeleteSensitiveWordShrinkRequest {
	s.WordListShrink = &v
	return s
}

type DeleteSensitiveWordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *DeleteSensitiveWordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordResponseBody) SetRequestId(v string) *DeleteSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSensitiveWordResponseBody) SetResult(v *DeleteSensitiveWordResponseBodyResult) *DeleteSensitiveWordResponseBody {
	s.Result = v
	return s
}

type DeleteSensitiveWordResponseBodyResult struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSensitiveWordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordResponseBodyResult) SetSuccess(v bool) *DeleteSensitiveWordResponseBodyResult {
	s.Success = &v
	return s
}

type DeleteSensitiveWordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *DeleteSensitiveWordResponse) SetHeaders(v map[string]*string) *DeleteSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *DeleteSensitiveWordResponse) SetStatusCode(v int32) *DeleteSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSensitiveWordResponse) SetBody(v *DeleteSensitiveWordResponseBody) *DeleteSensitiveWordResponse {
	s.Body = v
	return s
}

type DescribeMeterImpPlayBackTimeByLiveIdRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EndTs   *int64  `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	LiveId  *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	StartTs *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeMeterImpPlayBackTimeByLiveIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpPlayBackTimeByLiveIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdRequest) SetAppId(v string) *DescribeMeterImpPlayBackTimeByLiveIdRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdRequest) SetEndTs(v int64) *DescribeMeterImpPlayBackTimeByLiveIdRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdRequest) SetLiveId(v string) *DescribeMeterImpPlayBackTimeByLiveIdRequest {
	s.LiveId = &v
	return s
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdRequest) SetStartTs(v int64) *DescribeMeterImpPlayBackTimeByLiveIdRequest {
	s.StartTs = &v
	return s
}

type DescribeMeterImpPlayBackTimeByLiveIdResponseBody struct {
	Data []*DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImpPlayBackTimeByLiveIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpPlayBackTimeByLiveIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdResponseBody) SetData(v []*DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData) *DescribeMeterImpPlayBackTimeByLiveIdResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdResponseBody) SetRequestId(v string) *DescribeMeterImpPlayBackTimeByLiveIdResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData struct {
	WatchTime *int64 `json:"WatchTime,omitempty" xml:"WatchTime,omitempty"`
}

func (s DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData) SetWatchTime(v int64) *DescribeMeterImpPlayBackTimeByLiveIdResponseBodyData {
	s.WatchTime = &v
	return s
}

type DescribeMeterImpPlayBackTimeByLiveIdResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMeterImpPlayBackTimeByLiveIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMeterImpPlayBackTimeByLiveIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpPlayBackTimeByLiveIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdResponse) SetHeaders(v map[string]*string) *DescribeMeterImpPlayBackTimeByLiveIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdResponse) SetStatusCode(v int32) *DescribeMeterImpPlayBackTimeByLiveIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImpPlayBackTimeByLiveIdResponse) SetBody(v *DescribeMeterImpPlayBackTimeByLiveIdResponseBody) *DescribeMeterImpPlayBackTimeByLiveIdResponse {
	s.Body = v
	return s
}

type DescribeMeterImpWatchLiveTimeByLiveIdRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdRequest) SetAppId(v string) *DescribeMeterImpWatchLiveTimeByLiveIdRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdRequest) SetLiveId(v string) *DescribeMeterImpWatchLiveTimeByLiveIdRequest {
	s.LiveId = &v
	return s
}

type DescribeMeterImpWatchLiveTimeByLiveIdResponseBody struct {
	Data []*DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponseBody) SetData(v []*DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData) *DescribeMeterImpWatchLiveTimeByLiveIdResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponseBody) SetRequestId(v string) *DescribeMeterImpWatchLiveTimeByLiveIdResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData struct {
	WatchTimeInLatency    *int64 `json:"WatchTimeInLatency,omitempty" xml:"WatchTimeInLatency,omitempty"`
	WatchTimeInLowLatency *int64 `json:"WatchTimeInLowLatency,omitempty" xml:"WatchTimeInLowLatency,omitempty"`
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData) SetWatchTimeInLatency(v int64) *DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData {
	s.WatchTimeInLatency = &v
	return s
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData) SetWatchTimeInLowLatency(v int64) *DescribeMeterImpWatchLiveTimeByLiveIdResponseBodyData {
	s.WatchTimeInLowLatency = &v
	return s
}

type DescribeMeterImpWatchLiveTimeByLiveIdResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMeterImpWatchLiveTimeByLiveIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMeterImpWatchLiveTimeByLiveIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponse) SetHeaders(v map[string]*string) *DescribeMeterImpWatchLiveTimeByLiveIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponse) SetStatusCode(v int32) *DescribeMeterImpWatchLiveTimeByLiveIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterImpWatchLiveTimeByLiveIdResponse) SetBody(v *DescribeMeterImpWatchLiveTimeByLiveIdResponseBody) *DescribeMeterImpWatchLiveTimeByLiveIdResponse {
	s.Body = v
	return s
}

type GetAuthTokenRequest struct {
	// 用户的应用ID，在控制台创建应用时生成
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 终端设备类型,通过控制台创建和查询
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 终端设备ID
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 用户UserId,在AppId下单独唯一
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) SetAppId(v string) *GetAuthTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetAuthTokenRequest) SetAppKey(v string) *GetAuthTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetAuthTokenRequest) SetDeviceId(v string) *GetAuthTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetAuthTokenRequest) SetUserId(v string) *GetAuthTokenRequest {
	s.UserId = &v
	return s
}

type GetAuthTokenResponseBody struct {
	// Id of the request
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAuthTokenResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetResult(v *GetAuthTokenResponseBodyResult) *GetAuthTokenResponseBody {
	s.Result = v
	return s
}

type GetAuthTokenResponseBodyResult struct {
	// 用于长连接建连的token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 登录token过期时间(毫秒)
	AccessTokenExpiredTime *int64 `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// 更新Token，若AccessToken过期，则可以使用RefreshToken再次获取新Token
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s GetAuthTokenResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyResult) SetAccessToken(v string) *GetAuthTokenResponseBodyResult {
	s.AccessToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetAccessTokenExpiredTime(v int64) *GetAuthTokenResponseBodyResult {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *GetAuthTokenResponseBodyResult) SetRefreshToken(v string) *GetAuthTokenResponseBodyResult {
	s.RefreshToken = &v
	return s
}

type GetAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetStatusCode(v int32) *GetAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

type GetClassDetailRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作人用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetClassDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailRequest) GoString() string {
	return s.String()
}

func (s *GetClassDetailRequest) SetAppId(v string) *GetClassDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetClassDetailRequest) SetClassId(v string) *GetClassDetailRequest {
	s.ClassId = &v
	return s
}

func (s *GetClassDetailRequest) SetUserId(v string) *GetClassDetailRequest {
	s.UserId = &v
	return s
}

type GetClassDetailResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *GetClassDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetClassDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetClassDetailResponseBody) SetRequestId(v string) *GetClassDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClassDetailResponseBody) SetResult(v *GetClassDetailResponseBodyResult) *GetClassDetailResponseBody {
	s.Result = v
	return s
}

type GetClassDetailResponseBodyResult struct {
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 创建人昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 下课时间戳，毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 开始上课时间戳，毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 课堂状态，0:未开始 1:上课中 2:已下课。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 课堂标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 白板ID
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
	// 白板录制ID
	WhiteboardRecordId *string `json:"WhiteboardRecordId,omitempty" xml:"WhiteboardRecordId,omitempty"`
}

func (s GetClassDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClassDetailResponseBodyResult) SetClassId(v string) *GetClassDetailResponseBodyResult {
	s.ClassId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetConfId(v string) *GetClassDetailResponseBodyResult {
	s.ConfId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetCreateNickname(v string) *GetClassDetailResponseBodyResult {
	s.CreateNickname = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetCreateUserId(v string) *GetClassDetailResponseBodyResult {
	s.CreateUserId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetEndTime(v int64) *GetClassDetailResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetLiveId(v string) *GetClassDetailResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetRoomId(v string) *GetClassDetailResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetStartTime(v int64) *GetClassDetailResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetStatus(v int32) *GetClassDetailResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetTitle(v string) *GetClassDetailResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetWhiteboardId(v string) *GetClassDetailResponseBodyResult {
	s.WhiteboardId = &v
	return s
}

func (s *GetClassDetailResponseBodyResult) SetWhiteboardRecordId(v string) *GetClassDetailResponseBodyResult {
	s.WhiteboardRecordId = &v
	return s
}

type GetClassDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClassDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClassDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClassDetailResponse) GoString() string {
	return s.String()
}

func (s *GetClassDetailResponse) SetHeaders(v map[string]*string) *GetClassDetailResponse {
	s.Headers = v
	return s
}

func (s *GetClassDetailResponse) SetStatusCode(v int32) *GetClassDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClassDetailResponse) SetBody(v *GetClassDetailResponseBody) *GetClassDetailResponse {
	s.Body = v
	return s
}

type GetClassRecordRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课程唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作人用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetClassRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordRequest) GoString() string {
	return s.String()
}

func (s *GetClassRecordRequest) SetAppId(v string) *GetClassRecordRequest {
	s.AppId = &v
	return s
}

func (s *GetClassRecordRequest) SetClassId(v string) *GetClassRecordRequest {
	s.ClassId = &v
	return s
}

func (s *GetClassRecordRequest) SetUserId(v string) *GetClassRecordRequest {
	s.UserId = &v
	return s
}

type GetClassRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *GetClassRecordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetClassRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetClassRecordResponseBody) SetRequestId(v string) *GetClassRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClassRecordResponseBody) SetResult(v *GetClassRecordResponseBodyResult) *GetClassRecordResponseBody {
	s.Result = v
	return s
}

type GetClassRecordResponseBodyResult struct {
	PlaybackUrlMap map[string][]*string `json:"PlaybackUrlMap,omitempty" xml:"PlaybackUrlMap,omitempty"`
}

func (s GetClassRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClassRecordResponseBodyResult) SetPlaybackUrlMap(v map[string][]*string) *GetClassRecordResponseBodyResult {
	s.PlaybackUrlMap = v
	return s
}

type GetClassRecordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClassRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClassRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClassRecordResponse) GoString() string {
	return s.String()
}

func (s *GetClassRecordResponse) SetHeaders(v map[string]*string) *GetClassRecordResponse {
	s.Headers = v
	return s
}

func (s *GetClassRecordResponse) SetStatusCode(v int32) *GetClassRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClassRecordResponse) SetBody(v *GetClassRecordResponseBody) *GetClassRecordResponse {
	s.Body = v
	return s
}

type GetConferenceRequest struct {
	// 会议资源唯一标识。
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s GetConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceRequest) GoString() string {
	return s.String()
}

func (s *GetConferenceRequest) SetConferenceId(v string) *GetConferenceRequest {
	s.ConferenceId = &v
	return s
}

type GetConferenceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *GetConferenceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetConferenceResponseBody) SetRequestId(v string) *GetConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConferenceResponseBody) SetResult(v *GetConferenceResponseBodyResult) *GetConferenceResponseBody {
	s.Result = v
	return s
}

type GetConferenceResponseBodyResult struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 会议资源唯一标识。
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 会议创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 录制回放地址，m3u8格式，会议结束后10秒才会生成。
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 会议状态。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 会议标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 创建会议用户。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetConferenceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConferenceResponseBodyResult) SetAppId(v string) *GetConferenceResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetConferenceId(v string) *GetConferenceResponseBodyResult {
	s.ConferenceId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetCreateTime(v int64) *GetConferenceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetPlaybackUrl(v string) *GetConferenceResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetRoomId(v string) *GetConferenceResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetStatus(v string) *GetConferenceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetTitle(v string) *GetConferenceResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetConferenceResponseBodyResult) SetUserId(v string) *GetConferenceResponseBodyResult {
	s.UserId = &v
	return s
}

type GetConferenceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceResponse) GoString() string {
	return s.String()
}

func (s *GetConferenceResponse) SetHeaders(v map[string]*string) *GetConferenceResponse {
	s.Headers = v
	return s
}

func (s *GetConferenceResponse) SetStatusCode(v int32) *GetConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConferenceResponse) SetBody(v *GetConferenceResponseBody) *GetConferenceResponse {
	s.Body = v
	return s
}

type GetLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s GetLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRequest) SetLiveId(v string) *GetLiveRequest {
	s.LiveId = &v
	return s
}

type GetLiveResponseBody struct {
	// Id of the request
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetLiveResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBody) SetRequestId(v string) *GetLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveResponseBody) SetResult(v *GetLiveResponseBodyResult) *GetLiveResponseBody {
	s.Result = v
	return s
}

type GetLiveResponseBodyResult struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// rts播流信息
	ArtcInfo *GetLiveResponseBodyResultArtcInfo `json:"ArtcInfo,omitempty" xml:"ArtcInfo,omitempty" type:"Struct"`
	// 直播推送分辨率 -1:lld 1:lsd 2:lhd 3:lud
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// 封面图片
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播创建时间（毫秒ms）
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 直播持续时间（毫秒ms）
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 直播结束时间（毫秒ms）
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// hls播放地址
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// 直播简介
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 多分辨率多协议播放信息
	PlayUrlInfoList []*GetLiveResponseBodyResultPlayUrlInfoList `json:"PlayUrlInfoList,omitempty" xml:"PlayUrlInfoList,omitempty" type:"Repeated"`
	// 直播回放地址
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// 直播推流地址
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 房间id
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态：Created-已创建，未开播，Living-直播中，End-直播结束
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 直播标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 用户自定义数据存储
	UserDefineField *string `json:"UserDefineField,omitempty" xml:"UserDefineField,omitempty"`
	// 创建直播用户
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBodyResult) SetAnchorId(v string) *GetLiveResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetAppId(v string) *GetLiveResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetArtcInfo(v *GetLiveResponseBodyResultArtcInfo) *GetLiveResponseBodyResult {
	s.ArtcInfo = v
	return s
}

func (s *GetLiveResponseBodyResult) SetCodeLevel(v int32) *GetLiveResponseBodyResult {
	s.CodeLevel = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetCoverUrl(v string) *GetLiveResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetCreateTime(v int64) *GetLiveResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetDuration(v int64) *GetLiveResponseBodyResult {
	s.Duration = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetEndTime(v int64) *GetLiveResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetHlsUrl(v string) *GetLiveResponseBodyResult {
	s.HlsUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetIntroduction(v string) *GetLiveResponseBodyResult {
	s.Introduction = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetLiveId(v string) *GetLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetLiveUrl(v string) *GetLiveResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPlayUrlInfoList(v []*GetLiveResponseBodyResultPlayUrlInfoList) *GetLiveResponseBodyResult {
	s.PlayUrlInfoList = v
	return s
}

func (s *GetLiveResponseBodyResult) SetPlaybackUrl(v string) *GetLiveResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetPushUrl(v string) *GetLiveResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetRoomId(v string) *GetLiveResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetStatus(v string) *GetLiveResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetTitle(v string) *GetLiveResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetUserDefineField(v string) *GetLiveResponseBodyResult {
	s.UserDefineField = &v
	return s
}

func (s *GetLiveResponseBodyResult) SetUserId(v string) *GetLiveResponseBodyResult {
	s.UserId = &v
	return s
}

type GetLiveResponseBodyResultArtcInfo struct {
	// 原画转码地址
	ArtcH5Url *string `json:"ArtcH5Url,omitempty" xml:"ArtcH5Url,omitempty"`
	// 源码地址
	ArtcUrl *string `json:"ArtcUrl,omitempty" xml:"ArtcUrl,omitempty"`
}

func (s GetLiveResponseBodyResultArtcInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBodyResultArtcInfo) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBodyResultArtcInfo) SetArtcH5Url(v string) *GetLiveResponseBodyResultArtcInfo {
	s.ArtcH5Url = &v
	return s
}

func (s *GetLiveResponseBodyResultArtcInfo) SetArtcUrl(v string) *GetLiveResponseBodyResultArtcInfo {
	s.ArtcUrl = &v
	return s
}

type GetLiveResponseBodyResultPlayUrlInfoList struct {
	// 直播拉取分辨率 -1:lld 1:lsd 2:lhd 3:lud
	CodeLevel *int32 `json:"CodeLevel,omitempty" xml:"CodeLevel,omitempty"`
	// flv拉流地址
	FlvUrl *string `json:"FlvUrl,omitempty" xml:"FlvUrl,omitempty"`
	// hls拉流地址
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// rtmp拉流地址
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
}

func (s GetLiveResponseBodyResultPlayUrlInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponseBodyResultPlayUrlInfoList) GoString() string {
	return s.String()
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetCodeLevel(v int32) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.CodeLevel = &v
	return s
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetFlvUrl(v string) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.FlvUrl = &v
	return s
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetHlsUrl(v string) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.HlsUrl = &v
	return s
}

func (s *GetLiveResponseBodyResultPlayUrlInfoList) SetRtmpUrl(v string) *GetLiveResponseBodyResultPlayUrlInfoList {
	s.RtmpUrl = &v
	return s
}

type GetLiveResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveResponse) GoString() string {
	return s.String()
}

func (s *GetLiveResponse) SetHeaders(v map[string]*string) *GetLiveResponse {
	s.Headers = v
	return s
}

func (s *GetLiveResponse) SetStatusCode(v int32) *GetLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveResponse) SetBody(v *GetLiveResponseBody) *GetLiveResponse {
	s.Body = v
	return s
}

type GetLiveRecordRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播唯一标识，由调用CreateLiveRoom返回。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetLiveRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRecordRequest) SetAppId(v string) *GetLiveRecordRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRecordRequest) SetLiveId(v string) *GetLiveRecordRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveRecordRequest) SetUserId(v string) *GetLiveRecordRequest {
	s.UserId = &v
	return s
}

type GetLiveRecordResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *GetLiveRecordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRecordResponseBody) SetRequestId(v string) *GetLiveRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRecordResponseBody) SetResult(v *GetLiveRecordResponseBodyResult) *GetLiveRecordResponseBody {
	s.Result = v
	return s
}

type GetLiveRecordResponseBodyResult struct {
	PlaybackUrlMap map[string][]*string `json:"PlaybackUrlMap,omitempty" xml:"PlaybackUrlMap,omitempty"`
}

func (s GetLiveRecordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRecordResponseBodyResult) SetPlaybackUrlMap(v map[string][]*string) *GetLiveRecordResponseBodyResult {
	s.PlaybackUrlMap = v
	return s
}

type GetLiveRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRecordResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRecordResponse) SetHeaders(v map[string]*string) *GetLiveRecordResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRecordResponse) SetStatusCode(v int32) *GetLiveRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveRecordResponse) SetBody(v *GetLiveRecordResponseBody) *GetLiveRecordResponse {
	s.Body = v
	return s
}

type GetLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s GetLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRoomRequest) SetAppId(v string) *GetLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomRequest) SetLiveId(v string) *GetLiveRoomRequest {
	s.LiveId = &v
	return s
}

type GetLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *GetLiveRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBody) SetRequestId(v string) *GetLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRoomResponseBody) SetResult(v *GetLiveRoomResponseBodyResult) *GetLiveRoomResponseBody {
	s.Result = v
	return s
}

type GetLiveRoomResponseBodyResult struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// RTS低延迟播流信息。
	ArtcInfo *GetLiveRoomResponseBodyResultArtcInfo `json:"ArtcInfo,omitempty" xml:"ArtcInfo,omitempty" type:"Struct"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播创建时间，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 是否开启连麦。
	EnableLinkMic *bool `json:"EnableLinkMic,omitempty" xml:"EnableLinkMic,omitempty"`
	// 直播结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 原画HLS播放地址。
	HlsUrl *string `json:"HlsUrl,omitempty" xml:"HlsUrl,omitempty"`
	// https协议的原画HLS播放地址。
	HlsUrlHttps *string `json:"HlsUrlHttps,omitempty" xml:"HlsUrlHttps,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址。
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// https协议的直播拉流地址。
	LiveUrlHttps *string `json:"LiveUrlHttps,omitempty" xml:"LiveUrlHttps,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 直播回放地址。
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// https协议的直播回放地址
	PlaybackUrlHttps *string `json:"PlaybackUrlHttps,omitempty" xml:"PlaybackUrlHttps,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*GetLiveRoomResponseBodyResultPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 直播推流地址。
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// rtmp协议的播放地址
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
	// 直播开始时间，单位：毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 直播状态，0-在播 1-下播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s GetLiveRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBodyResult) SetAnchorId(v string) *GetLiveRoomResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetAnchorNick(v string) *GetLiveRoomResponseBodyResult {
	s.AnchorNick = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetAppId(v string) *GetLiveRoomResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetArtcInfo(v *GetLiveRoomResponseBodyResultArtcInfo) *GetLiveRoomResponseBodyResult {
	s.ArtcInfo = v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetChatId(v string) *GetLiveRoomResponseBodyResult {
	s.ChatId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetConfId(v string) *GetLiveRoomResponseBodyResult {
	s.ConfId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetCoverUrl(v string) *GetLiveRoomResponseBodyResult {
	s.CoverUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetCreateTime(v int64) *GetLiveRoomResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetEnableLinkMic(v bool) *GetLiveRoomResponseBodyResult {
	s.EnableLinkMic = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetEndTime(v int64) *GetLiveRoomResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetExtension(v map[string]*string) *GetLiveRoomResponseBodyResult {
	s.Extension = v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetHlsUrl(v string) *GetLiveRoomResponseBodyResult {
	s.HlsUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetHlsUrlHttps(v string) *GetLiveRoomResponseBodyResult {
	s.HlsUrlHttps = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetLiveId(v string) *GetLiveRoomResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetLiveUrl(v string) *GetLiveRoomResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetLiveUrlHttps(v string) *GetLiveRoomResponseBodyResult {
	s.LiveUrlHttps = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetNotice(v string) *GetLiveRoomResponseBodyResult {
	s.Notice = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetOnlineCount(v int64) *GetLiveRoomResponseBodyResult {
	s.OnlineCount = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPlaybackUrl(v string) *GetLiveRoomResponseBodyResult {
	s.PlaybackUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPlaybackUrlHttps(v string) *GetLiveRoomResponseBodyResult {
	s.PlaybackUrlHttps = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPluginInstanceInfoList(v []*GetLiveRoomResponseBodyResultPluginInstanceInfoList) *GetLiveRoomResponseBodyResult {
	s.PluginInstanceInfoList = v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPushUrl(v string) *GetLiveRoomResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetPv(v int64) *GetLiveRoomResponseBodyResult {
	s.Pv = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetRoomId(v string) *GetLiveRoomResponseBodyResult {
	s.RoomId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetRtmpUrl(v string) *GetLiveRoomResponseBodyResult {
	s.RtmpUrl = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetStartTime(v int64) *GetLiveRoomResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetStatus(v int32) *GetLiveRoomResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetTitle(v string) *GetLiveRoomResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetLiveRoomResponseBodyResult) SetUv(v int64) *GetLiveRoomResponseBodyResult {
	s.Uv = &v
	return s
}

type GetLiveRoomResponseBodyResultArtcInfo struct {
	// RTS转码流地址，推荐web端使用。
	ArtcH5Url *string `json:"ArtcH5Url,omitempty" xml:"ArtcH5Url,omitempty"`
	// RTS原码流地址，推荐移动端使用。
	ArtcUrl *string `json:"ArtcUrl,omitempty" xml:"ArtcUrl,omitempty"`
}

func (s GetLiveRoomResponseBodyResultArtcInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBodyResultArtcInfo) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBodyResultArtcInfo) SetArtcH5Url(v string) *GetLiveRoomResponseBodyResultArtcInfo {
	s.ArtcH5Url = &v
	return s
}

func (s *GetLiveRoomResponseBodyResultArtcInfo) SetArtcUrl(v string) *GetLiveRoomResponseBodyResultArtcInfo {
	s.ArtcUrl = &v
	return s
}

type GetLiveRoomResponseBodyResultPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s GetLiveRoomResponseBodyResultPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponseBodyResultPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetCreateTime(v int64) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetExtension(v map[string]*string) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginId(v string) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *GetLiveRoomResponseBodyResultPluginInstanceInfoList) SetPluginType(v string) *GetLiveRoomResponseBodyResultPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

type GetLiveRoomResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRoomResponse) SetHeaders(v map[string]*string) *GetLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRoomResponse) SetStatusCode(v int32) *GetLiveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveRoomResponse) SetBody(v *GetLiveRoomResponseBody) *GetLiveRoomResponse {
	s.Body = v
	return s
}

type GetLiveRoomStatisticsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
}

func (s GetLiveRoomStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsRequest) SetAppId(v string) *GetLiveRoomStatisticsRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomStatisticsRequest) SetLiveId(v string) *GetLiveRoomStatisticsRequest {
	s.LiveId = &v
	return s
}

type GetLiveRoomStatisticsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *GetLiveRoomStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRoomStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsResponseBody) SetRequestId(v string) *GetLiveRoomStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBody) SetResult(v *GetLiveRoomStatisticsResponseBodyResult) *GetLiveRoomStatisticsResponseBody {
	s.Result = v
	return s
}

type GetLiveRoomStatisticsResponseBodyResult struct {
	// 直播结束时间，单位：毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 点赞数。
	LikeCount *int64 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 互动消息数。
	MessageCount *int64 `json:"MessageCount,omitempty" xml:"MessageCount,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 直播开始时间，单位：毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 直播状态，0-已创建 1-直播中 2-已关闭。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
	// 总观看时长，单位：毫秒。
	WatchLiveTime *int64 `json:"WatchLiveTime,omitempty" xml:"WatchLiveTime,omitempty"`
}

func (s GetLiveRoomStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetEndTime(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetLikeCount(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.LikeCount = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetLiveId(v string) *GetLiveRoomStatisticsResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetMessageCount(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.MessageCount = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetOnlineCount(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.OnlineCount = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetPv(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.Pv = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetStartTime(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetStatus(v int32) *GetLiveRoomStatisticsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetUv(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.Uv = &v
	return s
}

func (s *GetLiveRoomStatisticsResponseBodyResult) SetWatchLiveTime(v int64) *GetLiveRoomStatisticsResponseBodyResult {
	s.WatchLiveTime = &v
	return s
}

type GetLiveRoomStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveRoomStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRoomStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRoomStatisticsResponse) SetHeaders(v map[string]*string) *GetLiveRoomStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRoomStatisticsResponse) SetStatusCode(v int32) *GetLiveRoomStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveRoomStatisticsResponse) SetBody(v *GetLiveRoomStatisticsResponseBody) *GetLiveRoomStatisticsResponse {
	s.Body = v
	return s
}

type GetLiveRoomUserStatisticsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetLiveRoomUserStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsRequest) SetAppId(v string) *GetLiveRoomUserStatisticsRequest {
	s.AppId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsRequest) SetLiveId(v string) *GetLiveRoomUserStatisticsRequest {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsRequest) SetPageNumber(v string) *GetLiveRoomUserStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLiveRoomUserStatisticsRequest) SetPageSize(v string) *GetLiveRoomUserStatisticsRequest {
	s.PageSize = &v
	return s
}

type GetLiveRoomUserStatisticsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *GetLiveRoomUserStatisticsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetLiveRoomUserStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponseBody) SetRequestId(v string) *GetLiveRoomUserStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBody) SetResult(v *GetLiveRoomUserStatisticsResponseBodyResult) *GetLiveRoomUserStatisticsResponseBody {
	s.Result = v
	return s
}

type GetLiveRoomUserStatisticsResponseBodyResult struct {
	// 是否还有下一页。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 用户总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 用户总数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 用户观看数据列表。
	UserStatisticsList []*GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList `json:"UserStatisticsList,omitempty" xml:"UserStatisticsList,omitempty" type:"Repeated"`
}

func (s GetLiveRoomUserStatisticsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetHasMore(v bool) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetLiveId(v string) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetPageTotal(v int32) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetTotalCount(v int32) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResult) SetUserStatisticsList(v []*GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) *GetLiveRoomUserStatisticsResponseBodyResult {
	s.UserStatisticsList = v
	return s
}

type GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList struct {
	CommentCount *int32 `json:"CommentCount,omitempty" xml:"CommentCount,omitempty"`
	LikeCount    *int32 `json:"LikeCount,omitempty" xml:"LikeCount,omitempty"`
	// 用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 观看时长，单位：毫秒。
	WatchLiveTime *int64 `json:"WatchLiveTime,omitempty" xml:"WatchLiveTime,omitempty"`
}

func (s GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) SetCommentCount(v int32) *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList {
	s.CommentCount = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) SetLikeCount(v int32) *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList {
	s.LikeCount = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) SetUserId(v string) *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList {
	s.UserId = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList) SetWatchLiveTime(v int64) *GetLiveRoomUserStatisticsResponseBodyResultUserStatisticsList {
	s.WatchLiveTime = &v
	return s
}

type GetLiveRoomUserStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveRoomUserStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveRoomUserStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveRoomUserStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetLiveRoomUserStatisticsResponse) SetHeaders(v map[string]*string) *GetLiveRoomUserStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetLiveRoomUserStatisticsResponse) SetStatusCode(v int32) *GetLiveRoomUserStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveRoomUserStatisticsResponse) SetBody(v *GetLiveRoomUserStatisticsResponseBody) *GetLiveRoomUserStatisticsResponse {
	s.Body = v
	return s
}

type GetRecordFileInfoRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetRecordFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordFileInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRecordFileInfoRequest) SetAppId(v string) *GetRecordFileInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetRecordFileInfoRequest) SetMaxResults(v int32) *GetRecordFileInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRecordFileInfoRequest) SetNextToken(v string) *GetRecordFileInfoRequest {
	s.NextToken = &v
	return s
}

type GetRecordFileInfoResponseBody struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetRecordFileInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRecordFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordFileInfoResponseBody) SetDownloadUrl(v string) *GetRecordFileInfoResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetRecordFileInfoResponseBody) SetMaxResults(v int32) *GetRecordFileInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetRecordFileInfoResponseBody) SetNextToken(v string) *GetRecordFileInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRecordFileInfoResponseBody) SetRequestId(v string) *GetRecordFileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecordFileInfoResponseBody) SetResult(v *GetRecordFileInfoResponseBodyResult) *GetRecordFileInfoResponseBody {
	s.Result = v
	return s
}

type GetRecordFileInfoResponseBodyResult struct {
	ETag       *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetRecordFileInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRecordFileInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRecordFileInfoResponseBodyResult) SetETag(v string) *GetRecordFileInfoResponseBodyResult {
	s.ETag = &v
	return s
}

func (s *GetRecordFileInfoResponseBodyResult) SetExpireTime(v string) *GetRecordFileInfoResponseBodyResult {
	s.ExpireTime = &v
	return s
}

func (s *GetRecordFileInfoResponseBodyResult) SetFileName(v string) *GetRecordFileInfoResponseBodyResult {
	s.FileName = &v
	return s
}

type GetRecordFileInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRecordFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecordFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordFileInfoResponse) GoString() string {
	return s.String()
}

func (s *GetRecordFileInfoResponse) SetHeaders(v map[string]*string) *GetRecordFileInfoResponse {
	s.Headers = v
	return s
}

func (s *GetRecordFileInfoResponse) SetStatusCode(v int32) *GetRecordFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordFileInfoResponse) SetBody(v *GetRecordFileInfoResponseBody) *GetRecordFileInfoResponse {
	s.Body = v
	return s
}

type GetRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s GetRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoomRequest) GoString() string {
	return s.String()
}

func (s *GetRoomRequest) SetAppId(v string) *GetRoomRequest {
	s.AppId = &v
	return s
}

func (s *GetRoomRequest) SetRoomId(v string) *GetRoomRequest {
	s.RoomId = &v
	return s
}

type GetRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 查询房间信息返回结果。
	Result *GetRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBody) SetRequestId(v string) *GetRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoomResponseBody) SetResult(v *GetRoomResponseBodyResult) *GetRoomResponseBody {
	s.Result = v
	return s
}

type GetRoomResponseBodyResult struct {
	// 房间信息。
	RoomInfo *GetRoomResponseBodyResultRoomInfo `json:"RoomInfo,omitempty" xml:"RoomInfo,omitempty" type:"Struct"`
}

func (s GetRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResult) SetRoomInfo(v *GetRoomResponseBodyResultRoomInfo) *GetRoomResponseBodyResult {
	s.RoomInfo = v
	return s
}

type GetRoomResponseBodyResultRoomInfo struct {
	// 管理员ID列表。
	AdminIdList []*string `json:"AdminIdList,omitempty" xml:"AdminIdList,omitempty" type:"Repeated"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 创建房间使用的模板id。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s GetRoomResponseBodyResultRoomInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResultRoomInfo) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResultRoomInfo) SetAdminIdList(v []*string) *GetRoomResponseBodyResultRoomInfo {
	s.AdminIdList = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetAppId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.AppId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetCreateTime(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.CreateTime = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetExtension(v map[string]*string) *GetRoomResponseBodyResultRoomInfo {
	s.Extension = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetNotice(v string) *GetRoomResponseBodyResultRoomInfo {
	s.Notice = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetOnlineCount(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.OnlineCount = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetPluginInstanceInfoList(v []*GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) *GetRoomResponseBodyResultRoomInfo {
	s.PluginInstanceInfoList = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetPv(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.Pv = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetRoomId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.RoomId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetRoomOwnerId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.RoomOwnerId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetTemplateId(v string) *GetRoomResponseBodyResultRoomInfo {
	s.TemplateId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetTitle(v string) *GetRoomResponseBodyResultRoomInfo {
	s.Title = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfo) SetUv(v int64) *GetRoomResponseBodyResultRoomInfo {
	s.Uv = &v
	return s
}

type GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetCreateTime(v int64) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetExtension(v map[string]*string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetPluginId(v string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList) SetPluginType(v string) *GetRoomResponseBodyResultRoomInfoPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

type GetRoomResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoomResponse) GoString() string {
	return s.String()
}

func (s *GetRoomResponse) SetHeaders(v map[string]*string) *GetRoomResponse {
	s.Headers = v
	return s
}

func (s *GetRoomResponse) SetStatusCode(v int32) *GetRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoomResponse) SetBody(v *GetRoomResponseBody) *GetRoomResponse {
	s.Body = v
	return s
}

type GetStandardRoomJumpUrlRequest struct {
	// 用户的应用ID，在控制台创建应用时生成
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 终端设备类型,通过控制台创建和查询
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// 资源ID：根据业务类型来定，比如直播ID，课堂ID等
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// 业务类型：互动直播live，互动课堂class
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// 平台：win, mac, android, ios, web
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// 用户UserId,在AppId下单独唯一
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户昵称
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s GetStandardRoomJumpUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlRequest) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlRequest) SetAppId(v string) *GetStandardRoomJumpUrlRequest {
	s.AppId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetAppKey(v string) *GetStandardRoomJumpUrlRequest {
	s.AppKey = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetBizId(v string) *GetStandardRoomJumpUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetBizType(v string) *GetStandardRoomJumpUrlRequest {
	s.BizType = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetPlatform(v string) *GetStandardRoomJumpUrlRequest {
	s.Platform = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetUserId(v string) *GetStandardRoomJumpUrlRequest {
	s.UserId = &v
	return s
}

func (s *GetStandardRoomJumpUrlRequest) SetUserNick(v string) *GetStandardRoomJumpUrlRequest {
	s.UserNick = &v
	return s
}

type GetStandardRoomJumpUrlResponseBody struct {
	// Id of the request
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetStandardRoomJumpUrlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetStandardRoomJumpUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlResponseBody) SetRequestId(v string) *GetStandardRoomJumpUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardRoomJumpUrlResponseBody) SetResult(v *GetStandardRoomJumpUrlResponseBodyResult) *GetStandardRoomJumpUrlResponseBody {
	s.Result = v
	return s
}

type GetStandardRoomJumpUrlResponseBodyResult struct {
	// 样板间跳转协议地址
	StandardRoomJumpUrl *string `json:"StandardRoomJumpUrl,omitempty" xml:"StandardRoomJumpUrl,omitempty"`
}

func (s GetStandardRoomJumpUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlResponseBodyResult) SetStandardRoomJumpUrl(v string) *GetStandardRoomJumpUrlResponseBodyResult {
	s.StandardRoomJumpUrl = &v
	return s
}

type GetStandardRoomJumpUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStandardRoomJumpUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStandardRoomJumpUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStandardRoomJumpUrlResponse) GoString() string {
	return s.String()
}

func (s *GetStandardRoomJumpUrlResponse) SetHeaders(v map[string]*string) *GetStandardRoomJumpUrlResponse {
	s.Headers = v
	return s
}

func (s *GetStandardRoomJumpUrlResponse) SetStatusCode(v int32) *GetStandardRoomJumpUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardRoomJumpUrlResponse) SetBody(v *GetStandardRoomJumpUrlResponseBody) *GetStandardRoomJumpUrlResponse {
	s.Body = v
	return s
}

type KickRoomUserRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BlockTime *int64  `json:"BlockTime,omitempty" xml:"BlockTime,omitempty"`
	// 被踢出房间的用户ID。
	KickUser *string `json:"KickUser,omitempty" xml:"KickUser,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位，传空则随机生成一个房间id。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 操作人的用户ID，用于表示谁执行了踢人操作。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s KickRoomUserRequest) String() string {
	return tea.Prettify(s)
}

func (s KickRoomUserRequest) GoString() string {
	return s.String()
}

func (s *KickRoomUserRequest) SetAppId(v string) *KickRoomUserRequest {
	s.AppId = &v
	return s
}

func (s *KickRoomUserRequest) SetBlockTime(v int64) *KickRoomUserRequest {
	s.BlockTime = &v
	return s
}

func (s *KickRoomUserRequest) SetKickUser(v string) *KickRoomUserRequest {
	s.KickUser = &v
	return s
}

func (s *KickRoomUserRequest) SetRoomId(v string) *KickRoomUserRequest {
	s.RoomId = &v
	return s
}

func (s *KickRoomUserRequest) SetUserId(v string) *KickRoomUserRequest {
	s.UserId = &v
	return s
}

type KickRoomUserResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KickRoomUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KickRoomUserResponseBody) GoString() string {
	return s.String()
}

func (s *KickRoomUserResponseBody) SetRequestId(v string) *KickRoomUserResponseBody {
	s.RequestId = &v
	return s
}

type KickRoomUserResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *KickRoomUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KickRoomUserResponse) String() string {
	return tea.Prettify(s)
}

func (s KickRoomUserResponse) GoString() string {
	return s.String()
}

func (s *KickRoomUserResponse) SetHeaders(v map[string]*string) *KickRoomUserResponse {
	s.Headers = v
	return s
}

func (s *KickRoomUserResponse) SetStatusCode(v int32) *KickRoomUserResponse {
	s.StatusCode = &v
	return s
}

func (s *KickRoomUserResponse) SetBody(v *KickRoomUserResponseBody) *KickRoomUserResponse {
	s.Body = v
	return s
}

type ListClassesRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 课程状态，0-未开课 1-上课中 2-已下课，不传则返回所有课程。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListClassesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClassesRequest) GoString() string {
	return s.String()
}

func (s *ListClassesRequest) SetAppId(v string) *ListClassesRequest {
	s.AppId = &v
	return s
}

func (s *ListClassesRequest) SetPageNumber(v int32) *ListClassesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClassesRequest) SetPageSize(v int32) *ListClassesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClassesRequest) SetStatus(v int32) *ListClassesRequest {
	s.Status = &v
	return s
}

type ListClassesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建课程返回的结果。
	Result *ListClassesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListClassesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBody) SetRequestId(v string) *ListClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClassesResponseBody) SetResult(v *ListClassesResponseBodyResult) *ListClassesResponseBody {
	s.Result = v
	return s
}

type ListClassesResponseBodyResult struct {
	// 课程列表信息。
	ClassList []*ListClassesResponseBodyResultClassList `json:"ClassList,omitempty" xml:"ClassList,omitempty" type:"Repeated"`
	// 是否还有下一页。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 课程总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 课程总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClassesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBodyResult) SetClassList(v []*ListClassesResponseBodyResultClassList) *ListClassesResponseBodyResult {
	s.ClassList = v
	return s
}

func (s *ListClassesResponseBodyResult) SetHasMore(v bool) *ListClassesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListClassesResponseBodyResult) SetPageTotal(v int32) *ListClassesResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListClassesResponseBodyResult) SetTotalCount(v int32) *ListClassesResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListClassesResponseBodyResultClassList struct {
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 连麦会议唯一标识。
	ConfId *string `json:"ConfId,omitempty" xml:"ConfId,omitempty"`
	// 创建人昵称。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人ID。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 下课时间戳，毫秒。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播的唯一标识ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 开始上课时间戳，毫秒。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 课堂状态，0:未开始 1:上课中 2:已下课。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 课堂标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 白板ID
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
	// 白板录制ID
	WhiteboardRecordId *string `json:"WhiteboardRecordId,omitempty" xml:"WhiteboardRecordId,omitempty"`
}

func (s ListClassesResponseBodyResultClassList) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponseBodyResultClassList) GoString() string {
	return s.String()
}

func (s *ListClassesResponseBodyResultClassList) SetClassId(v string) *ListClassesResponseBodyResultClassList {
	s.ClassId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetConfId(v string) *ListClassesResponseBodyResultClassList {
	s.ConfId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetCreateNickname(v string) *ListClassesResponseBodyResultClassList {
	s.CreateNickname = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetCreateUserId(v string) *ListClassesResponseBodyResultClassList {
	s.CreateUserId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetEndTime(v int64) *ListClassesResponseBodyResultClassList {
	s.EndTime = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetLiveId(v string) *ListClassesResponseBodyResultClassList {
	s.LiveId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetRoomId(v string) *ListClassesResponseBodyResultClassList {
	s.RoomId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetStartTime(v int64) *ListClassesResponseBodyResultClassList {
	s.StartTime = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetStatus(v int32) *ListClassesResponseBodyResultClassList {
	s.Status = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetTitle(v string) *ListClassesResponseBodyResultClassList {
	s.Title = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetWhiteboardId(v string) *ListClassesResponseBodyResultClassList {
	s.WhiteboardId = &v
	return s
}

func (s *ListClassesResponseBodyResultClassList) SetWhiteboardRecordId(v string) *ListClassesResponseBodyResultClassList {
	s.WhiteboardRecordId = &v
	return s
}

type ListClassesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClassesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClassesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClassesResponse) GoString() string {
	return s.String()
}

func (s *ListClassesResponse) SetHeaders(v map[string]*string) *ListClassesResponse {
	s.Headers = v
	return s
}

func (s *ListClassesResponse) SetStatusCode(v int32) *ListClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClassesResponse) SetBody(v *ListClassesResponseBody) *ListClassesResponse {
	s.Body = v
	return s
}

type ListCommentsRequest struct {
	// 用户的应用ID，在控制台创建应用时生成。包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 弹幕创建者ID。
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// 查询弹幕消息列表的分页页数。应该从1开始，每次分页拉取时递增。
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 查询弹幕消息列表的分页大小。最小不得小于1，最大不得超过100。如果超过100，会被截断为前100条。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 房间的唯一标识，在调用CreateRoom时返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 查询弹幕消息列表的排序方式。取值仅限0和1，其中0表示按照弹幕消息创建时间递增的顺序拉取，1表示按照弹幕消息创建时间递减的时间拉取。
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// 操作人用户ID，表示谁执行了查询弹幕消息列表的操作。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCommentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsRequest) GoString() string {
	return s.String()
}

func (s *ListCommentsRequest) SetAppId(v string) *ListCommentsRequest {
	s.AppId = &v
	return s
}

func (s *ListCommentsRequest) SetCreatorId(v string) *ListCommentsRequest {
	s.CreatorId = &v
	return s
}

func (s *ListCommentsRequest) SetPageNum(v int32) *ListCommentsRequest {
	s.PageNum = &v
	return s
}

func (s *ListCommentsRequest) SetPageSize(v int32) *ListCommentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommentsRequest) SetRoomId(v string) *ListCommentsRequest {
	s.RoomId = &v
	return s
}

func (s *ListCommentsRequest) SetSortType(v int32) *ListCommentsRequest {
	s.SortType = &v
	return s
}

func (s *ListCommentsRequest) SetUserId(v string) *ListCommentsRequest {
	s.UserId = &v
	return s
}

type ListCommentsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用查询弹幕消息列表的返回结果。
	Result *ListCommentsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListCommentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommentsResponseBody) SetRequestId(v string) *ListCommentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommentsResponseBody) SetResult(v *ListCommentsResponseBodyResult) *ListCommentsResponseBody {
	s.Result = v
	return s
}

type ListCommentsResponseBodyResult struct {
	// 弹幕消息列表。
	CommentVOList []*ListCommentsResponseBodyResultCommentVOList `json:"CommentVOList,omitempty" xml:"CommentVOList,omitempty" type:"Repeated"`
	// 是否还有下一页数据。true表示还有，false表示没有。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 分页查询弹幕消息列表的总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 弹幕消息的总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCommentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommentsResponseBodyResult) SetCommentVOList(v []*ListCommentsResponseBodyResultCommentVOList) *ListCommentsResponseBodyResult {
	s.CommentVOList = v
	return s
}

func (s *ListCommentsResponseBodyResult) SetHasMore(v bool) *ListCommentsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListCommentsResponseBodyResult) SetPageTotal(v int32) *ListCommentsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListCommentsResponseBodyResult) SetTotalCount(v int32) *ListCommentsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListCommentsResponseBodyResultCommentVOList struct {
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 弹幕消息的唯一ID标识。
	CommentId *string `json:"CommentId,omitempty" xml:"CommentId,omitempty"`
	// 弹幕消息的内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 弹幕消息的创建时间，Unix时间戳，单位：毫秒。
	CreateAt *int64 `json:"CreateAt,omitempty" xml:"CreateAt,omitempty"`
	// 扩展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕消息的发送者ID标识。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕消息发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
}

func (s ListCommentsResponseBodyResultCommentVOList) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponseBodyResultCommentVOList) GoString() string {
	return s.String()
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetAppId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.AppId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetCommentId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.CommentId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetContent(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.Content = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetCreateAt(v int64) *ListCommentsResponseBodyResultCommentVOList {
	s.CreateAt = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetExtension(v map[string]*string) *ListCommentsResponseBodyResultCommentVOList {
	s.Extension = v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetRoomId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.RoomId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetSenderId(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.SenderId = &v
	return s
}

func (s *ListCommentsResponseBodyResultCommentVOList) SetSenderNick(v string) *ListCommentsResponseBodyResultCommentVOList {
	s.SenderNick = &v
	return s
}

type ListCommentsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommentsResponse) GoString() string {
	return s.String()
}

func (s *ListCommentsResponse) SetHeaders(v map[string]*string) *ListCommentsResponse {
	s.Headers = v
	return s
}

func (s *ListCommentsResponse) SetStatusCode(v int32) *ListCommentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommentsResponse) SetBody(v *ListCommentsResponseBody) *ListCommentsResponse {
	s.Body = v
	return s
}

type ListConferenceUsersRequest struct {
	// 会议唯一标识符
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 查询页码，从第1页开始。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大显示个数为100。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListConferenceUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersRequest) SetConferenceId(v string) *ListConferenceUsersRequest {
	s.ConferenceId = &v
	return s
}

func (s *ListConferenceUsersRequest) SetPageNumber(v int32) *ListConferenceUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConferenceUsersRequest) SetPageSize(v int32) *ListConferenceUsersRequest {
	s.PageSize = &v
	return s
}

type ListConferenceUsersResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回结果
	Result *ListConferenceUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListConferenceUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBody) SetRequestId(v string) *ListConferenceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConferenceUsersResponseBody) SetResult(v *ListConferenceUsersResponseBodyResult) *ListConferenceUsersResponseBody {
	s.Result = v
	return s
}

type ListConferenceUsersResponseBodyResult struct {
	// 会议用户列表。
	ConferenceUserList []*ListConferenceUsersResponseBodyResultConferenceUserList `json:"ConferenceUserList,omitempty" xml:"ConferenceUserList,omitempty" type:"Repeated"`
	// 是否还有数据
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 总页数
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConferenceUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBodyResult) SetConferenceUserList(v []*ListConferenceUsersResponseBodyResultConferenceUserList) *ListConferenceUsersResponseBodyResult {
	s.ConferenceUserList = v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetHasMore(v bool) *ListConferenceUsersResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetPageTotal(v int32) *ListConferenceUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResult) SetTotalCount(v int32) *ListConferenceUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListConferenceUsersResponseBodyResultConferenceUserList struct {
	// 用户状态。
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListConferenceUsersResponseBodyResultConferenceUserList) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponseBodyResultConferenceUserList) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponseBodyResultConferenceUserList) SetStatus(v string) *ListConferenceUsersResponseBodyResultConferenceUserList {
	s.Status = &v
	return s
}

func (s *ListConferenceUsersResponseBodyResultConferenceUserList) SetUserId(v string) *ListConferenceUsersResponseBodyResultConferenceUserList {
	s.UserId = &v
	return s
}

type ListConferenceUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConferenceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConferenceUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConferenceUsersResponse) GoString() string {
	return s.String()
}

func (s *ListConferenceUsersResponse) SetHeaders(v map[string]*string) *ListConferenceUsersResponse {
	s.Headers = v
	return s
}

func (s *ListConferenceUsersResponse) SetStatusCode(v int32) *ListConferenceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConferenceUsersResponse) SetBody(v *ListConferenceUsersResponseBody) *ListConferenceUsersResponse {
	s.Body = v
	return s
}

type ListLiveFilesRequest struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLiveFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveFilesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveFilesRequest) SetAppId(v string) *ListLiveFilesRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveFilesRequest) SetLiveId(v string) *ListLiveFilesRequest {
	s.LiveId = &v
	return s
}

func (s *ListLiveFilesRequest) SetMaxResults(v int32) *ListLiveFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLiveFilesRequest) SetNextToken(v string) *ListLiveFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListLiveFilesRequest) SetPageNumber(v int64) *ListLiveFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLiveFilesRequest) SetPageSize(v int64) *ListLiveFilesRequest {
	s.PageSize = &v
	return s
}

type ListLiveFilesResponseBody struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveFilesResponseBody) SetFileName(v string) *ListLiveFilesResponseBody {
	s.FileName = &v
	return s
}

func (s *ListLiveFilesResponseBody) SetMaxResults(v int32) *ListLiveFilesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLiveFilesResponseBody) SetNextToken(v string) *ListLiveFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLiveFilesResponseBody) SetRequestId(v string) *ListLiveFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveFilesResponseBody) SetTotalCount(v int32) *ListLiveFilesResponseBody {
	s.TotalCount = &v
	return s
}

type ListLiveFilesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLiveFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveFilesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveFilesResponse) SetHeaders(v map[string]*string) *ListLiveFilesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveFilesResponse) SetStatusCode(v int32) *ListLiveFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveFilesResponse) SetBody(v *ListLiveFilesResponseBody) *ListLiveFilesResponse {
	s.Body = v
	return s
}

type ListLiveRoomsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 直播状态，0-在播 1-下播，不传则返回所有直播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLiveRoomsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsRequest) SetAppId(v string) *ListLiveRoomsRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsRequest) SetPageNumber(v int32) *ListLiveRoomsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLiveRoomsRequest) SetPageSize(v int32) *ListLiveRoomsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRoomsRequest) SetStatus(v int32) *ListLiveRoomsRequest {
	s.Status = &v
	return s
}

type ListLiveRoomsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *ListLiveRoomsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListLiveRoomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponseBody) SetRequestId(v string) *ListLiveRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRoomsResponseBody) SetResult(v *ListLiveRoomsResponseBodyResult) *ListLiveRoomsResponseBody {
	s.Result = v
	return s
}

type ListLiveRoomsResponseBodyResult struct {
	// 是否还有下一页。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 直播列表信息。
	LiveList []*ListLiveRoomsResponseBodyResultLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
	// 直播总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 直播总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveRoomsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponseBodyResult) SetHasMore(v bool) *ListLiveRoomsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResult) SetLiveList(v []*ListLiveRoomsResponseBodyResultLiveList) *ListLiveRoomsResponseBodyResult {
	s.LiveList = v
	return s
}

func (s *ListLiveRoomsResponseBodyResult) SetPageTotal(v int32) *ListLiveRoomsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResult) SetTotalCount(v int32) *ListLiveRoomsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListLiveRoomsResponseBodyResultLiveList struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播的创建时间。单位为ms。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 直播的结束时间。单位为ms。
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播的开始时间。单位为ms。
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 直播状态，0-在播 1-下播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s ListLiveRoomsResponseBodyResultLiveList) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponseBodyResultLiveList) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetAnchorId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.AnchorId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetAnchorNick(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.AnchorNick = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetAppId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetChatId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.ChatId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetCoverUrl(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.CoverUrl = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetCreateTime(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetEndTime(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.EndTime = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetExtension(v map[string]*string) *ListLiveRoomsResponseBodyResultLiveList {
	s.Extension = v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetLiveId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.LiveId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetNotice(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.Notice = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetOnlineCount(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.OnlineCount = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetPv(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.Pv = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetRoomId(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.RoomId = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetStartTime(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.StartTime = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetStatus(v int32) *ListLiveRoomsResponseBodyResultLiveList {
	s.Status = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetTitle(v string) *ListLiveRoomsResponseBodyResultLiveList {
	s.Title = &v
	return s
}

func (s *ListLiveRoomsResponseBodyResultLiveList) SetUv(v int64) *ListLiveRoomsResponseBodyResultLiveList {
	s.Uv = &v
	return s
}

type ListLiveRoomsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLiveRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveRoomsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsResponse) SetHeaders(v map[string]*string) *ListLiveRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRoomsResponse) SetStatusCode(v int32) *ListLiveRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRoomsResponse) SetBody(v *ListLiveRoomsResponseBody) *ListLiveRoomsResponse {
	s.Body = v
	return s
}

type ListLiveRoomsByIdRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID列表。
	LiveIdList []*string `json:"LiveIdList,omitempty" xml:"LiveIdList,omitempty" type:"Repeated"`
}

func (s ListLiveRoomsByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdRequest) SetAppId(v string) *ListLiveRoomsByIdRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsByIdRequest) SetLiveIdList(v []*string) *ListLiveRoomsByIdRequest {
	s.LiveIdList = v
	return s
}

type ListLiveRoomsByIdShrinkRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID列表。
	LiveIdListShrink *string `json:"LiveIdList,omitempty" xml:"LiveIdList,omitempty"`
}

func (s ListLiveRoomsByIdShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdShrinkRequest) SetAppId(v string) *ListLiveRoomsByIdShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsByIdShrinkRequest) SetLiveIdListShrink(v string) *ListLiveRoomsByIdShrinkRequest {
	s.LiveIdListShrink = &v
	return s
}

type ListLiveRoomsByIdResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *ListLiveRoomsByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListLiveRoomsByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponseBody) SetRequestId(v string) *ListLiveRoomsByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBody) SetResult(v *ListLiveRoomsByIdResponseBodyResult) *ListLiveRoomsByIdResponseBody {
	s.Result = v
	return s
}

type ListLiveRoomsByIdResponseBodyResult struct {
	// 直播列表信息。
	LiveList []*ListLiveRoomsByIdResponseBodyResultLiveList `json:"LiveList,omitempty" xml:"LiveList,omitempty" type:"Repeated"`
}

func (s ListLiveRoomsByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponseBodyResult) SetLiveList(v []*ListLiveRoomsByIdResponseBodyResultLiveList) *ListLiveRoomsByIdResponseBodyResult {
	s.LiveList = v
	return s
}

type ListLiveRoomsByIdResponseBodyResultLiveList struct {
	// 主播ID。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用ID。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 聊天ID。
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 封面。
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 直播拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 在线用户数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 访问用户人次。
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
	// 房间ID。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 直播状态，0-在播 1-下播。
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 访问用户数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s ListLiveRoomsByIdResponseBodyResultLiveList) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponseBodyResultLiveList) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetAnchorId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.AnchorId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetAnchorNick(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.AnchorNick = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetAppId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.AppId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetChatId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.ChatId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetCoverUrl(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.CoverUrl = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetExtension(v map[string]*string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Extension = v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetLiveId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.LiveId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetNotice(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Notice = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetOnlineCount(v int64) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.OnlineCount = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetPv(v int64) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Pv = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetRoomId(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.RoomId = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetStatus(v int32) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Status = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetTitle(v string) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Title = &v
	return s
}

func (s *ListLiveRoomsByIdResponseBodyResultLiveList) SetUv(v int64) *ListLiveRoomsByIdResponseBodyResultLiveList {
	s.Uv = &v
	return s
}

type ListLiveRoomsByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLiveRoomsByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveRoomsByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRoomsByIdResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRoomsByIdResponse) SetHeaders(v map[string]*string) *ListLiveRoomsByIdResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRoomsByIdResponse) SetStatusCode(v int32) *ListLiveRoomsByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRoomsByIdResponse) SetBody(v *ListLiveRoomsByIdResponseBody) *ListLiveRoomsByIdResponse {
	s.Body = v
	return s
}

type ListRoomUsersRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 房间ID，最大长度36个字符。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s ListRoomUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRoomUsersRequest) SetAppId(v string) *ListRoomUsersRequest {
	s.AppId = &v
	return s
}

func (s *ListRoomUsersRequest) SetPageNumber(v int32) *ListRoomUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoomUsersRequest) SetPageSize(v int32) *ListRoomUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoomUsersRequest) SetRoomId(v string) *ListRoomUsersRequest {
	s.RoomId = &v
	return s
}

type ListRoomUsersResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *ListRoomUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRoomUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBody) SetRequestId(v string) *ListRoomUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoomUsersResponseBody) SetResult(v *ListRoomUsersResponseBodyResult) *ListRoomUsersResponseBody {
	s.Result = v
	return s
}

type ListRoomUsersResponseBodyResult struct {
	// 是否还有下一页用户列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 该房间的用户总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 房间用户列表信息。
	RoomUserList []*ListRoomUsersResponseBodyResultRoomUserList `json:"RoomUserList,omitempty" xml:"RoomUserList,omitempty" type:"Repeated"`
	// 该房间的用户总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoomUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBodyResult) SetHasMore(v bool) *ListRoomUsersResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetPageTotal(v int32) *ListRoomUsersResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetRoomUserList(v []*ListRoomUsersResponseBodyResultRoomUserList) *ListRoomUsersResponseBodyResult {
	s.RoomUserList = v
	return s
}

func (s *ListRoomUsersResponseBodyResult) SetTotalCount(v int32) *ListRoomUsersResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListRoomUsersResponseBodyResultRoomUserList struct {
	// 用户拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 用户昵称。
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// 用户唯一标识。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRoomUsersResponseBodyResultRoomUserList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponseBodyResultRoomUserList) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponseBodyResultRoomUserList) SetExtension(v map[string]*string) *ListRoomUsersResponseBodyResultRoomUserList {
	s.Extension = v
	return s
}

func (s *ListRoomUsersResponseBodyResultRoomUserList) SetNick(v string) *ListRoomUsersResponseBodyResultRoomUserList {
	s.Nick = &v
	return s
}

func (s *ListRoomUsersResponseBodyResultRoomUserList) SetUserId(v string) *ListRoomUsersResponseBodyResultRoomUserList {
	s.UserId = &v
	return s
}

type ListRoomUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRoomUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoomUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoomUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRoomUsersResponse) SetHeaders(v map[string]*string) *ListRoomUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRoomUsersResponse) SetStatusCode(v int32) *ListRoomUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoomUsersResponse) SetBody(v *ListRoomUsersResponseBody) *ListRoomUsersResponse {
	s.Body = v
	return s
}

type ListRoomsRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 查询页码，从1开始，传空默认查询第1页。
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页显示个数，最大支持50，参数为空默认显示个数为10。
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRoomsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsRequest) GoString() string {
	return s.String()
}

func (s *ListRoomsRequest) SetAppId(v string) *ListRoomsRequest {
	s.AppId = &v
	return s
}

func (s *ListRoomsRequest) SetPageNumber(v int32) *ListRoomsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoomsRequest) SetPageSize(v int32) *ListRoomsRequest {
	s.PageSize = &v
	return s
}

type ListRoomsResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *ListRoomsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListRoomsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBody) SetRequestId(v string) *ListRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoomsResponseBody) SetResult(v *ListRoomsResponseBodyResult) *ListRoomsResponseBody {
	s.Result = v
	return s
}

type ListRoomsResponseBodyResult struct {
	// 是否还有下一页房间列表。
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// 该应用的房间总页数。
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// 房间列表信息。
	RoomInfoList []*ListRoomsResponseBodyResultRoomInfoList `json:"RoomInfoList,omitempty" xml:"RoomInfoList,omitempty" type:"Repeated"`
	// 该应用的房间总数。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoomsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResult) SetHasMore(v bool) *ListRoomsResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetPageTotal(v int32) *ListRoomsResponseBodyResult {
	s.PageTotal = &v
	return s
}

func (s *ListRoomsResponseBodyResult) SetRoomInfoList(v []*ListRoomsResponseBodyResultRoomInfoList) *ListRoomsResponseBodyResult {
	s.RoomInfoList = v
	return s
}

func (s *ListRoomsResponseBodyResult) SetTotalCount(v int32) *ListRoomsResponseBodyResult {
	s.TotalCount = &v
	return s
}

type ListRoomsResponseBodyResultRoomInfoList struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间创建时间戳，单位：毫秒。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 房间拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 用户在线数。
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 活跃插件列表。
	PluginInstanceInfoList []*ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList `json:"PluginInstanceInfoList,omitempty" xml:"PluginInstanceInfoList,omitempty" type:"Repeated"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 创建房间使用的模板id。
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 房间标题。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 用户访问数。
	Uv *int64 `json:"Uv,omitempty" xml:"Uv,omitempty"`
}

func (s ListRoomsResponseBodyResultRoomInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResultRoomInfoList) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetAppId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.AppId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetCreateTime(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetExtension(v map[string]*string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Extension = v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetNotice(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Notice = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetOnlineCount(v int64) *ListRoomsResponseBodyResultRoomInfoList {
	s.OnlineCount = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetPluginInstanceInfoList(v []*ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) *ListRoomsResponseBodyResultRoomInfoList {
	s.PluginInstanceInfoList = v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetRoomId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.RoomId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetRoomOwnerId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.RoomOwnerId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetTemplateId(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.TemplateId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetTitle(v string) *ListRoomsResponseBodyResultRoomInfoList {
	s.Title = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoList) SetUv(v int64) *ListRoomsResponseBodyResultRoomInfoList {
	s.Uv = &v
	return s
}

type ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList struct {
	// 插件实例创建时间戳，单位：毫秒。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 插件拓展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 插件实例唯一标识。
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// 插件唯一标识，取值：live-直播，wb-白板，chat-聊天，rtc。
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) GoString() string {
	return s.String()
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetCreateTime(v int64) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetExtension(v map[string]*string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.Extension = v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetPluginId(v string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.PluginId = &v
	return s
}

func (s *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList) SetPluginType(v string) *ListRoomsResponseBodyResultRoomInfoListPluginInstanceInfoList {
	s.PluginType = &v
	return s
}

type ListRoomsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoomsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListRoomsResponse) SetHeaders(v map[string]*string) *ListRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListRoomsResponse) SetStatusCode(v int32) *ListRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoomsResponse) SetBody(v *ListRoomsResponseBody) *ListRoomsResponse {
	s.Body = v
	return s
}

type ListSensitiveWordRequest struct {
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSensitiveWordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordRequest) SetAppId(v string) *ListSensitiveWordRequest {
	s.AppId = &v
	return s
}

func (s *ListSensitiveWordRequest) SetPageNum(v int32) *ListSensitiveWordRequest {
	s.PageNum = &v
	return s
}

func (s *ListSensitiveWordRequest) SetPageSize(v int32) *ListSensitiveWordRequest {
	s.PageSize = &v
	return s
}

type ListSensitiveWordResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *ListSensitiveWordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSensitiveWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBody) SetRequestId(v string) *ListSensitiveWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveWordResponseBody) SetResult(v *ListSensitiveWordResponseBodyResult) *ListSensitiveWordResponseBody {
	s.Result = v
	return s
}

type ListSensitiveWordResponseBodyResult struct {
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WordList   []*string `json:"WordList,omitempty" xml:"WordList,omitempty" type:"Repeated"`
}

func (s ListSensitiveWordResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponseBodyResult) SetTotalCount(v int32) *ListSensitiveWordResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListSensitiveWordResponseBodyResult) SetWordList(v []*string) *ListSensitiveWordResponseBodyResult {
	s.WordList = v
	return s
}

type ListSensitiveWordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSensitiveWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSensitiveWordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSensitiveWordResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveWordResponse) SetHeaders(v map[string]*string) *ListSensitiveWordResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveWordResponse) SetStatusCode(v int32) *ListSensitiveWordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitiveWordResponse) SetBody(v *ListSensitiveWordResponseBody) *ListSensitiveWordResponse {
	s.Body = v
	return s
}

type PublishLiveRequest struct {
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 当前用户Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PublishLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRequest) GoString() string {
	return s.String()
}

func (s *PublishLiveRequest) SetLiveId(v string) *PublishLiveRequest {
	s.LiveId = &v
	return s
}

func (s *PublishLiveRequest) SetUserId(v string) *PublishLiveRequest {
	s.UserId = &v
	return s
}

type PublishLiveResponseBody struct {
	// Id of the request
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *PublishLiveResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s PublishLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponseBody) GoString() string {
	return s.String()
}

func (s *PublishLiveResponseBody) SetRequestId(v string) *PublishLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishLiveResponseBody) SetResult(v *PublishLiveResponseBodyResult) *PublishLiveResponseBody {
	s.Result = v
	return s
}

type PublishLiveResponseBodyResult struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 直播推流地址
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// 直播状态：Created-已创建未开播，Living-直播中，End-直播结束
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PublishLiveResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PublishLiveResponseBodyResult) SetAnchorId(v string) *PublishLiveResponseBodyResult {
	s.AnchorId = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetLiveId(v string) *PublishLiveResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetLiveUrl(v string) *PublishLiveResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetPushUrl(v string) *PublishLiveResponseBodyResult {
	s.PushUrl = &v
	return s
}

func (s *PublishLiveResponseBodyResult) SetStatus(v string) *PublishLiveResponseBodyResult {
	s.Status = &v
	return s
}

type PublishLiveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveResponse) GoString() string {
	return s.String()
}

func (s *PublishLiveResponse) SetHeaders(v map[string]*string) *PublishLiveResponse {
	s.Headers = v
	return s
}

func (s *PublishLiveResponse) SetStatusCode(v int32) *PublishLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishLiveResponse) SetBody(v *PublishLiveResponseBody) *PublishLiveResponse {
	s.Body = v
	return s
}

type PublishLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PublishLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomRequest) SetAppId(v string) *PublishLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *PublishLiveRoomRequest) SetLiveId(v string) *PublishLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *PublishLiveRoomRequest) SetUserId(v string) *PublishLiveRoomRequest {
	s.UserId = &v
	return s
}

type PublishLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建场景化直播返回的结果。
	Result *PublishLiveRoomResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s PublishLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomResponseBody) SetRequestId(v string) *PublishLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishLiveRoomResponseBody) SetResult(v *PublishLiveRoomResponseBodyResult) *PublishLiveRoomResponseBody {
	s.Result = v
	return s
}

type PublishLiveRoomResponseBodyResult struct {
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播拉流地址。
	LiveUrl *string `json:"LiveUrl,omitempty" xml:"LiveUrl,omitempty"`
	// 直播推流地址。
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
}

func (s PublishLiveRoomResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomResponseBodyResult) SetLiveId(v string) *PublishLiveRoomResponseBodyResult {
	s.LiveId = &v
	return s
}

func (s *PublishLiveRoomResponseBodyResult) SetLiveUrl(v string) *PublishLiveRoomResponseBodyResult {
	s.LiveUrl = &v
	return s
}

func (s *PublishLiveRoomResponseBodyResult) SetPushUrl(v string) *PublishLiveRoomResponseBodyResult {
	s.PushUrl = &v
	return s
}

type PublishLiveRoomResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *PublishLiveRoomResponse) SetHeaders(v map[string]*string) *PublishLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *PublishLiveRoomResponse) SetStatusCode(v int32) *PublishLiveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishLiveRoomResponse) SetBody(v *PublishLiveRoomResponseBody) *PublishLiveRoomResponse {
	s.Body = v
	return s
}

type RemoveMemberRequest struct {
	// 会议唯一标识
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// 邀请者用户ID
	FromUserId *string `json:"FromUserId,omitempty" xml:"FromUserId,omitempty"`
	// 被邀请用户ID
	ToUserId *string `json:"ToUserId,omitempty" xml:"ToUserId,omitempty"`
}

func (s RemoveMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveMemberRequest) SetConferenceId(v string) *RemoveMemberRequest {
	s.ConferenceId = &v
	return s
}

func (s *RemoveMemberRequest) SetFromUserId(v string) *RemoveMemberRequest {
	s.FromUserId = &v
	return s
}

func (s *RemoveMemberRequest) SetToUserId(v string) *RemoveMemberRequest {
	s.ToUserId = &v
	return s
}

type RemoveMemberResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMemberResponseBody) SetRequestId(v string) *RemoveMemberResponseBody {
	s.RequestId = &v
	return s
}

type RemoveMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveMemberResponse) SetHeaders(v map[string]*string) *RemoveMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveMemberResponse) SetStatusCode(v int32) *RemoveMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMemberResponse) SetBody(v *RemoveMemberResponseBody) *RemoveMemberResponse {
	s.Body = v
	return s
}

type SendCommentRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 发送的文本内容。最大的长度不超过256个字节。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段，服务端仅做透传。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕消息发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
}

func (s SendCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCommentRequest) GoString() string {
	return s.String()
}

func (s *SendCommentRequest) SetAppId(v string) *SendCommentRequest {
	s.AppId = &v
	return s
}

func (s *SendCommentRequest) SetContent(v string) *SendCommentRequest {
	s.Content = &v
	return s
}

func (s *SendCommentRequest) SetExtension(v map[string]*string) *SendCommentRequest {
	s.Extension = v
	return s
}

func (s *SendCommentRequest) SetRoomId(v string) *SendCommentRequest {
	s.RoomId = &v
	return s
}

func (s *SendCommentRequest) SetSenderId(v string) *SendCommentRequest {
	s.SenderId = &v
	return s
}

func (s *SendCommentRequest) SetSenderNick(v string) *SendCommentRequest {
	s.SenderNick = &v
	return s
}

type SendCommentShrinkRequest struct {
	// 应用唯一标识，可以包含小写字母、数字，长度为6个字符。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 发送的文本内容。最大的长度不超过256个字节。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 扩展字段，服务端仅做透传。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播间唯一标识，在调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 弹幕发送者的用户ID，最大长度不超过32个字节。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕消息发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
}

func (s SendCommentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCommentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendCommentShrinkRequest) SetAppId(v string) *SendCommentShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SendCommentShrinkRequest) SetContent(v string) *SendCommentShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendCommentShrinkRequest) SetExtensionShrink(v string) *SendCommentShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *SendCommentShrinkRequest) SetRoomId(v string) *SendCommentShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *SendCommentShrinkRequest) SetSenderId(v string) *SendCommentShrinkRequest {
	s.SenderId = &v
	return s
}

func (s *SendCommentShrinkRequest) SetSenderNick(v string) *SendCommentShrinkRequest {
	s.SenderNick = &v
	return s
}

type SendCommentResponseBody struct {
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用发送直播间弹幕的返回结果。
	Result *SendCommentResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponseBody) GoString() string {
	return s.String()
}

func (s *SendCommentResponseBody) SetRequestId(v string) *SendCommentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCommentResponseBody) SetResult(v *SendCommentResponseBodyResult) *SendCommentResponseBody {
	s.Result = v
	return s
}

type SendCommentResponseBodyResult struct {
	// 返回的弹幕数据模型。
	CommentVO *SendCommentResponseBodyResultCommentVO `json:"CommentVO,omitempty" xml:"CommentVO,omitempty" type:"Struct"`
}

func (s SendCommentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCommentResponseBodyResult) SetCommentVO(v *SendCommentResponseBodyResultCommentVO) *SendCommentResponseBodyResult {
	s.CommentVO = v
	return s
}

type SendCommentResponseBodyResultCommentVO struct {
	// 弹幕的唯一ID。
	CommentId *string `json:"CommentId,omitempty" xml:"CommentId,omitempty"`
	// 弹幕的内容。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 弹幕的创建时间，Unix时间戳，单位：毫秒。
	CreateAt *int64 `json:"CreateAt,omitempty" xml:"CreateAt,omitempty"`
	// 扩展字段。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 弹幕的发送者ID标识。
	SenderId *string `json:"SenderId,omitempty" xml:"SenderId,omitempty"`
	// 弹幕发送者的昵称。
	SenderNick *string `json:"SenderNick,omitempty" xml:"SenderNick,omitempty"`
}

func (s SendCommentResponseBodyResultCommentVO) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponseBodyResultCommentVO) GoString() string {
	return s.String()
}

func (s *SendCommentResponseBodyResultCommentVO) SetCommentId(v string) *SendCommentResponseBodyResultCommentVO {
	s.CommentId = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetContent(v string) *SendCommentResponseBodyResultCommentVO {
	s.Content = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetCreateAt(v int64) *SendCommentResponseBodyResultCommentVO {
	s.CreateAt = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetExtension(v map[string]*string) *SendCommentResponseBodyResultCommentVO {
	s.Extension = v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetSenderId(v string) *SendCommentResponseBodyResultCommentVO {
	s.SenderId = &v
	return s
}

func (s *SendCommentResponseBodyResultCommentVO) SetSenderNick(v string) *SendCommentResponseBodyResultCommentVO {
	s.SenderNick = &v
	return s
}

type SendCommentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendCommentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCommentResponse) GoString() string {
	return s.String()
}

func (s *SendCommentResponse) SetHeaders(v map[string]*string) *SendCommentResponse {
	s.Headers = v
	return s
}

func (s *SendCommentResponse) SetStatusCode(v int32) *SendCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCommentResponse) SetBody(v *SendCommentResponseBody) *SendCommentResponse {
	s.Body = v
	return s
}

type SendCustomMessageToAllRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息体内容。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s SendCustomMessageToAllRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllRequest) SetAppId(v string) *SendCustomMessageToAllRequest {
	s.AppId = &v
	return s
}

func (s *SendCustomMessageToAllRequest) SetBody(v string) *SendCustomMessageToAllRequest {
	s.Body = &v
	return s
}

func (s *SendCustomMessageToAllRequest) SetRoomId(v string) *SendCustomMessageToAllRequest {
	s.RoomId = &v
	return s
}

type SendCustomMessageToAllResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *SendCustomMessageToAllResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCustomMessageToAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllResponseBody) SetRequestId(v string) *SendCustomMessageToAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomMessageToAllResponseBody) SetResult(v *SendCustomMessageToAllResponseBodyResult) *SendCustomMessageToAllResponseBody {
	s.Result = v
	return s
}

type SendCustomMessageToAllResponseBodyResult struct {
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendCustomMessageToAllResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllResponseBodyResult) SetMessageId(v string) *SendCustomMessageToAllResponseBodyResult {
	s.MessageId = &v
	return s
}

type SendCustomMessageToAllResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendCustomMessageToAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomMessageToAllResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToAllResponse) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToAllResponse) SetHeaders(v map[string]*string) *SendCustomMessageToAllResponse {
	s.Headers = v
	return s
}

func (s *SendCustomMessageToAllResponse) SetStatusCode(v int32) *SendCustomMessageToAllResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomMessageToAllResponse) SetBody(v *SendCustomMessageToAllResponseBody) *SendCustomMessageToAllResponse {
	s.Body = v
	return s
}

type SendCustomMessageToUsersRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息体内容。
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// 消息指定的接收人ID列表。
	ReceiverList []*string `json:"ReceiverList,omitempty" xml:"ReceiverList,omitempty" type:"Repeated"`
	// 房间唯一标识，由调用CreateRoom返回。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s SendCustomMessageToUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersRequest) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersRequest) SetAppId(v string) *SendCustomMessageToUsersRequest {
	s.AppId = &v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetBody(v string) *SendCustomMessageToUsersRequest {
	s.Body = &v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetReceiverList(v []*string) *SendCustomMessageToUsersRequest {
	s.ReceiverList = v
	return s
}

func (s *SendCustomMessageToUsersRequest) SetRoomId(v string) *SendCustomMessageToUsersRequest {
	s.RoomId = &v
	return s
}

type SendCustomMessageToUsersResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// API请求的返回结果结构体。
	Result *SendCustomMessageToUsersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s SendCustomMessageToUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersResponseBody) SetRequestId(v string) *SendCustomMessageToUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomMessageToUsersResponseBody) SetResult(v *SendCustomMessageToUsersResponseBodyResult) *SendCustomMessageToUsersResponseBody {
	s.Result = v
	return s
}

type SendCustomMessageToUsersResponseBodyResult struct {
	// 消息的唯一ID标识。由数字、大小写字母组成，长度不超过20。
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendCustomMessageToUsersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersResponseBodyResult) SetMessageId(v string) *SendCustomMessageToUsersResponseBodyResult {
	s.MessageId = &v
	return s
}

type SendCustomMessageToUsersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendCustomMessageToUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCustomMessageToUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomMessageToUsersResponse) GoString() string {
	return s.String()
}

func (s *SendCustomMessageToUsersResponse) SetHeaders(v map[string]*string) *SendCustomMessageToUsersResponse {
	s.Headers = v
	return s
}

func (s *SendCustomMessageToUsersResponse) SetStatusCode(v int32) *SendCustomMessageToUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomMessageToUsersResponse) SetBody(v *SendCustomMessageToUsersResponseBody) *SendCustomMessageToUsersResponse {
	s.Body = v
	return s
}

type SetUserAdminRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 房间唯一标识，由字母、数字、符号.和-组成，最大长度36位。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SetUserAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserAdminRequest) GoString() string {
	return s.String()
}

func (s *SetUserAdminRequest) SetAppId(v string) *SetUserAdminRequest {
	s.AppId = &v
	return s
}

func (s *SetUserAdminRequest) SetRoomId(v string) *SetUserAdminRequest {
	s.RoomId = &v
	return s
}

func (s *SetUserAdminRequest) SetUserId(v string) *SetUserAdminRequest {
	s.UserId = &v
	return s
}

type SetUserAdminResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserAdminResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserAdminResponseBody) SetRequestId(v string) *SetUserAdminResponseBody {
	s.RequestId = &v
	return s
}

type SetUserAdminResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetUserAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetUserAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserAdminResponse) GoString() string {
	return s.String()
}

func (s *SetUserAdminResponse) SetHeaders(v map[string]*string) *SetUserAdminResponse {
	s.Headers = v
	return s
}

func (s *SetUserAdminResponse) SetStatusCode(v int32) *SetUserAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserAdminResponse) SetBody(v *SetUserAdminResponseBody) *SetUserAdminResponse {
	s.Body = v
	return s
}

type StopClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 操作者用户ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopClassRequest) String() string {
	return tea.Prettify(s)
}

func (s StopClassRequest) GoString() string {
	return s.String()
}

func (s *StopClassRequest) SetAppId(v string) *StopClassRequest {
	s.AppId = &v
	return s
}

func (s *StopClassRequest) SetClassId(v string) *StopClassRequest {
	s.ClassId = &v
	return s
}

func (s *StopClassRequest) SetUserId(v string) *StopClassRequest {
	s.UserId = &v
	return s
}

type StopClassResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopClassResponseBody) GoString() string {
	return s.String()
}

func (s *StopClassResponseBody) SetRequestId(v string) *StopClassResponseBody {
	s.RequestId = &v
	return s
}

type StopClassResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopClassResponse) String() string {
	return tea.Prettify(s)
}

func (s StopClassResponse) GoString() string {
	return s.String()
}

func (s *StopClassResponse) SetHeaders(v map[string]*string) *StopClassResponse {
	s.Headers = v
	return s
}

func (s *StopClassResponse) SetStatusCode(v int32) *StopClassResponse {
	s.StatusCode = &v
	return s
}

func (s *StopClassResponse) SetBody(v *StopClassResponseBody) *StopClassResponse {
	s.Body = v
	return s
}

type StopLiveRequest struct {
	// 租户名
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 房间ID，最大长度36位
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 创建直播用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRequest) GoString() string {
	return s.String()
}

func (s *StopLiveRequest) SetAppId(v string) *StopLiveRequest {
	s.AppId = &v
	return s
}

func (s *StopLiveRequest) SetLiveId(v string) *StopLiveRequest {
	s.LiveId = &v
	return s
}

func (s *StopLiveRequest) SetRoomId(v string) *StopLiveRequest {
	s.RoomId = &v
	return s
}

func (s *StopLiveRequest) SetUserId(v string) *StopLiveRequest {
	s.UserId = &v
	return s
}

type StopLiveResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLiveResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveResponseBody) SetRequestId(v string) *StopLiveResponseBody {
	s.RequestId = &v
	return s
}

type StopLiveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveResponse) GoString() string {
	return s.String()
}

func (s *StopLiveResponse) SetHeaders(v map[string]*string) *StopLiveResponse {
	s.Headers = v
	return s
}

func (s *StopLiveResponse) SetStatusCode(v int32) *StopLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveResponse) SetBody(v *StopLiveResponseBody) *StopLiveResponse {
	s.Body = v
	return s
}

type StopLiveRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *StopLiveRoomRequest) SetAppId(v string) *StopLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *StopLiveRoomRequest) SetLiveId(v string) *StopLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *StopLiveRoomRequest) SetUserId(v string) *StopLiveRoomRequest {
	s.UserId = &v
	return s
}

type StopLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveRoomResponseBody) SetRequestId(v string) *StopLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

type StopLiveRoomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *StopLiveRoomResponse) SetHeaders(v map[string]*string) *StopLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *StopLiveRoomResponse) SetStatusCode(v int32) *StopLiveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveRoomResponse) SetBody(v *StopLiveRoomResponseBody) *StopLiveRoomResponse {
	s.Body = v
	return s
}

type UpdateClassRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 创建人用户昵称，1~32个字符。
	CreateNickname *string `json:"CreateNickname,omitempty" xml:"CreateNickname,omitempty"`
	// 创建人用户ID，仅支持中英文数字，下划线，中划线，1~36个字符。
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// 课堂标题，1~32个字符。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateClassRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClassRequest) GoString() string {
	return s.String()
}

func (s *UpdateClassRequest) SetAppId(v string) *UpdateClassRequest {
	s.AppId = &v
	return s
}

func (s *UpdateClassRequest) SetClassId(v string) *UpdateClassRequest {
	s.ClassId = &v
	return s
}

func (s *UpdateClassRequest) SetCreateNickname(v string) *UpdateClassRequest {
	s.CreateNickname = &v
	return s
}

func (s *UpdateClassRequest) SetCreateUserId(v string) *UpdateClassRequest {
	s.CreateUserId = &v
	return s
}

func (s *UpdateClassRequest) SetTitle(v string) *UpdateClassRequest {
	s.Title = &v
	return s
}

type UpdateClassResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClassResponseBody) SetRequestId(v string) *UpdateClassResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClassResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClassResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClassResponse) GoString() string {
	return s.String()
}

func (s *UpdateClassResponse) SetHeaders(v map[string]*string) *UpdateClassResponse {
	s.Headers = v
	return s
}

func (s *UpdateClassResponse) SetStatusCode(v int32) *UpdateClassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClassResponse) SetBody(v *UpdateClassResponseBody) *UpdateClassResponse {
	s.Body = v
	return s
}

type UpdateLiveRequest struct {
	// 直播简介，支持中英文，最大长度2048位
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// 直播资源的唯一标识ID
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 直播标题，支持中英文，最大长度256位
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequest) SetIntroduction(v string) *UpdateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateLiveRequest) SetLiveId(v string) *UpdateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRequest) SetTitle(v string) *UpdateLiveRequest {
	s.Title = &v
	return s
}

type UpdateLiveResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBody) SetRequestId(v string) *UpdateLiveResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLiveResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponse) SetHeaders(v map[string]*string) *UpdateLiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveResponse) SetStatusCode(v int32) *UpdateLiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveResponse) SetBody(v *UpdateLiveResponseBody) *UpdateLiveResponse {
	s.Body = v
	return s
}

type UpdateLiveRoomRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomRequest) SetAnchorId(v string) *UpdateLiveRoomRequest {
	s.AnchorId = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetAnchorNick(v string) *UpdateLiveRoomRequest {
	s.AnchorNick = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetAppId(v string) *UpdateLiveRoomRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetCoverUrl(v string) *UpdateLiveRoomRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetExtension(v map[string]*string) *UpdateLiveRoomRequest {
	s.Extension = v
	return s
}

func (s *UpdateLiveRoomRequest) SetLiveId(v string) *UpdateLiveRoomRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetNotice(v string) *UpdateLiveRoomRequest {
	s.Notice = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetTitle(v string) *UpdateLiveRoomRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRoomRequest) SetUserId(v string) *UpdateLiveRoomRequest {
	s.UserId = &v
	return s
}

type UpdateLiveRoomShrinkRequest struct {
	// 主播id，仅支持英文和数字，最大长度36位。
	AnchorId *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	// 主播昵称。
	AnchorNick *string `json:"AnchorNick,omitempty" xml:"AnchorNick,omitempty"`
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 封面，支持图片地址链接格式
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 直播ID。
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// 公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 操作人ID。
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomShrinkRequest) SetAnchorId(v string) *UpdateLiveRoomShrinkRequest {
	s.AnchorId = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetAnchorNick(v string) *UpdateLiveRoomShrinkRequest {
	s.AnchorNick = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetAppId(v string) *UpdateLiveRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetCoverUrl(v string) *UpdateLiveRoomShrinkRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetExtensionShrink(v string) *UpdateLiveRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetLiveId(v string) *UpdateLiveRoomShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetNotice(v string) *UpdateLiveRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetTitle(v string) *UpdateLiveRoomShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRoomShrinkRequest) SetUserId(v string) *UpdateLiveRoomShrinkRequest {
	s.UserId = &v
	return s
}

type UpdateLiveRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomResponseBody) SetRequestId(v string) *UpdateLiveRoomResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLiveRoomResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateLiveRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveRoomResponse) SetHeaders(v map[string]*string) *UpdateLiveRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveRoomResponse) SetStatusCode(v int32) *UpdateLiveRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveRoomResponse) SetBody(v *UpdateLiveRoomResponseBody) *UpdateLiveRoomResponse {
	s.Body = v
	return s
}

type UpdateRoomRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoomRequest) SetAppId(v string) *UpdateRoomRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRoomRequest) SetExtension(v map[string]*string) *UpdateRoomRequest {
	s.Extension = v
	return s
}

func (s *UpdateRoomRequest) SetNotice(v string) *UpdateRoomRequest {
	s.Notice = &v
	return s
}

func (s *UpdateRoomRequest) SetRoomId(v string) *UpdateRoomRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateRoomRequest) SetRoomOwnerId(v string) *UpdateRoomRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *UpdateRoomRequest) SetTitle(v string) *UpdateRoomRequest {
	s.Title = &v
	return s
}

type UpdateRoomShrinkRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 拓展字段，按需传递，需要额外记录的房间属性。
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 房间公告，支持中英文，最大长度256位。
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// 房间唯一标识。
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	// 房主用户id，仅支持英文和数字，最大长度36位。
	RoomOwnerId *string `json:"RoomOwnerId,omitempty" xml:"RoomOwnerId,omitempty"`
	// 房间标题，支持中英文，最大长度32位。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateRoomShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoomShrinkRequest) SetAppId(v string) *UpdateRoomShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetExtensionShrink(v string) *UpdateRoomShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetNotice(v string) *UpdateRoomShrinkRequest {
	s.Notice = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetRoomId(v string) *UpdateRoomShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetRoomOwnerId(v string) *UpdateRoomShrinkRequest {
	s.RoomOwnerId = &v
	return s
}

func (s *UpdateRoomShrinkRequest) SetTitle(v string) *UpdateRoomShrinkRequest {
	s.Title = &v
	return s
}

type UpdateRoomResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRoomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoomResponseBody) SetRequestId(v string) *UpdateRoomResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRoomResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoomResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoomResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoomResponse) SetHeaders(v map[string]*string) *UpdateRoomResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoomResponse) SetStatusCode(v int32) *UpdateRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoomResponse) SetBody(v *UpdateRoomResponseBody) *UpdateRoomResponse {
	s.Body = v
	return s
}

type UpdateShareScreenLayoutRequest struct {
	// 应用唯一标识，由6位小写字母、数字组成。
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 课堂唯一标识，由调用CreateClass返回。
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// 是否开启叠加老师画面
	EnableOverlay *bool `json:"EnableOverlay,omitempty" xml:"EnableOverlay,omitempty"`
	// 叠加画面高度，归一化为1
	OverlayHeight *float32 `json:"OverlayHeight,omitempty" xml:"OverlayHeight,omitempty"`
	// 叠加画面宽度，归一化为1
	OverlayWidth *float32 `json:"OverlayWidth,omitempty" xml:"OverlayWidth,omitempty"`
	// 叠加画面X坐标，归一化为1
	OverlayX *float32 `json:"OverlayX,omitempty" xml:"OverlayX,omitempty"`
	// 叠加画面Y坐标，归一化为1
	OverlayY *float32 `json:"OverlayY,omitempty" xml:"OverlayY,omitempty"`
}

func (s UpdateShareScreenLayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateShareScreenLayoutRequest) GoString() string {
	return s.String()
}

func (s *UpdateShareScreenLayoutRequest) SetAppId(v string) *UpdateShareScreenLayoutRequest {
	s.AppId = &v
	return s
}

func (s *UpdateShareScreenLayoutRequest) SetClassId(v string) *UpdateShareScreenLayoutRequest {
	s.ClassId = &v
	return s
}

func (s *UpdateShareScreenLayoutRequest) SetEnableOverlay(v bool) *UpdateShareScreenLayoutRequest {
	s.EnableOverlay = &v
	return s
}

func (s *UpdateShareScreenLayoutRequest) SetOverlayHeight(v float32) *UpdateShareScreenLayoutRequest {
	s.OverlayHeight = &v
	return s
}

func (s *UpdateShareScreenLayoutRequest) SetOverlayWidth(v float32) *UpdateShareScreenLayoutRequest {
	s.OverlayWidth = &v
	return s
}

func (s *UpdateShareScreenLayoutRequest) SetOverlayX(v float32) *UpdateShareScreenLayoutRequest {
	s.OverlayX = &v
	return s
}

func (s *UpdateShareScreenLayoutRequest) SetOverlayY(v float32) *UpdateShareScreenLayoutRequest {
	s.OverlayY = &v
	return s
}

type UpdateShareScreenLayoutResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateShareScreenLayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateShareScreenLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateShareScreenLayoutResponseBody) SetRequestId(v string) *UpdateShareScreenLayoutResponseBody {
	s.RequestId = &v
	return s
}

type UpdateShareScreenLayoutResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateShareScreenLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateShareScreenLayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateShareScreenLayoutResponse) GoString() string {
	return s.String()
}

func (s *UpdateShareScreenLayoutResponse) SetHeaders(v map[string]*string) *UpdateShareScreenLayoutResponse {
	s.Headers = v
	return s
}

func (s *UpdateShareScreenLayoutResponse) SetStatusCode(v int32) *UpdateShareScreenLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateShareScreenLayoutResponse) SetBody(v *UpdateShareScreenLayoutResponseBody) *UpdateShareScreenLayoutResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BanAllCommentWithOptions(request *BanAllCommentRequest, runtime *util.RuntimeOptions) (_result *BanAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BanAllComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BanAllCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BanAllComment(request *BanAllCommentRequest) (_result *BanAllCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BanAllCommentResponse{}
	_body, _err := client.BanAllCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BanCommentWithOptions(request *BanCommentRequest, runtime *util.RuntimeOptions) (_result *BanCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BanCommentTime)) {
		body["BanCommentTime"] = request.BanCommentTime
	}

	if !tea.BoolValue(util.IsUnset(request.BanCommentUser)) {
		body["BanCommentUser"] = request.BanCommentUser
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BanComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BanCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BanComment(request *BanCommentRequest) (_result *BanCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BanCommentResponse{}
	_body, _err := client.BanCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelBanAllCommentWithOptions(request *CancelBanAllCommentRequest, runtime *util.RuntimeOptions) (_result *CancelBanAllCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelBanAllComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelBanAllCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelBanAllComment(request *CancelBanAllCommentRequest) (_result *CancelBanAllCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelBanAllCommentResponse{}
	_body, _err := client.CancelBanAllCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelBanCommentWithOptions(request *CancelBanCommentRequest, runtime *util.RuntimeOptions) (_result *CancelBanCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BanCommentUser)) {
		body["BanCommentUser"] = request.BanCommentUser
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelBanComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelBanCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelBanComment(request *CancelBanCommentRequest) (_result *CancelBanCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelBanCommentResponse{}
	_body, _err := client.CancelBanCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelUserAdminWithOptions(request *CancelUserAdminRequest, runtime *util.RuntimeOptions) (_result *CancelUserAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelUserAdmin"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelUserAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelUserAdmin(request *CancelUserAdminRequest) (_result *CancelUserAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelUserAdminResponse{}
	_body, _err := client.CancelUserAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClassWithOptions(request *CreateClassRequest, runtime *util.RuntimeOptions) (_result *CreateClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateNickname)) {
		body["CreateNickname"] = request.CreateNickname
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClass(request *CreateClassRequest) (_result *CreateClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClassResponse{}
	_body, _err := client.CreateClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveWithOptions(request *CreateLiveRequest, runtime *util.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		body["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeLevel)) {
		body["CodeLevel"] = request.CodeLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLive(request *CreateLiveRequest) (_result *CreateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveResponse{}
	_body, _err := client.CreateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveRecordSliceFileWithOptions(request *CreateLiveRecordSliceFileRequest, runtime *util.RuntimeOptions) (_result *CreateLiveRecordSliceFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLiveRecordSliceFile"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveRecordSliceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLiveRecordSliceFile(request *CreateLiveRecordSliceFileRequest) (_result *CreateLiveRecordSliceFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveRecordSliceFileResponse{}
	_body, _err := client.CreateLiveRecordSliceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveRoomWithOptions(tmpReq *CreateLiveRoomRequest, runtime *util.RuntimeOptions) (_result *CreateLiveRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateLiveRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		body["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.AnchorNick)) {
		body["AnchorNick"] = request.AnchorNick
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["CoverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLinkMic)) {
		body["EnableLinkMic"] = request.EnableLinkMic
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLiveRoom(request *CreateLiveRoomRequest) (_result *CreateLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveRoomResponse{}
	_body, _err := client.CreateLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoomWithOptions(tmpReq *CreateRoomRequest, runtime *util.RuntimeOptions) (_result *CreateRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomOwnerId)) {
		body["RoomOwnerId"] = request.RoomOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRoom(request *CreateRoomRequest) (_result *CreateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoomResponse{}
	_body, _err := client.CreateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSensitiveWordWithOptions(tmpReq *CreateSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *CreateSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSensitiveWordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.WordList)) {
		request.WordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WordList, tea.String("WordList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.WordListShrink)) {
		body["WordList"] = request.WordListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSensitiveWord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSensitiveWord(request *CreateSensitiveWordRequest) (_result *CreateSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSensitiveWordResponse{}
	_body, _err := client.CreateSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClassWithOptions(request *DeleteClassRequest, runtime *util.RuntimeOptions) (_result *DeleteClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClass(request *DeleteClassRequest) (_result *DeleteClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClassResponse{}
	_body, _err := client.DeleteClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommentWithOptions(request *DeleteCommentRequest, runtime *util.RuntimeOptions) (_result *DeleteCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentIdList)) {
		bodyFlat["CommentIdList"] = request.CommentIdList
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComment(request *DeleteCommentRequest) (_result *DeleteCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCommentResponse{}
	_body, _err := client.DeleteCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommentByCreatorIdWithOptions(request *DeleteCommentByCreatorIdRequest, runtime *util.RuntimeOptions) (_result *DeleteCommentByCreatorIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentIdList)) {
		bodyFlat["CommentIdList"] = request.CommentIdList
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCommentByCreatorId"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCommentByCreatorIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCommentByCreatorId(request *DeleteCommentByCreatorIdRequest) (_result *DeleteCommentByCreatorIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCommentByCreatorIdResponse{}
	_body, _err := client.DeleteCommentByCreatorIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConferenceWithOptions(request *DeleteConferenceRequest, runtime *util.RuntimeOptions) (_result *DeleteConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConference"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConference(request *DeleteConferenceRequest) (_result *DeleteConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConferenceResponse{}
	_body, _err := client.DeleteConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveWithOptions(request *DeleteLiveRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLive(request *DeleteLiveRequest) (_result *DeleteLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveResponse{}
	_body, _err := client.DeleteLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveRoomWithOptions(request *DeleteLiveRoomRequest, runtime *util.RuntimeOptions) (_result *DeleteLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveRoom(request *DeleteLiveRoomRequest) (_result *DeleteLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLiveRoomResponse{}
	_body, _err := client.DeleteLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecordFileInfoWithOptions(request *DeleteRecordFileInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordFileInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecordFileInfo"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecordFileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecordFileInfo(request *DeleteRecordFileInfoRequest) (_result *DeleteRecordFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordFileInfoResponse{}
	_body, _err := client.DeleteRecordFileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoomWithOptions(request *DeleteRoomRequest, runtime *util.RuntimeOptions) (_result *DeleteRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRoom(request *DeleteRoomRequest) (_result *DeleteRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoomResponse{}
	_body, _err := client.DeleteRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSensitiveWordWithOptions(tmpReq *DeleteSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *DeleteSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteSensitiveWordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.WordList)) {
		request.WordListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WordList, tea.String("WordList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.WordListShrink)) {
		body["WordList"] = request.WordListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSensitiveWord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSensitiveWord(request *DeleteSensitiveWordRequest) (_result *DeleteSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSensitiveWordResponse{}
	_body, _err := client.DeleteSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterImpPlayBackTimeByLiveIdWithOptions(request *DescribeMeterImpPlayBackTimeByLiveIdRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterImpPlayBackTimeByLiveIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterImpPlayBackTimeByLiveId"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterImpPlayBackTimeByLiveIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterImpPlayBackTimeByLiveId(request *DescribeMeterImpPlayBackTimeByLiveIdRequest) (_result *DescribeMeterImpPlayBackTimeByLiveIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterImpPlayBackTimeByLiveIdResponse{}
	_body, _err := client.DescribeMeterImpPlayBackTimeByLiveIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMeterImpWatchLiveTimeByLiveIdWithOptions(request *DescribeMeterImpWatchLiveTimeByLiveIdRequest, runtime *util.RuntimeOptions) (_result *DescribeMeterImpWatchLiveTimeByLiveIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMeterImpWatchLiveTimeByLiveId"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMeterImpWatchLiveTimeByLiveIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMeterImpWatchLiveTimeByLiveId(request *DescribeMeterImpWatchLiveTimeByLiveIdRequest) (_result *DescribeMeterImpWatchLiveTimeByLiveIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMeterImpWatchLiveTimeByLiveIdResponse{}
	_body, _err := client.DescribeMeterImpWatchLiveTimeByLiveIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthToken"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (_result *GetAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.GetAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClassDetailWithOptions(request *GetClassDetailRequest, runtime *util.RuntimeOptions) (_result *GetClassDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClassDetail"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClassDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClassDetail(request *GetClassDetailRequest) (_result *GetClassDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClassDetailResponse{}
	_body, _err := client.GetClassDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClassRecordWithOptions(request *GetClassRecordRequest, runtime *util.RuntimeOptions) (_result *GetClassRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClassRecord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClassRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClassRecord(request *GetClassRecordRequest) (_result *GetClassRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClassRecordResponse{}
	_body, _err := client.GetClassRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConferenceWithOptions(request *GetConferenceRequest, runtime *util.RuntimeOptions) (_result *GetConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConference"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConference(request *GetConferenceRequest) (_result *GetConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConferenceResponse{}
	_body, _err := client.GetConferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveWithOptions(request *GetLiveRequest, runtime *util.RuntimeOptions) (_result *GetLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLive(request *GetLiveRequest) (_result *GetLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveResponse{}
	_body, _err := client.GetLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRecordWithOptions(request *GetLiveRecordRequest, runtime *util.RuntimeOptions) (_result *GetLiveRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRecord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRecord(request *GetLiveRecordRequest) (_result *GetLiveRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRecordResponse{}
	_body, _err := client.GetLiveRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRoomWithOptions(request *GetLiveRoomRequest, runtime *util.RuntimeOptions) (_result *GetLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRoom(request *GetLiveRoomRequest) (_result *GetLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRoomResponse{}
	_body, _err := client.GetLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRoomStatisticsWithOptions(request *GetLiveRoomStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetLiveRoomStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRoomStatistics"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRoomStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRoomStatistics(request *GetLiveRoomStatisticsRequest) (_result *GetLiveRoomStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRoomStatisticsResponse{}
	_body, _err := client.GetLiveRoomStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveRoomUserStatisticsWithOptions(request *GetLiveRoomUserStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetLiveRoomUserStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveRoomUserStatistics"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLiveRoomUserStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveRoomUserStatistics(request *GetLiveRoomUserStatisticsRequest) (_result *GetLiveRoomUserStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLiveRoomUserStatisticsResponse{}
	_body, _err := client.GetLiveRoomUserStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordFileInfoWithOptions(request *GetRecordFileInfoRequest, runtime *util.RuntimeOptions) (_result *GetRecordFileInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecordFileInfo"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordFileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecordFileInfo(request *GetRecordFileInfoRequest) (_result *GetRecordFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordFileInfoResponse{}
	_body, _err := client.GetRecordFileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoomWithOptions(request *GetRoomRequest, runtime *util.RuntimeOptions) (_result *GetRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoom(request *GetRoomRequest) (_result *GetRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRoomResponse{}
	_body, _err := client.GetRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStandardRoomJumpUrlWithOptions(request *GetStandardRoomJumpUrlRequest, runtime *util.RuntimeOptions) (_result *GetStandardRoomJumpUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserNick)) {
		body["UserNick"] = request.UserNick
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStandardRoomJumpUrl"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStandardRoomJumpUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStandardRoomJumpUrl(request *GetStandardRoomJumpUrlRequest) (_result *GetStandardRoomJumpUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStandardRoomJumpUrlResponse{}
	_body, _err := client.GetStandardRoomJumpUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KickRoomUserWithOptions(request *KickRoomUserRequest, runtime *util.RuntimeOptions) (_result *KickRoomUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BlockTime)) {
		body["BlockTime"] = request.BlockTime
	}

	if !tea.BoolValue(util.IsUnset(request.KickUser)) {
		body["KickUser"] = request.KickUser
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KickRoomUser"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KickRoomUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KickRoomUser(request *KickRoomUserRequest) (_result *KickRoomUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KickRoomUserResponse{}
	_body, _err := client.KickRoomUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClassesWithOptions(request *ListClassesRequest, runtime *util.RuntimeOptions) (_result *ListClassesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClasses"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClassesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClasses(request *ListClassesRequest) (_result *ListClassesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClassesResponse{}
	_body, _err := client.ListClassesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommentsWithOptions(request *ListCommentsRequest, runtime *util.RuntimeOptions) (_result *ListCommentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		body["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListComments"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComments(request *ListCommentsRequest) (_result *ListCommentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommentsResponse{}
	_body, _err := client.ListCommentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConferenceUsersWithOptions(request *ListConferenceUsersRequest, runtime *util.RuntimeOptions) (_result *ListConferenceUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConferenceUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConferenceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConferenceUsers(request *ListConferenceUsersRequest) (_result *ListConferenceUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConferenceUsersResponse{}
	_body, _err := client.ListConferenceUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveFilesWithOptions(request *ListLiveFilesRequest, runtime *util.RuntimeOptions) (_result *ListLiveFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		query["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLiveFiles"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLiveFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveFiles(request *ListLiveFilesRequest) (_result *ListLiveFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveFilesResponse{}
	_body, _err := client.ListLiveFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRoomsWithOptions(request *ListLiveRoomsRequest, runtime *util.RuntimeOptions) (_result *ListLiveRoomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLiveRooms"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLiveRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRooms(request *ListLiveRoomsRequest) (_result *ListLiveRoomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRoomsResponse{}
	_body, _err := client.ListLiveRoomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRoomsByIdWithOptions(tmpReq *ListLiveRoomsByIdRequest, runtime *util.RuntimeOptions) (_result *ListLiveRoomsByIdResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListLiveRoomsByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LiveIdList)) {
		request.LiveIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LiveIdList, tea.String("LiveIdList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveIdListShrink)) {
		body["LiveIdList"] = request.LiveIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLiveRoomsById"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLiveRoomsByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRoomsById(request *ListLiveRoomsByIdRequest) (_result *ListLiveRoomsByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRoomsByIdResponse{}
	_body, _err := client.ListLiveRoomsByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoomUsersWithOptions(request *ListRoomUsersRequest, runtime *util.RuntimeOptions) (_result *ListRoomUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoomUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoomUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoomUsers(request *ListRoomUsersRequest) (_result *ListRoomUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoomUsersResponse{}
	_body, _err := client.ListRoomUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoomsWithOptions(request *ListRoomsRequest, runtime *util.RuntimeOptions) (_result *ListRoomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRooms"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoomsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRooms(request *ListRoomsRequest) (_result *ListRoomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoomsResponse{}
	_body, _err := client.ListRoomsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSensitiveWordWithOptions(request *ListSensitiveWordRequest, runtime *util.RuntimeOptions) (_result *ListSensitiveWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSensitiveWord"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSensitiveWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSensitiveWord(request *ListSensitiveWordRequest) (_result *ListSensitiveWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSensitiveWordResponse{}
	_body, _err := client.ListSensitiveWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishLiveWithOptions(request *PublishLiveRequest, runtime *util.RuntimeOptions) (_result *PublishLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishLive(request *PublishLiveRequest) (_result *PublishLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishLiveResponse{}
	_body, _err := client.PublishLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishLiveRoomWithOptions(request *PublishLiveRoomRequest, runtime *util.RuntimeOptions) (_result *PublishLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishLiveRoom(request *PublishLiveRoomRequest) (_result *PublishLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishLiveRoomResponse{}
	_body, _err := client.PublishLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMemberWithOptions(request *RemoveMemberRequest, runtime *util.RuntimeOptions) (_result *RemoveMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceId)) {
		body["ConferenceId"] = request.ConferenceId
	}

	if !tea.BoolValue(util.IsUnset(request.FromUserId)) {
		body["FromUserId"] = request.FromUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserId)) {
		body["ToUserId"] = request.ToUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveMember"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMember(request *RemoveMemberRequest) (_result *RemoveMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveMemberResponse{}
	_body, _err := client.RemoveMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCommentWithOptions(tmpReq *SendCommentRequest, runtime *util.RuntimeOptions) (_result *SendCommentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendCommentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["SenderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNick)) {
		body["SenderNick"] = request.SenderNick
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendComment"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCommentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendComment(request *SendCommentRequest) (_result *SendCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCommentResponse{}
	_body, _err := client.SendCommentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCustomMessageToAllWithOptions(request *SendCustomMessageToAllRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomMessageToAll"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCustomMessageToAllResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomMessageToAll(request *SendCustomMessageToAllRequest) (_result *SendCustomMessageToAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomMessageToAllResponse{}
	_body, _err := client.SendCustomMessageToAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCustomMessageToUsersWithOptions(request *SendCustomMessageToUsersRequest, runtime *util.RuntimeOptions) (_result *SendCustomMessageToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReceiverList)) {
		bodyFlat["ReceiverList"] = request.ReceiverList
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomMessageToUsers"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCustomMessageToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCustomMessageToUsers(request *SendCustomMessageToUsersRequest) (_result *SendCustomMessageToUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomMessageToUsersResponse{}
	_body, _err := client.SendCustomMessageToUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetUserAdminWithOptions(request *SetUserAdminRequest, runtime *util.RuntimeOptions) (_result *SetUserAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserAdmin"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetUserAdmin(request *SetUserAdminRequest) (_result *SetUserAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserAdminResponse{}
	_body, _err := client.SetUserAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopClassWithOptions(request *StopClassRequest, runtime *util.RuntimeOptions) (_result *StopClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopClass(request *StopClassRequest) (_result *StopClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopClassResponse{}
	_body, _err := client.StopClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLiveWithOptions(request *StopLiveRequest, runtime *util.RuntimeOptions) (_result *StopLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLive(request *StopLiveRequest) (_result *StopLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveResponse{}
	_body, _err := client.StopLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLiveRoomWithOptions(request *StopLiveRoomRequest, runtime *util.RuntimeOptions) (_result *StopLiveRoomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLiveRoom(request *StopLiveRoomRequest) (_result *StopLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLiveRoomResponse{}
	_body, _err := client.StopLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClassWithOptions(request *UpdateClassRequest, runtime *util.RuntimeOptions) (_result *UpdateClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateNickname)) {
		body["CreateNickname"] = request.CreateNickname
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["CreateUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClass"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClass(request *UpdateClassRequest) (_result *UpdateClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClassResponse{}
	_body, _err := client.UpdateClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveWithOptions(request *UpdateLiveRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Introduction)) {
		body["Introduction"] = request.Introduction
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLive"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLive(request *UpdateLiveRequest) (_result *UpdateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveResponse{}
	_body, _err := client.UpdateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveRoomWithOptions(tmpReq *UpdateLiveRoomRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLiveRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		body["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.AnchorNick)) {
		body["AnchorNick"] = request.AnchorNick
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CoverUrl)) {
		body["CoverUrl"] = request.CoverUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LiveId)) {
		body["LiveId"] = request.LiveId
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLiveRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveRoom(request *UpdateLiveRoomRequest) (_result *UpdateLiveRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveRoomResponse{}
	_body, _err := client.UpdateLiveRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoomWithOptions(tmpReq *UpdateRoomRequest, runtime *util.RuntimeOptions) (_result *UpdateRoomResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRoomShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extension)) {
		request.ExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extension, tea.String("Extension"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtensionShrink)) {
		body["Extension"] = request.ExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Notice)) {
		body["Notice"] = request.Notice
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		body["RoomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.RoomOwnerId)) {
		body["RoomOwnerId"] = request.RoomOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoom"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoomResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRoom(request *UpdateRoomRequest) (_result *UpdateRoomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoomResponse{}
	_body, _err := client.UpdateRoomWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateShareScreenLayoutWithOptions(request *UpdateShareScreenLayoutRequest, runtime *util.RuntimeOptions) (_result *UpdateShareScreenLayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ClassId)) {
		body["ClassId"] = request.ClassId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOverlay)) {
		body["EnableOverlay"] = request.EnableOverlay
	}

	if !tea.BoolValue(util.IsUnset(request.OverlayHeight)) {
		body["OverlayHeight"] = request.OverlayHeight
	}

	if !tea.BoolValue(util.IsUnset(request.OverlayWidth)) {
		body["OverlayWidth"] = request.OverlayWidth
	}

	if !tea.BoolValue(util.IsUnset(request.OverlayX)) {
		body["OverlayX"] = request.OverlayX
	}

	if !tea.BoolValue(util.IsUnset(request.OverlayY)) {
		body["OverlayY"] = request.OverlayY
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateShareScreenLayout"),
		Version:     tea.String("2021-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateShareScreenLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateShareScreenLayout(request *UpdateShareScreenLayoutRequest) (_result *UpdateShareScreenLayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateShareScreenLayoutResponse{}
	_body, _err := client.UpdateShareScreenLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
