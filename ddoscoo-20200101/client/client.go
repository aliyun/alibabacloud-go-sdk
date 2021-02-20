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

type AddAutoCcBlacklistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Blacklist  *string `json:"Blacklist,omitempty" xml:"Blacklist,omitempty"`
	ExpireTime *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s AddAutoCcBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddAutoCcBlacklistRequest) SetSourceIp(v string) *AddAutoCcBlacklistRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAutoCcBlacklistRequest) SetInstanceId(v string) *AddAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

func (s *AddAutoCcBlacklistRequest) SetBlacklist(v string) *AddAutoCcBlacklistRequest {
	s.Blacklist = &v
	return s
}

func (s *AddAutoCcBlacklistRequest) SetExpireTime(v int32) *AddAutoCcBlacklistRequest {
	s.ExpireTime = &v
	return s
}

type AddAutoCcBlacklistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAutoCcBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddAutoCcBlacklistResponseBody) SetRequestId(v string) *AddAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type AddAutoCcBlacklistResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAutoCcBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *AddAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddAutoCcBlacklistResponse) SetBody(v *AddAutoCcBlacklistResponseBody) *AddAutoCcBlacklistResponse {
	s.Body = v
	return s
}

type AddAutoCcWhitelistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Whitelist  *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	ExpireTime *int32  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s AddAutoCcWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *AddAutoCcWhitelistRequest) SetSourceIp(v string) *AddAutoCcWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAutoCcWhitelistRequest) SetInstanceId(v string) *AddAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *AddAutoCcWhitelistRequest) SetWhitelist(v string) *AddAutoCcWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *AddAutoCcWhitelistRequest) SetExpireTime(v int32) *AddAutoCcWhitelistRequest {
	s.ExpireTime = &v
	return s
}

type AddAutoCcWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAutoCcWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *AddAutoCcWhitelistResponseBody) SetRequestId(v string) *AddAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type AddAutoCcWhitelistResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAutoCcWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *AddAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *AddAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *AddAutoCcWhitelistResponse) SetBody(v *AddAutoCcWhitelistResponseBody) *AddAutoCcWhitelistResponse {
	s.Body = v
	return s
}

type AssociateWebCertRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CertId          *int32  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cert            *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s AssociateWebCertRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateWebCertRequest) GoString() string {
	return s.String()
}

func (s *AssociateWebCertRequest) SetSourceIp(v string) *AssociateWebCertRequest {
	s.SourceIp = &v
	return s
}

func (s *AssociateWebCertRequest) SetResourceGroupId(v string) *AssociateWebCertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AssociateWebCertRequest) SetDomain(v string) *AssociateWebCertRequest {
	s.Domain = &v
	return s
}

func (s *AssociateWebCertRequest) SetCertId(v int32) *AssociateWebCertRequest {
	s.CertId = &v
	return s
}

func (s *AssociateWebCertRequest) SetCertName(v string) *AssociateWebCertRequest {
	s.CertName = &v
	return s
}

func (s *AssociateWebCertRequest) SetCert(v string) *AssociateWebCertRequest {
	s.Cert = &v
	return s
}

func (s *AssociateWebCertRequest) SetKey(v string) *AssociateWebCertRequest {
	s.Key = &v
	return s
}

type AssociateWebCertResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateWebCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateWebCertResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateWebCertResponseBody) SetRequestId(v string) *AssociateWebCertResponseBody {
	s.RequestId = &v
	return s
}

type AssociateWebCertResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateWebCertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateWebCertResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateWebCertResponse) GoString() string {
	return s.String()
}

func (s *AssociateWebCertResponse) SetHeaders(v map[string]*string) *AssociateWebCertResponse {
	s.Headers = v
	return s
}

func (s *AssociateWebCertResponse) SetBody(v *AssociateWebCertResponseBody) *AssociateWebCertResponse {
	s.Body = v
	return s
}

type AttachSceneDefenseObjectRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Objects    *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
}

func (s AttachSceneDefenseObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachSceneDefenseObjectRequest) GoString() string {
	return s.String()
}

func (s *AttachSceneDefenseObjectRequest) SetSourceIp(v string) *AttachSceneDefenseObjectRequest {
	s.SourceIp = &v
	return s
}

func (s *AttachSceneDefenseObjectRequest) SetPolicyId(v string) *AttachSceneDefenseObjectRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachSceneDefenseObjectRequest) SetObjectType(v string) *AttachSceneDefenseObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *AttachSceneDefenseObjectRequest) SetObjects(v string) *AttachSceneDefenseObjectRequest {
	s.Objects = &v
	return s
}

type AttachSceneDefenseObjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachSceneDefenseObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachSceneDefenseObjectResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSceneDefenseObjectResponseBody) SetRequestId(v string) *AttachSceneDefenseObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachSceneDefenseObjectResponseBody) SetSuccess(v bool) *AttachSceneDefenseObjectResponseBody {
	s.Success = &v
	return s
}

type AttachSceneDefenseObjectResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachSceneDefenseObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachSceneDefenseObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachSceneDefenseObjectResponse) GoString() string {
	return s.String()
}

func (s *AttachSceneDefenseObjectResponse) SetHeaders(v map[string]*string) *AttachSceneDefenseObjectResponse {
	s.Headers = v
	return s
}

func (s *AttachSceneDefenseObjectResponse) SetBody(v *AttachSceneDefenseObjectResponseBody) *AttachSceneDefenseObjectResponse {
	s.Body = v
	return s
}

type ConfigL7RsPolicyRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Policy          *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ConfigL7RsPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigL7RsPolicyRequest) GoString() string {
	return s.String()
}

func (s *ConfigL7RsPolicyRequest) SetSourceIp(v string) *ConfigL7RsPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) SetResourceGroupId(v string) *ConfigL7RsPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) SetDomain(v string) *ConfigL7RsPolicyRequest {
	s.Domain = &v
	return s
}

func (s *ConfigL7RsPolicyRequest) SetPolicy(v string) *ConfigL7RsPolicyRequest {
	s.Policy = &v
	return s
}

type ConfigL7RsPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigL7RsPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigL7RsPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigL7RsPolicyResponseBody) SetRequestId(v string) *ConfigL7RsPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ConfigL7RsPolicyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigL7RsPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigL7RsPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigL7RsPolicyResponse) GoString() string {
	return s.String()
}

func (s *ConfigL7RsPolicyResponse) SetHeaders(v map[string]*string) *ConfigL7RsPolicyResponse {
	s.Headers = v
	return s
}

func (s *ConfigL7RsPolicyResponse) SetBody(v *ConfigL7RsPolicyResponseBody) *ConfigL7RsPolicyResponse {
	s.Body = v
	return s
}

type ConfigNetworkRegionBlockRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config     *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ConfigNetworkRegionBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetworkRegionBlockRequest) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRegionBlockRequest) SetSourceIp(v string) *ConfigNetworkRegionBlockRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigNetworkRegionBlockRequest) SetInstanceId(v string) *ConfigNetworkRegionBlockRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigNetworkRegionBlockRequest) SetConfig(v string) *ConfigNetworkRegionBlockRequest {
	s.Config = &v
	return s
}

type ConfigNetworkRegionBlockResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigNetworkRegionBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetworkRegionBlockResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRegionBlockResponseBody) SetRequestId(v string) *ConfigNetworkRegionBlockResponseBody {
	s.RequestId = &v
	return s
}

type ConfigNetworkRegionBlockResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigNetworkRegionBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigNetworkRegionBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetworkRegionBlockResponse) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRegionBlockResponse) SetHeaders(v map[string]*string) *ConfigNetworkRegionBlockResponse {
	s.Headers = v
	return s
}

func (s *ConfigNetworkRegionBlockResponse) SetBody(v *ConfigNetworkRegionBlockResponseBody) *ConfigNetworkRegionBlockResponse {
	s.Body = v
	return s
}

type ConfigNetworkRulesRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s ConfigNetworkRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRulesRequest) SetSourceIp(v string) *ConfigNetworkRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigNetworkRulesRequest) SetNetworkRules(v string) *ConfigNetworkRulesRequest {
	s.NetworkRules = &v
	return s
}

type ConfigNetworkRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigNetworkRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRulesResponseBody) SetRequestId(v string) *ConfigNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

type ConfigNetworkRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigNetworkRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRulesResponse) SetHeaders(v map[string]*string) *ConfigNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *ConfigNetworkRulesResponse) SetBody(v *ConfigNetworkRulesResponseBody) *ConfigNetworkRulesResponse {
	s.Body = v
	return s
}

type ConfigWebCCTemplateRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Template        *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ConfigWebCCTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigWebCCTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigWebCCTemplateRequest) SetSourceIp(v string) *ConfigWebCCTemplateRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigWebCCTemplateRequest) SetResourceGroupId(v string) *ConfigWebCCTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigWebCCTemplateRequest) SetDomain(v string) *ConfigWebCCTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ConfigWebCCTemplateRequest) SetTemplate(v string) *ConfigWebCCTemplateRequest {
	s.Template = &v
	return s
}

type ConfigWebCCTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigWebCCTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigWebCCTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigWebCCTemplateResponseBody) SetRequestId(v string) *ConfigWebCCTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ConfigWebCCTemplateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigWebCCTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigWebCCTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigWebCCTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigWebCCTemplateResponse) SetHeaders(v map[string]*string) *ConfigWebCCTemplateResponse {
	s.Headers = v
	return s
}

func (s *ConfigWebCCTemplateResponse) SetBody(v *ConfigWebCCTemplateResponseBody) *ConfigWebCCTemplateResponse {
	s.Body = v
	return s
}

type ConfigWebIpSetRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	BlackList       []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	WhiteList       []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigWebIpSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigWebIpSetRequest) GoString() string {
	return s.String()
}

func (s *ConfigWebIpSetRequest) SetSourceIp(v string) *ConfigWebIpSetRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigWebIpSetRequest) SetResourceGroupId(v string) *ConfigWebIpSetRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigWebIpSetRequest) SetDomain(v string) *ConfigWebIpSetRequest {
	s.Domain = &v
	return s
}

func (s *ConfigWebIpSetRequest) SetBlackList(v []*string) *ConfigWebIpSetRequest {
	s.BlackList = v
	return s
}

func (s *ConfigWebIpSetRequest) SetWhiteList(v []*string) *ConfigWebIpSetRequest {
	s.WhiteList = v
	return s
}

type ConfigWebIpSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigWebIpSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigWebIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigWebIpSetResponseBody) SetRequestId(v string) *ConfigWebIpSetResponseBody {
	s.RequestId = &v
	return s
}

type ConfigWebIpSetResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigWebIpSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigWebIpSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigWebIpSetResponse) GoString() string {
	return s.String()
}

func (s *ConfigWebIpSetResponse) SetHeaders(v map[string]*string) *ConfigWebIpSetResponse {
	s.Headers = v
	return s
}

func (s *ConfigWebIpSetResponse) SetBody(v *ConfigWebIpSetResponseBody) *ConfigWebIpSetResponse {
	s.Body = v
	return s
}

type CreateAsyncTaskRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskType        *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskParams      *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
}

func (s CreateAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) SetSourceIp(v string) *CreateAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetLang(v string) *CreateAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetResourceGroupId(v string) *CreateAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskType(v int32) *CreateAsyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskParams(v string) *CreateAsyncTaskRequest {
	s.TaskParams = &v
	return s
}

type CreateAsyncTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponseBody) SetRequestId(v string) *CreateAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateAsyncTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponse) SetHeaders(v map[string]*string) *CreateAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncTaskResponse) SetBody(v *CreateAsyncTaskResponseBody) *CreateAsyncTaskResponse {
	s.Body = v
	return s
}

type CreateNetworkRulesRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s CreateNetworkRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRulesRequest) SetSourceIp(v string) *CreateNetworkRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateNetworkRulesRequest) SetNetworkRules(v string) *CreateNetworkRulesRequest {
	s.NetworkRules = &v
	return s
}

type CreateNetworkRulesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkRulesResponseBody) SetRequestId(v string) *CreateNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

type CreateNetworkRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNetworkRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkRulesResponse) SetHeaders(v map[string]*string) *CreateNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkRulesResponse) SetBody(v *CreateNetworkRulesResponseBody) *CreateNetworkRulesResponse {
	s.Body = v
	return s
}

type CreateSceneDefensePolicyRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Template  *string `json:"Template,omitempty" xml:"Template,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateSceneDefensePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneDefensePolicyRequest) SetSourceIp(v string) *CreateSceneDefensePolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetName(v string) *CreateSceneDefensePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetTemplate(v string) *CreateSceneDefensePolicyRequest {
	s.Template = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetStartTime(v int64) *CreateSceneDefensePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetEndTime(v int64) *CreateSceneDefensePolicyRequest {
	s.EndTime = &v
	return s
}

type CreateSceneDefensePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSceneDefensePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneDefensePolicyResponseBody) SetRequestId(v string) *CreateSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneDefensePolicyResponseBody) SetSuccess(v bool) *CreateSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

type CreateSceneDefensePolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSceneDefensePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *CreateSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneDefensePolicyResponse) SetBody(v *CreateSceneDefensePolicyResponseBody) *CreateSceneDefensePolicyResponse {
	s.Body = v
	return s
}

type CreateSchedulerRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Rules           *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType        *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Param           *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s CreateSchedulerRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedulerRuleRequest) SetSourceIp(v string) *CreateSchedulerRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetResourceGroupId(v string) *CreateSchedulerRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetRules(v string) *CreateSchedulerRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetRuleName(v string) *CreateSchedulerRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetRuleType(v int32) *CreateSchedulerRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateSchedulerRuleRequest) SetParam(v string) *CreateSchedulerRuleRequest {
	s.Param = &v
	return s
}

type CreateSchedulerRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Cname     *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	RuleName  *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateSchedulerRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedulerRuleResponseBody) SetRequestId(v string) *CreateSchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchedulerRuleResponseBody) SetCname(v string) *CreateSchedulerRuleResponseBody {
	s.Cname = &v
	return s
}

func (s *CreateSchedulerRuleResponseBody) SetRuleName(v string) *CreateSchedulerRuleResponseBody {
	s.RuleName = &v
	return s
}

type CreateSchedulerRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSchedulerRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedulerRuleResponse) SetHeaders(v map[string]*string) *CreateSchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedulerRuleResponse) SetBody(v *CreateSchedulerRuleResponseBody) *CreateSchedulerRuleResponse {
	s.Body = v
	return s
}

type CreateTagResourcesRequest struct {
	SourceIp        *string                          `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceIds     []*string                        `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	Tags            []*CreateTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesRequest) SetSourceIp(v string) *CreateTagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateTagResourcesRequest) SetRegionId(v string) *CreateTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTagResourcesRequest) SetResourceGroupId(v string) *CreateTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTagResourcesRequest) SetResourceType(v string) *CreateTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateTagResourcesRequest) SetResourceIds(v []*string) *CreateTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *CreateTagResourcesRequest) SetTags(v []*CreateTagResourcesRequestTags) *CreateTagResourcesRequest {
	s.Tags = v
	return s
}

type CreateTagResourcesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesRequestTags) SetKey(v string) *CreateTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTagResourcesRequestTags) SetValue(v string) *CreateTagResourcesRequestTags {
	s.Value = &v
	return s
}

type CreateTagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesResponseBody) SetRequestId(v string) *CreateTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type CreateTagResourcesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *CreateTagResourcesResponse) SetHeaders(v map[string]*string) *CreateTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *CreateTagResourcesResponse) SetBody(v *CreateTagResourcesResponseBody) *CreateTagResourcesResponse {
	s.Body = v
	return s
}

type CreateWebCCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Count           *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateWebCCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWebCCRuleRequest) SetSourceIp(v string) *CreateWebCCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetResourceGroupId(v string) *CreateWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetDomain(v string) *CreateWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetName(v string) *CreateWebCCRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetAct(v string) *CreateWebCCRuleRequest {
	s.Act = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetCount(v int32) *CreateWebCCRuleRequest {
	s.Count = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetInterval(v int32) *CreateWebCCRuleRequest {
	s.Interval = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetMode(v string) *CreateWebCCRuleRequest {
	s.Mode = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetTtl(v int32) *CreateWebCCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetUri(v string) *CreateWebCCRuleRequest {
	s.Uri = &v
	return s
}

type CreateWebCCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebCCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebCCRuleResponseBody) SetRequestId(v string) *CreateWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateWebCCRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWebCCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWebCCRuleResponse) SetHeaders(v map[string]*string) *CreateWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWebCCRuleResponse) SetBody(v *CreateWebCCRuleResponseBody) *CreateWebCCRuleResponse {
	s.Body = v
	return s
}

type CreateWebRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RsType          *int32    `json:"RsType,omitempty" xml:"RsType,omitempty"`
	Rules           *string   `json:"Rules,omitempty" xml:"Rules,omitempty"`
	HttpsExt        *string   `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	DefenseId       *string   `json:"DefenseId,omitempty" xml:"DefenseId,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateWebRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWebRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWebRuleRequest) SetSourceIp(v string) *CreateWebRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateWebRuleRequest) SetResourceGroupId(v string) *CreateWebRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWebRuleRequest) SetDomain(v string) *CreateWebRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateWebRuleRequest) SetRsType(v int32) *CreateWebRuleRequest {
	s.RsType = &v
	return s
}

func (s *CreateWebRuleRequest) SetRules(v string) *CreateWebRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateWebRuleRequest) SetHttpsExt(v string) *CreateWebRuleRequest {
	s.HttpsExt = &v
	return s
}

func (s *CreateWebRuleRequest) SetDefenseId(v string) *CreateWebRuleRequest {
	s.DefenseId = &v
	return s
}

func (s *CreateWebRuleRequest) SetInstanceIds(v []*string) *CreateWebRuleRequest {
	s.InstanceIds = v
	return s
}

type CreateWebRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWebRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebRuleResponseBody) SetRequestId(v string) *CreateWebRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateWebRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWebRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWebRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWebRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWebRuleResponse) SetHeaders(v map[string]*string) *CreateWebRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWebRuleResponse) SetBody(v *CreateWebRuleResponseBody) *CreateWebRuleResponse {
	s.Body = v
	return s
}

type DeleteAsyncTaskRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskId          *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskRequest) SetSourceIp(v string) *DeleteAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetLang(v string) *DeleteAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetResourceGroupId(v string) *DeleteAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetTaskId(v int32) *DeleteAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteAsyncTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponseBody) SetRequestId(v string) *DeleteAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAsyncTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponse) SetHeaders(v map[string]*string) *DeleteAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsyncTaskResponse) SetBody(v *DeleteAsyncTaskResponseBody) *DeleteAsyncTaskResponse {
	s.Body = v
	return s
}

type DeleteAutoCcBlacklistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Blacklist  *string `json:"Blacklist,omitempty" xml:"Blacklist,omitempty"`
}

func (s DeleteAutoCcBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcBlacklistRequest) SetSourceIp(v string) *DeleteAutoCcBlacklistRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAutoCcBlacklistRequest) SetInstanceId(v string) *DeleteAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoCcBlacklistRequest) SetBlacklist(v string) *DeleteAutoCcBlacklistRequest {
	s.Blacklist = &v
	return s
}

type DeleteAutoCcBlacklistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoCcBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcBlacklistResponseBody) SetRequestId(v string) *DeleteAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAutoCcBlacklistResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutoCcBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *DeleteAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoCcBlacklistResponse) SetBody(v *DeleteAutoCcBlacklistResponseBody) *DeleteAutoCcBlacklistResponse {
	s.Body = v
	return s
}

type DeleteAutoCcWhitelistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Whitelist  *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s DeleteAutoCcWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcWhitelistRequest) SetSourceIp(v string) *DeleteAutoCcWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAutoCcWhitelistRequest) SetInstanceId(v string) *DeleteAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoCcWhitelistRequest) SetWhitelist(v string) *DeleteAutoCcWhitelistRequest {
	s.Whitelist = &v
	return s
}

type DeleteAutoCcWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoCcWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcWhitelistResponseBody) SetRequestId(v string) *DeleteAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAutoCcWhitelistResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutoCcWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *DeleteAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoCcWhitelistResponse) SetBody(v *DeleteAutoCcWhitelistResponseBody) *DeleteAutoCcWhitelistResponse {
	s.Body = v
	return s
}

type DeleteNetworkRuleRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	NetworkRule *string `json:"NetworkRule,omitempty" xml:"NetworkRule,omitempty"`
}

func (s DeleteNetworkRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleRequest) SetSourceIp(v string) *DeleteNetworkRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteNetworkRuleRequest) SetNetworkRule(v string) *DeleteNetworkRuleRequest {
	s.NetworkRule = &v
	return s
}

type DeleteNetworkRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleResponseBody) SetRequestId(v string) *DeleteNetworkRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNetworkRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleResponse) SetHeaders(v map[string]*string) *DeleteNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkRuleResponse) SetBody(v *DeleteNetworkRuleResponseBody) *DeleteNetworkRuleResponse {
	s.Body = v
	return s
}

type DeleteSceneDefensePolicyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteSceneDefensePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSceneDefensePolicyRequest) SetSourceIp(v string) *DeleteSceneDefensePolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteSceneDefensePolicyRequest) SetPolicyId(v string) *DeleteSceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

type DeleteSceneDefensePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSceneDefensePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneDefensePolicyResponseBody) SetRequestId(v string) *DeleteSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSceneDefensePolicyResponseBody) SetSuccess(v bool) *DeleteSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

type DeleteSceneDefensePolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSceneDefensePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *DeleteSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneDefensePolicyResponse) SetBody(v *DeleteSceneDefensePolicyResponseBody) *DeleteSceneDefensePolicyResponse {
	s.Body = v
	return s
}

type DeleteSchedulerRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteSchedulerRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerRuleRequest) SetSourceIp(v string) *DeleteSchedulerRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteSchedulerRuleRequest) SetResourceGroupId(v string) *DeleteSchedulerRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteSchedulerRuleRequest) SetRuleName(v string) *DeleteSchedulerRuleRequest {
	s.RuleName = &v
	return s
}

type DeleteSchedulerRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchedulerRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerRuleResponseBody) SetRequestId(v string) *DeleteSchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSchedulerRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSchedulerRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerRuleResponse) SetHeaders(v map[string]*string) *DeleteSchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedulerRuleResponse) SetBody(v *DeleteSchedulerRuleResponseBody) *DeleteSchedulerRuleResponse {
	s.Body = v
	return s
}

type DeleteTagResourcesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All             *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceIds     []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	TagKey          []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DeleteTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagResourcesRequest) SetSourceIp(v string) *DeleteTagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetRegionId(v string) *DeleteTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetResourceGroupId(v string) *DeleteTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetResourceType(v string) *DeleteTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetAll(v bool) *DeleteTagResourcesRequest {
	s.All = &v
	return s
}

func (s *DeleteTagResourcesRequest) SetResourceIds(v []*string) *DeleteTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *DeleteTagResourcesRequest) SetTagKey(v []*string) *DeleteTagResourcesRequest {
	s.TagKey = v
	return s
}

type DeleteTagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResourcesResponseBody) SetRequestId(v string) *DeleteTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagResourcesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagResourcesResponse) SetHeaders(v map[string]*string) *DeleteTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagResourcesResponse) SetBody(v *DeleteTagResourcesResponseBody) *DeleteTagResourcesResponse {
	s.Body = v
	return s
}

type DeleteWebCacheCustomRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RuleNames       []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DeleteWebCacheCustomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebCacheCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebCacheCustomRuleRequest) SetSourceIp(v string) *DeleteWebCacheCustomRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteWebCacheCustomRuleRequest) SetResourceGroupId(v string) *DeleteWebCacheCustomRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebCacheCustomRuleRequest) SetDomain(v string) *DeleteWebCacheCustomRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebCacheCustomRuleRequest) SetRuleNames(v []*string) *DeleteWebCacheCustomRuleRequest {
	s.RuleNames = v
	return s
}

type DeleteWebCacheCustomRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebCacheCustomRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebCacheCustomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebCacheCustomRuleResponseBody) SetRequestId(v string) *DeleteWebCacheCustomRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebCacheCustomRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebCacheCustomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebCacheCustomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebCacheCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebCacheCustomRuleResponse) SetHeaders(v map[string]*string) *DeleteWebCacheCustomRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebCacheCustomRuleResponse) SetBody(v *DeleteWebCacheCustomRuleResponseBody) *DeleteWebCacheCustomRuleResponse {
	s.Body = v
	return s
}

type DeleteWebCCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteWebCCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleRequest) SetSourceIp(v string) *DeleteWebCCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteWebCCRuleRequest) SetResourceGroupId(v string) *DeleteWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebCCRuleRequest) SetDomain(v string) *DeleteWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebCCRuleRequest) SetName(v string) *DeleteWebCCRuleRequest {
	s.Name = &v
	return s
}

type DeleteWebCCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebCCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleResponseBody) SetRequestId(v string) *DeleteWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebCCRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebCCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleResponse) SetHeaders(v map[string]*string) *DeleteWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebCCRuleResponse) SetBody(v *DeleteWebCCRuleResponseBody) *DeleteWebCCRuleResponse {
	s.Body = v
	return s
}

type DeleteWebPreciseAccessRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RuleNames       []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DeleteWebPreciseAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebPreciseAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebPreciseAccessRuleRequest) SetSourceIp(v string) *DeleteWebPreciseAccessRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleRequest) SetResourceGroupId(v string) *DeleteWebPreciseAccessRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleRequest) SetDomain(v string) *DeleteWebPreciseAccessRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleRequest) SetRuleNames(v []*string) *DeleteWebPreciseAccessRuleRequest {
	s.RuleNames = v
	return s
}

type DeleteWebPreciseAccessRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebPreciseAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebPreciseAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebPreciseAccessRuleResponseBody) SetRequestId(v string) *DeleteWebPreciseAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebPreciseAccessRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebPreciseAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebPreciseAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebPreciseAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebPreciseAccessRuleResponse) SetHeaders(v map[string]*string) *DeleteWebPreciseAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebPreciseAccessRuleResponse) SetBody(v *DeleteWebPreciseAccessRuleResponseBody) *DeleteWebPreciseAccessRuleResponse {
	s.Body = v
	return s
}

type DeleteWebRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteWebRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebRuleRequest) SetSourceIp(v string) *DeleteWebRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteWebRuleRequest) SetResourceGroupId(v string) *DeleteWebRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebRuleRequest) SetDomain(v string) *DeleteWebRuleRequest {
	s.Domain = &v
	return s
}

type DeleteWebRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebRuleResponseBody) SetRequestId(v string) *DeleteWebRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWebRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebRuleResponse) SetHeaders(v map[string]*string) *DeleteWebRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebRuleResponse) SetBody(v *DeleteWebRuleResponseBody) *DeleteWebRuleResponse {
	s.Body = v
	return s
}

type DescribeAsyncTasksRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAsyncTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksRequest) SetSourceIp(v string) *DescribeAsyncTasksRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAsyncTasksRequest) SetLang(v string) *DescribeAsyncTasksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAsyncTasksRequest) SetResourceGroupId(v string) *DescribeAsyncTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAsyncTasksRequest) SetPageNumber(v int32) *DescribeAsyncTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAsyncTasksRequest) SetPageSize(v int32) *DescribeAsyncTasksRequest {
	s.PageSize = &v
	return s
}

type DescribeAsyncTasksResponseBody struct {
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AsyncTasks []*DescribeAsyncTasksResponseBodyAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" type:"Repeated"`
}

func (s DescribeAsyncTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksResponseBody) SetTotalCount(v int32) *DescribeAsyncTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAsyncTasksResponseBody) SetRequestId(v string) *DescribeAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAsyncTasksResponseBody) SetAsyncTasks(v []*DescribeAsyncTasksResponseBodyAsyncTasks) *DescribeAsyncTasksResponseBody {
	s.AsyncTasks = v
	return s
}

type DescribeAsyncTasksResponseBodyAsyncTasks struct {
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskType   *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	TaskStatus *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeAsyncTasksResponseBodyAsyncTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTasksResponseBodyAsyncTasks) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetEndTime(v int64) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.EndTime = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskType(v int32) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetStartTime(v int64) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.StartTime = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskParams(v string) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskParams = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskStatus(v int32) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskResult(v string) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskResult = &v
	return s
}

func (s *DescribeAsyncTasksResponseBodyAsyncTasks) SetTaskId(v int64) *DescribeAsyncTasksResponseBodyAsyncTasks {
	s.TaskId = &v
	return s
}

type DescribeAsyncTasksResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAsyncTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAsyncTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAsyncTasksResponse) SetHeaders(v map[string]*string) *DescribeAsyncTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAsyncTasksResponse) SetBody(v *DescribeAsyncTasksResponseBody) *DescribeAsyncTasksResponse {
	s.Body = v
	return s
}

type DescribeAutoCcBlacklistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeyWord    *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoCcBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistRequest) SetSourceIp(v string) *DescribeAutoCcBlacklistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetInstanceId(v string) *DescribeAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetKeyWord(v string) *DescribeAutoCcBlacklistRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetPageNumber(v int32) *DescribeAutoCcBlacklistRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetPageSize(v int32) *DescribeAutoCcBlacklistRequest {
	s.PageSize = &v
	return s
}

type DescribeAutoCcBlacklistResponseBody struct {
	TotalCount      *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AutoCcBlacklist []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist `json:"AutoCcBlacklist,omitempty" xml:"AutoCcBlacklist,omitempty" type:"Repeated"`
}

func (s DescribeAutoCcBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistResponseBody) SetTotalCount(v int64) *DescribeAutoCcBlacklistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBody) SetRequestId(v string) *DescribeAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBody) SetAutoCcBlacklist(v []*DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) *DescribeAutoCcBlacklistResponseBody {
	s.AutoCcBlacklist = v
	return s
}

type DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DestIp   *string `json:"DestIp,omitempty" xml:"DestIp,omitempty"`
}

func (s DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetType(v string) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.Type = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetEndTime(v int64) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.EndTime = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetSourceIp(v string) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist) SetDestIp(v string) *DescribeAutoCcBlacklistResponseBodyAutoCcBlacklist {
	s.DestIp = &v
	return s
}

type DescribeAutoCcBlacklistResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoCcBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *DescribeAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoCcBlacklistResponse) SetBody(v *DescribeAutoCcBlacklistResponseBody) *DescribeAutoCcBlacklistResponse {
	s.Body = v
	return s
}

type DescribeAutoCcListCountRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
}

func (s DescribeAutoCcListCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcListCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcListCountRequest) SetSourceIp(v string) *DescribeAutoCcListCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcListCountRequest) SetInstanceId(v string) *DescribeAutoCcListCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoCcListCountRequest) SetQueryType(v string) *DescribeAutoCcListCountRequest {
	s.QueryType = &v
	return s
}

