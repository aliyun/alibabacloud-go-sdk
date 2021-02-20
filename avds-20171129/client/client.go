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

type AddAssetsRequest struct {
	SourceIp           *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId    *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartAfterVerified *bool     `json:"StartAfterVerified,omitempty" xml:"StartAfterVerified,omitempty"`
	AssetGroupId       *string   `json:"AssetGroupId,omitempty" xml:"AssetGroupId,omitempty"`
	Assets             []*string `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	Tags               []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAssetsRequest) GoString() string {
	return s.String()
}

func (s *AddAssetsRequest) SetSourceIp(v string) *AddAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAssetsRequest) SetResourceOwnerId(v int64) *AddAssetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAssetsRequest) SetStartAfterVerified(v bool) *AddAssetsRequest {
	s.StartAfterVerified = &v
	return s
}

func (s *AddAssetsRequest) SetAssetGroupId(v string) *AddAssetsRequest {
	s.AssetGroupId = &v
	return s
}

func (s *AddAssetsRequest) SetAssets(v []*string) *AddAssetsRequest {
	s.Assets = v
	return s
}

func (s *AddAssetsRequest) SetTags(v []*string) *AddAssetsRequest {
	s.Tags = v
	return s
}

type AddAssetsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *AddAssetsResponseBody) SetRequestId(v string) *AddAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAssetsResponseBody) SetSuccess(v bool) *AddAssetsResponseBody {
	s.Success = &v
	return s
}

type AddAssetsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAssetsResponse) GoString() string {
	return s.String()
}

func (s *AddAssetsResponse) SetHeaders(v map[string]*string) *AddAssetsResponse {
	s.Headers = v
	return s
}

func (s *AddAssetsResponse) SetBody(v *AddAssetsResponseBody) *AddAssetsResponse {
	s.Body = v
	return s
}

type AddAssetTagsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Assets          []*string `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	Tags            []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddAssetTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAssetTagsRequest) GoString() string {
	return s.String()
}

func (s *AddAssetTagsRequest) SetSourceIp(v string) *AddAssetTagsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAssetTagsRequest) SetResourceOwnerId(v int64) *AddAssetTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAssetTagsRequest) SetLang(v string) *AddAssetTagsRequest {
	s.Lang = &v
	return s
}

func (s *AddAssetTagsRequest) SetAssets(v []*string) *AddAssetTagsRequest {
	s.Assets = v
	return s
}

func (s *AddAssetTagsRequest) SetTags(v []*string) *AddAssetTagsRequest {
	s.Tags = v
	return s
}

type AddAssetTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAssetTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAssetTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddAssetTagsResponseBody) SetRequestId(v string) *AddAssetTagsResponseBody {
	s.RequestId = &v
	return s
}

type AddAssetTagsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAssetTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAssetTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAssetTagsResponse) GoString() string {
	return s.String()
}

func (s *AddAssetTagsResponse) SetHeaders(v map[string]*string) *AddAssetTagsResponse {
	s.Headers = v
	return s
}

func (s *AddAssetTagsResponse) SetBody(v *AddAssetTagsResponseBody) *AddAssetTagsResponse {
	s.Body = v
	return s
}

type AddOrgDomainsRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId    *int32    `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Domains  []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s AddOrgDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgDomainsRequest) GoString() string {
	return s.String()
}

func (s *AddOrgDomainsRequest) SetSourceIp(v string) *AddOrgDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddOrgDomainsRequest) SetOrgId(v int32) *AddOrgDomainsRequest {
	s.OrgId = &v
	return s
}

func (s *AddOrgDomainsRequest) SetDomains(v []*string) *AddOrgDomainsRequest {
	s.Domains = v
	return s
}

type AddOrgDomainsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddOrgDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrgDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrgDomainsResponseBody) SetRequestId(v string) *AddOrgDomainsResponseBody {
	s.RequestId = &v
	return s
}

type AddOrgDomainsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOrgDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOrgDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrgDomainsResponse) GoString() string {
	return s.String()
}

func (s *AddOrgDomainsResponse) SetHeaders(v map[string]*string) *AddOrgDomainsResponse {
	s.Headers = v
	return s
}

func (s *AddOrgDomainsResponse) SetBody(v *AddOrgDomainsResponseBody) *AddOrgDomainsResponse {
	s.Body = v
	return s
}

type AddOrgHostsRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId    *int32    `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Hosts    []*string `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s AddOrgHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgHostsRequest) GoString() string {
	return s.String()
}

func (s *AddOrgHostsRequest) SetSourceIp(v string) *AddOrgHostsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddOrgHostsRequest) SetOrgId(v int32) *AddOrgHostsRequest {
	s.OrgId = &v
	return s
}

func (s *AddOrgHostsRequest) SetHosts(v []*string) *AddOrgHostsRequest {
	s.Hosts = v
	return s
}

type AddOrgHostsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddOrgHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrgHostsResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrgHostsResponseBody) SetRequestId(v string) *AddOrgHostsResponseBody {
	s.RequestId = &v
	return s
}

type AddOrgHostsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOrgHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOrgHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrgHostsResponse) GoString() string {
	return s.String()
}

func (s *AddOrgHostsResponse) SetHeaders(v map[string]*string) *AddOrgHostsResponse {
	s.Headers = v
	return s
}

func (s *AddOrgHostsResponse) SetBody(v *AddOrgHostsResponseBody) *AddOrgHostsResponse {
	s.Body = v
	return s
}

type AddOrgSubdomainsRequest struct {
	SourceIp   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId      *int32    `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Subdomains []*string `json:"Subdomains,omitempty" xml:"Subdomains,omitempty" type:"Repeated"`
}

func (s AddOrgSubdomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgSubdomainsRequest) GoString() string {
	return s.String()
}

func (s *AddOrgSubdomainsRequest) SetSourceIp(v string) *AddOrgSubdomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddOrgSubdomainsRequest) SetOrgId(v int32) *AddOrgSubdomainsRequest {
	s.OrgId = &v
	return s
}

func (s *AddOrgSubdomainsRequest) SetSubdomains(v []*string) *AddOrgSubdomainsRequest {
	s.Subdomains = v
	return s
}

type AddOrgSubdomainsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddOrgSubdomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrgSubdomainsResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrgSubdomainsResponseBody) SetRequestId(v string) *AddOrgSubdomainsResponseBody {
	s.RequestId = &v
	return s
}

type AddOrgSubdomainsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOrgSubdomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOrgSubdomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrgSubdomainsResponse) GoString() string {
	return s.String()
}

func (s *AddOrgSubdomainsResponse) SetHeaders(v map[string]*string) *AddOrgSubdomainsResponse {
	s.Headers = v
	return s
}

func (s *AddOrgSubdomainsResponse) SetBody(v *AddOrgSubdomainsResponseBody) *AddOrgSubdomainsResponse {
	s.Body = v
	return s
}

type AddOrgWebPathsRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId    *int32    `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	URLs     []*string `json:"URLs,omitempty" xml:"URLs,omitempty" type:"Repeated"`
}

func (s AddOrgWebPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgWebPathsRequest) GoString() string {
	return s.String()
}

func (s *AddOrgWebPathsRequest) SetSourceIp(v string) *AddOrgWebPathsRequest {
	s.SourceIp = &v
	return s
}

func (s *AddOrgWebPathsRequest) SetOrgId(v int32) *AddOrgWebPathsRequest {
	s.OrgId = &v
	return s
}

func (s *AddOrgWebPathsRequest) SetURLs(v []*string) *AddOrgWebPathsRequest {
	s.URLs = v
	return s
}

type AddOrgWebPathsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddOrgWebPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrgWebPathsResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrgWebPathsResponseBody) SetRequestId(v string) *AddOrgWebPathsResponseBody {
	s.RequestId = &v
	return s
}

type AddOrgWebPathsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddOrgWebPathsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOrgWebPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrgWebPathsResponse) GoString() string {
	return s.String()
}

func (s *AddOrgWebPathsResponse) SetHeaders(v map[string]*string) *AddOrgWebPathsResponse {
	s.Headers = v
	return s
}

func (s *AddOrgWebPathsResponse) SetBody(v *AddOrgWebPathsResponseBody) *AddOrgWebPathsResponse {
	s.Body = v
	return s
}

type CreateAvdsBagOrderRequest struct {
	Period        *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit    *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	FlowoutSpec   *string `json:"FlowoutSpec,omitempty" xml:"FlowoutSpec,omitempty"`
	Pack          *string `json:"Pack,omitempty" xml:"Pack,omitempty"`
}

func (s CreateAvdsBagOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsBagOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateAvdsBagOrderRequest) SetPeriod(v int32) *CreateAvdsBagOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateAvdsBagOrderRequest) SetPeriodUnit(v string) *CreateAvdsBagOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAvdsBagOrderRequest) SetAutoPay(v bool) *CreateAvdsBagOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAvdsBagOrderRequest) SetAutoUseCoupon(v bool) *CreateAvdsBagOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateAvdsBagOrderRequest) SetFlowoutSpec(v string) *CreateAvdsBagOrderRequest {
	s.FlowoutSpec = &v
	return s
}

func (s *CreateAvdsBagOrderRequest) SetPack(v string) *CreateAvdsBagOrderRequest {
	s.Pack = &v
	return s
}

type CreateAvdsBagOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAvdsBagOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsBagOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAvdsBagOrderResponseBody) SetRequestId(v string) *CreateAvdsBagOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAvdsBagOrderResponseBody) SetOrderId(v string) *CreateAvdsBagOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateAvdsBagOrderResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAvdsBagOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAvdsBagOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsBagOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateAvdsBagOrderResponse) SetHeaders(v map[string]*string) *CreateAvdsBagOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateAvdsBagOrderResponse) SetBody(v *CreateAvdsBagOrderResponseBody) *CreateAvdsBagOrderResponse {
	s.Body = v
	return s
}

type CreateAvdsOrderRequest struct {
	Period         *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit     *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay        *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon  *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	SiteNum        *string `json:"SiteNum,omitempty" xml:"SiteNum,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	UrlNum         *string `json:"UrlNum,omitempty" xml:"UrlNum,omitempty"`
	AddVulNum      *string `json:"AddVulNum,omitempty" xml:"AddVulNum,omitempty"`
}

func (s CreateAvdsOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateAvdsOrderRequest) SetPeriod(v int32) *CreateAvdsOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetPeriodUnit(v string) *CreateAvdsOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetAutoPay(v bool) *CreateAvdsOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetAutoUseCoupon(v bool) *CreateAvdsOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetSiteNum(v string) *CreateAvdsOrderRequest {
	s.SiteNum = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetServiceVersion(v string) *CreateAvdsOrderRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetUrlNum(v string) *CreateAvdsOrderRequest {
	s.UrlNum = &v
	return s
}

func (s *CreateAvdsOrderRequest) SetAddVulNum(v string) *CreateAvdsOrderRequest {
	s.AddVulNum = &v
	return s
}

type CreateAvdsOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAvdsOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAvdsOrderResponseBody) SetRequestId(v string) *CreateAvdsOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAvdsOrderResponseBody) SetOrderId(v string) *CreateAvdsOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateAvdsOrderResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAvdsOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAvdsOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateAvdsOrderResponse) SetHeaders(v map[string]*string) *CreateAvdsOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateAvdsOrderResponse) SetBody(v *CreateAvdsOrderResponseBody) *CreateAvdsOrderResponse {
	s.Body = v
	return s
}

type CreateAvdsPublicOrderRequest struct {
	AutoPay       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	NameTime      *string `json:"NameTime,omitempty" xml:"NameTime,omitempty"`
}

func (s CreateAvdsPublicOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsPublicOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateAvdsPublicOrderRequest) SetAutoPay(v bool) *CreateAvdsPublicOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAvdsPublicOrderRequest) SetAutoUseCoupon(v bool) *CreateAvdsPublicOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateAvdsPublicOrderRequest) SetNameTime(v string) *CreateAvdsPublicOrderRequest {
	s.NameTime = &v
	return s
}

type CreateAvdsPublicOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAvdsPublicOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsPublicOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAvdsPublicOrderResponseBody) SetRequestId(v string) *CreateAvdsPublicOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAvdsPublicOrderResponseBody) SetOrderId(v string) *CreateAvdsPublicOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateAvdsPublicOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAvdsPublicOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAvdsPublicOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAvdsPublicOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateAvdsPublicOrderResponse) SetHeaders(v map[string]*string) *CreateAvdsPublicOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateAvdsPublicOrderResponse) SetBody(v *CreateAvdsPublicOrderResponseBody) *CreateAvdsPublicOrderResponse {
	s.Body = v
	return s
}

type CreateOrganizationRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Domains     []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s CreateOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationRequest) GoString() string {
	return s.String()
}

func (s *CreateOrganizationRequest) SetSourceIp(v string) *CreateOrganizationRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateOrganizationRequest) SetName(v string) *CreateOrganizationRequest {
	s.Name = &v
	return s
}

func (s *CreateOrganizationRequest) SetDescription(v string) *CreateOrganizationRequest {
	s.Description = &v
	return s
}

func (s *CreateOrganizationRequest) SetDomains(v []*string) *CreateOrganizationRequest {
	s.Domains = v
	return s
}

type CreateOrganizationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationResponseBody) SetRequestId(v string) *CreateOrganizationResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrganizationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationResponse) GoString() string {
	return s.String()
}

func (s *CreateOrganizationResponse) SetHeaders(v map[string]*string) *CreateOrganizationResponse {
	s.Headers = v
	return s
}

func (s *CreateOrganizationResponse) SetBody(v *CreateOrganizationResponseBody) *CreateOrganizationResponse {
	s.Body = v
	return s
}

type CreateSessionRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Asset        *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	TTL          *int32  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	LoginSession *string `json:"LoginSession,omitempty" xml:"LoginSession,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) SetSourceIp(v string) *CreateSessionRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSessionRequest) SetAsset(v string) *CreateSessionRequest {
	s.Asset = &v
	return s
}

func (s *CreateSessionRequest) SetTTL(v int32) *CreateSessionRequest {
	s.TTL = &v
	return s
}

func (s *CreateSessionRequest) SetLoginSession(v string) *CreateSessionRequest {
	s.LoginSession = &v
	return s
}

type CreateSessionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) SetRequestId(v string) *CreateSessionResponseBody {
	s.RequestId = &v
	return s
}

type CreateSessionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionResponse) SetHeaders(v map[string]*string) *CreateSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionResponse) SetBody(v *CreateSessionResponseBody) *CreateSessionResponse {
	s.Body = v
	return s
}

type DeleteAssetsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AssetIds        []*string `json:"AssetIds,omitempty" xml:"AssetIds,omitempty" type:"Repeated"`
}

func (s DeleteAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAssetsRequest) SetSourceIp(v string) *DeleteAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAssetsRequest) SetResourceOwnerId(v int64) *DeleteAssetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAssetsRequest) SetAssetIds(v []*string) *DeleteAssetsRequest {
	s.AssetIds = v
	return s
}

type DeleteAssetsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAssetsResponseBody) SetRequestId(v string) *DeleteAssetsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAssetsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAssetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAssetsResponse) SetHeaders(v map[string]*string) *DeleteAssetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAssetsResponse) SetBody(v *DeleteAssetsResponseBody) *DeleteAssetsResponse {
	s.Body = v
	return s
}

type DeleteOrganizationsRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgIds   []*string `json:"OrgIds,omitempty" xml:"OrgIds,omitempty" type:"Repeated"`
}

func (s DeleteOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationsRequest) SetSourceIp(v string) *DeleteOrganizationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteOrganizationsRequest) SetOrgIds(v []*string) *DeleteOrganizationsRequest {
	s.OrgIds = v
	return s
}

type DeleteOrganizationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationsResponseBody) SetRequestId(v string) *DeleteOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOrganizationsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationsResponse) SetHeaders(v map[string]*string) *DeleteOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrganizationsResponse) SetBody(v *DeleteOrganizationsResponseBody) *DeleteOrganizationsResponse {
	s.Body = v
	return s
}

type DeleteOrgAttackSurfaceRecordsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	OrgId    *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Records  []*int  `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DeleteOrgAttackSurfaceRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgAttackSurfaceRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrgAttackSurfaceRecordsRequest) SetSourceIp(v string) *DeleteOrgAttackSurfaceRecordsRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteOrgAttackSurfaceRecordsRequest) SetSource(v string) *DeleteOrgAttackSurfaceRecordsRequest {
	s.Source = &v
	return s
}

func (s *DeleteOrgAttackSurfaceRecordsRequest) SetOrgId(v int32) *DeleteOrgAttackSurfaceRecordsRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteOrgAttackSurfaceRecordsRequest) SetRecords(v []*int) *DeleteOrgAttackSurfaceRecordsRequest {
	s.Records = v
	return s
}

type DeleteOrgAttackSurfaceRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOrgAttackSurfaceRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgAttackSurfaceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrgAttackSurfaceRecordsResponseBody) SetRequestId(v string) *DeleteOrgAttackSurfaceRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOrgAttackSurfaceRecordsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOrgAttackSurfaceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOrgAttackSurfaceRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgAttackSurfaceRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrgAttackSurfaceRecordsResponse) SetHeaders(v map[string]*string) *DeleteOrgAttackSurfaceRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrgAttackSurfaceRecordsResponse) SetBody(v *DeleteOrgAttackSurfaceRecordsResponseBody) *DeleteOrgAttackSurfaceRecordsResponse {
	s.Body = v
	return s
}

type DeleteSessionRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SessionId *int32  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSessionRequest) SetSourceIp(v string) *DeleteSessionRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteSessionRequest) SetSessionId(v int32) *DeleteSessionRequest {
	s.SessionId = &v
	return s
}

type DeleteSessionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponseBody) SetRequestId(v string) *DeleteSessionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSessionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponse) SetHeaders(v map[string]*string) *DeleteSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSessionResponse) SetBody(v *DeleteSessionResponseBody) *DeleteSessionResponse {
	s.Body = v
	return s
}

type DeleteUserAttackSurfaceRecordsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Records  []*int  `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DeleteUserAttackSurfaceRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserAttackSurfaceRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserAttackSurfaceRecordsRequest) SetSourceIp(v string) *DeleteUserAttackSurfaceRecordsRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteUserAttackSurfaceRecordsRequest) SetSource(v string) *DeleteUserAttackSurfaceRecordsRequest {
	s.Source = &v
	return s
}

func (s *DeleteUserAttackSurfaceRecordsRequest) SetRecords(v []*int) *DeleteUserAttackSurfaceRecordsRequest {
	s.Records = v
	return s
}

type DeleteUserAttackSurfaceRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserAttackSurfaceRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserAttackSurfaceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserAttackSurfaceRecordsResponseBody) SetRequestId(v string) *DeleteUserAttackSurfaceRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserAttackSurfaceRecordsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserAttackSurfaceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserAttackSurfaceRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserAttackSurfaceRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserAttackSurfaceRecordsResponse) SetHeaders(v map[string]*string) *DeleteUserAttackSurfaceRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserAttackSurfaceRecordsResponse) SetBody(v *DeleteUserAttackSurfaceRecordsResponseBody) *DeleteUserAttackSurfaceRecordsResponse {
	s.Body = v
	return s
}

type DescribeAllVulnerabilitiesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScanId          *string `json:"ScanId,omitempty" xml:"ScanId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Search          *string `json:"Search,omitempty" xml:"Search,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Severity        *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	BeginTime       *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Module          *string `json:"Module,omitempty" xml:"Module,omitempty"`
	VulType         *int64  `json:"VulType,omitempty" xml:"VulType,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAllVulnerabilitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllVulnerabilitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllVulnerabilitiesRequest) SetSourceIp(v string) *DescribeAllVulnerabilitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetResourceOwnerId(v int64) *DescribeAllVulnerabilitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetScanId(v string) *DescribeAllVulnerabilitiesRequest {
	s.ScanId = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetName(v string) *DescribeAllVulnerabilitiesRequest {
	s.Name = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetSearch(v string) *DescribeAllVulnerabilitiesRequest {
	s.Search = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetLang(v string) *DescribeAllVulnerabilitiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetSeverity(v int32) *DescribeAllVulnerabilitiesRequest {
	s.Severity = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetStatus(v string) *DescribeAllVulnerabilitiesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetBeginTime(v int64) *DescribeAllVulnerabilitiesRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetEndTime(v int64) *DescribeAllVulnerabilitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetTaskId(v int64) *DescribeAllVulnerabilitiesRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetCategory(v string) *DescribeAllVulnerabilitiesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetModule(v string) *DescribeAllVulnerabilitiesRequest {
	s.Module = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetVulType(v int64) *DescribeAllVulnerabilitiesRequest {
	s.VulType = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetCurrentPage(v int32) *DescribeAllVulnerabilitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAllVulnerabilitiesRequest) SetPageSize(v int32) *DescribeAllVulnerabilitiesRequest {
	s.PageSize = &v
	return s
}

type DescribeAllVulnerabilitiesResponseBody struct {
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount   *int32                                        `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*DescribeAllVulnerabilitiesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Count       *int32                                        `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAllVulnerabilitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllVulnerabilitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetTotalCount(v int32) *DescribeAllVulnerabilitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetRequestId(v string) *DescribeAllVulnerabilitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetPageSize(v int32) *DescribeAllVulnerabilitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetPageCount(v int32) *DescribeAllVulnerabilitiesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetCurrentPage(v int32) *DescribeAllVulnerabilitiesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetList(v []*DescribeAllVulnerabilitiesResponseBodyList) *DescribeAllVulnerabilitiesResponseBody {
	s.List = v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBody) SetCount(v int32) *DescribeAllVulnerabilitiesResponseBody {
	s.Count = &v
	return s
}

type DescribeAllVulnerabilitiesResponseBodyList struct {
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Severity             *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	LastDiscoveredAt     *int64  `json:"LastDiscoveredAt,omitempty" xml:"LastDiscoveredAt,omitempty"`
	Hostname             *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VulnerabilityTypeDes *string `json:"VulnerabilityTypeDes,omitempty" xml:"VulnerabilityTypeDes,omitempty"`
	Target               *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeAllVulnerabilitiesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllVulnerabilitiesResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetStatus(v int32) *DescribeAllVulnerabilitiesResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetSeverity(v int32) *DescribeAllVulnerabilitiesResponseBodyList {
	s.Severity = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetLastDiscoveredAt(v int64) *DescribeAllVulnerabilitiesResponseBodyList {
	s.LastDiscoveredAt = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetHostname(v string) *DescribeAllVulnerabilitiesResponseBodyList {
	s.Hostname = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetName(v string) *DescribeAllVulnerabilitiesResponseBodyList {
	s.Name = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetTaskId(v int64) *DescribeAllVulnerabilitiesResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetVulnerabilityTypeDes(v string) *DescribeAllVulnerabilitiesResponseBodyList {
	s.VulnerabilityTypeDes = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetTarget(v string) *DescribeAllVulnerabilitiesResponseBodyList {
	s.Target = &v
	return s
}

func (s *DescribeAllVulnerabilitiesResponseBodyList) SetId(v int64) *DescribeAllVulnerabilitiesResponseBodyList {
	s.Id = &v
	return s
}

type DescribeAllVulnerabilitiesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllVulnerabilitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllVulnerabilitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllVulnerabilitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllVulnerabilitiesResponse) SetHeaders(v map[string]*string) *DescribeAllVulnerabilitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllVulnerabilitiesResponse) SetBody(v *DescribeAllVulnerabilitiesResponseBody) *DescribeAllVulnerabilitiesResponse {
	s.Body = v
	return s
}

type DescribeAssetsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status          *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Search          *string   `json:"Search,omitempty" xml:"Search,omitempty"`
	AssetGroupId    *string   `json:"AssetGroupId,omitempty" xml:"AssetGroupId,omitempty"`
	Source          *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	GmtCreateFrom   *int64    `json:"GmtCreateFrom,omitempty" xml:"GmtCreateFrom,omitempty"`
	GmtCreateTo     *int64    `json:"GmtCreateTo,omitempty" xml:"GmtCreateTo,omitempty"`
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Types           []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	Assets          []*string `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
}

func (s DescribeAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetsRequest) SetSourceIp(v string) *DescribeAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetsRequest) SetResourceOwnerId(v int64) *DescribeAssetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAssetsRequest) SetStatus(v string) *DescribeAssetsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAssetsRequest) SetSearch(v string) *DescribeAssetsRequest {
	s.Search = &v
	return s
}

func (s *DescribeAssetsRequest) SetAssetGroupId(v string) *DescribeAssetsRequest {
	s.AssetGroupId = &v
	return s
}

func (s *DescribeAssetsRequest) SetSource(v string) *DescribeAssetsRequest {
	s.Source = &v
	return s
}

func (s *DescribeAssetsRequest) SetGmtCreateFrom(v int64) *DescribeAssetsRequest {
	s.GmtCreateFrom = &v
	return s
}

func (s *DescribeAssetsRequest) SetGmtCreateTo(v int64) *DescribeAssetsRequest {
	s.GmtCreateTo = &v
	return s
}

func (s *DescribeAssetsRequest) SetCurrentPage(v int32) *DescribeAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAssetsRequest) SetPageSize(v int32) *DescribeAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAssetsRequest) SetTypes(v []*string) *DescribeAssetsRequest {
	s.Types = v
	return s
}

func (s *DescribeAssetsRequest) SetAssets(v []*string) *DescribeAssetsRequest {
	s.Assets = v
	return s
}

type DescribeAssetsResponseBody struct {
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount   *int32                            `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*DescribeAssetsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Count       *int32                            `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetsResponseBody) SetTotalCount(v int32) *DescribeAssetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAssetsResponseBody) SetRequestId(v string) *DescribeAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetsResponseBody) SetPageSize(v int32) *DescribeAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAssetsResponseBody) SetPageCount(v int32) *DescribeAssetsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeAssetsResponseBody) SetCurrentPage(v int32) *DescribeAssetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAssetsResponseBody) SetList(v []*DescribeAssetsResponseBodyList) *DescribeAssetsResponseBody {
	s.List = v
	return s
}

func (s *DescribeAssetsResponseBody) SetCount(v int32) *DescribeAssetsResponseBody {
	s.Count = &v
	return s
}

type DescribeAssetsResponseBodyList struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	LastScanDate *int64  `json:"LastScanDate,omitempty" xml:"LastScanDate,omitempty"`
	ExpireTime   *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	AssetId      *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Asset        *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
}

func (s DescribeAssetsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeAssetsResponseBodyList) SetType(v string) *DescribeAssetsResponseBodyList {
	s.Type = &v
	return s
}

func (s *DescribeAssetsResponseBodyList) SetLastScanDate(v int64) *DescribeAssetsResponseBodyList {
	s.LastScanDate = &v
	return s
}

func (s *DescribeAssetsResponseBodyList) SetExpireTime(v int64) *DescribeAssetsResponseBodyList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAssetsResponseBodyList) SetAssetId(v string) *DescribeAssetsResponseBodyList {
	s.AssetId = &v
	return s
}

func (s *DescribeAssetsResponseBodyList) SetGmtCreate(v int64) *DescribeAssetsResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAssetsResponseBodyList) SetSource(v string) *DescribeAssetsResponseBodyList {
	s.Source = &v
	return s
}

func (s *DescribeAssetsResponseBodyList) SetAsset(v string) *DescribeAssetsResponseBodyList {
	s.Asset = &v
	return s
}

type DescribeAssetsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetsResponse) SetHeaders(v map[string]*string) *DescribeAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetsResponse) SetBody(v *DescribeAssetsResponseBody) *DescribeAssetsResponse {
	s.Body = v
	return s
}

type DescribeAttackSurfacesFacetsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId   *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAttackSurfacesFacetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackSurfacesFacetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackSurfacesFacetsRequest) SetSourceIp(v string) *DescribeAttackSurfacesFacetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsRequest) SetTaskId(v int32) *DescribeAttackSurfacesFacetsRequest {
	s.TaskId = &v
	return s
}

type DescribeAttackSurfacesFacetsResponseBody struct {
	Domains         *int32  `json:"Domains,omitempty" xml:"Domains,omitempty"`
	Hosts           *int32  `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	WebPaths        *int32  `json:"WebPaths,omitempty" xml:"WebPaths,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DNSMap          *int32  `json:"DNSMap,omitempty" xml:"DNSMap,omitempty"`
	WebServers      *int32  `json:"WebServers,omitempty" xml:"WebServers,omitempty"`
	Ports           *int32  `json:"Ports,omitempty" xml:"Ports,omitempty"`
	CrawlerRequests *int32  `json:"CrawlerRequests,omitempty" xml:"CrawlerRequests,omitempty"`
	WebTechs        *int32  `json:"WebTechs,omitempty" xml:"WebTechs,omitempty"`
	Subdomains      *int32  `json:"Subdomains,omitempty" xml:"Subdomains,omitempty"`
}

func (s DescribeAttackSurfacesFacetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackSurfacesFacetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetDomains(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.Domains = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetHosts(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.Hosts = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetWebPaths(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.WebPaths = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetRequestId(v string) *DescribeAttackSurfacesFacetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetDNSMap(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.DNSMap = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetWebServers(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.WebServers = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetPorts(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.Ports = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetCrawlerRequests(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.CrawlerRequests = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetWebTechs(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.WebTechs = &v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponseBody) SetSubdomains(v int32) *DescribeAttackSurfacesFacetsResponseBody {
	s.Subdomains = &v
	return s
}

type DescribeAttackSurfacesFacetsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAttackSurfacesFacetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAttackSurfacesFacetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackSurfacesFacetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackSurfacesFacetsResponse) SetHeaders(v map[string]*string) *DescribeAttackSurfacesFacetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackSurfacesFacetsResponse) SetBody(v *DescribeAttackSurfacesFacetsResponseBody) *DescribeAttackSurfacesFacetsResponse {
	s.Body = v
	return s
}

type DescribeCrawlerRequestsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCrawlerRequestsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrawlerRequestsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrawlerRequestsRequest) SetSourceIp(v string) *DescribeCrawlerRequestsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCrawlerRequestsRequest) SetTaskId(v int32) *DescribeCrawlerRequestsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCrawlerRequestsRequest) SetAsset(v string) *DescribeCrawlerRequestsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeCrawlerRequestsRequest) SetCurrentPage(v int32) *DescribeCrawlerRequestsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCrawlerRequestsRequest) SetPageSize(v int32) *DescribeCrawlerRequestsRequest {
	s.PageSize = &v
	return s
}

type DescribeCrawlerRequestsResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                        `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeCrawlerRequestsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeCrawlerRequestsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrawlerRequestsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrawlerRequestsResponseBody) SetRequestId(v string) *DescribeCrawlerRequestsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBody) SetTotal(v int32) *DescribeCrawlerRequestsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBody) SetRecords(v []*DescribeCrawlerRequestsResponseBodyRecords) *DescribeCrawlerRequestsResponseBody {
	s.Records = v
	return s
}

type DescribeCrawlerRequestsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	URL       *string `json:"URL,omitempty" xml:"URL,omitempty"`
	Method    *string `json:"Method,omitempty" xml:"Method,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Cookies   *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
}

func (s DescribeCrawlerRequestsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrawlerRequestsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeCrawlerRequestsResponseBodyRecords) SetIndex(v int32) *DescribeCrawlerRequestsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBodyRecords) SetData(v string) *DescribeCrawlerRequestsResponseBodyRecords {
	s.Data = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBodyRecords) SetURL(v string) *DescribeCrawlerRequestsResponseBodyRecords {
	s.URL = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBodyRecords) SetMethod(v string) *DescribeCrawlerRequestsResponseBodyRecords {
	s.Method = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBodyRecords) SetUpdatedAt(v int64) *DescribeCrawlerRequestsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeCrawlerRequestsResponseBodyRecords) SetCookies(v string) *DescribeCrawlerRequestsResponseBodyRecords {
	s.Cookies = &v
	return s
}

type DescribeCrawlerRequestsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCrawlerRequestsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCrawlerRequestsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrawlerRequestsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrawlerRequestsResponse) SetHeaders(v map[string]*string) *DescribeCrawlerRequestsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrawlerRequestsResponse) SetBody(v *DescribeCrawlerRequestsResponseBody) *DescribeCrawlerRequestsResponse {
	s.Body = v
	return s
}

type DescribeDNSMapRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDNSMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSMapRequest) GoString() string {
	return s.String()
}

func (s *DescribeDNSMapRequest) SetSourceIp(v string) *DescribeDNSMapRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDNSMapRequest) SetTaskId(v int32) *DescribeDNSMapRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDNSMapRequest) SetAsset(v string) *DescribeDNSMapRequest {
	s.Asset = &v
	return s
}

func (s *DescribeDNSMapRequest) SetCurrentPage(v int32) *DescribeDNSMapRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDNSMapRequest) SetPageSize(v int32) *DescribeDNSMapRequest {
	s.PageSize = &v
	return s
}

type DescribeDNSMapResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                               `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeDNSMapResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeDNSMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSMapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDNSMapResponseBody) SetRequestId(v string) *DescribeDNSMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDNSMapResponseBody) SetTotal(v int32) *DescribeDNSMapResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDNSMapResponseBody) SetRecords(v []*DescribeDNSMapResponseBodyRecords) *DescribeDNSMapResponseBody {
	s.Records = v
	return s
}

type DescribeDNSMapResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Record    *string `json:"Record,omitempty" xml:"Record,omitempty"`
}

func (s DescribeDNSMapResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSMapResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDNSMapResponseBodyRecords) SetIndex(v int32) *DescribeDNSMapResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeDNSMapResponseBodyRecords) SetType(v string) *DescribeDNSMapResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *DescribeDNSMapResponseBodyRecords) SetDomain(v string) *DescribeDNSMapResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *DescribeDNSMapResponseBodyRecords) SetUpdatedAt(v int64) *DescribeDNSMapResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeDNSMapResponseBodyRecords) SetRecord(v string) *DescribeDNSMapResponseBodyRecords {
	s.Record = &v
	return s
}

type DescribeDNSMapResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDNSMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDNSMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDNSMapResponse) GoString() string {
	return s.String()
}

func (s *DescribeDNSMapResponse) SetHeaders(v map[string]*string) *DescribeDNSMapResponse {
	s.Headers = v
	return s
}

func (s *DescribeDNSMapResponse) SetBody(v *DescribeDNSMapResponseBody) *DescribeDNSMapResponse {
	s.Body = v
	return s
}

type DescribeDomainAttackSurfacesFacetsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId   *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainAttackSurfacesFacetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackSurfacesFacetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackSurfacesFacetsRequest) SetSourceIp(v string) *DescribeDomainAttackSurfacesFacetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsRequest) SetTaskId(v int32) *DescribeDomainAttackSurfacesFacetsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsRequest) SetDomain(v string) *DescribeDomainAttackSurfacesFacetsRequest {
	s.Domain = &v
	return s
}

type DescribeDomainAttackSurfacesFacetsResponseBody struct {
	WebPaths        *int32  `json:"WebPaths,omitempty" xml:"WebPaths,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebServers      *int32  `json:"WebServers,omitempty" xml:"WebServers,omitempty"`
	CrawlerRequests *int32  `json:"CrawlerRequests,omitempty" xml:"CrawlerRequests,omitempty"`
	WebTechs        *int32  `json:"WebTechs,omitempty" xml:"WebTechs,omitempty"`
}

func (s DescribeDomainAttackSurfacesFacetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackSurfacesFacetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackSurfacesFacetsResponseBody) SetWebPaths(v int32) *DescribeDomainAttackSurfacesFacetsResponseBody {
	s.WebPaths = &v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsResponseBody) SetRequestId(v string) *DescribeDomainAttackSurfacesFacetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsResponseBody) SetWebServers(v int32) *DescribeDomainAttackSurfacesFacetsResponseBody {
	s.WebServers = &v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsResponseBody) SetCrawlerRequests(v int32) *DescribeDomainAttackSurfacesFacetsResponseBody {
	s.CrawlerRequests = &v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsResponseBody) SetWebTechs(v int32) *DescribeDomainAttackSurfacesFacetsResponseBody {
	s.WebTechs = &v
	return s
}

type DescribeDomainAttackSurfacesFacetsResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainAttackSurfacesFacetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainAttackSurfacesFacetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackSurfacesFacetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackSurfacesFacetsResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackSurfacesFacetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackSurfacesFacetsResponse) SetBody(v *DescribeDomainAttackSurfacesFacetsResponseBody) *DescribeDomainAttackSurfacesFacetsResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetSourceIp(v string) *DescribeDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainsRequest) SetTaskId(v int32) *DescribeDomainsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDomainsRequest) SetAsset(v string) *DescribeDomainsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeDomainsRequest) SetCurrentPage(v int32) *DescribeDomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int32) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeDomainsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotal(v int32) *DescribeDomainsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetRecords(v []*DescribeDomainsResponseBodyRecords) *DescribeDomainsResponseBody {
	s.Records = v
	return s
}

type DescribeDomainsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeDomainsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyRecords) SetIndex(v int32) *DescribeDomainsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeDomainsResponseBodyRecords) SetDomain(v string) *DescribeDomainsResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseBodyRecords) SetUpdatedAt(v int64) *DescribeDomainsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

type DescribeDomainsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponse) SetHeaders(v map[string]*string) *DescribeDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

type DescribeHostAttackSurfacesFacetsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId   *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Host     *string `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeHostAttackSurfacesFacetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAttackSurfacesFacetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostAttackSurfacesFacetsRequest) SetSourceIp(v string) *DescribeHostAttackSurfacesFacetsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsRequest) SetTaskId(v int32) *DescribeHostAttackSurfacesFacetsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsRequest) SetHost(v string) *DescribeHostAttackSurfacesFacetsRequest {
	s.Host = &v
	return s
}

type DescribeHostAttackSurfacesFacetsResponseBody struct {
	Hosts           *int32  `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	WebPaths        *int32  `json:"WebPaths,omitempty" xml:"WebPaths,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ports           *int32  `json:"Ports,omitempty" xml:"Ports,omitempty"`
	CrawlerRequests *int32  `json:"CrawlerRequests,omitempty" xml:"CrawlerRequests,omitempty"`
	WebTechs        *int32  `json:"WebTechs,omitempty" xml:"WebTechs,omitempty"`
}

func (s DescribeHostAttackSurfacesFacetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAttackSurfacesFacetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostAttackSurfacesFacetsResponseBody) SetHosts(v int32) *DescribeHostAttackSurfacesFacetsResponseBody {
	s.Hosts = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsResponseBody) SetWebPaths(v int32) *DescribeHostAttackSurfacesFacetsResponseBody {
	s.WebPaths = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsResponseBody) SetRequestId(v string) *DescribeHostAttackSurfacesFacetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsResponseBody) SetPorts(v int32) *DescribeHostAttackSurfacesFacetsResponseBody {
	s.Ports = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsResponseBody) SetCrawlerRequests(v int32) *DescribeHostAttackSurfacesFacetsResponseBody {
	s.CrawlerRequests = &v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsResponseBody) SetWebTechs(v int32) *DescribeHostAttackSurfacesFacetsResponseBody {
	s.WebTechs = &v
	return s
}

type DescribeHostAttackSurfacesFacetsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHostAttackSurfacesFacetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHostAttackSurfacesFacetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAttackSurfacesFacetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostAttackSurfacesFacetsResponse) SetHeaders(v map[string]*string) *DescribeHostAttackSurfacesFacetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostAttackSurfacesFacetsResponse) SetBody(v *DescribeHostAttackSurfacesFacetsResponseBody) *DescribeHostAttackSurfacesFacetsResponse {
	s.Body = v
	return s
}

type DescribeHostsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostsRequest) SetSourceIp(v string) *DescribeHostsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHostsRequest) SetTaskId(v int32) *DescribeHostsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeHostsRequest) SetAsset(v string) *DescribeHostsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeHostsRequest) SetCurrentPage(v int32) *DescribeHostsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHostsRequest) SetPageSize(v int32) *DescribeHostsRequest {
	s.PageSize = &v
	return s
}

type DescribeHostsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                              `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeHostsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostsResponseBody) SetRequestId(v string) *DescribeHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostsResponseBody) SetTotal(v int32) *DescribeHostsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHostsResponseBody) SetRecords(v []*DescribeHostsResponseBodyRecords) *DescribeHostsResponseBody {
	s.Records = v
	return s
}

type DescribeHostsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	OS        *string `json:"OS,omitempty" xml:"OS,omitempty"`
	Hostname  *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	IsUp      *string `json:"IsUp,omitempty" xml:"IsUp,omitempty"`
}

func (s DescribeHostsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeHostsResponseBodyRecords) SetIndex(v int32) *DescribeHostsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeHostsResponseBodyRecords) SetOS(v string) *DescribeHostsResponseBodyRecords {
	s.OS = &v
	return s
}

func (s *DescribeHostsResponseBodyRecords) SetHostname(v string) *DescribeHostsResponseBodyRecords {
	s.Hostname = &v
	return s
}

func (s *DescribeHostsResponseBodyRecords) SetUpdatedAt(v int64) *DescribeHostsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeHostsResponseBodyRecords) SetIP(v string) *DescribeHostsResponseBodyRecords {
	s.IP = &v
	return s
}

func (s *DescribeHostsResponseBodyRecords) SetIsUp(v string) *DescribeHostsResponseBodyRecords {
	s.IsUp = &v
	return s
}

type DescribeHostsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostsResponse) SetHeaders(v map[string]*string) *DescribeHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostsResponse) SetBody(v *DescribeHostsResponseBody) *DescribeHostsResponse {
	s.Body = v
	return s
}

type DescribeListSessionsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeListSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeListSessionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeListSessionsRequest) SetSourceIp(v string) *DescribeListSessionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeListSessionsRequest) SetAsset(v string) *DescribeListSessionsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeListSessionsRequest) SetCurrentPage(v int32) *DescribeListSessionsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeListSessionsRequest) SetPageSize(v int32) *DescribeListSessionsRequest {
	s.PageSize = &v
	return s
}

type DescribeListSessionsResponseBody struct {
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Sessions    []*DescribeListSessionsResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
}

func (s DescribeListSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeListSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeListSessionsResponseBody) SetTotalCount(v int32) *DescribeListSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeListSessionsResponseBody) SetPageSize(v int32) *DescribeListSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeListSessionsResponseBody) SetRequestId(v string) *DescribeListSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeListSessionsResponseBody) SetCurrentPage(v int32) *DescribeListSessionsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeListSessionsResponseBody) SetSessions(v []*DescribeListSessionsResponseBodySessions) *DescribeListSessionsResponseBody {
	s.Sessions = v
	return s
}

type DescribeListSessionsResponseBodySessions struct {
	TTL          *int32  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Expired      *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreatedAt    *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AliUid       *int32  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ModifiedAt   *int64  `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	LoginSession *string `json:"LoginSession,omitempty" xml:"LoginSession,omitempty"`
	SessionId    *int32  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Asset        *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
}

func (s DescribeListSessionsResponseBodySessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeListSessionsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeListSessionsResponseBodySessions) SetTTL(v int32) *DescribeListSessionsResponseBodySessions {
	s.TTL = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetExpired(v int32) *DescribeListSessionsResponseBodySessions {
	s.Expired = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetCreatedAt(v int64) *DescribeListSessionsResponseBodySessions {
	s.CreatedAt = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetAliUid(v int32) *DescribeListSessionsResponseBodySessions {
	s.AliUid = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetModifiedAt(v int64) *DescribeListSessionsResponseBodySessions {
	s.ModifiedAt = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetLoginSession(v string) *DescribeListSessionsResponseBodySessions {
	s.LoginSession = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetSessionId(v int32) *DescribeListSessionsResponseBodySessions {
	s.SessionId = &v
	return s
}

func (s *DescribeListSessionsResponseBodySessions) SetAsset(v string) *DescribeListSessionsResponseBodySessions {
	s.Asset = &v
	return s
}

type DescribeListSessionsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeListSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeListSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeListSessionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeListSessionsResponse) SetHeaders(v map[string]*string) *DescribeListSessionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeListSessionsResponse) SetBody(v *DescribeListSessionsResponseBody) *DescribeListSessionsResponse {
	s.Body = v
	return s
}

type DescribeOrgAttackSurfaceDetailsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RecordId *int32  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	OrgId    *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeOrgAttackSurfaceDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgAttackSurfaceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgAttackSurfaceDetailsRequest) SetSourceIp(v string) *DescribeOrgAttackSurfaceDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsRequest) SetRecordId(v int32) *DescribeOrgAttackSurfaceDetailsRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsRequest) SetOrgId(v int32) *DescribeOrgAttackSurfaceDetailsRequest {
	s.OrgId = &v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsRequest) SetSource(v string) *DescribeOrgAttackSurfaceDetailsRequest {
	s.Source = &v
	return s
}

type DescribeOrgAttackSurfaceDetailsResponseBody struct {
	LastScannedAt  *int64  `json:"LastScannedAt,omitempty" xml:"LastScannedAt,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FirstScannedAt *int64  `json:"FirstScannedAt,omitempty" xml:"FirstScannedAt,omitempty"`
	Occurrences    *int32  `json:"Occurrences,omitempty" xml:"Occurrences,omitempty"`
}

func (s DescribeOrgAttackSurfaceDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgAttackSurfaceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrgAttackSurfaceDetailsResponseBody) SetLastScannedAt(v int64) *DescribeOrgAttackSurfaceDetailsResponseBody {
	s.LastScannedAt = &v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsResponseBody) SetRequestId(v string) *DescribeOrgAttackSurfaceDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsResponseBody) SetFirstScannedAt(v int64) *DescribeOrgAttackSurfaceDetailsResponseBody {
	s.FirstScannedAt = &v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsResponseBody) SetOccurrences(v int32) *DescribeOrgAttackSurfaceDetailsResponseBody {
	s.Occurrences = &v
	return s
}

type DescribeOrgAttackSurfaceDetailsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOrgAttackSurfaceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOrgAttackSurfaceDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgAttackSurfaceDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrgAttackSurfaceDetailsResponse) SetHeaders(v map[string]*string) *DescribeOrgAttackSurfaceDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrgAttackSurfaceDetailsResponse) SetBody(v *DescribeOrgAttackSurfaceDetailsResponseBody) *DescribeOrgAttackSurfaceDetailsResponse {
	s.Body = v
	return s
}

type DescribeOrgInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId    *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s DescribeOrgInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgInfoRequest) SetSourceIp(v string) *DescribeOrgInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOrgInfoRequest) SetOrgId(v int32) *DescribeOrgInfoRequest {
	s.OrgId = &v
	return s
}

type DescribeOrgInfoResponseBody struct {
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreatedAt   *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrgInfoResponseBody) SetOrgId(v int32) *DescribeOrgInfoResponseBody {
	s.OrgId = &v
	return s
}

func (s *DescribeOrgInfoResponseBody) SetDescription(v string) *DescribeOrgInfoResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeOrgInfoResponseBody) SetRequestId(v string) *DescribeOrgInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrgInfoResponseBody) SetCreatedAt(v int64) *DescribeOrgInfoResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeOrgInfoResponseBody) SetAliUid(v int64) *DescribeOrgInfoResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeOrgInfoResponseBody) SetName(v string) *DescribeOrgInfoResponseBody {
	s.Name = &v
	return s
}

type DescribeOrgInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrgInfoResponse) SetHeaders(v map[string]*string) *DescribeOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrgInfoResponse) SetBody(v *DescribeOrgInfoResponseBody) *DescribeOrgInfoResponse {
	s.Body = v
	return s
}

type DescribePortsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePortsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortsRequest) SetSourceIp(v string) *DescribePortsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortsRequest) SetTaskId(v int32) *DescribePortsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribePortsRequest) SetAsset(v string) *DescribePortsRequest {
	s.Asset = &v
	return s
}

func (s *DescribePortsRequest) SetCurrentPage(v int32) *DescribePortsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePortsRequest) SetPageSize(v int32) *DescribePortsRequest {
	s.PageSize = &v
	return s
}

type DescribePortsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                              `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribePortsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribePortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortsResponseBody) SetRequestId(v string) *DescribePortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortsResponseBody) SetTotal(v int32) *DescribePortsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePortsResponseBody) SetRecords(v []*DescribePortsResponseBodyRecords) *DescribePortsResponseBody {
	s.Records = v
	return s
}

type DescribePortsResponseBodyRecords struct {
	Service     *string `json:"Service,omitempty" xml:"Service,omitempty"`
	Index       *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	UpdatedAt   *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Port        *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IP          *string `json:"IP,omitempty" xml:"IP,omitempty"`
}

func (s DescribePortsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribePortsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribePortsResponseBodyRecords) SetService(v string) *DescribePortsResponseBodyRecords {
	s.Service = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetIndex(v int32) *DescribePortsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetFingerprint(v string) *DescribePortsResponseBodyRecords {
	s.Fingerprint = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetVersion(v string) *DescribePortsResponseBodyRecords {
	s.Version = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetProduct(v string) *DescribePortsResponseBodyRecords {
	s.Product = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetProtocol(v string) *DescribePortsResponseBodyRecords {
	s.Protocol = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetUpdatedAt(v int64) *DescribePortsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetPort(v string) *DescribePortsResponseBodyRecords {
	s.Port = &v
	return s
}

func (s *DescribePortsResponseBodyRecords) SetIP(v string) *DescribePortsResponseBodyRecords {
	s.IP = &v
	return s
}

type DescribePortsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortsResponse) SetHeaders(v map[string]*string) *DescribePortsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortsResponse) SetBody(v *DescribePortsResponseBody) *DescribePortsResponse {
	s.Body = v
	return s
}

type DescribeScanSessionsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScanId          *string   `json:"ScanId,omitempty" xml:"ScanId,omitempty"`
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search          *string   `json:"Search,omitempty" xml:"Search,omitempty"`
	StatusList      []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeScanSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanSessionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanSessionsRequest) SetSourceIp(v string) *DescribeScanSessionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeScanSessionsRequest) SetResourceOwnerId(v int64) *DescribeScanSessionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScanSessionsRequest) SetScanId(v string) *DescribeScanSessionsRequest {
	s.ScanId = &v
	return s
}

func (s *DescribeScanSessionsRequest) SetCurrentPage(v int32) *DescribeScanSessionsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeScanSessionsRequest) SetPageSize(v int32) *DescribeScanSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScanSessionsRequest) SetSearch(v string) *DescribeScanSessionsRequest {
	s.Search = &v
	return s
}

func (s *DescribeScanSessionsRequest) SetStatusList(v []*string) *DescribeScanSessionsRequest {
	s.StatusList = v
	return s
}

type DescribeScanSessionsResponseBody struct {
	TotalCount  *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount   *int32                                  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*DescribeScanSessionsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Count       *int32                                  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeScanSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanSessionsResponseBody) SetTotalCount(v int32) *DescribeScanSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScanSessionsResponseBody) SetRequestId(v string) *DescribeScanSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanSessionsResponseBody) SetPageSize(v int32) *DescribeScanSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScanSessionsResponseBody) SetPageCount(v int32) *DescribeScanSessionsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeScanSessionsResponseBody) SetCurrentPage(v int32) *DescribeScanSessionsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeScanSessionsResponseBody) SetList(v []*DescribeScanSessionsResponseBodyList) *DescribeScanSessionsResponseBody {
	s.List = v
	return s
}

func (s *DescribeScanSessionsResponseBody) SetCount(v int32) *DescribeScanSessionsResponseBody {
	s.Count = &v
	return s
}

type DescribeScanSessionsResponseBodyList struct {
	ReportStatus *string   `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	FinishedAt   *int64    `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Targets      []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	CreatedAt    *int64    `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ScanId       *string   `json:"ScanId,omitempty" xml:"ScanId,omitempty"`
	Period       *string   `json:"Period,omitempty" xml:"Period,omitempty"`
	TriggerType  *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	ReportUrl    *string   `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	JobStatus    *string   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	RunPercent   *float32  `json:"RunPercent,omitempty" xml:"RunPercent,omitempty"`
	Interval     *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskId       *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScanSessionsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanSessionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeScanSessionsResponseBodyList) SetReportStatus(v string) *DescribeScanSessionsResponseBodyList {
	s.ReportStatus = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetFinishedAt(v int64) *DescribeScanSessionsResponseBodyList {
	s.FinishedAt = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetTargets(v []*string) *DescribeScanSessionsResponseBodyList {
	s.Targets = v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetCreatedAt(v int64) *DescribeScanSessionsResponseBodyList {
	s.CreatedAt = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetScanId(v string) *DescribeScanSessionsResponseBodyList {
	s.ScanId = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetPeriod(v string) *DescribeScanSessionsResponseBodyList {
	s.Period = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetTriggerType(v string) *DescribeScanSessionsResponseBodyList {
	s.TriggerType = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetReportUrl(v string) *DescribeScanSessionsResponseBodyList {
	s.ReportUrl = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetJobStatus(v string) *DescribeScanSessionsResponseBodyList {
	s.JobStatus = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetRunPercent(v float32) *DescribeScanSessionsResponseBodyList {
	s.RunPercent = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetInterval(v int32) *DescribeScanSessionsResponseBodyList {
	s.Interval = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetName(v string) *DescribeScanSessionsResponseBodyList {
	s.Name = &v
	return s
}

func (s *DescribeScanSessionsResponseBodyList) SetTaskId(v int64) *DescribeScanSessionsResponseBodyList {
	s.TaskId = &v
	return s
}

type DescribeScanSessionsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScanSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScanSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanSessionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanSessionsResponse) SetHeaders(v map[string]*string) *DescribeScanSessionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanSessionsResponse) SetBody(v *DescribeScanSessionsResponseBody) *DescribeScanSessionsResponse {
	s.Body = v
	return s
}

type DescribeSessionRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SessionId *int32  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSessionRequest) SetSourceIp(v string) *DescribeSessionRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSessionRequest) SetSessionId(v int32) *DescribeSessionRequest {
	s.SessionId = &v
	return s
}

type DescribeSessionResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Session   *DescribeSessionResponseBodySession `json:"Session,omitempty" xml:"Session,omitempty" type:"Struct"`
}

func (s DescribeSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSessionResponseBody) SetRequestId(v string) *DescribeSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSessionResponseBody) SetSession(v *DescribeSessionResponseBodySession) *DescribeSessionResponseBody {
	s.Session = v
	return s
}

type DescribeSessionResponseBodySession struct {
	TTL          *int32  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Expired      *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreatedAt    *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AliUid       *int32  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ModifiedAt   *int64  `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	LoginSession *string `json:"LoginSession,omitempty" xml:"LoginSession,omitempty"`
	SessionId    *int32  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Asset        *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
}

