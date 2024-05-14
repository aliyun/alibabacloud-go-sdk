// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateCertificateRequest struct {
	// example:
	//
	// -----BEGIN CERTIFICATE----- 62EcYPWd2Oy1vs6MTXcJSfN9Z7rZ9fmxWr2BFN2XbahgnsSXM48ixZJ4krc+1M+j2kcubVpsE2cgHdj4v8H6jUz9Ji4mr7vMNS6dXv8PUkl/qoDeNGCNdyTS5NIL5ir+g92cL8IGOkjgvhlqt9vc65Cgb4mL+n5+DV9uOyTZTW/MojmlgfUekC2xiXa54nxJf17Y1TADGSbyJbsC0Q9nIrHsPl8YKkvRWvIAqYxXZ7wRwWWmv4TMxFhWRiNY7yZIo2ZUhl02SIDNggIEeg== -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CertName
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- DADTPZoOHd9WtZ3UKHJTRgNQmioPQn2bqdKHop+B/dn/4VZL7Jt8zSDGM9sTMThLyvsmLQKBgQCr+ujntC1kN6pGBj2Fw2l/EA/W3rYEce2tyhjgmG7rZ+A/jVE9fld5sQra6ZdwBcQJaiygoIYoaMF2EjRwc0qwHaluq0C15f6ujSoHh2e+D5zdmkTg/3NKNjqNv6xA2gYpinVDzFdZ9Zujxvuh9o4Vqf0YF8bv5UK5G04RtKadOw== -----END RSA PRIVATE KEY-----
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *CreateCertificateRequest) SetRegionId(v string) *CreateCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCertificateRequest) SetResourceGroupId(v string) *CreateCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateCertificateResponseBody struct {
	// example:
	//
	// 2329260
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateCertificateResponse) SetStatusCode(v int32) *CreateCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateResponse) SetBody(v *CreateCertificateResponseBody) *CreateCertificateResponse {
	s.Body = v
	return s
}

type CreateCertificateByCertificateIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3384669
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *CreateCertificateByCertificateIdRequest) SetRegionId(v string) *CreateCertificateByCertificateIdRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCertificateByCertificateIdRequest) SetResourceGroupId(v string) *CreateCertificateByCertificateIdRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateCertificateByCertificateIdResponseBody struct {
	// example:
	//
	// 3384669
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCertificateByCertificateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateCertificateByCertificateIdResponse) SetStatusCode(v int32) *CreateCertificateByCertificateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertificateByCertificateIdResponse) SetBody(v *CreateCertificateByCertificateIdResponseBody) *CreateCertificateByCertificateIdResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// example:
	//
	// 0
	AccessHeaderMode *int32 `json:"AccessHeaderMode,omitempty" xml:"AccessHeaderMode,omitempty"`
	// example:
	//
	// ["X-Client-IP"]
	AccessHeaders *string `json:"AccessHeaders,omitempty" xml:"AccessHeaders,omitempty"`
	// example:
	//
	// waf-cloud-dns
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// [{"ProtocolPortConfigs":[{"Ports":[80],"Protocol":"http"}],"RedirectionTypeName":"ALB","InstanceId":"alb-s65nua68wdedsp****","IPAddressList":["182.XX.XX.113"],"CloudNativeProductName":"ALB"}]
	CloudNativeInstances *string `json:"CloudNativeInstances,omitempty" xml:"CloudNativeInstances,omitempty"`
	// example:
	//
	// 0
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 5
	ConnectionTime *int32 `json:"ConnectionTime,omitempty" xml:"ConnectionTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// [443]
	Http2Port *string `json:"Http2Port,omitempty" xml:"Http2Port,omitempty"`
	// example:
	//
	// [80]
	HttpPort *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// example:
	//
	// 0
	HttpToUserIp *int32 `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	// example:
	//
	// [443]
	HttpsPort *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// example:
	//
	// 0
	HttpsRedirect *int32 `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp26f1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	IpFollowStatus *int32 `json:"IpFollowStatus,omitempty" xml:"IpFollowStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsAccessProduct   *int32 `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	Keepalive         *bool  `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	KeepaliveTimeout  *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// example:
	//
	// 0
	LoadBalancing *int32 `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	// example:
	//
	// [{"k":"ALIWAF-TAG","v":"Yes"}]
	LogHeaders *string `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty"`
	// example:
	//
	// 120
	ReadTime *int32  `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Retry           *bool   `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// example:
	//
	// waf.example.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// example:
	//
	// 1
	SniStatus *int32 `json:"SniStatus,omitempty" xml:"SniStatus,omitempty"`
	// example:
	//
	// ["39.XX.XX.197"]
	SourceIps *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	// example:
	//
	// 120
	WriteTime *int32 `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
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

func (s *CreateDomainRequest) SetKeepalive(v bool) *CreateDomainRequest {
	s.Keepalive = &v
	return s
}

func (s *CreateDomainRequest) SetKeepaliveRequests(v int32) *CreateDomainRequest {
	s.KeepaliveRequests = &v
	return s
}

func (s *CreateDomainRequest) SetKeepaliveTimeout(v int32) *CreateDomainRequest {
	s.KeepaliveTimeout = &v
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

func (s *CreateDomainRequest) SetRegionId(v string) *CreateDomainRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainRequest) SetResourceGroupId(v string) *CreateDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDomainRequest) SetRetry(v bool) *CreateDomainRequest {
	s.Retry = &v
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
	// example:
	//
	// mmspx7qhfvnfzggheh1g2wnbhog66vcv.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateProtectionModuleRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac_custom
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"action":"monitor","name":"test","scene":"custom_acl","conditions":[{"opCode":1,"key":"URL","values":"/example"}]}
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
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

func (s *CreateProtectionModuleRuleRequest) SetRegionId(v string) *CreateProtectionModuleRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetResourceGroupId(v string) *CreateProtectionModuleRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetRule(v string) *CreateProtectionModuleRuleRequest {
	s.Rule = &v
	return s
}

type CreateProtectionModuleRuleResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProtectionModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateProtectionModuleRuleResponse) SetStatusCode(v int32) *CreateProtectionModuleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtectionModuleRuleResponse) SetBody(v *CreateProtectionModuleRuleResponseBody) *CreateProtectionModuleRuleResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DeleteDomainRequest) SetRegionId(v string) *DeleteDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDomainRequest) SetResourceGroupId(v string) *DeleteDomainRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteDomainResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
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

func (s *DeleteInstanceRequest) SetRegionId(v string) *DeleteInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceGroupId(v string) *DeleteInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760E840
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteProtectionModuleRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac_custom
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 42754
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

func (s *DeleteProtectionModuleRuleRequest) SetRegionId(v string) *DeleteProtectionModuleRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteProtectionModuleRuleRequest) SetResourceGroupId(v string) *DeleteProtectionModuleRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteProtectionModuleRuleRequest) SetRuleId(v int64) *DeleteProtectionModuleRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteProtectionModuleRuleResponseBody struct {
	// example:
	//
	// 1557B42F-B889-460A-B17F-1DE5C5AD7FF2
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProtectionModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteProtectionModuleRuleResponse) SetStatusCode(v int32) *DeleteProtectionModuleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProtectionModuleRuleResponse) SetBody(v *DeleteProtectionModuleRuleResponseBody) *DeleteProtectionModuleRuleResponse {
	s.Body = v
	return s
}

type DescribeCertMatchStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- 62EcYPWd2Oy1vs6MTXcJSfN9Z7rZ9fmxWr2BFN2XbahgnsSXM48ixZJ4krc+1M+j2kcubVpsE2cgHdj4v8H6jUz9Ji4mr7vMNS6dXv8PUkl/qoDeNGCNdyTS5NIL5ir+g92cL8IGOkjgvhlqt9vc65Cgb4mL+n5+DV9uOyTZTW/MojmlgfUekC2xiXa54nxJf17Y1TADGSbyJbsC0Q9nIrHsPl8YKkvRWvIAqYxXZ7wRwWWmv4TMxFhWRiNY7yZIo2ZUhl02SIDNggIEeg== -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- DADTPZoOHd9WtZ3UKHJTRgNQmioPQn2bqdKHop+B/dn/4VZL7Jt8zSDGM9sTMThLyvsmLQKBgQCr+ujntC1kN6pGBj2Fw2l/EA/W3rYEce2tyhjgmG7rZ+A/jVE9fld5sQra6ZdwBcQJaiygoIYoaMF2EjRwc0qwHaluq0C15f6ujSoHh2e+D5zdmkTg/3NKNjqNv6xA2gYpinVDzFdZ9Zujxvuh9o4Vqf0YF8bv5UK5G04RtKadOw== -----END RSA PRIVATE KEY-----
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeCertMatchStatusRequest) SetRegionId(v string) *DescribeCertMatchStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCertMatchStatusRequest) SetResourceGroupId(v string) *DescribeCertMatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeCertMatchStatusResponseBody struct {
	// example:
	//
	// false
	MatchStatus *bool `json:"MatchStatus,omitempty" xml:"MatchStatus,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertMatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeCertMatchStatusResponse) SetStatusCode(v int32) *DescribeCertMatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertMatchStatusResponse) SetBody(v *DescribeCertMatchStatusResponseBody) *DescribeCertMatchStatusResponse {
	s.Body = v
	return s
}