type DescribeAutoCcListCountResponseBody struct {
	BlackCount *int32  `json:"BlackCount,omitempty" xml:"BlackCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WhiteCount *int32  `json:"WhiteCount,omitempty" xml:"WhiteCount,omitempty"`
}

func (s DescribeAutoCcListCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcListCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcListCountResponseBody) SetBlackCount(v int32) *DescribeAutoCcListCountResponseBody {
	s.BlackCount = &v
	return s
}

func (s *DescribeAutoCcListCountResponseBody) SetRequestId(v string) *DescribeAutoCcListCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoCcListCountResponseBody) SetWhiteCount(v int32) *DescribeAutoCcListCountResponseBody {
	s.WhiteCount = &v
	return s
}

type DescribeAutoCcListCountResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoCcListCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoCcListCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcListCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcListCountResponse) SetHeaders(v map[string]*string) *DescribeAutoCcListCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoCcListCountResponse) SetBody(v *DescribeAutoCcListCountResponseBody) *DescribeAutoCcListCountResponse {
	s.Body = v
	return s
}

type DescribeAutoCcWhitelistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeyWord    *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoCcWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistRequest) SetSourceIp(v string) *DescribeAutoCcWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetInstanceId(v string) *DescribeAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetKeyWord(v string) *DescribeAutoCcWhitelistRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetPageNumber(v int32) *DescribeAutoCcWhitelistRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetPageSize(v int32) *DescribeAutoCcWhitelistRequest {
	s.PageSize = &v
	return s
}

type DescribeAutoCcWhitelistResponseBody struct {
	AutoCcWhitelist []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist `json:"AutoCcWhitelist,omitempty" xml:"AutoCcWhitelist,omitempty" type:"Repeated"`
	TotalCount      *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoCcWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistResponseBody) SetAutoCcWhitelist(v []*DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) *DescribeAutoCcWhitelistResponseBody {
	s.AutoCcWhitelist = v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBody) SetTotalCount(v int64) *DescribeAutoCcWhitelistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBody) SetRequestId(v string) *DescribeAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DestIp   *string `json:"DestIp,omitempty" xml:"DestIp,omitempty"`
}

func (s DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetType(v string) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.Type = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetEndTime(v int64) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.EndTime = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetSourceIp(v string) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist) SetDestIp(v string) *DescribeAutoCcWhitelistResponseBodyAutoCcWhitelist {
	s.DestIp = &v
	return s
}

type DescribeAutoCcWhitelistResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoCcWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *DescribeAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoCcWhitelistResponse) SetBody(v *DescribeAutoCcWhitelistResponseBody) *DescribeAutoCcWhitelistResponse {
	s.Body = v
	return s
}

type DescribeBackSourceCidrRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Line            *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s DescribeBackSourceCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrRequest) SetSourceIp(v string) *DescribeBackSourceCidrRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetResourceGroupId(v string) *DescribeBackSourceCidrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetLine(v string) *DescribeBackSourceCidrRequest {
	s.Line = &v
	return s
}

type DescribeBackSourceCidrResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Cidrs     []*string `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Repeated"`
}

func (s DescribeBackSourceCidrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponseBody) SetRequestId(v string) *DescribeBackSourceCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) SetCidrs(v []*string) *DescribeBackSourceCidrResponseBody {
	s.Cidrs = v
	return s
}

type DescribeBackSourceCidrResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackSourceCidrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackSourceCidrResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponse) SetHeaders(v map[string]*string) *DescribeBackSourceCidrResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackSourceCidrResponse) SetBody(v *DescribeBackSourceCidrResponseBody) *DescribeBackSourceCidrResponse {
	s.Body = v
	return s
}

type DescribeBlackholeStatusRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeBlackholeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackholeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusRequest) SetSourceIp(v string) *DescribeBlackholeStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBlackholeStatusRequest) SetInstanceIds(v []*string) *DescribeBlackholeStatusRequest {
	s.InstanceIds = v
	return s
}

type DescribeBlackholeStatusResponseBody struct {
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BlackholeStatus []*DescribeBlackholeStatusResponseBodyBlackholeStatus `json:"BlackholeStatus,omitempty" xml:"BlackholeStatus,omitempty" type:"Repeated"`
}

func (s DescribeBlackholeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackholeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusResponseBody) SetRequestId(v string) *DescribeBlackholeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBody) SetBlackholeStatus(v []*DescribeBlackholeStatusResponseBodyBlackholeStatus) *DescribeBlackholeStatusResponseBody {
	s.BlackholeStatus = v
	return s
}

type DescribeBlackholeStatusResponseBodyBlackholeStatus struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	BlackStatus *string `json:"BlackStatus,omitempty" xml:"BlackStatus,omitempty"`
}

func (s DescribeBlackholeStatusResponseBodyBlackholeStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackholeStatusResponseBodyBlackholeStatus) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetEndTime(v int64) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.EndTime = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetStartTime(v int64) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.StartTime = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetIp(v string) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.Ip = &v
	return s
}

func (s *DescribeBlackholeStatusResponseBodyBlackholeStatus) SetBlackStatus(v string) *DescribeBlackholeStatusResponseBodyBlackholeStatus {
	s.BlackStatus = &v
	return s
}

type DescribeBlackholeStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBlackholeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBlackholeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackholeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusResponse) SetHeaders(v map[string]*string) *DescribeBlackholeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlackholeStatusResponse) SetBody(v *DescribeBlackholeStatusResponseBody) *DescribeBlackholeStatusResponse {
	s.Body = v
	return s
}

type DescribeBlockStatusRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeBlockStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusRequest) SetSourceIp(v string) *DescribeBlockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBlockStatusRequest) SetLang(v string) *DescribeBlockStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBlockStatusRequest) SetResourceGroupId(v string) *DescribeBlockStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBlockStatusRequest) SetInstanceIds(v []*string) *DescribeBlockStatusRequest {
	s.InstanceIds = v
	return s
}

type DescribeBlockStatusResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusList []*DescribeBlockStatusResponseBodyStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeBlockStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponseBody) SetRequestId(v string) *DescribeBlockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlockStatusResponseBody) SetStatusList(v []*DescribeBlockStatusResponseBodyStatusList) *DescribeBlockStatusResponseBody {
	s.StatusList = v
	return s
}

type DescribeBlockStatusResponseBodyStatusList struct {
	Ip              *string                                                     `json:"Ip,omitempty" xml:"Ip,omitempty"`
	BlockStatusList []*DescribeBlockStatusResponseBodyStatusListBlockStatusList `json:"BlockStatusList,omitempty" xml:"BlockStatusList,omitempty" type:"Repeated"`
}

func (s DescribeBlockStatusResponseBodyStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockStatusResponseBodyStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponseBodyStatusList) SetIp(v string) *DescribeBlockStatusResponseBodyStatusList {
	s.Ip = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusList) SetBlockStatusList(v []*DescribeBlockStatusResponseBodyStatusListBlockStatusList) *DescribeBlockStatusResponseBodyStatusList {
	s.BlockStatusList = v
	return s
}

type DescribeBlockStatusResponseBodyStatusListBlockStatusList struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Line        *string `json:"Line,omitempty" xml:"Line,omitempty"`
	BlockStatus *string `json:"BlockStatus,omitempty" xml:"BlockStatus,omitempty"`
}

func (s DescribeBlockStatusResponseBodyStatusListBlockStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockStatusResponseBodyStatusListBlockStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetEndTime(v int64) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.EndTime = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetStartTime(v int64) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.StartTime = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetLine(v string) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.Line = &v
	return s
}

func (s *DescribeBlockStatusResponseBodyStatusListBlockStatusList) SetBlockStatus(v string) *DescribeBlockStatusResponseBodyStatusListBlockStatusList {
	s.BlockStatus = &v
	return s
}

type DescribeBlockStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBlockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBlockStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlockStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponse) SetHeaders(v map[string]*string) *DescribeBlockStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlockStatusResponse) SetBody(v *DescribeBlockStatusResponseBody) *DescribeBlockStatusResponse {
	s.Body = v
	return s
}

type DescribeCertsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertsRequest) SetSourceIp(v string) *DescribeCertsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCertsRequest) SetResourceGroupId(v string) *DescribeCertsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCertsRequest) SetDomain(v string) *DescribeCertsRequest {
	s.Domain = &v
	return s
}

type DescribeCertsResponseBody struct {
	Certs     []*DescribeCertsResponseBodyCerts `json:"Certs,omitempty" xml:"Certs,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DescribeCertsResponseBodyCerts struct {
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DomainRelated *bool   `json:"DomainRelated,omitempty" xml:"DomainRelated,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Issuer        *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Common        *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Id            *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCertsResponseBodyCerts) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertsResponseBodyCerts) GoString() string {
	return s.String()
}

func (s *DescribeCertsResponseBodyCerts) SetEndDate(v string) *DescribeCertsResponseBodyCerts {
	s.EndDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetDomainRelated(v bool) *DescribeCertsResponseBodyCerts {
	s.DomainRelated = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetStartDate(v string) *DescribeCertsResponseBodyCerts {
	s.StartDate = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetIssuer(v string) *DescribeCertsResponseBodyCerts {
	s.Issuer = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetName(v string) *DescribeCertsResponseBodyCerts {
	s.Name = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetCommon(v string) *DescribeCertsResponseBodyCerts {
	s.Common = &v
	return s
}

func (s *DescribeCertsResponseBodyCerts) SetId(v int32) *DescribeCertsResponseBodyCerts {
	s.Id = &v
	return s
}

type DescribeCertsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCertsResponse) SetBody(v *DescribeCertsResponseBody) *DescribeCertsResponse {
	s.Body = v
	return s
}

type DescribeCnameReusesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domains         []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeCnameReusesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameReusesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesRequest) SetSourceIp(v string) *DescribeCnameReusesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCnameReusesRequest) SetResourceGroupId(v string) *DescribeCnameReusesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCnameReusesRequest) SetDomains(v []*string) *DescribeCnameReusesRequest {
	s.Domains = v
	return s
}

type DescribeCnameReusesResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CnameReuses []*DescribeCnameReusesResponseBodyCnameReuses `json:"CnameReuses,omitempty" xml:"CnameReuses,omitempty" type:"Repeated"`
}

func (s DescribeCnameReusesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameReusesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesResponseBody) SetRequestId(v string) *DescribeCnameReusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCnameReusesResponseBody) SetCnameReuses(v []*DescribeCnameReusesResponseBodyCnameReuses) *DescribeCnameReusesResponseBody {
	s.CnameReuses = v
	return s
}

type DescribeCnameReusesResponseBodyCnameReuses struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Cname  *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Enable *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeCnameReusesResponseBodyCnameReuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameReusesResponseBodyCnameReuses) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) SetDomain(v string) *DescribeCnameReusesResponseBodyCnameReuses {
	s.Domain = &v
	return s
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) SetCname(v string) *DescribeCnameReusesResponseBodyCnameReuses {
	s.Cname = &v
	return s
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) SetEnable(v int32) *DescribeCnameReusesResponseBodyCnameReuses {
	s.Enable = &v
	return s
}

type DescribeCnameReusesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCnameReusesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCnameReusesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameReusesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesResponse) SetHeaders(v map[string]*string) *DescribeCnameReusesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCnameReusesResponse) SetBody(v *DescribeCnameReusesResponseBody) *DescribeCnameReusesResponse {
	s.Body = v
	return s
}

type DescribeDDosAllEventListRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EventType  *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDDosAllEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosAllEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListRequest) SetSourceIp(v string) *DescribeDDosAllEventListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetEventType(v string) *DescribeDDosAllEventListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetStartTime(v int64) *DescribeDDosAllEventListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetEndTime(v int64) *DescribeDDosAllEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetPageNumber(v int32) *DescribeDDosAllEventListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDosAllEventListRequest) SetPageSize(v int32) *DescribeDDosAllEventListRequest {
	s.PageSize = &v
	return s
}

type DescribeDDosAllEventListResponseBody struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	AttackEvents []*DescribeDDosAllEventListResponseBodyAttackEvents `json:"AttackEvents,omitempty" xml:"AttackEvents,omitempty" type:"Repeated"`
}

func (s DescribeDDosAllEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosAllEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListResponseBody) SetRequestId(v string) *DescribeDDosAllEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBody) SetTotal(v int64) *DescribeDDosAllEventListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBody) SetAttackEvents(v []*DescribeDDosAllEventListResponseBodyAttackEvents) *DescribeDDosAllEventListResponseBody {
	s.AttackEvents = v
	return s
}

type DescribeDDosAllEventListResponseBodyAttackEvents struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Mbps      *int64  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Pps       *int64  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeDDosAllEventListResponseBodyAttackEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosAllEventListResponseBodyAttackEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetEndTime(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetStartTime(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetEventType(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.EventType = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetMbps(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetIp(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetPort(v string) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Port = &v
	return s
}

func (s *DescribeDDosAllEventListResponseBodyAttackEvents) SetPps(v int64) *DescribeDDosAllEventListResponseBodyAttackEvents {
	s.Pps = &v
	return s
}

type DescribeDDosAllEventListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDosAllEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDosAllEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosAllEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListResponse) SetHeaders(v map[string]*string) *DescribeDDosAllEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosAllEventListResponse) SetBody(v *DescribeDDosAllEventListResponseBody) *DescribeDDosAllEventListResponse {
	s.Body = v
	return s
}

type DescribeDDosEventAreaRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDDosEventAreaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAreaRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaRequest) SetSourceIp(v string) *DescribeDDosEventAreaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) SetEventType(v string) *DescribeDDosEventAreaRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) SetStartTime(v int64) *DescribeDDosEventAreaRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventAreaRequest) SetIp(v string) *DescribeDDosEventAreaRequest {
	s.Ip = &v
	return s
}

type DescribeDDosEventAreaResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Areas     []*DescribeDDosEventAreaResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
}

func (s DescribeDDosEventAreaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAreaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaResponseBody) SetRequestId(v string) *DescribeDDosEventAreaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventAreaResponseBody) SetAreas(v []*DescribeDDosEventAreaResponseBodyAreas) *DescribeDDosEventAreaResponseBody {
	s.Areas = v
	return s
}

type DescribeDDosEventAreaResponseBodyAreas struct {
	InPkts *int64  `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
	Area   *string `json:"Area,omitempty" xml:"Area,omitempty"`
}

func (s DescribeDDosEventAreaResponseBodyAreas) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAreaResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaResponseBodyAreas) SetInPkts(v int64) *DescribeDDosEventAreaResponseBodyAreas {
	s.InPkts = &v
	return s
}

func (s *DescribeDDosEventAreaResponseBodyAreas) SetArea(v string) *DescribeDDosEventAreaResponseBodyAreas {
	s.Area = &v
	return s
}

type DescribeDDosEventAreaResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDosEventAreaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDosEventAreaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAreaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaResponse) SetHeaders(v map[string]*string) *DescribeDDosEventAreaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventAreaResponse) SetBody(v *DescribeDDosEventAreaResponseBody) *DescribeDDosEventAreaResponse {
	s.Body = v
	return s
}

type DescribeDDosEventAttackTypeRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDDosEventAttackTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAttackTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeRequest) SetSourceIp(v string) *DescribeDDosEventAttackTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDosEventAttackTypeRequest) SetEventType(v string) *DescribeDDosEventAttackTypeRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventAttackTypeRequest) SetStartTime(v int64) *DescribeDDosEventAttackTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventAttackTypeRequest) SetIp(v string) *DescribeDDosEventAttackTypeRequest {
	s.Ip = &v
	return s
}

type DescribeDDosEventAttackTypeResponseBody struct {
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AttackTypes []*DescribeDDosEventAttackTypeResponseBodyAttackTypes `json:"AttackTypes,omitempty" xml:"AttackTypes,omitempty" type:"Repeated"`
}

func (s DescribeDDosEventAttackTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAttackTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeResponseBody) SetRequestId(v string) *DescribeDDosEventAttackTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventAttackTypeResponseBody) SetAttackTypes(v []*DescribeDDosEventAttackTypeResponseBodyAttackTypes) *DescribeDDosEventAttackTypeResponseBody {
	s.AttackTypes = v
	return s
}

type DescribeDDosEventAttackTypeResponseBodyAttackTypes struct {
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	InPkts     *int64  `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
}

func (s DescribeDDosEventAttackTypeResponseBodyAttackTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAttackTypeResponseBodyAttackTypes) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) SetAttackType(v string) *DescribeDDosEventAttackTypeResponseBodyAttackTypes {
	s.AttackType = &v
	return s
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) SetInPkts(v int64) *DescribeDDosEventAttackTypeResponseBodyAttackTypes {
	s.InPkts = &v
	return s
}

type DescribeDDosEventAttackTypeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDosEventAttackTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDosEventAttackTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventAttackTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeResponse) SetHeaders(v map[string]*string) *DescribeDDosEventAttackTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventAttackTypeResponse) SetBody(v *DescribeDDosEventAttackTypeResponseBody) *DescribeDDosEventAttackTypeResponse {
	s.Body = v
	return s
}

type DescribeDDosEventIspRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDDosEventIspRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventIspRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspRequest) SetSourceIp(v string) *DescribeDDosEventIspRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDosEventIspRequest) SetEventType(v string) *DescribeDDosEventIspRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventIspRequest) SetStartTime(v int64) *DescribeDDosEventIspRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventIspRequest) SetIp(v string) *DescribeDDosEventIspRequest {
	s.Ip = &v
	return s
}

type DescribeDDosEventIspResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Isps      []*DescribeDDosEventIspResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Repeated"`
}

func (s DescribeDDosEventIspResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventIspResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspResponseBody) SetRequestId(v string) *DescribeDDosEventIspResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventIspResponseBody) SetIsps(v []*DescribeDDosEventIspResponseBodyIsps) *DescribeDDosEventIspResponseBody {
	s.Isps = v
	return s
}

type DescribeDDosEventIspResponseBodyIsps struct {
	InPkts *int64  `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
	Isp    *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeDDosEventIspResponseBodyIsps) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventIspResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspResponseBodyIsps) SetInPkts(v int64) *DescribeDDosEventIspResponseBodyIsps {
	s.InPkts = &v
	return s
}

func (s *DescribeDDosEventIspResponseBodyIsps) SetIsp(v string) *DescribeDDosEventIspResponseBodyIsps {
	s.Isp = &v
	return s
}

type DescribeDDosEventIspResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDosEventIspResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDosEventIspResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventIspResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspResponse) SetHeaders(v map[string]*string) *DescribeDDosEventIspResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventIspResponse) SetBody(v *DescribeDDosEventIspResponseBody) *DescribeDDosEventIspResponse {
	s.Body = v
	return s
}

type DescribeDDosEventMaxRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDDosEventMaxRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventMaxRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventMaxRequest) SetSourceIp(v string) *DescribeDDosEventMaxRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDosEventMaxRequest) SetStartTime(v int64) *DescribeDDosEventMaxRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventMaxRequest) SetEndTime(v int64) *DescribeDDosEventMaxRequest {
	s.EndTime = &v
	return s
}

type DescribeDDosEventMaxResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Qps       *int64  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Cps       *int64  `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Mbps      *int64  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
}

func (s DescribeDDosEventMaxResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventMaxResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventMaxResponseBody) SetRequestId(v string) *DescribeDDosEventMaxResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) SetQps(v int64) *DescribeDDosEventMaxResponseBody {
	s.Qps = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) SetCps(v int64) *DescribeDDosEventMaxResponseBody {
	s.Cps = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) SetMbps(v int64) *DescribeDDosEventMaxResponseBody {
	s.Mbps = &v
	return s
}

type DescribeDDosEventMaxResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDosEventMaxResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDosEventMaxResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventMaxResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventMaxResponse) SetHeaders(v map[string]*string) *DescribeDDosEventMaxResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventMaxResponse) SetBody(v *DescribeDDosEventMaxResponseBody) *DescribeDDosEventMaxResponse {
	s.Body = v
	return s
}

type DescribeDDoSEventsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeDDoSEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsRequest) SetSourceIp(v string) *DescribeDDoSEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetResourceGroupId(v string) *DescribeDDoSEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEndTime(v int64) *DescribeDDoSEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetStartTime(v int64) *DescribeDDoSEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageSize(v int32) *DescribeDDoSEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageNumber(v int32) *DescribeDDoSEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetInstanceIds(v []*string) *DescribeDDoSEventsRequest {
	s.InstanceIds = v
	return s
}

type DescribeDDoSEventsResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	DDoSEvents []*DescribeDDoSEventsResponseBodyDDoSEvents `json:"DDoSEvents,omitempty" xml:"DDoSEvents,omitempty" type:"Repeated"`
}

func (s DescribeDDoSEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBody) SetRequestId(v string) *DescribeDDoSEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetTotal(v int64) *DescribeDDoSEventsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetDDoSEvents(v []*DescribeDDoSEventsResponseBodyDDoSEvents) *DescribeDDoSEventsResponseBody {
	s.DDoSEvents = v
	return s
}

type DescribeDDoSEventsResponseBodyDDoSEvents struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port      *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Bps       *int64  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Pps       *int64  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeDDoSEventsResponseBodyDDoSEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseBodyDDoSEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetEndTime(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetStartTime(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetEventType(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetRegion(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Region = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetIp(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetPort(v string) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Port = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetBps(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Bps = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyDDoSEvents) SetPps(v int64) *DescribeDDoSEventsResponseBodyDDoSEvents {
	s.Pps = &v
	return s
}

type DescribeDDoSEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDoSEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDoSEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponse) SetHeaders(v map[string]*string) *DescribeDDoSEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSEventsResponse) SetBody(v *DescribeDDoSEventsResponseBody) *DescribeDDoSEventsResponse {
	s.Body = v
	return s
}

type DescribeDDosEventSrcIpRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Range     *int64  `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s DescribeDDosEventSrcIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventSrcIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpRequest) SetSourceIp(v string) *DescribeDDosEventSrcIpRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetEventType(v string) *DescribeDDosEventSrcIpRequest {
	s.EventType = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetStartTime(v int64) *DescribeDDosEventSrcIpRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetIp(v string) *DescribeDDosEventSrcIpRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDDosEventSrcIpRequest) SetRange(v int64) *DescribeDDosEventSrcIpRequest {
	s.Range = &v
	return s
}

type DescribeDDosEventSrcIpResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ips       []*DescribeDDosEventSrcIpResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s DescribeDDosEventSrcIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventSrcIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpResponseBody) SetRequestId(v string) *DescribeDDosEventSrcIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBody) SetIps(v []*DescribeDDosEventSrcIpResponseBodyIps) *DescribeDDosEventSrcIpResponseBody {
	s.Ips = v
	return s
}

type DescribeDDosEventSrcIpResponseBodyIps struct {
	SrcIp  *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	Isp    *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeDDosEventSrcIpResponseBodyIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventSrcIpResponseBodyIps) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) SetSrcIp(v string) *DescribeDDosEventSrcIpResponseBodyIps {
	s.SrcIp = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) SetAreaId(v string) *DescribeDDosEventSrcIpResponseBodyIps {
	s.AreaId = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) SetIsp(v string) *DescribeDDosEventSrcIpResponseBodyIps {
	s.Isp = &v
	return s
}

type DescribeDDosEventSrcIpResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDosEventSrcIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDosEventSrcIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDosEventSrcIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpResponse) SetHeaders(v map[string]*string) *DescribeDDosEventSrcIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventSrcIpResponse) SetBody(v *DescribeDDosEventSrcIpResponseBody) *DescribeDDosEventSrcIpResponse {
	s.Body = v
	return s
}

type DescribeDefenseCountStatisticsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDefenseCountStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsRequest) SetSourceIp(v string) *DescribeDefenseCountStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDefenseCountStatisticsResponseBody struct {
	RequestId              *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DefenseCountStatistics *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics `json:"DefenseCountStatistics,omitempty" xml:"DefenseCountStatistics,omitempty" type:"Struct"`
}

func (s DescribeDefenseCountStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetRequestId(v string) *DescribeDefenseCountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetDefenseCountStatistics(v *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) *DescribeDefenseCountStatisticsResponseBody {
	s.DefenseCountStatistics = v
	return s
}

type DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics struct {
	MaxUsableDefenseCountCurrentMonth    *int32 `json:"MaxUsableDefenseCountCurrentMonth,omitempty" xml:"MaxUsableDefenseCountCurrentMonth,omitempty"`
	FlowPackCountRemain                  *int32 `json:"FlowPackCountRemain,omitempty" xml:"FlowPackCountRemain,omitempty"`
	DefenseCountTotalUsageOfCurrentMonth *int32 `json:"DefenseCountTotalUsageOfCurrentMonth,omitempty" xml:"DefenseCountTotalUsageOfCurrentMonth,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetMaxUsableDefenseCountCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.MaxUsableDefenseCountCurrentMonth = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetFlowPackCountRemain(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.FlowPackCountRemain = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetDefenseCountTotalUsageOfCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.DefenseCountTotalUsageOfCurrentMonth = &v
	return s
}

type DescribeDefenseCountStatisticsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDefenseCountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseCountStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDefenseCountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) SetBody(v *DescribeDefenseCountStatisticsResponseBody) *DescribeDefenseCountStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDefenseRecordsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDefenseRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsRequest) SetSourceIp(v string) *DescribeDefenseRecordsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetLang(v string) *DescribeDefenseRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetResourceGroupId(v string) *DescribeDefenseRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetInstanceId(v string) *DescribeDefenseRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetStartTime(v int64) *DescribeDefenseRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetEndTime(v int64) *DescribeDefenseRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetPageNumber(v int32) *DescribeDefenseRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseRecordsRequest) SetPageSize(v int32) *DescribeDefenseRecordsRequest {
	s.PageSize = &v
	return s
}

type DescribeDefenseRecordsResponseBody struct {
	TotalCount     *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DefenseRecords []*DescribeDefenseRecordsResponseBodyDefenseRecords `json:"DefenseRecords,omitempty" xml:"DefenseRecords,omitempty" type:"Repeated"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsResponseBody) SetTotalCount(v int64) *DescribeDefenseRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBody) SetDefenseRecords(v []*DescribeDefenseRecordsResponseBodyDefenseRecords) *DescribeDefenseRecordsResponseBody {
	s.DefenseRecords = v
	return s
}

func (s *DescribeDefenseRecordsResponseBody) SetRequestId(v string) *DescribeDefenseRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDefenseRecordsResponseBodyDefenseRecords struct {
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EventCount *int32  `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AttackPeak *int64  `json:"AttackPeak,omitempty" xml:"AttackPeak,omitempty"`
}

func (s DescribeDefenseRecordsResponseBodyDefenseRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRecordsResponseBodyDefenseRecords) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetEndTime(v int64) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.EndTime = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetStatus(v int32) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetStartTime(v int64) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.StartTime = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetEventCount(v int32) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.EventCount = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetInstanceId(v string) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRecordsResponseBodyDefenseRecords) SetAttackPeak(v int64) *DescribeDefenseRecordsResponseBodyDefenseRecords {
	s.AttackPeak = &v
	return s
}

type DescribeDefenseRecordsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDefenseRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRecordsResponse) SetHeaders(v map[string]*string) *DescribeDefenseRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRecordsResponse) SetBody(v *DescribeDefenseRecordsResponseBody) *DescribeDefenseRecordsResponse {
	s.Body = v
	return s
}

type DescribeDomainAttackEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainAttackEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsRequest) SetSourceIp(v string) *DescribeDomainAttackEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetResourceGroupId(v string) *DescribeDomainAttackEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetStartTime(v int64) *DescribeDomainAttackEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetEndTime(v int64) *DescribeDomainAttackEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetDomain(v string) *DescribeDomainAttackEventsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetPageNumber(v int32) *DescribeDomainAttackEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetPageSize(v int32) *DescribeDomainAttackEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainAttackEventsResponseBody struct {
	TotalCount         *int64                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DomainAttackEvents []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents `json:"DomainAttackEvents,omitempty" xml:"DomainAttackEvents,omitempty" type:"Repeated"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainAttackEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseBody) SetTotalCount(v int64) *DescribeDomainAttackEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) SetDomainAttackEvents(v []*DescribeDomainAttackEventsResponseBodyDomainAttackEvents) *DescribeDomainAttackEventsResponseBody {
	s.DomainAttackEvents = v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) SetRequestId(v string) *DescribeDomainAttackEventsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainAttackEventsResponseBodyDomainAttackEvents struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	MaxQps    *int64  `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
}

func (s DescribeDomainAttackEventsResponseBodyDomainAttackEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseBodyDomainAttackEvents) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetEndTime(v int64) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetStartTime(v int64) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetDomain(v string) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyDomainAttackEvents) SetMaxQps(v int64) *DescribeDomainAttackEventsResponseBodyDomainAttackEvents {
	s.MaxQps = &v
	return s
}

type DescribeDomainAttackEventsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainAttackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainAttackEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackEventsResponse) SetBody(v *DescribeDomainAttackEventsResponseBody) *DescribeDomainAttackEventsResponse {
	s.Body = v
	return s
}

type DescribeDomainOverviewRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewRequest) SetSourceIp(v string) *DescribeDomainOverviewRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetResourceGroupId(v string) *DescribeDomainOverviewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetStartTime(v int64) *DescribeDomainOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetEndTime(v int64) *DescribeDomainOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetDomain(v string) *DescribeDomainOverviewRequest {
	s.Domain = &v
	return s
}

type DescribeDomainOverviewResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxHttps  *int64  `json:"MaxHttps,omitempty" xml:"MaxHttps,omitempty"`
	MaxHttp   *int64  `json:"MaxHttp,omitempty" xml:"MaxHttp,omitempty"`
}

func (s DescribeDomainOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewResponseBody) SetRequestId(v string) *DescribeDomainOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) SetMaxHttps(v int64) *DescribeDomainOverviewResponseBody {
	s.MaxHttps = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) SetMaxHttp(v int64) *DescribeDomainOverviewResponseBody {
	s.MaxHttp = &v
	return s
}

type DescribeDomainOverviewResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewResponse) SetHeaders(v map[string]*string) *DescribeDomainOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainOverviewResponse) SetBody(v *DescribeDomainOverviewResponseBody) *DescribeDomainOverviewResponse {
	s.Body = v
	return s
}

type DescribeDomainQPSListRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval        *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainQPSListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQPSListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListRequest) SetSourceIp(v string) *DescribeDomainQPSListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetResourceGroupId(v string) *DescribeDomainQPSListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetStartTime(v int64) *DescribeDomainQPSListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetEndTime(v int64) *DescribeDomainQPSListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetInterval(v int64) *DescribeDomainQPSListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetDomain(v string) *DescribeDomainQPSListRequest {
	s.Domain = &v
	return s
}

