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

type AddCdnDomainRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	CdnType         *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *AddCdnDomainRequest) SetOwnerId(v int64) *AddCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCdnDomainRequest) SetOwnerAccount(v string) *AddCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddCdnDomainRequest) SetSecurityToken(v string) *AddCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddCdnDomainRequest) SetCdnType(v string) *AddCdnDomainRequest {
	s.CdnType = &v
	return s
}

func (s *AddCdnDomainRequest) SetDomainName(v string) *AddCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddCdnDomainRequest) SetResourceGroupId(v string) *AddCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddCdnDomainRequest) SetSources(v string) *AddCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddCdnDomainRequest) SetCheckUrl(v string) *AddCdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddCdnDomainRequest) SetScope(v string) *AddCdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddCdnDomainRequest) SetTopLevelDomain(v string) *AddCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type AddCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddCdnDomainResponseBody) SetRequestId(v string) *AddCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddCdnDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *AddCdnDomainResponse) SetHeaders(v map[string]*string) *AddCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *AddCdnDomainResponse) SetBody(v *AddCdnDomainResponseBody) *AddCdnDomainResponse {
	s.Body = v
	return s
}

type AddFCTriggerRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TriggerARN       *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
	EventMetaName    *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
	SourceARN        *string `json:"SourceARN,omitempty" xml:"SourceARN,omitempty"`
	FunctionARN      *string `json:"FunctionARN,omitempty" xml:"FunctionARN,omitempty"`
	RoleARN          *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	Notes            *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
}

func (s AddFCTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *AddFCTriggerRequest) SetOwnerId(v int64) *AddFCTriggerRequest {
	s.OwnerId = &v
	return s
}

func (s *AddFCTriggerRequest) SetTriggerARN(v string) *AddFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetEventMetaName(v string) *AddFCTriggerRequest {
	s.EventMetaName = &v
	return s
}

func (s *AddFCTriggerRequest) SetEventMetaVersion(v string) *AddFCTriggerRequest {
	s.EventMetaVersion = &v
	return s
}

func (s *AddFCTriggerRequest) SetSourceARN(v string) *AddFCTriggerRequest {
	s.SourceARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetFunctionARN(v string) *AddFCTriggerRequest {
	s.FunctionARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetRoleARN(v string) *AddFCTriggerRequest {
	s.RoleARN = &v
	return s
}

func (s *AddFCTriggerRequest) SetNotes(v string) *AddFCTriggerRequest {
	s.Notes = &v
	return s
}

type AddFCTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFCTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *AddFCTriggerResponseBody) SetRequestId(v string) *AddFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

type AddFCTriggerResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddFCTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *AddFCTriggerResponse) SetHeaders(v map[string]*string) *AddFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *AddFCTriggerResponse) SetBody(v *AddFCTriggerResponseBody) *AddFCTriggerResponse {
	s.Body = v
	return s
}

type BatchAddCdnDomainRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	CdnType         *string `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchAddCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchAddCdnDomainRequest) SetOwnerId(v int64) *BatchAddCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetOwnerAccount(v string) *BatchAddCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetSecurityToken(v string) *BatchAddCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetCdnType(v string) *BatchAddCdnDomainRequest {
	s.CdnType = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetDomainName(v string) *BatchAddCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetResourceGroupId(v string) *BatchAddCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetSources(v string) *BatchAddCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetCheckUrl(v string) *BatchAddCdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetScope(v string) *BatchAddCdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *BatchAddCdnDomainRequest) SetTopLevelDomain(v string) *BatchAddCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type BatchAddCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAddCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddCdnDomainResponseBody) SetRequestId(v string) *BatchAddCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchAddCdnDomainResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAddCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAddCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchAddCdnDomainResponse) SetHeaders(v map[string]*string) *BatchAddCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchAddCdnDomainResponse) SetBody(v *BatchAddCdnDomainResponseBody) *BatchAddCdnDomainResponse {
	s.Body = v
	return s
}

type BatchSetCdnDomainConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
}

func (s BatchSetCdnDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetCdnDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigRequest) SetOwnerId(v int64) *BatchSetCdnDomainConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetOwnerAccount(v string) *BatchSetCdnDomainConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetSecurityToken(v string) *BatchSetCdnDomainConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetDomainNames(v string) *BatchSetCdnDomainConfigRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetCdnDomainConfigRequest) SetFunctions(v string) *BatchSetCdnDomainConfigRequest {
	s.Functions = &v
	return s
}

type BatchSetCdnDomainConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetCdnDomainConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetCdnDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigResponseBody) SetRequestId(v string) *BatchSetCdnDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetCdnDomainConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetCdnDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetCdnDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetCdnDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainConfigResponse) SetHeaders(v map[string]*string) *BatchSetCdnDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchSetCdnDomainConfigResponse) SetBody(v *BatchSetCdnDomainConfigResponseBody) *BatchSetCdnDomainConfigResponse {
	s.Body = v
	return s
}

type BatchSetCdnDomainServerCertificateRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType      *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SSLPri        *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ForceSet      *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
}

func (s BatchSetCdnDomainServerCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetCdnDomainServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetOwnerId(v int64) *BatchSetCdnDomainServerCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetSecurityToken(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetDomainName(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetCertName(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.CertName = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetCertType(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.CertType = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetSSLProtocol(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetSSLPub(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetSSLPri(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetRegion(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.Region = &v
	return s
}

func (s *BatchSetCdnDomainServerCertificateRequest) SetForceSet(v string) *BatchSetCdnDomainServerCertificateRequest {
	s.ForceSet = &v
	return s
}

type BatchSetCdnDomainServerCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetCdnDomainServerCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetCdnDomainServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainServerCertificateResponseBody) SetRequestId(v string) *BatchSetCdnDomainServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetCdnDomainServerCertificateResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetCdnDomainServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetCdnDomainServerCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetCdnDomainServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *BatchSetCdnDomainServerCertificateResponse) SetHeaders(v map[string]*string) *BatchSetCdnDomainServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *BatchSetCdnDomainServerCertificateResponse) SetBody(v *BatchSetCdnDomainServerCertificateResponseBody) *BatchSetCdnDomainServerCertificateResponse {
	s.Body = v
	return s
}

type BatchStartCdnDomainRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
}

func (s BatchStartCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartCdnDomainRequest) SetOwnerId(v int64) *BatchStartCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartCdnDomainRequest) SetSecurityToken(v string) *BatchStartCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStartCdnDomainRequest) SetDomainNames(v string) *BatchStartCdnDomainRequest {
	s.DomainNames = &v
	return s
}

type BatchStartCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStartCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartCdnDomainResponseBody) SetRequestId(v string) *BatchStartCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStartCdnDomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStartCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartCdnDomainResponse) SetHeaders(v map[string]*string) *BatchStartCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartCdnDomainResponse) SetBody(v *BatchStartCdnDomainResponseBody) *BatchStartCdnDomainResponse {
	s.Body = v
	return s
}

type BatchStopCdnDomainRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
}

func (s BatchStopCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopCdnDomainRequest) SetOwnerId(v int64) *BatchStopCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopCdnDomainRequest) SetSecurityToken(v string) *BatchStopCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchStopCdnDomainRequest) SetDomainNames(v string) *BatchStopCdnDomainRequest {
	s.DomainNames = &v
	return s
}

type BatchStopCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopCdnDomainResponseBody) SetRequestId(v string) *BatchStopCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStopCdnDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStopCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopCdnDomainResponse) SetHeaders(v map[string]*string) *BatchStopCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopCdnDomainResponse) SetBody(v *BatchStopCdnDomainResponseBody) *BatchStopCdnDomainResponse {
	s.Body = v
	return s
}

type BatchUpdateCdnDomainRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchUpdateCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateCdnDomainRequest) SetOwnerId(v int64) *BatchUpdateCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetSecurityToken(v string) *BatchUpdateCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetDomainName(v string) *BatchUpdateCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetSources(v string) *BatchUpdateCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetResourceGroupId(v string) *BatchUpdateCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchUpdateCdnDomainRequest) SetTopLevelDomain(v string) *BatchUpdateCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type BatchUpdateCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateCdnDomainResponseBody) SetRequestId(v string) *BatchUpdateCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchUpdateCdnDomainResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUpdateCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateCdnDomainResponse) SetHeaders(v map[string]*string) *BatchUpdateCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateCdnDomainResponse) SetBody(v *BatchUpdateCdnDomainResponseBody) *BatchUpdateCdnDomainResponse {
	s.Body = v
	return s
}

type CreateCdnCertificateSigningRequestRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	CommonName       *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	SANs             *string `json:"SANs,omitempty" xml:"SANs,omitempty"`
	Organization     *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
	City             *string `json:"City,omitempty" xml:"City,omitempty"`
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s CreateCdnCertificateSigningRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnCertificateSigningRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCdnCertificateSigningRequestRequest) SetOwnerId(v int64) *CreateCdnCertificateSigningRequestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetCommonName(v string) *CreateCdnCertificateSigningRequestRequest {
	s.CommonName = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetSANs(v string) *CreateCdnCertificateSigningRequestRequest {
	s.SANs = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetOrganization(v string) *CreateCdnCertificateSigningRequestRequest {
	s.Organization = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetOrganizationUnit(v string) *CreateCdnCertificateSigningRequestRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetCountry(v string) *CreateCdnCertificateSigningRequestRequest {
	s.Country = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetState(v string) *CreateCdnCertificateSigningRequestRequest {
	s.State = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetCity(v string) *CreateCdnCertificateSigningRequestRequest {
	s.City = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetEmail(v string) *CreateCdnCertificateSigningRequestRequest {
	s.Email = &v
	return s
}

type CreateCdnCertificateSigningRequestResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PubMd5     *string `json:"PubMd5,omitempty" xml:"PubMd5,omitempty"`
	Csr        *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
}

func (s CreateCdnCertificateSigningRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnCertificateSigningRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetRequestId(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetPubMd5(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.PubMd5 = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetCsr(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.Csr = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponseBody) SetCommonName(v string) *CreateCdnCertificateSigningRequestResponseBody {
	s.CommonName = &v
	return s
}

type CreateCdnCertificateSigningRequestResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCdnCertificateSigningRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCdnCertificateSigningRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnCertificateSigningRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCdnCertificateSigningRequestResponse) SetHeaders(v map[string]*string) *CreateCdnCertificateSigningRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCdnCertificateSigningRequestResponse) SetBody(v *CreateCdnCertificateSigningRequestResponseBody) *CreateCdnCertificateSigningRequestResponse {
	s.Body = v
	return s
}

type CreateCdnDeliverTaskRequest struct {
	OwnerId    *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Name       *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Status     *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Reports    *string                `json:"Reports,omitempty" xml:"Reports,omitempty"`
	DomainName *string                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Deliver    map[string]interface{} `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	Schedule   map[string]interface{} `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateCdnDeliverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCdnDeliverTaskRequest) SetOwnerId(v int64) *CreateCdnDeliverTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCdnDeliverTaskRequest) SetName(v string) *CreateCdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateCdnDeliverTaskRequest) SetStatus(v string) *CreateCdnDeliverTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateCdnDeliverTaskRequest) SetReports(v string) *CreateCdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *CreateCdnDeliverTaskRequest) SetDomainName(v string) *CreateCdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCdnDeliverTaskRequest) SetDeliver(v map[string]interface{}) *CreateCdnDeliverTaskRequest {
	s.Deliver = v
	return s
}

func (s *CreateCdnDeliverTaskRequest) SetSchedule(v map[string]interface{}) *CreateCdnDeliverTaskRequest {
	s.Schedule = v
	return s
}

type CreateCdnDeliverTaskShrinkRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Reports        *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DeliverShrink  *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateCdnDeliverTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnDeliverTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetOwnerId(v int64) *CreateCdnDeliverTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetName(v string) *CreateCdnDeliverTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetStatus(v string) *CreateCdnDeliverTaskShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetReports(v string) *CreateCdnDeliverTaskShrinkRequest {
	s.Reports = &v
	return s
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetDomainName(v string) *CreateCdnDeliverTaskShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetDeliverShrink(v string) *CreateCdnDeliverTaskShrinkRequest {
	s.DeliverShrink = &v
	return s
}

func (s *CreateCdnDeliverTaskShrinkRequest) SetScheduleShrink(v string) *CreateCdnDeliverTaskShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

type CreateCdnDeliverTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdnDeliverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdnDeliverTaskResponseBody) SetRequestId(v string) *CreateCdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateCdnDeliverTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCdnDeliverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCdnDeliverTaskResponse) SetHeaders(v map[string]*string) *CreateCdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCdnDeliverTaskResponse) SetBody(v *CreateCdnDeliverTaskResponseBody) *CreateCdnDeliverTaskResponse {
	s.Body = v
	return s
}

type CreateCdnSubTaskRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ReportIds  *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateCdnSubTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCdnSubTaskRequest) SetOwnerId(v int64) *CreateCdnSubTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCdnSubTaskRequest) SetStatus(v string) *CreateCdnSubTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateCdnSubTaskRequest) SetReportIds(v string) *CreateCdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

func (s *CreateCdnSubTaskRequest) SetDomainName(v string) *CreateCdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *CreateCdnSubTaskRequest) SetStartTime(v string) *CreateCdnSubTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCdnSubTaskRequest) SetEndTime(v string) *CreateCdnSubTaskRequest {
	s.EndTime = &v
	return s
}

type CreateCdnSubTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdnSubTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdnSubTaskResponseBody) SetRequestId(v string) *CreateCdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateCdnSubTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCdnSubTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCdnSubTaskResponse) SetHeaders(v map[string]*string) *CreateCdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCdnSubTaskResponse) SetBody(v *CreateCdnSubTaskResponseBody) *CreateCdnSubTaskResponse {
	s.Body = v
	return s
}

type CreateIllegalUrlExportTaskRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateIllegalUrlExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIllegalUrlExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateIllegalUrlExportTaskRequest) SetOwnerId(v int64) *CreateIllegalUrlExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIllegalUrlExportTaskRequest) SetTimePoint(v string) *CreateIllegalUrlExportTaskRequest {
	s.TimePoint = &v
	return s
}

func (s *CreateIllegalUrlExportTaskRequest) SetTaskName(v string) *CreateIllegalUrlExportTaskRequest {
	s.TaskName = &v
	return s
}

type CreateIllegalUrlExportTaskResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIllegalUrlExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIllegalUrlExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIllegalUrlExportTaskResponseBody) SetTaskId(v string) *CreateIllegalUrlExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateIllegalUrlExportTaskResponseBody) SetRequestId(v string) *CreateIllegalUrlExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateIllegalUrlExportTaskResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIllegalUrlExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIllegalUrlExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIllegalUrlExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateIllegalUrlExportTaskResponse) SetHeaders(v map[string]*string) *CreateIllegalUrlExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateIllegalUrlExportTaskResponse) SetBody(v *CreateIllegalUrlExportTaskResponseBody) *CreateIllegalUrlExportTaskResponse {
	s.Body = v
	return s
}

type CreateRealTimeLogDeliveryRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CreateRealTimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateRealTimeLogDeliveryRequest) SetOwnerId(v int64) *CreateRealTimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetProject(v string) *CreateRealTimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetLogstore(v string) *CreateRealTimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetRegion(v string) *CreateRealTimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetDomain(v string) *CreateRealTimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

type CreateRealTimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRealTimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type CreateRealTimeLogDeliveryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRealTimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateRealTimeLogDeliveryResponse) SetBody(v *CreateRealTimeLogDeliveryResponseBody) *CreateRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

type CreateUsageDetailDataExportTaskRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s CreateUsageDetailDataExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUsageDetailDataExportTaskRequest) SetOwnerId(v int64) *CreateUsageDetailDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetStartTime(v string) *CreateUsageDetailDataExportTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetEndTime(v string) *CreateUsageDetailDataExportTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetGroup(v string) *CreateUsageDetailDataExportTaskRequest {
	s.Group = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetDomainNames(v string) *CreateUsageDetailDataExportTaskRequest {
	s.DomainNames = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetType(v string) *CreateUsageDetailDataExportTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetTaskName(v string) *CreateUsageDetailDataExportTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskRequest) SetLanguage(v string) *CreateUsageDetailDataExportTaskRequest {
	s.Language = &v
	return s
}

type CreateUsageDetailDataExportTaskResponseBody struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateUsageDetailDataExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetEndTime(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetStartTime(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponseBody) SetTaskId(v string) *CreateUsageDetailDataExportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateUsageDetailDataExportTaskResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUsageDetailDataExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *CreateUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUsageDetailDataExportTaskResponse) SetBody(v *CreateUsageDetailDataExportTaskResponseBody) *CreateUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

type CreateUserUsageDataExportTaskRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Language  *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s CreateUserUsageDataExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserUsageDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUserUsageDataExportTaskRequest) SetOwnerId(v int64) *CreateUserUsageDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetStartTime(v string) *CreateUserUsageDataExportTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetEndTime(v string) *CreateUserUsageDataExportTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetTaskName(v string) *CreateUserUsageDataExportTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateUserUsageDataExportTaskRequest) SetLanguage(v string) *CreateUserUsageDataExportTaskRequest {
	s.Language = &v
	return s
}

type CreateUserUsageDataExportTaskResponseBody struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateUserUsageDataExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserUsageDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetEndTime(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetStartTime(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetRequestId(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserUsageDataExportTaskResponseBody) SetTaskId(v string) *CreateUserUsageDataExportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateUserUsageDataExportTaskResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserUsageDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserUsageDataExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserUsageDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUserUsageDataExportTaskResponse) SetHeaders(v map[string]*string) *CreateUserUsageDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUserUsageDataExportTaskResponse) SetBody(v *CreateUserUsageDataExportTaskResponseBody) *CreateUserUsageDataExportTaskResponse {
	s.Body = v
	return s
}

type DeleteCdnDeliverTaskRequest struct {
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DeliverId *int64 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
}

func (s DeleteCdnDeliverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteCdnDeliverTaskRequest) SetOwnerId(v int64) *DeleteCdnDeliverTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCdnDeliverTaskRequest) SetDeliverId(v int64) *DeleteCdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

type DeleteCdnDeliverTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCdnDeliverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdnDeliverTaskResponseBody) SetRequestId(v string) *DeleteCdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCdnDeliverTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCdnDeliverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdnDeliverTaskResponse) SetHeaders(v map[string]*string) *DeleteCdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdnDeliverTaskResponse) SetBody(v *DeleteCdnDeliverTaskResponseBody) *DeleteCdnDeliverTaskResponse {
	s.Body = v
	return s
}

type DeleteCdnDomainRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DeleteCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteCdnDomainRequest) SetOwnerId(v int64) *DeleteCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCdnDomainRequest) SetOwnerAccount(v string) *DeleteCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCdnDomainRequest) SetSecurityToken(v string) *DeleteCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteCdnDomainRequest) SetDomainName(v string) *DeleteCdnDomainRequest {
	s.DomainName = &v
	return s
}

type DeleteCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdnDomainResponseBody) SetRequestId(v string) *DeleteCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCdnDomainResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdnDomainResponse) SetHeaders(v map[string]*string) *DeleteCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdnDomainResponse) SetBody(v *DeleteCdnDomainResponseBody) *DeleteCdnDomainResponse {
	s.Body = v
	return s
}

type DeleteCdnSubTaskRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteCdnSubTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteCdnSubTaskRequest) SetOwnerId(v int64) *DeleteCdnSubTaskRequest {
	s.OwnerId = &v
	return s
}

type DeleteCdnSubTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCdnSubTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdnSubTaskResponseBody) SetRequestId(v string) *DeleteCdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCdnSubTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCdnSubTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdnSubTaskResponse) SetHeaders(v map[string]*string) *DeleteCdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdnSubTaskResponse) SetBody(v *DeleteCdnSubTaskResponseBody) *DeleteCdnSubTaskResponse {
	s.Body = v
	return s
}

type DeleteFCTriggerRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TriggerARN *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
}

func (s DeleteFCTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteFCTriggerRequest) SetOwnerId(v int64) *DeleteFCTriggerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFCTriggerRequest) SetTriggerARN(v string) *DeleteFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

type DeleteFCTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFCTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFCTriggerResponseBody) SetRequestId(v string) *DeleteFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFCTriggerResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFCTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteFCTriggerResponse) SetHeaders(v map[string]*string) *DeleteFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteFCTriggerResponse) SetBody(v *DeleteFCTriggerResponseBody) *DeleteFCTriggerResponse {
	s.Body = v
	return s
}

type DeleteRealtimeLogDeliveryRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DeleteRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetDomain(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetProject(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetLogstore(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetRegion(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

type DeleteRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DeleteRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DeleteRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRealtimeLogDeliveryResponse) SetBody(v *DeleteRealtimeLogDeliveryResponseBody) *DeleteRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type DeleteSpecificConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s DeleteSpecificConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpecificConfigRequest) SetOwnerId(v int64) *DeleteSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSpecificConfigRequest) SetSecurityToken(v string) *DeleteSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSpecificConfigRequest) SetDomainName(v string) *DeleteSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSpecificConfigRequest) SetConfigId(v string) *DeleteSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

type DeleteSpecificConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpecificConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpecificConfigResponseBody) SetRequestId(v string) *DeleteSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSpecificConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSpecificConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpecificConfigResponse) SetBody(v *DeleteSpecificConfigResponseBody) *DeleteSpecificConfigResponse {
	s.Body = v
	return s
}

type DeleteSpecificStagingConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s DeleteSpecificStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteSpecificStagingConfigRequest) SetOwnerId(v int64) *DeleteSpecificStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) SetSecurityToken(v string) *DeleteSpecificStagingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) SetDomainName(v string) *DeleteSpecificStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSpecificStagingConfigRequest) SetConfigId(v string) *DeleteSpecificStagingConfigRequest {
	s.ConfigId = &v
	return s
}

type DeleteSpecificStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpecificStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpecificStagingConfigResponseBody) SetRequestId(v string) *DeleteSpecificStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSpecificStagingConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSpecificStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSpecificStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSpecificStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpecificStagingConfigResponse) SetHeaders(v map[string]*string) *DeleteSpecificStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpecificStagingConfigResponse) SetBody(v *DeleteSpecificStagingConfigResponseBody) *DeleteSpecificStagingConfigResponse {
	s.Body = v
	return s
}

type DeleteUsageDetailDataExportTaskRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteUsageDetailDataExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsageDetailDataExportTaskRequest) SetOwnerId(v int64) *DeleteUsageDetailDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUsageDetailDataExportTaskRequest) SetTaskId(v string) *DeleteUsageDetailDataExportTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteUsageDetailDataExportTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsageDetailDataExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *DeleteUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUsageDetailDataExportTaskResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUsageDetailDataExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *DeleteUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsageDetailDataExportTaskResponse) SetBody(v *DeleteUsageDetailDataExportTaskResponseBody) *DeleteUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

type DeleteUserUsageDataExportTaskRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteUserUsageDataExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserUsageDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserUsageDataExportTaskRequest) SetOwnerId(v int64) *DeleteUserUsageDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteUserUsageDataExportTaskRequest) SetTaskId(v string) *DeleteUserUsageDataExportTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteUserUsageDataExportTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserUsageDataExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserUsageDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserUsageDataExportTaskResponseBody) SetRequestId(v string) *DeleteUserUsageDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserUsageDataExportTaskResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserUsageDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserUsageDataExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserUsageDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserUsageDataExportTaskResponse) SetHeaders(v map[string]*string) *DeleteUserUsageDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserUsageDataExportTaskResponse) SetBody(v *DeleteUserUsageDataExportTaskResponseBody) *DeleteUserUsageDataExportTaskResponse {
	s.Body = v
	return s
}

type DescribeActiveVersionOfConfigGroupRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ConfigGroupId *string `json:"ConfigGroupId,omitempty" xml:"ConfigGroupId,omitempty"`
	Env           *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s DescribeActiveVersionOfConfigGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveVersionOfConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveVersionOfConfigGroupRequest) SetOwnerId(v int64) *DescribeActiveVersionOfConfigGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupRequest) SetConfigGroupId(v string) *DescribeActiveVersionOfConfigGroupRequest {
	s.ConfigGroupId = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupRequest) SetEnv(v string) *DescribeActiveVersionOfConfigGroupRequest {
	s.Env = &v
	return s
}

type DescribeActiveVersionOfConfigGroupResponseBody struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SeqId         *int64  `json:"SeqId,omitempty" xml:"SeqId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ConfigGroupId *string `json:"ConfigGroupId,omitempty" xml:"ConfigGroupId,omitempty"`
	Operator      *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	BaseVersionId *string `json:"BaseVersionId,omitempty" xml:"BaseVersionId,omitempty"`
}

func (s DescribeActiveVersionOfConfigGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveVersionOfConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetStatus(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetUpdateTime(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetRequestId(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetSeqId(v int64) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.SeqId = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetDescription(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetCreateTime(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetVersionId(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.VersionId = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetConfigGroupId(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.ConfigGroupId = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetOperator(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.Operator = &v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponseBody) SetBaseVersionId(v string) *DescribeActiveVersionOfConfigGroupResponseBody {
	s.BaseVersionId = &v
	return s
}

type DescribeActiveVersionOfConfigGroupResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeActiveVersionOfConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeActiveVersionOfConfigGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveVersionOfConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveVersionOfConfigGroupResponse) SetHeaders(v map[string]*string) *DescribeActiveVersionOfConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveVersionOfConfigGroupResponse) SetBody(v *DescribeActiveVersionOfConfigGroupResponseBody) *DescribeActiveVersionOfConfigGroupResponse {
	s.Body = v
	return s
}

type DescribeCdnCertificateDetailRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
}

func (s DescribeCdnCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailRequest) SetOwnerId(v int64) *DescribeCdnCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnCertificateDetailRequest) SetSecurityToken(v string) *DescribeCdnCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnCertificateDetailRequest) SetCertName(v string) *DescribeCdnCertificateDetailRequest {
	s.CertName = &v
	return s
}

type DescribeCdnCertificateDetailResponseBody struct {
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailResponseBody) SetCertName(v string) *DescribeCdnCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetKey(v string) *DescribeCdnCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetCert(v string) *DescribeCdnCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetCertId(v int64) *DescribeCdnCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetRequestId(v string) *DescribeCdnCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnCertificateDetailResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeCdnCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnCertificateDetailResponse) SetBody(v *DescribeCdnCertificateDetailResponseBody) *DescribeCdnCertificateDetailResponse {
	s.Body = v
	return s
}

type DescribeCdnCertificateListRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListRequest) SetOwnerId(v int64) *DescribeCdnCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnCertificateListRequest) SetSecurityToken(v string) *DescribeCdnCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnCertificateListRequest) SetDomainName(v string) *DescribeCdnCertificateListRequest {
	s.DomainName = &v
	return s
}

type DescribeCdnCertificateListResponseBody struct {
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertificateListModel *DescribeCdnCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
}

func (s DescribeCdnCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBody) SetRequestId(v string) *DescribeCdnCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBody) SetCertificateListModel(v *DescribeCdnCertificateListResponseBodyCertificateListModel) *DescribeCdnCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

type DescribeCdnCertificateListResponseBodyCertificateListModel struct {
	Count    *int32                                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	CertList *DescribeCdnCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeCdnCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeCdnCertificateListResponseBodyCertificateListModelCertList) *DescribeCdnCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

type DescribeCdnCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) *DescribeCdnCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

type DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert struct {
	LastTime    *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	CertId      *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty"`
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

type DescribeCdnCertificateListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponse) SetHeaders(v map[string]*string) *DescribeCdnCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnCertificateListResponse) SetBody(v *DescribeCdnCertificateListResponseBody) *DescribeCdnCertificateListResponse {
	s.Body = v
	return s
}

type DescribeCdnDeliverListRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DeliverId *int64  `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnDeliverListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDeliverListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListRequest) SetOwnerId(v int64) *DescribeCdnDeliverListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDeliverListRequest) SetDeliverId(v int64) *DescribeCdnDeliverListRequest {
	s.DeliverId = &v
	return s
}

func (s *DescribeCdnDeliverListRequest) SetStatus(v string) *DescribeCdnDeliverListRequest {
	s.Status = &v
	return s
}

type DescribeCdnDeliverListResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDeliverListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBody) SetContent(v string) *DescribeCdnDeliverListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBody) SetRequestId(v string) *DescribeCdnDeliverListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnDeliverListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnDeliverListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDeliverListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDeliverListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponse) SetHeaders(v map[string]*string) *DescribeCdnDeliverListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDeliverListResponse) SetBody(v *DescribeCdnDeliverListResponseBody) *DescribeCdnDeliverListResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainByCertificateRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SSLPub  *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
}

func (s DescribeCdnDomainByCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainByCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateRequest) SetOwnerId(v int64) *DescribeCdnDomainByCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainByCertificateRequest) SetSSLPub(v string) *DescribeCdnDomainByCertificateRequest {
	s.SSLPub = &v
	return s
}

type DescribeCdnDomainByCertificateResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertInfos *DescribeCdnDomainByCertificateResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainByCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponseBody) SetRequestId(v string) *DescribeCdnDomainByCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBody) SetCertInfos(v *DescribeCdnDomainByCertificateResponseBodyCertInfos) *DescribeCdnDomainByCertificateResponseBody {
	s.CertInfos = v
	return s
}

type DescribeCdnDomainByCertificateResponseBodyCertInfos struct {
	CertInfo []*DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfos) SetCertInfo(v []*DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) *DescribeCdnDomainByCertificateResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo struct {
	CertStartTime         *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertExpireTime        *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertCaIsLegacy        *string `json:"CertCaIsLegacy,omitempty" xml:"CertCaIsLegacy,omitempty"`
	CertSubjectCommonName *string `json:"CertSubjectCommonName,omitempty" xml:"CertSubjectCommonName,omitempty"`
	CertType              *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainNames           *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	CertExpired           *string `json:"CertExpired,omitempty" xml:"CertExpired,omitempty"`
	Issuer                *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	DomainList            *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertCaIsLegacy(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertCaIsLegacy = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertSubjectCommonName(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertSubjectCommonName = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainNames(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainNames = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpired(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpired = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetIssuer(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.Issuer = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

type DescribeCdnDomainByCertificateResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnDomainByCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainByCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainByCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainByCertificateResponse) SetBody(v *DescribeCdnDomainByCertificateResponseBody) *DescribeCdnDomainByCertificateResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainConfigsRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
}

func (s DescribeCdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsRequest) SetOwnerId(v int64) *DescribeCdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetSecurityToken(v string) *DescribeCdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetDomainName(v string) *DescribeCdnDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainConfigsRequest) SetFunctionNames(v string) *DescribeCdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

type DescribeCdnDomainConfigsResponseBody struct {
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainConfigs *DescribeCdnDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBody) SetRequestId(v string) *DescribeCdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBody) SetDomainConfigs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigs) *DescribeCdnDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeCdnDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	Status       *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	ConfigId     *string                                                                    `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionName *string                                                                    `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionArgs *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeCdnDomainConfigsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainConfigsResponse) SetBody(v *DescribeCdnDomainConfigsResponseBody) *DescribeCdnDomainConfigsResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainDetailRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailRequest) SetOwnerId(v int64) *DescribeCdnDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) SetSecurityToken(v string) *DescribeCdnDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) SetDomainName(v string) *DescribeCdnDomainDetailRequest {
	s.DomainName = &v
	return s
}

type DescribeCdnDomainDetailResponseBody struct {
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GetDomainDetailModel *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel `json:"GetDomainDetailModel,omitempty" xml:"GetDomainDetailModel,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBody) SetRequestId(v string) *DescribeCdnDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBody) SetGetDomainDetailModel(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) *DescribeCdnDomainDetailResponseBody {
	s.GetDomainDetailModel = v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModel struct {
	GmtCreated              *string                                                              `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	Description             *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	HttpsCname              *string                                                              `json:"HttpsCname,omitempty" xml:"HttpsCname,omitempty"`
	ResourceGroupId         *string                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerCertificateStatus *string                                                              `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	Scope                   *string                                                              `json:"Scope,omitempty" xml:"Scope,omitempty"`
	DomainStatus            *string                                                              `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	Cname                   *string                                                              `json:"Cname,omitempty" xml:"Cname,omitempty"`
	GmtModified             *string                                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	CdnType                 *string                                                              `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	DomainName              *string                                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SourceModels            *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels `json:"SourceModels,omitempty" xml:"SourceModels,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetGmtCreated(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDescription(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Description = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetHttpsCname(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.HttpsCname = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetResourceGroupId(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetServerCertificateStatus(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetScope(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Scope = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDomainStatus(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.DomainStatus = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetCname(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.Cname = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetGmtModified(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetCdnType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.CdnType = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetDomainName(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel) SetSourceModels(v *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModel {
	s.SourceModels = v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels struct {
	SourceModel []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel `json:"SourceModel,omitempty" xml:"SourceModel,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels) SetSourceModel(v []*DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModels {
	s.SourceModel = v
	return s
}

type DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetType(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Type = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetWeight(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Weight = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetEnabled(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Enabled = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetPriority(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Priority = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetPort(v int32) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Port = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel) SetContent(v string) *DescribeCdnDomainDetailResponseBodyGetDomainDetailModelSourceModelsSourceModel {
	s.Content = &v
	return s
}

type DescribeCdnDomainDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainDetailResponse) SetBody(v *DescribeCdnDomainDetailResponseBody) *DescribeCdnDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainLogsRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCdnDomainLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsRequest) SetOwnerId(v int64) *DescribeCdnDomainLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetDomainName(v string) *DescribeCdnDomainLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageSize(v int64) *DescribeCdnDomainLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetPageNumber(v int64) *DescribeCdnDomainLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetStartTime(v string) *DescribeCdnDomainLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsRequest) SetEndTime(v string) *DescribeCdnDomainLogsRequest {
	s.EndTime = &v
	return s
}

type DescribeCdnDomainLogsResponseBody struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainLogDetails *DescribeCdnDomainLogsResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBody) SetRequestId(v string) *DescribeCdnDomainLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetDomainLogDetails(v *DescribeCdnDomainLogsResponseBodyDomainLogDetails) *DescribeCdnDomainLogsResponseBody {
	s.DomainLogDetails = v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) *DescribeCdnDomainLogsResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount   *int64                                                                     `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	DomainName *string                                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PageInfos  *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
	LogInfos   *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetDomainName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageInfoDetail []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail `json:"PageInfoDetail,omitempty" xml:"PageInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageInfoDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageInfoDetail = v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail) SetPageIndex(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail {
	s.PageIndex = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail) SetPageSize(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail) SetTotal(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfosPageInfoDetail {
	s.Total = &v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

type DescribeCdnDomainLogsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnDomainLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainLogsResponse) SetBody(v *DescribeCdnDomainLogsResponseBody) *DescribeCdnDomainLogsResponse {
	s.Body = v
	return s
}

type DescribeCdnDomainStagingConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
}

func (s DescribeCdnDomainStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigRequest) SetOwnerId(v int64) *DescribeCdnDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigRequest) SetDomainName(v string) *DescribeCdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigRequest) SetFunctionNames(v string) *DescribeCdnDomainStagingConfigRequest {
	s.FunctionNames = &v
	return s
}