type DescribeCertificatesRequest struct {
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeCertificatesRequest) SetRegionId(v string) *DescribeCertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCertificatesRequest) SetResourceGroupId(v string) *DescribeCertificatesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeCertificatesResponseBody struct {
	Certificates []*DescribeCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// example:
	//
	// ECF65091-3704-55D5-BC88-EC208B0E238C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2329260
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// *.aliyundoc.com
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// example:
	//
	// *.aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IsUsing *bool     `json:"IsUsing,omitempty" xml:"IsUsing,omitempty"`
	Sans    []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
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

func (s *DescribeCertificatesResponseBodyCertificates) SetEndTime(v int64) *DescribeCertificatesResponseBodyCertificates {
	s.EndTime = &v
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeCertificatesResponse) SetStatusCode(v int32) *DescribeCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificatesResponse) SetBody(v *DescribeCertificatesResponseBody) *DescribeCertificatesResponse {
	s.Body = v
	return s
}

type DescribeDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp26f1****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDomainRequest) SetRegionId(v string) *DescribeDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainRequest) SetResourceGroupId(v string) *DescribeDomainRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainResponseBody struct {
	Domain *DescribeDomainResponseBodyDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Struct"`
	// example:
	//
	// D827FCFE-90A7-4330-9326-D33C8B4C7726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	AccessHeaderMode *int32    `json:"AccessHeaderMode,omitempty" xml:"AccessHeaderMode,omitempty"`
	AccessHeaders    []*string `json:"AccessHeaders,omitempty" xml:"AccessHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// waf-cloud-dns
	AccessType           *string                                                 `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	CloudNativeInstances []*DescribeDomainResponseBodyDomainCloudNativeInstances `json:"CloudNativeInstances,omitempty" xml:"CloudNativeInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// kdmqyi3ck7xogegxpiyfpb0fj21mgkxn.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// 5
	ConnectionTime *int32   `json:"ConnectionTime,omitempty" xml:"ConnectionTime,omitempty"`
	Http2Port      []*int32 `json:"Http2Port,omitempty" xml:"Http2Port,omitempty" type:"Repeated"`
	HttpPort       []*int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	HttpToUserIp *int32   `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsPort    []*int32 `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	HttpsRedirect *int32 `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	// example:
	//
	// 1
	IpFollowStatus *int32 `json:"IpFollowStatus,omitempty" xml:"IpFollowStatus,omitempty"`
	// example:
	//
	// 1
	IsAccessProduct   *int32 `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	Keepalive         *bool  `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	KeepaliveTimeout  *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// example:
	//
	// 2
	LoadBalancing *int32                                        `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	LogHeaders    []*DescribeDomainResponseBodyDomainLogHeaders `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// 120
	ReadTime *int32 `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	// example:
	//
	// rg-acfm2mkrunv****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Retry           *bool   `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// example:
	//
	// waf.example.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// example:
	//
	// 1
	SniStatus *int32    `json:"SniStatus,omitempty" xml:"SniStatus,omitempty"`
	SourceIps []*string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty" type:"Repeated"`
	// example:
	//
	// 40
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 120
	WriteTime *int32 `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
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

func (s *DescribeDomainResponseBodyDomain) SetHttp2Port(v []*int32) *DescribeDomainResponseBodyDomain {
	s.Http2Port = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpPort(v []*int32) *DescribeDomainResponseBodyDomain {
	s.HttpPort = v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpToUserIp(v int32) *DescribeDomainResponseBodyDomain {
	s.HttpToUserIp = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetHttpsPort(v []*int32) *DescribeDomainResponseBodyDomain {
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

func (s *DescribeDomainResponseBodyDomain) SetKeepalive(v bool) *DescribeDomainResponseBodyDomain {
	s.Keepalive = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetKeepaliveRequests(v int32) *DescribeDomainResponseBodyDomain {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeDomainResponseBodyDomain) SetKeepaliveTimeout(v int32) *DescribeDomainResponseBodyDomain {
	s.KeepaliveTimeout = &v
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

func (s *DescribeDomainResponseBodyDomain) SetRetry(v bool) *DescribeDomainResponseBodyDomain {
	s.Retry = &v
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
	// example:
	//
	// ALB
	CloudNativeProductName *string `json:"CloudNativeProductName,omitempty" xml:"CloudNativeProductName,omitempty"`
	// example:
	//
	// ["39.XX.XX.197"]
	IPAddressList []*string `json:"IPAddressList,omitempty" xml:"IPAddressList,omitempty" type:"Repeated"`
	// example:
	//
	// alb-s65nua68wdedsp****
	InstanceId          *string                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProtocolPortConfigs []*DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs `json:"ProtocolPortConfigs,omitempty" xml:"ProtocolPortConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// ALB
	RedirectionTypeName *string `json:"RedirectionTypeName,omitempty" xml:"RedirectionTypeName,omitempty"`
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

func (s *DescribeDomainResponseBodyDomainCloudNativeInstances) SetIPAddressList(v []*string) *DescribeDomainResponseBodyDomainCloudNativeInstances {
	s.IPAddressList = v
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
	// example:
	//
	// [80]
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) SetPorts(v []*int32) *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs {
	s.Ports = v
	return s
}

func (s *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs) SetProtocol(v string) *DescribeDomainResponseBodyDomainCloudNativeInstancesProtocolPortConfigs {
	s.Protocol = &v
	return s
}

type DescribeDomainResponseBodyDomainLogHeaders struct {
	// example:
	//
	// ALIWAF-TAG
	K *string `json:"k,omitempty" xml:"k,omitempty"`
	// example:
	//
	// Yes
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainResponse) SetStatusCode(v int32) *DescribeDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResponse) SetBody(v *DescribeDomainResponseBody) *DescribeDomainResponse {
	s.Body = v
	return s
}

type DescribeDomainAdvanceConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-2r427ng****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
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

func (s *DescribeDomainAdvanceConfigsRequest) SetRegionId(v string) *DescribeDomainAdvanceConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsRequest) SetResourceGroupId(v string) *DescribeDomainAdvanceConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainAdvanceConfigsResponseBody struct {
	DomainConfigs []*DescribeDomainAdvanceConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// www.aliyundoc.com
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
	// example:
	//
	// 1
	CertStatus *int32 `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	// example:
	//
	// 0
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// ****dsbpkt75zeiok5mta2j5l7hggcrm.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// 0
	ExclusiveVipStatus *int32 `json:"ExclusiveVipStatus,omitempty" xml:"ExclusiveVipStatus,omitempty"`
	// example:
	//
	// on
	GSLBStatus *string  `json:"GSLBStatus,omitempty" xml:"GSLBStatus,omitempty"`
	Http2Port  []*int32 `json:"Http2Port,omitempty" xml:"Http2Port,omitempty" type:"Repeated"`
	HttpPort   []*int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty" type:"Repeated"`
	HttpsPort  []*int32 `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Ipv6Status *int32 `json:"Ipv6Status,omitempty" xml:"Ipv6Status,omitempty"`
	// example:
	//
	// 0
	ResolvedType *int32 `json:"ResolvedType,omitempty" xml:"ResolvedType,omitempty"`
	// example:
	//
	// ["39.XX.XX.197"]
	Rs []*string `json:"Rs,omitempty" xml:"Rs,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	VipServiceStatus *int32 `json:"VipServiceStatus,omitempty" xml:"VipServiceStatus,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAdvanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainAdvanceConfigsResponse) SetStatusCode(v int32) *DescribeDomainAdvanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAdvanceConfigsResponse) SetBody(v *DescribeDomainAdvanceConfigsResponseBody) *DescribeDomainAdvanceConfigsResponse {
	s.Body = v
	return s
}

type DescribeDomainBasicConfigsRequest struct {
	// example:
	//
	// waf-cloud-dns
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 0
	CloudNativeProductId *int32 `json:"CloudNativeProductId,omitempty" xml:"CloudNativeProductId,omitempty"`
	// example:
	//
	// aliyundoc
	DomainKey *string `json:"DomainKey,omitempty" xml:"DomainKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDomainBasicConfigsRequest) SetRegionId(v string) *DescribeDomainBasicConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainBasicConfigsRequest) SetResourceGroupId(v string) *DescribeDomainBasicConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainBasicConfigsResponseBody struct {
	DomainConfigs []*DescribeDomainBasicConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// waf-cloud-dns
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// 1
	AclStatus *int32 `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// example:
	//
	// 0
	CcMode *int32 `json:"CcMode,omitempty" xml:"CcMode,omitempty"`
	// example:
	//
	// 1
	CcStatus *int32 `json:"CcStatus,omitempty" xml:"CcStatus,omitempty"`
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// WAF
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 0
	WafMode *int32 `json:"WafMode,omitempty" xml:"WafMode,omitempty"`
	// example:
	//
	// 1
	WafStatus *int32 `json:"WafStatus,omitempty" xml:"WafStatus,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainBasicConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainBasicConfigsResponse) SetStatusCode(v int32) *DescribeDomainBasicConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBasicConfigsResponse) SetBody(v *DescribeDomainBasicConfigsResponseBody) *DescribeDomainBasicConfigsResponse {
	s.Body = v
	return s
}

type DescribeDomainListRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// example.com
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp26f1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	IsSub *int32 `json:"IsSub,omitempty" xml:"IsSub,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDomainListRequest) SetRegionId(v string) *DescribeDomainListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainListRequest) SetResourceGroupId(v string) *DescribeDomainListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainListResponseBody struct {
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// example:
	//
	// 592E866F-6C05-4E7C-81DE-B4D8E86B91EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainListResponse) SetStatusCode(v int32) *DescribeDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainListResponse) SetBody(v *DescribeDomainListResponseBody) *DescribeDomainListResponse {
	s.Body = v
	return s
}

type DescribeDomainNamesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
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

func (s *DescribeDomainNamesRequest) SetRegionId(v string) *DescribeDomainNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainNamesRequest) SetResourceGroupId(v string) *DescribeDomainNamesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainNamesResponseBody struct {
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeDomainRuleGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeDomainRuleGroupRequest) SetRegionId(v string) *DescribeDomainRuleGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainRuleGroupRequest) SetResourceGroupId(v string) *DescribeDomainRuleGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDomainRuleGroupResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1012
	RuleGroupId *int64 `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	// example:
	//
	// 1
	WafAiStatus *int32 `json:"WafAiStatus,omitempty" xml:"WafAiStatus,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainRuleGroupResponse) SetStatusCode(v int32) *DescribeDomainRuleGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRuleGroupResponse) SetBody(v *DescribeDomainRuleGroupResponseBody) *DescribeDomainRuleGroupResponse {
	s.Body = v
	return s
}

type DescribeInstanceInfoRequest struct {
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
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

func (s *DescribeInstanceInfoRequest) SetRegionId(v string) *DescribeInstanceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceInfoRequest) SetResourceGroupId(v string) *DescribeInstanceInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceInfoResponseBody struct {
	InstanceInfo *DescribeInstanceInfoResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1512921600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 1
	InDebt *int32 `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1
	RemainDay *int32 `json:"RemainDay,omitempty" xml:"RemainDay,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// example:
	//
	// 1
	Trial *int32 `json:"Trial,omitempty" xml:"Trial,omitempty"`
	// example:
	//
	// version_3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceInfoResponse) SetStatusCode(v int32) *DescribeInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceInfoResponse) SetBody(v *DescribeInstanceInfoResponseBody) *DescribeInstanceInfoResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecInfoRequest struct {
	// example:
	//
	// waf-cn-st2225l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
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

func (s *DescribeInstanceSpecInfoRequest) SetRegionId(v string) *DescribeInstanceSpecInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSpecInfoRequest) SetResourceGroupId(v string) *DescribeInstanceSpecInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceSpecInfoResponseBody struct {
	// example:
	//
	// 1677168000000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// waf-cn-st2225l****
	InstanceId        *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSpecInfos []*DescribeInstanceSpecInfoResponseBodyInstanceSpecInfos `json:"InstanceSpecInfos,omitempty" xml:"InstanceSpecInfos,omitempty" type:"Repeated"`
	// example:
	//
	// E906513E-F6B5-495E-98DC-7BA888671D76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	// example:
	//
	// 103
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceSpecInfoResponse) SetStatusCode(v int32) *DescribeInstanceSpecInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecInfoResponse) SetBody(v *DescribeInstanceSpecInfoResponseBody) *DescribeInstanceSpecInfoResponse {
	s.Body = v
	return s
}

