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

type AppOpenAckRequest struct {
	Ack             *string `json:"Ack,omitempty" xml:"Ack,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AsyncMethod     *string `json:"AsyncMethod,omitempty" xml:"AsyncMethod,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServiceId       *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s AppOpenAckRequest) String() string {
	return tea.Prettify(s)
}

func (s AppOpenAckRequest) GoString() string {
	return s.String()
}

func (s *AppOpenAckRequest) SetAck(v string) *AppOpenAckRequest {
	s.Ack = &v
	return s
}

func (s *AppOpenAckRequest) SetAppName(v string) *AppOpenAckRequest {
	s.AppName = &v
	return s
}

func (s *AppOpenAckRequest) SetAsyncMethod(v string) *AppOpenAckRequest {
	s.AsyncMethod = &v
	return s
}

func (s *AppOpenAckRequest) SetRegion(v string) *AppOpenAckRequest {
	s.Region = &v
	return s
}

func (s *AppOpenAckRequest) SetResourceOwnerId(v int64) *AppOpenAckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AppOpenAckRequest) SetServiceId(v string) *AppOpenAckRequest {
	s.ServiceId = &v
	return s
}

type AppOpenAckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppOpenAckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppOpenAckResponseBody) GoString() string {
	return s.String()
}

func (s *AppOpenAckResponseBody) SetRequestId(v string) *AppOpenAckResponseBody {
	s.RequestId = &v
	return s
}

type AppOpenAckResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AppOpenAckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppOpenAckResponse) String() string {
	return tea.Prettify(s)
}

func (s AppOpenAckResponse) GoString() string {
	return s.String()
}

func (s *AppOpenAckResponse) SetHeaders(v map[string]*string) *AppOpenAckResponse {
	s.Headers = v
	return s
}

func (s *AppOpenAckResponse) SetStatusCode(v int32) *AppOpenAckResponse {
	s.StatusCode = &v
	return s
}

func (s *AppOpenAckResponse) SetBody(v *AppOpenAckResponseBody) *AppOpenAckResponse {
	s.Body = v
	return s
}

type CreateDomainConfigRequest struct {
	Caller            *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Domain            *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HttpPort          *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpToUserIp      *int32  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsPort         *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	HttpsRedirect     *int32  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsAccessProduct   *int32  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	IsNonStandardPort *int32  `json:"IsNonStandardPort,omitempty" xml:"IsNonStandardPort,omitempty"`
	LoadBalancing     *int32  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	Protocols         *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RsType            *int32  `json:"RsType,omitempty" xml:"RsType,omitempty"`
	SourceIps         *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
}

func (s CreateDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainConfigRequest) SetCaller(v string) *CreateDomainConfigRequest {
	s.Caller = &v
	return s
}

func (s *CreateDomainConfigRequest) SetDomain(v string) *CreateDomainConfigRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpPort(v string) *CreateDomainConfigRequest {
	s.HttpPort = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpToUserIp(v int32) *CreateDomainConfigRequest {
	s.HttpToUserIp = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpsPort(v string) *CreateDomainConfigRequest {
	s.HttpsPort = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpsRedirect(v int32) *CreateDomainConfigRequest {
	s.HttpsRedirect = &v
	return s
}

func (s *CreateDomainConfigRequest) SetInstanceId(v string) *CreateDomainConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainConfigRequest) SetIsAccessProduct(v int32) *CreateDomainConfigRequest {
	s.IsAccessProduct = &v
	return s
}

func (s *CreateDomainConfigRequest) SetIsNonStandardPort(v int32) *CreateDomainConfigRequest {
	s.IsNonStandardPort = &v
	return s
}

func (s *CreateDomainConfigRequest) SetLoadBalancing(v int32) *CreateDomainConfigRequest {
	s.LoadBalancing = &v
	return s
}

func (s *CreateDomainConfigRequest) SetProtocols(v string) *CreateDomainConfigRequest {
	s.Protocols = &v
	return s
}

func (s *CreateDomainConfigRequest) SetRegion(v string) *CreateDomainConfigRequest {
	s.Region = &v
	return s
}

func (s *CreateDomainConfigRequest) SetResourceOwnerId(v int64) *CreateDomainConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDomainConfigRequest) SetRsType(v int32) *CreateDomainConfigRequest {
	s.RsType = &v
	return s
}

func (s *CreateDomainConfigRequest) SetSourceIps(v string) *CreateDomainConfigRequest {
	s.SourceIps = &v
	return s
}

type CreateDomainConfigResponseBody struct {
	Cname     *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainConfigResponseBody) SetCname(v string) *CreateDomainConfigResponseBody {
	s.Cname = &v
	return s
}

func (s *CreateDomainConfigResponseBody) SetRequestId(v string) *CreateDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainConfigResponse) SetHeaders(v map[string]*string) *CreateDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainConfigResponse) SetStatusCode(v int32) *CreateDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainConfigResponse) SetBody(v *CreateDomainConfigResponseBody) *CreateDomainConfigResponse {
	s.Body = v
	return s
}

type DeleteDomainConfigRequest struct {
	Caller          *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainConfigRequest) SetCaller(v string) *DeleteDomainConfigRequest {
	s.Caller = &v
	return s
}

func (s *DeleteDomainConfigRequest) SetDomain(v string) *DeleteDomainConfigRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainConfigRequest) SetInstanceId(v string) *DeleteDomainConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDomainConfigRequest) SetRegion(v string) *DeleteDomainConfigRequest {
	s.Region = &v
	return s
}

func (s *DeleteDomainConfigRequest) SetResourceOwnerId(v int64) *DeleteDomainConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteDomainConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainConfigResponseBody) SetRequestId(v string) *DeleteDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainConfigResponse) SetHeaders(v map[string]*string) *DeleteDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainConfigResponse) SetStatusCode(v int32) *DeleteDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainConfigResponse) SetBody(v *DeleteDomainConfigResponseBody) *DeleteDomainConfigResponse {
	s.Body = v
	return s
}

type DescribeDomainNamesRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDomainNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesRequest) SetInstanceId(v string) *DescribeDomainNamesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainNamesRequest) SetRegion(v string) *DescribeDomainNamesRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainNamesRequest) SetResourceOwnerId(v int64) *DescribeDomainNamesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDomainNamesResponseBody struct {
	DomainNames *DescribeDomainNamesResponseBodyDomainNames `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesResponseBody) SetDomainNames(v *DescribeDomainNamesResponseBodyDomainNames) *DescribeDomainNamesResponseBody {
	s.DomainNames = v
	return s
}

func (s *DescribeDomainNamesResponseBody) SetRequestId(v string) *DescribeDomainNamesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainNamesResponseBodyDomainNames struct {
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
}

func (s DescribeDomainNamesResponseBodyDomainNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesResponseBodyDomainNames) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesResponseBodyDomainNames) SetDomainName(v []*string) *DescribeDomainNamesResponseBodyDomainNames {
	s.DomainName = v
	return s
}

type DescribeDomainNamesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesResponse) SetHeaders(v map[string]*string) *DescribeDomainNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainNamesResponse) SetStatusCode(v int32) *DescribeDomainNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainNamesResponse) SetBody(v *DescribeDomainNamesResponseBody) *DescribeDomainNamesResponse {
	s.Body = v
	return s
}

type DescribeDomainTransferConfigRequest struct {
	Caller          *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDomainTransferConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTransferConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTransferConfigRequest) SetCaller(v string) *DescribeDomainTransferConfigRequest {
	s.Caller = &v
	return s
}

func (s *DescribeDomainTransferConfigRequest) SetDomain(v string) *DescribeDomainTransferConfigRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTransferConfigRequest) SetRegion(v string) *DescribeDomainTransferConfigRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainTransferConfigRequest) SetResourceOwnerId(v int64) *DescribeDomainTransferConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDomainTransferConfigResponseBody struct {
	Cname         *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	HttpPort      *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpsPort     *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	IsOwned       *int32  `json:"IsOwned,omitempty" xml:"IsOwned,omitempty"`
	IsWafActive   *int32  `json:"IsWafActive,omitempty" xml:"IsWafActive,omitempty"`
	ProtocolType  *int32  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Protocols     *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceIps     *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	WafAffectMode *int32  `json:"WafAffectMode,omitempty" xml:"WafAffectMode,omitempty"`
}

func (s DescribeDomainTransferConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTransferConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTransferConfigResponseBody) SetCname(v string) *DescribeDomainTransferConfigResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetHttpPort(v string) *DescribeDomainTransferConfigResponseBody {
	s.HttpPort = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetHttpsPort(v string) *DescribeDomainTransferConfigResponseBody {
	s.HttpsPort = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetIsOwned(v int32) *DescribeDomainTransferConfigResponseBody {
	s.IsOwned = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetIsWafActive(v int32) *DescribeDomainTransferConfigResponseBody {
	s.IsWafActive = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetProtocolType(v int32) *DescribeDomainTransferConfigResponseBody {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetProtocols(v string) *DescribeDomainTransferConfigResponseBody {
	s.Protocols = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetRequestId(v string) *DescribeDomainTransferConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetSourceIps(v string) *DescribeDomainTransferConfigResponseBody {
	s.SourceIps = &v
	return s
}

func (s *DescribeDomainTransferConfigResponseBody) SetWafAffectMode(v int32) *DescribeDomainTransferConfigResponseBody {
	s.WafAffectMode = &v
	return s
}

type DescribeDomainTransferConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainTransferConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainTransferConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTransferConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTransferConfigResponse) SetHeaders(v map[string]*string) *DescribeDomainTransferConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTransferConfigResponse) SetStatusCode(v int32) *DescribeDomainTransferConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainTransferConfigResponse) SetBody(v *DescribeDomainTransferConfigResponseBody) *DescribeDomainTransferConfigResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Page            *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceId(v string) *DescribeDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPage(v int32) *DescribeDomainsRequest {
	s.Page = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int32) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetRegion(v string) *DescribeDomainsRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceOwnerId(v int64) *DescribeDomainsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeDomainsResponseBody struct {
	Domains   *DescribeDomainsResponseBodyDomains  `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageInfo  *DescribeDomainsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetDomains(v *DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageInfo(v *DescribeDomainsResponseBodyPageInfo) *DescribeDomainsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainsResponseBodyDomains struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v []*string) *DescribeDomainsResponseBodyDomains {
	s.Domain = v
	return s
}

type DescribeDomainsResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDomainsResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeDomainsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDomainsResponseBodyPageInfo) SetPageSize(v int32) *DescribeDomainsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsResponseBodyPageInfo) SetTotal(v int32) *DescribeDomainsResponseBodyPageInfo {
	s.Total = &v
	return s
}

type DescribeDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDomainsResponse) SetStatusCode(v int32) *DescribeDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

type DescribeHttpsCertInUseRequest struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHttpsCertInUseRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpsCertInUseRequest) GoString() string {
	return s.String()
}

func (s *DescribeHttpsCertInUseRequest) SetDomain(v string) *DescribeHttpsCertInUseRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHttpsCertInUseRequest) SetRegion(v string) *DescribeHttpsCertInUseRequest {
	s.Region = &v
	return s
}

func (s *DescribeHttpsCertInUseRequest) SetResourceOwnerId(v int64) *DescribeHttpsCertInUseRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeHttpsCertInUseResponseBody struct {
	CertContent *string `json:"CertContent,omitempty" xml:"CertContent,omitempty"`
	CertId      *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertKey     *string `json:"CertKey,omitempty" xml:"CertKey,omitempty"`
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHttpsCertInUseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpsCertInUseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpsCertInUseResponseBody) SetCertContent(v string) *DescribeHttpsCertInUseResponseBody {
	s.CertContent = &v
	return s
}

func (s *DescribeHttpsCertInUseResponseBody) SetCertId(v string) *DescribeHttpsCertInUseResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeHttpsCertInUseResponseBody) SetCertKey(v string) *DescribeHttpsCertInUseResponseBody {
	s.CertKey = &v
	return s
}

func (s *DescribeHttpsCertInUseResponseBody) SetCertName(v string) *DescribeHttpsCertInUseResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeHttpsCertInUseResponseBody) SetRequestId(v string) *DescribeHttpsCertInUseResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHttpsCertInUseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHttpsCertInUseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHttpsCertInUseResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHttpsCertInUseResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpsCertInUseResponse) SetHeaders(v map[string]*string) *DescribeHttpsCertInUseResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpsCertInUseResponse) SetStatusCode(v int32) *DescribeHttpsCertInUseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpsCertInUseResponse) SetBody(v *DescribeHttpsCertInUseResponseBody) *DescribeHttpsCertInUseResponse {
	s.Body = v
	return s
}