func (s DescribeSessionResponseBodySession) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionResponseBodySession) GoString() string {
	return s.String()
}

func (s *DescribeSessionResponseBodySession) SetTTL(v int32) *DescribeSessionResponseBodySession {
	s.TTL = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetExpired(v int32) *DescribeSessionResponseBodySession {
	s.Expired = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetCreatedAt(v int64) *DescribeSessionResponseBodySession {
	s.CreatedAt = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetAliUid(v int32) *DescribeSessionResponseBodySession {
	s.AliUid = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetModifiedAt(v int64) *DescribeSessionResponseBodySession {
	s.ModifiedAt = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetLoginSession(v string) *DescribeSessionResponseBodySession {
	s.LoginSession = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetSessionId(v int32) *DescribeSessionResponseBodySession {
	s.SessionId = &v
	return s
}

func (s *DescribeSessionResponseBodySession) SetAsset(v string) *DescribeSessionResponseBodySession {
	s.Asset = &v
	return s
}

type DescribeSessionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSessionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSessionResponse) SetHeaders(v map[string]*string) *DescribeSessionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSessionResponse) SetBody(v *DescribeSessionResponseBody) *DescribeSessionResponse {
	s.Body = v
	return s
}

type DescribeSubdomainsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSubdomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubdomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubdomainsRequest) SetSourceIp(v string) *DescribeSubdomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSubdomainsRequest) SetTaskId(v int32) *DescribeSubdomainsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSubdomainsRequest) SetAsset(v string) *DescribeSubdomainsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeSubdomainsRequest) SetCurrentPage(v int32) *DescribeSubdomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSubdomainsRequest) SetPageSize(v int32) *DescribeSubdomainsRequest {
	s.PageSize = &v
	return s
}

type DescribeSubdomainsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeSubdomainsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeSubdomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubdomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubdomainsResponseBody) SetRequestId(v string) *DescribeSubdomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubdomainsResponseBody) SetTotal(v int32) *DescribeSubdomainsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSubdomainsResponseBody) SetRecords(v []*DescribeSubdomainsResponseBodyRecords) *DescribeSubdomainsResponseBody {
	s.Records = v
	return s
}

type DescribeSubdomainsResponseBodyRecords struct {
	Index      *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	UpdatedAt  *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	RootDomain *string `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
}

func (s DescribeSubdomainsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubdomainsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeSubdomainsResponseBodyRecords) SetIndex(v int32) *DescribeSubdomainsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeSubdomainsResponseBodyRecords) SetDomain(v string) *DescribeSubdomainsResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *DescribeSubdomainsResponseBodyRecords) SetUpdatedAt(v int64) *DescribeSubdomainsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeSubdomainsResponseBodyRecords) SetRootDomain(v string) *DescribeSubdomainsResponseBodyRecords {
	s.RootDomain = &v
	return s
}

type DescribeSubdomainsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubdomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubdomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubdomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubdomainsResponse) SetHeaders(v map[string]*string) *DescribeSubdomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubdomainsResponse) SetBody(v *DescribeSubdomainsResponseBody) *DescribeSubdomainsResponse {
	s.Body = v
	return s
}

type DescribeTaskBriefInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId   *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeTaskBriefInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskBriefInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskBriefInfoRequest) SetSourceIp(v string) *DescribeTaskBriefInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTaskBriefInfoRequest) SetTaskId(v int32) *DescribeTaskBriefInfoRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskBriefInfoRequest) SetLang(v string) *DescribeTaskBriefInfoRequest {
	s.Lang = &v
	return s
}

type DescribeTaskBriefInfoResponseBody struct {
	RisksFacets *DescribeTaskBriefInfoResponseBodyRisksFacets `json:"RisksFacets,omitempty" xml:"RisksFacets,omitempty" type:"Struct"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Brief       *DescribeTaskBriefInfoResponseBodyBrief       `json:"Brief,omitempty" xml:"Brief,omitempty" type:"Struct"`
}

func (s DescribeTaskBriefInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskBriefInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskBriefInfoResponseBody) SetRisksFacets(v *DescribeTaskBriefInfoResponseBodyRisksFacets) *DescribeTaskBriefInfoResponseBody {
	s.RisksFacets = v
	return s
}

func (s *DescribeTaskBriefInfoResponseBody) SetRequestId(v string) *DescribeTaskBriefInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBody) SetBrief(v *DescribeTaskBriefInfoResponseBodyBrief) *DescribeTaskBriefInfoResponseBody {
	s.Brief = v
	return s
}

type DescribeTaskBriefInfoResponseBodyRisksFacets struct {
	Medium *int32 `json:"Medium,omitempty" xml:"Medium,omitempty"`
	Low    *int32 `json:"Low,omitempty" xml:"Low,omitempty"`
	Total  *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	High   *int32 `json:"High,omitempty" xml:"High,omitempty"`
	Info   *int32 `json:"Info,omitempty" xml:"Info,omitempty"`
}

func (s DescribeTaskBriefInfoResponseBodyRisksFacets) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskBriefInfoResponseBodyRisksFacets) GoString() string {
	return s.String()
}

func (s *DescribeTaskBriefInfoResponseBodyRisksFacets) SetMedium(v int32) *DescribeTaskBriefInfoResponseBodyRisksFacets {
	s.Medium = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyRisksFacets) SetLow(v int32) *DescribeTaskBriefInfoResponseBodyRisksFacets {
	s.Low = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyRisksFacets) SetTotal(v int32) *DescribeTaskBriefInfoResponseBodyRisksFacets {
	s.Total = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyRisksFacets) SetHigh(v int32) *DescribeTaskBriefInfoResponseBodyRisksFacets {
	s.High = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyRisksFacets) SetInfo(v int32) *DescribeTaskBriefInfoResponseBodyRisksFacets {
	s.Info = &v
	return s
}

type DescribeTaskBriefInfoResponseBodyBrief struct {
	FinishedAt      *int64                                                 `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Progress        *float32                                               `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Description     *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedAt       *string                                                `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ActionRunsFacet *DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet `json:"ActionRunsFacet,omitempty" xml:"ActionRunsFacet,omitempty" type:"Struct"`
}

func (s DescribeTaskBriefInfoResponseBodyBrief) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskBriefInfoResponseBodyBrief) GoString() string {
	return s.String()
}

func (s *DescribeTaskBriefInfoResponseBodyBrief) SetFinishedAt(v int64) *DescribeTaskBriefInfoResponseBodyBrief {
	s.FinishedAt = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyBrief) SetProgress(v float32) *DescribeTaskBriefInfoResponseBodyBrief {
	s.Progress = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyBrief) SetDescription(v string) *DescribeTaskBriefInfoResponseBodyBrief {
	s.Description = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyBrief) SetCreatedAt(v string) *DescribeTaskBriefInfoResponseBodyBrief {
	s.CreatedAt = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyBrief) SetActionRunsFacet(v *DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet) *DescribeTaskBriefInfoResponseBodyBrief {
	s.ActionRunsFacet = v
	return s
}

type DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet struct {
	Completed *int32 `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet) GoString() string {
	return s.String()
}

func (s *DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet) SetCompleted(v int32) *DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet {
	s.Completed = &v
	return s
}

func (s *DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet) SetTotal(v int32) *DescribeTaskBriefInfoResponseBodyBriefActionRunsFacet {
	s.Total = &v
	return s
}

type DescribeTaskBriefInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTaskBriefInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTaskBriefInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskBriefInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskBriefInfoResponse) SetHeaders(v map[string]*string) *DescribeTaskBriefInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskBriefInfoResponse) SetBody(v *DescribeTaskBriefInfoResponseBody) *DescribeTaskBriefInfoResponse {
	s.Body = v
	return s
}

type DescribeUserTagsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Search          *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsRequest) SetSourceIp(v string) *DescribeUserTagsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserTagsRequest) SetResourceOwnerId(v int64) *DescribeUserTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserTagsRequest) SetLang(v string) *DescribeUserTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserTagsRequest) SetSearch(v string) *DescribeUserTagsRequest {
	s.Search = &v
	return s
}

func (s *DescribeUserTagsRequest) SetCurrentPage(v int32) *DescribeUserTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserTagsRequest) SetPageSize(v int32) *DescribeUserTagsRequest {
	s.PageSize = &v
	return s
}

type DescribeUserTagsResponseBody struct {
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount   *int32                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*DescribeUserTagsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Count       *int32                              `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponseBody) SetTotalCount(v int32) *DescribeUserTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetRequestId(v string) *DescribeUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetPageSize(v int32) *DescribeUserTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetPageCount(v int32) *DescribeUserTagsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetCurrentPage(v int32) *DescribeUserTagsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetList(v []*DescribeUserTagsResponseBodyList) *DescribeUserTagsResponseBody {
	s.List = v
	return s
}

func (s *DescribeUserTagsResponseBody) SetCount(v int32) *DescribeUserTagsResponseBody {
	s.Count = &v
	return s
}

type DescribeUserTagsResponseBodyList struct {
	AssetCount *int32  `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeUserTagsResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponseBodyList) SetAssetCount(v int32) *DescribeUserTagsResponseBodyList {
	s.AssetCount = &v
	return s
}

func (s *DescribeUserTagsResponseBodyList) SetTagKey(v string) *DescribeUserTagsResponseBodyList {
	s.TagKey = &v
	return s
}

type DescribeUserTagsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponse) SetHeaders(v map[string]*string) *DescribeUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTagsResponse) SetBody(v *DescribeUserTagsResponseBody) *DescribeUserTagsResponse {
	s.Body = v
	return s
}

type DescribeVulnerabilityRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeVulnerabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityRequest) SetSourceIp(v string) *DescribeVulnerabilityRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulnerabilityRequest) SetResourceOwnerId(v int64) *DescribeVulnerabilityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVulnerabilityRequest) SetId(v int64) *DescribeVulnerabilityRequest {
	s.Id = &v
	return s
}

func (s *DescribeVulnerabilityRequest) SetLang(v string) *DescribeVulnerabilityRequest {
	s.Lang = &v
	return s
}

type DescribeVulnerabilityResponseBody struct {
	Impact           *string                                         `json:"Impact,omitempty" xml:"Impact,omitempty"`
	Args             *string                                         `json:"Args,omitempty" xml:"Args,omitempty"`
	Description      *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LastDiscoveredAt *int64                                          `json:"LastDiscoveredAt,omitempty" xml:"LastDiscoveredAt,omitempty"`
	Poc              *string                                         `json:"Poc,omitempty" xml:"Poc,omitempty"`
	Reference        *string                                         `json:"Reference,omitempty" xml:"Reference,omitempty"`
	Hostname         *string                                         `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	Severity         *int32                                          `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Data             *string                                         `json:"Data,omitempty" xml:"Data,omitempty"`
	Vulnerability    *DescribeVulnerabilityResponseBodyVulnerability `json:"Vulnerability,omitempty" xml:"Vulnerability,omitempty" type:"Struct"`
	Success          *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Name             *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Target           *string                                         `json:"Target,omitempty" xml:"Target,omitempty"`
	Id               *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Solution         *string                                         `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s DescribeVulnerabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityResponseBody) SetImpact(v string) *DescribeVulnerabilityResponseBody {
	s.Impact = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetArgs(v string) *DescribeVulnerabilityResponseBody {
	s.Args = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetDescription(v string) *DescribeVulnerabilityResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetRequestId(v string) *DescribeVulnerabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetLastDiscoveredAt(v int64) *DescribeVulnerabilityResponseBody {
	s.LastDiscoveredAt = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetPoc(v string) *DescribeVulnerabilityResponseBody {
	s.Poc = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetReference(v string) *DescribeVulnerabilityResponseBody {
	s.Reference = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetHostname(v string) *DescribeVulnerabilityResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetSeverity(v int32) *DescribeVulnerabilityResponseBody {
	s.Severity = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetData(v string) *DescribeVulnerabilityResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetVulnerability(v *DescribeVulnerabilityResponseBodyVulnerability) *DescribeVulnerabilityResponseBody {
	s.Vulnerability = v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetSuccess(v bool) *DescribeVulnerabilityResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetName(v string) *DescribeVulnerabilityResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetTarget(v string) *DescribeVulnerabilityResponseBody {
	s.Target = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetId(v int64) *DescribeVulnerabilityResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeVulnerabilityResponseBody) SetSolution(v string) *DescribeVulnerabilityResponseBody {
	s.Solution = &v
	return s
}

type DescribeVulnerabilityResponseBodyVulnerability struct {
	Severity         *int32  `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Data             *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Args             *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Owasp            *string `json:"Owasp,omitempty" xml:"Owasp,omitempty"`
	Impact           *string `json:"Impact,omitempty" xml:"Impact,omitempty"`
	Cwe              *string `json:"Cwe,omitempty" xml:"Cwe,omitempty"`
	Hostname         *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	Reference        *string `json:"Reference,omitempty" xml:"Reference,omitempty"`
	Wasc             *string `json:"Wasc,omitempty" xml:"Wasc,omitempty"`
	LastDiscoveredAt *int64  `json:"LastDiscoveredAt,omitempty" xml:"LastDiscoveredAt,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CommonType       *string `json:"CommonType,omitempty" xml:"CommonType,omitempty"`
	Solution         *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Target           *string `json:"Target,omitempty" xml:"Target,omitempty"`
	Poc              *string `json:"Poc,omitempty" xml:"Poc,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeVulnerabilityResponseBodyVulnerability) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityResponseBodyVulnerability) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetSeverity(v int32) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Severity = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetStatus(v int32) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Status = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetData(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Data = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetArgs(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Args = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetOwasp(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Owasp = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetImpact(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Impact = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetCwe(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Cwe = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetHostname(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Hostname = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetReference(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Reference = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetWasc(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Wasc = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetLastDiscoveredAt(v int64) *DescribeVulnerabilityResponseBodyVulnerability {
	s.LastDiscoveredAt = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetDescription(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Description = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetCommonType(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.CommonType = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetSolution(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Solution = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetName(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Name = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetTarget(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Target = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetPoc(v string) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Poc = &v
	return s
}

func (s *DescribeVulnerabilityResponseBodyVulnerability) SetId(v int64) *DescribeVulnerabilityResponseBodyVulnerability {
	s.Id = &v
	return s
}

type DescribeVulnerabilityResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulnerabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulnerabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityResponse) SetHeaders(v map[string]*string) *DescribeVulnerabilityResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulnerabilityResponse) SetBody(v *DescribeVulnerabilityResponseBody) *DescribeVulnerabilityResponse {
	s.Body = v
	return s
}

type DescribeWebPathsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPathsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebPathsRequest) SetSourceIp(v string) *DescribeWebPathsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebPathsRequest) SetTaskId(v int32) *DescribeWebPathsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeWebPathsRequest) SetAsset(v string) *DescribeWebPathsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeWebPathsRequest) SetCurrentPage(v int32) *DescribeWebPathsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebPathsRequest) SetPageSize(v int32) *DescribeWebPathsRequest {
	s.PageSize = &v
	return s
}

type DescribeWebPathsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeWebPathsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeWebPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPathsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebPathsResponseBody) SetRequestId(v string) *DescribeWebPathsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebPathsResponseBody) SetTotal(v int32) *DescribeWebPathsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeWebPathsResponseBody) SetRecords(v []*DescribeWebPathsResponseBodyRecords) *DescribeWebPathsResponseBody {
	s.Records = v
	return s
}

type DescribeWebPathsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Site      *string `json:"Site,omitempty" xml:"Site,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeWebPathsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPathsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeWebPathsResponseBodyRecords) SetIndex(v int32) *DescribeWebPathsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeWebPathsResponseBodyRecords) SetExtension(v string) *DescribeWebPathsResponseBodyRecords {
	s.Extension = &v
	return s
}

func (s *DescribeWebPathsResponseBodyRecords) SetSite(v string) *DescribeWebPathsResponseBodyRecords {
	s.Site = &v
	return s
}

func (s *DescribeWebPathsResponseBodyRecords) SetPath(v string) *DescribeWebPathsResponseBodyRecords {
	s.Path = &v
	return s
}

func (s *DescribeWebPathsResponseBodyRecords) SetUpdatedAt(v int64) *DescribeWebPathsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

type DescribeWebPathsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebPathsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPathsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebPathsResponse) SetHeaders(v map[string]*string) *DescribeWebPathsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebPathsResponse) SetBody(v *DescribeWebPathsResponseBody) *DescribeWebPathsResponse {
	s.Body = v
	return s
}

type DescribeWebServersRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebServersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebServersRequest) SetSourceIp(v string) *DescribeWebServersRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebServersRequest) SetTaskId(v int32) *DescribeWebServersRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeWebServersRequest) SetAsset(v string) *DescribeWebServersRequest {
	s.Asset = &v
	return s
}

func (s *DescribeWebServersRequest) SetCurrentPage(v int32) *DescribeWebServersRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebServersRequest) SetPageSize(v int32) *DescribeWebServersRequest {
	s.PageSize = &v
	return s
}

type DescribeWebServersResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeWebServersResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeWebServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebServersResponseBody) SetRequestId(v string) *DescribeWebServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebServersResponseBody) SetTotal(v int32) *DescribeWebServersResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeWebServersResponseBody) SetRecords(v []*DescribeWebServersResponseBodyRecords) *DescribeWebServersResponseBody {
	s.Records = v
	return s
}

type DescribeWebServersResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Scheme    *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeWebServersResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebServersResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeWebServersResponseBodyRecords) SetIndex(v int32) *DescribeWebServersResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeWebServersResponseBodyRecords) SetHost(v string) *DescribeWebServersResponseBodyRecords {
	s.Host = &v
	return s
}

func (s *DescribeWebServersResponseBodyRecords) SetScheme(v string) *DescribeWebServersResponseBodyRecords {
	s.Scheme = &v
	return s
}

func (s *DescribeWebServersResponseBodyRecords) SetUpdatedAt(v int64) *DescribeWebServersResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeWebServersResponseBodyRecords) SetPort(v string) *DescribeWebServersResponseBodyRecords {
	s.Port = &v
	return s
}

type DescribeWebServersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebServersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebServersResponse) SetHeaders(v map[string]*string) *DescribeWebServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebServersResponse) SetBody(v *DescribeWebServersResponseBody) *DescribeWebServersResponse {
	s.Body = v
	return s
}

type DescribeWebTechsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId      *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebTechsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebTechsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebTechsRequest) SetSourceIp(v string) *DescribeWebTechsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebTechsRequest) SetTaskId(v int32) *DescribeWebTechsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeWebTechsRequest) SetAsset(v string) *DescribeWebTechsRequest {
	s.Asset = &v
	return s
}

func (s *DescribeWebTechsRequest) SetCurrentPage(v int32) *DescribeWebTechsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebTechsRequest) SetPageSize(v int32) *DescribeWebTechsRequest {
	s.PageSize = &v
	return s
}

type DescribeWebTechsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	Records   []*DescribeWebTechsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeWebTechsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebTechsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebTechsResponseBody) SetRequestId(v string) *DescribeWebTechsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebTechsResponseBody) SetTotal(v int32) *DescribeWebTechsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeWebTechsResponseBody) SetRecords(v []*DescribeWebTechsResponseBodyRecords) *DescribeWebTechsResponseBody {
	s.Records = v
	return s
}

type DescribeWebTechsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	PoweredBy *string `json:"PoweredBy,omitempty" xml:"PoweredBy,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	URL       *string `json:"URL,omitempty" xml:"URL,omitempty"`
	Server    *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeWebTechsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebTechsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeWebTechsResponseBodyRecords) SetIndex(v int32) *DescribeWebTechsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetPoweredBy(v string) *DescribeWebTechsResponseBodyRecords {
	s.PoweredBy = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetVersion(v string) *DescribeWebTechsResponseBodyRecords {
	s.Version = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetURL(v string) *DescribeWebTechsResponseBodyRecords {
	s.URL = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetServer(v string) *DescribeWebTechsResponseBodyRecords {
	s.Server = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetUpdatedAt(v int64) *DescribeWebTechsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetTitle(v string) *DescribeWebTechsResponseBodyRecords {
	s.Title = &v
	return s
}

func (s *DescribeWebTechsResponseBodyRecords) SetName(v string) *DescribeWebTechsResponseBodyRecords {
	s.Name = &v
	return s
}

type DescribeWebTechsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebTechsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebTechsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebTechsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebTechsResponse) SetHeaders(v map[string]*string) *DescribeWebTechsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebTechsResponse) SetBody(v *DescribeWebTechsResponseBody) *DescribeWebTechsResponse {
	s.Body = v
	return s
}

type EditSessionRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SessionId    *int32  `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Asset        *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	TTL          *int32  `json:"TTL,omitempty" xml:"TTL,omitempty"`
	LoginSession *string `json:"LoginSession,omitempty" xml:"LoginSession,omitempty"`
}

func (s EditSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s EditSessionRequest) GoString() string {
	return s.String()
}

func (s *EditSessionRequest) SetSourceIp(v string) *EditSessionRequest {
	s.SourceIp = &v
	return s
}

func (s *EditSessionRequest) SetSessionId(v int32) *EditSessionRequest {
	s.SessionId = &v
	return s
}

func (s *EditSessionRequest) SetAsset(v string) *EditSessionRequest {
	s.Asset = &v
	return s
}

func (s *EditSessionRequest) SetTTL(v int32) *EditSessionRequest {
	s.TTL = &v
	return s
}

func (s *EditSessionRequest) SetLoginSession(v string) *EditSessionRequest {
	s.LoginSession = &v
	return s
}

type EditSessionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EditSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditSessionResponseBody) GoString() string {
	return s.String()
}

func (s *EditSessionResponseBody) SetRequestId(v string) *EditSessionResponseBody {
	s.RequestId = &v
	return s
}

type EditSessionResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s EditSessionResponse) GoString() string {
	return s.String()
}

func (s *EditSessionResponse) SetHeaders(v map[string]*string) *EditSessionResponse {
	s.Headers = v
	return s
}

func (s *EditSessionResponse) SetBody(v *EditSessionResponseBody) *EditSessionResponse {
	s.Body = v
	return s
}

type GenerateVulReportRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Format          *string   `json:"format,omitempty" xml:"format,omitempty"`
	Lang            *string   `json:"lang,omitempty" xml:"lang,omitempty"`
	TaskIds         []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s GenerateVulReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVulReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateVulReportRequest) SetSourceIp(v string) *GenerateVulReportRequest {
	s.SourceIp = &v
	return s
}

func (s *GenerateVulReportRequest) SetResourceOwnerId(v int64) *GenerateVulReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GenerateVulReportRequest) SetFormat(v string) *GenerateVulReportRequest {
	s.Format = &v
	return s
}

func (s *GenerateVulReportRequest) SetLang(v string) *GenerateVulReportRequest {
	s.Lang = &v
	return s
}

func (s *GenerateVulReportRequest) SetTaskIds(v []*string) *GenerateVulReportRequest {
	s.TaskIds = v
	return s
}

type GenerateVulReportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateVulReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVulReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVulReportResponseBody) SetRequestId(v string) *GenerateVulReportResponseBody {
	s.RequestId = &v
	return s
}

type GenerateVulReportResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateVulReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVulReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVulReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateVulReportResponse) SetHeaders(v map[string]*string) *GenerateVulReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateVulReportResponse) SetBody(v *GenerateVulReportResponseBody) *GenerateVulReportResponse {
	s.Body = v
	return s
}

type ListOrgDNSMapRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgDNSMapRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDNSMapRequest) GoString() string {
	return s.String()
}

func (s *ListOrgDNSMapRequest) SetSourceIp(v string) *ListOrgDNSMapRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgDNSMapRequest) SetOrgId(v int32) *ListOrgDNSMapRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgDNSMapRequest) SetSearch(v string) *ListOrgDNSMapRequest {
	s.Search = &v
	return s
}

func (s *ListOrgDNSMapRequest) SetCurrentPage(v int32) *ListOrgDNSMapRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgDNSMapRequest) SetPageSize(v int32) *ListOrgDNSMapRequest {
	s.PageSize = &v
	return s
}

type ListOrgDNSMapResponseBody struct {
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgDNSMapResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgDNSMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDNSMapResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgDNSMapResponseBody) SetTotalCount(v int32) *ListOrgDNSMapResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgDNSMapResponseBody) SetPageSize(v int32) *ListOrgDNSMapResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgDNSMapResponseBody) SetRequestId(v string) *ListOrgDNSMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgDNSMapResponseBody) SetCurrentPage(v int32) *ListOrgDNSMapResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgDNSMapResponseBody) SetRecords(v []*ListOrgDNSMapResponseBodyRecords) *ListOrgDNSMapResponseBody {
	s.Records = v
	return s
}

type ListOrgDNSMapResponseBodyRecords struct {
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OrgId     *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Record    *string `json:"Record,omitempty" xml:"Record,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgDNSMapResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDNSMapResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgDNSMapResponseBodyRecords) SetType(v string) *ListOrgDNSMapResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetIndex(v int32) *ListOrgDNSMapResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetDomain(v string) *ListOrgDNSMapResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetOrgId(v int32) *ListOrgDNSMapResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetSubdomain(v string) *ListOrgDNSMapResponseBodyRecords {
	s.Subdomain = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgDNSMapResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetRecord(v string) *ListOrgDNSMapResponseBodyRecords {
	s.Record = &v
	return s
}

func (s *ListOrgDNSMapResponseBodyRecords) SetId(v int32) *ListOrgDNSMapResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgDNSMapResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgDNSMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgDNSMapResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDNSMapResponse) GoString() string {
	return s.String()
}

func (s *ListOrgDNSMapResponse) SetHeaders(v map[string]*string) *ListOrgDNSMapResponse {
	s.Headers = v
	return s
}

func (s *ListOrgDNSMapResponse) SetBody(v *ListOrgDNSMapResponseBody) *ListOrgDNSMapResponse {
	s.Body = v
	return s
}

type ListOrgDomainsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgDomainsRequest) SetSourceIp(v string) *ListOrgDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgDomainsRequest) SetSearch(v string) *ListOrgDomainsRequest {
	s.Search = &v
	return s
}

func (s *ListOrgDomainsRequest) SetOrgId(v int32) *ListOrgDomainsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgDomainsRequest) SetCurrentPage(v int32) *ListOrgDomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgDomainsRequest) SetPageSize(v int32) *ListOrgDomainsRequest {
	s.PageSize = &v
	return s
}

type ListOrgDomainsResponseBody struct {
	TotalCount  *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgDomainsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgDomainsResponseBody) SetTotalCount(v int32) *ListOrgDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgDomainsResponseBody) SetPageSize(v int32) *ListOrgDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgDomainsResponseBody) SetRequestId(v string) *ListOrgDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgDomainsResponseBody) SetCurrentPage(v int32) *ListOrgDomainsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgDomainsResponseBody) SetRecords(v []*ListOrgDomainsResponseBodyRecords) *ListOrgDomainsResponseBody {
	s.Records = v
	return s
}

type ListOrgDomainsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OrgId     *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgDomainsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDomainsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgDomainsResponseBodyRecords) SetIndex(v int32) *ListOrgDomainsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgDomainsResponseBodyRecords) SetDomain(v string) *ListOrgDomainsResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListOrgDomainsResponseBodyRecords) SetOrgId(v int32) *ListOrgDomainsResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgDomainsResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgDomainsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgDomainsResponseBodyRecords) SetId(v int32) *ListOrgDomainsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgDomainsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgDomainsResponse) SetHeaders(v map[string]*string) *ListOrgDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgDomainsResponse) SetBody(v *ListOrgDomainsResponseBody) *ListOrgDomainsResponse {
	s.Body = v
	return s
}

type ListOrgHostsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgHostsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgHostsRequest) SetSourceIp(v string) *ListOrgHostsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgHostsRequest) SetSearch(v string) *ListOrgHostsRequest {
	s.Search = &v
	return s
}

func (s *ListOrgHostsRequest) SetOrgId(v int32) *ListOrgHostsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgHostsRequest) SetCurrentPage(v int32) *ListOrgHostsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgHostsRequest) SetPageSize(v int32) *ListOrgHostsRequest {
	s.PageSize = &v
	return s
}

type ListOrgHostsResponseBody struct {
	TotalCount  *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgHostsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgHostsResponseBody) SetTotalCount(v int32) *ListOrgHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgHostsResponseBody) SetPageSize(v int32) *ListOrgHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgHostsResponseBody) SetRequestId(v string) *ListOrgHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgHostsResponseBody) SetCurrentPage(v int32) *ListOrgHostsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgHostsResponseBody) SetRecords(v []*ListOrgHostsResponseBodyRecords) *ListOrgHostsResponseBody {
	s.Records = v
	return s
}

type ListOrgHostsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	OrgId     *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OS        *string `json:"OS,omitempty" xml:"OS,omitempty"`
	Hostname  *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgHostsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgHostsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgHostsResponseBodyRecords) SetIndex(v int32) *ListOrgHostsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetState(v string) *ListOrgHostsResponseBodyRecords {
	s.State = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetOrgId(v int32) *ListOrgHostsResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetOS(v string) *ListOrgHostsResponseBodyRecords {
	s.OS = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetHostname(v string) *ListOrgHostsResponseBodyRecords {
	s.Hostname = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgHostsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetIP(v string) *ListOrgHostsResponseBodyRecords {
	s.IP = &v
	return s
}

func (s *ListOrgHostsResponseBodyRecords) SetId(v int32) *ListOrgHostsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgHostsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgHostsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgHostsResponse) SetHeaders(v map[string]*string) *ListOrgHostsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgHostsResponse) SetBody(v *ListOrgHostsResponseBody) *ListOrgHostsResponse {
	s.Body = v
	return s
}

type ListOrgPortsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgPortsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgPortsRequest) SetSourceIp(v string) *ListOrgPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgPortsRequest) SetSearch(v string) *ListOrgPortsRequest {
	s.Search = &v
	return s
}

func (s *ListOrgPortsRequest) SetOrgId(v int32) *ListOrgPortsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgPortsRequest) SetCurrentPage(v int32) *ListOrgPortsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgPortsRequest) SetPageSize(v int32) *ListOrgPortsRequest {
	s.PageSize = &v
	return s
}

type ListOrgPortsResponseBody struct {
	TotalCount  *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgPortsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgPortsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgPortsResponseBody) SetTotalCount(v int32) *ListOrgPortsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgPortsResponseBody) SetPageSize(v int32) *ListOrgPortsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgPortsResponseBody) SetRequestId(v string) *ListOrgPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgPortsResponseBody) SetCurrentPage(v int32) *ListOrgPortsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgPortsResponseBody) SetRecords(v []*ListOrgPortsResponseBodyRecords) *ListOrgPortsResponseBody {
	s.Records = v
	return s
}

type ListOrgPortsResponseBodyRecords struct {
	Service     *string `json:"Service,omitempty" xml:"Service,omitempty"`
	Index       *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UpdatedAt   *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	IP          *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Id          *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgPortsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgPortsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgPortsResponseBodyRecords) SetService(v string) *ListOrgPortsResponseBodyRecords {
	s.Service = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetIndex(v int32) *ListOrgPortsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetFingerprint(v string) *ListOrgPortsResponseBodyRecords {
	s.Fingerprint = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetVersion(v string) *ListOrgPortsResponseBodyRecords {
	s.Version = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetProduct(v string) *ListOrgPortsResponseBodyRecords {
	s.Product = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetProtocol(v string) *ListOrgPortsResponseBodyRecords {
	s.Protocol = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetOrgId(v int32) *ListOrgPortsResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgPortsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetPort(v int32) *ListOrgPortsResponseBodyRecords {
	s.Port = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetIP(v string) *ListOrgPortsResponseBodyRecords {
	s.IP = &v
	return s
}

func (s *ListOrgPortsResponseBodyRecords) SetId(v int32) *ListOrgPortsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgPortsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgPortsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgPortsResponse) SetHeaders(v map[string]*string) *ListOrgPortsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgPortsResponse) SetBody(v *ListOrgPortsResponseBody) *ListOrgPortsResponse {
	s.Body = v
	return s
}

type ListOrgRiskyAssetsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId    *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s ListOrgRiskyAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgRiskyAssetsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgRiskyAssetsRequest) SetSourceIp(v string) *ListOrgRiskyAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgRiskyAssetsRequest) SetOrgId(v int32) *ListOrgRiskyAssetsRequest {
	s.OrgId = &v
	return s
}

type ListOrgRiskyAssetsResponseBody struct {
	TotalCount  *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Assets      []*ListOrgRiskyAssetsResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	Total       *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOrgRiskyAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgRiskyAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgRiskyAssetsResponseBody) SetTotalCount(v int32) *ListOrgRiskyAssetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgRiskyAssetsResponseBody) SetPageSize(v int32) *ListOrgRiskyAssetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgRiskyAssetsResponseBody) SetRequestId(v string) *ListOrgRiskyAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgRiskyAssetsResponseBody) SetCurrentPage(v int32) *ListOrgRiskyAssetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgRiskyAssetsResponseBody) SetAssets(v []*ListOrgRiskyAssetsResponseBodyAssets) *ListOrgRiskyAssetsResponseBody {
	s.Assets = v
	return s
}

func (s *ListOrgRiskyAssetsResponseBody) SetTotal(v int32) *ListOrgRiskyAssetsResponseBody {
	s.Total = &v
	return s
}

type ListOrgRiskyAssetsResponseBodyAssets struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Asset *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListOrgRiskyAssetsResponseBodyAssets) String() string {
	return tea.Prettify(s)
}

func (s ListOrgRiskyAssetsResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *ListOrgRiskyAssetsResponseBodyAssets) SetType(v string) *ListOrgRiskyAssetsResponseBodyAssets {
	s.Type = &v
	return s
}

func (s *ListOrgRiskyAssetsResponseBodyAssets) SetAsset(v string) *ListOrgRiskyAssetsResponseBodyAssets {
	s.Asset = &v
	return s
}

func (s *ListOrgRiskyAssetsResponseBodyAssets) SetCount(v int32) *ListOrgRiskyAssetsResponseBodyAssets {
	s.Count = &v
	return s
}

type ListOrgRiskyAssetsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgRiskyAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgRiskyAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgRiskyAssetsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgRiskyAssetsResponse) SetHeaders(v map[string]*string) *ListOrgRiskyAssetsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgRiskyAssetsResponse) SetBody(v *ListOrgRiskyAssetsResponseBody) *ListOrgRiskyAssetsResponse {
	s.Body = v
	return s
}

type ListOrgSubdomainsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgSubdomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgSubdomainsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgSubdomainsRequest) SetSourceIp(v string) *ListOrgSubdomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgSubdomainsRequest) SetSearch(v string) *ListOrgSubdomainsRequest {
	s.Search = &v
	return s
}

func (s *ListOrgSubdomainsRequest) SetOrgId(v int32) *ListOrgSubdomainsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgSubdomainsRequest) SetCurrentPage(v int32) *ListOrgSubdomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgSubdomainsRequest) SetPageSize(v int32) *ListOrgSubdomainsRequest {
	s.PageSize = &v
	return s
}

type ListOrgSubdomainsResponseBody struct {
	TotalCount  *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgSubdomainsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgSubdomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgSubdomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgSubdomainsResponseBody) SetTotalCount(v int32) *ListOrgSubdomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgSubdomainsResponseBody) SetPageSize(v int32) *ListOrgSubdomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgSubdomainsResponseBody) SetRequestId(v string) *ListOrgSubdomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgSubdomainsResponseBody) SetCurrentPage(v int32) *ListOrgSubdomainsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgSubdomainsResponseBody) SetRecords(v []*ListOrgSubdomainsResponseBodyRecords) *ListOrgSubdomainsResponseBody {
	s.Records = v
	return s
}

type ListOrgSubdomainsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OrgId     *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgSubdomainsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgSubdomainsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgSubdomainsResponseBodyRecords) SetIndex(v int32) *ListOrgSubdomainsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgSubdomainsResponseBodyRecords) SetDomain(v string) *ListOrgSubdomainsResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListOrgSubdomainsResponseBodyRecords) SetOrgId(v int32) *ListOrgSubdomainsResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgSubdomainsResponseBodyRecords) SetSubdomain(v string) *ListOrgSubdomainsResponseBodyRecords {
	s.Subdomain = &v
	return s
}