type DescribeLogServiceStatusRequest struct {
	// example:
	//
	// www.aliyun.com
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// cn
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeLogServiceStatusRequest) SetRegionId(v string) *DescribeLogServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogServiceStatusRequest) SetResourceGroupId(v string) *DescribeLogServiceStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeLogServiceStatusResponseBody struct {
	DomainStatus []*DescribeLogServiceStatusResponseBodyDomainStatus `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty" type:"Repeated"`
	// example:
	//
	// C2E97B3F-1623-4CDF-A7E2-FD9D4CF1027A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1
	SlsLogActive *int32 `json:"SlsLogActive,omitempty" xml:"SlsLogActive,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeLogServiceStatusResponse) SetStatusCode(v int32) *DescribeLogServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogServiceStatusResponse) SetBody(v *DescribeLogServiceStatusResponseBody) *DescribeLogServiceStatusResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleCodeConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 14
	CodeType *int32 `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// example:
	//
	// 0
	CodeValue *int32 `json:"CodeValue,omitempty" xml:"CodeValue,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
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

func (s *DescribeProtectionModuleCodeConfigRequest) SetRegionId(v string) *DescribeProtectionModuleCodeConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProtectionModuleCodeConfigRequest) SetResourceGroupId(v string) *DescribeProtectionModuleCodeConfigRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeProtectionModuleCodeConfigResponseBody struct {
	// example:
	//
	// [{"code":0,"name":"310000,530000,150000,110000,TW_01,220000,510000,120000,640000,340000,370000,140000,440000,450000,650000,320000,360000,130000,410000,330000,460000,420000,430000,MO_01,620000,350000,540000,520000,210000,500000,610000,630000,HK_01,230000","env":"online"}]
	CodeConfigs *string `json:"CodeConfigs,omitempty" xml:"CodeConfigs,omitempty"`
	// example:
	//
	// BE3911B8-9D96-5B39-8875-503BBC9DA4BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProtectionModuleCodeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeProtectionModuleCodeConfigResponse) SetStatusCode(v int32) *DescribeProtectionModuleCodeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtectionModuleCodeConfigResponse) SetBody(v *DescribeProtectionModuleCodeConfigResponseBody) *DescribeProtectionModuleCodeConfigResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeProtectionModuleModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleModeRequest) SetDefenseType(v string) *DescribeProtectionModuleModeRequest {
	s.DefenseType = &v
	return s
}

func (s *DescribeProtectionModuleModeRequest) SetDomain(v string) *DescribeProtectionModuleModeRequest {
	s.Domain = &v
	return s
}

func (s *DescribeProtectionModuleModeRequest) SetInstanceId(v string) *DescribeProtectionModuleModeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProtectionModuleModeRequest) SetRegionId(v string) *DescribeProtectionModuleModeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProtectionModuleModeRequest) SetResourceGroupId(v string) *DescribeProtectionModuleModeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeProtectionModuleModeResponseBody struct {
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtectionModuleModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleModeResponseBody) SetMode(v int32) *DescribeProtectionModuleModeResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeProtectionModuleModeResponseBody) SetRequestId(v string) *DescribeProtectionModuleModeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProtectionModuleModeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProtectionModuleModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProtectionModuleModeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleModeResponse) SetHeaders(v map[string]*string) *DescribeProtectionModuleModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeProtectionModuleModeResponse) SetStatusCode(v int32) *DescribeProtectionModuleModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtectionModuleModeResponse) SetBody(v *DescribeProtectionModuleModeResponseBody) *DescribeProtectionModuleModeResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac_highfreq
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// e2ZpbHRlcjp7InJ1bGVJZCI6NDI3NTV9LG9yZGVyQnk6ImdtdF9tb2RpZmllZCIsZGVzYzp0cnVlfQ==
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
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

func (s *DescribeProtectionModuleRulesRequest) SetRegionId(v string) *DescribeProtectionModuleRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetResourceGroupId(v string) *DescribeProtectionModuleRulesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeProtectionModuleRulesResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeProtectionModuleRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// {"count":60,"interval":60,"ttl":300}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 42755
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1570700044
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 2
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProtectionModuleRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeProtectionModuleRulesResponse) SetStatusCode(v int32) *DescribeProtectionModuleRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponse) SetBody(v *DescribeProtectionModuleRulesResponseBody) *DescribeProtectionModuleRulesResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeProtectionModuleStatusRequest) SetRegionId(v string) *DescribeProtectionModuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProtectionModuleStatusRequest) SetResourceGroupId(v string) *DescribeProtectionModuleStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeProtectionModuleStatusResponseBody struct {
	// example:
	//
	// 1
	ModuleStatus *int32 `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProtectionModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeProtectionModuleStatusResponse) SetStatusCode(v int32) *DescribeProtectionModuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProtectionModuleStatusResponse) SetBody(v *DescribeProtectionModuleStatusResponseBody) *DescribeProtectionModuleStatusResponse {
	s.Body = v
	return s
}

type DescribeRuleGroupsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1011
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek23puu7m3kmea
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 60.208.111.213
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 10
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// ZH
	WafLang *string `json:"WafLang,omitempty" xml:"WafLang,omitempty"`
}

