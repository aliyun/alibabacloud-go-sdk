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

type CreateCertificateRequest struct {
	Certificate     *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s CreateCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequest) SetCertificate(v string) *CreateCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *CreateCertificateRequest) SetCertificateName(v string) *CreateCertificateRequest {
	s.CertificateName = &v
	return s
}

func (s *CreateCertificateRequest) SetDomain(v string) *CreateCertificateRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertificateRequest) SetInstanceId(v string) *CreateCertificateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCertificateRequest) SetPrivateKey(v string) *CreateCertificateRequest {
	s.PrivateKey = &v
	return s
}

type CreateCertificateResponseBody struct {
	CertificateId *int64  `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponseBody) SetCertificateId(v int64) *CreateCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CreateCertificateResponseBody) SetRequestId(v string) *CreateCertificateResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateResponse) SetHeaders(v map[string]*string) *CreateCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateResponse) SetBody(v *CreateCertificateResponseBody) *CreateCertificateResponse {
	s.Body = v
	return s
}

type CreateCertificateByCertificateIdRequest struct {
	CertificateId *int64  `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateCertificateByCertificateIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateByCertificateIdRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateByCertificateIdRequest) SetCertificateId(v int64) *CreateCertificateByCertificateIdRequest {
	s.CertificateId = &v
	return s
}

func (s *CreateCertificateByCertificateIdRequest) SetDomain(v string) *CreateCertificateByCertificateIdRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertificateByCertificateIdRequest) SetInstanceId(v string) *CreateCertificateByCertificateIdRequest {
	s.InstanceId = &v
	return s
}

type CreateCertificateByCertificateIdResponseBody struct {
	CertificateId *int64  `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCertificateByCertificateIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateByCertificateIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateByCertificateIdResponseBody) SetCertificateId(v int64) *CreateCertificateByCertificateIdResponseBody {
	s.CertificateId = &v
	return s
}

func (s *CreateCertificateByCertificateIdResponseBody) SetRequestId(v string) *CreateCertificateByCertificateIdResponseBody {
	s.RequestId = &v
	return s
}

type CreateCertificateByCertificateIdResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertificateByCertificateIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertificateByCertificateIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateByCertificateIdResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateByCertificateIdResponse) SetHeaders(v map[string]*string) *CreateCertificateByCertificateIdResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateByCertificateIdResponse) SetBody(v *CreateCertificateByCertificateIdResponseBody) *CreateCertificateByCertificateIdResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	AccessHeaderMode     *int32  `json:"AccessHeaderMode,omitempty" xml:"AccessHeaderMode,omitempty"`
	AccessHeaders        *string `json:"AccessHeaders,omitempty" xml:"AccessHeaders,omitempty"`
	AccessType           *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CloudNativeInstances *string `json:"CloudNativeInstances,omitempty" xml:"CloudNativeInstances,omitempty"`
	ClusterType          *int32  `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ConnectionTime       *int32  `json:"ConnectionTime,omitempty" xml:"ConnectionTime,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Http2Port            *string `json:"Http2Port,omitempty" xml:"Http2Port,omitempty"`
	HttpPort             *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpToUserIp         *int32  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsPort            *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	HttpsRedirect        *int32  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpFollowStatus       *int32  `json:"IpFollowStatus,omitempty" xml:"IpFollowStatus,omitempty"`
	IsAccessProduct      *int32  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	LoadBalancing        *int32  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	LogHeaders           *string `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty"`
	ReadTime             *int32  `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SniHost              *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	SniStatus            *int32  `json:"SniStatus,omitempty" xml:"SniStatus,omitempty"`
	SourceIps            *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	WriteTime            *int32  `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetAccessHeaderMode(v int32) *CreateDomainRequest {
	s.AccessHeaderMode = &v
	return s
}

func (s *CreateDomainRequest) SetAccessHeaders(v string) *CreateDomainRequest {
	s.AccessHeaders = &v
	return s
}

func (s *CreateDomainRequest) SetAccessType(v string) *CreateDomainRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainRequest) SetCloudNativeInstances(v string) *CreateDomainRequest {
	s.CloudNativeInstances = &v
	return s
}

func (s *CreateDomainRequest) SetClusterType(v int32) *CreateDomainRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateDomainRequest) SetConnectionTime(v int32) *CreateDomainRequest {
	s.ConnectionTime = &v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetHttp2Port(v string) *CreateDomainRequest {
	s.Http2Port = &v
	return s
}

func (s *CreateDomainRequest) SetHttpPort(v string) *CreateDomainRequest {
	s.HttpPort = &v
	return s
}

func (s *CreateDomainRequest) SetHttpToUserIp(v int32) *CreateDomainRequest {
	s.HttpToUserIp = &v
	return s
}

func (s *CreateDomainRequest) SetHttpsPort(v string) *CreateDomainRequest {
	s.HttpsPort = &v
	return s
}

func (s *CreateDomainRequest) SetHttpsRedirect(v int32) *CreateDomainRequest {
	s.HttpsRedirect = &v
	return s
}

func (s *CreateDomainRequest) SetInstanceId(v string) *CreateDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainRequest) SetIpFollowStatus(v int32) *CreateDomainRequest {
	s.IpFollowStatus = &v
	return s
}

func (s *CreateDomainRequest) SetIsAccessProduct(v int32) *CreateDomainRequest {
	s.IsAccessProduct = &v
	return s
}

func (s *CreateDomainRequest) SetLoadBalancing(v int32) *CreateDomainRequest {
	s.LoadBalancing = &v
	return s
}

func (s *CreateDomainRequest) SetLogHeaders(v string) *CreateDomainRequest {
	s.LogHeaders = &v
	return s
}

func (s *CreateDomainRequest) SetReadTime(v int32) *CreateDomainRequest {
	s.ReadTime = &v
	return s
}

func (s *CreateDomainRequest) SetResourceGroupId(v string) *CreateDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDomainRequest) SetSniHost(v string) *CreateDomainRequest {
	s.SniHost = &v
	return s
}

func (s *CreateDomainRequest) SetSniStatus(v int32) *CreateDomainRequest {
	s.SniStatus = &v
	return s
}

func (s *CreateDomainRequest) SetSourceIps(v string) *CreateDomainRequest {
	s.SourceIps = &v
	return s
}

func (s *CreateDomainRequest) SetWriteTime(v int32) *CreateDomainRequest {
	s.WriteTime = &v
	return s
}

type CreateDomainResponseBody struct {
	Cname     *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetCname(v string) *CreateDomainResponseBody {
	s.Cname = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateProtectionModuleRuleRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s CreateProtectionModuleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProtectionModuleRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateProtectionModuleRuleRequest) SetDefenseType(v string) *CreateProtectionModuleRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetDomain(v string) *CreateProtectionModuleRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetInstanceId(v string) *CreateProtectionModuleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetRule(v string) *CreateProtectionModuleRuleRequest {
	s.Rule = &v
	return s
}

type CreateProtectionModuleRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProtectionModuleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProtectionModuleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProtectionModuleRuleResponseBody) SetRequestId(v string) *CreateProtectionModuleRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateProtectionModuleRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProtectionModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProtectionModuleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProtectionModuleRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateProtectionModuleRuleResponse) SetHeaders(v map[string]*string) *CreateProtectionModuleRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateProtectionModuleRuleResponse) SetBody(v *CreateProtectionModuleRuleResponseBody) *CreateProtectionModuleRuleResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainRequest) SetInstanceId(v string) *DeleteDomainRequest {
	s.InstanceId = &v
	return s
}

type DeleteDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceGroupId(v string) *DeleteInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteProtectionModuleRuleRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteProtectionModuleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectionModuleRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtectionModuleRuleRequest) SetDefenseType(v string) *DeleteProtectionModuleRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *DeleteProtectionModuleRuleRequest) SetDomain(v string) *DeleteProtectionModuleRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteProtectionModuleRuleRequest) SetInstanceId(v string) *DeleteProtectionModuleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteProtectionModuleRuleRequest) SetRuleId(v int64) *DeleteProtectionModuleRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteProtectionModuleRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtectionModuleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectionModuleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtectionModuleRuleResponseBody) SetRequestId(v string) *DeleteProtectionModuleRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProtectionModuleRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProtectionModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProtectionModuleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProtectionModuleRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteProtectionModuleRuleResponse) SetHeaders(v map[string]*string) *DeleteProtectionModuleRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteProtectionModuleRuleResponse) SetBody(v *DeleteProtectionModuleRuleResponseBody) *DeleteProtectionModuleRuleResponse {
	s.Body = v
	return s
}