type DescribeDomainQPSListResponseBody struct {
	DomainQPSList []*DescribeDomainQPSListResponseBodyDomainQPSList `json:"DomainQPSList,omitempty" xml:"DomainQPSList,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainQPSListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQPSListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListResponseBody) SetDomainQPSList(v []*DescribeDomainQPSListResponseBodyDomainQPSList) *DescribeDomainQPSListResponseBody {
	s.DomainQPSList = v
	return s
}

func (s *DescribeDomainQPSListResponseBody) SetRequestId(v string) *DescribeDomainQPSListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainQPSListResponseBodyDomainQPSList struct {
	Index        *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	Time         *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	MaxAttackQps *int64 `json:"MaxAttackQps,omitempty" xml:"MaxAttackQps,omitempty"`
	AttackQps    *int64 `json:"AttackQps,omitempty" xml:"AttackQps,omitempty"`
	MaxQps       *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	MaxNormalQps *int64 `json:"MaxNormalQps,omitempty" xml:"MaxNormalQps,omitempty"`
	TotalQps     *int64 `json:"TotalQps,omitempty" xml:"TotalQps,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CacheHits    *int64 `json:"CacheHits,omitempty" xml:"CacheHits,omitempty"`
}

func (s DescribeDomainQPSListResponseBodyDomainQPSList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQPSListResponseBodyDomainQPSList) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetIndex(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.Index = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetTime(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.Time = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetMaxAttackQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.MaxAttackQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetAttackQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.AttackQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetMaxQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetMaxNormalQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.MaxNormalQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetTotalQps(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.TotalQps = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetTotalCount(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainQPSListResponseBodyDomainQPSList) SetCacheHits(v int64) *DescribeDomainQPSListResponseBodyDomainQPSList {
	s.CacheHits = &v
	return s
}

type DescribeDomainQPSListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainQPSListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQPSListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQPSListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListResponse) SetHeaders(v map[string]*string) *DescribeDomainQPSListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQPSListResponse) SetBody(v *DescribeDomainQPSListResponseBody) *DescribeDomainQPSListResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsWithCacheRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainQpsWithCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheRequest) SetSourceIp(v string) *DescribeDomainQpsWithCacheRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetResourceGroupId(v string) *DescribeDomainQpsWithCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetDomain(v string) *DescribeDomainQpsWithCacheRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetStartTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetEndTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainQpsWithCacheResponseBody struct {
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" type:"Repeated"`
	CcJsQps       []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" type:"Repeated"`
	Blocks        []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" type:"Repeated"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" type:"Repeated"`
	Totals        []*string `json:"Totals,omitempty" xml:"Totals,omitempty" type:"Repeated"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CcBlockQps    []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" type:"Repeated"`
	Interval      *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" type:"Repeated"`
	CacheHits     []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsWithCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetIpBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRequestId(v string) *DescribeDomainQpsWithCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetTotals(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Totals = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetStartTime(v int64) *DescribeDomainQpsWithCacheResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetInterval(v int32) *DescribeDomainQpsWithCacheResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRegionBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCacheHits(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CacheHits = v
	return s
}

type DescribeDomainQpsWithCacheResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainQpsWithCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQpsWithCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsWithCacheResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetBody(v *DescribeDomainQpsWithCacheResponseBody) *DescribeDomainQpsWithCacheResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s *DescribeDomainsRequest) SetLang(v string) *DescribeDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceIds(v []*string) *DescribeDomainsRequest {
	s.InstanceIds = v
	return s
}

type DescribeDomainsResponseBody struct {
	Domains   []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*string) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
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

type DescribeDomainStatusCodeCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainStatusCodeCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeCountRequest) SetSourceIp(v string) *DescribeDomainStatusCodeCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetResourceGroupId(v string) *DescribeDomainStatusCodeCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetStartTime(v int64) *DescribeDomainStatusCodeCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetEndTime(v int64) *DescribeDomainStatusCodeCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetDomain(v string) *DescribeDomainStatusCodeCountRequest {
	s.Domain = &v
	return s
}

type DescribeDomainStatusCodeCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status501 *int64  `json:"Status501,omitempty" xml:"Status501,omitempty"`
	Status502 *int64  `json:"Status502,omitempty" xml:"Status502,omitempty"`
	Status403 *int64  `json:"Status403,omitempty" xml:"Status403,omitempty"`
	Status503 *int64  `json:"Status503,omitempty" xml:"Status503,omitempty"`
	Status404 *int64  `json:"Status404,omitempty" xml:"Status404,omitempty"`
	Status405 *int64  `json:"Status405,omitempty" xml:"Status405,omitempty"`
	Status504 *int64  `json:"Status504,omitempty" xml:"Status504,omitempty"`
	Status2XX *int64  `json:"Status2XX,omitempty" xml:"Status2XX,omitempty"`
	Status200 *int64  `json:"Status200,omitempty" xml:"Status200,omitempty"`
	Status3XX *int64  `json:"Status3XX,omitempty" xml:"Status3XX,omitempty"`
	Status4XX *int64  `json:"Status4XX,omitempty" xml:"Status4XX,omitempty"`
	Status5XX *int64  `json:"Status5XX,omitempty" xml:"Status5XX,omitempty"`
}

func (s DescribeDomainStatusCodeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetRequestId(v string) *DescribeDomainStatusCodeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus501(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status501 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus502(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status502 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus403(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status403 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus503(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status503 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus404(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status404 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus405(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status405 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus504(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status504 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus2XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status2XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus200(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status200 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus3XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status3XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus4XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status4XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus5XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status5XX = &v
	return s
}

type DescribeDomainStatusCodeCountResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainStatusCodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainStatusCodeCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeCountResponse) SetHeaders(v map[string]*string) *DescribeDomainStatusCodeCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatusCodeCountResponse) SetBody(v *DescribeDomainStatusCodeCountResponseBody) *DescribeDomainStatusCodeCountResponse {
	s.Body = v
	return s
}

type DescribeDomainStatusCodeListRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval        *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	QueryType       *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
}

func (s DescribeDomainStatusCodeListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListRequest) SetSourceIp(v string) *DescribeDomainStatusCodeListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetResourceGroupId(v string) *DescribeDomainStatusCodeListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetStartTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetEndTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetInterval(v int64) *DescribeDomainStatusCodeListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetDomain(v string) *DescribeDomainStatusCodeListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetQueryType(v string) *DescribeDomainStatusCodeListRequest {
	s.QueryType = &v
	return s
}

type DescribeDomainStatusCodeListResponseBody struct {
	RequestId      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusCodeList []*DescribeDomainStatusCodeListResponseBodyStatusCodeList `json:"StatusCodeList,omitempty" xml:"StatusCodeList,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatusCodeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponseBody) SetRequestId(v string) *DescribeDomainStatusCodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBody) SetStatusCodeList(v []*DescribeDomainStatusCodeListResponseBodyStatusCodeList) *DescribeDomainStatusCodeListResponseBody {
	s.StatusCodeList = v
	return s
}

type DescribeDomainStatusCodeListResponseBodyStatusCodeList struct {
	Status502 *int64 `json:"Status502,omitempty" xml:"Status502,omitempty"`
	Index     *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	Time      *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	Status405 *int64 `json:"Status405,omitempty" xml:"Status405,omitempty"`
	Status3XX *int64 `json:"Status3XX,omitempty" xml:"Status3XX,omitempty"`
	Status503 *int64 `json:"Status503,omitempty" xml:"Status503,omitempty"`
	Status4XX *int64 `json:"Status4XX,omitempty" xml:"Status4XX,omitempty"`
	Status2XX *int64 `json:"Status2XX,omitempty" xml:"Status2XX,omitempty"`
	Status5XX *int64 `json:"Status5XX,omitempty" xml:"Status5XX,omitempty"`
	Status504 *int64 `json:"Status504,omitempty" xml:"Status504,omitempty"`
	Status403 *int64 `json:"Status403,omitempty" xml:"Status403,omitempty"`
	Status200 *int64 `json:"Status200,omitempty" xml:"Status200,omitempty"`
	Status404 *int64 `json:"Status404,omitempty" xml:"Status404,omitempty"`
	Status501 *int64 `json:"Status501,omitempty" xml:"Status501,omitempty"`
}

func (s DescribeDomainStatusCodeListResponseBodyStatusCodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponseBodyStatusCodeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus502(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status502 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetIndex(v int32) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Index = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetTime(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Time = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus405(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status405 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus3XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status3XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus503(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status503 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus4XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status4XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus2XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status2XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus5XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status5XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus504(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status504 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus403(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status403 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus200(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status200 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus404(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status404 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus501(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status501 = &v
	return s
}

type DescribeDomainStatusCodeListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainStatusCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainStatusCodeListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponse) SetHeaders(v map[string]*string) *DescribeDomainStatusCodeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatusCodeListResponse) SetBody(v *DescribeDomainStatusCodeListResponseBody) *DescribeDomainStatusCodeListResponse {
	s.Body = v
	return s
}

type DescribeDomainTopAttackListRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainTopAttackListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopAttackListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListRequest) SetSourceIp(v string) *DescribeDomainTopAttackListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) SetResourceGroupId(v string) *DescribeDomainTopAttackListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) SetStartTime(v int64) *DescribeDomainTopAttackListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopAttackListRequest) SetEndTime(v int64) *DescribeDomainTopAttackListRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainTopAttackListResponseBody struct {
	AttackList []*DescribeDomainTopAttackListResponseBodyAttackList `json:"AttackList,omitempty" xml:"AttackList,omitempty" type:"Repeated"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopAttackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopAttackListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListResponseBody) SetAttackList(v []*DescribeDomainTopAttackListResponseBodyAttackList) *DescribeDomainTopAttackListResponseBody {
	s.AttackList = v
	return s
}

func (s *DescribeDomainTopAttackListResponseBody) SetRequestId(v string) *DescribeDomainTopAttackListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainTopAttackListResponseBodyAttackList struct {
	Attack *int64  `json:"Attack,omitempty" xml:"Attack,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Count  *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainTopAttackListResponseBodyAttackList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopAttackListResponseBodyAttackList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) SetAttack(v int64) *DescribeDomainTopAttackListResponseBodyAttackList {
	s.Attack = &v
	return s
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) SetDomain(v string) *DescribeDomainTopAttackListResponseBodyAttackList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopAttackListResponseBodyAttackList) SetCount(v int64) *DescribeDomainTopAttackListResponseBodyAttackList {
	s.Count = &v
	return s
}

type DescribeDomainTopAttackListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainTopAttackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainTopAttackListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopAttackListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopAttackListResponse) SetHeaders(v map[string]*string) *DescribeDomainTopAttackListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopAttackListResponse) SetBody(v *DescribeDomainTopAttackListResponseBody) *DescribeDomainTopAttackListResponse {
	s.Body = v
	return s
}

type DescribeDomainViewSourceCountriesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainViewSourceCountriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesRequest) SetSourceIp(v string) *DescribeDomainViewSourceCountriesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetResourceGroupId(v string) *DescribeDomainViewSourceCountriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetStartTime(v int64) *DescribeDomainViewSourceCountriesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetEndTime(v int64) *DescribeDomainViewSourceCountriesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetDomain(v string) *DescribeDomainViewSourceCountriesRequest {
	s.Domain = &v
	return s
}

type DescribeDomainViewSourceCountriesResponseBody struct {
	RequestId      *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceCountrys []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys `json:"SourceCountrys,omitempty" xml:"SourceCountrys,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewSourceCountriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesResponseBody) SetRequestId(v string) *DescribeDomainViewSourceCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponseBody) SetSourceCountrys(v []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys) *DescribeDomainViewSourceCountriesResponseBody {
	s.SourceCountrys = v
	return s
}

type DescribeDomainViewSourceCountriesResponseBodySourceCountrys struct {
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainViewSourceCountriesResponseBodySourceCountrys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesResponseBodySourceCountrys) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) SetCountryId(v string) *DescribeDomainViewSourceCountriesResponseBodySourceCountrys {
	s.CountryId = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) SetCount(v int64) *DescribeDomainViewSourceCountriesResponseBodySourceCountrys {
	s.Count = &v
	return s
}

type DescribeDomainViewSourceCountriesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainViewSourceCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainViewSourceCountriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesResponse) SetHeaders(v map[string]*string) *DescribeDomainViewSourceCountriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponse) SetBody(v *DescribeDomainViewSourceCountriesResponseBody) *DescribeDomainViewSourceCountriesResponse {
	s.Body = v
	return s
}

type DescribeDomainViewSourceProvincesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainViewSourceProvincesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesRequest) SetSourceIp(v string) *DescribeDomainViewSourceProvincesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesRequest) SetResourceGroupId(v string) *DescribeDomainViewSourceProvincesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesRequest) SetStartTime(v int64) *DescribeDomainViewSourceProvincesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesRequest) SetEndTime(v int64) *DescribeDomainViewSourceProvincesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesRequest) SetDomain(v string) *DescribeDomainViewSourceProvincesRequest {
	s.Domain = &v
	return s
}

type DescribeDomainViewSourceProvincesResponseBody struct {
	SourceProvinces []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces `json:"SourceProvinces,omitempty" xml:"SourceProvinces,omitempty" type:"Repeated"`
	RequestId       *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainViewSourceProvincesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesResponseBody) SetSourceProvinces(v []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces) *DescribeDomainViewSourceProvincesResponseBody {
	s.SourceProvinces = v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponseBody) SetRequestId(v string) *DescribeDomainViewSourceProvincesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainViewSourceProvincesResponseBodySourceProvinces struct {
	ProvinceId *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainViewSourceProvincesResponseBodySourceProvinces) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesResponseBodySourceProvinces) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) SetProvinceId(v string) *DescribeDomainViewSourceProvincesResponseBodySourceProvinces {
	s.ProvinceId = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) SetCount(v int64) *DescribeDomainViewSourceProvincesResponseBodySourceProvinces {
	s.Count = &v
	return s
}

type DescribeDomainViewSourceProvincesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainViewSourceProvincesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainViewSourceProvincesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesResponse) SetHeaders(v map[string]*string) *DescribeDomainViewSourceProvincesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponse) SetBody(v *DescribeDomainViewSourceProvincesResponseBody) *DescribeDomainViewSourceProvincesResponse {
	s.Body = v
	return s
}

type DescribeDomainViewTopCostTimeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Top             *int32  `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeDomainViewTopCostTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeRequest) SetSourceIp(v string) *DescribeDomainViewTopCostTimeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetResourceGroupId(v string) *DescribeDomainViewTopCostTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetStartTime(v int64) *DescribeDomainViewTopCostTimeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetEndTime(v int64) *DescribeDomainViewTopCostTimeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetDomain(v string) *DescribeDomainViewTopCostTimeRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetTop(v int32) *DescribeDomainViewTopCostTimeRequest {
	s.Top = &v
	return s
}

type DescribeDomainViewTopCostTimeResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlList   []*DescribeDomainViewTopCostTimeResponseBodyUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewTopCostTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeResponseBody) SetRequestId(v string) *DescribeDomainViewTopCostTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBody) SetUrlList(v []*DescribeDomainViewTopCostTimeResponseBodyUrlList) *DescribeDomainViewTopCostTimeResponseBody {
	s.UrlList = v
	return s
}

type DescribeDomainViewTopCostTimeResponseBodyUrlList struct {
	Domain   *string  `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CostTime *float32 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	Url      *string  `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDomainViewTopCostTimeResponseBodyUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeResponseBodyUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) SetDomain(v string) *DescribeDomainViewTopCostTimeResponseBodyUrlList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) SetCostTime(v float32) *DescribeDomainViewTopCostTimeResponseBodyUrlList {
	s.CostTime = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) SetUrl(v string) *DescribeDomainViewTopCostTimeResponseBodyUrlList {
	s.Url = &v
	return s
}

type DescribeDomainViewTopCostTimeResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainViewTopCostTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainViewTopCostTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeResponse) SetHeaders(v map[string]*string) *DescribeDomainViewTopCostTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponse) SetBody(v *DescribeDomainViewTopCostTimeResponseBody) *DescribeDomainViewTopCostTimeResponse {
	s.Body = v
	return s
}

type DescribeDomainViewTopUrlRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Top             *int32  `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeDomainViewTopUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlRequest) SetSourceIp(v string) *DescribeDomainViewTopUrlRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetResourceGroupId(v string) *DescribeDomainViewTopUrlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetStartTime(v int64) *DescribeDomainViewTopUrlRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetEndTime(v int64) *DescribeDomainViewTopUrlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetDomain(v string) *DescribeDomainViewTopUrlRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetTop(v int32) *DescribeDomainViewTopUrlRequest {
	s.Top = &v
	return s
}

type DescribeDomainViewTopUrlResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlList   []*DescribeDomainViewTopUrlResponseBodyUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewTopUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlResponseBody) SetRequestId(v string) *DescribeDomainViewTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBody) SetUrlList(v []*DescribeDomainViewTopUrlResponseBodyUrlList) *DescribeDomainViewTopUrlResponseBody {
	s.UrlList = v
	return s
}

type DescribeDomainViewTopUrlResponseBodyUrlList struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Count  *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainViewTopUrlResponseBodyUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopUrlResponseBodyUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) SetDomain(v string) *DescribeDomainViewTopUrlResponseBodyUrlList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) SetUrl(v string) *DescribeDomainViewTopUrlResponseBodyUrlList {
	s.Url = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) SetCount(v int64) *DescribeDomainViewTopUrlResponseBodyUrlList {
	s.Count = &v
	return s
}

type DescribeDomainViewTopUrlResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainViewTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainViewTopUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainViewTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlResponse) SetHeaders(v map[string]*string) *DescribeDomainViewTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewTopUrlResponse) SetBody(v *DescribeDomainViewTopUrlResponseBody) *DescribeDomainViewTopUrlResponse {
	s.Body = v
	return s
}

type DescribeElasticBandwidthSpecRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeElasticBandwidthSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecRequest) SetSourceIp(v string) *DescribeElasticBandwidthSpecRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest {
	s.InstanceId = &v
	return s
}

type DescribeElasticBandwidthSpecResponseBody struct {
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ElasticBandwidthSpec []*string `json:"ElasticBandwidthSpec,omitempty" xml:"ElasticBandwidthSpec,omitempty" type:"Repeated"`
}

func (s DescribeElasticBandwidthSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetRequestId(v string) *DescribeElasticBandwidthSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetElasticBandwidthSpec(v []*string) *DescribeElasticBandwidthSpecResponseBody {
	s.ElasticBandwidthSpec = v
	return s
}

type DescribeElasticBandwidthSpecResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeElasticBandwidthSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticBandwidthSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponse) SetHeaders(v map[string]*string) *DescribeElasticBandwidthSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) SetBody(v *DescribeElasticBandwidthSpecResponseBody) *DescribeElasticBandwidthSpecResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckListRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s DescribeHealthCheckListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListRequest) SetSourceIp(v string) *DescribeHealthCheckListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckListRequest) SetNetworkRules(v string) *DescribeHealthCheckListRequest {
	s.NetworkRules = &v
	return s
}

type DescribeHealthCheckListResponseBody struct {
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HealthCheckList []*DescribeHealthCheckListResponseBodyHealthCheckList `json:"HealthCheckList,omitempty" xml:"HealthCheckList,omitempty" type:"Repeated"`
}

func (s DescribeHealthCheckListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBody) SetRequestId(v string) *DescribeHealthCheckListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBody) SetHealthCheckList(v []*DescribeHealthCheckListResponseBodyHealthCheckList) *DescribeHealthCheckListResponseBody {
	s.HealthCheckList = v
	return s
}

type DescribeHealthCheckListResponseBodyHealthCheckList struct {
	FrontendPort *int32                                                         `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Protocol     *string                                                        `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InstanceId   *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	HealthCheck  *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
}

func (s DescribeHealthCheckListResponseBodyHealthCheckList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyHealthCheckList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetFrontendPort(v int32) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetProtocol(v string) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetInstanceId(v string) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckList) SetHealthCheck(v *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) *DescribeHealthCheckListResponseBodyHealthCheckList {
	s.HealthCheck = v
	return s
}

type DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck struct {
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Up       *int32  `json:"Up,omitempty" xml:"Up,omitempty"`
	Down     *int32  `json:"Down,omitempty" xml:"Down,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetTimeout(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetType(v string) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetDomain(v string) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetInterval(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetUp(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Up = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetDown(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetPort(v int32) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck) SetUri(v string) *DescribeHealthCheckListResponseBodyHealthCheckListHealthCheck {
	s.Uri = &v
	return s
}

type DescribeHealthCheckListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHealthCheckListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthCheckListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckListResponse) SetBody(v *DescribeHealthCheckListResponseBody) *DescribeHealthCheckListResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckStatusRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s DescribeHealthCheckStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusRequest) SetSourceIp(v string) *DescribeHealthCheckStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckStatusRequest) SetNetworkRules(v string) *DescribeHealthCheckStatusRequest {
	s.NetworkRules = &v
	return s
}

type DescribeHealthCheckStatusResponseBody struct {
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HealthCheckStatus []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus `json:"HealthCheckStatus,omitempty" xml:"HealthCheckStatus,omitempty" type:"Repeated"`
}

func (s DescribeHealthCheckStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponseBody) SetRequestId(v string) *DescribeHealthCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBody) SetHealthCheckStatus(v []*DescribeHealthCheckStatusResponseBodyHealthCheckStatus) *DescribeHealthCheckStatusResponseBody {
	s.HealthCheckStatus = v
	return s
}

type DescribeHealthCheckStatusResponseBodyHealthCheckStatus struct {
	Status               *string                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	FrontendPort         *int32                                                                        `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Protocol             *string                                                                       `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InstanceId           *string                                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RealServerStatusList []*DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList `json:"RealServerStatusList,omitempty" xml:"RealServerStatusList,omitempty" type:"Repeated"`
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetStatus(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetFrontendPort(v int32) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetProtocol(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetInstanceId(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatus) SetRealServerStatusList(v []*DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) *DescribeHealthCheckStatusResponseBodyHealthCheckStatus {
	s.RealServerStatusList = v
	return s
}

type DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) SetStatus(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList) SetAddress(v string) *DescribeHealthCheckStatusResponseBodyHealthCheckStatusRealServerStatusList {
	s.Address = &v
	return s
}

type DescribeHealthCheckStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHealthCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthCheckStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckStatusResponse) SetBody(v *DescribeHealthCheckStatusResponseBody) *DescribeHealthCheckStatusResponse {
	s.Body = v
	return s
}

type DescribeInstanceDetailsRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsRequest) SetSourceIp(v string) *DescribeInstanceDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceDetailsRequest) SetInstanceIds(v []*string) *DescribeInstanceDetailsRequest {
	s.InstanceIds = v
	return s
}

type DescribeInstanceDetailsResponseBody struct {
	InstanceDetails []*DescribeInstanceDetailsResponseBodyInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBody) SetInstanceDetails(v []*DescribeInstanceDetailsResponseBodyInstanceDetails) *DescribeInstanceDetailsResponseBody {
	s.InstanceDetails = v
	return s
}

func (s *DescribeInstanceDetailsResponseBody) SetRequestId(v string) *DescribeInstanceDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceDetailsResponseBodyInstanceDetails struct {
	Line       *string                                                       `json:"Line,omitempty" xml:"Line,omitempty"`
	InstanceId *string                                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EipInfos   []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos `json:"EipInfos,omitempty" xml:"EipInfos,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetLine(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.Line = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetInstanceId(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetEipInfos(v []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.EipInfos = v
	return s
}

type DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Eip    *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetStatus(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.Status = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetEip(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.Eip = &v
	return s
}

type DescribeInstanceDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponse) SetHeaders(v map[string]*string) *DescribeInstanceDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDetailsResponse) SetBody(v *DescribeInstanceDetailsResponseBody) *DescribeInstanceDetailsResponse {
	s.Body = v
	return s
}

type DescribeInstanceIdsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Edition         *int32    `json:"Edition,omitempty" xml:"Edition,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIdsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsRequest) SetSourceIp(v string) *DescribeInstanceIdsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceIdsRequest) SetResourceGroupId(v string) *DescribeInstanceIdsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceIdsRequest) SetEdition(v int32) *DescribeInstanceIdsRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceIdsRequest) SetInstanceIds(v []*string) *DescribeInstanceIdsRequest {
	s.InstanceIds = v
	return s
}

type DescribeInstanceIdsResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceIds []*DescribeInstanceIdsResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsResponseBody) SetRequestId(v string) *DescribeInstanceIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceIdsResponseBody) SetInstanceIds(v []*DescribeInstanceIdsResponseBodyInstanceIds) *DescribeInstanceIdsResponseBody {
	s.InstanceIds = v
	return s
}

type DescribeInstanceIdsResponseBodyInstanceIds struct {
	Edition    *int32  `json:"Edition,omitempty" xml:"Edition,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceIdsResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIdsResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetEdition(v int32) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetRemark(v string) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceIdsResponseBodyInstanceIds) SetInstanceId(v string) *DescribeInstanceIdsResponseBodyInstanceIds {
	s.InstanceId = &v
	return s
}

type DescribeInstanceIdsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsResponse) SetHeaders(v map[string]*string) *DescribeInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceIdsResponse) SetBody(v *DescribeInstanceIdsResponseBody) *DescribeInstanceIdsResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	SourceIp        *string                        `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNumber      *string                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Ip              *string                        `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Remark          *string                        `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Edition         *int32                         `json:"Edition,omitempty" xml:"Edition,omitempty"`
	Enabled         *int32                         `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExpireStartTime *int64                         `json:"ExpireStartTime,omitempty" xml:"ExpireStartTime,omitempty"`
	ExpireEndTime   *int64                         `json:"ExpireEndTime,omitempty" xml:"ExpireEndTime,omitempty"`
	InstanceIds     []*string                      `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	Status          []*int                         `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetSourceIp(v string) *DescribeInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v string) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetIp(v string) *DescribeInstancesRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesRequest) SetRemark(v string) *DescribeInstancesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesRequest) SetEdition(v int32) *DescribeInstancesRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnabled(v int32) *DescribeInstancesRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireStartTime(v int64) *DescribeInstancesRequest {
	s.ExpireStartTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireEndTime(v int64) *DescribeInstancesRequest {
	s.ExpireEndTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v []*string) *DescribeInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v []*int) *DescribeInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DebtStatus     *int32  `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	Edition        *int32  `json:"Edition,omitempty" xml:"Edition,omitempty"`
	ExpireTime     *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Enabled        *int32  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ConnInstanceId *string `json:"ConnInstanceId,omitempty" xml:"ConnInstanceId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDebtStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEdition(v int32) *DescribeInstancesResponseBodyInstances {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRemark(v string) *DescribeInstancesResponseBodyInstances {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCreateTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnabled(v int32) *DescribeInstancesResponseBodyInstances {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetConnInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.ConnInstanceId = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetSourceIp(v string) *DescribeInstanceSpecsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIds(v []*string) *DescribeInstanceSpecsRequest {
	s.InstanceIds = v
	return s
}

type DescribeInstanceSpecsResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	BaseBandwidth    *int32  `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	QpsLimit         *int32  `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty"`
	BandwidthMbps    *int32  `json:"BandwidthMbps,omitempty" xml:"BandwidthMbps,omitempty"`
	DefenseCount     *int32  `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	SiteLimit        *int32  `json:"SiteLimit,omitempty" xml:"SiteLimit,omitempty"`
	PortLimit        *int32  `json:"PortLimit,omitempty" xml:"PortLimit,omitempty"`
	ElasticBandwidth *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	FunctionVersion  *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DomainLimit      *int32  `json:"DomainLimit,omitempty" xml:"DomainLimit,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBaseBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetQpsLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.QpsLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBandwidthMbps(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BandwidthMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetSiteLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.SiteLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPortLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PortLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetFunctionVersion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDomainLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DomainLimit = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeInstanceStatisticsRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetInstanceIds(v []*string) *DescribeInstanceStatisticsRequest {
	s.InstanceIds = v
	return s
}

type DescribeInstanceStatisticsResponseBody struct {
	InstanceStatistics []*DescribeInstanceStatisticsResponseBodyInstanceStatistics `json:"InstanceStatistics,omitempty" xml:"InstanceStatistics,omitempty" type:"Repeated"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBody) SetInstanceStatistics(v []*DescribeInstanceStatisticsResponseBodyInstanceStatistics) *DescribeInstanceStatisticsResponseBody {
	s.InstanceStatistics = v
	return s
}

func (s *DescribeInstanceStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceStatisticsResponseBodyInstanceStatistics struct {
	DomainUsage       *int32  `json:"DomainUsage,omitempty" xml:"DomainUsage,omitempty"`
	DefenseCountUsage *int32  `json:"DefenseCountUsage,omitempty" xml:"DefenseCountUsage,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SiteUsage         *int32  `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
	PortUsage         *int32  `json:"PortUsage,omitempty" xml:"PortUsage,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDomainUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DomainUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDefenseCountUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DefenseCountUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetInstanceId(v string) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetSiteUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.SiteUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetPortUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.PortUsage = &v
	return s
}

type DescribeInstanceStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetBody(v *DescribeInstanceStatisticsResponseBody) *DescribeInstanceStatisticsResponse {
	s.Body = v
	return s
}

type DescribeInstanceStatusRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductType *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusRequest) SetSourceIp(v string) *DescribeInstanceStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetInstanceId(v string) *DescribeInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetProductType(v int32) *DescribeInstanceStatusRequest {
	s.ProductType = &v
	return s
}

type DescribeInstanceStatusResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus *int32  `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
}

func (s DescribeInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponseBody) SetRequestId(v string) *DescribeInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetInstanceId(v string) *DescribeInstanceStatusResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatusResponseBody) SetInstanceStatus(v int32) *DescribeInstanceStatusResponseBody {
	s.InstanceStatus = &v
	return s
}

type DescribeInstanceStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStatusResponse) SetBody(v *DescribeInstanceStatusResponseBody) *DescribeInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeL7RsPolicyRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RealServers     []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
}

func (s DescribeL7RsPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeL7RsPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyRequest) SetSourceIp(v string) *DescribeL7RsPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeL7RsPolicyRequest) SetResourceGroupId(v string) *DescribeL7RsPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeL7RsPolicyRequest) SetDomain(v string) *DescribeL7RsPolicyRequest {
	s.Domain = &v
	return s
}

func (s *DescribeL7RsPolicyRequest) SetRealServers(v []*string) *DescribeL7RsPolicyRequest {
	s.RealServers = v
	return s
}

type DescribeL7RsPolicyResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Attributes []*DescribeL7RsPolicyResponseBodyAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	ProxyMode  *string                                     `json:"ProxyMode,omitempty" xml:"ProxyMode,omitempty"`
}

func (s DescribeL7RsPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeL7RsPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponseBody) SetRequestId(v string) *DescribeL7RsPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) SetAttributes(v []*DescribeL7RsPolicyResponseBodyAttributes) *DescribeL7RsPolicyResponseBody {
	s.Attributes = v
	return s
}

func (s *DescribeL7RsPolicyResponseBody) SetProxyMode(v string) *DescribeL7RsPolicyResponseBody {
	s.ProxyMode = &v
	return s
}

