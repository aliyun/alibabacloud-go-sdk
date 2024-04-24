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

type ClearMajorProtectionBlackIpRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the IP address blacklist rule template for major event protection.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ClearMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpRequest) SetInstanceId(v string) *ClearMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetRegionId(v string) *ClearMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *ClearMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetRuleId(v int64) *ClearMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetTemplateId(v int64) *ClearMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type ClearMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpResponseBody) SetRequestId(v string) *ClearMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type ClearMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *ClearMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) SetStatusCode(v int32) *ClearMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) SetBody(v *ClearMajorProtectionBlackIpResponseBody) *ClearMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type CopyDefenseTemplateRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection template that you want to copy.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CopyDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateRequest) SetInstanceId(v string) *CopyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetRegionId(v string) *CopyDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *CopyDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetTemplateId(v int64) *CopyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

type CopyDefenseTemplateResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the new protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CopyDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateResponseBody) SetRequestId(v string) *CopyDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyDefenseTemplateResponseBody) SetTemplateId(v int64) *CopyDefenseTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CopyDefenseTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateResponse) SetHeaders(v map[string]*string) *CopyDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *CopyDefenseTemplateResponse) SetStatusCode(v int32) *CopyDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDefenseTemplateResponse) SetBody(v *CopyDefenseTemplateResponseBody) *CopyDefenseTemplateResponse {
	s.Body = v
	return s
}

type CreateDefenseResourceGroupRequest struct {
	// The protected objects that you want to add to the protected object group. You can add multiple protected objects to a protected object group at the same time. You can specify multiple protected objects. Separate them with commas (,).
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	// The description of the protected object group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the protected object group that you want to create.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreateDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupRequest) SetAddList(v string) *CreateDefenseResourceGroupRequest {
	s.AddList = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetDescription(v string) *CreateDefenseResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetGroupName(v string) *CreateDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetInstanceId(v string) *CreateDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetRegionId(v string) *CreateDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type CreateDefenseResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupResponseBody) SetRequestId(v string) *CreateDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *CreateDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseResourceGroupResponse) SetStatusCode(v int32) *CreateDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseResourceGroupResponse) SetBody(v *CreateDefenseResourceGroupResponseBody) *CreateDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type CreateDefenseRuleRequest struct {
	// The module to which the protection rule that you want to create belongs.
	//
	// *   **waf_group:** the basic protection rule module.
	// *   **antiscan:** the scan protection module.
	// *   **ip_blacklist:** the IP address blacklist module.
	// *   **custom_acl:** the custom rule module.
	// *   **whitelist:** the whitelist module.
	// *   **region_block:** the region blacklist module.
	// *   **custom_response:** the custom response module.
	// *   **cc:** the HTTP flood protection module.
	// *   **tamperproof:** the website tamper-proofing module.
	// *   **dlp:** the data leakage prevention module.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The configurations of the protection rule. Specify a string that contains multiple parameters in the JSON format.
	//
	// >  The parameters vary based on the value of the **DefenseScene** parameter. For more information, see the "**Protection rule parameters**" section in this topic.
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the protection rule template for which you want to create a protection rule.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleRequest) SetDefenseScene(v string) *CreateDefenseRuleRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetInstanceId(v string) *CreateDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetRegionId(v string) *CreateDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetRules(v string) *CreateDefenseRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetTemplateId(v int64) *CreateDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type CreateDefenseRuleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleResponseBody) SetRequestId(v string) *CreateDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleResponse) SetHeaders(v map[string]*string) *CreateDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseRuleResponse) SetStatusCode(v int32) *CreateDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseRuleResponse) SetBody(v *CreateDefenseRuleResponseBody) *CreateDefenseRuleResponse {
	s.Body = v
	return s
}

type CreateDefenseTemplateRequest struct {
	// The scenario in which you want to use the protection rule template. For more information, see the description of the **DefenseScene** parameter in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The description of the protection rule template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The name of the protection rule template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection rule template that you want to create. Set the value to **custom**. The value specifies that the protection rule template is a custom template.
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection rule template. Valid values:
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection rule template. Valid values:
	//
	// *   **user_default:** default template.
	// *   **user_custom:** custom template.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateRequest) SetDefenseScene(v string) *CreateDefenseTemplateRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetDescription(v string) *CreateDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetInstanceId(v string) *CreateDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetRegionId(v string) *CreateDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateName(v string) *CreateDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateOrigin(v string) *CreateDefenseTemplateRequest {
	s.TemplateOrigin = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateStatus(v int32) *CreateDefenseTemplateRequest {
	s.TemplateStatus = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateType(v string) *CreateDefenseTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateDefenseTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateResponseBody) SetRequestId(v string) *CreateDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefenseTemplateResponseBody) SetTemplateId(v int64) *CreateDefenseTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateResponse) SetHeaders(v map[string]*string) *CreateDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseTemplateResponse) SetStatusCode(v int32) *CreateDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseTemplateResponse) SetBody(v *CreateDefenseTemplateResponseBody) *CreateDefenseTemplateResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// The mode in which you want to add the domain name to WAF. Valid values:
	//
	// *   **share:** adds the domain name to WAF in CNAME record mode. This is the default value.
	// *   **hybrid_cloud_cname:** adds the domain name to WAF in hybrid cloud reverse proxy mode.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that you want to add to WAF.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configurations of the listeners.
	Listen *CreateDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The configurations of the forwarding rule.
	Redirect *CreateDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou**: the Chinese mainland
	// *   **ap-southeast-1**: outside the Chinese mainland
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetAccessType(v string) *CreateDomainRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetInstanceId(v string) *CreateDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainRequest) SetListen(v *CreateDomainRequestListen) *CreateDomainRequest {
	s.Listen = v
	return s
}

func (s *CreateDomainRequest) SetRedirect(v *CreateDomainRequestRedirect) *CreateDomainRequest {
	s.Redirect = v
	return s
}

func (s *CreateDomainRequest) SetRegionId(v string) *CreateDomainRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainRequest) SetResourceManagerResourceGroupId(v string) *CreateDomainRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type CreateDomainRequestListen struct {
	// The ID of the certificate that you want to add. This parameter is available only if you specify **HttpsPorts**.
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of cipher suite that you want to add. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// *   **1:** all cipher suites.
	// *   **2:** strong cipher suites. You can select this value only if you set **TLSVersion** to **tlsv1.2**.
	// *   **99:** custom cipher suites.
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suite that you want to add.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// *   **true**
	// *   **false**
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable an exclusive IP address. This parameter is available only if you set **IPv6Enabled** to **false** and **ProtectionResource** to **share**. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable HTTP to HTTPS redirection. This parameter is available only if you specify HttpsPorts and leave HttpPorts empty. Valid values:
	//
	// *   **true**
	// *   **false**
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The HTTP listener port.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The HTTPS listener port.
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of the protection resource. Valid values:
	//
	// *   **share:** a shared cluster. This is the default value.
	// *   **gslb:** shared cluster-based intelligent load balancing.
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Specifies whether to allow access only from SM certificate-based clients. This parameter is available only if you set SM2Enabled to true.
	//
	// *   true
	// *   false
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate that you want to add. This parameter is available only if you set SM2Enabled to true.
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Specifies whether to enable the ShangMi (SM) certificate.
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The version of the Transport Layer Security (TLS) protocol. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// *   **tlsv1**
	// *   **tlsv1.1**
	// *   **tlsv1.2**
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that you want WAF to use to obtain the actual IP address of a client. Valid values:
	//
	// *   **0:** No Layer 7 proxies are deployed in front of WAF. This is the default value.
	// *   **1:** WAF reads the first value of the X-Forwarded-For (XFF) header field as the IP address of the client.
	// *   **2:** WAF reads the value of a custom header field as the IP address of the client.
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header field that you want WAF to use to obtain the actual IP address of a client.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s CreateDomainRequestListen) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestListen) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestListen) SetCertId(v string) *CreateDomainRequestListen {
	s.CertId = &v
	return s
}

func (s *CreateDomainRequestListen) SetCipherSuite(v int32) *CreateDomainRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *CreateDomainRequestListen) SetCustomCiphers(v []*string) *CreateDomainRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *CreateDomainRequestListen) SetEnableTLSv3(v bool) *CreateDomainRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *CreateDomainRequestListen) SetExclusiveIp(v bool) *CreateDomainRequestListen {
	s.ExclusiveIp = &v
	return s
}

func (s *CreateDomainRequestListen) SetFocusHttps(v bool) *CreateDomainRequestListen {
	s.FocusHttps = &v
	return s
}

func (s *CreateDomainRequestListen) SetHttp2Enabled(v bool) *CreateDomainRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetHttpPorts(v []*int32) *CreateDomainRequestListen {
	s.HttpPorts = v
	return s
}

func (s *CreateDomainRequestListen) SetHttpsPorts(v []*int32) *CreateDomainRequestListen {
	s.HttpsPorts = v
	return s
}

func (s *CreateDomainRequestListen) SetIPv6Enabled(v bool) *CreateDomainRequestListen {
	s.IPv6Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetProtectionResource(v string) *CreateDomainRequestListen {
	s.ProtectionResource = &v
	return s
}

func (s *CreateDomainRequestListen) SetSM2AccessOnly(v bool) *CreateDomainRequestListen {
	s.SM2AccessOnly = &v
	return s
}

func (s *CreateDomainRequestListen) SetSM2CertId(v string) *CreateDomainRequestListen {
	s.SM2CertId = &v
	return s
}

func (s *CreateDomainRequestListen) SetSM2Enabled(v bool) *CreateDomainRequestListen {
	s.SM2Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetTLSVersion(v string) *CreateDomainRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *CreateDomainRequestListen) SetXffHeaderMode(v int32) *CreateDomainRequestListen {
	s.XffHeaderMode = &v
	return s
}

func (s *CreateDomainRequestListen) SetXffHeaders(v []*string) *CreateDomainRequestListen {
	s.XffHeaders = v
	return s
}

type CreateDomainRequestRedirect struct {
	// The IP addresses or domain names of the origin server.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// Specifies whether to enable the public cloud disaster recovery feature. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The timeout period of connections. Unit: seconds. Valid values: 1 to 3600.
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Specifies whether to enable HTTPS to HTTP redirection for back-to-origin requests. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// *   **true**
	// *   **false**
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Specifies whether to enable the persistent connection feature. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter specifies the number of reused persistent connections after you enable the persistent connection feature.
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of idle persistent connections. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// >  This parameter specifies the period of time during which a reused persistent connection is allowed to remain in the Idle state before the persistent connection is released.
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that you want to use to forward requests to the origin server. Valid values:
	//
	// *   **iphash**
	// *   **roundRobin**
	// *   **leastTime** You can set the parameter to this value only if you set **ProtectionResource** to **gslb**.
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The timeout period of read connections. Unit: seconds. Valid values: 1 to 3600.
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The custom header field that you want to use to label requests that are processed by WAF.
	//
	// When a request passes through WAF, the custom header field is automatically used to label the request. This way, the backend service can identify requests that are processed by WAF.
	RequestHeaders []*CreateDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Specifies whether WAF retries forwarding requests to the origin server when the requests fail to be forwarded to the origin server. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules that you want to configure for the domain name that you want to add to WAF in hybrid cloud mode. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// *   **rs**: the back-to-origin IP addresses or CNAMEs. The value must be of the ARRAY type.
	// *   **location**: the name of the protection node. The value must be of the STRING type.
	// *   **locationId**: the ID of the protection node. The value must be of the LONG type.
	RoutingRules *string `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	// Specifies whether to enable origin Server Name Indication (SNI). This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the SNI field. If you do not specify this parameter, the value of the **Host** field is automatically used. This parameter is optional. If you want WAF to use an SNI field value that is different from the Host field value in back-to-origin requests, you can specify a custom value for the SNI field.
	//
	// >  This parameter is required only if you set **SniEnalbed** to **true**.
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// The timeout period of write connections. Unit: seconds. Valid values: 1 to 3600.
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Indicates whether the X-Forward-For-Proto header is used to identify the protocol used by WAF to forward requests to the origin server. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s CreateDomainRequestRedirect) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestRedirect) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirect) SetBackends(v []*string) *CreateDomainRequestRedirect {
	s.Backends = v
	return s
}

func (s *CreateDomainRequestRedirect) SetCnameEnabled(v bool) *CreateDomainRequestRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetConnectTimeout(v int32) *CreateDomainRequestRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetFocusHttpBackend(v bool) *CreateDomainRequestRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepalive(v bool) *CreateDomainRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepaliveRequests(v int32) *CreateDomainRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepaliveTimeout(v int32) *CreateDomainRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetLoadbalance(v string) *CreateDomainRequestRedirect {
	s.Loadbalance = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetReadTimeout(v int32) *CreateDomainRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetRequestHeaders(v []*CreateDomainRequestRedirectRequestHeaders) *CreateDomainRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *CreateDomainRequestRedirect) SetRetry(v bool) *CreateDomainRequestRedirect {
	s.Retry = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetRoutingRules(v string) *CreateDomainRequestRedirect {
	s.RoutingRules = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetSniEnabled(v bool) *CreateDomainRequestRedirect {
	s.SniEnabled = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetSniHost(v string) *CreateDomainRequestRedirect {
	s.SniHost = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetWriteTimeout(v int32) *CreateDomainRequestRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetXffProto(v bool) *CreateDomainRequestRedirect {
	s.XffProto = &v
	return s
}

type CreateDomainRequestRedirectRequestHeaders struct {
	// The custom header field.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom header field.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDomainRequestRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirectRequestHeaders) SetKey(v string) *CreateDomainRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *CreateDomainRequestRedirectRequestHeaders) SetValue(v string) *CreateDomainRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

type CreateDomainShrinkRequest struct {
	// The mode in which you want to add the domain name to WAF. Valid values:
	//
	// *   **share:** adds the domain name to WAF in CNAME record mode. This is the default value.
	// *   **hybrid_cloud_cname:** adds the domain name to WAF in hybrid cloud reverse proxy mode.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that you want to add to WAF.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configurations of the listeners.
	ListenShrink *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	// The configurations of the forwarding rule.
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou**: the Chinese mainland
	// *   **ap-southeast-1**: outside the Chinese mainland
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreateDomainShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainShrinkRequest) SetAccessType(v string) *CreateDomainShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetDomain(v string) *CreateDomainShrinkRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetInstanceId(v string) *CreateDomainShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetListenShrink(v string) *CreateDomainShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetRedirectShrink(v string) *CreateDomainShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetRegionId(v string) *CreateDomainShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetResourceManagerResourceGroupId(v string) *CreateDomainShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type CreateDomainResponseBody struct {
	// The information about the domain name.
	DomainInfo *CreateDomainResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetDomainInfo(v *CreateDomainResponseBodyDomainInfo) *CreateDomainResponseBody {
	s.DomainInfo = v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponseBodyDomainInfo struct {
	// The CNAME that is assigned by WAF to the domain name.
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name that you added to WAF.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CreateDomainResponseBodyDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBodyDomainInfo) SetCname(v string) *CreateDomainResponseBodyDomainInfo {
	s.Cname = &v
	return s
}

func (s *CreateDomainResponseBodyDomainInfo) SetDomain(v string) *CreateDomainResponseBodyDomainInfo {
	s.Domain = &v
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

type CreateMajorProtectionBlackIpRequest struct {
	// The description of the IP address blacklist.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time after which the IP address blacklist becomes invalid. Unit: seconds.
	//
	// >  If you set the value to **0**, the blacklist is permanently valid.
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP addresses that you want to add to the IP address blacklist. CIDR blocks and IP addresses are supported. IPv4 and IPv6 addresses are supported. Separate the CIDR blocks or IP addresses with commas (,). For more information, see [Protection for major events](~~425591~~).
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the IP address blacklist rule template for major event protection.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpRequest) SetDescription(v string) *CreateMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *CreateMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetInstanceId(v string) *CreateMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetIpList(v string) *CreateMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetRegionId(v string) *CreateMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *CreateMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetRuleId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetTemplateId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type CreateMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponseBody) SetRequestId(v string) *CreateMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type CreateMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetStatusCode(v int32) *CreateMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetBody(v *CreateMajorProtectionBlackIpResponseBody) *CreateMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type CreateMemberAccountsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account IDs of the members that you want to add. You can add up to 10 members at the same time.
	MemberAccountIds []*string `json:"MemberAccountIds,omitempty" xml:"MemberAccountIds,omitempty" type:"Repeated"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source IP address of the request. The system automatically obtains the value of this parameter.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateMemberAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberAccountsRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberAccountsRequest) SetInstanceId(v string) *CreateMemberAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMemberAccountsRequest) SetMemberAccountIds(v []*string) *CreateMemberAccountsRequest {
	s.MemberAccountIds = v
	return s
}

func (s *CreateMemberAccountsRequest) SetRegionId(v string) *CreateMemberAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMemberAccountsRequest) SetResourceManagerResourceGroupId(v string) *CreateMemberAccountsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateMemberAccountsRequest) SetSourceIp(v string) *CreateMemberAccountsRequest {
	s.SourceIp = &v
	return s
}

type CreateMemberAccountsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMemberAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberAccountsResponseBody) SetRequestId(v string) *CreateMemberAccountsResponseBody {
	s.RequestId = &v
	return s
}

type CreateMemberAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemberAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberAccountsResponse) GoString() string {
	return s.String()
}

func (s *CreateMemberAccountsResponse) SetHeaders(v map[string]*string) *CreateMemberAccountsResponse {
	s.Headers = v
	return s
}

func (s *CreateMemberAccountsResponse) SetStatusCode(v int32) *CreateMemberAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemberAccountsResponse) SetBody(v *CreateMemberAccountsResponseBody) *CreateMemberAccountsResponse {
	s.Body = v
	return s
}

type CreatePostpaidInstanceRequest struct {
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreatePostpaidInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePostpaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePostpaidInstanceRequest) SetRegionId(v string) *CreatePostpaidInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostpaidInstanceRequest) SetResourceManagerResourceGroupId(v string) *CreatePostpaidInstanceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type CreatePostpaidInstanceResponseBody struct {
	// The ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePostpaidInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePostpaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostpaidInstanceResponseBody) SetInstanceId(v string) *CreatePostpaidInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreatePostpaidInstanceResponseBody) SetRequestId(v string) *CreatePostpaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreatePostpaidInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostpaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostpaidInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePostpaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePostpaidInstanceResponse) SetHeaders(v map[string]*string) *CreatePostpaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePostpaidInstanceResponse) SetStatusCode(v int32) *CreatePostpaidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostpaidInstanceResponse) SetBody(v *CreatePostpaidInstanceResponseBody) *CreatePostpaidInstanceResponse {
	s.Body = v
	return s
}

type DeleteDefenseResourceGroupRequest struct {
	// The name of the protected object group that you want to delete.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupRequest) SetGroupName(v string) *DeleteDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetInstanceId(v string) *DeleteDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetRegionId(v string) *DeleteDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DeleteDefenseResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupResponseBody) SetRequestId(v string) *DeleteDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) SetStatusCode(v int32) *DeleteDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) SetBody(v *DeleteDefenseResourceGroupResponseBody) *DeleteDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteDefenseRuleRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The IDs of the protection rules that you want to delete. Separate the IDs with commas (,).
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The ID of the protection rule template to which the protection rule that you want to delete belongs.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleRequest) SetInstanceId(v string) *DeleteDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetRegionId(v string) *DeleteDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetRuleIds(v string) *DeleteDefenseRuleRequest {
	s.RuleIds = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetTemplateId(v int64) *DeleteDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type DeleteDefenseRuleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleResponseBody) SetRequestId(v string) *DeleteDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleResponse) SetHeaders(v map[string]*string) *DeleteDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseRuleResponse) SetStatusCode(v int32) *DeleteDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseRuleResponse) SetBody(v *DeleteDefenseRuleResponseBody) *DeleteDefenseRuleResponse {
	s.Body = v
	return s
}

type DeleteDefenseTemplateRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template that you want to delete.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateRequest) SetInstanceId(v string) *DeleteDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetRegionId(v string) *DeleteDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetTemplateId(v int64) *DeleteDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteDefenseTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateResponseBody) SetRequestId(v string) *DeleteDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateResponse) SetHeaders(v map[string]*string) *DeleteDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseTemplateResponse) SetStatusCode(v int32) *DeleteDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseTemplateResponse) SetBody(v *DeleteDefenseTemplateResponseBody) *DeleteDefenseTemplateResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// The mode in which the domain name is added to WAF. Valid values:
	//
	// *   **share:** CNAME record mode. This is the default value.
	// *   **hybrid_cloud_cname:** hybrid cloud reverse proxy mode.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that you want to delete.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the domain name.
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetAccessType(v string) *DeleteDomainRequest {
	s.AccessType = &v
	return s
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainRequest) SetDomainId(v string) *DeleteDomainRequest {
	s.DomainId = &v
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

type DeleteDomainResponseBody struct {
	// The ID of the request.
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

type DeleteMajorProtectionBlackIpRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address blacklist for major event protection that you want to delete. You can specify multiple CIDR blocks or IP addresses. IPv4 and IPv6 addresses are supported. Separate the CIDR blocks or IP addresses with commas (,). For more information, see [Protection for major events](~~425591~~).
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the IP address blacklist rule template for major event protection.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpRequest) SetInstanceId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetIpList(v string) *DeleteMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetRegionId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetRuleId(v int64) *DeleteMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetTemplateId(v int64) *DeleteMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type DeleteMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpResponseBody) SetRequestId(v string) *DeleteMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *DeleteMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) SetStatusCode(v int32) *DeleteMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) SetBody(v *DeleteMajorProtectionBlackIpResponseBody) *DeleteMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type DeleteMemberAccountRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID of the managed member.
	MemberAccountId *string `json:"MemberAccountId,omitempty" xml:"MemberAccountId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source IP address of the request. The system automatically obtains the value of this parameter.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteMemberAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemberAccountRequest) SetInstanceId(v string) *DeleteMemberAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetMemberAccountId(v string) *DeleteMemberAccountRequest {
	s.MemberAccountId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetRegionId(v string) *DeleteMemberAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetResourceManagerResourceGroupId(v string) *DeleteMemberAccountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteMemberAccountRequest) SetSourceIp(v string) *DeleteMemberAccountRequest {
	s.SourceIp = &v
	return s
}

type DeleteMemberAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMemberAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemberAccountResponseBody) SetRequestId(v string) *DeleteMemberAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMemberAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemberAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemberAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemberAccountResponse) SetHeaders(v map[string]*string) *DeleteMemberAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemberAccountResponse) SetStatusCode(v int32) *DeleteMemberAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemberAccountResponse) SetBody(v *DeleteMemberAccountResponseBody) *DeleteMemberAccountResponse {
	s.Body = v
	return s
}

type DescribeAccountDelegatedStatusRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAccountDelegatedStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDelegatedStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountDelegatedStatusRequest) SetInstanceId(v string) *DescribeAccountDelegatedStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusRequest) SetRegionId(v string) *DescribeAccountDelegatedStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeAccountDelegatedStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeAccountDelegatedStatusResponseBody struct {
	// The ID of the Alibaba Cloud account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the Alibaba Cloud account. This parameter is returned only if the account is the delegated administrator account.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the Alibaba Cloud account is the delegated administrator account of the WAF instance.
	//
	// *   **true**
	// *   **false**
	DelegatedStatus *bool `json:"DelegatedStatus,omitempty" xml:"DelegatedStatus,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountDelegatedStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDelegatedStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetAccountId(v string) *DescribeAccountDelegatedStatusResponseBody {
	s.AccountId = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetAccountName(v string) *DescribeAccountDelegatedStatusResponseBody {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetDelegatedStatus(v bool) *DescribeAccountDelegatedStatusResponseBody {
	s.DelegatedStatus = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponseBody) SetRequestId(v string) *DescribeAccountDelegatedStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAccountDelegatedStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountDelegatedStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountDelegatedStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountDelegatedStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountDelegatedStatusResponse) SetHeaders(v map[string]*string) *DescribeAccountDelegatedStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountDelegatedStatusResponse) SetStatusCode(v int32) *DescribeAccountDelegatedStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountDelegatedStatusResponse) SetBody(v *DescribeAccountDelegatedStatusResponseBody) *DescribeAccountDelegatedStatusResponse {
	s.Body = v
	return s
}

type DescribeCertDetailRequest struct {
	// The ID of the certificate.
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCertDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailRequest) SetCertIdentifier(v string) *DescribeCertDetailRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertDetailRequest) SetInstanceId(v string) *DescribeCertDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCertDetailRequest) SetRegionId(v string) *DescribeCertDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCertDetailRequest) SetResourceManagerResourceGroupId(v string) *DescribeCertDetailRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeCertDetailResponseBody struct {
	// The details of the certificate.
	CertDetail *DescribeCertDetailResponseBodyCertDetail `json:"CertDetail,omitempty" xml:"CertDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailResponseBody) SetCertDetail(v *DescribeCertDetailResponseBodyCertDetail) *DescribeCertDetailResponseBody {
	s.CertDetail = v
	return s
}

func (s *DescribeCertDetailResponseBody) SetRequestId(v string) *DescribeCertDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCertDetailResponseBodyCertDetail struct {
	// The time when the certificate expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The time when the certificate was issued. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The ID of the certificate.
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The primary domain name, which is a common name.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain name that is associated with the certificate.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The other domain names that are associated with the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
}

func (s DescribeCertDetailResponseBodyCertDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertDetailResponseBodyCertDetail) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetAfterDate(v int64) *DescribeCertDetailResponseBodyCertDetail {
	s.AfterDate = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetBeforeDate(v int64) *DescribeCertDetailResponseBodyCertDetail {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetCertIdentifier(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetCertName(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.CertName = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetCommonName(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.CommonName = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetDomain(v string) *DescribeCertDetailResponseBodyCertDetail {
	s.Domain = &v
	return s
}

func (s *DescribeCertDetailResponseBodyCertDetail) SetSans(v []*string) *DescribeCertDetailResponseBodyCertDetail {
	s.Sans = v
	return s
}

type DescribeCertDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailResponse) SetHeaders(v map[string]*string) *DescribeCertDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertDetailResponse) SetStatusCode(v int32) *DescribeCertDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertDetailResponse) SetBody(v *DescribeCertDetailResponseBody) *DescribeCertDetailResponse {
	s.Body = v
	return s
}

type DescribeCertsRequest struct {
	// The type of the encryption algorithm. Valid values:
	//
	// *   **NotSM2**: The encryption algorithm is not the SM2 algorithm. This is the default value.
	// *   **SM2**: The encryption algorithm is the SM2 algorithm.
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The domain name.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: **10**.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertsRequest) SetAlgorithm(v string) *DescribeCertsRequest {
	s.Algorithm = &v
	return s
}

func (s *DescribeCertsRequest) SetDomain(v string) *DescribeCertsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCertsRequest) SetInstanceId(v string) *DescribeCertsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCertsRequest) SetPageNumber(v int64) *DescribeCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCertsRequest) SetPageSize(v int64) *DescribeCertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCertsRequest) SetRegionId(v string) *DescribeCertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCertsRequest) SetResourceManagerResourceGroupId(v string) *DescribeCertsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeCertsResponseBody struct {
	// The certificates.
	Certs []*DescribeCertsResponseBodyCerts `json:"Certs,omitempty" xml:"Certs,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponseBody) SetCerts(v []*DescribeCertsResponseBodyCerts) *DescribeCertsResponseBody {
	s.Certs = v
	return s
}