type DescribeCertMatchStatusRequest struct {
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s DescribeCertMatchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertMatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertMatchStatusRequest) SetCertificate(v string) *DescribeCertMatchStatusRequest {
	s.Certificate = &v
	return s
}

func (s *DescribeCertMatchStatusRequest) SetDomain(v string) *DescribeCertMatchStatusRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCertMatchStatusRequest) SetInstanceId(v string) *DescribeCertMatchStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCertMatchStatusRequest) SetPrivateKey(v string) *DescribeCertMatchStatusRequest {
	s.PrivateKey = &v
	return s
}

type DescribeCertMatchStatusResponseBody struct {
	MatchStatus *bool   `json:"MatchStatus,omitempty" xml:"MatchStatus,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertMatchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertMatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertMatchStatusResponseBody) SetMatchStatus(v bool) *DescribeCertMatchStatusResponseBody {
	s.MatchStatus = &v
	return s
}

func (s *DescribeCertMatchStatusResponseBody) SetRequestId(v string) *DescribeCertMatchStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCertMatchStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertMatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCertMatchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertMatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertMatchStatusResponse) SetHeaders(v map[string]*string) *DescribeCertMatchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertMatchStatusResponse) SetBody(v *DescribeCertMatchStatusResponseBody) *DescribeCertMatchStatusResponse {
	s.Body = v
	return s
}

type DescribeCertificatesRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificatesRequest) SetDomain(v string) *DescribeCertificatesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCertificatesRequest) SetInstanceId(v string) *DescribeCertificatesRequest {
	s.InstanceId = &v
	return s
}

type DescribeCertificatesResponseBody struct {
	Certificates []*DescribeCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificatesResponseBody) SetCertificates(v []*DescribeCertificatesResponseBodyCertificates) *DescribeCertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *DescribeCertificatesResponseBody) SetRequestId(v string) *DescribeCertificatesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCertificatesResponseBodyCertificates struct {
	CertificateId   *int64    `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName *string   `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CommonName      *string   `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	IsUsing         *bool     `json:"IsUsing,omitempty" xml:"IsUsing,omitempty"`
	Sans            []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
}

func (s DescribeCertificatesResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *DescribeCertificatesResponseBodyCertificates) SetCertificateId(v int64) *DescribeCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *DescribeCertificatesResponseBodyCertificates) SetCertificateName(v string) *DescribeCertificatesResponseBodyCertificates {
	s.CertificateName = &v
	return s
}

func (s *DescribeCertificatesResponseBodyCertificates) SetCommonName(v string) *DescribeCertificatesResponseBodyCertificates {
	s.CommonName = &v
	return s
}

func (s *DescribeCertificatesResponseBodyCertificates) SetIsUsing(v bool) *DescribeCertificatesResponseBodyCertificates {
	s.IsUsing = &v
	return s
}

func (s *DescribeCertificatesResponseBodyCertificates) SetSans(v []*string) *DescribeCertificatesResponseBodyCertificates {
	s.Sans = v
	return s
}

type DescribeCertificatesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificatesResponse) SetHeaders(v map[string]*string) *DescribeCertificatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificatesResponse) SetBody(v *DescribeCertificatesResponseBody) *DescribeCertificatesResponse {
	s.Body = v
	return s
}

type DescribeDomainRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRequest) SetDomain(v string) *DescribeDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainRequest) SetInstanceId(v string) *DescribeDomainRequest {
	s.InstanceId = &v
	return s
}

type DescribeDomainResponseBody struct {
	Domain    *DescribeDomainResponseBodyDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBody) SetDomain(v *DescribeDomainResponseBodyDomain) *DescribeDomainResponseBody {
	s.Domain = v
	return s
}

func (s *DescribeDomainResponseBody) SetRequestId(v string) *DescribeDomainResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainResponseBodyDomain struct {
	AccessHeaderMode     *int32                                                  `json:"AccessHeaderMode,omitempty" xml:"AccessHeaderMode,omitempty"`
	AccessHeaders        []*string                                               `json:"AccessHeaders,omitempty" xml:"AccessHeaders,omitempty" type:"Repeated"`
	AccessType           *string                                                 `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CloudNativeInstances []*DescribeDomainResponseBodyDomainCloudNativeInstances `json:"CloudNativeInstances,omitempty" xml:"CloudNativeInstances,omitempty" type:"Repeated"`
	ClusterType          *int32                                                  `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Cname                *string                                                 `json:"Cname,omitempty" xml:"Cname,omitempty"`
	ConnectionTime       *int32                                                  `json:"ConnectionTime,omitempty" xml:"ConnectionTime,omitempty"`
	Http2Port            []*string                                               `json:"Http2Port,omitempty" xml:"Http2Port,omitempty" type:"Repeated"`
	HttpPort             []*string                                               `json:"HttpPort,omitempty" xml:"HttpPort,omitempty" type:"Repeated"`
	HttpToUserIp         *int32                                                  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsPort            []*string                                               `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty" type:"Repeated"`
	HttpsRedirect        *int32                                                  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	IpFollowStatus       *int32                                                  `json:"IpFollowStatus,omitempty" xml:"IpFollowStatus,omitempty"`
	IsAccessProduct      *int32                                                  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	LoadBalancing        *int32                                                  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	LogHeaders           []*DescribeDomainResponseBodyDomainLogHeaders           `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty" type:"Repeated"`
	ReadTime             *int32                                                  `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	ResourceGroupId      *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SniHost              *string                                                 `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	SniStatus            *int32                                                  `json:"SniStatus,omitempty" xml:"SniStatus,omitempty"`
	SourceIps            []*string                                               `json:"SourceIps,omitempty" xml:"SourceIps,omitempty" type:"Repeated"`
	Version              *int64                                                  `json:"Version,omitempty" xml:"Version,omitempty"`
	WriteTime            *int32                                                  `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
}

func (s DescribeDomainResponseBodyDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBodyDomain) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBodyDomain) SetAccessHeaderMode(v int32) *DescribeDomainResponseBodyDomain {
	s.AccessHeaderMode = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetAccessHeaders(v []*string) *DescribeDomainResponseBodyDomain {
	s.AccessHeaders = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetAccessType(v string) *DescribeDomainResponseBodyDomain {
	s.AccessType = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetCloudNativeInstances(v []*DescribeDomainResponseBodyDomainCloudNativeInstances) *DescribeDomainResponseBodyDomain {
	s.CloudNativeInstances = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetClusterType(v int32) *DescribeDomainResponseBodyDomain {
	s.ClusterType = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetCname(v string) *DescribeDomainResponseBodyDomain {
	s.Cname = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetConnectionTime(v int32) *DescribeDomainResponseBodyDomain {
	s.ConnectionTime = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttp2Port(v []*string) *DescribeDomainResponseBodyDomain {
	s.Http2Port = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpPort(v []*string) *DescribeDomainResponseBodyDomain {
	s.HttpPort = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpToUserIp(v int32) *DescribeDomainResponseBodyDomain {
	s.HttpToUserIp = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpsPort(v []*string) *DescribeDomainResponseBodyDomain {
	s.HttpsPort = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpsRedirect(v int32) *DescribeDomainResponseBodyDomain {
	s.HttpsRedirect = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetIpFollowStatus(v int32) *DescribeDomainResponseBodyDomain {
	s.IpFollowStatus = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetIsAccessProduct(v int32) *DescribeDomainResponseBodyDomain {
	s.IsAccessProduct = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetLoadBalancing(v int32) *DescribeDomainResponseBodyDomain {
	s.LoadBalancing = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetLogHeaders(v []*DescribeDomainResponseBodyDomainLogHeaders) *DescribeDomainResponseBodyDomain {
	s.LogHeaders = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetReadTime(v int32) *DescribeDomainResponseBodyDomain {
	s.ReadTime = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetResourceGroupId(v string) *DescribeDomainResponseBodyDomain {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetSniHost(v string) *DescribeDomainResponseBodyDomain {
	s.SniHost = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetSniStatus(v int32) *DescribeDomainResponseBodyDomain {
	s.SniStatus = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetSourceIps(v []*string) *DescribeDomainResponseBodyDomain {
	s.SourceIps = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetVersion(v int64) *DescribeDomainResponseBodyDomain {
	s.Version = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetWriteTime(v int32) *DescribeDomainResponseBodyDomain {
	s.WriteTime = &v
	return s
}

type DescribeDomainResponseBodyDomainCloudNativeInstances struct {
	CloudNativeProductName *string                                                                    `json:"CloudNativeProductName,omitempty" xml:"CloudNativeProductName,omitempty"`
	IPAddressList          *string                                                                    `json:"IPAddressList,omitempty" xml:"IPAddressList,omitempty"`
	InstanceId             *string                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProtocolPortConfigs    []*DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs `json:"ProtocolPortConfigs,omitempty" xml:"ProtocolPortConfigs,omitempty" type:"Repeated"`
	RedirectionTypeName    *string                                                                    `json:"RedirectionTypeName,omitempty" xml:"RedirectionTypeName,omitempty"`
}

func (s DescribeDomainResponseBodyDomainCloudNativeInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBodyDomainCloudNativeInstances) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstances) SetCloudNativeProductName(v string) *DescribeDomainResponseBodyDomainCloudNativeInstances {
	s.CloudNativeProductName = &v
	return s
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstances) SetIPAddressList(v string) *DescribeDomainResponseBodyDomainCloudNativeInstances {
	s.IPAddressList = &v
	return s
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstances) SetInstanceId(v string) *DescribeDomainResponseBodyDomainCloudNativeInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstances) SetProtocolPortConfigs(v []*DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) *DescribeDomainResponseBodyDomainCloudNativeInstances {
	s.ProtocolPortConfigs = v
	return s
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstances) SetRedirectionTypeName(v string) *DescribeDomainResponseBodyDomainCloudNativeInstances {
	s.RedirectionTypeName = &v
	return s
}

type DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs struct {
	Ports    *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) SetPorts(v string) *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs {
	s.Ports = &v
	return s
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) SetProtocol(v string) *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs {
	s.Protocol = &v
	return s
}

type DescribeDomainResponseBodyDomainLogHeaders struct {
	K *string `json:"k,omitempty" xml:"k,omitempty"`
	V *string `json:"v,omitempty" xml:"v,omitempty"`
}

func (s DescribeDomainResponseBodyDomainLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBodyDomainLogHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBodyDomainLogHeaders) SetK(v string) *DescribeDomainResponseBodyDomainLogHeaders {
	s.K = &v
	return s
}

func (s *DescribeDomainResponseBodyDomainLogHeaders) SetV(v string) *DescribeDomainResponseBodyDomainLogHeaders {
	s.V = &v
	return s
}

type DescribeDomainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponse) SetHeaders(v map[string]*string) *DescribeDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResponse) SetBody(v *DescribeDomainResponseBody) *DescribeDomainResponse {
	s.Body = v
	return s
}

type DescribeDomainAdvanceConfigsRequest struct {
	DomainList      *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDomainAdvanceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAdvanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAdvanceConfigsRequest) SetDomainList(v string) *DescribeDomainAdvanceConfigsRequest {
	s.DomainList = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsRequest) SetInstanceId(v string) *DescribeDomainAdvanceConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsRequest) SetResourceGroupId(v string) *DescribeDomainAdvanceConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainAdvanceConfigsResponseBody struct {
	DomainConfigs []*DescribeDomainAdvanceConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainAdvanceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAdvanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAdvanceConfigsResponseBody) SetDomainConfigs(v []*DescribeDomainAdvanceConfigsResponseBodyDomainConfigs) *DescribeDomainAdvanceConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBody) SetRequestId(v string) *DescribeDomainAdvanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainAdvanceConfigsResponseBodyDomainConfigs struct {
	Domain  *string                                                       `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Profile *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile `json:"Profile,omitempty" xml:"Profile,omitempty" type:"Struct"`
}

func (s DescribeDomainAdvanceConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAdvanceConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigs) SetDomain(v string) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigs {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigs) SetProfile(v *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigs {
	s.Profile = v
	return s
}

type DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile struct {
	CertStatus         *int32    `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	ClusterType        *int32    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Cname              *string   `json:"Cname,omitempty" xml:"Cname,omitempty"`
	ExclusiveVipStatus *int32    `json:"ExclusiveVipStatus,omitempty" xml:"ExclusiveVipStatus,omitempty"`
	GSLBStatus         *string   `json:"GSLBStatus,omitempty" xml:"GSLBStatus,omitempty"`
	Http2Port          []*int32  `json:"Http2Port,omitempty" xml:"Http2Port,omitempty" type:"Repeated"`
	HttpPort           []*int32  `json:"HttpPort,omitempty" xml:"HttpPort,omitempty" type:"Repeated"`
	HttpsPort          []*int32  `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty" type:"Repeated"`
	Ipv6Status         *int32    `json:"Ipv6Status,omitempty" xml:"Ipv6Status,omitempty"`
	ResolvedType       *int32    `json:"ResolvedType,omitempty" xml:"ResolvedType,omitempty"`
	Rs                 []*string `json:"Rs,omitempty" xml:"Rs,omitempty" type:"Repeated"`
	VipServiceStatus   *int32    `json:"VipServiceStatus,omitempty" xml:"VipServiceStatus,omitempty"`
}

func (s DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) GoString() string {
	return s.String()
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetCertStatus(v int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.CertStatus = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetClusterType(v int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.ClusterType = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetCname(v string) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.Cname = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetExclusiveVipStatus(v int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.ExclusiveVipStatus = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetGSLBStatus(v string) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.GSLBStatus = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetHttp2Port(v []*int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.Http2Port = v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetHttpPort(v []*int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.HttpPort = v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetHttpsPort(v []*int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.HttpsPort = v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetIpv6Status(v int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.Ipv6Status = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetResolvedType(v int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.ResolvedType = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetRs(v []*string) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.Rs = v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile) SetVipServiceStatus(v int32) *DescribeDomainAdvanceConfigsResponseBodyDomainConfigsProfile {
	s.VipServiceStatus = &v
	return s
}

type DescribeDomainAdvanceConfigsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainAdvanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainAdvanceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAdvanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAdvanceConfigsResponse) SetHeaders(v map[string]*string) *DescribeDomainAdvanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponse) SetBody(v *DescribeDomainAdvanceConfigsResponseBody) *DescribeDomainAdvanceConfigsResponse {
	s.Body = v
	return s
}

type DescribeDomainBasicConfigsRequest struct {
	AccessType           *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CloudNativeProductId *int32  `json:"CloudNativeProductId,omitempty" xml:"CloudNativeProductId,omitempty"`
	DomainKey            *string `json:"DomainKey,omitempty" xml:"DomainKey,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDomainBasicConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBasicConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBasicConfigsRequest) SetAccessType(v string) *DescribeDomainBasicConfigsRequest {
	s.AccessType = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetCloudNativeProductId(v int32) *DescribeDomainBasicConfigsRequest {
	s.CloudNativeProductId = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetDomainKey(v string) *DescribeDomainBasicConfigsRequest {
	s.DomainKey = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetInstanceId(v string) *DescribeDomainBasicConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetPageNumber(v int32) *DescribeDomainBasicConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetPageSize(v int32) *DescribeDomainBasicConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetResourceGroupId(v string) *DescribeDomainBasicConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainBasicConfigsResponseBody struct {
	DomainConfigs []*DescribeDomainBasicConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainBasicConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBasicConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBasicConfigsResponseBody) SetDomainConfigs(v []*DescribeDomainBasicConfigsResponseBodyDomainConfigs) *DescribeDomainBasicConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBody) SetRequestId(v string) *DescribeDomainBasicConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBody) SetTotalCount(v int32) *DescribeDomainBasicConfigsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDomainBasicConfigsResponseBodyDomainConfigs struct {
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	AclStatus  *int32  `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	CcMode     *int32  `json:"CcMode,omitempty" xml:"CcMode,omitempty"`
	CcStatus   *int32  `json:"CcStatus,omitempty" xml:"CcStatus,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Version    *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
	WafMode    *int32  `json:"WafMode,omitempty" xml:"WafMode,omitempty"`
	WafStatus  *int32  `json:"WafStatus,omitempty" xml:"WafStatus,omitempty"`
}

func (s DescribeDomainBasicConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBasicConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetAccessType(v string) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.AccessType = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetAclStatus(v int32) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.AclStatus = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetCcMode(v int32) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.CcMode = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetCcStatus(v int32) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.CcStatus = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetDomain(v string) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.Domain = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetOwner(v string) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.Owner = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetStatus(v int32) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetVersion(v int64) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.Version = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetWafMode(v int32) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.WafMode = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponseBodyDomainConfigs) SetWafStatus(v int32) *DescribeDomainBasicConfigsResponseBodyDomainConfigs {
	s.WafStatus = &v
	return s
}

type DescribeDomainBasicConfigsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainBasicConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainBasicConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBasicConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBasicConfigsResponse) SetHeaders(v map[string]*string) *DescribeDomainBasicConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBasicConfigsResponse) SetBody(v *DescribeDomainBasicConfigsResponseBody) *DescribeDomainBasicConfigsResponse {
	s.Body = v
	return s
}

type DescribeDomainListRequest struct {
	DomainName      *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNames     []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSub           *int32    `json:"IsSub,omitempty" xml:"IsSub,omitempty"`
	PageNumber      *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainListRequest) SetDomainName(v string) *DescribeDomainListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainListRequest) SetDomainNames(v []*string) *DescribeDomainListRequest {
	s.DomainNames = v
	return s
}

func (s *DescribeDomainListRequest) SetInstanceId(v string) *DescribeDomainListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainListRequest) SetIsSub(v int32) *DescribeDomainListRequest {
	s.IsSub = &v
	return s
}

func (s *DescribeDomainListRequest) SetPageNumber(v int32) *DescribeDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainListRequest) SetPageSize(v int32) *DescribeDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainListRequest) SetResourceGroupId(v string) *DescribeDomainListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainListResponseBody struct {
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBody) SetDomainNames(v []*string) *DescribeDomainListResponseBody {
	s.DomainNames = v
	return s
}

func (s *DescribeDomainListResponseBody) SetRequestId(v string) *DescribeDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainListResponseBody) SetTotalCount(v int32) *DescribeDomainListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDomainListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponse) SetHeaders(v map[string]*string) *DescribeDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainListResponse) SetBody(v *DescribeDomainListResponseBody) *DescribeDomainListResponse {
	s.Body = v
	return s
}

type DescribeDomainNamesRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDomainNamesRequest) SetResourceGroupId(v string) *DescribeDomainNamesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainNamesResponseBody struct {
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesResponseBody) SetDomainNames(v []*string) *DescribeDomainNamesResponseBody {
	s.DomainNames = v
	return s
}

func (s *DescribeDomainNamesResponseBody) SetRequestId(v string) *DescribeDomainNamesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainNamesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDomainNamesResponse) SetBody(v *DescribeDomainNamesResponseBody) *DescribeDomainNamesResponse {
	s.Body = v
	return s
}

type DescribeDomainRuleGroupRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDomainRuleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRuleGroupRequest) SetDomain(v string) *DescribeDomainRuleGroupRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainRuleGroupRequest) SetInstanceId(v string) *DescribeDomainRuleGroupRequest {
	s.InstanceId = &v
	return s
}

type DescribeDomainRuleGroupResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleGroupId *int64  `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	WafAiStatus *int32  `json:"WafAiStatus,omitempty" xml:"WafAiStatus,omitempty"`
}

func (s DescribeDomainRuleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRuleGroupResponseBody) SetRequestId(v string) *DescribeDomainRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRuleGroupResponseBody) SetRuleGroupId(v int64) *DescribeDomainRuleGroupResponseBody {
	s.RuleGroupId = &v
	return s
}

func (s *DescribeDomainRuleGroupResponseBody) SetWafAiStatus(v int32) *DescribeDomainRuleGroupResponseBody {
	s.WafAiStatus = &v
	return s
}

type DescribeDomainRuleGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRuleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRuleGroupResponse) SetHeaders(v map[string]*string) *DescribeDomainRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRuleGroupResponse) SetBody(v *DescribeDomainRuleGroupResponseBody) *DescribeDomainRuleGroupResponse {
	s.Body = v
	return s
}

type DescribeInstanceInfoRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceInfoRequest) SetInstanceId(v string) *DescribeInstanceInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceInfoRequest) SetResourceGroupId(v string) *DescribeInstanceInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceInfoResponseBody struct {
	InstanceInfo *DescribeInstanceInfoResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceInfoResponseBody) SetInstanceInfo(v *DescribeInstanceInfoResponseBodyInstanceInfo) *DescribeInstanceInfoResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *DescribeInstanceInfoResponseBody) SetRequestId(v string) *DescribeInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceInfoResponseBodyInstanceInfo struct {
	EndDate          *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InDebt           *int32  `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PayType          *int32  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RemainDay        *int32  `json:"RemainDay,omitempty" xml:"RemainDay,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Trial            *int32  `json:"Trial,omitempty" xml:"Trial,omitempty"`
	Version          *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceInfoResponseBodyInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceInfoResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetEndDate(v int64) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetInDebt(v int32) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.InDebt = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetInstanceId(v string) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetPayType(v int32) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetRegion(v string) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.Region = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetRemainDay(v int32) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.RemainDay = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetStatus(v int32) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetSubscriptionType(v string) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetTrial(v int32) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.Trial = &v
	return s
}

func (s *DescribeInstanceInfoResponseBodyInstanceInfo) SetVersion(v string) *DescribeInstanceInfoResponseBodyInstanceInfo {
	s.Version = &v
	return s
}

type DescribeInstanceInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceInfoResponse) SetHeaders(v map[string]*string) *DescribeInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceInfoResponse) SetBody(v *DescribeInstanceInfoResponseBody) *DescribeInstanceInfoResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecInfoRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceSpecInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecInfoRequest) SetInstanceId(v string) *DescribeInstanceSpecInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecInfoRequest) SetResourceGroupId(v string) *DescribeInstanceSpecInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceSpecInfoResponseBody struct {
	ExpireTime        *int64                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId        *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSpecInfos []*DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos `json:"InstanceSpecInfos,omitempty" xml:"InstanceSpecInfos,omitempty" type:"Repeated"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version           *string                                                  `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceSpecInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecInfoResponseBody) SetExpireTime(v int64) *DescribeInstanceSpecInfoResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceSpecInfoResponseBody) SetInstanceId(v string) *DescribeInstanceSpecInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecInfoResponseBody) SetInstanceSpecInfos(v []*DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos) *DescribeInstanceSpecInfoResponseBody {
	s.InstanceSpecInfos = v
	return s
}