type DescribeL7RsPolicyResponseBodyAttributes struct {
	RsType     *int32                                             `json:"RsType,omitempty" xml:"RsType,omitempty"`
	Attribute  *DescribeL7RsPolicyResponseBodyAttributesAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	RealServer *string                                            `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
}

func (s DescribeL7RsPolicyResponseBodyAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeL7RsPolicyResponseBodyAttributes) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) SetRsType(v int32) *DescribeL7RsPolicyResponseBodyAttributes {
	s.RsType = &v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) SetAttribute(v *DescribeL7RsPolicyResponseBodyAttributesAttribute) *DescribeL7RsPolicyResponseBodyAttributes {
	s.Attribute = v
	return s
}

func (s *DescribeL7RsPolicyResponseBodyAttributes) SetRealServer(v string) *DescribeL7RsPolicyResponseBodyAttributes {
	s.RealServer = &v
	return s
}

type DescribeL7RsPolicyResponseBodyAttributesAttribute struct {
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeL7RsPolicyResponseBodyAttributesAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeL7RsPolicyResponseBodyAttributesAttribute) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponseBodyAttributesAttribute) SetWeight(v int32) *DescribeL7RsPolicyResponseBodyAttributesAttribute {
	s.Weight = &v
	return s
}

type DescribeL7RsPolicyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeL7RsPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeL7RsPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeL7RsPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyResponse) SetHeaders(v map[string]*string) *DescribeL7RsPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeL7RsPolicyResponse) SetBody(v *DescribeL7RsPolicyResponseBody) *DescribeL7RsPolicyResponse {
	s.Body = v
	return s
}

type DescribeLogStoreExistStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeLogStoreExistStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusRequest) SetSourceIp(v string) *DescribeLogStoreExistStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetLang(v string) *DescribeLogStoreExistStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeLogStoreExistStatusResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExistStatus *bool   `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
}

func (s DescribeLogStoreExistStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponseBody) SetRequestId(v string) *DescribeLogStoreExistStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponseBody) SetExistStatus(v bool) *DescribeLogStoreExistStatusResponseBody {
	s.ExistStatus = &v
	return s
}

type DescribeLogStoreExistStatusResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogStoreExistStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogStoreExistStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponse) SetHeaders(v map[string]*string) *DescribeLogStoreExistStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) SetBody(v *DescribeLogStoreExistStatusResponseBody) *DescribeLogStoreExistStatusResponse {
	s.Body = v
	return s
}

type DescribeNetworkRegionBlockRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNetworkRegionBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRegionBlockRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockRequest) SetSourceIp(v string) *DescribeNetworkRegionBlockRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNetworkRegionBlockRequest) SetInstanceId(v string) *DescribeNetworkRegionBlockRequest {
	s.InstanceId = &v
	return s
}

type DescribeNetworkRegionBlockResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Config    *DescribeNetworkRegionBlockResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
}

func (s DescribeNetworkRegionBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRegionBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockResponseBody) SetRequestId(v string) *DescribeNetworkRegionBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBody) SetConfig(v *DescribeNetworkRegionBlockResponseBodyConfig) *DescribeNetworkRegionBlockResponseBody {
	s.Config = v
	return s
}

type DescribeNetworkRegionBlockResponseBodyConfig struct {
	RegionBlockSwitch *string   `json:"RegionBlockSwitch,omitempty" xml:"RegionBlockSwitch,omitempty"`
	Provinces         []*string `json:"Provinces,omitempty" xml:"Provinces,omitempty" type:"Repeated"`
	Countries         []*string `json:"Countries,omitempty" xml:"Countries,omitempty" type:"Repeated"`
}

func (s DescribeNetworkRegionBlockResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRegionBlockResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) SetRegionBlockSwitch(v string) *DescribeNetworkRegionBlockResponseBodyConfig {
	s.RegionBlockSwitch = &v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) SetProvinces(v []*string) *DescribeNetworkRegionBlockResponseBodyConfig {
	s.Provinces = v
	return s
}

func (s *DescribeNetworkRegionBlockResponseBodyConfig) SetCountries(v []*string) *DescribeNetworkRegionBlockResponseBodyConfig {
	s.Countries = v
	return s
}

type DescribeNetworkRegionBlockResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkRegionBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkRegionBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRegionBlockResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRegionBlockResponse) SetHeaders(v map[string]*string) *DescribeNetworkRegionBlockResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRegionBlockResponse) SetBody(v *DescribeNetworkRegionBlockResponseBody) *DescribeNetworkRegionBlockResponse {
	s.Body = v
	return s
}

type DescribeNetworkRuleAttributesRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s DescribeNetworkRuleAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesRequest) SetSourceIp(v string) *DescribeNetworkRuleAttributesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNetworkRuleAttributesRequest) SetNetworkRules(v string) *DescribeNetworkRuleAttributesRequest {
	s.NetworkRules = &v
	return s
}

type DescribeNetworkRuleAttributesResponseBody struct {
	NetworkRuleAttributes []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes `json:"NetworkRuleAttributes,omitempty" xml:"NetworkRuleAttributes,omitempty" type:"Repeated"`
	RequestId             *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBody) SetNetworkRuleAttributes(v []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) *DescribeNetworkRuleAttributesResponseBody {
	s.NetworkRuleAttributes = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBody) SetRequestId(v string) *DescribeNetworkRuleAttributesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes struct {
	FrontendPort *int32                                                                `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Protocol     *string                                                               `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InstanceId   *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config       *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetFrontendPort(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.FrontendPort = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetProtocol(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetInstanceId(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetConfig(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.Config = v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig struct {
	Cc                 *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc         `json:"Cc,omitempty" xml:"Cc,omitempty" type:"Struct"`
	PayloadLen         *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty" type:"Struct"`
	PersistenceTimeout *int32                                                                          `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	Sla                *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla        `json:"Sla,omitempty" xml:"Sla,omitempty" type:"Struct"`
	Slimit             *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit     `json:"Slimit,omitempty" xml:"Slimit,omitempty" type:"Struct"`
	NodataConn         *string                                                                         `json:"NodataConn,omitempty" xml:"NodataConn,omitempty"`
	Synproxy           *string                                                                         `json:"Synproxy,omitempty" xml:"Synproxy,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetCc(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Cc = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetPayloadLen(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.PayloadLen = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetPersistenceTimeout(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetSla(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Sla = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetSlimit(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Slimit = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetNodataConn(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.NodataConn = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetSynproxy(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Synproxy = &v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc struct {
	Sblack []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack `json:"Sblack,omitempty" xml:"Sblack,omitempty" type:"Repeated"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) SetSblack(v []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc {
	s.Sblack = v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack struct {
	Type    *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	During  *int32 `json:"During,omitempty" xml:"During,omitempty"`
	Cnt     *int32 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetType(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.Type = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetExpires(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.Expires = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetDuring(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.During = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetCnt(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.Cnt = &v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen struct {
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) SetMax(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen {
	s.Max = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) SetMin(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen {
	s.Min = &v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla struct {
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	CpsEnable     *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	Cps           *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Maxconn       *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetMaxconnEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetCpsEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetCps(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetMaxconn(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.Maxconn = &v
	return s
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit struct {
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	CpsEnable     *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	Cps           *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Pps           *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	Bps           *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Maxconn       *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	CpsMode       *int32 `json:"CpsMode,omitempty" xml:"CpsMode,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetMaxconnEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetCpsEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetCps(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetPps(v int64) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Pps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetBps(v int64) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Bps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetMaxconn(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Maxconn = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetCpsMode(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.CpsMode = &v
	return s
}

type DescribeNetworkRuleAttributesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkRuleAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkRuleAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponse) SetHeaders(v map[string]*string) *DescribeNetworkRuleAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponse) SetBody(v *DescribeNetworkRuleAttributesResponseBody) *DescribeNetworkRuleAttributesResponse {
	s.Body = v
	return s
}

type DescribeNetworkRulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int32  `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsOffset        *bool   `json:"IsOffset,omitempty" xml:"IsOffset,omitempty"`
}

func (s DescribeNetworkRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesRequest) SetSourceIp(v string) *DescribeNetworkRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetInstanceId(v string) *DescribeNetworkRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetForwardProtocol(v string) *DescribeNetworkRulesRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetFrontendPort(v int32) *DescribeNetworkRulesRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetPageNumber(v int32) *DescribeNetworkRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetPageSize(v int32) *DescribeNetworkRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkRulesRequest) SetIsOffset(v bool) *DescribeNetworkRulesRequest {
	s.IsOffset = &v
	return s
}

type DescribeNetworkRulesResponseBody struct {
	TotalCount   *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NetworkRules []*DescribeNetworkRulesResponseBodyNetworkRules `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesResponseBody) SetTotalCount(v int64) *DescribeNetworkRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkRulesResponseBody) SetNetworkRules(v []*DescribeNetworkRulesResponseBodyNetworkRules) *DescribeNetworkRulesResponseBody {
	s.NetworkRules = v
	return s
}

func (s *DescribeNetworkRulesResponseBody) SetRequestId(v string) *DescribeNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNetworkRulesResponseBodyNetworkRules struct {
	FrontendPort *int32    `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	IsAutoCreate *bool     `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	Protocol     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServers  []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BackendPort  *int32    `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
}

func (s DescribeNetworkRulesResponseBodyNetworkRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRulesResponseBodyNetworkRules) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetFrontendPort(v int32) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.FrontendPort = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetIsAutoCreate(v bool) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetProtocol(v string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetRealServers(v []*string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.RealServers = v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetInstanceId(v string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetBackendPort(v int32) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.BackendPort = &v
	return s
}

type DescribeNetworkRulesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesResponse) SetHeaders(v map[string]*string) *DescribeNetworkRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRulesResponse) SetBody(v *DescribeNetworkRulesResponseBody) *DescribeNetworkRulesResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EntityType      *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityObject    *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityType(v int32) *DescribeOpEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityObject(v string) *DescribeOpEntitiesRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageNumber(v int32) *DescribeOpEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int64) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	EntityType   *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OpAction     *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribePortAttackMaxFlowRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortAttackMaxFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAttackMaxFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribePortAttackMaxFlowRequest) SetSourceIp(v string) *DescribePortAttackMaxFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetResourceGroupId(v string) *DescribePortAttackMaxFlowRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetEndTime(v int64) *DescribePortAttackMaxFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetStartTime(v int64) *DescribePortAttackMaxFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortAttackMaxFlowRequest) SetInstanceIds(v []*string) *DescribePortAttackMaxFlowRequest {
	s.InstanceIds = v
	return s
}

type DescribePortAttackMaxFlowResponseBody struct {
	Pps       *int64  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Bps       *int64  `json:"Bps,omitempty" xml:"Bps,omitempty"`
}

func (s DescribePortAttackMaxFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAttackMaxFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortAttackMaxFlowResponseBody) SetPps(v int64) *DescribePortAttackMaxFlowResponseBody {
	s.Pps = &v
	return s
}

func (s *DescribePortAttackMaxFlowResponseBody) SetRequestId(v string) *DescribePortAttackMaxFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortAttackMaxFlowResponseBody) SetBps(v int64) *DescribePortAttackMaxFlowResponseBody {
	s.Bps = &v
	return s
}

type DescribePortAttackMaxFlowResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortAttackMaxFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortAttackMaxFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAttackMaxFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribePortAttackMaxFlowResponse) SetHeaders(v map[string]*string) *DescribePortAttackMaxFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribePortAttackMaxFlowResponse) SetBody(v *DescribePortAttackMaxFlowResponseBody) *DescribePortAttackMaxFlowResponse {
	s.Body = v
	return s
}

type DescribePortAutoCcStatusRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortAutoCcStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAutoCcStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusRequest) SetSourceIp(v string) *DescribePortAutoCcStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortAutoCcStatusRequest) SetInstanceIds(v []*string) *DescribePortAutoCcStatusRequest {
	s.InstanceIds = v
	return s
}

type DescribePortAutoCcStatusResponseBody struct {
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PortAutoCcStatus []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus `json:"PortAutoCcStatus,omitempty" xml:"PortAutoCcStatus,omitempty" type:"Repeated"`
}

func (s DescribePortAutoCcStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAutoCcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusResponseBody) SetRequestId(v string) *DescribePortAutoCcStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBody) SetPortAutoCcStatus(v []*DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) *DescribePortAutoCcStatusResponseBody {
	s.PortAutoCcStatus = v
	return s
}

type DescribePortAutoCcStatusResponseBodyPortAutoCcStatus struct {
	Switch    *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
	Mode      *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	WebSwitch *string `json:"WebSwitch,omitempty" xml:"WebSwitch,omitempty"`
	WebMode   *string `json:"WebMode,omitempty" xml:"WebMode,omitempty"`
}

func (s DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetSwitch(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.Switch = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetMode(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.Mode = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetWebSwitch(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.WebSwitch = &v
	return s
}

func (s *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus) SetWebMode(v string) *DescribePortAutoCcStatusResponseBodyPortAutoCcStatus {
	s.WebMode = &v
	return s
}

type DescribePortAutoCcStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortAutoCcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortAutoCcStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortAutoCcStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePortAutoCcStatusResponse) SetHeaders(v map[string]*string) *DescribePortAutoCcStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePortAutoCcStatusResponse) SetBody(v *DescribePortAutoCcStatusResponseBody) *DescribePortAutoCcStatusResponse {
	s.Body = v
	return s
}

type DescribePortConnsCountRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Port            *string   `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortConnsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsCountRequest) GoString() string {
	return s.String()
}

func (s *DescribePortConnsCountRequest) SetSourceIp(v string) *DescribePortConnsCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetResourceGroupId(v string) *DescribePortConnsCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetEndTime(v int64) *DescribePortConnsCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetStartTime(v int64) *DescribePortConnsCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetPort(v string) *DescribePortConnsCountRequest {
	s.Port = &v
	return s
}

func (s *DescribePortConnsCountRequest) SetInstanceIds(v []*string) *DescribePortConnsCountRequest {
	s.InstanceIds = v
	return s
}

type DescribePortConnsCountResponseBody struct {
	Conns      *int64  `json:"Conns,omitempty" xml:"Conns,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Cps        *int64  `json:"Cps,omitempty" xml:"Cps,omitempty"`
	InActConns *int64  `json:"InActConns,omitempty" xml:"InActConns,omitempty"`
	ActConns   *int64  `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
}

func (s DescribePortConnsCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortConnsCountResponseBody) SetConns(v int64) *DescribePortConnsCountResponseBody {
	s.Conns = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetRequestId(v string) *DescribePortConnsCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetCps(v int64) *DescribePortConnsCountResponseBody {
	s.Cps = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetInActConns(v int64) *DescribePortConnsCountResponseBody {
	s.InActConns = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetActConns(v int64) *DescribePortConnsCountResponseBody {
	s.ActConns = &v
	return s
}

type DescribePortConnsCountResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortConnsCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortConnsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsCountResponse) GoString() string {
	return s.String()
}

func (s *DescribePortConnsCountResponse) SetHeaders(v map[string]*string) *DescribePortConnsCountResponse {
	s.Headers = v
	return s
}

func (s *DescribePortConnsCountResponse) SetBody(v *DescribePortConnsCountResponseBody) *DescribePortConnsCountResponse {
	s.Body = v
	return s
}

type DescribePortConnsListRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval        *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Port            *string   `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortConnsListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsListRequest) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListRequest) SetSourceIp(v string) *DescribePortConnsListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortConnsListRequest) SetResourceGroupId(v string) *DescribePortConnsListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortConnsListRequest) SetEndTime(v int64) *DescribePortConnsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortConnsListRequest) SetStartTime(v int64) *DescribePortConnsListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortConnsListRequest) SetInterval(v int32) *DescribePortConnsListRequest {
	s.Interval = &v
	return s
}

func (s *DescribePortConnsListRequest) SetPort(v string) *DescribePortConnsListRequest {
	s.Port = &v
	return s
}

func (s *DescribePortConnsListRequest) SetInstanceIds(v []*string) *DescribePortConnsListRequest {
	s.InstanceIds = v
	return s
}

type DescribePortConnsListResponseBody struct {
	ConnsList []*DescribePortConnsListResponseBodyConnsList `json:"ConnsList,omitempty" xml:"ConnsList,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortConnsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListResponseBody) SetConnsList(v []*DescribePortConnsListResponseBodyConnsList) *DescribePortConnsListResponseBody {
	s.ConnsList = v
	return s
}

func (s *DescribePortConnsListResponseBody) SetRequestId(v string) *DescribePortConnsListResponseBody {
	s.RequestId = &v
	return s
}

type DescribePortConnsListResponseBodyConnsList struct {
	ActConns   *int64 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	InActConns *int64 `json:"InActConns,omitempty" xml:"InActConns,omitempty"`
	Index      *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	Time       *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	Cps        *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Conns      *int64 `json:"Conns,omitempty" xml:"Conns,omitempty"`
}

func (s DescribePortConnsListResponseBodyConnsList) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsListResponseBodyConnsList) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListResponseBodyConnsList) SetActConns(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.ActConns = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetInActConns(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.InActConns = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetIndex(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Index = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetTime(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Time = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetCps(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Cps = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetConns(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Conns = &v
	return s
}

type DescribePortConnsListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortConnsListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortConnsListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortConnsListResponse) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListResponse) SetHeaders(v map[string]*string) *DescribePortConnsListResponse {
	s.Headers = v
	return s
}

func (s *DescribePortConnsListResponse) SetBody(v *DescribePortConnsListResponseBody) *DescribePortConnsListResponse {
	s.Body = v
	return s
}

type DescribePortFlowListRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval        *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortFlowListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortFlowListRequest) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListRequest) SetSourceIp(v string) *DescribePortFlowListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortFlowListRequest) SetResourceGroupId(v string) *DescribePortFlowListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortFlowListRequest) SetEndTime(v int64) *DescribePortFlowListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortFlowListRequest) SetStartTime(v int64) *DescribePortFlowListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortFlowListRequest) SetInterval(v int32) *DescribePortFlowListRequest {
	s.Interval = &v
	return s
}

func (s *DescribePortFlowListRequest) SetInstanceIds(v []*string) *DescribePortFlowListRequest {
	s.InstanceIds = v
	return s
}

type DescribePortFlowListResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PortFlowList []*DescribePortFlowListResponseBodyPortFlowList `json:"PortFlowList,omitempty" xml:"PortFlowList,omitempty" type:"Repeated"`
}

func (s DescribePortFlowListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortFlowListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListResponseBody) SetRequestId(v string) *DescribePortFlowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortFlowListResponseBody) SetPortFlowList(v []*DescribePortFlowListResponseBodyPortFlowList) *DescribePortFlowListResponseBody {
	s.PortFlowList = v
	return s
}

type DescribePortFlowListResponseBodyPortFlowList struct {
	Index     *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	Time      *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	InPps     *int64  `json:"InPps,omitempty" xml:"InPps,omitempty"`
	InBps     *int64  `json:"InBps,omitempty" xml:"InBps,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	OutPps    *int64  `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	AttackPps *int64  `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	OutBps    *int64  `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	AttackBps *int64  `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
}

func (s DescribePortFlowListResponseBodyPortFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribePortFlowListResponseBodyPortFlowList) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetIndex(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.Index = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetTime(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.Time = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetInPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.InPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetInBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.InBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetRegion(v string) *DescribePortFlowListResponseBodyPortFlowList {
	s.Region = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetOutPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.OutPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetAttackPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetOutBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.OutBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetAttackBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.AttackBps = &v
	return s
}

type DescribePortFlowListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortFlowListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortFlowListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortFlowListResponse) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListResponse) SetHeaders(v map[string]*string) *DescribePortFlowListResponse {
	s.Headers = v
	return s
}

func (s *DescribePortFlowListResponse) SetBody(v *DescribePortFlowListResponseBody) *DescribePortFlowListResponse {
	s.Body = v
	return s
}

type DescribePortMaxConnsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortMaxConnsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortMaxConnsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsRequest) SetSourceIp(v string) *DescribePortMaxConnsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortMaxConnsRequest) SetResourceGroupId(v string) *DescribePortMaxConnsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortMaxConnsRequest) SetEndTime(v int64) *DescribePortMaxConnsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortMaxConnsRequest) SetStartTime(v int64) *DescribePortMaxConnsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortMaxConnsRequest) SetInstanceIds(v []*string) *DescribePortMaxConnsRequest {
	s.InstanceIds = v
	return s
}

type DescribePortMaxConnsResponseBody struct {
	PortMaxConns []*DescribePortMaxConnsResponseBodyPortMaxConns `json:"PortMaxConns,omitempty" xml:"PortMaxConns,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortMaxConnsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortMaxConnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsResponseBody) SetPortMaxConns(v []*DescribePortMaxConnsResponseBodyPortMaxConns) *DescribePortMaxConnsResponseBody {
	s.PortMaxConns = v
	return s
}

func (s *DescribePortMaxConnsResponseBody) SetRequestId(v string) *DescribePortMaxConnsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePortMaxConnsResponseBodyPortMaxConns struct {
	Cps  *int64  `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribePortMaxConnsResponseBodyPortMaxConns) String() string {
	return tea.Prettify(s)
}

func (s DescribePortMaxConnsResponseBodyPortMaxConns) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) SetCps(v int64) *DescribePortMaxConnsResponseBodyPortMaxConns {
	s.Cps = &v
	return s
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) SetIp(v string) *DescribePortMaxConnsResponseBodyPortMaxConns {
	s.Ip = &v
	return s
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) SetPort(v string) *DescribePortMaxConnsResponseBodyPortMaxConns {
	s.Port = &v
	return s
}

type DescribePortMaxConnsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortMaxConnsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortMaxConnsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortMaxConnsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsResponse) SetHeaders(v map[string]*string) *DescribePortMaxConnsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortMaxConnsResponse) SetBody(v *DescribePortMaxConnsResponseBody) *DescribePortMaxConnsResponse {
	s.Body = v
	return s
}

type DescribePortViewSourceCountriesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceCountriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceCountriesRequest) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesRequest) SetSourceIp(v string) *DescribePortViewSourceCountriesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetResourceGroupId(v string) *DescribePortViewSourceCountriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetEndTime(v int64) *DescribePortViewSourceCountriesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetStartTime(v int64) *DescribePortViewSourceCountriesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortViewSourceCountriesRequest) SetInstanceIds(v []*string) *DescribePortViewSourceCountriesRequest {
	s.InstanceIds = v
	return s
}

type DescribePortViewSourceCountriesResponseBody struct {
	RequestId      *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceCountrys []*DescribePortViewSourceCountriesResponseBodySourceCountrys `json:"SourceCountrys,omitempty" xml:"SourceCountrys,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceCountriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesResponseBody) SetRequestId(v string) *DescribePortViewSourceCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortViewSourceCountriesResponseBody) SetSourceCountrys(v []*DescribePortViewSourceCountriesResponseBodySourceCountrys) *DescribePortViewSourceCountriesResponseBody {
	s.SourceCountrys = v
	return s
}

type DescribePortViewSourceCountriesResponseBodySourceCountrys struct {
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePortViewSourceCountriesResponseBodySourceCountrys) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceCountriesResponseBodySourceCountrys) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) SetCountryId(v string) *DescribePortViewSourceCountriesResponseBodySourceCountrys {
	s.CountryId = &v
	return s
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) SetCount(v int64) *DescribePortViewSourceCountriesResponseBodySourceCountrys {
	s.Count = &v
	return s
}

type DescribePortViewSourceCountriesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortViewSourceCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortViewSourceCountriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceCountriesResponse) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesResponse) SetHeaders(v map[string]*string) *DescribePortViewSourceCountriesResponse {
	s.Headers = v
	return s
}

func (s *DescribePortViewSourceCountriesResponse) SetBody(v *DescribePortViewSourceCountriesResponseBody) *DescribePortViewSourceCountriesResponse {
	s.Body = v
	return s
}

type DescribePortViewSourceIspsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceIspsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceIspsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsRequest) SetSourceIp(v string) *DescribePortViewSourceIspsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetResourceGroupId(v string) *DescribePortViewSourceIspsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetEndTime(v int64) *DescribePortViewSourceIspsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetStartTime(v int64) *DescribePortViewSourceIspsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortViewSourceIspsRequest) SetInstanceIds(v []*string) *DescribePortViewSourceIspsRequest {
	s.InstanceIds = v
	return s
}

type DescribePortViewSourceIspsResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Isps      []*DescribePortViewSourceIspsResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceIspsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceIspsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsResponseBody) SetRequestId(v string) *DescribePortViewSourceIspsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortViewSourceIspsResponseBody) SetIsps(v []*DescribePortViewSourceIspsResponseBodyIsps) *DescribePortViewSourceIspsResponseBody {
	s.Isps = v
	return s
}

type DescribePortViewSourceIspsResponseBodyIsps struct {
	IspId *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePortViewSourceIspsResponseBodyIsps) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceIspsResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) SetIspId(v string) *DescribePortViewSourceIspsResponseBodyIsps {
	s.IspId = &v
	return s
}

func (s *DescribePortViewSourceIspsResponseBodyIsps) SetCount(v int64) *DescribePortViewSourceIspsResponseBodyIsps {
	s.Count = &v
	return s
}

type DescribePortViewSourceIspsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortViewSourceIspsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortViewSourceIspsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceIspsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceIspsResponse) SetHeaders(v map[string]*string) *DescribePortViewSourceIspsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortViewSourceIspsResponse) SetBody(v *DescribePortViewSourceIspsResponseBody) *DescribePortViewSourceIspsResponse {
	s.Body = v
	return s
}

type DescribePortViewSourceProvincesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EndTime         *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceProvincesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceProvincesRequest) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesRequest) SetSourceIp(v string) *DescribePortViewSourceProvincesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetResourceGroupId(v string) *DescribePortViewSourceProvincesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetEndTime(v int64) *DescribePortViewSourceProvincesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetStartTime(v int64) *DescribePortViewSourceProvincesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePortViewSourceProvincesRequest) SetInstanceIds(v []*string) *DescribePortViewSourceProvincesRequest {
	s.InstanceIds = v
	return s
}

type DescribePortViewSourceProvincesResponseBody struct {
	SourceProvinces []*DescribePortViewSourceProvincesResponseBodySourceProvinces `json:"SourceProvinces,omitempty" xml:"SourceProvinces,omitempty" type:"Repeated"`
	RequestId       *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortViewSourceProvincesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceProvincesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesResponseBody) SetSourceProvinces(v []*DescribePortViewSourceProvincesResponseBodySourceProvinces) *DescribePortViewSourceProvincesResponseBody {
	s.SourceProvinces = v
	return s
}

func (s *DescribePortViewSourceProvincesResponseBody) SetRequestId(v string) *DescribePortViewSourceProvincesResponseBody {
	s.RequestId = &v
	return s
}

type DescribePortViewSourceProvincesResponseBodySourceProvinces struct {
	ProvinceId *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePortViewSourceProvincesResponseBodySourceProvinces) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceProvincesResponseBodySourceProvinces) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) SetProvinceId(v string) *DescribePortViewSourceProvincesResponseBodySourceProvinces {
	s.ProvinceId = &v
	return s
}

func (s *DescribePortViewSourceProvincesResponseBodySourceProvinces) SetCount(v int64) *DescribePortViewSourceProvincesResponseBodySourceProvinces {
	s.Count = &v
	return s
}

type DescribePortViewSourceProvincesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortViewSourceProvincesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortViewSourceProvincesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortViewSourceProvincesResponse) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesResponse) SetHeaders(v map[string]*string) *DescribePortViewSourceProvincesResponse {
	s.Headers = v
	return s
}

func (s *DescribePortViewSourceProvincesResponse) SetBody(v *DescribePortViewSourceProvincesResponseBody) *DescribePortViewSourceProvincesResponse {
	s.Body = v
	return s
}

type DescribeSceneDefenseObjectsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId        *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSceneDefenseObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefenseObjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsRequest) SetSourceIp(v string) *DescribeSceneDefenseObjectsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSceneDefenseObjectsRequest) SetPolicyId(v string) *DescribeSceneDefenseObjectsRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsRequest) SetResourceGroupId(v string) *DescribeSceneDefenseObjectsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSceneDefenseObjectsResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Objects   []*DescribeSceneDefenseObjectsResponseBodyObjects `json:"Objects,omitempty" xml:"Objects,omitempty" type:"Repeated"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSceneDefenseObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefenseObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsResponseBody) SetRequestId(v string) *DescribeSceneDefenseObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBody) SetObjects(v []*DescribeSceneDefenseObjectsResponseBodyObjects) *DescribeSceneDefenseObjectsResponseBody {
	s.Objects = v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBody) SetSuccess(v bool) *DescribeSceneDefenseObjectsResponseBody {
	s.Success = &v
	return s
}

type DescribeSceneDefenseObjectsResponseBodyObjects struct {
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Vip      *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
}

func (s DescribeSceneDefenseObjectsResponseBodyObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefenseObjectsResponseBodyObjects) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) SetDomain(v string) *DescribeSceneDefenseObjectsResponseBodyObjects {
	s.Domain = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) SetPolicyId(v string) *DescribeSceneDefenseObjectsResponseBodyObjects {
	s.PolicyId = &v
	return s
}

func (s *DescribeSceneDefenseObjectsResponseBodyObjects) SetVip(v string) *DescribeSceneDefenseObjectsResponseBodyObjects {
	s.Vip = &v
	return s
}

type DescribeSceneDefenseObjectsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSceneDefenseObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSceneDefenseObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefenseObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefenseObjectsResponse) SetHeaders(v map[string]*string) *DescribeSceneDefenseObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneDefenseObjectsResponse) SetBody(v *DescribeSceneDefenseObjectsResponseBody) *DescribeSceneDefenseObjectsResponse {
	s.Body = v
	return s
}

type DescribeSceneDefensePoliciesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Template        *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSceneDefensePoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefensePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesRequest) SetSourceIp(v string) *DescribeSceneDefensePoliciesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSceneDefensePoliciesRequest) SetTemplate(v string) *DescribeSceneDefensePoliciesRequest {
	s.Template = &v
	return s
}

func (s *DescribeSceneDefensePoliciesRequest) SetStatus(v string) *DescribeSceneDefensePoliciesRequest {
	s.Status = &v
	return s
}

func (s *DescribeSceneDefensePoliciesRequest) SetResourceGroupId(v string) *DescribeSceneDefensePoliciesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSceneDefensePoliciesResponseBody struct {
	Policies  []*DescribeSceneDefensePoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSceneDefensePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponseBody) SetPolicies(v []*DescribeSceneDefensePoliciesResponseBodyPolicies) *DescribeSceneDefensePoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBody) SetRequestId(v string) *DescribeSceneDefensePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBody) SetSuccess(v bool) *DescribeSceneDefensePoliciesResponseBody {
	s.Success = &v
	return s
}

type DescribeSceneDefensePoliciesResponseBodyPolicies struct {
	Done            *int32                                                             `json:"Done,omitempty" xml:"Done,omitempty"`
	EndTime         *int64                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status          *int32                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime       *int64                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ObjectCount     *int32                                                             `json:"ObjectCount,omitempty" xml:"ObjectCount,omitempty"`
	Template        *string                                                            `json:"Template,omitempty" xml:"Template,omitempty"`
	PolicyId        *string                                                            `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Name            *string                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	RuntimePolicies []*DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies `json:"RuntimePolicies,omitempty" xml:"RuntimePolicies,omitempty" type:"Repeated"`
}

