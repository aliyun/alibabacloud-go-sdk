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

type CreateAclRuleRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Rules      *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateAclRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRuleRequest) SetSourceIp(v string) *CreateAclRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAclRuleRequest) SetLang(v string) *CreateAclRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateAclRuleRequest) SetRules(v string) *CreateAclRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateAclRuleRequest) SetDomain(v string) *CreateAclRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateAclRuleRequest) SetInstanceId(v string) *CreateAclRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAclRuleRequest) SetRegion(v string) *CreateAclRuleRequest {
	s.Region = &v
	return s
}

type CreateAclRuleResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAclRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAclRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclRuleResponseBody) SetRequestId(v string) *CreateAclRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclRuleResponseBody) SetResult(v *CreateAclRuleResponseBodyResult) *CreateAclRuleResponseBody {
	s.Result = v
	return s
}

type CreateAclRuleResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s CreateAclRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAclRuleResponseBodyResult) SetStatus(v int32) *CreateAclRuleResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAclRuleResponseBodyResult) SetWafTaskId(v string) *CreateAclRuleResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type CreateAclRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAclRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAclRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAclRuleResponse) SetHeaders(v map[string]*string) *CreateAclRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAclRuleResponse) SetBody(v *CreateAclRuleResponseBody) *CreateAclRuleResponse {
	s.Body = v
	return s
}

type CreateCertAndKeyRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Cert          *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	HttpsCertName *string `json:"HttpsCertName,omitempty" xml:"HttpsCertName,omitempty"`
}

func (s CreateCertAndKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertAndKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateCertAndKeyRequest) SetSourceIp(v string) *CreateCertAndKeyRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetLang(v string) *CreateCertAndKeyRequest {
	s.Lang = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetDomain(v string) *CreateCertAndKeyRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetRegion(v string) *CreateCertAndKeyRequest {
	s.Region = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetInstanceId(v string) *CreateCertAndKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetCert(v string) *CreateCertAndKeyRequest {
	s.Cert = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetKey(v string) *CreateCertAndKeyRequest {
	s.Key = &v
	return s
}

func (s *CreateCertAndKeyRequest) SetHttpsCertName(v string) *CreateCertAndKeyRequest {
	s.HttpsCertName = &v
	return s
}

type CreateCertAndKeyResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateCertAndKeyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateCertAndKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertAndKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertAndKeyResponseBody) SetRequestId(v string) *CreateCertAndKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertAndKeyResponseBody) SetResult(v *CreateCertAndKeyResponseBodyResult) *CreateCertAndKeyResponseBody {
	s.Result = v
	return s
}

type CreateCertAndKeyResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s CreateCertAndKeyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateCertAndKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateCertAndKeyResponseBodyResult) SetStatus(v int32) *CreateCertAndKeyResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateCertAndKeyResponseBodyResult) SetWafTaskId(v string) *CreateCertAndKeyResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type CreateCertAndKeyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertAndKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertAndKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertAndKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateCertAndKeyResponse) SetHeaders(v map[string]*string) *CreateCertAndKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateCertAndKeyResponse) SetBody(v *CreateCertAndKeyResponseBody) *CreateCertAndKeyResponse {
	s.Body = v
	return s
}

type CreateDomainConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SourceIps       *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	HttpPort        *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpsPort       *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IsAccessProduct *int32  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	Protocols       *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	LoadBalancing   *int32  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	HttpToUserIp    *int32  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsRedirect   *int32  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
	RsType          *int32  `json:"RsType,omitempty" xml:"RsType,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainConfigRequest) SetSourceIp(v string) *CreateDomainConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateDomainConfigRequest) SetLang(v string) *CreateDomainConfigRequest {
	s.Lang = &v
	return s
}

func (s *CreateDomainConfigRequest) SetDomain(v string) *CreateDomainConfigRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainConfigRequest) SetSourceIps(v string) *CreateDomainConfigRequest {
	s.SourceIps = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpPort(v string) *CreateDomainConfigRequest {
	s.HttpPort = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpsPort(v string) *CreateDomainConfigRequest {
	s.HttpsPort = &v
	return s
}

func (s *CreateDomainConfigRequest) SetInstanceId(v string) *CreateDomainConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainConfigRequest) SetRegion(v string) *CreateDomainConfigRequest {
	s.Region = &v
	return s
}

func (s *CreateDomainConfigRequest) SetIsAccessProduct(v int32) *CreateDomainConfigRequest {
	s.IsAccessProduct = &v
	return s
}

func (s *CreateDomainConfigRequest) SetProtocols(v string) *CreateDomainConfigRequest {
	s.Protocols = &v
	return s
}

func (s *CreateDomainConfigRequest) SetLoadBalancing(v int32) *CreateDomainConfigRequest {
	s.LoadBalancing = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpToUserIp(v int32) *CreateDomainConfigRequest {
	s.HttpToUserIp = &v
	return s
}

func (s *CreateDomainConfigRequest) SetHttpsRedirect(v int32) *CreateDomainConfigRequest {
	s.HttpsRedirect = &v
	return s
}

func (s *CreateDomainConfigRequest) SetRsType(v int32) *CreateDomainConfigRequest {
	s.RsType = &v
	return s
}

func (s *CreateDomainConfigRequest) SetResourceGroupId(v string) *CreateDomainConfigRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateDomainConfigResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateDomainConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainConfigResponseBody) SetRequestId(v string) *CreateDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainConfigResponseBody) SetResult(v *CreateDomainConfigResponseBodyResult) *CreateDomainConfigResponseBody {
	s.Result = v
	return s
}

type CreateDomainConfigResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s CreateDomainConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDomainConfigResponseBodyResult) SetStatus(v int32) *CreateDomainConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateDomainConfigResponseBodyResult) SetWafTaskId(v string) *CreateDomainConfigResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type CreateDomainConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDomainConfigResponse) SetBody(v *CreateDomainConfigResponseBody) *CreateDomainConfigResponse {
	s.Body = v
	return s
}

type CreateProtectionModuleRuleRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Defense    *string `json:"Defense,omitempty" xml:"Defense,omitempty"`
	Rule       *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateProtectionModuleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProtectionModuleRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateProtectionModuleRuleRequest) SetSourceIp(v string) *CreateProtectionModuleRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetLang(v string) *CreateProtectionModuleRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetDomain(v string) *CreateProtectionModuleRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetDefense(v string) *CreateProtectionModuleRuleRequest {
	s.Defense = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetRule(v string) *CreateProtectionModuleRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetInstanceId(v string) *CreateProtectionModuleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateProtectionModuleRuleRequest) SetRegion(v string) *CreateProtectionModuleRuleRequest {
	s.Region = &v
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

type DeleteAclRuleRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteAclRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclRuleRequest) SetSourceIp(v string) *DeleteAclRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAclRuleRequest) SetLang(v string) *DeleteAclRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAclRuleRequest) SetRuleId(v int64) *DeleteAclRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteAclRuleRequest) SetDomain(v string) *DeleteAclRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteAclRuleRequest) SetInstanceId(v string) *DeleteAclRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAclRuleRequest) SetRegion(v string) *DeleteAclRuleRequest {
	s.Region = &v
	return s
}

type DeleteAclRuleResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteAclRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteAclRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclRuleResponseBody) SetRequestId(v string) *DeleteAclRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclRuleResponseBody) SetResult(v *DeleteAclRuleResponseBodyResult) *DeleteAclRuleResponseBody {
	s.Result = v
	return s
}

type DeleteAclRuleResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s DeleteAclRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAclRuleResponseBodyResult) SetStatus(v int32) *DeleteAclRuleResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DeleteAclRuleResponseBodyResult) SetWafTaskId(v string) *DeleteAclRuleResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type DeleteAclRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAclRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAclRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclRuleResponse) SetHeaders(v map[string]*string) *DeleteAclRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclRuleResponse) SetBody(v *DeleteAclRuleResponseBody) *DeleteAclRuleResponse {
	s.Body = v
	return s
}

type DeleteDomainConfigRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainConfigRequest) SetSourceIp(v string) *DeleteDomainConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteDomainConfigRequest) SetLang(v string) *DeleteDomainConfigRequest {
	s.Lang = &v
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

type DeleteDomainConfigResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DeleteDomainConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s *DeleteDomainConfigResponseBody) SetResult(v *DeleteDomainConfigResponseBodyResult) *DeleteDomainConfigResponseBody {
	s.Result = v
	return s
}

type DeleteDomainConfigResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s DeleteDomainConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteDomainConfigResponseBodyResult) SetStatus(v int32) *DeleteDomainConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DeleteDomainConfigResponseBodyResult) SetWafTaskId(v string) *DeleteDomainConfigResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type DeleteDomainConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDomainConfigResponse) SetBody(v *DeleteDomainConfigResponseBody) *DeleteDomainConfigResponse {
	s.Body = v
	return s
}

type DescribeAclRulesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAclRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesRequest) SetSourceIp(v string) *DescribeAclRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAclRulesRequest) SetLang(v string) *DescribeAclRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclRulesRequest) SetRegion(v string) *DescribeAclRulesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAclRulesRequest) SetInstanceId(v string) *DescribeAclRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAclRulesRequest) SetDomain(v string) *DescribeAclRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeAclRulesRequest) SetCurrentPage(v int32) *DescribeAclRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAclRulesRequest) SetPageSize(v int32) *DescribeAclRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeAclRulesResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAclRulesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAclRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponseBody) SetRequestId(v string) *DescribeAclRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclRulesResponseBody) SetResult(v *DescribeAclRulesResponseBodyResult) *DescribeAclRulesResponseBody {
	s.Result = v
	return s
}

type DescribeAclRulesResponseBodyResult struct {
	AclRules *DescribeAclRulesResponseBodyResultAclRules `json:"AclRules,omitempty" xml:"AclRules,omitempty" type:"Struct"`
	Total    *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAclRulesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponseBodyResult) SetAclRules(v *DescribeAclRulesResponseBodyResultAclRules) *DescribeAclRulesResponseBodyResult {
	s.AclRules = v
	return s
}

func (s *DescribeAclRulesResponseBodyResult) SetTotal(v int32) *DescribeAclRulesResponseBodyResult {
	s.Total = &v
	return s
}

type DescribeAclRulesResponseBodyResultAclRules struct {
	AclRule []*DescribeAclRulesResponseBodyResultAclRulesAclRule `json:"AclRule,omitempty" xml:"AclRule,omitempty" type:"Repeated"`
}

func (s DescribeAclRulesResponseBodyResultAclRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponseBodyResultAclRules) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponseBodyResultAclRules) SetAclRule(v []*DescribeAclRulesResponseBodyResultAclRulesAclRule) *DescribeAclRulesResponseBodyResultAclRules {
	s.AclRule = v
	return s
}

type DescribeAclRulesResponseBodyResultAclRulesAclRule struct {
	Action                  *int32                                                       `json:"Action,omitempty" xml:"Action,omitempty"`
	IsDefault               *int32                                                       `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ContinueWaf             *int32                                                       `json:"ContinueWaf,omitempty" xml:"ContinueWaf,omitempty"`
	Order                   *int32                                                       `json:"Order,omitempty" xml:"Order,omitempty"`
	Conditions              *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Struct"`
	ContinueDataRiskControl *int32                                                       `json:"ContinueDataRiskControl,omitempty" xml:"ContinueDataRiskControl,omitempty"`
	ContinueSdk             *int32                                                       `json:"ContinueSdk,omitempty" xml:"ContinueSdk,omitempty"`
	ContinueCc              *int32                                                       `json:"ContinueCc,omitempty" xml:"ContinueCc,omitempty"`
	ContinueSA              *int32                                                       `json:"ContinueSA,omitempty" xml:"ContinueSA,omitempty"`
	ContinueBlockGeo        *int32                                                       `json:"ContinueBlockGeo,omitempty" xml:"ContinueBlockGeo,omitempty"`
	Name                    *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                      *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeAclRulesResponseBodyResultAclRulesAclRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponseBodyResultAclRulesAclRule) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetAction(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.Action = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetIsDefault(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.IsDefault = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetContinueWaf(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.ContinueWaf = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetOrder(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.Order = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetConditions(v *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.Conditions = v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetContinueDataRiskControl(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.ContinueDataRiskControl = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetContinueSdk(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.ContinueSdk = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetContinueCc(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.ContinueCc = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetContinueSA(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.ContinueSA = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetContinueBlockGeo(v int32) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.ContinueBlockGeo = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetName(v string) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.Name = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRule) SetId(v int64) *DescribeAclRulesResponseBodyResultAclRulesAclRule {
	s.Id = &v
	return s
}

type DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions struct {
	Condition []*DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Repeated"`
}

func (s DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions) SetCondition(v []*DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition) *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditions {
	s.Condition = v
	return s
}

type DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition struct {
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value   *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Contain *string `json:"Contain,omitempty" xml:"Contain,omitempty"`
}

func (s DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition) SetKey(v string) *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition {
	s.Key = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition) SetValue(v string) *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition {
	s.Value = &v
	return s
}

func (s *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition) SetContain(v string) *DescribeAclRulesResponseBodyResultAclRulesAclRuleConditionsCondition {
	s.Contain = &v
	return s
}

type DescribeAclRulesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAclRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAclRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclRulesResponse) SetHeaders(v map[string]*string) *DescribeAclRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclRulesResponse) SetBody(v *DescribeAclRulesResponseBody) *DescribeAclRulesResponse {
	s.Body = v
	return s
}

type DescribeAsyncTaskStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	WafRequestId    *string `json:"WafRequestId,omitempty" xml:"WafRequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeAsyncTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTaskStatusRequest) SetSourceIp(v string) *DescribeAsyncTaskStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAsyncTaskStatusRequest) SetLang(v string) *DescribeAsyncTaskStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAsyncTaskStatusRequest) SetRegion(v string) *DescribeAsyncTaskStatusRequest {
	s.Region = &v
	return s
}

func (s *DescribeAsyncTaskStatusRequest) SetInstanceId(v string) *DescribeAsyncTaskStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAsyncTaskStatusRequest) SetWafRequestId(v string) *DescribeAsyncTaskStatusRequest {
	s.WafRequestId = &v
	return s
}

func (s *DescribeAsyncTaskStatusRequest) SetResourceGroupId(v string) *DescribeAsyncTaskStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeAsyncTaskStatusResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeAsyncTaskStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeAsyncTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTaskStatusResponseBody) SetRequestId(v string) *DescribeAsyncTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAsyncTaskStatusResponseBody) SetResult(v *DescribeAsyncTaskStatusResponseBodyResult) *DescribeAsyncTaskStatusResponseBody {
	s.Result = v
	return s
}

type DescribeAsyncTaskStatusResponseBodyResult struct {
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Progress        *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ErrCode         *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMsg          *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	AsyncTaskStatus *string `json:"AsyncTaskStatus,omitempty" xml:"AsyncTaskStatus,omitempty"`
}

func (s DescribeAsyncTaskStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTaskStatusResponseBodyResult) SetData(v string) *DescribeAsyncTaskStatusResponseBodyResult {
	s.Data = &v
	return s
}

func (s *DescribeAsyncTaskStatusResponseBodyResult) SetProgress(v int32) *DescribeAsyncTaskStatusResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *DescribeAsyncTaskStatusResponseBodyResult) SetErrCode(v string) *DescribeAsyncTaskStatusResponseBodyResult {
	s.ErrCode = &v
	return s
}

func (s *DescribeAsyncTaskStatusResponseBodyResult) SetErrMsg(v string) *DescribeAsyncTaskStatusResponseBodyResult {
	s.ErrMsg = &v
	return s
}

func (s *DescribeAsyncTaskStatusResponseBodyResult) SetAsyncTaskStatus(v string) *DescribeAsyncTaskStatusResponseBodyResult {
	s.AsyncTaskStatus = &v
	return s
}

type DescribeAsyncTaskStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAsyncTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAsyncTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTaskStatusResponse) SetHeaders(v map[string]*string) *DescribeAsyncTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAsyncTaskStatusResponse) SetBody(v *DescribeAsyncTaskStatusResponseBody) *DescribeAsyncTaskStatusResponse {
	s.Body = v
	return s
}

type DescribeDomainConfigRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigRequest) SetSourceIp(v string) *DescribeDomainConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainConfigRequest) SetLang(v string) *DescribeDomainConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainConfigRequest) SetRegion(v string) *DescribeDomainConfigRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainConfigRequest) SetInstanceId(v string) *DescribeDomainConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainConfigRequest) SetDomain(v string) *DescribeDomainConfigRequest {
	s.Domain = &v
	return s
}

type DescribeDomainConfigResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDomainConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigResponseBody) SetRequestId(v string) *DescribeDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainConfigResponseBody) SetResult(v *DescribeDomainConfigResponseBodyResult) *DescribeDomainConfigResponseBody {
	s.Result = v
	return s
}

type DescribeDomainConfigResponseBodyResult struct {
	Status       *int32                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId    *string                                             `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
	DomainConfig *DescribeDomainConfigResponseBodyResultDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
}

func (s DescribeDomainConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigResponseBodyResult) SetStatus(v int32) *DescribeDomainConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeDomainConfigResponseBodyResult) SetWafTaskId(v string) *DescribeDomainConfigResponseBodyResult {
	s.WafTaskId = &v
	return s
}

func (s *DescribeDomainConfigResponseBodyResult) SetDomainConfig(v *DescribeDomainConfigResponseBodyResultDomainConfig) *DescribeDomainConfigResponseBodyResult {
	s.DomainConfig = v
	return s
}

type DescribeDomainConfigResponseBodyResultDomainConfig struct {
	Cname        *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	ProtocolType *int32  `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	SourceIps    *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
}

func (s DescribeDomainConfigResponseBodyResultDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigResponseBodyResultDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigResponseBodyResultDomainConfig) SetCname(v string) *DescribeDomainConfigResponseBodyResultDomainConfig {
	s.Cname = &v
	return s
}

func (s *DescribeDomainConfigResponseBodyResultDomainConfig) SetProtocolType(v int32) *DescribeDomainConfigResponseBodyResultDomainConfig {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDomainConfigResponseBodyResultDomainConfig) SetSourceIps(v string) *DescribeDomainConfigResponseBodyResultDomainConfig {
	s.SourceIps = &v
	return s
}

type DescribeDomainConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigResponse) SetHeaders(v map[string]*string) *DescribeDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainConfigResponse) SetBody(v *DescribeDomainConfigResponseBody) *DescribeDomainConfigResponse {
	s.Body = v
	return s
}

type DescribeDomainConfigStatusRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainConfigStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigStatusRequest) SetSourceIp(v string) *DescribeDomainConfigStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainConfigStatusRequest) SetLang(v string) *DescribeDomainConfigStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainConfigStatusRequest) SetRegion(v string) *DescribeDomainConfigStatusRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainConfigStatusRequest) SetInstanceId(v string) *DescribeDomainConfigStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainConfigStatusRequest) SetDomain(v string) *DescribeDomainConfigStatusRequest {
	s.Domain = &v
	return s
}

type DescribeDomainConfigStatusResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDomainConfigStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDomainConfigStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigStatusResponseBody) SetRequestId(v string) *DescribeDomainConfigStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainConfigStatusResponseBody) SetResult(v *DescribeDomainConfigStatusResponseBodyResult) *DescribeDomainConfigStatusResponseBody {
	s.Result = v
	return s
}

type DescribeDomainConfigStatusResponseBodyResult struct {
	Status       *int32                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId    *string                                                   `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
	DomainConfig *DescribeDomainConfigStatusResponseBodyResultDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
}

func (s DescribeDomainConfigStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigStatusResponseBodyResult) SetStatus(v int32) *DescribeDomainConfigStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeDomainConfigStatusResponseBodyResult) SetWafTaskId(v string) *DescribeDomainConfigStatusResponseBodyResult {
	s.WafTaskId = &v
	return s
}

func (s *DescribeDomainConfigStatusResponseBodyResult) SetDomainConfig(v *DescribeDomainConfigStatusResponseBodyResultDomainConfig) *DescribeDomainConfigStatusResponseBodyResult {
	s.DomainConfig = v
	return s
}

type DescribeDomainConfigStatusResponseBodyResultDomainConfig struct {
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
}

func (s DescribeDomainConfigStatusResponseBodyResultDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigStatusResponseBodyResultDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigStatusResponseBodyResultDomainConfig) SetConfigStatus(v string) *DescribeDomainConfigStatusResponseBodyResultDomainConfig {
	s.ConfigStatus = &v
	return s
}

type DescribeDomainConfigStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainConfigStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainConfigStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigStatusResponse) SetHeaders(v map[string]*string) *DescribeDomainConfigStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainConfigStatusResponse) SetBody(v *DescribeDomainConfigStatusResponseBody) *DescribeDomainConfigStatusResponse {
	s.Body = v
	return s
}

type DescribeDomainNamesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDomainNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesRequest) SetSourceIp(v string) *DescribeDomainNamesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainNamesRequest) SetLang(v string) *DescribeDomainNamesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainNamesRequest) SetRegion(v string) *DescribeDomainNamesRequest {
	s.Region = &v
	return s
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
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeDomainNamesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDomainNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesResponseBody) SetRequestId(v string) *DescribeDomainNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainNamesResponseBody) SetResult(v *DescribeDomainNamesResponseBodyResult) *DescribeDomainNamesResponseBody {
	s.Result = v
	return s
}