func (s *DescribeInstanceSpecInfoResponseBody) SetRequestId(v string) *DescribeInstanceSpecInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecInfoResponseBody) SetVersion(v string) *DescribeInstanceSpecInfoResponseBody {
	s.Version = &v
	return s
}

type DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos) SetCode(v string) *DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos {
	s.Code = &v
	return s
}

func (s *DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos) SetValue(v string) *DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos {
	s.Value = &v
	return s
}

type DescribeInstanceSpecInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecInfoResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecInfoResponse) SetBody(v *DescribeInstanceSpecInfoResponseBody) *DescribeInstanceSpecInfoResponse {
	s.Body = v
	return s
}

type DescribeLogServiceStatusRequest struct {
	DomainNames     []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber      *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Region          *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeLogServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogServiceStatusRequest) SetDomainNames(v []*string) *DescribeLogServiceStatusRequest {
	s.DomainNames = v
	return s
}

func (s *DescribeLogServiceStatusRequest) SetInstanceId(v string) *DescribeLogServiceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogServiceStatusRequest) SetPageNumber(v int32) *DescribeLogServiceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogServiceStatusRequest) SetPageSize(v int32) *DescribeLogServiceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogServiceStatusRequest) SetRegion(v string) *DescribeLogServiceStatusRequest {
	s.Region = &v
	return s
}