type DescribeCdnDomainStagingConfigResponseBody struct {
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainConfigs []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponseBody) SetRequestId(v string) *DescribeCdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBody) SetDomainConfigs(v []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) *DescribeCdnDomainStagingConfigResponseBody {
	s.DomainConfigs = v
	return s
}

type DescribeCdnDomainStagingConfigResponseBodyDomainConfigs struct {
	Status       *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	ConfigId     *string                                                                `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionName *string                                                                `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionArgs []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetConfigId(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

type DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

type DescribeCdnDomainStagingConfigResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnDomainStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponse) SetBody(v *DescribeCdnDomainStagingConfigResponseBody) *DescribeCdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

type DescribeCdnHttpsDomainListRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s DescribeCdnHttpsDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnHttpsDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListRequest) SetOwnerId(v int64) *DescribeCdnHttpsDomainListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnHttpsDomainListRequest) SetPageNumber(v int32) *DescribeCdnHttpsDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnHttpsDomainListRequest) SetPageSize(v int32) *DescribeCdnHttpsDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnHttpsDomainListRequest) SetKeyword(v string) *DescribeCdnHttpsDomainListRequest {
	s.Keyword = &v
	return s
}

type DescribeCdnHttpsDomainListResponseBody struct {
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertInfos  *DescribeCdnHttpsDomainListResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnHttpsDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponseBody) SetTotalCount(v int32) *DescribeCdnHttpsDomainListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBody) SetRequestId(v string) *DescribeCdnHttpsDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBody) SetCertInfos(v *DescribeCdnHttpsDomainListResponseBodyCertInfos) *DescribeCdnHttpsDomainListResponseBody {
	s.CertInfos = v
	return s
}

type DescribeCdnHttpsDomainListResponseBodyCertInfos struct {
	CertInfo []*DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfos) SetCertInfo(v []*DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) *DescribeCdnHttpsDomainListResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo struct {
	CertStartTime  *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertUpdateTime *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertStatus     *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStatus(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStatus = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertCommonName(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertCommonName = &v
	return s
}

type DescribeCdnHttpsDomainListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnHttpsDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnHttpsDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponse) SetHeaders(v map[string]*string) *DescribeCdnHttpsDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnHttpsDomainListResponse) SetBody(v *DescribeCdnHttpsDomainListResponseBody) *DescribeCdnHttpsDomainListResponse {
	s.Body = v
	return s
}

type DescribeCdnRegionAndIspRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnRegionAndIspRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspRequest) SetOwnerId(v int64) *DescribeCdnRegionAndIspRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnRegionAndIspRequest) SetSecurityToken(v string) *DescribeCdnRegionAndIspRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCdnRegionAndIspResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeCdnRegionAndIspResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	Isps      *DescribeCdnRegionAndIspResponseBodyIsps    `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Struct"`
}

func (s DescribeCdnRegionAndIspResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBody) SetRequestId(v string) *DescribeCdnRegionAndIspResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBody) SetRegions(v *DescribeCdnRegionAndIspResponseBodyRegions) *DescribeCdnRegionAndIspResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBody) SetIsps(v *DescribeCdnRegionAndIspResponseBodyIsps) *DescribeCdnRegionAndIspResponseBody {
	s.Isps = v
	return s
}

type DescribeCdnRegionAndIspResponseBodyRegions struct {
	Region []*DescribeCdnRegionAndIspResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeCdnRegionAndIspResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyRegions) SetRegion(v []*DescribeCdnRegionAndIspResponseBodyRegionsRegion) *DescribeCdnRegionAndIspResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeCdnRegionAndIspResponseBodyRegionsRegion struct {
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeCdnRegionAndIspResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) SetNameEn(v string) *DescribeCdnRegionAndIspResponseBodyRegionsRegion {
	s.NameEn = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) SetNameZh(v string) *DescribeCdnRegionAndIspResponseBodyRegionsRegion {
	s.NameZh = &v
	return s
}

type DescribeCdnRegionAndIspResponseBodyIsps struct {
	Isp []*DescribeCdnRegionAndIspResponseBodyIspsIsp `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Repeated"`
}

func (s DescribeCdnRegionAndIspResponseBodyIsps) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyIsps) SetIsp(v []*DescribeCdnRegionAndIspResponseBodyIspsIsp) *DescribeCdnRegionAndIspResponseBodyIsps {
	s.Isp = v
	return s
}

type DescribeCdnRegionAndIspResponseBodyIspsIsp struct {
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeCdnRegionAndIspResponseBodyIspsIsp) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyIspsIsp) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) SetNameEn(v string) *DescribeCdnRegionAndIspResponseBodyIspsIsp {
	s.NameEn = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) SetNameZh(v string) *DescribeCdnRegionAndIspResponseBodyIspsIsp {
	s.NameZh = &v
	return s
}

type DescribeCdnRegionAndIspResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnRegionAndIspResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnRegionAndIspResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponse) SetHeaders(v map[string]*string) *DescribeCdnRegionAndIspResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnRegionAndIspResponse) SetBody(v *DescribeCdnRegionAndIspResponseBody) *DescribeCdnRegionAndIspResponse {
	s.Body = v
	return s
}

type DescribeCdnReportRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ReportId   *int64  `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	Area       *string `json:"Area,omitempty" xml:"Area,omitempty"`
	IsOverseas *string `json:"IsOverseas,omitempty" xml:"IsOverseas,omitempty"`
	HttpCode   *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCdnReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportRequest) SetOwnerId(v int64) *DescribeCdnReportRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnReportRequest) SetDomainName(v string) *DescribeCdnReportRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnReportRequest) SetReportId(v int64) *DescribeCdnReportRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeCdnReportRequest) SetArea(v string) *DescribeCdnReportRequest {
	s.Area = &v
	return s
}

func (s *DescribeCdnReportRequest) SetIsOverseas(v string) *DescribeCdnReportRequest {
	s.IsOverseas = &v
	return s
}

func (s *DescribeCdnReportRequest) SetHttpCode(v string) *DescribeCdnReportRequest {
	s.HttpCode = &v
	return s
}

func (s *DescribeCdnReportRequest) SetStartTime(v string) *DescribeCdnReportRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnReportRequest) SetEndTime(v string) *DescribeCdnReportRequest {
	s.EndTime = &v
	return s
}

type DescribeCdnReportResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponseBody) SetContent(v string) *DescribeCdnReportResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnReportResponseBody) SetRequestId(v string) *DescribeCdnReportResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnReportResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportResponse) SetHeaders(v map[string]*string) *DescribeCdnReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnReportResponse) SetBody(v *DescribeCdnReportResponseBody) *DescribeCdnReportResponse {
	s.Body = v
	return s
}

type DescribeCdnReportListRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ReportId   *int64  `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s DescribeCdnReportListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListRequest) SetOwnerId(v int64) *DescribeCdnReportListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnReportListRequest) SetReportId(v int64) *DescribeCdnReportListRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeCdnReportListRequest) SetStatus(v string) *DescribeCdnReportListRequest {
	s.Status = &v
	return s
}

func (s *DescribeCdnReportListRequest) SetPermission(v string) *DescribeCdnReportListRequest {
	s.Permission = &v
	return s
}

type DescribeCdnReportListResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnReportListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponseBody) SetContent(v string) *DescribeCdnReportListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnReportListResponseBody) SetRequestId(v string) *DescribeCdnReportListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnReportListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnReportListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnReportListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponse) SetHeaders(v map[string]*string) *DescribeCdnReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnReportListResponse) SetBody(v *DescribeCdnReportListResponseBody) *DescribeCdnReportListResponse {
	s.Body = v
	return s
}

type DescribeCdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceRequest) SetOwnerId(v int64) *DescribeCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnServiceRequest) SetSecurityToken(v string) *DescribeCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCdnServiceResponseBody struct {
	ChangingAffectTime *string                                       `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChangingChargeType *string                                       `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	OpeningTime        *string                                       `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	InternetChargeType *string                                       `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InstanceId         *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperationLocks     *DescribeCdnServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
}

func (s DescribeCdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBody) SetChangingAffectTime(v string) *DescribeCdnServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetRequestId(v string) *DescribeCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetChangingChargeType(v string) *DescribeCdnServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetOpeningTime(v string) *DescribeCdnServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetInternetChargeType(v string) *DescribeCdnServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetInstanceId(v string) *DescribeCdnServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetOperationLocks(v *DescribeCdnServiceResponseBodyOperationLocks) *DescribeCdnServiceResponseBody {
	s.OperationLocks = v
	return s
}

type DescribeCdnServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeCdnServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeCdnServiceResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeCdnServiceResponseBodyOperationLocksLockReason) *DescribeCdnServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

type DescribeCdnServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeCdnServiceResponseBodyOperationLocksLockReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeCdnServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

type DescribeCdnServiceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponse) SetHeaders(v map[string]*string) *DescribeCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnServiceResponse) SetBody(v *DescribeCdnServiceResponseBody) *DescribeCdnServiceResponse {
	s.Body = v
	return s
}

type DescribeCdnSubListRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnSubListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnSubListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListRequest) SetOwnerId(v int64) *DescribeCdnSubListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnSubListRequest) SetStatus(v string) *DescribeCdnSubListRequest {
	s.Status = &v
	return s
}

type DescribeCdnSubListResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnSubListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnSubListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListResponseBody) SetContent(v string) *DescribeCdnSubListResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCdnSubListResponseBody) SetRequestId(v string) *DescribeCdnSubListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCdnSubListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnSubListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnSubListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnSubListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnSubListResponse) SetHeaders(v map[string]*string) *DescribeCdnSubListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnSubListResponse) SetBody(v *DescribeCdnSubListResponseBody) *DescribeCdnSubListResponse {
	s.Body = v
	return s
}

type DescribeCdnUserBillHistoryRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCdnUserBillHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryRequest) SetOwnerId(v int64) *DescribeCdnUserBillHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserBillHistoryRequest) SetStartTime(v string) *DescribeCdnUserBillHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillHistoryRequest) SetEndTime(v string) *DescribeCdnUserBillHistoryRequest {
	s.EndTime = &v
	return s
}

type DescribeCdnUserBillHistoryResponseBody struct {
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BillHistoryData *DescribeCdnUserBillHistoryResponseBodyBillHistoryData `json:"BillHistoryData,omitempty" xml:"BillHistoryData,omitempty" type:"Struct"`
}

func (s DescribeCdnUserBillHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBody) SetRequestId(v string) *DescribeCdnUserBillHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBody) SetBillHistoryData(v *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) *DescribeCdnUserBillHistoryResponseBody {
	s.BillHistoryData = v
	return s
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryData struct {
	BillHistoryDataItem []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem `json:"BillHistoryDataItem,omitempty" xml:"BillHistoryDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) SetBillHistoryDataItem(v []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) *DescribeCdnUserBillHistoryResponseBodyBillHistoryData {
	s.BillHistoryDataItem = v
	return s
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem struct {
	Dimension   *string                                                                              `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	BillType    *string                                                                              `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BillTime    *string                                                                              `json:"BillTime,omitempty" xml:"BillTime,omitempty"`
	BillingData *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData `json:"BillingData,omitempty" xml:"BillingData,omitempty" type:"Struct"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetDimension(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillType(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillTime(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillTime = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillingData(v *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillingData = v
	return s
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData struct {
	BillingDataItem []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem `json:"BillingDataItem,omitempty" xml:"BillingDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) SetBillingDataItem(v []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData {
	s.BillingDataItem = v
	return s
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem struct {
	Flow       *float32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	Bandwidth  *float32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Count      *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	CdnRegion  *string  `json:"CdnRegion,omitempty" xml:"CdnRegion,omitempty"`
	ChargeType *string  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetFlow(v float32) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Flow = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetBandwidth(v float32) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCount(v float32) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Count = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCdnRegion(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.CdnRegion = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetChargeType(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.ChargeType = &v
	return s
}

type DescribeCdnUserBillHistoryResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserBillHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserBillHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponse) SetHeaders(v map[string]*string) *DescribeCdnUserBillHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponse) SetBody(v *DescribeCdnUserBillHistoryResponseBody) *DescribeCdnUserBillHistoryResponse {
	s.Body = v
	return s
}

type DescribeCdnUserBillPredictionRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	Area      *string `json:"Area,omitempty" xml:"Area,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCdnUserBillPredictionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillPredictionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionRequest) SetOwnerId(v int64) *DescribeCdnUserBillPredictionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetDimension(v string) *DescribeCdnUserBillPredictionRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetArea(v string) *DescribeCdnUserBillPredictionRequest {
	s.Area = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetStartTime(v string) *DescribeCdnUserBillPredictionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionRequest) SetEndTime(v string) *DescribeCdnUserBillPredictionRequest {
	s.EndTime = &v
	return s
}

type DescribeCdnUserBillPredictionResponseBody struct {
	EndTime            *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime          *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BillType           *string                                                      `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BillPredictionData *DescribeCdnUserBillPredictionResponseBodyBillPredictionData `json:"BillPredictionData,omitempty" xml:"BillPredictionData,omitempty" type:"Struct"`
}

func (s DescribeCdnUserBillPredictionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetEndTime(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetStartTime(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetRequestId(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetBillType(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetBillPredictionData(v *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) *DescribeCdnUserBillPredictionResponseBody {
	s.BillPredictionData = v
	return s
}

type DescribeCdnUserBillPredictionResponseBodyBillPredictionData struct {
	BillPredictionDataItem []*DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem `json:"BillPredictionDataItem,omitempty" xml:"BillPredictionDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) SetBillPredictionDataItem(v []*DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) *DescribeCdnUserBillPredictionResponseBodyBillPredictionData {
	s.BillPredictionDataItem = v
	return s
}

type DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem struct {
	Value   *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStp *string  `json:"TimeStp,omitempty" xml:"TimeStp,omitempty"`
	Area    *string  `json:"Area,omitempty" xml:"Area,omitempty"`
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetValue(v float32) *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Value = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetTimeStp(v string) *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.TimeStp = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetArea(v string) *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Area = &v
	return s
}

type DescribeCdnUserBillPredictionResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserBillPredictionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserBillPredictionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponse) SetHeaders(v map[string]*string) *DescribeCdnUserBillPredictionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserBillPredictionResponse) SetBody(v *DescribeCdnUserBillPredictionResponseBody) *DescribeCdnUserBillPredictionResponse {
	s.Body = v
	return s
}

type DescribeCdnUserBillTypeRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCdnUserBillTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeRequest) SetOwnerId(v int64) *DescribeCdnUserBillTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserBillTypeRequest) SetStartTime(v string) *DescribeCdnUserBillTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeRequest) SetEndTime(v string) *DescribeCdnUserBillTypeRequest {
	s.EndTime = &v
	return s
}

type DescribeCdnUserBillTypeResponseBody struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BillTypeData *DescribeCdnUserBillTypeResponseBodyBillTypeData `json:"BillTypeData,omitempty" xml:"BillTypeData,omitempty" type:"Struct"`
}

func (s DescribeCdnUserBillTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBody) SetRequestId(v string) *DescribeCdnUserBillTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBody) SetBillTypeData(v *DescribeCdnUserBillTypeResponseBodyBillTypeData) *DescribeCdnUserBillTypeResponseBody {
	s.BillTypeData = v
	return s
}

type DescribeCdnUserBillTypeResponseBodyBillTypeData struct {
	BillTypeDataItem []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem `json:"BillTypeDataItem,omitempty" xml:"BillTypeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) SetBillTypeDataItem(v []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) *DescribeCdnUserBillTypeResponseBodyBillTypeData {
	s.BillTypeDataItem = v
	return s
}

type DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	BillType     *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	Dimension    *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetEndTime(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetStartTime(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillingCycle(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillingCycle = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetProduct(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Product = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillType(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetDimension(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Dimension = &v
	return s
}

type DescribeCdnUserBillTypeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserBillTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserBillTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponse) SetHeaders(v map[string]*string) *DescribeCdnUserBillTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserBillTypeResponse) SetBody(v *DescribeCdnUserBillTypeResponseBody) *DescribeCdnUserBillTypeResponse {
	s.Body = v
	return s
}

type DescribeCdnUserConfigsRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeCdnUserConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsRequest) SetOwnerId(v int64) *DescribeCdnUserConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserConfigsRequest) SetFunctionName(v string) *DescribeCdnUserConfigsRequest {
	s.FunctionName = &v
	return s
}

type DescribeCdnUserConfigsResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Configs   []*DescribeCdnUserConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsResponseBody) SetRequestId(v string) *DescribeCdnUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBody) SetConfigs(v []*DescribeCdnUserConfigsResponseBodyConfigs) *DescribeCdnUserConfigsResponseBody {
	s.Configs = v
	return s
}

type DescribeCdnUserConfigsResponseBodyConfigs struct {
	ArgValue     *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
	ArgName      *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeCdnUserConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) SetArgValue(v string) *DescribeCdnUserConfigsResponseBodyConfigs {
	s.ArgValue = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) SetArgName(v string) *DescribeCdnUserConfigsResponseBodyConfigs {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnUserConfigsResponseBodyConfigs) SetFunctionName(v string) *DescribeCdnUserConfigsResponseBodyConfigs {
	s.FunctionName = &v
	return s
}

type DescribeCdnUserConfigsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsResponse) SetHeaders(v map[string]*string) *DescribeCdnUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserConfigsResponse) SetBody(v *DescribeCdnUserConfigsResponseBody) *DescribeCdnUserConfigsResponse {
	s.Body = v
	return s
}

type DescribeCdnUserDomainsByFuncRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	FuncId          *int32  `json:"FuncId,omitempty" xml:"FuncId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCdnUserDomainsByFuncRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetOwnerId(v int64) *DescribeCdnUserDomainsByFuncRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetFuncId(v int32) *DescribeCdnUserDomainsByFuncRequest {
	s.FuncId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetPageNumber(v int32) *DescribeCdnUserDomainsByFuncRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetPageSize(v int32) *DescribeCdnUserDomainsByFuncRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetResourceGroupId(v string) *DescribeCdnUserDomainsByFuncRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeCdnUserDomainsByFuncResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Domains    *DescribeCdnUserDomainsByFuncResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
}

func (s DescribeCdnUserDomainsByFuncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetRequestId(v string) *DescribeCdnUserDomainsByFuncResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetPageNumber(v int64) *DescribeCdnUserDomainsByFuncResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetPageSize(v int64) *DescribeCdnUserDomainsByFuncResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetTotalCount(v int64) *DescribeCdnUserDomainsByFuncResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBody) SetDomains(v *DescribeCdnUserDomainsByFuncResponseBodyDomains) *DescribeCdnUserDomainsByFuncResponseBody {
	s.Domains = v
	return s
}

type DescribeCdnUserDomainsByFuncResponseBodyDomains struct {
	PageData []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomains) SetPageData(v []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) *DescribeCdnUserDomainsByFuncResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData struct {
	GmtCreated      *string                                                         `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	SslProtocol     *string                                                         `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
	Description     *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sandbox         *string                                                         `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	DomainStatus    *string                                                         `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	Cname           *string                                                         `json:"Cname,omitempty" xml:"Cname,omitempty"`
	GmtModified     *string                                                         `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	CdnType         *string                                                         `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	DomainName      *string                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Sources         *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetDescription(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetSandbox(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetCname(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetCdnType(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.CdnType = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetDomainName(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData) SetSources(v *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

type DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources struct {
	Source []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources) SetSource(v []*DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeCdnUserDomainsByFuncResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

type DescribeCdnUserDomainsByFuncResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserDomainsByFuncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserDomainsByFuncResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncResponse) SetHeaders(v map[string]*string) *DescribeCdnUserDomainsByFuncResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserDomainsByFuncResponse) SetBody(v *DescribeCdnUserDomainsByFuncResponseBody) *DescribeCdnUserDomainsByFuncResponse {
	s.Body = v
	return s
}

type DescribeCdnUserQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserQuotaRequest) SetOwnerId(v int64) *DescribeCdnUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserQuotaRequest) SetSecurityToken(v string) *DescribeCdnUserQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCdnUserQuotaResponseBody struct {
	BlockQuota       *int32  `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	RefreshUrlRemain *int32  `json:"RefreshUrlRemain,omitempty" xml:"RefreshUrlRemain,omitempty"`
	DomainQuota      *int32  `json:"DomainQuota,omitempty" xml:"DomainQuota,omitempty"`
	BlockRemain      *int32  `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	PreloadRemain    *int32  `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RefreshUrlQuota  *int32  `json:"RefreshUrlQuota,omitempty" xml:"RefreshUrlQuota,omitempty"`
	PreloadQuota     *int32  `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	RefreshDirQuota  *int32  `json:"RefreshDirQuota,omitempty" xml:"RefreshDirQuota,omitempty"`
	RefreshDirRemain *int32  `json:"RefreshDirRemain,omitempty" xml:"RefreshDirRemain,omitempty"`
}

func (s DescribeCdnUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserQuotaResponseBody) SetBlockQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshUrlRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshUrlRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetDomainQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.DomainQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetBlockRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetPreloadRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRequestId(v string) *DescribeCdnUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshUrlQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshUrlQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetPreloadQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshDirQuota(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshDirQuota = &v
	return s
}

func (s *DescribeCdnUserQuotaResponseBody) SetRefreshDirRemain(v int32) *DescribeCdnUserQuotaResponseBody {
	s.RefreshDirRemain = &v
	return s
}

type DescribeCdnUserQuotaResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeCdnUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserQuotaResponse) SetBody(v *DescribeCdnUserQuotaResponseBody) *DescribeCdnUserQuotaResponse {
	s.Body = v
	return s
}

type DescribeCdnUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnUserResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageRequest) SetOwnerId(v int64) *DescribeCdnUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserResourcePackageRequest) SetSecurityToken(v string) *DescribeCdnUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

type DescribeCdnUserResourcePackageResponseBody struct {
	RequestId            *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePackageInfos *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnUserResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponseBody) SetRequestId(v string) *DescribeCdnUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBody) SetResourcePackageInfos(v *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) *DescribeCdnUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

type DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos struct {
	ResourcePackageInfo []*DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo `json:"ResourcePackageInfo,omitempty" xml:"ResourcePackageInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos) SetResourcePackageInfo(v []*DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfos {
	s.ResourcePackageInfo = v
	return s
}

type DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CurrCapacity  *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	InitCapacity  *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetEndTime(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStatus(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.Status = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetDisplayName(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetStartTime(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCommodityCode(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetCurrCapacity(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInitCapacity(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InitCapacity = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetInstanceId(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo) SetTemplateName(v string) *DescribeCdnUserResourcePackageResponseBodyResourcePackageInfosResourcePackageInfo {
	s.TemplateName = &v
	return s
}

type DescribeCdnUserResourcePackageResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnUserResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeCdnUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnUserResourcePackageResponse) SetBody(v *DescribeCdnUserResourcePackageResponseBody) *DescribeCdnUserResourcePackageResponse {
	s.Body = v
	return s
}

type DescribeCdnWafDomainRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnWafDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnWafDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainRequest) SetOwnerId(v int64) *DescribeCdnWafDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnWafDomainRequest) SetRegionId(v string) *DescribeCdnWafDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCdnWafDomainRequest) SetDomainName(v string) *DescribeCdnWafDomainRequest {
	s.DomainName = &v
	return s
}

type DescribeCdnWafDomainResponseBody struct {
	TotalCount    *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OutPutDomains []*DescribeCdnWafDomainResponseBodyOutPutDomains `json:"OutPutDomains,omitempty" xml:"OutPutDomains,omitempty" type:"Repeated"`
}

func (s DescribeCdnWafDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnWafDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainResponseBody) SetTotalCount(v int32) *DescribeCdnWafDomainResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBody) SetRequestId(v string) *DescribeCdnWafDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBody) SetOutPutDomains(v []*DescribeCdnWafDomainResponseBodyOutPutDomains) *DescribeCdnWafDomainResponseBody {
	s.OutPutDomains = v
	return s
}

type DescribeCdnWafDomainResponseBodyOutPutDomains struct {
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CcStatus  *string `json:"CcStatus,omitempty" xml:"CcStatus,omitempty"`
	WafStatus *string `json:"WafStatus,omitempty" xml:"WafStatus,omitempty"`
}

func (s DescribeCdnWafDomainResponseBodyOutPutDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnWafDomainResponseBodyOutPutDomains) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetAclStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.AclStatus = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.Status = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetDomain(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.Domain = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetCcStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.CcStatus = &v
	return s
}

func (s *DescribeCdnWafDomainResponseBodyOutPutDomains) SetWafStatus(v string) *DescribeCdnWafDomainResponseBodyOutPutDomains {
	s.WafStatus = &v
	return s
}

type DescribeCdnWafDomainResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCdnWafDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCdnWafDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCdnWafDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainResponse) SetHeaders(v map[string]*string) *DescribeCdnWafDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnWafDomainResponse) SetBody(v *DescribeCdnWafDomainResponseBody) *DescribeCdnWafDomainResponse {
	s.Body = v
	return s
}

type DescribeCertificateInfoByIDRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	CertId  *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s DescribeCertificateInfoByIDRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateInfoByIDRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDRequest) SetOwnerId(v int64) *DescribeCertificateInfoByIDRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCertificateInfoByIDRequest) SetCertId(v string) *DescribeCertificateInfoByIDRequest {
	s.CertId = &v
	return s
}

type DescribeCertificateInfoByIDResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertInfos *DescribeCertificateInfoByIDResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
}

func (s DescribeCertificateInfoByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponseBody) SetRequestId(v string) *DescribeCertificateInfoByIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBody) SetCertInfos(v *DescribeCertificateInfoByIDResponseBodyCertInfos) *DescribeCertificateInfoByIDResponseBody {
	s.CertInfos = v
	return s
}

type DescribeCertificateInfoByIDResponseBodyCertInfos struct {
	CertInfo []*DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfos) SetCertInfo(v []*DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) *DescribeCertificateInfoByIDResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo struct {
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertId         *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	DomainList     *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	HttpsCrt       *string `json:"HttpsCrt,omitempty" xml:"HttpsCrt,omitempty"`
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCreateTime(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetHttpsCrt(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.HttpsCrt = &v
	return s
}

type DescribeCertificateInfoByIDResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertificateInfoByIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCertificateInfoByIDResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponse) SetHeaders(v map[string]*string) *DescribeCertificateInfoByIDResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateInfoByIDResponse) SetBody(v *DescribeCertificateInfoByIDResponseBody) *DescribeCertificateInfoByIDResponse {
	s.Body = v
	return s
}

type DescribeConfigOfVersionRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	FunctionId    *int32  `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName  *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	GroupId       *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeConfigOfVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionRequest) SetOwnerId(v int64) *DescribeConfigOfVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConfigOfVersionRequest) SetSecurityToken(v string) *DescribeConfigOfVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeConfigOfVersionRequest) SetVersionId(v string) *DescribeConfigOfVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DescribeConfigOfVersionRequest) SetFunctionId(v int32) *DescribeConfigOfVersionRequest {
	s.FunctionId = &v
	return s
}

func (s *DescribeConfigOfVersionRequest) SetFunctionName(v string) *DescribeConfigOfVersionRequest {
	s.FunctionName = &v
	return s
}

func (s *DescribeConfigOfVersionRequest) SetGroupId(v int64) *DescribeConfigOfVersionRequest {
	s.GroupId = &v
	return s
}

type DescribeConfigOfVersionResponseBody struct {
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionConfigs *DescribeConfigOfVersionResponseBodyVersionConfigs `json:"VersionConfigs,omitempty" xml:"VersionConfigs,omitempty" type:"Struct"`
}

func (s DescribeConfigOfVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionResponseBody) SetRequestId(v string) *DescribeConfigOfVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigOfVersionResponseBody) SetVersionConfigs(v *DescribeConfigOfVersionResponseBodyVersionConfigs) *DescribeConfigOfVersionResponseBody {
	s.VersionConfigs = v
	return s
}

type DescribeConfigOfVersionResponseBodyVersionConfigs struct {
	VersionConfig []*DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Repeated"`
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigs) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigs) SetVersionConfig(v []*DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) *DescribeConfigOfVersionResponseBodyVersionConfigs {
	s.VersionConfig = v
	return s
}

type DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig struct {
	Status       *string                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	ConfigId     *string                                                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionName *string                                                                     `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionArgs *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetStatus(v string) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.Status = &v
	return s
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetConfigId(v string) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetFunctionName(v string) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig) SetFunctionArgs(v *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfig {
	s.FunctionArgs = v
	return s
}

type DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs struct {
	FunctionArg []*DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs) SetFunctionArg(v []*DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeConfigOfVersionResponseBodyVersionConfigsVersionConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeConfigOfVersionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConfigOfVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigOfVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigOfVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigOfVersionResponse) SetHeaders(v map[string]*string) *DescribeConfigOfVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigOfVersionResponse) SetBody(v *DescribeConfigOfVersionResponseBody) *DescribeConfigOfVersionResponse {
	s.Body = v
	return s
}

type DescribeCustomLogConfigRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s DescribeCustomLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomLogConfigRequest) SetOwnerId(v int64) *DescribeCustomLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCustomLogConfigRequest) SetConfigId(v string) *DescribeCustomLogConfigRequest {
	s.ConfigId = &v
	return s
}

type DescribeCustomLogConfigResponseBody struct {
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Sample    *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
}

func (s DescribeCustomLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLogConfigResponseBody) SetTag(v string) *DescribeCustomLogConfigResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) SetRequestId(v string) *DescribeCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) SetRemark(v string) *DescribeCustomLogConfigResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeCustomLogConfigResponseBody) SetSample(v string) *DescribeCustomLogConfigResponseBody {
	s.Sample = &v
	return s
}

type DescribeCustomLogConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLogConfigResponse) SetHeaders(v map[string]*string) *DescribeCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLogConfigResponse) SetBody(v *DescribeCustomLogConfigResponseBody) *DescribeCustomLogConfigResponse {
	s.Body = v
	return s
}

type DescribeDomainAverageResponseTimeRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TimeMerge      *string `json:"TimeMerge,omitempty" xml:"TimeMerge,omitempty"`
	DomainType     *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeDomainAverageResponseTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeRequest) SetOwnerId(v int64) *DescribeDomainAverageResponseTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetTimeMerge(v string) *DescribeDomainAverageResponseTimeRequest {
	s.TimeMerge = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetDomainType(v string) *DescribeDomainAverageResponseTimeRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetDomainName(v string) *DescribeDomainAverageResponseTimeRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetStartTime(v string) *DescribeDomainAverageResponseTimeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetEndTime(v string) *DescribeDomainAverageResponseTimeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetInterval(v string) *DescribeDomainAverageResponseTimeRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetIspNameEn(v string) *DescribeDomainAverageResponseTimeRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeRequest) SetLocationNameEn(v string) *DescribeDomainAverageResponseTimeRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeDomainAverageResponseTimeResponseBody struct {
	EndTime          *string                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime        *string                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId        *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName       *string                                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval     *string                                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	AvgRTPerInterval *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval `json:"AvgRTPerInterval,omitempty" xml:"AvgRTPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainAverageResponseTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetEndTime(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetStartTime(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetRequestId(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetDomainName(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetDataInterval(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetAvgRTPerInterval(v *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) *DescribeDomainAverageResponseTimeResponseBody {
	s.AvgRTPerInterval = v
	return s
}

type DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval struct {
	DataModule []*DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) SetDataModule(v []*DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) SetValue(v string) *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDomainAverageResponseTimeResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainAverageResponseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainAverageResponseTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponse) SetHeaders(v map[string]*string) *DescribeDomainAverageResponseTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponse) SetBody(v *DescribeDomainAverageResponseTimeResponseBody) *DescribeDomainAverageResponseTimeResponse {
	s.Body = v
	return s
}

type DescribeDomainBpsDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataRequest) SetOwnerId(v int64) *DescribeDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetDomainName(v string) *DescribeDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetStartTime(v string) *DescribeDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetEndTime(v string) *DescribeDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetInterval(v string) *DescribeDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetIspNameEn(v string) *DescribeDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeDomainBpsDataResponseBody struct {
	EndTime            *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime          *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IspNameEn          *string                                              `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn     *string                                              `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	DomainName         *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval       *string                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	BpsDataPerInterval *DescribeDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBody) SetEndTime(v string) *DescribeDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetStartTime(v string) *DescribeDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetRequestId(v string) *DescribeDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetIspNameEn(v string) *DescribeDomainBpsDataResponseBody {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetLocationNameEn(v string) *DescribeDomainBpsDataResponseBody {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDomainName(v string) *DescribeDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

type DescribeDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	OverseasValue      *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	HttpsValue         *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	DomesticValue      *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

type DescribeDomainBpsDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataResponse) SetBody(v *DescribeDomainBpsDataResponseBody) *DescribeDomainBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainBpsDataByLayerRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	Layer          *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
}

func (s DescribeDomainBpsDataByLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerRequest) SetOwnerId(v int64) *DescribeDomainBpsDataByLayerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetDomainName(v string) *DescribeDomainBpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetStartTime(v string) *DescribeDomainBpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetEndTime(v string) *DescribeDomainBpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetInterval(v string) *DescribeDomainBpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainBpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainBpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerRequest) SetLayer(v string) *DescribeDomainBpsDataByLayerRequest {
	s.Layer = &v
	return s
}

type DescribeDomainBpsDataByLayerResponseBody struct {
	DataInterval    *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BpsDataInterval *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval `json:"BpsDataInterval,omitempty" xml:"BpsDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainBpsDataByLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeDomainBpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainBpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBody) SetBpsDataInterval(v *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) *DescribeDomainBpsDataByLayerResponseBody {
	s.BpsDataInterval = v
	return s
}

type DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval struct {
	DataModule []*DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval) SetDataModule(v []*DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) *DescribeDomainBpsDataByLayerResponseBodyBpsDataInterval {
	s.DataModule = v
	return s
}

type DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule struct {
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetValue(v string) *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTrafficValue(v string) *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainBpsDataByLayerResponseBodyBpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDomainBpsDataByLayerResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainBpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainBpsDataByLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataByLayerResponse) SetBody(v *DescribeDomainBpsDataByLayerResponseBody) *DescribeDomainBpsDataByLayerResponse {
	s.Body = v
	return s
}

type DescribeDomainBpsDataByTimeStampRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	TimePoint     *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	IspNames      *string `json:"IspNames,omitempty" xml:"IspNames,omitempty"`
	LocationNames *string `json:"LocationNames,omitempty" xml:"LocationNames,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetOwnerId(v int64) *DescribeDomainBpsDataByTimeStampRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetDomainName(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetTimePoint(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetIspNames(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.IspNames = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampRequest) SetLocationNames(v string) *DescribeDomainBpsDataByTimeStampRequest {
	s.LocationNames = &v
	return s
}

type DescribeDomainBpsDataByTimeStampResponseBody struct {
	TimeStamp   *string                                                  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	RequestId   *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName  *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	BpsDataList *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList `json:"BpsDataList,omitempty" xml:"BpsDataList,omitempty" type:"Struct"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetRequestId(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetDomainName(v string) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBody) SetBpsDataList(v *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) *DescribeDomainBpsDataByTimeStampResponseBody {
	s.BpsDataList = v
	return s
}

type DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList struct {
	BpsDataModel []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel `json:"BpsDataModel,omitempty" xml:"BpsDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList) SetBpsDataModel(v []*DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataList {
	s.BpsDataModel = v
	return s
}

type DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel struct {
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	IspName      *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	Bps          *int64  `json:"Bps,omitempty" xml:"Bps,omitempty"`
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetLocationName(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.LocationName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetTimeStamp(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetIspName(v string) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.IspName = &v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel) SetBps(v int64) *DescribeDomainBpsDataByTimeStampResponseBodyBpsDataListBpsDataModel {
	s.Bps = &v
	return s
}

type DescribeDomainBpsDataByTimeStampResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainBpsDataByTimeStampResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainBpsDataByTimeStampResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainBpsDataByTimeStampResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataByTimeStampResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataByTimeStampResponse) SetBody(v *DescribeDomainBpsDataByTimeStampResponseBody) *DescribeDomainBpsDataByTimeStampResponse {
	s.Body = v
	return s
}

type DescribeDomainCcActivityLogRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDomainCcActivityLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCcActivityLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogRequest) SetOwnerId(v int64) *DescribeDomainCcActivityLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetDomainName(v string) *DescribeDomainCcActivityLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetStartTime(v string) *DescribeDomainCcActivityLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetEndTime(v string) *DescribeDomainCcActivityLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetTriggerObject(v string) *DescribeDomainCcActivityLogRequest {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetValue(v string) *DescribeDomainCcActivityLogRequest {
	s.Value = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetRuleName(v string) *DescribeDomainCcActivityLogRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetPageSize(v int64) *DescribeDomainCcActivityLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainCcActivityLogRequest) SetPageNumber(v int64) *DescribeDomainCcActivityLogRequest {
	s.PageNumber = &v
	return s
}

type DescribeDomainCcActivityLogResponseBody struct {
	PageIndex   *int64                                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int64                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int64                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	ActivityLog []*DescribeDomainCcActivityLogResponseBodyActivityLog `json:"ActivityLog,omitempty" xml:"ActivityLog,omitempty" type:"Repeated"`
}

func (s DescribeDomainCcActivityLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCcActivityLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogResponseBody) SetPageIndex(v int64) *DescribeDomainCcActivityLogResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetRequestId(v string) *DescribeDomainCcActivityLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetPageSize(v int64) *DescribeDomainCcActivityLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetTotal(v int64) *DescribeDomainCcActivityLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBody) SetActivityLog(v []*DescribeDomainCcActivityLogResponseBodyActivityLog) *DescribeDomainCcActivityLogResponseBody {
	s.ActivityLog = v
	return s
}

type DescribeDomainCcActivityLogResponseBodyActivityLog struct {
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Ttl           *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	TriggerObject *string `json:"TriggerObject,omitempty" xml:"TriggerObject,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeDomainCcActivityLogResponseBodyActivityLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCcActivityLogResponseBodyActivityLog) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetValue(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.Value = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetTtl(v int64) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.Ttl = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetAction(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.Action = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetTriggerObject(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.TriggerObject = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetTimeStamp(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetDomainName(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCcActivityLogResponseBodyActivityLog) SetRuleName(v string) *DescribeDomainCcActivityLogResponseBodyActivityLog {
	s.RuleName = &v
	return s
}

type DescribeDomainCcActivityLogResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainCcActivityLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainCcActivityLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCcActivityLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcActivityLogResponse) SetHeaders(v map[string]*string) *DescribeDomainCcActivityLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCcActivityLogResponse) SetBody(v *DescribeDomainCcActivityLogResponseBody) *DescribeDomainCcActivityLogResponse {
	s.Body = v
	return s
}

type DescribeDomainCertificateInfoRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainCertificateInfoRequest) SetDomainName(v string) *DescribeDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

type DescribeDomainCertificateInfoResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertInfos *DescribeDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
}

func (s DescribeDomainCertificateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeDomainCertificateInfoResponseBodyCertInfos) *DescribeDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

type DescribeDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertExpireTime          *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertLife                *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CertUpdateTime          *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	CertDomainName          *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	CertOrg                 *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CertStartTime           *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	DomainCnameStatus       *string `json:"DomainCnameStatus,omitempty" xml:"DomainCnameStatus,omitempty"`
	ServerCertificate       *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainCnameStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainCnameStatus = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificate(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificate = &v
	return s
}

type DescribeDomainCertificateInfoResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCertificateInfoResponse) SetBody(v *DescribeDomainCertificateInfoResponseBody) *DescribeDomainCertificateInfoResponse {
	s.Body = v
	return s
}

type DescribeDomainCustomLogConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainCustomLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCustomLogConfigRequest) SetOwnerId(v int64) *DescribeDomainCustomLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainCustomLogConfigRequest) SetDomainName(v string) *DescribeDomainCustomLogConfigRequest {
	s.DomainName = &v
	return s
}

type DescribeDomainCustomLogConfigResponseBody struct {
	ConfigId  *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Sample    *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
}

func (s DescribeDomainCustomLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetConfigId(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.ConfigId = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetTag(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetRequestId(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetRemark(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponseBody) SetSample(v string) *DescribeDomainCustomLogConfigResponseBody {
	s.Sample = &v
	return s
}

type DescribeDomainCustomLogConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainCustomLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCustomLogConfigResponse) SetHeaders(v map[string]*string) *DescribeDomainCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCustomLogConfigResponse) SetBody(v *DescribeDomainCustomLogConfigResponseBody) *DescribeDomainCustomLogConfigResponse {
	s.Body = v
	return s
}

type DescribeDomainDetailDataByLayerRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Field          *string `json:"Field,omitempty" xml:"Field,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	Layer          *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
}

func (s DescribeDomainDetailDataByLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerRequest) SetOwnerId(v int64) *DescribeDomainDetailDataByLayerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetField(v string) *DescribeDomainDetailDataByLayerRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetDomainName(v string) *DescribeDomainDetailDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetStartTime(v string) *DescribeDomainDetailDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetEndTime(v string) *DescribeDomainDetailDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainDetailDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainDetailDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerRequest) SetLayer(v string) *DescribeDomainDetailDataByLayerRequest {
	s.Layer = &v
	return s
}

type DescribeDomainDetailDataByLayerResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDomainDetailDataByLayerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDomainDetailDataByLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainDetailDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBody) SetData(v *DescribeDomainDetailDataByLayerResponseBodyData) *DescribeDomainDetailDataByLayerResponseBody {
	s.Data = v
	return s
}

type DescribeDomainDetailDataByLayerResponseBodyData struct {
	DataModule []*DescribeDomainDetailDataByLayerResponseBodyDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainDetailDataByLayerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponseBodyData) SetDataModule(v []*DescribeDomainDetailDataByLayerResponseBodyDataDataModule) *DescribeDomainDetailDataByLayerResponseBodyData {
	s.DataModule = v
	return s
}

type DescribeDomainDetailDataByLayerResponseBodyDataDataModule struct {
	Traf       *int64   `json:"Traf,omitempty" xml:"Traf,omitempty"`
	Qps        *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Ipv6Qps    *float32 `json:"Ipv6Qps,omitempty" xml:"Ipv6Qps,omitempty"`
	Ipv6Bps    *float32 `json:"Ipv6Bps,omitempty" xml:"Ipv6Bps,omitempty"`
	Acc        *int64   `json:"Acc,omitempty" xml:"Acc,omitempty"`
	Ipv6Traf   *int64   `json:"Ipv6Traf,omitempty" xml:"Ipv6Traf,omitempty"`
	Ipv6Acc    *int64   `json:"Ipv6Acc,omitempty" xml:"Ipv6Acc,omitempty"`
	TimeStamp  *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	HttpCode   *string  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Bps        *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	DomainName *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainDetailDataByLayerResponseBodyDataDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetTraf(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Traf = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetQps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Qps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Qps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Qps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Bps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Bps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetAcc(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Acc = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Traf(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Traf = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Acc(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Acc = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetTimeStamp(v string) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetHttpCode(v string) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.HttpCode = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetBps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetDomainName(v string) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.DomainName = &v
	return s
}

type DescribeDomainDetailDataByLayerResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainDetailDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainDetailDataByLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainDetailDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponse) SetBody(v *DescribeDomainDetailDataByLayerResponseBody) *DescribeDomainDetailDataByLayerResponse {
	s.Body = v
	return s
}

type DescribeDomainFileSizeProportionDataRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainFileSizeProportionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetOwnerId(v int64) *DescribeDomainFileSizeProportionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetSecurityToken(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetDomainName(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetStartTime(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataRequest) SetEndTime(v string) *DescribeDomainFileSizeProportionDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBody struct {
	EndTime                        *string                                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime                      *string                                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId                      *string                                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName                     *string                                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval                   *string                                                                         `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	FileSizeProportionDataInterval *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval `json:"FileSizeProportionDataInterval,omitempty" xml:"FileSizeProportionDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainFileSizeProportionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetEndTime(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetStartTime(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetRequestId(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetDomainName(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetDataInterval(v string) *DescribeDomainFileSizeProportionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBody) SetFileSizeProportionDataInterval(v *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) *DescribeDomainFileSizeProportionDataResponseBody {
	s.FileSizeProportionDataInterval = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval struct {
	UsageData []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval) SetUsageData(v []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataInterval {
	s.UsageData = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData struct {
	TimeStamp *string                                                                                       `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData) SetValue(v *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageData {
	s.Value = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue struct {
	FileSizeProportionData []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData `json:"FileSizeProportionData,omitempty" xml:"FileSizeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue) SetFileSizeProportionData(v []*DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValue {
	s.FileSizeProportionData = v
	return s
}

type DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData struct {
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	FileSize   *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) SetProportion(v string) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData) SetFileSize(v string) *DescribeDomainFileSizeProportionDataResponseBodyFileSizeProportionDataIntervalUsageDataValueFileSizeProportionData {
	s.FileSize = &v
	return s
}

type DescribeDomainFileSizeProportionDataResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainFileSizeProportionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainFileSizeProportionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainFileSizeProportionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainFileSizeProportionDataResponse) SetHeaders(v map[string]*string) *DescribeDomainFileSizeProportionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainFileSizeProportionDataResponse) SetBody(v *DescribeDomainFileSizeProportionDataResponseBody) *DescribeDomainFileSizeProportionDataResponse {
	s.Body = v
	return s
}

type DescribeDomainHitRateDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataRequest) SetOwnerId(v int64) *DescribeDomainHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetDomainName(v string) *DescribeDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetStartTime(v string) *DescribeDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetEndTime(v string) *DescribeDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHitRateDataRequest) SetInterval(v string) *DescribeDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainHitRateDataResponseBody struct {
	EndTime         *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName      *string                                               `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval    *string                                               `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	HitRateInterval *DescribeDomainHitRateDataResponseBodyHitRateInterval `json:"HitRateInterval,omitempty" xml:"HitRateInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetHitRateInterval(v *DescribeDomainHitRateDataResponseBodyHitRateInterval) *DescribeDomainHitRateDataResponseBody {
	s.HitRateInterval = v
	return s
}

type DescribeDomainHitRateDataResponseBodyHitRateInterval struct {
	DataModule []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainHitRateDataResponseBodyHitRateInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBodyHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateInterval) SetDataModule(v []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) *DescribeDomainHitRateDataResponseBodyHitRateInterval {
	s.DataModule = v
	return s
}

type DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
}

func (s DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetValue(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetHttpsValue(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.HttpsValue = &v
	return s
}

type DescribeDomainHitRateDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHitRateDataResponse) SetBody(v *DescribeDomainHitRateDataResponseBody) *DescribeDomainHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDomainHttpCodeDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDomainHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataRequest) SetInterval(v string) *DescribeDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainHttpCodeDataResponseBody struct {
	EndTime      *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval *string                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	HttpCodeData *DescribeDomainHttpCodeDataResponseBodyHttpCodeData `json:"HttpCodeData,omitempty" xml:"HttpCodeData,omitempty" type:"Struct"`
}

func (s DescribeDomainHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetHttpCodeData(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainHttpCodeDataResponseBody {
	s.HttpCodeData = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeData struct {
	UsageData []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) SetUsageData(v []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) *DescribeDomainHttpCodeDataResponseBodyHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData struct {
	TimeStamp *string                                                           `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetValue(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue struct {
	CodeProportionData []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData `json:"CodeProportionData,omitempty" xml:"CodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) SetCodeProportionData(v []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	s.CodeProportionData = v
	return s
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCode(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetProportion(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCount(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Count = &v
	return s
}

type DescribeDomainHttpCodeDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponse) SetBody(v *DescribeDomainHttpCodeDataResponseBody) *DescribeDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDomainHttpCodeDataByLayerRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	Layer          *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
}

func (s DescribeDomainHttpCodeDataByLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetOwnerId(v int64) *DescribeDomainHttpCodeDataByLayerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetDomainName(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetStartTime(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetEndTime(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetInterval(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerRequest) SetLayer(v string) *DescribeDomainHttpCodeDataByLayerRequest {
	s.Layer = &v
	return s
}

type DescribeDomainHttpCodeDataByLayerResponseBody struct {
	DataInterval         *string                                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RequestId            *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCodeDataInterval *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval `json:"HttpCodeDataInterval,omitempty" xml:"HttpCodeDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainHttpCodeDataByLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) SetDataInterval(v string) *DescribeDomainHttpCodeDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainHttpCodeDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBody) SetHttpCodeDataInterval(v *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) *DescribeDomainHttpCodeDataByLayerResponseBody {
	s.HttpCodeDataInterval = v
	return s
}

type DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval struct {
	DataModule []*DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval) SetDataModule(v []*DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataInterval {
	s.DataModule = v
	return s
}

type DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TotalValue *string `json:"TotalValue,omitempty" xml:"TotalValue,omitempty"`
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetValue(v string) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule) SetTotalValue(v string) *DescribeDomainHttpCodeDataByLayerResponseBodyHttpCodeDataIntervalDataModule {
	s.TotalValue = &v
	return s
}

type DescribeDomainHttpCodeDataByLayerResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainHttpCodeDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainHttpCodeDataByLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainHttpCodeDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainHttpCodeDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainHttpCodeDataByLayerResponse) SetBody(v *DescribeDomainHttpCodeDataByLayerResponseBody) *DescribeDomainHttpCodeDataByLayerResponse {
	s.Body = v
	return s
}

type DescribeDomainISPDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainISPDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataRequest) SetOwnerId(v int64) *DescribeDomainISPDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetDomainName(v string) *DescribeDomainISPDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetStartTime(v string) *DescribeDomainISPDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainISPDataRequest) SetEndTime(v string) *DescribeDomainISPDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainISPDataResponseBody struct {
	EndTime      *string                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval *string                                 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	Value        *DescribeDomainISPDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainISPDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBody) SetEndTime(v string) *DescribeDomainISPDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetStartTime(v string) *DescribeDomainISPDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetRequestId(v string) *DescribeDomainISPDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetDomainName(v string) *DescribeDomainISPDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetDataInterval(v string) *DescribeDomainISPDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetValue(v *DescribeDomainISPDataResponseBodyValue) *DescribeDomainISPDataResponseBody {
	s.Value = v
	return s
}

type DescribeDomainISPDataResponseBodyValue struct {
	ISPProportionData []*DescribeDomainISPDataResponseBodyValueISPProportionData `json:"ISPProportionData,omitempty" xml:"ISPProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainISPDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBodyValue) SetISPProportionData(v []*DescribeDomainISPDataResponseBodyValueISPProportionData) *DescribeDomainISPDataResponseBodyValue {
	s.ISPProportionData = v
	return s
}

type DescribeDomainISPDataResponseBodyValueISPProportionData struct {
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	IspEname        *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	ISP             *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
}

func (s DescribeDomainISPDataResponseBodyValueISPProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponseBodyValueISPProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetTotalQuery(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetTotalBytes(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseTime(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetReqErrRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgObjectSize(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetBps(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetQps(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetProportion(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetIspEname(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetISP(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ISP = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetBytesProportion(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.BytesProportion = &v
	return s
}

type DescribeDomainISPDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainISPDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainISPDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainISPDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponse) SetHeaders(v map[string]*string) *DescribeDomainISPDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainISPDataResponse) SetBody(v *DescribeDomainISPDataResponseBody) *DescribeDomainISPDataResponse {
	s.Body = v
	return s
}

type DescribeDomainMax95BpsDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TimePoint  *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	Cycle      *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
}

func (s DescribeDomainMax95BpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainMax95BpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataRequest) SetOwnerId(v int64) *DescribeDomainMax95BpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetDomainName(v string) *DescribeDomainMax95BpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetStartTime(v string) *DescribeDomainMax95BpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetEndTime(v string) *DescribeDomainMax95BpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetTimePoint(v string) *DescribeDomainMax95BpsDataRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeDomainMax95BpsDataRequest) SetCycle(v string) *DescribeDomainMax95BpsDataRequest {
	s.Cycle = &v
	return s
}

type DescribeDomainMax95BpsDataResponseBody struct {
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomesticMax95Bps *string `json:"DomesticMax95Bps,omitempty" xml:"DomesticMax95Bps,omitempty"`
	OverseasMax95Bps *string `json:"OverseasMax95Bps,omitempty" xml:"OverseasMax95Bps,omitempty"`
	Max95Bps         *string `json:"Max95Bps,omitempty" xml:"Max95Bps,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainMax95BpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainMax95BpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetEndTime(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetStartTime(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetRequestId(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetDomesticMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.DomesticMax95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetOverseasMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.OverseasMax95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.Max95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetDomainName(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.DomainName = &v
	return s
}

type DescribeDomainMax95BpsDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainMax95BpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainMax95BpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainMax95BpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainMax95BpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainMax95BpsDataResponse) SetBody(v *DescribeDomainMax95BpsDataResponseBody) *DescribeDomainMax95BpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainNamesOfVersionRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainNamesOfVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesOfVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesOfVersionRequest) SetOwnerId(v int64) *DescribeDomainNamesOfVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainNamesOfVersionRequest) SetVersionId(v string) *DescribeDomainNamesOfVersionRequest {
	s.VersionId = &v
	return s
}

func (s *DescribeDomainNamesOfVersionRequest) SetPageIndex(v int32) *DescribeDomainNamesOfVersionRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeDomainNamesOfVersionRequest) SetPageSize(v string) *DescribeDomainNamesOfVersionRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainNamesOfVersionResponseBody struct {
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Contents   []*DescribeDomainNamesOfVersionResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s DescribeDomainNamesOfVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesOfVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesOfVersionResponseBody) SetTotalCount(v int32) *DescribeDomainNamesOfVersionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainNamesOfVersionResponseBody) SetRequestId(v string) *DescribeDomainNamesOfVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainNamesOfVersionResponseBody) SetContents(v []*DescribeDomainNamesOfVersionResponseBodyContents) *DescribeDomainNamesOfVersionResponseBody {
	s.Contents = v
	return s
}

type DescribeDomainNamesOfVersionResponseBodyContents struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainId   *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
}

func (s DescribeDomainNamesOfVersionResponseBodyContents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesOfVersionResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesOfVersionResponseBodyContents) SetDomainName(v string) *DescribeDomainNamesOfVersionResponseBodyContents {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainNamesOfVersionResponseBodyContents) SetDomainId(v string) *DescribeDomainNamesOfVersionResponseBodyContents {
	s.DomainId = &v
	return s
}

type DescribeDomainNamesOfVersionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainNamesOfVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainNamesOfVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainNamesOfVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainNamesOfVersionResponse) SetHeaders(v map[string]*string) *DescribeDomainNamesOfVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainNamesOfVersionResponse) SetBody(v *DescribeDomainNamesOfVersionResponseBody) *DescribeDomainNamesOfVersionResponse {
	s.Body = v
	return s
}

type DescribeDomainPathDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainPathDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPathDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataRequest) SetOwnerId(v int64) *DescribeDomainPathDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetPageNumber(v int32) *DescribeDomainPathDataRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetPageSize(v int32) *DescribeDomainPathDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetPath(v string) *DescribeDomainPathDataRequest {
	s.Path = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetStartTime(v string) *DescribeDomainPathDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetEndTime(v string) *DescribeDomainPathDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPathDataRequest) SetDomainName(v string) *DescribeDomainPathDataRequest {
	s.DomainName = &v
	return s
}

type DescribeDomainPathDataResponseBody struct {
	EndTime             *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime           *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PageSize            *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber          *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount          *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DomainName          *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval        *string                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	PathDataPerInterval *DescribeDomainPathDataResponseBodyPathDataPerInterval `json:"PathDataPerInterval,omitempty" xml:"PathDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainPathDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPathDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponseBody) SetEndTime(v string) *DescribeDomainPathDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetStartTime(v string) *DescribeDomainPathDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetPageSize(v int32) *DescribeDomainPathDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetPageNumber(v int32) *DescribeDomainPathDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetTotalCount(v int32) *DescribeDomainPathDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetDomainName(v string) *DescribeDomainPathDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetDataInterval(v string) *DescribeDomainPathDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainPathDataResponseBody) SetPathDataPerInterval(v *DescribeDomainPathDataResponseBodyPathDataPerInterval) *DescribeDomainPathDataResponseBody {
	s.PathDataPerInterval = v
	return s
}

type DescribeDomainPathDataResponseBodyPathDataPerInterval struct {
	UsageData []*DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainPathDataResponseBodyPathDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPathDataResponseBodyPathDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerInterval) SetUsageData(v []*DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) *DescribeDomainPathDataResponseBodyPathDataPerInterval {
	s.UsageData = v
	return s
}

type DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Time    *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Acc     *int32  `json:"Acc,omitempty" xml:"Acc,omitempty"`
	Traffic *int32  `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetPath(v string) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Path = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetTime(v string) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Time = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetAcc(v int32) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Acc = &v
	return s
}

func (s *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData) SetTraffic(v int32) *DescribeDomainPathDataResponseBodyPathDataPerIntervalUsageData {
	s.Traffic = &v
	return s
}

type DescribeDomainPathDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainPathDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainPathDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPathDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainPathDataResponse) SetHeaders(v map[string]*string) *DescribeDomainPathDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainPathDataResponse) SetBody(v *DescribeDomainPathDataResponseBody) *DescribeDomainPathDataResponse {
	s.Body = v
	return s
}

type DescribeDomainPvDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainPvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataRequest) SetOwnerId(v int64) *DescribeDomainPvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainPvDataRequest) SetDomainName(v string) *DescribeDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPvDataRequest) SetStartTime(v string) *DescribeDomainPvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPvDataRequest) SetEndTime(v string) *DescribeDomainPvDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainPvDataResponseBody struct {
	EndTime        *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName     *string                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval   *string                                         `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	PvDataInterval *DescribeDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainPvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponseBody) SetEndTime(v string) *DescribeDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetStartTime(v string) *DescribeDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetRequestId(v string) *DescribeDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetDomainName(v string) *DescribeDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetDataInterval(v string) *DescribeDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetPvDataInterval(v *DescribeDomainPvDataResponseBodyPvDataInterval) *DescribeDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

type DescribeDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainPvDataResponseBodyPvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

type DescribeDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

type DescribeDomainPvDataResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainPvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainPvDataResponse) SetBody(v *DescribeDomainPvDataResponseBody) *DescribeDomainPvDataResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeDomainQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataRequest) SetOwnerId(v int64) *DescribeDomainQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetDomainName(v string) *DescribeDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetStartTime(v string) *DescribeDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetEndTime(v string) *DescribeDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetInterval(v string) *DescribeDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetIspNameEn(v string) *DescribeDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeDomainQpsDataResponseBody struct {
	EndTime         *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName      *string                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval    *string                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	QpsDataInterval *DescribeDomainQpsDataResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBody) SetEndTime(v string) *DescribeDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetStartTime(v string) *DescribeDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetRequestId(v string) *DescribeDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetDomainName(v string) *DescribeDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBody) SetQpsDataInterval(v *DescribeDomainQpsDataResponseBodyQpsDataInterval) *DescribeDomainQpsDataResponseBody {
	s.QpsDataInterval = v
	return s
}

type DescribeDomainQpsDataResponseBodyQpsDataInterval struct {
	DataModule []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsDataResponseBodyQpsDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataInterval) SetDataModule(v []*DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) *DescribeDomainQpsDataResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

type DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule struct {
	AccValue              *string `json:"AccValue,omitempty" xml:"AccValue,omitempty"`
	AccDomesticValue      *string `json:"AccDomesticValue,omitempty" xml:"AccDomesticValue,omitempty"`
	AccOverseasValue      *string `json:"AccOverseasValue,omitempty" xml:"AccOverseasValue,omitempty"`
	HttpsValue            *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	HttpsOverseasValue    *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	DomesticValue         *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	HttpsAccOverseasValue *string `json:"HttpsAccOverseasValue,omitempty" xml:"HttpsAccOverseasValue,omitempty"`
	HttpsDomesticValue    *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	HttpsAccValue         *string `json:"HttpsAccValue,omitempty" xml:"HttpsAccValue,omitempty"`
	Value                 *string `json:"Value,omitempty" xml:"Value,omitempty"`
	OverseasValue         *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	TimeStamp             *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	HttpsAccDomesticValue *string `json:"HttpsAccDomesticValue,omitempty" xml:"HttpsAccDomesticValue,omitempty"`
}

func (s DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule) SetHttpsAccDomesticValue(v string) *DescribeDomainQpsDataResponseBodyQpsDataIntervalDataModule {
	s.HttpsAccDomesticValue = &v
	return s
}

type DescribeDomainQpsDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsDataResponse) SetBody(v *DescribeDomainQpsDataResponseBody) *DescribeDomainQpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsDataByLayerRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	Layer          *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
}

func (s DescribeDomainQpsDataByLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerRequest) SetOwnerId(v int64) *DescribeDomainQpsDataByLayerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetDomainName(v string) *DescribeDomainQpsDataByLayerRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetStartTime(v string) *DescribeDomainQpsDataByLayerRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetEndTime(v string) *DescribeDomainQpsDataByLayerRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetInterval(v string) *DescribeDomainQpsDataByLayerRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetIspNameEn(v string) *DescribeDomainQpsDataByLayerRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetLocationNameEn(v string) *DescribeDomainQpsDataByLayerRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerRequest) SetLayer(v string) *DescribeDomainQpsDataByLayerRequest {
	s.Layer = &v
	return s
}

type DescribeDomainQpsDataByLayerResponseBody struct {
	EndTime         *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Layer           *string                                                  `json:"Layer,omitempty" xml:"Layer,omitempty"`
	DomainName      *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval    *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	QpsDataInterval *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval `json:"QpsDataInterval,omitempty" xml:"QpsDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainQpsDataByLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetEndTime(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetStartTime(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetLayer(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.Layer = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetDomainName(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetDataInterval(v string) *DescribeDomainQpsDataByLayerResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBody) SetQpsDataInterval(v *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) *DescribeDomainQpsDataByLayerResponseBody {
	s.QpsDataInterval = v
	return s
}

type DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval struct {
	DataModule []*DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval) SetDataModule(v []*DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) *DescribeDomainQpsDataByLayerResponseBodyQpsDataInterval {
	s.DataModule = v
	return s
}

type DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule struct {
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
	AccValue         *string `json:"AccValue,omitempty" xml:"AccValue,omitempty"`
	AccDomesticValue *string `json:"AccDomesticValue,omitempty" xml:"AccDomesticValue,omitempty"`
	OverseasValue    *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	AccOverseasValue *string `json:"AccOverseasValue,omitempty" xml:"AccOverseasValue,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	DomesticValue    *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccDomesticValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccDomesticValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetOverseasValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetAccOverseasValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.AccOverseasValue = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetTimeStamp(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule) SetDomesticValue(v string) *DescribeDomainQpsDataByLayerResponseBodyQpsDataIntervalDataModule {
	s.DomesticValue = &v
	return s
}

type DescribeDomainQpsDataByLayerResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainQpsDataByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQpsDataByLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsDataByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsDataByLayerResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsDataByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsDataByLayerResponse) SetBody(v *DescribeDomainQpsDataByLayerResponseBody) *DescribeDomainQpsDataByLayerResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeBpsDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeBpsDataResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDomainRealTimeBpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBody) SetData(v *DescribeDomainRealTimeBpsDataResponseBodyData) *DescribeDomainRealTimeBpsDataResponseBody {
	s.Data = v
	return s
}

type DescribeDomainRealTimeBpsDataResponseBodyData struct {
	BpsModel []*DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel `json:"BpsModel,omitempty" xml:"BpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeBpsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyData) SetBpsModel(v []*DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) *DescribeDomainRealTimeBpsDataResponseBodyData {
	s.BpsModel = v
	return s
}

type DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel struct {
	Bps       *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) SetBps(v float32) *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.Bps = &v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel) SetTimeStamp(v string) *DescribeDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeBpsDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeBpsDataResponse) SetBody(v *DescribeDomainRealTimeBpsDataResponseBody) *DescribeDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeByteHitRateDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeByteHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetDomainName(v string) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetStartTime(v string) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataRequest) SetEndTime(v string) *DescribeDomainRealTimeByteHitRateDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeByteHitRateDataResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDomainRealTimeByteHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeByteHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBody) SetData(v *DescribeDomainRealTimeByteHitRateDataResponseBodyData) *DescribeDomainRealTimeByteHitRateDataResponseBody {
	s.Data = v
	return s
}

type DescribeDomainRealTimeByteHitRateDataResponseBodyData struct {
	ByteHitRateDataModel []*DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel `json:"ByteHitRateDataModel,omitempty" xml:"ByteHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyData) SetByteHitRateDataModel(v []*DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) *DescribeDomainRealTimeByteHitRateDataResponseBodyData {
	s.ByteHitRateDataModel = v
	return s
}

type DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel struct {
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	TimeStamp   *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetByteHitRate(v float32) *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetTimeStamp(v string) *DescribeDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeByteHitRateDataResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeByteHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeByteHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeByteHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeByteHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeByteHitRateDataResponse) SetBody(v *DescribeDomainRealTimeByteHitRateDataResponseBody) *DescribeDomainRealTimeByteHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeDetailDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Field          *string `json:"Field,omitempty" xml:"Field,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	Merge          *string `json:"Merge,omitempty" xml:"Merge,omitempty"`
	MergeLocIsp    *string `json:"MergeLocIsp,omitempty" xml:"MergeLocIsp,omitempty"`
}

func (s DescribeDomainRealTimeDetailDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeDetailDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeDetailDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetDomainName(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetStartTime(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetEndTime(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetField(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetMerge(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.Merge = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataRequest) SetMergeLocIsp(v string) *DescribeDomainRealTimeDetailDataRequest {
	s.MergeLocIsp = &v
	return s
}

type DescribeDomainRealTimeDetailDataResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainRealTimeDetailDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) SetData(v string) *DescribeDomainRealTimeDetailDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeDetailDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainRealTimeDetailDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeDetailDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeDetailDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeDetailDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeDetailDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeDetailDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponse) SetBody(v *DescribeDomainRealTimeDetailDataResponseBody) *DescribeDomainRealTimeDetailDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeHttpCodeDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeDomainRealTimeHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeDomainRealTimeHttpCodeDataResponseBody struct {
	EndTime              *string                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime            *string                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId            *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName           *string                                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval         *string                                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RealTimeHttpCodeData *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData `json:"RealTimeHttpCodeData,omitempty" xml:"RealTimeHttpCodeData,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBody) SetRealTimeHttpCodeData(v *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeDomainRealTimeHttpCodeDataResponseBody {
	s.RealTimeHttpCodeData = v
	return s
}

type DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData struct {
	UsageData []*DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) SetUsageData(v []*DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData struct {
	TimeStamp *string                                                                           `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetValue(v *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue struct {
	RealTimeCodeProportionData []*DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData `json:"RealTimeCodeProportionData,omitempty" xml:"RealTimeCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) SetRealTimeCodeProportionData(v []*DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	s.RealTimeCodeProportionData = v
	return s
}

type DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCode(v string) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetProportion(v string) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCount(v string) *DescribeDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Count = &v
	return s
}

type DescribeDomainRealTimeHttpCodeDataResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeDomainRealTimeHttpCodeDataResponseBody) *DescribeDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealtimeLogDeliveryRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DescribeDomainRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryRequest) SetDomain(v string) *DescribeDomainRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

type DescribeDomainRealtimeLogDeliveryResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Logstore  *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeDomainRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetStatus(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetLogstore(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Logstore = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetProject(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetRegion(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Region = &v
	return s
}

type DescribeDomainRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeDomainRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponse) SetBody(v *DescribeDomainRealtimeLogDeliveryResponseBody) *DescribeDomainRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeQpsDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetDomainName(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetStartTime(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataRequest) SetEndTime(v string) *DescribeDomainRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeQpsDataResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDomainRealTimeQpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBody) SetData(v *DescribeDomainRealTimeQpsDataResponseBodyData) *DescribeDomainRealTimeQpsDataResponseBody {
	s.Data = v
	return s
}

type DescribeDomainRealTimeQpsDataResponseBodyData struct {
	QpsModel []*DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel `json:"QpsModel,omitempty" xml:"QpsModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeQpsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyData) SetQpsModel(v []*DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) *DescribeDomainRealTimeQpsDataResponseBodyData {
	s.QpsModel = v
	return s
}

type DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel struct {
	Qps       *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) SetQps(v float32) *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.Qps = &v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel) SetTimeStamp(v string) *DescribeDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeQpsDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeQpsDataResponse) SetBody(v *DescribeDomainRealTimeQpsDataResponseBody) *DescribeDomainRealTimeQpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeReqHitRateDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeReqHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetDomainName(v string) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetStartTime(v string) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataRequest) SetEndTime(v string) *DescribeDomainRealTimeReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeReqHitRateDataResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDomainRealTimeReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBody) SetData(v *DescribeDomainRealTimeReqHitRateDataResponseBodyData) *DescribeDomainRealTimeReqHitRateDataResponseBody {
	s.Data = v
	return s
}

type DescribeDomainRealTimeReqHitRateDataResponseBodyData struct {
	ReqHitRateDataModel []*DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel `json:"ReqHitRateDataModel,omitempty" xml:"ReqHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyData) SetReqHitRateDataModel(v []*DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) *DescribeDomainRealTimeReqHitRateDataResponseBodyData {
	s.ReqHitRateDataModel = v
	return s
}

type DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel struct {
	ReqHitRate *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TimeStamp  *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetReqHitRate(v float32) *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetTimeStamp(v string) *DescribeDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeReqHitRateDataResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeReqHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeReqHitRateDataResponse) SetBody(v *DescribeDomainRealTimeReqHitRateDataResponseBody) *DescribeDomainRealTimeReqHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeSrcBpsDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetDomainName(v string) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetStartTime(v string) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataRequest) SetEndTime(v string) *DescribeDomainRealTimeSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeSrcBpsDataResponseBody struct {
	EndTime                       *string                                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime                     *string                                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId                     *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName                    *string                                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval                  *string                                                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RealTimeSrcBpsDataPerInterval *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval,omitempty" xml:"RealTimeSrcBpsDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetRealTimeSrcBpsDataPerInterval(v *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.RealTimeSrcBpsDataPerInterval = v
	return s
}

type DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval struct {
	DataModule []*DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) SetDataModule(v []*DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeSrcBpsDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponse) SetBody(v *DescribeDomainRealTimeSrcBpsDataResponseBody) *DescribeDomainRealTimeSrcBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeSrcHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBody struct {
	EndTime                 *string                                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime               *string                                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId               *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName              *string                                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval            *string                                                                   `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RealTimeSrcHttpCodeData *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData `json:"RealTimeSrcHttpCodeData,omitempty" xml:"RealTimeSrcHttpCodeData,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetRealTimeSrcHttpCodeData(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RealTimeSrcHttpCodeData = v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData struct {
	UsageData []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) SetUsageData(v []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData struct {
	TimeStamp *string                                                                                 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetValue(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue struct {
	RealTimeSrcCodeProportionData []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData `json:"RealTimeSrcCodeProportionData,omitempty" xml:"RealTimeSrcCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) SetRealTimeSrcCodeProportionData(v []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue {
	s.RealTimeSrcCodeProportionData = v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCode(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetProportion(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCount(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Count = &v
	return s
}

type DescribeDomainRealTimeSrcHttpCodeDataResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeSrcHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponse) SetBody(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) *DescribeDomainRealTimeSrcHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeSrcTrafficDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetDomainName(v string) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetStartTime(v string) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataRequest) SetEndTime(v string) *DescribeDomainRealTimeSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeSrcTrafficDataResponseBody struct {
	EndTime                           *string                                                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime                         *string                                                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId                         *string                                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName                        *string                                                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval                      *string                                                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RealTimeSrcTrafficDataPerInterval *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval,omitempty" xml:"RealTimeSrcTrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetRealTimeSrcTrafficDataPerInterval(v *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.RealTimeSrcTrafficDataPerInterval = v
	return s
}

type DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval struct {
	DataModule []*DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) SetDataModule(v []*DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeSrcTrafficDataResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponse) SetBody(v *DescribeDomainRealTimeSrcTrafficDataResponseBody) *DescribeDomainRealTimeSrcTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRealTimeTrafficDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRealTimeTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetOwnerId(v int64) *DescribeDomainRealTimeTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetIspNameEn(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetLocationNameEn(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRealTimeTrafficDataResponseBody struct {
	EndTime                        *string                                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime                      *string                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId                      *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName                     *string                                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval                   *string                                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RealTimeTrafficDataPerInterval *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

type DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDomainRealTimeTrafficDataResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRealTimeTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponse) SetBody(v *DescribeDomainRealTimeTrafficDataResponseBody) *DescribeDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDomainRegionDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainRegionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataRequest) SetOwnerId(v int64) *DescribeDomainRegionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetDomainName(v string) *DescribeDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetStartTime(v string) *DescribeDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRegionDataRequest) SetEndTime(v string) *DescribeDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainRegionDataResponseBody struct {
	EndTime      *string                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval *string                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	Value        *DescribeDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainRegionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBody) SetEndTime(v string) *DescribeDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetStartTime(v string) *DescribeDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetRequestId(v string) *DescribeDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetDomainName(v string) *DescribeDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetValue(v *DescribeDomainRegionDataResponseBodyValue) *DescribeDomainRegionDataResponseBody {
	s.Value = v
	return s
}

type DescribeDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRegionDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeDomainRegionDataResponseBodyValueRegionProportionData) *DescribeDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

type DescribeDomainRegionDataResponseBodyValueRegionProportionData struct {
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	RegionEname     *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
}

func (s DescribeDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

type DescribeDomainRegionDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainRegionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRegionDataResponse) SetBody(v *DescribeDomainRegionDataResponseBody) *DescribeDomainRegionDataResponse {
	s.Body = v
	return s
}

type DescribeDomainReqHitRateDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainReqHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataRequest) SetOwnerId(v int64) *DescribeDomainReqHitRateDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetDomainName(v string) *DescribeDomainReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetStartTime(v string) *DescribeDomainReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetEndTime(v string) *DescribeDomainReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataRequest) SetInterval(v string) *DescribeDomainReqHitRateDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainReqHitRateDataResponseBody struct {
	EndTime            *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime          *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName         *string                                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval       *string                                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	ReqHitRateInterval *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval `json:"ReqHitRateInterval,omitempty" xml:"ReqHitRateInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainReqHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetEndTime(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetStartTime(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetDomainName(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetDataInterval(v string) *DescribeDomainReqHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBody) SetReqHitRateInterval(v *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) *DescribeDomainReqHitRateDataResponseBody {
	s.ReqHitRateInterval = v
	return s
}

type DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval struct {
	DataModule []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval) SetDataModule(v []*DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) *DescribeDomainReqHitRateDataResponseBodyReqHitRateInterval {
	s.DataModule = v
	return s
}

type DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetValue(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule) SetHttpsValue(v string) *DescribeDomainReqHitRateDataResponseBodyReqHitRateIntervalDataModule {
	s.HttpsValue = &v
	return s
}

type DescribeDomainReqHitRateDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainReqHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeDomainReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainReqHitRateDataResponse) SetBody(v *DescribeDomainReqHitRateDataResponseBody) *DescribeDomainReqHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeDomainsBySourceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources       *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s DescribeDomainsBySourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceRequest) SetOwnerId(v int64) *DescribeDomainsBySourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) SetSecurityToken(v string) *DescribeDomainsBySourceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDomainsBySourceRequest) SetSources(v string) *DescribeDomainsBySourceRequest {
	s.Sources = &v
	return s
}

type DescribeDomainsBySourceResponseBody struct {
	Sources     *string                                         `json:"Sources,omitempty" xml:"Sources,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainsList *DescribeDomainsBySourceResponseBodyDomainsList `json:"DomainsList,omitempty" xml:"DomainsList,omitempty" type:"Struct"`
}

func (s DescribeDomainsBySourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBody) SetSources(v string) *DescribeDomainsBySourceResponseBody {
	s.Sources = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) SetRequestId(v string) *DescribeDomainsBySourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBody) SetDomainsList(v *DescribeDomainsBySourceResponseBodyDomainsList) *DescribeDomainsBySourceResponseBody {
	s.DomainsList = v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsList struct {
	DomainsData []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData `json:"DomainsData,omitempty" xml:"DomainsData,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsList) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsList) SetDomainsData(v []*DescribeDomainsBySourceResponseBodyDomainsListDomainsData) *DescribeDomainsBySourceResponseBodyDomainsList {
	s.DomainsData = v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsData struct {
	Source      *string                                                               `json:"Source,omitempty" xml:"Source,omitempty"`
	DomainInfos *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos `json:"DomainInfos,omitempty" xml:"DomainInfos,omitempty" type:"Struct"`
	Domains     *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains     `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsData) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetSource(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.Source = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetDomainInfos(v *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.DomainInfos = v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsData) SetDomains(v *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) *DescribeDomainsBySourceResponseBodyDomainsListDomainsData {
	s.Domains = v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos struct {
	DomainInfo []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo `json:"domainInfo,omitempty" xml:"domainInfo,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos) SetDomainInfo(v []*DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfos {
	s.DomainInfo = v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DomainCname *string `json:"DomainCname,omitempty" xml:"DomainCname,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetStatus(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.Status = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetUpdateTime(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetDomainCname(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.DomainCname = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetDomainName(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo) SetCreateTime(v string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomainInfosDomainInfo {
	s.CreateTime = &v
	return s
}

type DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains struct {
	DomainNames []*string `json:"domainNames,omitempty" xml:"domainNames,omitempty" type:"Repeated"`
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains) SetDomainNames(v []*string) *DescribeDomainsBySourceResponseBodyDomainsListDomainsDataDomains {
	s.DomainNames = v
	return s
}

type DescribeDomainsBySourceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainsBySourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsBySourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsBySourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsBySourceResponse) SetHeaders(v map[string]*string) *DescribeDomainsBySourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsBySourceResponse) SetBody(v *DescribeDomainsBySourceResponseBody) *DescribeDomainsBySourceResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcBpsDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainSrcBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataRequest) SetOwnerId(v int64) *DescribeDomainSrcBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetDomainName(v string) *DescribeDomainSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetStartTime(v string) *DescribeDomainSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetEndTime(v string) *DescribeDomainSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataRequest) SetInterval(v string) *DescribeDomainSrcBpsDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainSrcBpsDataResponseBody struct {
	EndTime               *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime             *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName            *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval          *string                                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	SrcBpsDataPerInterval *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval `json:"SrcBpsDataPerInterval,omitempty" xml:"SrcBpsDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetSrcBpsDataPerInterval(v *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) *DescribeDomainSrcBpsDataResponseBody {
	s.SrcBpsDataPerInterval = v
	return s
}

type DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval struct {
	DataModule []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) SetDataModule(v []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

type DescribeDomainSrcBpsDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponse) SetBody(v *DescribeDomainSrcBpsDataResponseBody) *DescribeDomainSrcBpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcHttpCodeDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainSrcHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetOwnerId(v int64) *DescribeDomainSrcHttpCodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetDomainName(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetStartTime(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetEndTime(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataRequest) SetInterval(v string) *DescribeDomainSrcHttpCodeDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainSrcHttpCodeDataResponseBody struct {
	EndTime      *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval *string                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	HttpCodeData *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData `json:"HttpCodeData,omitempty" xml:"HttpCodeData,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetHttpCodeData(v *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.HttpCodeData = v
	return s
}

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData struct {
	UsageData []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) SetUsageData(v []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData struct {
	TimeStamp *string                                                              `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) SetValue(v *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue struct {
	CodeProportionData []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData `json:"CodeProportionData,omitempty" xml:"CodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) SetCodeProportionData(v []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	s.CodeProportionData = v
	return s
}

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCode(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetProportion(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCount(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Count = &v
	return s
}

type DescribeDomainSrcHttpCodeDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSrcHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponse) SetBody(v *DescribeDomainSrcHttpCodeDataResponseBody) *DescribeDomainSrcHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcQpsDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainSrcQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataRequest) SetOwnerId(v int64) *DescribeDomainSrcQpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetDomainName(v string) *DescribeDomainSrcQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetStartTime(v string) *DescribeDomainSrcQpsDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetEndTime(v string) *DescribeDomainSrcQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataRequest) SetInterval(v string) *DescribeDomainSrcQpsDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainSrcQpsDataResponseBody struct {
	EndTime               *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime             *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName            *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval          *string                                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	SrcQpsDataPerInterval *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval `json:"SrcQpsDataPerInterval,omitempty" xml:"SrcQpsDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetEndTime(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetStartTime(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetRequestId(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetDomainName(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBody) SetSrcQpsDataPerInterval(v *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) *DescribeDomainSrcQpsDataResponseBody {
	s.SrcQpsDataPerInterval = v
	return s
}

type DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval struct {
	DataModule []*DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval) SetDataModule(v []*DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcQpsDataResponseBodySrcQpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeDomainSrcQpsDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSrcQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcQpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcQpsDataResponse) SetBody(v *DescribeDomainSrcQpsDataResponseBody) *DescribeDomainSrcQpsDataResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcTopUrlVisitRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetOwnerId(v int64) *DescribeDomainSrcTopUrlVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetDomainName(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetStartTime(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetEndTime(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitRequest) SetSortBy(v string) *DescribeDomainSrcTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBody struct {
	StartTime  *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	AllUrlList *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
	Url200List *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	Url300List *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	Url400List *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	Url500List *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetStartTime(v string) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetRequestId(v string) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetDomainName(v string) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetAllUrlList(v *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl200List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl300List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl400List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBody) SetUrl500List(v *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) *DescribeDomainSrcTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

type DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeDomainSrcTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainSrcTopUrlVisitResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSrcTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcTopUrlVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcTopUrlVisitResponse) SetBody(v *DescribeDomainSrcTopUrlVisitResponseBody) *DescribeDomainSrcTopUrlVisitResponse {
	s.Body = v
	return s
}

type DescribeDomainSrcTrafficDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainSrcTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataRequest) SetOwnerId(v int64) *DescribeDomainSrcTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetDomainName(v string) *DescribeDomainSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetStartTime(v string) *DescribeDomainSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetEndTime(v string) *DescribeDomainSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataRequest) SetInterval(v string) *DescribeDomainSrcTrafficDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainSrcTrafficDataResponseBody struct {
	EndTime                   *string                                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime                 *string                                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId                 *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName                *string                                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval              *string                                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	SrcTrafficDataPerInterval *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval `json:"SrcTrafficDataPerInterval,omitempty" xml:"SrcTrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetSrcTrafficDataPerInterval(v *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) *DescribeDomainSrcTrafficDataResponseBody {
	s.SrcTrafficDataPerInterval = v
	return s
}

type DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval struct {
	DataModule []*DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) SetDataModule(v []*DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

type DescribeDomainSrcTrafficDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSrcTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponse) SetBody(v *DescribeDomainSrcTrafficDataResponseBody) *DescribeDomainSrcTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDomainsUsageByDayRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainsUsageByDayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayRequest) SetOwnerId(v int64) *DescribeDomainsUsageByDayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetDomainName(v string) *DescribeDomainsUsageByDayRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetStartTime(v string) *DescribeDomainsUsageByDayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayRequest) SetEndTime(v string) *DescribeDomainsUsageByDayRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainsUsageByDayResponseBody struct {
	EndTime      *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string                                           `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval *string                                           `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	UsageTotal   *DescribeDomainsUsageByDayResponseBodyUsageTotal  `json:"UsageTotal,omitempty" xml:"UsageTotal,omitempty" type:"Struct"`
	UsageByDays  *DescribeDomainsUsageByDayResponseBodyUsageByDays `json:"UsageByDays,omitempty" xml:"UsageByDays,omitempty" type:"Struct"`
}

func (s DescribeDomainsUsageByDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBody) SetEndTime(v string) *DescribeDomainsUsageByDayResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetStartTime(v string) *DescribeDomainsUsageByDayResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetRequestId(v string) *DescribeDomainsUsageByDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetDomainName(v string) *DescribeDomainsUsageByDayResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetDataInterval(v string) *DescribeDomainsUsageByDayResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetUsageTotal(v *DescribeDomainsUsageByDayResponseBodyUsageTotal) *DescribeDomainsUsageByDayResponseBody {
	s.UsageTotal = v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetUsageByDays(v *DescribeDomainsUsageByDayResponseBodyUsageByDays) *DescribeDomainsUsageByDayResponseBody {
	s.UsageByDays = v
	return s
}

type DescribeDomainsUsageByDayResponseBodyUsageTotal struct {
	MaxSrcBpsTime  *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	MaxBps         *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	TotalAccess    *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	BytesHitRate   *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	MaxSrcBps      *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageTotal) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageTotal) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetRequestHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetTotalAccess(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetBytesHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetTotalTraffic(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBps = &v
	return s
}

type DescribeDomainsUsageByDayResponseBodyUsageByDays struct {
	UsageByDay []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay `json:"UsageByDay,omitempty" xml:"UsageByDay,omitempty" type:"Repeated"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDays) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDays) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDays) SetUsageByDay(v []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) *DescribeDomainsUsageByDayResponseBodyUsageByDays {
	s.UsageByDay = v
	return s
}

type DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay struct {
	MaxSrcBpsTime  *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	Qps            *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	MaxBps         *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	TotalAccess    *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TimeStamp      *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	BytesHitRate   *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	MaxSrcBps      *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetQps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.Qps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetRequestHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalAccess(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTimeStamp(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetBytesHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalTraffic(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBpsTime = &v
	return s
}

type DescribeDomainsUsageByDayResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainsUsageByDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsUsageByDayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponse) SetHeaders(v map[string]*string) *DescribeDomainsUsageByDayResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsUsageByDayResponse) SetBody(v *DescribeDomainsUsageByDayResponseBody) *DescribeDomainsUsageByDayResponse {
	s.Body = v
	return s
}

type DescribeDomainTopClientIpVisitRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SortBy         *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Limit          *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s DescribeDomainTopClientIpVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitRequest) SetOwnerId(v int64) *DescribeDomainTopClientIpVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetDomainName(v string) *DescribeDomainTopClientIpVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetLocationNameEn(v string) *DescribeDomainTopClientIpVisitRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetStartTime(v string) *DescribeDomainTopClientIpVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetEndTime(v string) *DescribeDomainTopClientIpVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetSortBy(v string) *DescribeDomainTopClientIpVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitRequest) SetLimit(v string) *DescribeDomainTopClientIpVisitRequest {
	s.Limit = &v
	return s
}

type DescribeDomainTopClientIpVisitResponseBody struct {
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClientIpList []*DescribeDomainTopClientIpVisitResponseBodyClientIpList `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopClientIpVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitResponseBody) SetRequestId(v string) *DescribeDomainTopClientIpVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBody) SetClientIpList(v []*DescribeDomainTopClientIpVisitResponseBodyClientIpList) *DescribeDomainTopClientIpVisitResponseBody {
	s.ClientIpList = v
	return s
}

type DescribeDomainTopClientIpVisitResponseBodyClientIpList struct {
	Rank     *int32  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Acc      *int64  `json:"Acc,omitempty" xml:"Acc,omitempty"`
	Traffic  *int64  `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeDomainTopClientIpVisitResponseBodyClientIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitResponseBodyClientIpList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetRank(v int32) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.Rank = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetClientIp(v string) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.ClientIp = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetAcc(v int64) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.Acc = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetTraffic(v int64) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.Traffic = &v
	return s
}

type DescribeDomainTopClientIpVisitResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainTopClientIpVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainTopClientIpVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainTopClientIpVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponse) SetBody(v *DescribeDomainTopClientIpVisitResponseBody) *DescribeDomainTopClientIpVisitResponse {
	s.Body = v
	return s
}

type DescribeDomainTopReferVisitRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Percent    *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
}

func (s DescribeDomainTopReferVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopReferVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitRequest) SetOwnerId(v int64) *DescribeDomainTopReferVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetDomainName(v string) *DescribeDomainTopReferVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetStartTime(v string) *DescribeDomainTopReferVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetEndTime(v string) *DescribeDomainTopReferVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetSortBy(v string) *DescribeDomainTopReferVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDomainTopReferVisitRequest) SetPercent(v string) *DescribeDomainTopReferVisitRequest {
	s.Percent = &v
	return s
}

type DescribeDomainTopReferVisitResponseBody struct {
	StartTime    *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName   *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	TopReferList *DescribeDomainTopReferVisitResponseBodyTopReferList `json:"TopReferList,omitempty" xml:"TopReferList,omitempty" type:"Struct"`
}

func (s DescribeDomainTopReferVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponseBody) SetStartTime(v string) *DescribeDomainTopReferVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) SetRequestId(v string) *DescribeDomainTopReferVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) SetDomainName(v string) *DescribeDomainTopReferVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBody) SetTopReferList(v *DescribeDomainTopReferVisitResponseBodyTopReferList) *DescribeDomainTopReferVisitResponseBody {
	s.TopReferList = v
	return s
}

type DescribeDomainTopReferVisitResponseBodyTopReferList struct {
	ReferList []*DescribeDomainTopReferVisitResponseBodyTopReferListReferList `json:"ReferList,omitempty" xml:"ReferList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferList) SetReferList(v []*DescribeDomainTopReferVisitResponseBodyTopReferListReferList) *DescribeDomainTopReferVisitResponseBodyTopReferList {
	s.ReferList = v
	return s
}

type DescribeDomainTopReferVisitResponseBodyTopReferListReferList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	ReferDetail     *string  `json:"ReferDetail,omitempty" xml:"ReferDetail,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferListReferList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponseBodyTopReferListReferList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetFlow(v string) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetFlowProportion(v float32) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitData(v string) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetReferDetail(v string) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.ReferDetail = &v
	return s
}

func (s *DescribeDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitProportion(v float32) *DescribeDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainTopReferVisitResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainTopReferVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainTopReferVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopReferVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopReferVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainTopReferVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopReferVisitResponse) SetBody(v *DescribeDomainTopReferVisitResponseBody) *DescribeDomainTopReferVisitResponse {
	s.Body = v
	return s
}

type DescribeDomainTopUrlVisitRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s DescribeDomainTopUrlVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitRequest) SetOwnerId(v int64) *DescribeDomainTopUrlVisitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetDomainName(v string) *DescribeDomainTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetStartTime(v string) *DescribeDomainTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetEndTime(v string) *DescribeDomainTopUrlVisitRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopUrlVisitRequest) SetSortBy(v string) *DescribeDomainTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

type DescribeDomainTopUrlVisitResponseBody struct {
	StartTime  *string                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName *string                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	AllUrlList *DescribeDomainTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
	Url200List *DescribeDomainTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	Url300List *DescribeDomainTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	Url400List *DescribeDomainTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	Url500List *DescribeDomainTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeDomainTopUrlVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetStartTime(v string) *DescribeDomainTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetRequestId(v string) *DescribeDomainTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetDomainName(v string) *DescribeDomainTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetAllUrlList(v *DescribeDomainTopUrlVisitResponseBodyAllUrlList) *DescribeDomainTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl200List(v *DescribeDomainTopUrlVisitResponseBodyUrl200List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl300List(v *DescribeDomainTopUrlVisitResponseBodyUrl300List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl400List(v *DescribeDomainTopUrlVisitResponseBodyUrl400List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBody) SetUrl500List(v *DescribeDomainTopUrlVisitResponseBodyUrl500List) *DescribeDomainTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeDomainTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500List) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeDomainTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

type DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeDomainTopUrlVisitResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainTopUrlVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeDomainTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTopUrlVisitResponse) SetBody(v *DescribeDomainTopUrlVisitResponseBody) *DescribeDomainTopUrlVisitResponse {
	s.Body = v
	return s
}

type DescribeDomainTrafficDataRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
}

func (s DescribeDomainTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetDomainName(v string) *DescribeDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetStartTime(v string) *DescribeDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetEndTime(v string) *DescribeDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetInterval(v string) *DescribeDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

type DescribeDomainTrafficDataResponseBody struct {
	EndTime                *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime              *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName             *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval           *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	TrafficDataPerInterval *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	OverseasValue      *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	HttpsValue         *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	DomesticValue      *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

type DescribeDomainTrafficDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainTrafficDataResponse) SetBody(v *DescribeDomainTrafficDataResponseBody) *DescribeDomainTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeDomainUsageDataRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DataProtocol *string `json:"DataProtocol,omitempty" xml:"DataProtocol,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Area         *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Field        *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDomainUsageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataRequest) SetOwnerId(v int64) *DescribeDomainUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetDataProtocol(v string) *DescribeDomainUsageDataRequest {
	s.DataProtocol = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetDomainName(v string) *DescribeDomainUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetStartTime(v string) *DescribeDomainUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetEndTime(v string) *DescribeDomainUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetArea(v string) *DescribeDomainUsageDataRequest {
	s.Area = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetField(v string) *DescribeDomainUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainUsageDataRequest) SetInterval(v string) *DescribeDomainUsageDataRequest {
	s.Interval = &v
	return s
}

type DescribeDomainUsageDataResponseBody struct {
	EndTime              *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type                 *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	StartTime            *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Area                 *string                                                  `json:"Area,omitempty" xml:"Area,omitempty"`
	DomainName           *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval         *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	UsageDataPerInterval *DescribeDomainUsageDataResponseBodyUsageDataPerInterval `json:"UsageDataPerInterval,omitempty" xml:"UsageDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainUsageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBody) SetEndTime(v string) *DescribeDomainUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetType(v string) *DescribeDomainUsageDataResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetStartTime(v string) *DescribeDomainUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetRequestId(v string) *DescribeDomainUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetArea(v string) *DescribeDomainUsageDataResponseBody {
	s.Area = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetDomainName(v string) *DescribeDomainUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetDataInterval(v string) *DescribeDomainUsageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetUsageDataPerInterval(v *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeDomainUsageDataResponseBody {
	s.UsageDataPerInterval = v
	return s
}

type DescribeDomainUsageDataResponseBodyUsageDataPerInterval struct {
	DataModule []*DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) SetDataModule(v []*DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) *DescribeDomainUsageDataResponseBodyUsageDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule struct {
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	PeakTime     *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	SpecialValue *string `json:"SpecialValue,omitempty" xml:"SpecialValue,omitempty"`
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetValue(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetPeakTime(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.PeakTime = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetSpecialValue(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.SpecialValue = &v
	return s
}

type DescribeDomainUsageDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainUsageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDomainUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainUsageDataResponse) SetBody(v *DescribeDomainUsageDataResponseBody) *DescribeDomainUsageDataResponse {
	s.Body = v
	return s
}

type DescribeDomainUvDataRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainUvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataRequest) SetOwnerId(v int64) *DescribeDomainUvDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetDomainName(v string) *DescribeDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetStartTime(v string) *DescribeDomainUvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUvDataRequest) SetEndTime(v string) *DescribeDomainUvDataRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainUvDataResponseBody struct {
	EndTime        *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainName     *string                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DataInterval   *string                                         `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	UvDataInterval *DescribeDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainUvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBody) SetEndTime(v string) *DescribeDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetStartTime(v string) *DescribeDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetRequestId(v string) *DescribeDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetDomainName(v string) *DescribeDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetDataInterval(v string) *DescribeDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetUvDataInterval(v *DescribeDomainUvDataResponseBodyUvDataInterval) *DescribeDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

type DescribeDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainUvDataResponseBodyUvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

type DescribeDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

type DescribeDomainUvDataResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainUvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainUvDataResponse) SetBody(v *DescribeDomainUvDataResponseBody) *DescribeDomainUvDataResponse {
	s.Body = v
	return s
}

type DescribeEsExceptionDataRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeEsExceptionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExceptionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataRequest) SetOwnerId(v int64) *DescribeEsExceptionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEsExceptionDataRequest) SetStartTime(v string) *DescribeEsExceptionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEsExceptionDataRequest) SetEndTime(v string) *DescribeEsExceptionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEsExceptionDataRequest) SetRuleId(v string) *DescribeEsExceptionDataRequest {
	s.RuleId = &v
	return s
}

type DescribeEsExceptionDataResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Contents  []*DescribeEsExceptionDataResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s DescribeEsExceptionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExceptionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataResponseBody) SetRequestId(v string) *DescribeEsExceptionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEsExceptionDataResponseBody) SetContents(v []*DescribeEsExceptionDataResponseBodyContents) *DescribeEsExceptionDataResponseBody {
	s.Contents = v
	return s
}

type DescribeEsExceptionDataResponseBodyContents struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Points  []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s DescribeEsExceptionDataResponseBodyContents) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExceptionDataResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataResponseBodyContents) SetName(v string) *DescribeEsExceptionDataResponseBodyContents {
	s.Name = &v
	return s
}

func (s *DescribeEsExceptionDataResponseBodyContents) SetPoints(v []*string) *DescribeEsExceptionDataResponseBodyContents {
	s.Points = v
	return s
}

func (s *DescribeEsExceptionDataResponseBodyContents) SetColumns(v []*string) *DescribeEsExceptionDataResponseBodyContents {
	s.Columns = v
	return s
}

type DescribeEsExceptionDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEsExceptionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEsExceptionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExceptionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataResponse) SetHeaders(v map[string]*string) *DescribeEsExceptionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEsExceptionDataResponse) SetBody(v *DescribeEsExceptionDataResponseBody) *DescribeEsExceptionDataResponse {
	s.Body = v
	return s
}

type DescribeEsExecuteDataRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeEsExecuteDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExecuteDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataRequest) SetOwnerId(v int64) *DescribeEsExecuteDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEsExecuteDataRequest) SetStartTime(v string) *DescribeEsExecuteDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEsExecuteDataRequest) SetEndTime(v string) *DescribeEsExecuteDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEsExecuteDataRequest) SetRuleId(v string) *DescribeEsExecuteDataRequest {
	s.RuleId = &v
	return s
}

type DescribeEsExecuteDataResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Contents  []*DescribeEsExecuteDataResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s DescribeEsExecuteDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExecuteDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataResponseBody) SetRequestId(v string) *DescribeEsExecuteDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEsExecuteDataResponseBody) SetContents(v []*DescribeEsExecuteDataResponseBodyContents) *DescribeEsExecuteDataResponseBody {
	s.Contents = v
	return s
}

type DescribeEsExecuteDataResponseBodyContents struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Points  []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s DescribeEsExecuteDataResponseBodyContents) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExecuteDataResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataResponseBodyContents) SetName(v string) *DescribeEsExecuteDataResponseBodyContents {
	s.Name = &v
	return s
}

func (s *DescribeEsExecuteDataResponseBodyContents) SetPoints(v []*string) *DescribeEsExecuteDataResponseBodyContents {
	s.Points = v
	return s
}

func (s *DescribeEsExecuteDataResponseBodyContents) SetColumns(v []*string) *DescribeEsExecuteDataResponseBodyContents {
	s.Columns = v
	return s
}

type DescribeEsExecuteDataResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEsExecuteDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEsExecuteDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEsExecuteDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataResponse) SetHeaders(v map[string]*string) *DescribeEsExecuteDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEsExecuteDataResponse) SetBody(v *DescribeEsExecuteDataResponseBody) *DescribeEsExecuteDataResponse {
	s.Body = v
	return s
}

type DescribeFCTriggerRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TriggerARN *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
}

func (s DescribeFCTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerRequest) SetOwnerId(v int64) *DescribeFCTriggerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFCTriggerRequest) SetTriggerARN(v string) *DescribeFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

type DescribeFCTriggerResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FCTrigger *DescribeFCTriggerResponseBodyFCTrigger `json:"FCTrigger,omitempty" xml:"FCTrigger,omitempty" type:"Struct"`
}

func (s DescribeFCTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerResponseBody) SetRequestId(v string) *DescribeFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFCTriggerResponseBody) SetFCTrigger(v *DescribeFCTriggerResponseBodyFCTrigger) *DescribeFCTriggerResponseBody {
	s.FCTrigger = v
	return s
}

type DescribeFCTriggerResponseBodyFCTrigger struct {
	TriggerARN       *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
	RoleARN          *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	SourceArn        *string `json:"SourceArn,omitempty" xml:"SourceArn,omitempty"`
	Notes            *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	EventMetaName    *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
}

func (s DescribeFCTriggerResponseBodyFCTrigger) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCTriggerResponseBodyFCTrigger) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetTriggerARN(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.TriggerARN = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetRoleARN(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.RoleARN = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetSourceArn(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.SourceArn = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetNotes(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.Notes = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetEventMetaName(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.EventMetaName = &v
	return s
}

func (s *DescribeFCTriggerResponseBodyFCTrigger) SetEventMetaVersion(v string) *DescribeFCTriggerResponseBodyFCTrigger {
	s.EventMetaVersion = &v
	return s
}

type DescribeFCTriggerResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFCTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerResponse) SetHeaders(v map[string]*string) *DescribeFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *DescribeFCTriggerResponse) SetBody(v *DescribeFCTriggerResponseBody) *DescribeFCTriggerResponse {
	s.Body = v
	return s
}

type DescribeIllegalUrlExportTaskRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeIllegalUrlExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIllegalUrlExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeIllegalUrlExportTaskRequest) SetOwnerId(v int64) *DescribeIllegalUrlExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIllegalUrlExportTaskRequest) SetTaskId(v string) *DescribeIllegalUrlExportTaskRequest {
	s.TaskId = &v
	return s
}

type DescribeIllegalUrlExportTaskResponseBody struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s DescribeIllegalUrlExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIllegalUrlExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIllegalUrlExportTaskResponseBody) SetStatus(v string) *DescribeIllegalUrlExportTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeIllegalUrlExportTaskResponseBody) SetRequestId(v string) *DescribeIllegalUrlExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIllegalUrlExportTaskResponseBody) SetDownloadUrl(v string) *DescribeIllegalUrlExportTaskResponseBody {
	s.DownloadUrl = &v
	return s
}

type DescribeIllegalUrlExportTaskResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIllegalUrlExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIllegalUrlExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIllegalUrlExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeIllegalUrlExportTaskResponse) SetHeaders(v map[string]*string) *DescribeIllegalUrlExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeIllegalUrlExportTaskResponse) SetBody(v *DescribeIllegalUrlExportTaskResponseBody) *DescribeIllegalUrlExportTaskResponse {
	s.Body = v
	return s
}

type DescribeIpInfoRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	IP            *string `json:"IP,omitempty" xml:"IP,omitempty"`
}

func (s DescribeIpInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoRequest) SetOwnerId(v int64) *DescribeIpInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpInfoRequest) SetSecurityToken(v string) *DescribeIpInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeIpInfoRequest) SetIP(v string) *DescribeIpInfoRequest {
	s.IP = &v
	return s
}

type DescribeIpInfoResponseBody struct {
	CdnIp       *string `json:"CdnIp,omitempty" xml:"CdnIp,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IspEname    *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	ISP         *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
}

func (s DescribeIpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoResponseBody) SetCdnIp(v string) *DescribeIpInfoResponseBody {
	s.CdnIp = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRequestId(v string) *DescribeIpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRegionEname(v string) *DescribeIpInfoResponseBody {
	s.RegionEname = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRegion(v string) *DescribeIpInfoResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIspEname(v string) *DescribeIpInfoResponseBody {
	s.IspEname = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetISP(v string) *DescribeIpInfoResponseBody {
	s.ISP = &v
	return s
}

type DescribeIpInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoResponse) SetHeaders(v map[string]*string) *DescribeIpInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpInfoResponse) SetBody(v *DescribeIpInfoResponseBody) *DescribeIpInfoResponse {
	s.Body = v
	return s
}

type DescribeL2VipsByDomainRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeL2VipsByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeL2VipsByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainRequest) SetOwnerId(v int64) *DescribeL2VipsByDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeL2VipsByDomainRequest) SetSecurityToken(v string) *DescribeL2VipsByDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeL2VipsByDomainRequest) SetDomainName(v string) *DescribeL2VipsByDomainRequest {
	s.DomainName = &v
	return s
}

type DescribeL2VipsByDomainResponseBody struct {
	DomainName *string                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Vips       *DescribeL2VipsByDomainResponseBodyVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
}

func (s DescribeL2VipsByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeL2VipsByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainResponseBody) SetDomainName(v string) *DescribeL2VipsByDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeL2VipsByDomainResponseBody) SetRequestId(v string) *DescribeL2VipsByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL2VipsByDomainResponseBody) SetVips(v *DescribeL2VipsByDomainResponseBodyVips) *DescribeL2VipsByDomainResponseBody {
	s.Vips = v
	return s
}

type DescribeL2VipsByDomainResponseBodyVips struct {
	Vip []*string `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeL2VipsByDomainResponseBodyVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeL2VipsByDomainResponseBodyVips) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainResponseBodyVips) SetVip(v []*string) *DescribeL2VipsByDomainResponseBodyVips {
	s.Vip = v
	return s
}

type DescribeL2VipsByDomainResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeL2VipsByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeL2VipsByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeL2VipsByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainResponse) SetHeaders(v map[string]*string) *DescribeL2VipsByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeL2VipsByDomainResponse) SetBody(v *DescribeL2VipsByDomainResponseBody) *DescribeL2VipsByDomainResponse {
	s.Body = v
	return s
}

type DescribeRangeDataByLocateAndIspServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNames      *string `json:"IspNames,omitempty" xml:"IspNames,omitempty"`
	LocationNames *string `json:"LocationNames,omitempty" xml:"LocationNames,omitempty"`
}

func (s DescribeRangeDataByLocateAndIspServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRangeDataByLocateAndIspServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetOwnerId(v int64) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetDomainNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.DomainNames = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetStartTime(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetEndTime(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetIspNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.IspNames = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceRequest) SetLocationNames(v string) *DescribeRangeDataByLocateAndIspServiceRequest {
	s.LocationNames = &v
	return s
}

type DescribeRangeDataByLocateAndIspServiceResponseBody struct {
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRangeDataByLocateAndIspServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRangeDataByLocateAndIspServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) SetJsonResult(v string) *DescribeRangeDataByLocateAndIspServiceResponseBody {
	s.JsonResult = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) SetRequestId(v string) *DescribeRangeDataByLocateAndIspServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRangeDataByLocateAndIspServiceResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRangeDataByLocateAndIspServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRangeDataByLocateAndIspServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRangeDataByLocateAndIspServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) SetHeaders(v map[string]*string) *DescribeRangeDataByLocateAndIspServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponse) SetBody(v *DescribeRangeDataByLocateAndIspServiceResponseBody) *DescribeRangeDataByLocateAndIspServiceResponse {
	s.Body = v
	return s
}

type DescribeRealtimeDeliveryAccRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	LogStore  *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s DescribeRealtimeDeliveryAccRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccRequest) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccRequest) SetOwnerId(v int64) *DescribeRealtimeDeliveryAccRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetStartTime(v string) *DescribeRealtimeDeliveryAccRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetEndTime(v string) *DescribeRealtimeDeliveryAccRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetInterval(v string) *DescribeRealtimeDeliveryAccRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetProject(v string) *DescribeRealtimeDeliveryAccRequest {
	s.Project = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccRequest) SetLogStore(v string) *DescribeRealtimeDeliveryAccRequest {
	s.LogStore = &v
	return s
}

type DescribeRealtimeDeliveryAccResponseBody struct {
	RequestId               *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReatTimeDeliveryAccData *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData `json:"ReatTimeDeliveryAccData,omitempty" xml:"ReatTimeDeliveryAccData,omitempty" type:"Struct"`
}

func (s DescribeRealtimeDeliveryAccResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponseBody) SetRequestId(v string) *DescribeRealtimeDeliveryAccResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBody) SetReatTimeDeliveryAccData(v *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) *DescribeRealtimeDeliveryAccResponseBody {
	s.ReatTimeDeliveryAccData = v
	return s
}

type DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData struct {
	AccData []*DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData `json:"AccData,omitempty" xml:"AccData,omitempty" type:"Repeated"`
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData) SetAccData(v []*DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccData {
	s.AccData = v
	return s
}

type DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData struct {
	TimeStamp  *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	FailedNum  *int32  `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	SuccessNum *int32  `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) SetTimeStamp(v string) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) SetFailedNum(v int32) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	s.FailedNum = &v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData) SetSuccessNum(v int32) *DescribeRealtimeDeliveryAccResponseBodyReatTimeDeliveryAccDataAccData {
	s.SuccessNum = &v
	return s
}

type DescribeRealtimeDeliveryAccResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRealtimeDeliveryAccResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRealtimeDeliveryAccResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRealtimeDeliveryAccResponse) GoString() string {
	return s.String()
}

func (s *DescribeRealtimeDeliveryAccResponse) SetHeaders(v map[string]*string) *DescribeRealtimeDeliveryAccResponse {
	s.Headers = v
	return s
}

func (s *DescribeRealtimeDeliveryAccResponse) SetBody(v *DescribeRealtimeDeliveryAccResponseBody) *DescribeRealtimeDeliveryAccResponse {
	s.Body = v
	return s
}

type DescribeRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRefreshQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaRequest) SetOwnerId(v int64) *DescribeRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshQuotaRequest) SetSecurityToken(v string) *DescribeRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRefreshQuotaResponseBody struct {
	PreloadRemain     *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	BlockRemain       *string `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	RegexRemain       *string `json:"RegexRemain,omitempty" xml:"RegexRemain,omitempty"`
	UrlRemain         *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
	DirRemain         *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	UrlQuota          *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	BlockQuota        *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DirQuota          *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	PreloadEdgeQuota  *string `json:"PreloadEdgeQuota,omitempty" xml:"PreloadEdgeQuota,omitempty"`
	PreloadEdgeRemain *string `json:"PreloadEdgeRemain,omitempty" xml:"PreloadEdgeRemain,omitempty"`
	PreloadQuota      *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	RegexQuota        *string `json:"RegexQuota,omitempty" xml:"RegexQuota,omitempty"`
}

func (s DescribeRefreshQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRegexRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.RegexRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRequestId(v string) *DescribeRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadEdgeQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadEdgeQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadEdgeRemain(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadEdgeRemain = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeRefreshQuotaResponseBody) SetRegexQuota(v string) *DescribeRefreshQuotaResponseBody {
	s.RegexQuota = &v
	return s
}

type DescribeRefreshQuotaResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRefreshQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshQuotaResponse) SetBody(v *DescribeRefreshQuotaResponseBody) *DescribeRefreshQuotaResponse {
	s.Body = v
	return s
}

type DescribeRefreshTaskByIdRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTaskByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTaskByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdRequest) SetOwnerId(v int64) *DescribeRefreshTaskByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshTaskByIdRequest) SetTaskId(v string) *DescribeRefreshTaskByIdRequest {
	s.TaskId = &v
	return s
}

type DescribeRefreshTaskByIdResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      []*DescribeRefreshTaskByIdResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DescribeRefreshTaskByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTaskByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdResponseBody) SetTotalCount(v int64) *DescribeRefreshTaskByIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBody) SetRequestId(v string) *DescribeRefreshTaskByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBody) SetTasks(v []*DescribeRefreshTaskByIdResponseBodyTasks) *DescribeRefreshTaskByIdResponseBody {
	s.Tasks = v
	return s
}

type DescribeRefreshTaskByIdResponseBodyTasks struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTaskByIdResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTaskByIdResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetStatus(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetCreationTime(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.CreationTime = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetObjectType(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetProcess(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.Process = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetDescription(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetObjectPath(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTaskByIdResponseBodyTasks) SetTaskId(v string) *DescribeRefreshTaskByIdResponseBodyTasks {
	s.TaskId = &v
	return s
}

type DescribeRefreshTaskByIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRefreshTaskByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRefreshTaskByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTaskByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTaskByIdResponse) SetHeaders(v map[string]*string) *DescribeRefreshTaskByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshTaskByIdResponse) SetBody(v *DescribeRefreshTaskByIdResponseBody) *DescribeRefreshTaskByIdResponse {
	s.Body = v
	return s
}

type DescribeRefreshTasksRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ObjectPath      *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ObjectType      *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeRefreshTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksRequest) SetOwnerId(v int64) *DescribeRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetSecurityToken(v string) *DescribeRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetTaskId(v string) *DescribeRefreshTasksRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetObjectPath(v string) *DescribeRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetPageNumber(v int32) *DescribeRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetObjectType(v string) *DescribeRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetDomainName(v string) *DescribeRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetStatus(v string) *DescribeRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetPageSize(v int32) *DescribeRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetStartTime(v string) *DescribeRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetEndTime(v string) *DescribeRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRefreshTasksRequest) SetResourceGroupId(v string) *DescribeRefreshTasksRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeRefreshTasksResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Tasks      *DescribeRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeRefreshTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBody) SetRequestId(v string) *DescribeRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetPageSize(v int64) *DescribeRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRefreshTasksResponseBody) SetTasks(v *DescribeRefreshTasksResponseBodyTasks) *DescribeRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

type DescribeRefreshTasksResponseBodyTasks struct {
	CDNTask []*DescribeRefreshTasksResponseBodyTasksCDNTask `json:"CDNTask,omitempty" xml:"CDNTask,omitempty" type:"Repeated"`
}

func (s DescribeRefreshTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBodyTasks) SetCDNTask(v []*DescribeRefreshTasksResponseBodyTasksCDNTask) *DescribeRefreshTasksResponseBodyTasks {
	s.CDNTask = v
	return s
}

type DescribeRefreshTasksResponseBodyTasksCDNTask struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRefreshTasksResponseBodyTasksCDNTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTasksResponseBodyTasksCDNTask) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetStatus(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.Status = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetCreationTime(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetObjectType(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetProcess(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.Process = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetDescription(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.Description = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetObjectPath(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeRefreshTasksResponseBodyTasksCDNTask) SetTaskId(v string) *DescribeRefreshTasksResponseBodyTasksCDNTask {
	s.TaskId = &v
	return s
}

type DescribeRefreshTasksResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRefreshTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshTasksResponse) SetBody(v *DescribeRefreshTasksResponseBody) *DescribeRefreshTasksResponse {
	s.Body = v
	return s
}

type DescribeStagingIpRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeStagingIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStagingIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpRequest) SetOwnerId(v int64) *DescribeStagingIpRequest {
	s.OwnerId = &v
	return s
}

type DescribeStagingIpResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IPV4s     *DescribeStagingIpResponseBodyIPV4s `json:"IPV4s,omitempty" xml:"IPV4s,omitempty" type:"Struct"`
}

func (s DescribeStagingIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStagingIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpResponseBody) SetRequestId(v string) *DescribeStagingIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStagingIpResponseBody) SetIPV4s(v *DescribeStagingIpResponseBodyIPV4s) *DescribeStagingIpResponseBody {
	s.IPV4s = v
	return s
}

type DescribeStagingIpResponseBodyIPV4s struct {
	IPV4 []*string `json:"IPV4,omitempty" xml:"IPV4,omitempty" type:"Repeated"`
}

func (s DescribeStagingIpResponseBodyIPV4s) String() string {
	return tea.Prettify(s)
}

func (s DescribeStagingIpResponseBodyIPV4s) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpResponseBodyIPV4s) SetIPV4(v []*string) *DescribeStagingIpResponseBodyIPV4s {
	s.IPV4 = v
	return s
}

type DescribeStagingIpResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStagingIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStagingIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStagingIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeStagingIpResponse) SetHeaders(v map[string]*string) *DescribeStagingIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeStagingIpResponse) SetBody(v *DescribeStagingIpResponseBody) *DescribeStagingIpResponse {
	s.Body = v
	return s
}

type DescribeTagResourcesRequest struct {
	OwnerId      *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceType *string                           `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                         `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*DescribeTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequest) SetOwnerId(v int64) *DescribeTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceType(v string) *DescribeTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceId(v []*string) *DescribeTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeTagResourcesRequest) SetTag(v []*DescribeTagResourcesRequestTag) *DescribeTagResourcesRequest {
	s.Tag = v
	return s
}

type DescribeTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequestTag) SetKey(v string) *DescribeTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTagResourcesRequestTag) SetValue(v string) *DescribeTagResourcesRequestTag {
	s.Value = &v
	return s
}

type DescribeTagResourcesResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*DescribeTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBody) SetRequestId(v string) *DescribeTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagResourcesResponseBody) SetTagResources(v []*DescribeTagResourcesResponseBodyTagResources) *DescribeTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type DescribeTagResourcesResponseBodyTagResources struct {
	ResourceId *string                                            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Tag        []*DescribeTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResources) SetTag(v []*DescribeTagResourcesResponseBodyTagResourcesTag) *DescribeTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

type DescribeTagResourcesResponseBodyTagResourcesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagResourcesResponseBodyTagResourcesTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
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

type DescribeTopDomainsByFlowRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s DescribeTopDomainsByFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowRequest) SetOwnerId(v int64) *DescribeTopDomainsByFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetStartTime(v string) *DescribeTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetEndTime(v string) *DescribeTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowRequest) SetLimit(v int64) *DescribeTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

type DescribeTopDomainsByFlowResponseBody struct {
	DomainOnlineCount *int64                                          `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	EndTime           *string                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime         *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainCount       *int64                                          `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	TopDomains        *DescribeTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeTopDomainsByFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeTopDomainsByFlowResponseBodyTopDomains) *DescribeTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

type DescribeTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

type DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	MaxBps         *float32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	Rank           *int64   `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64   `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TrafficPercent *string  `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
	DomainName     *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	TotalTraffic   *string  `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	MaxBpsTime     *string  `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v float32) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

type DescribeTopDomainsByFlowResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopDomainsByFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopDomainsByFlowResponse) SetBody(v *DescribeTopDomainsByFlowResponseBody) *DescribeTopDomainsByFlowResponse {
	s.Body = v
	return s
}

type DescribeUserCertificateExpireCountRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeUserCertificateExpireCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateExpireCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateExpireCountRequest) SetOwnerId(v int64) *DescribeUserCertificateExpireCountRequest {
	s.OwnerId = &v
	return s
}

type DescribeUserCertificateExpireCountResponseBody struct {
	ExpireWithin30DaysCount *int32  `json:"ExpireWithin30DaysCount,omitempty" xml:"ExpireWithin30DaysCount,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExpiredCount            *int32  `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
}

func (s DescribeUserCertificateExpireCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateExpireCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateExpireCountResponseBody) SetExpireWithin30DaysCount(v int32) *DescribeUserCertificateExpireCountResponseBody {
	s.ExpireWithin30DaysCount = &v
	return s
}

func (s *DescribeUserCertificateExpireCountResponseBody) SetRequestId(v string) *DescribeUserCertificateExpireCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCertificateExpireCountResponseBody) SetExpiredCount(v int32) *DescribeUserCertificateExpireCountResponseBody {
	s.ExpiredCount = &v
	return s
}

type DescribeUserCertificateExpireCountResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserCertificateExpireCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserCertificateExpireCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateExpireCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateExpireCountResponse) SetHeaders(v map[string]*string) *DescribeUserCertificateExpireCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserCertificateExpireCountResponse) SetBody(v *DescribeUserCertificateExpireCountResponseBody) *DescribeUserCertificateExpireCountResponse {
	s.Body = v
	return s
}

type DescribeUserConfigsRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Config        *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s DescribeUserConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsRequest) SetOwnerId(v int64) *DescribeUserConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserConfigsRequest) SetSecurityToken(v string) *DescribeUserConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserConfigsRequest) SetConfig(v string) *DescribeUserConfigsRequest {
	s.Config = &v
	return s
}

type DescribeUserConfigsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Configs   *DescribeUserConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
}

func (s DescribeUserConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBody) SetRequestId(v string) *DescribeUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserConfigsResponseBody) SetConfigs(v *DescribeUserConfigsResponseBodyConfigs) *DescribeUserConfigsResponseBody {
	s.Configs = v
	return s
}

type DescribeUserConfigsResponseBodyConfigs struct {
	OssLogConfig       *DescribeUserConfigsResponseBodyConfigsOssLogConfig       `json:"OssLogConfig,omitempty" xml:"OssLogConfig,omitempty" type:"Struct"`
	GreenManagerConfig *DescribeUserConfigsResponseBodyConfigsGreenManagerConfig `json:"GreenManagerConfig,omitempty" xml:"GreenManagerConfig,omitempty" type:"Struct"`
	WafConfig          *DescribeUserConfigsResponseBodyConfigsWafConfig          `json:"WafConfig,omitempty" xml:"WafConfig,omitempty" type:"Struct"`
}

func (s DescribeUserConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigs) SetOssLogConfig(v *DescribeUserConfigsResponseBodyConfigsOssLogConfig) *DescribeUserConfigsResponseBodyConfigs {
	s.OssLogConfig = v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigs) SetGreenManagerConfig(v *DescribeUserConfigsResponseBodyConfigsGreenManagerConfig) *DescribeUserConfigsResponseBodyConfigs {
	s.GreenManagerConfig = v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigs) SetWafConfig(v *DescribeUserConfigsResponseBodyConfigsWafConfig) *DescribeUserConfigsResponseBodyConfigs {
	s.WafConfig = v
	return s
}

type DescribeUserConfigsResponseBodyConfigsOssLogConfig struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s DescribeUserConfigsResponseBodyConfigsOssLogConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigsOssLogConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) SetPrefix(v string) *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	s.Prefix = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) SetEnable(v string) *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	s.Enable = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) SetBucket(v string) *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	s.Bucket = &v
	return s
}

type DescribeUserConfigsResponseBodyConfigsGreenManagerConfig struct {
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s DescribeUserConfigsResponseBodyConfigsGreenManagerConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigsGreenManagerConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigsGreenManagerConfig) SetRatio(v string) *DescribeUserConfigsResponseBodyConfigsGreenManagerConfig {
	s.Ratio = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsGreenManagerConfig) SetQuota(v string) *DescribeUserConfigsResponseBodyConfigsGreenManagerConfig {
	s.Quota = &v
	return s
}

type DescribeUserConfigsResponseBodyConfigsWafConfig struct {
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeUserConfigsResponseBodyConfigsWafConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigsWafConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigsWafConfig) SetEnable(v string) *DescribeUserConfigsResponseBodyConfigsWafConfig {
	s.Enable = &v
	return s
}

type DescribeUserConfigsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponse) SetHeaders(v map[string]*string) *DescribeUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserConfigsResponse) SetBody(v *DescribeUserConfigsResponseBody) *DescribeUserConfigsResponse {
	s.Body = v
	return s
}

type DescribeUserDomainsRequest struct {
	OwnerId          *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken    *string                          `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	PageSize         *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DomainName       *string                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus     *string                          `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	DomainSearchType *string                          `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	CdnType          *string                          `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	CheckDomainShow  *bool                            `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	ResourceGroupId  *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ChangeStartTime  *string                          `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	ChangeEndTime    *string                          `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	Coverage         *string                          `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	Tag              []*DescribeUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsRequest) SetOwnerId(v int64) *DescribeUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetSecurityToken(v string) *DescribeUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetPageSize(v int32) *DescribeUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetPageNumber(v int32) *DescribeUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainName(v string) *DescribeUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainStatus(v string) *DescribeUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetDomainSearchType(v string) *DescribeUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetCdnType(v string) *DescribeUserDomainsRequest {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetResourceGroupId(v string) *DescribeUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetChangeStartTime(v string) *DescribeUserDomainsRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetChangeEndTime(v string) *DescribeUserDomainsRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetCoverage(v string) *DescribeUserDomainsRequest {
	s.Coverage = &v
	return s
}

func (s *DescribeUserDomainsRequest) SetTag(v []*DescribeUserDomainsRequestTag) *DescribeUserDomainsRequest {
	s.Tag = v
	return s
}

type DescribeUserDomainsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUserDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsRequestTag) SetKey(v string) *DescribeUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeUserDomainsRequestTag) SetValue(v string) *DescribeUserDomainsRequestTag {
	s.Value = &v
	return s
}

type DescribeUserDomainsResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Domains    *DescribeUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
}

func (s DescribeUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBody) SetRequestId(v string) *DescribeUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetPageNumber(v int64) *DescribeUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetPageSize(v int64) *DescribeUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetTotalCount(v int64) *DescribeUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserDomainsResponseBody) SetDomains(v *DescribeUserDomainsResponseBodyDomains) *DescribeUserDomainsResponseBody {
	s.Domains = v
	return s
}

type DescribeUserDomainsResponseBodyDomains struct {
	PageData []*DescribeUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomains) SetPageData(v []*DescribeUserDomainsResponseBodyDomainsPageData) *DescribeUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeUserDomainsResponseBodyDomainsPageData struct {
	GmtCreated      *string                                                `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	SslProtocol     *string                                                `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
	Description     *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Coverage        *string                                                `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	ResourceGroupId *string                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Sandbox         *string                                                `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	DomainStatus    *string                                                `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	Cname           *string                                                `json:"Cname,omitempty" xml:"Cname,omitempty"`
	GmtModified     *string                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	CdnType         *string                                                `json:"CdnType,omitempty" xml:"CdnType,omitempty"`
	DomainName      *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Sources         *DescribeUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCoverage(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Coverage = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetCdnType(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.CdnType = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeUserDomainsResponseBodyDomainsPageDataSources) *DescribeUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

type DescribeUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight   *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetWeight(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

type DescribeUserDomainsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDomainsResponse) SetBody(v *DescribeUserDomainsResponseBody) *DescribeUserDomainsResponse {
	s.Body = v
	return s
}

type DescribeUserTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsRequest) SetOwnerId(v int64) *DescribeUserTagsRequest {
	s.OwnerId = &v
	return s
}

type DescribeUserTagsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DescribeUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponseBody) SetRequestId(v string) *DescribeUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTagsResponseBody) SetTags(v []*DescribeUserTagsResponseBodyTags) *DescribeUserTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeUserTagsResponseBodyTags struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeUserTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponseBodyTags) SetKey(v string) *DescribeUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeUserTagsResponseBodyTags) SetValue(v []*string) *DescribeUserTagsResponseBodyTags {
	s.Value = v
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

type DescribeUserUsageDataExportTaskRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeUserUsageDataExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskRequest) SetOwnerId(v int64) *DescribeUserUsageDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskRequest) SetPageSize(v string) *DescribeUserUsageDataExportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskRequest) SetPageNumber(v string) *DescribeUserUsageDataExportTaskRequest {
	s.PageNumber = &v
	return s
}

type DescribeUserUsageDataExportTaskResponseBody struct {
	RequestId        *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageDataPerPage *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage `json:"UsageDataPerPage,omitempty" xml:"UsageDataPerPage,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDataExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBody) SetRequestId(v string) *DescribeUserUsageDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBody) SetUsageDataPerPage(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) *DescribeUserUsageDataExportTaskResponseBody {
	s.UsageDataPerPage = v
	return s
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage struct {
	PageSize   *int32                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetPageSize(v int32) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetPageNumber(v int32) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetTotalCount(v int32) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage) SetData(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPage {
	s.Data = v
	return s
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData struct {
	DataItem []*DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData) SetDataItem(v []*DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageData {
	s.DataItem = v
	return s
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem struct {
	Status      *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string                                                                            `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DownloadUrl *string                                                                            `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	CreateTime  *string                                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskName    *string                                                                            `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskId      *string                                                                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskConfig  *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetStatus(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.Status = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetUpdateTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetDownloadUrl(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetCreateTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskName(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskName = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskId(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskId = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskConfig(v *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskConfig = v
	return s
}

type DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetEndTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetStartTime(v string) *DescribeUserUsageDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.StartTime = &v
	return s
}

type DescribeUserUsageDataExportTaskResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserUsageDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserUsageDataExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponse) SetHeaders(v map[string]*string) *DescribeUserUsageDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponse) SetBody(v *DescribeUserUsageDataExportTaskResponseBody) *DescribeUserUsageDataExportTaskResponse {
	s.Body = v
	return s
}

type DescribeUserUsageDetailDataExportTaskRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) SetOwnerId(v int64) *DescribeUserUsageDetailDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) SetPageSize(v string) *DescribeUserUsageDetailDataExportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) SetPageNumber(v string) *DescribeUserUsageDetailDataExportTaskRequest {
	s.PageNumber = &v
	return s
}

type DescribeUserUsageDetailDataExportTaskResponseBody struct {
	RequestId        *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageDataPerPage *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage `json:"UsageDataPerPage,omitempty" xml:"UsageDataPerPage,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *DescribeUserUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBody) SetUsageDataPerPage(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) *DescribeUserUsageDetailDataExportTaskResponseBody {
	s.UsageDataPerPage = v
	return s
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage struct {
	PageSize   *int32                                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Data       *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetPageSize(v int32) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetPageNumber(v int32) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetTotalCount(v int32) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage) SetData(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPage {
	s.Data = v
	return s
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData struct {
	DataItem []*DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData) SetDataItem(v []*DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageData {
	s.DataItem = v
	return s
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem struct {
	Status      *string                                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *string                                                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DownloadUrl *string                                                                                  `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	CreateTime  *string                                                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskName    *string                                                                                  `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskId      *string                                                                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskConfig  *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty" type:"Struct"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetStatus(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.Status = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetUpdateTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetDownloadUrl(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetCreateTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskName(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskName = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskId(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskId = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem) SetTaskConfig(v *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItem {
	s.TaskConfig = v
	return s
}

type DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetEndTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig) SetStartTime(v string) *DescribeUserUsageDetailDataExportTaskResponseBodyUsageDataPerPageDataDataItemTaskConfig {
	s.StartTime = &v
	return s
}

type DescribeUserUsageDetailDataExportTaskResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserUsageDetailDataExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *DescribeUserUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) SetBody(v *DescribeUserUsageDetailDataExportTaskResponseBody) *DescribeUserUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

type DescribeUserVipsByDomainRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Available  *string `json:"Available,omitempty" xml:"Available,omitempty"`
}

func (s DescribeUserVipsByDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVipsByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainRequest) SetOwnerId(v int64) *DescribeUserVipsByDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserVipsByDomainRequest) SetDomainName(v string) *DescribeUserVipsByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeUserVipsByDomainRequest) SetAvailable(v string) *DescribeUserVipsByDomainRequest {
	s.Available = &v
	return s
}

type DescribeUserVipsByDomainResponseBody struct {
	DomainName *string                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Vips       *DescribeUserVipsByDomainResponseBodyVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
}

func (s DescribeUserVipsByDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVipsByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainResponseBody) SetDomainName(v string) *DescribeUserVipsByDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeUserVipsByDomainResponseBody) SetRequestId(v string) *DescribeUserVipsByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserVipsByDomainResponseBody) SetVips(v *DescribeUserVipsByDomainResponseBodyVips) *DescribeUserVipsByDomainResponseBody {
	s.Vips = v
	return s
}

type DescribeUserVipsByDomainResponseBodyVips struct {
	Vip []*string `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeUserVipsByDomainResponseBodyVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVipsByDomainResponseBodyVips) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainResponseBodyVips) SetVip(v []*string) *DescribeUserVipsByDomainResponseBodyVips {
	s.Vip = v
	return s
}

type DescribeUserVipsByDomainResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserVipsByDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserVipsByDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVipsByDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainResponse) SetHeaders(v map[string]*string) *DescribeUserVipsByDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserVipsByDomainResponse) SetBody(v *DescribeUserVipsByDomainResponseBody) *DescribeUserVipsByDomainResponse {
	s.Body = v
	return s
}

type DescribeVerifyContentRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeVerifyContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentRequest) SetOwnerId(v int64) *DescribeVerifyContentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVerifyContentRequest) SetDomainName(v string) *DescribeVerifyContentRequest {
	s.DomainName = &v
	return s
}

type DescribeVerifyContentResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVerifyContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponseBody) SetContent(v string) *DescribeVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeVerifyContentResponseBody) SetRequestId(v string) *DescribeVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVerifyContentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVerifyContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyContentResponse) SetBody(v *DescribeVerifyContentResponseBody) *DescribeVerifyContentResponse {
	s.Body = v
	return s
}

type DisableRealtimeLogDeliveryRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DisableRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DisableRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DisableRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableRealtimeLogDeliveryRequest) SetDomain(v string) *DisableRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

type DisableRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DisableRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type DisableRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DisableRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DisableRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DisableRealtimeLogDeliveryResponse) SetBody(v *DisableRealtimeLogDeliveryResponseBody) *DisableRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type EnableRealtimeLogDeliveryRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s EnableRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *EnableRealtimeLogDeliveryRequest) SetOwnerId(v int64) *EnableRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableRealtimeLogDeliveryRequest) SetDomain(v string) *EnableRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

type EnableRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRealtimeLogDeliveryResponseBody) SetRequestId(v string) *EnableRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type EnableRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *EnableRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *EnableRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *EnableRealtimeLogDeliveryResponse) SetBody(v *EnableRealtimeLogDeliveryResponseBody) *EnableRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type ListDomainsByLogConfigIdRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s ListDomainsByLogConfigIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsByLogConfigIdRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdRequest) SetOwnerId(v int64) *ListDomainsByLogConfigIdRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDomainsByLogConfigIdRequest) SetConfigId(v string) *ListDomainsByLogConfigIdRequest {
	s.ConfigId = &v
	return s
}

type ListDomainsByLogConfigIdResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Domains   *ListDomainsByLogConfigIdResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
}

func (s ListDomainsByLogConfigIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsByLogConfigIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdResponseBody) SetRequestId(v string) *ListDomainsByLogConfigIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsByLogConfigIdResponseBody) SetDomains(v *ListDomainsByLogConfigIdResponseBodyDomains) *ListDomainsByLogConfigIdResponseBody {
	s.Domains = v
	return s
}

type ListDomainsByLogConfigIdResponseBodyDomains struct {
	Domain []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s ListDomainsByLogConfigIdResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsByLogConfigIdResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdResponseBodyDomains) SetDomain(v []*string) *ListDomainsByLogConfigIdResponseBodyDomains {
	s.Domain = v
	return s
}

type ListDomainsByLogConfigIdResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDomainsByLogConfigIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDomainsByLogConfigIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsByLogConfigIdResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsByLogConfigIdResponse) SetHeaders(v map[string]*string) *ListDomainsByLogConfigIdResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsByLogConfigIdResponse) SetBody(v *ListDomainsByLogConfigIdResponseBody) *ListDomainsByLogConfigIdResponse {
	s.Body = v
	return s
}

type ListFCTriggerRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	EventMetaName    *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
}

func (s ListFCTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *ListFCTriggerRequest) SetOwnerId(v int64) *ListFCTriggerRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFCTriggerRequest) SetEventMetaName(v string) *ListFCTriggerRequest {
	s.EventMetaName = &v
	return s
}

func (s *ListFCTriggerRequest) SetEventMetaVersion(v string) *ListFCTriggerRequest {
	s.EventMetaVersion = &v
	return s
}

type ListFCTriggerResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FCTriggers []*ListFCTriggerResponseBodyFCTriggers `json:"FCTriggers,omitempty" xml:"FCTriggers,omitempty" type:"Repeated"`
}

func (s ListFCTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *ListFCTriggerResponseBody) SetRequestId(v string) *ListFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFCTriggerResponseBody) SetFCTriggers(v []*ListFCTriggerResponseBodyFCTriggers) *ListFCTriggerResponseBody {
	s.FCTriggers = v
	return s
}

type ListFCTriggerResponseBodyFCTriggers struct {
	TriggerARN       *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
	RoleARN          *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	SourceArn        *string `json:"SourceArn,omitempty" xml:"SourceArn,omitempty"`
	Notes            *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	EventMetaName    *string `json:"EventMetaName,omitempty" xml:"EventMetaName,omitempty"`
	EventMetaVersion *string `json:"EventMetaVersion,omitempty" xml:"EventMetaVersion,omitempty"`
}

func (s ListFCTriggerResponseBodyFCTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListFCTriggerResponseBodyFCTriggers) GoString() string {
	return s.String()
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetTriggerARN(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.TriggerARN = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetRoleARN(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.RoleARN = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetSourceArn(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.SourceArn = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetNotes(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.Notes = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetEventMetaName(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.EventMetaName = &v
	return s
}

func (s *ListFCTriggerResponseBodyFCTriggers) SetEventMetaVersion(v string) *ListFCTriggerResponseBodyFCTriggers {
	s.EventMetaVersion = &v
	return s
}

type ListFCTriggerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFCTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *ListFCTriggerResponse) SetHeaders(v map[string]*string) *ListFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *ListFCTriggerResponse) SetBody(v *ListFCTriggerResponseBody) *ListFCTriggerResponse {
	s.Body = v
	return s
}

type ListRealtimeLogDeliveryDomainsRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListRealtimeLogDeliveryDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetOwnerId(v int64) *ListRealtimeLogDeliveryDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetProject(v string) *ListRealtimeLogDeliveryDomainsRequest {
	s.Project = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetLogstore(v string) *ListRealtimeLogDeliveryDomainsRequest {
	s.Logstore = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetRegion(v string) *ListRealtimeLogDeliveryDomainsRequest {
	s.Region = &v
	return s
}

type ListRealtimeLogDeliveryDomainsResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content   *ListRealtimeLogDeliveryDomainsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
}

func (s ListRealtimeLogDeliveryDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) SetRequestId(v string) *ListRealtimeLogDeliveryDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBody) SetContent(v *ListRealtimeLogDeliveryDomainsResponseBodyContent) *ListRealtimeLogDeliveryDomainsResponseBody {
	s.Content = v
	return s
}

type ListRealtimeLogDeliveryDomainsResponseBodyContent struct {
	Domains []*ListRealtimeLogDeliveryDomainsResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContent) SetDomains(v []*ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) *ListRealtimeLogDeliveryDomainsResponseBodyContent {
	s.Domains = v
	return s
}

type ListRealtimeLogDeliveryDomainsResponseBodyContentDomains struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetStatus(v string) *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.Status = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetDomainName(v string) *ListRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

type ListRealtimeLogDeliveryDomainsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRealtimeLogDeliveryDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRealtimeLogDeliveryDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponse) SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponse) SetBody(v *ListRealtimeLogDeliveryDomainsResponseBody) *ListRealtimeLogDeliveryDomainsResponse {
	s.Body = v
	return s
}

type ListRealtimeLogDeliveryInfosRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListRealtimeLogDeliveryInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosRequest) SetOwnerId(v int64) *ListRealtimeLogDeliveryInfosRequest {
	s.OwnerId = &v
	return s
}

type ListRealtimeLogDeliveryInfosResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content   *ListRealtimeLogDeliveryInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
}

func (s ListRealtimeLogDeliveryInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) SetRequestId(v string) *ListRealtimeLogDeliveryInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBody) SetContent(v *ListRealtimeLogDeliveryInfosResponseBodyContent) *ListRealtimeLogDeliveryInfosResponseBody {
	s.Content = v
	return s
}

type ListRealtimeLogDeliveryInfosResponseBodyContent struct {
	RealtimeLogDeliveryInfos []*ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos `json:"RealtimeLogDeliveryInfos,omitempty" xml:"RealtimeLogDeliveryInfos,omitempty" type:"Repeated"`
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContent) SetRealtimeLogDeliveryInfos(v []*ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) *ListRealtimeLogDeliveryInfosResponseBodyContent {
	s.RealtimeLogDeliveryInfos = v
	return s
}

type ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos struct {
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetLogstore(v string) *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Logstore = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetProject(v string) *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Project = &v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetRegion(v string) *ListRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Region = &v
	return s
}

type ListRealtimeLogDeliveryInfosResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRealtimeLogDeliveryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRealtimeLogDeliveryInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeLogDeliveryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryInfosResponse) SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryInfosResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeLogDeliveryInfosResponse) SetBody(v *ListRealtimeLogDeliveryInfosResponseBody) *ListRealtimeLogDeliveryInfosResponse {
	s.Body = v
	return s
}

type ListUserCustomLogConfigRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListUserCustomLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigRequest) SetOwnerId(v int64) *ListUserCustomLogConfigRequest {
	s.OwnerId = &v
	return s
}

type ListUserCustomLogConfigResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConfigIds *ListUserCustomLogConfigResponseBodyConfigIds `json:"ConfigIds,omitempty" xml:"ConfigIds,omitempty" type:"Struct"`
}

func (s ListUserCustomLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigResponseBody) SetRequestId(v string) *ListUserCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserCustomLogConfigResponseBody) SetConfigIds(v *ListUserCustomLogConfigResponseBodyConfigIds) *ListUserCustomLogConfigResponseBody {
	s.ConfigIds = v
	return s
}

type ListUserCustomLogConfigResponseBodyConfigIds struct {
	ConfigId []*string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty" type:"Repeated"`
}

func (s ListUserCustomLogConfigResponseBodyConfigIds) String() string {
	return tea.Prettify(s)
}

func (s ListUserCustomLogConfigResponseBodyConfigIds) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigResponseBodyConfigIds) SetConfigId(v []*string) *ListUserCustomLogConfigResponseBodyConfigIds {
	s.ConfigId = v
	return s
}

type ListUserCustomLogConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserCustomLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ListUserCustomLogConfigResponse) SetHeaders(v map[string]*string) *ListUserCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ListUserCustomLogConfigResponse) SetBody(v *ListUserCustomLogConfigResponseBody) *ListUserCustomLogConfigResponse {
	s.Body = v
	return s
}

type ModifyCdnDomainRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s ModifyCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainRequest) SetOwnerId(v int64) *ModifyCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetSecurityToken(v string) *ModifyCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetDomainName(v string) *ModifyCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetSources(v string) *ModifyCdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetResourceGroupId(v string) *ModifyCdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyCdnDomainRequest) SetTopLevelDomain(v string) *ModifyCdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type ModifyCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainResponseBody) SetRequestId(v string) *ModifyCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCdnDomainResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainResponse) SetHeaders(v map[string]*string) *ModifyCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdnDomainResponse) SetBody(v *ModifyCdnDomainResponseBody) *ModifyCdnDomainResponse {
	s.Body = v
	return s
}

type ModifyCdnDomainSchdmByPropertyRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Property   *string `json:"Property,omitempty" xml:"Property,omitempty"`
}

func (s ModifyCdnDomainSchdmByPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCdnDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) SetOwnerId(v int64) *ModifyCdnDomainSchdmByPropertyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyCdnDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyCdnDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

type ModifyCdnDomainSchdmByPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdnDomainSchdmByPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCdnDomainSchdmByPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainSchdmByPropertyResponseBody) SetRequestId(v string) *ModifyCdnDomainSchdmByPropertyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCdnDomainSchdmByPropertyResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCdnDomainSchdmByPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCdnDomainSchdmByPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCdnDomainSchdmByPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) SetHeaders(v map[string]*string) *ModifyCdnDomainSchdmByPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyResponse) SetBody(v *ModifyCdnDomainSchdmByPropertyResponseBody) *ModifyCdnDomainSchdmByPropertyResponse {
	s.Body = v
	return s
}

type ModifyDomainCustomLogConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId   *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s ModifyDomainCustomLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainCustomLogConfigRequest) SetOwnerId(v int64) *ModifyDomainCustomLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDomainCustomLogConfigRequest) SetDomainName(v string) *ModifyDomainCustomLogConfigRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyDomainCustomLogConfigRequest) SetConfigId(v string) *ModifyDomainCustomLogConfigRequest {
	s.ConfigId = &v
	return s
}

type ModifyDomainCustomLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainCustomLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainCustomLogConfigResponseBody) SetRequestId(v string) *ModifyDomainCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainCustomLogConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDomainCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainCustomLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainCustomLogConfigResponse) SetHeaders(v map[string]*string) *ModifyDomainCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainCustomLogConfigResponse) SetBody(v *ModifyDomainCustomLogConfigResponseBody) *ModifyDomainCustomLogConfigResponse {
	s.Body = v
	return s
}

type ModifyRealtimeLogDeliveryRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s ModifyRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ModifyRealtimeLogDeliveryRequest) SetOwnerId(v int64) *ModifyRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetProject(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetLogstore(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetRegion(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetDomain(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

type ModifyRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRealtimeLogDeliveryResponseBody) SetRequestId(v string) *ModifyRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ModifyRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *ModifyRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ModifyRealtimeLogDeliveryResponse) SetBody(v *ModifyRealtimeLogDeliveryResponseBody) *ModifyRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type ModifyUserCustomLogConfigRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ModifyUserCustomLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserCustomLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserCustomLogConfigRequest) SetOwnerId(v int64) *ModifyUserCustomLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUserCustomLogConfigRequest) SetConfigId(v string) *ModifyUserCustomLogConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyUserCustomLogConfigRequest) SetTag(v string) *ModifyUserCustomLogConfigRequest {
	s.Tag = &v
	return s
}

type ModifyUserCustomLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserCustomLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserCustomLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserCustomLogConfigResponseBody) SetRequestId(v string) *ModifyUserCustomLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserCustomLogConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserCustomLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserCustomLogConfigResponse) SetHeaders(v map[string]*string) *ModifyUserCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserCustomLogConfigResponse) SetBody(v *ModifyUserCustomLogConfigResponseBody) *ModifyUserCustomLogConfigResponse {
	s.Body = v
	return s
}

type OpenCdnServiceRequest struct {
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
}

func (s OpenCdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceRequest) SetOwnerId(v int64) *OpenCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenCdnServiceRequest) SetSecurityToken(v string) *OpenCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenCdnServiceRequest) SetInternetChargeType(v string) *OpenCdnServiceRequest {
	s.InternetChargeType = &v
	return s
}

type OpenCdnServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenCdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceResponseBody) SetRequestId(v string) *OpenCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenCdnServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenCdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenCdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCdnServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceResponse) SetHeaders(v map[string]*string) *OpenCdnServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCdnServiceResponse) SetBody(v *OpenCdnServiceResponseBody) *OpenCdnServiceResponse {
	s.Body = v
	return s
}

type PublishStagingConfigToProductionRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s PublishStagingConfigToProductionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishStagingConfigToProductionRequest) GoString() string {
	return s.String()
}

func (s *PublishStagingConfigToProductionRequest) SetOwnerId(v int64) *PublishStagingConfigToProductionRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishStagingConfigToProductionRequest) SetDomainName(v string) *PublishStagingConfigToProductionRequest {
	s.DomainName = &v
	return s
}

func (s *PublishStagingConfigToProductionRequest) SetFunctionName(v string) *PublishStagingConfigToProductionRequest {
	s.FunctionName = &v
	return s
}

type PublishStagingConfigToProductionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishStagingConfigToProductionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishStagingConfigToProductionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishStagingConfigToProductionResponseBody) SetRequestId(v string) *PublishStagingConfigToProductionResponseBody {
	s.RequestId = &v
	return s
}

type PublishStagingConfigToProductionResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishStagingConfigToProductionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishStagingConfigToProductionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishStagingConfigToProductionResponse) GoString() string {
	return s.String()
}

func (s *PublishStagingConfigToProductionResponse) SetHeaders(v map[string]*string) *PublishStagingConfigToProductionResponse {
	s.Headers = v
	return s
}

func (s *PublishStagingConfigToProductionResponse) SetBody(v *PublishStagingConfigToProductionResponseBody) *PublishStagingConfigToProductionResponse {
	s.Body = v
	return s
}

type PushObjectCacheRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	Area          *string `json:"Area,omitempty" xml:"Area,omitempty"`
}

func (s PushObjectCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s PushObjectCacheRequest) GoString() string {
	return s.String()
}

func (s *PushObjectCacheRequest) SetOwnerId(v int64) *PushObjectCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *PushObjectCacheRequest) SetSecurityToken(v string) *PushObjectCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *PushObjectCacheRequest) SetObjectPath(v string) *PushObjectCacheRequest {
	s.ObjectPath = &v
	return s
}

func (s *PushObjectCacheRequest) SetArea(v string) *PushObjectCacheRequest {
	s.Area = &v
	return s
}

type PushObjectCacheResponseBody struct {
	PushTaskId *string `json:"PushTaskId,omitempty" xml:"PushTaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushObjectCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushObjectCacheResponseBody) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponseBody) SetPushTaskId(v string) *PushObjectCacheResponseBody {
	s.PushTaskId = &v
	return s
}

func (s *PushObjectCacheResponseBody) SetRequestId(v string) *PushObjectCacheResponseBody {
	s.RequestId = &v
	return s
}

type PushObjectCacheResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PushObjectCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushObjectCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s PushObjectCacheResponse) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponse) SetHeaders(v map[string]*string) *PushObjectCacheResponse {
	s.Headers = v
	return s
}

func (s *PushObjectCacheResponse) SetBody(v *PushObjectCacheResponseBody) *PushObjectCacheResponse {
	s.Body = v
	return s
}

type RefreshObjectCachesRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s RefreshObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesRequest) SetOwnerId(v int64) *RefreshObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetSecurityToken(v string) *RefreshObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectPath(v string) *RefreshObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectType(v string) *RefreshObjectCachesRequest {
	s.ObjectType = &v
	return s
}

type RefreshObjectCachesResponseBody struct {
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshObjectCachesResponseBody) SetRequestId(v string) *RefreshObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type RefreshObjectCachesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshObjectCachesResponse) SetBody(v *RefreshObjectCachesResponseBody) *RefreshObjectCachesResponse {
	s.Body = v
	return s
}

type RollbackStagingConfigRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s RollbackStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *RollbackStagingConfigRequest) SetOwnerId(v int64) *RollbackStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *RollbackStagingConfigRequest) SetDomainName(v string) *RollbackStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *RollbackStagingConfigRequest) SetFunctionName(v string) *RollbackStagingConfigRequest {
	s.FunctionName = &v
	return s
}

type RollbackStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackStagingConfigResponseBody) SetRequestId(v string) *RollbackStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type RollbackStagingConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *RollbackStagingConfigResponse) SetHeaders(v map[string]*string) *RollbackStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *RollbackStagingConfigResponse) SetBody(v *RollbackStagingConfigResponseBody) *RollbackStagingConfigResponse {
	s.Body = v
	return s
}

type SetCcConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	AllowIps      *string `json:"AllowIps,omitempty" xml:"AllowIps,omitempty"`
	BlockIps      *string `json:"BlockIps,omitempty" xml:"BlockIps,omitempty"`
}

func (s SetCcConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCcConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCcConfigRequest) SetOwnerId(v int64) *SetCcConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCcConfigRequest) SetSecurityToken(v string) *SetCcConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetCcConfigRequest) SetDomainName(v string) *SetCcConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetCcConfigRequest) SetAllowIps(v string) *SetCcConfigRequest {
	s.AllowIps = &v
	return s
}

func (s *SetCcConfigRequest) SetBlockIps(v string) *SetCcConfigRequest {
	s.BlockIps = &v
	return s
}

type SetCcConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCcConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCcConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetCcConfigResponseBody) SetRequestId(v string) *SetCcConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetCcConfigResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetCcConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCcConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCcConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCcConfigResponse) SetHeaders(v map[string]*string) *SetCcConfigResponse {
	s.Headers = v
	return s
}

func (s *SetCcConfigResponse) SetBody(v *SetCcConfigResponseBody) *SetCcConfigResponse {
	s.Body = v
	return s
}

type SetCdnDomainCSRCertificateRequest struct {
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	DomainName        *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SetCdnDomainCSRCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCdnDomainCSRCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetCdnDomainCSRCertificateRequest) SetOwnerId(v int64) *SetCdnDomainCSRCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCdnDomainCSRCertificateRequest) SetServerCertificate(v string) *SetCdnDomainCSRCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *SetCdnDomainCSRCertificateRequest) SetDomainName(v string) *SetCdnDomainCSRCertificateRequest {
	s.DomainName = &v
	return s
}

type SetCdnDomainCSRCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnDomainCSRCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCdnDomainCSRCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnDomainCSRCertificateResponseBody) SetRequestId(v string) *SetCdnDomainCSRCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetCdnDomainCSRCertificateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetCdnDomainCSRCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCdnDomainCSRCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCdnDomainCSRCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetCdnDomainCSRCertificateResponse) SetHeaders(v map[string]*string) *SetCdnDomainCSRCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetCdnDomainCSRCertificateResponse) SetBody(v *SetCdnDomainCSRCertificateResponseBody) *SetCdnDomainCSRCertificateResponse {
	s.Body = v
	return s
}

type SetCdnDomainStagingConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Functions  *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
}

func (s SetCdnDomainStagingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigRequest) SetOwnerId(v int64) *SetCdnDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCdnDomainStagingConfigRequest) SetDomainName(v string) *SetCdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetCdnDomainStagingConfigRequest) SetFunctions(v string) *SetCdnDomainStagingConfigRequest {
	s.Functions = &v
	return s
}

type SetCdnDomainStagingConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnDomainStagingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigResponseBody) SetRequestId(v string) *SetCdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetCdnDomainStagingConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetCdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCdnDomainStagingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *SetCdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *SetCdnDomainStagingConfigResponse) SetBody(v *SetCdnDomainStagingConfigResponseBody) *SetCdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

type SetConfigOfVersionRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VersionId       *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	ConfigId        *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionId      *int64  `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	FunctionName    *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionArgs    *string `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty"`
	FunctionMatches *string `json:"FunctionMatches,omitempty" xml:"FunctionMatches,omitempty"`
}

func (s SetConfigOfVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetConfigOfVersionRequest) GoString() string {
	return s.String()
}

func (s *SetConfigOfVersionRequest) SetOwnerId(v int64) *SetConfigOfVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetOwnerAccount(v string) *SetConfigOfVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetSecurityToken(v string) *SetConfigOfVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetVersionId(v string) *SetConfigOfVersionRequest {
	s.VersionId = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetConfigId(v string) *SetConfigOfVersionRequest {
	s.ConfigId = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetFunctionId(v int64) *SetConfigOfVersionRequest {
	s.FunctionId = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetFunctionName(v string) *SetConfigOfVersionRequest {
	s.FunctionName = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetFunctionArgs(v string) *SetConfigOfVersionRequest {
	s.FunctionArgs = &v
	return s
}

func (s *SetConfigOfVersionRequest) SetFunctionMatches(v string) *SetConfigOfVersionRequest {
	s.FunctionMatches = &v
	return s
}

type SetConfigOfVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetConfigOfVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetConfigOfVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SetConfigOfVersionResponseBody) SetRequestId(v string) *SetConfigOfVersionResponseBody {
	s.RequestId = &v
	return s
}

type SetConfigOfVersionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetConfigOfVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetConfigOfVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetConfigOfVersionResponse) GoString() string {
	return s.String()
}

func (s *SetConfigOfVersionResponse) SetHeaders(v map[string]*string) *SetConfigOfVersionResponse {
	s.Headers = v
	return s
}

func (s *SetConfigOfVersionResponse) SetBody(v *SetConfigOfVersionResponseBody) *SetConfigOfVersionResponse {
	s.Body = v
	return s
}

type SetDomainGreenManagerConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s SetDomainGreenManagerConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainGreenManagerConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDomainGreenManagerConfigRequest) SetOwnerId(v int64) *SetDomainGreenManagerConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDomainGreenManagerConfigRequest) SetDomainName(v string) *SetDomainGreenManagerConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainGreenManagerConfigRequest) SetEnable(v string) *SetDomainGreenManagerConfigRequest {
	s.Enable = &v
	return s
}

type SetDomainGreenManagerConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainGreenManagerConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainGreenManagerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainGreenManagerConfigResponseBody) SetRequestId(v string) *SetDomainGreenManagerConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainGreenManagerConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainGreenManagerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainGreenManagerConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainGreenManagerConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDomainGreenManagerConfigResponse) SetHeaders(v map[string]*string) *SetDomainGreenManagerConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDomainGreenManagerConfigResponse) SetBody(v *SetDomainGreenManagerConfigResponseBody) *SetDomainGreenManagerConfigResponse {
	s.Body = v
	return s
}

type SetDomainServerCertificateRequest struct {
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken           *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	ServerCertificate       *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	PrivateKey              *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	ForceSet                *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
}

func (s SetDomainServerCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDomainServerCertificateRequest) SetOwnerId(v int64) *SetDomainServerCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetSecurityToken(v string) *SetDomainServerCertificateRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetDomainName(v string) *SetDomainServerCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetCertName(v string) *SetDomainServerCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetCertType(v string) *SetDomainServerCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetServerCertificateStatus(v string) *SetDomainServerCertificateRequest {
	s.ServerCertificateStatus = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetServerCertificate(v string) *SetDomainServerCertificateRequest {
	s.ServerCertificate = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetPrivateKey(v string) *SetDomainServerCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *SetDomainServerCertificateRequest) SetForceSet(v string) *SetDomainServerCertificateRequest {
	s.ForceSet = &v
	return s
}

type SetDomainServerCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainServerCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainServerCertificateResponseBody) SetRequestId(v string) *SetDomainServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainServerCertificateResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainServerCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDomainServerCertificateResponse) SetHeaders(v map[string]*string) *SetDomainServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDomainServerCertificateResponse) SetBody(v *SetDomainServerCertificateResponseBody) *SetDomainServerCertificateResponse {
	s.Body = v
	return s
}

type SetErrorPageConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PageType      *string `json:"PageType,omitempty" xml:"PageType,omitempty"`
	CustomPageUrl *string `json:"CustomPageUrl,omitempty" xml:"CustomPageUrl,omitempty"`
}

func (s SetErrorPageConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetErrorPageConfigRequest) GoString() string {
	return s.String()
}

func (s *SetErrorPageConfigRequest) SetOwnerId(v int64) *SetErrorPageConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetErrorPageConfigRequest) SetSecurityToken(v string) *SetErrorPageConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetErrorPageConfigRequest) SetDomainName(v string) *SetErrorPageConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetErrorPageConfigRequest) SetPageType(v string) *SetErrorPageConfigRequest {
	s.PageType = &v
	return s
}

func (s *SetErrorPageConfigRequest) SetCustomPageUrl(v string) *SetErrorPageConfigRequest {
	s.CustomPageUrl = &v
	return s
}

type SetErrorPageConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetErrorPageConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetErrorPageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetErrorPageConfigResponseBody) SetRequestId(v string) *SetErrorPageConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetErrorPageConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetErrorPageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetErrorPageConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetErrorPageConfigResponse) GoString() string {
	return s.String()
}

func (s *SetErrorPageConfigResponse) SetHeaders(v map[string]*string) *SetErrorPageConfigResponse {
	s.Headers = v
	return s
}

func (s *SetErrorPageConfigResponse) SetBody(v *SetErrorPageConfigResponseBody) *SetErrorPageConfigResponse {
	s.Body = v
	return s
}

type SetFileCacheExpiredConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	CacheContent  *string `json:"CacheContent,omitempty" xml:"CacheContent,omitempty"`
	TTL           *string `json:"TTL,omitempty" xml:"TTL,omitempty"`
	Weight        *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SetFileCacheExpiredConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFileCacheExpiredConfigRequest) GoString() string {
	return s.String()
}

func (s *SetFileCacheExpiredConfigRequest) SetOwnerId(v int64) *SetFileCacheExpiredConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetFileCacheExpiredConfigRequest) SetSecurityToken(v string) *SetFileCacheExpiredConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetFileCacheExpiredConfigRequest) SetDomainName(v string) *SetFileCacheExpiredConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetFileCacheExpiredConfigRequest) SetCacheContent(v string) *SetFileCacheExpiredConfigRequest {
	s.CacheContent = &v
	return s
}

func (s *SetFileCacheExpiredConfigRequest) SetTTL(v string) *SetFileCacheExpiredConfigRequest {
	s.TTL = &v
	return s
}

func (s *SetFileCacheExpiredConfigRequest) SetWeight(v string) *SetFileCacheExpiredConfigRequest {
	s.Weight = &v
	return s
}

type SetFileCacheExpiredConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFileCacheExpiredConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFileCacheExpiredConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetFileCacheExpiredConfigResponseBody) SetRequestId(v string) *SetFileCacheExpiredConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetFileCacheExpiredConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetFileCacheExpiredConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetFileCacheExpiredConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFileCacheExpiredConfigResponse) GoString() string {
	return s.String()
}

func (s *SetFileCacheExpiredConfigResponse) SetHeaders(v map[string]*string) *SetFileCacheExpiredConfigResponse {
	s.Headers = v
	return s
}

func (s *SetFileCacheExpiredConfigResponse) SetBody(v *SetFileCacheExpiredConfigResponseBody) *SetFileCacheExpiredConfigResponse {
	s.Body = v
	return s
}

type SetForceRedirectConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RedirectType  *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
}

func (s SetForceRedirectConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetForceRedirectConfigRequest) GoString() string {
	return s.String()
}

func (s *SetForceRedirectConfigRequest) SetOwnerId(v int64) *SetForceRedirectConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetForceRedirectConfigRequest) SetSecurityToken(v string) *SetForceRedirectConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetForceRedirectConfigRequest) SetDomainName(v string) *SetForceRedirectConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetForceRedirectConfigRequest) SetRedirectType(v string) *SetForceRedirectConfigRequest {
	s.RedirectType = &v
	return s
}

type SetForceRedirectConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetForceRedirectConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetForceRedirectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetForceRedirectConfigResponseBody) SetRequestId(v string) *SetForceRedirectConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetForceRedirectConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetForceRedirectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetForceRedirectConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetForceRedirectConfigResponse) GoString() string {
	return s.String()
}

func (s *SetForceRedirectConfigResponse) SetHeaders(v map[string]*string) *SetForceRedirectConfigResponse {
	s.Headers = v
	return s
}

func (s *SetForceRedirectConfigResponse) SetBody(v *SetForceRedirectConfigResponseBody) *SetForceRedirectConfigResponse {
	s.Body = v
	return s
}

type SetForwardSchemeConfigRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId         *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enable           *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	SchemeOrigin     *string `json:"SchemeOrigin,omitempty" xml:"SchemeOrigin,omitempty"`
	SchemeOriginPort *string `json:"SchemeOriginPort,omitempty" xml:"SchemeOriginPort,omitempty"`
}

func (s SetForwardSchemeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetForwardSchemeConfigRequest) GoString() string {
	return s.String()
}

func (s *SetForwardSchemeConfigRequest) SetOwnerId(v int64) *SetForwardSchemeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetForwardSchemeConfigRequest) SetDomainName(v string) *SetForwardSchemeConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetForwardSchemeConfigRequest) SetConfigId(v int64) *SetForwardSchemeConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetForwardSchemeConfigRequest) SetEnable(v string) *SetForwardSchemeConfigRequest {
	s.Enable = &v
	return s
}

func (s *SetForwardSchemeConfigRequest) SetSchemeOrigin(v string) *SetForwardSchemeConfigRequest {
	s.SchemeOrigin = &v
	return s
}

func (s *SetForwardSchemeConfigRequest) SetSchemeOriginPort(v string) *SetForwardSchemeConfigRequest {
	s.SchemeOriginPort = &v
	return s
}

type SetForwardSchemeConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetForwardSchemeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetForwardSchemeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetForwardSchemeConfigResponseBody) SetRequestId(v string) *SetForwardSchemeConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetForwardSchemeConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetForwardSchemeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetForwardSchemeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetForwardSchemeConfigResponse) GoString() string {
	return s.String()
}

func (s *SetForwardSchemeConfigResponse) SetHeaders(v map[string]*string) *SetForwardSchemeConfigResponse {
	s.Headers = v
	return s
}

func (s *SetForwardSchemeConfigResponse) SetBody(v *SetForwardSchemeConfigResponseBody) *SetForwardSchemeConfigResponse {
	s.Body = v
	return s
}

type SetHttpErrorPageConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	PageUrl    *string `json:"PageUrl,omitempty" xml:"PageUrl,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s SetHttpErrorPageConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetHttpErrorPageConfigRequest) GoString() string {
	return s.String()
}

func (s *SetHttpErrorPageConfigRequest) SetOwnerId(v int64) *SetHttpErrorPageConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetHttpErrorPageConfigRequest) SetDomainName(v string) *SetHttpErrorPageConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetHttpErrorPageConfigRequest) SetErrorCode(v string) *SetHttpErrorPageConfigRequest {
	s.ErrorCode = &v
	return s
}

func (s *SetHttpErrorPageConfigRequest) SetPageUrl(v string) *SetHttpErrorPageConfigRequest {
	s.PageUrl = &v
	return s
}

func (s *SetHttpErrorPageConfigRequest) SetConfigId(v int64) *SetHttpErrorPageConfigRequest {
	s.ConfigId = &v
	return s
}

type SetHttpErrorPageConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetHttpErrorPageConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetHttpErrorPageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpErrorPageConfigResponseBody) SetRequestId(v string) *SetHttpErrorPageConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetHttpErrorPageConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetHttpErrorPageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetHttpErrorPageConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetHttpErrorPageConfigResponse) GoString() string {
	return s.String()
}

func (s *SetHttpErrorPageConfigResponse) SetHeaders(v map[string]*string) *SetHttpErrorPageConfigResponse {
	s.Headers = v
	return s
}

func (s *SetHttpErrorPageConfigResponse) SetBody(v *SetHttpErrorPageConfigResponseBody) *SetHttpErrorPageConfigResponse {
	s.Body = v
	return s
}

type SetHttpHeaderConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	HeaderKey     *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	HeaderValue   *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
	ConfigId      *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s SetHttpHeaderConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetHttpHeaderConfigRequest) GoString() string {
	return s.String()
}

func (s *SetHttpHeaderConfigRequest) SetOwnerId(v int64) *SetHttpHeaderConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetHttpHeaderConfigRequest) SetSecurityToken(v string) *SetHttpHeaderConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetHttpHeaderConfigRequest) SetDomainName(v string) *SetHttpHeaderConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetHttpHeaderConfigRequest) SetHeaderKey(v string) *SetHttpHeaderConfigRequest {
	s.HeaderKey = &v
	return s
}

func (s *SetHttpHeaderConfigRequest) SetHeaderValue(v string) *SetHttpHeaderConfigRequest {
	s.HeaderValue = &v
	return s
}

func (s *SetHttpHeaderConfigRequest) SetConfigId(v int64) *SetHttpHeaderConfigRequest {
	s.ConfigId = &v
	return s
}

type SetHttpHeaderConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetHttpHeaderConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetHttpHeaderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpHeaderConfigResponseBody) SetRequestId(v string) *SetHttpHeaderConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetHttpHeaderConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetHttpHeaderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetHttpHeaderConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetHttpHeaderConfigResponse) GoString() string {
	return s.String()
}

func (s *SetHttpHeaderConfigResponse) SetHeaders(v map[string]*string) *SetHttpHeaderConfigResponse {
	s.Headers = v
	return s
}

func (s *SetHttpHeaderConfigResponse) SetBody(v *SetHttpHeaderConfigResponseBody) *SetHttpHeaderConfigResponse {
	s.Body = v
	return s
}

type SetHttpsOptionConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Http2      *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
}

func (s SetHttpsOptionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetHttpsOptionConfigRequest) GoString() string {
	return s.String()
}

func (s *SetHttpsOptionConfigRequest) SetOwnerId(v int64) *SetHttpsOptionConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetHttpsOptionConfigRequest) SetDomainName(v string) *SetHttpsOptionConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetHttpsOptionConfigRequest) SetConfigId(v int64) *SetHttpsOptionConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetHttpsOptionConfigRequest) SetHttp2(v string) *SetHttpsOptionConfigRequest {
	s.Http2 = &v
	return s
}

type SetHttpsOptionConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetHttpsOptionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetHttpsOptionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpsOptionConfigResponseBody) SetRequestId(v string) *SetHttpsOptionConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetHttpsOptionConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetHttpsOptionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetHttpsOptionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetHttpsOptionConfigResponse) GoString() string {
	return s.String()
}

func (s *SetHttpsOptionConfigResponse) SetHeaders(v map[string]*string) *SetHttpsOptionConfigResponse {
	s.Headers = v
	return s
}

func (s *SetHttpsOptionConfigResponse) SetBody(v *SetHttpsOptionConfigResponseBody) *SetHttpsOptionConfigResponse {
	s.Body = v
	return s
}

type SetIgnoreQueryStringConfigRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId    *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enable      *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	HashKeyArgs *string `json:"HashKeyArgs,omitempty" xml:"HashKeyArgs,omitempty"`
	KeepOssArgs *string `json:"KeepOssArgs,omitempty" xml:"KeepOssArgs,omitempty"`
}

func (s SetIgnoreQueryStringConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIgnoreQueryStringConfigRequest) GoString() string {
	return s.String()
}

func (s *SetIgnoreQueryStringConfigRequest) SetOwnerId(v int64) *SetIgnoreQueryStringConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetIgnoreQueryStringConfigRequest) SetDomainName(v string) *SetIgnoreQueryStringConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetIgnoreQueryStringConfigRequest) SetConfigId(v int64) *SetIgnoreQueryStringConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetIgnoreQueryStringConfigRequest) SetEnable(v string) *SetIgnoreQueryStringConfigRequest {
	s.Enable = &v
	return s
}

func (s *SetIgnoreQueryStringConfigRequest) SetHashKeyArgs(v string) *SetIgnoreQueryStringConfigRequest {
	s.HashKeyArgs = &v
	return s
}

func (s *SetIgnoreQueryStringConfigRequest) SetKeepOssArgs(v string) *SetIgnoreQueryStringConfigRequest {
	s.KeepOssArgs = &v
	return s
}

type SetIgnoreQueryStringConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIgnoreQueryStringConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetIgnoreQueryStringConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetIgnoreQueryStringConfigResponseBody) SetRequestId(v string) *SetIgnoreQueryStringConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetIgnoreQueryStringConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetIgnoreQueryStringConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetIgnoreQueryStringConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIgnoreQueryStringConfigResponse) GoString() string {
	return s.String()
}

func (s *SetIgnoreQueryStringConfigResponse) SetHeaders(v map[string]*string) *SetIgnoreQueryStringConfigResponse {
	s.Headers = v
	return s
}

func (s *SetIgnoreQueryStringConfigResponse) SetBody(v *SetIgnoreQueryStringConfigResponseBody) *SetIgnoreQueryStringConfigResponse {
	s.Body = v
	return s
}

type SetIpAllowListConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	AllowIps      *string `json:"AllowIps,omitempty" xml:"AllowIps,omitempty"`
}

func (s SetIpAllowListConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIpAllowListConfigRequest) GoString() string {
	return s.String()
}

func (s *SetIpAllowListConfigRequest) SetOwnerId(v int64) *SetIpAllowListConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetIpAllowListConfigRequest) SetSecurityToken(v string) *SetIpAllowListConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetIpAllowListConfigRequest) SetDomainName(v string) *SetIpAllowListConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetIpAllowListConfigRequest) SetAllowIps(v string) *SetIpAllowListConfigRequest {
	s.AllowIps = &v
	return s
}

type SetIpAllowListConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIpAllowListConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetIpAllowListConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetIpAllowListConfigResponseBody) SetRequestId(v string) *SetIpAllowListConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetIpAllowListConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetIpAllowListConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetIpAllowListConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIpAllowListConfigResponse) GoString() string {
	return s.String()
}

func (s *SetIpAllowListConfigResponse) SetHeaders(v map[string]*string) *SetIpAllowListConfigResponse {
	s.Headers = v
	return s
}

func (s *SetIpAllowListConfigResponse) SetBody(v *SetIpAllowListConfigResponseBody) *SetIpAllowListConfigResponse {
	s.Body = v
	return s
}

type SetIpBlackListConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	BlockIps   *string `json:"BlockIps,omitempty" xml:"BlockIps,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s SetIpBlackListConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIpBlackListConfigRequest) GoString() string {
	return s.String()
}

func (s *SetIpBlackListConfigRequest) SetOwnerId(v int64) *SetIpBlackListConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetIpBlackListConfigRequest) SetDomainName(v string) *SetIpBlackListConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetIpBlackListConfigRequest) SetBlockIps(v string) *SetIpBlackListConfigRequest {
	s.BlockIps = &v
	return s
}

func (s *SetIpBlackListConfigRequest) SetConfigId(v int64) *SetIpBlackListConfigRequest {
	s.ConfigId = &v
	return s
}

type SetIpBlackListConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIpBlackListConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetIpBlackListConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetIpBlackListConfigResponseBody) SetRequestId(v string) *SetIpBlackListConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetIpBlackListConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetIpBlackListConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetIpBlackListConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIpBlackListConfigResponse) GoString() string {
	return s.String()
}

func (s *SetIpBlackListConfigResponse) SetHeaders(v map[string]*string) *SetIpBlackListConfigResponse {
	s.Headers = v
	return s
}

func (s *SetIpBlackListConfigResponse) SetBody(v *SetIpBlackListConfigResponseBody) *SetIpBlackListConfigResponse {
	s.Body = v
	return s
}

type SetOptimizeConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s SetOptimizeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOptimizeConfigRequest) GoString() string {
	return s.String()
}

func (s *SetOptimizeConfigRequest) SetOwnerId(v int64) *SetOptimizeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetOptimizeConfigRequest) SetDomainName(v string) *SetOptimizeConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetOptimizeConfigRequest) SetConfigId(v int64) *SetOptimizeConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetOptimizeConfigRequest) SetEnable(v string) *SetOptimizeConfigRequest {
	s.Enable = &v
	return s
}

type SetOptimizeConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetOptimizeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetOptimizeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetOptimizeConfigResponseBody) SetRequestId(v string) *SetOptimizeConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetOptimizeConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetOptimizeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetOptimizeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOptimizeConfigResponse) GoString() string {
	return s.String()
}

func (s *SetOptimizeConfigResponse) SetHeaders(v map[string]*string) *SetOptimizeConfigResponse {
	s.Headers = v
	return s
}

func (s *SetOptimizeConfigResponse) SetBody(v *SetOptimizeConfigResponseBody) *SetOptimizeConfigResponse {
	s.Body = v
	return s
}

type SetPageCompressConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s SetPageCompressConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPageCompressConfigRequest) GoString() string {
	return s.String()
}

func (s *SetPageCompressConfigRequest) SetOwnerId(v int64) *SetPageCompressConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetPageCompressConfigRequest) SetDomainName(v string) *SetPageCompressConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetPageCompressConfigRequest) SetConfigId(v int64) *SetPageCompressConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetPageCompressConfigRequest) SetEnable(v string) *SetPageCompressConfigRequest {
	s.Enable = &v
	return s
}

type SetPageCompressConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPageCompressConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPageCompressConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetPageCompressConfigResponseBody) SetRequestId(v string) *SetPageCompressConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetPageCompressConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetPageCompressConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetPageCompressConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPageCompressConfigResponse) GoString() string {
	return s.String()
}

func (s *SetPageCompressConfigResponse) SetHeaders(v map[string]*string) *SetPageCompressConfigResponse {
	s.Headers = v
	return s
}

func (s *SetPageCompressConfigResponse) SetBody(v *SetPageCompressConfigResponseBody) *SetPageCompressConfigResponse {
	s.Body = v
	return s
}

type SetRangeConfigRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s SetRangeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRangeConfigRequest) GoString() string {
	return s.String()
}

func (s *SetRangeConfigRequest) SetOwnerId(v int64) *SetRangeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRangeConfigRequest) SetDomainName(v string) *SetRangeConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetRangeConfigRequest) SetEnable(v string) *SetRangeConfigRequest {
	s.Enable = &v
	return s
}

func (s *SetRangeConfigRequest) SetConfigId(v int64) *SetRangeConfigRequest {
	s.ConfigId = &v
	return s
}

type SetRangeConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRangeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRangeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetRangeConfigResponseBody) SetRequestId(v string) *SetRangeConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetRangeConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRangeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRangeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRangeConfigResponse) GoString() string {
	return s.String()
}

func (s *SetRangeConfigResponse) SetHeaders(v map[string]*string) *SetRangeConfigResponse {
	s.Headers = v
	return s
}

func (s *SetRangeConfigResponse) SetBody(v *SetRangeConfigResponseBody) *SetRangeConfigResponse {
	s.Body = v
	return s
}

type SetRefererConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ReferType     *string `json:"ReferType,omitempty" xml:"ReferType,omitempty"`
	ReferList     *string `json:"ReferList,omitempty" xml:"ReferList,omitempty"`
	AllowEmpty    *string `json:"AllowEmpty,omitempty" xml:"AllowEmpty,omitempty"`
	DisableAst    *string `json:"DisableAst,omitempty" xml:"DisableAst,omitempty"`
}