func (s *DescribeCertsResponseBody) SetRequestId(v string) *DescribeCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertsResponseBody) SetTotalCount(v int64) *DescribeCertsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCertsResponseBodyCerts struct {
	// The time when the certificate becomes valid.
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The time when the certificate expires.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The globally unique ID of the certificate. The value is in the "Certificate ID-cn-hangzhou" format. For example, if the ID of the certificate is 123, the value of CertIdentifier is 123-cn-hangzhou.
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain name that is added to WAF.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the certificate chain is complete. Valid values:
	//
	// *   **true**
	// *   **false**
	IsChainCompleted *bool `json:"IsChainCompleted,omitempty" xml:"IsChainCompleted,omitempty"`
}

func (s DescribeCertsResponseBodyCerts) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertsResponseBodyCerts) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponseBodyCerts) SetAfterDate(v int64) *DescribeCertsResponseBodyCerts {
	s.AfterDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetBeforeDate(v int64) *DescribeCertsResponseBodyCerts {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCertIdentifier(v string) *DescribeCertsResponseBodyCerts {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCertName(v string) *DescribeCertsResponseBodyCerts {
	s.CertName = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCommonName(v string) *DescribeCertsResponseBodyCerts {
	s.CommonName = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetDomain(v string) *DescribeCertsResponseBodyCerts {
	s.Domain = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetIsChainCompleted(v bool) *DescribeCertsResponseBodyCerts {
	s.IsChainCompleted = &v
	return s
}

type DescribeCertsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponse) SetHeaders(v map[string]*string) *DescribeCertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertsResponse) SetStatusCode(v int32) *DescribeCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertsResponse) SetBody(v *DescribeCertsResponseBody) *DescribeCertsResponse {
	s.Body = v
	return s
}

type DescribeCloudResourcesRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The domain name. You can use this parameter if you set ResourceProduct to fc or sae.
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The function name. You can use this parameter if you set ResourceProduct to fc.
	ResourceFunction *string `json:"ResourceFunction,omitempty" xml:"ResourceFunction,omitempty"`
	// The ID of the resource.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The cloud service to which the resource belongs. Valid values:
	//
	// *   **alb**: Application Load Balancer (ALB).
	// *   **mse**: Microservices Engine (MSE).
	// *   **fc**: Function Compute.
	// *   **sae**: Serverless App Engine (SAE).
	//
	// >  Different cloud services are available in different regions. The specified cloud service must be available in the specified region.
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the resource. For information about region IDs, see the following table.
	//
	// >  Different cloud services are available in different regions. The specified cloud service must be available in the specified region.
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The route name. You can use this parameter if you set ResourceProduct to mse.
	ResourceRouteName *string `json:"ResourceRouteName,omitempty" xml:"ResourceRouteName,omitempty"`
}

func (s DescribeCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesRequest) SetInstanceId(v string) *DescribeCloudResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetOwnerUserId(v string) *DescribeCloudResourcesRequest {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetPageNumber(v int64) *DescribeCloudResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetPageSize(v int64) *DescribeCloudResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetRegionId(v string) *DescribeCloudResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceDomain(v string) *DescribeCloudResourcesRequest {
	s.ResourceDomain = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceFunction(v string) *DescribeCloudResourcesRequest {
	s.ResourceFunction = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceInstanceId(v string) *DescribeCloudResourcesRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeCloudResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceName(v string) *DescribeCloudResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceProduct(v string) *DescribeCloudResourcesRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceRegionId(v string) *DescribeCloudResourcesRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeCloudResourcesRequest) SetResourceRouteName(v string) *DescribeCloudResourcesRequest {
	s.ResourceRouteName = &v
	return s
}

type DescribeCloudResourcesResponseBody struct {
	// The cloud service resources that are added to WAF.
	CloudResources []*DescribeCloudResourcesResponseBodyCloudResources `json:"CloudResources,omitempty" xml:"CloudResources,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud service resources returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesResponseBody) SetCloudResources(v []*DescribeCloudResourcesResponseBodyCloudResources) *DescribeCloudResourcesResponseBody {
	s.CloudResources = v
	return s
}

func (s *DescribeCloudResourcesResponseBody) SetRequestId(v string) *DescribeCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBody) SetTotalCount(v int64) *DescribeCloudResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCloudResourcesResponseBodyCloudResources struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The domain name. This parameter has a value only if the value of ResourceProduct is fc or sae.
	ResourceDomain *string `json:"ResourceDomain,omitempty" xml:"ResourceDomain,omitempty"`
	// The function name. This parameter has a value only if the value of ResourceProduct is fc.
	ResourceFunction *string `json:"ResourceFunction,omitempty" xml:"ResourceFunction,omitempty"`
	// The ID of the resource.
	ResourceInstance *string `json:"ResourceInstance,omitempty" xml:"ResourceInstance,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The cloud service to which the resource belongs. Valid values:
	//
	// *   **alb**: ALB.
	// *   **mse**: MSE.
	// *   **fc**: Function Compute.
	// *   **sae**: SAE.
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the resource.
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The route name. This parameter has a value only if the value of ResourceProduct is mse.
	ResourceRouteName *string `json:"ResourceRouteName,omitempty" xml:"ResourceRouteName,omitempty"`
	// The service name. This parameter has a value only if the value of ResourceProduct is fc.
	ResourceService *string `json:"ResourceService,omitempty" xml:"ResourceService,omitempty"`
}

func (s DescribeCloudResourcesResponseBodyCloudResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourcesResponseBodyCloudResources) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetOwnerUserId(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceDomain(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceDomain = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceFunction(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceFunction = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceInstance(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceInstance = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceName(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceName = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceProduct(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceRegionId(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceRouteName(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceRouteName = &v
	return s
}

func (s *DescribeCloudResourcesResponseBodyCloudResources) SetResourceService(v string) *DescribeCloudResourcesResponseBodyCloudResources {
	s.ResourceService = &v
	return s
}

type DescribeCloudResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesResponse) SetHeaders(v map[string]*string) *DescribeCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourcesResponse) SetStatusCode(v int32) *DescribeCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourcesResponse) SetBody(v *DescribeCloudResourcesResponseBody) *DescribeCloudResourcesResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object that you want to query. Only exact queries are supported.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceRequest) SetInstanceId(v string) *DescribeDefenseResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceRequest) SetRegionId(v string) *DescribeDefenseResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceRequest) SetResource(v string) *DescribeDefenseResourceRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeDefenseResourceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the protected object.
	Resource *DescribeDefenseResourceResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s DescribeDefenseResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponseBody) SetRequestId(v string) *DescribeDefenseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceResponseBody) SetResource(v *DescribeDefenseResourceResponseBodyResource) *DescribeDefenseResourceResponseBody {
	s.Resource = v
	return s
}

type DescribeDefenseResourceResponseBodyResource struct {
	// The status of the tracking cookie.
	//
	// *   **0**: disabled.
	// *   **1**: enabled.
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the secure attribute of the tracking cookie.
	//
	// *   **0**: disabled.
	// *   **1**: enabled.
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the secure attribute of the slider CAPTCHA cookie.
	//
	// *   **0**: disabled.
	// *   **1**: enabled.
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The custom header fields.
	//
	// >  If the value of XffStatus is 1, the first IP address in the specified header field is used as the originating IP address of the client to prevent X-Forwarded-For (XFF) forgery. If you specify multiple header fields, WAF reads the values of the header fields in sequence until the originating IP address is obtained. If the originating IP address cannot be obtained, the first IP address in the XFF header field is used as the originating IP address of the client.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	// The description of the protected object.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the protected object. Different key-value pairs indicate different attributes of the protected object.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the protected object was created. Unit: milliseconds.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the protected object was modified. Unit: milliseconds.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user ID (UID) of the Alibaba Cloud account to which the protected object belongs.
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The pattern used for the protected object.
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The name of the cloud service.
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The name of the protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the protected object group to which the protected object belongs.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The origin of the protected object. Valid values:
	//
	// *   **custom**
	// *   **access**
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	// Indicates whether a Layer 7 proxy is deployed in front of WAF, such as Anti-DDoS Proxy and Alibaba Cloud CDN. Valid values:
	//
	// *   **0**: No Layer 7 proxy is deployed.
	// *   **1**: A Layer 7 proxy is deployed.
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s DescribeDefenseResourceResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponseBodyResource) SetAcwCookieStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.AcwCookieStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetAcwSecureStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.AcwSecureStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetAcwV3SecureStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetCustomHeaders(v []*string) *DescribeDefenseResourceResponseBodyResource {
	s.CustomHeaders = v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetDescription(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetDetail(v map[string]interface{}) *DescribeDefenseResourceResponseBodyResource {
	s.Detail = v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetGmtCreate(v int64) *DescribeDefenseResourceResponseBodyResource {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetGmtModified(v int64) *DescribeDefenseResourceResponseBodyResource {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetOwnerUserId(v string) *DescribeDefenseResourceResponseBodyResource {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetPattern(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Pattern = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetProduct(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Product = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResource(v string) *DescribeDefenseResourceResponseBodyResource {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResourceGroup(v string) *DescribeDefenseResourceResponseBodyResource {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceResponseBodyResource {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetResourceOrigin(v string) *DescribeDefenseResourceResponseBodyResource {
	s.ResourceOrigin = &v
	return s
}

func (s *DescribeDefenseResourceResponseBodyResource) SetXffStatus(v int32) *DescribeDefenseResourceResponseBodyResource {
	s.XffStatus = &v
	return s
}

type DescribeDefenseResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceResponse) SetStatusCode(v int32) *DescribeDefenseResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceResponse) SetBody(v *DescribeDefenseResourceResponseBody) *DescribeDefenseResourceResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceGroupRequest struct {
	// The name of the protected object group whose information you want to query.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupRequest) SetGroupName(v string) *DescribeDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetRegionId(v string) *DescribeDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeDefenseResourceGroupResponseBody struct {
	// The information about the protected object group.
	Group *DescribeDefenseResourceGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponseBody) SetGroup(v *DescribeDefenseResourceGroupResponseBodyGroup) *DescribeDefenseResourceGroupResponseBody {
	s.Group = v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDefenseResourceGroupResponseBodyGroup struct {
	// The description of the protected object group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protected object group was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The most recent time when the protected object group was modified.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the protected object group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The protected objects in the protected object group. The protected objects are separated with commas (,).
	ResourceList *string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty"`
}

func (s DescribeDefenseResourceGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetDescription(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGmtCreate(v int64) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGmtModified(v int64) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGroupName(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetResourceList(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.ResourceList = &v
	return s
}

type DescribeDefenseResourceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) SetBody(v *DescribeDefenseResourceGroupResponseBody) *DescribeDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceGroupNamesRequest struct {
	// The name of the protected object group. Fuzzy queries are supported.
	GroupNameLike *string `json:"GroupNameLike,omitempty" xml:"GroupNameLike,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceGroupNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetGroupNameLike(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.GroupNameLike = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetPageNumber(v int32) *DescribeDefenseResourceGroupNamesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetPageSize(v int32) *DescribeDefenseResourceGroupNamesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetRegionId(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupNamesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeDefenseResourceGroupNamesResponseBody struct {
	// The names of the protected object groups.
	GroupNames []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourceGroupNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) SetGroupNames(v []*string) *DescribeDefenseResourceGroupNamesResponseBody {
	s.GroupNames = v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourceGroupNamesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseResourceGroupNamesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceGroupNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceGroupNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupNamesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupNamesResponse) SetBody(v *DescribeDefenseResourceGroupNamesResponseBody) *DescribeDefenseResourceGroupNamesResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceGroupsRequest struct {
	// The name of the protected object group that you want to query. Fuzzy queries are supported.
	GroupNameLike *string `json:"GroupNameLike,omitempty" xml:"GroupNameLike,omitempty"`
	// The names of the protected object groups that you want to query. Separate multiple names with commas (,).
	GroupNames *string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsRequest) SetGroupNameLike(v string) *DescribeDefenseResourceGroupsRequest {
	s.GroupNameLike = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetGroupNames(v string) *DescribeDefenseResourceGroupsRequest {
	s.GroupNames = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetPageNumber(v int32) *DescribeDefenseResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetPageSize(v int32) *DescribeDefenseResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetRegionId(v string) *DescribeDefenseResourceGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeDefenseResourceGroupsResponseBody struct {
	// The list of protected object groups.
	Groups []*DescribeDefenseResourceGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsResponseBody) SetGroups(v []*DescribeDefenseResourceGroupsResponseBodyGroups) *DescribeDefenseResourceGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBody) SetTotalCount(v int64) *DescribeDefenseResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseResourceGroupsResponseBodyGroups struct {
	// The description of the protected object group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protected object group was created. Unit: milliseconds.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The most recent time when the protected object group was modified. Unit: milliseconds.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the protected object group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The names of the protected objects that are added to the protected object group. Separate multiple protected objects with commas (,).
	ResourceList *string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty"`
}

func (s DescribeDefenseResourceGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetDescription(v string) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetGmtCreate(v int64) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetGmtModified(v int64) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetGroupName(v string) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponseBodyGroups) SetResourceList(v string) *DescribeDefenseResourceGroupsResponseBodyGroups {
	s.ResourceList = &v
	return s
}

type DescribeDefenseResourceGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupsResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponse) SetBody(v *DescribeDefenseResourceGroupsResponseBody) *DescribeDefenseResourceGroupsResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceNamesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object that you want to query.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceNamesRequest) SetInstanceId(v string) *DescribeDefenseResourceNamesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetPageNumber(v int32) *DescribeDefenseResourceNamesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetPageSize(v int32) *DescribeDefenseResourceNamesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetRegionId(v string) *DescribeDefenseResourceNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetResource(v string) *DescribeDefenseResourceNamesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceNamesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceNamesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeDefenseResourceNamesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The names of the protected objects.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourceNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceNamesResponseBody) SetRequestId(v string) *DescribeDefenseResourceNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceNamesResponseBody) SetResources(v []*string) *DescribeDefenseResourceNamesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseResourceNamesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourceNamesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseResourceNamesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceNamesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceNamesResponse) SetStatusCode(v int32) *DescribeDefenseResourceNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceNamesResponse) SetBody(v *DescribeDefenseResourceNamesResponseBody) *DescribeDefenseResourceNamesResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceTemplatesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object or protected object group that you want to query.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protected resource. Valid values:
	//
	// *   **single**: protected object. This is the default value.
	// *   **group**: protected object group.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the protection rule.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The type of the protection rule. Valid values:
	//
	// *   **defense**: defense rule. This is the default value.
	// *   **whitelist**: whitelist rule.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeDefenseResourceTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesRequest) SetInstanceId(v string) *DescribeDefenseResourceTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRegionId(v string) *DescribeDefenseResourceTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetResource(v string) *DescribeDefenseResourceTemplatesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceTemplatesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetResourceType(v string) *DescribeDefenseResourceTemplatesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRuleId(v int64) *DescribeDefenseResourceTemplatesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRuleType(v string) *DescribeDefenseResourceTemplatesRequest {
	s.RuleType = &v
	return s
}

type DescribeDefenseResourceTemplatesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protection templates.
	Templates []*DescribeDefenseResourceTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeDefenseResourceTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesResponseBody) SetRequestId(v string) *DescribeDefenseResourceTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBody) SetTemplates(v []*DescribeDefenseResourceTemplatesResponseBodyTemplates) *DescribeDefenseResourceTemplatesResponseBody {
	s.Templates = v
	return s
}

type DescribeDefenseResourceTemplatesResponseBodyTemplates struct {
	// The scenario in which the protection template is used.
	//
	// *   **waf_group**: basic protection.
	// *   **antiscan**: scan protection.
	// *   **ip_blacklist**: IP address blacklist.
	// *   **custom_acl**: custom rule.
	// *   **whitelist**: whitelist.
	// *   **region_block**: region blacklist.
	// *   **custom_response**: custom response.
	// *   **cc**: HTTP flood protection.
	// *   **tamperproof**: website tamper-proofing.
	// *   **dlp**: data leakage prevention.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the template is used. Valid values:
	//
	// *   **web**: bot management for website protection.
	// *   **app**: bot management for app protection.
	// *   **basic**: bot management for basic protection.
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protection template was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection template. The value custom indicates that the template is a custom template created by the user.
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection template. Valid values:
	//
	// *   **0**: disabled.
	// *   **1**: enabled.
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection template. Valid values:
	//
	// *   **user_default**: default template.
	// *   **user_custom**: custom template.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseResourceTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetDefenseScene(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetDefenseSubScene(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetGmtModified(v int64) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateId(v int64) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateName(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateOrigin(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateStatus(v int32) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeDefenseResourceTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

type DescribeDefenseResourceTemplatesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponse) SetStatusCode(v int32) *DescribeDefenseResourceTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesResponse) SetBody(v *DescribeDefenseResourceTemplatesResponseBody) *DescribeDefenseResourceTemplatesResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourcesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. Specify the value of this parameter as a string in the JSON format.
	//
	// >  The results vary based on the query condition. For more information, see the "**Query parameters**" section in this topic.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tags of the resources that you want to query. You can specify up to 20 tags.
	Tag []*DescribeDefenseResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDefenseResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesRequest) SetInstanceId(v string) *DescribeDefenseResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetPageNumber(v int32) *DescribeDefenseResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetPageSize(v int32) *DescribeDefenseResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetQuery(v string) *DescribeDefenseResourcesRequest {
	s.Query = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetRegionId(v string) *DescribeDefenseResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetTag(v []*DescribeDefenseResourcesRequestTag) *DescribeDefenseResourcesRequest {
	s.Tag = v
	return s
}

type DescribeDefenseResourcesRequestTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDefenseResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesRequestTag) SetKey(v string) *DescribeDefenseResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDefenseResourcesRequestTag) SetValue(v string) *DescribeDefenseResourcesRequestTag {
	s.Value = &v
	return s
}

type DescribeDefenseResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protected objects.
	Resources []*DescribeDefenseResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBody) SetRequestId(v string) *DescribeDefenseResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) SetResources(v []*DescribeDefenseResourcesResponseBodyResources) *DescribeDefenseResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseResourcesResponseBodyResources struct {
	// The status of the tracking cookie.
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the secure attribute in the tracking cookie.
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the secure attribute in the slider CAPTCHA cookie.
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The custom XFF headers that are used to identify the originating IP addresses of clients. If the value of XffStatus is 1 and CustomHeaders is left empty, the first IP addresses in the XFF headers are used as the originating IP addresses of clients.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	// The description of the protected object.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the protected object. Different key-value pairs indicate different attributes of the protected object.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the protected object was created. Unit: milliseconds.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the protected object was modified. Unit: milliseconds.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The pattern in which the protected object is protected.
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The name of the cloud service.
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The name of the protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the protected object group to which the protected object belongs.
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The origin of the protected object.
	ResourceOrigin *string `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	// Indicates whether the X-Forwarded-For (XFF) proxy is enabled.
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s DescribeDefenseResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetAcwCookieStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.AcwCookieStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetAcwSecureStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.AcwSecureStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetAcwV3SecureStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetCustomHeaders(v []*string) *DescribeDefenseResourcesResponseBodyResources {
	s.CustomHeaders = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetDescription(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetDetail(v map[string]interface{}) *DescribeDefenseResourcesResponseBodyResources {
	s.Detail = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetGmtCreate(v int64) *DescribeDefenseResourcesResponseBodyResources {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetGmtModified(v int64) *DescribeDefenseResourcesResponseBodyResources {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetOwnerUserId(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetPattern(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Pattern = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetProduct(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Product = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResource(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceGroup(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceOrigin(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceOrigin = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetXffStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.XffStatus = &v
	return s
}

type DescribeDefenseResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourcesResponse) SetStatusCode(v int32) *DescribeDefenseResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourcesResponse) SetBody(v *DescribeDefenseResourcesResponseBody) *DescribeDefenseResourcesResponse {
	s.Body = v
	return s
}

type DescribeDefenseRuleRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule that you want to query.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the protection rule template to which the protection rule that you want to query belongs.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleRequest) SetInstanceId(v string) *DescribeDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetRegionId(v string) *DescribeDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetRuleId(v int64) *DescribeDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetTemplateId(v int64) *DescribeDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type DescribeDefenseRuleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the protection rule. The value is a JSON string that contains multiple parameters.
	Rule *DescribeDefenseRuleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s DescribeDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponseBody) SetRequestId(v string) *DescribeDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBody) SetRule(v *DescribeDefenseRuleResponseBodyRule) *DescribeDefenseRuleResponseBody {
	s.Rule = v
	return s
}

type DescribeDefenseRuleResponseBodyRule struct {
	// The details of the protection rule. The value is a JSON string that contains multiple parameters. For more information, see the "**Protection rule parameters**" section of the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The origin of the protection rule. Valid values:
	//
	// *   **custom:** The protection rule is created by the user.
	// *   **system:** The protection rule is automatically generated by the system.
	DefenseOrigin *string `json:"DefenseOrigin,omitempty" xml:"DefenseOrigin,omitempty"`
	// The scenario in which the protection rule is used. For more information, see the description of **DefenseScene** in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The most recent time when the protection rule was modified.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection rule.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the protection rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the protection rule. Valid values:
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRuleResponseBodyRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponseBodyRule) SetConfig(v string) *DescribeDefenseRuleResponseBodyRule {
	s.Config = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseOrigin(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseOrigin = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseScene(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetGmtModified(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetRuleId(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetRuleName(v string) *DescribeDefenseRuleResponseBodyRule {
	s.RuleName = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetStatus(v int32) *DescribeDefenseRuleResponseBodyRule {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetTemplateId(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.TemplateId = &v
	return s
}

type DescribeDefenseRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponse) SetHeaders(v map[string]*string) *DescribeDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRuleResponse) SetStatusCode(v int32) *DescribeDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRuleResponse) SetBody(v *DescribeDefenseRuleResponseBody) *DescribeDefenseRuleResponse {
	s.Body = v
	return s
}

type DescribeDefenseRulesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. Specify a string that contains multiple parameters in the JSON format.
	//
	// >  The results vary based on the query conditions. For more information, see the "**Query parameters**" section in this topic.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of protection rule that you want to query. Valid values:
	//
	// *   **whitelist:** whitelist rule.
	// *   **defense:** defense rule. This is the default value.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeDefenseRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesRequest) SetInstanceId(v string) *DescribeDefenseRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetPageNumber(v int32) *DescribeDefenseRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetPageSize(v int32) *DescribeDefenseRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetQuery(v string) *DescribeDefenseRulesRequest {
	s.Query = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetRegionId(v string) *DescribeDefenseRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseRulesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetRuleType(v string) *DescribeDefenseRulesRequest {
	s.RuleType = &v
	return s
}

type DescribeDefenseRulesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of protection rules.
	Rules []*DescribeDefenseRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of returned entries.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponseBody) SetRequestId(v string) *DescribeDefenseRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRulesResponseBody) SetRules(v []*DescribeDefenseRulesResponseBodyRules) *DescribeDefenseRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeDefenseRulesResponseBody) SetTotalCount(v int64) *DescribeDefenseRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseRulesResponseBodyRules struct {
	// The details of the protection rule. The value is a string that contains multiple parameters in the JSON format. For more information, see the "**Rule parameters**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The origin of the protection rule. Valid values:
	//
	// *   **custom:** The protection rule is created by the user.
	// *   **system:** The protection rule is automatically generated by the system.
	DefenseOrigin *string `json:"DefenseOrigin,omitempty" xml:"DefenseOrigin,omitempty"`
	// The scenario in which the protection rule is used. For more information, see the description of the **DefenseScene** parameter in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The most recent time when the protection rule was modified.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection rule.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the protection rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the protection rule. Valid values:
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponseBodyRules) SetConfig(v string) *DescribeDefenseRulesResponseBodyRules {
	s.Config = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetDefenseOrigin(v string) *DescribeDefenseRulesResponseBodyRules {
	s.DefenseOrigin = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetDefenseScene(v string) *DescribeDefenseRulesResponseBodyRules {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetGmtModified(v int64) *DescribeDefenseRulesResponseBodyRules {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetRuleId(v int64) *DescribeDefenseRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetRuleName(v string) *DescribeDefenseRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetStatus(v int32) *DescribeDefenseRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetTemplateId(v int64) *DescribeDefenseRulesResponseBodyRules {
	s.TemplateId = &v
	return s
}

type DescribeDefenseRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponse) SetHeaders(v map[string]*string) *DescribeDefenseRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRulesResponse) SetStatusCode(v int32) *DescribeDefenseRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRulesResponse) SetBody(v *DescribeDefenseRulesResponseBody) *DescribeDefenseRulesResponse {
	s.Body = v
	return s
}

type DescribeDefenseTemplateRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateRequest) SetInstanceId(v string) *DescribeDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetRegionId(v string) *DescribeDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetTemplateId(v int64) *DescribeDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

type DescribeDefenseTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the template.
	Template *DescribeDefenseTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s DescribeDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponseBody) SetRequestId(v string) *DescribeDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBody) SetTemplate(v *DescribeDefenseTemplateResponseBodyTemplate) *DescribeDefenseTemplateResponseBody {
	s.Template = v
	return s
}

type DescribeDefenseTemplateResponseBodyTemplate struct {
	// The scenario in which the template is used. For more information, see the description of the **DefenseScene** parameter in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the template is used. Valid values:
	//
	// *   **web**: The template is a bot management template that is used for website protection.
	// *   **app**: The template is a bot management template that is used for app protection.
	// *   **basic**: The template is a bot management template that is used for basic protection.
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection rule template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The most recent time when the protection rule template was modified.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection rule template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection rule template. If the value of this parameter is custom, the protection rule template is created by the user.
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection rule template. Valid values:
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection rule template. Valid values:
	//
	// *   **user_default:** default template.
	// *   **user_custom:** custom template.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDefenseScene(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDefenseSubScene(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDescription(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetGmtModified(v int64) *DescribeDefenseTemplateResponseBodyTemplate {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateId(v int64) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateName(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateOrigin(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateStatus(v int32) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateType(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateType = &v
	return s
}

type DescribeDefenseTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplateResponse) SetStatusCode(v int32) *DescribeDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplateResponse) SetBody(v *DescribeDefenseTemplateResponseBody) *DescribeDefenseTemplateResponse {
	s.Body = v
	return s
}

type DescribeDefenseTemplateValidGroupsRequest struct {
	// The scenario in which the protection template is used.
	//
	// *   **waf_group**: basic protection.
	// *   **antiscan**: scan protection.
	// *   **ip_blacklist**: IP address blacklist.
	// *   **custom_acl**: custom rule.
	// *   **whitelist**: whitelist.
	// *   **region_block**: region blacklist.
	// *   **custom_response**: custom response.
	// *   **cc**: HTTP flood protection.
	// *   **tamperproof**: website tamper-proofing.
	// *   **dlp**: data leakage prevention.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The name of the protected object group that you want to query.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateValidGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateValidGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetDefenseScene(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetGroupName(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetInstanceId(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetPageNumber(v int32) *DescribeDefenseTemplateValidGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetPageSize(v int32) *DescribeDefenseTemplateValidGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetRegionId(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsRequest) SetTemplateId(v int64) *DescribeDefenseTemplateValidGroupsRequest {
	s.TemplateId = &v
	return s
}

type DescribeDefenseTemplateValidGroupsResponseBody struct {
	// The names of the protected object groups.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseTemplateValidGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateValidGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) SetGroups(v []*string) *DescribeDefenseTemplateValidGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) SetRequestId(v string) *DescribeDefenseTemplateValidGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponseBody) SetTotalCount(v int64) *DescribeDefenseTemplateValidGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseTemplateValidGroupsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplateValidGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplateValidGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateValidGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidGroupsResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplateValidGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponse) SetStatusCode(v int32) *DescribeDefenseTemplateValidGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponse) SetBody(v *DescribeDefenseTemplateValidGroupsResponseBody) *DescribeDefenseTemplateValidGroupsResponse {
	s.Body = v
	return s
}

type DescribeDefenseTemplatesRequest struct {
	// The scenario in which the protection template is used.
	//
	// *   **waf_group**: basic protection.
	// *   **antiscan**: scan protection.
	// *   **ip_blacklist**: IP address blacklist.
	// *   **custom_acl**: custom rule.
	// *   **whitelist**: whitelist.
	// *   **region_block**: region blacklist.
	// *   **custom_response**: custom response.
	// *   **cc**: HTTP flood protection.
	// *   **tamperproof**: website tamper-proofing.
	// *   **dlp**: data leakage prevention.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the protection template is used. Valid values:
	//
	// *   **web**: bot management for website protection.
	// *   **app**: bot management for app protection.
	// *   **basic**: bot management for basic protection.
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object or protected object group.
	//
	// >  If you specify ResourceType, you must specify this parameter.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protected resource. Valid values:
	//
	// *   **single**: protected object. This is the default value.
	// *   **group**: protected object group.
	//
	// >  If you specify Resource, you must specify this parameter.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the protection template. Valid values:
	//
	// *   **user_default**: default template.
	// *   **user_custom**: custom template.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesRequest) SetDefenseScene(v string) *DescribeDefenseTemplatesRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetDefenseSubScene(v string) *DescribeDefenseTemplatesRequest {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetInstanceId(v string) *DescribeDefenseTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetPageNumber(v int32) *DescribeDefenseTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetPageSize(v int32) *DescribeDefenseTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetRegionId(v string) *DescribeDefenseTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetResource(v string) *DescribeDefenseTemplatesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplatesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetResourceType(v string) *DescribeDefenseTemplatesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetTemplateId(v int64) *DescribeDefenseTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplatesRequest) SetTemplateType(v string) *DescribeDefenseTemplatesRequest {
	s.TemplateType = &v
	return s
}

type DescribeDefenseTemplatesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The protection templates.
	Templates []*DescribeDefenseTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesResponseBody) SetRequestId(v string) *DescribeDefenseTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBody) SetTemplates(v []*DescribeDefenseTemplatesResponseBodyTemplates) *DescribeDefenseTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeDefenseTemplatesResponseBody) SetTotalCount(v int64) *DescribeDefenseTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseTemplatesResponseBodyTemplates struct {
	// The scenario in which the protection template is used.
	//
	// *   **waf_group**: basic protection.
	// *   **antiscan**: scan protection.
	// *   **ip_blacklist**: IP address blacklist.
	// *   **custom_acl**: custom rule.
	// *   **whitelist**: whitelist.
	// *   **region_block**: region blacklist.
	// *   **custom_response**: custom response.
	// *   **cc**: HTTP flood protection.
	// *   **tamperproof**: website tamper-proofing.
	// *   **dlp**: data leakage prevention.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The sub-scenario in which the protection template is used. Valid values:
	//
	// *   **web**: bot management for website protection.
	// *   **app**: bot management for app protection.
	// *   **basic**: bot management for basic protection.
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the protection template was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection template. The value custom indicates that the protection template is a custom template created by the user.
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection template. Valid values:
	//
	// *   **0**: disabled.
	// *   **1**: enabled.
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection template. Valid values:
	//
	// *   **user_default**: default template.
	// *   **user_custom**: custom template.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetDefenseScene(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetDefenseSubScene(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.DefenseSubScene = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetGmtModified(v int64) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateId(v int64) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateName(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateOrigin(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateStatus(v int32) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseTemplatesResponseBodyTemplates) SetTemplateType(v string) *DescribeDefenseTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

type DescribeDefenseTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplatesResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplatesResponse) SetStatusCode(v int32) *DescribeDefenseTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplatesResponse) SetBody(v *DescribeDefenseTemplatesResponseBody) *DescribeDefenseTemplatesResponse {
	s.Body = v
	return s
}

type DescribeDomainDNSRecordRequest struct {
	// The domain name whose DNS settings you want to check.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDomainDNSRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDNSRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDNSRecordRequest) SetDomain(v string) *DescribeDomainDNSRecordRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) SetInstanceId(v string) *DescribeDomainDNSRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) SetRegionId(v string) *DescribeDomainDNSRecordRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) SetResourceManagerResourceGroupId(v string) *DescribeDomainDNSRecordRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeDomainDNSRecordResponseBody struct {
	// The status of the DNS settings. Valid values:
	//
	// *   **cnameMatched**: The DNS settings are properly configured.
	// *   **vipMatched**: An A record maps the domain name to the WAF virtual IP address (VIP).
	// *   **wafVip**: An A record maps the domain name to another WAF VIP.
	// *   **unRecord**: The domain name does not have a DNS record.
	// *   **unUsed**: The domain name is not pointed to WAF.
	// *   **checkTimeout**: The check times out.
	DNSStatus *string `json:"DNSStatus,omitempty" xml:"DNSStatus,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainDNSRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDNSRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDNSRecordResponseBody) SetDNSStatus(v string) *DescribeDomainDNSRecordResponseBody {
	s.DNSStatus = &v
	return s
}

func (s *DescribeDomainDNSRecordResponseBody) SetRequestId(v string) *DescribeDomainDNSRecordResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainDNSRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainDNSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainDNSRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDNSRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDNSRecordResponse) SetHeaders(v map[string]*string) *DescribeDomainDNSRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDNSRecordResponse) SetStatusCode(v int32) *DescribeDomainDNSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDNSRecordResponse) SetBody(v *DescribeDomainDNSRecordResponseBody) *DescribeDomainDNSRecordResponse {
	s.Body = v
	return s
}