func (s *ListOrgSubdomainsResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgSubdomainsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgSubdomainsResponseBodyRecords) SetId(v int32) *ListOrgSubdomainsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgSubdomainsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgSubdomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgSubdomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgSubdomainsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgSubdomainsResponse) SetHeaders(v map[string]*string) *ListOrgSubdomainsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgSubdomainsResponse) SetBody(v *ListOrgSubdomainsResponseBody) *ListOrgSubdomainsResponse {
	s.Body = v
	return s
}

type ListOrgVulFacetsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgVulFacetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgVulFacetsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgVulFacetsRequest) SetSourceIp(v string) *ListOrgVulFacetsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgVulFacetsRequest) SetOrgId(v int32) *ListOrgVulFacetsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgVulFacetsRequest) SetAsset(v string) *ListOrgVulFacetsRequest {
	s.Asset = &v
	return s
}

func (s *ListOrgVulFacetsRequest) SetCurrentPage(v int32) *ListOrgVulFacetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgVulFacetsRequest) SetLang(v string) *ListOrgVulFacetsRequest {
	s.Lang = &v
	return s
}

func (s *ListOrgVulFacetsRequest) SetPageSize(v int32) *ListOrgVulFacetsRequest {
	s.PageSize = &v
	return s
}

type ListOrgVulFacetsResponseBody struct {
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Total       *int32                              `json:"Total,omitempty" xml:"Total,omitempty"`
	Vuls        []*ListOrgVulFacetsResponseBodyVuls `json:"Vuls,omitempty" xml:"Vuls,omitempty" type:"Repeated"`
}

func (s ListOrgVulFacetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgVulFacetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgVulFacetsResponseBody) SetTotalCount(v int32) *ListOrgVulFacetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgVulFacetsResponseBody) SetPageSize(v int32) *ListOrgVulFacetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgVulFacetsResponseBody) SetRequestId(v string) *ListOrgVulFacetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgVulFacetsResponseBody) SetCurrentPage(v int32) *ListOrgVulFacetsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgVulFacetsResponseBody) SetTotal(v int32) *ListOrgVulFacetsResponseBody {
	s.Total = &v
	return s
}

func (s *ListOrgVulFacetsResponseBody) SetVuls(v []*ListOrgVulFacetsResponseBodyVuls) *ListOrgVulFacetsResponseBody {
	s.Vuls = v
	return s
}

type ListOrgVulFacetsResponseBodyVuls struct {
	Index          *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Severity       *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	ModuleID       *int32  `json:"ModuleID,omitempty" xml:"ModuleID,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListOrgVulFacetsResponseBodyVuls) String() string {
	return tea.Prettify(s)
}

func (s ListOrgVulFacetsResponseBodyVuls) GoString() string {
	return s.String()
}

func (s *ListOrgVulFacetsResponseBodyVuls) SetIndex(v int32) *ListOrgVulFacetsResponseBodyVuls {
	s.Index = &v
	return s
}

func (s *ListOrgVulFacetsResponseBodyVuls) SetSeverity(v string) *ListOrgVulFacetsResponseBodyVuls {
	s.Severity = &v
	return s
}

func (s *ListOrgVulFacetsResponseBodyVuls) SetModuleID(v int32) *ListOrgVulFacetsResponseBodyVuls {
	s.ModuleID = &v
	return s
}

func (s *ListOrgVulFacetsResponseBodyVuls) SetClassification(v string) *ListOrgVulFacetsResponseBodyVuls {
	s.Classification = &v
	return s
}

func (s *ListOrgVulFacetsResponseBodyVuls) SetName(v string) *ListOrgVulFacetsResponseBodyVuls {
	s.Name = &v
	return s
}

func (s *ListOrgVulFacetsResponseBodyVuls) SetCount(v int32) *ListOrgVulFacetsResponseBodyVuls {
	s.Count = &v
	return s
}

type ListOrgVulFacetsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgVulFacetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgVulFacetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgVulFacetsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgVulFacetsResponse) SetHeaders(v map[string]*string) *ListOrgVulFacetsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgVulFacetsResponse) SetBody(v *ListOrgVulFacetsResponseBody) *ListOrgVulFacetsResponse {
	s.Body = v
	return s
}

type ListOrgWebPathsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgWebPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebPathsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgWebPathsRequest) SetSourceIp(v string) *ListOrgWebPathsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgWebPathsRequest) SetSearch(v string) *ListOrgWebPathsRequest {
	s.Search = &v
	return s
}

func (s *ListOrgWebPathsRequest) SetOrgId(v int32) *ListOrgWebPathsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgWebPathsRequest) SetCurrentPage(v int32) *ListOrgWebPathsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgWebPathsRequest) SetPageSize(v int32) *ListOrgWebPathsRequest {
	s.PageSize = &v
	return s
}

type ListOrgWebPathsResponseBody struct {
	TotalCount  *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgWebPathsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgWebPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebPathsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgWebPathsResponseBody) SetTotalCount(v int32) *ListOrgWebPathsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgWebPathsResponseBody) SetPageSize(v int32) *ListOrgWebPathsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgWebPathsResponseBody) SetRequestId(v string) *ListOrgWebPathsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgWebPathsResponseBody) SetCurrentPage(v int32) *ListOrgWebPathsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgWebPathsResponseBody) SetRecords(v []*ListOrgWebPathsResponseBodyRecords) *ListOrgWebPathsResponseBody {
	s.Records = v
	return s
}

type ListOrgWebPathsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	OrgId     *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Site      *string `json:"Site,omitempty" xml:"Site,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgWebPathsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebPathsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgWebPathsResponseBodyRecords) SetIndex(v int32) *ListOrgWebPathsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgWebPathsResponseBodyRecords) SetOrgId(v int32) *ListOrgWebPathsResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgWebPathsResponseBodyRecords) SetSite(v string) *ListOrgWebPathsResponseBodyRecords {
	s.Site = &v
	return s
}

func (s *ListOrgWebPathsResponseBodyRecords) SetPath(v string) *ListOrgWebPathsResponseBodyRecords {
	s.Path = &v
	return s
}

func (s *ListOrgWebPathsResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgWebPathsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgWebPathsResponseBodyRecords) SetId(v int32) *ListOrgWebPathsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgWebPathsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgWebPathsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgWebPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebPathsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgWebPathsResponse) SetHeaders(v map[string]*string) *ListOrgWebPathsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgWebPathsResponse) SetBody(v *ListOrgWebPathsResponseBody) *ListOrgWebPathsResponse {
	s.Body = v
	return s
}

type ListOrgWebTechsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrgWebTechsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebTechsRequest) GoString() string {
	return s.String()
}

func (s *ListOrgWebTechsRequest) SetSourceIp(v string) *ListOrgWebTechsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListOrgWebTechsRequest) SetSearch(v string) *ListOrgWebTechsRequest {
	s.Search = &v
	return s
}

func (s *ListOrgWebTechsRequest) SetOrgId(v int32) *ListOrgWebTechsRequest {
	s.OrgId = &v
	return s
}

func (s *ListOrgWebTechsRequest) SetCurrentPage(v int32) *ListOrgWebTechsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgWebTechsRequest) SetPageSize(v int32) *ListOrgWebTechsRequest {
	s.PageSize = &v
	return s
}

type ListOrgWebTechsResponseBody struct {
	TotalCount  *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListOrgWebTechsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListOrgWebTechsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebTechsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgWebTechsResponseBody) SetTotalCount(v int32) *ListOrgWebTechsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrgWebTechsResponseBody) SetPageSize(v int32) *ListOrgWebTechsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrgWebTechsResponseBody) SetRequestId(v string) *ListOrgWebTechsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrgWebTechsResponseBody) SetCurrentPage(v int32) *ListOrgWebTechsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOrgWebTechsResponseBody) SetRecords(v []*ListOrgWebTechsResponseBodyRecords) *ListOrgWebTechsResponseBody {
	s.Records = v
	return s
}

type ListOrgWebTechsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	PoweredBy *string `json:"PoweredBy,omitempty" xml:"PoweredBy,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	OrgId     *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	URL       *string `json:"URL,omitempty" xml:"URL,omitempty"`
	Server    *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOrgWebTechsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebTechsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListOrgWebTechsResponseBodyRecords) SetIndex(v int32) *ListOrgWebTechsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetPoweredBy(v string) *ListOrgWebTechsResponseBodyRecords {
	s.PoweredBy = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetVersion(v string) *ListOrgWebTechsResponseBodyRecords {
	s.Version = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetOrgId(v int32) *ListOrgWebTechsResponseBodyRecords {
	s.OrgId = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetURL(v string) *ListOrgWebTechsResponseBodyRecords {
	s.URL = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetServer(v string) *ListOrgWebTechsResponseBodyRecords {
	s.Server = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetUpdatedAt(v int64) *ListOrgWebTechsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetTitle(v string) *ListOrgWebTechsResponseBodyRecords {
	s.Title = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetName(v string) *ListOrgWebTechsResponseBodyRecords {
	s.Name = &v
	return s
}

func (s *ListOrgWebTechsResponseBodyRecords) SetId(v int32) *ListOrgWebTechsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListOrgWebTechsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrgWebTechsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrgWebTechsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgWebTechsResponse) GoString() string {
	return s.String()
}

func (s *ListOrgWebTechsResponse) SetHeaders(v map[string]*string) *ListOrgWebTechsResponse {
	s.Headers = v
	return s
}

func (s *ListOrgWebTechsResponse) SetBody(v *ListOrgWebTechsResponseBody) *ListOrgWebTechsResponse {
	s.Body = v
	return s
}

type ListUserDNSMapRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDNSMapRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapRequest) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapRequest) SetSourceIp(v string) *ListUserDNSMapRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserDNSMapRequest) SetSearch(v string) *ListUserDNSMapRequest {
	s.Search = &v
	return s
}

func (s *ListUserDNSMapRequest) SetCurrentPage(v int32) *ListUserDNSMapRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDNSMapRequest) SetPageSize(v int32) *ListUserDNSMapRequest {
	s.PageSize = &v
	return s
}

type ListUserDNSMapResponseBody struct {
	TotalCount  *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserDNSMapResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserDNSMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapResponseBody) SetTotalCount(v int32) *ListUserDNSMapResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserDNSMapResponseBody) SetPageSize(v int32) *ListUserDNSMapResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDNSMapResponseBody) SetRequestId(v string) *ListUserDNSMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDNSMapResponseBody) SetCurrentPage(v int32) *ListUserDNSMapResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDNSMapResponseBody) SetRecords(v []*ListUserDNSMapResponseBodyRecords) *ListUserDNSMapResponseBody {
	s.Records = v
	return s
}

type ListUserDNSMapResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Record    *string `json:"Record,omitempty" xml:"Record,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserDNSMapResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapResponseBodyRecords) SetIndex(v int32) *ListUserDNSMapResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserDNSMapResponseBodyRecords) SetType(v string) *ListUserDNSMapResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *ListUserDNSMapResponseBodyRecords) SetDomain(v string) *ListUserDNSMapResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListUserDNSMapResponseBodyRecords) SetSubdomain(v string) *ListUserDNSMapResponseBodyRecords {
	s.Subdomain = &v
	return s
}

func (s *ListUserDNSMapResponseBodyRecords) SetUpdatedAt(v int64) *ListUserDNSMapResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserDNSMapResponseBodyRecords) SetRecord(v string) *ListUserDNSMapResponseBodyRecords {
	s.Record = &v
	return s
}

func (s *ListUserDNSMapResponseBodyRecords) SetId(v int32) *ListUserDNSMapResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserDNSMapResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserDNSMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserDNSMapResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapResponse) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapResponse) SetHeaders(v map[string]*string) *ListUserDNSMapResponse {
	s.Headers = v
	return s
}

func (s *ListUserDNSMapResponse) SetBody(v *ListUserDNSMapResponseBody) *ListUserDNSMapResponse {
	s.Body = v
	return s
}

type ListUserDNSMapHistoriesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Id          *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDNSMapHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapHistoriesRequest) SetSourceIp(v string) *ListUserDNSMapHistoriesRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserDNSMapHistoriesRequest) SetId(v int32) *ListUserDNSMapHistoriesRequest {
	s.Id = &v
	return s
}

func (s *ListUserDNSMapHistoriesRequest) SetCurrentPage(v int32) *ListUserDNSMapHistoriesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDNSMapHistoriesRequest) SetPageSize(v int32) *ListUserDNSMapHistoriesRequest {
	s.PageSize = &v
	return s
}

type ListUserDNSMapHistoriesResponseBody struct {
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserDNSMapHistoriesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserDNSMapHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapHistoriesResponseBody) SetTotalCount(v int32) *ListUserDNSMapHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBody) SetPageSize(v int32) *ListUserDNSMapHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBody) SetRequestId(v string) *ListUserDNSMapHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBody) SetCurrentPage(v int32) *ListUserDNSMapHistoriesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBody) SetRecords(v []*ListUserDNSMapHistoriesResponseBodyRecords) *ListUserDNSMapHistoriesResponseBody {
	s.Records = v
	return s
}

type ListUserDNSMapHistoriesResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
	Record    *string `json:"Record,omitempty" xml:"Record,omitempty"`
	TaskId    *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListUserDNSMapHistoriesResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapHistoriesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetIndex(v int32) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetType(v string) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetDomain(v string) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetCreatedAt(v int64) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.CreatedAt = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetSubdomain(v string) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.Subdomain = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetRecord(v string) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.Record = &v
	return s
}

func (s *ListUserDNSMapHistoriesResponseBodyRecords) SetTaskId(v int32) *ListUserDNSMapHistoriesResponseBodyRecords {
	s.TaskId = &v
	return s
}

type ListUserDNSMapHistoriesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserDNSMapHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserDNSMapHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDNSMapHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListUserDNSMapHistoriesResponse) SetHeaders(v map[string]*string) *ListUserDNSMapHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListUserDNSMapHistoriesResponse) SetBody(v *ListUserDNSMapHistoriesResponseBody) *ListUserDNSMapHistoriesResponse {
	s.Body = v
	return s
}

type ListUserDomainsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListUserDomainsRequest) SetSourceIp(v string) *ListUserDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserDomainsRequest) SetSearch(v string) *ListUserDomainsRequest {
	s.Search = &v
	return s
}

func (s *ListUserDomainsRequest) SetCurrentPage(v int32) *ListUserDomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDomainsRequest) SetPageSize(v int32) *ListUserDomainsRequest {
	s.PageSize = &v
	return s
}

type ListUserDomainsResponseBody struct {
	TotalCount  *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserDomainsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDomainsResponseBody) SetTotalCount(v int32) *ListUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserDomainsResponseBody) SetPageSize(v int32) *ListUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDomainsResponseBody) SetRequestId(v string) *ListUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDomainsResponseBody) SetCurrentPage(v int32) *ListUserDomainsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDomainsResponseBody) SetRecords(v []*ListUserDomainsResponseBodyRecords) *ListUserDomainsResponseBody {
	s.Records = v
	return s
}

type ListUserDomainsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserDomainsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserDomainsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserDomainsResponseBodyRecords) SetIndex(v int32) *ListUserDomainsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserDomainsResponseBodyRecords) SetDomain(v string) *ListUserDomainsResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListUserDomainsResponseBodyRecords) SetUpdatedAt(v int64) *ListUserDomainsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserDomainsResponseBodyRecords) SetId(v int32) *ListUserDomainsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserDomainsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListUserDomainsResponse) SetHeaders(v map[string]*string) *ListUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListUserDomainsResponse) SetBody(v *ListUserDomainsResponseBody) *ListUserDomainsResponse {
	s.Body = v
	return s
}

type ListUserHostsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserHostsRequest) GoString() string {
	return s.String()
}

func (s *ListUserHostsRequest) SetSourceIp(v string) *ListUserHostsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserHostsRequest) SetSearch(v string) *ListUserHostsRequest {
	s.Search = &v
	return s
}

func (s *ListUserHostsRequest) SetCurrentPage(v int32) *ListUserHostsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserHostsRequest) SetPageSize(v int32) *ListUserHostsRequest {
	s.PageSize = &v
	return s
}

type ListUserHostsResponseBody struct {
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserHostsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserHostsResponseBody) SetTotalCount(v int32) *ListUserHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserHostsResponseBody) SetPageSize(v int32) *ListUserHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserHostsResponseBody) SetRequestId(v string) *ListUserHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserHostsResponseBody) SetCurrentPage(v int32) *ListUserHostsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserHostsResponseBody) SetRecords(v []*ListUserHostsResponseBodyRecords) *ListUserHostsResponseBody {
	s.Records = v
	return s
}

type ListUserHostsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	OS        *string `json:"OS,omitempty" xml:"OS,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	Hostname  *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	IP        *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserHostsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserHostsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserHostsResponseBodyRecords) SetIndex(v int32) *ListUserHostsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserHostsResponseBodyRecords) SetOS(v string) *ListUserHostsResponseBodyRecords {
	s.OS = &v
	return s
}

func (s *ListUserHostsResponseBodyRecords) SetState(v string) *ListUserHostsResponseBodyRecords {
	s.State = &v
	return s
}

func (s *ListUserHostsResponseBodyRecords) SetHostname(v string) *ListUserHostsResponseBodyRecords {
	s.Hostname = &v
	return s
}

func (s *ListUserHostsResponseBodyRecords) SetUpdatedAt(v int64) *ListUserHostsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserHostsResponseBodyRecords) SetIP(v string) *ListUserHostsResponseBodyRecords {
	s.IP = &v
	return s
}

func (s *ListUserHostsResponseBodyRecords) SetId(v int32) *ListUserHostsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserHostsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserHostsResponse) GoString() string {
	return s.String()
}

func (s *ListUserHostsResponse) SetHeaders(v map[string]*string) *ListUserHostsResponse {
	s.Headers = v
	return s
}

func (s *ListUserHostsResponse) SetBody(v *ListUserHostsResponseBody) *ListUserHostsResponse {
	s.Body = v
	return s
}

type ListUserOrganizationsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationsRequest) SetSourceIp(v string) *ListUserOrganizationsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserOrganizationsRequest) SetSearch(v string) *ListUserOrganizationsRequest {
	s.Search = &v
	return s
}

func (s *ListUserOrganizationsRequest) SetCurrentPage(v int32) *ListUserOrganizationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserOrganizationsRequest) SetPageSize(v int32) *ListUserOrganizationsRequest {
	s.PageSize = &v
	return s
}

type ListUserOrganizationsResponseBody struct {
	TotalCount    *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage   *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Organizations []*ListUserOrganizationsResponseBodyOrganizations `json:"Organizations,omitempty" xml:"Organizations,omitempty" type:"Repeated"`
}

func (s ListUserOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationsResponseBody) SetTotalCount(v int32) *ListUserOrganizationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserOrganizationsResponseBody) SetPageSize(v int32) *ListUserOrganizationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserOrganizationsResponseBody) SetRequestId(v string) *ListUserOrganizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserOrganizationsResponseBody) SetCurrentPage(v int32) *ListUserOrganizationsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserOrganizationsResponseBody) SetOrganizations(v []*ListUserOrganizationsResponseBodyOrganizations) *ListUserOrganizationsResponseBody {
	s.Organizations = v
	return s
}

type ListUserOrganizationsResponseBodyOrganizations struct {
	Index       *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	CreatedAt   *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt   *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUserOrganizationsResponseBodyOrganizations) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationsResponseBodyOrganizations) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationsResponseBodyOrganizations) SetIndex(v int32) *ListUserOrganizationsResponseBodyOrganizations {
	s.Index = &v
	return s
}

func (s *ListUserOrganizationsResponseBodyOrganizations) SetDescription(v string) *ListUserOrganizationsResponseBodyOrganizations {
	s.Description = &v
	return s
}

func (s *ListUserOrganizationsResponseBodyOrganizations) SetOrgId(v int32) *ListUserOrganizationsResponseBodyOrganizations {
	s.OrgId = &v
	return s
}

func (s *ListUserOrganizationsResponseBodyOrganizations) SetCreatedAt(v int64) *ListUserOrganizationsResponseBodyOrganizations {
	s.CreatedAt = &v
	return s
}

func (s *ListUserOrganizationsResponseBodyOrganizations) SetUpdatedAt(v int64) *ListUserOrganizationsResponseBodyOrganizations {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserOrganizationsResponseBodyOrganizations) SetName(v string) *ListUserOrganizationsResponseBodyOrganizations {
	s.Name = &v
	return s
}

type ListUserOrganizationsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationsResponse) SetHeaders(v map[string]*string) *ListUserOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *ListUserOrganizationsResponse) SetBody(v *ListUserOrganizationsResponseBody) *ListUserOrganizationsResponse {
	s.Body = v
	return s
}

type ListUserPortsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserPortsRequest) GoString() string {
	return s.String()
}

func (s *ListUserPortsRequest) SetSourceIp(v string) *ListUserPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserPortsRequest) SetSearch(v string) *ListUserPortsRequest {
	s.Search = &v
	return s
}

func (s *ListUserPortsRequest) SetCurrentPage(v int32) *ListUserPortsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserPortsRequest) SetPageSize(v int32) *ListUserPortsRequest {
	s.PageSize = &v
	return s
}

type ListUserPortsResponseBody struct {
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserPortsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserPortsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPortsResponseBody) SetTotalCount(v int32) *ListUserPortsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPortsResponseBody) SetPageSize(v int32) *ListUserPortsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserPortsResponseBody) SetRequestId(v string) *ListUserPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPortsResponseBody) SetCurrentPage(v int32) *ListUserPortsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserPortsResponseBody) SetRecords(v []*ListUserPortsResponseBodyRecords) *ListUserPortsResponseBody {
	s.Records = v
	return s
}

type ListUserPortsResponseBodyRecords struct {
	Service     *string `json:"Service,omitempty" xml:"Service,omitempty"`
	Index       *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	UpdatedAt   *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Port        *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	IP          *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Id          *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserPortsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserPortsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserPortsResponseBodyRecords) SetService(v string) *ListUserPortsResponseBodyRecords {
	s.Service = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetIndex(v int32) *ListUserPortsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetFingerprint(v string) *ListUserPortsResponseBodyRecords {
	s.Fingerprint = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetVersion(v string) *ListUserPortsResponseBodyRecords {
	s.Version = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetProduct(v string) *ListUserPortsResponseBodyRecords {
	s.Product = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetProtocol(v string) *ListUserPortsResponseBodyRecords {
	s.Protocol = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetUpdatedAt(v int64) *ListUserPortsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetPort(v int32) *ListUserPortsResponseBodyRecords {
	s.Port = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetIP(v string) *ListUserPortsResponseBodyRecords {
	s.IP = &v
	return s
}

func (s *ListUserPortsResponseBodyRecords) SetId(v int32) *ListUserPortsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserPortsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserPortsResponse) GoString() string {
	return s.String()
}

func (s *ListUserPortsResponse) SetHeaders(v map[string]*string) *ListUserPortsResponse {
	s.Headers = v
	return s
}

func (s *ListUserPortsResponse) SetBody(v *ListUserPortsResponseBody) *ListUserPortsResponse {
	s.Body = v
	return s
}

type ListUserSubdomainsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserSubdomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserSubdomainsRequest) GoString() string {
	return s.String()
}

func (s *ListUserSubdomainsRequest) SetSourceIp(v string) *ListUserSubdomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserSubdomainsRequest) SetSearch(v string) *ListUserSubdomainsRequest {
	s.Search = &v
	return s
}

func (s *ListUserSubdomainsRequest) SetCurrentPage(v int32) *ListUserSubdomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserSubdomainsRequest) SetPageSize(v int32) *ListUserSubdomainsRequest {
	s.PageSize = &v
	return s
}

type ListUserSubdomainsResponseBody struct {
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserSubdomainsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserSubdomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserSubdomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSubdomainsResponseBody) SetTotalCount(v int32) *ListUserSubdomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserSubdomainsResponseBody) SetPageSize(v int32) *ListUserSubdomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSubdomainsResponseBody) SetRequestId(v string) *ListUserSubdomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSubdomainsResponseBody) SetCurrentPage(v int32) *ListUserSubdomainsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserSubdomainsResponseBody) SetRecords(v []*ListUserSubdomainsResponseBodyRecords) *ListUserSubdomainsResponseBody {
	s.Records = v
	return s
}

type ListUserSubdomainsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Subdomain *string `json:"Subdomain,omitempty" xml:"Subdomain,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserSubdomainsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserSubdomainsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserSubdomainsResponseBodyRecords) SetIndex(v int32) *ListUserSubdomainsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserSubdomainsResponseBodyRecords) SetDomain(v string) *ListUserSubdomainsResponseBodyRecords {
	s.Domain = &v
	return s
}

func (s *ListUserSubdomainsResponseBodyRecords) SetSubdomain(v string) *ListUserSubdomainsResponseBodyRecords {
	s.Subdomain = &v
	return s
}

func (s *ListUserSubdomainsResponseBodyRecords) SetUpdatedAt(v int64) *ListUserSubdomainsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserSubdomainsResponseBodyRecords) SetId(v int32) *ListUserSubdomainsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserSubdomainsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserSubdomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserSubdomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserSubdomainsResponse) GoString() string {
	return s.String()
}

func (s *ListUserSubdomainsResponse) SetHeaders(v map[string]*string) *ListUserSubdomainsResponse {
	s.Headers = v
	return s
}

func (s *ListUserSubdomainsResponse) SetBody(v *ListUserSubdomainsResponseBody) *ListUserSubdomainsResponse {
	s.Body = v
	return s
}

type ListUserWebPathsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserWebPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebPathsRequest) GoString() string {
	return s.String()
}

func (s *ListUserWebPathsRequest) SetSourceIp(v string) *ListUserWebPathsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserWebPathsRequest) SetSearch(v string) *ListUserWebPathsRequest {
	s.Search = &v
	return s
}

func (s *ListUserWebPathsRequest) SetCurrentPage(v int32) *ListUserWebPathsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserWebPathsRequest) SetPageSize(v int32) *ListUserWebPathsRequest {
	s.PageSize = &v
	return s
}

type ListUserWebPathsResponseBody struct {
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserWebPathsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserWebPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebPathsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserWebPathsResponseBody) SetTotalCount(v int32) *ListUserWebPathsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserWebPathsResponseBody) SetPageSize(v int32) *ListUserWebPathsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserWebPathsResponseBody) SetRequestId(v string) *ListUserWebPathsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserWebPathsResponseBody) SetCurrentPage(v int32) *ListUserWebPathsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserWebPathsResponseBody) SetRecords(v []*ListUserWebPathsResponseBodyRecords) *ListUserWebPathsResponseBody {
	s.Records = v
	return s
}

type ListUserWebPathsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Site      *string `json:"Site,omitempty" xml:"Site,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserWebPathsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebPathsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserWebPathsResponseBodyRecords) SetIndex(v int32) *ListUserWebPathsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserWebPathsResponseBodyRecords) SetSite(v string) *ListUserWebPathsResponseBodyRecords {
	s.Site = &v
	return s
}

func (s *ListUserWebPathsResponseBodyRecords) SetPath(v string) *ListUserWebPathsResponseBodyRecords {
	s.Path = &v
	return s
}

func (s *ListUserWebPathsResponseBodyRecords) SetUpdatedAt(v int64) *ListUserWebPathsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserWebPathsResponseBodyRecords) SetId(v int32) *ListUserWebPathsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserWebPathsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserWebPathsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserWebPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebPathsResponse) GoString() string {
	return s.String()
}

func (s *ListUserWebPathsResponse) SetHeaders(v map[string]*string) *ListUserWebPathsResponse {
	s.Headers = v
	return s
}

func (s *ListUserWebPathsResponse) SetBody(v *ListUserWebPathsResponseBody) *ListUserWebPathsResponse {
	s.Body = v
	return s
}

type ListUserWebTechsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Search      *string `json:"Search,omitempty" xml:"Search,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserWebTechsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebTechsRequest) GoString() string {
	return s.String()
}

func (s *ListUserWebTechsRequest) SetSourceIp(v string) *ListUserWebTechsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserWebTechsRequest) SetSearch(v string) *ListUserWebTechsRequest {
	s.Search = &v
	return s
}

func (s *ListUserWebTechsRequest) SetCurrentPage(v int32) *ListUserWebTechsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserWebTechsRequest) SetPageSize(v int32) *ListUserWebTechsRequest {
	s.PageSize = &v
	return s
}

type ListUserWebTechsResponseBody struct {
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Records     []*ListUserWebTechsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s ListUserWebTechsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebTechsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserWebTechsResponseBody) SetTotalCount(v int32) *ListUserWebTechsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserWebTechsResponseBody) SetPageSize(v int32) *ListUserWebTechsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserWebTechsResponseBody) SetRequestId(v string) *ListUserWebTechsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserWebTechsResponseBody) SetCurrentPage(v int32) *ListUserWebTechsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserWebTechsResponseBody) SetRecords(v []*ListUserWebTechsResponseBodyRecords) *ListUserWebTechsResponseBody {
	s.Records = v
	return s
}

type ListUserWebTechsResponseBodyRecords struct {
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	PoweredBy *string `json:"PoweredBy,omitempty" xml:"PoweredBy,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
	URL       *string `json:"URL,omitempty" xml:"URL,omitempty"`
	Server    *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserWebTechsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebTechsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserWebTechsResponseBodyRecords) SetIndex(v int32) *ListUserWebTechsResponseBodyRecords {
	s.Index = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetPoweredBy(v string) *ListUserWebTechsResponseBodyRecords {
	s.PoweredBy = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetVersion(v string) *ListUserWebTechsResponseBodyRecords {
	s.Version = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetURL(v string) *ListUserWebTechsResponseBodyRecords {
	s.URL = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetServer(v string) *ListUserWebTechsResponseBodyRecords {
	s.Server = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetUpdatedAt(v int64) *ListUserWebTechsResponseBodyRecords {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetTitle(v string) *ListUserWebTechsResponseBodyRecords {
	s.Title = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetName(v string) *ListUserWebTechsResponseBodyRecords {
	s.Name = &v
	return s
}

func (s *ListUserWebTechsResponseBodyRecords) SetId(v int32) *ListUserWebTechsResponseBodyRecords {
	s.Id = &v
	return s
}

type ListUserWebTechsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserWebTechsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserWebTechsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserWebTechsResponse) GoString() string {
	return s.String()
}

func (s *ListUserWebTechsResponse) SetHeaders(v map[string]*string) *ListUserWebTechsResponse {
	s.Headers = v
	return s
}

func (s *ListUserWebTechsResponse) SetBody(v *ListUserWebTechsResponseBody) *ListUserWebTechsResponse {
	s.Body = v
	return s
}

type ModifyOrganizationRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OrgId       *int32  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s ModifyOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOrganizationRequest) GoString() string {
	return s.String()
}

func (s *ModifyOrganizationRequest) SetSourceIp(v string) *ModifyOrganizationRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyOrganizationRequest) SetName(v string) *ModifyOrganizationRequest {
	s.Name = &v
	return s
}

func (s *ModifyOrganizationRequest) SetDescription(v string) *ModifyOrganizationRequest {
	s.Description = &v
	return s
}

func (s *ModifyOrganizationRequest) SetOrgId(v int32) *ModifyOrganizationRequest {
	s.OrgId = &v
	return s
}

type ModifyOrganizationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOrganizationResponseBody) SetRequestId(v string) *ModifyOrganizationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOrganizationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOrganizationResponse) GoString() string {
	return s.String()
}

func (s *ModifyOrganizationResponse) SetHeaders(v map[string]*string) *ModifyOrganizationResponse {
	s.Headers = v
	return s
}

func (s *ModifyOrganizationResponse) SetBody(v *ModifyOrganizationResponseBody) *ModifyOrganizationResponse {
	s.Body = v
	return s
}

type TagAssetsByRecordsRequest struct {
	SourceIp  *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Source    *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	AssetType *string   `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	RecordIds []*int    `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
}

func (s TagAssetsByRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s TagAssetsByRecordsRequest) GoString() string {
	return s.String()
}

func (s *TagAssetsByRecordsRequest) SetSourceIp(v string) *TagAssetsByRecordsRequest {
	s.SourceIp = &v
	return s
}

func (s *TagAssetsByRecordsRequest) SetSource(v string) *TagAssetsByRecordsRequest {
	s.Source = &v
	return s
}

func (s *TagAssetsByRecordsRequest) SetAssetType(v string) *TagAssetsByRecordsRequest {
	s.AssetType = &v
	return s
}

func (s *TagAssetsByRecordsRequest) SetTags(v []*string) *TagAssetsByRecordsRequest {
	s.Tags = v
	return s
}

func (s *TagAssetsByRecordsRequest) SetRecordIds(v []*int) *TagAssetsByRecordsRequest {
	s.RecordIds = v
	return s
}

type TagAssetsByRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagAssetsByRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagAssetsByRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *TagAssetsByRecordsResponseBody) SetRequestId(v string) *TagAssetsByRecordsResponseBody {
	s.RequestId = &v
	return s
}

type TagAssetsByRecordsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagAssetsByRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagAssetsByRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s TagAssetsByRecordsResponse) GoString() string {
	return s.String()
}

func (s *TagAssetsByRecordsResponse) SetHeaders(v map[string]*string) *TagAssetsByRecordsResponse {
	s.Headers = v
	return s
}

func (s *TagAssetsByRecordsResponse) SetBody(v *TagAssetsByRecordsResponseBody) *TagAssetsByRecordsResponse {
	s.Body = v
	return s
}

type TagAssetsBySearchRequest struct {
	SourceIp  *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Source    *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Search    *string   `json:"Search,omitempty" xml:"Search,omitempty"`
	AssetType *string   `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagAssetsBySearchRequest) String() string {
	return tea.Prettify(s)
}

func (s TagAssetsBySearchRequest) GoString() string {
	return s.String()
}

func (s *TagAssetsBySearchRequest) SetSourceIp(v string) *TagAssetsBySearchRequest {
	s.SourceIp = &v
	return s
}

func (s *TagAssetsBySearchRequest) SetSource(v string) *TagAssetsBySearchRequest {
	s.Source = &v
	return s
}

func (s *TagAssetsBySearchRequest) SetSearch(v string) *TagAssetsBySearchRequest {
	s.Search = &v
	return s
}

func (s *TagAssetsBySearchRequest) SetAssetType(v string) *TagAssetsBySearchRequest {
	s.AssetType = &v
	return s
}

func (s *TagAssetsBySearchRequest) SetTags(v []*string) *TagAssetsBySearchRequest {
	s.Tags = v
	return s
}

type TagAssetsBySearchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagAssetsBySearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagAssetsBySearchResponseBody) GoString() string {
	return s.String()
}

func (s *TagAssetsBySearchResponseBody) SetRequestId(v string) *TagAssetsBySearchResponseBody {
	s.RequestId = &v
	return s
}

type TagAssetsBySearchResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagAssetsBySearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagAssetsBySearchResponse) String() string {
	return tea.Prettify(s)
}

func (s TagAssetsBySearchResponse) GoString() string {
	return s.String()
}

func (s *TagAssetsBySearchResponse) SetHeaders(v map[string]*string) *TagAssetsBySearchResponse {
	s.Headers = v
	return s
}