type DescribeDomainNamesResponseBodyResult struct {
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
}

func (s DescribeDomainNamesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesResponseBodyResult) SetDomainNames(v []*string) *DescribeDomainNamesResponseBodyResult {
	s.DomainNames = v
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

type DescribePayInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceSource  *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePayInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribePayInfoRequest) SetSourceIp(v string) *DescribePayInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePayInfoRequest) SetLang(v string) *DescribePayInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribePayInfoRequest) SetRegion(v string) *DescribePayInfoRequest {
	s.Region = &v
	return s
}

func (s *DescribePayInfoRequest) SetInstanceSource(v string) *DescribePayInfoRequest {
	s.InstanceSource = &v
	return s
}

func (s *DescribePayInfoRequest) SetResourceGroupId(v string) *DescribePayInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePayInfoResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePayInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribePayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePayInfoResponseBody) SetRequestId(v string) *DescribePayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePayInfoResponseBody) SetResult(v *DescribePayInfoResponseBodyResult) *DescribePayInfoResponseBody {
	s.Result = v
	return s
}

type DescribePayInfoResponseBodyResult struct {
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	EndDate    *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	RemainDay  *int32  `json:"RemainDay,omitempty" xml:"RemainDay,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	PayType    *int32  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	InDebt     *int32  `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Trial      *int32  `json:"Trial,omitempty" xml:"Trial,omitempty"`
}