type DescribeDomainDetailRequest struct {
	// The domain name that you want to query.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailRequest) SetDomain(v string) *DescribeDomainDetailRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetInstanceId(v string) *DescribeDomainDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetRegionId(v string) *DescribeDomainDetailRequest {
	s.RegionId = &v
	return s
}

type DescribeDomainDetailResponseBody struct {
	// The details of the SSL certificate.
	CertDetail *DescribeDomainDetailResponseBodyCertDetail `json:"CertDetail,omitempty" xml:"CertDetail,omitempty" type:"Struct"`
	// The CNAME that is assigned by WAF to the domain name.
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The configurations of the listeners.
	Listen *DescribeDomainDetailResponseBodyListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The configurations of the forwarding rule.
	Redirect *DescribeDomainDetailResponseBodyRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The information about the SM certificate.
	SM2CertDetail *DescribeDomainDetailResponseBodySM2CertDetail `json:"SM2CertDetail,omitempty" xml:"SM2CertDetail,omitempty" type:"Struct"`
	// The status of the domain name. Valid values:
	//
	// *   **1:** The domain name is in a normal state.
	// *   **2:** The domain name is being created.
	// *   **3:** The domain name is being modified.
	// *   **4:** The domain name is being released.
	// *   **5:** WAF no longer forwards traffic of the domain name.
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBody) SetCertDetail(v *DescribeDomainDetailResponseBodyCertDetail) *DescribeDomainDetailResponseBody {
	s.CertDetail = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetCname(v string) *DescribeDomainDetailResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomain(v string) *DescribeDomainDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetListen(v *DescribeDomainDetailResponseBodyListen) *DescribeDomainDetailResponseBody {
	s.Listen = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRedirect(v *DescribeDomainDetailResponseBodyRedirect) *DescribeDomainDetailResponseBody {
	s.Redirect = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRequestId(v string) *DescribeDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetResourceManagerResourceGroupId(v string) *DescribeDomainDetailResponseBody {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetSM2CertDetail(v *DescribeDomainDetailResponseBodySM2CertDetail) *DescribeDomainDetailResponseBody {
	s.SM2CertDetail = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetStatus(v int64) *DescribeDomainDetailResponseBody {
	s.Status = &v
	return s
}

type DescribeDomainDetailResponseBodyCertDetail struct {
	// The domain name of your website.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The end of the validity period of the SSL certificate. The value is in the UNIX timestamp format. Unit: milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the SSL certificate.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the SSL certificate.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// All domain names that are bound to the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	// The beginning of the validity period of the SSL certificate. The value is in the UNIX timestamp format. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainDetailResponseBodyCertDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyCertDetail) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetCommonName(v string) *DescribeDomainDetailResponseBodyCertDetail {
	s.CommonName = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetEndTime(v int64) *DescribeDomainDetailResponseBodyCertDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetId(v string) *DescribeDomainDetailResponseBodyCertDetail {
	s.Id = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetName(v string) *DescribeDomainDetailResponseBodyCertDetail {
	s.Name = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetSans(v []*string) *DescribeDomainDetailResponseBodyCertDetail {
	s.Sans = v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetStartTime(v int64) *DescribeDomainDetailResponseBodyCertDetail {
	s.StartTime = &v
	return s
}

type DescribeDomainDetailResponseBodyListen struct {
	// The ID of the certificate.
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suites. Valid values:
	//
	// *   **1:** all cipher suites.
	// *   **2:** strong cipher suites.
	// *   **99:** custom cipher suites.
	CipherSuite *int64 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// An array of custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// *   **true:** TLS 1.3 is supported.
	// *   **false:** TLS 1.3 is not supported.
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether an exclusive IP address is enabled. Valid values:
	//
	// *   **true:** An exclusive IP address is enabled for the domain name.
	// *   **false:** No exclusive IP addresses are enabled for the domain name.
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether HTTP to HTTPS redirection is enabled for the domain name. Valid values:
	//
	// *   **true:** HTTP to HTTPS redirection is enabled.
	// *   **false:** HTTP to HTTPS redirection is disabled.
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// *   **true:** HTTP/2 is enabled.
	// *   **false:** HTTP/2 is disabled.
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// An array of HTTP listener ports.
	HttpPorts []*int64 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// An array of HTTPS listener ports.
	HttpsPorts []*int64 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// *   **true:** IPv6 is enabled.
	// *   **false:** IPv6 is disabled.
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of protection resource that is used. Valid values:
	//
	// *   **share:** shared cluster.
	// *   **gslb:** shared cluster-based intelligent load balancing.
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Indicates whether only SM certificate-based clients can access the domain name. This parameter is returned only if the value of SM2Enabled is true. Valid values:
	//
	// *   true
	// *   false
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate that is added. This parameter is returned only if the value of SM2Enabled is true.
	SM2CertId *bool `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Indicates whether SM certificate-based verification is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The version of the Transport Layer Security (TLS) protocol. Valid values:
	//
	// *   **tlsv1**
	// *   **tlsv1.1**
	// *   **tlsv1.2**
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that WAF uses to obtain the actual IP address of a client. Valid values:
	//
	// *   **0:** No Layer 7 proxies are deployed in front of WAF.
	// *   **1:** WAF reads the first value of the X-Forwarded-For (XFF) header field as the actual IP address of the client.
	// *   **2:** WAF reads the value of a custom header field as the actual IP address of the client.
	XffHeaderMode *int64 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// An array of custom header fields that are used to obtain the actual IP address of a client.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s DescribeDomainDetailResponseBodyListen) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyListen) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyListen) SetCertId(v int64) *DescribeDomainDetailResponseBodyListen {
	s.CertId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetCipherSuite(v int64) *DescribeDomainDetailResponseBodyListen {
	s.CipherSuite = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetCustomCiphers(v []*string) *DescribeDomainDetailResponseBodyListen {
	s.CustomCiphers = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetEnableTLSv3(v bool) *DescribeDomainDetailResponseBodyListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetExclusiveIp(v bool) *DescribeDomainDetailResponseBodyListen {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetFocusHttps(v bool) *DescribeDomainDetailResponseBodyListen {
	s.FocusHttps = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttp2Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttpPorts(v []*int64) *DescribeDomainDetailResponseBodyListen {
	s.HttpPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttpsPorts(v []*int64) *DescribeDomainDetailResponseBodyListen {
	s.HttpsPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetIPv6Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.IPv6Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetProtectionResource(v string) *DescribeDomainDetailResponseBodyListen {
	s.ProtectionResource = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetSM2AccessOnly(v bool) *DescribeDomainDetailResponseBodyListen {
	s.SM2AccessOnly = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetSM2CertId(v bool) *DescribeDomainDetailResponseBodyListen {
	s.SM2CertId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetSM2Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.SM2Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetTLSVersion(v string) *DescribeDomainDetailResponseBodyListen {
	s.TLSVersion = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetXffHeaderMode(v int64) *DescribeDomainDetailResponseBodyListen {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetXffHeaders(v []*string) *DescribeDomainDetailResponseBodyListen {
	s.XffHeaders = v
	return s
}

type DescribeDomainDetailResponseBodyRedirect struct {
	// An array of addresses of origin servers.
	Backends []*DescribeDomainDetailResponseBodyRedirectBackends `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The timeout period of the connection. Unit: seconds. Valid values: 5 to 120.
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Indicates whether HTTPS to HTTP redirection is enabled for back-to-origin requests of the domain name. Valid values:
	//
	// *   **true:** HTTPS to HTTP redirection for back-to-origin requests of the domain name is enabled.
	// *   **false:** HTTPS to HTTP redirection for back-to-origin requests of the domain name is disabled.
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Indicates whether the persistent connection feature is enabled. Valid values:
	//
	// *   **true:** The persistent connection feature is enabled. This is the default value.
	// *   **false:** The persistent connection feature is disabled.
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter specifies the number of reused persistent connections when you enable the persistent connection feature.
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of persistent connections that are in the Idle state. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// >  This parameter specifies the period of time during which a reused persistent connection is allowed to remain in the Idle state before the persistent connection is released.
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that is used when WAF forwards requests to the origin server. Valid values:
	//
	// *   **ip_hash:** the IP hash algorithm.
	// *   **roundRobin:** the round-robin algorithm.
	// *   **leastTime:** the least response time algorithm.
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The read timeout period. Unit: seconds. Valid values: 5 to 1800.
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// An array of key-value pairs that are used to mark the requests that pass through the WAF instance.
	RequestHeaders []*DescribeDomainDetailResponseBodyRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Indicates whether WAF retries to forward requests when requests fail to be forwarded to the origin server. Valid values:
	//
	// *   **true:** WAF retries to forward requests. This is the default value.
	// *   **false:** WAF does not retry to forward requests.
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// Indicates whether origin Server Name Indication (SNI) is enabled. Valid values:
	//
	// *   **true:** Origin SNI is enabled.
	// *   **false:** Origin SNI is disabled. This is the default value.
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI field.
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// The write timeout period. Unit: seconds. Valid values: 5 to 1800.
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Indicates whether the X-Forward-For-Proto header is used to identify the protocol used by WAF to forward requests to the origin server. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirect) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirect) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackends(v []*DescribeDomainDetailResponseBodyRedirectBackends) *DescribeDomainDetailResponseBodyRedirect {
	s.Backends = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetConnectTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetFocusHttpBackend(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepalive(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Keepalive = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepaliveRequests(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepaliveTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetLoadbalance(v string) *DescribeDomainDetailResponseBodyRedirect {
	s.Loadbalance = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetReadTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetRequestHeaders(v []*DescribeDomainDetailResponseBodyRedirectRequestHeaders) *DescribeDomainDetailResponseBodyRedirect {
	s.RequestHeaders = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetRetry(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Retry = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetSniEnabled(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.SniEnabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetSniHost(v string) *DescribeDomainDetailResponseBodyRedirect {
	s.SniHost = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetWriteTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetXffProto(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.XffProto = &v
	return s
}

type DescribeDomainDetailResponseBodyRedirectBackends struct {
	// The IP address or domain name of the origin server.
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectBackends) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectBackends) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectBackends) SetBackend(v string) *DescribeDomainDetailResponseBodyRedirectBackends {
	s.Backend = &v
	return s
}

type DescribeDomainDetailResponseBodyRedirectRequestHeaders struct {
	// The custom header field.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom header field.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) SetKey(v string) *DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) SetValue(v string) *DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	s.Value = &v
	return s
}

type DescribeDomainDetailResponseBodySM2CertDetail struct {
	// The domain name of your website.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The end of the validity period of the SSL certificate. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the SSL certificate.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the SSL certificate.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// All domain names that are bound to the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	// The beginning of the validity period of the SSL certificate. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainDetailResponseBodySM2CertDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodySM2CertDetail) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetCommonName(v string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.CommonName = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetEndTime(v int64) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetId(v string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.Id = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetName(v string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.Name = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetSans(v []*string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.Sans = v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetStartTime(v int64) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.StartTime = &v
	return s
}

type DescribeDomainDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDetailResponse) SetStatusCode(v int32) *DescribeDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDetailResponse) SetBody(v *DescribeDomainDetailResponseBody) *DescribeDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	// An array of HTTPS listener ports.
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// The ID of the request.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The page number of the page to return. Default value: 1.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The HTTPS address of the origin server.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Queries the list of a domain name that is added to Web Application Firewall (WAF).
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tag of the resource. You can specify up to 20 tags.
	Tag []*DescribeDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetBackend(v string) *DescribeDomainsRequest {
	s.Backend = &v
	return s
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceId(v string) *DescribeDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetRegionId(v string) *DescribeDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetTag(v []*DescribeDomainsRequestTag) *DescribeDomainsRequest {
	s.Tag = v
	return s
}

type DescribeDomainsRequestTag struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequestTag) SetKey(v string) *DescribeDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDomainsRequestTag) SetValue(v string) *DescribeDomainsRequestTag {
	s.Value = &v
	return s
}

type DescribeDomainsResponseBody struct {
	// The domain names that are added to WAF in CNAME record mode.
	Domains []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDomainsResponseBodyDomains struct {
	// The back-to-origin settings.
	Backeds *DescribeDomainsResponseBodyDomainsBackeds `json:"Backeds,omitempty" xml:"Backeds,omitempty" type:"Struct"`
	// The CNAME assigned by WAF to the domain name.
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name that is added to WAF in CNAME record mode.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The configurations of the listeners.
	ListenPorts *DescribeDomainsResponseBodyDomainsListenPorts `json:"ListenPorts,omitempty" xml:"ListenPorts,omitempty" type:"Struct"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The status of the domain name. Valid values:
	//
	// *   **1:** The domain name is in a normal state.
	// *   **2:** The domain name is being created.
	// *   **3:** The domain name is being modified.
	// *   **4:** The domain name is being released.
	// *   **5:** WAF no longer forwards traffic that is sent to the domain name.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetBackeds(v *DescribeDomainsResponseBodyDomainsBackeds) *DescribeDomainsResponseBodyDomains {
	s.Backeds = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCname(v string) *DescribeDomainsResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v string) *DescribeDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetListenPorts(v *DescribeDomainsResponseBodyDomainsListenPorts) *DescribeDomainsResponseBodyDomains {
	s.ListenPorts = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetResourceManagerResourceGroupId(v string) *DescribeDomainsResponseBodyDomains {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetStatus(v int32) *DescribeDomainsResponseBodyDomains {
	s.Status = &v
	return s
}

type DescribeDomainsResponseBodyDomainsBackeds struct {
	// The HTTP addresses of the origin server.
	Http []*DescribeDomainsResponseBodyDomainsBackedsHttp `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	// The HTTPS addresses of the origin server.
	Https []*DescribeDomainsResponseBodyDomainsBackedsHttps `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsBackeds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackeds) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) SetHttp(v []*DescribeDomainsResponseBodyDomainsBackedsHttp) *DescribeDomainsResponseBodyDomainsBackeds {
	s.Http = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) SetHttps(v []*DescribeDomainsResponseBodyDomainsBackedsHttps) *DescribeDomainsResponseBodyDomainsBackeds {
	s.Https = v
	return s
}

type DescribeDomainsResponseBodyDomainsBackedsHttp struct {
	// The HTTP address of the origin server.
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttp) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttp) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttp) SetBackend(v string) *DescribeDomainsResponseBodyDomainsBackedsHttp {
	s.Backend = &v
	return s
}

type DescribeDomainsResponseBodyDomainsBackedsHttps struct {
	// The HTTPS address of the origin server.
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttps) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttps) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttps) SetBackend(v string) *DescribeDomainsResponseBodyDomainsBackedsHttps {
	s.Backend = &v
	return s
}

type DescribeDomainsResponseBodyDomainsListenPorts struct {
	// The HTTP listener ports.
	Http []*int64 `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	// The HTTPS listener ports.
	Https []*int64 `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsListenPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsListenPorts) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) SetHttp(v []*int64) *DescribeDomainsResponseBodyDomainsListenPorts {
	s.Http = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) SetHttps(v []*int64) *DescribeDomainsResponseBodyDomainsListenPorts {
	s.Https = v
	return s
}

type DescribeDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeFlowChartRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds. The value must be an integral multiple of 60.
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowChartRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartRequest) SetEndTimestamp(v string) *DescribeFlowChartRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowChartRequest) SetInstanceId(v string) *DescribeFlowChartRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetInterval(v string) *DescribeFlowChartRequest {
	s.Interval = &v
	return s
}

func (s *DescribeFlowChartRequest) SetRegionId(v string) *DescribeFlowChartRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetResource(v string) *DescribeFlowChartRequest {
	s.Resource = &v
	return s
}

func (s *DescribeFlowChartRequest) SetResourceManagerResourceGroupId(v string) *DescribeFlowChartRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetStartTimestamp(v string) *DescribeFlowChartRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeFlowChartResponseBody struct {
	// The traffic statistics.
	FlowChart []*DescribeFlowChartResponseBodyFlowChart `json:"FlowChart,omitempty" xml:"FlowChart,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowChartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponseBody) SetFlowChart(v []*DescribeFlowChartResponseBodyFlowChart) *DescribeFlowChartResponseBody {
	s.FlowChart = v
	return s
}

func (s *DescribeFlowChartResponseBody) SetRequestId(v string) *DescribeFlowChartResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowChartResponseBodyFlowChart struct {
	// The number of requests that are blocked by custom access control list (ACL) rules.
	AclCustomBlockSum *int64 `json:"AclCustomBlockSum,omitempty" xml:"AclCustomBlockSum,omitempty"`
	// The number of requests that are monitored by custom ACL rules.
	AclCustomReportsSum *int64 `json:"AclCustomReportsSum,omitempty" xml:"AclCustomReportsSum,omitempty"`
	// The number of requests that are blocked by scan protection rules.
	AntiScanBlockSum *int64 `json:"AntiScanBlockSum,omitempty" xml:"AntiScanBlockSum,omitempty"`
	// The number of requests that are blocked by bot management rules.
	AntibotBlockSum *int64 `json:"AntibotBlockSum,omitempty" xml:"AntibotBlockSum,omitempty"`
	// The number of requests that are monitored by bot management rules.
	AntibotReportSum *string `json:"AntibotReportSum,omitempty" xml:"AntibotReportSum,omitempty"`
	// The number of requests that are monitored by scan protection rules.
	AntiscanReportsSum *int64 `json:"AntiscanReportsSum,omitempty" xml:"AntiscanReportsSum,omitempty"`
	// The number of requests that are blocked by the IP address blacklist.
	BlacklistBlockSum *string `json:"BlacklistBlockSum,omitempty" xml:"BlacklistBlockSum,omitempty"`
	// The number of requests that are monitored by the IP address blacklist.
	BlacklistReportsSum *int64 `json:"BlacklistReportsSum,omitempty" xml:"BlacklistReportsSum,omitempty"`
	// The number of requests that are blocked by custom HTTP flood protection rules.
	CcCustomBlockSum *int64 `json:"CcCustomBlockSum,omitempty" xml:"CcCustomBlockSum,omitempty"`
	// The number of requests that are monitored by custom HTTP flood protection rules.
	CcCustomReportsSum *int64 `json:"CcCustomReportsSum,omitempty" xml:"CcCustomReportsSum,omitempty"`
	// The number of requests that are blocked by HTTP flood protection rules created by the system.
	CcSystemBlocksSum *int64 `json:"CcSystemBlocksSum,omitempty" xml:"CcSystemBlocksSum,omitempty"`
	// The number of requests that are monitored by HTTP flood protection rules created by the system.
	CcSystemReportsSum *int64 `json:"CcSystemReportsSum,omitempty" xml:"CcSystemReportsSum,omitempty"`
	// The total number of requests.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The total number of requests that are redirected to the WAF instance.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The serial number of the time interval. The serial numbers are arranged in chronological order.
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The peak traffic.
	MaxPv *int64 `json:"MaxPv,omitempty" xml:"MaxPv,omitempty"`
	// The total number of requests that are forwarded by the WAF instance.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The number of requests that are blocked by rate limiting rules.
	RatelimitBlockSum *int64 `json:"RatelimitBlockSum,omitempty" xml:"RatelimitBlockSum,omitempty"`
	// The number of requests that are monitored by rate limiting rules.
	RatelimitReportSum *int64 `json:"RatelimitReportSum,omitempty" xml:"RatelimitReportSum,omitempty"`
	// The number of requests that are blocked by region blacklist rules.
	RegionBlockBlocksSum *int64 `json:"RegionBlockBlocksSum,omitempty" xml:"RegionBlockBlocksSum,omitempty"`
	// The number of requests that are monitored by region blacklist rules.
	RegionBlockReportsSum *int64 `json:"RegionBlockReportsSum,omitempty" xml:"RegionBlockReportsSum,omitempty"`
	// The total number of bot requests.
	RobotCount *int64 `json:"RobotCount,omitempty" xml:"RobotCount,omitempty"`
	// The number of requests that are blocked by basic protection rules.
	WafBlockSum *int64 `json:"WafBlockSum,omitempty" xml:"WafBlockSum,omitempty"`
	// The number of requests that are monitored by basic protection rules.
	WafReportSum *string `json:"WafReportSum,omitempty" xml:"WafReportSum,omitempty"`
}

func (s DescribeFlowChartResponseBodyFlowChart) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartResponseBodyFlowChart) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAclCustomBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AclCustomBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAclCustomReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AclCustomReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntiScanBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntiScanBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntibotBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntibotBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntibotReportSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.AntibotReportSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntiscanReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntiscanReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetBlacklistBlockSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.BlacklistBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetBlacklistReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.BlacklistReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcCustomBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcCustomBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcCustomReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcCustomReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcSystemBlocksSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcSystemBlocksSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcSystemReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcSystemReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCount(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.Count = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetInBytes(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.InBytes = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetIndex(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.Index = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetMaxPv(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.MaxPv = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetOutBytes(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.OutBytes = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRatelimitBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RatelimitBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRatelimitReportSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RatelimitReportSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRegionBlockBlocksSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RegionBlockBlocksSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRegionBlockReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RegionBlockReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRobotCount(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RobotCount = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetWafBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.WafBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetWafReportSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.WafReportSum = &v
	return s
}

type DescribeFlowChartResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowChartResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponse) SetHeaders(v map[string]*string) *DescribeFlowChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowChartResponse) SetStatusCode(v int32) *DescribeFlowChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowChartResponse) SetBody(v *DescribeFlowChartResponseBody) *DescribeFlowChartResponse {
	s.Body = v
	return s
}

type DescribeFlowTopResourceRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowTopResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceRequest) SetEndTimestamp(v string) *DescribeFlowTopResourceRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetInstanceId(v string) *DescribeFlowTopResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetRegionId(v string) *DescribeFlowTopResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetResourceManagerResourceGroupId(v string) *DescribeFlowTopResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetStartTimestamp(v string) *DescribeFlowTopResourceRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeFlowTopResourceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 protected objects that receive requests.
	RuleHitsTopResource []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource `json:"RuleHitsTopResource,omitempty" xml:"RuleHitsTopResource,omitempty" type:"Repeated"`
}

func (s DescribeFlowTopResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponseBody) SetRequestId(v string) *DescribeFlowTopResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBody) SetRuleHitsTopResource(v []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource) *DescribeFlowTopResourceResponseBody {
	s.RuleHitsTopResource = v
	return s
}

type DescribeFlowTopResourceResponseBodyRuleHitsTopResource struct {
	// The total number of requests received by the protected object in a specified time range.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeFlowTopResourceResponseBodyRuleHitsTopResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceResponseBodyRuleHitsTopResource) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) SetCount(v int64) *DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	s.Count = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) SetResource(v string) *DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	s.Resource = &v
	return s
}

type DescribeFlowTopResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowTopResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowTopResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponse) SetHeaders(v map[string]*string) *DescribeFlowTopResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTopResourceResponse) SetStatusCode(v int32) *DescribeFlowTopResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowTopResourceResponse) SetBody(v *DescribeFlowTopResourceResponseBody) *DescribeFlowTopResourceResponse {
	s.Body = v
	return s
}

type DescribeFlowTopUrlRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowTopUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlRequest) SetEndTimestamp(v string) *DescribeFlowTopUrlRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetInstanceId(v string) *DescribeFlowTopUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetRegionId(v string) *DescribeFlowTopUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetResource(v string) *DescribeFlowTopUrlRequest {
	s.Resource = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetResourceManagerResourceGroupId(v string) *DescribeFlowTopUrlRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetStartTimestamp(v string) *DescribeFlowTopUrlRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeFlowTopUrlResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 URLs that are used to initiate requests.
	RuleHitsTopUrl []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl `json:"RuleHitsTopUrl,omitempty" xml:"RuleHitsTopUrl,omitempty" type:"Repeated"`
}

func (s DescribeFlowTopUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponseBody) SetRequestId(v string) *DescribeFlowTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBody) SetRuleHitsTopUrl(v []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) *DescribeFlowTopUrlResponseBody {
	s.RuleHitsTopUrl = v
	return s
}

type DescribeFlowTopUrlResponseBodyRuleHitsTopUrl struct {
	// The total number of requests that are initiated by using the URL.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The URL that is used to initiate requests.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) SetCount(v int64) *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	s.Count = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) SetUrl(v string) *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	s.Url = &v
	return s
}

type DescribeFlowTopUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowTopUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponse) SetHeaders(v map[string]*string) *DescribeFlowTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTopUrlResponse) SetStatusCode(v int32) *DescribeFlowTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowTopUrlResponse) SetBody(v *DescribeFlowTopUrlResponseBody) *DescribeFlowTopUrlResponse {
	s.Body = v
	return s
}

type DescribeHybridCloudGroupsRequest struct {
	// The ID of the hybrid cloud cluster.
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of proxy cluster that is used. Valid values:
	//
	// *   **service**: service-based traffic mirroring.
	// *   **cname**: reverse proxy.
	ClusterProxyType *string `json:"ClusterProxyType,omitempty" xml:"ClusterProxyType,omitempty"`
	// The name of the node group that you want to query.
	GroupName *int32 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// *   **protect**
	// *   **control**
	// *   **storage**
	// *   **controlStorage**
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsRequest) SetClusterId(v int64) *DescribeHybridCloudGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetClusterProxyType(v string) *DescribeHybridCloudGroupsRequest {
	s.ClusterProxyType = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetGroupName(v int32) *DescribeHybridCloudGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetGroupType(v string) *DescribeHybridCloudGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetInstanceId(v string) *DescribeHybridCloudGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetPageNumber(v int32) *DescribeHybridCloudGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetPageSize(v int32) *DescribeHybridCloudGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetRegionId(v string) *DescribeHybridCloudGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeHybridCloudGroupsResponseBody struct {
	// The node groups.
	Groups []*DescribeHybridCloudGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsResponseBody) SetGroups(v []*DescribeHybridCloudGroupsResponseBodyGroups) *DescribeHybridCloudGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBody) SetRequestId(v string) *DescribeHybridCloudGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBody) SetTotalCount(v int32) *DescribeHybridCloudGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHybridCloudGroupsResponseBodyGroups struct {
	// The back-to-origin mark of the protected cluster. The value is in the {ISP name}-{Continent name}-{City name}-{Back-to-origin identifier} format. The back-to-origin identifier is optional.
	//
	// >  For more information about ISP names, continent names, city names, and back-to-origin identifiers, see the following sections.
	BackSourceMark *string `json:"BackSourceMark,omitempty" xml:"BackSourceMark,omitempty"`
	// The continent code of the protected cluster.
	//
	// >  For more information about continent codes, see Continent codes in this topic.
	ContinentsValue *int32 `json:"ContinentsValue,omitempty" xml:"ContinentsValue,omitempty"`
	// The ID of the node group.
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the node group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the node group. Valid values:
	//
	// *   **protect**
	// *   **control**
	// *   **storage**
	// *   **controlStorage**
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The IP address of the server used for load balancing.
	LoadBalanceIp *string `json:"LoadBalanceIp,omitempty" xml:"LoadBalanceIp,omitempty"`
	// The ID of the protection node.
	LocationId *int64 `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The ISP code of the protected cluster.
	//
	// >  For more information about ISP codes, see ISP codes in this topic.
	OperatorValue *int32 `json:"OperatorValue,omitempty" xml:"OperatorValue,omitempty"`
	// The port that is used by the hybrid cloud cluster. The value of this parameter is a string. If multiple ports are returned, the value is in the **port1,port2,port3** format.
	Ports *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
	// The city code of the protected cluster.
	//
	// >  For more information about city codes, see City codes in this topic.
	RegionCodeValue *int32 `json:"RegionCodeValue,omitempty" xml:"RegionCodeValue,omitempty"`
	// The description of the node group.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeHybridCloudGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetBackSourceMark(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.BackSourceMark = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetContinentsValue(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.ContinentsValue = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetGroupId(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetGroupName(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetGroupType(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.GroupType = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetLoadBalanceIp(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.LoadBalanceIp = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetLocationId(v int64) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.LocationId = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetOperatorValue(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.OperatorValue = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetPorts(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.Ports = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetRegionCodeValue(v int32) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.RegionCodeValue = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponseBodyGroups) SetRemark(v string) *DescribeHybridCloudGroupsResponseBodyGroups {
	s.Remark = &v
	return s
}

type DescribeHybridCloudGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudGroupsResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudGroupsResponse) SetStatusCode(v int32) *DescribeHybridCloudGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudGroupsResponse) SetBody(v *DescribeHybridCloudGroupsResponseBody) *DescribeHybridCloudGroupsResponse {
	s.Body = v
	return s
}

type DescribeHybridCloudResourcesRequest struct {
	// The back-to-origin IP address or domain name.
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// Specifies whether the public cloud disaster recovery feature is enabled for the domain name. Valid values:
	//
	// *   **true**
	// *   **false**
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The domain name that you want to query.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesRequest) SetBackend(v string) *DescribeHybridCloudResourcesRequest {
	s.Backend = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetCnameEnabled(v bool) *DescribeHybridCloudResourcesRequest {
	s.CnameEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetDomain(v string) *DescribeHybridCloudResourcesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetInstanceId(v string) *DescribeHybridCloudResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetPageNumber(v int64) *DescribeHybridCloudResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetPageSize(v int64) *DescribeHybridCloudResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetRegionId(v string) *DescribeHybridCloudResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeHybridCloudResourcesResponseBody struct {
	// The domain names.
	Domains []*DescribeHybridCloudResourcesResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBody) SetDomains(v []*DescribeHybridCloudResourcesResponseBodyDomains) *DescribeHybridCloudResourcesResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBody) SetRequestId(v string) *DescribeHybridCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBody) SetTotalCount(v int64) *DescribeHybridCloudResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHybridCloudResourcesResponseBodyDomains struct {
	// The CNAME assigned by WAF.
	//
	// >  This parameter is returned only if the value of **CnameEnabled** is true.
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The access ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The listeners.
	Listen *DescribeHybridCloudResourcesResponseBodyDomainsListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The configurations of the forwarding rule.
	Redirect *DescribeHybridCloudResourcesResponseBodyDomainsRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The status of the domain name. Valid values:
	//
	// *   **1:** The domain name is in a normal state.
	// *   **2:** The domain name is being created.
	// *   **3:** The domain name is being modified.
	// *   **4:** The domain name is being released.
	// *   **5:** WAF no longer forwards the traffic of the domain name.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetCname(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetDomain(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetId(v int64) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Id = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetListen(v *DescribeHybridCloudResourcesResponseBodyDomainsListen) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Listen = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetRedirect(v *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Redirect = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetStatus(v int32) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Status = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetUid(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Uid = &v
	return s
}

type DescribeHybridCloudResourcesResponseBodyDomainsListen struct {
	// The ID of the certificate.
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The types of cipher suites that are added. Valid values:
	//
	// *   **1:** all cipher suites.
	// *   **2:** strong cipher suites.
	// *   **99:** custom cipher suites.
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites.
	//
	// >  This parameter is returned only if the value of **CipherSuite** is **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// *   **true**
	// *   **false**
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether exclusive IP addresses are supported. Valid values:
	//
	// *   **true**
	// *   **false**
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether the HTTP to HTTPS redirection feature is enabled for the domain name. Valid values:
	//
	// *   **true**
	// *   **false**
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The HTTP listener ports.
	HttpPorts []*int64 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The HTTPS listener ports.
	HttpsPorts []*int64 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// *   **true**
	// *   **false**
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of the protection resource. Valid values:
	//
	// *   **share:** shared cluster.
	// *   **gslb:** shared cluster-based intelligent load balancing.
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// The version of the Transport Layer Security (TLS) protocol. Valid values:
	//
	// *   **tlsv1**
	// *   **tlsv1.1**
	// *   **tlsv1.2**
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that is used to obtain the actual IP address of a client. Valid values:
	//
	// *   **0**: No Layer 7 proxies are deployed in front of WAF.
	// *   **1**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the actual IP address of the client.
	// *   **2**: WAF reads the value of a custom header field as the actual IP address of the client.
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the actual IP addresses of clients. The value is in the \["header1","header2",...] format.
	//
	// >  This parameter is returned only if the value of **XffHeaderMode** is 2.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsListen) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsListen) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetCertId(v string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.CertId = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetCipherSuite(v int32) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.CipherSuite = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetCustomCiphers(v []*string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.CustomCiphers = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetEnableTLSv3(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetExclusiveIp(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetFocusHttps(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.FocusHttps = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetHttp2Enabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetHttpPorts(v []*int64) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.HttpPorts = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetHttpsPorts(v []*int64) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.HttpsPorts = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetIPv6Enabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.IPv6Enabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetProtectionResource(v string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.ProtectionResource = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetTLSVersion(v string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.TLSVersion = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetXffHeaderMode(v int32) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetXffHeaders(v []*string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.XffHeaders = v
	return s
}

type DescribeHybridCloudResourcesResponseBodyDomainsRedirect struct {
	// The IP addresses or domain names of the origin server.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// Indicates whether the public cloud disaster recovery feature is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The timeout period for connections. Unit: seconds. Valid values: 5 to 120.
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Indicates whether the HTTPS to HTTP redirection feature is enabled for back-to-origin requests. Valid values:
	//
	// *   **true**
	// *   **false**
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Indicates whether the persistent connection feature is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter indicates the number of reused persistent connections after the persistent connection feature is enabled.
	KeepaliveRequests *int64 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period for persistent connections that are in the Idle state. Unit: seconds. Valid values: 1 to 60. Default value: 15.
	//
	// >  This parameter indicates the period of time during which a reused persistent connection can remain in the Idle state before the persistent connection is released.
	KeepaliveTimeout *int64 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that is used to forward requests to the origin server. Valid values:
	//
	// *   **iphash**
	// *   **roundRobin**
	// *   **leastTime**
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The timeout period for read connections. Unit: seconds. Valid values: 5 to 1800.
	ReadTimeout *int64 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The key-value pair that is used to label requests that pass through WAF.
	RequestHeaders []*DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Indicates whether WAF retries forwarding requests if requests fail to be forwarded to the origin server. Valid values:
	//
	// *   **true**
	// *   **false**
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules that are configured for the domain name. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// *   **rs**: the back-to-origin IP addresses or CNAMEs. The value is of the ARRAY type.
	// *   **location**: the name of the protection node. The value is of the STRING type.
	// *   **locationId**: the ID of the protection node. The value is of the LONG type.
	RoutingRules *string `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	// Indicates whether the origin Server Name Indication (SNI) feature is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI field. If the parameter is left empty, the value of the **Host** field in the request header is automatically used as the value of the SNI field.
	//
	// >  This parameter is returned only if the value of **SniEnabled** is **true**.
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// The timeout period for write connections. Unit: seconds. Valid values: 5 to 1800.
	WriteTimeout *int64 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirect) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetBackends(v []*string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Backends = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetCnameEnabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetConnectTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetFocusHttpBackend(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetKeepalive(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Keepalive = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetKeepaliveRequests(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetKeepaliveTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetLoadbalance(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Loadbalance = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetReadTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetRequestHeaders(v []*DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.RequestHeaders = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetRetry(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Retry = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetRoutingRules(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.RoutingRules = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetSniEnabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.SniEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetSniHost(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.SniHost = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetWriteTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.WriteTimeout = &v
	return s
}

type DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders struct {
	// The key of the custom header field.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom header field.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) SetKey(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) SetValue(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders {
	s.Value = &v
	return s
}

type DescribeHybridCloudResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudResourcesResponse) SetStatusCode(v int32) *DescribeHybridCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponse) SetBody(v *DescribeHybridCloudResourcesResponseBody) *DescribeHybridCloudResourcesResponse {
	s.Body = v
	return s
}

type DescribeHybridCloudUserRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserRequest) SetInstanceId(v string) *DescribeHybridCloudUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudUserRequest) SetRegionId(v string) *DescribeHybridCloudUserRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudUserRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUserRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeHybridCloudUserResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the ports that can be used by a hybrid cloud cluster.
	UserInfo *DescribeHybridCloudUserResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DescribeHybridCloudUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserResponseBody) SetRequestId(v string) *DescribeHybridCloudUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudUserResponseBody) SetUserInfo(v *DescribeHybridCloudUserResponseBodyUserInfo) *DescribeHybridCloudUserResponseBody {
	s.UserInfo = v
	return s
}

type DescribeHybridCloudUserResponseBodyUserInfo struct {
	// The HTTP ports. The value is a string. If multiple ports are returned, the value is in the **port1,port2,port3** format.
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS ports. The value is a string. If multiple ports are returned, the value is in the **port1,port2,port3** format.
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
}

func (s DescribeHybridCloudUserResponseBodyUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudUserResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) SetHttpPorts(v string) *DescribeHybridCloudUserResponseBodyUserInfo {
	s.HttpPorts = &v
	return s
}

func (s *DescribeHybridCloudUserResponseBodyUserInfo) SetHttpsPorts(v string) *DescribeHybridCloudUserResponseBodyUserInfo {
	s.HttpsPorts = &v
	return s
}

type DescribeHybridCloudUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHybridCloudUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudUserResponse) SetStatusCode(v int32) *DescribeHybridCloudUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudUserResponse) SetBody(v *DescribeHybridCloudUserResponseBody) *DescribeHybridCloudUserResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetRegionId(v string) *DescribeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetResourceManagerResourceGroupId(v string) *DescribeInstanceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	// The details of the WAF instance.
	Details *DescribeInstanceResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	// The edition of the WAF instance.
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The expiration time of the WAF instance.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the WAF instance has overdue payments. Valid values:
	//
	// *   **0**: The WAF instance does not have overdue payments.
	// *   **1**: The WAF instance has overdue payments.
	InDebt *string `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	// The ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The billing method of the WAF instance. Valid values:
	//
	// *   **POSTPAY:** The WAF instance uses the pay-as-you-go billing method.
	// *   **PREPAY:** The WAF instance uses the subscription billing method.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The purchase time of the WAF instance. The time is in the UNIX timestamp format. The time is displayed in UTC. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the WAF instance. Valid values:
	//
	// *   **1:** The WAF instance is in a normal state.
	// *   **2:** The WAF instance has expired.
	// *   **3:** The WAF instance has been released.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetDetails(v *DescribeInstanceResponseBodyDetails) *DescribeInstanceResponseBody {
	s.Details = v
	return s
}

func (s *DescribeInstanceResponseBody) SetEdition(v string) *DescribeInstanceResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEndTime(v int64) *DescribeInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInDebt(v string) *DescribeInstanceResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPayType(v string) *DescribeInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStartTime(v int64) *DescribeInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v int32) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

type DescribeInstanceResponseBodyDetails struct {
	// The maximum number of IP addresses that can be added to the match content of a match condition. For more information, see [Match conditions](~~374354~~).
	AclRuleMaxIpCount *int64 `json:"AclRuleMaxIpCount,omitempty" xml:"AclRuleMaxIpCount,omitempty"`
	// Indicates whether the scan protection module is supported. Valid values:
	//
	// *   **true:** The scan protection module is supported.
	// *   **false:** The scan protection module is not supported.
	AntiScan *bool `json:"AntiScan,omitempty" xml:"AntiScan,omitempty"`
	// The maximum number of scan protection rule templates that can be configured.
	AntiScanTemplateMaxCount *int64 `json:"AntiScanTemplateMaxCount,omitempty" xml:"AntiScanTemplateMaxCount,omitempty"`
	// The maximum number of back-to-origin IP addresses that can be configured.
	BackendMaxCount *int64 `json:"BackendMaxCount,omitempty" xml:"BackendMaxCount,omitempty"`
	// Indicates whether the basic protection rule module is supported. Valid values:
	//
	// *   **true:** The basic protection rule module is supported.
	// *   **false:** The basic protection rule module is not supported.
	BaseWafGroup *bool `json:"BaseWafGroup,omitempty" xml:"BaseWafGroup,omitempty"`
	// The maximum number of protection rules that can be included in a basic protection rule template.
	BaseWafGroupRuleInTemplateMaxCount *int64 `json:"BaseWafGroupRuleInTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleInTemplateMaxCount,omitempty"`
	// The maximum number of basic protection rule templates that can be configured.
	BaseWafGroupRuleTemplateMaxCount *int64 `json:"BaseWafGroupRuleTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleTemplateMaxCount,omitempty"`
	// Indicates whether the bot management module is supported. Valid values:
	//
	// *   **true:** The bot management module is supported.
	// *   **false:** The bot management module is not supported.
	Bot *bool `json:"Bot,omitempty" xml:"Bot,omitempty"`
	// Indicates whether bot management for app protection is supported. Valid values:
	//
	// *   **true:** Bot management for app protection is supported.
	// *   **false:** Bot management for app protection is not supported.
	BotApp *string `json:"BotApp,omitempty" xml:"BotApp,omitempty"`
	// The maximum number of bot management rule templates that can be configured.
	BotTemplateMaxCount *int64 `json:"BotTemplateMaxCount,omitempty" xml:"BotTemplateMaxCount,omitempty"`
	// Indicates whether bot management for website protection is supported. Valid values:
	//
	// *   **true:** Bot management for website protection is supported.
	// *   **false:** Bot management for website protection is not supported.
	BotWeb *string `json:"BotWeb,omitempty" xml:"BotWeb,omitempty"`
	// The maximum number of CNAMEs that can be added.
	CnameResourceMaxCount *int64 `json:"CnameResourceMaxCount,omitempty" xml:"CnameResourceMaxCount,omitempty"`
	// Indicates whether the custom response module is supported. Valid values:
	//
	// *   **true:** The custom response module is supported.
	// *   **false:** The custom response module is not supported.
	CustomResponse *bool `json:"CustomResponse,omitempty" xml:"CustomResponse,omitempty"`
	// The maximum number of rules that can be included in a custom response rule template.
	CustomResponseRuleInTemplateMaxCount *int64 `json:"CustomResponseRuleInTemplateMaxCount,omitempty" xml:"CustomResponseRuleInTemplateMaxCount,omitempty"`
	// The maximum number of custom response rule templates that can be configured.
	CustomResponseTemplateMaxCount *int64 `json:"CustomResponseTemplateMaxCount,omitempty" xml:"CustomResponseTemplateMaxCount,omitempty"`
	// Indicates whether the custom rule module is supported. Valid values:
	//
	// *   **true:** The custom rule module is supported.
	// *   **false:** The custom rule module is not supported.
	CustomRule *bool `json:"CustomRule,omitempty" xml:"CustomRule,omitempty"`
	// The action that can be included in a custom rule.
	CustomRuleAction *string `json:"CustomRuleAction,omitempty" xml:"CustomRuleAction,omitempty"`
	// The match conditions that can be used in a custom rule. For more information, see **Match condition parameters** in the "**Parameters of custom rules (custom_acl)**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	CustomRuleCondition *string `json:"CustomRuleCondition,omitempty" xml:"CustomRuleCondition,omitempty"`
	// The maximum number of rules that can be included in a custom rule template.
	CustomRuleInTemplateMaxCount *int64 `json:"CustomRuleInTemplateMaxCount,omitempty" xml:"CustomRuleInTemplateMaxCount,omitempty"`
	// The statistical object for rate limiting in a custom rule.
	CustomRuleRatelimitor *string `json:"CustomRuleRatelimitor,omitempty" xml:"CustomRuleRatelimitor,omitempty"`
	// The maximum number of custom rule templates that can be configured.
	CustomRuleTemplateMaxCount *int64 `json:"CustomRuleTemplateMaxCount,omitempty" xml:"CustomRuleTemplateMaxCount,omitempty"`
	// The maximum number of protected object groups that can be configured.
	DefenseGroupMaxCount *int64 `json:"DefenseGroupMaxCount,omitempty" xml:"DefenseGroupMaxCount,omitempty"`
	// The maximum number of protected objects that can be included in a protected object group.
	DefenseObjectInGroupMaxCount *int64 `json:"DefenseObjectInGroupMaxCount,omitempty" xml:"DefenseObjectInGroupMaxCount,omitempty"`
	// The maximum number of protected objects to which a protection rule template can be applied.
	DefenseObjectInTemplateMaxCount *int64 `json:"DefenseObjectInTemplateMaxCount,omitempty" xml:"DefenseObjectInTemplateMaxCount,omitempty"`
	// The maximum number of protected objects that can be configured.
	DefenseObjectMaxCount *int64 `json:"DefenseObjectMaxCount,omitempty" xml:"DefenseObjectMaxCount,omitempty"`
	// Indicates whether the data leakage prevention module is supported. Valid values:
	//
	// *   **true:** The data leakage prevention module is supported.
	// *   **false:** The data leakage prevention module is not supported.
	Dlp *bool `json:"Dlp,omitempty" xml:"Dlp,omitempty"`
	// The maximum number of rules that can be included in a data leakage prevention rule template.
	DlpRuleInTemplateMaxCount *int64 `json:"DlpRuleInTemplateMaxCount,omitempty" xml:"DlpRuleInTemplateMaxCount,omitempty"`
	// The maximum number of data leakage prevention rule templates that can be configured.
	DlpTemplateMaxCount *int64 `json:"DlpTemplateMaxCount,omitempty" xml:"DlpTemplateMaxCount,omitempty"`
	// Indicates whether exclusive IP addresses are supported. Valid values:
	//
	// *   **true:** Exclusive IP addresses are supported.
	// *   **false:** Exclusive IP addresses are not supported.
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether global server load balancing (GSLB) is supported. Valid values:
	//
	// *   **true:** GSLB is supported.
	// *   **false:** GSLB is not supported.
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// The HTTP port range that is supported. For more information, see [View supported ports](~~385578~~).
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS port range that is supported. For more information, see [View supported ports](~~385578~~).
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// Indicates whether the IP address blacklist module is supported. Valid values:
	//
	// *   **true:** The IP address blacklist module is supported.
	// *   **false:** The IP address blacklist module is not supported.
	IpBlacklist *bool `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The maximum number of IP addresses that can be added to an IP address blacklist rule.
	IpBlacklistIpInRuleMaxCount *int64 `json:"IpBlacklistIpInRuleMaxCount,omitempty" xml:"IpBlacklistIpInRuleMaxCount,omitempty"`
	// The maximum number of rules that can be included in an IP address blacklist rule template.
	IpBlacklistRuleInTemplateMaxCount *int64 `json:"IpBlacklistRuleInTemplateMaxCount,omitempty" xml:"IpBlacklistRuleInTemplateMaxCount,omitempty"`
	// The maximum number of IP address blacklist rule templates that can be configured.
	IpBlacklistTemplateMaxCount *int64 `json:"IpBlacklistTemplateMaxCount,omitempty" xml:"IpBlacklistTemplateMaxCount,omitempty"`
	// Indicates whether IPv6 is supported. Valid values:
	//
	// *   **true:** IPv6 is supported.
	// *   **false:** IPv6 is not supported.
	Ipv6 *bool `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Indicates whether the log collection feature is supported. Valid values:
	//
	// *   **true:** The log collection feature is supported.
	// *   **false:** The log collection feature is not supported.
	LogService *bool `json:"LogService,omitempty" xml:"LogService,omitempty"`
	// Indicates whether major event protection is supported. Valid values:
	//
	// *   **true:** Major event protection is supported.
	// *   **false:** Major event protection is not supported.
	MajorProtection *bool `json:"MajorProtection,omitempty" xml:"MajorProtection,omitempty"`
	// The maximum number of major event protection rule templates that can be configured.
	MajorProtectionTemplateMaxCount *int64 `json:"MajorProtectionTemplateMaxCount,omitempty" xml:"MajorProtectionTemplateMaxCount,omitempty"`
	// Indicates whether the website tamper-proofing module is supported. Valid values:
	//
	// *   **true:** The website tamper-proofing module is supported.
	// *   **false:** The website tamper-proofing module is not supported.
	Tamperproof *bool `json:"Tamperproof,omitempty" xml:"Tamperproof,omitempty"`
	// The maximum number of rules that can be included in a website tamper-proofing rule template.
	TamperproofRuleInTemplateMaxCount *int64 `json:"TamperproofRuleInTemplateMaxCount,omitempty" xml:"TamperproofRuleInTemplateMaxCount,omitempty"`
	// The maximum number of website tamper-proofing rule templates that can be configured.
	TamperproofTemplateMaxCount *int64 `json:"TamperproofTemplateMaxCount,omitempty" xml:"TamperproofTemplateMaxCount,omitempty"`
	// The maximum number of IP addresses or CIDR blocks that can be added to an IP address blacklist in a batch.
	VastIpBlacklistInFileMaxCount *int64 `json:"VastIpBlacklistInFileMaxCount,omitempty" xml:"VastIpBlacklistInFileMaxCount,omitempty"`
	// The maximum number of IP addresses or CIDR blocks that can be added to an IP address blacklist on a page.
	VastIpBlacklistInOperationMaxCount *int64 `json:"VastIpBlacklistInOperationMaxCount,omitempty" xml:"VastIpBlacklistInOperationMaxCount,omitempty"`
	// The maximum number of IP addresses or CIDR blocks that can be added to an IP address blacklist per Alibaba Cloud account.
	VastIpBlacklistMaxCount *int64 `json:"VastIpBlacklistMaxCount,omitempty" xml:"VastIpBlacklistMaxCount,omitempty"`
	// Indicates whether the whitelist module is supported. Valid values:
	//
	// *   **true:** The whitelist module is supported.
	// *   **false:** The whitelist module is not supported.
	Whitelist *bool `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// The logical operators that can be used in a whitelist rule. For more information, see **Match condition parameters** in the "**Parameters of whitelist rules (whitelist)**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	WhitelistLogical *string `json:"WhitelistLogical,omitempty" xml:"WhitelistLogical,omitempty"`
	// The match fields that can be used in a whitelist rule. For more information, see **Match condition parameters** in the "**Parameters of whitelist rules (whitelist)**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	WhitelistRuleCondition *string `json:"WhitelistRuleCondition,omitempty" xml:"WhitelistRuleCondition,omitempty"`
	// The maximum number of rules that can be included in a whitelist rule template.
	WhitelistRuleInTemplateMaxCount *int64 `json:"WhitelistRuleInTemplateMaxCount,omitempty" xml:"WhitelistRuleInTemplateMaxCount,omitempty"`
	// The maximum number of whitelist rule templates that can be configured.
	WhitelistTemplateMaxCount *int64 `json:"WhitelistTemplateMaxCount,omitempty" xml:"WhitelistTemplateMaxCount,omitempty"`
}

func (s DescribeInstanceResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyDetails) SetAclRuleMaxIpCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AclRuleMaxIpCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScan(v bool) *DescribeInstanceResponseBodyDetails {
	s.AntiScan = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScanTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AntiScanTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBackendMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BackendMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroup(v bool) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroup = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBot(v bool) *DescribeInstanceResponseBodyDetails {
	s.Bot = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotApp(v string) *DescribeInstanceResponseBodyDetails {
	s.BotApp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BotTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotWeb(v string) *DescribeInstanceResponseBodyDetails {
	s.BotWeb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCnameResourceMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CnameResourceMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponse(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomResponse = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRule(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomRule = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleAction(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleAction = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleRatelimitor(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleRatelimitor = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlp(v bool) *DescribeInstanceResponseBodyDetails {
	s.Dlp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlpRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DlpRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlpTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DlpTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetExclusiveIp(v bool) *DescribeInstanceResponseBodyDetails {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetGslb(v bool) *DescribeInstanceResponseBodyDetails {
	s.Gslb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpsPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklist(v bool) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistIpInRuleMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistIpInRuleMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpv6(v bool) *DescribeInstanceResponseBodyDetails {
	s.Ipv6 = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetLogService(v bool) *DescribeInstanceResponseBodyDetails {
	s.LogService = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtection(v bool) *DescribeInstanceResponseBodyDetails {
	s.MajorProtection = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtectionTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.MajorProtectionTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproof(v bool) *DescribeInstanceResponseBodyDetails {
	s.Tamperproof = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproofRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.TamperproofRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproofTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.TamperproofTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInFileMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInFileMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInOperationMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInOperationMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelist(v bool) *DescribeInstanceResponseBodyDetails {
	s.Whitelist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistLogical(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistLogical = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistTemplateMaxCount = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeMajorProtectionBlackIpsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address that you want to query. You can specify this parameter to query an IP address in the IP address blacklist for major event protection by using fuzzy matching.
	IpLike *string `json:"IpLike,omitempty" xml:"IpLike,omitempty"`
	// The method that you want to use to sort the IP addresses **in descending order**. Valid values:
	//
	// *   **gmtModified:** sorts the IP addresses by most recent modification time.
	// *   **ip:** sorts the IP addresses by IP address.
	// *   **templateId:** sorts the IP addresses by template ID.
	// *   **id:** sorts the IP addresses by primary key.
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the rule template for major event protection.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetInstanceId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetIpLike(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.IpLike = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetOrderBy(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetPageNumber(v int32) *DescribeMajorProtectionBlackIpsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetPageSize(v int32) *DescribeMajorProtectionBlackIpsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetRegionId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetResourceManagerResourceGroupId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetRuleId(v int64) *DescribeMajorProtectionBlackIpsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsRequest {
	s.TemplateId = &v
	return s
}

type DescribeMajorProtectionBlackIpsResponseBody struct {
	// An array of IP addresses in the IP address blacklist.
	IpList []*DescribeMajorProtectionBlackIpsResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of IP addresses in the blacklist.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetIpList(v []*DescribeMajorProtectionBlackIpsResponseBodyIpList) *DescribeMajorProtectionBlackIpsResponseBody {
	s.IpList = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetRequestId(v string) *DescribeMajorProtectionBlackIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetTotalCount(v int64) *DescribeMajorProtectionBlackIpsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeMajorProtectionBlackIpsResponseBodyIpList struct {
	// The description of the IP address in the blacklist.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time after which the IP address blacklist becomes invalid. Unit: seconds.
	//
	// >  If the value of this parameter is **0**, the blacklist is permanently valid.
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The most recent time when the IP address blacklist was modified.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The IP address in the IP address blacklist.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the rule template for major event protection.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetDescription(v string) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.Description = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetExpiredTime(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetGmtModified(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.GmtModified = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetIp(v string) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetRuleId(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.RuleId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.TemplateId = &v
	return s
}

type DescribeMajorProtectionBlackIpsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMajorProtectionBlackIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetHeaders(v map[string]*string) *DescribeMajorProtectionBlackIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetStatusCode(v int32) *DescribeMajorProtectionBlackIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetBody(v *DescribeMajorProtectionBlackIpsResponseBody) *DescribeMajorProtectionBlackIpsResponse {
	s.Body = v
	return s
}

type DescribeMemberAccountsRequest struct {
	// The status of the member that you want to query.
	//
	// *   **enabled**: managed.
	// *   **disabled**: not managed.
	// *   **disabling**: being deleted.
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source IP address of the request. The system specifies this parameter.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeMemberAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsRequest) SetAccountStatus(v string) *DescribeMemberAccountsRequest {
	s.AccountStatus = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetInstanceId(v string) *DescribeMemberAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetRegionId(v string) *DescribeMemberAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetResourceManagerResourceGroupId(v string) *DescribeMemberAccountsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeMemberAccountsRequest) SetSourceIp(v string) *DescribeMemberAccountsRequest {
	s.SourceIp = &v
	return s
}

type DescribeMemberAccountsResponseBody struct {
	// The information about the member.
	AccountInfos []*DescribeMemberAccountsResponseBodyAccountInfos `json:"AccountInfos,omitempty" xml:"AccountInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMemberAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsResponseBody) SetAccountInfos(v []*DescribeMemberAccountsResponseBodyAccountInfos) *DescribeMemberAccountsResponseBody {
	s.AccountInfos = v
	return s
}

func (s *DescribeMemberAccountsResponseBody) SetRequestId(v string) *DescribeMemberAccountsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMemberAccountsResponseBodyAccountInfos struct {
	// The ID of the member.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the member.
	//
	// *   **enabled**: managed.
	// *   **disabled**: not managed.
	// *   **disabling**: being deleted.
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The description of the member.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the member was added.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s DescribeMemberAccountsResponseBodyAccountInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberAccountsResponseBodyAccountInfos) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetAccountId(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.AccountId = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetAccountName(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.AccountName = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetAccountStatus(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.AccountStatus = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetDescription(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.Description = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetGmtCreate(v int64) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.GmtCreate = &v
	return s
}

type DescribeMemberAccountsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMemberAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMemberAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMemberAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsResponse) SetHeaders(v map[string]*string) *DescribeMemberAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMemberAccountsResponse) SetStatusCode(v int32) *DescribeMemberAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMemberAccountsResponse) SetBody(v *DescribeMemberAccountsResponseBody) *DescribeMemberAccountsResponse {
	s.Body = v
	return s
}

type DescribePeakTrendRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds. The value must be an integral multiple of 60.
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribePeakTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendRequest) SetEndTimestamp(v string) *DescribePeakTrendRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribePeakTrendRequest) SetInstanceId(v string) *DescribePeakTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetInterval(v string) *DescribePeakTrendRequest {
	s.Interval = &v
	return s
}

func (s *DescribePeakTrendRequest) SetRegionId(v string) *DescribePeakTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetResource(v string) *DescribePeakTrendRequest {
	s.Resource = &v
	return s
}

func (s *DescribePeakTrendRequest) SetResourceManagerResourceGroupId(v string) *DescribePeakTrendRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetStartTimestamp(v string) *DescribePeakTrendRequest {
	s.StartTimestamp = &v
	return s
}

type DescribePeakTrendResponseBody struct {
	// An array of the QPS statistics of the WAF instance.
	FlowChart []*DescribePeakTrendResponseBodyFlowChart `json:"FlowChart,omitempty" xml:"FlowChart,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePeakTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponseBody) SetFlowChart(v []*DescribePeakTrendResponseBodyFlowChart) *DescribePeakTrendResponseBody {
	s.FlowChart = v
	return s
}

func (s *DescribePeakTrendResponseBody) SetRequestId(v string) *DescribePeakTrendResponseBody {
	s.RequestId = &v
	return s
}

type DescribePeakTrendResponseBodyFlowChart struct {
	// The number of requests that are monitored or blocked by the custom rule (access control) module.
	AclSum *int64 `json:"AclSum,omitempty" xml:"AclSum,omitempty"`
	// The number of requests that are monitored or blocked by the scan protection module.
	AntiScanSum *int64 `json:"AntiScanSum,omitempty" xml:"AntiScanSum,omitempty"`
	// The number of requests that are monitored or blocked by the HTTP flood protection module.
	CcSum *int64 `json:"CcSum,omitempty" xml:"CcSum,omitempty"`
	// The total number of requests.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The serial number of the time interval. The serial numbers are arranged in chronological order.
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of requests that are monitored or blocked by the regular expression protection engine.
	WafSum *int64 `json:"WafSum,omitempty" xml:"WafSum,omitempty"`
}

func (s DescribePeakTrendResponseBodyFlowChart) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendResponseBodyFlowChart) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetAclSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.AclSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetAntiScanSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.AntiScanSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetCcSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.CcSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetCount(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.Count = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetIndex(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.Index = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetWafSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.WafSum = &v
	return s
}

type DescribePeakTrendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePeakTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePeakTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponse) SetHeaders(v map[string]*string) *DescribePeakTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribePeakTrendResponse) SetStatusCode(v int32) *DescribePeakTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePeakTrendResponse) SetBody(v *DescribePeakTrendResponseBody) *DescribePeakTrendResponse {
	s.Body = v
	return s
}

type DescribeProductInstancesRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The public IP address of the instance.
	ResourceIp *string `json:"ResourceIp,omitempty" xml:"ResourceIp,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The cloud service to which the instance belongs. Valid values:
	//
	// *   **clb4**: Layer 4 Classic Load Balancer (CLB).
	// *   **clb7**: Layer 7 CLB.
	// *   **ecs**: Elastic Compute Service (ECS).
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the instance. Valid values:
	//
	// *   **cn-chengdu**: China (Chengdu).
	// *   **cn-beijing**: China (Beijing).
	// *   **cn-zhangjiakou**: China (Zhangjiakou).
	// *   **cn-hangzhou**: China (Hangzhou).
	// *   **cn-shanghai**: China (Shanghai).
	// *   **cn-shenzhen**: China (Shenzhen).
	// *   **cn-qingdao**: China (Qingdao).
	// *   **cn-hongkong**: China (Hong Kong).
	// *   **ap-southeast-3**: Malaysia (Kuala Lumpur).
	// *   **ap-southeast-5**: Indonesia (Jakarta).
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeProductInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesRequest) SetInstanceId(v string) *DescribeProductInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetOwnerUserId(v string) *DescribeProductInstancesRequest {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetPageNumber(v int64) *DescribeProductInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetPageSize(v int64) *DescribeProductInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetRegionId(v string) *DescribeProductInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceInstanceId(v string) *DescribeProductInstancesRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceIp(v string) *DescribeProductInstancesRequest {
	s.ResourceIp = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceManagerResourceGroupId(v string) *DescribeProductInstancesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceName(v string) *DescribeProductInstancesRequest {
	s.ResourceName = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceProduct(v string) *DescribeProductInstancesRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeProductInstancesRequest) SetResourceRegionId(v string) *DescribeProductInstancesRequest {
	s.ResourceRegionId = &v
	return s
}

type DescribeProductInstancesResponseBody struct {
	// The information about the instances.
	ProductInstances []*DescribeProductInstancesResponseBodyProductInstances `json:"ProductInstances,omitempty" xml:"ProductInstances,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBody) SetProductInstances(v []*DescribeProductInstancesResponseBodyProductInstances) *DescribeProductInstancesResponseBody {
	s.ProductInstances = v
	return s
}

func (s *DescribeProductInstancesResponseBody) SetRequestId(v string) *DescribeProductInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductInstancesResponseBody) SetTotalCount(v int64) *DescribeProductInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProductInstancesResponseBodyProductInstances struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The ID of the instance.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The public IP address of the instance.
	ResourceIp *string `json:"ResourceIp,omitempty" xml:"ResourceIp,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The information about the ports.
	ResourcePorts []*DescribeProductInstancesResponseBodyProductInstancesResourcePorts `json:"ResourcePorts,omitempty" xml:"ResourcePorts,omitempty" type:"Repeated"`
	// The cloud service to which the instance belongs. Valid values:
	//
	// *   **clb4**: Layer 4 CLB.
	// *   **clb7**: Layer 7 CLB.
	// *   **ecs**: ECS.
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the instance. Valid values:
	//
	// *   **cn-chengdu**: China (Chengdu).
	// *   **cn-beijing**: China (Beijing).
	// *   **cn-zhangjiakou**: China (Zhangjiakou).
	// *   **cn-hangzhou**: China (Hangzhou).
	// *   **cn-shanghai**: China (Shanghai).
	// *   **cn-shenzhen**: China (Shenzhen).
	// *   **cn-qingdao**: China (Qingdao).
	// *   **cn-hongkong**: China (Hong Kong).
	// *   **ap-southeast-3**: Malaysia (Kuala Lumpur).
	// *   **ap-southeast-5**: Indonesia (Jakarta).
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstances) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetOwnerUserId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceInstanceId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceIp(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceIp = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceName(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceName = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourcePorts(v []*DescribeProductInstancesResponseBodyProductInstancesResourcePorts) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourcePorts = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceProduct(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceRegionId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceRegionId = &v
	return s
}

type DescribeProductInstancesResponseBodyProductInstancesResourcePorts struct {
	// The information about the certificates.
	Certificates []*DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The port number.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// *   **http**
	// *   **https**
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePorts) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) SetCertificates(v []*DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) *DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	s.Certificates = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) SetPort(v int32) *DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	s.Port = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) SetProtocol(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	s.Protocol = &v
	return s
}

type DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates struct {
	// The ID of the certificate.
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The name of the certificate.
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) SetCertificateId(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	s.CertificateId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) SetCertificateName(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	s.CertificateName = &v
	return s
}

type DescribeProductInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponse) SetHeaders(v map[string]*string) *DescribeProductInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductInstancesResponse) SetStatusCode(v int32) *DescribeProductInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductInstancesResponse) SetBody(v *DescribeProductInstancesResponseBody) *DescribeProductInstancesResponse {
	s.Body = v
	return s
}

type DescribePunishedDomainsRequest struct {
	// The domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribePunishedDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePunishedDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribePunishedDomainsRequest) SetDomains(v []*string) *DescribePunishedDomainsRequest {
	s.Domains = v
	return s
}

func (s *DescribePunishedDomainsRequest) SetInstanceId(v string) *DescribePunishedDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePunishedDomainsRequest) SetRegionId(v string) *DescribePunishedDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePunishedDomainsRequest) SetResourceManagerResourceGroupId(v string) *DescribePunishedDomainsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribePunishedDomainsResponseBody struct {
	// The domain names that are penalized for failing to obtain an ICP filing.
	PunishedDomains []*string `json:"PunishedDomains,omitempty" xml:"PunishedDomains,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePunishedDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePunishedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePunishedDomainsResponseBody) SetPunishedDomains(v []*string) *DescribePunishedDomainsResponseBody {
	s.PunishedDomains = v
	return s
}

func (s *DescribePunishedDomainsResponseBody) SetRequestId(v string) *DescribePunishedDomainsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePunishedDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePunishedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePunishedDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePunishedDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribePunishedDomainsResponse) SetHeaders(v map[string]*string) *DescribePunishedDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribePunishedDomainsResponse) SetStatusCode(v int32) *DescribePunishedDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePunishedDomainsResponse) SetBody(v *DescribePunishedDomainsResponseBody) *DescribePunishedDomainsResponse {
	s.Body = v
	return s
}

type DescribeResourceInstanceCertsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourceInstanceCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInstanceCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsRequest) SetInstanceId(v string) *DescribeResourceInstanceCertsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetPageNumber(v int64) *DescribeResourceInstanceCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetPageSize(v int64) *DescribeResourceInstanceCertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetRegionId(v string) *DescribeResourceInstanceCertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetResourceInstanceId(v string) *DescribeResourceInstanceCertsRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceInstanceCertsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeResourceInstanceCertsResponseBody struct {
	// The certificates.
	Certs []*DescribeResourceInstanceCertsResponseBodyCerts `json:"Certs,omitempty" xml:"Certs,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceInstanceCertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInstanceCertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsResponseBody) SetCerts(v []*DescribeResourceInstanceCertsResponseBodyCerts) *DescribeResourceInstanceCertsResponseBody {
	s.Certs = v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBody) SetRequestId(v string) *DescribeResourceInstanceCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBody) SetTotalCount(v int64) *DescribeResourceInstanceCertsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeResourceInstanceCertsResponseBodyCerts struct {
	// The time when the certificate expires.
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The time when the certificate was issued.
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The globally unique ID of the certificate. The value is in the "Certificate ID-cn-hangzhou" format. For example, if the ID of the certificate is 123, the value of CertIdentifier is 123-cn-hangzhou.
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name.
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The domain name for which the certificate is issued.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the certificate chain is complete.
	IsChainCompleted *bool `json:"IsChainCompleted,omitempty" xml:"IsChainCompleted,omitempty"`
}

func (s DescribeResourceInstanceCertsResponseBodyCerts) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInstanceCertsResponseBodyCerts) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetAfterDate(v int64) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.AfterDate = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetBeforeDate(v int64) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.BeforeDate = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetCertIdentifier(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetCertName(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.CertName = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetCommonName(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.CommonName = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetDomain(v string) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.Domain = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponseBodyCerts) SetIsChainCompleted(v bool) *DescribeResourceInstanceCertsResponseBodyCerts {
	s.IsChainCompleted = &v
	return s
}

type DescribeResourceInstanceCertsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceInstanceCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceInstanceCertsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceInstanceCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsResponse) SetHeaders(v map[string]*string) *DescribeResourceInstanceCertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceInstanceCertsResponse) SetStatusCode(v int32) *DescribeResourceInstanceCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponse) SetBody(v *DescribeResourceInstanceCertsResponseBody) *DescribeResourceInstanceCertsResponse {
	s.Body = v
	return s
}

type DescribeResourceLogStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The protected object that you want to query. You can specify multiple protected objects. Separate the protected objects with commas (,).
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DescribeResourceLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusRequest) SetInstanceId(v string) *DescribeResourceLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetRegionId(v string) *DescribeResourceLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetResources(v string) *DescribeResourceLogStatusRequest {
	s.Resources = &v
	return s
}

type DescribeResourceLogStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*DescribeResourceLogStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeResourceLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBody) SetRequestId(v string) *DescribeResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBody) SetResult(v []*DescribeResourceLogStatusResponseBodyResult) *DescribeResourceLogStatusResponseBody {
	s.Result = v
	return s
}

type DescribeResourceLogStatusResponseBodyResult struct {
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Indicates whether the log collection feature is enabled for the protected object. Valid values:
	//
	// *   **true:** The log collection feature is enabled.
	// *   **false:** The log collection feature is disabled.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetResource(v string) *DescribeResourceLogStatusResponseBodyResult {
	s.Resource = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetStatus(v bool) *DescribeResourceLogStatusResponseBodyResult {
	s.Status = &v
	return s
}

type DescribeResourceLogStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponse) SetHeaders(v map[string]*string) *DescribeResourceLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogStatusResponse) SetStatusCode(v int32) *DescribeResourceLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogStatusResponse) SetBody(v *DescribeResourceLogStatusResponseBody) *DescribeResourceLogStatusResponse {
	s.Body = v
	return s
}

type DescribeResourcePortRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud service instance.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourcePortRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePortRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortRequest) SetInstanceId(v string) *DescribeResourcePortRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetRegionId(v string) *DescribeResourcePortRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetResourceInstanceId(v string) *DescribeResourcePortRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourcePortRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeResourcePortResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of HTTP and HTTPS listener ports that are added to the WAF instance.
	ResourcePorts []*string `json:"ResourcePorts,omitempty" xml:"ResourcePorts,omitempty" type:"Repeated"`
}

func (s DescribeResourcePortResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePortResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortResponseBody) SetRequestId(v string) *DescribeResourcePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePortResponseBody) SetResourcePorts(v []*string) *DescribeResourcePortResponseBody {
	s.ResourcePorts = v
	return s
}

type DescribeResourcePortResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcePortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcePortResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePortResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortResponse) SetHeaders(v map[string]*string) *DescribeResourcePortResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePortResponse) SetStatusCode(v int32) *DescribeResourcePortResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePortResponse) SetBody(v *DescribeResourcePortResponseBody) *DescribeResourcePortResponse {
	s.Body = v
	return s
}

type DescribeResourceRegionIdRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourceRegionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceRegionIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceRegionIdRequest) SetInstanceId(v string) *DescribeResourceRegionIdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceRegionIdRequest) SetRegionId(v string) *DescribeResourceRegionIdRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceRegionIdRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceRegionIdRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeResourceRegionIdResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The region IDs.
	ResourceRegionIds []*string `json:"ResourceRegionIds,omitempty" xml:"ResourceRegionIds,omitempty" type:"Repeated"`
}

func (s DescribeResourceRegionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceRegionIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceRegionIdResponseBody) SetRequestId(v string) *DescribeResourceRegionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceRegionIdResponseBody) SetResourceRegionIds(v []*string) *DescribeResourceRegionIdResponseBody {
	s.ResourceRegionIds = v
	return s
}

type DescribeResourceRegionIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceRegionIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceRegionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceRegionIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceRegionIdResponse) SetHeaders(v map[string]*string) *DescribeResourceRegionIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceRegionIdResponse) SetStatusCode(v int32) *DescribeResourceRegionIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceRegionIdResponse) SetBody(v *DescribeResourceRegionIdResponseBody) *DescribeResourceRegionIdResponse {
	s.Body = v
	return s
}

type DescribeResourceSupportRegionsRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourceSupportRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceSupportRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceSupportRegionsRequest) SetInstanceId(v string) *DescribeResourceSupportRegionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceSupportRegionsRequest) SetRegionId(v string) *DescribeResourceSupportRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceSupportRegionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceSupportRegionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeResourceSupportRegionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The region IDs.
	SupportRegions []*string `json:"SupportRegions,omitempty" xml:"SupportRegions,omitempty" type:"Repeated"`
}

func (s DescribeResourceSupportRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceSupportRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceSupportRegionsResponseBody) SetRequestId(v string) *DescribeResourceSupportRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceSupportRegionsResponseBody) SetSupportRegions(v []*string) *DescribeResourceSupportRegionsResponseBody {
	s.SupportRegions = v
	return s
}

type DescribeResourceSupportRegionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceSupportRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceSupportRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceSupportRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceSupportRegionsResponse) SetHeaders(v map[string]*string) *DescribeResourceSupportRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceSupportRegionsResponse) SetStatusCode(v int32) *DescribeResourceSupportRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceSupportRegionsResponse) SetBody(v *DescribeResourceSupportRegionsResponseBody) *DescribeResourceSupportRegionsResponse {
	s.Body = v
	return s
}

type DescribeResponseCodeTrendGraphRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds. The value must be an integral multiple of 60.
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The type of the error codes. Valid values:
	//
	// *   **waf:** error codes that are returned to clients from WAF.
	// *   **upstream:** error codes that are returned to WAF from the origin server.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResponseCodeTrendGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphRequest) SetEndTimestamp(v string) *DescribeResponseCodeTrendGraphRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetInstanceId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetInterval(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Interval = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetRegionId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetResource(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Resource = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetResourceManagerResourceGroupId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetStartTimestamp(v string) *DescribeResponseCodeTrendGraphRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetType(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Type = &v
	return s
}

type DescribeResponseCodeTrendGraphResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the statistics of the error codes.
	ResponseCodes []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes `json:"ResponseCodes,omitempty" xml:"ResponseCodes,omitempty" type:"Repeated"`
}

func (s DescribeResponseCodeTrendGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponseBody) SetRequestId(v string) *DescribeResponseCodeTrendGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBody) SetResponseCodes(v []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes) *DescribeResponseCodeTrendGraphResponseBody {
	s.ResponseCodes = v
	return s
}

type DescribeResponseCodeTrendGraphResponseBodyResponseCodes struct {
	// The number of 302 error codes that are returned.
	Code302Pv *int64 `json:"302Pv,omitempty" xml:"302Pv,omitempty"`
	// The number of 405 error codes that are returned.
	Code405Pv *int64 `json:"405Pv,omitempty" xml:"405Pv,omitempty"`
	// The number of 444 error codes that are returned.
	Code444Pv *int64 `json:"444Pv,omitempty" xml:"444Pv,omitempty"`
	// The number of 499 error codes that are returned.
	Code499Pv *int64 `json:"499Pv,omitempty" xml:"499Pv,omitempty"`
	// The number of 5xx error codes that are returned.
	Code5xxPv *int64 `json:"5xxPv,omitempty" xml:"5xxPv,omitempty"`
	// The serial number of the time interval. The serial numbers are arranged in chronological order.
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s DescribeResponseCodeTrendGraphResponseBodyResponseCodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode302Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code302Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode405Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code405Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode444Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code444Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode499Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code499Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode5xxPv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code5xxPv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetIndex(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Index = &v
	return s
}

type DescribeResponseCodeTrendGraphResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResponseCodeTrendGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResponseCodeTrendGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponse) SetHeaders(v map[string]*string) *DescribeResponseCodeTrendGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) SetStatusCode(v int32) *DescribeResponseCodeTrendGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) SetBody(v *DescribeResponseCodeTrendGraphResponseBody) *DescribeResponseCodeTrendGraphResponse {
	s.Body = v
	return s
}

type DescribeRuleGroupsRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the query condition. Valid values:
	//
	// *   **id:** queries regular expression rule groups by ID.
	// *   **name:** queries regular expression rule groups by name.
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The query condition.
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeRuleGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsRequest) SetInstanceId(v string) *DescribeRuleGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageNumber(v int32) *DescribeRuleGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageSize(v int32) *DescribeRuleGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetRegionId(v string) *DescribeRuleGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSearchType(v string) *DescribeRuleGroupsRequest {
	s.SearchType = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSearchValue(v string) *DescribeRuleGroupsRequest {
	s.SearchValue = &v
	return s
}

type DescribeRuleGroupsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of regular expression rule groups.
	RuleGroups []*DescribeRuleGroupsResponseBodyRuleGroups `json:"RuleGroups,omitempty" xml:"RuleGroups,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeRuleGroupsResponseBody) SetTotalCount(v int64) *DescribeRuleGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRuleGroupsResponseBodyRuleGroups struct {
	// The most recent time when the rule group was modified.
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the automatic update feature is enabled for the rule group.
	//
	// *   1: The automatic update feature is enabled for the rule group.
	// *   2: The automatic update feature is disabled for the rule group.
	IsSubscribe *int32 `json:"IsSubscribe,omitempty" xml:"IsSubscribe,omitempty"`
	// The ID of the rule group.
	//
	// *   0: The rule group is created from scratch.
	// *   1011: The rule group is a strict rule group.
	// *   1012: The rule group is a medium rule group.
	// *   1013: The rue group is a loose rule group.
	ParentRuleGroupId *int64 `json:"ParentRuleGroupId,omitempty" xml:"ParentRuleGroupId,omitempty"`
	// The ID of the regular expression rule group.
	RuleGroupId *int64 `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	// The name of the rule group.
	RuleGroupName *string `json:"RuleGroupName,omitempty" xml:"RuleGroupName,omitempty"`
	// The number of built-in rules in the rule group.
	RuleTotalCount *int32 `json:"RuleTotalCount,omitempty" xml:"RuleTotalCount,omitempty"`
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetGmtModified(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetIsSubscribe(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.IsSubscribe = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetParentRuleGroupId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.ParentRuleGroupId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupName(v string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupName = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleTotalCount(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleTotalCount = &v
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

type DescribeRuleHitsTopClientIpRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of rules that are triggered by the protected object. By default, this parameter is not specified and all types of rules are queried.
	//
	// *   **blacklist:** IP address blacklist rules.
	// *   **custom:** custom rules.
	// *   **antiscan:** scan protection rules.
	// *   **cc_system:** HTTP flood protection rules.
	// *   **region_block:** region blacklist rules.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopClientIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopClientIpRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetInstanceId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetRegionId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetResource(v string) *DescribeRuleHitsTopClientIpRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetRuleType(v string) *DescribeRuleHitsTopClientIpRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopClientIpRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopClientIpResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 IP addresses from which attacks are initiated.
	RuleHitsTopClientIp []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp `json:"RuleHitsTopClientIp,omitempty" xml:"RuleHitsTopClientIp,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopClientIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponseBody) SetRequestId(v string) *DescribeRuleHitsTopClientIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBody) SetRuleHitsTopClientIp(v []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) *DescribeRuleHitsTopClientIpResponseBody {
	s.RuleHitsTopClientIp = v
	return s
}

type DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp struct {
	// The IP address of the service client.
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The number of attacks that are initiated from the IP address.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) SetClientIp(v string) *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	s.ClientIp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) SetCount(v int64) *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	s.Count = &v
	return s
}

type DescribeRuleHitsTopClientIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopClientIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopClientIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopClientIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) SetStatusCode(v int32) *DescribeRuleHitsTopClientIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) SetBody(v *DescribeRuleHitsTopClientIpResponseBody) *DescribeRuleHitsTopClientIpResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopResourceRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of rules that are triggered by the protected object. By default, this parameter is not specified and all types of rules are queried.
	//
	// *   **blacklist:** IP address blacklist rules.
	// *   **custom:** custom rules.
	// *   **antiscan:** scan protection rules.
	// *   **cc_system:** HTTP flood protection rules.
	// *   **region_block:** region blacklist rules.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopResourceRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetInstanceId(v string) *DescribeRuleHitsTopResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetRegionId(v string) *DescribeRuleHitsTopResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetRuleType(v string) *DescribeRuleHitsTopResourceRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopResourceRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopResourceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 protected objects that trigger protection rules.
	RuleHitsTopResource []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource `json:"RuleHitsTopResource,omitempty" xml:"RuleHitsTopResource,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponseBody) SetRequestId(v string) *DescribeRuleHitsTopResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBody) SetRuleHitsTopResource(v []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) *DescribeRuleHitsTopResourceResponseBody {
	s.RuleHitsTopResource = v
	return s
}

type DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource struct {
	// The number of requests that match protection rules.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) SetCount(v int64) *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) SetResource(v string) *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	s.Resource = &v
	return s
}

type DescribeRuleHitsTopResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) SetStatusCode(v int32) *DescribeRuleHitsTopResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) SetBody(v *DescribeRuleHitsTopResourceResponseBody) *DescribeRuleHitsTopResourceResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopRuleIdRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether protected objects that trigger protection rules are returned in the response. Valid values
	//
	// - **true**: returns only the number of times each protection rule is triggered. If you set IsGroupResource to true, Resource is left empty.
	// - **false**: returns the number of times each protection rule is triggered by each protected object.
	IsGroupResource *string `json:"IsGroupResource,omitempty" xml:"IsGroupResource,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of rules that are triggered by the protected object. By default, this parameter is not specified and all types of rules are queried.
	//
	// *   **blacklist:** IP address blacklist rules.
	// *   **custom:** custom rules.
	// *   **antiscan:** scan protection rules.
	// *   **cc_system:** HTTP flood protection rules.
	// *   **region_block:** region blacklist rules.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetInstanceId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetIsGroupResource(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.IsGroupResource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetRegionId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetResource(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetRuleType(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopRuleIdResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the IDs of the top 10 rules that are matched by requests.
	RuleHitsTopRuleId []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId `json:"RuleHitsTopRuleId,omitempty" xml:"RuleHitsTopRuleId,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopRuleIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) SetRequestId(v string) *DescribeRuleHitsTopRuleIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) SetRuleHitsTopRuleId(v []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) *DescribeRuleHitsTopRuleIdResponseBody {
	s.RuleHitsTopRuleId = v
	return s
}

type DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId struct {
	// The number of requests that match the rule.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetCount(v int64) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetResource(v string) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetRuleId(v string) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.RuleId = &v
	return s
}

type DescribeRuleHitsTopRuleIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopRuleIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopRuleIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetStatusCode(v int32) *DescribeRuleHitsTopRuleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetBody(v *DescribeRuleHitsTopRuleIdResponseBody) *DescribeRuleHitsTopRuleIdResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopTuleTypeRequest struct {
	// The end point of the time period for which to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou**: the Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The start point of the time period for which to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetInstanceId(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetRegionId(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetResource(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopTuleTypeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The top 10 protection modules that are matched.
	RuleHitsTopTuleType []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType `json:"RuleHitsTopTuleType,omitempty" xml:"RuleHitsTopTuleType,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopTuleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) SetRequestId(v string) *DescribeRuleHitsTopTuleTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) SetRuleHitsTopTuleType(v []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) *DescribeRuleHitsTopTuleTypeResponseBody {
	s.RuleHitsTopTuleType = v
	return s
}

type DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType struct {
	// The number of requests that match protection rules.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of rule that is matched. By default, this parameter is not returned. This indicates that all types of rules that are matched are returned.
	//
	// *   **waf:** basic protection rules.
	// *   **blacklist:** IP address blacklist rules.
	// *   **custom:** custom rules.
	// *   **antiscan:** scan protection rules.
	// *   **cc_system:** HTTP flood protection rules.
	// *   **region_block:** region blacklist rules.
	// *   **scene:** bot management rules.
	// *   **dlp:** data leakage prevention rules.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) SetCount(v int64) *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) SetRuleType(v string) *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	s.RuleType = &v
	return s
}

type DescribeRuleHitsTopTuleTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopTuleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopTuleTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetStatusCode(v int32) *DescribeRuleHitsTopTuleTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetBody(v *DescribeRuleHitsTopTuleTypeResponseBody) *DescribeRuleHitsTopTuleTypeResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopUaRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopUaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopUaRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetInstanceId(v string) *DescribeRuleHitsTopUaRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetRegionId(v string) *DescribeRuleHitsTopUaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetResource(v string) *DescribeRuleHitsTopUaRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopUaRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopUaRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopUaResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 user agents that are used to initiate attacks.
	RuleHitsTopUa []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa `json:"RuleHitsTopUa,omitempty" xml:"RuleHitsTopUa,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopUaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponseBody) SetRequestId(v string) *DescribeRuleHitsTopUaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBody) SetRuleHitsTopUa(v []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) *DescribeRuleHitsTopUaResponseBody {
	s.RuleHitsTopUa = v
	return s
}

type DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa struct {
	// The number of attacks that are initiated from the IP address.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The user agent.
	Ua *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) SetCount(v int64) *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) SetUa(v string) *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	s.Ua = &v
	return s
}

type DescribeRuleHitsTopUaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopUaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopUaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopUaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) SetStatusCode(v int32) *DescribeRuleHitsTopUaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) SetBody(v *DescribeRuleHitsTopUaResponseBody) *DescribeRuleHitsTopUaResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopUrlRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of rules that are triggered by the protected object. By default, this parameter is not specified and all types of rules are queried.
	//
	// *   **blacklist:** IP address blacklist rules.
	// *   **custom:** custom rules.
	// *   **antiscan:** scan protection rules.
	// *   **cc_system:** HTTP flood protection rules.
	// *   **region_block:** region blacklist rules.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopUrlRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetInstanceId(v string) *DescribeRuleHitsTopUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetRegionId(v string) *DescribeRuleHitsTopUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetResource(v string) *DescribeRuleHitsTopUrlRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopUrlRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetRuleType(v string) *DescribeRuleHitsTopUrlRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopUrlRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopUrlResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 URLs that trigger protection rules.
	RuleHitsTopUrl []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl `json:"RuleHitsTopUrl,omitempty" xml:"RuleHitsTopUrl,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponseBody) SetRequestId(v string) *DescribeRuleHitsTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBody) SetRuleHitsTopUrl(v []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) *DescribeRuleHitsTopUrlResponseBody {
	s.RuleHitsTopUrl = v
	return s
}

type DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl struct {
	// The number of requests from the URL that match protection rules.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request URL.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) SetCount(v int64) *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) SetUrl(v string) *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	s.Url = &v
	return s
}

type DescribeRuleHitsTopUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) SetStatusCode(v int32) *DescribeRuleHitsTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) SetBody(v *DescribeRuleHitsTopUrlResponseBody) *DescribeRuleHitsTopUrlResponse {
	s.Body = v
	return s
}

type DescribeSlsAuthStatusRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) SetInstanceId(v string) *DescribeSlsAuthStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetRegionId(v string) *DescribeSlsAuthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeSlsAuthStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether WAF is authorized to access Logstores. Valid values:
	//
	// *   **true**
	// *   **false**
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSlsAuthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponseBody) SetRequestId(v string) *DescribeSlsAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsAuthStatusResponseBody) SetStatus(v bool) *DescribeSlsAuthStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeSlsAuthStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsAuthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetStatusCode(v int32) *DescribeSlsAuthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetBody(v *DescribeSlsAuthStatusResponseBody) *DescribeSlsAuthStatusResponse {
	s.Body = v
	return s
}

type DescribeSlsLogStoreRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSlsLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogStoreRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreRequest) SetInstanceId(v string) *DescribeSlsLogStoreRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlsLogStoreRequest) SetRegionId(v string) *DescribeSlsLogStoreRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlsLogStoreRequest) SetResourceManagerResourceGroupId(v string) *DescribeSlsLogStoreRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeSlsLogStoreResponseBody struct {
	// The name of the Logstore.
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the Simple Log Service project.
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The capacity of the Logstore. Unit: bytes.
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage duration of the Logstore. Unit: days.
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The used capacity of the Logstore. Unit: bytes.
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeSlsLogStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreResponseBody) SetLogStoreName(v string) *DescribeSlsLogStoreResponseBody {
	s.LogStoreName = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetProjectName(v string) *DescribeSlsLogStoreResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetQuota(v int64) *DescribeSlsLogStoreResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetRequestId(v string) *DescribeSlsLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetTtl(v int32) *DescribeSlsLogStoreResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeSlsLogStoreResponseBody) SetUsed(v int64) *DescribeSlsLogStoreResponseBody {
	s.Used = &v
	return s
}

type DescribeSlsLogStoreResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogStoreResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreResponse) SetHeaders(v map[string]*string) *DescribeSlsLogStoreResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogStoreResponse) SetStatusCode(v int32) *DescribeSlsLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsLogStoreResponse) SetBody(v *DescribeSlsLogStoreResponseBody) *DescribeSlsLogStoreResponse {
	s.Body = v
	return s
}

type DescribeSlsLogStoreStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSlsLogStoreStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogStoreStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreStatusRequest) SetInstanceId(v string) *DescribeSlsLogStoreStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlsLogStoreStatusRequest) SetRegionId(v string) *DescribeSlsLogStoreStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlsLogStoreStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeSlsLogStoreStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeSlsLogStoreStatusResponseBody struct {
	// Indicates whether a Logstore is created for WAF. Valid values:
	//
	// *   **true**
	// *   **false**
	ExistStatus *bool `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlsLogStoreStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogStoreStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreStatusResponseBody) SetExistStatus(v bool) *DescribeSlsLogStoreStatusResponseBody {
	s.ExistStatus = &v
	return s
}

func (s *DescribeSlsLogStoreStatusResponseBody) SetRequestId(v string) *DescribeSlsLogStoreStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSlsLogStoreStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsLogStoreStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsLogStoreStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogStoreStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsLogStoreStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogStoreStatusResponse) SetStatusCode(v int32) *DescribeSlsLogStoreStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsLogStoreStatusResponse) SetBody(v *DescribeSlsLogStoreStatusResponseBody) *DescribeSlsLogStoreStatusResponse {
	s.Body = v
	return s
}

type DescribeTemplateResourceCountRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The IDs of the protection templates that you want to query. Separate multiple template IDs with commas (,).
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
}

func (s DescribeTemplateResourceCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourceCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountRequest) SetInstanceId(v string) *DescribeTemplateResourceCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) SetRegionId(v string) *DescribeTemplateResourceCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) SetResourceManagerResourceGroupId(v string) *DescribeTemplateResourceCountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeTemplateResourceCountRequest) SetTemplateIds(v string) *DescribeTemplateResourceCountRequest {
	s.TemplateIds = &v
	return s
}

type DescribeTemplateResourceCountResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of protected objects or protected object groups for which the protection template takes effect.
	ResourceCount []*DescribeTemplateResourceCountResponseBodyResourceCount `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" type:"Repeated"`
}