func (s SetRefererConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRefererConfigRequest) GoString() string {
	return s.String()
}

func (s *SetRefererConfigRequest) SetOwnerId(v int64) *SetRefererConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRefererConfigRequest) SetSecurityToken(v string) *SetRefererConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetRefererConfigRequest) SetDomainName(v string) *SetRefererConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetRefererConfigRequest) SetReferType(v string) *SetRefererConfigRequest {
	s.ReferType = &v
	return s
}

func (s *SetRefererConfigRequest) SetReferList(v string) *SetRefererConfigRequest {
	s.ReferList = &v
	return s
}

func (s *SetRefererConfigRequest) SetAllowEmpty(v string) *SetRefererConfigRequest {
	s.AllowEmpty = &v
	return s
}

func (s *SetRefererConfigRequest) SetDisableAst(v string) *SetRefererConfigRequest {
	s.DisableAst = &v
	return s
}

type SetRefererConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRefererConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRefererConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetRefererConfigResponseBody) SetRequestId(v string) *SetRefererConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetRefererConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRefererConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRefererConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRefererConfigResponse) GoString() string {
	return s.String()
}

func (s *SetRefererConfigResponse) SetHeaders(v map[string]*string) *SetRefererConfigResponse {
	s.Headers = v
	return s
}

func (s *SetRefererConfigResponse) SetBody(v *SetRefererConfigResponseBody) *SetRefererConfigResponse {
	s.Body = v
	return s
}

type SetRemoveQueryStringConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ConfigId      *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	AliRemoveArgs *string `json:"AliRemoveArgs,omitempty" xml:"AliRemoveArgs,omitempty"`
	KeepOssArgs   *string `json:"KeepOssArgs,omitempty" xml:"KeepOssArgs,omitempty"`
}

func (s SetRemoveQueryStringConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRemoveQueryStringConfigRequest) GoString() string {
	return s.String()
}

func (s *SetRemoveQueryStringConfigRequest) SetOwnerId(v int64) *SetRemoveQueryStringConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRemoveQueryStringConfigRequest) SetDomainName(v string) *SetRemoveQueryStringConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetRemoveQueryStringConfigRequest) SetConfigId(v int64) *SetRemoveQueryStringConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *SetRemoveQueryStringConfigRequest) SetAliRemoveArgs(v string) *SetRemoveQueryStringConfigRequest {
	s.AliRemoveArgs = &v
	return s
}

func (s *SetRemoveQueryStringConfigRequest) SetKeepOssArgs(v string) *SetRemoveQueryStringConfigRequest {
	s.KeepOssArgs = &v
	return s
}

type SetRemoveQueryStringConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRemoveQueryStringConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRemoveQueryStringConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetRemoveQueryStringConfigResponseBody) SetRequestId(v string) *SetRemoveQueryStringConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetRemoveQueryStringConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRemoveQueryStringConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRemoveQueryStringConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRemoveQueryStringConfigResponse) GoString() string {
	return s.String()
}

func (s *SetRemoveQueryStringConfigResponse) SetHeaders(v map[string]*string) *SetRemoveQueryStringConfigResponse {
	s.Headers = v
	return s
}

func (s *SetRemoveQueryStringConfigResponse) SetBody(v *SetRemoveQueryStringConfigResponseBody) *SetRemoveQueryStringConfigResponse {
	s.Body = v
	return s
}

type SetReqAuthConfigRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	AuthType       *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Key1           *string `json:"Key1,omitempty" xml:"Key1,omitempty"`
	Key2           *string `json:"Key2,omitempty" xml:"Key2,omitempty"`
	TimeOut        *string `json:"TimeOut,omitempty" xml:"TimeOut,omitempty"`
	AuthRemoteDesc *string `json:"AuthRemoteDesc,omitempty" xml:"AuthRemoteDesc,omitempty"`
}

func (s SetReqAuthConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetReqAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *SetReqAuthConfigRequest) SetOwnerId(v int64) *SetReqAuthConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetSecurityToken(v string) *SetReqAuthConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetDomainName(v string) *SetReqAuthConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetAuthType(v string) *SetReqAuthConfigRequest {
	s.AuthType = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetKey1(v string) *SetReqAuthConfigRequest {
	s.Key1 = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetKey2(v string) *SetReqAuthConfigRequest {
	s.Key2 = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetTimeOut(v string) *SetReqAuthConfigRequest {
	s.TimeOut = &v
	return s
}

func (s *SetReqAuthConfigRequest) SetAuthRemoteDesc(v string) *SetReqAuthConfigRequest {
	s.AuthRemoteDesc = &v
	return s
}

type SetReqAuthConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetReqAuthConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetReqAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetReqAuthConfigResponseBody) SetRequestId(v string) *SetReqAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetReqAuthConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetReqAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetReqAuthConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetReqAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *SetReqAuthConfigResponse) SetHeaders(v map[string]*string) *SetReqAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *SetReqAuthConfigResponse) SetBody(v *SetReqAuthConfigResponseBody) *SetReqAuthConfigResponse {
	s.Body = v
	return s
}

type SetReqHeaderConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ConfigId      *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s SetReqHeaderConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetReqHeaderConfigRequest) GoString() string {
	return s.String()
}

func (s *SetReqHeaderConfigRequest) SetOwnerId(v int64) *SetReqHeaderConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetSecurityToken(v string) *SetReqHeaderConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetDomainName(v string) *SetReqHeaderConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetKey(v string) *SetReqHeaderConfigRequest {
	s.Key = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetValue(v string) *SetReqHeaderConfigRequest {
	s.Value = &v
	return s
}

func (s *SetReqHeaderConfigRequest) SetConfigId(v int64) *SetReqHeaderConfigRequest {
	s.ConfigId = &v
	return s
}

type SetReqHeaderConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetReqHeaderConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetReqHeaderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetReqHeaderConfigResponseBody) SetRequestId(v string) *SetReqHeaderConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetReqHeaderConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetReqHeaderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetReqHeaderConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetReqHeaderConfigResponse) GoString() string {
	return s.String()
}

func (s *SetReqHeaderConfigResponse) SetHeaders(v map[string]*string) *SetReqHeaderConfigResponse {
	s.Headers = v
	return s
}

func (s *SetReqHeaderConfigResponse) SetBody(v *SetReqHeaderConfigResponseBody) *SetReqHeaderConfigResponse {
	s.Body = v
	return s
}

type SetSourceHostConfigRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Enable        *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	BackSrcDomain *string `json:"BackSrcDomain,omitempty" xml:"BackSrcDomain,omitempty"`
}

func (s SetSourceHostConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSourceHostConfigRequest) GoString() string {
	return s.String()
}

func (s *SetSourceHostConfigRequest) SetOwnerId(v int64) *SetSourceHostConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetSourceHostConfigRequest) SetSecurityToken(v string) *SetSourceHostConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetSourceHostConfigRequest) SetDomainName(v string) *SetSourceHostConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetSourceHostConfigRequest) SetEnable(v string) *SetSourceHostConfigRequest {
	s.Enable = &v
	return s
}

func (s *SetSourceHostConfigRequest) SetBackSrcDomain(v string) *SetSourceHostConfigRequest {
	s.BackSrcDomain = &v
	return s
}

type SetSourceHostConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSourceHostConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSourceHostConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetSourceHostConfigResponseBody) SetRequestId(v string) *SetSourceHostConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetSourceHostConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetSourceHostConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetSourceHostConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSourceHostConfigResponse) GoString() string {
	return s.String()
}

func (s *SetSourceHostConfigResponse) SetHeaders(v map[string]*string) *SetSourceHostConfigResponse {
	s.Headers = v
	return s
}

func (s *SetSourceHostConfigResponse) SetBody(v *SetSourceHostConfigResponseBody) *SetSourceHostConfigResponse {
	s.Body = v
	return s
}

type SetWaitingRoomConfigRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName  *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	WaitUri     *string `json:"WaitUri,omitempty" xml:"WaitUri,omitempty"`
	AllowPct    *int32  `json:"AllowPct,omitempty" xml:"AllowPct,omitempty"`
	MaxTimeWait *int32  `json:"MaxTimeWait,omitempty" xml:"MaxTimeWait,omitempty"`
	GapTime     *int32  `json:"GapTime,omitempty" xml:"GapTime,omitempty"`
	WaitUrl     *string `json:"WaitUrl,omitempty" xml:"WaitUrl,omitempty"`
}

func (s SetWaitingRoomConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetWaitingRoomConfigRequest) GoString() string {
	return s.String()
}

func (s *SetWaitingRoomConfigRequest) SetOwnerId(v int64) *SetWaitingRoomConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetDomainName(v string) *SetWaitingRoomConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetWaitUri(v string) *SetWaitingRoomConfigRequest {
	s.WaitUri = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetAllowPct(v int32) *SetWaitingRoomConfigRequest {
	s.AllowPct = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetMaxTimeWait(v int32) *SetWaitingRoomConfigRequest {
	s.MaxTimeWait = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetGapTime(v int32) *SetWaitingRoomConfigRequest {
	s.GapTime = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetWaitUrl(v string) *SetWaitingRoomConfigRequest {
	s.WaitUrl = &v
	return s
}

type SetWaitingRoomConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetWaitingRoomConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetWaitingRoomConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetWaitingRoomConfigResponseBody) SetRequestId(v string) *SetWaitingRoomConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetWaitingRoomConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetWaitingRoomConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetWaitingRoomConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetWaitingRoomConfigResponse) GoString() string {
	return s.String()
}

func (s *SetWaitingRoomConfigResponse) SetHeaders(v map[string]*string) *SetWaitingRoomConfigResponse {
	s.Headers = v
	return s
}

func (s *SetWaitingRoomConfigResponse) SetBody(v *SetWaitingRoomConfigResponseBody) *SetWaitingRoomConfigResponse {
	s.Body = v
	return s
}

type StartCdnDomainRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s StartCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StartCdnDomainRequest) SetOwnerId(v int64) *StartCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartCdnDomainRequest) SetSecurityToken(v string) *StartCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StartCdnDomainRequest) SetDomainName(v string) *StartCdnDomainRequest {
	s.DomainName = &v
	return s
}

type StartCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartCdnDomainResponseBody) SetRequestId(v string) *StartCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type StartCdnDomainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StartCdnDomainResponse) SetHeaders(v map[string]*string) *StartCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StartCdnDomainResponse) SetBody(v *StartCdnDomainResponseBody) *StartCdnDomainResponse {
	s.Body = v
	return s
}

type StopCdnDomainRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s StopCdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StopCdnDomainRequest) SetOwnerId(v int64) *StopCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopCdnDomainRequest) SetSecurityToken(v string) *StopCdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *StopCdnDomainRequest) SetDomainName(v string) *StopCdnDomainRequest {
	s.DomainName = &v
	return s
}

type StopCdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopCdnDomainResponseBody) SetRequestId(v string) *StopCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type StopCdnDomainResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StopCdnDomainResponse) SetHeaders(v map[string]*string) *StopCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StopCdnDomainResponse) SetBody(v *StopCdnDomainResponseBody) *StopCdnDomainResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId      *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateCdnDeliverTaskRequest struct {
	OwnerId    *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DeliverId  *int64                 `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	Name       *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Status     *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Reports    *string                `json:"Reports,omitempty" xml:"Reports,omitempty"`
	DomainName *string                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Deliver    map[string]interface{} `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	Schedule   map[string]interface{} `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateCdnDeliverTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnDeliverTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskRequest) SetOwnerId(v int64) *UpdateCdnDeliverTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetDeliverId(v int64) *UpdateCdnDeliverTaskRequest {
	s.DeliverId = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetName(v string) *UpdateCdnDeliverTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetStatus(v string) *UpdateCdnDeliverTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetReports(v string) *UpdateCdnDeliverTaskRequest {
	s.Reports = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetDomainName(v string) *UpdateCdnDeliverTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetDeliver(v map[string]interface{}) *UpdateCdnDeliverTaskRequest {
	s.Deliver = v
	return s
}

func (s *UpdateCdnDeliverTaskRequest) SetSchedule(v map[string]interface{}) *UpdateCdnDeliverTaskRequest {
	s.Schedule = v
	return s
}

type UpdateCdnDeliverTaskShrinkRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DeliverId      *int64  `json:"DeliverId,omitempty" xml:"DeliverId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Reports        *string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DeliverShrink  *string `json:"Deliver,omitempty" xml:"Deliver,omitempty"`
	ScheduleShrink *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateCdnDeliverTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnDeliverTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetOwnerId(v int64) *UpdateCdnDeliverTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetDeliverId(v int64) *UpdateCdnDeliverTaskShrinkRequest {
	s.DeliverId = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetName(v string) *UpdateCdnDeliverTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetStatus(v string) *UpdateCdnDeliverTaskShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetReports(v string) *UpdateCdnDeliverTaskShrinkRequest {
	s.Reports = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetDomainName(v string) *UpdateCdnDeliverTaskShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetDeliverShrink(v string) *UpdateCdnDeliverTaskShrinkRequest {
	s.DeliverShrink = &v
	return s
}

func (s *UpdateCdnDeliverTaskShrinkRequest) SetScheduleShrink(v string) *UpdateCdnDeliverTaskShrinkRequest {
	s.ScheduleShrink = &v
	return s
}

type UpdateCdnDeliverTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCdnDeliverTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnDeliverTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskResponseBody) SetRequestId(v string) *UpdateCdnDeliverTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCdnDeliverTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCdnDeliverTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCdnDeliverTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnDeliverTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateCdnDeliverTaskResponse) SetHeaders(v map[string]*string) *UpdateCdnDeliverTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateCdnDeliverTaskResponse) SetBody(v *UpdateCdnDeliverTaskResponseBody) *UpdateCdnDeliverTaskResponse {
	s.Body = v
	return s
}

type UpdateCdnSubTaskRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ReportIds  *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s UpdateCdnSubTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateCdnSubTaskRequest) SetOwnerId(v int64) *UpdateCdnSubTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetStatus(v string) *UpdateCdnSubTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetReportIds(v string) *UpdateCdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetDomainName(v string) *UpdateCdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetStartTime(v string) *UpdateCdnSubTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetEndTime(v string) *UpdateCdnSubTaskRequest {
	s.EndTime = &v
	return s
}

type UpdateCdnSubTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCdnSubTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCdnSubTaskResponseBody) SetRequestId(v string) *UpdateCdnSubTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCdnSubTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCdnSubTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateCdnSubTaskResponse) SetHeaders(v map[string]*string) *UpdateCdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateCdnSubTaskResponse) SetBody(v *UpdateCdnSubTaskResponseBody) *UpdateCdnSubTaskResponse {
	s.Body = v
	return s
}

type UpdateFCTriggerRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TriggerARN  *string `json:"TriggerARN,omitempty" xml:"TriggerARN,omitempty"`
	SourceARN   *string `json:"SourceARN,omitempty" xml:"SourceARN,omitempty"`
	FunctionARN *string `json:"FunctionARN,omitempty" xml:"FunctionARN,omitempty"`
	RoleARN     *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	Notes       *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
}

func (s UpdateFCTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFCTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateFCTriggerRequest) SetOwnerId(v int64) *UpdateFCTriggerRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetTriggerARN(v string) *UpdateFCTriggerRequest {
	s.TriggerARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetSourceARN(v string) *UpdateFCTriggerRequest {
	s.SourceARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetFunctionARN(v string) *UpdateFCTriggerRequest {
	s.FunctionARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetRoleARN(v string) *UpdateFCTriggerRequest {
	s.RoleARN = &v
	return s
}

func (s *UpdateFCTriggerRequest) SetNotes(v string) *UpdateFCTriggerRequest {
	s.Notes = &v
	return s
}

type UpdateFCTriggerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFCTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFCTriggerResponseBody) SetRequestId(v string) *UpdateFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFCTriggerResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFCTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateFCTriggerResponse) SetHeaders(v map[string]*string) *UpdateFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateFCTriggerResponse) SetBody(v *UpdateFCTriggerResponseBody) *UpdateFCTriggerResponse {
	s.Body = v
	return s
}

type VerifyDomainOwnerRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerRequest) SetOwnerId(v int64) *VerifyDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyDomainOwnerRequest) SetDomainName(v string) *VerifyDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyDomainOwnerRequest) SetVerifyType(v string) *VerifyDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

type VerifyDomainOwnerResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponseBody) SetContent(v string) *VerifyDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyDomainOwnerResponseBody) SetRequestId(v string) *VerifyDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyDomainOwnerResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyDomainOwnerResponse) SetBody(v *VerifyDomainOwnerResponseBody) *VerifyDomainOwnerResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":     tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":   tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"me-east-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"us-east-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
		"us-west-1":      tea.String("cdn.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddCdnDomainWithOptions(request *AddCdnDomainRequest, runtime *util.RuntimeOptions) (_result *AddCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCdnDomain(request *AddCdnDomainRequest) (_result *AddCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCdnDomainResponse{}
	_body, _err := client.AddCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFCTriggerWithOptions(request *AddFCTriggerRequest, runtime *util.RuntimeOptions) (_result *AddFCTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddFCTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddFCTrigger"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddFCTrigger(request *AddFCTriggerRequest) (_result *AddFCTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFCTriggerResponse{}
	_body, _err := client.AddFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddCdnDomainWithOptions(request *BatchAddCdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchAddCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchAddCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchAddCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddCdnDomain(request *BatchAddCdnDomainRequest) (_result *BatchAddCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddCdnDomainResponse{}
	_body, _err := client.BatchAddCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetCdnDomainConfigWithOptions(request *BatchSetCdnDomainConfigRequest, runtime *util.RuntimeOptions) (_result *BatchSetCdnDomainConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchSetCdnDomainConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchSetCdnDomainConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetCdnDomainConfig(request *BatchSetCdnDomainConfigRequest) (_result *BatchSetCdnDomainConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetCdnDomainConfigResponse{}
	_body, _err := client.BatchSetCdnDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetCdnDomainServerCertificateWithOptions(request *BatchSetCdnDomainServerCertificateRequest, runtime *util.RuntimeOptions) (_result *BatchSetCdnDomainServerCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchSetCdnDomainServerCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchSetCdnDomainServerCertificate"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetCdnDomainServerCertificate(request *BatchSetCdnDomainServerCertificateRequest) (_result *BatchSetCdnDomainServerCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetCdnDomainServerCertificateResponse{}
	_body, _err := client.BatchSetCdnDomainServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartCdnDomainWithOptions(request *BatchStartCdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStartCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchStartCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchStartCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartCdnDomain(request *BatchStartCdnDomainRequest) (_result *BatchStartCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStartCdnDomainResponse{}
	_body, _err := client.BatchStartCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopCdnDomainWithOptions(request *BatchStopCdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStopCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchStopCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchStopCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopCdnDomain(request *BatchStopCdnDomainRequest) (_result *BatchStopCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopCdnDomainResponse{}
	_body, _err := client.BatchStopCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateCdnDomainWithOptions(request *BatchUpdateCdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchUpdateCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchUpdateCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchUpdateCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateCdnDomain(request *BatchUpdateCdnDomainRequest) (_result *BatchUpdateCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUpdateCdnDomainResponse{}
	_body, _err := client.BatchUpdateCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCdnCertificateSigningRequestWithOptions(request *CreateCdnCertificateSigningRequestRequest, runtime *util.RuntimeOptions) (_result *CreateCdnCertificateSigningRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCdnCertificateSigningRequestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCdnCertificateSigningRequest"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCdnCertificateSigningRequest(request *CreateCdnCertificateSigningRequestRequest) (_result *CreateCdnCertificateSigningRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCdnCertificateSigningRequestResponse{}
	_body, _err := client.CreateCdnCertificateSigningRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCdnDeliverTaskWithOptions(tmpReq *CreateCdnDeliverTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCdnDeliverTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCdnDeliverTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Deliver)) {
		request.DeliverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Deliver, tea.String("Deliver"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Schedule)) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, tea.String("Schedule"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCdnDeliverTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCdnDeliverTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCdnDeliverTask(request *CreateCdnDeliverTaskRequest) (_result *CreateCdnDeliverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCdnDeliverTaskResponse{}
	_body, _err := client.CreateCdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCdnSubTaskWithOptions(request *CreateCdnSubTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCdnSubTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCdnSubTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCdnSubTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCdnSubTask(request *CreateCdnSubTaskRequest) (_result *CreateCdnSubTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCdnSubTaskResponse{}
	_body, _err := client.CreateCdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIllegalUrlExportTaskWithOptions(request *CreateIllegalUrlExportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateIllegalUrlExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateIllegalUrlExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateIllegalUrlExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIllegalUrlExportTask(request *CreateIllegalUrlExportTaskRequest) (_result *CreateIllegalUrlExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIllegalUrlExportTaskResponse{}
	_body, _err := client.CreateIllegalUrlExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRealTimeLogDeliveryWithOptions(request *CreateRealTimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *CreateRealTimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateRealTimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRealTimeLogDelivery"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRealTimeLogDelivery(request *CreateRealTimeLogDeliveryRequest) (_result *CreateRealTimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRealTimeLogDeliveryResponse{}
	_body, _err := client.CreateRealTimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUsageDetailDataExportTaskWithOptions(request *CreateUsageDetailDataExportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateUsageDetailDataExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUsageDetailDataExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUsageDetailDataExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUsageDetailDataExportTask(request *CreateUsageDetailDataExportTaskRequest) (_result *CreateUsageDetailDataExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUsageDetailDataExportTaskResponse{}
	_body, _err := client.CreateUsageDetailDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserUsageDataExportTaskWithOptions(request *CreateUserUsageDataExportTaskRequest, runtime *util.RuntimeOptions) (_result *CreateUserUsageDataExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserUsageDataExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUserUsageDataExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserUsageDataExportTask(request *CreateUserUsageDataExportTaskRequest) (_result *CreateUserUsageDataExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserUsageDataExportTaskResponse{}
	_body, _err := client.CreateUserUsageDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCdnDeliverTaskWithOptions(request *DeleteCdnDeliverTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteCdnDeliverTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCdnDeliverTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCdnDeliverTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCdnDeliverTask(request *DeleteCdnDeliverTaskRequest) (_result *DeleteCdnDeliverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCdnDeliverTaskResponse{}
	_body, _err := client.DeleteCdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCdnDomainWithOptions(request *DeleteCdnDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) (_result *DeleteCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCdnDomainResponse{}
	_body, _err := client.DeleteCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCdnSubTaskWithOptions(request *DeleteCdnSubTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteCdnSubTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCdnSubTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCdnSubTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCdnSubTask(request *DeleteCdnSubTaskRequest) (_result *DeleteCdnSubTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCdnSubTaskResponse{}
	_body, _err := client.DeleteCdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFCTriggerWithOptions(request *DeleteFCTriggerRequest, runtime *util.RuntimeOptions) (_result *DeleteFCTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFCTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFCTrigger"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFCTrigger(request *DeleteFCTriggerRequest) (_result *DeleteFCTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFCTriggerResponse{}
	_body, _err := client.DeleteFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRealtimeLogDeliveryWithOptions(request *DeleteRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DeleteRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRealtimeLogDelivery"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRealtimeLogDelivery(request *DeleteRealtimeLogDeliveryRequest) (_result *DeleteRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRealtimeLogDeliveryResponse{}
	_body, _err := client.DeleteRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSpecificConfigWithOptions(request *DeleteSpecificConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteSpecificConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSpecificConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSpecificConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpecificConfig(request *DeleteSpecificConfigRequest) (_result *DeleteSpecificConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSpecificConfigResponse{}
	_body, _err := client.DeleteSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSpecificStagingConfigWithOptions(request *DeleteSpecificStagingConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteSpecificStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSpecificStagingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSpecificStagingConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSpecificStagingConfig(request *DeleteSpecificStagingConfigRequest) (_result *DeleteSpecificStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSpecificStagingConfigResponse{}
	_body, _err := client.DeleteSpecificStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUsageDetailDataExportTaskWithOptions(request *DeleteUsageDetailDataExportTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteUsageDetailDataExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUsageDetailDataExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUsageDetailDataExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUsageDetailDataExportTask(request *DeleteUsageDetailDataExportTaskRequest) (_result *DeleteUsageDetailDataExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUsageDetailDataExportTaskResponse{}
	_body, _err := client.DeleteUsageDetailDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserUsageDataExportTaskWithOptions(request *DeleteUserUsageDataExportTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteUserUsageDataExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserUsageDataExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserUsageDataExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserUsageDataExportTask(request *DeleteUserUsageDataExportTaskRequest) (_result *DeleteUserUsageDataExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserUsageDataExportTaskResponse{}
	_body, _err := client.DeleteUserUsageDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeActiveVersionOfConfigGroupWithOptions(request *DescribeActiveVersionOfConfigGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveVersionOfConfigGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeActiveVersionOfConfigGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeActiveVersionOfConfigGroup"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeActiveVersionOfConfigGroup(request *DescribeActiveVersionOfConfigGroupRequest) (_result *DescribeActiveVersionOfConfigGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveVersionOfConfigGroupResponse{}
	_body, _err := client.DescribeActiveVersionOfConfigGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnCertificateDetailWithOptions(request *DescribeCdnCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnCertificateDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnCertificateDetail"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnCertificateDetail(request *DescribeCdnCertificateDetailRequest) (_result *DescribeCdnCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnCertificateDetailResponse{}
	_body, _err := client.DescribeCdnCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnCertificateListWithOptions(request *DescribeCdnCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnCertificateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnCertificateList"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnCertificateList(request *DescribeCdnCertificateListRequest) (_result *DescribeCdnCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnCertificateListResponse{}
	_body, _err := client.DescribeCdnCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDeliverListWithOptions(request *DescribeCdnDeliverListRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDeliverListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnDeliverListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnDeliverList"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDeliverList(request *DescribeCdnDeliverListRequest) (_result *DescribeCdnDeliverListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDeliverListResponse{}
	_body, _err := client.DescribeCdnDeliverListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainByCertificateWithOptions(request *DescribeCdnDomainByCertificateRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainByCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnDomainByCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnDomainByCertificate"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainByCertificate(request *DescribeCdnDomainByCertificateRequest) (_result *DescribeCdnDomainByCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainByCertificateResponse{}
	_body, _err := client.DescribeCdnDomainByCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainConfigsWithOptions(request *DescribeCdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnDomainConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnDomainConfigs"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainConfigs(request *DescribeCdnDomainConfigsRequest) (_result *DescribeCdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainConfigsResponse{}
	_body, _err := client.DescribeCdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainDetailWithOptions(request *DescribeCdnDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnDomainDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnDomainDetail"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainDetail(request *DescribeCdnDomainDetailRequest) (_result *DescribeCdnDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainDetailResponse{}
	_body, _err := client.DescribeCdnDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainLogsWithOptions(request *DescribeCdnDomainLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnDomainLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnDomainLogs"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (_result *DescribeCdnDomainLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainLogsResponse{}
	_body, _err := client.DescribeCdnDomainLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnDomainStagingConfigWithOptions(request *DescribeCdnDomainStagingConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnDomainStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnDomainStagingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnDomainStagingConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnDomainStagingConfig(request *DescribeCdnDomainStagingConfigRequest) (_result *DescribeCdnDomainStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnDomainStagingConfigResponse{}
	_body, _err := client.DescribeCdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnHttpsDomainListWithOptions(request *DescribeCdnHttpsDomainListRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnHttpsDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnHttpsDomainListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnHttpsDomainList"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnHttpsDomainList(request *DescribeCdnHttpsDomainListRequest) (_result *DescribeCdnHttpsDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnHttpsDomainListResponse{}
	_body, _err := client.DescribeCdnHttpsDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnRegionAndIspWithOptions(request *DescribeCdnRegionAndIspRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnRegionAndIspResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnRegionAndIspResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnRegionAndIsp"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnRegionAndIsp(request *DescribeCdnRegionAndIspRequest) (_result *DescribeCdnRegionAndIspResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnRegionAndIspResponse{}
	_body, _err := client.DescribeCdnRegionAndIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnReportWithOptions(request *DescribeCdnReportRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnReport"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnReport(request *DescribeCdnReportRequest) (_result *DescribeCdnReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnReportResponse{}
	_body, _err := client.DescribeCdnReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnReportListWithOptions(request *DescribeCdnReportListRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnReportListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnReportListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnReportList"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnReportList(request *DescribeCdnReportListRequest) (_result *DescribeCdnReportListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnReportListResponse{}
	_body, _err := client.DescribeCdnReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnServiceWithOptions(request *DescribeCdnServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnService"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnService(request *DescribeCdnServiceRequest) (_result *DescribeCdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnServiceResponse{}
	_body, _err := client.DescribeCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnSubListWithOptions(request *DescribeCdnSubListRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnSubListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnSubListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnSubList"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnSubList(request *DescribeCdnSubListRequest) (_result *DescribeCdnSubListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnSubListResponse{}
	_body, _err := client.DescribeCdnSubListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserBillHistoryWithOptions(request *DescribeCdnUserBillHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserBillHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserBillHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserBillHistory"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserBillHistory(request *DescribeCdnUserBillHistoryRequest) (_result *DescribeCdnUserBillHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserBillHistoryResponse{}
	_body, _err := client.DescribeCdnUserBillHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserBillPredictionWithOptions(request *DescribeCdnUserBillPredictionRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserBillPredictionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserBillPredictionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserBillPrediction"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserBillPrediction(request *DescribeCdnUserBillPredictionRequest) (_result *DescribeCdnUserBillPredictionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserBillPredictionResponse{}
	_body, _err := client.DescribeCdnUserBillPredictionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserBillTypeWithOptions(request *DescribeCdnUserBillTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserBillTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserBillTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserBillType"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserBillType(request *DescribeCdnUserBillTypeRequest) (_result *DescribeCdnUserBillTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserBillTypeResponse{}
	_body, _err := client.DescribeCdnUserBillTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserConfigsWithOptions(request *DescribeCdnUserConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserConfigs"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserConfigs(request *DescribeCdnUserConfigsRequest) (_result *DescribeCdnUserConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserConfigsResponse{}
	_body, _err := client.DescribeCdnUserConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserDomainsByFuncWithOptions(request *DescribeCdnUserDomainsByFuncRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserDomainsByFuncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserDomainsByFuncResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserDomainsByFunc"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserDomainsByFunc(request *DescribeCdnUserDomainsByFuncRequest) (_result *DescribeCdnUserDomainsByFuncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserDomainsByFuncResponse{}
	_body, _err := client.DescribeCdnUserDomainsByFuncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserQuotaWithOptions(request *DescribeCdnUserQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserQuota"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserQuota(request *DescribeCdnUserQuotaRequest) (_result *DescribeCdnUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserQuotaResponse{}
	_body, _err := client.DescribeCdnUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnUserResourcePackageWithOptions(request *DescribeCdnUserResourcePackageRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnUserResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnUserResourcePackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnUserResourcePackage"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnUserResourcePackage(request *DescribeCdnUserResourcePackageRequest) (_result *DescribeCdnUserResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnUserResourcePackageResponse{}
	_body, _err := client.DescribeCdnUserResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCdnWafDomainWithOptions(request *DescribeCdnWafDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeCdnWafDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCdnWafDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCdnWafDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCdnWafDomain(request *DescribeCdnWafDomainRequest) (_result *DescribeCdnWafDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCdnWafDomainResponse{}
	_body, _err := client.DescribeCdnWafDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertificateInfoByIDWithOptions(request *DescribeCertificateInfoByIDRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificateInfoByIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeCertificateInfoByIDResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCertificateInfoByID"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCertificateInfoByID(request *DescribeCertificateInfoByIDRequest) (_result *DescribeCertificateInfoByIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificateInfoByIDResponse{}
	_body, _err := client.DescribeCertificateInfoByIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigOfVersionWithOptions(request *DescribeConfigOfVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigOfVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConfigOfVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConfigOfVersion"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigOfVersion(request *DescribeConfigOfVersionRequest) (_result *DescribeConfigOfVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigOfVersionResponse{}
	_body, _err := client.DescribeConfigOfVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomLogConfigWithOptions(request *DescribeCustomLogConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeCustomLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomLogConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomLogConfig(request *DescribeCustomLogConfigRequest) (_result *DescribeCustomLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomLogConfigResponse{}
	_body, _err := client.DescribeCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAverageResponseTimeWithOptions(request *DescribeDomainAverageResponseTimeRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAverageResponseTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainAverageResponseTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainAverageResponseTime"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAverageResponseTime(request *DescribeDomainAverageResponseTimeRequest) (_result *DescribeDomainAverageResponseTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAverageResponseTimeResponse{}
	_body, _err := client.DescribeDomainAverageResponseTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataWithOptions(request *DescribeDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainBpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainBpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainBpsData(request *DescribeDomainBpsDataRequest) (_result *DescribeDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainBpsDataResponse{}
	_body, _err := client.DescribeDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataByLayerWithOptions(request *DescribeDomainBpsDataByLayerRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainBpsDataByLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainBpsDataByLayerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainBpsDataByLayer"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataByLayer(request *DescribeDomainBpsDataByLayerRequest) (_result *DescribeDomainBpsDataByLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainBpsDataByLayerResponse{}
	_body, _err := client.DescribeDomainBpsDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataByTimeStampWithOptions(request *DescribeDomainBpsDataByTimeStampRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainBpsDataByTimeStampResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainBpsDataByTimeStampResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainBpsDataByTimeStamp"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainBpsDataByTimeStamp(request *DescribeDomainBpsDataByTimeStampRequest) (_result *DescribeDomainBpsDataByTimeStampResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainBpsDataByTimeStampResponse{}
	_body, _err := client.DescribeDomainBpsDataByTimeStampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainCcActivityLogWithOptions(request *DescribeDomainCcActivityLogRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainCcActivityLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainCcActivityLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainCcActivityLog"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainCcActivityLog(request *DescribeDomainCcActivityLogRequest) (_result *DescribeDomainCcActivityLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainCcActivityLogResponse{}
	_body, _err := client.DescribeDomainCcActivityLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainCertificateInfoWithOptions(request *DescribeDomainCertificateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainCertificateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainCertificateInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainCertificateInfo"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainCertificateInfo(request *DescribeDomainCertificateInfoRequest) (_result *DescribeDomainCertificateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainCertificateInfoResponse{}
	_body, _err := client.DescribeDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainCustomLogConfigWithOptions(request *DescribeDomainCustomLogConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainCustomLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainCustomLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainCustomLogConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainCustomLogConfig(request *DescribeDomainCustomLogConfigRequest) (_result *DescribeDomainCustomLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainCustomLogConfigResponse{}
	_body, _err := client.DescribeDomainCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainDetailDataByLayerWithOptions(request *DescribeDomainDetailDataByLayerRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainDetailDataByLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainDetailDataByLayerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainDetailDataByLayer"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainDetailDataByLayer(request *DescribeDomainDetailDataByLayerRequest) (_result *DescribeDomainDetailDataByLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainDetailDataByLayerResponse{}
	_body, _err := client.DescribeDomainDetailDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainFileSizeProportionDataWithOptions(request *DescribeDomainFileSizeProportionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainFileSizeProportionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainFileSizeProportionDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainFileSizeProportionData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainFileSizeProportionData(request *DescribeDomainFileSizeProportionDataRequest) (_result *DescribeDomainFileSizeProportionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainFileSizeProportionDataResponse{}
	_body, _err := client.DescribeDomainFileSizeProportionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainHitRateDataWithOptions(request *DescribeDomainHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainHitRateDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainHitRateData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainHitRateData(request *DescribeDomainHitRateDataRequest) (_result *DescribeDomainHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainHitRateDataResponse{}
	_body, _err := client.DescribeDomainHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainHttpCodeDataWithOptions(request *DescribeDomainHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainHttpCodeDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainHttpCodeData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainHttpCodeData(request *DescribeDomainHttpCodeDataRequest) (_result *DescribeDomainHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainHttpCodeDataByLayerWithOptions(request *DescribeDomainHttpCodeDataByLayerRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainHttpCodeDataByLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainHttpCodeDataByLayer"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainHttpCodeDataByLayer(request *DescribeDomainHttpCodeDataByLayerRequest) (_result *DescribeDomainHttpCodeDataByLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.DescribeDomainHttpCodeDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainISPDataWithOptions(request *DescribeDomainISPDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainISPDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainISPDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainISPData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainISPData(request *DescribeDomainISPDataRequest) (_result *DescribeDomainISPDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainISPDataResponse{}
	_body, _err := client.DescribeDomainISPDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainMax95BpsDataWithOptions(request *DescribeDomainMax95BpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainMax95BpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainMax95BpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainMax95BpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainMax95BpsData(request *DescribeDomainMax95BpsDataRequest) (_result *DescribeDomainMax95BpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainMax95BpsDataResponse{}
	_body, _err := client.DescribeDomainMax95BpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainNamesOfVersionWithOptions(request *DescribeDomainNamesOfVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainNamesOfVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainNamesOfVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainNamesOfVersion"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainNamesOfVersion(request *DescribeDomainNamesOfVersionRequest) (_result *DescribeDomainNamesOfVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainNamesOfVersionResponse{}
	_body, _err := client.DescribeDomainNamesOfVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainPathDataWithOptions(request *DescribeDomainPathDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainPathDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainPathDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainPathData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainPathData(request *DescribeDomainPathDataRequest) (_result *DescribeDomainPathDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainPathDataResponse{}
	_body, _err := client.DescribeDomainPathDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainPvDataWithOptions(request *DescribeDomainPvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainPvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainPvDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainPvData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainPvData(request *DescribeDomainPvDataRequest) (_result *DescribeDomainPvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainPvDataResponse{}
	_body, _err := client.DescribeDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsDataWithOptions(request *DescribeDomainQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainQpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainQpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQpsData(request *DescribeDomainQpsDataRequest) (_result *DescribeDomainQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsDataResponse{}
	_body, _err := client.DescribeDomainQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsDataByLayerWithOptions(request *DescribeDomainQpsDataByLayerRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsDataByLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainQpsDataByLayerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainQpsDataByLayer"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQpsDataByLayer(request *DescribeDomainQpsDataByLayerRequest) (_result *DescribeDomainQpsDataByLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsDataByLayerResponse{}
	_body, _err := client.DescribeDomainQpsDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeBpsDataWithOptions(request *DescribeDomainRealTimeBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainRealTimeBpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeBpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeBpsData(request *DescribeDomainRealTimeBpsDataRequest) (_result *DescribeDomainRealTimeBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeBpsDataResponse{}
	_body, _err := client.DescribeDomainRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeByteHitRateDataWithOptions(request *DescribeDomainRealTimeByteHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeByteHitRateData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeByteHitRateData(request *DescribeDomainRealTimeByteHitRateDataRequest) (_result *DescribeDomainRealTimeByteHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.DescribeDomainRealTimeByteHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeDetailDataWithOptions(request *DescribeDomainRealTimeDetailDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeDetailDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainRealTimeDetailDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeDetailData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeDetailData(request *DescribeDomainRealTimeDetailDataRequest) (_result *DescribeDomainRealTimeDetailDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeDetailDataResponse{}
	_body, _err := client.DescribeDomainRealTimeDetailDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeHttpCodeDataWithOptions(request *DescribeDomainRealTimeHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeHttpCodeData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeHttpCodeData(request *DescribeDomainRealTimeHttpCodeDataRequest) (_result *DescribeDomainRealTimeHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainRealTimeHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealtimeLogDeliveryWithOptions(request *DescribeDomainRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealtimeLogDelivery"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealtimeLogDelivery(request *DescribeDomainRealtimeLogDeliveryRequest) (_result *DescribeDomainRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DescribeDomainRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeQpsDataWithOptions(request *DescribeDomainRealTimeQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainRealTimeQpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeQpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeQpsData(request *DescribeDomainRealTimeQpsDataRequest) (_result *DescribeDomainRealTimeQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeQpsDataResponse{}
	_body, _err := client.DescribeDomainRealTimeQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeReqHitRateDataWithOptions(request *DescribeDomainRealTimeReqHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeReqHitRateData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeReqHitRateData(request *DescribeDomainRealTimeReqHitRateDataRequest) (_result *DescribeDomainRealTimeReqHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.DescribeDomainRealTimeReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeSrcBpsDataWithOptions(request *DescribeDomainRealTimeSrcBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeSrcBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeSrcBpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeSrcBpsData(request *DescribeDomainRealTimeSrcBpsDataRequest) (_result *DescribeDomainRealTimeSrcBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.DescribeDomainRealTimeSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeSrcHttpCodeDataWithOptions(request *DescribeDomainRealTimeSrcHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeSrcHttpCodeData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeSrcHttpCodeData(request *DescribeDomainRealTimeSrcHttpCodeDataRequest) (_result *DescribeDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainRealTimeSrcHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeSrcTrafficDataWithOptions(request *DescribeDomainRealTimeSrcTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeSrcTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeSrcTrafficData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeSrcTrafficData(request *DescribeDomainRealTimeSrcTrafficDataRequest) (_result *DescribeDomainRealTimeSrcTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.DescribeDomainRealTimeSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeTrafficDataWithOptions(request *DescribeDomainRealTimeTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRealTimeTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRealTimeTrafficData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRealTimeTrafficData(request *DescribeDomainRealTimeTrafficDataRequest) (_result *DescribeDomainRealTimeTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DescribeDomainRealTimeTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainRegionDataWithOptions(request *DescribeDomainRegionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainRegionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainRegionDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainRegionData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainRegionData(request *DescribeDomainRegionDataRequest) (_result *DescribeDomainRegionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainRegionDataResponse{}
	_body, _err := client.DescribeDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainReqHitRateDataWithOptions(request *DescribeDomainReqHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainReqHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainReqHitRateDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainReqHitRateData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainReqHitRateData(request *DescribeDomainReqHitRateDataRequest) (_result *DescribeDomainReqHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainReqHitRateDataResponse{}
	_body, _err := client.DescribeDomainReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsBySourceWithOptions(request *DescribeDomainsBySourceRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsBySourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainsBySourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainsBySource"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainsBySource(request *DescribeDomainsBySourceRequest) (_result *DescribeDomainsBySourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsBySourceResponse{}
	_body, _err := client.DescribeDomainsBySourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcBpsDataWithOptions(request *DescribeDomainSrcBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSrcBpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSrcBpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcBpsData(request *DescribeDomainSrcBpsDataRequest) (_result *DescribeDomainSrcBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcBpsDataResponse{}
	_body, _err := client.DescribeDomainSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcHttpCodeDataWithOptions(request *DescribeDomainSrcHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSrcHttpCodeDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSrcHttpCodeData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcHttpCodeData(request *DescribeDomainSrcHttpCodeDataRequest) (_result *DescribeDomainSrcHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainSrcHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcQpsDataWithOptions(request *DescribeDomainSrcQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSrcQpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSrcQpsData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcQpsData(request *DescribeDomainSrcQpsDataRequest) (_result *DescribeDomainSrcQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcQpsDataResponse{}
	_body, _err := client.DescribeDomainSrcQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcTopUrlVisitWithOptions(request *DescribeDomainSrcTopUrlVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcTopUrlVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSrcTopUrlVisitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSrcTopUrlVisit"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcTopUrlVisit(request *DescribeDomainSrcTopUrlVisitRequest) (_result *DescribeDomainSrcTopUrlVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcTopUrlVisitResponse{}
	_body, _err := client.DescribeDomainSrcTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSrcTrafficDataWithOptions(request *DescribeDomainSrcTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSrcTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSrcTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSrcTrafficData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSrcTrafficData(request *DescribeDomainSrcTrafficDataRequest) (_result *DescribeDomainSrcTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSrcTrafficDataResponse{}
	_body, _err := client.DescribeDomainSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsUsageByDayWithOptions(request *DescribeDomainsUsageByDayRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsUsageByDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainsUsageByDayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainsUsageByDay"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainsUsageByDay(request *DescribeDomainsUsageByDayRequest) (_result *DescribeDomainsUsageByDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsUsageByDayResponse{}
	_body, _err := client.DescribeDomainsUsageByDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainTopClientIpVisitWithOptions(request *DescribeDomainTopClientIpVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainTopClientIpVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainTopClientIpVisitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainTopClientIpVisit"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainTopClientIpVisit(request *DescribeDomainTopClientIpVisitRequest) (_result *DescribeDomainTopClientIpVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainTopClientIpVisitResponse{}
	_body, _err := client.DescribeDomainTopClientIpVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainTopReferVisitWithOptions(request *DescribeDomainTopReferVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainTopReferVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainTopReferVisitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainTopReferVisit"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainTopReferVisit(request *DescribeDomainTopReferVisitRequest) (_result *DescribeDomainTopReferVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainTopReferVisitResponse{}
	_body, _err := client.DescribeDomainTopReferVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainTopUrlVisitWithOptions(request *DescribeDomainTopUrlVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainTopUrlVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainTopUrlVisitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainTopUrlVisit"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainTopUrlVisit(request *DescribeDomainTopUrlVisitRequest) (_result *DescribeDomainTopUrlVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainTopUrlVisitResponse{}
	_body, _err := client.DescribeDomainTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainTrafficDataWithOptions(request *DescribeDomainTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainTrafficData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainTrafficData(request *DescribeDomainTrafficDataRequest) (_result *DescribeDomainTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainTrafficDataResponse{}
	_body, _err := client.DescribeDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainUsageDataWithOptions(request *DescribeDomainUsageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainUsageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainUsageDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainUsageData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainUsageData(request *DescribeDomainUsageDataRequest) (_result *DescribeDomainUsageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainUsageDataResponse{}
	_body, _err := client.DescribeDomainUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainUvDataWithOptions(request *DescribeDomainUvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainUvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainUvDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainUvData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainUvData(request *DescribeDomainUvDataRequest) (_result *DescribeDomainUvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainUvDataResponse{}
	_body, _err := client.DescribeDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEsExceptionDataWithOptions(request *DescribeEsExceptionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEsExceptionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEsExceptionDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEsExceptionData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEsExceptionData(request *DescribeEsExceptionDataRequest) (_result *DescribeEsExceptionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEsExceptionDataResponse{}
	_body, _err := client.DescribeEsExceptionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEsExecuteDataWithOptions(request *DescribeEsExecuteDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEsExecuteDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEsExecuteDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEsExecuteData"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEsExecuteData(request *DescribeEsExecuteDataRequest) (_result *DescribeEsExecuteDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEsExecuteDataResponse{}
	_body, _err := client.DescribeEsExecuteDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFCTriggerWithOptions(request *DescribeFCTriggerRequest, runtime *util.RuntimeOptions) (_result *DescribeFCTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeFCTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFCTrigger"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFCTrigger(request *DescribeFCTriggerRequest) (_result *DescribeFCTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFCTriggerResponse{}
	_body, _err := client.DescribeFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIllegalUrlExportTaskWithOptions(request *DescribeIllegalUrlExportTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeIllegalUrlExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIllegalUrlExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIllegalUrlExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIllegalUrlExportTask(request *DescribeIllegalUrlExportTaskRequest) (_result *DescribeIllegalUrlExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIllegalUrlExportTaskResponse{}
	_body, _err := client.DescribeIllegalUrlExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpInfoWithOptions(request *DescribeIpInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeIpInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpInfo"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpInfo(request *DescribeIpInfoRequest) (_result *DescribeIpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpInfoResponse{}
	_body, _err := client.DescribeIpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeL2VipsByDomainWithOptions(request *DescribeL2VipsByDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeL2VipsByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeL2VipsByDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeL2VipsByDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeL2VipsByDomain(request *DescribeL2VipsByDomainRequest) (_result *DescribeL2VipsByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeL2VipsByDomainResponse{}
	_body, _err := client.DescribeL2VipsByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRangeDataByLocateAndIspServiceWithOptions(request *DescribeRangeDataByLocateAndIspServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeRangeDataByLocateAndIspServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRangeDataByLocateAndIspServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRangeDataByLocateAndIspService"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRangeDataByLocateAndIspService(request *DescribeRangeDataByLocateAndIspServiceRequest) (_result *DescribeRangeDataByLocateAndIspServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRangeDataByLocateAndIspServiceResponse{}
	_body, _err := client.DescribeRangeDataByLocateAndIspServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRealtimeDeliveryAccWithOptions(request *DescribeRealtimeDeliveryAccRequest, runtime *util.RuntimeOptions) (_result *DescribeRealtimeDeliveryAccResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRealtimeDeliveryAccResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRealtimeDeliveryAcc"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRealtimeDeliveryAcc(request *DescribeRealtimeDeliveryAccRequest) (_result *DescribeRealtimeDeliveryAccResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRealtimeDeliveryAccResponse{}
	_body, _err := client.DescribeRealtimeDeliveryAccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRefreshQuotaWithOptions(request *DescribeRefreshQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeRefreshQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRefreshQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRefreshQuota"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRefreshQuota(request *DescribeRefreshQuotaRequest) (_result *DescribeRefreshQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRefreshQuotaResponse{}
	_body, _err := client.DescribeRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRefreshTaskByIdWithOptions(request *DescribeRefreshTaskByIdRequest, runtime *util.RuntimeOptions) (_result *DescribeRefreshTaskByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRefreshTaskByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRefreshTaskById"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRefreshTaskById(request *DescribeRefreshTaskByIdRequest) (_result *DescribeRefreshTaskByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRefreshTaskByIdResponse{}
	_body, _err := client.DescribeRefreshTaskByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRefreshTasksWithOptions(request *DescribeRefreshTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeRefreshTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRefreshTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRefreshTasks"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRefreshTasks(request *DescribeRefreshTasksRequest) (_result *DescribeRefreshTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRefreshTasksResponse{}
	_body, _err := client.DescribeRefreshTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStagingIpWithOptions(request *DescribeStagingIpRequest, runtime *util.RuntimeOptions) (_result *DescribeStagingIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStagingIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStagingIp"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStagingIp(request *DescribeStagingIpRequest) (_result *DescribeStagingIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStagingIpResponse{}
	_body, _err := client.DescribeStagingIpWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeTagResources"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTopDomainsByFlowWithOptions(request *DescribeTopDomainsByFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeTopDomainsByFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTopDomainsByFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTopDomainsByFlow"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopDomainsByFlow(request *DescribeTopDomainsByFlowRequest) (_result *DescribeTopDomainsByFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopDomainsByFlowResponse{}
	_body, _err := client.DescribeTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserCertificateExpireCountWithOptions(request *DescribeUserCertificateExpireCountRequest, runtime *util.RuntimeOptions) (_result *DescribeUserCertificateExpireCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserCertificateExpireCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserCertificateExpireCount"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserCertificateExpireCount(request *DescribeUserCertificateExpireCountRequest) (_result *DescribeUserCertificateExpireCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserCertificateExpireCountResponse{}
	_body, _err := client.DescribeUserCertificateExpireCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserConfigsWithOptions(request *DescribeUserConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserConfigs"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserConfigs(request *DescribeUserConfigsRequest) (_result *DescribeUserConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserConfigsResponse{}
	_body, _err := client.DescribeUserConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserDomainsWithOptions(request *DescribeUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserDomains"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserDomains(request *DescribeUserDomainsRequest) (_result *DescribeUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserDomainsResponse{}
	_body, _err := client.DescribeUserDomainsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserTags"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeUserUsageDataExportTaskWithOptions(request *DescribeUserUsageDataExportTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeUserUsageDataExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserUsageDataExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserUsageDataExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserUsageDataExportTask(request *DescribeUserUsageDataExportTaskRequest) (_result *DescribeUserUsageDataExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserUsageDataExportTaskResponse{}
	_body, _err := client.DescribeUserUsageDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserUsageDetailDataExportTaskWithOptions(request *DescribeUserUsageDetailDataExportTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeUserUsageDetailDataExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserUsageDetailDataExportTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserUsageDetailDataExportTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserUsageDetailDataExportTask(request *DescribeUserUsageDetailDataExportTaskRequest) (_result *DescribeUserUsageDetailDataExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserUsageDetailDataExportTaskResponse{}
	_body, _err := client.DescribeUserUsageDetailDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserVipsByDomainWithOptions(request *DescribeUserVipsByDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeUserVipsByDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeUserVipsByDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserVipsByDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserVipsByDomain(request *DescribeUserVipsByDomainRequest) (_result *DescribeUserVipsByDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserVipsByDomainResponse{}
	_body, _err := client.DescribeUserVipsByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVerifyContentWithOptions(request *DescribeVerifyContentRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifyContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVerifyContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVerifyContent"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVerifyContent(request *DescribeVerifyContentRequest) (_result *DescribeVerifyContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifyContentResponse{}
	_body, _err := client.DescribeVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableRealtimeLogDeliveryWithOptions(request *DisableRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DisableRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DisableRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableRealtimeLogDelivery"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableRealtimeLogDelivery(request *DisableRealtimeLogDeliveryRequest) (_result *DisableRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableRealtimeLogDeliveryResponse{}
	_body, _err := client.DisableRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRealtimeLogDeliveryWithOptions(request *EnableRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *EnableRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &EnableRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableRealtimeLogDelivery"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRealtimeLogDelivery(request *EnableRealtimeLogDeliveryRequest) (_result *EnableRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRealtimeLogDeliveryResponse{}
	_body, _err := client.EnableRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDomainsByLogConfigIdWithOptions(request *ListDomainsByLogConfigIdRequest, runtime *util.RuntimeOptions) (_result *ListDomainsByLogConfigIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDomainsByLogConfigIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDomainsByLogConfigId"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDomainsByLogConfigId(request *ListDomainsByLogConfigIdRequest) (_result *ListDomainsByLogConfigIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDomainsByLogConfigIdResponse{}
	_body, _err := client.ListDomainsByLogConfigIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFCTriggerWithOptions(request *ListFCTriggerRequest, runtime *util.RuntimeOptions) (_result *ListFCTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListFCTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFCTrigger"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFCTrigger(request *ListFCTriggerRequest) (_result *ListFCTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFCTriggerResponse{}
	_body, _err := client.ListFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRealtimeLogDeliveryDomainsWithOptions(request *ListRealtimeLogDeliveryDomainsRequest, runtime *util.RuntimeOptions) (_result *ListRealtimeLogDeliveryDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRealtimeLogDeliveryDomains"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRealtimeLogDeliveryDomains(request *ListRealtimeLogDeliveryDomainsRequest) (_result *ListRealtimeLogDeliveryDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.ListRealtimeLogDeliveryDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRealtimeLogDeliveryInfosWithOptions(request *ListRealtimeLogDeliveryInfosRequest, runtime *util.RuntimeOptions) (_result *ListRealtimeLogDeliveryInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRealtimeLogDeliveryInfos"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRealtimeLogDeliveryInfos(request *ListRealtimeLogDeliveryInfosRequest) (_result *ListRealtimeLogDeliveryInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.ListRealtimeLogDeliveryInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserCustomLogConfigWithOptions(request *ListUserCustomLogConfigRequest, runtime *util.RuntimeOptions) (_result *ListUserCustomLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListUserCustomLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserCustomLogConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserCustomLogConfig(request *ListUserCustomLogConfigRequest) (_result *ListUserCustomLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserCustomLogConfigResponse{}
	_body, _err := client.ListUserCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCdnDomainWithOptions(request *ModifyCdnDomainRequest, runtime *util.RuntimeOptions) (_result *ModifyCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCdnDomain(request *ModifyCdnDomainRequest) (_result *ModifyCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCdnDomainResponse{}
	_body, _err := client.ModifyCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCdnDomainSchdmByPropertyWithOptions(request *ModifyCdnDomainSchdmByPropertyRequest, runtime *util.RuntimeOptions) (_result *ModifyCdnDomainSchdmByPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCdnDomainSchdmByProperty"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCdnDomainSchdmByProperty(request *ModifyCdnDomainSchdmByPropertyRequest) (_result *ModifyCdnDomainSchdmByPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.ModifyCdnDomainSchdmByPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainCustomLogConfigWithOptions(request *ModifyDomainCustomLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainCustomLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyDomainCustomLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDomainCustomLogConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainCustomLogConfig(request *ModifyDomainCustomLogConfigRequest) (_result *ModifyDomainCustomLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainCustomLogConfigResponse{}
	_body, _err := client.ModifyDomainCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRealtimeLogDeliveryWithOptions(request *ModifyRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *ModifyRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRealtimeLogDelivery"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRealtimeLogDelivery(request *ModifyRealtimeLogDeliveryRequest) (_result *ModifyRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRealtimeLogDeliveryResponse{}
	_body, _err := client.ModifyRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserCustomLogConfigWithOptions(request *ModifyUserCustomLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyUserCustomLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyUserCustomLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserCustomLogConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserCustomLogConfig(request *ModifyUserCustomLogConfigRequest) (_result *ModifyUserCustomLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserCustomLogConfigResponse{}
	_body, _err := client.ModifyUserCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenCdnServiceWithOptions(request *OpenCdnServiceRequest, runtime *util.RuntimeOptions) (_result *OpenCdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenCdnServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenCdnService"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenCdnService(request *OpenCdnServiceRequest) (_result *OpenCdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCdnServiceResponse{}
	_body, _err := client.OpenCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishStagingConfigToProductionWithOptions(request *PublishStagingConfigToProductionRequest, runtime *util.RuntimeOptions) (_result *PublishStagingConfigToProductionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PublishStagingConfigToProductionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PublishStagingConfigToProduction"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishStagingConfigToProduction(request *PublishStagingConfigToProductionRequest) (_result *PublishStagingConfigToProductionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PublishStagingConfigToProductionResponse{}
	_body, _err := client.PublishStagingConfigToProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushObjectCacheWithOptions(request *PushObjectCacheRequest, runtime *util.RuntimeOptions) (_result *PushObjectCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PushObjectCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PushObjectCache"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushObjectCache(request *PushObjectCacheRequest) (_result *PushObjectCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushObjectCacheResponse{}
	_body, _err := client.PushObjectCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshObjectCachesWithOptions(request *RefreshObjectCachesRequest, runtime *util.RuntimeOptions) (_result *RefreshObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshObjectCachesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshObjectCaches"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshObjectCaches(request *RefreshObjectCachesRequest) (_result *RefreshObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshObjectCachesResponse{}
	_body, _err := client.RefreshObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackStagingConfigWithOptions(request *RollbackStagingConfigRequest, runtime *util.RuntimeOptions) (_result *RollbackStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollbackStagingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollbackStagingConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackStagingConfig(request *RollbackStagingConfigRequest) (_result *RollbackStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackStagingConfigResponse{}
	_body, _err := client.RollbackStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCcConfigWithOptions(request *SetCcConfigRequest, runtime *util.RuntimeOptions) (_result *SetCcConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetCcConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetCcConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCcConfig(request *SetCcConfigRequest) (_result *SetCcConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCcConfigResponse{}
	_body, _err := client.SetCcConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCdnDomainCSRCertificateWithOptions(request *SetCdnDomainCSRCertificateRequest, runtime *util.RuntimeOptions) (_result *SetCdnDomainCSRCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetCdnDomainCSRCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetCdnDomainCSRCertificate"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCdnDomainCSRCertificate(request *SetCdnDomainCSRCertificateRequest) (_result *SetCdnDomainCSRCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCdnDomainCSRCertificateResponse{}
	_body, _err := client.SetCdnDomainCSRCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCdnDomainStagingConfigWithOptions(request *SetCdnDomainStagingConfigRequest, runtime *util.RuntimeOptions) (_result *SetCdnDomainStagingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetCdnDomainStagingConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetCdnDomainStagingConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCdnDomainStagingConfig(request *SetCdnDomainStagingConfigRequest) (_result *SetCdnDomainStagingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCdnDomainStagingConfigResponse{}
	_body, _err := client.SetCdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetConfigOfVersionWithOptions(request *SetConfigOfVersionRequest, runtime *util.RuntimeOptions) (_result *SetConfigOfVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetConfigOfVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetConfigOfVersion"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetConfigOfVersion(request *SetConfigOfVersionRequest) (_result *SetConfigOfVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetConfigOfVersionResponse{}
	_body, _err := client.SetConfigOfVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainGreenManagerConfigWithOptions(request *SetDomainGreenManagerConfigRequest, runtime *util.RuntimeOptions) (_result *SetDomainGreenManagerConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDomainGreenManagerConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDomainGreenManagerConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainGreenManagerConfig(request *SetDomainGreenManagerConfigRequest) (_result *SetDomainGreenManagerConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainGreenManagerConfigResponse{}
	_body, _err := client.SetDomainGreenManagerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainServerCertificateWithOptions(request *SetDomainServerCertificateRequest, runtime *util.RuntimeOptions) (_result *SetDomainServerCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDomainServerCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDomainServerCertificate"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainServerCertificate(request *SetDomainServerCertificateRequest) (_result *SetDomainServerCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainServerCertificateResponse{}
	_body, _err := client.SetDomainServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetErrorPageConfigWithOptions(request *SetErrorPageConfigRequest, runtime *util.RuntimeOptions) (_result *SetErrorPageConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetErrorPageConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetErrorPageConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetErrorPageConfig(request *SetErrorPageConfigRequest) (_result *SetErrorPageConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetErrorPageConfigResponse{}
	_body, _err := client.SetErrorPageConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetFileCacheExpiredConfigWithOptions(request *SetFileCacheExpiredConfigRequest, runtime *util.RuntimeOptions) (_result *SetFileCacheExpiredConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetFileCacheExpiredConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetFileCacheExpiredConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetFileCacheExpiredConfig(request *SetFileCacheExpiredConfigRequest) (_result *SetFileCacheExpiredConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFileCacheExpiredConfigResponse{}
	_body, _err := client.SetFileCacheExpiredConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetForceRedirectConfigWithOptions(request *SetForceRedirectConfigRequest, runtime *util.RuntimeOptions) (_result *SetForceRedirectConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetForceRedirectConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetForceRedirectConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetForceRedirectConfig(request *SetForceRedirectConfigRequest) (_result *SetForceRedirectConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetForceRedirectConfigResponse{}
	_body, _err := client.SetForceRedirectConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetForwardSchemeConfigWithOptions(request *SetForwardSchemeConfigRequest, runtime *util.RuntimeOptions) (_result *SetForwardSchemeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetForwardSchemeConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetForwardSchemeConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetForwardSchemeConfig(request *SetForwardSchemeConfigRequest) (_result *SetForwardSchemeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetForwardSchemeConfigResponse{}
	_body, _err := client.SetForwardSchemeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetHttpErrorPageConfigWithOptions(request *SetHttpErrorPageConfigRequest, runtime *util.RuntimeOptions) (_result *SetHttpErrorPageConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetHttpErrorPageConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetHttpErrorPageConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetHttpErrorPageConfig(request *SetHttpErrorPageConfigRequest) (_result *SetHttpErrorPageConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetHttpErrorPageConfigResponse{}
	_body, _err := client.SetHttpErrorPageConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetHttpHeaderConfigWithOptions(request *SetHttpHeaderConfigRequest, runtime *util.RuntimeOptions) (_result *SetHttpHeaderConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetHttpHeaderConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetHttpHeaderConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetHttpHeaderConfig(request *SetHttpHeaderConfigRequest) (_result *SetHttpHeaderConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetHttpHeaderConfigResponse{}
	_body, _err := client.SetHttpHeaderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetHttpsOptionConfigWithOptions(request *SetHttpsOptionConfigRequest, runtime *util.RuntimeOptions) (_result *SetHttpsOptionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetHttpsOptionConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetHttpsOptionConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetHttpsOptionConfig(request *SetHttpsOptionConfigRequest) (_result *SetHttpsOptionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetHttpsOptionConfigResponse{}
	_body, _err := client.SetHttpsOptionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIgnoreQueryStringConfigWithOptions(request *SetIgnoreQueryStringConfigRequest, runtime *util.RuntimeOptions) (_result *SetIgnoreQueryStringConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetIgnoreQueryStringConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetIgnoreQueryStringConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIgnoreQueryStringConfig(request *SetIgnoreQueryStringConfigRequest) (_result *SetIgnoreQueryStringConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIgnoreQueryStringConfigResponse{}
	_body, _err := client.SetIgnoreQueryStringConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIpAllowListConfigWithOptions(request *SetIpAllowListConfigRequest, runtime *util.RuntimeOptions) (_result *SetIpAllowListConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetIpAllowListConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetIpAllowListConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIpAllowListConfig(request *SetIpAllowListConfigRequest) (_result *SetIpAllowListConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIpAllowListConfigResponse{}
	_body, _err := client.SetIpAllowListConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIpBlackListConfigWithOptions(request *SetIpBlackListConfigRequest, runtime *util.RuntimeOptions) (_result *SetIpBlackListConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetIpBlackListConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetIpBlackListConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIpBlackListConfig(request *SetIpBlackListConfigRequest) (_result *SetIpBlackListConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIpBlackListConfigResponse{}
	_body, _err := client.SetIpBlackListConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetOptimizeConfigWithOptions(request *SetOptimizeConfigRequest, runtime *util.RuntimeOptions) (_result *SetOptimizeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetOptimizeConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetOptimizeConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetOptimizeConfig(request *SetOptimizeConfigRequest) (_result *SetOptimizeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetOptimizeConfigResponse{}
	_body, _err := client.SetOptimizeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetPageCompressConfigWithOptions(request *SetPageCompressConfigRequest, runtime *util.RuntimeOptions) (_result *SetPageCompressConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetPageCompressConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetPageCompressConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPageCompressConfig(request *SetPageCompressConfigRequest) (_result *SetPageCompressConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPageCompressConfigResponse{}
	_body, _err := client.SetPageCompressConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRangeConfigWithOptions(request *SetRangeConfigRequest, runtime *util.RuntimeOptions) (_result *SetRangeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetRangeConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetRangeConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRangeConfig(request *SetRangeConfigRequest) (_result *SetRangeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRangeConfigResponse{}
	_body, _err := client.SetRangeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRefererConfigWithOptions(request *SetRefererConfigRequest, runtime *util.RuntimeOptions) (_result *SetRefererConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetRefererConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetRefererConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRefererConfig(request *SetRefererConfigRequest) (_result *SetRefererConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRefererConfigResponse{}
	_body, _err := client.SetRefererConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRemoveQueryStringConfigWithOptions(request *SetRemoveQueryStringConfigRequest, runtime *util.RuntimeOptions) (_result *SetRemoveQueryStringConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetRemoveQueryStringConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetRemoveQueryStringConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRemoveQueryStringConfig(request *SetRemoveQueryStringConfigRequest) (_result *SetRemoveQueryStringConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRemoveQueryStringConfigResponse{}
	_body, _err := client.SetRemoveQueryStringConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetReqAuthConfigWithOptions(request *SetReqAuthConfigRequest, runtime *util.RuntimeOptions) (_result *SetReqAuthConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetReqAuthConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetReqAuthConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetReqAuthConfig(request *SetReqAuthConfigRequest) (_result *SetReqAuthConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetReqAuthConfigResponse{}
	_body, _err := client.SetReqAuthConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetReqHeaderConfigWithOptions(request *SetReqHeaderConfigRequest, runtime *util.RuntimeOptions) (_result *SetReqHeaderConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetReqHeaderConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetReqHeaderConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetReqHeaderConfig(request *SetReqHeaderConfigRequest) (_result *SetReqHeaderConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetReqHeaderConfigResponse{}
	_body, _err := client.SetReqHeaderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSourceHostConfigWithOptions(request *SetSourceHostConfigRequest, runtime *util.RuntimeOptions) (_result *SetSourceHostConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetSourceHostConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetSourceHostConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSourceHostConfig(request *SetSourceHostConfigRequest) (_result *SetSourceHostConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSourceHostConfigResponse{}
	_body, _err := client.SetSourceHostConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetWaitingRoomConfigWithOptions(request *SetWaitingRoomConfigRequest, runtime *util.RuntimeOptions) (_result *SetWaitingRoomConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetWaitingRoomConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetWaitingRoomConfig"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetWaitingRoomConfig(request *SetWaitingRoomConfigRequest) (_result *SetWaitingRoomConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetWaitingRoomConfigResponse{}
	_body, _err := client.SetWaitingRoomConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCdnDomainWithOptions(request *StartCdnDomainRequest, runtime *util.RuntimeOptions) (_result *StartCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCdnDomain(request *StartCdnDomainRequest) (_result *StartCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCdnDomainResponse{}
	_body, _err := client.StartCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCdnDomainWithOptions(request *StopCdnDomainRequest, runtime *util.RuntimeOptions) (_result *StopCdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopCdnDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopCdnDomain"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCdnDomain(request *StopCdnDomainRequest) (_result *StopCdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCdnDomainResponse{}
	_body, _err := client.StopCdnDomainWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateCdnDeliverTaskWithOptions(tmpReq *UpdateCdnDeliverTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateCdnDeliverTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCdnDeliverTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Deliver)) {
		request.DeliverShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Deliver, tea.String("Deliver"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Schedule)) {
		request.ScheduleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Schedule, tea.String("Schedule"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCdnDeliverTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCdnDeliverTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCdnDeliverTask(request *UpdateCdnDeliverTaskRequest) (_result *UpdateCdnDeliverTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCdnDeliverTaskResponse{}
	_body, _err := client.UpdateCdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCdnSubTaskWithOptions(request *UpdateCdnSubTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateCdnSubTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCdnSubTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCdnSubTask"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCdnSubTask(request *UpdateCdnSubTaskRequest) (_result *UpdateCdnSubTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCdnSubTaskResponse{}
	_body, _err := client.UpdateCdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFCTriggerWithOptions(request *UpdateFCTriggerRequest, runtime *util.RuntimeOptions) (_result *UpdateFCTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFCTriggerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFCTrigger"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFCTrigger(request *UpdateFCTriggerRequest) (_result *UpdateFCTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFCTriggerResponse{}
	_body, _err := client.UpdateFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyDomainOwnerWithOptions(request *VerifyDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyDomainOwner"), tea.String("2018-05-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyDomainOwner(request *VerifyDomainOwnerRequest) (_result *VerifyDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.VerifyDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