func (s *DescribeLogServiceStatusRequest) SetResourceGroupId(v string) *DescribeLogServiceStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeLogServiceStatusResponseBody struct {
	DomainStatus []*DescribeLogServiceStatusResponseBodyDomainStatus `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty" type:"Repeated"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLogServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogServiceStatusResponseBody) SetDomainStatus(v []*DescribeLogServiceStatusResponseBodyDomainStatus) *DescribeLogServiceStatusResponseBody {
	s.DomainStatus = v
	return s
}

func (s *DescribeLogServiceStatusResponseBody) SetRequestId(v string) *DescribeLogServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogServiceStatusResponseBody) SetTotalCount(v int32) *DescribeLogServiceStatusResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeLogServiceStatusResponseBodyDomainStatus struct {
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SlsLogActive *int32  `json:"SlsLogActive,omitempty" xml:"SlsLogActive,omitempty"`
}

func (s DescribeLogServiceStatusResponseBodyDomainStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogServiceStatusResponseBodyDomainStatus) GoString() string {
	return s.String()
}

func (s *DescribeLogServiceStatusResponseBodyDomainStatus) SetDomain(v string) *DescribeLogServiceStatusResponseBodyDomainStatus {
	s.Domain = &v
	return s
}

func (s *DescribeLogServiceStatusResponseBodyDomainStatus) SetSlsLogActive(v int32) *DescribeLogServiceStatusResponseBodyDomainStatus {
	s.SlsLogActive = &v
	return s
}

type DescribeLogServiceStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeLogServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogServiceStatusResponse) SetBody(v *DescribeLogServiceStatusResponseBody) *DescribeLogServiceStatusResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleCodeConfigRequest struct {
	CodeType        *int32  `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	CodeValue       *int32  `json:"CodeValue,omitempty" xml:"CodeValue,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeProtectionModuleCodeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleCodeConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleCodeConfigRequest) SetCodeType(v int32) *DescribeProtectionModuleCodeConfigRequest {
	s.CodeType = &v
	return s
}

func (s *DescribeProtectionModuleCodeConfigRequest) SetCodeValue(v int32) *DescribeProtectionModuleCodeConfigRequest {
	s.CodeValue = &v
	return s
}

func (s *DescribeProtectionModuleCodeConfigRequest) SetInstanceId(v string) *DescribeProtectionModuleCodeConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProtectionModuleCodeConfigRequest) SetResourceGroupId(v string) *DescribeProtectionModuleCodeConfigRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeProtectionModuleCodeConfigResponseBody struct {
	CodeConfigs *string `json:"CodeConfigs,omitempty" xml:"CodeConfigs,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtectionModuleCodeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleCodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleCodeConfigResponseBody) SetCodeConfigs(v string) *DescribeProtectionModuleCodeConfigResponseBody {
	s.CodeConfigs = &v
	return s
}

func (s *DescribeProtectionModuleCodeConfigResponseBody) SetRequestId(v string) *DescribeProtectionModuleCodeConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtectionModuleCodeConfigResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProtectionModuleCodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtectionModuleCodeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleCodeConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleCodeConfigResponse) SetHeaders(v map[string]*string) *DescribeProtectionModuleCodeConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtectionModuleCodeConfigResponse) SetBody(v *DescribeProtectionModuleCodeConfigResponseBody) *DescribeProtectionModuleCodeConfigResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleRulesRequest struct {
	DefenseType     *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query           *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeProtectionModuleRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesRequest) SetDefenseType(v string) *DescribeProtectionModuleRulesRequest {
	s.DefenseType = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetDomain(v string) *DescribeProtectionModuleRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetInstanceId(v string) *DescribeProtectionModuleRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetLang(v string) *DescribeProtectionModuleRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetPageNumber(v int32) *DescribeProtectionModuleRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetPageSize(v int32) *DescribeProtectionModuleRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetQuery(v string) *DescribeProtectionModuleRulesRequest {
	s.Query = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetResourceGroupId(v string) *DescribeProtectionModuleRulesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeProtectionModuleRulesResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules      []*DescribeProtectionModuleRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProtectionModuleRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesResponseBody) SetRequestId(v string) *DescribeProtectionModuleRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBody) SetRules(v []*DescribeProtectionModuleRulesResponseBodyRules) *DescribeProtectionModuleRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBody) SetTotalCount(v int32) *DescribeProtectionModuleRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProtectionModuleRulesResponseBodyRules struct {
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	RuleId  *int64                 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Status  *int64                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Time    *int64                 `json:"Time,omitempty" xml:"Time,omitempty"`
	Version *int64                 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeProtectionModuleRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesResponseBodyRules) SetContent(v map[string]interface{}) *DescribeProtectionModuleRulesResponseBodyRules {
	s.Content = v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyRules) SetRuleId(v int64) *DescribeProtectionModuleRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyRules) SetStatus(v int64) *DescribeProtectionModuleRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyRules) SetTime(v int64) *DescribeProtectionModuleRulesResponseBodyRules {
	s.Time = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyRules) SetVersion(v int64) *DescribeProtectionModuleRulesResponseBodyRules {
	s.Version = &v
	return s
}

type DescribeProtectionModuleRulesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProtectionModuleRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtectionModuleRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesResponse) SetHeaders(v map[string]*string) *DescribeProtectionModuleRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtectionModuleRulesResponse) SetBody(v *DescribeProtectionModuleRulesResponseBody) *DescribeProtectionModuleRulesResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleStatusRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProtectionModuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleStatusRequest) SetDefenseType(v string) *DescribeProtectionModuleStatusRequest {
	s.DefenseType = &v
	return s
}

func (s *DescribeProtectionModuleStatusRequest) SetDomain(v string) *DescribeProtectionModuleStatusRequest {
	s.Domain = &v
	return s
}

func (s *DescribeProtectionModuleStatusRequest) SetInstanceId(v string) *DescribeProtectionModuleStatusRequest {
	s.InstanceId = &v
	return s
}

type DescribeProtectionModuleStatusResponseBody struct {
	ModuleStatus *int32  `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtectionModuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleStatusResponseBody) SetModuleStatus(v int32) *DescribeProtectionModuleStatusResponseBody {
	s.ModuleStatus = &v
	return s
}

func (s *DescribeProtectionModuleStatusResponseBody) SetRequestId(v string) *DescribeProtectionModuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtectionModuleStatusResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProtectionModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProtectionModuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleStatusResponse) SetHeaders(v map[string]*string) *DescribeProtectionModuleStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtectionModuleStatusResponse) SetBody(v *DescribeProtectionModuleStatusResponseBody) *DescribeProtectionModuleStatusResponse {
	s.Body = v
	return s
}

type DescribeWafSourceIpSegmentRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWafSourceIpSegmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentRequest) SetInstanceId(v string) *DescribeWafSourceIpSegmentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetResourceGroupId(v string) *DescribeWafSourceIpSegmentRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeWafSourceIpSegmentResponseBody struct {
	IpV6s     *string `json:"IpV6s,omitempty" xml:"IpV6s,omitempty"`
	Ips       *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWafSourceIpSegmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetIpV6s(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.IpV6s = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetIps(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.Ips = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetRequestId(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWafSourceIpSegmentResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWafSourceIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWafSourceIpSegmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponse) SetHeaders(v map[string]*string) *DescribeWafSourceIpSegmentResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) SetBody(v *DescribeWafSourceIpSegmentResponseBody) *DescribeWafSourceIpSegmentResponse {
	s.Body = v
	return s
}

type ModifyDomainRequest struct {
	AccessHeaderMode     *int32  `json:"AccessHeaderMode,omitempty" xml:"AccessHeaderMode,omitempty"`
	AccessHeaders        *string `json:"AccessHeaders,omitempty" xml:"AccessHeaders,omitempty"`
	AccessType           *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CloudNativeInstances *string `json:"CloudNativeInstances,omitempty" xml:"CloudNativeInstances,omitempty"`
	ClusterType          *int32  `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ConnectionTime       *int32  `json:"ConnectionTime,omitempty" xml:"ConnectionTime,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Http2Port            *string `json:"Http2Port,omitempty" xml:"Http2Port,omitempty"`
	HttpPort             *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpToUserIp         *int32  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsPort            *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	HttpsRedirect        *int32  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpFollowStatus       *int32  `json:"IpFollowStatus,omitempty" xml:"IpFollowStatus,omitempty"`
	IsAccessProduct      *int32  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	LoadBalancing        *int32  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	LogHeaders           *string `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty"`
	ReadTime             *int32  `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	SniHost              *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	SniStatus            *int32  `json:"SniStatus,omitempty" xml:"SniStatus,omitempty"`
	SourceIps            *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	WriteTime            *int32  `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
}

func (s ModifyDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequest) SetAccessHeaderMode(v int32) *ModifyDomainRequest {
	s.AccessHeaderMode = &v
	return s
}

func (s *ModifyDomainRequest) SetAccessHeaders(v string) *ModifyDomainRequest {
	s.AccessHeaders = &v
	return s
}

func (s *ModifyDomainRequest) SetAccessType(v string) *ModifyDomainRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainRequest) SetCloudNativeInstances(v string) *ModifyDomainRequest {
	s.CloudNativeInstances = &v
	return s
}

func (s *ModifyDomainRequest) SetClusterType(v int32) *ModifyDomainRequest {
	s.ClusterType = &v
	return s
}

func (s *ModifyDomainRequest) SetConnectionTime(v int32) *ModifyDomainRequest {
	s.ConnectionTime = &v
	return s
}