func (s *TagAssetsBySearchResponse) SetBody(v *TagAssetsBySearchResponseBody) *TagAssetsBySearchResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("avds"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddAssetsWithOptions(request *AddAssetsRequest, runtime *util.RuntimeOptions) (_result *AddAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAssetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAssets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAssets(request *AddAssetsRequest) (_result *AddAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAssetsResponse{}
	_body, _err := client.AddAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAssetTagsWithOptions(request *AddAssetTagsRequest, runtime *util.RuntimeOptions) (_result *AddAssetTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAssetTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAssetTags"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAssetTags(request *AddAssetTagsRequest) (_result *AddAssetTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAssetTagsResponse{}
	_body, _err := client.AddAssetTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOrgDomainsWithOptions(request *AddOrgDomainsRequest, runtime *util.RuntimeOptions) (_result *AddOrgDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddOrgDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddOrgDomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOrgDomains(request *AddOrgDomainsRequest) (_result *AddOrgDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddOrgDomainsResponse{}
	_body, _err := client.AddOrgDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOrgHostsWithOptions(request *AddOrgHostsRequest, runtime *util.RuntimeOptions) (_result *AddOrgHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddOrgHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddOrgHosts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOrgHosts(request *AddOrgHostsRequest) (_result *AddOrgHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddOrgHostsResponse{}
	_body, _err := client.AddOrgHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOrgSubdomainsWithOptions(request *AddOrgSubdomainsRequest, runtime *util.RuntimeOptions) (_result *AddOrgSubdomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddOrgSubdomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddOrgSubdomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOrgSubdomains(request *AddOrgSubdomainsRequest) (_result *AddOrgSubdomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddOrgSubdomainsResponse{}
	_body, _err := client.AddOrgSubdomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddOrgWebPathsWithOptions(request *AddOrgWebPathsRequest, runtime *util.RuntimeOptions) (_result *AddOrgWebPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddOrgWebPathsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddOrgWebPaths"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOrgWebPaths(request *AddOrgWebPathsRequest) (_result *AddOrgWebPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddOrgWebPathsResponse{}
	_body, _err := client.AddOrgWebPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAvdsBagOrderWithOptions(request *CreateAvdsBagOrderRequest, runtime *util.RuntimeOptions) (_result *CreateAvdsBagOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAvdsBagOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAvdsBagOrder"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAvdsBagOrder(request *CreateAvdsBagOrderRequest) (_result *CreateAvdsBagOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAvdsBagOrderResponse{}
	_body, _err := client.CreateAvdsBagOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAvdsOrderWithOptions(request *CreateAvdsOrderRequest, runtime *util.RuntimeOptions) (_result *CreateAvdsOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAvdsOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAvdsOrder"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAvdsOrder(request *CreateAvdsOrderRequest) (_result *CreateAvdsOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAvdsOrderResponse{}
	_body, _err := client.CreateAvdsOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAvdsPublicOrderWithOptions(request *CreateAvdsPublicOrderRequest, runtime *util.RuntimeOptions) (_result *CreateAvdsPublicOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAvdsPublicOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAvdsPublicOrder"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAvdsPublicOrder(request *CreateAvdsPublicOrderRequest) (_result *CreateAvdsPublicOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAvdsPublicOrderResponse{}
	_body, _err := client.CreateAvdsPublicOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrganizationWithOptions(request *CreateOrganizationRequest, runtime *util.RuntimeOptions) (_result *CreateOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOrganizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOrganization"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrganization(request *CreateOrganizationRequest) (_result *CreateOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrganizationResponse{}
	_body, _err := client.CreateOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, runtime *util.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSession"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSession(request *CreateSessionRequest) (_result *CreateSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSessionResponse{}
	_body, _err := client.CreateSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAssetsWithOptions(request *DeleteAssetsRequest, runtime *util.RuntimeOptions) (_result *DeleteAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAssetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAssets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAssets(request *DeleteAssetsRequest) (_result *DeleteAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAssetsResponse{}
	_body, _err := client.DeleteAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOrganizationsWithOptions(request *DeleteOrganizationsRequest, runtime *util.RuntimeOptions) (_result *DeleteOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteOrganizationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOrganizations"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOrganizations(request *DeleteOrganizationsRequest) (_result *DeleteOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOrganizationsResponse{}
	_body, _err := client.DeleteOrganizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOrgAttackSurfaceRecordsWithOptions(request *DeleteOrgAttackSurfaceRecordsRequest, runtime *util.RuntimeOptions) (_result *DeleteOrgAttackSurfaceRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteOrgAttackSurfaceRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOrgAttackSurfaceRecords"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOrgAttackSurfaceRecords(request *DeleteOrgAttackSurfaceRecordsRequest) (_result *DeleteOrgAttackSurfaceRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOrgAttackSurfaceRecordsResponse{}
	_body, _err := client.DeleteOrgAttackSurfaceRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSessionWithOptions(request *DeleteSessionRequest, runtime *util.RuntimeOptions) (_result *DeleteSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSessionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSession"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSession(request *DeleteSessionRequest) (_result *DeleteSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSessionResponse{}
	_body, _err := client.DeleteSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserAttackSurfaceRecordsWithOptions(request *DeleteUserAttackSurfaceRecordsRequest, runtime *util.RuntimeOptions) (_result *DeleteUserAttackSurfaceRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserAttackSurfaceRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserAttackSurfaceRecords"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserAttackSurfaceRecords(request *DeleteUserAttackSurfaceRecordsRequest) (_result *DeleteUserAttackSurfaceRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserAttackSurfaceRecordsResponse{}
	_body, _err := client.DeleteUserAttackSurfaceRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllVulnerabilitiesWithOptions(request *DescribeAllVulnerabilitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeAllVulnerabilitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllVulnerabilitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllVulnerabilities"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllVulnerabilities(request *DescribeAllVulnerabilitiesRequest) (_result *DescribeAllVulnerabilitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllVulnerabilitiesResponse{}
	_body, _err := client.DescribeAllVulnerabilitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssetsWithOptions(request *DescribeAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAssetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAssets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssets(request *DescribeAssetsRequest) (_result *DescribeAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetsResponse{}
	_body, _err := client.DescribeAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAttackSurfacesFacetsWithOptions(request *DescribeAttackSurfacesFacetsRequest, runtime *util.RuntimeOptions) (_result *DescribeAttackSurfacesFacetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAttackSurfacesFacetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAttackSurfacesFacets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAttackSurfacesFacets(request *DescribeAttackSurfacesFacetsRequest) (_result *DescribeAttackSurfacesFacetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAttackSurfacesFacetsResponse{}
	_body, _err := client.DescribeAttackSurfacesFacetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCrawlerRequestsWithOptions(request *DescribeCrawlerRequestsRequest, runtime *util.RuntimeOptions) (_result *DescribeCrawlerRequestsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCrawlerRequestsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCrawlerRequests"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCrawlerRequests(request *DescribeCrawlerRequestsRequest) (_result *DescribeCrawlerRequestsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCrawlerRequestsResponse{}
	_body, _err := client.DescribeCrawlerRequestsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDNSMapWithOptions(request *DescribeDNSMapRequest, runtime *util.RuntimeOptions) (_result *DescribeDNSMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDNSMapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDNSMap"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDNSMap(request *DescribeDNSMapRequest) (_result *DescribeDNSMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDNSMapResponse{}
	_body, _err := client.DescribeDNSMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAttackSurfacesFacetsWithOptions(request *DescribeDomainAttackSurfacesFacetsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAttackSurfacesFacetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainAttackSurfacesFacetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainAttackSurfacesFacets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAttackSurfacesFacets(request *DescribeDomainAttackSurfacesFacetsRequest) (_result *DescribeDomainAttackSurfacesFacetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAttackSurfacesFacetsResponse{}
	_body, _err := client.DescribeDomainAttackSurfacesFacetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHostAttackSurfacesFacetsWithOptions(request *DescribeHostAttackSurfacesFacetsRequest, runtime *util.RuntimeOptions) (_result *DescribeHostAttackSurfacesFacetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHostAttackSurfacesFacetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHostAttackSurfacesFacets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHostAttackSurfacesFacets(request *DescribeHostAttackSurfacesFacetsRequest) (_result *DescribeHostAttackSurfacesFacetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHostAttackSurfacesFacetsResponse{}
	_body, _err := client.DescribeHostAttackSurfacesFacetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHostsWithOptions(request *DescribeHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHosts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHosts(request *DescribeHostsRequest) (_result *DescribeHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHostsResponse{}
	_body, _err := client.DescribeHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeListSessionsWithOptions(request *DescribeListSessionsRequest, runtime *util.RuntimeOptions) (_result *DescribeListSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeListSessionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeListSessions"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeListSessions(request *DescribeListSessionsRequest) (_result *DescribeListSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeListSessionsResponse{}
	_body, _err := client.DescribeListSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOrgAttackSurfaceDetailsWithOptions(request *DescribeOrgAttackSurfaceDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeOrgAttackSurfaceDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOrgAttackSurfaceDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOrgAttackSurfaceDetails"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOrgAttackSurfaceDetails(request *DescribeOrgAttackSurfaceDetailsRequest) (_result *DescribeOrgAttackSurfaceDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrgAttackSurfaceDetailsResponse{}
	_body, _err := client.DescribeOrgAttackSurfaceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOrgInfoWithOptions(request *DescribeOrgInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeOrgInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOrgInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOrgInfo"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOrgInfo(request *DescribeOrgInfoRequest) (_result *DescribeOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrgInfoResponse{}
	_body, _err := client.DescribeOrgInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortsWithOptions(request *DescribePortsRequest, runtime *util.RuntimeOptions) (_result *DescribePortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePorts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePorts(request *DescribePortsRequest) (_result *DescribePortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortsResponse{}
	_body, _err := client.DescribePortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScanSessionsWithOptions(request *DescribeScanSessionsRequest, runtime *util.RuntimeOptions) (_result *DescribeScanSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScanSessionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScanSessions"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScanSessions(request *DescribeScanSessionsRequest) (_result *DescribeScanSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScanSessionsResponse{}
	_body, _err := client.DescribeScanSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSessionWithOptions(request *DescribeSessionRequest, runtime *util.RuntimeOptions) (_result *DescribeSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSessionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSession"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSession(request *DescribeSessionRequest) (_result *DescribeSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSessionResponse{}
	_body, _err := client.DescribeSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubdomainsWithOptions(request *DescribeSubdomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeSubdomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubdomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubdomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubdomains(request *DescribeSubdomainsRequest) (_result *DescribeSubdomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubdomainsResponse{}
	_body, _err := client.DescribeSubdomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskBriefInfoWithOptions(request *DescribeTaskBriefInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskBriefInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTaskBriefInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTaskBriefInfo"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTaskBriefInfo(request *DescribeTaskBriefInfoRequest) (_result *DescribeTaskBriefInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskBriefInfoResponse{}
	_body, _err := client.DescribeTaskBriefInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserTagsWithOptions(request *DescribeUserTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserTags"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserTags(request *DescribeUserTagsRequest) (_result *DescribeUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserTagsResponse{}
	_body, _err := client.DescribeUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulnerabilityWithOptions(request *DescribeVulnerabilityRequest, runtime *util.RuntimeOptions) (_result *DescribeVulnerabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulnerabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulnerability"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulnerability(request *DescribeVulnerabilityRequest) (_result *DescribeVulnerabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulnerabilityResponse{}
	_body, _err := client.DescribeVulnerabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebPathsWithOptions(request *DescribeWebPathsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebPathsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebPaths"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebPaths(request *DescribeWebPathsRequest) (_result *DescribeWebPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebPathsResponse{}
	_body, _err := client.DescribeWebPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebServersWithOptions(request *DescribeWebServersRequest, runtime *util.RuntimeOptions) (_result *DescribeWebServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebServersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebServers"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebServers(request *DescribeWebServersRequest) (_result *DescribeWebServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebServersResponse{}
	_body, _err := client.DescribeWebServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebTechsWithOptions(request *DescribeWebTechsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebTechsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebTechsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebTechs"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebTechs(request *DescribeWebTechsRequest) (_result *DescribeWebTechsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebTechsResponse{}
	_body, _err := client.DescribeWebTechsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditSessionWithOptions(request *EditSessionRequest, runtime *util.RuntimeOptions) (_result *EditSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditSessionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditSession"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditSession(request *EditSessionRequest) (_result *EditSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditSessionResponse{}
	_body, _err := client.EditSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateVulReportWithOptions(request *GenerateVulReportRequest, runtime *util.RuntimeOptions) (_result *GenerateVulReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateVulReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateVulReport"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVulReport(request *GenerateVulReportRequest) (_result *GenerateVulReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVulReportResponse{}
	_body, _err := client.GenerateVulReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgDNSMapWithOptions(request *ListOrgDNSMapRequest, runtime *util.RuntimeOptions) (_result *ListOrgDNSMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgDNSMapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgDNSMap"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgDNSMap(request *ListOrgDNSMapRequest) (_result *ListOrgDNSMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgDNSMapResponse{}
	_body, _err := client.ListOrgDNSMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgDomainsWithOptions(request *ListOrgDomainsRequest, runtime *util.RuntimeOptions) (_result *ListOrgDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgDomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgDomains(request *ListOrgDomainsRequest) (_result *ListOrgDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgDomainsResponse{}
	_body, _err := client.ListOrgDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgHostsWithOptions(request *ListOrgHostsRequest, runtime *util.RuntimeOptions) (_result *ListOrgHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgHosts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgHosts(request *ListOrgHostsRequest) (_result *ListOrgHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgHostsResponse{}
	_body, _err := client.ListOrgHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgPortsWithOptions(request *ListOrgPortsRequest, runtime *util.RuntimeOptions) (_result *ListOrgPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgPorts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgPorts(request *ListOrgPortsRequest) (_result *ListOrgPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgPortsResponse{}
	_body, _err := client.ListOrgPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgRiskyAssetsWithOptions(request *ListOrgRiskyAssetsRequest, runtime *util.RuntimeOptions) (_result *ListOrgRiskyAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgRiskyAssetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgRiskyAssets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgRiskyAssets(request *ListOrgRiskyAssetsRequest) (_result *ListOrgRiskyAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgRiskyAssetsResponse{}
	_body, _err := client.ListOrgRiskyAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgSubdomainsWithOptions(request *ListOrgSubdomainsRequest, runtime *util.RuntimeOptions) (_result *ListOrgSubdomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgSubdomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgSubdomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgSubdomains(request *ListOrgSubdomainsRequest) (_result *ListOrgSubdomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgSubdomainsResponse{}
	_body, _err := client.ListOrgSubdomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgVulFacetsWithOptions(request *ListOrgVulFacetsRequest, runtime *util.RuntimeOptions) (_result *ListOrgVulFacetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgVulFacetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgVulFacets"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgVulFacets(request *ListOrgVulFacetsRequest) (_result *ListOrgVulFacetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgVulFacetsResponse{}
	_body, _err := client.ListOrgVulFacetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgWebPathsWithOptions(request *ListOrgWebPathsRequest, runtime *util.RuntimeOptions) (_result *ListOrgWebPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgWebPathsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgWebPaths"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgWebPaths(request *ListOrgWebPathsRequest) (_result *ListOrgWebPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgWebPathsResponse{}
	_body, _err := client.ListOrgWebPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrgWebTechsWithOptions(request *ListOrgWebTechsRequest, runtime *util.RuntimeOptions) (_result *ListOrgWebTechsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrgWebTechsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrgWebTechs"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrgWebTechs(request *ListOrgWebTechsRequest) (_result *ListOrgWebTechsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrgWebTechsResponse{}
	_body, _err := client.ListOrgWebTechsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserDNSMapWithOptions(request *ListUserDNSMapRequest, runtime *util.RuntimeOptions) (_result *ListUserDNSMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserDNSMapResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserDNSMap"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserDNSMap(request *ListUserDNSMapRequest) (_result *ListUserDNSMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDNSMapResponse{}
	_body, _err := client.ListUserDNSMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserDNSMapHistoriesWithOptions(request *ListUserDNSMapHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListUserDNSMapHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserDNSMapHistoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserDNSMapHistories"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserDNSMapHistories(request *ListUserDNSMapHistoriesRequest) (_result *ListUserDNSMapHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDNSMapHistoriesResponse{}
	_body, _err := client.ListUserDNSMapHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserDomainsWithOptions(request *ListUserDomainsRequest, runtime *util.RuntimeOptions) (_result *ListUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserDomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserDomains(request *ListUserDomainsRequest) (_result *ListUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDomainsResponse{}
	_body, _err := client.ListUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserHostsWithOptions(request *ListUserHostsRequest, runtime *util.RuntimeOptions) (_result *ListUserHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserHosts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserHosts(request *ListUserHostsRequest) (_result *ListUserHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserHostsResponse{}
	_body, _err := client.ListUserHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserOrganizationsWithOptions(request *ListUserOrganizationsRequest, runtime *util.RuntimeOptions) (_result *ListUserOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserOrganizationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserOrganizations"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserOrganizations(request *ListUserOrganizationsRequest) (_result *ListUserOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserOrganizationsResponse{}
	_body, _err := client.ListUserOrganizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserPortsWithOptions(request *ListUserPortsRequest, runtime *util.RuntimeOptions) (_result *ListUserPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserPorts"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserPorts(request *ListUserPortsRequest) (_result *ListUserPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserPortsResponse{}
	_body, _err := client.ListUserPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserSubdomainsWithOptions(request *ListUserSubdomainsRequest, runtime *util.RuntimeOptions) (_result *ListUserSubdomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserSubdomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserSubdomains"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserSubdomains(request *ListUserSubdomainsRequest) (_result *ListUserSubdomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserSubdomainsResponse{}
	_body, _err := client.ListUserSubdomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserWebPathsWithOptions(request *ListUserWebPathsRequest, runtime *util.RuntimeOptions) (_result *ListUserWebPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserWebPathsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserWebPaths"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserWebPaths(request *ListUserWebPathsRequest) (_result *ListUserWebPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserWebPathsResponse{}
	_body, _err := client.ListUserWebPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserWebTechsWithOptions(request *ListUserWebTechsRequest, runtime *util.RuntimeOptions) (_result *ListUserWebTechsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserWebTechsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserWebTechs"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserWebTechs(request *ListUserWebTechsRequest) (_result *ListUserWebTechsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserWebTechsResponse{}
	_body, _err := client.ListUserWebTechsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOrganizationWithOptions(request *ModifyOrganizationRequest, runtime *util.RuntimeOptions) (_result *ModifyOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyOrganizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyOrganization"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOrganization(request *ModifyOrganizationRequest) (_result *ModifyOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOrganizationResponse{}
	_body, _err := client.ModifyOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagAssetsByRecordsWithOptions(request *TagAssetsByRecordsRequest, runtime *util.RuntimeOptions) (_result *TagAssetsByRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagAssetsByRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagAssetsByRecords"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagAssetsByRecords(request *TagAssetsByRecordsRequest) (_result *TagAssetsByRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagAssetsByRecordsResponse{}
	_body, _err := client.TagAssetsByRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagAssetsBySearchWithOptions(request *TagAssetsBySearchRequest, runtime *util.RuntimeOptions) (_result *TagAssetsBySearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagAssetsBySearchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagAssetsBySearch"), tea.String("2017-11-29"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagAssetsBySearch(request *TagAssetsBySearchRequest) (_result *TagAssetsBySearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagAssetsBySearchResponse{}
	_body, _err := client.TagAssetsBySearchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