func (s DescribePayInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribePayInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePayInfoResponseBodyResult) SetStatus(v int32) *DescribePayInfoResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetEndDate(v int64) *DescribePayInfoResponseBodyResult {
	s.EndDate = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetRemainDay(v int32) *DescribePayInfoResponseBodyResult {
	s.RemainDay = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetRegion(v string) *DescribePayInfoResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetPayType(v int32) *DescribePayInfoResponseBodyResult {
	s.PayType = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetInDebt(v int32) *DescribePayInfoResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetInstanceId(v string) *DescribePayInfoResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribePayInfoResponseBodyResult) SetTrial(v int32) *DescribePayInfoResponseBodyResult {
	s.Trial = &v
	return s
}

type DescribePayInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePayInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribePayInfoResponse) SetHeaders(v map[string]*string) *DescribePayInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribePayInfoResponse) SetBody(v *DescribePayInfoResponseBody) *DescribePayInfoResponse {
	s.Body = v
	return s
}

type DescribeProtectionModuleRulesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Defense     *string `json:"Defense,omitempty" xml:"Defense,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeProtectionModuleRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesRequest) SetSourceIp(v string) *DescribeProtectionModuleRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetLang(v string) *DescribeProtectionModuleRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetPageSize(v int32) *DescribeProtectionModuleRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetCurrentPage(v int32) *DescribeProtectionModuleRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetDomain(v string) *DescribeProtectionModuleRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetDefense(v string) *DescribeProtectionModuleRulesRequest {
	s.Defense = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetInstanceId(v string) *DescribeProtectionModuleRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProtectionModuleRulesRequest) SetRegion(v string) *DescribeProtectionModuleRulesRequest {
	s.Region = &v
	return s
}

type DescribeProtectionModuleRulesResponseBody struct {
	ModuleRules []*DescribeProtectionModuleRulesResponseBodyModuleRules `json:"ModuleRules,omitempty" xml:"ModuleRules,omitempty" type:"Repeated"`
	RequestId   *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *int32                                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	TaskStatus  *int32                                                  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeProtectionModuleRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesResponseBody) SetModuleRules(v []*DescribeProtectionModuleRulesResponseBodyModuleRules) *DescribeProtectionModuleRulesResponseBody {
	s.ModuleRules = v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBody) SetRequestId(v string) *DescribeProtectionModuleRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBody) SetTotal(v int32) *DescribeProtectionModuleRulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBody) SetTaskStatus(v int32) *DescribeProtectionModuleRulesResponseBody {
	s.TaskStatus = &v
	return s
}

type DescribeProtectionModuleRulesResponseBodyModuleRules struct {
	Time    *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Version *int64  `json:"Version,omitempty" xml:"Version,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeProtectionModuleRulesResponseBodyModuleRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeProtectionModuleRulesResponseBodyModuleRules) GoString() string {
	return s.String()
}