func (s DescribeSceneDefensePoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetDone(v int32) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Done = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetEndTime(v int64) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.EndTime = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetStatus(v int32) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Status = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetStartTime(v int64) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.StartTime = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetObjectCount(v int32) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.ObjectCount = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetTemplate(v string) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Template = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetPolicyId(v string) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetName(v string) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPolicies) SetRuntimePolicies(v []*DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) *DescribeSceneDefensePoliciesResponseBodyPolicies {
	s.RuntimePolicies = v
	return s
}

type DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies struct {
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	OldValue   *string `json:"oldValue,omitempty" xml:"oldValue,omitempty"`
	NewValue   *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	PolicyType *int32  `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetStatus(v int32) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.Status = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetOldValue(v string) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.OldValue = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetNewValue(v string) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.NewValue = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies) SetPolicyType(v int32) *DescribeSceneDefensePoliciesResponseBodyPoliciesRuntimePolicies {
	s.PolicyType = &v
	return s
}

type DescribeSceneDefensePoliciesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSceneDefensePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSceneDefensePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponse) SetHeaders(v map[string]*string) *DescribeSceneDefensePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneDefensePoliciesResponse) SetBody(v *DescribeSceneDefensePoliciesResponseBody) *DescribeSceneDefensePoliciesResponse {
	s.Body = v
	return s
}

type DescribeSchedulerRulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Offset          *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsOffset        *bool   `json:"IsOffset,omitempty" xml:"IsOffset,omitempty"`
}

func (s DescribeSchedulerRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesRequest) SetSourceIp(v string) *DescribeSchedulerRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetResourceGroupId(v string) *DescribeSchedulerRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetRuleName(v string) *DescribeSchedulerRulesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetOffset(v int32) *DescribeSchedulerRulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetPageNumber(v int32) *DescribeSchedulerRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetPageSize(v int32) *DescribeSchedulerRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSchedulerRulesRequest) SetIsOffset(v bool) *DescribeSchedulerRulesRequest {
	s.IsOffset = &v
	return s
}

type DescribeSchedulerRulesResponseBody struct {
	TotalCount     *string                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchedulerRules []*DescribeSchedulerRulesResponseBodySchedulerRules `json:"SchedulerRules,omitempty" xml:"SchedulerRules,omitempty" type:"Repeated"`
}

func (s DescribeSchedulerRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBody) SetTotalCount(v string) *DescribeSchedulerRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBody) SetRequestId(v string) *DescribeSchedulerRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBody) SetSchedulerRules(v []*DescribeSchedulerRulesResponseBodySchedulerRules) *DescribeSchedulerRulesResponseBody {
	s.SchedulerRules = v
	return s
}

type DescribeSchedulerRulesResponseBodySchedulerRules struct {
	RuleType *string                                                  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Param    *DescribeSchedulerRulesResponseBodySchedulerRulesParam   `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	Cname    *string                                                  `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Rules    []*DescribeSchedulerRulesResponseBodySchedulerRulesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	RuleName *string                                                  `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRules) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetRuleType(v string) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.RuleType = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetParam(v *DescribeSchedulerRulesResponseBodySchedulerRulesParam) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.Param = v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetCname(v string) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.Cname = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetRules(v []*DescribeSchedulerRulesResponseBodySchedulerRulesRules) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.Rules = v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRules) SetRuleName(v string) *DescribeSchedulerRulesResponseBodySchedulerRules {
	s.RuleName = &v
	return s
}

type DescribeSchedulerRulesResponseBodySchedulerRulesParam struct {
	ParamData *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData `json:"ParamData,omitempty" xml:"ParamData,omitempty" type:"Struct"`
	ParamType *string                                                         `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParam) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) SetParamData(v *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) *DescribeSchedulerRulesResponseBodySchedulerRulesParam {
	s.ParamData = v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParam) SetParamType(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesParam {
	s.ParamType = &v
	return s
}

type DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData struct {
	CloudInstanceId *string `json:"CloudInstanceId,omitempty" xml:"CloudInstanceId,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData) SetCloudInstanceId(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesParamParamData {
	s.CloudInstanceId = &v
	return s
}

type DescribeSchedulerRulesResponseBodySchedulerRulesRules struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueType    *int32  `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RestoreDelay *int32  `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesResponseBodySchedulerRulesRules) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetType(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Type = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetStatus(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Status = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetValue(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Value = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetValueType(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.ValueType = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetPriority(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.Priority = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetRestoreDelay(v int32) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.RestoreDelay = &v
	return s
}

func (s *DescribeSchedulerRulesResponseBodySchedulerRulesRules) SetRegionId(v string) *DescribeSchedulerRulesResponseBodySchedulerRulesRules {
	s.RegionId = &v
	return s
}

type DescribeSchedulerRulesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSchedulerRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSchedulerRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSchedulerRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSchedulerRulesResponse) SetHeaders(v map[string]*string) *DescribeSchedulerRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSchedulerRulesResponse) SetBody(v *DescribeSchedulerRulesResponseBody) *DescribeSchedulerRulesResponse {
	s.Body = v
	return s
}

type DescribeSlsAuthStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) SetSourceIp(v string) *DescribeSlsAuthStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetLang(v string) *DescribeSlsAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsAuthStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsAuthStatus *bool   `json:"SlsAuthStatus,omitempty" xml:"SlsAuthStatus,omitempty"`
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

func (s *DescribeSlsAuthStatusResponseBody) SetSlsAuthStatus(v bool) *DescribeSlsAuthStatusResponseBody {
	s.SlsAuthStatus = &v
	return s
}

type DescribeSlsAuthStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSlsAuthStatusResponse) SetBody(v *DescribeSlsAuthStatusResponseBody) *DescribeSlsAuthStatusResponse {
	s.Body = v
	return s
}

type DescribeSlsLogstoreInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsLogstoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoRequest) SetSourceIp(v string) *DescribeSlsLogstoreInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetLang(v string) *DescribeSlsLogstoreInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsLogstoreInfoResponseBody struct {
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Quota     *int64  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	LogStore  *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Used      *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	Ttl       *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeSlsLogstoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetProject(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetRequestId(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetQuota(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetLogStore(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetUsed(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetTtl(v int32) *DescribeSlsLogstoreInfoResponseBody {
	s.Ttl = &v
	return s
}

type DescribeSlsLogstoreInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsLogstoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlsLogstoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponse) SetHeaders(v map[string]*string) *DescribeSlsLogstoreInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetBody(v *DescribeSlsLogstoreInfoResponseBody) *DescribeSlsLogstoreInfoResponse {
	s.Body = v
	return s
}

type DescribeSlsOpenStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusRequest) SetSourceIp(v string) *DescribeSlsOpenStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetLang(v string) *DescribeSlsOpenStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsOpenStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsOpenStatus *bool   `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s DescribeSlsOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponseBody) SetRequestId(v string) *DescribeSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *DescribeSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

type DescribeSlsOpenStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsOpenStatusResponse) SetBody(v *DescribeSlsOpenStatusResponseBody) *DescribeSlsOpenStatusResponse {
	s.Body = v
	return s
}

type DescribeStsGrantStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeStsGrantStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStsGrantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusRequest) SetSourceIp(v string) *DescribeStsGrantStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStsGrantStatusRequest) SetResourceGroupId(v string) *DescribeStsGrantStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeStsGrantStatusRequest) SetRole(v string) *DescribeStsGrantStatusRequest {
	s.Role = &v
	return s
}

type DescribeStsGrantStatusResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StsGrant  *DescribeStsGrantStatusResponseBodyStsGrant `json:"StsGrant,omitempty" xml:"StsGrant,omitempty" type:"Struct"`
}

func (s DescribeStsGrantStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStsGrantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusResponseBody) SetRequestId(v string) *DescribeStsGrantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStsGrantStatusResponseBody) SetStsGrant(v *DescribeStsGrantStatusResponseBodyStsGrant) *DescribeStsGrantStatusResponseBody {
	s.StsGrant = v
	return s
}

type DescribeStsGrantStatusResponseBodyStsGrant struct {
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeStsGrantStatusResponseBodyStsGrant) String() string {
	return tea.Prettify(s)
}

func (s DescribeStsGrantStatusResponseBodyStsGrant) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusResponseBodyStsGrant) SetStatus(v int32) *DescribeStsGrantStatusResponseBodyStsGrant {
	s.Status = &v
	return s
}

type DescribeStsGrantStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStsGrantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStsGrantStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStsGrantStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusResponse) SetHeaders(v map[string]*string) *DescribeStsGrantStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeStsGrantStatusResponse) SetBody(v *DescribeStsGrantStatusResponseBody) *DescribeStsGrantStatusResponse {
	s.Body = v
	return s
}

type DescribeTagKeysRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysRequest) SetSourceIp(v string) *DescribeTagKeysRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTagKeysRequest) SetRegionId(v string) *DescribeTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceGroupId(v string) *DescribeTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagKeysRequest) SetResourceType(v string) *DescribeTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagKeysRequest) SetPageSize(v int32) *DescribeTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagKeysRequest) SetPageNumber(v int32) *DescribeTagKeysRequest {
	s.PageNumber = &v
	return s
}

type DescribeTagKeysResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TagKeys    []*DescribeTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s DescribeTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBody) SetTotalCount(v int32) *DescribeTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetRequestId(v string) *DescribeTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetPageSize(v int32) *DescribeTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetPageNumber(v int32) *DescribeTagKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetTagKeys(v []*DescribeTagKeysResponseBodyTagKeys) *DescribeTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type DescribeTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBodyTagKeys) SetTagCount(v int32) *DescribeTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *DescribeTagKeysResponseBodyTagKeys) SetTagKey(v string) *DescribeTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type DescribeTagKeysResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponse) SetHeaders(v map[string]*string) *DescribeTagKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagKeysResponse) SetBody(v *DescribeTagKeysResponseBody) *DescribeTagKeysResponse {
	s.Body = v
	return s
}

type DescribeTagResourcesRequest struct {
	SourceIp        *string                            `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string                            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken       *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceIds     []*string                          `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	Tags            []*DescribeTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequest) SetSourceIp(v string) *DescribeTagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetRegionId(v string) *DescribeTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceGroupId(v string) *DescribeTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceType(v string) *DescribeTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetNextToken(v string) *DescribeTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceIds(v []*string) *DescribeTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeTagResourcesRequest) SetTags(v []*DescribeTagResourcesRequestTags) *DescribeTagResourcesRequest {
	s.Tags = v
	return s
}

type DescribeTagResourcesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequestTags) SetKey(v string) *DescribeTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeTagResourcesRequestTags) SetValue(v string) *DescribeTagResourcesRequestTags {
	s.Value = &v
	return s
}

type DescribeTagResourcesResponseBody struct {
	NextToken    *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *DescribeTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s DescribeTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBody) SetNextToken(v string) *DescribeTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTagResourcesResponseBody) SetRequestId(v string) *DescribeTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagResourcesResponseBody) SetTagResources(v *DescribeTagResourcesResponseBodyTagResources) *DescribeTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type DescribeTagResourcesResponseBodyTagResources struct {
	TagResource []*DescribeTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResources) SetTagResource(v []*DescribeTagResourcesResponseBodyTagResourcesTagResource) *DescribeTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type DescribeTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type DescribeTagResourcesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagResourcesResponse) SetBody(v *DescribeTagResourcesResponseBody) *DescribeTagResourcesResponse {
	s.Body = v
	return s
}

type DescribeUnBlackholeCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeUnBlackholeCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnBlackholeCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnBlackholeCountRequest) SetSourceIp(v string) *DescribeUnBlackholeCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUnBlackholeCountRequest) SetResourceGroupId(v string) *DescribeUnBlackholeCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeUnBlackholeCountResponseBody struct {
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RemainCount *int32  `json:"RemainCount,omitempty" xml:"RemainCount,omitempty"`
}

func (s DescribeUnBlackholeCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnBlackholeCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnBlackholeCountResponseBody) SetTotalCount(v int32) *DescribeUnBlackholeCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUnBlackholeCountResponseBody) SetRequestId(v string) *DescribeUnBlackholeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnBlackholeCountResponseBody) SetRemainCount(v int32) *DescribeUnBlackholeCountResponseBody {
	s.RemainCount = &v
	return s
}

type DescribeUnBlackholeCountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUnBlackholeCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUnBlackholeCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnBlackholeCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnBlackholeCountResponse) SetHeaders(v map[string]*string) *DescribeUnBlackholeCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnBlackholeCountResponse) SetBody(v *DescribeUnBlackholeCountResponseBody) *DescribeUnBlackholeCountResponse {
	s.Body = v
	return s
}

type DescribeUnBlockCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeUnBlockCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnBlockCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnBlockCountRequest) SetSourceIp(v string) *DescribeUnBlockCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUnBlockCountRequest) SetLang(v string) *DescribeUnBlockCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUnBlockCountRequest) SetResourceGroupId(v string) *DescribeUnBlockCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeUnBlockCountResponseBody struct {
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RemainCount *int32  `json:"RemainCount,omitempty" xml:"RemainCount,omitempty"`
}

func (s DescribeUnBlockCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnBlockCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnBlockCountResponseBody) SetTotalCount(v int32) *DescribeUnBlockCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUnBlockCountResponseBody) SetRequestId(v string) *DescribeUnBlockCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnBlockCountResponseBody) SetRemainCount(v int32) *DescribeUnBlockCountResponseBody {
	s.RemainCount = &v
	return s
}

type DescribeUnBlockCountResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUnBlockCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUnBlockCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnBlockCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnBlockCountResponse) SetHeaders(v map[string]*string) *DescribeUnBlockCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnBlockCountResponse) SetBody(v *DescribeUnBlockCountResponseBody) *DescribeUnBlockCountResponse {
	s.Body = v
	return s
}

type DescribeWebAccessLogDispatchStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebAccessLogDispatchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetSourceIp(v string) *DescribeWebAccessLogDispatchStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetLang(v string) *DescribeWebAccessLogDispatchStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetResourceGroupId(v string) *DescribeWebAccessLogDispatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetPageNumber(v int32) *DescribeWebAccessLogDispatchStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusRequest) SetPageSize(v int32) *DescribeWebAccessLogDispatchStatusRequest {
	s.PageSize = &v
	return s
}

type DescribeWebAccessLogDispatchStatusResponseBody struct {
	TotalCount      *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsConfigStatus []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" type:"Repeated"`
}

func (s DescribeWebAccessLogDispatchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) SetTotalCount(v int32) *DescribeWebAccessLogDispatchStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) SetRequestId(v string) *DescribeWebAccessLogDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBody) SetSlsConfigStatus(v []*DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) *DescribeWebAccessLogDispatchStatusResponseBody {
	s.SlsConfigStatus = v
	return s
}

type DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) SetDomain(v string) *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus {
	s.Domain = &v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus) SetEnable(v bool) *DescribeWebAccessLogDispatchStatusResponseBodySlsConfigStatus {
	s.Enable = &v
	return s
}

type DescribeWebAccessLogDispatchStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebAccessLogDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebAccessLogDispatchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogDispatchStatusResponse) SetHeaders(v map[string]*string) *DescribeWebAccessLogDispatchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessLogDispatchStatusResponse) SetBody(v *DescribeWebAccessLogDispatchStatusResponseBody) *DescribeWebAccessLogDispatchStatusResponse {
	s.Body = v
	return s
}

type DescribeWebAccessLogEmptyCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebAccessLogEmptyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogEmptyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogEmptyCountRequest) SetSourceIp(v string) *DescribeWebAccessLogEmptyCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountRequest) SetLang(v string) *DescribeWebAccessLogEmptyCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountRequest) SetResourceGroupId(v string) *DescribeWebAccessLogEmptyCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeWebAccessLogEmptyCountResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableCount *int32  `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty"`
}

func (s DescribeWebAccessLogEmptyCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogEmptyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) SetRequestId(v string) *DescribeWebAccessLogEmptyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponseBody) SetAvailableCount(v int32) *DescribeWebAccessLogEmptyCountResponseBody {
	s.AvailableCount = &v
	return s
}

type DescribeWebAccessLogEmptyCountResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebAccessLogEmptyCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebAccessLogEmptyCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogEmptyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogEmptyCountResponse) SetHeaders(v map[string]*string) *DescribeWebAccessLogEmptyCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessLogEmptyCountResponse) SetBody(v *DescribeWebAccessLogEmptyCountResponseBody) *DescribeWebAccessLogEmptyCountResponse {
	s.Body = v
	return s
}

type DescribeWebAccessLogStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeWebAccessLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogStatusRequest) SetSourceIp(v string) *DescribeWebAccessLogStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebAccessLogStatusRequest) SetLang(v string) *DescribeWebAccessLogStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebAccessLogStatusRequest) SetResourceGroupId(v string) *DescribeWebAccessLogStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAccessLogStatusRequest) SetDomain(v string) *DescribeWebAccessLogStatusRequest {
	s.Domain = &v
	return s
}

type DescribeWebAccessLogStatusResponseBody struct {
	SlsProject  *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	SlsStatus   *bool   `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
}

func (s DescribeWebAccessLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogStatusResponseBody) SetSlsProject(v string) *DescribeWebAccessLogStatusResponseBody {
	s.SlsProject = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) SetSlsStatus(v bool) *DescribeWebAccessLogStatusResponseBody {
	s.SlsStatus = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) SetRequestId(v string) *DescribeWebAccessLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) SetSlsLogstore(v string) *DescribeWebAccessLogStatusResponseBody {
	s.SlsLogstore = &v
	return s
}

type DescribeWebAccessLogStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebAccessLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebAccessLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogStatusResponse) SetHeaders(v map[string]*string) *DescribeWebAccessLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessLogStatusResponse) SetBody(v *DescribeWebAccessLogStatusResponseBody) *DescribeWebAccessLogStatusResponse {
	s.Body = v
	return s
}

type DescribeWebAccessModeRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Domains  []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeRequest) SetSourceIp(v string) *DescribeWebAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebAccessModeRequest) SetDomains(v []*string) *DescribeWebAccessModeRequest {
	s.Domains = v
	return s
}

type DescribeWebAccessModeResponseBody struct {
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainModes []*DescribeWebAccessModeResponseBodyDomainModes `json:"DomainModes,omitempty" xml:"DomainModes,omitempty" type:"Repeated"`
}

func (s DescribeWebAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeResponseBody) SetRequestId(v string) *DescribeWebAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessModeResponseBody) SetDomainModes(v []*DescribeWebAccessModeResponseBodyDomainModes) *DescribeWebAccessModeResponseBody {
	s.DomainModes = v
	return s
}

type DescribeWebAccessModeResponseBodyDomainModes struct {
	AccessMode *int32  `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeWebAccessModeResponseBodyDomainModes) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessModeResponseBodyDomainModes) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) SetAccessMode(v int32) *DescribeWebAccessModeResponseBodyDomainModes {
	s.AccessMode = &v
	return s
}

func (s *DescribeWebAccessModeResponseBodyDomainModes) SetDomain(v string) *DescribeWebAccessModeResponseBodyDomainModes {
	s.Domain = &v
	return s
}

type DescribeWebAccessModeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAccessModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessModeResponse) SetHeaders(v map[string]*string) *DescribeWebAccessModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessModeResponse) SetBody(v *DescribeWebAccessModeResponseBody) *DescribeWebAccessModeResponse {
	s.Body = v
	return s
}

type DescribeWebAreaBlockConfigsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domains         []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebAreaBlockConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsRequest) SetSourceIp(v string) *DescribeWebAreaBlockConfigsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsRequest) SetResourceGroupId(v string) *DescribeWebAreaBlockConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsRequest) SetDomains(v []*string) *DescribeWebAreaBlockConfigsRequest {
	s.Domains = v
	return s
}

type DescribeWebAreaBlockConfigsResponseBody struct {
	AreaBlockConfigs []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs `json:"AreaBlockConfigs,omitempty" xml:"AreaBlockConfigs,omitempty" type:"Repeated"`
	RequestId        *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebAreaBlockConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponseBody) SetAreaBlockConfigs(v []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) *DescribeWebAreaBlockConfigsResponseBody {
	s.AreaBlockConfigs = v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBody) SetRequestId(v string) *DescribeWebAreaBlockConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs struct {
	Domain     *string                                                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RegionList []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) SetDomain(v string) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs {
	s.Domain = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs) SetRegionList(v []*DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigs {
	s.RegionList = v
	return s
}

type DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList struct {
	Block  *int32  `json:"Block,omitempty" xml:"Block,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) SetBlock(v int32) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList {
	s.Block = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList) SetRegion(v string) *DescribeWebAreaBlockConfigsResponseBodyAreaBlockConfigsRegionList {
	s.Region = &v
	return s
}

type DescribeWebAreaBlockConfigsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebAreaBlockConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebAreaBlockConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponse) SetHeaders(v map[string]*string) *DescribeWebAreaBlockConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponse) SetBody(v *DescribeWebAreaBlockConfigsResponseBody) *DescribeWebAreaBlockConfigsResponse {
	s.Body = v
	return s
}

type DescribeWebCacheConfigsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domains         []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebCacheConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCacheConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsRequest) SetSourceIp(v string) *DescribeWebCacheConfigsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebCacheConfigsRequest) SetResourceGroupId(v string) *DescribeWebCacheConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCacheConfigsRequest) SetDomains(v []*string) *DescribeWebCacheConfigsRequest {
	s.Domains = v
	return s
}

type DescribeWebCacheConfigsResponseBody struct {
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainCacheConfigs []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs `json:"DomainCacheConfigs,omitempty" xml:"DomainCacheConfigs,omitempty" type:"Repeated"`
}

func (s DescribeWebCacheConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCacheConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponseBody) SetRequestId(v string) *DescribeWebCacheConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBody) SetDomainCacheConfigs(v []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) *DescribeWebCacheConfigsResponseBody {
	s.DomainCacheConfigs = v
	return s
}

type DescribeWebCacheConfigsResponseBodyDomainCacheConfigs struct {
	Domain      *string                                                             `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CustomRules []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules `json:"CustomRules,omitempty" xml:"CustomRules,omitempty" type:"Repeated"`
	Mode        *string                                                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Enable      *int32                                                              `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetDomain(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.Domain = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetCustomRules(v []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.CustomRules = v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetMode(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.Mode = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetEnable(v int32) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.Enable = &v
	return s
}

type DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules struct {
	CacheTtl *int64  `json:"CacheTtl,omitempty" xml:"CacheTtl,omitempty"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetCacheTtl(v int64) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.CacheTtl = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetMode(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.Mode = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetName(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.Name = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetUri(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.Uri = &v
	return s
}

type DescribeWebCacheConfigsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebCacheConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebCacheConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCacheConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponse) SetHeaders(v map[string]*string) *DescribeWebCacheConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCacheConfigsResponse) SetBody(v *DescribeWebCacheConfigsResponseBody) *DescribeWebCacheConfigsResponse {
	s.Body = v
	return s
}

type DescribeWebCcProtectSwitchRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domains         []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebCcProtectSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCcProtectSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchRequest) SetSourceIp(v string) *DescribeWebCcProtectSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebCcProtectSwitchRequest) SetResourceGroupId(v string) *DescribeWebCcProtectSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCcProtectSwitchRequest) SetDomains(v []*string) *DescribeWebCcProtectSwitchRequest {
	s.Domains = v
	return s
}

type DescribeWebCcProtectSwitchResponseBody struct {
	RequestId         *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProtectSwitchList []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList `json:"ProtectSwitchList,omitempty" xml:"ProtectSwitchList,omitempty" type:"Repeated"`
}

func (s DescribeWebCcProtectSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCcProtectSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchResponseBody) SetRequestId(v string) *DescribeWebCcProtectSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBody) SetProtectSwitchList(v []*DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) *DescribeWebCcProtectSwitchResponseBody {
	s.ProtectSwitchList = v
	return s
}

type DescribeWebCcProtectSwitchResponseBodyProtectSwitchList struct {
	BlackWhiteListEnable *int32  `json:"BlackWhiteListEnable,omitempty" xml:"BlackWhiteListEnable,omitempty"`
	AiTemplate           *string `json:"AiTemplate,omitempty" xml:"AiTemplate,omitempty"`
	PreciseRuleEnable    *int32  `json:"PreciseRuleEnable,omitempty" xml:"PreciseRuleEnable,omitempty"`
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	AiMode               *string `json:"AiMode,omitempty" xml:"AiMode,omitempty"`
	AiRuleEnable         *int32  `json:"AiRuleEnable,omitempty" xml:"AiRuleEnable,omitempty"`
	RegionBlockEnable    *int32  `json:"RegionBlockEnable,omitempty" xml:"RegionBlockEnable,omitempty"`
	CcTemplate           *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	CcCustomRuleEnable   *int32  `json:"CcCustomRuleEnable,omitempty" xml:"CcCustomRuleEnable,omitempty"`
	CcEnable             *int32  `json:"CcEnable,omitempty" xml:"CcEnable,omitempty"`
}

func (s DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetBlackWhiteListEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.BlackWhiteListEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetAiTemplate(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiTemplate = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetPreciseRuleEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.PreciseRuleEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetDomain(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.Domain = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetAiMode(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiMode = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetAiRuleEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.AiRuleEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetRegionBlockEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.RegionBlockEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcTemplate(v string) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcTemplate = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcCustomRuleEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcCustomRuleEnable = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList) SetCcEnable(v int32) *DescribeWebCcProtectSwitchResponseBodyProtectSwitchList {
	s.CcEnable = &v
	return s
}

type DescribeWebCcProtectSwitchResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebCcProtectSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebCcProtectSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCcProtectSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchResponse) SetHeaders(v map[string]*string) *DescribeWebCcProtectSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCcProtectSwitchResponse) SetBody(v *DescribeWebCcProtectSwitchResponseBody) *DescribeWebCcProtectSwitchResponse {
	s.Body = v
	return s
}

type DescribeWebCCRulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Offset          *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsOffset        *bool   `json:"IsOffset,omitempty" xml:"IsOffset,omitempty"`
}

func (s DescribeWebCCRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCCRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesRequest) SetSourceIp(v string) *DescribeWebCCRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetResourceGroupId(v string) *DescribeWebCCRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetDomain(v string) *DescribeWebCCRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetOffset(v int32) *DescribeWebCCRulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetPageNumber(v int32) *DescribeWebCCRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetPageSize(v string) *DescribeWebCCRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebCCRulesRequest) SetIsOffset(v bool) *DescribeWebCCRulesRequest {
	s.IsOffset = &v
	return s
}

type DescribeWebCCRulesResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebCCRules []*DescribeWebCCRulesResponseBodyWebCCRules `json:"WebCCRules,omitempty" xml:"WebCCRules,omitempty" type:"Repeated"`
}

func (s DescribeWebCCRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCCRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesResponseBody) SetTotalCount(v int64) *DescribeWebCCRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebCCRulesResponseBody) SetRequestId(v string) *DescribeWebCCRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCCRulesResponseBody) SetWebCCRules(v []*DescribeWebCCRulesResponseBodyWebCCRules) *DescribeWebCCRulesResponseBody {
	s.WebCCRules = v
	return s
}

type DescribeWebCCRulesResponseBodyWebCCRules struct {
	Ttl      *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Act      *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeWebCCRulesResponseBodyWebCCRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCCRulesResponseBodyWebCCRules) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetTtl(v int32) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Ttl = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetAct(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Act = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetInterval(v int32) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Interval = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetMode(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Mode = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetName(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Name = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetUri(v string) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Uri = &v
	return s
}

func (s *DescribeWebCCRulesResponseBodyWebCCRules) SetCount(v int32) *DescribeWebCCRulesResponseBodyWebCCRules {
	s.Count = &v
	return s
}

type DescribeWebCCRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebCCRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebCCRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCCRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesResponse) SetHeaders(v map[string]*string) *DescribeWebCCRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCCRulesResponse) SetBody(v *DescribeWebCCRulesResponseBody) *DescribeWebCCRulesResponse {
	s.Body = v
	return s
}

type DescribeWebCustomPortsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebCustomPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCustomPortsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsRequest) SetSourceIp(v string) *DescribeWebCustomPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebCustomPortsRequest) SetLang(v string) *DescribeWebCustomPortsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebCustomPortsRequest) SetResourceGroupId(v string) *DescribeWebCustomPortsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeWebCustomPortsResponseBody struct {
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebCustomPorts []*DescribeWebCustomPortsResponseBodyWebCustomPorts `json:"WebCustomPorts,omitempty" xml:"WebCustomPorts,omitempty" type:"Repeated"`
}

func (s DescribeWebCustomPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCustomPortsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsResponseBody) SetRequestId(v string) *DescribeWebCustomPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCustomPortsResponseBody) SetWebCustomPorts(v []*DescribeWebCustomPortsResponseBodyWebCustomPorts) *DescribeWebCustomPortsResponseBody {
	s.WebCustomPorts = v
	return s
}

type DescribeWebCustomPortsResponseBodyWebCustomPorts struct {
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	ProxyType  *string   `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeWebCustomPortsResponseBodyWebCustomPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCustomPortsResponseBodyWebCustomPorts) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) SetProxyPorts(v []*string) *DescribeWebCustomPortsResponseBodyWebCustomPorts {
	s.ProxyPorts = v
	return s
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) SetProxyType(v string) *DescribeWebCustomPortsResponseBodyWebCustomPorts {
	s.ProxyType = &v
	return s
}

type DescribeWebCustomPortsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebCustomPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebCustomPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebCustomPortsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsResponse) SetHeaders(v map[string]*string) *DescribeWebCustomPortsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCustomPortsResponse) SetBody(v *DescribeWebCustomPortsResponseBody) *DescribeWebCustomPortsResponse {
	s.Body = v
	return s
}

type DescribeWebInstanceRelationsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domains         []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebInstanceRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebInstanceRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsRequest) SetSourceIp(v string) *DescribeWebInstanceRelationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebInstanceRelationsRequest) SetResourceGroupId(v string) *DescribeWebInstanceRelationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebInstanceRelationsRequest) SetDomains(v []*string) *DescribeWebInstanceRelationsRequest {
	s.Domains = v
	return s
}

type DescribeWebInstanceRelationsResponseBody struct {
	RequestId            *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebInstanceRelations []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations `json:"WebInstanceRelations,omitempty" xml:"WebInstanceRelations,omitempty" type:"Repeated"`
}

func (s DescribeWebInstanceRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponseBody) SetRequestId(v string) *DescribeWebInstanceRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBody) SetWebInstanceRelations(v []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) *DescribeWebInstanceRelationsResponseBody {
	s.WebInstanceRelations = v
	return s
}

type DescribeWebInstanceRelationsResponseBodyWebInstanceRelations struct {
	Domain          *string                                                                        `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceDetails []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) SetDomain(v string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations {
	s.Domain = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations) SetInstanceDetails(v []*DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelations {
	s.InstanceDetails = v
	return s
}

type DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails struct {
	EipList         []*string `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	FunctionVersion *string   `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) SetEipList(v []*string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	s.EipList = v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) SetFunctionVersion(v string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails) SetInstanceId(v string) *DescribeWebInstanceRelationsResponseBodyWebInstanceRelationsInstanceDetails {
	s.InstanceId = &v
	return s
}

type DescribeWebInstanceRelationsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebInstanceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebInstanceRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebInstanceRelationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebInstanceRelationsResponse) SetHeaders(v map[string]*string) *DescribeWebInstanceRelationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebInstanceRelationsResponse) SetBody(v *DescribeWebInstanceRelationsResponseBody) *DescribeWebInstanceRelationsResponse {
	s.Body = v
	return s
}

type DescribeWebPreciseAccessRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domains         []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s DescribeWebPreciseAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleRequest) SetSourceIp(v string) *DescribeWebPreciseAccessRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleRequest) SetResourceGroupId(v string) *DescribeWebPreciseAccessRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleRequest) SetDomains(v []*string) *DescribeWebPreciseAccessRuleRequest {
	s.Domains = v
	return s
}

type DescribeWebPreciseAccessRuleResponseBody struct {
	PreciseAccessConfigList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList `json:"PreciseAccessConfigList,omitempty" xml:"PreciseAccessConfigList,omitempty" type:"Repeated"`
	RequestId               *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebPreciseAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBody) SetPreciseAccessConfigList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) *DescribeWebPreciseAccessRuleResponseBody {
	s.PreciseAccessConfigList = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBody) SetRequestId(v string) *DescribeWebPreciseAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList struct {
	Domain   *string                                                                    `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RuleList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) SetDomain(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList {
	s.Domain = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList) SetRuleList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigList {
	s.RuleList = v
	return s
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList struct {
	Action        *string                                                                                 `json:"Action,omitempty" xml:"Action,omitempty"`
	Owner         *string                                                                                 `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Expires       *int64                                                                                  `json:"Expires,omitempty" xml:"Expires,omitempty"`
	Name          *string                                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ConditionList []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetAction(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Action = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetOwner(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Owner = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetExpires(v int64) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Expires = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetName(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.Name = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList) SetConditionList(v []*DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleList {
	s.ConditionList = v
	return s
}

type DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList struct {
	MatchMethod *string `json:"MatchMethod,omitempty" xml:"MatchMethod,omitempty"`
	Field       *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	HeaderName  *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetMatchMethod(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.MatchMethod = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetField(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.Field = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetContent(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.Content = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList) SetHeaderName(v string) *DescribeWebPreciseAccessRuleResponseBodyPreciseAccessConfigListRuleListConditionList {
	s.HeaderName = &v
	return s
}

type DescribeWebPreciseAccessRuleResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebPreciseAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebPreciseAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleResponse) SetHeaders(v map[string]*string) *DescribeWebPreciseAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebPreciseAccessRuleResponse) SetBody(v *DescribeWebPreciseAccessRuleResponseBody) *DescribeWebPreciseAccessRuleResponse {
	s.Body = v
	return s
}

type DescribeWebRulesRequest struct {
	SourceIp           *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId    *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain             *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	QueryDomainPattern *string   `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
	PageNumber         *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeWebRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesRequest) SetSourceIp(v string) *DescribeWebRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebRulesRequest) SetResourceGroupId(v string) *DescribeWebRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebRulesRequest) SetDomain(v string) *DescribeWebRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebRulesRequest) SetQueryDomainPattern(v string) *DescribeWebRulesRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeWebRulesRequest) SetPageNumber(v int32) *DescribeWebRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWebRulesRequest) SetPageSize(v int32) *DescribeWebRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebRulesRequest) SetInstanceIds(v []*string) *DescribeWebRulesRequest {
	s.InstanceIds = v
	return s
}

type DescribeWebRulesResponseBody struct {
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WebRules   []*DescribeWebRulesResponseBodyWebRules `json:"WebRules,omitempty" xml:"WebRules,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBody) SetTotalCount(v int64) *DescribeWebRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebRulesResponseBody) SetWebRules(v []*DescribeWebRulesResponseBodyWebRules) *DescribeWebRulesResponseBody {
	s.WebRules = v
	return s
}

func (s *DescribeWebRulesResponseBody) SetRequestId(v string) *DescribeWebRulesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWebRulesResponseBodyWebRules struct {
	Domain           *string                                            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Http2HttpsEnable *bool                                              `json:"Http2HttpsEnable,omitempty" xml:"Http2HttpsEnable,omitempty"`
	CustomCiphers    []*string                                          `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	BlackList        []*string                                          `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	WhiteList        []*string                                          `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
	RealServers      []*DescribeWebRulesResponseBodyWebRulesRealServers `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	ProxyTypes       []*DescribeWebRulesResponseBodyWebRulesProxyTypes  `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	SslProtocols     *string                                            `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	CcTemplate       *string                                            `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	CcEnabled        *bool                                              `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	SslCiphers       *string                                            `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	CcRuleEnabled    *bool                                              `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	Ssl13Enabled     *bool                                              `json:"Ssl13Enabled,omitempty" xml:"Ssl13Enabled,omitempty"`
	ProxyEnabled     *bool                                              `json:"ProxyEnabled,omitempty" xml:"ProxyEnabled,omitempty"`
	CertName         *string                                            `json:"CertName,omitempty" xml:"CertName,omitempty"`
	PolicyMode       *string                                            `json:"PolicyMode,omitempty" xml:"PolicyMode,omitempty"`
	Cname            *string                                            `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Http2Enable      *bool                                              `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
	Https2HttpEnable *bool                                              `json:"Https2HttpEnable,omitempty" xml:"Https2HttpEnable,omitempty"`
}

func (s DescribeWebRulesResponseBodyWebRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRules) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRules) SetDomain(v string) *DescribeWebRulesResponseBodyWebRules {
	s.Domain = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetHttp2HttpsEnable(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Http2HttpsEnable = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCustomCiphers(v []*string) *DescribeWebRulesResponseBodyWebRules {
	s.CustomCiphers = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetBlackList(v []*string) *DescribeWebRulesResponseBodyWebRules {
	s.BlackList = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetWhiteList(v []*string) *DescribeWebRulesResponseBodyWebRules {
	s.WhiteList = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetRealServers(v []*DescribeWebRulesResponseBodyWebRulesRealServers) *DescribeWebRulesResponseBodyWebRules {
	s.RealServers = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetProxyTypes(v []*DescribeWebRulesResponseBodyWebRulesProxyTypes) *DescribeWebRulesResponseBodyWebRules {
	s.ProxyTypes = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetSslProtocols(v string) *DescribeWebRulesResponseBodyWebRules {
	s.SslProtocols = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCcTemplate(v string) *DescribeWebRulesResponseBodyWebRules {
	s.CcTemplate = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCcEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.CcEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetSslCiphers(v string) *DescribeWebRulesResponseBodyWebRules {
	s.SslCiphers = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCcRuleEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetSsl13Enabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Ssl13Enabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetProxyEnabled(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.ProxyEnabled = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCertName(v string) *DescribeWebRulesResponseBodyWebRules {
	s.CertName = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetPolicyMode(v string) *DescribeWebRulesResponseBodyWebRules {
	s.PolicyMode = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetCname(v string) *DescribeWebRulesResponseBodyWebRules {
	s.Cname = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetHttp2Enable(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Http2Enable = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRules) SetHttps2HttpEnable(v bool) *DescribeWebRulesResponseBodyWebRules {
	s.Https2HttpEnable = &v
	return s
}

type DescribeWebRulesResponseBodyWebRulesRealServers struct {
	RsType     *int32  `json:"RsType,omitempty" xml:"RsType,omitempty"`
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
}

func (s DescribeWebRulesResponseBodyWebRulesRealServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRulesRealServers) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) SetRsType(v int32) *DescribeWebRulesResponseBodyWebRulesRealServers {
	s.RsType = &v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesRealServers) SetRealServer(v string) *DescribeWebRulesResponseBodyWebRulesRealServers {
	s.RealServer = &v
	return s
}

type DescribeWebRulesResponseBodyWebRulesProxyTypes struct {
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	ProxyType  *string   `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeWebRulesResponseBodyWebRulesProxyTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebRulesResponseBodyWebRulesProxyTypes) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) SetProxyPorts(v []*string) *DescribeWebRulesResponseBodyWebRulesProxyTypes {
	s.ProxyPorts = v
	return s
}

func (s *DescribeWebRulesResponseBodyWebRulesProxyTypes) SetProxyType(v string) *DescribeWebRulesResponseBodyWebRulesProxyTypes {
	s.ProxyType = &v
	return s
}

type DescribeWebRulesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponse) SetHeaders(v map[string]*string) *DescribeWebRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebRulesResponse) SetBody(v *DescribeWebRulesResponseBody) *DescribeWebRulesResponse {
	s.Body = v
	return s
}

type DetachSceneDefenseObjectRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Objects    *string `json:"Objects,omitempty" xml:"Objects,omitempty"`
}

func (s DetachSceneDefenseObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachSceneDefenseObjectRequest) GoString() string {
	return s.String()
}

func (s *DetachSceneDefenseObjectRequest) SetSourceIp(v string) *DetachSceneDefenseObjectRequest {
	s.SourceIp = &v
	return s
}

func (s *DetachSceneDefenseObjectRequest) SetPolicyId(v string) *DetachSceneDefenseObjectRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachSceneDefenseObjectRequest) SetObjectType(v string) *DetachSceneDefenseObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *DetachSceneDefenseObjectRequest) SetObjects(v string) *DetachSceneDefenseObjectRequest {
	s.Objects = &v
	return s
}

type DetachSceneDefenseObjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachSceneDefenseObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachSceneDefenseObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetachSceneDefenseObjectResponseBody) SetRequestId(v string) *DetachSceneDefenseObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachSceneDefenseObjectResponseBody) SetSuccess(v bool) *DetachSceneDefenseObjectResponseBody {
	s.Success = &v
	return s
}

type DetachSceneDefenseObjectResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachSceneDefenseObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachSceneDefenseObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachSceneDefenseObjectResponse) GoString() string {
	return s.String()
}

func (s *DetachSceneDefenseObjectResponse) SetHeaders(v map[string]*string) *DetachSceneDefenseObjectResponse {
	s.Headers = v
	return s
}

func (s *DetachSceneDefenseObjectResponse) SetBody(v *DetachSceneDefenseObjectResponseBody) *DetachSceneDefenseObjectResponse {
	s.Body = v
	return s
}

type DisableSceneDefensePolicyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DisableSceneDefensePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableSceneDefensePolicyRequest) SetSourceIp(v string) *DisableSceneDefensePolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableSceneDefensePolicyRequest) SetPolicyId(v string) *DisableSceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

type DisableSceneDefensePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSceneDefensePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSceneDefensePolicyResponseBody) SetRequestId(v string) *DisableSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSceneDefensePolicyResponseBody) SetSuccess(v bool) *DisableSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

type DisableSceneDefensePolicyResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSceneDefensePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *DisableSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableSceneDefensePolicyResponse) SetBody(v *DisableSceneDefensePolicyResponseBody) *DisableSceneDefensePolicyResponse {
	s.Body = v
	return s
}

type DisableWebAccessLogConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DisableWebAccessLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableWebAccessLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableWebAccessLogConfigRequest) SetSourceIp(v string) *DisableWebAccessLogConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableWebAccessLogConfigRequest) SetLang(v string) *DisableWebAccessLogConfigRequest {
	s.Lang = &v
	return s
}

func (s *DisableWebAccessLogConfigRequest) SetResourceGroupId(v string) *DisableWebAccessLogConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableWebAccessLogConfigRequest) SetDomain(v string) *DisableWebAccessLogConfigRequest {
	s.Domain = &v
	return s
}

type DisableWebAccessLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableWebAccessLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableWebAccessLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWebAccessLogConfigResponseBody) SetRequestId(v string) *DisableWebAccessLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type DisableWebAccessLogConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableWebAccessLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableWebAccessLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableWebAccessLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DisableWebAccessLogConfigResponse) SetHeaders(v map[string]*string) *DisableWebAccessLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DisableWebAccessLogConfigResponse) SetBody(v *DisableWebAccessLogConfigResponseBody) *DisableWebAccessLogConfigResponse {
	s.Body = v
	return s
}

type DisableWebCCRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DisableWebCCRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableWebCCRequest) GoString() string {
	return s.String()
}

func (s *DisableWebCCRequest) SetSourceIp(v string) *DisableWebCCRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableWebCCRequest) SetResourceGroupId(v string) *DisableWebCCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableWebCCRequest) SetDomain(v string) *DisableWebCCRequest {
	s.Domain = &v
	return s
}

type DisableWebCCResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableWebCCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableWebCCResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWebCCResponseBody) SetRequestId(v string) *DisableWebCCResponseBody {
	s.RequestId = &v
	return s
}

type DisableWebCCResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableWebCCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableWebCCResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableWebCCResponse) GoString() string {
	return s.String()
}

func (s *DisableWebCCResponse) SetHeaders(v map[string]*string) *DisableWebCCResponse {
	s.Headers = v
	return s
}

func (s *DisableWebCCResponse) SetBody(v *DisableWebCCResponseBody) *DisableWebCCResponse {
	s.Body = v
	return s
}

type DisableWebCCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DisableWebCCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableWebCCRuleRequest) SetSourceIp(v string) *DisableWebCCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableWebCCRuleRequest) SetResourceGroupId(v string) *DisableWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableWebCCRuleRequest) SetDomain(v string) *DisableWebCCRuleRequest {
	s.Domain = &v
	return s
}

type DisableWebCCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableWebCCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWebCCRuleResponseBody) SetRequestId(v string) *DisableWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

type DisableWebCCRuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableWebCCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableWebCCRuleResponse) SetHeaders(v map[string]*string) *DisableWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableWebCCRuleResponse) SetBody(v *DisableWebCCRuleResponseBody) *DisableWebCCRuleResponse {
	s.Body = v
	return s
}

type EmptyAutoCcBlacklistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EmptyAutoCcBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s EmptyAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *EmptyAutoCcBlacklistRequest) SetSourceIp(v string) *EmptyAutoCcBlacklistRequest {
	s.SourceIp = &v
	return s
}

func (s *EmptyAutoCcBlacklistRequest) SetInstanceId(v string) *EmptyAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

type EmptyAutoCcBlacklistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptyAutoCcBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmptyAutoCcBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *EmptyAutoCcBlacklistResponseBody) SetRequestId(v string) *EmptyAutoCcBlacklistResponseBody {
	s.RequestId = &v
	return s
}

type EmptyAutoCcBlacklistResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EmptyAutoCcBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EmptyAutoCcBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s EmptyAutoCcBlacklistResponse) GoString() string {
	return s.String()
}

func (s *EmptyAutoCcBlacklistResponse) SetHeaders(v map[string]*string) *EmptyAutoCcBlacklistResponse {
	s.Headers = v
	return s
}

func (s *EmptyAutoCcBlacklistResponse) SetBody(v *EmptyAutoCcBlacklistResponseBody) *EmptyAutoCcBlacklistResponse {
	s.Body = v
	return s
}

type EmptyAutoCcWhitelistRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EmptyAutoCcWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s EmptyAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *EmptyAutoCcWhitelistRequest) SetSourceIp(v string) *EmptyAutoCcWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *EmptyAutoCcWhitelistRequest) SetInstanceId(v string) *EmptyAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

type EmptyAutoCcWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptyAutoCcWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmptyAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *EmptyAutoCcWhitelistResponseBody) SetRequestId(v string) *EmptyAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type EmptyAutoCcWhitelistResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EmptyAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EmptyAutoCcWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s EmptyAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *EmptyAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *EmptyAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *EmptyAutoCcWhitelistResponse) SetBody(v *EmptyAutoCcWhitelistResponseBody) *EmptyAutoCcWhitelistResponse {
	s.Body = v
	return s
}

type EmptySlsLogstoreRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s EmptySlsLogstoreRequest) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreRequest) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreRequest) SetSourceIp(v string) *EmptySlsLogstoreRequest {
	s.SourceIp = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetLang(v string) *EmptySlsLogstoreRequest {
	s.Lang = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetResourceGroupId(v string) *EmptySlsLogstoreRequest {
	s.ResourceGroupId = &v
	return s
}

type EmptySlsLogstoreResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptySlsLogstoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreResponseBody) SetRequestId(v string) *EmptySlsLogstoreResponseBody {
	s.RequestId = &v
	return s
}

type EmptySlsLogstoreResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EmptySlsLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EmptySlsLogstoreResponse) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreResponse) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreResponse) SetHeaders(v map[string]*string) *EmptySlsLogstoreResponse {
	s.Headers = v
	return s
}

func (s *EmptySlsLogstoreResponse) SetBody(v *EmptySlsLogstoreResponseBody) *EmptySlsLogstoreResponse {
	s.Body = v
	return s
}

type EnableSceneDefensePolicyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s EnableSceneDefensePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *EnableSceneDefensePolicyRequest) SetSourceIp(v string) *EnableSceneDefensePolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableSceneDefensePolicyRequest) SetPolicyId(v string) *EnableSceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

type EnableSceneDefensePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSceneDefensePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSceneDefensePolicyResponseBody) SetRequestId(v string) *EnableSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSceneDefensePolicyResponseBody) SetSuccess(v bool) *EnableSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

type EnableSceneDefensePolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSceneDefensePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *EnableSceneDefensePolicyResponse) SetHeaders(v map[string]*string) *EnableSceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *EnableSceneDefensePolicyResponse) SetBody(v *EnableSceneDefensePolicyResponseBody) *EnableSceneDefensePolicyResponse {
	s.Body = v
	return s
}

type EnableWebAccessLogConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s EnableWebAccessLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableWebAccessLogConfigRequest) GoString() string {
	return s.String()
}

func (s *EnableWebAccessLogConfigRequest) SetSourceIp(v string) *EnableWebAccessLogConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableWebAccessLogConfigRequest) SetLang(v string) *EnableWebAccessLogConfigRequest {
	s.Lang = &v
	return s
}

func (s *EnableWebAccessLogConfigRequest) SetResourceGroupId(v string) *EnableWebAccessLogConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableWebAccessLogConfigRequest) SetDomain(v string) *EnableWebAccessLogConfigRequest {
	s.Domain = &v
	return s
}

type EnableWebAccessLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWebAccessLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableWebAccessLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *EnableWebAccessLogConfigResponseBody) SetRequestId(v string) *EnableWebAccessLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type EnableWebAccessLogConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableWebAccessLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableWebAccessLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableWebAccessLogConfigResponse) GoString() string {
	return s.String()
}

func (s *EnableWebAccessLogConfigResponse) SetHeaders(v map[string]*string) *EnableWebAccessLogConfigResponse {
	s.Headers = v
	return s
}

func (s *EnableWebAccessLogConfigResponse) SetBody(v *EnableWebAccessLogConfigResponseBody) *EnableWebAccessLogConfigResponse {
	s.Body = v
	return s
}

type EnableWebCCRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s EnableWebCCRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableWebCCRequest) GoString() string {
	return s.String()
}

func (s *EnableWebCCRequest) SetSourceIp(v string) *EnableWebCCRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableWebCCRequest) SetResourceGroupId(v string) *EnableWebCCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableWebCCRequest) SetDomain(v string) *EnableWebCCRequest {
	s.Domain = &v
	return s
}

type EnableWebCCResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWebCCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableWebCCResponseBody) GoString() string {
	return s.String()
}

func (s *EnableWebCCResponseBody) SetRequestId(v string) *EnableWebCCResponseBody {
	s.RequestId = &v
	return s
}

type EnableWebCCResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableWebCCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableWebCCResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableWebCCResponse) GoString() string {
	return s.String()
}

func (s *EnableWebCCResponse) SetHeaders(v map[string]*string) *EnableWebCCResponse {
	s.Headers = v
	return s
}

func (s *EnableWebCCResponse) SetBody(v *EnableWebCCResponseBody) *EnableWebCCResponse {
	s.Body = v
	return s
}

type EnableWebCCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s EnableWebCCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableWebCCRuleRequest) SetSourceIp(v string) *EnableWebCCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableWebCCRuleRequest) SetResourceGroupId(v string) *EnableWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableWebCCRuleRequest) SetDomain(v string) *EnableWebCCRuleRequest {
	s.Domain = &v
	return s
}

type EnableWebCCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWebCCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableWebCCRuleResponseBody) SetRequestId(v string) *EnableWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableWebCCRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableWebCCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableWebCCRuleResponse) SetHeaders(v map[string]*string) *EnableWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableWebCCRuleResponse) SetBody(v *EnableWebCCRuleResponseBody) *EnableWebCCRuleResponse {
	s.Body = v
	return s
}

type ModifyBlackholeStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BlackholeStatus *string `json:"BlackholeStatus,omitempty" xml:"BlackholeStatus,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyBlackholeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBlackholeStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyBlackholeStatusRequest) SetSourceIp(v string) *ModifyBlackholeStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyBlackholeStatusRequest) SetBlackholeStatus(v string) *ModifyBlackholeStatusRequest {
	s.BlackholeStatus = &v
	return s
}

func (s *ModifyBlackholeStatusRequest) SetInstanceId(v string) *ModifyBlackholeStatusRequest {
	s.InstanceId = &v
	return s
}

type ModifyBlackholeStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBlackholeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBlackholeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBlackholeStatusResponseBody) SetRequestId(v string) *ModifyBlackholeStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBlackholeStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBlackholeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBlackholeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBlackholeStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyBlackholeStatusResponse) SetHeaders(v map[string]*string) *ModifyBlackholeStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyBlackholeStatusResponse) SetBody(v *ModifyBlackholeStatusResponseBody) *ModifyBlackholeStatusResponse {
	s.Body = v
	return s
}

type ModifyBlockStatusRequest struct {
	SourceIp   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status     *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Duration   *int32    `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lines      []*string `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s ModifyBlockStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBlockStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyBlockStatusRequest) SetSourceIp(v string) *ModifyBlockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetLang(v string) *ModifyBlockStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetStatus(v string) *ModifyBlockStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetDuration(v int32) *ModifyBlockStatusRequest {
	s.Duration = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetInstanceId(v string) *ModifyBlockStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBlockStatusRequest) SetLines(v []*string) *ModifyBlockStatusRequest {
	s.Lines = v
	return s
}

type ModifyBlockStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBlockStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBlockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBlockStatusResponseBody) SetRequestId(v string) *ModifyBlockStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBlockStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBlockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBlockStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBlockStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyBlockStatusResponse) SetHeaders(v map[string]*string) *ModifyBlockStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyBlockStatusResponse) SetBody(v *ModifyBlockStatusResponseBody) *ModifyBlockStatusResponse {
	s.Body = v
	return s
}

type ModifyCnameReuseRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Cname           *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Enable          *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyCnameReuseRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCnameReuseRequest) GoString() string {
	return s.String()
}

func (s *ModifyCnameReuseRequest) SetSourceIp(v string) *ModifyCnameReuseRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetResourceGroupId(v string) *ModifyCnameReuseRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetDomain(v string) *ModifyCnameReuseRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetCname(v string) *ModifyCnameReuseRequest {
	s.Cname = &v
	return s
}

func (s *ModifyCnameReuseRequest) SetEnable(v int32) *ModifyCnameReuseRequest {
	s.Enable = &v
	return s
}

type ModifyCnameReuseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCnameReuseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCnameReuseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCnameReuseResponseBody) SetRequestId(v string) *ModifyCnameReuseResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCnameReuseResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCnameReuseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCnameReuseResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCnameReuseResponse) GoString() string {
	return s.String()
}

func (s *ModifyCnameReuseResponse) SetHeaders(v map[string]*string) *ModifyCnameReuseResponse {
	s.Headers = v
	return s
}

func (s *ModifyCnameReuseResponse) SetBody(v *ModifyCnameReuseResponseBody) *ModifyCnameReuseResponse {
	s.Body = v
	return s
}

type ModifyElasticBandWidthRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ElasticBandwidth *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyElasticBandWidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthRequest) SetSourceIp(v string) *ModifyElasticBandWidthRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetInstanceId(v string) *ModifyElasticBandWidthRequest {
	s.InstanceId = &v
	return s
}

type ModifyElasticBandWidthResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticBandWidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponseBody) SetRequestId(v string) *ModifyElasticBandWidthResponseBody {
	s.RequestId = &v
	return s
}

type ModifyElasticBandWidthResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyElasticBandWidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyElasticBandWidthResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponse) SetHeaders(v map[string]*string) *ModifyElasticBandWidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticBandWidthResponse) SetBody(v *ModifyElasticBandWidthResponseBody) *ModifyElasticBandWidthResponse {
	s.Body = v
	return s
}

type ModifyFullLogTtlRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyFullLogTtlRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlRequest) SetSourceIp(v string) *ModifyFullLogTtlRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetLang(v string) *ModifyFullLogTtlRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetTtl(v int32) *ModifyFullLogTtlRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetResourceGroupId(v string) *ModifyFullLogTtlRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyFullLogTtlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFullLogTtlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponseBody) SetRequestId(v string) *ModifyFullLogTtlResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFullLogTtlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFullLogTtlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFullLogTtlResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlResponse) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponse) SetHeaders(v map[string]*string) *ModifyFullLogTtlResponse {
	s.Headers = v
	return s
}

func (s *ModifyFullLogTtlResponse) SetBody(v *ModifyFullLogTtlResponseBody) *ModifyFullLogTtlResponse {
	s.Body = v
	return s
}

type ModifyHealthCheckConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int32  `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	HealthCheck     *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
}

func (s ModifyHealthCheckConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHealthCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigRequest) SetSourceIp(v string) *ModifyHealthCheckConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetInstanceId(v string) *ModifyHealthCheckConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetForwardProtocol(v string) *ModifyHealthCheckConfigRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetFrontendPort(v int32) *ModifyHealthCheckConfigRequest {
	s.FrontendPort = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetHealthCheck(v string) *ModifyHealthCheckConfigRequest {
	s.HealthCheck = &v
	return s
}

type ModifyHealthCheckConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHealthCheckConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHealthCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigResponseBody) SetRequestId(v string) *ModifyHealthCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHealthCheckConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHealthCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHealthCheckConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHealthCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigResponse) SetHeaders(v map[string]*string) *ModifyHealthCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyHealthCheckConfigResponse) SetBody(v *ModifyHealthCheckConfigResponseBody) *ModifyHealthCheckConfigResponse {
	s.Body = v
	return s
}

type ModifyHttp2EnableRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enable          *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyHttp2EnableRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHttp2EnableRequest) GoString() string {
	return s.String()
}

func (s *ModifyHttp2EnableRequest) SetSourceIp(v string) *ModifyHttp2EnableRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHttp2EnableRequest) SetResourceGroupId(v string) *ModifyHttp2EnableRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyHttp2EnableRequest) SetDomain(v string) *ModifyHttp2EnableRequest {
	s.Domain = &v
	return s
}

func (s *ModifyHttp2EnableRequest) SetEnable(v int32) *ModifyHttp2EnableRequest {
	s.Enable = &v
	return s
}

type ModifyHttp2EnableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHttp2EnableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHttp2EnableResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHttp2EnableResponseBody) SetRequestId(v string) *ModifyHttp2EnableResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHttp2EnableResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHttp2EnableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHttp2EnableResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHttp2EnableResponse) GoString() string {
	return s.String()
}

func (s *ModifyHttp2EnableResponse) SetHeaders(v map[string]*string) *ModifyHttp2EnableResponse {
	s.Headers = v
	return s
}

func (s *ModifyHttp2EnableResponse) SetBody(v *ModifyHttp2EnableResponseBody) *ModifyHttp2EnableResponse {
	s.Body = v
	return s
}

type ModifyInstanceRemarkRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyInstanceRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkRequest) SetSourceIp(v string) *ModifyInstanceRemarkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetInstanceId(v string) *ModifyInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetRemark(v string) *ModifyInstanceRemarkRequest {
	s.Remark = &v
	return s
}

type ModifyInstanceRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponseBody) SetRequestId(v string) *ModifyInstanceRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceRemarkResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponse) SetHeaders(v map[string]*string) *ModifyInstanceRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceRemarkResponse) SetBody(v *ModifyInstanceRemarkResponseBody) *ModifyInstanceRemarkResponse {
	s.Body = v
	return s
}

type ModifyNetworkRuleAttributeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int32  `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyNetworkRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkRuleAttributeRequest) SetSourceIp(v string) *ModifyNetworkRuleAttributeRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetInstanceId(v string) *ModifyNetworkRuleAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetForwardProtocol(v string) *ModifyNetworkRuleAttributeRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetFrontendPort(v int32) *ModifyNetworkRuleAttributeRequest {
	s.FrontendPort = &v
	return s
}

func (s *ModifyNetworkRuleAttributeRequest) SetConfig(v string) *ModifyNetworkRuleAttributeRequest {
	s.Config = &v
	return s
}

type ModifyNetworkRuleAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkRuleAttributeResponseBody) SetRequestId(v string) *ModifyNetworkRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkRuleAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkRuleAttributeResponse) SetHeaders(v map[string]*string) *ModifyNetworkRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkRuleAttributeResponse) SetBody(v *ModifyNetworkRuleAttributeResponseBody) *ModifyNetworkRuleAttributeResponse {
	s.Body = v
	return s
}

type ModifyPortAutoCcStatusRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Switch     *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyPortAutoCcStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPortAutoCcStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyPortAutoCcStatusRequest) SetSourceIp(v string) *ModifyPortAutoCcStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) SetLang(v string) *ModifyPortAutoCcStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) SetInstanceId(v string) *ModifyPortAutoCcStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) SetSwitch(v string) *ModifyPortAutoCcStatusRequest {
	s.Switch = &v
	return s
}

func (s *ModifyPortAutoCcStatusRequest) SetMode(v string) *ModifyPortAutoCcStatusRequest {
	s.Mode = &v
	return s
}

type ModifyPortAutoCcStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPortAutoCcStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPortAutoCcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPortAutoCcStatusResponseBody) SetRequestId(v string) *ModifyPortAutoCcStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPortAutoCcStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPortAutoCcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPortAutoCcStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPortAutoCcStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyPortAutoCcStatusResponse) SetHeaders(v map[string]*string) *ModifyPortAutoCcStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyPortAutoCcStatusResponse) SetBody(v *ModifyPortAutoCcStatusResponseBody) *ModifyPortAutoCcStatusResponse {
	s.Body = v
	return s
}

type ModifySceneDefensePolicyRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PolicyId  *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Template  *string `json:"Template,omitempty" xml:"Template,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ModifySceneDefensePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySceneDefensePolicyRequest) SetSourceIp(v string) *ModifySceneDefensePolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetPolicyId(v string) *ModifySceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetName(v string) *ModifySceneDefensePolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetTemplate(v string) *ModifySceneDefensePolicyRequest {
	s.Template = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetStartTime(v int64) *ModifySceneDefensePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetEndTime(v int64) *ModifySceneDefensePolicyRequest {
	s.EndTime = &v
	return s
}

type ModifySceneDefensePolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySceneDefensePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySceneDefensePolicyResponseBody) SetRequestId(v string) *ModifySceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySceneDefensePolicyResponseBody) SetSuccess(v bool) *ModifySceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

type ModifySceneDefensePolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySceneDefensePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySceneDefensePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySceneDefensePolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySceneDefensePolicyResponse) SetHeaders(v map[string]*string) *ModifySceneDefensePolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySceneDefensePolicyResponse) SetBody(v *ModifySceneDefensePolicyResponseBody) *ModifySceneDefensePolicyResponse {
	s.Body = v
	return s
}

type ModifySchedulerRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Rules           *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType        *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Param           *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s ModifySchedulerRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifySchedulerRuleRequest) SetSourceIp(v string) *ModifySchedulerRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifySchedulerRuleRequest) SetResourceGroupId(v string) *ModifySchedulerRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySchedulerRuleRequest) SetRules(v string) *ModifySchedulerRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifySchedulerRuleRequest) SetRuleName(v string) *ModifySchedulerRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifySchedulerRuleRequest) SetRuleType(v int32) *ModifySchedulerRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifySchedulerRuleRequest) SetParam(v string) *ModifySchedulerRuleRequest {
	s.Param = &v
	return s
}

type ModifySchedulerRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Cname     *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	RuleName  *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifySchedulerRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySchedulerRuleResponseBody) SetRequestId(v string) *ModifySchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySchedulerRuleResponseBody) SetCname(v string) *ModifySchedulerRuleResponseBody {
	s.Cname = &v
	return s
}

func (s *ModifySchedulerRuleResponseBody) SetRuleName(v string) *ModifySchedulerRuleResponseBody {
	s.RuleName = &v
	return s
}

type ModifySchedulerRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySchedulerRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifySchedulerRuleResponse) SetHeaders(v map[string]*string) *ModifySchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifySchedulerRuleResponse) SetBody(v *ModifySchedulerRuleResponseBody) *ModifySchedulerRuleResponse {
	s.Body = v
	return s
}

type ModifyTlsConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyTlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTlsConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTlsConfigRequest) SetSourceIp(v string) *ModifyTlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyTlsConfigRequest) SetResourceGroupId(v string) *ModifyTlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyTlsConfigRequest) SetDomain(v string) *ModifyTlsConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyTlsConfigRequest) SetConfig(v string) *ModifyTlsConfigRequest {
	s.Config = &v
	return s
}

type ModifyTlsConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTlsConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTlsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTlsConfigResponseBody) SetRequestId(v string) *ModifyTlsConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTlsConfigResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTlsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTlsConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTlsConfigResponse) SetHeaders(v map[string]*string) *ModifyTlsConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTlsConfigResponse) SetBody(v *ModifyTlsConfigResponseBody) *ModifyTlsConfigResponse {
	s.Body = v
	return s
}

type ModifyWebAccessModeRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	AccessMode *int32  `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s ModifyWebAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAccessModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAccessModeRequest) SetSourceIp(v string) *ModifyWebAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebAccessModeRequest) SetDomain(v string) *ModifyWebAccessModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAccessModeRequest) SetAccessMode(v int32) *ModifyWebAccessModeRequest {
	s.AccessMode = &v
	return s
}

type ModifyWebAccessModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAccessModeResponseBody) SetRequestId(v string) *ModifyWebAccessModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebAccessModeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAccessModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAccessModeResponse) SetHeaders(v map[string]*string) *ModifyWebAccessModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAccessModeResponse) SetBody(v *ModifyWebAccessModeResponseBody) *ModifyWebAccessModeResponse {
	s.Body = v
	return s
}

type ModifyWebAIProtectModeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyWebAIProtectModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAIProtectModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectModeRequest) SetSourceIp(v string) *ModifyWebAIProtectModeRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebAIProtectModeRequest) SetResourceGroupId(v string) *ModifyWebAIProtectModeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAIProtectModeRequest) SetDomain(v string) *ModifyWebAIProtectModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAIProtectModeRequest) SetConfig(v string) *ModifyWebAIProtectModeRequest {
	s.Config = &v
	return s
}

type ModifyWebAIProtectModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAIProtectModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAIProtectModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectModeResponseBody) SetRequestId(v string) *ModifyWebAIProtectModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebAIProtectModeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebAIProtectModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebAIProtectModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAIProtectModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectModeResponse) SetHeaders(v map[string]*string) *ModifyWebAIProtectModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAIProtectModeResponse) SetBody(v *ModifyWebAIProtectModeResponseBody) *ModifyWebAIProtectModeResponse {
	s.Body = v
	return s
}

type ModifyWebAIProtectSwitchRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyWebAIProtectSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAIProtectSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectSwitchRequest) SetSourceIp(v string) *ModifyWebAIProtectSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebAIProtectSwitchRequest) SetResourceGroupId(v string) *ModifyWebAIProtectSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAIProtectSwitchRequest) SetDomain(v string) *ModifyWebAIProtectSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAIProtectSwitchRequest) SetConfig(v string) *ModifyWebAIProtectSwitchRequest {
	s.Config = &v
	return s
}

type ModifyWebAIProtectSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAIProtectSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAIProtectSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectSwitchResponseBody) SetRequestId(v string) *ModifyWebAIProtectSwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebAIProtectSwitchResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebAIProtectSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebAIProtectSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAIProtectSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebAIProtectSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAIProtectSwitchResponse) SetBody(v *ModifyWebAIProtectSwitchResponseBody) *ModifyWebAIProtectSwitchResponse {
	s.Body = v
	return s
}

type ModifyWebAreaBlockRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Regions         []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s ModifyWebAreaBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAreaBlockRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockRequest) SetSourceIp(v string) *ModifyWebAreaBlockRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebAreaBlockRequest) SetResourceGroupId(v string) *ModifyWebAreaBlockRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAreaBlockRequest) SetDomain(v string) *ModifyWebAreaBlockRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAreaBlockRequest) SetRegions(v []*string) *ModifyWebAreaBlockRequest {
	s.Regions = v
	return s
}

type ModifyWebAreaBlockResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAreaBlockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAreaBlockResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockResponseBody) SetRequestId(v string) *ModifyWebAreaBlockResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebAreaBlockResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebAreaBlockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebAreaBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAreaBlockResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockResponse) SetHeaders(v map[string]*string) *ModifyWebAreaBlockResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAreaBlockResponse) SetBody(v *ModifyWebAreaBlockResponseBody) *ModifyWebAreaBlockResponse {
	s.Body = v
	return s
}

type ModifyWebAreaBlockSwitchRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyWebAreaBlockSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAreaBlockSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockSwitchRequest) SetSourceIp(v string) *ModifyWebAreaBlockSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchRequest) SetResourceGroupId(v string) *ModifyWebAreaBlockSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchRequest) SetDomain(v string) *ModifyWebAreaBlockSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchRequest) SetConfig(v string) *ModifyWebAreaBlockSwitchRequest {
	s.Config = &v
	return s
}

type ModifyWebAreaBlockSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAreaBlockSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAreaBlockSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockSwitchResponseBody) SetRequestId(v string) *ModifyWebAreaBlockSwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebAreaBlockSwitchResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebAreaBlockSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebAreaBlockSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebAreaBlockSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebAreaBlockSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebAreaBlockSwitchResponse) SetBody(v *ModifyWebAreaBlockSwitchResponseBody) *ModifyWebAreaBlockSwitchResponse {
	s.Body = v
	return s
}

type ModifyWebCacheCustomRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Rules           *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyWebCacheCustomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheCustomRuleRequest) SetSourceIp(v string) *ModifyWebCacheCustomRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebCacheCustomRuleRequest) SetResourceGroupId(v string) *ModifyWebCacheCustomRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCacheCustomRuleRequest) SetDomain(v string) *ModifyWebCacheCustomRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCacheCustomRuleRequest) SetRules(v string) *ModifyWebCacheCustomRuleRequest {
	s.Rules = &v
	return s
}

type ModifyWebCacheCustomRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCacheCustomRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheCustomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheCustomRuleResponseBody) SetRequestId(v string) *ModifyWebCacheCustomRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebCacheCustomRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebCacheCustomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebCacheCustomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheCustomRuleResponse) SetHeaders(v map[string]*string) *ModifyWebCacheCustomRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCacheCustomRuleResponse) SetBody(v *ModifyWebCacheCustomRuleResponseBody) *ModifyWebCacheCustomRuleResponse {
	s.Body = v
	return s
}

type ModifyWebCacheModeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyWebCacheModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheModeRequest) SetSourceIp(v string) *ModifyWebCacheModeRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebCacheModeRequest) SetResourceGroupId(v string) *ModifyWebCacheModeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCacheModeRequest) SetDomain(v string) *ModifyWebCacheModeRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCacheModeRequest) SetMode(v string) *ModifyWebCacheModeRequest {
	s.Mode = &v
	return s
}

type ModifyWebCacheModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCacheModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheModeResponseBody) SetRequestId(v string) *ModifyWebCacheModeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebCacheModeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebCacheModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebCacheModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheModeResponse) SetHeaders(v map[string]*string) *ModifyWebCacheModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCacheModeResponse) SetBody(v *ModifyWebCacheModeResponseBody) *ModifyWebCacheModeResponse {
	s.Body = v
	return s
}

type ModifyWebCacheSwitchRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enable          *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyWebCacheSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheSwitchRequest) SetSourceIp(v string) *ModifyWebCacheSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebCacheSwitchRequest) SetResourceGroupId(v string) *ModifyWebCacheSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCacheSwitchRequest) SetDomain(v string) *ModifyWebCacheSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCacheSwitchRequest) SetEnable(v int32) *ModifyWebCacheSwitchRequest {
	s.Enable = &v
	return s
}

type ModifyWebCacheSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCacheSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheSwitchResponseBody) SetRequestId(v string) *ModifyWebCacheSwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebCacheSwitchResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebCacheSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebCacheSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCacheSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebCacheSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCacheSwitchResponse) SetBody(v *ModifyWebCacheSwitchResponseBody) *ModifyWebCacheSwitchResponse {
	s.Body = v
	return s
}

type ModifyWebCCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Count           *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ModifyWebCCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCCRuleRequest) SetSourceIp(v string) *ModifyWebCCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetResourceGroupId(v string) *ModifyWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetDomain(v string) *ModifyWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetName(v string) *ModifyWebCCRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetAct(v string) *ModifyWebCCRuleRequest {
	s.Act = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetCount(v int32) *ModifyWebCCRuleRequest {
	s.Count = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetInterval(v int32) *ModifyWebCCRuleRequest {
	s.Interval = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetMode(v string) *ModifyWebCCRuleRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetTtl(v int32) *ModifyWebCCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyWebCCRuleRequest) SetUri(v string) *ModifyWebCCRuleRequest {
	s.Uri = &v
	return s
}

type ModifyWebCCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCCRuleResponseBody) SetRequestId(v string) *ModifyWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebCCRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebCCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCCRuleResponse) SetHeaders(v map[string]*string) *ModifyWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCCRuleResponse) SetBody(v *ModifyWebCCRuleResponseBody) *ModifyWebCCRuleResponse {
	s.Body = v
	return s
}

type ModifyWebIpSetSwitchRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyWebIpSetSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebIpSetSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebIpSetSwitchRequest) SetSourceIp(v string) *ModifyWebIpSetSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebIpSetSwitchRequest) SetResourceGroupId(v string) *ModifyWebIpSetSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebIpSetSwitchRequest) SetDomain(v string) *ModifyWebIpSetSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebIpSetSwitchRequest) SetConfig(v string) *ModifyWebIpSetSwitchRequest {
	s.Config = &v
	return s
}

type ModifyWebIpSetSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebIpSetSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebIpSetSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebIpSetSwitchResponseBody) SetRequestId(v string) *ModifyWebIpSetSwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebIpSetSwitchResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebIpSetSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebIpSetSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebIpSetSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebIpSetSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebIpSetSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebIpSetSwitchResponse) SetBody(v *ModifyWebIpSetSwitchResponseBody) *ModifyWebIpSetSwitchResponse {
	s.Body = v
	return s
}

type ModifyWebPreciseAccessRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Rules           *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	Expires         *int32  `json:"Expires,omitempty" xml:"Expires,omitempty"`
}

func (s ModifyWebPreciseAccessRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebPreciseAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessRuleRequest) SetSourceIp(v string) *ModifyWebPreciseAccessRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetResourceGroupId(v string) *ModifyWebPreciseAccessRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetDomain(v string) *ModifyWebPreciseAccessRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetRules(v string) *ModifyWebPreciseAccessRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetExpires(v int32) *ModifyWebPreciseAccessRuleRequest {
	s.Expires = &v
	return s
}

type ModifyWebPreciseAccessRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebPreciseAccessRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebPreciseAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessRuleResponseBody) SetRequestId(v string) *ModifyWebPreciseAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebPreciseAccessRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebPreciseAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebPreciseAccessRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebPreciseAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessRuleResponse) SetHeaders(v map[string]*string) *ModifyWebPreciseAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebPreciseAccessRuleResponse) SetBody(v *ModifyWebPreciseAccessRuleResponseBody) *ModifyWebPreciseAccessRuleResponse {
	s.Body = v
	return s
}

type ModifyWebPreciseAccessSwitchRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyWebPreciseAccessSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebPreciseAccessSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetSourceIp(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetResourceGroupId(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetDomain(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetConfig(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.Config = &v
	return s
}

type ModifyWebPreciseAccessSwitchResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebPreciseAccessSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebPreciseAccessSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessSwitchResponseBody) SetRequestId(v string) *ModifyWebPreciseAccessSwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebPreciseAccessSwitchResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebPreciseAccessSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebPreciseAccessSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebPreciseAccessSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessSwitchResponse) SetHeaders(v map[string]*string) *ModifyWebPreciseAccessSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebPreciseAccessSwitchResponse) SetBody(v *ModifyWebPreciseAccessSwitchResponseBody) *ModifyWebPreciseAccessSwitchResponse {
	s.Body = v
	return s
}

type ModifyWebRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyTypes      *string   `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty"`
	RsType          *int32    `json:"RsType,omitempty" xml:"RsType,omitempty"`
	HttpsExt        *string   `json:"HttpsExt,omitempty" xml:"HttpsExt,omitempty"`
	RealServers     []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s ModifyWebRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebRuleRequest) SetSourceIp(v string) *ModifyWebRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebRuleRequest) SetResourceGroupId(v string) *ModifyWebRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebRuleRequest) SetDomain(v string) *ModifyWebRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebRuleRequest) SetProxyTypes(v string) *ModifyWebRuleRequest {
	s.ProxyTypes = &v
	return s
}

func (s *ModifyWebRuleRequest) SetRsType(v int32) *ModifyWebRuleRequest {
	s.RsType = &v
	return s
}

func (s *ModifyWebRuleRequest) SetHttpsExt(v string) *ModifyWebRuleRequest {
	s.HttpsExt = &v
	return s
}

func (s *ModifyWebRuleRequest) SetRealServers(v []*string) *ModifyWebRuleRequest {
	s.RealServers = v
	return s
}

func (s *ModifyWebRuleRequest) SetInstanceIds(v []*string) *ModifyWebRuleRequest {
	s.InstanceIds = v
	return s
}

type ModifyWebRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebRuleResponseBody) SetRequestId(v string) *ModifyWebRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebRuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebRuleResponse) SetHeaders(v map[string]*string) *ModifyWebRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebRuleResponse) SetBody(v *ModifyWebRuleResponseBody) *ModifyWebRuleResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetSourceIp(v string) *ReleaseInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type SwitchSchedulerRuleRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType   *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SwitchData *string `json:"SwitchData,omitempty" xml:"SwitchData,omitempty"`
}

func (s SwitchSchedulerRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchSchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *SwitchSchedulerRuleRequest) SetSourceIp(v string) *SwitchSchedulerRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *SwitchSchedulerRuleRequest) SetRuleName(v string) *SwitchSchedulerRuleRequest {
	s.RuleName = &v
	return s
}

func (s *SwitchSchedulerRuleRequest) SetRuleType(v int32) *SwitchSchedulerRuleRequest {
	s.RuleType = &v
	return s
}

func (s *SwitchSchedulerRuleRequest) SetSwitchData(v string) *SwitchSchedulerRuleRequest {
	s.SwitchData = &v
	return s
}

type SwitchSchedulerRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchSchedulerRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchSchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSchedulerRuleResponseBody) SetRequestId(v string) *SwitchSchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

type SwitchSchedulerRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchSchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchSchedulerRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchSchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *SwitchSchedulerRuleResponse) SetHeaders(v map[string]*string) *SwitchSchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *SwitchSchedulerRuleResponse) SetBody(v *SwitchSchedulerRuleResponseBody) *SwitchSchedulerRuleResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddoscoo"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddAutoCcBlacklistWithOptions(request *AddAutoCcBlacklistRequest, runtime *util.RuntimeOptions) (_result *AddAutoCcBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAutoCcBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAutoCcBlacklist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAutoCcBlacklist(request *AddAutoCcBlacklistRequest) (_result *AddAutoCcBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAutoCcBlacklistResponse{}
	_body, _err := client.AddAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAutoCcWhitelistWithOptions(request *AddAutoCcWhitelistRequest, runtime *util.RuntimeOptions) (_result *AddAutoCcWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAutoCcWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAutoCcWhitelist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAutoCcWhitelist(request *AddAutoCcWhitelistRequest) (_result *AddAutoCcWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAutoCcWhitelistResponse{}
	_body, _err := client.AddAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateWebCertWithOptions(request *AssociateWebCertRequest, runtime *util.RuntimeOptions) (_result *AssociateWebCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateWebCertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateWebCert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateWebCert(request *AssociateWebCertRequest) (_result *AssociateWebCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateWebCertResponse{}
	_body, _err := client.AssociateWebCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachSceneDefenseObjectWithOptions(request *AttachSceneDefenseObjectRequest, runtime *util.RuntimeOptions) (_result *AttachSceneDefenseObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachSceneDefenseObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachSceneDefenseObject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachSceneDefenseObject(request *AttachSceneDefenseObjectRequest) (_result *AttachSceneDefenseObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachSceneDefenseObjectResponse{}
	_body, _err := client.AttachSceneDefenseObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigL7RsPolicyWithOptions(request *ConfigL7RsPolicyRequest, runtime *util.RuntimeOptions) (_result *ConfigL7RsPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigL7RsPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigL7RsPolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigL7RsPolicy(request *ConfigL7RsPolicyRequest) (_result *ConfigL7RsPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigL7RsPolicyResponse{}
	_body, _err := client.ConfigL7RsPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigNetworkRegionBlockWithOptions(request *ConfigNetworkRegionBlockRequest, runtime *util.RuntimeOptions) (_result *ConfigNetworkRegionBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigNetworkRegionBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigNetworkRegionBlock"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigNetworkRegionBlock(request *ConfigNetworkRegionBlockRequest) (_result *ConfigNetworkRegionBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigNetworkRegionBlockResponse{}
	_body, _err := client.ConfigNetworkRegionBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigNetworkRulesWithOptions(request *ConfigNetworkRulesRequest, runtime *util.RuntimeOptions) (_result *ConfigNetworkRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigNetworkRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigNetworkRules"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigNetworkRules(request *ConfigNetworkRulesRequest) (_result *ConfigNetworkRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigNetworkRulesResponse{}
	_body, _err := client.ConfigNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigWebCCTemplateWithOptions(request *ConfigWebCCTemplateRequest, runtime *util.RuntimeOptions) (_result *ConfigWebCCTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigWebCCTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigWebCCTemplate"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigWebCCTemplate(request *ConfigWebCCTemplateRequest) (_result *ConfigWebCCTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigWebCCTemplateResponse{}
	_body, _err := client.ConfigWebCCTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigWebIpSetWithOptions(request *ConfigWebIpSetRequest, runtime *util.RuntimeOptions) (_result *ConfigWebIpSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigWebIpSetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigWebIpSet"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigWebIpSet(request *ConfigWebIpSetRequest) (_result *ConfigWebIpSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigWebIpSetResponse{}
	_body, _err := client.ConfigWebIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAsyncTaskWithOptions(request *CreateAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAsyncTask"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsyncTask(request *CreateAsyncTaskRequest) (_result *CreateAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CreateAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNetworkRulesWithOptions(request *CreateNetworkRulesRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNetworkRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNetworkRules"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNetworkRules(request *CreateNetworkRulesRequest) (_result *CreateNetworkRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkRulesResponse{}
	_body, _err := client.CreateNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSceneDefensePolicyWithOptions(request *CreateSceneDefensePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateSceneDefensePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSceneDefensePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSceneDefensePolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSceneDefensePolicy(request *CreateSceneDefensePolicyRequest) (_result *CreateSceneDefensePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSceneDefensePolicyResponse{}
	_body, _err := client.CreateSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSchedulerRuleWithOptions(request *CreateSchedulerRuleRequest, runtime *util.RuntimeOptions) (_result *CreateSchedulerRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSchedulerRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSchedulerRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchedulerRule(request *CreateSchedulerRuleRequest) (_result *CreateSchedulerRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchedulerRuleResponse{}
	_body, _err := client.CreateSchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTagResourcesWithOptions(request *CreateTagResourcesRequest, runtime *util.RuntimeOptions) (_result *CreateTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTagResources"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTagResources(request *CreateTagResourcesRequest) (_result *CreateTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTagResourcesResponse{}
	_body, _err := client.CreateTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWebCCRuleWithOptions(request *CreateWebCCRuleRequest, runtime *util.RuntimeOptions) (_result *CreateWebCCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWebCCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWebCCRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWebCCRule(request *CreateWebCCRuleRequest) (_result *CreateWebCCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWebCCRuleResponse{}
	_body, _err := client.CreateWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWebRuleWithOptions(request *CreateWebRuleRequest, runtime *util.RuntimeOptions) (_result *CreateWebRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWebRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWebRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWebRule(request *CreateWebRuleRequest) (_result *CreateWebRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWebRuleResponse{}
	_body, _err := client.CreateWebRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAsyncTaskWithOptions(request *DeleteAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAsyncTask"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAsyncTask(request *DeleteAsyncTaskRequest) (_result *DeleteAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DeleteAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAutoCcBlacklistWithOptions(request *DeleteAutoCcBlacklistRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoCcBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAutoCcBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAutoCcBlacklist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAutoCcBlacklist(request *DeleteAutoCcBlacklistRequest) (_result *DeleteAutoCcBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoCcBlacklistResponse{}
	_body, _err := client.DeleteAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAutoCcWhitelistWithOptions(request *DeleteAutoCcWhitelistRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoCcWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAutoCcWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAutoCcWhitelist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAutoCcWhitelist(request *DeleteAutoCcWhitelistRequest) (_result *DeleteAutoCcWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoCcWhitelistResponse{}
	_body, _err := client.DeleteAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNetworkRuleWithOptions(request *DeleteNetworkRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNetworkRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNetworkRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNetworkRule(request *DeleteNetworkRuleRequest) (_result *DeleteNetworkRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkRuleResponse{}
	_body, _err := client.DeleteNetworkRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSceneDefensePolicyWithOptions(request *DeleteSceneDefensePolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteSceneDefensePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSceneDefensePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSceneDefensePolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSceneDefensePolicy(request *DeleteSceneDefensePolicyRequest) (_result *DeleteSceneDefensePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSceneDefensePolicyResponse{}
	_body, _err := client.DeleteSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSchedulerRuleWithOptions(request *DeleteSchedulerRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteSchedulerRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSchedulerRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSchedulerRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSchedulerRule(request *DeleteSchedulerRuleRequest) (_result *DeleteSchedulerRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSchedulerRuleResponse{}
	_body, _err := client.DeleteSchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagResourcesWithOptions(request *DeleteTagResourcesRequest, runtime *util.RuntimeOptions) (_result *DeleteTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTagResources"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTagResources(request *DeleteTagResourcesRequest) (_result *DeleteTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagResourcesResponse{}
	_body, _err := client.DeleteTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebCacheCustomRuleWithOptions(request *DeleteWebCacheCustomRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteWebCacheCustomRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWebCacheCustomRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWebCacheCustomRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebCacheCustomRule(request *DeleteWebCacheCustomRuleRequest) (_result *DeleteWebCacheCustomRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebCacheCustomRuleResponse{}
	_body, _err := client.DeleteWebCacheCustomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebCCRuleWithOptions(request *DeleteWebCCRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteWebCCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWebCCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWebCCRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebCCRule(request *DeleteWebCCRuleRequest) (_result *DeleteWebCCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebCCRuleResponse{}
	_body, _err := client.DeleteWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebPreciseAccessRuleWithOptions(request *DeleteWebPreciseAccessRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteWebPreciseAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWebPreciseAccessRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWebPreciseAccessRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebPreciseAccessRule(request *DeleteWebPreciseAccessRuleRequest) (_result *DeleteWebPreciseAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebPreciseAccessRuleResponse{}
	_body, _err := client.DeleteWebPreciseAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebRuleWithOptions(request *DeleteWebRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteWebRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWebRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWebRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebRule(request *DeleteWebRuleRequest) (_result *DeleteWebRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebRuleResponse{}
	_body, _err := client.DeleteWebRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAsyncTasksWithOptions(request *DescribeAsyncTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeAsyncTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAsyncTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAsyncTasks"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAsyncTasks(request *DescribeAsyncTasksRequest) (_result *DescribeAsyncTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAsyncTasksResponse{}
	_body, _err := client.DescribeAsyncTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoCcBlacklistWithOptions(request *DescribeAutoCcBlacklistRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoCcBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoCcBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoCcBlacklist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoCcBlacklist(request *DescribeAutoCcBlacklistRequest) (_result *DescribeAutoCcBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoCcBlacklistResponse{}
	_body, _err := client.DescribeAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoCcListCountWithOptions(request *DescribeAutoCcListCountRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoCcListCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoCcListCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoCcListCount"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoCcListCount(request *DescribeAutoCcListCountRequest) (_result *DescribeAutoCcListCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoCcListCountResponse{}
	_body, _err := client.DescribeAutoCcListCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoCcWhitelistWithOptions(request *DescribeAutoCcWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoCcWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoCcWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoCcWhitelist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoCcWhitelist(request *DescribeAutoCcWhitelistRequest) (_result *DescribeAutoCcWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoCcWhitelistResponse{}
	_body, _err := client.DescribeAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackSourceCidrWithOptions(request *DescribeBackSourceCidrRequest, runtime *util.RuntimeOptions) (_result *DescribeBackSourceCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackSourceCidr"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackSourceCidr(request *DescribeBackSourceCidrRequest) (_result *DescribeBackSourceCidrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DescribeBackSourceCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBlackholeStatusWithOptions(request *DescribeBlackholeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBlackholeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBlackholeStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBlackholeStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBlackholeStatus(request *DescribeBlackholeStatusRequest) (_result *DescribeBlackholeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBlackholeStatusResponse{}
	_body, _err := client.DescribeBlackholeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBlockStatusWithOptions(request *DescribeBlockStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBlockStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBlockStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBlockStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBlockStatus(request *DescribeBlockStatusRequest) (_result *DescribeBlockStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBlockStatusResponse{}
	_body, _err := client.DescribeBlockStatusWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCertsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCerts"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeCnameReusesWithOptions(request *DescribeCnameReusesRequest, runtime *util.RuntimeOptions) (_result *DescribeCnameReusesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCnameReusesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCnameReuses"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCnameReuses(request *DescribeCnameReusesRequest) (_result *DescribeCnameReusesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCnameReusesResponse{}
	_body, _err := client.DescribeCnameReusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDosAllEventListWithOptions(request *DescribeDDosAllEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeDDosAllEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDosAllEventListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDosAllEventList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDosAllEventList(request *DescribeDDosAllEventListRequest) (_result *DescribeDDosAllEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDosAllEventListResponse{}
	_body, _err := client.DescribeDDosAllEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDosEventAreaWithOptions(request *DescribeDDosEventAreaRequest, runtime *util.RuntimeOptions) (_result *DescribeDDosEventAreaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDosEventAreaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDosEventArea"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDosEventArea(request *DescribeDDosEventAreaRequest) (_result *DescribeDDosEventAreaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDosEventAreaResponse{}
	_body, _err := client.DescribeDDosEventAreaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDosEventAttackTypeWithOptions(request *DescribeDDosEventAttackTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeDDosEventAttackTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDosEventAttackTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDosEventAttackType"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDosEventAttackType(request *DescribeDDosEventAttackTypeRequest) (_result *DescribeDDosEventAttackTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDosEventAttackTypeResponse{}
	_body, _err := client.DescribeDDosEventAttackTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDosEventIspWithOptions(request *DescribeDDosEventIspRequest, runtime *util.RuntimeOptions) (_result *DescribeDDosEventIspResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDosEventIspResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDosEventIsp"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDosEventIsp(request *DescribeDDosEventIspRequest) (_result *DescribeDDosEventIspResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDosEventIspResponse{}
	_body, _err := client.DescribeDDosEventIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDosEventMaxWithOptions(request *DescribeDDosEventMaxRequest, runtime *util.RuntimeOptions) (_result *DescribeDDosEventMaxResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDosEventMaxResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDosEventMax"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDosEventMax(request *DescribeDDosEventMaxRequest) (_result *DescribeDDosEventMaxResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDosEventMaxResponse{}
	_body, _err := client.DescribeDDosEventMaxWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDoSEventsWithOptions(request *DescribeDDoSEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDoSEvents"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDoSEvents(request *DescribeDDoSEventsRequest) (_result *DescribeDDoSEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DescribeDDoSEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDosEventSrcIpWithOptions(request *DescribeDDosEventSrcIpRequest, runtime *util.RuntimeOptions) (_result *DescribeDDosEventSrcIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDosEventSrcIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDosEventSrcIp"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDosEventSrcIp(request *DescribeDDosEventSrcIpRequest) (_result *DescribeDDosEventSrcIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDosEventSrcIpResponse{}
	_body, _err := client.DescribeDDosEventSrcIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseCountStatisticsWithOptions(request *DescribeDefenseCountStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDefenseCountStatistics"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseCountStatistics(request *DescribeDefenseCountStatisticsRequest) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DescribeDefenseCountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseRecordsWithOptions(request *DescribeDefenseRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDefenseRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDefenseRecords"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseRecords(request *DescribeDefenseRecordsRequest) (_result *DescribeDefenseRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseRecordsResponse{}
	_body, _err := client.DescribeDefenseRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAttackEventsWithOptions(request *DescribeDomainAttackEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAttackEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainAttackEvents"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAttackEvents(request *DescribeDomainAttackEventsRequest) (_result *DescribeDomainAttackEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DescribeDomainAttackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainOverviewWithOptions(request *DescribeDomainOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainOverviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainOverview"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainOverview(request *DescribeDomainOverviewRequest) (_result *DescribeDomainOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainOverviewResponse{}
	_body, _err := client.DescribeDomainOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQPSListWithOptions(request *DescribeDomainQPSListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQPSListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainQPSListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainQPSList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQPSList(request *DescribeDomainQPSListRequest) (_result *DescribeDomainQPSListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQPSListResponse{}
	_body, _err := client.DescribeDomainQPSListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithCacheWithOptions(request *DescribeDomainQpsWithCacheRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainQpsWithCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainQpsWithCache"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithCache(request *DescribeDomainQpsWithCacheRequest) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsWithCacheResponse{}
	_body, _err := client.DescribeDomainQpsWithCacheWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomains"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDomainStatusCodeCountWithOptions(request *DescribeDomainStatusCodeCountRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainStatusCodeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainStatusCodeCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainStatusCodeCount"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainStatusCodeCount(request *DescribeDomainStatusCodeCountRequest) (_result *DescribeDomainStatusCodeCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainStatusCodeCountResponse{}
	_body, _err := client.DescribeDomainStatusCodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainStatusCodeListWithOptions(request *DescribeDomainStatusCodeListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainStatusCodeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainStatusCodeListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainStatusCodeList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainStatusCodeList(request *DescribeDomainStatusCodeListRequest) (_result *DescribeDomainStatusCodeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainStatusCodeListResponse{}
	_body, _err := client.DescribeDomainStatusCodeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainTopAttackListWithOptions(request *DescribeDomainTopAttackListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainTopAttackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainTopAttackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainTopAttackList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainTopAttackList(request *DescribeDomainTopAttackListRequest) (_result *DescribeDomainTopAttackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainTopAttackListResponse{}
	_body, _err := client.DescribeDomainTopAttackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainViewSourceCountriesWithOptions(request *DescribeDomainViewSourceCountriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainViewSourceCountriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainViewSourceCountriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainViewSourceCountries"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainViewSourceCountries(request *DescribeDomainViewSourceCountriesRequest) (_result *DescribeDomainViewSourceCountriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainViewSourceCountriesResponse{}
	_body, _err := client.DescribeDomainViewSourceCountriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainViewSourceProvincesWithOptions(request *DescribeDomainViewSourceProvincesRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainViewSourceProvincesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainViewSourceProvincesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainViewSourceProvinces"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainViewSourceProvinces(request *DescribeDomainViewSourceProvincesRequest) (_result *DescribeDomainViewSourceProvincesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainViewSourceProvincesResponse{}
	_body, _err := client.DescribeDomainViewSourceProvincesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainViewTopCostTimeWithOptions(request *DescribeDomainViewTopCostTimeRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainViewTopCostTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainViewTopCostTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainViewTopCostTime"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainViewTopCostTime(request *DescribeDomainViewTopCostTimeRequest) (_result *DescribeDomainViewTopCostTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainViewTopCostTimeResponse{}
	_body, _err := client.DescribeDomainViewTopCostTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainViewTopUrlWithOptions(request *DescribeDomainViewTopUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainViewTopUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainViewTopUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainViewTopUrl"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainViewTopUrl(request *DescribeDomainViewTopUrlRequest) (_result *DescribeDomainViewTopUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainViewTopUrlResponse{}
	_body, _err := client.DescribeDomainViewTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticBandwidthSpecWithOptions(request *DescribeElasticBandwidthSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeElasticBandwidthSpec"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticBandwidthSpec(request *DescribeElasticBandwidthSpecRequest) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DescribeElasticBandwidthSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckListWithOptions(request *DescribeHealthCheckListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHealthCheckList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckList(request *DescribeHealthCheckListRequest) (_result *DescribeHealthCheckListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DescribeHealthCheckListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckStatusWithOptions(request *DescribeHealthCheckStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHealthCheckStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHealthCheckStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckStatus(request *DescribeHealthCheckStatusRequest) (_result *DescribeHealthCheckStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckStatusResponse{}
	_body, _err := client.DescribeHealthCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceDetailsWithOptions(request *DescribeInstanceDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceDetails"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceDetails(request *DescribeInstanceDetailsRequest) (_result *DescribeInstanceDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DescribeInstanceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceIdsWithOptions(request *DescribeInstanceIdsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceIds"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceIds(request *DescribeInstanceIdsRequest) (_result *DescribeInstanceIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceIdsResponse{}
	_body, _err := client.DescribeInstanceIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSpecs"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStatistics"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (_result *DescribeInstanceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DescribeInstanceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStatusWithOptions(request *DescribeInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStatus(request *DescribeInstanceStatusRequest) (_result *DescribeInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStatusResponse{}
	_body, _err := client.DescribeInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeL7RsPolicyWithOptions(request *DescribeL7RsPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeL7RsPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeL7RsPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeL7RsPolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeL7RsPolicy(request *DescribeL7RsPolicyRequest) (_result *DescribeL7RsPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeL7RsPolicyResponse{}
	_body, _err := client.DescribeL7RsPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogStoreExistStatusWithOptions(request *DescribeLogStoreExistStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogStoreExistStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogStoreExistStatus(request *DescribeLogStoreExistStatusRequest) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DescribeLogStoreExistStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkRegionBlockWithOptions(request *DescribeNetworkRegionBlockRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkRegionBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworkRegionBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworkRegionBlock"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkRegionBlock(request *DescribeNetworkRegionBlockRequest) (_result *DescribeNetworkRegionBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkRegionBlockResponse{}
	_body, _err := client.DescribeNetworkRegionBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkRuleAttributesWithOptions(request *DescribeNetworkRuleAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkRuleAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworkRuleAttributesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworkRuleAttributes"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkRuleAttributes(request *DescribeNetworkRuleAttributesRequest) (_result *DescribeNetworkRuleAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkRuleAttributesResponse{}
	_body, _err := client.DescribeNetworkRuleAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkRulesWithOptions(request *DescribeNetworkRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworkRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworkRules"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkRules(request *DescribeNetworkRulesRequest) (_result *DescribeNetworkRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkRulesResponse{}
	_body, _err := client.DescribeNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOpEntities"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortAttackMaxFlowWithOptions(request *DescribePortAttackMaxFlowRequest, runtime *util.RuntimeOptions) (_result *DescribePortAttackMaxFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortAttackMaxFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortAttackMaxFlow"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortAttackMaxFlow(request *DescribePortAttackMaxFlowRequest) (_result *DescribePortAttackMaxFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortAttackMaxFlowResponse{}
	_body, _err := client.DescribePortAttackMaxFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortAutoCcStatusWithOptions(request *DescribePortAutoCcStatusRequest, runtime *util.RuntimeOptions) (_result *DescribePortAutoCcStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortAutoCcStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortAutoCcStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortAutoCcStatus(request *DescribePortAutoCcStatusRequest) (_result *DescribePortAutoCcStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortAutoCcStatusResponse{}
	_body, _err := client.DescribePortAutoCcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortConnsCountWithOptions(request *DescribePortConnsCountRequest, runtime *util.RuntimeOptions) (_result *DescribePortConnsCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortConnsCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortConnsCount"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortConnsCount(request *DescribePortConnsCountRequest) (_result *DescribePortConnsCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortConnsCountResponse{}
	_body, _err := client.DescribePortConnsCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortConnsListWithOptions(request *DescribePortConnsListRequest, runtime *util.RuntimeOptions) (_result *DescribePortConnsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortConnsListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortConnsList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortConnsList(request *DescribePortConnsListRequest) (_result *DescribePortConnsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortConnsListResponse{}
	_body, _err := client.DescribePortConnsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortFlowListWithOptions(request *DescribePortFlowListRequest, runtime *util.RuntimeOptions) (_result *DescribePortFlowListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortFlowListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortFlowList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortFlowList(request *DescribePortFlowListRequest) (_result *DescribePortFlowListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortFlowListResponse{}
	_body, _err := client.DescribePortFlowListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortMaxConnsWithOptions(request *DescribePortMaxConnsRequest, runtime *util.RuntimeOptions) (_result *DescribePortMaxConnsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortMaxConnsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortMaxConns"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortMaxConns(request *DescribePortMaxConnsRequest) (_result *DescribePortMaxConnsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortMaxConnsResponse{}
	_body, _err := client.DescribePortMaxConnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortViewSourceCountriesWithOptions(request *DescribePortViewSourceCountriesRequest, runtime *util.RuntimeOptions) (_result *DescribePortViewSourceCountriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortViewSourceCountriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortViewSourceCountries"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortViewSourceCountries(request *DescribePortViewSourceCountriesRequest) (_result *DescribePortViewSourceCountriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortViewSourceCountriesResponse{}
	_body, _err := client.DescribePortViewSourceCountriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortViewSourceIspsWithOptions(request *DescribePortViewSourceIspsRequest, runtime *util.RuntimeOptions) (_result *DescribePortViewSourceIspsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortViewSourceIspsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortViewSourceIsps"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortViewSourceIsps(request *DescribePortViewSourceIspsRequest) (_result *DescribePortViewSourceIspsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortViewSourceIspsResponse{}
	_body, _err := client.DescribePortViewSourceIspsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortViewSourceProvincesWithOptions(request *DescribePortViewSourceProvincesRequest, runtime *util.RuntimeOptions) (_result *DescribePortViewSourceProvincesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortViewSourceProvincesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortViewSourceProvinces"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortViewSourceProvinces(request *DescribePortViewSourceProvincesRequest) (_result *DescribePortViewSourceProvincesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortViewSourceProvincesResponse{}
	_body, _err := client.DescribePortViewSourceProvincesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSceneDefenseObjectsWithOptions(request *DescribeSceneDefenseObjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeSceneDefenseObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSceneDefenseObjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSceneDefenseObjects"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSceneDefenseObjects(request *DescribeSceneDefenseObjectsRequest) (_result *DescribeSceneDefenseObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSceneDefenseObjectsResponse{}
	_body, _err := client.DescribeSceneDefenseObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSceneDefensePoliciesWithOptions(request *DescribeSceneDefensePoliciesRequest, runtime *util.RuntimeOptions) (_result *DescribeSceneDefensePoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSceneDefensePoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSceneDefensePolicies"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSceneDefensePolicies(request *DescribeSceneDefensePoliciesRequest) (_result *DescribeSceneDefensePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSceneDefensePoliciesResponse{}
	_body, _err := client.DescribeSceneDefensePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSchedulerRulesWithOptions(request *DescribeSchedulerRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeSchedulerRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSchedulerRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSchedulerRules"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSchedulerRules(request *DescribeSchedulerRulesRequest) (_result *DescribeSchedulerRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSchedulerRulesResponse{}
	_body, _err := client.DescribeSchedulerRulesWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsAuthStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSlsLogstoreInfoWithOptions(request *DescribeSlsLogstoreInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsLogstoreInfo"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsLogstoreInfo(request *DescribeSlsLogstoreInfoRequest) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DescribeSlsLogstoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsOpenStatusWithOptions(request *DescribeSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsOpenStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsOpenStatus(request *DescribeSlsOpenStatusRequest) (_result *DescribeSlsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DescribeSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStsGrantStatusWithOptions(request *DescribeStsGrantStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeStsGrantStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStsGrantStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStsGrantStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStsGrantStatus(request *DescribeStsGrantStatusRequest) (_result *DescribeStsGrantStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStsGrantStatusResponse{}
	_body, _err := client.DescribeStsGrantStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagKeysWithOptions(request *DescribeTagKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTagKeys"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTagKeys(request *DescribeTagKeysRequest) (_result *DescribeTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagKeysResponse{}
	_body, _err := client.DescribeTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagResourcesWithOptions(request *DescribeTagResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTagResources"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTagResources(request *DescribeTagResourcesRequest) (_result *DescribeTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagResourcesResponse{}
	_body, _err := client.DescribeTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUnBlackholeCountWithOptions(request *DescribeUnBlackholeCountRequest, runtime *util.RuntimeOptions) (_result *DescribeUnBlackholeCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUnBlackholeCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUnBlackholeCount"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUnBlackholeCount(request *DescribeUnBlackholeCountRequest) (_result *DescribeUnBlackholeCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUnBlackholeCountResponse{}
	_body, _err := client.DescribeUnBlackholeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUnBlockCountWithOptions(request *DescribeUnBlockCountRequest, runtime *util.RuntimeOptions) (_result *DescribeUnBlockCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUnBlockCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUnBlockCount"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUnBlockCount(request *DescribeUnBlockCountRequest) (_result *DescribeUnBlockCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUnBlockCountResponse{}
	_body, _err := client.DescribeUnBlockCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebAccessLogDispatchStatusWithOptions(request *DescribeWebAccessLogDispatchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeWebAccessLogDispatchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebAccessLogDispatchStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebAccessLogDispatchStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebAccessLogDispatchStatus(request *DescribeWebAccessLogDispatchStatusRequest) (_result *DescribeWebAccessLogDispatchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebAccessLogDispatchStatusResponse{}
	_body, _err := client.DescribeWebAccessLogDispatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebAccessLogEmptyCountWithOptions(request *DescribeWebAccessLogEmptyCountRequest, runtime *util.RuntimeOptions) (_result *DescribeWebAccessLogEmptyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebAccessLogEmptyCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebAccessLogEmptyCount"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebAccessLogEmptyCount(request *DescribeWebAccessLogEmptyCountRequest) (_result *DescribeWebAccessLogEmptyCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebAccessLogEmptyCountResponse{}
	_body, _err := client.DescribeWebAccessLogEmptyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebAccessLogStatusWithOptions(request *DescribeWebAccessLogStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeWebAccessLogStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebAccessLogStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebAccessLogStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebAccessLogStatus(request *DescribeWebAccessLogStatusRequest) (_result *DescribeWebAccessLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebAccessLogStatusResponse{}
	_body, _err := client.DescribeWebAccessLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebAccessModeWithOptions(request *DescribeWebAccessModeRequest, runtime *util.RuntimeOptions) (_result *DescribeWebAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebAccessModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebAccessMode"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebAccessMode(request *DescribeWebAccessModeRequest) (_result *DescribeWebAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebAccessModeResponse{}
	_body, _err := client.DescribeWebAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebAreaBlockConfigsWithOptions(request *DescribeWebAreaBlockConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebAreaBlockConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebAreaBlockConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebAreaBlockConfigs"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebAreaBlockConfigs(request *DescribeWebAreaBlockConfigsRequest) (_result *DescribeWebAreaBlockConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebAreaBlockConfigsResponse{}
	_body, _err := client.DescribeWebAreaBlockConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebCacheConfigsWithOptions(request *DescribeWebCacheConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebCacheConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebCacheConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebCacheConfigs"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebCacheConfigs(request *DescribeWebCacheConfigsRequest) (_result *DescribeWebCacheConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebCacheConfigsResponse{}
	_body, _err := client.DescribeWebCacheConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebCcProtectSwitchWithOptions(request *DescribeWebCcProtectSwitchRequest, runtime *util.RuntimeOptions) (_result *DescribeWebCcProtectSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebCcProtectSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebCcProtectSwitch"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebCcProtectSwitch(request *DescribeWebCcProtectSwitchRequest) (_result *DescribeWebCcProtectSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebCcProtectSwitchResponse{}
	_body, _err := client.DescribeWebCcProtectSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebCCRulesWithOptions(request *DescribeWebCCRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeWebCCRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebCCRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebCCRules"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebCCRules(request *DescribeWebCCRulesRequest) (_result *DescribeWebCCRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebCCRulesResponse{}
	_body, _err := client.DescribeWebCCRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebCustomPortsWithOptions(request *DescribeWebCustomPortsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebCustomPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebCustomPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebCustomPorts"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebCustomPorts(request *DescribeWebCustomPortsRequest) (_result *DescribeWebCustomPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebCustomPortsResponse{}
	_body, _err := client.DescribeWebCustomPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebInstanceRelationsWithOptions(request *DescribeWebInstanceRelationsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebInstanceRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebInstanceRelationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebInstanceRelations"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebInstanceRelations(request *DescribeWebInstanceRelationsRequest) (_result *DescribeWebInstanceRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebInstanceRelationsResponse{}
	_body, _err := client.DescribeWebInstanceRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebPreciseAccessRuleWithOptions(request *DescribeWebPreciseAccessRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeWebPreciseAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebPreciseAccessRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebPreciseAccessRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebPreciseAccessRule(request *DescribeWebPreciseAccessRuleRequest) (_result *DescribeWebPreciseAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebPreciseAccessRuleResponse{}
	_body, _err := client.DescribeWebPreciseAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebRulesWithOptions(request *DescribeWebRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeWebRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebRules"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebRules(request *DescribeWebRulesRequest) (_result *DescribeWebRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebRulesResponse{}
	_body, _err := client.DescribeWebRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachSceneDefenseObjectWithOptions(request *DetachSceneDefenseObjectRequest, runtime *util.RuntimeOptions) (_result *DetachSceneDefenseObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachSceneDefenseObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachSceneDefenseObject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachSceneDefenseObject(request *DetachSceneDefenseObjectRequest) (_result *DetachSceneDefenseObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachSceneDefenseObjectResponse{}
	_body, _err := client.DetachSceneDefenseObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSceneDefensePolicyWithOptions(request *DisableSceneDefensePolicyRequest, runtime *util.RuntimeOptions) (_result *DisableSceneDefensePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableSceneDefensePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableSceneDefensePolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSceneDefensePolicy(request *DisableSceneDefensePolicyRequest) (_result *DisableSceneDefensePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSceneDefensePolicyResponse{}
	_body, _err := client.DisableSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableWebAccessLogConfigWithOptions(request *DisableWebAccessLogConfigRequest, runtime *util.RuntimeOptions) (_result *DisableWebAccessLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableWebAccessLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableWebAccessLogConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableWebAccessLogConfig(request *DisableWebAccessLogConfigRequest) (_result *DisableWebAccessLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableWebAccessLogConfigResponse{}
	_body, _err := client.DisableWebAccessLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableWebCCWithOptions(request *DisableWebCCRequest, runtime *util.RuntimeOptions) (_result *DisableWebCCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableWebCCResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableWebCC"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableWebCC(request *DisableWebCCRequest) (_result *DisableWebCCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableWebCCResponse{}
	_body, _err := client.DisableWebCCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableWebCCRuleWithOptions(request *DisableWebCCRuleRequest, runtime *util.RuntimeOptions) (_result *DisableWebCCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableWebCCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableWebCCRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableWebCCRule(request *DisableWebCCRuleRequest) (_result *DisableWebCCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableWebCCRuleResponse{}
	_body, _err := client.DisableWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EmptyAutoCcBlacklistWithOptions(request *EmptyAutoCcBlacklistRequest, runtime *util.RuntimeOptions) (_result *EmptyAutoCcBlacklistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EmptyAutoCcBlacklistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EmptyAutoCcBlacklist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EmptyAutoCcBlacklist(request *EmptyAutoCcBlacklistRequest) (_result *EmptyAutoCcBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EmptyAutoCcBlacklistResponse{}
	_body, _err := client.EmptyAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EmptyAutoCcWhitelistWithOptions(request *EmptyAutoCcWhitelistRequest, runtime *util.RuntimeOptions) (_result *EmptyAutoCcWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EmptyAutoCcWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EmptyAutoCcWhitelist"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EmptyAutoCcWhitelist(request *EmptyAutoCcWhitelistRequest) (_result *EmptyAutoCcWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EmptyAutoCcWhitelistResponse{}
	_body, _err := client.EmptyAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EmptySlsLogstoreWithOptions(request *EmptySlsLogstoreRequest, runtime *util.RuntimeOptions) (_result *EmptySlsLogstoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EmptySlsLogstore"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EmptySlsLogstore(request *EmptySlsLogstoreRequest) (_result *EmptySlsLogstoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.EmptySlsLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSceneDefensePolicyWithOptions(request *EnableSceneDefensePolicyRequest, runtime *util.RuntimeOptions) (_result *EnableSceneDefensePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableSceneDefensePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableSceneDefensePolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSceneDefensePolicy(request *EnableSceneDefensePolicyRequest) (_result *EnableSceneDefensePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSceneDefensePolicyResponse{}
	_body, _err := client.EnableSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableWebAccessLogConfigWithOptions(request *EnableWebAccessLogConfigRequest, runtime *util.RuntimeOptions) (_result *EnableWebAccessLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableWebAccessLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableWebAccessLogConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableWebAccessLogConfig(request *EnableWebAccessLogConfigRequest) (_result *EnableWebAccessLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableWebAccessLogConfigResponse{}
	_body, _err := client.EnableWebAccessLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableWebCCWithOptions(request *EnableWebCCRequest, runtime *util.RuntimeOptions) (_result *EnableWebCCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableWebCCResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableWebCC"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableWebCC(request *EnableWebCCRequest) (_result *EnableWebCCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableWebCCResponse{}
	_body, _err := client.EnableWebCCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableWebCCRuleWithOptions(request *EnableWebCCRuleRequest, runtime *util.RuntimeOptions) (_result *EnableWebCCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableWebCCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableWebCCRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableWebCCRule(request *EnableWebCCRuleRequest) (_result *EnableWebCCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableWebCCRuleResponse{}
	_body, _err := client.EnableWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBlackholeStatusWithOptions(request *ModifyBlackholeStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyBlackholeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBlackholeStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBlackholeStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBlackholeStatus(request *ModifyBlackholeStatusRequest) (_result *ModifyBlackholeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBlackholeStatusResponse{}
	_body, _err := client.ModifyBlackholeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBlockStatusWithOptions(request *ModifyBlockStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyBlockStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBlockStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBlockStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBlockStatus(request *ModifyBlockStatusRequest) (_result *ModifyBlockStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBlockStatusResponse{}
	_body, _err := client.ModifyBlockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCnameReuseWithOptions(request *ModifyCnameReuseRequest, runtime *util.RuntimeOptions) (_result *ModifyCnameReuseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCnameReuseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCnameReuse"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCnameReuse(request *ModifyCnameReuseRequest) (_result *ModifyCnameReuseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCnameReuseResponse{}
	_body, _err := client.ModifyCnameReuseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElasticBandWidthWithOptions(request *ModifyElasticBandWidthRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticBandWidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyElasticBandWidth"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElasticBandWidth(request *ModifyElasticBandWidthRequest) (_result *ModifyElasticBandWidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.ModifyElasticBandWidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFullLogTtlWithOptions(request *ModifyFullLogTtlRequest, runtime *util.RuntimeOptions) (_result *ModifyFullLogTtlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFullLogTtl"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFullLogTtl(request *ModifyFullLogTtlRequest) (_result *ModifyFullLogTtlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.ModifyFullLogTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHealthCheckConfigWithOptions(request *ModifyHealthCheckConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyHealthCheckConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHealthCheckConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHealthCheckConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHealthCheckConfig(request *ModifyHealthCheckConfigRequest) (_result *ModifyHealthCheckConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHealthCheckConfigResponse{}
	_body, _err := client.ModifyHealthCheckConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHttp2EnableWithOptions(request *ModifyHttp2EnableRequest, runtime *util.RuntimeOptions) (_result *ModifyHttp2EnableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHttp2EnableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHttp2Enable"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHttp2Enable(request *ModifyHttp2EnableRequest) (_result *ModifyHttp2EnableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHttp2EnableResponse{}
	_body, _err := client.ModifyHttp2EnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceRemarkWithOptions(request *ModifyInstanceRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceRemark"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceRemark(request *ModifyInstanceRemarkRequest) (_result *ModifyInstanceRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.ModifyInstanceRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkRuleAttributeWithOptions(request *ModifyNetworkRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNetworkRuleAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNetworkRuleAttribute"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkRuleAttribute(request *ModifyNetworkRuleAttributeRequest) (_result *ModifyNetworkRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkRuleAttributeResponse{}
	_body, _err := client.ModifyNetworkRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPortAutoCcStatusWithOptions(request *ModifyPortAutoCcStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyPortAutoCcStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPortAutoCcStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPortAutoCcStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPortAutoCcStatus(request *ModifyPortAutoCcStatusRequest) (_result *ModifyPortAutoCcStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPortAutoCcStatusResponse{}
	_body, _err := client.ModifyPortAutoCcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySceneDefensePolicyWithOptions(request *ModifySceneDefensePolicyRequest, runtime *util.RuntimeOptions) (_result *ModifySceneDefensePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySceneDefensePolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySceneDefensePolicy"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySceneDefensePolicy(request *ModifySceneDefensePolicyRequest) (_result *ModifySceneDefensePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySceneDefensePolicyResponse{}
	_body, _err := client.ModifySceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySchedulerRuleWithOptions(request *ModifySchedulerRuleRequest, runtime *util.RuntimeOptions) (_result *ModifySchedulerRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySchedulerRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySchedulerRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySchedulerRule(request *ModifySchedulerRuleRequest) (_result *ModifySchedulerRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySchedulerRuleResponse{}
	_body, _err := client.ModifySchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTlsConfigWithOptions(request *ModifyTlsConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyTlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyTlsConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyTlsConfig"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTlsConfig(request *ModifyTlsConfigRequest) (_result *ModifyTlsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTlsConfigResponse{}
	_body, _err := client.ModifyTlsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebAccessModeWithOptions(request *ModifyWebAccessModeRequest, runtime *util.RuntimeOptions) (_result *ModifyWebAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebAccessModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebAccessMode"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebAccessMode(request *ModifyWebAccessModeRequest) (_result *ModifyWebAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebAccessModeResponse{}
	_body, _err := client.ModifyWebAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebAIProtectModeWithOptions(request *ModifyWebAIProtectModeRequest, runtime *util.RuntimeOptions) (_result *ModifyWebAIProtectModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebAIProtectModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebAIProtectMode"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebAIProtectMode(request *ModifyWebAIProtectModeRequest) (_result *ModifyWebAIProtectModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebAIProtectModeResponse{}
	_body, _err := client.ModifyWebAIProtectModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebAIProtectSwitchWithOptions(request *ModifyWebAIProtectSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWebAIProtectSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebAIProtectSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebAIProtectSwitch"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebAIProtectSwitch(request *ModifyWebAIProtectSwitchRequest) (_result *ModifyWebAIProtectSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebAIProtectSwitchResponse{}
	_body, _err := client.ModifyWebAIProtectSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebAreaBlockWithOptions(request *ModifyWebAreaBlockRequest, runtime *util.RuntimeOptions) (_result *ModifyWebAreaBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebAreaBlockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebAreaBlock"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebAreaBlock(request *ModifyWebAreaBlockRequest) (_result *ModifyWebAreaBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebAreaBlockResponse{}
	_body, _err := client.ModifyWebAreaBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebAreaBlockSwitchWithOptions(request *ModifyWebAreaBlockSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWebAreaBlockSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebAreaBlockSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebAreaBlockSwitch"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebAreaBlockSwitch(request *ModifyWebAreaBlockSwitchRequest) (_result *ModifyWebAreaBlockSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebAreaBlockSwitchResponse{}
	_body, _err := client.ModifyWebAreaBlockSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebCacheCustomRuleWithOptions(request *ModifyWebCacheCustomRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyWebCacheCustomRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebCacheCustomRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebCacheCustomRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebCacheCustomRule(request *ModifyWebCacheCustomRuleRequest) (_result *ModifyWebCacheCustomRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebCacheCustomRuleResponse{}
	_body, _err := client.ModifyWebCacheCustomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebCacheModeWithOptions(request *ModifyWebCacheModeRequest, runtime *util.RuntimeOptions) (_result *ModifyWebCacheModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebCacheModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebCacheMode"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebCacheMode(request *ModifyWebCacheModeRequest) (_result *ModifyWebCacheModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebCacheModeResponse{}
	_body, _err := client.ModifyWebCacheModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebCacheSwitchWithOptions(request *ModifyWebCacheSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWebCacheSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebCacheSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebCacheSwitch"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebCacheSwitch(request *ModifyWebCacheSwitchRequest) (_result *ModifyWebCacheSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebCacheSwitchResponse{}
	_body, _err := client.ModifyWebCacheSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebCCRuleWithOptions(request *ModifyWebCCRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyWebCCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebCCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebCCRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebCCRule(request *ModifyWebCCRuleRequest) (_result *ModifyWebCCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebCCRuleResponse{}
	_body, _err := client.ModifyWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebIpSetSwitchWithOptions(request *ModifyWebIpSetSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWebIpSetSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebIpSetSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebIpSetSwitch"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebIpSetSwitch(request *ModifyWebIpSetSwitchRequest) (_result *ModifyWebIpSetSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebIpSetSwitchResponse{}
	_body, _err := client.ModifyWebIpSetSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebPreciseAccessRuleWithOptions(request *ModifyWebPreciseAccessRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyWebPreciseAccessRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebPreciseAccessRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebPreciseAccessRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebPreciseAccessRule(request *ModifyWebPreciseAccessRuleRequest) (_result *ModifyWebPreciseAccessRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebPreciseAccessRuleResponse{}
	_body, _err := client.ModifyWebPreciseAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebPreciseAccessSwitchWithOptions(request *ModifyWebPreciseAccessSwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyWebPreciseAccessSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebPreciseAccessSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebPreciseAccessSwitch"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebPreciseAccessSwitch(request *ModifyWebPreciseAccessSwitchRequest) (_result *ModifyWebPreciseAccessSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebPreciseAccessSwitchResponse{}
	_body, _err := client.ModifyWebPreciseAccessSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebRuleWithOptions(request *ModifyWebRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyWebRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebRule(request *ModifyWebRuleRequest) (_result *ModifyWebRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebRuleResponse{}
	_body, _err := client.ModifyWebRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchSchedulerRuleWithOptions(request *SwitchSchedulerRuleRequest, runtime *util.RuntimeOptions) (_result *SwitchSchedulerRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchSchedulerRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchSchedulerRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchSchedulerRule(request *SwitchSchedulerRuleRequest) (_result *SwitchSchedulerRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchSchedulerRuleResponse{}
	_body, _err := client.SwitchSchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