type DescribeNeedUpgradeDomainLimitRequest struct {
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeNeedUpgradeDomainLimitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNeedUpgradeDomainLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeNeedUpgradeDomainLimitRequest) SetDomain(v string) *DescribeNeedUpgradeDomainLimitRequest {
	s.Domain = &v
	return s
}

func (s *DescribeNeedUpgradeDomainLimitRequest) SetInstanceId(v string) *DescribeNeedUpgradeDomainLimitRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNeedUpgradeDomainLimitRequest) SetRegion(v string) *DescribeNeedUpgradeDomainLimitRequest {
	s.Region = &v
	return s
}

func (s *DescribeNeedUpgradeDomainLimitRequest) SetResourceOwnerId(v int64) *DescribeNeedUpgradeDomainLimitRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeNeedUpgradeDomainLimitResponseBody struct {
	NeedUpgrade *bool   `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNeedUpgradeDomainLimitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNeedUpgradeDomainLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNeedUpgradeDomainLimitResponseBody) SetNeedUpgrade(v bool) *DescribeNeedUpgradeDomainLimitResponseBody {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeNeedUpgradeDomainLimitResponseBody) SetRequestId(v string) *DescribeNeedUpgradeDomainLimitResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNeedUpgradeDomainLimitResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNeedUpgradeDomainLimitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNeedUpgradeDomainLimitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNeedUpgradeDomainLimitResponse) GoString() string {
	return s.String()
}

func (s *DescribeNeedUpgradeDomainLimitResponse) SetHeaders(v map[string]*string) *DescribeNeedUpgradeDomainLimitResponse {
	s.Headers = v
	return s
}

func (s *DescribeNeedUpgradeDomainLimitResponse) SetStatusCode(v int32) *DescribeNeedUpgradeDomainLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNeedUpgradeDomainLimitResponse) SetBody(v *DescribeNeedUpgradeDomainLimitResponseBody) *DescribeNeedUpgradeDomainLimitResponse {
	s.Body = v
	return s
}

type DescribePackageRequest struct {
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageRequest) SetRegion(v string) *DescribePackageRequest {
	s.Region = &v
	return s
}

func (s *DescribePackageRequest) SetResourceOwnerId(v int64) *DescribePackageRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribePackageResponseBody struct {
	ExpireTime *int64                            `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId *string                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules      *DescribePackageResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	Version    *string                           `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageResponseBody) SetExpireTime(v int64) *DescribePackageResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribePackageResponseBody) SetInstanceId(v string) *DescribePackageResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribePackageResponseBody) SetRequestId(v string) *DescribePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageResponseBody) SetRules(v *DescribePackageResponseBodyRules) *DescribePackageResponseBody {
	s.Rules = v
	return s
}

func (s *DescribePackageResponseBody) SetVersion(v string) *DescribePackageResponseBody {
	s.Version = &v
	return s
}

type DescribePackageResponseBodyRules struct {
	Rule []*string `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribePackageResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribePackageResponseBodyRules) SetRule(v []*string) *DescribePackageResponseBodyRules {
	s.Rule = v
	return s
}

type DescribePackageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageResponse) SetHeaders(v map[string]*string) *DescribePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageResponse) SetStatusCode(v int32) *DescribePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackageResponse) SetBody(v *DescribePackageResponseBody) *DescribePackageResponse {
	s.Body = v
	return s
}

type DescribeQpsRequest struct {
	Domain           *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndMillisecond   *int64    `json:"EndMillisecond,omitempty" xml:"EndMillisecond,omitempty"`
	Field            []*string `json:"Field,omitempty" xml:"Field,omitempty" type:"Repeated"`
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval         *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Region           *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId  *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartMillisecond *int64    `json:"StartMillisecond,omitempty" xml:"StartMillisecond,omitempty"`
}

func (s DescribeQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeQpsRequest) SetDomain(v string) *DescribeQpsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeQpsRequest) SetEndMillisecond(v int64) *DescribeQpsRequest {
	s.EndMillisecond = &v
	return s
}

func (s *DescribeQpsRequest) SetField(v []*string) *DescribeQpsRequest {
	s.Field = v
	return s
}

func (s *DescribeQpsRequest) SetInstanceId(v string) *DescribeQpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeQpsRequest) SetInterval(v int32) *DescribeQpsRequest {
	s.Interval = &v
	return s
}

func (s *DescribeQpsRequest) SetRegion(v string) *DescribeQpsRequest {
	s.Region = &v
	return s
}

func (s *DescribeQpsRequest) SetResourceOwnerId(v int64) *DescribeQpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeQpsRequest) SetStartMillisecond(v int64) *DescribeQpsRequest {
	s.StartMillisecond = &v
	return s
}

type DescribeQpsResponseBody struct {
	Items     *DescribeQpsResponseBodyItems     `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeScope *DescribeQpsResponseBodyTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Struct"`
}

func (s DescribeQpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQpsResponseBody) SetItems(v *DescribeQpsResponseBodyItems) *DescribeQpsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeQpsResponseBody) SetRequestId(v string) *DescribeQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQpsResponseBody) SetTimeScope(v *DescribeQpsResponseBodyTimeScope) *DescribeQpsResponseBody {
	s.TimeScope = v
	return s
}

type DescribeQpsResponseBodyItems struct {
	Qps []*string `json:"Qps,omitempty" xml:"Qps,omitempty" type:"Repeated"`
}

func (s DescribeQpsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeQpsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeQpsResponseBodyItems) SetQps(v []*string) *DescribeQpsResponseBodyItems {
	s.Qps = v
	return s
}

type DescribeQpsResponseBodyTimeScope struct {
	End   *int64 `json:"End,omitempty" xml:"End,omitempty"`
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Step  *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeQpsResponseBodyTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeQpsResponseBodyTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeQpsResponseBodyTimeScope) SetEnd(v int64) *DescribeQpsResponseBodyTimeScope {
	s.End = &v
	return s
}

func (s *DescribeQpsResponseBodyTimeScope) SetStart(v int64) *DescribeQpsResponseBodyTimeScope {
	s.Start = &v
	return s
}

func (s *DescribeQpsResponseBodyTimeScope) SetStep(v int32) *DescribeQpsResponseBodyTimeScope {
	s.Step = &v
	return s
}

type DescribeQpsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeQpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeQpsResponse) SetHeaders(v map[string]*string) *DescribeQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeQpsResponse) SetStatusCode(v int32) *DescribeQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQpsResponse) SetBody(v *DescribeQpsResponseBody) *DescribeQpsResponse {
	s.Body = v
	return s
}

type DescribeRegionStatusRequest struct {
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRegionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionStatusRequest) SetInstanceSource(v string) *DescribeRegionStatusRequest {
	s.InstanceSource = &v
	return s
}

func (s *DescribeRegionStatusRequest) SetLang(v string) *DescribeRegionStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionStatusRequest) SetRegion(v string) *DescribeRegionStatusRequest {
	s.Region = &v
	return s
}

func (s *DescribeRegionStatusRequest) SetSourceIp(v string) *DescribeRegionStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeRegionStatusResponseBody struct {
	InDebt    *int32  `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	PayType   *int32  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionStatusResponseBody) SetInDebt(v int32) *DescribeRegionStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeRegionStatusResponseBody) SetPayType(v int32) *DescribeRegionStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeRegionStatusResponseBody) SetRequestId(v string) *DescribeRegionStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionStatusResponse) SetHeaders(v map[string]*string) *DescribeRegionStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionStatusResponse) SetStatusCode(v int32) *DescribeRegionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionStatusResponse) SetBody(v *DescribeRegionStatusResponseBody) *DescribeRegionStatusResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*string `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*string) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTransferConfigInWorkRequest struct {
	CheckRequestId  *string `json:"CheckRequestId,omitempty" xml:"CheckRequestId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeTransferConfigInWorkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferConfigInWorkRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransferConfigInWorkRequest) SetCheckRequestId(v string) *DescribeTransferConfigInWorkRequest {
	s.CheckRequestId = &v
	return s
}

func (s *DescribeTransferConfigInWorkRequest) SetDomain(v string) *DescribeTransferConfigInWorkRequest {
	s.Domain = &v
	return s
}

func (s *DescribeTransferConfigInWorkRequest) SetRegion(v string) *DescribeTransferConfigInWorkRequest {
	s.Region = &v
	return s
}

func (s *DescribeTransferConfigInWorkRequest) SetResourceOwnerId(v int64) *DescribeTransferConfigInWorkRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeTransferConfigInWorkResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WafRequestId *string `json:"WafRequestId,omitempty" xml:"WafRequestId,omitempty"`
}

func (s DescribeTransferConfigInWorkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferConfigInWorkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransferConfigInWorkResponseBody) SetRequestId(v string) *DescribeTransferConfigInWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransferConfigInWorkResponseBody) SetStatus(v string) *DescribeTransferConfigInWorkResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTransferConfigInWorkResponseBody) SetWafRequestId(v string) *DescribeTransferConfigInWorkResponseBody {
	s.WafRequestId = &v
	return s
}

type DescribeTransferConfigInWorkResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTransferConfigInWorkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTransferConfigInWorkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTransferConfigInWorkResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransferConfigInWorkResponse) SetHeaders(v map[string]*string) *DescribeTransferConfigInWorkResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransferConfigInWorkResponse) SetStatusCode(v int32) *DescribeTransferConfigInWorkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransferConfigInWorkResponse) SetBody(v *DescribeTransferConfigInWorkResponseBody) *DescribeTransferConfigInWorkResponse {
	s.Body = v
	return s
}

type GetQpsRequest struct {
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	E               *int64    `json:"e,omitempty" xml:"e,omitempty"`
	F               []*string `json:"f,omitempty" xml:"f,omitempty" type:"Repeated"`
	N               *string   `json:"n,omitempty" xml:"n,omitempty"`
	S               *int64    `json:"s,omitempty" xml:"s,omitempty"`
	X               *int32    `json:"x,omitempty" xml:"x,omitempty"`
}

func (s GetQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQpsRequest) GoString() string {
	return s.String()
}

func (s *GetQpsRequest) SetInstanceId(v string) *GetQpsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQpsRequest) SetRegion(v string) *GetQpsRequest {
	s.Region = &v
	return s
}

func (s *GetQpsRequest) SetResourceOwnerId(v int64) *GetQpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetQpsRequest) SetE(v int64) *GetQpsRequest {
	s.E = &v
	return s
}

func (s *GetQpsRequest) SetF(v []*string) *GetQpsRequest {
	s.F = v
	return s
}

func (s *GetQpsRequest) SetN(v string) *GetQpsRequest {
	s.N = &v
	return s
}

func (s *GetQpsRequest) SetS(v int64) *GetQpsRequest {
	s.S = &v
	return s
}

func (s *GetQpsRequest) SetX(v int32) *GetQpsRequest {
	s.X = &v
	return s
}

type GetQpsResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Items     *GetQpsResponseBodyItems     `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeScope *GetQpsResponseBodyTimeScope `json:"timeScope,omitempty" xml:"timeScope,omitempty" type:"Struct"`
}

func (s GetQpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQpsResponseBody) GoString() string {
	return s.String()
}

func (s *GetQpsResponseBody) SetCode(v string) *GetQpsResponseBody {
	s.Code = &v
	return s
}

func (s *GetQpsResponseBody) SetItems(v *GetQpsResponseBodyItems) *GetQpsResponseBody {
	s.Items = v
	return s
}

func (s *GetQpsResponseBody) SetMessage(v string) *GetQpsResponseBody {
	s.Message = &v
	return s
}

func (s *GetQpsResponseBody) SetRequestId(v string) *GetQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQpsResponseBody) SetSuccess(v bool) *GetQpsResponseBody {
	s.Success = &v
	return s
}

func (s *GetQpsResponseBody) SetTimeScope(v *GetQpsResponseBodyTimeScope) *GetQpsResponseBody {
	s.TimeScope = v
	return s
}

type GetQpsResponseBodyItems struct {
	Qps []*GetQpsResponseBodyItemsQps `json:"Qps,omitempty" xml:"Qps,omitempty" type:"Repeated"`
}

func (s GetQpsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s GetQpsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetQpsResponseBodyItems) SetQps(v []*GetQpsResponseBodyItemsQps) *GetQpsResponseBodyItems {
	s.Qps = v
	return s
}

type GetQpsResponseBodyItemsQps struct {
	Data *GetQpsResponseBodyItemsQpsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Name *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetQpsResponseBodyItemsQps) String() string {
	return tea.Prettify(s)
}

func (s GetQpsResponseBodyItemsQps) GoString() string {
	return s.String()
}