func (s *DescribeProtectionModuleRulesResponseBodyModuleRules) SetTime(v int64) *DescribeProtectionModuleRulesResponseBodyModuleRules {
	s.Time = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyModuleRules) SetVersion(v int64) *DescribeProtectionModuleRulesResponseBodyModuleRules {
	s.Version = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyModuleRules) SetContent(v string) *DescribeProtectionModuleRulesResponseBodyModuleRules {
	s.Content = &v
	return s
}

func (s *DescribeProtectionModuleRulesResponseBodyModuleRules) SetId(v int64) *DescribeProtectionModuleRulesResponseBodyModuleRules {
	s.Id = &v
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

type DescribeRegionsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSourceIp(v string) *DescribeRegionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegion(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.Region = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetDisplay(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.Display = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeWafSourceIpSegmentRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWafSourceIpSegmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentRequest) SetSourceIp(v string) *DescribeWafSourceIpSegmentRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetLang(v string) *DescribeWafSourceIpSegmentRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetInstanceId(v string) *DescribeWafSourceIpSegmentRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetRegion(v string) *DescribeWafSourceIpSegmentRequest {
	s.Region = &v
	return s
}

func (s *DescribeWafSourceIpSegmentRequest) SetResourceGroupId(v string) *DescribeWafSourceIpSegmentRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeWafSourceIpSegmentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ips       *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
}

func (s DescribeWafSourceIpSegmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetRequestId(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetIps(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.Ips = &v
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

type ModifyAclRuleRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Rules      *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyAclRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAclRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAclRuleRequest) SetSourceIp(v string) *ModifyAclRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAclRuleRequest) SetLang(v string) *ModifyAclRuleRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAclRuleRequest) SetDomain(v string) *ModifyAclRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyAclRuleRequest) SetRules(v string) *ModifyAclRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyAclRuleRequest) SetInstanceId(v string) *ModifyAclRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAclRuleRequest) SetRegion(v string) *ModifyAclRuleRequest {
	s.Region = &v
	return s
}

type ModifyAclRuleResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifyAclRuleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ModifyAclRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAclRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAclRuleResponseBody) SetRequestId(v string) *ModifyAclRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAclRuleResponseBody) SetResult(v *ModifyAclRuleResponseBodyResult) *ModifyAclRuleResponseBody {
	s.Result = v
	return s
}

type ModifyAclRuleResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s ModifyAclRuleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyAclRuleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAclRuleResponseBodyResult) SetStatus(v int32) *ModifyAclRuleResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAclRuleResponseBodyResult) SetWafTaskId(v string) *ModifyAclRuleResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type ModifyAclRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAclRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAclRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAclRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAclRuleResponse) SetHeaders(v map[string]*string) *ModifyAclRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAclRuleResponse) SetBody(v *ModifyAclRuleResponseBody) *ModifyAclRuleResponse {
	s.Body = v
	return s
}

type ModifyDomainConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	SourceIps       *string `json:"SourceIps,omitempty" xml:"SourceIps,omitempty"`
	HttpPort        *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	HttpsPort       *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IsAccessProduct *int32  `json:"IsAccessProduct,omitempty" xml:"IsAccessProduct,omitempty"`
	Protocols       *string `json:"Protocols,omitempty" xml:"Protocols,omitempty"`
	LoadBalancing   *int32  `json:"LoadBalancing,omitempty" xml:"LoadBalancing,omitempty"`
	HttpToUserIp    *int32  `json:"HttpToUserIp,omitempty" xml:"HttpToUserIp,omitempty"`
	HttpsRedirect   *int32  `json:"HttpsRedirect,omitempty" xml:"HttpsRedirect,omitempty"`
}