func (s *ModifyDomainRequest) SetDomain(v string) *ModifyDomainRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainRequest) SetHttp2Port(v string) *ModifyDomainRequest {
	s.Http2Port = &v
	return s
}

func (s *ModifyDomainRequest) SetHttpPort(v string) *ModifyDomainRequest {
	s.HttpPort = &v
	return s
}

func (s *ModifyDomainRequest) SetHttpToUserIp(v int32) *ModifyDomainRequest {
	s.HttpToUserIp = &v
	return s
}

func (s *ModifyDomainRequest) SetHttpsPort(v string) *ModifyDomainRequest {
	s.HttpsPort = &v
	return s
}

func (s *ModifyDomainRequest) SetHttpsRedirect(v int32) *ModifyDomainRequest {
	s.HttpsRedirect = &v
	return s
}

func (s *ModifyDomainRequest) SetInstanceId(v string) *ModifyDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainRequest) SetIpFollowStatus(v int32) *ModifyDomainRequest {
	s.IpFollowStatus = &v
	return s
}

func (s *ModifyDomainRequest) SetIsAccessProduct(v int32) *ModifyDomainRequest {
	s.IsAccessProduct = &v
	return s
}

func (s *ModifyDomainRequest) SetLoadBalancing(v int32) *ModifyDomainRequest {
	s.LoadBalancing = &v
	return s
}

func (s *ModifyDomainRequest) SetLogHeaders(v string) *ModifyDomainRequest {
	s.LogHeaders = &v
	return s
}

func (s *ModifyDomainRequest) SetReadTime(v int32) *ModifyDomainRequest {
	s.ReadTime = &v
	return s
}

func (s *ModifyDomainRequest) SetSniHost(v string) *ModifyDomainRequest {
	s.SniHost = &v
	return s
}

func (s *ModifyDomainRequest) SetSniStatus(v int32) *ModifyDomainRequest {
	s.SniStatus = &v
	return s
}

func (s *ModifyDomainRequest) SetSourceIps(v string) *ModifyDomainRequest {
	s.SourceIps = &v
	return s
}

func (s *ModifyDomainRequest) SetWriteTime(v int32) *ModifyDomainRequest {
	s.WriteTime = &v
	return s
}

type ModifyDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBody) SetRequestId(v string) *ModifyDomainResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponse) SetHeaders(v map[string]*string) *ModifyDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainResponse) SetBody(v *ModifyDomainResponseBody) *ModifyDomainResponse {
	s.Body = v
	return s
}

type ModifyDomainIpv6StatusRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled    *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyDomainIpv6StatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainIpv6StatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainIpv6StatusRequest) SetDomain(v string) *ModifyDomainIpv6StatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainIpv6StatusRequest) SetEnabled(v string) *ModifyDomainIpv6StatusRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyDomainIpv6StatusRequest) SetInstanceId(v string) *ModifyDomainIpv6StatusRequest {
	s.InstanceId = &v
	return s
}

type ModifyDomainIpv6StatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainIpv6StatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainIpv6StatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainIpv6StatusResponseBody) SetRequestId(v string) *ModifyDomainIpv6StatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainIpv6StatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDomainIpv6StatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainIpv6StatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainIpv6StatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainIpv6StatusResponse) SetHeaders(v map[string]*string) *ModifyDomainIpv6StatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainIpv6StatusResponse) SetBody(v *ModifyDomainIpv6StatusResponseBody) *ModifyDomainIpv6StatusResponse {
	s.Body = v
	return s
}

type ModifyLogRetrievalStatusRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled    *int32  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyLogRetrievalStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogRetrievalStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogRetrievalStatusRequest) SetDomain(v string) *ModifyLogRetrievalStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyLogRetrievalStatusRequest) SetEnabled(v int32) *ModifyLogRetrievalStatusRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyLogRetrievalStatusRequest) SetInstanceId(v string) *ModifyLogRetrievalStatusRequest {
	s.InstanceId = &v
	return s
}

type ModifyLogRetrievalStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogRetrievalStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogRetrievalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogRetrievalStatusResponseBody) SetRequestId(v string) *ModifyLogRetrievalStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogRetrievalStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLogRetrievalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogRetrievalStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogRetrievalStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogRetrievalStatusResponse) SetHeaders(v map[string]*string) *ModifyLogRetrievalStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogRetrievalStatusResponse) SetBody(v *ModifyLogRetrievalStatusResponseBody) *ModifyLogRetrievalStatusResponse {
	s.Body = v
	return s
}

type ModifyLogServiceStatusRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enabled    *int32  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyLogServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogServiceStatusRequest) SetDomain(v string) *ModifyLogServiceStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyLogServiceStatusRequest) SetEnabled(v int32) *ModifyLogServiceStatusRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyLogServiceStatusRequest) SetInstanceId(v string) *ModifyLogServiceStatusRequest {
	s.InstanceId = &v
	return s
}

type ModifyLogServiceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogServiceStatusResponseBody) SetRequestId(v string) *ModifyLogServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogServiceStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLogServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogServiceStatusResponse) SetHeaders(v map[string]*string) *ModifyLogServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogServiceStatusResponse) SetBody(v *ModifyLogServiceStatusResponseBody) *ModifyLogServiceStatusResponse {
	s.Body = v
	return s
}

type ModifyProtectionModuleModeRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode        *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyProtectionModuleModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleModeRequest) SetDefenseType(v string) *ModifyProtectionModuleModeRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyProtectionModuleModeRequest) SetDomain(v string) *ModifyProtectionModuleModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionModuleModeRequest) SetInstanceId(v string) *ModifyProtectionModuleModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionModuleModeRequest) SetMode(v int32) *ModifyProtectionModuleModeRequest {
	s.Mode = &v
	return s
}

type ModifyProtectionModuleModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtectionModuleModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleModeResponseBody) SetRequestId(v string) *ModifyProtectionModuleModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtectionModuleModeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyProtectionModuleModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtectionModuleModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleModeResponse) SetHeaders(v map[string]*string) *ModifyProtectionModuleModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtectionModuleModeResponse) SetBody(v *ModifyProtectionModuleModeResponseBody) *ModifyProtectionModuleModeResponse {
	s.Body = v
	return s
}

type ModifyProtectionModuleRuleRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LockVersion *int64  `json:"LockVersion,omitempty" xml:"LockVersion,omitempty"`
	Rule        *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyProtectionModuleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleRuleRequest) SetDefenseType(v string) *ModifyProtectionModuleRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyProtectionModuleRuleRequest) SetDomain(v string) *ModifyProtectionModuleRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionModuleRuleRequest) SetInstanceId(v string) *ModifyProtectionModuleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionModuleRuleRequest) SetLockVersion(v int64) *ModifyProtectionModuleRuleRequest {
	s.LockVersion = &v
	return s
}

func (s *ModifyProtectionModuleRuleRequest) SetRule(v string) *ModifyProtectionModuleRuleRequest {
	s.Rule = &v
	return s
}

func (s *ModifyProtectionModuleRuleRequest) SetRuleId(v int64) *ModifyProtectionModuleRuleRequest {
	s.RuleId = &v
	return s
}

type ModifyProtectionModuleRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtectionModuleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleRuleResponseBody) SetRequestId(v string) *ModifyProtectionModuleRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtectionModuleRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyProtectionModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtectionModuleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleRuleResponse) SetHeaders(v map[string]*string) *ModifyProtectionModuleRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtectionModuleRuleResponse) SetBody(v *ModifyProtectionModuleRuleResponseBody) *ModifyProtectionModuleRuleResponse {
	s.Body = v
	return s
}