func (s DescribeRuleGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsRequest) SetCurrentPage(v int32) *DescribeRuleGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetInstanceId(v string) *DescribeRuleGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetLang(v string) *DescribeRuleGroupsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageSize(v int32) *DescribeRuleGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPolicyId(v int64) *DescribeRuleGroupsRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetRegion(v string) *DescribeRuleGroupsRequest {
	s.Region = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetRegionId(v string) *DescribeRuleGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetResourceGroupId(v string) *DescribeRuleGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSourceIp(v string) *DescribeRuleGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetType(v int32) *DescribeRuleGroupsRequest {
	s.Type = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetWafLang(v string) *DescribeRuleGroupsRequest {
	s.WafLang = &v
	return s
}

type DescribeRuleGroupsResponseBody struct {
	// example:
	//
	// 02E9A4B8-90FB-5F41-A049-C82277EB82FB
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleGroups []*DescribeRuleGroupsResponseBodyRuleGroups `json:"RuleGroups,omitempty" xml:"RuleGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 123
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s DescribeRuleGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBody) SetRequestId(v string) *DescribeRuleGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetRuleGroups(v []*DescribeRuleGroupsResponseBodyRuleGroups) *DescribeRuleGroupsResponseBody {
	s.RuleGroups = v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetTaskStatus(v int32) *DescribeRuleGroupsResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetTotal(v int32) *DescribeRuleGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetWafTaskId(v string) *DescribeRuleGroupsResponseBody {
	s.WafTaskId = &v
	return s
}

type DescribeRuleGroupsResponseBodyRuleGroups struct {
	// example:
	//
	// desc
	Desc       *string   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// rule_group_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 116562
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	RuleCnt *int32 `json:"RuleCnt,omitempty" xml:"RuleCnt,omitempty"`
	// example:
	//
	// rule_group_test
	RuleGroupTemplateName *string `json:"RuleGroupTemplateName,omitempty" xml:"RuleGroupTemplateName,omitempty"`
	// example:
	//
	// 1711445265
	RuleGroupUpdateTime *int64 `json:"RuleGroupUpdateTime,omitempty" xml:"RuleGroupUpdateTime,omitempty"`
	// example:
	//
	// 1102
	TemplatePolicyId *int64 `json:"TemplatePolicyId,omitempty" xml:"TemplatePolicyId,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 11
	WafVersion *int64 `json:"WafVersion,omitempty" xml:"WafVersion,omitempty"`
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetDesc(v string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.Desc = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetDomainList(v []*string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.DomainList = v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetName(v string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.Name = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetPolicyId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.PolicyId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleCnt(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleCnt = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupTemplateName(v string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupTemplateName = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupUpdateTime(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupUpdateTime = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetTemplatePolicyId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.TemplatePolicyId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetType(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.Type = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetWafVersion(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.WafVersion = &v
	return s
}

type DescribeRuleGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponse) SetHeaders(v map[string]*string) *DescribeRuleGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleGroupsResponse) SetStatusCode(v int32) *DescribeRuleGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleGroupsResponse) SetBody(v *DescribeRuleGroupsResponseBody) *DescribeRuleGroupsResponse {
	s.Body = v
	return s
}

type DescribeRulesRequest struct {
	// example:
	//
	// 0
	ApplicationType *int32 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// CVE ID
	//
	// example:
	//
	// CVE-*-*5
	CveIdKey *string `json:"CveIdKey,omitempty" xml:"CveIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// 11
	ProtectionType *int32 `json:"ProtectionType,omitempty" xml:"ProtectionType,omitempty"`
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-*
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 0
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 1012
	RuleGroupId *int64 `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	// example:
	//
	// 1131*0
	RuleIdKey *string `json:"RuleIdKey,omitempty" xml:"RuleIdKey,omitempty"`
	// example:
	//
	// 42.84.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) SetApplicationType(v int32) *DescribeRulesRequest {
	s.ApplicationType = &v
	return s
}

func (s *DescribeRulesRequest) SetCveIdKey(v string) *DescribeRulesRequest {
	s.CveIdKey = &v
	return s
}

func (s *DescribeRulesRequest) SetInstanceId(v string) *DescribeRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRulesRequest) SetLang(v string) *DescribeRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRulesRequest) SetPageNumber(v int32) *DescribeRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRulesRequest) SetPageSize(v int32) *DescribeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRulesRequest) SetProtectionType(v int32) *DescribeRulesRequest {
	s.ProtectionType = &v
	return s
}

func (s *DescribeRulesRequest) SetRegion(v string) *DescribeRulesRequest {
	s.Region = &v
	return s
}

func (s *DescribeRulesRequest) SetRegionId(v string) *DescribeRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceGroupId(v string) *DescribeRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRulesRequest) SetRiskLevel(v int32) *DescribeRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRulesRequest) SetRuleGroupId(v int64) *DescribeRulesRequest {
	s.RuleGroupId = &v
	return s
}

func (s *DescribeRulesRequest) SetRuleIdKey(v string) *DescribeRulesRequest {
	s.RuleIdKey = &v
	return s
}

func (s *DescribeRulesRequest) SetSourceIp(v string) *DescribeRulesRequest {
	s.SourceIp = &v
	return s
}

type DescribeRulesResponseBody struct {
	// example:
	//
	// 1
	IsSubscribe *int64 `json:"IsSubscribe,omitempty" xml:"IsSubscribe,omitempty"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-*
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test111
	RuleGroupName *string `json:"RuleGroupName,omitempty" xml:"RuleGroupName,omitempty"`
	// example:
	//
	// 1012
	RuleGroupTemplateId *string `json:"RuleGroupTemplateId,omitempty" xml:"RuleGroupTemplateId,omitempty"`
	// example:
	//
	// rule_group_test
	RuleGroupTemplateName *string                           `json:"RuleGroupTemplateName,omitempty" xml:"RuleGroupTemplateName,omitempty"`
	Rules                 []*DescribeRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) SetIsSubscribe(v int64) *DescribeRulesResponseBody {
	s.IsSubscribe = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRuleGroupName(v string) *DescribeRulesResponseBody {
	s.RuleGroupName = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRuleGroupTemplateId(v string) *DescribeRulesResponseBody {
	s.RuleGroupTemplateId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRuleGroupTemplateName(v string) *DescribeRulesResponseBody {
	s.RuleGroupTemplateName = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRules(v []*DescribeRulesResponseBodyRules) *DescribeRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeRulesResponseBody) SetTotalCount(v int32) *DescribeRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRulesResponseBodyRules struct {
	// example:
	//
	// 15
	ApplicationType *int32 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// CVE ID。
	//
	// example:
	//
	// CVE-2021-*
	CveId       *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	CveUrl      *string `json:"CveUrl,omitempty" xml:"CveUrl,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 11
	ProtectionType *int32 `json:"ProtectionType,omitempty" xml:"ProtectionType,omitempty"`
	// example:
	//
	// 0
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 42755
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// rules_41
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1684120148.0
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyRules) SetApplicationType(v int32) *DescribeRulesResponseBodyRules {
	s.ApplicationType = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetCveId(v string) *DescribeRulesResponseBodyRules {
	s.CveId = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetCveUrl(v string) *DescribeRulesResponseBodyRules {
	s.CveUrl = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetDescription(v string) *DescribeRulesResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetProtectionType(v int32) *DescribeRulesResponseBodyRules {
	s.ProtectionType = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetRiskLevel(v int32) *DescribeRulesResponseBodyRules {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetRuleId(v int64) *DescribeRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetRuleName(v string) *DescribeRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeRulesResponseBodyRules) SetUpdateTime(v int64) *DescribeRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

type DescribeRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponse) SetHeaders(v map[string]*string) *DescribeRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulesResponse) SetStatusCode(v int32) *DescribeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulesResponse) SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse {
	s.Body = v
	return s
}

type DescribeWafSourceIpSegmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
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

func (s *DescribeWafSourceIpSegmentRequest) SetRegionId(v string) *DescribeWafSourceIpSegmentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetResourceGroupId(v string) *DescribeWafSourceIpSegmentRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeWafSourceIpSegmentResponseBody struct {
	// example:
	//
	// 39.XXX.XXX.0/24,……,2408:400a:XXXX:XXXX::/56
	IpV6s *string `json:"IpV6s,omitempty" xml:"IpV6s,omitempty"`
	// example:
	//
	// 47.XXX.XXX.192/26,……,47.XXX.XXX.0/24
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// example:
	//
	// AB2F5E31-EE96-4FD7-9560-45FF5D5377FF
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWafSourceIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeWafSourceIpSegmentResponse) SetStatusCode(v int32) *DescribeWafSourceIpSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) SetBody(v *DescribeWafSourceIpSegmentResponseBody) *DescribeWafSourceIpSegmentResponse {
	s.Body = v
	return s
}

type ModifyDomainRequest struct {
	// The method that WAF uses to obtain the actual IP address of a client. Valid values:
	//
	// 	- **0**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the actual IP address of the client. This is the default value.
	//
	// 	- **1**: WAF reads the value of a custom header field as the actual IP address of the client.
	//
	// >  You need to specify the parameter only when the **IsAccessProduct*	- parameter is set to **1**.
	//
	// example:
	//
	// 0
	AccessHeaderMode *int32 `json:"AccessHeaderMode,omitempty" xml:"AccessHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the actual IP address of a client. Specify the value in the `["header1","header2",...]` format.
	//
	// >  You need to specify the parameter only when the **AccessHeaderMode*	- parameter is set to **1**.
	//
	// example:
	//
	// ["X-Client-IP"]
	AccessHeaders *string `json:"AccessHeaders,omitempty" xml:"AccessHeaders,omitempty"`
	// The mode that is used to add the domain name. Valid values:
	//
	// 	- **waf-cloud-dns**: CNAME record mode. This is the default value.
	//
	// 	- **waf-cloud-native**: transparent proxy mode.
	//
	// example:
	//
	// waf-cloud-dns
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The list of server and port configurations for the transparent proxy mode. The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **ProtocolPortConfigs**: the list of protocol and port configurations. This field is required. Data type: array. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	//     	- **Ports**: the list of ports. This field is required. Data type: array. The value is in the `[port1,port2,……]` format.
	//
	//     	- **Protocol**: the protocol. This field is required. Data type: string. Valid values: **http*	- and **https**.
	//
	// 	- **CloudNativeProductName**: the type of the cloud service instance. This field is required. Data type: string. Valid values: **ECS**, **SLB**, and **ALB**.
	//
	// 	- **RedirectionTypeName**: the type of traffic redirection port. This field is required. Data type: string. Valid values: **ECS**, **SLB-L4**, **SLB-L7**, and **ALB**.
	//
	// 	- **InstanceId**: the ID of the cloud service instance. This field is required. Data type: string.
	//
	// 	- **IPAddressList**: the list of public IP addresses of the cloud service instance. This field is required. Data type: array. The value is in the `["ip1","ip2",...]` format.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-native**.
	//
	// example:
	//
	// [{"ProtocolPortConfigs":[{"Ports":[80],"Protocol":"http"}],"RedirectionTypeName":"ALB","InstanceId":"alb-s65nua68wdedsp****","IPAddressList":["182.XX.XX.113"],"CloudNativeProductName":"ALB"}]
	CloudNativeInstances *string `json:"CloudNativeInstances,omitempty" xml:"CloudNativeInstances,omitempty"`
	// The type of WAF protection cluster. Valid values:
	//
	// 	- **0**: shared cluster. This is the default value.
	//
	// 	- **1**: exclusive cluster.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns**.
	//
	// example:
	//
	// 0
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The timeout period for connections of WAF exclusive clusters. Unit: seconds.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the value of the **ClusterType*	- parameter is set to **1**.
	//
	// example:
	//
	// 5
	ConnectionTime *int32 `json:"ConnectionTime,omitempty" xml:"ConnectionTime,omitempty"`
	// The domain name whose configurations you want to modify.
	//
	// >  You can call the [DescribeDomainNames](https://help.aliyun.com/document_detail/86373.html) operation to query the domain names that are added to Web Application Firewall (WAF).
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The HTTP/2 ports. Specify the value in the `["port1","port2",...]` format.
	//
	// >  You need to specify this parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the **HttpsPort*	- parameter is not empty. If the HttpsPort parameter is not empty, your website uses HTTPS.
	//
	// example:
	//
	// [443]
	Http2Port *string `json:"Http2Port,omitempty" xml:"Http2Port,omitempty"`
	// The HTTP ports. Specify the value in the `["port1","port2",...]` format.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns**. If you specify this parameter, your website uses HTTP. You must specify at least one of the **HttpPort*	- and **HttpsPort*	- parameters.
	//
	// example:
	//
	// [80]
	HttpPort *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// Specifies whether to enable the feature of redirecting HTTPS requests to HTTP requests. If you enable the feature, HTTPS requests are redirected to HTTP requests on port 80, which is used by default. Valid values:
	//
	// 	- **0**: disables the feature. This is the default value.
	//
	// 	- **1**: enables the feature.
	//
	// >  You need to specify this parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the **HttpsPort*	- parameter is not empty. If the HttpsPort parameter is not empty, your website uses HTTPS.
	//
	// example:
	//
	// 0
	HttpToUserIp *int32 `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	// The HTTPS ports. Specify the value in the `["port1","port2",...]` format.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns**. If you specify this parameter, your website uses HTTPS. You must specify at least one of the **HttpPort*	- and **HttpsPort*	- parameters.
	//
	// example:
	//
	// [443]
	HttpsPort *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// Specifies whether to enable the feature of redirecting HTTP requests to HTTPS requests. If you enable the feature, HTTP requests are redirected to HTTPS requests on port 443, which is used by default. Valid values:
	//
	// 	- **0**: disables the feature. This is the default value.
	//
	// 	- **1**: enables the feature.
	//
	// >  You need to specify this parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the **HttpsPort*	- parameter is not empty. If the HttpsPort parameter is not empty, your website uses HTTPS.
	//
	// example:
	//
	// 0
	HttpsRedirect *int32 `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp26f1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable the feature of forwarding requests to the origin servers that use the IP address type specified in the requests. If you enable the feature, WAF forwards requests from IPv4 addresses to origin servers that use IPv4 addresses and requests from IPv6 addresses to origin servers that use IPv6 addresses. Valid values:
	//
	// 	- **0**: disables the feature. This is the default value.
	//
	// 	- **1**: enables the feature.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns**.
	//
	// example:
	//
	// 0
	IpFollowStatus *int32 `json:"IpFollowStatus,omitempty" xml:"IpFollowStatus,omitempty"`
	// Specifies whether to deploy a Layer 7 proxy, which is used to filter inbound traffic before the traffic reaches the WAF instance. The supported Layer 7 proxies include Anti-DDoS Pro, Anti-DDoS Premium, and Alibaba Cloud CDN. Valid values:
	//
	// 	- **0**: does not configure a Layer 7 proxy
	//
	// 	- **1**: configures a Layer 7 proxy
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsAccessProduct *int32 `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// example:
	//
	// 60
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that is used when WAF forwards requests to the origin server. Valid values:
	//
	// 	- **0**: IP hash
	//
	// 	- **1**: round-robin
	//
	// 	- **2**: least time
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns**.
	//
	// example:
	//
	// 0
	LoadBalancing *int32 `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	// The key-value pair that is used to mark the requests that pass through the WAF instance.
	//
	// Specify the key-value pair in the `[{"k":"_key_","v":"_value_"}]` format. `_key_` specifies a header field in a custom request. `_value_` specifies the value of the field.
	//
	// WAF automatically adds the key-value pair to the headers of requests. This way, the requests that pass through WAF are identified.
	//
	// >  If requests contain the custom header field, WAF overwrites the original value of the field with the specified value.
	//
	// example:
	//
	// [{"k":"ALIWAF-TAG","v":"Yes"}]
	LogHeaders *string `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty"`
	// The timeout period for read connections of WAF exclusive clusters. Unit: seconds.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the value of the **ClusterType*	- parameter is set to **1**.
	//
	// example:
	//
	// 120
	ReadTime *int32 `json:"ReadTime,omitempty" xml:"ReadTime,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The value of the custom SNI field. If this parameter is not specified, the value of the **Host*	- field in the request header is automatically used as the value of the SNI field.
	//
	// If you want WAF to use an SNI field whose value is different from the value of the Host field, you can specify a custom value for the SNI field.
	//
	// >  This parameter needs to be set only when the value of the **SniStatus*	- parameter is set to **1**.
	//
	// example:
	//
	// waf.example.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Specifies whether to enable origin SNI. Origin Server Name Indication (SNI) specifies the domain name to which an HTTPS connection needs to be established at the start of the TLS handshaking process when WAF forwards requests to the origin server. If the origin server hosts multiple domain names, you must enable this feature. Valid values:
	//
	// 	- **0**: disables origin SNI.
	//
	// 	- **1**: enables origin SNI.
	//
	// By default, origin SNI is disabled for WAF instances in the Chinese mainland and enabled for WAF instances outside the Chinese mainland.
	//
	// >  You need to specify this parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the **HttpsPort*	- parameter is not empty. If the HttpsPort parameter is not empty, your website uses HTTPS.
	//
	// example:
	//
	// 1
	SniStatus *int32 `json:"SniStatus,omitempty" xml:"SniStatus,omitempty"`
	// The address type of the origin server. The address can be an IP address or a domain name. You can specify only one type of address.
	//
	// 	- If you use the IP address type, specify the value in the `["ip1","ip2",...]` format. You can add up to 20 IP addresses.
	//
	// 	- If you use the domain name type, specify the value in the `["domain"]` format. You can enter only one domain name.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns**.
	//
	// example:
	//
	// ["39.XX.XX.197"]
	SourceIps *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	// The timeout period for write connections of WAF exclusive clusters. Unit: seconds.
	//
	// >  You need to specify the parameter only when the value of the **AccessType*	- parameter is set to **waf-cloud-dns*	- and the value of the **ClusterType*	- parameter is set to **1**.
	//
	// example:
	//
	// 120
	WriteTime *int32 `json:"WriteTime,omitempty" xml:"WriteTime,omitempty"`
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

func (s *ModifyDomainRequest) SetKeepalive(v bool) *ModifyDomainRequest {
	s.Keepalive = &v
	return s
}

func (s *ModifyDomainRequest) SetKeepaliveRequests(v int32) *ModifyDomainRequest {
	s.KeepaliveRequests = &v
	return s
}

func (s *ModifyDomainRequest) SetKeepaliveTimeout(v int32) *ModifyDomainRequest {
	s.KeepaliveTimeout = &v
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

func (s *ModifyDomainRequest) SetRegionId(v string) *ModifyDomainRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainRequest) SetResourceGroupId(v string) *ModifyDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDomainRequest) SetRetry(v bool) *ModifyDomainRequest {
	s.Retry = &v
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
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDomainResponse) SetStatusCode(v int32) *ModifyDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainResponse) SetBody(v *ModifyDomainResponseBody) *ModifyDomainResponse {
	s.Body = v
	return s
}

type ModifyDomainIpv6StatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ModifyDomainIpv6StatusRequest) SetRegionId(v string) *ModifyDomainIpv6StatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainIpv6StatusRequest) SetResourceGroupId(v string) *ModifyDomainIpv6StatusRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyDomainIpv6StatusResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainIpv6StatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyDomainIpv6StatusResponse) SetStatusCode(v int32) *ModifyDomainIpv6StatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainIpv6StatusResponse) SetBody(v *ModifyDomainIpv6StatusResponseBody) *ModifyDomainIpv6StatusResponse {
	s.Body = v
	return s
}

type ModifyLogRetrievalStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ModifyLogRetrievalStatusRequest) SetRegionId(v string) *ModifyLogRetrievalStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLogRetrievalStatusRequest) SetResourceGroupId(v string) *ModifyLogRetrievalStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyLogRetrievalStatusResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogRetrievalStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyLogRetrievalStatusResponse) SetStatusCode(v int32) *ModifyLogRetrievalStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogRetrievalStatusResponse) SetBody(v *ModifyLogRetrievalStatusResponseBody) *ModifyLogRetrievalStatusResponse {
	s.Body = v
	return s
}

type ModifyLogServiceStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ModifyLogServiceStatusRequest) SetRegionId(v string) *ModifyLogServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLogServiceStatusRequest) SetResourceGroupId(v string) *ModifyLogServiceStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyLogServiceStatusResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyLogServiceStatusResponse) SetStatusCode(v int32) *ModifyLogServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogServiceStatusResponse) SetBody(v *ModifyLogServiceStatusResponseBody) *ModifyLogServiceStatusResponse {
	s.Body = v
	return s
}

type ModifyProtectionModuleModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Mode            *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ModifyProtectionModuleModeRequest) SetRegionId(v string) *ModifyProtectionModuleModeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProtectionModuleModeRequest) SetResourceGroupId(v string) *ModifyProtectionModuleModeRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyProtectionModuleModeResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtectionModuleModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyProtectionModuleModeResponse) SetStatusCode(v int32) *ModifyProtectionModuleModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtectionModuleModeResponse) SetBody(v *ModifyProtectionModuleModeResponseBody) *ModifyProtectionModuleModeResponse {
	s.Body = v
	return s
}

type ModifyProtectionModuleRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac_custom
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LockVersion     *int64  `json:"LockVersion,omitempty" xml:"LockVersion,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"action":"monitor","name":"test","scene":"custom_acl","conditions":[{"opCode":1,"key":"URL","values":"/example"}]}
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 369998
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

func (s *ModifyProtectionModuleRuleRequest) SetRegionId(v string) *ModifyProtectionModuleRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProtectionModuleRuleRequest) SetResourceGroupId(v string) *ModifyProtectionModuleRuleRequest {
	s.ResourceGroupId = &v
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
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtectionModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyProtectionModuleRuleResponse) SetStatusCode(v int32) *ModifyProtectionModuleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtectionModuleRuleResponse) SetBody(v *ModifyProtectionModuleRuleResponseBody) *ModifyProtectionModuleRuleResponse {
	s.Body = v
	return s
}

type ModifyProtectionModuleStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ModuleStatus    *int32  `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ModifyProtectionModuleStatusRequest) SetRegionId(v string) *ModifyProtectionModuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProtectionModuleStatusRequest) SetResourceGroupId(v string) *ModifyProtectionModuleStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyProtectionModuleStatusResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtectionModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyProtectionModuleStatusResponse) SetStatusCode(v int32) *ModifyProtectionModuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtectionModuleStatusResponse) SetBody(v *ModifyProtectionModuleStatusResponseBody) *ModifyProtectionModuleStatusResponse {
	s.Body = v
	return s
}

type ModifyProtectionRuleCacheStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tamperproof
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 42755
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

func (s *ModifyProtectionRuleCacheStatusRequest) SetRegionId(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetResourceGroupId(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetRuleId(v int64) *ModifyProtectionRuleCacheStatusRequest {
	s.RuleId = &v
	return s
}

type ModifyProtectionRuleCacheStatusResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtectionRuleCacheStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyProtectionRuleCacheStatusResponse) SetStatusCode(v int32) *ModifyProtectionRuleCacheStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusResponse) SetBody(v *ModifyProtectionRuleCacheStatusResponseBody) *ModifyProtectionRuleCacheStatusResponse {
	s.Body = v
	return s
}

type ModifyProtectionRuleStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tamperproof
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LockVersion     *int64  `json:"LockVersion,omitempty" xml:"LockVersion,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 42755
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
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