func (s *GetQpsResponseBodyItemsQps) SetData(v *GetQpsResponseBodyItemsQpsData) *GetQpsResponseBodyItemsQps {
	s.Data = v
	return s
}

func (s *GetQpsResponseBodyItemsQps) SetName(v string) *GetQpsResponseBodyItemsQps {
	s.Name = &v
	return s
}

type GetQpsResponseBodyItemsQpsData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetQpsResponseBodyItemsQpsData) String() string {
	return tea.Prettify(s)
}

func (s GetQpsResponseBodyItemsQpsData) GoString() string {
	return s.String()
}

func (s *GetQpsResponseBodyItemsQpsData) SetData(v []*string) *GetQpsResponseBodyItemsQpsData {
	s.Data = v
	return s
}

type GetQpsResponseBodyTimeScope struct {
	End   *int64 `json:"End,omitempty" xml:"End,omitempty"`
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Step  *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s GetQpsResponseBodyTimeScope) String() string {
	return tea.Prettify(s)
}

func (s GetQpsResponseBodyTimeScope) GoString() string {
	return s.String()
}

func (s *GetQpsResponseBodyTimeScope) SetEnd(v int64) *GetQpsResponseBodyTimeScope {
	s.End = &v
	return s
}

func (s *GetQpsResponseBodyTimeScope) SetStart(v int64) *GetQpsResponseBodyTimeScope {
	s.Start = &v
	return s
}

func (s *GetQpsResponseBodyTimeScope) SetStep(v int32) *GetQpsResponseBodyTimeScope {
	s.Step = &v
	return s
}

type GetQpsResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQpsResponse) GoString() string {
	return s.String()
}

func (s *GetQpsResponse) SetHeaders(v map[string]*string) *GetQpsResponse {
	s.Headers = v
	return s
}

func (s *GetQpsResponse) SetStatusCode(v int32) *GetQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQpsResponse) SetBody(v *GetQpsResponseBody) *GetQpsResponse {
	s.Body = v
	return s
}

type GetQpsTotalRequest struct {
	Region          *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	E               *int64    `json:"e,omitempty" xml:"e,omitempty"`
	F               []*string `json:"f,omitempty" xml:"f,omitempty" type:"Repeated"`
	InstanceId      *string   `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	N               *string   `json:"n,omitempty" xml:"n,omitempty"`
	S               *int64    `json:"s,omitempty" xml:"s,omitempty"`
	X               *int32    `json:"x,omitempty" xml:"x,omitempty"`
}

func (s GetQpsTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalRequest) GoString() string {
	return s.String()
}

func (s *GetQpsTotalRequest) SetRegion(v string) *GetQpsTotalRequest {
	s.Region = &v
	return s
}

func (s *GetQpsTotalRequest) SetResourceOwnerId(v int64) *GetQpsTotalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetQpsTotalRequest) SetE(v int64) *GetQpsTotalRequest {
	s.E = &v
	return s
}

func (s *GetQpsTotalRequest) SetF(v []*string) *GetQpsTotalRequest {
	s.F = v
	return s
}

func (s *GetQpsTotalRequest) SetInstanceId(v string) *GetQpsTotalRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQpsTotalRequest) SetN(v string) *GetQpsTotalRequest {
	s.N = &v
	return s
}

func (s *GetQpsTotalRequest) SetS(v int64) *GetQpsTotalRequest {
	s.S = &v
	return s
}

func (s *GetQpsTotalRequest) SetX(v int32) *GetQpsTotalRequest {
	s.X = &v
	return s
}

type GetQpsTotalResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Items     *GetQpsTotalResponseBodyItems     `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeScope *GetQpsTotalResponseBodyTimeScope `json:"timeScope,omitempty" xml:"timeScope,omitempty" type:"Struct"`
}

func (s GetQpsTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalResponseBody) GoString() string {
	return s.String()
}

func (s *GetQpsTotalResponseBody) SetCode(v string) *GetQpsTotalResponseBody {
	s.Code = &v
	return s
}

func (s *GetQpsTotalResponseBody) SetItems(v *GetQpsTotalResponseBodyItems) *GetQpsTotalResponseBody {
	s.Items = v
	return s
}

func (s *GetQpsTotalResponseBody) SetMessage(v string) *GetQpsTotalResponseBody {
	s.Message = &v
	return s
}

func (s *GetQpsTotalResponseBody) SetRequestId(v string) *GetQpsTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQpsTotalResponseBody) SetSuccess(v bool) *GetQpsTotalResponseBody {
	s.Success = &v
	return s
}

func (s *GetQpsTotalResponseBody) SetTimeScope(v *GetQpsTotalResponseBodyTimeScope) *GetQpsTotalResponseBody {
	s.TimeScope = v
	return s
}

type GetQpsTotalResponseBodyItems struct {
	Qps []*GetQpsTotalResponseBodyItemsQps `json:"Qps,omitempty" xml:"Qps,omitempty" type:"Repeated"`
}

func (s GetQpsTotalResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetQpsTotalResponseBodyItems) SetQps(v []*GetQpsTotalResponseBodyItemsQps) *GetQpsTotalResponseBodyItems {
	s.Qps = v
	return s
}

type GetQpsTotalResponseBodyItemsQps struct {
	Data *GetQpsTotalResponseBodyItemsQpsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Name *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetQpsTotalResponseBodyItemsQps) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalResponseBodyItemsQps) GoString() string {
	return s.String()
}

func (s *GetQpsTotalResponseBodyItemsQps) SetData(v *GetQpsTotalResponseBodyItemsQpsData) *GetQpsTotalResponseBodyItemsQps {
	s.Data = v
	return s
}

func (s *GetQpsTotalResponseBodyItemsQps) SetName(v string) *GetQpsTotalResponseBodyItemsQps {
	s.Name = &v
	return s
}

type GetQpsTotalResponseBodyItemsQpsData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetQpsTotalResponseBodyItemsQpsData) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalResponseBodyItemsQpsData) GoString() string {
	return s.String()
}

func (s *GetQpsTotalResponseBodyItemsQpsData) SetData(v []*string) *GetQpsTotalResponseBodyItemsQpsData {
	s.Data = v
	return s
}

type GetQpsTotalResponseBodyTimeScope struct {
	End   *int64 `json:"End,omitempty" xml:"End,omitempty"`
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Step  *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s GetQpsTotalResponseBodyTimeScope) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalResponseBodyTimeScope) GoString() string {
	return s.String()
}