func (s DescribeTemplateResourceCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourceCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountResponseBody) SetRequestId(v string) *DescribeTemplateResourceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBody) SetResourceCount(v []*DescribeTemplateResourceCountResponseBodyResourceCount) *DescribeTemplateResourceCountResponseBody {
	s.ResourceCount = v
	return s
}

type DescribeTemplateResourceCountResponseBodyResourceCount struct {
	// The number of protected object groups.
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The number of protected objects.
	SingleCount *int32 `json:"SingleCount,omitempty" xml:"SingleCount,omitempty"`
	// The ID of the protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateResourceCountResponseBodyResourceCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourceCountResponseBodyResourceCount) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetGroupCount(v int32) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.GroupCount = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetSingleCount(v int32) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.SingleCount = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetTemplateId(v int64) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.TemplateId = &v
	return s
}

type DescribeTemplateResourceCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateResourceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateResourceCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourceCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountResponse) SetHeaders(v map[string]*string) *DescribeTemplateResourceCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResourceCountResponse) SetStatusCode(v int32) *DescribeTemplateResourceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateResourceCountResponse) SetBody(v *DescribeTemplateResourceCountResponseBody) *DescribeTemplateResourceCountResponse {
	s.Body = v
	return s
}

type DescribeTemplateResourcesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protected resource. Valid values:
	//
	// *   **single:** protected object.
	// *   **group:** protected object group.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesRequest) SetInstanceId(v string) *DescribeTemplateResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetRegionId(v string) *DescribeTemplateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeTemplateResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetResourceType(v string) *DescribeTemplateResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetTemplateId(v int64) *DescribeTemplateResourcesRequest {
	s.TemplateId = &v
	return s
}

type DescribeTemplateResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of protected objects or protected object groups that are associated to the protection rule template.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeTemplateResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesResponseBody) SetRequestId(v string) *DescribeTemplateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetResources(v []*string) *DescribeTemplateResourcesResponseBody {
	s.Resources = v
	return s
}

type DescribeTemplateResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesResponse) SetHeaders(v map[string]*string) *DescribeTemplateResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResourcesResponse) SetStatusCode(v int32) *DescribeTemplateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateResourcesResponse) SetBody(v *DescribeTemplateResourcesResponseBody) *DescribeTemplateResourcesResponse {
	s.Body = v
	return s
}

type DescribeUserSlsLogRegionsRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou:** Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeUserSlsLogRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSlsLogRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserSlsLogRegionsRequest) SetInstanceId(v string) *DescribeUserSlsLogRegionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserSlsLogRegionsRequest) SetRegionId(v string) *DescribeUserSlsLogRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserSlsLogRegionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserSlsLogRegionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeUserSlsLogRegionsResponseBody struct {
	// The region IDs.
	LogRegions []*string `json:"LogRegions,omitempty" xml:"LogRegions,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserSlsLogRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSlsLogRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserSlsLogRegionsResponseBody) SetLogRegions(v []*string) *DescribeUserSlsLogRegionsResponseBody {
	s.LogRegions = v
	return s
}

func (s *DescribeUserSlsLogRegionsResponseBody) SetRequestId(v string) *DescribeUserSlsLogRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserSlsLogRegionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserSlsLogRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserSlsLogRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserSlsLogRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserSlsLogRegionsResponse) SetHeaders(v map[string]*string) *DescribeUserSlsLogRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserSlsLogRegionsResponse) SetStatusCode(v int32) *DescribeUserSlsLogRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserSlsLogRegionsResponse) SetBody(v *DescribeUserSlsLogRegionsResponseBody) *DescribeUserSlsLogRegionsResponse {
	s.Body = v
	return s
}

type DescribeUserWafLogStatusRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeUserWafLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWafLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserWafLogStatusRequest) SetInstanceId(v string) *DescribeUserWafLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserWafLogStatusRequest) SetRegionId(v string) *DescribeUserWafLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserWafLogStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserWafLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeUserWafLogStatusResponseBody struct {
	// The ID of the region where WAF logs are stored. Valid values:
	//
	// *   **cn-hangzhou**: China (Hangzhou).
	// *   **cn-beijing**: China (Beijing).
	// *   **cn-hongkong**: China (Hong Kong).
	// *   **ap-southeast-1**: Singapore.
	// *   **ap-southeast-2**: Australia (Sydney).
	// *   **ap-southeast-3**: Malaysia (Kuala Lumpur).
	// *   **ap-southeast-5**: Indonesia (Jakarta).
	// *   **ap-southeast-6**: Philippines (Manila).
	// *   **ap-southeast-7**: Thailand (Bangkok).
	// *   **me-east-1**: UAE (Dubai).
	// *   **eu-central-1**: Germany (Frankfurt).
	// *   **us-east-1**: US (Virginia).
	// *   **us-west-1**: US (Silicon Valley).
	// *   **ap-northeast-1**: Japan (Tokyo).
	// *   **ap-northeast-2**: South Korea (Seoul).
	// *   **ap-south-1**: India (Mumbai).
	// *   **eu-west-1**: UK (London).
	// *   **cn-hangzhou-finance**: China East 1 Finance.
	// *   **cn-shanghai-finance-1**: China East 2 Finance.
	// *   **cn-shenzhen-finance**: China South 1 Finance.
	//
	// >  The China East 1 Finance, China East 2 Finance, and China South 1 Finance regions are available only for Alibaba Finance Cloud users. Alibaba Finance Cloud users are also limited to storing logs within these specific regions.
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The status of WAF logs.
	//
	// *   **initializing**
	// *   **initialize_failed**
	// *   **normal**
	// *   **releasing**
	// *   **release_failed**
	LogStatus *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the log status was modified. Unit: milliseconds.
	StatusUpdateTime *int64 `json:"StatusUpdateTime,omitempty" xml:"StatusUpdateTime,omitempty"`
}

func (s DescribeUserWafLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWafLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserWafLogStatusResponseBody) SetLogRegionId(v string) *DescribeUserWafLogStatusResponseBody {
	s.LogRegionId = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) SetLogStatus(v string) *DescribeUserWafLogStatusResponseBody {
	s.LogStatus = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) SetRequestId(v string) *DescribeUserWafLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserWafLogStatusResponseBody) SetStatusUpdateTime(v int64) *DescribeUserWafLogStatusResponseBody {
	s.StatusUpdateTime = &v
	return s
}

type DescribeUserWafLogStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserWafLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserWafLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWafLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserWafLogStatusResponse) SetHeaders(v map[string]*string) *DescribeUserWafLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserWafLogStatusResponse) SetStatusCode(v int32) *DescribeUserWafLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserWafLogStatusResponse) SetBody(v *DescribeUserWafLogStatusResponseBody) *DescribeUserWafLogStatusResponse {
	s.Body = v
	return s
}

type DescribeVisitTopIpRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeVisitTopIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpRequest) SetEndTimestamp(v string) *DescribeVisitTopIpRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetInstanceId(v string) *DescribeVisitTopIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetRegionId(v string) *DescribeVisitTopIpRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetResource(v string) *DescribeVisitTopIpRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetResourceManagerResourceGroupId(v string) *DescribeVisitTopIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetStartTimestamp(v string) *DescribeVisitTopIpRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeVisitTopIpResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 IP addresses from which requests are sent.
	TopIp []*DescribeVisitTopIpResponseBodyTopIp `json:"TopIp,omitempty" xml:"TopIp,omitempty" type:"Repeated"`
}

func (s DescribeVisitTopIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponseBody) SetRequestId(v string) *DescribeVisitTopIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVisitTopIpResponseBody) SetTopIp(v []*DescribeVisitTopIpResponseBodyTopIp) *DescribeVisitTopIpResponseBody {
	s.TopIp = v
	return s
}

type DescribeVisitTopIpResponseBodyTopIp struct {
	// The ordinal number of the area to which the IP address belongs.
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The total number of requests that are sent from the IP address.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IP address.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ISP.
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeVisitTopIpResponseBodyTopIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpResponseBodyTopIp) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetArea(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Area = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetCount(v int64) *DescribeVisitTopIpResponseBodyTopIp {
	s.Count = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetIp(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Ip = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetIsp(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Isp = &v
	return s
}

type DescribeVisitTopIpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVisitTopIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVisitTopIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponse) SetHeaders(v map[string]*string) *DescribeVisitTopIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeVisitTopIpResponse) SetStatusCode(v int32) *DescribeVisitTopIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVisitTopIpResponse) SetBody(v *DescribeVisitTopIpResponseBody) *DescribeVisitTopIpResponse {
	s.Body = v
	return s
}

type DescribeVisitUasRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeVisitUasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasRequest) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasRequest) SetEndTimestamp(v string) *DescribeVisitUasRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeVisitUasRequest) SetInstanceId(v string) *DescribeVisitUasRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVisitUasRequest) SetRegionId(v string) *DescribeVisitUasRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVisitUasRequest) SetResource(v string) *DescribeVisitUasRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVisitUasRequest) SetStartTimestamp(v string) *DescribeVisitUasRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeVisitUasResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the top 10 user agents that are used to initiate requests.
	Uas []*DescribeVisitUasResponseBodyUas `json:"Uas,omitempty" xml:"Uas,omitempty" type:"Repeated"`
}

func (s DescribeVisitUasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponseBody) SetRequestId(v string) *DescribeVisitUasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVisitUasResponseBody) SetUas(v []*DescribeVisitUasResponseBodyUas) *DescribeVisitUasResponseBody {
	s.Uas = v
	return s
}

type DescribeVisitUasResponseBodyUas struct {
	// The number of requests that use the user agent.
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The user agent.
	Ua *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s DescribeVisitUasResponseBodyUas) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasResponseBodyUas) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponseBodyUas) SetCount(v int64) *DescribeVisitUasResponseBodyUas {
	s.Count = &v
	return s
}

func (s *DescribeVisitUasResponseBodyUas) SetUa(v string) *DescribeVisitUasResponseBodyUas {
	s.Ua = &v
	return s
}

type DescribeVisitUasResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVisitUasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVisitUasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasResponse) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponse) SetHeaders(v map[string]*string) *DescribeVisitUasResponse {
	s.Headers = v
	return s
}

func (s *DescribeVisitUasResponse) SetStatusCode(v int32) *DescribeVisitUasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVisitUasResponse) SetBody(v *DescribeVisitUasResponseBody) *DescribeVisitUasResponse {
	s.Body = v
	return s
}

type DescribeWafSourceIpSegmentRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
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

func (s *DescribeWafSourceIpSegmentRequest) SetResourceManagerResourceGroupId(v string) *DescribeWafSourceIpSegmentRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type DescribeWafSourceIpSegmentResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The back-to-origin CIDR blocks that are used by the protection cluster.
	WafSourceIp *DescribeWafSourceIpSegmentResponseBodyWafSourceIp `json:"WafSourceIp,omitempty" xml:"WafSourceIp,omitempty" type:"Struct"`
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

func (s *DescribeWafSourceIpSegmentResponseBody) SetWafSourceIp(v *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) *DescribeWafSourceIpSegmentResponseBody {
	s.WafSourceIp = v
	return s
}

type DescribeWafSourceIpSegmentResponseBodyWafSourceIp struct {
	// An array of back-to-origin IPv4 CIDR blocks.
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	// An array of back-to-origin IPv6 CIDR blocks.
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s DescribeWafSourceIpSegmentResponseBodyWafSourceIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBodyWafSourceIp) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) SetIPv4(v []*string) *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	s.IPv4 = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) SetIPv6(v []*string) *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	s.IPv6 = v
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

type ListTagKeysRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Set the value to ALIYUN::WAF::DEFENSERESOURCE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetInstanceId(v string) *ListTagKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The keys and types of the tags.
	Keys []*ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponseBodyKeys struct {
	// The type of the tag. Valid values:
	//
	// *   custom
	// *   system
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) SetCategory(v string) *ListTagKeysResponseBodyKeys {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v string) *ListTagKeysResponseBodyKeys {
	s.Key = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to ALIYUN::WAF::DEFENSERESOURCE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that are added to the resource.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N that is added to the resource. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the resource. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. ALIYUN::WAF::DEFENSERESOURCE is returned.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N that is added to the resource.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N that is added to the resource.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Set the value to ALIYUN::WAF::DEFENSERESOURCE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ModifyDefenseResourceGroupRequest struct {
	// The protected objects that you want to add to the protected object group. Separate the protected objects with commas (,). If you leave this parameter empty, no protected objects are added to the protected object group.
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	// The protected objects that you want to remove from the protected object group. Separate the protected objects with commas (,). If you leave this parameter empty, no protected objects are removed from the protected object group.
	DeleteList *string `json:"DeleteList,omitempty" xml:"DeleteList,omitempty"`
	// The description of the protected object group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the protected object group whose configurations you want to modify.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupRequest) SetAddList(v string) *ModifyDefenseResourceGroupRequest {
	s.AddList = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetDeleteList(v string) *ModifyDefenseResourceGroupRequest {
	s.DeleteList = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetDescription(v string) *ModifyDefenseResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetGroupName(v string) *ModifyDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetInstanceId(v string) *ModifyDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetRegionId(v string) *ModifyDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type ModifyDefenseResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupResponseBody) SetRequestId(v string) *ModifyDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) SetStatusCode(v int32) *ModifyDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) SetBody(v *ModifyDefenseResourceGroupResponseBody) *ModifyDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDefenseResourceXffRequest struct {
	// The status of the tracking cookie.
	//
	// *   **0**: disabled
	// *   **1**: enabled. This is the default value.
	AcwCookieStatus *int32 `json:"AcwCookieStatus,omitempty" xml:"AcwCookieStatus,omitempty"`
	// The status of the secure attribute of the tracking cookie.
	//
	// *   **0**: disabled. This is the default value.
	// *   **1**: enabled.
	AcwSecureStatus *int32 `json:"AcwSecureStatus,omitempty" xml:"AcwSecureStatus,omitempty"`
	// The status of the secure attribute of the slider CAPTCHA cookie.
	//
	// *   **0**: disabled. This is the default value.
	// *   **1**: enabled.
	AcwV3SecureStatus *int32 `json:"AcwV3SecureStatus,omitempty" xml:"AcwV3SecureStatus,omitempty"`
	// The custom header fields.
	//
	// >  The first IP address in the specified custom header field is used as the originating IP address of the client to prevent X-Forwarded-For forgery. If you specify multiple header fields, WAF reads the values of the header fields in sequence until the originating IP address is obtained. If the originating IP address cannot be obtained, the first IP address in the X-Forwarded-For header is used as the originating IP address of the client.
	CustomHeaders []*string `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Specifies whether a Layer 7 proxy is deployed in front of WAF. Layer 7 proxies include Anti-DDoS Proxy and Alibaba Cloud CDN. Valid values:
	//
	// *   **0**: No Layer 7 proxies are deployed. This is the default value.
	// *   **1**: A Layer 7 proxy is deployed.
	XffStatus *int32 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s ModifyDefenseResourceXffRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceXffRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffRequest) SetAcwCookieStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.AcwCookieStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetAcwSecureStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.AcwSecureStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetAcwV3SecureStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.AcwV3SecureStatus = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetCustomHeaders(v []*string) *ModifyDefenseResourceXffRequest {
	s.CustomHeaders = v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetInstanceId(v string) *ModifyDefenseResourceXffRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetRegionId(v string) *ModifyDefenseResourceXffRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetResource(v string) *ModifyDefenseResourceXffRequest {
	s.Resource = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceXffRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseResourceXffRequest) SetXffStatus(v int32) *ModifyDefenseResourceXffRequest {
	s.XffStatus = &v
	return s
}

type ModifyDefenseResourceXffResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseResourceXffResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceXffResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffResponseBody) SetRequestId(v string) *ModifyDefenseResourceXffResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseResourceXffResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseResourceXffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseResourceXffResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceXffResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffResponse) SetHeaders(v map[string]*string) *ModifyDefenseResourceXffResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseResourceXffResponse) SetStatusCode(v int32) *ModifyDefenseResourceXffResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseResourceXffResponse) SetBody(v *ModifyDefenseResourceXffResponseBody) *ModifyDefenseResourceXffResponse {
	s.Body = v
	return s
}