func (s ModifyDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainConfigRequest) SetSourceIp(v string) *ModifyDomainConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetLang(v string) *ModifyDomainConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetDomain(v string) *ModifyDomainConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetSourceIps(v string) *ModifyDomainConfigRequest {
	s.SourceIps = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpPort(v string) *ModifyDomainConfigRequest {
	s.HttpPort = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpsPort(v string) *ModifyDomainConfigRequest {
	s.HttpsPort = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetInstanceId(v string) *ModifyDomainConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetRegion(v string) *ModifyDomainConfigRequest {
	s.Region = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetIsAccessProduct(v int32) *ModifyDomainConfigRequest {
	s.IsAccessProduct = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetProtocols(v string) *ModifyDomainConfigRequest {
	s.Protocols = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetLoadBalancing(v int32) *ModifyDomainConfigRequest {
	s.LoadBalancing = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpToUserIp(v int32) *ModifyDomainConfigRequest {
	s.HttpToUserIp = &v
	return s
}

func (s *ModifyDomainConfigRequest) SetHttpsRedirect(v int32) *ModifyDomainConfigRequest {
	s.HttpsRedirect = &v
	return s
}

type ModifyDomainConfigResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifyDomainConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s *ModifyDomainConfigResponseBody) SetResult(v *ModifyDomainConfigResponseBodyResult) *ModifyDomainConfigResponseBody {
	s.Result = v
	return s
}

type ModifyDomainConfigResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s ModifyDomainConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyDomainConfigResponseBodyResult) SetStatus(v int32) *ModifyDomainConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyDomainConfigResponseBodyResult) SetWafTaskId(v string) *ModifyDomainConfigResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type ModifyDomainConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyDomainConfigResponse) SetBody(v *ModifyDomainConfigResponseBody) *ModifyDomainConfigResponse {
	s.Body = v
	return s
}

type ModifyProtectionRuleCacheStatusRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Defense    *string `json:"Defense,omitempty" xml:"Defense,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyProtectionRuleCacheStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleCacheStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetSourceIp(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetLang(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetDomain(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetId(v int64) *ModifyProtectionRuleCacheStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetDefense(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.Defense = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetInstanceId(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusRequest) SetRegion(v string) *ModifyProtectionRuleCacheStatusRequest {
	s.Region = &v
	return s
}

type ModifyProtectionRuleCacheStatusResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStatus *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	WafTaskId  *int32  `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
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

func (s *ModifyProtectionRuleCacheStatusResponseBody) SetTaskStatus(v int32) *ModifyProtectionRuleCacheStatusResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *ModifyProtectionRuleCacheStatusResponseBody) SetWafTaskId(v int32) *ModifyProtectionRuleCacheStatusResponseBody {
	s.WafTaskId = &v
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
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Defense     *string `json:"Defense,omitempty" xml:"Defense,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleStatus  *int32  `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	LockVersion *int64  `json:"LockVersion,omitempty" xml:"LockVersion,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyProtectionRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProtectionRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtectionRuleStatusRequest) SetSourceIp(v string) *ModifyProtectionRuleStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetLang(v string) *ModifyProtectionRuleStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetDomain(v string) *ModifyProtectionRuleStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetDefense(v string) *ModifyProtectionRuleStatusRequest {
	s.Defense = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetId(v int64) *ModifyProtectionRuleStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetRuleStatus(v int32) *ModifyProtectionRuleStatusRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetLockVersion(v int64) *ModifyProtectionRuleStatusRequest {
	s.LockVersion = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetInstanceId(v string) *ModifyProtectionRuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProtectionRuleStatusRequest) SetRegion(v string) *ModifyProtectionRuleStatusRequest {
	s.Region = &v
	return s
}

type ModifyProtectionRuleStatusResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStatus *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	WafTaskId  *int32  `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
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

func (s *ModifyProtectionRuleStatusResponseBody) SetTaskStatus(v int32) *ModifyProtectionRuleStatusResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *ModifyProtectionRuleStatusResponseBody) SetWafTaskId(v int32) *ModifyProtectionRuleStatusResponseBody {
	s.WafTaskId = &v
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

type ModifyWafSwitchRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServiceOn  *int32  `json:"ServiceOn,omitempty" xml:"ServiceOn,omitempty"`
}

func (s ModifyWafSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWafSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWafSwitchRequest) SetSourceIp(v string) *ModifyWafSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetLang(v string) *ModifyWafSwitchRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWafSwitchRequest) SetRegion(v string) *ModifyWafSwitchRequest {
	s.Region = &v
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

func (s *ModifyWafSwitchRequest) SetServiceOn(v int32) *ModifyWafSwitchRequest {
	s.ServiceOn = &v
	return s
}

type ModifyWafSwitchResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ModifyWafSwitchResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s *ModifyWafSwitchResponseBody) SetResult(v *ModifyWafSwitchResponseBodyResult) *ModifyWafSwitchResponseBody {
	s.Result = v
	return s
}

type ModifyWafSwitchResponseBodyResult struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	WafTaskId *string `json:"WafTaskId,omitempty" xml:"WafTaskId,omitempty"`
}