type ModifyProtectionModuleStatusRequest struct {
	DefenseType  *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModuleStatus *int32  `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
}

func (s ModifyProtectionModuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleStatusRequest) SetDefenseType(v string) *ModifyProtectionModuleStatusRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyProtectionModuleStatusRequest) SetDomain(v string) *ModifyProtectionModuleStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionModuleStatusRequest) SetInstanceId(v string) *ModifyProtectionModuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionModuleStatusRequest) SetModuleStatus(v int32) *ModifyProtectionModuleStatusRequest {
	s.ModuleStatus = &v
	return s
}

type ModifyProtectionModuleStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtectionModuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleStatusResponseBody) SetRequestId(v string) *ModifyProtectionModuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtectionModuleStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyProtectionModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtectionModuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionModuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtectionModuleStatusResponse) SetHeaders(v map[string]*string) *ModifyProtectionModuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtectionModuleStatusResponse) SetBody(v *ModifyProtectionModuleStatusResponseBody) *ModifyProtectionModuleStatusResponse {
	s.Body = v
	return s
}

type ModifyProtectionRuleCacheStatusRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ModifyProtectionRuleCacheStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleCacheStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetDefenseType(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetDomain(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetInstanceId(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetRuleId(v int64) *ModifyProtectionRuleCacheStatusRequest {
	s.RuleId = &v
	return s
}

type ModifyProtectionRuleCacheStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtectionRuleCacheStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleCacheStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleCacheStatusResponseBody) SetRequestId(v string) *ModifyProtectionRuleCacheStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtectionRuleCacheStatusResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyProtectionRuleCacheStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtectionRuleCacheStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleCacheStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleCacheStatusResponse) SetHeaders(v map[string]*string) *ModifyProtectionRuleCacheStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtectionRuleCacheStatusResponse) SetBody(v *ModifyProtectionRuleCacheStatusResponseBody) *ModifyProtectionRuleCacheStatusResponse {
	s.Body = v
	return s
}

type ModifyProtectionRuleStatusRequest struct {
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LockVersion *int64  `json:"LockVersion,omitempty" xml:"LockVersion,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleStatus  *int32  `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ModifyProtectionRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleStatusRequest) SetDefenseType(v string) *ModifyProtectionRuleStatusRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetDomain(v string) *ModifyProtectionRuleStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetInstanceId(v string) *ModifyProtectionRuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetLockVersion(v int64) *ModifyProtectionRuleStatusRequest {
	s.LockVersion = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetRuleId(v int64) *ModifyProtectionRuleStatusRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetRuleStatus(v int32) *ModifyProtectionRuleStatusRequest {
	s.RuleStatus = &v
	return s
}

type ModifyProtectionRuleStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProtectionRuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleStatusResponseBody) SetRequestId(v string) *ModifyProtectionRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProtectionRuleStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyProtectionRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProtectionRuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleStatusResponse) SetHeaders(v map[string]*string) *ModifyProtectionRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtectionRuleStatusResponse) SetBody(v *ModifyProtectionRuleStatusResponseBody) *ModifyProtectionRuleStatusResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type SetDomainRuleGroupRequest struct {
	Domains         *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RuleGroupId     *int64  `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	WafAiStatus     *int32  `json:"WafAiStatus,omitempty" xml:"WafAiStatus,omitempty"`
	WafVersion      *int64  `json:"WafVersion,omitempty" xml:"WafVersion,omitempty"`
}

func (s SetDomainRuleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *SetDomainRuleGroupRequest) SetDomains(v string) *SetDomainRuleGroupRequest {
	s.Domains = &v
	return s
}

func (s *SetDomainRuleGroupRequest) SetInstanceId(v string) *SetDomainRuleGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SetDomainRuleGroupRequest) SetResourceGroupId(v string) *SetDomainRuleGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SetDomainRuleGroupRequest) SetRuleGroupId(v int64) *SetDomainRuleGroupRequest {
	s.RuleGroupId = &v
	return s
}

func (s *SetDomainRuleGroupRequest) SetWafAiStatus(v int32) *SetDomainRuleGroupRequest {
	s.WafAiStatus = &v
	return s
}

func (s *SetDomainRuleGroupRequest) SetWafVersion(v int64) *SetDomainRuleGroupRequest {
	s.WafVersion = &v
	return s
}

type SetDomainRuleGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainRuleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainRuleGroupResponseBody) SetRequestId(v string) *SetDomainRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainRuleGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainRuleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *SetDomainRuleGroupResponse) SetHeaders(v map[string]*string) *SetDomainRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *SetDomainRuleGroupResponse) SetBody(v *SetDomainRuleGroupResponseBody) *SetDomainRuleGroupResponse {
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

func (client *Client) CreateCertificateWithOptions(request *CreateCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		query["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateName)) {
		query["CertificateName"] = request.CertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertificate(request *CreateCertificateRequest) (_result *CreateCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateResponse{}
	_body, _err := client.CreateCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCertificateByCertificateIdWithOptions(request *CreateCertificateByCertificateIdRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateByCertificateIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertificateByCertificateId"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertificateByCertificateIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertificateByCertificateId(request *CreateCertificateByCertificateIdRequest) (_result *CreateCertificateByCertificateIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateByCertificateIdResponse{}
	_body, _err := client.CreateCertificateByCertificateIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessHeaderMode)) {
		query["AccessHeaderMode"] = request.AccessHeaderMode
	}

	if !tea.BoolValue(util.IsUnset(request.AccessHeaders)) {
		query["AccessHeaders"] = request.AccessHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.CloudNativeInstances)) {
		query["CloudNativeInstances"] = request.CloudNativeInstances
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionTime)) {
		query["ConnectionTime"] = request.ConnectionTime
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Http2Port)) {
		query["Http2Port"] = request.Http2Port
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

	if !tea.BoolValue(util.IsUnset(request.IpFollowStatus)) {
		query["IpFollowStatus"] = request.IpFollowStatus
	}

	if !tea.BoolValue(util.IsUnset(request.IsAccessProduct)) {
		query["IsAccessProduct"] = request.IsAccessProduct
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancing)) {
		query["LoadBalancing"] = request.LoadBalancing
	}

	if !tea.BoolValue(util.IsUnset(request.LogHeaders)) {
		query["LogHeaders"] = request.LogHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.ReadTime)) {
		query["ReadTime"] = request.ReadTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SniHost)) {
		query["SniHost"] = request.SniHost
	}

	if !tea.BoolValue(util.IsUnset(request.SniStatus)) {
		query["SniStatus"] = request.SniStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIps)) {
		query["SourceIps"] = request.SourceIps
	}

	if !tea.BoolValue(util.IsUnset(request.WriteTime)) {
		query["WriteTime"] = request.WriteTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProtectionModuleRuleWithOptions(request *CreateProtectionModuleRuleRequest, runtime *util.RuntimeOptions) (_result *CreateProtectionModuleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		query["Rule"] = request.Rule
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProtectionModuleRule"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProtectionModuleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProtectionModuleRule(request *CreateProtectionModuleRuleRequest) (_result *CreateProtectionModuleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProtectionModuleRuleResponse{}
	_body, _err := client.CreateProtectionModuleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProtectionModuleRuleWithOptions(request *DeleteProtectionModuleRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteProtectionModuleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProtectionModuleRule"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProtectionModuleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProtectionModuleRule(request *DeleteProtectionModuleRuleRequest) (_result *DeleteProtectionModuleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProtectionModuleRuleResponse{}
	_body, _err := client.DeleteProtectionModuleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertMatchStatusWithOptions(request *DescribeCertMatchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeCertMatchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		query["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCertMatchStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCertMatchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCertMatchStatus(request *DescribeCertMatchStatusRequest) (_result *DescribeCertMatchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertMatchStatusResponse{}
	_body, _err := client.DescribeCertMatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertificatesWithOptions(request *DescribeCertificatesRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificatesResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCertificates"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCertificates(request *DescribeCertificatesRequest) (_result *DescribeCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificatesResponse{}
	_body, _err := client.DescribeCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainWithOptions(request *DescribeDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomain"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomain(request *DescribeDomainRequest) (_result *DescribeDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainResponse{}
	_body, _err := client.DescribeDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAdvanceConfigsWithOptions(request *DescribeDomainAdvanceConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAdvanceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainList)) {
		query["DomainList"] = request.DomainList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainAdvanceConfigs"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainAdvanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAdvanceConfigs(request *DescribeDomainAdvanceConfigsRequest) (_result *DescribeDomainAdvanceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAdvanceConfigsResponse{}
	_body, _err := client.DescribeDomainAdvanceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainBasicConfigsWithOptions(request *DescribeDomainBasicConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainBasicConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.CloudNativeProductId)) {
		query["CloudNativeProductId"] = request.CloudNativeProductId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainKey)) {
		query["DomainKey"] = request.DomainKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainBasicConfigs"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainBasicConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainBasicConfigs(request *DescribeDomainBasicConfigsRequest) (_result *DescribeDomainBasicConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainBasicConfigsResponse{}
	_body, _err := client.DescribeDomainBasicConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainListWithOptions(request *DescribeDomainListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSub)) {
		query["IsSub"] = request.IsSub
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainList"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainList(request *DescribeDomainListRequest) (_result *DescribeDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainListResponse{}
	_body, _err := client.DescribeDomainListWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainNames"),
		Version:     tea.String("2019-09-10"),
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

func (client *Client) DescribeDomainRuleGroupWithOptions(request *DescribeDomainRuleGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRuleGroupResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainRuleGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainRuleGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRuleGroup(request *DescribeDomainRuleGroupRequest) (_result *DescribeDomainRuleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRuleGroupResponse{}
	_body, _err := client.DescribeDomainRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceInfoWithOptions(request *DescribeInstanceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceInfo"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceInfo(request *DescribeInstanceInfoRequest) (_result *DescribeInstanceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceInfoResponse{}
	_body, _err := client.DescribeInstanceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecInfoWithOptions(request *DescribeInstanceSpecInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSpecInfo"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSpecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecInfo(request *DescribeInstanceSpecInfoRequest) (_result *DescribeInstanceSpecInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecInfoResponse{}
	_body, _err := client.DescribeInstanceSpecInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogServiceStatusWithOptions(request *DescribeLogServiceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeLogServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogServiceStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogServiceStatus(request *DescribeLogServiceStatusRequest) (_result *DescribeLogServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogServiceStatusResponse{}
	_body, _err := client.DescribeLogServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtectionModuleCodeConfigWithOptions(request *DescribeProtectionModuleCodeConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeProtectionModuleCodeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeType)) {
		query["CodeType"] = request.CodeType
	}

	if !tea.BoolValue(util.IsUnset(request.CodeValue)) {
		query["CodeValue"] = request.CodeValue
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtectionModuleCodeConfig"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtectionModuleCodeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtectionModuleCodeConfig(request *DescribeProtectionModuleCodeConfigRequest) (_result *DescribeProtectionModuleCodeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtectionModuleCodeConfigResponse{}
	_body, _err := client.DescribeProtectionModuleCodeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtectionModuleRulesWithOptions(request *DescribeProtectionModuleRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeProtectionModuleRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtectionModuleRules"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtectionModuleRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtectionModuleRules(request *DescribeProtectionModuleRulesRequest) (_result *DescribeProtectionModuleRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtectionModuleRulesResponse{}
	_body, _err := client.DescribeProtectionModuleRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProtectionModuleStatusWithOptions(request *DescribeProtectionModuleStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeProtectionModuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtectionModuleStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtectionModuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProtectionModuleStatus(request *DescribeProtectionModuleStatusRequest) (_result *DescribeProtectionModuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtectionModuleStatusResponse{}
	_body, _err := client.DescribeProtectionModuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWafSourceIpSegmentWithOptions(request *DescribeWafSourceIpSegmentRequest, runtime *util.RuntimeOptions) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWafSourceIpSegment"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWafSourceIpSegment(request *DescribeWafSourceIpSegmentRequest) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.DescribeWafSourceIpSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainWithOptions(request *ModifyDomainRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessHeaderMode)) {
		query["AccessHeaderMode"] = request.AccessHeaderMode
	}

	if !tea.BoolValue(util.IsUnset(request.AccessHeaders)) {
		query["AccessHeaders"] = request.AccessHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.CloudNativeInstances)) {
		query["CloudNativeInstances"] = request.CloudNativeInstances
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionTime)) {
		query["ConnectionTime"] = request.ConnectionTime
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Http2Port)) {
		query["Http2Port"] = request.Http2Port
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

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpFollowStatus)) {
		query["IpFollowStatus"] = request.IpFollowStatus
	}

	if !tea.BoolValue(util.IsUnset(request.IsAccessProduct)) {
		query["IsAccessProduct"] = request.IsAccessProduct
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancing)) {
		query["LoadBalancing"] = request.LoadBalancing
	}

	if !tea.BoolValue(util.IsUnset(request.LogHeaders)) {
		query["LogHeaders"] = request.LogHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.ReadTime)) {
		query["ReadTime"] = request.ReadTime
	}

	if !tea.BoolValue(util.IsUnset(request.SniHost)) {
		query["SniHost"] = request.SniHost
	}

	if !tea.BoolValue(util.IsUnset(request.SniStatus)) {
		query["SniStatus"] = request.SniStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIps)) {
		query["SourceIps"] = request.SourceIps
	}

	if !tea.BoolValue(util.IsUnset(request.WriteTime)) {
		query["WriteTime"] = request.WriteTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomain"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomain(request *ModifyDomainRequest) (_result *ModifyDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainResponse{}
	_body, _err := client.ModifyDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainIpv6StatusWithOptions(request *ModifyDomainIpv6StatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainIpv6StatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomainIpv6Status"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDomainIpv6StatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainIpv6Status(request *ModifyDomainIpv6StatusRequest) (_result *ModifyDomainIpv6StatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainIpv6StatusResponse{}
	_body, _err := client.ModifyDomainIpv6StatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogRetrievalStatusWithOptions(request *ModifyLogRetrievalStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyLogRetrievalStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLogRetrievalStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLogRetrievalStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogRetrievalStatus(request *ModifyLogRetrievalStatusRequest) (_result *ModifyLogRetrievalStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogRetrievalStatusResponse{}
	_body, _err := client.ModifyLogRetrievalStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogServiceStatusWithOptions(request *ModifyLogServiceStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyLogServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLogServiceStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLogServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogServiceStatus(request *ModifyLogServiceStatusRequest) (_result *ModifyLogServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogServiceStatusResponse{}
	_body, _err := client.ModifyLogServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtectionModuleModeWithOptions(request *ModifyProtectionModuleModeRequest, runtime *util.RuntimeOptions) (_result *ModifyProtectionModuleModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtectionModuleMode"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtectionModuleModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtectionModuleMode(request *ModifyProtectionModuleModeRequest) (_result *ModifyProtectionModuleModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtectionModuleModeResponse{}
	_body, _err := client.ModifyProtectionModuleModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtectionModuleRuleWithOptions(request *ModifyProtectionModuleRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyProtectionModuleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LockVersion)) {
		query["LockVersion"] = request.LockVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Rule)) {
		query["Rule"] = request.Rule
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtectionModuleRule"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtectionModuleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtectionModuleRule(request *ModifyProtectionModuleRuleRequest) (_result *ModifyProtectionModuleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtectionModuleRuleResponse{}
	_body, _err := client.ModifyProtectionModuleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtectionModuleStatusWithOptions(request *ModifyProtectionModuleStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyProtectionModuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleStatus)) {
		query["ModuleStatus"] = request.ModuleStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtectionModuleStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtectionModuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtectionModuleStatus(request *ModifyProtectionModuleStatusRequest) (_result *ModifyProtectionModuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtectionModuleStatusResponse{}
	_body, _err := client.ModifyProtectionModuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtectionRuleCacheStatusWithOptions(request *ModifyProtectionRuleCacheStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyProtectionRuleCacheStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtectionRuleCacheStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtectionRuleCacheStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtectionRuleCacheStatus(request *ModifyProtectionRuleCacheStatusRequest) (_result *ModifyProtectionRuleCacheStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtectionRuleCacheStatusResponse{}
	_body, _err := client.ModifyProtectionRuleCacheStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProtectionRuleStatusWithOptions(request *ModifyProtectionRuleStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyProtectionRuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseType)) {
		query["DefenseType"] = request.DefenseType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LockVersion)) {
		query["LockVersion"] = request.LockVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleStatus)) {
		query["RuleStatus"] = request.RuleStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProtectionRuleStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProtectionRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProtectionRuleStatus(request *ModifyProtectionRuleStatusRequest) (_result *ModifyProtectionRuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProtectionRuleStatusResponse{}
	_body, _err := client.ModifyProtectionRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainRuleGroupWithOptions(request *SetDomainRuleGroupRequest, runtime *util.RuntimeOptions) (_result *SetDomainRuleGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domains)) {
		query["Domains"] = request.Domains
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleGroupId)) {
		query["RuleGroupId"] = request.RuleGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.WafAiStatus)) {
		query["WafAiStatus"] = request.WafAiStatus
	}

	if !tea.BoolValue(util.IsUnset(request.WafVersion)) {
		query["WafVersion"] = request.WafVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomainRuleGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDomainRuleGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainRuleGroup(request *SetDomainRuleGroupRequest) (_result *SetDomainRuleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainRuleGroupResponse{}
	_body, _err := client.SetDomainRuleGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