func (s *ModifyProtectionRuleStatusRequest) SetRegionId(v string) *ModifyProtectionRuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetResourceGroupId(v string) *ModifyProtectionRuleStatusRequest {
	s.ResourceGroupId = &v
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
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtectionRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyProtectionRuleStatusResponse) SetStatusCode(v int32) *ModifyProtectionRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtectionRuleStatusResponse) SetBody(v *ModifyProtectionRuleStatusResponseBody) *ModifyProtectionRuleStatusResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-09k1rd5****~www.example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// domain
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
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
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type SetDomainRuleGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["www.aliyundoc.com"]
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1012
	RuleGroupId *int64 `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	// example:
	//
	// 1
	WafAiStatus *int32 `json:"WafAiStatus,omitempty" xml:"WafAiStatus,omitempty"`
	// example:
	//
	// 1
	WafVersion *int64 `json:"WafVersion,omitempty" xml:"WafVersion,omitempty"`
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

func (s *SetDomainRuleGroupRequest) SetRegionId(v string) *SetDomainRuleGroupRequest {
	s.RegionId = &v
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
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SetDomainRuleGroupResponse) SetStatusCode(v int32) *SetDomainRuleGroupResponse {
	s.StatusCode = &v
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

// @param request - CreateCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - CreateCertificateRequest
//
// @return CreateCertificateResponse
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

// @param request - CreateCertificateByCertificateIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertificateByCertificateIdResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - CreateCertificateByCertificateIdRequest
//
// @return CreateCertificateByCertificateIdResponse
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

// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
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

	if !tea.BoolValue(util.IsUnset(request.Keepalive)) {
		query["Keepalive"] = request.Keepalive
	}

	if !tea.BoolValue(util.IsUnset(request.KeepaliveRequests)) {
		query["KeepaliveRequests"] = request.KeepaliveRequests
	}

	if !tea.BoolValue(util.IsUnset(request.KeepaliveTimeout)) {
		query["KeepaliveTimeout"] = request.KeepaliveTimeout
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Retry)) {
		query["Retry"] = request.Retry
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

// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
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

// @param request - CreateProtectionModuleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtectionModuleRuleResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - CreateProtectionModuleRuleRequest
//
// @return CreateProtectionModuleRuleResponse
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

// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
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

// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
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

// @param request - DeleteProtectionModuleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtectionModuleRuleResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DeleteProtectionModuleRuleRequest
//
// @return DeleteProtectionModuleRuleResponse
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

// @param request - DescribeCertMatchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertMatchStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DescribeCertMatchStatusRequest
//
// @return DescribeCertMatchStatusResponse
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

// @param request - DescribeCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificatesResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DescribeCertificatesRequest
//
// @return DescribeCertificatesResponse
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

// @param request - DescribeDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DescribeDomainRequest
//
// @return DescribeDomainResponse
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

// @param request - DescribeDomainAdvanceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainAdvanceConfigsResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeDomainAdvanceConfigsRequest
//
// @return DescribeDomainAdvanceConfigsResponse
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

// @param request - DescribeDomainBasicConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainBasicConfigsResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeDomainBasicConfigsRequest
//
// @return DescribeDomainBasicConfigsResponse
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

// @param request - DescribeDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainListResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeDomainListRequest
//
// @return DescribeDomainListResponse
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

// @param request - DescribeDomainNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainNamesResponse
func (client *Client) DescribeDomainNamesWithOptions(request *DescribeDomainNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeDomainNamesRequest
//
// @return DescribeDomainNamesResponse
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

// @param request - DescribeDomainRuleGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRuleGroupResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DescribeDomainRuleGroupRequest
//
// @return DescribeDomainRuleGroupResponse
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

// @param request - DescribeInstanceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceInfoResponse
func (client *Client) DescribeInstanceInfoWithOptions(request *DescribeInstanceInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeInstanceInfoRequest
//
// @return DescribeInstanceInfoResponse
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

// @param request - DescribeInstanceSpecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSpecInfoResponse
func (client *Client) DescribeInstanceSpecInfoWithOptions(request *DescribeInstanceSpecInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeInstanceSpecInfoRequest
//
// @return DescribeInstanceSpecInfoResponse
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

// @param request - DescribeLogServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogServiceStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeLogServiceStatusRequest
//
// @return DescribeLogServiceStatusResponse
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

// @param request - DescribeProtectionModuleCodeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtectionModuleCodeConfigResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeProtectionModuleCodeConfigRequest
//
// @return DescribeProtectionModuleCodeConfigResponse
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

// @param request - DescribeProtectionModuleModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtectionModuleModeResponse
func (client *Client) DescribeProtectionModuleModeWithOptions(request *DescribeProtectionModuleModeRequest, runtime *util.RuntimeOptions) (_result *DescribeProtectionModuleModeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProtectionModuleMode"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProtectionModuleModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeProtectionModuleModeRequest
//
// @return DescribeProtectionModuleModeResponse
func (client *Client) DescribeProtectionModuleMode(request *DescribeProtectionModuleModeRequest) (_result *DescribeProtectionModuleModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProtectionModuleModeResponse{}
	_body, _err := client.DescribeProtectionModuleModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeProtectionModuleRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtectionModuleRulesResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeProtectionModuleRulesRequest
//
// @return DescribeProtectionModuleRulesResponse
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

// @param request - DescribeProtectionModuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtectionModuleStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - DescribeProtectionModuleStatusRequest
//
// @return DescribeProtectionModuleStatusResponse
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

// @param request - DescribeRuleGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleGroupsResponse
func (client *Client) DescribeRuleGroupsWithOptions(request *DescribeRuleGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WafLang)) {
		query["WafLang"] = request.WafLang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleGroups"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRuleGroupsRequest
//
// @return DescribeRuleGroupsResponse
func (client *Client) DescribeRuleGroups(request *DescribeRuleGroupsRequest) (_result *DescribeRuleGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleGroupsResponse{}
	_body, _err := client.DescribeRuleGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulesResponse
func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.CveIdKey)) {
		query["CveIdKey"] = request.CveIdKey
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

	if !tea.BoolValue(util.IsUnset(request.ProtectionType)) {
		query["ProtectionType"] = request.ProtectionType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.RuleGroupId)) {
		query["RuleGroupId"] = request.RuleGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIdKey)) {
		query["RuleIdKey"] = request.RuleIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRules"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRulesRequest
//
// @return DescribeRulesResponse
func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeWafSourceIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWafSourceIpSegmentResponse
func (client *Client) DescribeWafSourceIpSegmentWithOptions(request *DescribeWafSourceIpSegmentRequest, runtime *util.RuntimeOptions) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - DescribeWafSourceIpSegmentRequest
//
// @return DescribeWafSourceIpSegmentResponse
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

// Summary:
//
// Modifies the configurations of a domain name.
//
// @param request - ModifyDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDomainResponse
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

	if !tea.BoolValue(util.IsUnset(request.Keepalive)) {
		query["Keepalive"] = request.Keepalive
	}

	if !tea.BoolValue(util.IsUnset(request.KeepaliveRequests)) {
		query["KeepaliveRequests"] = request.KeepaliveRequests
	}

	if !tea.BoolValue(util.IsUnset(request.KeepaliveTimeout)) {
		query["KeepaliveTimeout"] = request.KeepaliveTimeout
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Retry)) {
		query["Retry"] = request.Retry
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

// Summary:
//
// Modifies the configurations of a domain name.
//
// @param request - ModifyDomainRequest
//
// @return ModifyDomainResponse
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

// @param request - ModifyDomainIpv6StatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDomainIpv6StatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyDomainIpv6StatusRequest
//
// @return ModifyDomainIpv6StatusResponse
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

// @param request - ModifyLogRetrievalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLogRetrievalStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyLogRetrievalStatusRequest
//
// @return ModifyLogRetrievalStatusResponse
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

// @param request - ModifyLogServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLogServiceStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyLogServiceStatusRequest
//
// @return ModifyLogServiceStatusResponse
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

// @param request - ModifyProtectionModuleModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtectionModuleModeResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyProtectionModuleModeRequest
//
// @return ModifyProtectionModuleModeResponse
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

// @param request - ModifyProtectionModuleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtectionModuleRuleResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyProtectionModuleRuleRequest
//
// @return ModifyProtectionModuleRuleResponse
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

// @param request - ModifyProtectionModuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtectionModuleStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyProtectionModuleStatusRequest
//
// @return ModifyProtectionModuleStatusResponse
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

// @param request - ModifyProtectionRuleCacheStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtectionRuleCacheStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyProtectionRuleCacheStatusRequest
//
// @return ModifyProtectionRuleCacheStatusResponse
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

// @param request - ModifyProtectionRuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtectionRuleStatusResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

// @param request - ModifyProtectionRuleStatusRequest
//
// @return ModifyProtectionRuleStatusResponse
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

// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

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

// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
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

// @param request - SetDomainRuleGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainRuleGroupResponse
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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

// @param request - SetDomainRuleGroupRequest
//
// @return SetDomainRuleGroupResponse
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