func (s ModifyWafSwitchResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyWafSwitchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyWafSwitchResponseBodyResult) SetStatus(v int32) *ModifyWafSwitchResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyWafSwitchResponseBodyResult) SetWafTaskId(v string) *ModifyWafSwitchResponseBodyResult {
	s.WafTaskId = &v
	return s
}

type ModifyWafSwitchResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWafSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CreateAclRuleWithOptions(request *CreateAclRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAclRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAclRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAclRule"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAclRule(request *CreateAclRuleRequest) (_result *CreateAclRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAclRuleResponse{}
	_body, _err := client.CreateAclRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCertAndKeyWithOptions(request *CreateCertAndKeyRequest, runtime *util.RuntimeOptions) (_result *CreateCertAndKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCertAndKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCertAndKey"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertAndKey(request *CreateCertAndKeyRequest) (_result *CreateCertAndKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertAndKeyResponse{}
	_body, _err := client.CreateCertAndKeyWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDomainConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDomainConfig"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateProtectionModuleRuleWithOptions(request *CreateProtectionModuleRuleRequest, runtime *util.RuntimeOptions) (_result *CreateProtectionModuleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProtectionModuleRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProtectionModuleRule"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteAclRuleWithOptions(request *DeleteAclRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAclRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAclRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAclRule"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAclRule(request *DeleteAclRuleRequest) (_result *DeleteAclRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAclRuleResponse{}
	_body, _err := client.DeleteAclRuleWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDomainConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomainConfig"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeAclRulesWithOptions(request *DescribeAclRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeAclRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAclRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAclRules"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAclRules(request *DescribeAclRulesRequest) (_result *DescribeAclRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAclRulesResponse{}
	_body, _err := client.DescribeAclRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAsyncTaskStatusWithOptions(request *DescribeAsyncTaskStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeAsyncTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAsyncTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAsyncTaskStatus"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAsyncTaskStatus(request *DescribeAsyncTaskStatusRequest) (_result *DescribeAsyncTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAsyncTaskStatusResponse{}
	_body, _err := client.DescribeAsyncTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainConfigWithOptions(request *DescribeDomainConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainConfig"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainConfig(request *DescribeDomainConfigRequest) (_result *DescribeDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainConfigResponse{}
	_body, _err := client.DescribeDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainConfigStatusWithOptions(request *DescribeDomainConfigStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainConfigStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainConfigStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainConfigStatus"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainConfigStatus(request *DescribeDomainConfigStatusRequest) (_result *DescribeDomainConfigStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainConfigStatusResponse{}
	_body, _err := client.DescribeDomainConfigStatusWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainNamesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainNames"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribePayInfoWithOptions(request *DescribePayInfoRequest, runtime *util.RuntimeOptions) (_result *DescribePayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePayInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePayInfo"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePayInfo(request *DescribePayInfoRequest) (_result *DescribePayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePayInfoResponse{}
	_body, _err := client.DescribePayInfoWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProtectionModuleRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProtectionModuleRules"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeWafSourceIpSegmentWithOptions(request *DescribeWafSourceIpSegmentRequest, runtime *util.RuntimeOptions) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWafSourceIpSegment"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyAclRuleWithOptions(request *ModifyAclRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyAclRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAclRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAclRule"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAclRule(request *ModifyAclRuleRequest) (_result *ModifyAclRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAclRuleResponse{}
	_body, _err := client.ModifyAclRuleWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDomainConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDomainConfig"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyProtectionRuleCacheStatusWithOptions(request *ModifyProtectionRuleCacheStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyProtectionRuleCacheStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyProtectionRuleCacheStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyProtectionRuleCacheStatus"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyProtectionRuleStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyProtectionRuleStatus"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyWafSwitchWithOptions(request *ModifyWafSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWafSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWafSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWafSwitch"), tea.String("2018-01-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