func (s *GetQpsTotalResponseBodyTimeScope) SetEnd(v int64) *GetQpsTotalResponseBodyTimeScope {
	s.End = &v
	return s
}

func (s *GetQpsTotalResponseBodyTimeScope) SetStart(v int64) *GetQpsTotalResponseBodyTimeScope {
	s.Start = &v
	return s
}

func (s *GetQpsTotalResponseBodyTimeScope) SetStep(v int32) *GetQpsTotalResponseBodyTimeScope {
	s.Step = &v
	return s
}

type GetQpsTotalResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQpsTotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQpsTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQpsTotalResponse) GoString() string {
	return s.String()
}

func (s *GetQpsTotalResponse) SetHeaders(v map[string]*string) *GetQpsTotalResponse {
	s.Headers = v
	return s
}

func (s *GetQpsTotalResponse) SetStatusCode(v int32) *GetQpsTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQpsTotalResponse) SetBody(v *GetQpsTotalResponseBody) *GetQpsTotalResponse {
	s.Body = v
	return s
}

type GetRegionListRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetRegionListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListRequest) GoString() string {
	return s.String()
}

func (s *GetRegionListRequest) SetResourceOwnerId(v int64) *GetRegionListRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetRegionListResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetRegionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRegionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionListResponseBody) SetCode(v string) *GetRegionListResponseBody {
	s.Code = &v
	return s
}

func (s *GetRegionListResponseBody) SetData(v *GetRegionListResponseBodyData) *GetRegionListResponseBody {
	s.Data = v
	return s
}

func (s *GetRegionListResponseBody) SetMessage(v string) *GetRegionListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRegionListResponseBody) SetRequestId(v string) *GetRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionListResponseBody) SetSuccess(v bool) *GetRegionListResponseBody {
	s.Success = &v
	return s
}

type GetRegionListResponseBodyData struct {
	RegionInfo []*GetRegionListResponseBodyDataRegionInfo `json:"RegionInfo,omitempty" xml:"RegionInfo,omitempty" type:"Repeated"`
}

func (s GetRegionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRegionListResponseBodyData) SetRegionInfo(v []*GetRegionListResponseBodyDataRegionInfo) *GetRegionListResponseBodyData {
	s.RegionInfo = v
	return s
}

type GetRegionListResponseBodyDataRegionInfo struct {
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetRegionListResponseBodyDataRegionInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponseBodyDataRegionInfo) GoString() string {
	return s.String()
}

func (s *GetRegionListResponseBodyDataRegionInfo) SetDisplay(v string) *GetRegionListResponseBodyDataRegionInfo {
	s.Display = &v
	return s
}

func (s *GetRegionListResponseBodyDataRegionInfo) SetRegion(v string) *GetRegionListResponseBodyDataRegionInfo {
	s.Region = &v
	return s
}

type GetRegionListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegionListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionListResponse) GoString() string {
	return s.String()
}

func (s *GetRegionListResponse) SetHeaders(v map[string]*string) *GetRegionListResponse {
	s.Headers = v
	return s
}

func (s *GetRegionListResponse) SetStatusCode(v int32) *GetRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionListResponse) SetBody(v *GetRegionListResponseBody) *GetRegionListResponse {
	s.Body = v
	return s
}

type ModifyCertAndKeyRequest struct {
	Caller          *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Cert            *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HttpsCertId     *string `json:"HttpsCertId,omitempty" xml:"HttpsCertId,omitempty"`
	HttpsCertName   *string `json:"HttpsCertName,omitempty" xml:"HttpsCertName,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCertAndKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCertAndKeyRequest) GoString() string {
	return s.String()
}

func (s *ModifyCertAndKeyRequest) SetCaller(v string) *ModifyCertAndKeyRequest {
	s.Caller = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetCert(v string) *ModifyCertAndKeyRequest {
	s.Cert = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetDomain(v string) *ModifyCertAndKeyRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetHttpsCertId(v string) *ModifyCertAndKeyRequest {
	s.HttpsCertId = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetHttpsCertName(v string) *ModifyCertAndKeyRequest {
	s.HttpsCertName = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetKey(v string) *ModifyCertAndKeyRequest {
	s.Key = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetRegion(v string) *ModifyCertAndKeyRequest {
	s.Region = &v
	return s
}

func (s *ModifyCertAndKeyRequest) SetResourceOwnerId(v int64) *ModifyCertAndKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyCertAndKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCertAndKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCertAndKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCertAndKeyResponseBody) SetRequestId(v string) *ModifyCertAndKeyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCertAndKeyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyCertAndKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCertAndKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCertAndKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifyCertAndKeyResponse) SetHeaders(v map[string]*string) *ModifyCertAndKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifyCertAndKeyResponse) SetStatusCode(v int32) *ModifyCertAndKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCertAndKeyResponse) SetBody(v *ModifyCertAndKeyResponseBody) *ModifyCertAndKeyResponse {
	s.Body = v
	return s
}

type ModifyDomainConfigRequest struct {
	Caller            *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Domain            *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HttpPort          *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpToUserIp      *int32  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsPort         *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	HttpsRedirect     *int32  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsAccessProduct   *int32  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	IsNonStandardPort *int32  `json:"IsNonStandardPort,omitempty" xml:"IsNonStandardPort,omitempty"`
	LoadBalancing     *int32  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	Protocols         *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RsType            *int32  `json:"RsType,omitempty" xml:"RsType,omitempty"`
	SourceIps         *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
}

func (s ModifyDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainConfigRequest) SetCaller(v string) *ModifyDomainConfigRequest {
	s.Caller = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetDomain(v string) *ModifyDomainConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpPort(v string) *ModifyDomainConfigRequest {
	s.HttpPort = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpToUserIp(v int32) *ModifyDomainConfigRequest {
	s.HttpToUserIp = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpsPort(v string) *ModifyDomainConfigRequest {
	s.HttpsPort = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpsRedirect(v int32) *ModifyDomainConfigRequest {
	s.HttpsRedirect = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetInstanceId(v string) *ModifyDomainConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetIsAccessProduct(v int32) *ModifyDomainConfigRequest {
	s.IsAccessProduct = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetIsNonStandardPort(v int32) *ModifyDomainConfigRequest {
	s.IsNonStandardPort = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetLoadBalancing(v int32) *ModifyDomainConfigRequest {
	s.LoadBalancing = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetProtocols(v string) *ModifyDomainConfigRequest {
	s.Protocols = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetRegion(v string) *ModifyDomainConfigRequest {
	s.Region = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetResourceOwnerId(v int64) *ModifyDomainConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetRsType(v int32) *ModifyDomainConfigRequest {
	s.RsType = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetSourceIps(v string) *ModifyDomainConfigRequest {
	s.SourceIps = &v
	return s
}

type ModifyDomainConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainConfigResponseBody) SetRequestId(v string) *ModifyDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainConfigResponse) SetHeaders(v map[string]*string) *ModifyDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainConfigResponse) SetStatusCode(v int32) *ModifyDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainConfigResponse) SetBody(v *ModifyDomainConfigResponseBody) *ModifyDomainConfigResponse {
	s.Body = v
	return s
}

type ModifyDomainPackageCountRequest struct {
	DomainPackageCount *int32  `json:"DomainPackageCount,omitempty" xml:"DomainPackageCount,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId    *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDomainPackageCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainPackageCountRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainPackageCountRequest) SetDomainPackageCount(v int32) *ModifyDomainPackageCountRequest {
	s.DomainPackageCount = &v
	return s
}