type ModifyDefenseRuleRequest struct {
	// The scenario in which you want to use the protection rule. For more information, see the description of the **DefenseScene** parameter in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The details of the protection rule. Specify a string that contains multiple parameters in the JSON format. You must specify the ID and the new configurations of the protection rule.
	//
	// *   **id:** The ID of the protection rule. Data type: long. You must specify this parameter.
	// *   The protection rule configurations: The role of this parameter is the same as that of the **Rules** parameter in the **CreateDefenseRule** topic. For more information, see the "**Protection rule parameters**" section in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the protection rule template to which the protection rule whose configurations you want to modify belongs.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleRequest) SetDefenseScene(v string) *ModifyDefenseRuleRequest {
	s.DefenseScene = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetInstanceId(v string) *ModifyDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetRegionId(v string) *ModifyDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetRules(v string) *ModifyDefenseRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetTemplateId(v int64) *ModifyDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type ModifyDefenseRuleResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleResponseBody) SetRequestId(v string) *ModifyDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleResponse) SetStatusCode(v int32) *ModifyDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleResponse) SetBody(v *ModifyDefenseRuleResponseBody) *ModifyDefenseRuleResponse {
	s.Body = v
	return s
}

type ModifyDefenseRuleCacheRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the protection template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleCacheRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleCacheRequest) SetInstanceId(v string) *ModifyDefenseRuleCacheRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetRegionId(v string) *ModifyDefenseRuleCacheRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleCacheRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetRuleId(v int64) *ModifyDefenseRuleCacheRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetTemplateId(v int64) *ModifyDefenseRuleCacheRequest {
	s.TemplateId = &v
	return s
}

type ModifyDefenseRuleCacheResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleCacheResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleCacheResponseBody) SetRequestId(v string) *ModifyDefenseRuleCacheResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseRuleCacheResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseRuleCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseRuleCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleCacheResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleCacheResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleCacheResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleCacheResponse) SetStatusCode(v int32) *ModifyDefenseRuleCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleCacheResponse) SetBody(v *ModifyDefenseRuleCacheResponseBody) *ModifyDefenseRuleCacheResponse {
	s.Body = v
	return s
}

type ModifyDefenseRuleStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule whose status you want to change.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The new status of the protection rule. Valid values:
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The ID of the protection rule template to which the protection rule whose status you want to change belongs.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusRequest) SetInstanceId(v string) *ModifyDefenseRuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRegionId(v string) *ModifyDefenseRuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRuleId(v int64) *ModifyDefenseRuleStatusRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRuleStatus(v int32) *ModifyDefenseRuleStatusRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetTemplateId(v int64) *ModifyDefenseRuleStatusRequest {
	s.TemplateId = &v
	return s
}

type ModifyDefenseRuleStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusResponseBody) SetRequestId(v string) *ModifyDefenseRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseRuleStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseRuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) SetStatusCode(v int32) *ModifyDefenseRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) SetBody(v *ModifyDefenseRuleStatusResponseBody) *ModifyDefenseRuleStatusResponse {
	s.Body = v
	return s
}

type ModifyDefenseTemplateRequest struct {
	// The description of the protection rule template whose configurations you want to modify.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template whose configurations you want to modify.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection rule template whose configurations you want to modify.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateRequest) SetDescription(v string) *ModifyDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetInstanceId(v string) *ModifyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetRegionId(v string) *ModifyDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateId(v int64) *ModifyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateName(v string) *ModifyDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

type ModifyDefenseTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateResponseBody) SetRequestId(v string) *ModifyDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateResponse) SetHeaders(v map[string]*string) *ModifyDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseTemplateResponse) SetStatusCode(v int32) *ModifyDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseTemplateResponse) SetBody(v *ModifyDefenseTemplateResponseBody) *ModifyDefenseTemplateResponse {
	s.Body = v
	return s
}

type ModifyDefenseTemplateStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template whose status you want to change.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The new status of the protection rule template. Valid values:
	//
	// *   **0:** disabled.
	// *   **1:** enabled.
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
}

func (s ModifyDefenseTemplateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusRequest) SetInstanceId(v string) *ModifyDefenseTemplateStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetRegionId(v string) *ModifyDefenseTemplateStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetTemplateId(v int64) *ModifyDefenseTemplateStatusRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetTemplateStatus(v int32) *ModifyDefenseTemplateStatusRequest {
	s.TemplateStatus = &v
	return s
}

type ModifyDefenseTemplateStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseTemplateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusResponseBody) SetRequestId(v string) *ModifyDefenseTemplateStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseTemplateStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseTemplateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseTemplateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusResponse) SetHeaders(v map[string]*string) *ModifyDefenseTemplateStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) SetStatusCode(v int32) *ModifyDefenseTemplateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) SetBody(v *ModifyDefenseTemplateStatusResponseBody) *ModifyDefenseTemplateStatusResponse {
	s.Body = v
	return s
}

type ModifyDomainRequest struct {
	// The mode in which you want to add the domain name to WAF. Set the value to share.
	//
	// *   **share:** adds the domain name to WAF in CNAME record mode. This is the default value.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name whose access configurations you want to modify.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configurations of the listeners.
	Listen *ModifyDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The configurations of the forwarding rule.
	Redirect *ModifyDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequest) SetAccessType(v string) *ModifyDomainRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainRequest) SetDomain(v string) *ModifyDomainRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainRequest) SetInstanceId(v string) *ModifyDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainRequest) SetListen(v *ModifyDomainRequestListen) *ModifyDomainRequest {
	s.Listen = v
	return s
}

func (s *ModifyDomainRequest) SetRedirect(v *ModifyDomainRequestRedirect) *ModifyDomainRequest {
	s.Redirect = v
	return s
}

func (s *ModifyDomainRequest) SetRegionId(v string) *ModifyDomainRequest {
	s.RegionId = &v
	return s
}

type ModifyDomainRequestListen struct {
	// The ID of the certificate that you want to add.
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of cipher suite that you want to add. This parameter is available only when you specify the **HttpsPorts** parameter. Valid values:
	//
	// *   **1:** all cipher suites.
	// *   **2:** strong cipher suites. You can select this value only when you set the **TLSVersion** parameter to **tlsv1.2**.
	// *   **99:** custom cipher suites.
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites that you want to add. This parameter is available only when you set the **CipherSuite** parameter to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. This parameter is available only when you specify the **HttpsPorts** parameter. Valid values:
	//
	// *   **true:** supports TLS 1.3.
	// *   **false:** does not support TLS 1.3.
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable an exclusive IP address for the domain name. This parameter is available only when you set the **IPv6Enabled** parameter to false and the **ProtectionResource** parameter to **share**. Valid values:
	//
	// *   **true:** enables an exclusive IP address for the domain name.
	// *   **false:** does not enable an exclusive IP address for the domain name. This is the default value.
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable HTTP to HTTPS redirection for the domain name. This parameter is available only when you specify the **HttpsPorts** parameter and leave the **HttpPorts** parameter empty. Valid values:
	//
	// *   **true:** enables HTTP to HTTPS redirection.
	// *   **false:** disables HTTP to HTTPS redirection.
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter is available only when you specify the **HttpsPorts** parameter. Valid values:
	//
	// *   **true:** enables HTTP/2.
	// *   **false:** disables HTTP/2. This is the default value.
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// An array of HTTP listener ports. Specify the value of this parameter in the \[port1,port2,...] format.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// An array of HTTPS listener ports. Specify the value of this parameter in the \[port1,port2,...] format.
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// *   **true:** enables IPv6.
	// *   **false:** disables IPv6. This is the default value.
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of the protection resource that you want to use. Valid values:
	//
	// *   **share:** shared cluster. This is the default value.
	// *   **gslb:** shared cluster-based intelligent load balancing.
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Specifies whether to allow access only from SM certificate-based clients. This parameter is available only if you set SM2Enabled to true.
	//
	// *   true
	// *   false
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate that you want to add. This parameter is available only if you set SM2Enabled to true.
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Indicates whether SM certificate-based verification is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The version of the Transport Layer Security (TLS) protocol. This parameter is available only when you specify the **HttpsPorts** parameter. Valid values:
	//
	// *   **tlsv1**
	// *   **tlsv1.1**
	// *   **tlsv1.2**
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that you want WAF to use to obtain the actual IP address of a client. Valid values:
	//
	// *   **0:** No Layer 7 proxies are deployed in front of WAF. This is the default value.
	// *   **1:** WAF reads the first value of the X-Forwarded-For (XFF) header field as the actual IP address of the client.
	// *   **2:** WAF reads the value of a custom header field as the actual IP address of the client.
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that you want to use to obtain the actual IP address of a client. Specify the value of this parameter in the \["header1","header2",...] format.
	//
	// >  If you set the **XffHeaderMode** parameter to 2, this parameter is required.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s ModifyDomainRequestListen) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequestListen) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestListen) SetCertId(v string) *ModifyDomainRequestListen {
	s.CertId = &v
	return s
}

func (s *ModifyDomainRequestListen) SetCipherSuite(v int32) *ModifyDomainRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *ModifyDomainRequestListen) SetCustomCiphers(v []*string) *ModifyDomainRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *ModifyDomainRequestListen) SetEnableTLSv3(v bool) *ModifyDomainRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *ModifyDomainRequestListen) SetExclusiveIp(v bool) *ModifyDomainRequestListen {
	s.ExclusiveIp = &v
	return s
}

func (s *ModifyDomainRequestListen) SetFocusHttps(v bool) *ModifyDomainRequestListen {
	s.FocusHttps = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHttp2Enabled(v bool) *ModifyDomainRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHttpPorts(v []*int32) *ModifyDomainRequestListen {
	s.HttpPorts = v
	return s
}

func (s *ModifyDomainRequestListen) SetHttpsPorts(v []*int32) *ModifyDomainRequestListen {
	s.HttpsPorts = v
	return s
}

func (s *ModifyDomainRequestListen) SetIPv6Enabled(v bool) *ModifyDomainRequestListen {
	s.IPv6Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetProtectionResource(v string) *ModifyDomainRequestListen {
	s.ProtectionResource = &v
	return s
}

func (s *ModifyDomainRequestListen) SetSM2AccessOnly(v bool) *ModifyDomainRequestListen {
	s.SM2AccessOnly = &v
	return s
}

func (s *ModifyDomainRequestListen) SetSM2CertId(v string) *ModifyDomainRequestListen {
	s.SM2CertId = &v
	return s
}

func (s *ModifyDomainRequestListen) SetSM2Enabled(v bool) *ModifyDomainRequestListen {
	s.SM2Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetTLSVersion(v string) *ModifyDomainRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *ModifyDomainRequestListen) SetXffHeaderMode(v int32) *ModifyDomainRequestListen {
	s.XffHeaderMode = &v
	return s
}

func (s *ModifyDomainRequestListen) SetXffHeaders(v []*string) *ModifyDomainRequestListen {
	s.XffHeaders = v
	return s
}

type ModifyDomainRequestRedirect struct {
	// An array of the IP addresses or domain names of the origin servers. You can specify only one type of address. If you use the domain name type, only IPv4 is supported.
	//
	// *   If you use the IP address type, specify the value of this parameter in the \["ip1","ip2",...] format. You can add up to 20 IP addresses.
	// *   If you use the domain name type, specify the value of this parameter in the \["domain"] format. You can add up to 20 domain names.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// Specifies whether to enable the public cloud disaster recovery feature. Valid values:
	//
	// *   **true**
	// *   **false** (default)
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The connection timeout period. Unit: seconds. Valid values: 1 to 3600.
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Specifies whether to enable HTTPS to HTTP redirection for back-to-origin requests of the domain name. This parameter is available only when you specify the **HttpsPorts** parameter. Valid values:
	//
	// *   **true:** enables HTTPS to HTTP redirection for back-to-origin requests of the domain name.
	// *   **false:** disables HTTPS to HTTP redirection for back-to-origin requests of the domain name.
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Specifies whether to enable the persistent connection feature. Valid values:
	//
	// *   **true:** enables the persistent connection feature. This is the default value.
	// *   **false:** disables the persistent connection feature.
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter specifies the number of reused persistent connections when you enable the persistent connection feature.
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of persistent connections that are in the Idle state. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// >  This parameter specifies the period of time during which a reused persistent connection is allowed to remain in the Idle state before the persistent connection is released.
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that you want to use when WAF forwards requests to the origin server. Valid values:
	//
	// *   **ip_hash:** the IP hash algorithm.
	// *   **roundRobin:** the round-robin algorithm.
	// *   **leastTime:** the least response time algorithm. You can select this value only when you set the **ProtectionResource** parameter to **gslb**.
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The read timeout period. Unit: seconds. Valid values: 1 to 3600.
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The key-value pairs that you want to use to mark the requests that pass through the WAF instance.
	//
	// WAF automatically adds the key-value pairs to the request headers to identify the requests that pass through WAF.
	RequestHeaders []*ModifyDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Specifies whether WAF retries to forward requests when requests fail to be forwarded to the origin server. Valid values:
	//
	// *   **true:** WAF retries to forward requests. This is the default value.
	// *   **false:** WAF does not retry to forward requests.
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules that you want to configure for the domain name that you want to add to WAF in hybrid cloud mode. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// *   **rs**: the back-to-origin IP addresses or CNAMEs. The value must be of the ARRAY type.
	// *   **location**: the name of the protection node. The value must be of the STRING type.
	// *   **locationId**: the ID of the protection node. The value must be of the LONG type.
	RoutingRules *string `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	// Specifies whether to enable origin Server Name Indication (SNI). This parameter is available only when you specify the **HttpsPorts** parameter. Valid values:
	//
	// *   **true:** enables origin SNI.
	// *   **false:** disables origin SNI. This is the default value.
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI field. If you do not specify this parameter, the value of the **Host** field in the request header is automatically used. If you want WAF to use an SNI field value that is different from the value of the Host field in back-to-origin requests, you can specify a custom value for the SNI field.
	//
	// >  If you set the **SniEnabled** parameter to true, this parameter is required.
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// The write timeout period. Unit: seconds. Valid values: 1 to 3600.
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Indicates whether the X-Forward-For-Proto header is used to identify the protocol used by WAF to forward requests to the origin server. Valid values:
	//
	// *   **true** (default)
	// *   **false**
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s ModifyDomainRequestRedirect) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequestRedirect) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirect) SetBackends(v []*string) *ModifyDomainRequestRedirect {
	s.Backends = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetCnameEnabled(v bool) *ModifyDomainRequestRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetConnectTimeout(v int32) *ModifyDomainRequestRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetFocusHttpBackend(v bool) *ModifyDomainRequestRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepalive(v bool) *ModifyDomainRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepaliveRequests(v int32) *ModifyDomainRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepaliveTimeout(v int32) *ModifyDomainRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetLoadbalance(v string) *ModifyDomainRequestRedirect {
	s.Loadbalance = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetReadTimeout(v int32) *ModifyDomainRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRequestHeaders(v []*ModifyDomainRequestRedirectRequestHeaders) *ModifyDomainRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRetry(v bool) *ModifyDomainRequestRedirect {
	s.Retry = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRoutingRules(v string) *ModifyDomainRequestRedirect {
	s.RoutingRules = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetSniEnabled(v bool) *ModifyDomainRequestRedirect {
	s.SniEnabled = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetSniHost(v string) *ModifyDomainRequestRedirect {
	s.SniHost = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetWriteTimeout(v int32) *ModifyDomainRequestRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetXffProto(v bool) *ModifyDomainRequestRedirect {
	s.XffProto = &v
	return s
}

type ModifyDomainRequestRedirectRequestHeaders struct {
	// The key of the custom header field.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom header field.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDomainRequestRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirectRequestHeaders) SetKey(v string) *ModifyDomainRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *ModifyDomainRequestRedirectRequestHeaders) SetValue(v string) *ModifyDomainRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

type ModifyDomainShrinkRequest struct {
	// The mode in which you want to add the domain name to WAF. Set the value to share.
	//
	// *   **share:** adds the domain name to WAF in CNAME record mode. This is the default value.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name whose access configurations you want to modify.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configurations of the listeners.
	ListenShrink *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	// The configurations of the forwarding rule.
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDomainShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainShrinkRequest) SetAccessType(v string) *ModifyDomainShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetDomain(v string) *ModifyDomainShrinkRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetInstanceId(v string) *ModifyDomainShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetListenShrink(v string) *ModifyDomainShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetRedirectShrink(v string) *ModifyDomainShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetRegionId(v string) *ModifyDomainShrinkRequest {
	s.RegionId = &v
	return s
}

type ModifyDomainResponseBody struct {
	// The information about the domain name.
	DomainInfo *ModifyDomainResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBody) SetDomainInfo(v *ModifyDomainResponseBodyDomainInfo) *ModifyDomainResponseBody {
	s.DomainInfo = v
	return s
}

func (s *ModifyDomainResponseBody) SetRequestId(v string) *ModifyDomainResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainResponseBodyDomainInfo struct {
	// The CNAME that is assigned by WAF to the domain name.
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name whose access configurations you modified.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the domain name.
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s ModifyDomainResponseBodyDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBodyDomainInfo) SetCname(v string) *ModifyDomainResponseBodyDomainInfo {
	s.Cname = &v
	return s
}

func (s *ModifyDomainResponseBodyDomainInfo) SetDomain(v string) *ModifyDomainResponseBodyDomainInfo {
	s.Domain = &v
	return s
}

func (s *ModifyDomainResponseBodyDomainInfo) SetDomainId(v string) *ModifyDomainResponseBodyDomainInfo {
	s.DomainId = &v
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

type ModifyDomainPunishStatusRequest struct {
	// The domain name that is penalized for failing to obtain an ICP filing.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyDomainPunishStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainPunishStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainPunishStatusRequest) SetDomain(v string) *ModifyDomainPunishStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) SetInstanceId(v string) *ModifyDomainPunishStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) SetRegionId(v string) *ModifyDomainPunishStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyDomainPunishStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type ModifyDomainPunishStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainPunishStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainPunishStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainPunishStatusResponseBody) SetRequestId(v string) *ModifyDomainPunishStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainPunishStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDomainPunishStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDomainPunishStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainPunishStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainPunishStatusResponse) SetHeaders(v map[string]*string) *ModifyDomainPunishStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainPunishStatusResponse) SetStatusCode(v int32) *ModifyDomainPunishStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainPunishStatusResponse) SetBody(v *ModifyDomainPunishStatusResponseBody) *ModifyDomainPunishStatusResponse {
	s.Body = v
	return s
}

type ModifyHybridCloudClusterBypassStatusRequest struct {
	// The ID of the hybrid cloud cluster.
	ClusterResourceId *string `json:"ClusterResourceId,omitempty" xml:"ClusterResourceId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// **
	//
	// **You can call the **DescribeInstanceInfo[ operation to obtain the ID of the WAF instance.](~~140857~~)
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of manual bypass. Valid values:
	//
	// *   **on**: enabled.
	// *   **off**: disabled. This is the default value.
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ModifyHybridCloudClusterBypassStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHybridCloudClusterBypassStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetClusterResourceId(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.ClusterResourceId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetInstanceId(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusRequest) SetRuleStatus(v string) *ModifyHybridCloudClusterBypassStatusRequest {
	s.RuleStatus = &v
	return s
}

type ModifyHybridCloudClusterBypassStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudClusterBypassStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHybridCloudClusterBypassStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterBypassStatusResponseBody) SetRequestId(v string) *ModifyHybridCloudClusterBypassStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHybridCloudClusterBypassStatusResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudClusterBypassStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudClusterBypassStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHybridCloudClusterBypassStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudClusterBypassStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) SetStatusCode(v int32) *ModifyHybridCloudClusterBypassStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudClusterBypassStatusResponse) SetBody(v *ModifyHybridCloudClusterBypassStatusResponseBody) *ModifyHybridCloudClusterBypassStatusResponse {
	s.Body = v
	return s
}

type ModifyMajorProtectionBlackIpRequest struct {
	// The description of the IP address blacklist.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time after which the IP address blacklist becomes invalid. Unit: seconds.
	//
	// >  If you set this parameter to **0**, the blacklist is permanently valid.
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP addresses that you want to add to the IP address blacklist. You can specify multiple CIDR blocks or IP addresses. IPv4 and IPv6 addresses are supported. Separate the CIDR blocks or IP addresses with commas (,). For more information, see [Protection for major events](~~425591~~).
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the IP address blacklist rule for major event protection.
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the IP address blacklist rule template for major event protection.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpRequest) SetDescription(v string) *ModifyMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetInstanceId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetIpList(v string) *ModifyMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetRegionId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetRuleId(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetTemplateId(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type ModifyMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpResponseBody) SetRequestId(v string) *ModifyMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *ModifyMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) SetStatusCode(v int32) *ModifyMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) SetBody(v *ModifyMajorProtectionBlackIpResponseBody) *ModifyMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type ModifyMemberAccountRequest struct {
	// The description of the member. The description must be 1 to 256 characters in length, and can contain letters, digits, periods (.), underscores (\_), hyphens (-), and asterisks (\*).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID of the managed member.
	MemberAccountId *string `json:"MemberAccountId,omitempty" xml:"MemberAccountId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The source IP address of the request. The system automatically obtains the value of this parameter.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyMemberAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyMemberAccountRequest) SetDescription(v string) *ModifyMemberAccountRequest {
	s.Description = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetInstanceId(v string) *ModifyMemberAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetMemberAccountId(v string) *ModifyMemberAccountRequest {
	s.MemberAccountId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetRegionId(v string) *ModifyMemberAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetResourceManagerResourceGroupId(v string) *ModifyMemberAccountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyMemberAccountRequest) SetSourceIp(v string) *ModifyMemberAccountRequest {
	s.SourceIp = &v
	return s
}

type ModifyMemberAccountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMemberAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMemberAccountResponseBody) SetRequestId(v string) *ModifyMemberAccountResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMemberAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMemberAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMemberAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyMemberAccountResponse) SetHeaders(v map[string]*string) *ModifyMemberAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyMemberAccountResponse) SetStatusCode(v int32) *ModifyMemberAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMemberAccountResponse) SetBody(v *ModifyMemberAccountResponseBody) *ModifyMemberAccountResponse {
	s.Body = v
	return s
}

type ModifyResourceLogStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object on which you want to manage the log collection feature.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Specifies whether to enable the log collection feature for the protected object. Valid values:
	//
	// *   **true:** enables the log collection feature.
	// *   **false:** disables the log collection feature.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceLogStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusRequest) SetInstanceId(v string) *ModifyResourceLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetRegionId(v string) *ModifyResourceLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetResource(v string) *ModifyResourceLogStatusRequest {
	s.Resource = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyResourceLogStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetStatus(v bool) *ModifyResourceLogStatusRequest {
	s.Status = &v
	return s
}

type ModifyResourceLogStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the log collection feature is enabled for the protected object. Valid values:
	//
	// *   **true**
	// *   **false**
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusResponseBody) SetRequestId(v string) *ModifyResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceLogStatusResponseBody) SetStatus(v bool) *ModifyResourceLogStatusResponseBody {
	s.Status = &v
	return s
}

type ModifyResourceLogStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceLogStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusResponse) SetHeaders(v map[string]*string) *ModifyResourceLogStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceLogStatusResponse) SetStatusCode(v int32) *ModifyResourceLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceLogStatusResponse) SetBody(v *ModifyResourceLogStatusResponseBody) *ModifyResourceLogStatusResponse {
	s.Body = v
	return s
}

type ModifyTemplateResourcesRequest struct {
	// The protected object groups that you want to associate with the protection rule template. Specify the value of this parameter in the \["group1","group2",...] format.
	BindResourceGroups []*string `json:"BindResourceGroups,omitempty" xml:"BindResourceGroups,omitempty" type:"Repeated"`
	// The protected objects that you want to associate with the protection rule template. Specify the value of this parameter in the \["XX1","XX2",...] format.
	BindResources []*string `json:"BindResources,omitempty" xml:"BindResources,omitempty" type:"Repeated"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to obtain the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// *   **cn-hangzhou:** the Chinese mainland.
	// *   **ap-southeast-1:** outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template.
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The protected object groups that you want to disassociate from the protection rule template. Specify the value of this parameter in the \["group1","group2",...] format.
	UnbindResourceGroups []*string `json:"UnbindResourceGroups,omitempty" xml:"UnbindResourceGroups,omitempty" type:"Repeated"`
	// The protected objects that you want to disassociate from the protection rule template. Specify the value of this parameter in the \["XX1","XX2",...] format.
	UnbindResources []*string `json:"UnbindResources,omitempty" xml:"UnbindResources,omitempty" type:"Repeated"`
}

func (s ModifyTemplateResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesRequest) SetBindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetBindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResources = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetInstanceId(v string) *ModifyTemplateResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetRegionId(v string) *ModifyTemplateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetResourceManagerResourceGroupId(v string) *ModifyTemplateResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetTemplateId(v int64) *ModifyTemplateResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetUnbindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetUnbindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindResources = v
	return s
}

type ModifyTemplateResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesResponseBody) SetRequestId(v string) *ModifyTemplateResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTemplateResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResourcesResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesResponse) SetHeaders(v map[string]*string) *ModifyTemplateResourcesResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateResourcesResponse) SetStatusCode(v int32) *ModifyTemplateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateResourcesResponse) SetBody(v *ModifyTemplateResourcesResponseBody) *ModifyTemplateResourcesResponse {
	s.Body = v
	return s
}

type SyncProductInstanceRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](~~433756~~) operation to query the ID of the WAF instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s SyncProductInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncProductInstanceRequest) GoString() string {
	return s.String()
}

func (s *SyncProductInstanceRequest) SetInstanceId(v string) *SyncProductInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncProductInstanceRequest) SetRegionId(v string) *SyncProductInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *SyncProductInstanceRequest) SetResourceManagerResourceGroupId(v string) *SyncProductInstanceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

type SyncProductInstanceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncProductInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncProductInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SyncProductInstanceResponseBody) SetRequestId(v string) *SyncProductInstanceResponseBody {
	s.RequestId = &v
	return s
}

type SyncProductInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncProductInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncProductInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncProductInstanceResponse) GoString() string {
	return s.String()
}

func (s *SyncProductInstanceResponse) SetHeaders(v map[string]*string) *SyncProductInstanceResponse {
	s.Headers = v
	return s
}

func (s *SyncProductInstanceResponse) SetStatusCode(v int32) *SyncProductInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncProductInstanceResponse) SetBody(v *SyncProductInstanceResponseBody) *SyncProductInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources. You can specify up to 50 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to ALIYUN::WAF::DEFENSERESOURCE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add to the resource.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of tag N to add to the resource. Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource groups or members. Valid values:
	//
	// *   false (default)
	// *   true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// *   **cn-hangzhou**: Chinese mainland.
	// *   **ap-southeast-1**: outside the Chinese mainland.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to ALIYUN::WAF::DEFENSERESOURCE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. You can specify up to 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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

func (client *Client) ClearMajorProtectionBlackIpWithOptions(request *ClearMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *ClearMajorProtectionBlackIpResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearMajorProtectionBlackIp(request *ClearMajorProtectionBlackIpRequest) (_result *ClearMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearMajorProtectionBlackIpResponse{}
	_body, _err := client.ClearMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyDefenseTemplateWithOptions(request *CopyDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *CopyDefenseTemplateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyDefenseTemplate(request *CopyDefenseTemplateRequest) (_result *CopyDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyDefenseTemplateResponse{}
	_body, _err := client.CopyDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefenseResourceGroupWithOptions(request *CreateDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddList)) {
		query["AddList"] = request.AddList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDefenseResourceGroup(request *CreateDefenseResourceGroupRequest) (_result *CreateDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDefenseResourceGroupResponse{}
	_body, _err := client.CreateDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefenseRuleWithOptions(request *CreateDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDefenseRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDefenseRule(request *CreateDefenseRuleRequest) (_result *CreateDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDefenseRuleResponse{}
	_body, _err := client.CreateDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefenseTemplateWithOptions(request *CreateDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateDefenseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateOrigin)) {
		query["TemplateOrigin"] = request.TemplateOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateStatus)) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDefenseTemplate(request *CreateDefenseTemplateRequest) (_result *CreateDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDefenseTemplateResponse{}
	_body, _err := client.CreateDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(tmpReq *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Listen)) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, tea.String("Listen"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Redirect)) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, tea.String("Redirect"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenShrink)) {
		query["Listen"] = request.ListenShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectShrink)) {
		query["Redirect"] = request.RedirectShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2021-10-01"),
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

/**
 * This operation is available only on the China site (aliyun.com).
 *
 * @param request CreateMajorProtectionBlackIpRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateMajorProtectionBlackIpResponse
 */
func (client *Client) CreateMajorProtectionBlackIpWithOptions(request *CreateMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This operation is available only on the China site (aliyun.com).
 *
 * @param request CreateMajorProtectionBlackIpRequest
 * @return CreateMajorProtectionBlackIpResponse
 */
func (client *Client) CreateMajorProtectionBlackIp(request *CreateMajorProtectionBlackIpRequest) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CreateMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMemberAccountsWithOptions(request *CreateMemberAccountsRequest, runtime *util.RuntimeOptions) (_result *CreateMemberAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberAccountIds)) {
		query["MemberAccountIds"] = request.MemberAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMemberAccounts"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMemberAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMemberAccounts(request *CreateMemberAccountsRequest) (_result *CreateMemberAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMemberAccountsResponse{}
	_body, _err := client.CreateMemberAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePostpaidInstanceWithOptions(request *CreatePostpaidInstanceRequest, runtime *util.RuntimeOptions) (_result *CreatePostpaidInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePostpaidInstance"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePostpaidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePostpaidInstance(request *CreatePostpaidInstanceRequest) (_result *CreatePostpaidInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePostpaidInstanceResponse{}
	_body, _err := client.CreatePostpaidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDefenseResourceGroupWithOptions(request *DeleteDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDefenseResourceGroup(request *DeleteDefenseResourceGroupRequest) (_result *DeleteDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDefenseResourceGroupResponse{}
	_body, _err := client.DeleteDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDefenseRuleWithOptions(request *DeleteDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDefenseRuleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIds)) {
		query["RuleIds"] = request.RuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDefenseRule(request *DeleteDefenseRuleRequest) (_result *DeleteDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDefenseRuleResponse{}
	_body, _err := client.DeleteDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDefenseTemplateWithOptions(request *DeleteDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteDefenseTemplateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDefenseTemplate(request *DeleteDefenseTemplateRequest) (_result *DeleteDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDefenseTemplateResponse{}
	_body, _err := client.DeleteDefenseTemplateWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2021-10-01"),
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

func (client *Client) DeleteMajorProtectionBlackIpWithOptions(request *DeleteMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *DeleteMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMajorProtectionBlackIp(request *DeleteMajorProtectionBlackIpRequest) (_result *DeleteMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMajorProtectionBlackIpResponse{}
	_body, _err := client.DeleteMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMemberAccountWithOptions(request *DeleteMemberAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteMemberAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberAccountId)) {
		query["MemberAccountId"] = request.MemberAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMemberAccount"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMemberAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMemberAccount(request *DeleteMemberAccountRequest) (_result *DeleteMemberAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMemberAccountResponse{}
	_body, _err := client.DeleteMemberAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountDelegatedStatusWithOptions(request *DescribeAccountDelegatedStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountDelegatedStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountDelegatedStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountDelegatedStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountDelegatedStatus(request *DescribeAccountDelegatedStatusRequest) (_result *DescribeAccountDelegatedStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountDelegatedStatusResponse{}
	_body, _err := client.DescribeAccountDelegatedStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertDetailWithOptions(request *DescribeCertDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCertDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCertDetail"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCertDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCertDetail(request *DescribeCertDetailRequest) (_result *DescribeCertDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertDetailResponse{}
	_body, _err := client.DescribeCertDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertsWithOptions(request *DescribeCertsRequest, runtime *util.RuntimeOptions) (_result *DescribeCertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCerts"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCerts(request *DescribeCertsRequest) (_result *DescribeCertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertsResponse{}
	_body, _err := client.DescribeCertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudResourcesWithOptions(request *DescribeCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["OwnerUserId"] = request.OwnerUserId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceDomain)) {
		query["ResourceDomain"] = request.ResourceDomain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceFunction)) {
		query["ResourceFunction"] = request.ResourceFunction
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceInstanceId)) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceProduct)) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRouteName)) {
		query["ResourceRouteName"] = request.ResourceRouteName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudResources(request *DescribeCloudResourcesRequest) (_result *DescribeCloudResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudResourcesResponse{}
	_body, _err := client.DescribeCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceWithOptions(request *DescribeDefenseResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResource"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResource(request *DescribeDefenseResourceRequest) (_result *DescribeDefenseResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceResponse{}
	_body, _err := client.DescribeDefenseResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroupWithOptions(request *DescribeDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroup(request *DescribeDefenseResourceGroupRequest) (_result *DescribeDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupResponse{}
	_body, _err := client.DescribeDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroupNamesWithOptions(request *DescribeDefenseResourceGroupNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceGroupNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupNameLike)) {
		query["GroupNameLike"] = request.GroupNameLike
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResourceGroupNames"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceGroupNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroupNames(request *DescribeDefenseResourceGroupNamesRequest) (_result *DescribeDefenseResourceGroupNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupNamesResponse{}
	_body, _err := client.DescribeDefenseResourceGroupNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroupsWithOptions(request *DescribeDefenseResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupNameLike)) {
		query["GroupNameLike"] = request.GroupNameLike
	}

	if !tea.BoolValue(util.IsUnset(request.GroupNames)) {
		query["GroupNames"] = request.GroupNames
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResourceGroups"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroups(request *DescribeDefenseResourceGroupsRequest) (_result *DescribeDefenseResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupsResponse{}
	_body, _err := client.DescribeDefenseResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceNamesWithOptions(request *DescribeDefenseResourceNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResourceNames"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResourceNames(request *DescribeDefenseResourceNamesRequest) (_result *DescribeDefenseResourceNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceNamesResponse{}
	_body, _err := client.DescribeDefenseResourceNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceTemplatesWithOptions(request *DescribeDefenseResourceTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceTemplatesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResourceTemplates"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResourceTemplates(request *DescribeDefenseResourceTemplatesRequest) (_result *DescribeDefenseResourceTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceTemplatesResponse{}
	_body, _err := client.DescribeDefenseResourceTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourcesWithOptions(request *DescribeDefenseResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResources(request *DescribeDefenseResourcesRequest) (_result *DescribeDefenseResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourcesResponse{}
	_body, _err := client.DescribeDefenseResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseRuleWithOptions(request *DescribeDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseRuleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseRule(request *DescribeDefenseRuleRequest) (_result *DescribeDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseRuleResponse{}
	_body, _err := client.DescribeDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseRulesWithOptions(request *DescribeDefenseRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseRules"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseRules(request *DescribeDefenseRulesRequest) (_result *DescribeDefenseRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseRulesResponse{}
	_body, _err := client.DescribeDefenseRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseTemplateWithOptions(request *DescribeDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseTemplateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseTemplate(request *DescribeDefenseTemplateRequest) (_result *DescribeDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseTemplateResponse{}
	_body, _err := client.DescribeDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseTemplateValidGroupsWithOptions(request *DescribeDefenseTemplateValidGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseTemplateValidGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseTemplateValidGroups"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseTemplateValidGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseTemplateValidGroups(request *DescribeDefenseTemplateValidGroupsRequest) (_result *DescribeDefenseTemplateValidGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseTemplateValidGroupsResponse{}
	_body, _err := client.DescribeDefenseTemplateValidGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseTemplatesWithOptions(request *DescribeDefenseTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.DefenseSubScene)) {
		query["DefenseSubScene"] = request.DefenseSubScene
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

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseTemplates"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseTemplates(request *DescribeDefenseTemplatesRequest) (_result *DescribeDefenseTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseTemplatesResponse{}
	_body, _err := client.DescribeDefenseTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainDNSRecordWithOptions(request *DescribeDomainDNSRecordRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainDNSRecordResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainDNSRecord"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainDNSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainDNSRecord(request *DescribeDomainDNSRecordRequest) (_result *DescribeDomainDNSRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainDNSRecordResponse{}
	_body, _err := client.DescribeDomainDNSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainDetailWithOptions(request *DescribeDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainDetailResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainDetail"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainDetail(request *DescribeDomainDetailRequest) (_result *DescribeDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.DescribeDomainDetailWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Backend)) {
		query["Backend"] = request.Backend
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomains"),
		Version:     tea.String("2021-10-01"),
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

func (client *Client) DescribeFlowChartWithOptions(request *DescribeFlowChartRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowChartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowChart"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowChart(request *DescribeFlowChartRequest) (_result *DescribeFlowChartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowChartResponse{}
	_body, _err := client.DescribeFlowChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowTopResourceWithOptions(request *DescribeFlowTopResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowTopResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowTopResource"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowTopResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowTopResource(request *DescribeFlowTopResourceRequest) (_result *DescribeFlowTopResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowTopResourceResponse{}
	_body, _err := client.DescribeFlowTopResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowTopUrlWithOptions(request *DescribeFlowTopUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowTopUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowTopUrl"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowTopUrl(request *DescribeFlowTopUrlRequest) (_result *DescribeFlowTopUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowTopUrlResponse{}
	_body, _err := client.DescribeFlowTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHybridCloudGroupsWithOptions(request *DescribeHybridCloudGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeHybridCloudGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterProxyType)) {
		query["ClusterProxyType"] = request.ClusterProxyType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHybridCloudGroups"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHybridCloudGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHybridCloudGroups(request *DescribeHybridCloudGroupsRequest) (_result *DescribeHybridCloudGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHybridCloudGroupsResponse{}
	_body, _err := client.DescribeHybridCloudGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHybridCloudResourcesWithOptions(request *DescribeHybridCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeHybridCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Backend)) {
		query["Backend"] = request.Backend
	}

	if !tea.BoolValue(util.IsUnset(request.CnameEnabled)) {
		query["CnameEnabled"] = request.CnameEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHybridCloudResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHybridCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHybridCloudResources(request *DescribeHybridCloudResourcesRequest) (_result *DescribeHybridCloudResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHybridCloudResourcesResponse{}
	_body, _err := client.DescribeHybridCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHybridCloudUserWithOptions(request *DescribeHybridCloudUserRequest, runtime *util.RuntimeOptions) (_result *DescribeHybridCloudUserResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHybridCloudUser"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHybridCloudUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHybridCloudUser(request *DescribeHybridCloudUserRequest) (_result *DescribeHybridCloudUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHybridCloudUserResponse{}
	_body, _err := client.DescribeHybridCloudUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMajorProtectionBlackIpsWithOptions(request *DescribeMajorProtectionBlackIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeMajorProtectionBlackIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpLike)) {
		query["IpLike"] = request.IpLike
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMajorProtectionBlackIps"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMajorProtectionBlackIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMajorProtectionBlackIps(request *DescribeMajorProtectionBlackIpsRequest) (_result *DescribeMajorProtectionBlackIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMajorProtectionBlackIpsResponse{}
	_body, _err := client.DescribeMajorProtectionBlackIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMemberAccountsWithOptions(request *DescribeMemberAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeMemberAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountStatus)) {
		query["AccountStatus"] = request.AccountStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMemberAccounts"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMemberAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMemberAccounts(request *DescribeMemberAccountsRequest) (_result *DescribeMemberAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMemberAccountsResponse{}
	_body, _err := client.DescribeMemberAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePeakTrendWithOptions(request *DescribePeakTrendRequest, runtime *util.RuntimeOptions) (_result *DescribePeakTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePeakTrend"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePeakTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePeakTrend(request *DescribePeakTrendRequest) (_result *DescribePeakTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePeakTrendResponse{}
	_body, _err := client.DescribePeakTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProductInstancesWithOptions(request *DescribeProductInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeProductInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		query["OwnerUserId"] = request.OwnerUserId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceInstanceId)) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIp)) {
		query["ResourceIp"] = request.ResourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceProduct)) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProductInstances"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProductInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProductInstances(request *DescribeProductInstancesRequest) (_result *DescribeProductInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProductInstancesResponse{}
	_body, _err := client.DescribeProductInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePunishedDomainsWithOptions(request *DescribePunishedDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribePunishedDomainsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePunishedDomains"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePunishedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePunishedDomains(request *DescribePunishedDomainsRequest) (_result *DescribePunishedDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePunishedDomainsResponse{}
	_body, _err := client.DescribePunishedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceInstanceCertsWithOptions(request *DescribeResourceInstanceCertsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceInstanceCertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.ResourceInstanceId)) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceInstanceCerts"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceInstanceCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceInstanceCerts(request *DescribeResourceInstanceCertsRequest) (_result *DescribeResourceInstanceCertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceInstanceCertsResponse{}
	_body, _err := client.DescribeResourceInstanceCertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceLogStatusWithOptions(request *DescribeResourceLogStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceLogStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceLogStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceLogStatus(request *DescribeResourceLogStatusRequest) (_result *DescribeResourceLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceLogStatusResponse{}
	_body, _err := client.DescribeResourceLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePortWithOptions(request *DescribeResourcePortRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePortResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceInstanceId)) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourcePort"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourcePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePort(request *DescribeResourcePortRequest) (_result *DescribeResourcePortResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePortResponse{}
	_body, _err := client.DescribeResourcePortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceRegionIdWithOptions(request *DescribeResourceRegionIdRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceRegionIdResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceRegionId"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceRegionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceRegionId(request *DescribeResourceRegionIdRequest) (_result *DescribeResourceRegionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceRegionIdResponse{}
	_body, _err := client.DescribeResourceRegionIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceSupportRegionsWithOptions(request *DescribeResourceSupportRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceSupportRegionsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceSupportRegions"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceSupportRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceSupportRegions(request *DescribeResourceSupportRegionsRequest) (_result *DescribeResourceSupportRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceSupportRegionsResponse{}
	_body, _err := client.DescribeResourceSupportRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResponseCodeTrendGraphWithOptions(request *DescribeResponseCodeTrendGraphRequest, runtime *util.RuntimeOptions) (_result *DescribeResponseCodeTrendGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResponseCodeTrendGraph"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResponseCodeTrendGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResponseCodeTrendGraph(request *DescribeResponseCodeTrendGraphRequest) (_result *DescribeResponseCodeTrendGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResponseCodeTrendGraphResponse{}
	_body, _err := client.DescribeResponseCodeTrendGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleGroupsWithOptions(request *DescribeRuleGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		query["SearchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleGroups"),
		Version:     tea.String("2021-10-01"),
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

func (client *Client) DescribeRuleHitsTopClientIpWithOptions(request *DescribeRuleHitsTopClientIpRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopClientIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopClientIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopClientIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopClientIp(request *DescribeRuleHitsTopClientIpRequest) (_result *DescribeRuleHitsTopClientIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopClientIpResponse{}
	_body, _err := client.DescribeRuleHitsTopClientIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopResourceWithOptions(request *DescribeRuleHitsTopResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopResource"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopResource(request *DescribeRuleHitsTopResourceRequest) (_result *DescribeRuleHitsTopResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopResourceResponse{}
	_body, _err := client.DescribeRuleHitsTopResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopRuleIdWithOptions(request *DescribeRuleHitsTopRuleIdRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopRuleIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsGroupResource)) {
		query["IsGroupResource"] = request.IsGroupResource
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopRuleId"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopRuleIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopRuleId(request *DescribeRuleHitsTopRuleIdRequest) (_result *DescribeRuleHitsTopRuleIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopRuleIdResponse{}
	_body, _err := client.DescribeRuleHitsTopRuleIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopTuleTypeWithOptions(request *DescribeRuleHitsTopTuleTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopTuleTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopTuleType"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopTuleTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopTuleType(request *DescribeRuleHitsTopTuleTypeRequest) (_result *DescribeRuleHitsTopTuleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopTuleTypeResponse{}
	_body, _err := client.DescribeRuleHitsTopTuleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUaWithOptions(request *DescribeRuleHitsTopUaRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopUaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopUa"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopUaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUa(request *DescribeRuleHitsTopUaRequest) (_result *DescribeRuleHitsTopUaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopUaResponse{}
	_body, _err := client.DescribeRuleHitsTopUaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUrlWithOptions(request *DescribeRuleHitsTopUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopUrl"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUrl(request *DescribeRuleHitsTopUrlRequest) (_result *DescribeRuleHitsTopUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopUrlResponse{}
	_body, _err := client.DescribeRuleHitsTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsAuthStatusWithOptions(request *DescribeSlsAuthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsAuthStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsAuthStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsAuthStatus(request *DescribeSlsAuthStatusRequest) (_result *DescribeSlsAuthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DescribeSlsAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsLogStoreWithOptions(request *DescribeSlsLogStoreRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsLogStoreResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsLogStore"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlsLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsLogStore(request *DescribeSlsLogStoreRequest) (_result *DescribeSlsLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsLogStoreResponse{}
	_body, _err := client.DescribeSlsLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsLogStoreStatusWithOptions(request *DescribeSlsLogStoreStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsLogStoreStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsLogStoreStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlsLogStoreStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsLogStoreStatus(request *DescribeSlsLogStoreStatusRequest) (_result *DescribeSlsLogStoreStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsLogStoreStatusResponse{}
	_body, _err := client.DescribeSlsLogStoreStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateResourceCountWithOptions(request *DescribeTemplateResourceCountRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateResourceCountResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIds)) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplateResourceCount"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplateResourceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplateResourceCount(request *DescribeTemplateResourceCountRequest) (_result *DescribeTemplateResourceCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateResourceCountResponse{}
	_body, _err := client.DescribeTemplateResourceCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateResourcesWithOptions(request *DescribeTemplateResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplateResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplateResources(request *DescribeTemplateResourcesRequest) (_result *DescribeTemplateResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateResourcesResponse{}
	_body, _err := client.DescribeTemplateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserSlsLogRegionsWithOptions(request *DescribeUserSlsLogRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserSlsLogRegionsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserSlsLogRegions"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserSlsLogRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserSlsLogRegions(request *DescribeUserSlsLogRegionsRequest) (_result *DescribeUserSlsLogRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserSlsLogRegionsResponse{}
	_body, _err := client.DescribeUserSlsLogRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserWafLogStatusWithOptions(request *DescribeUserWafLogStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserWafLogStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserWafLogStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserWafLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserWafLogStatus(request *DescribeUserWafLogStatusRequest) (_result *DescribeUserWafLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserWafLogStatusResponse{}
	_body, _err := client.DescribeUserWafLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVisitTopIpWithOptions(request *DescribeVisitTopIpRequest, runtime *util.RuntimeOptions) (_result *DescribeVisitTopIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVisitTopIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVisitTopIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVisitTopIp(request *DescribeVisitTopIpRequest) (_result *DescribeVisitTopIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVisitTopIpResponse{}
	_body, _err := client.DescribeVisitTopIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVisitUasWithOptions(request *DescribeVisitUasRequest, runtime *util.RuntimeOptions) (_result *DescribeVisitUasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVisitUas"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVisitUasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVisitUas(request *DescribeVisitUasRequest) (_result *DescribeVisitUasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVisitUasResponse{}
	_body, _err := client.DescribeVisitUasWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWafSourceIpSegment"),
		Version:     tea.String("2021-10-01"),
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

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseResourceGroupWithOptions(request *ModifyDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddList)) {
		query["AddList"] = request.AddList
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteList)) {
		query["DeleteList"] = request.DeleteList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseResourceGroup(request *ModifyDefenseResourceGroupRequest) (_result *ModifyDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseResourceGroupResponse{}
	_body, _err := client.ModifyDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseResourceXffWithOptions(request *ModifyDefenseResourceXffRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseResourceXffResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcwCookieStatus)) {
		query["AcwCookieStatus"] = request.AcwCookieStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AcwSecureStatus)) {
		query["AcwSecureStatus"] = request.AcwSecureStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AcwV3SecureStatus)) {
		query["AcwV3SecureStatus"] = request.AcwV3SecureStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CustomHeaders)) {
		query["CustomHeaders"] = request.CustomHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.XffStatus)) {
		query["XffStatus"] = request.XffStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseResourceXff"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseResourceXffResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseResourceXff(request *ModifyDefenseResourceXffRequest) (_result *ModifyDefenseResourceXffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseResourceXffResponse{}
	_body, _err := client.ModifyDefenseResourceXffWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseRuleWithOptions(request *ModifyDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseRule(request *ModifyDefenseRuleRequest) (_result *ModifyDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseRuleResponse{}
	_body, _err := client.ModifyDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseRuleCacheWithOptions(request *ModifyDefenseRuleCacheRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseRuleCacheResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseRuleCache"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseRuleCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseRuleCache(request *ModifyDefenseRuleCacheRequest) (_result *ModifyDefenseRuleCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseRuleCacheResponse{}
	_body, _err := client.ModifyDefenseRuleCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseRuleStatusWithOptions(request *ModifyDefenseRuleStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseRuleStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleStatus)) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseRuleStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseRuleStatus(request *ModifyDefenseRuleStatusRequest) (_result *ModifyDefenseRuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseRuleStatusResponse{}
	_body, _err := client.ModifyDefenseRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseTemplateWithOptions(request *ModifyDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseTemplate(request *ModifyDefenseTemplateRequest) (_result *ModifyDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseTemplateResponse{}
	_body, _err := client.ModifyDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseTemplateStatusWithOptions(request *ModifyDefenseTemplateStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseTemplateStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateStatus)) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseTemplateStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseTemplateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseTemplateStatus(request *ModifyDefenseTemplateStatusRequest) (_result *ModifyDefenseTemplateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseTemplateStatusResponse{}
	_body, _err := client.ModifyDefenseTemplateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainWithOptions(tmpReq *ModifyDomainRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Listen)) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, tea.String("Listen"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Redirect)) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, tea.String("Redirect"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenShrink)) {
		query["Listen"] = request.ListenShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectShrink)) {
		query["Redirect"] = request.RedirectShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomain"),
		Version:     tea.String("2021-10-01"),
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

func (client *Client) ModifyDomainPunishStatusWithOptions(request *ModifyDomainPunishStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainPunishStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomainPunishStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDomainPunishStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainPunishStatus(request *ModifyDomainPunishStatusRequest) (_result *ModifyDomainPunishStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainPunishStatusResponse{}
	_body, _err := client.ModifyDomainPunishStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHybridCloudClusterBypassStatusWithOptions(request *ModifyHybridCloudClusterBypassStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyHybridCloudClusterBypassStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterResourceId)) {
		query["ClusterResourceId"] = request.ClusterResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleStatus)) {
		query["RuleStatus"] = request.RuleStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHybridCloudClusterBypassStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHybridCloudClusterBypassStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHybridCloudClusterBypassStatus(request *ModifyHybridCloudClusterBypassStatusRequest) (_result *ModifyHybridCloudClusterBypassStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHybridCloudClusterBypassStatusResponse{}
	_body, _err := client.ModifyHybridCloudClusterBypassStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMajorProtectionBlackIpWithOptions(request *ModifyMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *ModifyMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMajorProtectionBlackIp(request *ModifyMajorProtectionBlackIpRequest) (_result *ModifyMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMajorProtectionBlackIpResponse{}
	_body, _err := client.ModifyMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMemberAccountWithOptions(request *ModifyMemberAccountRequest, runtime *util.RuntimeOptions) (_result *ModifyMemberAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberAccountId)) {
		query["MemberAccountId"] = request.MemberAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMemberAccount"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMemberAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMemberAccount(request *ModifyMemberAccountRequest) (_result *ModifyMemberAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMemberAccountResponse{}
	_body, _err := client.ModifyMemberAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyResourceLogStatusWithOptions(request *ModifyResourceLogStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyResourceLogStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyResourceLogStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyResourceLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyResourceLogStatus(request *ModifyResourceLogStatusRequest) (_result *ModifyResourceLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyResourceLogStatusResponse{}
	_body, _err := client.ModifyResourceLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTemplateResourcesWithOptions(request *ModifyTemplateResourcesRequest, runtime *util.RuntimeOptions) (_result *ModifyTemplateResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindResourceGroups)) {
		query["BindResourceGroups"] = request.BindResourceGroups
	}

	if !tea.BoolValue(util.IsUnset(request.BindResources)) {
		query["BindResources"] = request.BindResources
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UnbindResourceGroups)) {
		query["UnbindResourceGroups"] = request.UnbindResourceGroups
	}

	if !tea.BoolValue(util.IsUnset(request.UnbindResources)) {
		query["UnbindResources"] = request.UnbindResources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTemplateResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTemplateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTemplateResources(request *ModifyTemplateResourcesRequest) (_result *ModifyTemplateResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTemplateResourcesResponse{}
	_body, _err := client.ModifyTemplateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * SyncProductInstance is an asynchronous operation. You can call the [DescribeProductInstances](~~2743168~~) operation to query the status of the task.
 *
 * @param request SyncProductInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SyncProductInstanceResponse
 */
func (client *Client) SyncProductInstanceWithOptions(request *SyncProductInstanceRequest, runtime *util.RuntimeOptions) (_result *SyncProductInstanceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceManagerResourceGroupId)) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncProductInstance"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncProductInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * SyncProductInstance is an asynchronous operation. You can call the [DescribeProductInstances](~~2743168~~) operation to query the status of the task.
 *
 * @param request SyncProductInstanceRequest
 * @return SyncProductInstanceResponse
 */
func (client *Client) SyncProductInstance(request *SyncProductInstanceRequest) (_result *SyncProductInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncProductInstanceResponse{}
	_body, _err := client.SyncProductInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