func (s *ModifyDomainPackageCountRequest) SetInstanceId(v string) *ModifyDomainPackageCountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainPackageCountRequest) SetRegion(v string) *ModifyDomainPackageCountRequest {
	s.Region = &v
	return s
}

func (s *ModifyDomainPackageCountRequest) SetResourceOwnerId(v int64) *ModifyDomainPackageCountRequest {
	s.ResourceOwnerId = &v
	return s
}

type ModifyDomainPackageCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainPackageCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainPackageCountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainPackageCountResponseBody) SetRequestId(v string) *ModifyDomainPackageCountResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainPackageCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDomainPackageCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainPackageCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainPackageCountResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainPackageCountResponse) SetHeaders(v map[string]*string) *ModifyDomainPackageCountResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainPackageCountResponse) SetStatusCode(v int32) *ModifyDomainPackageCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainPackageCountResponse) SetBody(v *ModifyDomainPackageCountResponseBody) *ModifyDomainPackageCountResponse {
	s.Body = v
	return s
}

type ModifyWafSwitchRequest struct {
	Caller          *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServiceOn       *int32  `json:"ServiceOn,omitempty" xml:"ServiceOn,omitempty"`
}

func (s ModifyWafSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWafSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWafSwitchRequest) SetCaller(v string) *ModifyWafSwitchRequest {
	s.Caller = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetDomain(v string) *ModifyWafSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetInstanceId(v string) *ModifyWafSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetRegion(v string) *ModifyWafSwitchRequest {
	s.Region = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetResourceOwnerId(v int64) *ModifyWafSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetServiceOn(v int32) *ModifyWafSwitchRequest {
	s.ServiceOn = &v
	return s
}

type ModifyWafSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWafSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWafSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWafSwitchResponseBody) SetRequestId(v string) *ModifyWafSwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWafSwitchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyWafSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWafSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWafSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWafSwitchResponse) SetHeaders(v map[string]*string) *ModifyWafSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWafSwitchResponse) SetStatusCode(v int32) *ModifyWafSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWafSwitchResponse) SetBody(v *ModifyWafSwitchResponseBody) *ModifyWafSwitchResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-heyuan":             tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-wulanchabu":         tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":          tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("waf-openapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AppOpenAckWithOptions(request *AppOpenAckRequest, runtime *util.RuntimeOptions) (_result *AppOpenAckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ack)) {
		query["Ack"] = request.Ack
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AsyncMethod)) {
		query["AsyncMethod"] = request.AsyncMethod
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AppOpenAck"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppOpenAckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppOpenAck(request *AppOpenAckRequest) (_result *AppOpenAckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppOpenAckResponse{}
	_body, _err := client.AppOpenAckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainConfigWithOptions(request *CreateDomainConfigRequest, runtime *util.RuntimeOptions) (_result *CreateDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.HttpPort)) {
		query["HttpPort"] = request.HttpPort
	}

	if !tea.BoolValue(util.IsUnset(request.HttpToUserIp)) {
		query["HttpToUserIp"] = request.HttpToUserIp
	}

	if !tea.BoolValue(util.IsUnset(request.HttpsPort)) {
		query["HttpsPort"] = request.HttpsPort
	}

	if !tea.BoolValue(util.IsUnset(request.HttpsRedirect)) {
		query["HttpsRedirect"] = request.HttpsRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAccessProduct)) {
		query["IsAccessProduct"] = request.IsAccessProduct
	}

	if !tea.BoolValue(util.IsUnset(request.IsNonStandardPort)) {
		query["IsNonStandardPort"] = request.IsNonStandardPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancing)) {
		query["LoadBalancing"] = request.LoadBalancing
	}

	if !tea.BoolValue(util.IsUnset(request.Protocols)) {
		query["Protocols"] = request.Protocols
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RsType)) {
		query["RsType"] = request.RsType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIps)) {
		query["SourceIps"] = request.SourceIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomainConfig"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomainConfig(request *CreateDomainConfigRequest) (_result *CreateDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainConfigResponse{}
	_body, _err := client.CreateDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainConfigWithOptions(request *DeleteDomainConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainConfig"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainConfig(request *DeleteDomainConfigRequest) (_result *DeleteDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainConfigResponse{}
	_body, _err := client.DeleteDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainNamesWithOptions(request *DescribeDomainNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainNames"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainNames(request *DescribeDomainNamesRequest) (_result *DescribeDomainNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainNamesResponse{}
	_body, _err := client.DescribeDomainNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainTransferConfigWithOptions(request *DescribeDomainTransferConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainTransferConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainTransferConfig"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainTransferConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainTransferConfig(request *DescribeDomainTransferConfigRequest) (_result *DescribeDomainTransferConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainTransferConfigResponse{}
	_body, _err := client.DescribeDomainTransferConfigWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomains"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeHttpsCertInUseWithOptions(request *DescribeHttpsCertInUseRequest, runtime *util.RuntimeOptions) (_result *DescribeHttpsCertInUseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHttpsCertInUse"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHttpsCertInUseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHttpsCertInUse(request *DescribeHttpsCertInUseRequest) (_result *DescribeHttpsCertInUseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHttpsCertInUseResponse{}
	_body, _err := client.DescribeHttpsCertInUseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNeedUpgradeDomainLimitWithOptions(request *DescribeNeedUpgradeDomainLimitRequest, runtime *util.RuntimeOptions) (_result *DescribeNeedUpgradeDomainLimitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNeedUpgradeDomainLimit"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNeedUpgradeDomainLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNeedUpgradeDomainLimit(request *DescribeNeedUpgradeDomainLimitRequest) (_result *DescribeNeedUpgradeDomainLimitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNeedUpgradeDomainLimitResponse{}
	_body, _err := client.DescribeNeedUpgradeDomainLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackageWithOptions(request *DescribePackageRequest, runtime *util.RuntimeOptions) (_result *DescribePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackage"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackage(request *DescribePackageRequest) (_result *DescribePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackageResponse{}
	_body, _err := client.DescribePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQpsWithOptions(request *DescribeQpsRequest, runtime *util.RuntimeOptions) (_result *DescribeQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndMillisecond)) {
		query["EndMillisecond"] = request.EndMillisecond
	}

	if !tea.BoolValue(util.IsUnset(request.Field)) {
		query["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartMillisecond)) {
		query["StartMillisecond"] = request.StartMillisecond
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQps"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQps(request *DescribeQpsRequest) (_result *DescribeQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQpsResponse{}
	_body, _err := client.DescribeQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionStatusWithOptions(request *DescribeRegionStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceSource)) {
		query["InstanceSource"] = request.InstanceSource
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegionStatus"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegionStatus(request *DescribeRegionStatusRequest) (_result *DescribeRegionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionStatusResponse{}
	_body, _err := client.DescribeRegionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTransferConfigInWorkWithOptions(request *DescribeTransferConfigInWorkRequest, runtime *util.RuntimeOptions) (_result *DescribeTransferConfigInWorkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckRequestId)) {
		query["CheckRequestId"] = request.CheckRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTransferConfigInWork"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTransferConfigInWorkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTransferConfigInWork(request *DescribeTransferConfigInWorkRequest) (_result *DescribeTransferConfigInWorkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTransferConfigInWorkResponse{}
	_body, _err := client.DescribeTransferConfigInWorkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQpsWithOptions(request *GetQpsRequest, runtime *util.RuntimeOptions) (_result *GetQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.E)) {
		query["e"] = request.E
	}

	if !tea.BoolValue(util.IsUnset(request.F)) {
		query["f"] = request.F
	}

	if !tea.BoolValue(util.IsUnset(request.N)) {
		query["n"] = request.N
	}

	if !tea.BoolValue(util.IsUnset(request.S)) {
		query["s"] = request.S
	}

	if !tea.BoolValue(util.IsUnset(request.X)) {
		query["x"] = request.X
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQps"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQps(request *GetQpsRequest) (_result *GetQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQpsResponse{}
	_body, _err := client.GetQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQpsTotalWithOptions(request *GetQpsTotalRequest, runtime *util.RuntimeOptions) (_result *GetQpsTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.E)) {
		query["e"] = request.E
	}

	if !tea.BoolValue(util.IsUnset(request.F)) {
		query["f"] = request.F
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.N)) {
		query["n"] = request.N
	}

	if !tea.BoolValue(util.IsUnset(request.S)) {
		query["s"] = request.S
	}

	if !tea.BoolValue(util.IsUnset(request.X)) {
		query["x"] = request.X
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQpsTotal"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQpsTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQpsTotal(request *GetQpsTotalRequest) (_result *GetQpsTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQpsTotalResponse{}
	_body, _err := client.GetQpsTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegionListWithOptions(request *GetRegionListRequest, runtime *util.RuntimeOptions) (_result *GetRegionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegionList"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegionList(request *GetRegionListRequest) (_result *GetRegionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegionListResponse{}
	_body, _err := client.GetRegionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCertAndKeyWithOptions(request *ModifyCertAndKeyRequest, runtime *util.RuntimeOptions) (_result *ModifyCertAndKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Cert)) {
		query["Cert"] = request.Cert
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.HttpsCertId)) {
		query["HttpsCertId"] = request.HttpsCertId
	}

	if !tea.BoolValue(util.IsUnset(request.HttpsCertName)) {
		query["HttpsCertName"] = request.HttpsCertName
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCertAndKey"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCertAndKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCertAndKey(request *ModifyCertAndKeyRequest) (_result *ModifyCertAndKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCertAndKeyResponse{}
	_body, _err := client.ModifyCertAndKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainConfigWithOptions(request *ModifyDomainConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.HttpPort)) {
		query["HttpPort"] = request.HttpPort
	}

	if !tea.BoolValue(util.IsUnset(request.HttpToUserIp)) {
		query["HttpToUserIp"] = request.HttpToUserIp
	}

	if !tea.BoolValue(util.IsUnset(request.HttpsPort)) {
		query["HttpsPort"] = request.HttpsPort
	}

	if !tea.BoolValue(util.IsUnset(request.HttpsRedirect)) {
		query["HttpsRedirect"] = request.HttpsRedirect
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAccessProduct)) {
		query["IsAccessProduct"] = request.IsAccessProduct
	}

	if !tea.BoolValue(util.IsUnset(request.IsNonStandardPort)) {
		query["IsNonStandardPort"] = request.IsNonStandardPort
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancing)) {
		query["LoadBalancing"] = request.LoadBalancing
	}

	if !tea.BoolValue(util.IsUnset(request.Protocols)) {
		query["Protocols"] = request.Protocols
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RsType)) {
		query["RsType"] = request.RsType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIps)) {
		query["SourceIps"] = request.SourceIps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomainConfig"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainConfig(request *ModifyDomainConfigRequest) (_result *ModifyDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainConfigResponse{}
	_body, _err := client.ModifyDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainPackageCountWithOptions(request *ModifyDomainPackageCountRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainPackageCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainPackageCount)) {
		query["DomainPackageCount"] = request.DomainPackageCount
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomainPackageCount"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDomainPackageCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainPackageCount(request *ModifyDomainPackageCountRequest) (_result *ModifyDomainPackageCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainPackageCountResponse{}
	_body, _err := client.ModifyDomainPackageCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWafSwitchWithOptions(request *ModifyWafSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWafSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceOn)) {
		query["ServiceOn"] = request.ServiceOn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyWafSwitch"),
		Version:     tea.String("2017-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyWafSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWafSwitch(request *ModifyWafSwitchRequest) (_result *ModifyWafSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWafSwitchResponse{}
	_body, _err := client.ModifyWafSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
