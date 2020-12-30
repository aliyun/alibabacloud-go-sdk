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

type CleanFlexAccFwdRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s CleanFlexAccFwdRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CleanFlexAccFwdRulesRequest) GoString() string {
	return s.String()
}

func (s *CleanFlexAccFwdRulesRequest) SetSourceIp(v string) *CleanFlexAccFwdRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *CleanFlexAccFwdRulesRequest) SetEsnBizId(v int64) *CleanFlexAccFwdRulesRequest {
	s.EsnBizId = &v
	return s
}

type CleanFlexAccFwdRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CleanFlexAccFwdRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CleanFlexAccFwdRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CleanFlexAccFwdRulesResponseBody) SetRequestId(v string) *CleanFlexAccFwdRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CleanFlexAccFwdRulesResponseBody) SetPromptInfo(v map[string]interface{}) *CleanFlexAccFwdRulesResponseBody {
	s.PromptInfo = v
	return s
}

type CleanFlexAccFwdRulesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CleanFlexAccFwdRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CleanFlexAccFwdRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CleanFlexAccFwdRulesResponse) GoString() string {
	return s.String()
}

func (s *CleanFlexAccFwdRulesResponse) SetHeaders(v map[string]*string) *CleanFlexAccFwdRulesResponse {
	s.Headers = v
	return s
}

func (s *CleanFlexAccFwdRulesResponse) SetBody(v *CleanFlexAccFwdRulesResponseBody) *CleanFlexAccFwdRulesResponse {
	s.Body = v
	return s
}

type CleanFlexFwdRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s CleanFlexFwdRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CleanFlexFwdRulesRequest) GoString() string {
	return s.String()
}

func (s *CleanFlexFwdRulesRequest) SetSourceIp(v string) *CleanFlexFwdRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *CleanFlexFwdRulesRequest) SetLang(v string) *CleanFlexFwdRulesRequest {
	s.Lang = &v
	return s
}

func (s *CleanFlexFwdRulesRequest) SetEsnBizId(v int64) *CleanFlexFwdRulesRequest {
	s.EsnBizId = &v
	return s
}

type CleanFlexFwdRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CleanFlexFwdRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CleanFlexFwdRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CleanFlexFwdRulesResponseBody) SetRequestId(v string) *CleanFlexFwdRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CleanFlexFwdRulesResponseBody) SetPromptInfo(v map[string]interface{}) *CleanFlexFwdRulesResponseBody {
	s.PromptInfo = v
	return s
}

type CleanFlexFwdRulesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CleanFlexFwdRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CleanFlexFwdRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CleanFlexFwdRulesResponse) GoString() string {
	return s.String()
}

func (s *CleanFlexFwdRulesResponse) SetHeaders(v map[string]*string) *CleanFlexFwdRulesResponse {
	s.Headers = v
	return s
}

func (s *CleanFlexFwdRulesResponse) SetBody(v *CleanFlexFwdRulesResponseBody) *CleanFlexFwdRulesResponse {
	s.Body = v
	return s
}

type ClearCcRouteRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s ClearCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *ClearCcRouteRulesRequest) SetSourceIp(v string) *ClearCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *ClearCcRouteRulesRequest) SetLang(v string) *ClearCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *ClearCcRouteRulesRequest) SetBizId(v int64) *ClearCcRouteRulesRequest {
	s.BizId = &v
	return s
}

type ClearCcRouteRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s ClearCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ClearCcRouteRulesResponseBody) SetRequestId(v string) *ClearCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *ClearCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type ClearCcRouteRulesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClearCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *ClearCcRouteRulesResponse) SetHeaders(v map[string]*string) *ClearCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *ClearCcRouteRulesResponse) SetBody(v *ClearCcRouteRulesResponseBody) *ClearCcRouteRulesResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetSourceIp(v string) *CreateAppRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAppRequest) SetLang(v string) *CreateAppRequest {
	s.Lang = &v
	return s
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

type CreateAppResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppId      *int64                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetAppId(v int64) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetPromptInfo(v map[string]interface{}) *CreateAppResponseBody {
	s.PromptInfo = v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateAppKeyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateAppKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAppKeyRequest) SetSourceIp(v string) *CreateAppKeyRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAppKeyRequest) SetLang(v string) *CreateAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *CreateAppKeyRequest) SetAppId(v int64) *CreateAppKeyRequest {
	s.AppId = &v
	return s
}

type CreateAppKeyResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateAppKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppKeyResponseBody) SetRequestId(v string) *CreateAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppKeyResponseBody) SetPromptInfo(v map[string]interface{}) *CreateAppKeyResponseBody {
	s.PromptInfo = v
	return s
}

type CreateAppKeyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAppKeyResponse) SetHeaders(v map[string]*string) *CreateAppKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAppKeyResponse) SetBody(v *CreateAppKeyResponseBody) *CreateAppKeyResponse {
	s.Body = v
	return s
}

type CreateBizRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName  *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizType  *int32  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	UseCc    *int32  `json:"UseCc,omitempty" xml:"UseCc,omitempty"`
	CcQps    *int32  `json:"CcQps,omitempty" xml:"CcQps,omitempty"`
	L4Rules  *string `json:"L4Rules,omitempty" xml:"L4Rules,omitempty"`
	Groups   *string `json:"Groups,omitempty" xml:"Groups,omitempty"`
}

func (s CreateBizRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBizRequest) GoString() string {
	return s.String()
}

func (s *CreateBizRequest) SetSourceIp(v string) *CreateBizRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateBizRequest) SetLang(v string) *CreateBizRequest {
	s.Lang = &v
	return s
}

func (s *CreateBizRequest) SetAppId(v int64) *CreateBizRequest {
	s.AppId = &v
	return s
}

func (s *CreateBizRequest) SetAppName(v string) *CreateBizRequest {
	s.AppName = &v
	return s
}

func (s *CreateBizRequest) SetBizName(v string) *CreateBizRequest {
	s.BizName = &v
	return s
}

func (s *CreateBizRequest) SetBizType(v int32) *CreateBizRequest {
	s.BizType = &v
	return s
}

func (s *CreateBizRequest) SetUseCc(v int32) *CreateBizRequest {
	s.UseCc = &v
	return s
}

func (s *CreateBizRequest) SetCcQps(v int32) *CreateBizRequest {
	s.CcQps = &v
	return s
}

func (s *CreateBizRequest) SetL4Rules(v string) *CreateBizRequest {
	s.L4Rules = &v
	return s
}

func (s *CreateBizRequest) SetGroups(v string) *CreateBizRequest {
	s.Groups = &v
	return s
}

type CreateBizResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateBizResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBizResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizResponseBody) SetRequestId(v string) *CreateBizResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBizResponseBody) SetPromptInfo(v map[string]interface{}) *CreateBizResponseBody {
	s.PromptInfo = v
	return s
}

type CreateBizResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBizResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBizResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBizResponse) GoString() string {
	return s.String()
}

func (s *CreateBizResponse) SetHeaders(v map[string]*string) *CreateBizResponse {
	s.Headers = v
	return s
}

func (s *CreateBizResponse) SetBody(v *CreateBizResponseBody) *CreateBizResponse {
	s.Body = v
	return s
}

type CreateCcRouteRulesRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RouteList *string `json:"RouteList,omitempty" xml:"RouteList,omitempty"`
}

func (s CreateCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateCcRouteRulesRequest) SetSourceIp(v string) *CreateCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateCcRouteRulesRequest) SetLang(v string) *CreateCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *CreateCcRouteRulesRequest) SetBizId(v int64) *CreateCcRouteRulesRequest {
	s.BizId = &v
	return s
}

func (s *CreateCcRouteRulesRequest) SetRouteList(v string) *CreateCcRouteRulesRequest {
	s.RouteList = &v
	return s
}

type CreateCcRouteRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCcRouteRulesResponseBody) SetRequestId(v string) *CreateCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *CreateCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type CreateCcRouteRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateCcRouteRulesResponse) SetHeaders(v map[string]*string) *CreateCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateCcRouteRulesResponse) SetBody(v *CreateCcRouteRulesResponseBody) *CreateCcRouteRulesResponse {
	s.Body = v
	return s
}

type CreateFlexAccFwdRuleRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BizId        *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	IpList       *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	ProtocolEx   *string `json:"ProtocolEx,omitempty" xml:"ProtocolEx,omitempty"`
	DomainList   *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	AccType      *int32  `json:"AccType,omitempty" xml:"AccType,omitempty"`
}

func (s CreateFlexAccFwdRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexAccFwdRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFlexAccFwdRuleRequest) SetSourceIp(v string) *CreateFlexAccFwdRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetBizId(v int64) *CreateFlexAccFwdRuleRequest {
	s.BizId = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetIdentity(v string) *CreateFlexAccFwdRuleRequest {
	s.Identity = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetIpList(v string) *CreateFlexAccFwdRuleRequest {
	s.IpList = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetProtocolEx(v string) *CreateFlexAccFwdRuleRequest {
	s.ProtocolEx = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetDomainList(v string) *CreateFlexAccFwdRuleRequest {
	s.DomainList = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetRemark(v string) *CreateFlexAccFwdRuleRequest {
	s.Remark = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetMasterIpList(v string) *CreateFlexAccFwdRuleRequest {
	s.MasterIpList = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetSlaveIpList(v string) *CreateFlexAccFwdRuleRequest {
	s.SlaveIpList = &v
	return s
}

func (s *CreateFlexAccFwdRuleRequest) SetAccType(v int32) *CreateFlexAccFwdRuleRequest {
	s.AccType = &v
	return s
}

type CreateFlexAccFwdRuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateFlexAccFwdRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexAccFwdRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlexAccFwdRuleResponseBody) SetRequestId(v string) *CreateFlexAccFwdRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlexAccFwdRuleResponseBody) SetPromptInfo(v map[string]interface{}) *CreateFlexAccFwdRuleResponseBody {
	s.PromptInfo = v
	return s
}

type CreateFlexAccFwdRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlexAccFwdRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlexAccFwdRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexAccFwdRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFlexAccFwdRuleResponse) SetHeaders(v map[string]*string) *CreateFlexAccFwdRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFlexAccFwdRuleResponse) SetBody(v *CreateFlexAccFwdRuleResponseBody) *CreateFlexAccFwdRuleResponse {
	s.Body = v
	return s
}

type CreateFlexAccFwdRuleListRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FwdRuleList *string `json:"FwdRuleList,omitempty" xml:"FwdRuleList,omitempty"`
}

func (s CreateFlexAccFwdRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexAccFwdRuleListRequest) GoString() string {
	return s.String()
}

func (s *CreateFlexAccFwdRuleListRequest) SetSourceIp(v string) *CreateFlexAccFwdRuleListRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateFlexAccFwdRuleListRequest) SetBizId(v int64) *CreateFlexAccFwdRuleListRequest {
	s.BizId = &v
	return s
}

func (s *CreateFlexAccFwdRuleListRequest) SetFwdRuleList(v string) *CreateFlexAccFwdRuleListRequest {
	s.FwdRuleList = &v
	return s
}

type CreateFlexAccFwdRuleListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateFlexAccFwdRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexAccFwdRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlexAccFwdRuleListResponseBody) SetRequestId(v string) *CreateFlexAccFwdRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlexAccFwdRuleListResponseBody) SetPromptInfo(v map[string]interface{}) *CreateFlexAccFwdRuleListResponseBody {
	s.PromptInfo = v
	return s
}

type CreateFlexAccFwdRuleListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlexAccFwdRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlexAccFwdRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexAccFwdRuleListResponse) GoString() string {
	return s.String()
}

func (s *CreateFlexAccFwdRuleListResponse) SetHeaders(v map[string]*string) *CreateFlexAccFwdRuleListResponse {
	s.Headers = v
	return s
}

func (s *CreateFlexAccFwdRuleListResponse) SetBody(v *CreateFlexAccFwdRuleListResponseBody) *CreateFlexAccFwdRuleListResponse {
	s.Body = v
	return s
}

type CreateFlexFwdRuleRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId        *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateFlexFwdRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexFwdRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFlexFwdRuleRequest) SetSourceIp(v string) *CreateFlexFwdRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateFlexFwdRuleRequest) SetLang(v string) *CreateFlexFwdRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateFlexFwdRuleRequest) SetBizId(v int64) *CreateFlexFwdRuleRequest {
	s.BizId = &v
	return s
}

func (s *CreateFlexFwdRuleRequest) SetIdentity(v string) *CreateFlexFwdRuleRequest {
	s.Identity = &v
	return s
}

func (s *CreateFlexFwdRuleRequest) SetMasterIpList(v string) *CreateFlexFwdRuleRequest {
	s.MasterIpList = &v
	return s
}

func (s *CreateFlexFwdRuleRequest) SetSlaveIpList(v string) *CreateFlexFwdRuleRequest {
	s.SlaveIpList = &v
	return s
}

func (s *CreateFlexFwdRuleRequest) SetRemark(v string) *CreateFlexFwdRuleRequest {
	s.Remark = &v
	return s
}

type CreateFlexFwdRuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateFlexFwdRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexFwdRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlexFwdRuleResponseBody) SetRequestId(v string) *CreateFlexFwdRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlexFwdRuleResponseBody) SetPromptInfo(v map[string]interface{}) *CreateFlexFwdRuleResponseBody {
	s.PromptInfo = v
	return s
}

type CreateFlexFwdRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlexFwdRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlexFwdRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlexFwdRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFlexFwdRuleResponse) SetHeaders(v map[string]*string) *CreateFlexFwdRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFlexFwdRuleResponse) SetBody(v *CreateFlexFwdRuleResponseBody) *CreateFlexFwdRuleResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupDesc *string `json:"GroupDesc,omitempty" xml:"GroupDesc,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetSourceIp(v string) *CreateGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateGroupRequest) SetLang(v string) *CreateGroupRequest {
	s.Lang = &v
	return s
}

func (s *CreateGroupRequest) SetBizId(v int64) *CreateGroupRequest {
	s.BizId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupId(v string) *CreateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupDesc(v string) *CreateGroupRequest {
	s.GroupDesc = &v
	return s
}

func (s *CreateGroupRequest) SetIpList(v string) *CreateGroupRequest {
	s.IpList = &v
	return s
}

type CreateGroupResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) SetPromptInfo(v map[string]interface{}) *CreateGroupResponseBody {
	s.PromptInfo = v
	return s
}

type CreateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateLayer4RuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FrontPort *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	BackPort  *int32  `json:"BackPort,omitempty" xml:"BackPort,omitempty"`
	Rservers  *string `json:"Rservers,omitempty" xml:"Rservers,omitempty"`
}

func (s CreateLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleRequest) SetSourceIp(v string) *CreateLayer4RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetLang(v string) *CreateLayer4RuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetAppId(v int64) *CreateLayer4RuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetBizId(v int64) *CreateLayer4RuleRequest {
	s.BizId = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetFrontPort(v int32) *CreateLayer4RuleRequest {
	s.FrontPort = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetBackPort(v int32) *CreateLayer4RuleRequest {
	s.BackPort = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetRservers(v string) *CreateLayer4RuleRequest {
	s.Rservers = &v
	return s
}

type CreateLayer4RuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateLayer4RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleResponseBody) SetRequestId(v string) *CreateLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLayer4RuleResponseBody) SetPromptInfo(v map[string]interface{}) *CreateLayer4RuleResponseBody {
	s.PromptInfo = v
	return s
}

type CreateLayer4RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleResponse) SetHeaders(v map[string]*string) *CreateLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *CreateLayer4RuleResponse) SetBody(v *CreateLayer4RuleResponseBody) *CreateLayer4RuleResponse {
	s.Body = v
	return s
}

type CreateLayer4RulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Rules    *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	Flush    *int32  `json:"Flush,omitempty" xml:"Flush,omitempty"`
}

func (s CreateLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RulesRequest) SetSourceIp(v string) *CreateLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateLayer4RulesRequest) SetLang(v string) *CreateLayer4RulesRequest {
	s.Lang = &v
	return s
}

func (s *CreateLayer4RulesRequest) SetBizId(v int64) *CreateLayer4RulesRequest {
	s.BizId = &v
	return s
}

func (s *CreateLayer4RulesRequest) SetRules(v string) *CreateLayer4RulesRequest {
	s.Rules = &v
	return s
}

func (s *CreateLayer4RulesRequest) SetFlush(v int32) *CreateLayer4RulesRequest {
	s.Flush = &v
	return s
}

type CreateLayer4RulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s CreateLayer4RulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayer4RulesResponseBody) SetRequestId(v string) *CreateLayer4RulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLayer4RulesResponseBody) SetPromptInfo(v map[string]interface{}) *CreateLayer4RulesResponseBody {
	s.PromptInfo = v
	return s
}

type CreateLayer4RulesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLayer4RulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer4RulesResponse) SetHeaders(v map[string]*string) *CreateLayer4RulesResponse {
	s.Headers = v
	return s
}

func (s *CreateLayer4RulesResponse) SetBody(v *CreateLayer4RulesResponseBody) *CreateLayer4RulesResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetSourceIp(v string) *DeleteAppRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAppRequest) SetLang(v string) *DeleteAppRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAppRequest) SetAppId(v int64) *DeleteAppRequest {
	s.AppId = &v
	return s
}

type DeleteAppResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteAppResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteAppKeyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppKeyRequest) SetSourceIp(v string) *DeleteAppKeyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAppKeyRequest) SetLang(v string) *DeleteAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAppKeyRequest) SetAppId(v int64) *DeleteAppKeyRequest {
	s.AppId = &v
	return s
}

type DeleteAppKeyResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteAppKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppKeyResponseBody) SetRequestId(v string) *DeleteAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppKeyResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteAppKeyResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteAppKeyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppKeyResponse) SetHeaders(v map[string]*string) *DeleteAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppKeyResponse) SetBody(v *DeleteAppKeyResponseBody) *DeleteAppKeyResponse {
	s.Body = v
	return s
}

type DeleteBizRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DeleteBizRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizRequest) SetSourceIp(v string) *DeleteBizRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteBizRequest) SetLang(v string) *DeleteBizRequest {
	s.Lang = &v
	return s
}

func (s *DeleteBizRequest) SetBizId(v int64) *DeleteBizRequest {
	s.BizId = &v
	return s
}

type DeleteBizResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteBizResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizResponseBody) SetRequestId(v string) *DeleteBizResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBizResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteBizResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteBizResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBizResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBizResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizResponse) SetHeaders(v map[string]*string) *DeleteBizResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizResponse) SetBody(v *DeleteBizResponseBody) *DeleteBizResponse {
	s.Body = v
	return s
}

type DeleteCcRouteRulesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RouteIdList *string `json:"RouteIdList,omitempty" xml:"RouteIdList,omitempty"`
}

func (s DeleteCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCcRouteRulesRequest) SetSourceIp(v string) *DeleteCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteCcRouteRulesRequest) SetLang(v string) *DeleteCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *DeleteCcRouteRulesRequest) SetBizId(v int64) *DeleteCcRouteRulesRequest {
	s.BizId = &v
	return s
}

func (s *DeleteCcRouteRulesRequest) SetRouteIdList(v string) *DeleteCcRouteRulesRequest {
	s.RouteIdList = &v
	return s
}

type DeleteCcRouteRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCcRouteRulesResponseBody) SetRequestId(v string) *DeleteCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteCcRouteRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCcRouteRulesResponse) SetHeaders(v map[string]*string) *DeleteCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCcRouteRulesResponse) SetBody(v *DeleteCcRouteRulesResponseBody) *DeleteCcRouteRulesResponse {
	s.Body = v
	return s
}

type DeleteFlexAccFwdRuleRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
}

func (s DeleteFlexAccFwdRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlexAccFwdRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlexAccFwdRuleRequest) SetSourceIp(v string) *DeleteFlexAccFwdRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteFlexAccFwdRuleRequest) SetEsnBizId(v int64) *DeleteFlexAccFwdRuleRequest {
	s.EsnBizId = &v
	return s
}

func (s *DeleteFlexAccFwdRuleRequest) SetIdentity(v string) *DeleteFlexAccFwdRuleRequest {
	s.Identity = &v
	return s
}

type DeleteFlexAccFwdRuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteFlexAccFwdRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlexAccFwdRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlexAccFwdRuleResponseBody) SetRequestId(v string) *DeleteFlexAccFwdRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlexAccFwdRuleResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteFlexAccFwdRuleResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteFlexAccFwdRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlexAccFwdRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlexAccFwdRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlexAccFwdRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlexAccFwdRuleResponse) SetHeaders(v map[string]*string) *DeleteFlexAccFwdRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlexAccFwdRuleResponse) SetBody(v *DeleteFlexAccFwdRuleResponseBody) *DeleteFlexAccFwdRuleResponse {
	s.Body = v
	return s
}

type DeleteFlexFwdRuleRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
}

func (s DeleteFlexFwdRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlexFwdRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlexFwdRuleRequest) SetSourceIp(v string) *DeleteFlexFwdRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteFlexFwdRuleRequest) SetLang(v string) *DeleteFlexFwdRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteFlexFwdRuleRequest) SetEsnBizId(v int64) *DeleteFlexFwdRuleRequest {
	s.EsnBizId = &v
	return s
}

func (s *DeleteFlexFwdRuleRequest) SetIdentity(v string) *DeleteFlexFwdRuleRequest {
	s.Identity = &v
	return s
}

type DeleteFlexFwdRuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteFlexFwdRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlexFwdRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlexFwdRuleResponseBody) SetRequestId(v string) *DeleteFlexFwdRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlexFwdRuleResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteFlexFwdRuleResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteFlexFwdRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlexFwdRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlexFwdRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlexFwdRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlexFwdRuleResponse) SetHeaders(v map[string]*string) *DeleteFlexFwdRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlexFwdRuleResponse) SetBody(v *DeleteFlexFwdRuleResponseBody) *DeleteFlexFwdRuleResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetSourceIp(v string) *DeleteGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteGroupRequest) SetLang(v string) *DeleteGroupRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGroupRequest) SetBizId(v int64) *DeleteGroupRequest {
	s.BizId = &v
	return s
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteGroupResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteGroupResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteLayer4RuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FrontPort *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
}

func (s DeleteLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleRequest) SetSourceIp(v string) *DeleteLayer4RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteLayer4RuleRequest) SetLang(v string) *DeleteLayer4RuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteLayer4RuleRequest) SetBizId(v int64) *DeleteLayer4RuleRequest {
	s.BizId = &v
	return s
}

func (s *DeleteLayer4RuleRequest) SetFrontPort(v int32) *DeleteLayer4RuleRequest {
	s.FrontPort = &v
	return s
}

type DeleteLayer4RuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DeleteLayer4RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponseBody) SetRequestId(v string) *DeleteLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLayer4RuleResponseBody) SetPromptInfo(v map[string]interface{}) *DeleteLayer4RuleResponseBody {
	s.PromptInfo = v
	return s
}

type DeleteLayer4RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponse) SetHeaders(v map[string]*string) *DeleteLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer4RuleResponse) SetBody(v *DeleteLayer4RuleResponseBody) *DeleteLayer4RuleResponse {
	s.Body = v
	return s
}

type DescribeAccResSummaryRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAccResSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccResSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccResSummaryRequest) SetSourceIp(v string) *DescribeAccResSummaryRequest {
	s.SourceIp = &v
	return s
}

type DescribeAccResSummaryResponseBody struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullNodesSummary *DescribeAccResSummaryResponseBodyFullNodesSummary `json:"FullNodesSummary,omitempty" xml:"FullNodesSummary,omitempty" type:"Struct"`
	PromptInfo       map[string]interface{}                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeAccResSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccResSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccResSummaryResponseBody) SetRequestId(v string) *DescribeAccResSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccResSummaryResponseBody) SetFullNodesSummary(v *DescribeAccResSummaryResponseBodyFullNodesSummary) *DescribeAccResSummaryResponseBody {
	s.FullNodesSummary = v
	return s
}

func (s *DescribeAccResSummaryResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeAccResSummaryResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeAccResSummaryResponseBodyFullNodesSummary struct {
	Nodes       []*DescribeAccResSummaryResponseBodyFullNodesSummaryNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	CleanCount  *int32                                                    `json:"CleanCount,omitempty" xml:"CleanCount,omitempty"`
	HoleCount   *int32                                                    `json:"HoleCount,omitempty" xml:"HoleCount,omitempty"`
	UsedCount   *int32                                                    `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	UnUsedCount *int32                                                    `json:"UnUsedCount,omitempty" xml:"UnUsedCount,omitempty"`
	TotalCount  *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccResSummaryResponseBodyFullNodesSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccResSummaryResponseBodyFullNodesSummary) GoString() string {
	return s.String()
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummary) SetNodes(v []*DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) *DescribeAccResSummaryResponseBodyFullNodesSummary {
	s.Nodes = v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummary) SetCleanCount(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummary {
	s.CleanCount = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummary) SetHoleCount(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummary {
	s.HoleCount = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummary) SetUsedCount(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummary {
	s.UsedCount = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummary) SetUnUsedCount(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummary {
	s.UnUsedCount = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummary) SetTotalCount(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummary {
	s.TotalCount = &v
	return s
}

type DescribeAccResSummaryResponseBodyFullNodesSummaryNodes struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) GoString() string {
	return s.String()
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetType(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.Type = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetBizId(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizId = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetIp(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.Ip = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetRegionId(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetAppName(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppName = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetRegionName(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionName = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetGroupId(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetAppId(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppId = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetInUse(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.InUse = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetBizName(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizName = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetDdosStatus(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.DdosStatus = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetIsDisabled(v bool) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.IsDisabled = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetThreshold(v int32) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.Threshold = &v
	return s
}

func (s *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes) SetEsnBizId(v string) *DescribeAccResSummaryResponseBodyFullNodesSummaryNodes {
	s.EsnBizId = &v
	return s
}

type DescribeAccResSummaryResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccResSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccResSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccResSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccResSummaryResponse) SetHeaders(v map[string]*string) *DescribeAccResSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccResSummaryResponse) SetBody(v *DescribeAccResSummaryResponseBody) *DescribeAccResSummaryResponse {
	s.Body = v
	return s
}

type DescribeAllLocalIpsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAllLocalIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllLocalIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllLocalIpsRequest) SetSourceIp(v string) *DescribeAllLocalIpsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAllLocalIpsRequest) SetLang(v string) *DescribeAllLocalIpsRequest {
	s.Lang = &v
	return s
}

type DescribeAllLocalIpsResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LocalIps   *DescribeAllLocalIpsResponseBodyLocalIps `json:"LocalIps,omitempty" xml:"LocalIps,omitempty" type:"Struct"`
	PromptInfo map[string]interface{}                   `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeAllLocalIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllLocalIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllLocalIpsResponseBody) SetRequestId(v string) *DescribeAllLocalIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllLocalIpsResponseBody) SetLocalIps(v *DescribeAllLocalIpsResponseBodyLocalIps) *DescribeAllLocalIpsResponseBody {
	s.LocalIps = v
	return s
}

func (s *DescribeAllLocalIpsResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeAllLocalIpsResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeAllLocalIpsResponseBodyLocalIps struct {
	Bgp []*string `json:"Bgp,omitempty" xml:"Bgp,omitempty" type:"Repeated"`
	Gf  []*string `json:"Gf,omitempty" xml:"Gf,omitempty" type:"Repeated"`
}

func (s DescribeAllLocalIpsResponseBodyLocalIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllLocalIpsResponseBodyLocalIps) GoString() string {
	return s.String()
}

func (s *DescribeAllLocalIpsResponseBodyLocalIps) SetBgp(v []*string) *DescribeAllLocalIpsResponseBodyLocalIps {
	s.Bgp = v
	return s
}

func (s *DescribeAllLocalIpsResponseBodyLocalIps) SetGf(v []*string) *DescribeAllLocalIpsResponseBodyLocalIps {
	s.Gf = v
	return s
}

type DescribeAllLocalIpsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllLocalIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllLocalIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllLocalIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllLocalIpsResponse) SetHeaders(v map[string]*string) *DescribeAllLocalIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllLocalIpsResponse) SetBody(v *DescribeAllLocalIpsResponseBody) *DescribeAllLocalIpsResponse {
	s.Body = v
	return s
}

type DescribeAppDailyDetailsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnAppId  *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeAppDailyDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailyDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppDailyDetailsRequest) SetSourceIp(v string) *DescribeAppDailyDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAppDailyDetailsRequest) SetLang(v string) *DescribeAppDailyDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppDailyDetailsRequest) SetEsnAppId(v string) *DescribeAppDailyDetailsRequest {
	s.EsnAppId = &v
	return s
}

func (s *DescribeAppDailyDetailsRequest) SetStartTime(v int64) *DescribeAppDailyDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAppDailyDetailsRequest) SetEndTime(v int64) *DescribeAppDailyDetailsRequest {
	s.EndTime = &v
	return s
}

type DescribeAppDailyDetailsResponseBody struct {
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppDailyDetails []*DescribeAppDailyDetailsResponseBodyAppDailyDetails `json:"AppDailyDetails,omitempty" xml:"AppDailyDetails,omitempty" type:"Repeated"`
}

func (s DescribeAppDailyDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailyDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppDailyDetailsResponseBody) SetRequestId(v string) *DescribeAppDailyDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBody) SetAppDailyDetails(v []*DescribeAppDailyDetailsResponseBodyAppDailyDetails) *DescribeAppDailyDetailsResponseBody {
	s.AppDailyDetails = v
	return s
}

type DescribeAppDailyDetailsResponseBodyAppDailyDetails struct {
	IosActiveCount     *int64 `json:"IosActiveCount,omitempty" xml:"IosActiveCount,omitempty"`
	Index              *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	AndroidCount       *int64 `json:"AndroidCount,omitempty" xml:"AndroidCount,omitempty"`
	IosCount           *int64 `json:"IosCount,omitempty" xml:"IosCount,omitempty"`
	NewCount           *int64 `json:"NewCount,omitempty" xml:"NewCount,omitempty"`
	IpActiveCount      *int64 `json:"IpActiveCount,omitempty" xml:"IpActiveCount,omitempty"`
	ActiveCount        *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	Count              *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	AndroidActiveCount *int64 `json:"AndroidActiveCount,omitempty" xml:"AndroidActiveCount,omitempty"`
}

func (s DescribeAppDailyDetailsResponseBodyAppDailyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailyDetailsResponseBodyAppDailyDetails) GoString() string {
	return s.String()
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetIosActiveCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.IosActiveCount = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetIndex(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.Index = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetAndroidCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.AndroidCount = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetIosCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.IosCount = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetNewCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.NewCount = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetIpActiveCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.IpActiveCount = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetActiveCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.ActiveCount = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.Count = &v
	return s
}

func (s *DescribeAppDailyDetailsResponseBodyAppDailyDetails) SetAndroidActiveCount(v int64) *DescribeAppDailyDetailsResponseBodyAppDailyDetails {
	s.AndroidActiveCount = &v
	return s
}

type DescribeAppDailyDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppDailyDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppDailyDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailyDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppDailyDetailsResponse) SetHeaders(v map[string]*string) *DescribeAppDailyDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppDailyDetailsResponse) SetBody(v *DescribeAppDailyDetailsResponseBody) *DescribeAppDailyDetailsResponse {
	s.Body = v
	return s
}

type DescribeAppDailySdkVersionsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnAppId  *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeAppDailySdkVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailySdkVersionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppDailySdkVersionsRequest) SetSourceIp(v string) *DescribeAppDailySdkVersionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAppDailySdkVersionsRequest) SetLang(v string) *DescribeAppDailySdkVersionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppDailySdkVersionsRequest) SetEsnAppId(v string) *DescribeAppDailySdkVersionsRequest {
	s.EsnAppId = &v
	return s
}

func (s *DescribeAppDailySdkVersionsRequest) SetStartTime(v int64) *DescribeAppDailySdkVersionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAppDailySdkVersionsRequest) SetEndTime(v int64) *DescribeAppDailySdkVersionsRequest {
	s.EndTime = &v
	return s
}

type DescribeAppDailySdkVersionsResponseBody struct {
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppDailySdkVersions []*DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions `json:"AppDailySdkVersions,omitempty" xml:"AppDailySdkVersions,omitempty" type:"Repeated"`
}

func (s DescribeAppDailySdkVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailySdkVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppDailySdkVersionsResponseBody) SetRequestId(v string) *DescribeAppDailySdkVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppDailySdkVersionsResponseBody) SetAppDailySdkVersions(v []*DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions) *DescribeAppDailySdkVersionsResponseBody {
	s.AppDailySdkVersions = v
	return s
}

type DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions struct {
	Index   *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Count   *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions) GoString() string {
	return s.String()
}

func (s *DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions) SetIndex(v int64) *DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions {
	s.Index = &v
	return s
}

func (s *DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions) SetVersion(v string) *DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions {
	s.Version = &v
	return s
}

func (s *DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions) SetCount(v int64) *DescribeAppDailySdkVersionsResponseBodyAppDailySdkVersions {
	s.Count = &v
	return s
}

type DescribeAppDailySdkVersionsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppDailySdkVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppDailySdkVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppDailySdkVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppDailySdkVersionsResponse) SetHeaders(v map[string]*string) *DescribeAppDailySdkVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppDailySdkVersionsResponse) SetBody(v *DescribeAppDailySdkVersionsResponseBody) *DescribeAppDailySdkVersionsResponse {
	s.Body = v
	return s
}

type DescribeAppListRequest struct {
	SourceIp  *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId     *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName   *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
}

func (s DescribeAppListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppListRequest) SetSourceIp(v string) *DescribeAppListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAppListRequest) SetLang(v string) *DescribeAppListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppListRequest) SetAppId(v int64) *DescribeAppListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppListRequest) SetAppName(v string) *DescribeAppListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAppListRequest) SetAppIdList(v []*string) *DescribeAppListRequest {
	s.AppIdList = v
	return s
}

type DescribeAppListResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	AppList    []*DescribeAppListResponseBodyAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
}

func (s DescribeAppListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppListResponseBody) SetRequestId(v string) *DescribeAppListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppListResponseBody) SetTotal(v int64) *DescribeAppListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAppListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeAppListResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeAppListResponseBody) SetAppList(v []*DescribeAppListResponseBodyAppList) *DescribeAppListResponseBody {
	s.AppList = v
	return s
}

type DescribeAppListResponseBodyAppList struct {
	OldAppKey    *string `json:"OldAppKey,omitempty" xml:"OldAppKey,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppKey       *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	HasOldAppKey *bool   `json:"HasOldAppKey,omitempty" xml:"HasOldAppKey,omitempty"`
	SdkVersion   *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	BizNum       *int32  `json:"BizNum,omitempty" xml:"BizNum,omitempty"`
	GroupNum     *int32  `json:"GroupNum,omitempty" xml:"GroupNum,omitempty"`
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
}

func (s DescribeAppListResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppListResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *DescribeAppListResponseBodyAppList) SetOldAppKey(v string) *DescribeAppListResponseBodyAppList {
	s.OldAppKey = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetAppName(v string) *DescribeAppListResponseBodyAppList {
	s.AppName = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetAppKey(v string) *DescribeAppListResponseBodyAppList {
	s.AppKey = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetAppId(v int64) *DescribeAppListResponseBodyAppList {
	s.AppId = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetHasOldAppKey(v bool) *DescribeAppListResponseBodyAppList {
	s.HasOldAppKey = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetSdkVersion(v string) *DescribeAppListResponseBodyAppList {
	s.SdkVersion = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetBizNum(v int32) *DescribeAppListResponseBodyAppList {
	s.BizNum = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetGroupNum(v int32) *DescribeAppListResponseBodyAppList {
	s.GroupNum = &v
	return s
}

func (s *DescribeAppListResponseBodyAppList) SetNodeNum(v int32) *DescribeAppListResponseBodyAppList {
	s.NodeNum = &v
	return s
}

type DescribeAppListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppListResponse) SetHeaders(v map[string]*string) *DescribeAppListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppListResponse) SetBody(v *DescribeAppListResponseBody) *DescribeAppListResponse {
	s.Body = v
	return s
}

type DescribeAppSimpleListRequest struct {
	SourceIp  *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
}

func (s DescribeAppSimpleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSimpleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSimpleListRequest) SetSourceIp(v string) *DescribeAppSimpleListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAppSimpleListRequest) SetLang(v string) *DescribeAppSimpleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppSimpleListRequest) SetAppIdList(v []*string) *DescribeAppSimpleListRequest {
	s.AppIdList = v
	return s
}

type DescribeAppSimpleListResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppSimpleList []*DescribeAppSimpleListResponseBodyAppSimpleList `json:"AppSimpleList,omitempty" xml:"AppSimpleList,omitempty" type:"Repeated"`
	Total         *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo    map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeAppSimpleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSimpleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppSimpleListResponseBody) SetRequestId(v string) *DescribeAppSimpleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppSimpleListResponseBody) SetAppSimpleList(v []*DescribeAppSimpleListResponseBodyAppSimpleList) *DescribeAppSimpleListResponseBody {
	s.AppSimpleList = v
	return s
}

func (s *DescribeAppSimpleListResponseBody) SetTotal(v int64) *DescribeAppSimpleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAppSimpleListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeAppSimpleListResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeAppSimpleListResponseBodyAppSimpleList struct {
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	NgAppId  *string `json:"NgAppId,omitempty" xml:"NgAppId,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EsnAppId *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
}

func (s DescribeAppSimpleListResponseBodyAppSimpleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSimpleListResponseBodyAppSimpleList) GoString() string {
	return s.String()
}

func (s *DescribeAppSimpleListResponseBodyAppSimpleList) SetAppName(v string) *DescribeAppSimpleListResponseBodyAppSimpleList {
	s.AppName = &v
	return s
}

func (s *DescribeAppSimpleListResponseBodyAppSimpleList) SetNgAppId(v string) *DescribeAppSimpleListResponseBodyAppSimpleList {
	s.NgAppId = &v
	return s
}

func (s *DescribeAppSimpleListResponseBodyAppSimpleList) SetAppId(v int64) *DescribeAppSimpleListResponseBodyAppSimpleList {
	s.AppId = &v
	return s
}

func (s *DescribeAppSimpleListResponseBodyAppSimpleList) SetEsnAppId(v string) *DescribeAppSimpleListResponseBodyAppSimpleList {
	s.EsnAppId = &v
	return s
}

type DescribeAppSimpleListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppSimpleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppSimpleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSimpleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppSimpleListResponse) SetHeaders(v map[string]*string) *DescribeAppSimpleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppSimpleListResponse) SetBody(v *DescribeAppSimpleListResponseBody) *DescribeAppSimpleListResponse {
	s.Body = v
	return s
}

type DescribeAsyncOperationRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeAsyncOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncOperationRequest) GoString() string {
	return s.String()
}

func (s *DescribeAsyncOperationRequest) SetSourceIp(v string) *DescribeAsyncOperationRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAsyncOperationRequest) SetLang(v string) *DescribeAsyncOperationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAsyncOperationRequest) SetBizId(v int64) *DescribeAsyncOperationRequest {
	s.BizId = &v
	return s
}

type DescribeAsyncOperationResponseBody struct {
	AsyncOperation *DescribeAsyncOperationResponseBodyAsyncOperation `json:"AsyncOperation,omitempty" xml:"AsyncOperation,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo     map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeAsyncOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncOperationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAsyncOperationResponseBody) SetAsyncOperation(v *DescribeAsyncOperationResponseBodyAsyncOperation) *DescribeAsyncOperationResponseBody {
	s.AsyncOperation = v
	return s
}

func (s *DescribeAsyncOperationResponseBody) SetRequestId(v string) *DescribeAsyncOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAsyncOperationResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeAsyncOperationResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeAsyncOperationResponseBodyAsyncOperation struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAsyncOperationResponseBodyAsyncOperation) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncOperationResponseBodyAsyncOperation) GoString() string {
	return s.String()
}

func (s *DescribeAsyncOperationResponseBodyAsyncOperation) SetStatus(v string) *DescribeAsyncOperationResponseBodyAsyncOperation {
	s.Status = &v
	return s
}

type DescribeAsyncOperationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAsyncOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAsyncOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAsyncOperationResponse) GoString() string {
	return s.String()
}

func (s *DescribeAsyncOperationResponse) SetHeaders(v map[string]*string) *DescribeAsyncOperationResponse {
	s.Headers = v
	return s
}

func (s *DescribeAsyncOperationResponse) SetBody(v *DescribeAsyncOperationResponseBody) *DescribeAsyncOperationResponse {
	s.Body = v
	return s
}

type DescribeAutomaticTraceabilitySlsLogRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EsnAppId  *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeAutomaticTraceabilitySlsLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomaticTraceabilitySlsLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetSourceIp(v string) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetLang(v string) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetStartTime(v int64) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetEndTime(v int64) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetPage(v int32) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.Page = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetPageSize(v int32) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetEsnAppId(v string) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.EsnAppId = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogRequest) SetSource(v string) *DescribeAutomaticTraceabilitySlsLogRequest {
	s.Source = &v
	return s
}

type DescribeAutomaticTraceabilitySlsLogResponseBody struct {
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic []*DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total     *int64                                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAutomaticTraceabilitySlsLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomaticTraceabilitySlsLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBody) SetRequestId(v string) *DescribeAutomaticTraceabilitySlsLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBody) SetStatistic(v []*DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) *DescribeAutomaticTraceabilitySlsLogResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBody) SetTotal(v int64) *DescribeAutomaticTraceabilitySlsLogResponseBody {
	s.Total = &v
	return s
}

type DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Time             *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Token            *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Point            *string `json:"Point,omitempty" xml:"Point,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	SourceIpLocation *string `json:"SourceIpLocation,omitempty" xml:"SourceIpLocation,omitempty"`
}

func (s DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) SetSourceIp(v string) *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic {
	s.SourceIp = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) SetTime(v string) *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic {
	s.Time = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) SetToken(v string) *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic {
	s.Token = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) SetPoint(v string) *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic {
	s.Point = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) SetIp(v string) *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic {
	s.Ip = &v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic) SetSourceIpLocation(v string) *DescribeAutomaticTraceabilitySlsLogResponseBodyStatistic {
	s.SourceIpLocation = &v
	return s
}

type DescribeAutomaticTraceabilitySlsLogResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutomaticTraceabilitySlsLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutomaticTraceabilitySlsLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomaticTraceabilitySlsLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomaticTraceabilitySlsLogResponse) SetHeaders(v map[string]*string) *DescribeAutomaticTraceabilitySlsLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomaticTraceabilitySlsLogResponse) SetBody(v *DescribeAutomaticTraceabilitySlsLogResponseBody) *DescribeAutomaticTraceabilitySlsLogResponse {
	s.Body = v
	return s
}

type DescribeBgpResSummaryRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeBgpResSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpResSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpResSummaryRequest) SetSourceIp(v string) *DescribeBgpResSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBgpResSummaryRequest) SetLang(v string) *DescribeBgpResSummaryRequest {
	s.Lang = &v
	return s
}

type DescribeBgpResSummaryResponseBody struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullNodesSummary *DescribeBgpResSummaryResponseBodyFullNodesSummary `json:"FullNodesSummary,omitempty" xml:"FullNodesSummary,omitempty" type:"Struct"`
	PromptInfo       map[string]interface{}                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeBgpResSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpResSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpResSummaryResponseBody) SetRequestId(v string) *DescribeBgpResSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBody) SetFullNodesSummary(v *DescribeBgpResSummaryResponseBodyFullNodesSummary) *DescribeBgpResSummaryResponseBody {
	s.FullNodesSummary = v
	return s
}

func (s *DescribeBgpResSummaryResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeBgpResSummaryResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeBgpResSummaryResponseBodyFullNodesSummary struct {
	Nodes       []*DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	CleanCount  *int32                                                    `json:"CleanCount,omitempty" xml:"CleanCount,omitempty"`
	HoleCount   *int32                                                    `json:"HoleCount,omitempty" xml:"HoleCount,omitempty"`
	UsedCount   *int32                                                    `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	UnUsedCount *int32                                                    `json:"UnUsedCount,omitempty" xml:"UnUsedCount,omitempty"`
	TotalCount  *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBgpResSummaryResponseBodyFullNodesSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpResSummaryResponseBodyFullNodesSummary) GoString() string {
	return s.String()
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummary) SetNodes(v []*DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) *DescribeBgpResSummaryResponseBodyFullNodesSummary {
	s.Nodes = v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummary) SetCleanCount(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummary {
	s.CleanCount = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummary) SetHoleCount(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummary {
	s.HoleCount = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummary) SetUsedCount(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummary {
	s.UsedCount = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummary) SetUnUsedCount(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummary {
	s.UnUsedCount = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummary) SetTotalCount(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummary {
	s.TotalCount = &v
	return s
}

type DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) GoString() string {
	return s.String()
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetType(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.Type = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetBizId(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizId = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetIp(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.Ip = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetRegionId(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetAppName(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppName = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetRegionName(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionName = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetGroupId(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetAppId(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppId = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetInUse(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.InUse = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetBizName(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizName = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetDdosStatus(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.DdosStatus = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetIsDisabled(v bool) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.IsDisabled = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetThreshold(v int32) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.Threshold = &v
	return s
}

func (s *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes) SetEsnBizId(v string) *DescribeBgpResSummaryResponseBodyFullNodesSummaryNodes {
	s.EsnBizId = &v
	return s
}

type DescribeBgpResSummaryResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBgpResSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBgpResSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBgpResSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpResSummaryResponse) SetHeaders(v map[string]*string) *DescribeBgpResSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpResSummaryResponse) SetBody(v *DescribeBgpResSummaryResponseBody) *DescribeBgpResSummaryResponse {
	s.Body = v
	return s
}

type DescribeBizListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizName  *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
}

func (s DescribeBizListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizListRequest) SetSourceIp(v string) *DescribeBizListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBizListRequest) SetLang(v string) *DescribeBizListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBizListRequest) SetAppId(v int64) *DescribeBizListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeBizListRequest) SetBizId(v int64) *DescribeBizListRequest {
	s.BizId = &v
	return s
}

func (s *DescribeBizListRequest) SetBizName(v string) *DescribeBizListRequest {
	s.BizName = &v
	return s
}

type DescribeBizListResponseBody struct {
	BizList    []*DescribeBizListResponseBodyBizList `json:"BizList,omitempty" xml:"BizList,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeBizListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizListResponseBody) SetBizList(v []*DescribeBizListResponseBodyBizList) *DescribeBizListResponseBody {
	s.BizList = v
	return s
}

func (s *DescribeBizListResponseBody) SetRequestId(v string) *DescribeBizListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizListResponseBody) SetTotal(v int64) *DescribeBizListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBizListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeBizListResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeBizListResponseBodyBizList struct {
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	NodeInUseNum *int32  `json:"NodeInUseNum,omitempty" xml:"NodeInUseNum,omitempty"`
	NgGroupId    *string `json:"NgGroupId,omitempty" xml:"NgGroupId,omitempty"`
	UseCc        *int32  `json:"UseCc,omitempty" xml:"UseCc,omitempty"`
	BizId        *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupNum     *int32  `json:"GroupNum,omitempty" xml:"GroupNum,omitempty"`
	NodeUnUseNum *int32  `json:"NodeUnUseNum,omitempty" xml:"NodeUnUseNum,omitempty"`
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizName      *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	NodeNum      *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	EsnBizId     *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	CcQps        *int32  `json:"CcQps,omitempty" xml:"CcQps,omitempty"`
}

func (s DescribeBizListResponseBodyBizList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizListResponseBodyBizList) GoString() string {
	return s.String()
}

func (s *DescribeBizListResponseBodyBizList) SetType(v int32) *DescribeBizListResponseBodyBizList {
	s.Type = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetNodeInUseNum(v int32) *DescribeBizListResponseBodyBizList {
	s.NodeInUseNum = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetNgGroupId(v string) *DescribeBizListResponseBodyBizList {
	s.NgGroupId = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetUseCc(v int32) *DescribeBizListResponseBodyBizList {
	s.UseCc = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetBizId(v int64) *DescribeBizListResponseBodyBizList {
	s.BizId = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetGroupNum(v int32) *DescribeBizListResponseBodyBizList {
	s.GroupNum = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetNodeUnUseNum(v int32) *DescribeBizListResponseBodyBizList {
	s.NodeUnUseNum = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetAppId(v int64) *DescribeBizListResponseBodyBizList {
	s.AppId = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetBizName(v string) *DescribeBizListResponseBodyBizList {
	s.BizName = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetNodeNum(v int32) *DescribeBizListResponseBodyBizList {
	s.NodeNum = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetEsnBizId(v string) *DescribeBizListResponseBodyBizList {
	s.EsnBizId = &v
	return s
}

func (s *DescribeBizListResponseBodyBizList) SetCcQps(v int32) *DescribeBizListResponseBodyBizList {
	s.CcQps = &v
	return s
}

type DescribeBizListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBizListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBizListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizListResponse) SetHeaders(v map[string]*string) *DescribeBizListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizListResponse) SetBody(v *DescribeBizListResponseBody) *DescribeBizListResponse {
	s.Body = v
	return s
}

type DescribeBizSimpleListRequest struct {
	SourceIp  *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId     *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
}

func (s DescribeBizSimpleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizSimpleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizSimpleListRequest) SetSourceIp(v string) *DescribeBizSimpleListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBizSimpleListRequest) SetLang(v string) *DescribeBizSimpleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBizSimpleListRequest) SetAppId(v int64) *DescribeBizSimpleListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeBizSimpleListRequest) SetAppIdList(v []*string) *DescribeBizSimpleListRequest {
	s.AppIdList = v
	return s
}

type DescribeBizSimpleListResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo    map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	BizSimpleList []*DescribeBizSimpleListResponseBodyBizSimpleList `json:"BizSimpleList,omitempty" xml:"BizSimpleList,omitempty" type:"Repeated"`
}

func (s DescribeBizSimpleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizSimpleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizSimpleListResponseBody) SetRequestId(v string) *DescribeBizSimpleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizSimpleListResponseBody) SetTotal(v int64) *DescribeBizSimpleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBizSimpleListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeBizSimpleListResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeBizSimpleListResponseBody) SetBizSimpleList(v []*DescribeBizSimpleListResponseBodyBizSimpleList) *DescribeBizSimpleListResponseBody {
	s.BizSimpleList = v
	return s
}

type DescribeBizSimpleListResponseBodyBizSimpleList struct {
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	NgGroupId *string `json:"NgGroupId,omitempty" xml:"NgGroupId,omitempty"`
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	UseCc     *int32  `json:"UseCc,omitempty" xml:"UseCc,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizName   *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	EsnBizId  *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	CcQps     *int32  `json:"CcQps,omitempty" xml:"CcQps,omitempty"`
}

func (s DescribeBizSimpleListResponseBodyBizSimpleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizSimpleListResponseBodyBizSimpleList) GoString() string {
	return s.String()
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetType(v int32) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.Type = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetNgGroupId(v string) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.NgGroupId = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetAppId(v int64) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.AppId = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetUseCc(v int32) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.UseCc = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetBizId(v int64) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.BizId = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetBizName(v string) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.BizName = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetEsnBizId(v string) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.EsnBizId = &v
	return s
}

func (s *DescribeBizSimpleListResponseBodyBizSimpleList) SetCcQps(v int32) *DescribeBizSimpleListResponseBodyBizSimpleList {
	s.CcQps = &v
	return s
}

type DescribeBizSimpleListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBizSimpleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBizSimpleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizSimpleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizSimpleListResponse) SetHeaders(v map[string]*string) *DescribeBizSimpleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizSimpleListResponse) SetBody(v *DescribeBizSimpleListResponseBody) *DescribeBizSimpleListResponse {
	s.Body = v
	return s
}

type DescribeBpsFlowRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeBpsFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBpsFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeBpsFlowRequest) SetSourceIp(v string) *DescribeBpsFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBpsFlowRequest) SetLang(v string) *DescribeBpsFlowRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBpsFlowRequest) SetBizId(v int64) *DescribeBpsFlowRequest {
	s.BizId = &v
	return s
}

func (s *DescribeBpsFlowRequest) SetGroupId(v string) *DescribeBpsFlowRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeBpsFlowRequest) SetBeginTime(v int64) *DescribeBpsFlowRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeBpsFlowRequest) SetEndTime(v int64) *DescribeBpsFlowRequest {
	s.EndTime = &v
	return s
}

type DescribeBpsFlowResponseBody struct {
	FlowData   *DescribeBpsFlowResponseBodyFlowData `json:"FlowData,omitempty" xml:"FlowData,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}               `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeBpsFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBpsFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBpsFlowResponseBody) SetFlowData(v *DescribeBpsFlowResponseBodyFlowData) *DescribeBpsFlowResponseBody {
	s.FlowData = v
	return s
}

func (s *DescribeBpsFlowResponseBody) SetRequestId(v string) *DescribeBpsFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBpsFlowResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeBpsFlowResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeBpsFlowResponseBodyFlowData struct {
	Items          []*DescribeBpsFlowResponseBodyFlowDataItems   `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Categories     []*string                                     `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	TimeScope      *DescribeBpsFlowResponseBodyFlowDataTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Struct"`
	CategoriesType *string                                       `json:"CategoriesType,omitempty" xml:"CategoriesType,omitempty"`
}

func (s DescribeBpsFlowResponseBodyFlowData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBpsFlowResponseBodyFlowData) GoString() string {
	return s.String()
}

func (s *DescribeBpsFlowResponseBodyFlowData) SetItems(v []*DescribeBpsFlowResponseBodyFlowDataItems) *DescribeBpsFlowResponseBodyFlowData {
	s.Items = v
	return s
}

func (s *DescribeBpsFlowResponseBodyFlowData) SetCategories(v []*string) *DescribeBpsFlowResponseBodyFlowData {
	s.Categories = v
	return s
}

func (s *DescribeBpsFlowResponseBodyFlowData) SetTimeScope(v *DescribeBpsFlowResponseBodyFlowDataTimeScope) *DescribeBpsFlowResponseBodyFlowData {
	s.TimeScope = v
	return s
}

func (s *DescribeBpsFlowResponseBodyFlowData) SetCategoriesType(v string) *DescribeBpsFlowResponseBodyFlowData {
	s.CategoriesType = &v
	return s
}

type DescribeBpsFlowResponseBodyFlowDataItems struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBpsFlowResponseBodyFlowDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBpsFlowResponseBodyFlowDataItems) GoString() string {
	return s.String()
}

func (s *DescribeBpsFlowResponseBodyFlowDataItems) SetValues(v []*string) *DescribeBpsFlowResponseBodyFlowDataItems {
	s.Values = v
	return s
}

func (s *DescribeBpsFlowResponseBodyFlowDataItems) SetName(v string) *DescribeBpsFlowResponseBodyFlowDataItems {
	s.Name = &v
	return s
}

type DescribeBpsFlowResponseBodyFlowDataTimeScope struct {
	Start    *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeBpsFlowResponseBodyFlowDataTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeBpsFlowResponseBodyFlowDataTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeBpsFlowResponseBodyFlowDataTimeScope) SetStart(v int64) *DescribeBpsFlowResponseBodyFlowDataTimeScope {
	s.Start = &v
	return s
}

func (s *DescribeBpsFlowResponseBodyFlowDataTimeScope) SetInterval(v int64) *DescribeBpsFlowResponseBodyFlowDataTimeScope {
	s.Interval = &v
	return s
}

type DescribeBpsFlowResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBpsFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBpsFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBpsFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeBpsFlowResponse) SetHeaders(v map[string]*string) *DescribeBpsFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeBpsFlowResponse) SetBody(v *DescribeBpsFlowResponseBody) *DescribeBpsFlowResponse {
	s.Body = v
	return s
}

type DescribeCcBlackListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcBlackListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcBlackListRequest) SetSourceIp(v string) *DescribeCcBlackListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcBlackListRequest) SetLang(v string) *DescribeCcBlackListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcBlackListRequest) SetBizId(v int64) *DescribeCcBlackListRequest {
	s.BizId = &v
	return s
}

type DescribeCcBlackListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                 `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	BlackList  []*string              `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
}

func (s DescribeCcBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcBlackListResponseBody) SetRequestId(v string) *DescribeCcBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcBlackListResponseBody) SetTotal(v int64) *DescribeCcBlackListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCcBlackListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcBlackListResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcBlackListResponseBody) SetBlackList(v []*string) *DescribeCcBlackListResponseBody {
	s.BlackList = v
	return s
}

type DescribeCcBlackListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcBlackListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcBlackListResponse) SetHeaders(v map[string]*string) *DescribeCcBlackListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcBlackListResponse) SetBody(v *DescribeCcBlackListResponseBody) *DescribeCcBlackListResponse {
	s.Body = v
	return s
}

type DescribeCcDevieIpListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeCcDevieIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcDevieIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcDevieIpListRequest) SetSourceIp(v string) *DescribeCcDevieIpListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcDevieIpListRequest) SetLang(v string) *DescribeCcDevieIpListRequest {
	s.Lang = &v
	return s
}

type DescribeCcDevieIpListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpList     []*string              `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	Total      *int32                 `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcDevieIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcDevieIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcDevieIpListResponseBody) SetRequestId(v string) *DescribeCcDevieIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcDevieIpListResponseBody) SetIpList(v []*string) *DescribeCcDevieIpListResponseBody {
	s.IpList = v
	return s
}

func (s *DescribeCcDevieIpListResponseBody) SetTotal(v int32) *DescribeCcDevieIpListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCcDevieIpListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcDevieIpListResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcDevieIpListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcDevieIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcDevieIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcDevieIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcDevieIpListResponse) SetHeaders(v map[string]*string) *DescribeCcDevieIpListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcDevieIpListResponse) SetBody(v *DescribeCcDevieIpListResponseBody) *DescribeCcDevieIpListResponse {
	s.Body = v
	return s
}

type DescribeCcFlowRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId      *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
}

func (s DescribeCcFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcFlowRequest) SetSourceIp(v string) *DescribeCcFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcFlowRequest) SetLang(v string) *DescribeCcFlowRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcFlowRequest) SetBizId(v int64) *DescribeCcFlowRequest {
	s.BizId = &v
	return s
}

func (s *DescribeCcFlowRequest) SetBeginTime(v int64) *DescribeCcFlowRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeCcFlowRequest) SetEndTime(v int64) *DescribeCcFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCcFlowRequest) SetInterval(v int64) *DescribeCcFlowRequest {
	s.Interval = &v
	return s
}

func (s *DescribeCcFlowRequest) SetApiVersion(v string) *DescribeCcFlowRequest {
	s.ApiVersion = &v
	return s
}

type DescribeCcFlowResponseBody struct {
	FlowData   *DescribeCcFlowResponseBodyFlowData `json:"FlowData,omitempty" xml:"FlowData,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}              `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcFlowResponseBody) SetFlowData(v *DescribeCcFlowResponseBodyFlowData) *DescribeCcFlowResponseBody {
	s.FlowData = v
	return s
}

func (s *DescribeCcFlowResponseBody) SetRequestId(v string) *DescribeCcFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcFlowResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcFlowResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcFlowResponseBodyFlowData struct {
	Items          []*DescribeCcFlowResponseBodyFlowDataItems   `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Categories     []*string                                    `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	TimeScope      *DescribeCcFlowResponseBodyFlowDataTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Struct"`
	CategoriesType *string                                      `json:"CategoriesType,omitempty" xml:"CategoriesType,omitempty"`
}

func (s DescribeCcFlowResponseBodyFlowData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcFlowResponseBodyFlowData) GoString() string {
	return s.String()
}

func (s *DescribeCcFlowResponseBodyFlowData) SetItems(v []*DescribeCcFlowResponseBodyFlowDataItems) *DescribeCcFlowResponseBodyFlowData {
	s.Items = v
	return s
}

func (s *DescribeCcFlowResponseBodyFlowData) SetCategories(v []*string) *DescribeCcFlowResponseBodyFlowData {
	s.Categories = v
	return s
}

func (s *DescribeCcFlowResponseBodyFlowData) SetTimeScope(v *DescribeCcFlowResponseBodyFlowDataTimeScope) *DescribeCcFlowResponseBodyFlowData {
	s.TimeScope = v
	return s
}

func (s *DescribeCcFlowResponseBodyFlowData) SetCategoriesType(v string) *DescribeCcFlowResponseBodyFlowData {
	s.CategoriesType = &v
	return s
}

type DescribeCcFlowResponseBodyFlowDataItems struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCcFlowResponseBodyFlowDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcFlowResponseBodyFlowDataItems) GoString() string {
	return s.String()
}

func (s *DescribeCcFlowResponseBodyFlowDataItems) SetValues(v []*string) *DescribeCcFlowResponseBodyFlowDataItems {
	s.Values = v
	return s
}

func (s *DescribeCcFlowResponseBodyFlowDataItems) SetName(v string) *DescribeCcFlowResponseBodyFlowDataItems {
	s.Name = &v
	return s
}

type DescribeCcFlowResponseBodyFlowDataTimeScope struct {
	Start    *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeCcFlowResponseBodyFlowDataTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcFlowResponseBodyFlowDataTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeCcFlowResponseBodyFlowDataTimeScope) SetStart(v int64) *DescribeCcFlowResponseBodyFlowDataTimeScope {
	s.Start = &v
	return s
}

func (s *DescribeCcFlowResponseBodyFlowDataTimeScope) SetInterval(v int64) *DescribeCcFlowResponseBodyFlowDataTimeScope {
	s.Interval = &v
	return s
}

type DescribeCcFlowResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcFlowResponse) SetHeaders(v map[string]*string) *DescribeCcFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcFlowResponse) SetBody(v *DescribeCcFlowResponseBody) *DescribeCcFlowResponse {
	s.Body = v
	return s
}

type DescribeCcIDCBlockSwitchRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcIDCBlockSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcIDCBlockSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcIDCBlockSwitchRequest) SetSourceIp(v string) *DescribeCcIDCBlockSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcIDCBlockSwitchRequest) SetLang(v string) *DescribeCcIDCBlockSwitchRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcIDCBlockSwitchRequest) SetBizId(v int64) *DescribeCcIDCBlockSwitchRequest {
	s.BizId = &v
	return s
}

type DescribeCcIDCBlockSwitchResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                        `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	CcSwitch   *DescribeCcIDCBlockSwitchResponseBodyCcSwitch `json:"CcSwitch,omitempty" xml:"CcSwitch,omitempty" type:"Struct"`
}

func (s DescribeCcIDCBlockSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcIDCBlockSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcIDCBlockSwitchResponseBody) SetRequestId(v string) *DescribeCcIDCBlockSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcIDCBlockSwitchResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcIDCBlockSwitchResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcIDCBlockSwitchResponseBody) SetCcSwitch(v *DescribeCcIDCBlockSwitchResponseBodyCcSwitch) *DescribeCcIDCBlockSwitchResponseBody {
	s.CcSwitch = v
	return s
}

type DescribeCcIDCBlockSwitchResponseBodyCcSwitch struct {
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeCcIDCBlockSwitchResponseBodyCcSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcIDCBlockSwitchResponseBodyCcSwitch) GoString() string {
	return s.String()
}

func (s *DescribeCcIDCBlockSwitchResponseBodyCcSwitch) SetEnable(v int32) *DescribeCcIDCBlockSwitchResponseBodyCcSwitch {
	s.Enable = &v
	return s
}

type DescribeCcIDCBlockSwitchResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcIDCBlockSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcIDCBlockSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcIDCBlockSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcIDCBlockSwitchResponse) SetHeaders(v map[string]*string) *DescribeCcIDCBlockSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcIDCBlockSwitchResponse) SetBody(v *DescribeCcIDCBlockSwitchResponseBody) *DescribeCcIDCBlockSwitchResponse {
	s.Body = v
	return s
}

type DescribeCcMaxFwRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	DestPort  *int32  `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCcMaxFwRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcMaxFwRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcMaxFwRequest) SetSourceIp(v string) *DescribeCcMaxFwRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetLang(v string) *DescribeCcMaxFwRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetBizId(v int64) *DescribeCcMaxFwRequest {
	s.BizId = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetBeginTime(v int64) *DescribeCcMaxFwRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetEndTime(v int64) *DescribeCcMaxFwRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetSource(v string) *DescribeCcMaxFwRequest {
	s.Source = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetDestPort(v int32) *DescribeCcMaxFwRequest {
	s.DestPort = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetPage(v int32) *DescribeCcMaxFwRequest {
	s.Page = &v
	return s
}

func (s *DescribeCcMaxFwRequest) SetPageSize(v int32) *DescribeCcMaxFwRequest {
	s.PageSize = &v
	return s
}

type DescribeCcMaxFwResponseBody struct {
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FwMaxInfoList []*DescribeCcMaxFwResponseBodyFwMaxInfoList `json:"FwMaxInfoList,omitempty" xml:"FwMaxInfoList,omitempty" type:"Repeated"`
	Total         *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo    map[string]interface{}                      `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcMaxFwResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcMaxFwResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcMaxFwResponseBody) SetRequestId(v string) *DescribeCcMaxFwResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcMaxFwResponseBody) SetFwMaxInfoList(v []*DescribeCcMaxFwResponseBodyFwMaxInfoList) *DescribeCcMaxFwResponseBody {
	s.FwMaxInfoList = v
	return s
}

func (s *DescribeCcMaxFwResponseBody) SetTotal(v int64) *DescribeCcMaxFwResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCcMaxFwResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcMaxFwResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcMaxFwResponseBodyFwMaxInfoList struct {
	InWhiteList *bool   `json:"InWhiteList,omitempty" xml:"InWhiteList,omitempty"`
	Time        *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DestPort    *int32  `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	Action      *int32  `json:"Action,omitempty" xml:"Action,omitempty"`
	Protocol    *int32  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InBlackList *bool   `json:"InBlackList,omitempty" xml:"InBlackList,omitempty"`
	Count       *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeCcMaxFwResponseBodyFwMaxInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcMaxFwResponseBodyFwMaxInfoList) GoString() string {
	return s.String()
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetInWhiteList(v bool) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.InWhiteList = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetTime(v int64) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.Time = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetSourceIp(v string) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetDestPort(v int32) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.DestPort = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetAction(v int32) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.Action = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetProtocol(v int32) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.Protocol = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetInBlackList(v bool) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.InBlackList = &v
	return s
}

func (s *DescribeCcMaxFwResponseBodyFwMaxInfoList) SetCount(v int32) *DescribeCcMaxFwResponseBodyFwMaxInfoList {
	s.Count = &v
	return s
}

type DescribeCcMaxFwResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcMaxFwResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcMaxFwResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcMaxFwResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcMaxFwResponse) SetHeaders(v map[string]*string) *DescribeCcMaxFwResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcMaxFwResponse) SetBody(v *DescribeCcMaxFwResponseBody) *DescribeCcMaxFwResponse {
	s.Body = v
	return s
}

type DescribeCcResSummaryRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeCcResSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcResSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcResSummaryRequest) SetSourceIp(v string) *DescribeCcResSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcResSummaryRequest) SetLang(v string) *DescribeCcResSummaryRequest {
	s.Lang = &v
	return s
}

type DescribeCcResSummaryResponseBody struct {
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullNodesSummary *DescribeCcResSummaryResponseBodyFullNodesSummary `json:"FullNodesSummary,omitempty" xml:"FullNodesSummary,omitempty" type:"Struct"`
	PromptInfo       map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcResSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcResSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcResSummaryResponseBody) SetRequestId(v string) *DescribeCcResSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcResSummaryResponseBody) SetFullNodesSummary(v *DescribeCcResSummaryResponseBodyFullNodesSummary) *DescribeCcResSummaryResponseBody {
	s.FullNodesSummary = v
	return s
}

func (s *DescribeCcResSummaryResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcResSummaryResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcResSummaryResponseBodyFullNodesSummary struct {
	UnUsedQps     *int32                                                           `json:"UnUsedQps,omitempty" xml:"UnUsedQps,omitempty"`
	UsedQps       *int32                                                           `json:"UsedQps,omitempty" xml:"UsedQps,omitempty"`
	TotalQps      *int32                                                           `json:"TotalQps,omitempty" xml:"TotalQps,omitempty"`
	NGResInfoList []*DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList `json:"NGResInfoList,omitempty" xml:"NGResInfoList,omitempty" type:"Repeated"`
}

func (s DescribeCcResSummaryResponseBodyFullNodesSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcResSummaryResponseBodyFullNodesSummary) GoString() string {
	return s.String()
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummary) SetUnUsedQps(v int32) *DescribeCcResSummaryResponseBodyFullNodesSummary {
	s.UnUsedQps = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummary) SetUsedQps(v int32) *DescribeCcResSummaryResponseBodyFullNodesSummary {
	s.UsedQps = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummary) SetTotalQps(v int32) *DescribeCcResSummaryResponseBodyFullNodesSummary {
	s.TotalQps = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummary) SetNGResInfoList(v []*DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) *DescribeCcResSummaryResponseBodyFullNodesSummary {
	s.NGResInfoList = v
	return s
}

type DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	UsedQps *int32  `json:"UsedQps,omitempty" xml:"UsedQps,omitempty"`
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) GoString() string {
	return s.String()
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) SetAppName(v string) *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList {
	s.AppName = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) SetAppId(v string) *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList {
	s.AppId = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) SetUsedQps(v int32) *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList {
	s.UsedQps = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) SetBizName(v string) *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList {
	s.BizName = &v
	return s
}

func (s *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList) SetBizId(v string) *DescribeCcResSummaryResponseBodyFullNodesSummaryNGResInfoList {
	s.BizId = &v
	return s
}

type DescribeCcResSummaryResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcResSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcResSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcResSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcResSummaryResponse) SetHeaders(v map[string]*string) *DescribeCcResSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcResSummaryResponse) SetBody(v *DescribeCcResSummaryResponseBody) *DescribeCcResSummaryResponse {
	s.Body = v
	return s
}

type DescribeCcRouteRulesRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId        *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RouteComment *string `json:"RouteComment,omitempty" xml:"RouteComment,omitempty"`
	RouteId      *string `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	RouteIp      *string `json:"RouteIp,omitempty" xml:"RouteIp,omitempty"`
}

func (s DescribeCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteRulesRequest) SetSourceIp(v string) *DescribeCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcRouteRulesRequest) SetLang(v string) *DescribeCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcRouteRulesRequest) SetBizId(v int64) *DescribeCcRouteRulesRequest {
	s.BizId = &v
	return s
}

func (s *DescribeCcRouteRulesRequest) SetRouteComment(v string) *DescribeCcRouteRulesRequest {
	s.RouteComment = &v
	return s
}

func (s *DescribeCcRouteRulesRequest) SetRouteId(v string) *DescribeCcRouteRulesRequest {
	s.RouteId = &v
	return s
}

func (s *DescribeCcRouteRulesRequest) SetRouteIp(v string) *DescribeCcRouteRulesRequest {
	s.RouteIp = &v
	return s
}

type DescribeCcRouteRulesResponseBody struct {
	RuleQueryResult *DescribeCcRouteRulesResponseBodyRuleQueryResult `json:"RuleQueryResult,omitempty" xml:"RuleQueryResult,omitempty" type:"Struct"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo      map[string]interface{}                           `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteRulesResponseBody) SetRuleQueryResult(v *DescribeCcRouteRulesResponseBodyRuleQueryResult) *DescribeCcRouteRulesResponseBody {
	s.RuleQueryResult = v
	return s
}

func (s *DescribeCcRouteRulesResponseBody) SetRequestId(v string) *DescribeCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcRouteRulesResponseBodyRuleQueryResult struct {
	AppId      *string                                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RouteRules []*DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules `json:"RouteRules,omitempty" xml:"RouteRules,omitempty" type:"Repeated"`
	BizId      *string                                                      `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcRouteRulesResponseBodyRuleQueryResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteRulesResponseBodyRuleQueryResult) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteRulesResponseBodyRuleQueryResult) SetAppId(v string) *DescribeCcRouteRulesResponseBodyRuleQueryResult {
	s.AppId = &v
	return s
}

func (s *DescribeCcRouteRulesResponseBodyRuleQueryResult) SetRouteRules(v []*DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules) *DescribeCcRouteRulesResponseBodyRuleQueryResult {
	s.RouteRules = v
	return s
}

func (s *DescribeCcRouteRulesResponseBodyRuleQueryResult) SetBizId(v string) *DescribeCcRouteRulesResponseBodyRuleQueryResult {
	s.BizId = &v
	return s
}

type DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules struct {
	Comment  *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Rservers []*string `json:"Rservers,omitempty" xml:"Rservers,omitempty" type:"Repeated"`
	Id       *string   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules) SetComment(v string) *DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules {
	s.Comment = &v
	return s
}

func (s *DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules) SetRservers(v []*string) *DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules {
	s.Rservers = v
	return s
}

func (s *DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules) SetId(v string) *DescribeCcRouteRulesResponseBodyRuleQueryResultRouteRules {
	s.Id = &v
	return s
}

type DescribeCcRouteRulesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteRulesResponse) SetHeaders(v map[string]*string) *DescribeCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcRouteRulesResponse) SetBody(v *DescribeCcRouteRulesResponseBody) *DescribeCcRouteRulesResponse {
	s.Body = v
	return s
}

type DescribeCcRouteSwitchRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcRouteSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteSwitchRequest) SetSourceIp(v string) *DescribeCcRouteSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcRouteSwitchRequest) SetLang(v string) *DescribeCcRouteSwitchRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcRouteSwitchRequest) SetBizId(v int64) *DescribeCcRouteSwitchRequest {
	s.BizId = &v
	return s
}

type DescribeCcRouteSwitchResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Switch     *DescribeCcRouteSwitchResponseBodySwitch `json:"Switch,omitempty" xml:"Switch,omitempty" type:"Struct"`
	PromptInfo map[string]interface{}                   `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcRouteSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteSwitchResponseBody) SetRequestId(v string) *DescribeCcRouteSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcRouteSwitchResponseBody) SetSwitch(v *DescribeCcRouteSwitchResponseBodySwitch) *DescribeCcRouteSwitchResponseBody {
	s.Switch = v
	return s
}

func (s *DescribeCcRouteSwitchResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcRouteSwitchResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcRouteSwitchResponseBodySwitch struct {
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeCcRouteSwitchResponseBodySwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteSwitchResponseBodySwitch) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteSwitchResponseBodySwitch) SetEnable(v int32) *DescribeCcRouteSwitchResponseBodySwitch {
	s.Enable = &v
	return s
}

type DescribeCcRouteSwitchResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcRouteSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcRouteSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcRouteSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcRouteSwitchResponse) SetHeaders(v map[string]*string) *DescribeCcRouteSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcRouteSwitchResponse) SetBody(v *DescribeCcRouteSwitchResponseBody) *DescribeCcRouteSwitchResponse {
	s.Body = v
	return s
}

type DescribeCcSocketDetailRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeCcSocketDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcSocketDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcSocketDetailRequest) SetSourceIp(v string) *DescribeCcSocketDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcSocketDetailRequest) SetLang(v string) *DescribeCcSocketDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcSocketDetailRequest) SetBizId(v int64) *DescribeCcSocketDetailRequest {
	s.BizId = &v
	return s
}

func (s *DescribeCcSocketDetailRequest) SetBeginTime(v int64) *DescribeCcSocketDetailRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeCcSocketDetailRequest) SetEndTime(v int64) *DescribeCcSocketDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCcSocketDetailRequest) SetInterval(v int64) *DescribeCcSocketDetailRequest {
	s.Interval = &v
	return s
}

type DescribeCcSocketDetailResponseBody struct {
	FlowData   *DescribeCcSocketDetailResponseBodyFlowData `json:"FlowData,omitempty" xml:"FlowData,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                      `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeCcSocketDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcSocketDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcSocketDetailResponseBody) SetFlowData(v *DescribeCcSocketDetailResponseBodyFlowData) *DescribeCcSocketDetailResponseBody {
	s.FlowData = v
	return s
}

func (s *DescribeCcSocketDetailResponseBody) SetRequestId(v string) *DescribeCcSocketDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcSocketDetailResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcSocketDetailResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeCcSocketDetailResponseBodyFlowData struct {
	Items          []*DescribeCcSocketDetailResponseBodyFlowDataItems   `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Categories     []*string                                            `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	TimeScope      *DescribeCcSocketDetailResponseBodyFlowDataTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Struct"`
	CategoriesType *string                                              `json:"CategoriesType,omitempty" xml:"CategoriesType,omitempty"`
}

func (s DescribeCcSocketDetailResponseBodyFlowData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcSocketDetailResponseBodyFlowData) GoString() string {
	return s.String()
}

func (s *DescribeCcSocketDetailResponseBodyFlowData) SetItems(v []*DescribeCcSocketDetailResponseBodyFlowDataItems) *DescribeCcSocketDetailResponseBodyFlowData {
	s.Items = v
	return s
}

func (s *DescribeCcSocketDetailResponseBodyFlowData) SetCategories(v []*string) *DescribeCcSocketDetailResponseBodyFlowData {
	s.Categories = v
	return s
}

func (s *DescribeCcSocketDetailResponseBodyFlowData) SetTimeScope(v *DescribeCcSocketDetailResponseBodyFlowDataTimeScope) *DescribeCcSocketDetailResponseBodyFlowData {
	s.TimeScope = v
	return s
}

func (s *DescribeCcSocketDetailResponseBodyFlowData) SetCategoriesType(v string) *DescribeCcSocketDetailResponseBodyFlowData {
	s.CategoriesType = &v
	return s
}

type DescribeCcSocketDetailResponseBodyFlowDataItems struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCcSocketDetailResponseBodyFlowDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcSocketDetailResponseBodyFlowDataItems) GoString() string {
	return s.String()
}

func (s *DescribeCcSocketDetailResponseBodyFlowDataItems) SetValues(v []*string) *DescribeCcSocketDetailResponseBodyFlowDataItems {
	s.Values = v
	return s
}

func (s *DescribeCcSocketDetailResponseBodyFlowDataItems) SetName(v string) *DescribeCcSocketDetailResponseBodyFlowDataItems {
	s.Name = &v
	return s
}

type DescribeCcSocketDetailResponseBodyFlowDataTimeScope struct {
	Start    *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeCcSocketDetailResponseBodyFlowDataTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcSocketDetailResponseBodyFlowDataTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeCcSocketDetailResponseBodyFlowDataTimeScope) SetStart(v int64) *DescribeCcSocketDetailResponseBodyFlowDataTimeScope {
	s.Start = &v
	return s
}

func (s *DescribeCcSocketDetailResponseBodyFlowDataTimeScope) SetInterval(v int64) *DescribeCcSocketDetailResponseBodyFlowDataTimeScope {
	s.Interval = &v
	return s
}

type DescribeCcSocketDetailResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcSocketDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcSocketDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcSocketDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcSocketDetailResponse) SetHeaders(v map[string]*string) *DescribeCcSocketDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcSocketDetailResponse) SetBody(v *DescribeCcSocketDetailResponseBody) *DescribeCcSocketDetailResponse {
	s.Body = v
	return s
}

type DescribeCcTotalFwRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCcTotalFwRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTotalFwRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcTotalFwRequest) SetSourceIp(v string) *DescribeCcTotalFwRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetLang(v string) *DescribeCcTotalFwRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetBizId(v int64) *DescribeCcTotalFwRequest {
	s.BizId = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetBeginTime(v int64) *DescribeCcTotalFwRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetEndTime(v int64) *DescribeCcTotalFwRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetSource(v string) *DescribeCcTotalFwRequest {
	s.Source = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetPage(v int32) *DescribeCcTotalFwRequest {
	s.Page = &v
	return s
}

func (s *DescribeCcTotalFwRequest) SetPageSize(v int32) *DescribeCcTotalFwRequest {
	s.PageSize = &v
	return s
}

type DescribeCcTotalFwResponseBody struct {
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total           *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo      map[string]interface{}                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	FwTotalInfoList []*DescribeCcTotalFwResponseBodyFwTotalInfoList `json:"FwTotalInfoList,omitempty" xml:"FwTotalInfoList,omitempty" type:"Repeated"`
}

func (s DescribeCcTotalFwResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTotalFwResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcTotalFwResponseBody) SetRequestId(v string) *DescribeCcTotalFwResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcTotalFwResponseBody) SetTotal(v int64) *DescribeCcTotalFwResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCcTotalFwResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcTotalFwResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcTotalFwResponseBody) SetFwTotalInfoList(v []*DescribeCcTotalFwResponseBodyFwTotalInfoList) *DescribeCcTotalFwResponseBody {
	s.FwTotalInfoList = v
	return s
}

type DescribeCcTotalFwResponseBodyFwTotalInfoList struct {
	InWhiteList    *bool   `json:"InWhiteList,omitempty" xml:"InWhiteList,omitempty"`
	Time           *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	AttackType     *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	InBlackList    *bool   `json:"InBlackList,omitempty" xml:"InBlackList,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	SourceLocation *string `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
}

func (s DescribeCcTotalFwResponseBodyFwTotalInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTotalFwResponseBodyFwTotalInfoList) GoString() string {
	return s.String()
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetInWhiteList(v bool) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.InWhiteList = &v
	return s
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetTime(v int64) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.Time = &v
	return s
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetSourceIp(v string) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetAttackType(v string) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.AttackType = &v
	return s
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetInBlackList(v bool) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.InBlackList = &v
	return s
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetCount(v int32) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.Count = &v
	return s
}

func (s *DescribeCcTotalFwResponseBodyFwTotalInfoList) SetSourceLocation(v string) *DescribeCcTotalFwResponseBodyFwTotalInfoList {
	s.SourceLocation = &v
	return s
}

type DescribeCcTotalFwResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcTotalFwResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcTotalFwResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTotalFwResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcTotalFwResponse) SetHeaders(v map[string]*string) *DescribeCcTotalFwResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcTotalFwResponse) SetBody(v *DescribeCcTotalFwResponseBody) *DescribeCcTotalFwResponse {
	s.Body = v
	return s
}

type DescribeCcTunnelRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcTunnelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTunnelRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcTunnelRequest) SetSourceIp(v string) *DescribeCcTunnelRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcTunnelRequest) SetLang(v string) *DescribeCcTunnelRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcTunnelRequest) SetBizId(v string) *DescribeCcTunnelRequest {
	s.BizId = &v
	return s
}

type DescribeCcTunnelResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	CcTunnel   *DescribeCcTunnelResponseBodyCcTunnel `json:"CcTunnel,omitempty" xml:"CcTunnel,omitempty" type:"Struct"`
}

func (s DescribeCcTunnelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcTunnelResponseBody) SetRequestId(v string) *DescribeCcTunnelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcTunnelResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcTunnelResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcTunnelResponseBody) SetCcTunnel(v *DescribeCcTunnelResponseBodyCcTunnel) *DescribeCcTunnelResponseBody {
	s.CcTunnel = v
	return s
}

type DescribeCcTunnelResponseBodyCcTunnel struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OnlyAllow *bool   `json:"OnlyAllow,omitempty" xml:"OnlyAllow,omitempty"`
	Gray      *int32  `json:"Gray,omitempty" xml:"Gray,omitempty"`
}

func (s DescribeCcTunnelResponseBodyCcTunnel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTunnelResponseBodyCcTunnel) GoString() string {
	return s.String()
}

func (s *DescribeCcTunnelResponseBodyCcTunnel) SetStatus(v string) *DescribeCcTunnelResponseBodyCcTunnel {
	s.Status = &v
	return s
}

func (s *DescribeCcTunnelResponseBodyCcTunnel) SetOnlyAllow(v bool) *DescribeCcTunnelResponseBodyCcTunnel {
	s.OnlyAllow = &v
	return s
}

func (s *DescribeCcTunnelResponseBodyCcTunnel) SetGray(v int32) *DescribeCcTunnelResponseBodyCcTunnel {
	s.Gray = &v
	return s
}

type DescribeCcTunnelResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcTunnelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcTunnelResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcTunnelResponse) SetHeaders(v map[string]*string) *DescribeCcTunnelResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcTunnelResponse) SetBody(v *DescribeCcTunnelResponseBody) *DescribeCcTunnelResponse {
	s.Body = v
	return s
}

type DescribeCcWhiteListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcWhiteListRequest) SetSourceIp(v string) *DescribeCcWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcWhiteListRequest) SetLang(v string) *DescribeCcWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcWhiteListRequest) SetBizId(v int64) *DescribeCcWhiteListRequest {
	s.BizId = &v
	return s
}

type DescribeCcWhiteListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                 `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	WhiteList  []*string              `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeCcWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcWhiteListResponseBody) SetRequestId(v string) *DescribeCcWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcWhiteListResponseBody) SetTotal(v int64) *DescribeCcWhiteListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCcWhiteListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcWhiteListResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcWhiteListResponseBody) SetWhiteList(v []*string) *DescribeCcWhiteListResponseBody {
	s.WhiteList = v
	return s
}

type DescribeCcWhiteListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcWhiteListResponse) SetHeaders(v map[string]*string) *DescribeCcWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcWhiteListResponse) SetBody(v *DescribeCcWhiteListResponseBody) *DescribeCcWhiteListResponse {
	s.Body = v
	return s
}

type DescribeCcZoneBlockConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcZoneBlockConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZoneBlockConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcZoneBlockConfigRequest) SetSourceIp(v string) *DescribeCcZoneBlockConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcZoneBlockConfigRequest) SetLang(v string) *DescribeCcZoneBlockConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcZoneBlockConfigRequest) SetBizId(v int64) *DescribeCcZoneBlockConfigRequest {
	s.BizId = &v
	return s
}

type DescribeCcZoneBlockConfigResponseBody struct {
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo  map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	BlockConfig *DescribeCcZoneBlockConfigResponseBodyBlockConfig `json:"BlockConfig,omitempty" xml:"BlockConfig,omitempty" type:"Struct"`
}

func (s DescribeCcZoneBlockConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZoneBlockConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcZoneBlockConfigResponseBody) SetRequestId(v string) *DescribeCcZoneBlockConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcZoneBlockConfigResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcZoneBlockConfigResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcZoneBlockConfigResponseBody) SetBlockConfig(v *DescribeCcZoneBlockConfigResponseBodyBlockConfig) *DescribeCcZoneBlockConfigResponseBody {
	s.BlockConfig = v
	return s
}

type DescribeCcZoneBlockConfigResponseBodyBlockConfig struct {
	CcZones []*DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones `json:"CcZones,omitempty" xml:"CcZones,omitempty" type:"Repeated"`
	Enable  *int32                                                     `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeCcZoneBlockConfigResponseBodyBlockConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZoneBlockConfigResponseBodyBlockConfig) GoString() string {
	return s.String()
}

func (s *DescribeCcZoneBlockConfigResponseBodyBlockConfig) SetCcZones(v []*DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones) *DescribeCcZoneBlockConfigResponseBodyBlockConfig {
	s.CcZones = v
	return s
}

func (s *DescribeCcZoneBlockConfigResponseBodyBlockConfig) SetEnable(v int32) *DescribeCcZoneBlockConfigResponseBodyBlockConfig {
	s.Enable = &v
	return s
}

type DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones struct {
	Provinces []*string `json:"Provinces,omitempty" xml:"Provinces,omitempty" type:"Repeated"`
	Country   *string   `json:"Country,omitempty" xml:"Country,omitempty"`
}

func (s DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones) GoString() string {
	return s.String()
}

func (s *DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones) SetProvinces(v []*string) *DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones {
	s.Provinces = v
	return s
}

func (s *DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones) SetCountry(v string) *DescribeCcZoneBlockConfigResponseBodyBlockConfigCcZones {
	s.Country = &v
	return s
}

type DescribeCcZoneBlockConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcZoneBlockConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcZoneBlockConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZoneBlockConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcZoneBlockConfigResponse) SetHeaders(v map[string]*string) *DescribeCcZoneBlockConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcZoneBlockConfigResponse) SetBody(v *DescribeCcZoneBlockConfigResponseBody) *DescribeCcZoneBlockConfigResponse {
	s.Body = v
	return s
}

type DescribeCcZonesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeCcZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcZonesRequest) SetSourceIp(v string) *DescribeCcZonesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcZonesRequest) SetLang(v string) *DescribeCcZonesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcZonesRequest) SetBizId(v int64) *DescribeCcZonesRequest {
	s.BizId = &v
	return s
}

type DescribeCcZonesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	ZoneList   []*string              `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
}

func (s DescribeCcZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcZonesResponseBody) SetRequestId(v string) *DescribeCcZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcZonesResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeCcZonesResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeCcZonesResponseBody) SetZoneList(v []*string) *DescribeCcZonesResponseBody {
	s.ZoneList = v
	return s
}

type DescribeCcZonesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcZonesResponse) SetHeaders(v map[string]*string) *DescribeCcZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcZonesResponse) SetBody(v *DescribeCcZonesResponseBody) *DescribeCcZonesResponse {
	s.Body = v
	return s
}

type DescribeDailyDetailsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDailyDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDailyDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDailyDetailsRequest) SetSourceIp(v string) *DescribeDailyDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDailyDetailsRequest) SetLang(v string) *DescribeDailyDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDailyDetailsRequest) SetStartTime(v int64) *DescribeDailyDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDailyDetailsRequest) SetEndTime(v int64) *DescribeDailyDetailsRequest {
	s.EndTime = &v
	return s
}

type DescribeDailyDetailsResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DailyDetails []*DescribeDailyDetailsResponseBodyDailyDetails `json:"DailyDetails,omitempty" xml:"DailyDetails,omitempty" type:"Repeated"`
}

func (s DescribeDailyDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDailyDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDailyDetailsResponseBody) SetRequestId(v string) *DescribeDailyDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDailyDetailsResponseBody) SetDailyDetails(v []*DescribeDailyDetailsResponseBodyDailyDetails) *DescribeDailyDetailsResponseBody {
	s.DailyDetails = v
	return s
}

type DescribeDailyDetailsResponseBodyDailyDetails struct {
	IosActiveCount     *int64 `json:"IosActiveCount,omitempty" xml:"IosActiveCount,omitempty"`
	Index              *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	AndroidCount       *int64 `json:"AndroidCount,omitempty" xml:"AndroidCount,omitempty"`
	IosCount           *int64 `json:"IosCount,omitempty" xml:"IosCount,omitempty"`
	NewCount           *int64 `json:"NewCount,omitempty" xml:"NewCount,omitempty"`
	IpActiveCount      *int64 `json:"IpActiveCount,omitempty" xml:"IpActiveCount,omitempty"`
	ActiveCount        *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	Count              *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	AndroidActiveCount *int64 `json:"AndroidActiveCount,omitempty" xml:"AndroidActiveCount,omitempty"`
}

func (s DescribeDailyDetailsResponseBodyDailyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeDailyDetailsResponseBodyDailyDetails) GoString() string {
	return s.String()
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetIosActiveCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.IosActiveCount = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetIndex(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.Index = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetAndroidCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.AndroidCount = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetIosCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.IosCount = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetNewCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.NewCount = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetIpActiveCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.IpActiveCount = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetActiveCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.ActiveCount = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.Count = &v
	return s
}

func (s *DescribeDailyDetailsResponseBodyDailyDetails) SetAndroidActiveCount(v int64) *DescribeDailyDetailsResponseBodyDailyDetails {
	s.AndroidActiveCount = &v
	return s
}

type DescribeDailyDetailsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDailyDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDailyDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDailyDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDailyDetailsResponse) SetHeaders(v map[string]*string) *DescribeDailyDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDailyDetailsResponse) SetBody(v *DescribeDailyDetailsResponseBody) *DescribeDailyDetailsResponse {
	s.Body = v
	return s
}

type DescribeDownloadUrlsForAppKeyRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppKeyVersion *string `json:"AppKeyVersion,omitempty" xml:"AppKeyVersion,omitempty"`
}

func (s DescribeDownloadUrlsForAppKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForAppKeyRequest) SetSourceIp(v string) *DescribeDownloadUrlsForAppKeyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyRequest) SetLang(v string) *DescribeDownloadUrlsForAppKeyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyRequest) SetAppId(v int64) *DescribeDownloadUrlsForAppKeyRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyRequest) SetAppKeyVersion(v string) *DescribeDownloadUrlsForAppKeyRequest {
	s.AppKeyVersion = &v
	return s
}

type DescribeDownloadUrlsForAppKeyResponseBody struct {
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                              `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	UrlResult  *DescribeDownloadUrlsForAppKeyResponseBodyUrlResult `json:"UrlResult,omitempty" xml:"UrlResult,omitempty" type:"Struct"`
}

func (s DescribeDownloadUrlsForAppKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForAppKeyResponseBody) SetRequestId(v string) *DescribeDownloadUrlsForAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeDownloadUrlsForAppKeyResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyResponseBody) SetUrlResult(v *DescribeDownloadUrlsForAppKeyResponseBodyUrlResult) *DescribeDownloadUrlsForAppKeyResponseBody {
	s.UrlResult = v
	return s
}

type DescribeDownloadUrlsForAppKeyResponseBodyUrlResult struct {
	AppId *int64                                                  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Urls  *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Struct"`
}

func (s DescribeDownloadUrlsForAppKeyResponseBodyUrlResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForAppKeyResponseBodyUrlResult) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForAppKeyResponseBodyUrlResult) SetAppId(v int64) *DescribeDownloadUrlsForAppKeyResponseBodyUrlResult {
	s.AppId = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyResponseBodyUrlResult) SetUrls(v *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls) *DescribeDownloadUrlsForAppKeyResponseBodyUrlResult {
	s.Urls = v
	return s
}

type DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls struct {
	Android *string `json:"Android,omitempty" xml:"Android,omitempty"`
	Wins    *string `json:"Wins,omitempty" xml:"Wins,omitempty"`
	IOS     *string `json:"IOS,omitempty" xml:"IOS,omitempty"`
}

func (s DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls) SetAndroid(v string) *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls {
	s.Android = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls) SetWins(v string) *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls {
	s.Wins = &v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls) SetIOS(v string) *DescribeDownloadUrlsForAppKeyResponseBodyUrlResultUrls {
	s.IOS = &v
	return s
}

type DescribeDownloadUrlsForAppKeyResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDownloadUrlsForAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadUrlsForAppKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForAppKeyResponse) SetHeaders(v map[string]*string) *DescribeDownloadUrlsForAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadUrlsForAppKeyResponse) SetBody(v *DescribeDownloadUrlsForAppKeyResponseBody) *DescribeDownloadUrlsForAppKeyResponse {
	s.Body = v
	return s
}

type DescribeDownloadUrlsForSDKRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s DescribeDownloadUrlsForSDKRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForSDKRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForSDKRequest) SetSourceIp(v string) *DescribeDownloadUrlsForSDKRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKRequest) SetLang(v string) *DescribeDownloadUrlsForSDKRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKRequest) SetAppId(v int64) *DescribeDownloadUrlsForSDKRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKRequest) SetSdkVersion(v string) *DescribeDownloadUrlsForSDKRequest {
	s.SdkVersion = &v
	return s
}

type DescribeDownloadUrlsForSDKResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                           `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	UrlResult  *DescribeDownloadUrlsForSDKResponseBodyUrlResult `json:"UrlResult,omitempty" xml:"UrlResult,omitempty" type:"Struct"`
}

func (s DescribeDownloadUrlsForSDKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForSDKResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForSDKResponseBody) SetRequestId(v string) *DescribeDownloadUrlsForSDKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeDownloadUrlsForSDKResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeDownloadUrlsForSDKResponseBody) SetUrlResult(v *DescribeDownloadUrlsForSDKResponseBodyUrlResult) *DescribeDownloadUrlsForSDKResponseBody {
	s.UrlResult = v
	return s
}

type DescribeDownloadUrlsForSDKResponseBodyUrlResult struct {
	AppId *int64                                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Urls  *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Struct"`
}

func (s DescribeDownloadUrlsForSDKResponseBodyUrlResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForSDKResponseBodyUrlResult) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForSDKResponseBodyUrlResult) SetAppId(v int64) *DescribeDownloadUrlsForSDKResponseBodyUrlResult {
	s.AppId = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKResponseBodyUrlResult) SetUrls(v *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls) *DescribeDownloadUrlsForSDKResponseBodyUrlResult {
	s.Urls = v
	return s
}

type DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls struct {
	Android *string `json:"Android,omitempty" xml:"Android,omitempty"`
	Wins    *string `json:"Wins,omitempty" xml:"Wins,omitempty"`
	IOS     *string `json:"IOS,omitempty" xml:"IOS,omitempty"`
}

func (s DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls) SetAndroid(v string) *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls {
	s.Android = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls) SetWins(v string) *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls {
	s.Wins = &v
	return s
}

func (s *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls) SetIOS(v string) *DescribeDownloadUrlsForSDKResponseBodyUrlResultUrls {
	s.IOS = &v
	return s
}

type DescribeDownloadUrlsForSDKResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDownloadUrlsForSDKResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDownloadUrlsForSDKResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadUrlsForSDKResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadUrlsForSDKResponse) SetHeaders(v map[string]*string) *DescribeDownloadUrlsForSDKResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadUrlsForSDKResponse) SetBody(v *DescribeDownloadUrlsForSDKResponseBody) *DescribeDownloadUrlsForSDKResponse {
	s.Body = v
	return s
}

type DescribeFlexAccConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeFlexAccConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccConfigRequest) SetSourceIp(v string) *DescribeFlexAccConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexAccConfigRequest) SetEsnBizId(v int64) *DescribeFlexAccConfigRequest {
	s.EsnBizId = &v
	return s
}

type DescribeFlexAccConfigResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlexAccConfig *DescribeFlexAccConfigResponseBodyFlexAccConfig `json:"FlexAccConfig,omitempty" xml:"FlexAccConfig,omitempty" type:"Struct"`
	PromptInfo    map[string]interface{}                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeFlexAccConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccConfigResponseBody) SetRequestId(v string) *DescribeFlexAccConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexAccConfigResponseBody) SetFlexAccConfig(v *DescribeFlexAccConfigResponseBodyFlexAccConfig) *DescribeFlexAccConfigResponseBody {
	s.FlexAccConfig = v
	return s
}

func (s *DescribeFlexAccConfigResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexAccConfigResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeFlexAccConfigResponseBodyFlexAccConfig struct {
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TcpPorts *string `json:"TcpPorts,omitempty" xml:"TcpPorts,omitempty"`
	UdpPorts *string `json:"UdpPorts,omitempty" xml:"UdpPorts,omitempty"`
}

func (s DescribeFlexAccConfigResponseBodyFlexAccConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccConfigResponseBodyFlexAccConfig) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccConfigResponseBodyFlexAccConfig) SetStatus(v int32) *DescribeFlexAccConfigResponseBodyFlexAccConfig {
	s.Status = &v
	return s
}

func (s *DescribeFlexAccConfigResponseBodyFlexAccConfig) SetTcpPorts(v string) *DescribeFlexAccConfigResponseBodyFlexAccConfig {
	s.TcpPorts = &v
	return s
}

func (s *DescribeFlexAccConfigResponseBodyFlexAccConfig) SetUdpPorts(v string) *DescribeFlexAccConfigResponseBodyFlexAccConfig {
	s.UdpPorts = &v
	return s
}

type DescribeFlexAccConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexAccConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexAccConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccConfigResponse) SetHeaders(v map[string]*string) *DescribeFlexAccConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexAccConfigResponse) SetBody(v *DescribeFlexAccConfigResponseBody) *DescribeFlexAccConfigResponse {
	s.Body = v
	return s
}

type DescribeFlexAccFlowRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BizId      *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
}

func (s DescribeFlexAccFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFlowRequest) SetSourceIp(v string) *DescribeFlexAccFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexAccFlowRequest) SetBizId(v int64) *DescribeFlexAccFlowRequest {
	s.BizId = &v
	return s
}

func (s *DescribeFlexAccFlowRequest) SetBeginTime(v int64) *DescribeFlexAccFlowRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeFlexAccFlowRequest) SetEndTime(v int64) *DescribeFlexAccFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFlexAccFlowRequest) SetInterval(v int64) *DescribeFlexAccFlowRequest {
	s.Interval = &v
	return s
}

func (s *DescribeFlexAccFlowRequest) SetApiVersion(v string) *DescribeFlexAccFlowRequest {
	s.ApiVersion = &v
	return s
}

type DescribeFlexAccFlowResponseBody struct {
	FlowData   *DescribeFlexAccFlowResponseBodyFlowData `json:"FlowData,omitempty" xml:"FlowData,omitempty" type:"Struct"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                   `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeFlexAccFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFlowResponseBody) SetFlowData(v *DescribeFlexAccFlowResponseBodyFlowData) *DescribeFlexAccFlowResponseBody {
	s.FlowData = v
	return s
}

func (s *DescribeFlexAccFlowResponseBody) SetRequestId(v string) *DescribeFlexAccFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexAccFlowResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexAccFlowResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeFlexAccFlowResponseBodyFlowData struct {
	Items          []*DescribeFlexAccFlowResponseBodyFlowDataItems   `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Categories     []*string                                         `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	TimeScope      *DescribeFlexAccFlowResponseBodyFlowDataTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Struct"`
	CategoriesType *string                                           `json:"CategoriesType,omitempty" xml:"CategoriesType,omitempty"`
}

func (s DescribeFlexAccFlowResponseBodyFlowData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFlowResponseBodyFlowData) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFlowResponseBodyFlowData) SetItems(v []*DescribeFlexAccFlowResponseBodyFlowDataItems) *DescribeFlexAccFlowResponseBodyFlowData {
	s.Items = v
	return s
}

func (s *DescribeFlexAccFlowResponseBodyFlowData) SetCategories(v []*string) *DescribeFlexAccFlowResponseBodyFlowData {
	s.Categories = v
	return s
}

func (s *DescribeFlexAccFlowResponseBodyFlowData) SetTimeScope(v *DescribeFlexAccFlowResponseBodyFlowDataTimeScope) *DescribeFlexAccFlowResponseBodyFlowData {
	s.TimeScope = v
	return s
}

func (s *DescribeFlexAccFlowResponseBodyFlowData) SetCategoriesType(v string) *DescribeFlexAccFlowResponseBodyFlowData {
	s.CategoriesType = &v
	return s
}

type DescribeFlexAccFlowResponseBodyFlowDataItems struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlexAccFlowResponseBodyFlowDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFlowResponseBodyFlowDataItems) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFlowResponseBodyFlowDataItems) SetValues(v []*string) *DescribeFlexAccFlowResponseBodyFlowDataItems {
	s.Values = v
	return s
}

func (s *DescribeFlexAccFlowResponseBodyFlowDataItems) SetName(v string) *DescribeFlexAccFlowResponseBodyFlowDataItems {
	s.Name = &v
	return s
}

type DescribeFlexAccFlowResponseBodyFlowDataTimeScope struct {
	Start    *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeFlexAccFlowResponseBodyFlowDataTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFlowResponseBodyFlowDataTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFlowResponseBodyFlowDataTimeScope) SetStart(v int64) *DescribeFlexAccFlowResponseBodyFlowDataTimeScope {
	s.Start = &v
	return s
}

func (s *DescribeFlexAccFlowResponseBodyFlowDataTimeScope) SetInterval(v int64) *DescribeFlexAccFlowResponseBodyFlowDataTimeScope {
	s.Interval = &v
	return s
}

type DescribeFlexAccFlowResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexAccFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexAccFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFlowResponse) SetHeaders(v map[string]*string) *DescribeFlexAccFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexAccFlowResponse) SetBody(v *DescribeFlexAccFlowResponseBody) *DescribeFlexAccFlowResponse {
	s.Body = v
	return s
}

type DescribeFlexAccFwdRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFlexAccFwdRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFwdRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFwdRulesRequest) SetSourceIp(v string) *DescribeFlexAccFwdRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexAccFwdRulesRequest) SetEsnBizId(v int64) *DescribeFlexAccFwdRulesRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeFlexAccFwdRulesRequest) SetKeyword(v string) *DescribeFlexAccFwdRulesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeFlexAccFwdRulesRequest) SetPage(v int32) *DescribeFlexAccFwdRulesRequest {
	s.Page = &v
	return s
}

func (s *DescribeFlexAccFwdRulesRequest) SetPageSize(v int32) *DescribeFlexAccFwdRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeFlexAccFwdRulesResponseBody struct {
	FlexAccFwdRules []*DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules `json:"FlexAccFwdRules,omitempty" xml:"FlexAccFwdRules,omitempty" type:"Repeated"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total           *int64                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo      map[string]interface{}                                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeFlexAccFwdRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFwdRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFwdRulesResponseBody) SetFlexAccFwdRules(v []*DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) *DescribeFlexAccFwdRulesResponseBody {
	s.FlexAccFwdRules = v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBody) SetRequestId(v string) *DescribeFlexAccFwdRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBody) SetTotal(v int64) *DescribeFlexAccFwdRulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexAccFwdRulesResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules struct {
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Protocol     *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	AccType      *int32  `json:"AccType,omitempty" xml:"AccType,omitempty"`
	IpList       *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	DomainList   *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
}

func (s DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetMasterIpList(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.MasterIpList = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetRemark(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.Remark = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetSlaveIpList(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.SlaveIpList = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetIdentity(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.Identity = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetProtocol(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.Protocol = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetAccType(v int32) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.AccType = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetIpList(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.IpList = &v
	return s
}

func (s *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules) SetDomainList(v string) *DescribeFlexAccFwdRulesResponseBodyFlexAccFwdRules {
	s.DomainList = &v
	return s
}

type DescribeFlexAccFwdRulesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexAccFwdRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexAccFwdRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexAccFwdRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexAccFwdRulesResponse) SetHeaders(v map[string]*string) *DescribeFlexAccFwdRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexAccFwdRulesResponse) SetBody(v *DescribeFlexAccFwdRulesResponseBody) *DescribeFlexAccFwdRulesResponse {
	s.Body = v
	return s
}

type DescribeFlexBackupGroupsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeFlexBackupGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexBackupGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexBackupGroupsRequest) SetSourceIp(v string) *DescribeFlexBackupGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexBackupGroupsRequest) SetLang(v string) *DescribeFlexBackupGroupsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlexBackupGroupsRequest) SetBizId(v int64) *DescribeFlexBackupGroupsRequest {
	s.BizId = &v
	return s
}

func (s *DescribeFlexBackupGroupsRequest) SetGroupId(v string) *DescribeFlexBackupGroupsRequest {
	s.GroupId = &v
	return s
}

type DescribeFlexBackupGroupsResponseBody struct {
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo      map[string]interface{}                               `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	AllBackupGroups *DescribeFlexBackupGroupsResponseBodyAllBackupGroups `json:"AllBackupGroups,omitempty" xml:"AllBackupGroups,omitempty" type:"Struct"`
}

func (s DescribeFlexBackupGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexBackupGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexBackupGroupsResponseBody) SetRequestId(v string) *DescribeFlexBackupGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexBackupGroupsResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexBackupGroupsResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeFlexBackupGroupsResponseBody) SetAllBackupGroups(v *DescribeFlexBackupGroupsResponseBodyAllBackupGroups) *DescribeFlexBackupGroupsResponseBody {
	s.AllBackupGroups = v
	return s
}

type DescribeFlexBackupGroupsResponseBodyAllBackupGroups struct {
	SharedGroups []*DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups `json:"SharedGroups,omitempty" xml:"SharedGroups,omitempty" type:"Repeated"`
	BackupGroups []*DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups `json:"BackupGroups,omitempty" xml:"BackupGroups,omitempty" type:"Repeated"`
}

func (s DescribeFlexBackupGroupsResponseBodyAllBackupGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexBackupGroupsResponseBodyAllBackupGroups) GoString() string {
	return s.String()
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroups) SetSharedGroups(v []*DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups) *DescribeFlexBackupGroupsResponseBodyAllBackupGroups {
	s.SharedGroups = v
	return s
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroups) SetBackupGroups(v []*DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups) *DescribeFlexBackupGroupsResponseBodyAllBackupGroups {
	s.BackupGroups = v
	return s
}

type DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups struct {
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SpecDesc *string `json:"SpecDesc,omitempty" xml:"SpecDesc,omitempty"`
}

func (s DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups) GoString() string {
	return s.String()
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups) SetStatus(v int32) *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups {
	s.Status = &v
	return s
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups) SetGroupId(v string) *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups) SetSpecDesc(v string) *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsSharedGroups {
	s.SpecDesc = &v
	return s
}

type DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups struct {
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups) GoString() string {
	return s.String()
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups) SetStatus(v int32) *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups {
	s.Status = &v
	return s
}

func (s *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups) SetGroupId(v string) *DescribeFlexBackupGroupsResponseBodyAllBackupGroupsBackupGroups {
	s.GroupId = &v
	return s
}

type DescribeFlexBackupGroupsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexBackupGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexBackupGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexBackupGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexBackupGroupsResponse) SetHeaders(v map[string]*string) *DescribeFlexBackupGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexBackupGroupsResponse) SetBody(v *DescribeFlexBackupGroupsResponseBody) *DescribeFlexBackupGroupsResponse {
	s.Body = v
	return s
}

type DescribeFlexConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeFlexConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexConfigRequest) SetSourceIp(v string) *DescribeFlexConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexConfigRequest) SetLang(v string) *DescribeFlexConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlexConfigRequest) SetEsnBizId(v int64) *DescribeFlexConfigRequest {
	s.EsnBizId = &v
	return s
}

type DescribeFlexConfigResponseBody struct {
	FlexConfig *DescribeFlexConfigResponseBodyFlexConfig `json:"FlexConfig,omitempty" xml:"FlexConfig,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                    `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeFlexConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexConfigResponseBody) SetFlexConfig(v *DescribeFlexConfigResponseBodyFlexConfig) *DescribeFlexConfigResponseBody {
	s.FlexConfig = v
	return s
}

func (s *DescribeFlexConfigResponseBody) SetRequestId(v string) *DescribeFlexConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexConfigResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexConfigResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeFlexConfigResponseBodyFlexConfig struct {
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	LinkType *int32  `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	Ports    *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
	Rate     *int32  `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s DescribeFlexConfigResponseBodyFlexConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexConfigResponseBodyFlexConfig) GoString() string {
	return s.String()
}

func (s *DescribeFlexConfigResponseBodyFlexConfig) SetStatus(v int32) *DescribeFlexConfigResponseBodyFlexConfig {
	s.Status = &v
	return s
}

func (s *DescribeFlexConfigResponseBodyFlexConfig) SetLinkType(v int32) *DescribeFlexConfigResponseBodyFlexConfig {
	s.LinkType = &v
	return s
}

func (s *DescribeFlexConfigResponseBodyFlexConfig) SetPorts(v string) *DescribeFlexConfigResponseBodyFlexConfig {
	s.Ports = &v
	return s
}

func (s *DescribeFlexConfigResponseBodyFlexConfig) SetRate(v int32) *DescribeFlexConfigResponseBodyFlexConfig {
	s.Rate = &v
	return s
}

type DescribeFlexConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexConfigResponse) SetHeaders(v map[string]*string) *DescribeFlexConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexConfigResponse) SetBody(v *DescribeFlexConfigResponseBody) *DescribeFlexConfigResponse {
	s.Body = v
	return s
}

type DescribeFlexFwdRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeFlexFwdRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexFwdRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexFwdRulesRequest) SetSourceIp(v string) *DescribeFlexFwdRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexFwdRulesRequest) SetLang(v string) *DescribeFlexFwdRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlexFwdRulesRequest) SetEsnBizId(v int64) *DescribeFlexFwdRulesRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeFlexFwdRulesRequest) SetKeyword(v string) *DescribeFlexFwdRulesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeFlexFwdRulesRequest) SetPage(v int32) *DescribeFlexFwdRulesRequest {
	s.Page = &v
	return s
}

func (s *DescribeFlexFwdRulesRequest) SetPageSize(v int32) *DescribeFlexFwdRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeFlexFwdRulesResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo   map[string]interface{}                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	FlexFwdRules []*DescribeFlexFwdRulesResponseBodyFlexFwdRules `json:"FlexFwdRules,omitempty" xml:"FlexFwdRules,omitempty" type:"Repeated"`
}

func (s DescribeFlexFwdRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexFwdRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexFwdRulesResponseBody) SetRequestId(v string) *DescribeFlexFwdRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexFwdRulesResponseBody) SetTotal(v int64) *DescribeFlexFwdRulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeFlexFwdRulesResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexFwdRulesResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeFlexFwdRulesResponseBody) SetFlexFwdRules(v []*DescribeFlexFwdRulesResponseBodyFlexFwdRules) *DescribeFlexFwdRulesResponseBody {
	s.FlexFwdRules = v
	return s
}

type DescribeFlexFwdRulesResponseBodyFlexFwdRules struct {
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Protocol     *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
}

func (s DescribeFlexFwdRulesResponseBodyFlexFwdRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexFwdRulesResponseBodyFlexFwdRules) GoString() string {
	return s.String()
}

func (s *DescribeFlexFwdRulesResponseBodyFlexFwdRules) SetMasterIpList(v string) *DescribeFlexFwdRulesResponseBodyFlexFwdRules {
	s.MasterIpList = &v
	return s
}

func (s *DescribeFlexFwdRulesResponseBodyFlexFwdRules) SetRemark(v string) *DescribeFlexFwdRulesResponseBodyFlexFwdRules {
	s.Remark = &v
	return s
}

func (s *DescribeFlexFwdRulesResponseBodyFlexFwdRules) SetSlaveIpList(v string) *DescribeFlexFwdRulesResponseBodyFlexFwdRules {
	s.SlaveIpList = &v
	return s
}

func (s *DescribeFlexFwdRulesResponseBodyFlexFwdRules) SetProtocol(v string) *DescribeFlexFwdRulesResponseBodyFlexFwdRules {
	s.Protocol = &v
	return s
}

func (s *DescribeFlexFwdRulesResponseBodyFlexFwdRules) SetIdentity(v string) *DescribeFlexFwdRulesResponseBodyFlexFwdRules {
	s.Identity = &v
	return s
}

type DescribeFlexFwdRulesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexFwdRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexFwdRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexFwdRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexFwdRulesResponse) SetHeaders(v map[string]*string) *DescribeFlexFwdRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexFwdRulesResponse) SetBody(v *DescribeFlexFwdRulesResponseBody) *DescribeFlexFwdRulesResponse {
	s.Body = v
	return s
}

type DescribeFlexSechedPolicyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeFlexSechedPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexSechedPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlexSechedPolicyRequest) SetSourceIp(v string) *DescribeFlexSechedPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlexSechedPolicyRequest) SetLang(v string) *DescribeFlexSechedPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlexSechedPolicyRequest) SetEsnBizId(v int64) *DescribeFlexSechedPolicyRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeFlexSechedPolicyRequest) SetGroupId(v string) *DescribeFlexSechedPolicyRequest {
	s.GroupId = &v
	return s
}

type DescribeFlexSechedPolicyResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo       map[string]interface{}                                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	FlexSechedPolicy *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy `json:"FlexSechedPolicy,omitempty" xml:"FlexSechedPolicy,omitempty" type:"Struct"`
}

func (s DescribeFlexSechedPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexSechedPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlexSechedPolicyResponseBody) SetRequestId(v string) *DescribeFlexSechedPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlexSechedPolicyResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFlexSechedPolicyResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeFlexSechedPolicyResponseBody) SetFlexSechedPolicy(v *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy) *DescribeFlexSechedPolicyResponseBody {
	s.FlexSechedPolicy = v
	return s
}

type DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy struct {
	GreyGroupPolicy   *int32 `json:"GreyGroupPolicy,omitempty" xml:"GreyGroupPolicy,omitempty"`
	GroupInnerPolicy  *int32 `json:"GroupInnerPolicy,omitempty" xml:"GroupInnerPolicy,omitempty"`
	BackupGroupPolicy *int32 `json:"BackupGroupPolicy,omitempty" xml:"BackupGroupPolicy,omitempty"`
}

func (s DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy) GoString() string {
	return s.String()
}

func (s *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy) SetGreyGroupPolicy(v int32) *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy {
	s.GreyGroupPolicy = &v
	return s
}

func (s *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy) SetGroupInnerPolicy(v int32) *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy {
	s.GroupInnerPolicy = &v
	return s
}

func (s *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy) SetBackupGroupPolicy(v int32) *DescribeFlexSechedPolicyResponseBodyFlexSechedPolicy {
	s.BackupGroupPolicy = &v
	return s
}

type DescribeFlexSechedPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlexSechedPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlexSechedPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlexSechedPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlexSechedPolicyResponse) SetHeaders(v map[string]*string) *DescribeFlexSechedPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlexSechedPolicyResponse) SetBody(v *DescribeFlexSechedPolicyResponseBody) *DescribeFlexSechedPolicyResponse {
	s.Body = v
	return s
}

type DescribeFlowReportRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EsnAppId  *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeFlowReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowReportRequest) SetSourceIp(v string) *DescribeFlowReportRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFlowReportRequest) SetLang(v string) *DescribeFlowReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFlowReportRequest) SetInterval(v int32) *DescribeFlowReportRequest {
	s.Interval = &v
	return s
}

func (s *DescribeFlowReportRequest) SetStartTime(v int64) *DescribeFlowReportRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowReportRequest) SetEndTime(v int64) *DescribeFlowReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowReportRequest) SetEsnAppId(v string) *DescribeFlowReportRequest {
	s.EsnAppId = &v
	return s
}

func (s *DescribeFlowReportRequest) SetIp(v string) *DescribeFlowReportRequest {
	s.Ip = &v
	return s
}

type DescribeFlowReportResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowReport []*DescribeFlowReportResponseBodyFlowReport `json:"FlowReport,omitempty" xml:"FlowReport,omitempty" type:"Repeated"`
}

func (s DescribeFlowReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowReportResponseBody) SetRequestId(v string) *DescribeFlowReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowReportResponseBody) SetFlowReport(v []*DescribeFlowReportResponseBodyFlowReport) *DescribeFlowReportResponseBody {
	s.FlowReport = v
	return s
}

type DescribeFlowReportResponseBodyFlowReport struct {
	Index      *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	ActConns   *int64 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	InactConns *int64 `json:"InactConns,omitempty" xml:"InactConns,omitempty"`
	InBps      *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	OutBps     *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
}

func (s DescribeFlowReportResponseBodyFlowReport) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowReportResponseBodyFlowReport) GoString() string {
	return s.String()
}

func (s *DescribeFlowReportResponseBodyFlowReport) SetIndex(v int64) *DescribeFlowReportResponseBodyFlowReport {
	s.Index = &v
	return s
}

func (s *DescribeFlowReportResponseBodyFlowReport) SetActConns(v int64) *DescribeFlowReportResponseBodyFlowReport {
	s.ActConns = &v
	return s
}

func (s *DescribeFlowReportResponseBodyFlowReport) SetInactConns(v int64) *DescribeFlowReportResponseBodyFlowReport {
	s.InactConns = &v
	return s
}

func (s *DescribeFlowReportResponseBodyFlowReport) SetInBps(v int64) *DescribeFlowReportResponseBodyFlowReport {
	s.InBps = &v
	return s
}

func (s *DescribeFlowReportResponseBodyFlowReport) SetOutBps(v int64) *DescribeFlowReportResponseBodyFlowReport {
	s.OutBps = &v
	return s
}

type DescribeFlowReportResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowReportResponse) SetHeaders(v map[string]*string) *DescribeFlowReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowReportResponse) SetBody(v *DescribeFlowReportResponseBody) *DescribeFlowReportResponse {
	s.Body = v
	return s
}

type DescribeFullNodesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InUse    *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
}

func (s DescribeFullNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesRequest) SetSourceIp(v string) *DescribeFullNodesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFullNodesRequest) SetLang(v string) *DescribeFullNodesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFullNodesRequest) SetAppId(v int64) *DescribeFullNodesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFullNodesRequest) SetGroupId(v string) *DescribeFullNodesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeFullNodesRequest) SetInUse(v int32) *DescribeFullNodesRequest {
	s.InUse = &v
	return s
}

type DescribeFullNodesResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	NodeList   []*DescribeFullNodesResponseBodyNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	PromptInfo map[string]interface{}                   `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeFullNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesResponseBody) SetRequestId(v string) *DescribeFullNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFullNodesResponseBody) SetTotal(v int64) *DescribeFullNodesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeFullNodesResponseBody) SetNodeList(v []*DescribeFullNodesResponseBodyNodeList) *DescribeFullNodesResponseBody {
	s.NodeList = v
	return s
}

func (s *DescribeFullNodesResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFullNodesResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeFullNodesResponseBodyNodeList struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeFullNodesResponseBodyNodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesResponseBodyNodeList) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesResponseBodyNodeList) SetType(v int32) *DescribeFullNodesResponseBodyNodeList {
	s.Type = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetBizId(v string) *DescribeFullNodesResponseBodyNodeList {
	s.BizId = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetIp(v string) *DescribeFullNodesResponseBodyNodeList {
	s.Ip = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetRegionId(v string) *DescribeFullNodesResponseBodyNodeList {
	s.RegionId = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetAppName(v string) *DescribeFullNodesResponseBodyNodeList {
	s.AppName = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetRegionName(v string) *DescribeFullNodesResponseBodyNodeList {
	s.RegionName = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetGroupId(v string) *DescribeFullNodesResponseBodyNodeList {
	s.GroupId = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetAppId(v string) *DescribeFullNodesResponseBodyNodeList {
	s.AppId = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetInUse(v int32) *DescribeFullNodesResponseBodyNodeList {
	s.InUse = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetBizName(v string) *DescribeFullNodesResponseBodyNodeList {
	s.BizName = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetDdosStatus(v int32) *DescribeFullNodesResponseBodyNodeList {
	s.DdosStatus = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetIsDisabled(v bool) *DescribeFullNodesResponseBodyNodeList {
	s.IsDisabled = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetThreshold(v int32) *DescribeFullNodesResponseBodyNodeList {
	s.Threshold = &v
	return s
}

func (s *DescribeFullNodesResponseBodyNodeList) SetEsnBizId(v string) *DescribeFullNodesResponseBodyNodeList {
	s.EsnBizId = &v
	return s
}

type DescribeFullNodesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFullNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFullNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesResponse) SetHeaders(v map[string]*string) *DescribeFullNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFullNodesResponse) SetBody(v *DescribeFullNodesResponseBody) *DescribeFullNodesResponse {
	s.Body = v
	return s
}

type DescribeFullNodesSummaryRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeFullNodesSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesSummaryRequest) SetSourceIp(v string) *DescribeFullNodesSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFullNodesSummaryRequest) SetLang(v string) *DescribeFullNodesSummaryRequest {
	s.Lang = &v
	return s
}

type DescribeFullNodesSummaryResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullNodesSummary *DescribeFullNodesSummaryResponseBodyFullNodesSummary `json:"FullNodesSummary,omitempty" xml:"FullNodesSummary,omitempty" type:"Struct"`
	PromptInfo       map[string]interface{}                                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeFullNodesSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesSummaryResponseBody) SetRequestId(v string) *DescribeFullNodesSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBody) SetFullNodesSummary(v *DescribeFullNodesSummaryResponseBodyFullNodesSummary) *DescribeFullNodesSummaryResponseBody {
	s.FullNodesSummary = v
	return s
}

func (s *DescribeFullNodesSummaryResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeFullNodesSummaryResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeFullNodesSummaryResponseBodyFullNodesSummary struct {
	CleanCount    *int32                                                               `json:"CleanCount,omitempty" xml:"CleanCount,omitempty"`
	HolingNodes   []*DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes   `json:"HolingNodes,omitempty" xml:"HolingNodes,omitempty" type:"Repeated"`
	HoleCount     *int32                                                               `json:"HoleCount,omitempty" xml:"HoleCount,omitempty"`
	CleaningNodes []*DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes `json:"CleaningNodes,omitempty" xml:"CleaningNodes,omitempty" type:"Repeated"`
	UsedCount     *int32                                                               `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	UnUsedCount   *int32                                                               `json:"UnUsedCount,omitempty" xml:"UnUsedCount,omitempty"`
	TotalCount    *int32                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFullNodesSummaryResponseBodyFullNodesSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesSummaryResponseBodyFullNodesSummary) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetCleanCount(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.CleanCount = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetHolingNodes(v []*DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.HolingNodes = v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetHoleCount(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.HoleCount = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetCleaningNodes(v []*DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.CleaningNodes = v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetUsedCount(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.UsedCount = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetUnUsedCount(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.UnUsedCount = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummary) SetTotalCount(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummary {
	s.TotalCount = &v
	return s
}

type DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetType(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.Type = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetBizId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.BizId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetIp(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.Ip = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetRegionId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetAppName(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.AppName = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetRegionName(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.RegionName = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetGroupId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetAppId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.AppId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetInUse(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.InUse = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetBizName(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.BizName = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetDdosStatus(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.DdosStatus = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetIsDisabled(v bool) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.IsDisabled = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetThreshold(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.Threshold = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes) SetEsnBizId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryHolingNodes {
	s.EsnBizId = &v
	return s
}

type DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetType(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.Type = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetBizId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.BizId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetIp(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.Ip = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetRegionId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetAppName(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.AppName = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetRegionName(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.RegionName = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetGroupId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetAppId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.AppId = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetInUse(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.InUse = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetBizName(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.BizName = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetDdosStatus(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.DdosStatus = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetIsDisabled(v bool) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.IsDisabled = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetThreshold(v int32) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.Threshold = &v
	return s
}

func (s *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes) SetEsnBizId(v string) *DescribeFullNodesSummaryResponseBodyFullNodesSummaryCleaningNodes {
	s.EsnBizId = &v
	return s
}

type DescribeFullNodesSummaryResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFullNodesSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFullNodesSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFullNodesSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeFullNodesSummaryResponse) SetHeaders(v map[string]*string) *DescribeFullNodesSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeFullNodesSummaryResponse) SetBody(v *DescribeFullNodesSummaryResponseBody) *DescribeFullNodesSummaryResponse {
	s.Body = v
	return s
}

type DescribeGfResSummaryRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGfResSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGfResSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeGfResSummaryRequest) SetSourceIp(v string) *DescribeGfResSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGfResSummaryRequest) SetLang(v string) *DescribeGfResSummaryRequest {
	s.Lang = &v
	return s
}

type DescribeGfResSummaryResponseBody struct {
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullNodesSummary *DescribeGfResSummaryResponseBodyFullNodesSummary `json:"FullNodesSummary,omitempty" xml:"FullNodesSummary,omitempty" type:"Struct"`
	PromptInfo       map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeGfResSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGfResSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGfResSummaryResponseBody) SetRequestId(v string) *DescribeGfResSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGfResSummaryResponseBody) SetFullNodesSummary(v *DescribeGfResSummaryResponseBodyFullNodesSummary) *DescribeGfResSummaryResponseBody {
	s.FullNodesSummary = v
	return s
}

func (s *DescribeGfResSummaryResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeGfResSummaryResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeGfResSummaryResponseBodyFullNodesSummary struct {
	Nodes       []*DescribeGfResSummaryResponseBodyFullNodesSummaryNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	CleanCount  *int32                                                   `json:"CleanCount,omitempty" xml:"CleanCount,omitempty"`
	HoleCount   *int32                                                   `json:"HoleCount,omitempty" xml:"HoleCount,omitempty"`
	UsedCount   *int32                                                   `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	UnUsedCount *int32                                                   `json:"UnUsedCount,omitempty" xml:"UnUsedCount,omitempty"`
	TotalCount  *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGfResSummaryResponseBodyFullNodesSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeGfResSummaryResponseBodyFullNodesSummary) GoString() string {
	return s.String()
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummary) SetNodes(v []*DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) *DescribeGfResSummaryResponseBodyFullNodesSummary {
	s.Nodes = v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummary) SetCleanCount(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummary {
	s.CleanCount = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummary) SetHoleCount(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummary {
	s.HoleCount = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummary) SetUsedCount(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummary {
	s.UsedCount = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummary) SetUnUsedCount(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummary {
	s.UnUsedCount = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummary) SetTotalCount(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummary {
	s.TotalCount = &v
	return s
}

type DescribeGfResSummaryResponseBodyFullNodesSummaryNodes struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) GoString() string {
	return s.String()
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetType(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.Type = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetBizId(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizId = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetIp(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.Ip = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetRegionId(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetAppName(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppName = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetRegionName(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionName = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetGroupId(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetAppId(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppId = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetInUse(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.InUse = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetBizName(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizName = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetDdosStatus(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.DdosStatus = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetIsDisabled(v bool) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.IsDisabled = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetThreshold(v int32) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.Threshold = &v
	return s
}

func (s *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes) SetEsnBizId(v string) *DescribeGfResSummaryResponseBodyFullNodesSummaryNodes {
	s.EsnBizId = &v
	return s
}

type DescribeGfResSummaryResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGfResSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGfResSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGfResSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeGfResSummaryResponse) SetHeaders(v map[string]*string) *DescribeGfResSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeGfResSummaryResponse) SetBody(v *DescribeGfResSummaryResponseBody) *DescribeGfResSummaryResponse {
	s.Body = v
	return s
}

type DescribeGroupListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupListRequest) SetSourceIp(v string) *DescribeGroupListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGroupListRequest) SetLang(v string) *DescribeGroupListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupListRequest) SetBizId(v int64) *DescribeGroupListRequest {
	s.BizId = &v
	return s
}

type DescribeGroupListResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupList  []*DescribeGroupListResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	Total      *int64                                    `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                    `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupListResponseBody) SetRequestId(v string) *DescribeGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupListResponseBody) SetGroupList(v []*DescribeGroupListResponseBodyGroupList) *DescribeGroupListResponseBody {
	s.GroupList = v
	return s
}

func (s *DescribeGroupListResponseBody) SetTotal(v int64) *DescribeGroupListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeGroupListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeGroupListResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeGroupListResponseBodyGroupList struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	IsDnsEnabled *bool   `json:"IsDnsEnabled,omitempty" xml:"IsDnsEnabled,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BlackIp      *int32  `json:"BlackIp,omitempty" xml:"BlackIp,omitempty"`
	NormalIp     *int32  `json:"NormalIp,omitempty" xml:"NormalIp,omitempty"`
	CleanIp      *int32  `json:"CleanIp,omitempty" xml:"CleanIp,omitempty"`
	TotalIp      *int32  `json:"TotalIp,omitempty" xml:"TotalIp,omitempty"`
	IsDisabled   *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	DnsStatus    *string `json:"DnsStatus,omitempty" xml:"DnsStatus,omitempty"`
	GroupDesc    *string `json:"GroupDesc,omitempty" xml:"GroupDesc,omitempty"`
}

func (s DescribeGroupListResponseBodyGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupListResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *DescribeGroupListResponseBodyGroupList) SetStatus(v string) *DescribeGroupListResponseBodyGroupList {
	s.Status = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetIsDnsEnabled(v bool) *DescribeGroupListResponseBodyGroupList {
	s.IsDnsEnabled = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetGroupId(v string) *DescribeGroupListResponseBodyGroupList {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetBlackIp(v int32) *DescribeGroupListResponseBodyGroupList {
	s.BlackIp = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetNormalIp(v int32) *DescribeGroupListResponseBodyGroupList {
	s.NormalIp = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetCleanIp(v int32) *DescribeGroupListResponseBodyGroupList {
	s.CleanIp = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetTotalIp(v int32) *DescribeGroupListResponseBodyGroupList {
	s.TotalIp = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetIsDisabled(v bool) *DescribeGroupListResponseBodyGroupList {
	s.IsDisabled = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetDnsStatus(v string) *DescribeGroupListResponseBodyGroupList {
	s.DnsStatus = &v
	return s
}

func (s *DescribeGroupListResponseBodyGroupList) SetGroupDesc(v string) *DescribeGroupListResponseBodyGroupList {
	s.GroupDesc = &v
	return s
}

type DescribeGroupListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupListResponse) SetHeaders(v map[string]*string) *DescribeGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupListResponse) SetBody(v *DescribeGroupListResponseBody) *DescribeGroupListResponse {
	s.Body = v
	return s
}

type DescribeGroupSimpleListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeGroupSimpleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupSimpleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupSimpleListRequest) SetSourceIp(v string) *DescribeGroupSimpleListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGroupSimpleListRequest) SetLang(v string) *DescribeGroupSimpleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupSimpleListRequest) SetBizId(v int64) *DescribeGroupSimpleListRequest {
	s.BizId = &v
	return s
}

type DescribeGroupSimpleListResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupList  []*DescribeGroupSimpleListResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	Total      *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeGroupSimpleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupSimpleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupSimpleListResponseBody) SetRequestId(v string) *DescribeGroupSimpleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBody) SetGroupList(v []*DescribeGroupSimpleListResponseBodyGroupList) *DescribeGroupSimpleListResponseBody {
	s.GroupList = v
	return s
}

func (s *DescribeGroupSimpleListResponseBody) SetTotal(v int64) *DescribeGroupSimpleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeGroupSimpleListResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeGroupSimpleListResponseBodyGroupList struct {
	IsDnsEnabled *bool   `json:"IsDnsEnabled,omitempty" xml:"IsDnsEnabled,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IsDisabled   *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
}

func (s DescribeGroupSimpleListResponseBodyGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupSimpleListResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *DescribeGroupSimpleListResponseBodyGroupList) SetIsDnsEnabled(v bool) *DescribeGroupSimpleListResponseBodyGroupList {
	s.IsDnsEnabled = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBodyGroupList) SetGroupId(v string) *DescribeGroupSimpleListResponseBodyGroupList {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBodyGroupList) SetGroupName(v string) *DescribeGroupSimpleListResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBodyGroupList) SetAppId(v string) *DescribeGroupSimpleListResponseBodyGroupList {
	s.AppId = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBodyGroupList) SetBizId(v string) *DescribeGroupSimpleListResponseBodyGroupList {
	s.BizId = &v
	return s
}

func (s *DescribeGroupSimpleListResponseBodyGroupList) SetIsDisabled(v bool) *DescribeGroupSimpleListResponseBodyGroupList {
	s.IsDisabled = &v
	return s
}

type DescribeGroupSimpleListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupSimpleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupSimpleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupSimpleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupSimpleListResponse) SetHeaders(v map[string]*string) *DescribeGroupSimpleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupSimpleListResponse) SetBody(v *DescribeGroupSimpleListResponseBody) *DescribeGroupSimpleListResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetSourceIp(v string) *DescribeInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceRequest) SetLang(v string) *DescribeInstanceRequest {
	s.Lang = &v
	return s
}

type DescribeInstanceResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo   map[string]interface{}                    `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	InstanceInfo *DescribeInstanceResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeInstanceResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceInfo(v *DescribeInstanceResponseBodyInstanceInfo) *DescribeInstanceResponseBody {
	s.InstanceInfo = v
	return s
}

type DescribeInstanceResponseBodyInstanceInfo struct {
	EndTime     *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status      *int32                 `json:"Status,omitempty" xml:"Status,omitempty"`
	PackageCode *string                `json:"PackageCode,omitempty" xml:"PackageCode,omitempty"`
	Region      *string                `json:"Region,omitempty" xml:"Region,omitempty"`
	SpecMap     map[string]interface{} `json:"SpecMap,omitempty" xml:"SpecMap,omitempty"`
	InstanceId  *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceInfo) SetEndTime(v string) *DescribeInstanceResponseBodyInstanceInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceInfo) SetStatus(v int32) *DescribeInstanceResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceInfo) SetPackageCode(v string) *DescribeInstanceResponseBodyInstanceInfo {
	s.PackageCode = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceInfo) SetRegion(v string) *DescribeInstanceResponseBodyInstanceInfo {
	s.Region = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceInfo) SetSpecMap(v map[string]interface{}) *DescribeInstanceResponseBodyInstanceInfo {
	s.SpecMap = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceInfo) SetInstanceId(v string) *DescribeInstanceResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeIpFlowReportRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FPort      *int64  `json:"FPort,omitempty" xml:"FPort,omitempty"`
	Interval   *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ProtocolEx *int64  `json:"ProtocolEx,omitempty" xml:"ProtocolEx,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeIpFlowReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpFlowReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpFlowReportRequest) SetSourceIp(v string) *DescribeIpFlowReportRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIpFlowReportRequest) SetIp(v string) *DescribeIpFlowReportRequest {
	s.Ip = &v
	return s
}

func (s *DescribeIpFlowReportRequest) SetEndTime(v int64) *DescribeIpFlowReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeIpFlowReportRequest) SetFPort(v int64) *DescribeIpFlowReportRequest {
	s.FPort = &v
	return s
}

func (s *DescribeIpFlowReportRequest) SetInterval(v int64) *DescribeIpFlowReportRequest {
	s.Interval = &v
	return s
}

func (s *DescribeIpFlowReportRequest) SetProtocolEx(v int64) *DescribeIpFlowReportRequest {
	s.ProtocolEx = &v
	return s
}

func (s *DescribeIpFlowReportRequest) SetStartTime(v int64) *DescribeIpFlowReportRequest {
	s.StartTime = &v
	return s
}

type DescribeIpFlowReportResponseBody struct {
	IpFlowReportList []*DescribeIpFlowReportResponseBodyIpFlowReportList `json:"IpFlowReportList,omitempty" xml:"IpFlowReportList,omitempty" type:"Repeated"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo       map[string]interface{}                              `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeIpFlowReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpFlowReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpFlowReportResponseBody) SetIpFlowReportList(v []*DescribeIpFlowReportResponseBodyIpFlowReportList) *DescribeIpFlowReportResponseBody {
	s.IpFlowReportList = v
	return s
}

func (s *DescribeIpFlowReportResponseBody) SetRequestId(v string) *DescribeIpFlowReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpFlowReportResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeIpFlowReportResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeIpFlowReportResponseBodyIpFlowReportList struct {
	Outpps      *string `json:"Outpps,omitempty" xml:"Outpps,omitempty"`
	Cps         *string `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Inkbps      *string `json:"Inkbps,omitempty" xml:"Inkbps,omitempty"`
	Inpps       *string `json:"Inpps,omitempty" xml:"Inpps,omitempty"`
	Inbps       *string `json:"Inbps,omitempty" xml:"Inbps,omitempty"`
	Conns       *bool   `json:"Conns,omitempty" xml:"Conns,omitempty"`
	Inpkts      *string `json:"Inpkts,omitempty" xml:"Inpkts,omitempty"`
	Inbytes     *string `json:"Inbytes,omitempty" xml:"Inbytes,omitempty"`
	Outbytes    *string `json:"Outbytes,omitempty" xml:"Outbytes,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Actconns    *string `json:"Actconns,omitempty" xml:"Actconns,omitempty"`
	Outbps      *string `json:"Outbps,omitempty" xml:"Outbps,omitempty"`
	Outpkts     *string `json:"Outpkts,omitempty" xml:"Outpkts,omitempty"`
	Vip         *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	Qtime       *string `json:"Qtime,omitempty" xml:"Qtime,omitempty"`
	Outkbps     *string `json:"Outkbps,omitempty" xml:"Outkbps,omitempty"`
	Inactconns  *string `json:"Inactconns,omitempty" xml:"Inactconns,omitempty"`
}

func (s DescribeIpFlowReportResponseBodyIpFlowReportList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpFlowReportResponseBodyIpFlowReportList) GoString() string {
	return s.String()
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetOutpps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Outpps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetCps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Cps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetInkbps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Inkbps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetInpps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Inpps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetInbps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Inbps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetConns(v bool) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Conns = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetInpkts(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Inpkts = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetInbytes(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Inbytes = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetOutbytes(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Outbytes = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetClusterName(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.ClusterName = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetActconns(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Actconns = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetOutbps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Outbps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetOutpkts(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Outpkts = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetVip(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Vip = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetQtime(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Qtime = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetOutkbps(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Outkbps = &v
	return s
}

func (s *DescribeIpFlowReportResponseBodyIpFlowReportList) SetInactconns(v string) *DescribeIpFlowReportResponseBodyIpFlowReportList {
	s.Inactconns = &v
	return s
}

type DescribeIpFlowReportResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpFlowReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpFlowReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpFlowReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpFlowReportResponse) SetHeaders(v map[string]*string) *DescribeIpFlowReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpFlowReportResponse) SetBody(v *DescribeIpFlowReportResponseBody) *DescribeIpFlowReportResponse {
	s.Body = v
	return s
}

type DescribeJianYuTestGetRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeJianYuTestGetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestGetRequest) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestGetRequest) SetSourceIp(v string) *DescribeJianYuTestGetRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeJianYuTestGetRequest) SetLang(v string) *DescribeJianYuTestGetRequest {
	s.Lang = &v
	return s
}

func (s *DescribeJianYuTestGetRequest) SetStartTime(v string) *DescribeJianYuTestGetRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeJianYuTestGetRequest) SetEndTime(v string) *DescribeJianYuTestGetRequest {
	s.EndTime = &v
	return s
}

type DescribeJianYuTestGetResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                 `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	Data       *DescribeJianYuTestGetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeJianYuTestGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestGetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestGetResponseBody) SetRequestId(v string) *DescribeJianYuTestGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJianYuTestGetResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeJianYuTestGetResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeJianYuTestGetResponseBody) SetData(v *DescribeJianYuTestGetResponseBodyData) *DescribeJianYuTestGetResponseBody {
	s.Data = v
	return s
}

type DescribeJianYuTestGetResponseBodyData struct {
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id         *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeJianYuTestGetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestGetResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestGetResponseBodyData) SetCreateTime(v int64) *DescribeJianYuTestGetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeJianYuTestGetResponseBodyData) SetId(v int32) *DescribeJianYuTestGetResponseBodyData {
	s.Id = &v
	return s
}

type DescribeJianYuTestGetResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJianYuTestGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJianYuTestGetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestGetResponse) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestGetResponse) SetHeaders(v map[string]*string) *DescribeJianYuTestGetResponse {
	s.Headers = v
	return s
}

func (s *DescribeJianYuTestGetResponse) SetBody(v *DescribeJianYuTestGetResponseBody) *DescribeJianYuTestGetResponse {
	s.Body = v
	return s
}

type DescribeJianYuTestListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeJianYuTestListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestListRequest) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestListRequest) SetSourceIp(v string) *DescribeJianYuTestListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeJianYuTestListRequest) SetLang(v string) *DescribeJianYuTestListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeJianYuTestListRequest) SetStartTime(v string) *DescribeJianYuTestListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeJianYuTestListRequest) SetEndTime(v string) *DescribeJianYuTestListRequest {
	s.EndTime = &v
	return s
}

type DescribeJianYuTestListResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                    `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	List       []*DescribeJianYuTestListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeJianYuTestListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestListResponseBody) SetRequestId(v string) *DescribeJianYuTestListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJianYuTestListResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeJianYuTestListResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeJianYuTestListResponseBody) SetList(v []*DescribeJianYuTestListResponseBodyList) *DescribeJianYuTestListResponseBody {
	s.List = v
	return s
}

type DescribeJianYuTestListResponseBodyList struct {
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id         *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeJianYuTestListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestListResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestListResponseBodyList) SetCreateTime(v int64) *DescribeJianYuTestListResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *DescribeJianYuTestListResponseBodyList) SetId(v int32) *DescribeJianYuTestListResponseBodyList {
	s.Id = &v
	return s
}

type DescribeJianYuTestListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJianYuTestListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJianYuTestListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestListResponse) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestListResponse) SetHeaders(v map[string]*string) *DescribeJianYuTestListResponse {
	s.Headers = v
	return s
}

func (s *DescribeJianYuTestListResponse) SetBody(v *DescribeJianYuTestListResponseBody) *DescribeJianYuTestListResponse {
	s.Body = v
	return s
}

type DescribeJianYuTestPaginationRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeJianYuTestPaginationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestPaginationRequest) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestPaginationRequest) SetSourceIp(v string) *DescribeJianYuTestPaginationRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeJianYuTestPaginationRequest) SetLang(v string) *DescribeJianYuTestPaginationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeJianYuTestPaginationRequest) SetStartTime(v string) *DescribeJianYuTestPaginationRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeJianYuTestPaginationRequest) SetEndTime(v string) *DescribeJianYuTestPaginationRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeJianYuTestPaginationRequest) SetPage(v string) *DescribeJianYuTestPaginationRequest {
	s.Page = &v
	return s
}

func (s *DescribeJianYuTestPaginationRequest) SetPageSize(v string) *DescribeJianYuTestPaginationRequest {
	s.PageSize = &v
	return s
}

type DescribeJianYuTestPaginationResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	List       []*DescribeJianYuTestPaginationResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeJianYuTestPaginationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestPaginationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestPaginationResponseBody) SetRequestId(v string) *DescribeJianYuTestPaginationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJianYuTestPaginationResponseBody) SetTotal(v int64) *DescribeJianYuTestPaginationResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeJianYuTestPaginationResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeJianYuTestPaginationResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeJianYuTestPaginationResponseBody) SetList(v []*DescribeJianYuTestPaginationResponseBodyList) *DescribeJianYuTestPaginationResponseBody {
	s.List = v
	return s
}

type DescribeJianYuTestPaginationResponseBodyList struct {
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id         *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeJianYuTestPaginationResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestPaginationResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestPaginationResponseBodyList) SetCreateTime(v int64) *DescribeJianYuTestPaginationResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *DescribeJianYuTestPaginationResponseBodyList) SetId(v int32) *DescribeJianYuTestPaginationResponseBodyList {
	s.Id = &v
	return s
}

type DescribeJianYuTestPaginationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJianYuTestPaginationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJianYuTestPaginationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJianYuTestPaginationResponse) GoString() string {
	return s.String()
}

func (s *DescribeJianYuTestPaginationResponse) SetHeaders(v map[string]*string) *DescribeJianYuTestPaginationResponse {
	s.Headers = v
	return s
}

func (s *DescribeJianYuTestPaginationResponse) SetBody(v *DescribeJianYuTestPaginationResponseBody) *DescribeJianYuTestPaginationResponse {
	s.Body = v
	return s
}

type DescribeL4EventsSelectiveRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BeginTime *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EsnAppId  *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	EsnBizId  *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeL4EventsSelectiveRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeL4EventsSelectiveRequest) GoString() string {
	return s.String()
}

func (s *DescribeL4EventsSelectiveRequest) SetSourceIp(v string) *DescribeL4EventsSelectiveRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetLang(v string) *DescribeL4EventsSelectiveRequest {
	s.Lang = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetBeginTime(v int64) *DescribeL4EventsSelectiveRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetEndTime(v int64) *DescribeL4EventsSelectiveRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetPage(v int32) *DescribeL4EventsSelectiveRequest {
	s.Page = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetPageSize(v int32) *DescribeL4EventsSelectiveRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetEsnAppId(v string) *DescribeL4EventsSelectiveRequest {
	s.EsnAppId = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetEsnBizId(v string) *DescribeL4EventsSelectiveRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeL4EventsSelectiveRequest) SetGroupId(v string) *DescribeL4EventsSelectiveRequest {
	s.GroupId = &v
	return s
}

type DescribeL4EventsSelectiveResponseBody struct {
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EsnL4EventList []*DescribeL4EventsSelectiveResponseBodyEsnL4EventList `json:"EsnL4EventList,omitempty" xml:"EsnL4EventList,omitempty" type:"Repeated"`
	Total          *int64                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo     map[string]interface{}                                 `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeL4EventsSelectiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeL4EventsSelectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL4EventsSelectiveResponseBody) SetRequestId(v string) *DescribeL4EventsSelectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBody) SetEsnL4EventList(v []*DescribeL4EventsSelectiveResponseBodyEsnL4EventList) *DescribeL4EventsSelectiveResponseBody {
	s.EsnL4EventList = v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBody) SetTotal(v int64) *DescribeL4EventsSelectiveResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeL4EventsSelectiveResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeL4EventsSelectiveResponseBodyEsnL4EventList struct {
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetRegion *string `json:"TargetRegion,omitempty" xml:"TargetRegion,omitempty"`
	AttackType   *int32  `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	Result       *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	AttackSrc    *string `json:"AttackSrc,omitempty" xml:"AttackSrc,omitempty"`
	Duration     *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	AttackTime   *int64  `json:"AttackTime,omitempty" xml:"AttackTime,omitempty"`
	MaxAttack    *int64  `json:"MaxAttack,omitempty" xml:"MaxAttack,omitempty"`
	Target       *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetType   *int32  `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeL4EventsSelectiveResponseBodyEsnL4EventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeL4EventsSelectiveResponseBodyEsnL4EventList) GoString() string {
	return s.String()
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetEndTime(v int64) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.EndTime = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetStartTime(v int64) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.StartTime = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetTargetRegion(v string) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.TargetRegion = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetAttackType(v int32) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.AttackType = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetResult(v int32) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.Result = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetAttackSrc(v string) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.AttackSrc = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetDuration(v int64) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.Duration = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetAttackTime(v int64) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.AttackTime = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetMaxAttack(v int64) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.MaxAttack = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetTarget(v string) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.Target = &v
	return s
}

func (s *DescribeL4EventsSelectiveResponseBodyEsnL4EventList) SetTargetType(v int32) *DescribeL4EventsSelectiveResponseBodyEsnL4EventList {
	s.TargetType = &v
	return s
}

type DescribeL4EventsSelectiveResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeL4EventsSelectiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeL4EventsSelectiveResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeL4EventsSelectiveResponse) GoString() string {
	return s.String()
}

func (s *DescribeL4EventsSelectiveResponse) SetHeaders(v map[string]*string) *DescribeL4EventsSelectiveResponse {
	s.Headers = v
	return s
}

func (s *DescribeL4EventsSelectiveResponse) SetBody(v *DescribeL4EventsSelectiveResponseBody) *DescribeL4EventsSelectiveResponse {
	s.Body = v
	return s
}

type DescribeLayer4RulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesRequest) SetSourceIp(v string) *DescribeLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetLang(v string) *DescribeLayer4RulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetBizId(v int64) *DescribeLayer4RulesRequest {
	s.BizId = &v
	return s
}

type DescribeLayer4RulesResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExtInfo    *DescribeLayer4RulesResponseBodyExtInfo      `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	Total      *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                       `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	L4RuleList []*DescribeLayer4RulesResponseBodyL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
}

func (s DescribeLayer4RulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBody) SetRequestId(v string) *DescribeLayer4RulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetExtInfo(v *DescribeLayer4RulesResponseBodyExtInfo) *DescribeLayer4RulesResponseBody {
	s.ExtInfo = v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetTotal(v int64) *DescribeLayer4RulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeLayer4RulesResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetL4RuleList(v []*DescribeLayer4RulesResponseBodyL4RuleList) *DescribeLayer4RulesResponseBody {
	s.L4RuleList = v
	return s
}

type DescribeLayer4RulesResponseBodyExtInfo struct {
	UseCc         *int32 `json:"UseCc,omitempty" xml:"UseCc,omitempty"`
	NgRouteEnable *int32 `json:"NgRouteEnable,omitempty" xml:"NgRouteEnable,omitempty"`
}

func (s DescribeLayer4RulesResponseBodyExtInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseBodyExtInfo) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBodyExtInfo) SetUseCc(v int32) *DescribeLayer4RulesResponseBodyExtInfo {
	s.UseCc = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyExtInfo) SetNgRouteEnable(v int32) *DescribeLayer4RulesResponseBodyExtInfo {
	s.NgRouteEnable = &v
	return s
}

type DescribeLayer4RulesResponseBodyL4RuleList struct {
	Protocol  *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	AppId     *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BackPort  *int32    `json:"BackPort,omitempty" xml:"BackPort,omitempty"`
	BizId     *int64    `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Rservers  []*string `json:"Rservers,omitempty" xml:"Rservers,omitempty" type:"Repeated"`
	LvsPolicy *string   `json:"LvsPolicy,omitempty" xml:"LvsPolicy,omitempty"`
	FrontPort *int32    `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	RuleId    *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeLayer4RulesResponseBodyL4RuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseBodyL4RuleList) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetProtocol(v string) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetAppId(v int64) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.AppId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetBackPort(v int32) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.BackPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetBizId(v int64) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.BizId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetRservers(v []*string) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.Rservers = v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetLvsPolicy(v string) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.LvsPolicy = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetFrontPort(v int32) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.FrontPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyL4RuleList) SetRuleId(v string) *DescribeLayer4RulesResponseBodyL4RuleList {
	s.RuleId = &v
	return s
}

type DescribeLayer4RulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLayer4RulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponse) SetHeaders(v map[string]*string) *DescribeLayer4RulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer4RulesResponse) SetBody(v *DescribeLayer4RulesResponseBody) *DescribeLayer4RulesResponse {
	s.Body = v
	return s
}

type DescribeMianLiuResSummaryRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeMianLiuResSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMianLiuResSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeMianLiuResSummaryRequest) SetSourceIp(v string) *DescribeMianLiuResSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeMianLiuResSummaryRequest) SetLang(v string) *DescribeMianLiuResSummaryRequest {
	s.Lang = &v
	return s
}

type DescribeMianLiuResSummaryResponseBody struct {
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FullNodesSummary *DescribeMianLiuResSummaryResponseBodyFullNodesSummary `json:"FullNodesSummary,omitempty" xml:"FullNodesSummary,omitempty" type:"Struct"`
	PromptInfo       map[string]interface{}                                 `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeMianLiuResSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMianLiuResSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMianLiuResSummaryResponseBody) SetRequestId(v string) *DescribeMianLiuResSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBody) SetFullNodesSummary(v *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) *DescribeMianLiuResSummaryResponseBody {
	s.FullNodesSummary = v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeMianLiuResSummaryResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeMianLiuResSummaryResponseBodyFullNodesSummary struct {
	Nodes       []*DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	CleanCount  *int32                                                        `json:"CleanCount,omitempty" xml:"CleanCount,omitempty"`
	HoleCount   *int32                                                        `json:"HoleCount,omitempty" xml:"HoleCount,omitempty"`
	UsedCount   *int32                                                        `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	UnUsedCount *int32                                                        `json:"UnUsedCount,omitempty" xml:"UnUsedCount,omitempty"`
	TotalCount  *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMianLiuResSummaryResponseBodyFullNodesSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeMianLiuResSummaryResponseBodyFullNodesSummary) GoString() string {
	return s.String()
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) SetNodes(v []*DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) *DescribeMianLiuResSummaryResponseBodyFullNodesSummary {
	s.Nodes = v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) SetCleanCount(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummary {
	s.CleanCount = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) SetHoleCount(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummary {
	s.HoleCount = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) SetUsedCount(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummary {
	s.UsedCount = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) SetUnUsedCount(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummary {
	s.UnUsedCount = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummary) SetTotalCount(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummary {
	s.TotalCount = &v
	return s
}

type DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes struct {
	Type       *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	InUse      *int32  `json:"InUse,omitempty" xml:"InUse,omitempty"`
	BizName    *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	DdosStatus *int32  `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	IsDisabled *bool   `json:"IsDisabled,omitempty" xml:"IsDisabled,omitempty"`
	Threshold  *int32  `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EsnBizId   *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) GoString() string {
	return s.String()
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetType(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.Type = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetBizId(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizId = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetIp(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.Ip = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetRegionId(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetAppName(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppName = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetRegionName(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.RegionName = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetGroupId(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.GroupId = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetAppId(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.AppId = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetInUse(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.InUse = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetBizName(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.BizName = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetDdosStatus(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.DdosStatus = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetIsDisabled(v bool) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.IsDisabled = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetThreshold(v int32) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.Threshold = &v
	return s
}

func (s *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes) SetEsnBizId(v string) *DescribeMianLiuResSummaryResponseBodyFullNodesSummaryNodes {
	s.EsnBizId = &v
	return s
}

type DescribeMianLiuResSummaryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMianLiuResSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMianLiuResSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMianLiuResSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeMianLiuResSummaryResponse) SetHeaders(v map[string]*string) *DescribeMianLiuResSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeMianLiuResSummaryResponse) SetBody(v *DescribeMianLiuResSummaryResponseBody) *DescribeMianLiuResSummaryResponse {
	s.Body = v
	return s
}

type DescribeNgSourceDiagnosisLogRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page       *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	DetectPort *string `json:"DetectPort,omitempty" xml:"DetectPort,omitempty"`
	DetectIp   *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
}

func (s DescribeNgSourceDiagnosisLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNgSourceDiagnosisLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetSourceIp(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetLang(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetStartTime(v int64) *DescribeNgSourceDiagnosisLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetEndTime(v int64) *DescribeNgSourceDiagnosisLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetPage(v int32) *DescribeNgSourceDiagnosisLogRequest {
	s.Page = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetPageSize(v int32) *DescribeNgSourceDiagnosisLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetAppId(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.AppId = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetBizId(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.BizId = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetSource(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.Source = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetSourcePort(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.SourcePort = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetDetectPort(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.DetectPort = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogRequest) SetDetectIp(v string) *DescribeNgSourceDiagnosisLogRequest {
	s.DetectIp = &v
	return s
}

type DescribeNgSourceDiagnosisLogResponseBody struct {
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeNgSourceDiagnosisLogResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                               `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeNgSourceDiagnosisLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNgSourceDiagnosisLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNgSourceDiagnosisLogResponseBody) SetRequestId(v string) *DescribeNgSourceDiagnosisLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBody) SetStatistic(v []*DescribeNgSourceDiagnosisLogResponseBodyStatistic) *DescribeNgSourceDiagnosisLogResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBody) SetTotal(v int64) *DescribeNgSourceDiagnosisLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeNgSourceDiagnosisLogResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeNgSourceDiagnosisLogResponseBodyStatistic struct {
	Time             *string `json:"Time,omitempty" xml:"Time,omitempty"`
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	SourcePort       *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	DetectPort       *string `json:"DetectPort,omitempty" xml:"DetectPort,omitempty"`
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DetectIp         *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizName          *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SourceIpLocation *string `json:"SourceIpLocation,omitempty" xml:"SourceIpLocation,omitempty"`
}

func (s DescribeNgSourceDiagnosisLogResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeNgSourceDiagnosisLogResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetTime(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.Time = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetSourceIp(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.SourceIp = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetAppName(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.AppName = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetSourcePort(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.SourcePort = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetDetectPort(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.DetectPort = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetAppId(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.AppId = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetDetectIp(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.DetectIp = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetUserId(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.UserId = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetBizName(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.BizName = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetBizId(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.BizId = &v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponseBodyStatistic) SetSourceIpLocation(v string) *DescribeNgSourceDiagnosisLogResponseBodyStatistic {
	s.SourceIpLocation = &v
	return s
}

type DescribeNgSourceDiagnosisLogResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNgSourceDiagnosisLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNgSourceDiagnosisLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNgSourceDiagnosisLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeNgSourceDiagnosisLogResponse) SetHeaders(v map[string]*string) *DescribeNgSourceDiagnosisLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeNgSourceDiagnosisLogResponse) SetBody(v *DescribeNgSourceDiagnosisLogResponseBody) *DescribeNgSourceDiagnosisLogResponse {
	s.Body = v
	return s
}

type DescribeRemainQpsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeRemainQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRemainQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRemainQpsRequest) SetSourceIp(v string) *DescribeRemainQpsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRemainQpsRequest) SetLang(v string) *DescribeRemainQpsRequest {
	s.Lang = &v
	return s
}

type DescribeRemainQpsResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Qps        *int32                 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRemainQpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRemainQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRemainQpsResponseBody) SetRequestId(v string) *DescribeRemainQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRemainQpsResponseBody) SetQps(v int32) *DescribeRemainQpsResponseBody {
	s.Qps = &v
	return s
}

func (s *DescribeRemainQpsResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRemainQpsResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRemainQpsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRemainQpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRemainQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRemainQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRemainQpsResponse) SetHeaders(v map[string]*string) *DescribeRemainQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRemainQpsResponse) SetBody(v *DescribeRemainQpsResponseBody) *DescribeRemainQpsResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticByEsnBizIntervalRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp    *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz    *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Province  *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Country   *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRequestStatisticByEsnBizIntervalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticByEsnBizIntervalRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetSourceIp(v string) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetLang(v string) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetEsnApp(v string) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetEsnBiz(v string) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetStartTime(v int64) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetEndTime(v int64) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetInterval(v int32) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetProvince(v string) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.Province = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetCountry(v string) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.Country = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetPage(v int32) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.Page = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalRequest) SetPageSize(v int32) *DescribeRequestStatisticByEsnBizIntervalRequest {
	s.PageSize = &v
	return s
}

type DescribeRequestStatisticByEsnBizIntervalResponseBody struct {
	RequestId  *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                           `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                           `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticByEsnBizIntervalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticByEsnBizIntervalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBody) SetRequestId(v string) *DescribeRequestStatisticByEsnBizIntervalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBody) SetStatistic(v []*DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) *DescribeRequestStatisticByEsnBizIntervalResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBody) SetTotal(v int64) *DescribeRequestStatisticByEsnBizIntervalResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticByEsnBizIntervalResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic struct {
	Index        *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	Time         *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	FailureCount *int64 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) SetIndex(v int32) *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic {
	s.Index = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) SetTime(v int64) *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic {
	s.Time = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) SetFailureCount(v int64) *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic {
	s.FailureCount = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) SetTotalCount(v int64) *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic {
	s.TotalCount = &v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic) SetSuccessCount(v int64) *DescribeRequestStatisticByEsnBizIntervalResponseBodyStatistic {
	s.SuccessCount = &v
	return s
}

type DescribeRequestStatisticByEsnBizIntervalResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticByEsnBizIntervalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticByEsnBizIntervalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticByEsnBizIntervalResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticByEsnBizIntervalResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticByEsnBizIntervalResponse) SetBody(v *DescribeRequestStatisticByEsnBizIntervalResponseBody) *DescribeRequestStatisticByEsnBizIntervalResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp    *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz    *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayRequest) SetSourceIp(v string) *DescribeRequestStatisticCountByEsnBiz1DayRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayRequest) SetLang(v string) *DescribeRequestStatisticCountByEsnBiz1DayRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayRequest) SetEsnApp(v string) *DescribeRequestStatisticCountByEsnBiz1DayRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayRequest) SetEsnBiz(v string) *DescribeRequestStatisticCountByEsnBiz1DayRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayRequest) SetStartTime(v int64) *DescribeRequestStatisticCountByEsnBiz1DayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayRequest) SetEndTime(v int64) *DescribeRequestStatisticCountByEsnBiz1DayRequest {
	s.EndTime = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayResponseBody struct {
	RequestId  *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	PromptInfo map[string]interface{}                                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponseBody) SetRequestId(v string) *DescribeRequestStatisticCountByEsnBiz1DayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponseBody) SetStatistic(v *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic) *DescribeRequestStatisticCountByEsnBiz1DayResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticCountByEsnBiz1DayResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic struct {
	FailureCount *int64 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic) SetFailureCount(v int64) *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic {
	s.FailureCount = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic) SetSuccessCount(v int64) *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic {
	s.SuccessCount = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic) SetTotalCount(v int64) *DescribeRequestStatisticCountByEsnBiz1DayResponseBodyStatistic {
	s.TotalCount = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticCountByEsnBiz1DayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticCountByEsnBiz1DayResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayResponse) SetBody(v *DescribeRequestStatisticCountByEsnBiz1DayResponseBody) *DescribeRequestStatisticCountByEsnBiz1DayResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp    *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz    *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetSourceIp(v string) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetLang(v string) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetEsnApp(v string) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetEsnBiz(v string) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetStartTime(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetEndTime(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetPage(v int32) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.Page = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) SetPageSize(v int32) *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest {
	s.PageSize = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody struct {
	RequestId  *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                                    `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) SetRequestId(v string) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) SetStatistic(v []*DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) SetTotal(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic struct {
	FailureCount *int64 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	ProvinceId   *int64 `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) SetFailureCount(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic {
	s.FailureCount = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) SetProvinceId(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic {
	s.ProvinceId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) SetTotalCount(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic {
	s.TotalCount = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic) SetSuccessCount(v int64) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBodyStatistic {
	s.SuccessCount = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse struct {
	Headers map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse) SetBody(v *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponseBody) *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1M30MRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp   *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz   *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	Time     *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetSourceIp(v string) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetLang(v string) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetEsnApp(v string) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetEsnBiz(v string) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetTime(v int64) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.Time = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetPage(v int32) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.Page = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MRequest) SetPageSize(v int32) *DescribeRequestStatisticCountByEsnBiz1M30MRequest {
	s.PageSize = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1M30MResponseBody struct {
	RequestId  *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) SetRequestId(v string) *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) SetStatistic(v []*DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic) *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) SetTotal(v int64) *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic struct {
	Time  *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic) SetTime(v int64) *DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic {
	s.Time = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic) SetCount(v int64) *DescribeRequestStatisticCountByEsnBiz1M30MResponseBodyStatistic {
	s.Count = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz1M30MResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz1M30MResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticCountByEsnBiz1M30MResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz1M30MResponse) SetBody(v *DescribeRequestStatisticCountByEsnBiz1M30MResponseBody) *DescribeRequestStatisticCountByEsnBiz1M30MResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz30MRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp   *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz   *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	Time     *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz30MRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz30MRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz30MRequest) SetSourceIp(v string) *DescribeRequestStatisticCountByEsnBiz30MRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MRequest) SetLang(v string) *DescribeRequestStatisticCountByEsnBiz30MRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MRequest) SetEsnApp(v string) *DescribeRequestStatisticCountByEsnBiz30MRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MRequest) SetEsnBiz(v string) *DescribeRequestStatisticCountByEsnBiz30MRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MRequest) SetTime(v int64) *DescribeRequestStatisticCountByEsnBiz30MRequest {
	s.Time = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz30MResponseBody struct {
	RequestId  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	PromptInfo map[string]interface{}                                         `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz30MResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz30MResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponseBody) SetRequestId(v string) *DescribeRequestStatisticCountByEsnBiz30MResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponseBody) SetStatistic(v *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic) *DescribeRequestStatisticCountByEsnBiz30MResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticCountByEsnBiz30MResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic struct {
	FailureCount *int64 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic) SetFailureCount(v int64) *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic {
	s.FailureCount = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic) SetSuccessCount(v int64) *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic {
	s.SuccessCount = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic) SetTotalCount(v int64) *DescribeRequestStatisticCountByEsnBiz30MResponseBodyStatistic {
	s.TotalCount = &v
	return s
}

type DescribeRequestStatisticCountByEsnBiz30MResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticCountByEsnBiz30MResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticCountByEsnBiz30MResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBiz30MResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticCountByEsnBiz30MResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBiz30MResponse) SetBody(v *DescribeRequestStatisticCountByEsnBiz30MResponseBody) *DescribeRequestStatisticCountByEsnBiz30MResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticCountByEsnBizGroup30MRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp   *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz   *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	Time     *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetSourceIp(v string) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetLang(v string) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetEsnApp(v string) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetEsnBiz(v string) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetTime(v int64) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.Time = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetPage(v int32) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.Page = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MRequest) SetPageSize(v int32) *DescribeRequestStatisticCountByEsnBizGroup30MRequest {
	s.PageSize = &v
	return s
}

type DescribeRequestStatisticCountByEsnBizGroup30MResponseBody struct {
	RequestId  *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) SetRequestId(v string) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) SetStatistic(v []*DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) SetTotal(v int64) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic struct {
	DetectIp   *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
	EsnGroupId *string `json:"EsnGroupId,omitempty" xml:"EsnGroupId,omitempty"`
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic) SetDetectIp(v string) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic {
	s.DetectIp = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic) SetEsnGroupId(v string) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic {
	s.EsnGroupId = &v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic) SetCount(v int64) *DescribeRequestStatisticCountByEsnBizGroup30MResponseBodyStatistic {
	s.Count = &v
	return s
}

type DescribeRequestStatisticCountByEsnBizGroup30MResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticCountByEsnBizGroup30MResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticCountByEsnBizGroup30MResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticCountByEsnBizGroup30MResponse) SetBody(v *DescribeRequestStatisticCountByEsnBizGroup30MResponseBody) *DescribeRequestStatisticCountByEsnBizGroup30MResponse {
	s.Body = v
	return s
}

type DescribeRequestStatisticLogRequest struct {
	SourceIp           *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime          *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime            *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page               *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EsnApp             *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz             *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	EsnGroup           *string `json:"EsnGroup,omitempty" xml:"EsnGroup,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Country            *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Province           *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City               *string `json:"City,omitempty" xml:"City,omitempty"`
	ISP                *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	Token              *string `json:"Token,omitempty" xml:"Token,omitempty"`
	SDK                *string `json:"SDK,omitempty" xml:"SDK,omitempty"`
	CallResult         *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	OsType             *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	EscapedLessThan    *int64  `json:"EscapedLessThan,omitempty" xml:"EscapedLessThan,omitempty"`
	EscapedGreaterThan *int64  `json:"EscapedGreaterThan,omitempty" xml:"EscapedGreaterThan,omitempty"`
	DetectIp           *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
}

func (s DescribeRequestStatisticLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticLogRequest) SetSourceIp(v string) *DescribeRequestStatisticLogRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetLang(v string) *DescribeRequestStatisticLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetStartTime(v int64) *DescribeRequestStatisticLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetEndTime(v int64) *DescribeRequestStatisticLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetPage(v int32) *DescribeRequestStatisticLogRequest {
	s.Page = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetPageSize(v int32) *DescribeRequestStatisticLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetEsnApp(v string) *DescribeRequestStatisticLogRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetEsnBiz(v string) *DescribeRequestStatisticLogRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetEsnGroup(v string) *DescribeRequestStatisticLogRequest {
	s.EsnGroup = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetSource(v string) *DescribeRequestStatisticLogRequest {
	s.Source = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetCountry(v string) *DescribeRequestStatisticLogRequest {
	s.Country = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetProvince(v string) *DescribeRequestStatisticLogRequest {
	s.Province = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetCity(v string) *DescribeRequestStatisticLogRequest {
	s.City = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetISP(v string) *DescribeRequestStatisticLogRequest {
	s.ISP = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetToken(v string) *DescribeRequestStatisticLogRequest {
	s.Token = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetSDK(v string) *DescribeRequestStatisticLogRequest {
	s.SDK = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetCallResult(v string) *DescribeRequestStatisticLogRequest {
	s.CallResult = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetOsType(v string) *DescribeRequestStatisticLogRequest {
	s.OsType = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetEscapedLessThan(v int64) *DescribeRequestStatisticLogRequest {
	s.EscapedLessThan = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetEscapedGreaterThan(v int64) *DescribeRequestStatisticLogRequest {
	s.EscapedGreaterThan = &v
	return s
}

func (s *DescribeRequestStatisticLogRequest) SetDetectIp(v string) *DescribeRequestStatisticLogRequest {
	s.DetectIp = &v
	return s
}

type DescribeRequestStatisticLogResponseBody struct {
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeRequestStatisticLogResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                              `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeRequestStatisticLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticLogResponseBody) SetRequestId(v string) *DescribeRequestStatisticLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBody) SetStatistic(v []*DescribeRequestStatisticLogResponseBodyStatistic) *DescribeRequestStatisticLogResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeRequestStatisticLogResponseBody) SetTotal(v int64) *DescribeRequestStatisticLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeRequestStatisticLogResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeRequestStatisticLogResponseBodyStatistic struct {
	ServerTime  *string `json:"ServerTime,omitempty" xml:"ServerTime,omitempty"`
	CallResult  *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CountryId   *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	DetectIp    *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
	EsnAppId    *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	IspId       *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
	Isp         *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	EscapedTime *string `json:"EscapedTime,omitempty" xml:"EscapedTime,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	CityId      *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	ProvinceId  *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	OsType      *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	SdkVersion  *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EsnGroupId  *string `json:"EsnGroupId,omitempty" xml:"EsnGroupId,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	EsnBizId    *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeRequestStatisticLogResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticLogResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetServerTime(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.ServerTime = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetCallResult(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.CallResult = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetToken(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.Token = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetUserId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.UserId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetCountryId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.CountryId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetDetectIp(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.DetectIp = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetEsnAppId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.EsnAppId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetCity(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.City = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetIspId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.IspId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetIsp(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.Isp = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetEscapedTime(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.EscapedTime = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetSourceIp(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.SourceIp = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetCityId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.CityId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetProvinceId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.ProvinceId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetOsType(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.OsType = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetSdkVersion(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.SdkVersion = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetCountry(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.Country = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetEsnGroupId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.EsnGroupId = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetProvince(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.Province = &v
	return s
}

func (s *DescribeRequestStatisticLogResponseBodyStatistic) SetEsnBizId(v string) *DescribeRequestStatisticLogResponseBodyStatistic {
	s.EsnBizId = &v
	return s
}

type DescribeRequestStatisticLogResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRequestStatisticLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRequestStatisticLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestStatisticLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestStatisticLogResponse) SetHeaders(v map[string]*string) *DescribeRequestStatisticLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestStatisticLogResponse) SetBody(v *DescribeRequestStatisticLogResponseBody) *DescribeRequestStatisticLogResponse {
	s.Body = v
	return s
}

type DescribeSDKStatisticLogRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EsnApp    *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz    *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	EsnGroup  *string `json:"EsnGroup,omitempty" xml:"EsnGroup,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Country   *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Province  *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City      *string `json:"City,omitempty" xml:"City,omitempty"`
	ISP       *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	OsType    *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	DetectIp  *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
}

func (s DescribeSDKStatisticLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticLogRequest) SetSourceIp(v string) *DescribeSDKStatisticLogRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetLang(v string) *DescribeSDKStatisticLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetStartTime(v int64) *DescribeSDKStatisticLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetEndTime(v int64) *DescribeSDKStatisticLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetPage(v int32) *DescribeSDKStatisticLogRequest {
	s.Page = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetPageSize(v int32) *DescribeSDKStatisticLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetEsnApp(v string) *DescribeSDKStatisticLogRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetEsnBiz(v string) *DescribeSDKStatisticLogRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetEsnGroup(v string) *DescribeSDKStatisticLogRequest {
	s.EsnGroup = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetSource(v string) *DescribeSDKStatisticLogRequest {
	s.Source = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetCountry(v string) *DescribeSDKStatisticLogRequest {
	s.Country = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetProvince(v string) *DescribeSDKStatisticLogRequest {
	s.Province = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetCity(v string) *DescribeSDKStatisticLogRequest {
	s.City = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetISP(v string) *DescribeSDKStatisticLogRequest {
	s.ISP = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetToken(v string) *DescribeSDKStatisticLogRequest {
	s.Token = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetResult(v string) *DescribeSDKStatisticLogRequest {
	s.Result = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetOsType(v string) *DescribeSDKStatisticLogRequest {
	s.OsType = &v
	return s
}

func (s *DescribeSDKStatisticLogRequest) SetDetectIp(v string) *DescribeSDKStatisticLogRequest {
	s.DetectIp = &v
	return s
}

type DescribeSDKStatisticLogResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeSDKStatisticLogResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                          `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeSDKStatisticLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticLogResponseBody) SetRequestId(v string) *DescribeSDKStatisticLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBody) SetStatistic(v []*DescribeSDKStatisticLogResponseBodyStatistic) *DescribeSDKStatisticLogResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeSDKStatisticLogResponseBody) SetTotal(v int64) *DescribeSDKStatisticLogResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeSDKStatisticLogResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeSDKStatisticLogResponseBodyStatistic struct {
	ServerTime    *string `json:"ServerTime,omitempty" xml:"ServerTime,omitempty"`
	CallResult    *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
	PingResult    *string `json:"PingResult,omitempty" xml:"PingResult,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CountryId     *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	City          *string `json:"City,omitempty" xml:"City,omitempty"`
	ConnectResult *string `json:"ConnectResult,omitempty" xml:"ConnectResult,omitempty"`
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProvinceId    *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	DetectIp      *string `json:"DetectIp,omitempty" xml:"DetectIp,omitempty"`
	TraceResult   *string `json:"TraceResult,omitempty" xml:"TraceResult,omitempty"`
	EsnAppId      *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
	MType         *string `json:"MType,omitempty" xml:"MType,omitempty"`
	IspId         *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
	Isp           *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	CityId        *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	VerdictCode   *string `json:"VerdictCode,omitempty" xml:"VerdictCode,omitempty"`
	SdkVersion    *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	Country       *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EsnGroupId    *string `json:"EsnGroupId,omitempty" xml:"EsnGroupId,omitempty"`
	Province      *string `json:"Province,omitempty" xml:"Province,omitempty"`
	EsnBizId      *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
}

func (s DescribeSDKStatisticLogResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticLogResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetServerTime(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.ServerTime = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetCallResult(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.CallResult = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetToken(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.Token = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetPingResult(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.PingResult = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetUserId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.UserId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetCountryId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.CountryId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetCity(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.City = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetConnectResult(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.ConnectResult = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetSourceIp(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.SourceIp = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetProvinceId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.ProvinceId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetDetectIp(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.DetectIp = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetTraceResult(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.TraceResult = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetEsnAppId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.EsnAppId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetMType(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.MType = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetIspId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.IspId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetIsp(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.Isp = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetCityId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.CityId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetOsType(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.OsType = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetVerdictCode(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.VerdictCode = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetSdkVersion(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.SdkVersion = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetCountry(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.Country = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetEsnGroupId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.EsnGroupId = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetProvince(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.Province = &v
	return s
}

func (s *DescribeSDKStatisticLogResponseBodyStatistic) SetEsnBizId(v string) *DescribeSDKStatisticLogResponseBodyStatistic {
	s.EsnBizId = &v
	return s
}

type DescribeSDKStatisticLogResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSDKStatisticLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSDKStatisticLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticLogResponse) SetHeaders(v map[string]*string) *DescribeSDKStatisticLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDKStatisticLogResponse) SetBody(v *DescribeSDKStatisticLogResponseBody) *DescribeSDKStatisticLogResponse {
	s.Body = v
	return s
}

type DescribeSDKStatisticResultByEsnBizISP1M30MRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp   *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz   *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	Time     *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetSourceIp(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetLang(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetEsnApp(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetEsnBiz(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetTime(v int64) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.Time = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetPage(v int32) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.Page = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetPageSize(v int32) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) SetType(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MRequest {
	s.Type = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody struct {
	RequestId  *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) SetRequestId(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) SetStatistic(v []*DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) SetTotal(v int64) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic struct {
	Time   *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	IspId  *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
	Count  *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) SetTime(v int64) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic {
	s.Time = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) SetResult(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic {
	s.Result = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) SetIspId(v string) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic {
	s.IspId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic) SetCount(v int64) *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBodyStatistic {
	s.Count = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizISP1M30MResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISP1M30MResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponse) SetHeaders(v map[string]*string) *DescribeSDKStatisticResultByEsnBizISP1M30MResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISP1M30MResponse) SetBody(v *DescribeSDKStatisticResultByEsnBizISP1M30MResponseBody) *DescribeSDKStatisticResultByEsnBizISP1M30MResponse {
	s.Body = v
	return s
}

type DescribeSDKStatisticResultByEsnBizISPIntervalRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp    *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz    *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Interval  *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Province  *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetSourceIp(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetLang(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetEsnApp(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetEsnBiz(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetStartTime(v int64) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetEndTime(v int64) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetPage(v int32) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.Page = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetPageSize(v int32) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetType(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.Type = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetInterval(v int64) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.Interval = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) SetProvince(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalRequest {
	s.Province = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody struct {
	RequestId  *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) SetRequestId(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) SetStatistic(v []*DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) SetTotal(v int64) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic struct {
	Index  *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Time   *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	IspId  *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
	Count  *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) SetIndex(v int32) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic {
	s.Index = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) SetTime(v int64) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic {
	s.Time = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) SetResult(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic {
	s.Result = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) SetIspId(v string) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic {
	s.IspId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic) SetCount(v int64) *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBodyStatistic {
	s.Count = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizISPIntervalResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizISPIntervalResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponse) SetHeaders(v map[string]*string) *DescribeSDKStatisticResultByEsnBizISPIntervalResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizISPIntervalResponse) SetBody(v *DescribeSDKStatisticResultByEsnBizISPIntervalResponseBody) *DescribeSDKStatisticResultByEsnBizISPIntervalResponse {
	s.Body = v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince1DayRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp    *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz    *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetSourceIp(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetLang(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetEsnApp(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetEsnBiz(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetStartTime(v int64) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetEndTime(v int64) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetPage(v int32) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.Page = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetPageSize(v int32) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) SetType(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayRequest {
	s.Type = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody struct {
	RequestId  *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                                 `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) SetRequestId(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) SetStatistic(v []*DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) SetTotal(v int64) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic struct {
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ProvinceId *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic) SetResult(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic {
	s.Result = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic) SetProvinceId(v string) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic {
	s.ProvinceId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic) SetCount(v int64) *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBodyStatistic {
	s.Count = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince1DayResponse struct {
	Headers map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince1DayResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponse) SetHeaders(v map[string]*string) *DescribeSDKStatisticResultByEsnBizProvince1DayResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince1DayResponse) SetBody(v *DescribeSDKStatisticResultByEsnBizProvince1DayResponseBody) *DescribeSDKStatisticResultByEsnBizProvince1DayResponse {
	s.Body = v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince30MRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnApp   *string `json:"EsnApp,omitempty" xml:"EsnApp,omitempty"`
	EsnBiz   *string `json:"EsnBiz,omitempty" xml:"EsnBiz,omitempty"`
	Time     *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetSourceIp(v string) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetLang(v string) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetEsnApp(v string) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.EsnApp = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetEsnBiz(v string) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.EsnBiz = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetTime(v int64) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.Time = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetPage(v int32) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.Page = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetPageSize(v int32) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MRequest) SetType(v string) *DescribeSDKStatisticResultByEsnBizProvince30MRequest {
	s.Type = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince30MResponseBody struct {
	RequestId  *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistic  []*DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
	Total      *int64                                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                                                `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) SetRequestId(v string) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) SetStatistic(v []*DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody {
	s.Statistic = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) SetTotal(v int64) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic struct {
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ProvinceId *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	Count      *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic) SetResult(v string) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic {
	s.Result = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic) SetProvinceId(v string) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic {
	s.ProvinceId = &v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic) SetCount(v int64) *DescribeSDKStatisticResultByEsnBizProvince30MResponseBodyStatistic {
	s.Count = &v
	return s
}

type DescribeSDKStatisticResultByEsnBizProvince30MResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSDKStatisticResultByEsnBizProvince30MResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponse) SetHeaders(v map[string]*string) *DescribeSDKStatisticResultByEsnBizProvince30MResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDKStatisticResultByEsnBizProvince30MResponse) SetBody(v *DescribeSDKStatisticResultByEsnBizProvince30MResponseBody) *DescribeSDKStatisticResultByEsnBizProvince30MResponse {
	s.Body = v
	return s
}

type DescribeSourceFailureTopIpRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId  *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s DescribeSourceFailureTopIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTopIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTopIpRequest) SetSourceIp(v string) *DescribeSourceFailureTopIpRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSourceFailureTopIpRequest) SetLang(v string) *DescribeSourceFailureTopIpRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSourceFailureTopIpRequest) SetEsnBizId(v string) *DescribeSourceFailureTopIpRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeSourceFailureTopIpRequest) SetStartTime(v int64) *DescribeSourceFailureTopIpRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSourceFailureTopIpRequest) SetEndTime(v int64) *DescribeSourceFailureTopIpRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSourceFailureTopIpRequest) SetLimit(v int32) *DescribeSourceFailureTopIpRequest {
	s.Limit = &v
	return s
}

type DescribeSourceFailureTopIpResponseBody struct {
	TopIps    []*DescribeSourceFailureTopIpResponseBodyTopIps `json:"TopIps,omitempty" xml:"TopIps,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSourceFailureTopIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTopIpResponseBody) SetTopIps(v []*DescribeSourceFailureTopIpResponseBodyTopIps) *DescribeSourceFailureTopIpResponseBody {
	s.TopIps = v
	return s
}

func (s *DescribeSourceFailureTopIpResponseBody) SetRequestId(v string) *DescribeSourceFailureTopIpResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSourceFailureTopIpResponseBodyTopIps struct {
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port       *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSourceFailureTopIpResponseBodyTopIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTopIpResponseBodyTopIps) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTopIpResponseBodyTopIps) SetIp(v string) *DescribeSourceFailureTopIpResponseBodyTopIps {
	s.Ip = &v
	return s
}

func (s *DescribeSourceFailureTopIpResponseBodyTopIps) SetPort(v int64) *DescribeSourceFailureTopIpResponseBodyTopIps {
	s.Port = &v
	return s
}

func (s *DescribeSourceFailureTopIpResponseBodyTopIps) SetTotalCount(v int64) *DescribeSourceFailureTopIpResponseBodyTopIps {
	s.TotalCount = &v
	return s
}

type DescribeSourceFailureTopIpResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSourceFailureTopIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSourceFailureTopIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTopIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTopIpResponse) SetHeaders(v map[string]*string) *DescribeSourceFailureTopIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeSourceFailureTopIpResponse) SetBody(v *DescribeSourceFailureTopIpResponseBody) *DescribeSourceFailureTopIpResponse {
	s.Body = v
	return s
}

type DescribeSourceFailureTrendGraphRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId  *string `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeSourceFailureTrendGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTrendGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTrendGraphRequest) SetSourceIp(v string) *DescribeSourceFailureTrendGraphRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphRequest) SetLang(v string) *DescribeSourceFailureTrendGraphRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphRequest) SetEsnBizId(v string) *DescribeSourceFailureTrendGraphRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphRequest) SetStartTime(v int64) *DescribeSourceFailureTrendGraphRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphRequest) SetEndTime(v int64) *DescribeSourceFailureTrendGraphRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphRequest) SetInterval(v int32) *DescribeSourceFailureTrendGraphRequest {
	s.Interval = &v
	return s
}

type DescribeSourceFailureTrendGraphResponseBody struct {
	RequestId  *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrendGraph []*DescribeSourceFailureTrendGraphResponseBodyTrendGraph `json:"TrendGraph,omitempty" xml:"TrendGraph,omitempty" type:"Repeated"`
}

func (s DescribeSourceFailureTrendGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTrendGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTrendGraphResponseBody) SetRequestId(v string) *DescribeSourceFailureTrendGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphResponseBody) SetTrendGraph(v []*DescribeSourceFailureTrendGraphResponseBodyTrendGraph) *DescribeSourceFailureTrendGraphResponseBody {
	s.TrendGraph = v
	return s
}

type DescribeSourceFailureTrendGraphResponseBodyTrendGraph struct {
	Index      *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSourceFailureTrendGraphResponseBodyTrendGraph) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTrendGraphResponseBodyTrendGraph) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTrendGraphResponseBodyTrendGraph) SetIndex(v int64) *DescribeSourceFailureTrendGraphResponseBodyTrendGraph {
	s.Index = &v
	return s
}

func (s *DescribeSourceFailureTrendGraphResponseBodyTrendGraph) SetTotalCount(v int64) *DescribeSourceFailureTrendGraphResponseBodyTrendGraph {
	s.TotalCount = &v
	return s
}

type DescribeSourceFailureTrendGraphResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSourceFailureTrendGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSourceFailureTrendGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceFailureTrendGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeSourceFailureTrendGraphResponse) SetHeaders(v map[string]*string) *DescribeSourceFailureTrendGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeSourceFailureTrendGraphResponse) SetBody(v *DescribeSourceFailureTrendGraphResponseBody) *DescribeSourceFailureTrendGraphResponse {
	s.Body = v
	return s
}

type DescribeUploadPreSignRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUploadPreSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadPreSignRequest) GoString() string {
	return s.String()
}

func (s *DescribeUploadPreSignRequest) SetSourceIp(v string) *DescribeUploadPreSignRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUploadPreSignRequest) SetLang(v string) *DescribeUploadPreSignRequest {
	s.Lang = &v
	return s
}

type DescribeUploadPreSignResponseBody struct {
	Policy     *string                `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessId   *string                `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Signature  *string                `json:"Signature,omitempty" xml:"Signature,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	Host       *string                `json:"Host,omitempty" xml:"Host,omitempty"`
	Key        *string                `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeUploadPreSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadPreSignResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUploadPreSignResponseBody) SetPolicy(v string) *DescribeUploadPreSignResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetRequestId(v string) *DescribeUploadPreSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetAccessId(v string) *DescribeUploadPreSignResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetSignature(v string) *DescribeUploadPreSignResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeUploadPreSignResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetHost(v string) *DescribeUploadPreSignResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetKey(v string) *DescribeUploadPreSignResponseBody {
	s.Key = &v
	return s
}

type DescribeUploadPreSignResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUploadPreSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUploadPreSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadPreSignResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadPreSignResponse) SetHeaders(v map[string]*string) *DescribeUploadPreSignResponse {
	s.Headers = v
	return s
}

func (s *DescribeUploadPreSignResponse) SetBody(v *DescribeUploadPreSignResponseBody) *DescribeUploadPreSignResponse {
	s.Body = v
	return s
}

type DescribeUserFlowInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserFlowInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowInfoRequest) SetSourceIp(v string) *DescribeUserFlowInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserFlowInfoRequest) SetLang(v string) *DescribeUserFlowInfoRequest {
	s.Lang = &v
	return s
}

type DescribeUserFlowInfoResponseBody struct {
	FlowInfo   *DescribeUserFlowInfoResponseBodyFlowInfo `json:"FlowInfo,omitempty" xml:"FlowInfo,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                    `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DescribeUserFlowInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowInfoResponseBody) SetFlowInfo(v *DescribeUserFlowInfoResponseBodyFlowInfo) *DescribeUserFlowInfoResponseBody {
	s.FlowInfo = v
	return s
}

func (s *DescribeUserFlowInfoResponseBody) SetRequestId(v string) *DescribeUserFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserFlowInfoResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeUserFlowInfoResponseBody {
	s.PromptInfo = v
	return s
}

type DescribeUserFlowInfoResponseBodyFlowInfo struct {
	BizBandWidth *int64 `json:"BizBandWidth,omitempty" xml:"BizBandWidth,omitempty"`
}

func (s DescribeUserFlowInfoResponseBodyFlowInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowInfoResponseBodyFlowInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowInfoResponseBodyFlowInfo) SetBizBandWidth(v int64) *DescribeUserFlowInfoResponseBodyFlowInfo {
	s.BizBandWidth = &v
	return s
}

type DescribeUserFlowInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserFlowInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowInfoResponse) SetHeaders(v map[string]*string) *DescribeUserFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserFlowInfoResponse) SetBody(v *DescribeUserFlowInfoResponseBody) *DescribeUserFlowInfoResponse {
	s.Body = v
	return s
}

type DescribeUserFlowReportRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EsnGroupId *string `json:"EsnGroupId,omitempty" xml:"EsnGroupId,omitempty"`
	EsnBizId   *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	EsnAppId   *string `json:"EsnAppId,omitempty" xml:"EsnAppId,omitempty"`
}

func (s DescribeUserFlowReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowReportRequest) SetSourceIp(v string) *DescribeUserFlowReportRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetLang(v string) *DescribeUserFlowReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetEndTime(v int64) *DescribeUserFlowReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetInterval(v int64) *DescribeUserFlowReportRequest {
	s.Interval = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetStartTime(v int64) *DescribeUserFlowReportRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetEsnGroupId(v string) *DescribeUserFlowReportRequest {
	s.EsnGroupId = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetEsnBizId(v int64) *DescribeUserFlowReportRequest {
	s.EsnBizId = &v
	return s
}

func (s *DescribeUserFlowReportRequest) SetEsnAppId(v string) *DescribeUserFlowReportRequest {
	s.EsnAppId = &v
	return s
}

type DescribeUserFlowReportResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                        `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	FlowReport *DescribeUserFlowReportResponseBodyFlowReport `json:"FlowReport,omitempty" xml:"FlowReport,omitempty" type:"Struct"`
}

func (s DescribeUserFlowReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowReportResponseBody) SetRequestId(v string) *DescribeUserFlowReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserFlowReportResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeUserFlowReportResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeUserFlowReportResponseBody) SetFlowReport(v *DescribeUserFlowReportResponseBodyFlowReport) *DescribeUserFlowReportResponseBody {
	s.FlowReport = v
	return s
}

type DescribeUserFlowReportResponseBodyFlowReport struct {
	Time       []*string `json:"Time,omitempty" xml:"Time,omitempty" type:"Repeated"`
	MaxInflow  []*string `json:"MaxInflow,omitempty" xml:"MaxInflow,omitempty" type:"Repeated"`
	AttackFlow []*string `json:"AttackFlow,omitempty" xml:"AttackFlow,omitempty" type:"Repeated"`
	MaxOutFlow []*string `json:"MaxOutFlow,omitempty" xml:"MaxOutFlow,omitempty" type:"Repeated"`
}

func (s DescribeUserFlowReportResponseBodyFlowReport) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowReportResponseBodyFlowReport) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowReportResponseBodyFlowReport) SetTime(v []*string) *DescribeUserFlowReportResponseBodyFlowReport {
	s.Time = v
	return s
}

func (s *DescribeUserFlowReportResponseBodyFlowReport) SetMaxInflow(v []*string) *DescribeUserFlowReportResponseBodyFlowReport {
	s.MaxInflow = v
	return s
}

func (s *DescribeUserFlowReportResponseBodyFlowReport) SetAttackFlow(v []*string) *DescribeUserFlowReportResponseBodyFlowReport {
	s.AttackFlow = v
	return s
}

func (s *DescribeUserFlowReportResponseBodyFlowReport) SetMaxOutFlow(v []*string) *DescribeUserFlowReportResponseBodyFlowReport {
	s.MaxOutFlow = v
	return s
}

type DescribeUserFlowReportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserFlowReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserFlowReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserFlowReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowReportResponse) SetHeaders(v map[string]*string) *DescribeUserFlowReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserFlowReportResponse) SetBody(v *DescribeUserFlowReportResponseBody) *DescribeUserFlowReportResponse {
	s.Body = v
	return s
}

type DescribeUserTotalFlowReportRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserTotalFlowReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTotalFlowReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTotalFlowReportRequest) SetSourceIp(v string) *DescribeUserTotalFlowReportRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserTotalFlowReportRequest) SetLang(v string) *DescribeUserTotalFlowReportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserTotalFlowReportRequest) SetEndTime(v int64) *DescribeUserTotalFlowReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserTotalFlowReportRequest) SetInterval(v int64) *DescribeUserTotalFlowReportRequest {
	s.Interval = &v
	return s
}

func (s *DescribeUserTotalFlowReportRequest) SetStartTime(v int64) *DescribeUserTotalFlowReportRequest {
	s.StartTime = &v
	return s
}

type DescribeUserTotalFlowReportResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{}                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	FlowReport *DescribeUserTotalFlowReportResponseBodyFlowReport `json:"FlowReport,omitempty" xml:"FlowReport,omitempty" type:"Struct"`
}

func (s DescribeUserTotalFlowReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTotalFlowReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTotalFlowReportResponseBody) SetRequestId(v string) *DescribeUserTotalFlowReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTotalFlowReportResponseBody) SetPromptInfo(v map[string]interface{}) *DescribeUserTotalFlowReportResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DescribeUserTotalFlowReportResponseBody) SetFlowReport(v *DescribeUserTotalFlowReportResponseBodyFlowReport) *DescribeUserTotalFlowReportResponseBody {
	s.FlowReport = v
	return s
}

type DescribeUserTotalFlowReportResponseBodyFlowReport struct {
	Time       []*string `json:"Time,omitempty" xml:"Time,omitempty" type:"Repeated"`
	MaxInflow  []*string `json:"MaxInflow,omitempty" xml:"MaxInflow,omitempty" type:"Repeated"`
	AttackFlow []*string `json:"AttackFlow,omitempty" xml:"AttackFlow,omitempty" type:"Repeated"`
	MaxOutFlow []*string `json:"MaxOutFlow,omitempty" xml:"MaxOutFlow,omitempty" type:"Repeated"`
}

func (s DescribeUserTotalFlowReportResponseBodyFlowReport) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTotalFlowReportResponseBodyFlowReport) GoString() string {
	return s.String()
}

func (s *DescribeUserTotalFlowReportResponseBodyFlowReport) SetTime(v []*string) *DescribeUserTotalFlowReportResponseBodyFlowReport {
	s.Time = v
	return s
}

func (s *DescribeUserTotalFlowReportResponseBodyFlowReport) SetMaxInflow(v []*string) *DescribeUserTotalFlowReportResponseBodyFlowReport {
	s.MaxInflow = v
	return s
}

func (s *DescribeUserTotalFlowReportResponseBodyFlowReport) SetAttackFlow(v []*string) *DescribeUserTotalFlowReportResponseBodyFlowReport {
	s.AttackFlow = v
	return s
}

func (s *DescribeUserTotalFlowReportResponseBodyFlowReport) SetMaxOutFlow(v []*string) *DescribeUserTotalFlowReportResponseBodyFlowReport {
	s.MaxOutFlow = v
	return s
}

type DescribeUserTotalFlowReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserTotalFlowReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserTotalFlowReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserTotalFlowReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTotalFlowReportResponse) SetHeaders(v map[string]*string) *DescribeUserTotalFlowReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTotalFlowReportResponse) SetBody(v *DescribeUserTotalFlowReportResponseBody) *DescribeUserTotalFlowReportResponse {
	s.Body = v
	return s
}

type DownloadAppKeyFileRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	KeyVersion *string `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty"`
}

func (s DownloadAppKeyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadAppKeyFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadAppKeyFileRequest) SetSourceIp(v string) *DownloadAppKeyFileRequest {
	s.SourceIp = &v
	return s
}

func (s *DownloadAppKeyFileRequest) SetLang(v string) *DownloadAppKeyFileRequest {
	s.Lang = &v
	return s
}

func (s *DownloadAppKeyFileRequest) SetAppId(v int64) *DownloadAppKeyFileRequest {
	s.AppId = &v
	return s
}

func (s *DownloadAppKeyFileRequest) SetPlatform(v string) *DownloadAppKeyFileRequest {
	s.Platform = &v
	return s
}

func (s *DownloadAppKeyFileRequest) SetKeyVersion(v string) *DownloadAppKeyFileRequest {
	s.KeyVersion = &v
	return s
}

type DownloadAppKeyFileResponseBody struct {
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo         map[string]interface{}                            `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	DownloadLinkResult *DownloadAppKeyFileResponseBodyDownloadLinkResult `json:"DownloadLinkResult,omitempty" xml:"DownloadLinkResult,omitempty" type:"Struct"`
}

func (s DownloadAppKeyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadAppKeyFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadAppKeyFileResponseBody) SetRequestId(v string) *DownloadAppKeyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadAppKeyFileResponseBody) SetPromptInfo(v map[string]interface{}) *DownloadAppKeyFileResponseBody {
	s.PromptInfo = v
	return s
}

func (s *DownloadAppKeyFileResponseBody) SetDownloadLinkResult(v *DownloadAppKeyFileResponseBodyDownloadLinkResult) *DownloadAppKeyFileResponseBody {
	s.DownloadLinkResult = v
	return s
}

type DownloadAppKeyFileResponseBodyDownloadLinkResult struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s DownloadAppKeyFileResponseBodyDownloadLinkResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadAppKeyFileResponseBodyDownloadLinkResult) GoString() string {
	return s.String()
}

func (s *DownloadAppKeyFileResponseBodyDownloadLinkResult) SetDownloadLink(v string) *DownloadAppKeyFileResponseBodyDownloadLinkResult {
	s.DownloadLink = &v
	return s
}

type DownloadAppKeyFileResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadAppKeyFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadAppKeyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadAppKeyFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadAppKeyFileResponse) SetHeaders(v map[string]*string) *DownloadAppKeyFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadAppKeyFileResponse) SetBody(v *DownloadAppKeyFileResponseBody) *DownloadAppKeyFileResponse {
	s.Body = v
	return s
}

type DownloadCcRouteRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DownloadCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *DownloadCcRouteRulesRequest) SetSourceIp(v string) *DownloadCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DownloadCcRouteRulesRequest) SetLang(v string) *DownloadCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *DownloadCcRouteRulesRequest) SetBizId(v int64) *DownloadCcRouteRulesRequest {
	s.BizId = &v
	return s
}

type DownloadCcRouteRulesResponseBody struct {
	DownloadFileResult *DownloadCcRouteRulesResponseBodyDownloadFileResult `json:"DownloadFileResult,omitempty" xml:"DownloadFileResult,omitempty" type:"Struct"`
	RequestId          *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo         map[string]interface{}                              `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DownloadCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadCcRouteRulesResponseBody) SetDownloadFileResult(v *DownloadCcRouteRulesResponseBodyDownloadFileResult) *DownloadCcRouteRulesResponseBody {
	s.DownloadFileResult = v
	return s
}

func (s *DownloadCcRouteRulesResponseBody) SetRequestId(v string) *DownloadCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *DownloadCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type DownloadCcRouteRulesResponseBodyDownloadFileResult struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s DownloadCcRouteRulesResponseBodyDownloadFileResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadCcRouteRulesResponseBodyDownloadFileResult) GoString() string {
	return s.String()
}

func (s *DownloadCcRouteRulesResponseBodyDownloadFileResult) SetDownloadLink(v string) *DownloadCcRouteRulesResponseBodyDownloadFileResult {
	s.DownloadLink = &v
	return s
}

type DownloadCcRouteRulesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *DownloadCcRouteRulesResponse) SetHeaders(v map[string]*string) *DownloadCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *DownloadCcRouteRulesResponse) SetBody(v *DownloadCcRouteRulesResponseBody) *DownloadCcRouteRulesResponse {
	s.Body = v
	return s
}

type DownloadFlexAccRulesFileRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DownloadFlexAccRulesFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadFlexAccRulesFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadFlexAccRulesFileRequest) SetSourceIp(v string) *DownloadFlexAccRulesFileRequest {
	s.SourceIp = &v
	return s
}

func (s *DownloadFlexAccRulesFileRequest) SetBizId(v int64) *DownloadFlexAccRulesFileRequest {
	s.BizId = &v
	return s
}

type DownloadFlexAccRulesFileResponseBody struct {
	DownloadFileResult *DownloadFlexAccRulesFileResponseBodyDownloadFileResult `json:"DownloadFileResult,omitempty" xml:"DownloadFileResult,omitempty" type:"Struct"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo         map[string]interface{}                                  `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DownloadFlexAccRulesFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadFlexAccRulesFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadFlexAccRulesFileResponseBody) SetDownloadFileResult(v *DownloadFlexAccRulesFileResponseBodyDownloadFileResult) *DownloadFlexAccRulesFileResponseBody {
	s.DownloadFileResult = v
	return s
}

func (s *DownloadFlexAccRulesFileResponseBody) SetRequestId(v string) *DownloadFlexAccRulesFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadFlexAccRulesFileResponseBody) SetPromptInfo(v map[string]interface{}) *DownloadFlexAccRulesFileResponseBody {
	s.PromptInfo = v
	return s
}

type DownloadFlexAccRulesFileResponseBodyDownloadFileResult struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s DownloadFlexAccRulesFileResponseBodyDownloadFileResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadFlexAccRulesFileResponseBodyDownloadFileResult) GoString() string {
	return s.String()
}

func (s *DownloadFlexAccRulesFileResponseBodyDownloadFileResult) SetDownloadLink(v string) *DownloadFlexAccRulesFileResponseBodyDownloadFileResult {
	s.DownloadLink = &v
	return s
}

type DownloadFlexAccRulesFileResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadFlexAccRulesFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadFlexAccRulesFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadFlexAccRulesFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadFlexAccRulesFileResponse) SetHeaders(v map[string]*string) *DownloadFlexAccRulesFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadFlexAccRulesFileResponse) SetBody(v *DownloadFlexAccRulesFileResponseBody) *DownloadFlexAccRulesFileResponse {
	s.Body = v
	return s
}

type DownloadLayer4RulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DownloadLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *DownloadLayer4RulesRequest) SetSourceIp(v string) *DownloadLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DownloadLayer4RulesRequest) SetLang(v string) *DownloadLayer4RulesRequest {
	s.Lang = &v
	return s
}

func (s *DownloadLayer4RulesRequest) SetBizId(v int64) *DownloadLayer4RulesRequest {
	s.BizId = &v
	return s
}

type DownloadLayer4RulesResponseBody struct {
	DownloadFileResult *DownloadLayer4RulesResponseBodyDownloadFileResult `json:"DownloadFileResult,omitempty" xml:"DownloadFileResult,omitempty" type:"Struct"`
	RequestId          *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo         map[string]interface{}                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DownloadLayer4RulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadLayer4RulesResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadLayer4RulesResponseBody) SetDownloadFileResult(v *DownloadLayer4RulesResponseBodyDownloadFileResult) *DownloadLayer4RulesResponseBody {
	s.DownloadFileResult = v
	return s
}

func (s *DownloadLayer4RulesResponseBody) SetRequestId(v string) *DownloadLayer4RulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadLayer4RulesResponseBody) SetPromptInfo(v map[string]interface{}) *DownloadLayer4RulesResponseBody {
	s.PromptInfo = v
	return s
}

type DownloadLayer4RulesResponseBodyDownloadFileResult struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s DownloadLayer4RulesResponseBodyDownloadFileResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadLayer4RulesResponseBodyDownloadFileResult) GoString() string {
	return s.String()
}

func (s *DownloadLayer4RulesResponseBodyDownloadFileResult) SetDownloadLink(v string) *DownloadLayer4RulesResponseBodyDownloadFileResult {
	s.DownloadLink = &v
	return s
}

type DownloadLayer4RulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadLayer4RulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *DownloadLayer4RulesResponse) SetHeaders(v map[string]*string) *DownloadLayer4RulesResponse {
	s.Headers = v
	return s
}

func (s *DownloadLayer4RulesResponse) SetBody(v *DownloadLayer4RulesResponseBody) *DownloadLayer4RulesResponse {
	s.Body = v
	return s
}

type DownloadSdkFileRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s DownloadSdkFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadSdkFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadSdkFileRequest) SetSourceIp(v string) *DownloadSdkFileRequest {
	s.SourceIp = &v
	return s
}

func (s *DownloadSdkFileRequest) SetLang(v string) *DownloadSdkFileRequest {
	s.Lang = &v
	return s
}

func (s *DownloadSdkFileRequest) SetAppId(v int64) *DownloadSdkFileRequest {
	s.AppId = &v
	return s
}

func (s *DownloadSdkFileRequest) SetSdkVersion(v string) *DownloadSdkFileRequest {
	s.SdkVersion = &v
	return s
}

func (s *DownloadSdkFileRequest) SetPlatform(v string) *DownloadSdkFileRequest {
	s.Platform = &v
	return s
}

type DownloadSdkFileResponseBody struct {
	DownloadFileResult *DownloadSdkFileResponseBodyDownloadFileResult `json:"DownloadFileResult,omitempty" xml:"DownloadFileResult,omitempty" type:"Struct"`
	RequestId          *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo         map[string]interface{}                         `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DownloadSdkFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadSdkFileResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadSdkFileResponseBody) SetDownloadFileResult(v *DownloadSdkFileResponseBodyDownloadFileResult) *DownloadSdkFileResponseBody {
	s.DownloadFileResult = v
	return s
}

func (s *DownloadSdkFileResponseBody) SetRequestId(v string) *DownloadSdkFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadSdkFileResponseBody) SetPromptInfo(v map[string]interface{}) *DownloadSdkFileResponseBody {
	s.PromptInfo = v
	return s
}

type DownloadSdkFileResponseBodyDownloadFileResult struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s DownloadSdkFileResponseBodyDownloadFileResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadSdkFileResponseBodyDownloadFileResult) GoString() string {
	return s.String()
}

func (s *DownloadSdkFileResponseBodyDownloadFileResult) SetDownloadLink(v string) *DownloadSdkFileResponseBodyDownloadFileResult {
	s.DownloadLink = &v
	return s
}

type DownloadSdkFileResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadSdkFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadSdkFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadSdkFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadSdkFileResponse) SetHeaders(v map[string]*string) *DownloadSdkFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadSdkFileResponse) SetBody(v *DownloadSdkFileResponseBody) *DownloadSdkFileResponse {
	s.Body = v
	return s
}

type DownloadUserTotalFlowReportRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DownloadUserTotalFlowReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadUserTotalFlowReportRequest) GoString() string {
	return s.String()
}

func (s *DownloadUserTotalFlowReportRequest) SetSourceIp(v string) *DownloadUserTotalFlowReportRequest {
	s.SourceIp = &v
	return s
}

func (s *DownloadUserTotalFlowReportRequest) SetLang(v string) *DownloadUserTotalFlowReportRequest {
	s.Lang = &v
	return s
}

func (s *DownloadUserTotalFlowReportRequest) SetStartTime(v int64) *DownloadUserTotalFlowReportRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadUserTotalFlowReportRequest) SetEndTime(v int64) *DownloadUserTotalFlowReportRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadUserTotalFlowReportRequest) SetInterval(v int32) *DownloadUserTotalFlowReportRequest {
	s.Interval = &v
	return s
}

type DownloadUserTotalFlowReportResponseBody struct {
	DownloadFileResult *DownloadUserTotalFlowReportResponseBodyDownloadFileResult `json:"DownloadFileResult,omitempty" xml:"DownloadFileResult,omitempty" type:"Struct"`
	RequestId          *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo         map[string]interface{}                                     `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s DownloadUserTotalFlowReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadUserTotalFlowReportResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadUserTotalFlowReportResponseBody) SetDownloadFileResult(v *DownloadUserTotalFlowReportResponseBodyDownloadFileResult) *DownloadUserTotalFlowReportResponseBody {
	s.DownloadFileResult = v
	return s
}

func (s *DownloadUserTotalFlowReportResponseBody) SetRequestId(v string) *DownloadUserTotalFlowReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadUserTotalFlowReportResponseBody) SetPromptInfo(v map[string]interface{}) *DownloadUserTotalFlowReportResponseBody {
	s.PromptInfo = v
	return s
}

type DownloadUserTotalFlowReportResponseBodyDownloadFileResult struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
}

func (s DownloadUserTotalFlowReportResponseBodyDownloadFileResult) String() string {
	return tea.Prettify(s)
}

func (s DownloadUserTotalFlowReportResponseBodyDownloadFileResult) GoString() string {
	return s.String()
}

func (s *DownloadUserTotalFlowReportResponseBodyDownloadFileResult) SetDownloadLink(v string) *DownloadUserTotalFlowReportResponseBodyDownloadFileResult {
	s.DownloadLink = &v
	return s
}

type DownloadUserTotalFlowReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadUserTotalFlowReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadUserTotalFlowReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadUserTotalFlowReportResponse) GoString() string {
	return s.String()
}

func (s *DownloadUserTotalFlowReportResponse) SetHeaders(v map[string]*string) *DownloadUserTotalFlowReportResponse {
	s.Headers = v
	return s
}

func (s *DownloadUserTotalFlowReportResponse) SetBody(v *DownloadUserTotalFlowReportResponseBody) *DownloadUserTotalFlowReportResponse {
	s.Body = v
	return s
}

type FlushLayer4RulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s FlushLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s FlushLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *FlushLayer4RulesRequest) SetSourceIp(v string) *FlushLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *FlushLayer4RulesRequest) SetLang(v string) *FlushLayer4RulesRequest {
	s.Lang = &v
	return s
}

func (s *FlushLayer4RulesRequest) SetBizId(v int64) *FlushLayer4RulesRequest {
	s.BizId = &v
	return s
}

type FlushLayer4RulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s FlushLayer4RulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlushLayer4RulesResponseBody) GoString() string {
	return s.String()
}

func (s *FlushLayer4RulesResponseBody) SetRequestId(v string) *FlushLayer4RulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlushLayer4RulesResponseBody) SetPromptInfo(v map[string]interface{}) *FlushLayer4RulesResponseBody {
	s.PromptInfo = v
	return s
}

type FlushLayer4RulesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FlushLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FlushLayer4RulesResponse) String() string {
	return tea.Prettify(s)
}

func (s FlushLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *FlushLayer4RulesResponse) SetHeaders(v map[string]*string) *FlushLayer4RulesResponse {
	s.Headers = v
	return s
}

func (s *FlushLayer4RulesResponse) SetBody(v *FlushLayer4RulesResponseBody) *FlushLayer4RulesResponse {
	s.Body = v
	return s
}

type ReallocNgResourceRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CcQps    *int32  `json:"CcQps,omitempty" xml:"CcQps,omitempty"`
}

func (s ReallocNgResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReallocNgResourceRequest) GoString() string {
	return s.String()
}

func (s *ReallocNgResourceRequest) SetSourceIp(v string) *ReallocNgResourceRequest {
	s.SourceIp = &v
	return s
}

func (s *ReallocNgResourceRequest) SetLang(v string) *ReallocNgResourceRequest {
	s.Lang = &v
	return s
}

func (s *ReallocNgResourceRequest) SetBizId(v string) *ReallocNgResourceRequest {
	s.BizId = &v
	return s
}

func (s *ReallocNgResourceRequest) SetCcQps(v int32) *ReallocNgResourceRequest {
	s.CcQps = &v
	return s
}

type ReallocNgResourceResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s ReallocNgResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReallocNgResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ReallocNgResourceResponseBody) SetRequestId(v string) *ReallocNgResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReallocNgResourceResponseBody) SetPromptInfo(v map[string]interface{}) *ReallocNgResourceResponseBody {
	s.PromptInfo = v
	return s
}

type ReallocNgResourceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReallocNgResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReallocNgResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReallocNgResourceResponse) GoString() string {
	return s.String()
}

func (s *ReallocNgResourceResponse) SetHeaders(v map[string]*string) *ReallocNgResourceResponse {
	s.Headers = v
	return s
}

func (s *ReallocNgResourceResponse) SetBody(v *ReallocNgResourceResponseBody) *ReallocNgResourceResponse {
	s.Body = v
	return s
}

type ReplaceCcRouteRulesRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	RouteList *string `json:"RouteList,omitempty" xml:"RouteList,omitempty"`
}

func (s ReplaceCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *ReplaceCcRouteRulesRequest) SetSourceIp(v string) *ReplaceCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *ReplaceCcRouteRulesRequest) SetLang(v string) *ReplaceCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *ReplaceCcRouteRulesRequest) SetBizId(v int64) *ReplaceCcRouteRulesRequest {
	s.BizId = &v
	return s
}

func (s *ReplaceCcRouteRulesRequest) SetRouteList(v string) *ReplaceCcRouteRulesRequest {
	s.RouteList = &v
	return s
}

type ReplaceCcRouteRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s ReplaceCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceCcRouteRulesResponseBody) SetRequestId(v string) *ReplaceCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *ReplaceCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type ReplaceCcRouteRulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplaceCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *ReplaceCcRouteRulesResponse) SetHeaders(v map[string]*string) *ReplaceCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *ReplaceCcRouteRulesResponse) SetBody(v *ReplaceCcRouteRulesResponseBody) *ReplaceCcRouteRulesResponse {
	s.Body = v
	return s
}

type SearchFlexFwdRulesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s SearchFlexFwdRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFlexFwdRulesRequest) GoString() string {
	return s.String()
}

func (s *SearchFlexFwdRulesRequest) SetSourceIp(v string) *SearchFlexFwdRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *SearchFlexFwdRulesRequest) SetLang(v string) *SearchFlexFwdRulesRequest {
	s.Lang = &v
	return s
}

func (s *SearchFlexFwdRulesRequest) SetEsnBizId(v int64) *SearchFlexFwdRulesRequest {
	s.EsnBizId = &v
	return s
}

func (s *SearchFlexFwdRulesRequest) SetKeyword(v string) *SearchFlexFwdRulesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchFlexFwdRulesRequest) SetPage(v int32) *SearchFlexFwdRulesRequest {
	s.Page = &v
	return s
}

func (s *SearchFlexFwdRulesRequest) SetPageSize(v int32) *SearchFlexFwdRulesRequest {
	s.PageSize = &v
	return s
}

type SearchFlexFwdRulesResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                        `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo   map[string]interface{}                        `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	FlexFwdRules []*SearchFlexFwdRulesResponseBodyFlexFwdRules `json:"FlexFwdRules,omitempty" xml:"FlexFwdRules,omitempty" type:"Repeated"`
}

func (s SearchFlexFwdRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFlexFwdRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFlexFwdRulesResponseBody) SetRequestId(v string) *SearchFlexFwdRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFlexFwdRulesResponseBody) SetTotal(v int64) *SearchFlexFwdRulesResponseBody {
	s.Total = &v
	return s
}

func (s *SearchFlexFwdRulesResponseBody) SetPromptInfo(v map[string]interface{}) *SearchFlexFwdRulesResponseBody {
	s.PromptInfo = v
	return s
}

func (s *SearchFlexFwdRulesResponseBody) SetFlexFwdRules(v []*SearchFlexFwdRulesResponseBodyFlexFwdRules) *SearchFlexFwdRulesResponseBody {
	s.FlexFwdRules = v
	return s
}

type SearchFlexFwdRulesResponseBodyFlexFwdRules struct {
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Protocol     *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
}

func (s SearchFlexFwdRulesResponseBodyFlexFwdRules) String() string {
	return tea.Prettify(s)
}

func (s SearchFlexFwdRulesResponseBodyFlexFwdRules) GoString() string {
	return s.String()
}

func (s *SearchFlexFwdRulesResponseBodyFlexFwdRules) SetMasterIpList(v string) *SearchFlexFwdRulesResponseBodyFlexFwdRules {
	s.MasterIpList = &v
	return s
}

func (s *SearchFlexFwdRulesResponseBodyFlexFwdRules) SetRemark(v string) *SearchFlexFwdRulesResponseBodyFlexFwdRules {
	s.Remark = &v
	return s
}

func (s *SearchFlexFwdRulesResponseBodyFlexFwdRules) SetSlaveIpList(v string) *SearchFlexFwdRulesResponseBodyFlexFwdRules {
	s.SlaveIpList = &v
	return s
}

func (s *SearchFlexFwdRulesResponseBodyFlexFwdRules) SetProtocol(v string) *SearchFlexFwdRulesResponseBodyFlexFwdRules {
	s.Protocol = &v
	return s
}

func (s *SearchFlexFwdRulesResponseBodyFlexFwdRules) SetIdentity(v string) *SearchFlexFwdRulesResponseBodyFlexFwdRules {
	s.Identity = &v
	return s
}

type SearchFlexFwdRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFlexFwdRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFlexFwdRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFlexFwdRulesResponse) GoString() string {
	return s.String()
}

func (s *SearchFlexFwdRulesResponse) SetHeaders(v map[string]*string) *SearchFlexFwdRulesResponse {
	s.Headers = v
	return s
}

func (s *SearchFlexFwdRulesResponse) SetBody(v *SearchFlexFwdRulesResponseBody) *SearchFlexFwdRulesResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppId    *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetSourceIp(v string) *UpdateAppRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateAppRequest) SetLang(v string) *UpdateAppRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAppRequest) SetAppName(v string) *UpdateAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppRequest) SetAppId(v int64) *UpdateAppRequest {
	s.AppId = &v
	return s
}

type UpdateAppResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateAppResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateCcBlackListRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId      *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ips        *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s UpdateCcBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcBlackListRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcBlackListRequest) SetSourceIp(v string) *UpdateCcBlackListRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcBlackListRequest) SetLang(v string) *UpdateCcBlackListRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcBlackListRequest) SetBizId(v int64) *UpdateCcBlackListRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcBlackListRequest) SetIps(v string) *UpdateCcBlackListRequest {
	s.Ips = &v
	return s
}

func (s *UpdateCcBlackListRequest) SetActionType(v string) *UpdateCcBlackListRequest {
	s.ActionType = &v
	return s
}

type UpdateCcBlackListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcBlackListResponseBody) SetRequestId(v string) *UpdateCcBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcBlackListResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcBlackListResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcBlackListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcBlackListResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcBlackListResponse) SetHeaders(v map[string]*string) *UpdateCcBlackListResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcBlackListResponse) SetBody(v *UpdateCcBlackListResponseBody) *UpdateCcBlackListResponse {
	s.Body = v
	return s
}

type UpdateCcIDCBlockSwitchRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Enable   *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateCcIDCBlockSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcIDCBlockSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcIDCBlockSwitchRequest) SetSourceIp(v string) *UpdateCcIDCBlockSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcIDCBlockSwitchRequest) SetLang(v string) *UpdateCcIDCBlockSwitchRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcIDCBlockSwitchRequest) SetBizId(v int64) *UpdateCcIDCBlockSwitchRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcIDCBlockSwitchRequest) SetEnable(v int32) *UpdateCcIDCBlockSwitchRequest {
	s.Enable = &v
	return s
}

type UpdateCcIDCBlockSwitchResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcIDCBlockSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcIDCBlockSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcIDCBlockSwitchResponseBody) SetRequestId(v string) *UpdateCcIDCBlockSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcIDCBlockSwitchResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcIDCBlockSwitchResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcIDCBlockSwitchResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcIDCBlockSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcIDCBlockSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcIDCBlockSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcIDCBlockSwitchResponse) SetHeaders(v map[string]*string) *UpdateCcIDCBlockSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcIDCBlockSwitchResponse) SetBody(v *UpdateCcIDCBlockSwitchResponseBody) *UpdateCcIDCBlockSwitchResponse {
	s.Body = v
	return s
}

type UpdateCcRouteRulesRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Routelist *string `json:"Routelist,omitempty" xml:"Routelist,omitempty"`
}

func (s UpdateCcRouteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcRouteRulesRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcRouteRulesRequest) SetSourceIp(v string) *UpdateCcRouteRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcRouteRulesRequest) SetLang(v string) *UpdateCcRouteRulesRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcRouteRulesRequest) SetBizId(v int64) *UpdateCcRouteRulesRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcRouteRulesRequest) SetRoutelist(v string) *UpdateCcRouteRulesRequest {
	s.Routelist = &v
	return s
}

type UpdateCcRouteRulesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcRouteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcRouteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcRouteRulesResponseBody) SetRequestId(v string) *UpdateCcRouteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcRouteRulesResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcRouteRulesResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcRouteRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcRouteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcRouteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcRouteRulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcRouteRulesResponse) SetHeaders(v map[string]*string) *UpdateCcRouteRulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcRouteRulesResponse) SetBody(v *UpdateCcRouteRulesResponseBody) *UpdateCcRouteRulesResponse {
	s.Body = v
	return s
}

type UpdateCcRouteSwitchRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Enable   *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateCcRouteSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcRouteSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcRouteSwitchRequest) SetSourceIp(v string) *UpdateCcRouteSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcRouteSwitchRequest) SetLang(v string) *UpdateCcRouteSwitchRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcRouteSwitchRequest) SetBizId(v int64) *UpdateCcRouteSwitchRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcRouteSwitchRequest) SetEnable(v int32) *UpdateCcRouteSwitchRequest {
	s.Enable = &v
	return s
}

type UpdateCcRouteSwitchResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcRouteSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcRouteSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcRouteSwitchResponseBody) SetRequestId(v string) *UpdateCcRouteSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcRouteSwitchResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcRouteSwitchResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcRouteSwitchResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcRouteSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcRouteSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcRouteSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcRouteSwitchResponse) SetHeaders(v map[string]*string) *UpdateCcRouteSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcRouteSwitchResponse) SetBody(v *UpdateCcRouteSwitchResponseBody) *UpdateCcRouteSwitchResponse {
	s.Body = v
	return s
}

type UpdateCcTunnelGrayAndOnlyAllowRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OnlyAllow *int32  `json:"OnlyAllow,omitempty" xml:"OnlyAllow,omitempty"`
	Gray      *int32  `json:"Gray,omitempty" xml:"Gray,omitempty"`
}

func (s UpdateCcTunnelGrayAndOnlyAllowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcTunnelGrayAndOnlyAllowRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcTunnelGrayAndOnlyAllowRequest) SetSourceIp(v string) *UpdateCcTunnelGrayAndOnlyAllowRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcTunnelGrayAndOnlyAllowRequest) SetLang(v string) *UpdateCcTunnelGrayAndOnlyAllowRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcTunnelGrayAndOnlyAllowRequest) SetBizId(v string) *UpdateCcTunnelGrayAndOnlyAllowRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcTunnelGrayAndOnlyAllowRequest) SetOnlyAllow(v int32) *UpdateCcTunnelGrayAndOnlyAllowRequest {
	s.OnlyAllow = &v
	return s
}

func (s *UpdateCcTunnelGrayAndOnlyAllowRequest) SetGray(v int32) *UpdateCcTunnelGrayAndOnlyAllowRequest {
	s.Gray = &v
	return s
}

type UpdateCcTunnelGrayAndOnlyAllowResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcTunnelGrayAndOnlyAllowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcTunnelGrayAndOnlyAllowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcTunnelGrayAndOnlyAllowResponseBody) SetRequestId(v string) *UpdateCcTunnelGrayAndOnlyAllowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcTunnelGrayAndOnlyAllowResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcTunnelGrayAndOnlyAllowResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcTunnelGrayAndOnlyAllowResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcTunnelGrayAndOnlyAllowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcTunnelGrayAndOnlyAllowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcTunnelGrayAndOnlyAllowResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcTunnelGrayAndOnlyAllowResponse) SetHeaders(v map[string]*string) *UpdateCcTunnelGrayAndOnlyAllowResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcTunnelGrayAndOnlyAllowResponse) SetBody(v *UpdateCcTunnelGrayAndOnlyAllowResponseBody) *UpdateCcTunnelGrayAndOnlyAllowResponse {
	s.Body = v
	return s
}

type UpdateCcTunnelStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateCcTunnelStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcTunnelStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcTunnelStatusRequest) SetSourceIp(v string) *UpdateCcTunnelStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcTunnelStatusRequest) SetLang(v string) *UpdateCcTunnelStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcTunnelStatusRequest) SetBizId(v string) *UpdateCcTunnelStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcTunnelStatusRequest) SetStatus(v string) *UpdateCcTunnelStatusRequest {
	s.Status = &v
	return s
}

type UpdateCcTunnelStatusResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcTunnelStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcTunnelStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcTunnelStatusResponseBody) SetRequestId(v string) *UpdateCcTunnelStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcTunnelStatusResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcTunnelStatusResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcTunnelStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcTunnelStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcTunnelStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcTunnelStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcTunnelStatusResponse) SetHeaders(v map[string]*string) *UpdateCcTunnelStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcTunnelStatusResponse) SetBody(v *UpdateCcTunnelStatusResponseBody) *UpdateCcTunnelStatusResponse {
	s.Body = v
	return s
}

type UpdateCcUseRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId      *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	CcQps      *int32  `json:"CcQps,omitempty" xml:"CcQps,omitempty"`
}

func (s UpdateCcUseRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcUseRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcUseRequest) SetSourceIp(v string) *UpdateCcUseRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcUseRequest) SetLang(v string) *UpdateCcUseRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcUseRequest) SetBizId(v int64) *UpdateCcUseRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcUseRequest) SetActionType(v string) *UpdateCcUseRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateCcUseRequest) SetCcQps(v int32) *UpdateCcUseRequest {
	s.CcQps = &v
	return s
}

type UpdateCcUseResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcUseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcUseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcUseResponseBody) SetRequestId(v string) *UpdateCcUseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcUseResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcUseResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcUseResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcUseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcUseResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcUseResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcUseResponse) SetHeaders(v map[string]*string) *UpdateCcUseResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcUseResponse) SetBody(v *UpdateCcUseResponseBody) *UpdateCcUseResponse {
	s.Body = v
	return s
}

type UpdateCcWhiteListRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId      *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ips        *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s UpdateCcWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcWhiteListRequest) SetSourceIp(v string) *UpdateCcWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcWhiteListRequest) SetLang(v string) *UpdateCcWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcWhiteListRequest) SetBizId(v int64) *UpdateCcWhiteListRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcWhiteListRequest) SetIps(v string) *UpdateCcWhiteListRequest {
	s.Ips = &v
	return s
}

func (s *UpdateCcWhiteListRequest) SetActionType(v string) *UpdateCcWhiteListRequest {
	s.ActionType = &v
	return s
}

type UpdateCcWhiteListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcWhiteListResponseBody) SetRequestId(v string) *UpdateCcWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcWhiteListResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcWhiteListResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcWhiteListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcWhiteListResponse) SetHeaders(v map[string]*string) *UpdateCcWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcWhiteListResponse) SetBody(v *UpdateCcWhiteListResponseBody) *UpdateCcWhiteListResponse {
	s.Body = v
	return s
}

type UpdateCcZoneBlockConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Config   *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s UpdateCcZoneBlockConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcZoneBlockConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcZoneBlockConfigRequest) SetSourceIp(v string) *UpdateCcZoneBlockConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcZoneBlockConfigRequest) SetLang(v string) *UpdateCcZoneBlockConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcZoneBlockConfigRequest) SetBizId(v int64) *UpdateCcZoneBlockConfigRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcZoneBlockConfigRequest) SetConfig(v string) *UpdateCcZoneBlockConfigRequest {
	s.Config = &v
	return s
}

type UpdateCcZoneBlockConfigResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcZoneBlockConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcZoneBlockConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcZoneBlockConfigResponseBody) SetRequestId(v string) *UpdateCcZoneBlockConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcZoneBlockConfigResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcZoneBlockConfigResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcZoneBlockConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcZoneBlockConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcZoneBlockConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcZoneBlockConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcZoneBlockConfigResponse) SetHeaders(v map[string]*string) *UpdateCcZoneBlockConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcZoneBlockConfigResponse) SetBody(v *UpdateCcZoneBlockConfigResponseBody) *UpdateCcZoneBlockConfigResponse {
	s.Body = v
	return s
}

type UpdateCcZoneBlockStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Enable   *int32  `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateCcZoneBlockStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcZoneBlockStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcZoneBlockStatusRequest) SetSourceIp(v string) *UpdateCcZoneBlockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcZoneBlockStatusRequest) SetLang(v string) *UpdateCcZoneBlockStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcZoneBlockStatusRequest) SetBizId(v int64) *UpdateCcZoneBlockStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateCcZoneBlockStatusRequest) SetEnable(v int32) *UpdateCcZoneBlockStatusRequest {
	s.Enable = &v
	return s
}

type UpdateCcZoneBlockStatusResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateCcZoneBlockStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcZoneBlockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcZoneBlockStatusResponseBody) SetRequestId(v string) *UpdateCcZoneBlockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCcZoneBlockStatusResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateCcZoneBlockStatusResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateCcZoneBlockStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcZoneBlockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcZoneBlockStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcZoneBlockStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcZoneBlockStatusResponse) SetHeaders(v map[string]*string) *UpdateCcZoneBlockStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcZoneBlockStatusResponse) SetBody(v *UpdateCcZoneBlockStatusResponseBody) *UpdateCcZoneBlockStatusResponse {
	s.Body = v
	return s
}

type UpdateFlexAccFwdRuleRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId   *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Identity   *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	ProtocolEx *string `json:"ProtocolEx,omitempty" xml:"ProtocolEx,omitempty"`
	IpList     *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateFlexAccFwdRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccFwdRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccFwdRuleRequest) SetSourceIp(v string) *UpdateFlexAccFwdRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexAccFwdRuleRequest) SetEsnBizId(v int64) *UpdateFlexAccFwdRuleRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexAccFwdRuleRequest) SetIdentity(v string) *UpdateFlexAccFwdRuleRequest {
	s.Identity = &v
	return s
}

func (s *UpdateFlexAccFwdRuleRequest) SetProtocolEx(v string) *UpdateFlexAccFwdRuleRequest {
	s.ProtocolEx = &v
	return s
}

func (s *UpdateFlexAccFwdRuleRequest) SetIpList(v string) *UpdateFlexAccFwdRuleRequest {
	s.IpList = &v
	return s
}

func (s *UpdateFlexAccFwdRuleRequest) SetDomainList(v string) *UpdateFlexAccFwdRuleRequest {
	s.DomainList = &v
	return s
}

func (s *UpdateFlexAccFwdRuleRequest) SetRemark(v string) *UpdateFlexAccFwdRuleRequest {
	s.Remark = &v
	return s
}

type UpdateFlexAccFwdRuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexAccFwdRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccFwdRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccFwdRuleResponseBody) SetRequestId(v string) *UpdateFlexAccFwdRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexAccFwdRuleResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexAccFwdRuleResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexAccFwdRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexAccFwdRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexAccFwdRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccFwdRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccFwdRuleResponse) SetHeaders(v map[string]*string) *UpdateFlexAccFwdRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexAccFwdRuleResponse) SetBody(v *UpdateFlexAccFwdRuleResponseBody) *UpdateFlexAccFwdRuleResponse {
	s.Body = v
	return s
}

type UpdateFlexAccFwdRuleV2Request struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BizId        *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	IpList       *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	ProtocolEx   *string `json:"ProtocolEx,omitempty" xml:"ProtocolEx,omitempty"`
	DomainList   *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	AccType      *int32  `json:"AccType,omitempty" xml:"AccType,omitempty"`
}

func (s UpdateFlexAccFwdRuleV2Request) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccFwdRuleV2Request) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccFwdRuleV2Request) SetSourceIp(v string) *UpdateFlexAccFwdRuleV2Request {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetBizId(v int64) *UpdateFlexAccFwdRuleV2Request {
	s.BizId = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetIdentity(v string) *UpdateFlexAccFwdRuleV2Request {
	s.Identity = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetIpList(v string) *UpdateFlexAccFwdRuleV2Request {
	s.IpList = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetProtocolEx(v string) *UpdateFlexAccFwdRuleV2Request {
	s.ProtocolEx = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetDomainList(v string) *UpdateFlexAccFwdRuleV2Request {
	s.DomainList = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetRemark(v string) *UpdateFlexAccFwdRuleV2Request {
	s.Remark = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetMasterIpList(v string) *UpdateFlexAccFwdRuleV2Request {
	s.MasterIpList = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetSlaveIpList(v string) *UpdateFlexAccFwdRuleV2Request {
	s.SlaveIpList = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Request) SetAccType(v int32) *UpdateFlexAccFwdRuleV2Request {
	s.AccType = &v
	return s
}

type UpdateFlexAccFwdRuleV2ResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexAccFwdRuleV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccFwdRuleV2ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccFwdRuleV2ResponseBody) SetRequestId(v string) *UpdateFlexAccFwdRuleV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexAccFwdRuleV2ResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexAccFwdRuleV2ResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexAccFwdRuleV2Response struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexAccFwdRuleV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexAccFwdRuleV2Response) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccFwdRuleV2Response) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccFwdRuleV2Response) SetHeaders(v map[string]*string) *UpdateFlexAccFwdRuleV2Response {
	s.Headers = v
	return s
}

func (s *UpdateFlexAccFwdRuleV2Response) SetBody(v *UpdateFlexAccFwdRuleV2ResponseBody) *UpdateFlexAccFwdRuleV2Response {
	s.Body = v
	return s
}

type UpdateFlexAccStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFlexAccStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccStatusRequest) SetSourceIp(v string) *UpdateFlexAccStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexAccStatusRequest) SetBizId(v int64) *UpdateFlexAccStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateFlexAccStatusRequest) SetStatus(v string) *UpdateFlexAccStatusRequest {
	s.Status = &v
	return s
}

type UpdateFlexAccStatusResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexAccStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccStatusResponseBody) SetRequestId(v string) *UpdateFlexAccStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexAccStatusResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexAccStatusResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexAccStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexAccStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexAccStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccStatusResponse) SetHeaders(v map[string]*string) *UpdateFlexAccStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexAccStatusResponse) SetBody(v *UpdateFlexAccStatusResponseBody) *UpdateFlexAccStatusResponse {
	s.Body = v
	return s
}

type UpdateFlexAccTcpPortsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Ports    *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
}

func (s UpdateFlexAccTcpPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccTcpPortsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccTcpPortsRequest) SetSourceIp(v string) *UpdateFlexAccTcpPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexAccTcpPortsRequest) SetEsnBizId(v int64) *UpdateFlexAccTcpPortsRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexAccTcpPortsRequest) SetPorts(v string) *UpdateFlexAccTcpPortsRequest {
	s.Ports = &v
	return s
}

type UpdateFlexAccTcpPortsResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexAccTcpPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccTcpPortsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccTcpPortsResponseBody) SetRequestId(v string) *UpdateFlexAccTcpPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexAccTcpPortsResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexAccTcpPortsResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexAccTcpPortsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexAccTcpPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexAccTcpPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccTcpPortsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccTcpPortsResponse) SetHeaders(v map[string]*string) *UpdateFlexAccTcpPortsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexAccTcpPortsResponse) SetBody(v *UpdateFlexAccTcpPortsResponseBody) *UpdateFlexAccTcpPortsResponse {
	s.Body = v
	return s
}

type UpdateFlexAccUdpPortsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Ports    *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
}

func (s UpdateFlexAccUdpPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccUdpPortsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccUdpPortsRequest) SetSourceIp(v string) *UpdateFlexAccUdpPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexAccUdpPortsRequest) SetEsnBizId(v int64) *UpdateFlexAccUdpPortsRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexAccUdpPortsRequest) SetPorts(v string) *UpdateFlexAccUdpPortsRequest {
	s.Ports = &v
	return s
}

type UpdateFlexAccUdpPortsResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexAccUdpPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccUdpPortsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccUdpPortsResponseBody) SetRequestId(v string) *UpdateFlexAccUdpPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexAccUdpPortsResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexAccUdpPortsResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexAccUdpPortsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexAccUdpPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexAccUdpPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexAccUdpPortsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexAccUdpPortsResponse) SetHeaders(v map[string]*string) *UpdateFlexAccUdpPortsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexAccUdpPortsResponse) SetBody(v *UpdateFlexAccUdpPortsResponseBody) *UpdateFlexAccUdpPortsResponse {
	s.Body = v
	return s
}

type UpdateFlexBackupGroupsRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId          *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BackupGroupPolicy *int32  `json:"BackupGroupPolicy,omitempty" xml:"BackupGroupPolicy,omitempty"`
	Backups           *string `json:"Backups,omitempty" xml:"Backups,omitempty"`
	SharedBackups     *string `json:"SharedBackups,omitempty" xml:"SharedBackups,omitempty"`
}

func (s UpdateFlexBackupGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexBackupGroupsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexBackupGroupsRequest) SetSourceIp(v string) *UpdateFlexBackupGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexBackupGroupsRequest) SetLang(v string) *UpdateFlexBackupGroupsRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFlexBackupGroupsRequest) SetEsnBizId(v int64) *UpdateFlexBackupGroupsRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexBackupGroupsRequest) SetGroupId(v string) *UpdateFlexBackupGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateFlexBackupGroupsRequest) SetBackupGroupPolicy(v int32) *UpdateFlexBackupGroupsRequest {
	s.BackupGroupPolicy = &v
	return s
}

func (s *UpdateFlexBackupGroupsRequest) SetBackups(v string) *UpdateFlexBackupGroupsRequest {
	s.Backups = &v
	return s
}

func (s *UpdateFlexBackupGroupsRequest) SetSharedBackups(v string) *UpdateFlexBackupGroupsRequest {
	s.SharedBackups = &v
	return s
}

type UpdateFlexBackupGroupsResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexBackupGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexBackupGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexBackupGroupsResponseBody) SetRequestId(v string) *UpdateFlexBackupGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexBackupGroupsResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexBackupGroupsResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexBackupGroupsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexBackupGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexBackupGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexBackupGroupsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexBackupGroupsResponse) SetHeaders(v map[string]*string) *UpdateFlexBackupGroupsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexBackupGroupsResponse) SetBody(v *UpdateFlexBackupGroupsResponseBody) *UpdateFlexBackupGroupsResponse {
	s.Body = v
	return s
}

type UpdateFlexFwdRuleRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId        *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateFlexFwdRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexFwdRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexFwdRuleRequest) SetSourceIp(v string) *UpdateFlexFwdRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexFwdRuleRequest) SetLang(v string) *UpdateFlexFwdRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFlexFwdRuleRequest) SetBizId(v int64) *UpdateFlexFwdRuleRequest {
	s.BizId = &v
	return s
}

func (s *UpdateFlexFwdRuleRequest) SetIdentity(v string) *UpdateFlexFwdRuleRequest {
	s.Identity = &v
	return s
}

func (s *UpdateFlexFwdRuleRequest) SetMasterIpList(v string) *UpdateFlexFwdRuleRequest {
	s.MasterIpList = &v
	return s
}

func (s *UpdateFlexFwdRuleRequest) SetSlaveIpList(v string) *UpdateFlexFwdRuleRequest {
	s.SlaveIpList = &v
	return s
}

func (s *UpdateFlexFwdRuleRequest) SetRemark(v string) *UpdateFlexFwdRuleRequest {
	s.Remark = &v
	return s
}

type UpdateFlexFwdRuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexFwdRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexFwdRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexFwdRuleResponseBody) SetRequestId(v string) *UpdateFlexFwdRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexFwdRuleResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexFwdRuleResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexFwdRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexFwdRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexFwdRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexFwdRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexFwdRuleResponse) SetHeaders(v map[string]*string) *UpdateFlexFwdRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexFwdRuleResponse) SetBody(v *UpdateFlexFwdRuleResponseBody) *UpdateFlexFwdRuleResponse {
	s.Body = v
	return s
}

type UpdateFlexInnerPolicyRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId         *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupInnerPolicy *int32  `json:"GroupInnerPolicy,omitempty" xml:"GroupInnerPolicy,omitempty"`
}

func (s UpdateFlexInnerPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexInnerPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexInnerPolicyRequest) SetSourceIp(v string) *UpdateFlexInnerPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexInnerPolicyRequest) SetLang(v string) *UpdateFlexInnerPolicyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFlexInnerPolicyRequest) SetEsnBizId(v int64) *UpdateFlexInnerPolicyRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexInnerPolicyRequest) SetGroupId(v string) *UpdateFlexInnerPolicyRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateFlexInnerPolicyRequest) SetGroupInnerPolicy(v int32) *UpdateFlexInnerPolicyRequest {
	s.GroupInnerPolicy = &v
	return s
}

type UpdateFlexInnerPolicyResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexInnerPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexInnerPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexInnerPolicyResponseBody) SetRequestId(v string) *UpdateFlexInnerPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexInnerPolicyResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexInnerPolicyResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexInnerPolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexInnerPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexInnerPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexInnerPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexInnerPolicyResponse) SetHeaders(v map[string]*string) *UpdateFlexInnerPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexInnerPolicyResponse) SetBody(v *UpdateFlexInnerPolicyResponseBody) *UpdateFlexInnerPolicyResponse {
	s.Body = v
	return s
}

type UpdateFlexLinkTypeRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	LinkType *int32  `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
}

func (s UpdateFlexLinkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexLinkTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexLinkTypeRequest) SetSourceIp(v string) *UpdateFlexLinkTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexLinkTypeRequest) SetLang(v string) *UpdateFlexLinkTypeRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFlexLinkTypeRequest) SetEsnBizId(v int64) *UpdateFlexLinkTypeRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexLinkTypeRequest) SetLinkType(v int32) *UpdateFlexLinkTypeRequest {
	s.LinkType = &v
	return s
}

type UpdateFlexLinkTypeResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexLinkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexLinkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexLinkTypeResponseBody) SetRequestId(v string) *UpdateFlexLinkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexLinkTypeResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexLinkTypeResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexLinkTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexLinkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexLinkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexLinkTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexLinkTypeResponse) SetHeaders(v map[string]*string) *UpdateFlexLinkTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexLinkTypeResponse) SetBody(v *UpdateFlexLinkTypeResponseBody) *UpdateFlexLinkTypeResponse {
	s.Body = v
	return s
}

type UpdateFlexPortsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	EsnBizId *int64  `json:"EsnBizId,omitempty" xml:"EsnBizId,omitempty"`
	Ports    *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
}

func (s UpdateFlexPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexPortsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexPortsRequest) SetSourceIp(v string) *UpdateFlexPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexPortsRequest) SetLang(v string) *UpdateFlexPortsRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFlexPortsRequest) SetEsnBizId(v int64) *UpdateFlexPortsRequest {
	s.EsnBizId = &v
	return s
}

func (s *UpdateFlexPortsRequest) SetPorts(v string) *UpdateFlexPortsRequest {
	s.Ports = &v
	return s
}

type UpdateFlexPortsResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexPortsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexPortsResponseBody) SetRequestId(v string) *UpdateFlexPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexPortsResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexPortsResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexPortsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexPortsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexPortsResponse) SetHeaders(v map[string]*string) *UpdateFlexPortsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexPortsResponse) SetBody(v *UpdateFlexPortsResponseBody) *UpdateFlexPortsResponse {
	s.Body = v
	return s
}

type UpdateFlexStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFlexStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlexStatusRequest) SetSourceIp(v string) *UpdateFlexStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateFlexStatusRequest) SetLang(v string) *UpdateFlexStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateFlexStatusRequest) SetBizId(v int64) *UpdateFlexStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateFlexStatusRequest) SetStatus(v string) *UpdateFlexStatusRequest {
	s.Status = &v
	return s
}

type UpdateFlexStatusResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateFlexStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlexStatusResponseBody) SetRequestId(v string) *UpdateFlexStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlexStatusResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateFlexStatusResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateFlexStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlexStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlexStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlexStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlexStatusResponse) SetHeaders(v map[string]*string) *UpdateFlexStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlexStatusResponse) SetBody(v *UpdateFlexStatusResponseBody) *UpdateFlexStatusResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupDesc *string `json:"GroupDesc,omitempty" xml:"GroupDesc,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetSourceIp(v string) *UpdateGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateGroupRequest) SetLang(v string) *UpdateGroupRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGroupRequest) SetBizId(v int64) *UpdateGroupRequest {
	s.BizId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupId(v string) *UpdateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupDesc(v string) *UpdateGroupRequest {
	s.GroupDesc = &v
	return s
}

type UpdateGroupResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateGroupResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateGroupDnsStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateGroupDnsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDnsStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDnsStatusRequest) SetSourceIp(v string) *UpdateGroupDnsStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateGroupDnsStatusRequest) SetLang(v string) *UpdateGroupDnsStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGroupDnsStatusRequest) SetBizId(v int64) *UpdateGroupDnsStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateGroupDnsStatusRequest) SetGroupId(v string) *UpdateGroupDnsStatusRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupDnsStatusRequest) SetStatus(v string) *UpdateGroupDnsStatusRequest {
	s.Status = &v
	return s
}

type UpdateGroupDnsStatusResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateGroupDnsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDnsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDnsStatusResponseBody) SetRequestId(v string) *UpdateGroupDnsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupDnsStatusResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateGroupDnsStatusResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateGroupDnsStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupDnsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupDnsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDnsStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDnsStatusResponse) SetHeaders(v map[string]*string) *UpdateGroupDnsStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDnsStatusResponse) SetBody(v *UpdateGroupDnsStatusResponseBody) *UpdateGroupDnsStatusResponse {
	s.Body = v
	return s
}

type UpdateGroupNodesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IpList   *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s UpdateGroupNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNodesRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupNodesRequest) SetSourceIp(v string) *UpdateGroupNodesRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateGroupNodesRequest) SetLang(v string) *UpdateGroupNodesRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGroupNodesRequest) SetBizId(v int64) *UpdateGroupNodesRequest {
	s.BizId = &v
	return s
}

func (s *UpdateGroupNodesRequest) SetGroupId(v string) *UpdateGroupNodesRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupNodesRequest) SetIpList(v string) *UpdateGroupNodesRequest {
	s.IpList = &v
	return s
}

type UpdateGroupNodesResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateGroupNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNodesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupNodesResponseBody) SetRequestId(v string) *UpdateGroupNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupNodesResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateGroupNodesResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateGroupNodesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNodesResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupNodesResponse) SetHeaders(v map[string]*string) *UpdateGroupNodesResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupNodesResponse) SetBody(v *UpdateGroupNodesResponseBody) *UpdateGroupNodesResponse {
	s.Body = v
	return s
}

type UpdateGroupStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateGroupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupStatusRequest) SetSourceIp(v string) *UpdateGroupStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateGroupStatusRequest) SetLang(v string) *UpdateGroupStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGroupStatusRequest) SetBizId(v int64) *UpdateGroupStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateGroupStatusRequest) SetGroupId(v string) *UpdateGroupStatusRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupStatusRequest) SetStatus(v string) *UpdateGroupStatusRequest {
	s.Status = &v
	return s
}

type UpdateGroupStatusResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateGroupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupStatusResponseBody) SetRequestId(v string) *UpdateGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupStatusResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateGroupStatusResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateGroupStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupStatusResponse) SetHeaders(v map[string]*string) *UpdateGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupStatusResponse) SetBody(v *UpdateGroupStatusResponseBody) *UpdateGroupStatusResponse {
	s.Body = v
	return s
}

type UpdateLayer4RuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId     *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FrontPort *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	Rservers  *string `json:"Rservers,omitempty" xml:"Rservers,omitempty"`
}

func (s UpdateLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateLayer4RuleRequest) SetSourceIp(v string) *UpdateLayer4RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateLayer4RuleRequest) SetLang(v string) *UpdateLayer4RuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateLayer4RuleRequest) SetBizId(v int64) *UpdateLayer4RuleRequest {
	s.BizId = &v
	return s
}

func (s *UpdateLayer4RuleRequest) SetFrontPort(v int32) *UpdateLayer4RuleRequest {
	s.FrontPort = &v
	return s
}

func (s *UpdateLayer4RuleRequest) SetRservers(v string) *UpdateLayer4RuleRequest {
	s.Rservers = &v
	return s
}

type UpdateLayer4RuleResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UpdateLayer4RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLayer4RuleResponseBody) SetRequestId(v string) *UpdateLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLayer4RuleResponseBody) SetPromptInfo(v map[string]interface{}) *UpdateLayer4RuleResponseBody {
	s.PromptInfo = v
	return s
}

type UpdateLayer4RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateLayer4RuleResponse) SetHeaders(v map[string]*string) *UpdateLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateLayer4RuleResponse) SetBody(v *UpdateLayer4RuleResponseBody) *UpdateLayer4RuleResponse {
	s.Body = v
	return s
}

type UploadCcRouteFileForParseRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
}

func (s UploadCcRouteFileForParseRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCcRouteFileForParseRequest) GoString() string {
	return s.String()
}

func (s *UploadCcRouteFileForParseRequest) SetSourceIp(v string) *UploadCcRouteFileForParseRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadCcRouteFileForParseRequest) SetLang(v string) *UploadCcRouteFileForParseRequest {
	s.Lang = &v
	return s
}

func (s *UploadCcRouteFileForParseRequest) SetFileKey(v string) *UploadCcRouteFileForParseRequest {
	s.FileKey = &v
	return s
}

type UploadCcRouteFileForParseResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	RouteRules []*UploadCcRouteFileForParseResponseBodyRouteRules `json:"RouteRules,omitempty" xml:"RouteRules,omitempty" type:"Repeated"`
}

func (s UploadCcRouteFileForParseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCcRouteFileForParseResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCcRouteFileForParseResponseBody) SetRequestId(v string) *UploadCcRouteFileForParseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCcRouteFileForParseResponseBody) SetTotal(v int64) *UploadCcRouteFileForParseResponseBody {
	s.Total = &v
	return s
}

func (s *UploadCcRouteFileForParseResponseBody) SetPromptInfo(v map[string]interface{}) *UploadCcRouteFileForParseResponseBody {
	s.PromptInfo = v
	return s
}

func (s *UploadCcRouteFileForParseResponseBody) SetRouteRules(v []*UploadCcRouteFileForParseResponseBodyRouteRules) *UploadCcRouteFileForParseResponseBody {
	s.RouteRules = v
	return s
}

type UploadCcRouteFileForParseResponseBodyRouteRules struct {
	Comment  *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Rservers []*string `json:"Rservers,omitempty" xml:"Rservers,omitempty" type:"Repeated"`
	Id       *string   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UploadCcRouteFileForParseResponseBodyRouteRules) String() string {
	return tea.Prettify(s)
}

func (s UploadCcRouteFileForParseResponseBodyRouteRules) GoString() string {
	return s.String()
}

func (s *UploadCcRouteFileForParseResponseBodyRouteRules) SetComment(v string) *UploadCcRouteFileForParseResponseBodyRouteRules {
	s.Comment = &v
	return s
}

func (s *UploadCcRouteFileForParseResponseBodyRouteRules) SetRservers(v []*string) *UploadCcRouteFileForParseResponseBodyRouteRules {
	s.Rservers = v
	return s
}

func (s *UploadCcRouteFileForParseResponseBodyRouteRules) SetId(v string) *UploadCcRouteFileForParseResponseBodyRouteRules {
	s.Id = &v
	return s
}

type UploadCcRouteFileForParseResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadCcRouteFileForParseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadCcRouteFileForParseResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCcRouteFileForParseResponse) GoString() string {
	return s.String()
}

func (s *UploadCcRouteFileForParseResponse) SetHeaders(v map[string]*string) *UploadCcRouteFileForParseResponse {
	s.Headers = v
	return s
}

func (s *UploadCcRouteFileForParseResponse) SetBody(v *UploadCcRouteFileForParseResponseBody) *UploadCcRouteFileForParseResponse {
	s.Body = v
	return s
}

type UploadCcWhiteBlackListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	BizId    *int64  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BWList   *string `json:"BWList,omitempty" xml:"BWList,omitempty"`
}

func (s UploadCcWhiteBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCcWhiteBlackListRequest) GoString() string {
	return s.String()
}

func (s *UploadCcWhiteBlackListRequest) SetSourceIp(v string) *UploadCcWhiteBlackListRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadCcWhiteBlackListRequest) SetLang(v string) *UploadCcWhiteBlackListRequest {
	s.Lang = &v
	return s
}

func (s *UploadCcWhiteBlackListRequest) SetBizId(v int64) *UploadCcWhiteBlackListRequest {
	s.BizId = &v
	return s
}

func (s *UploadCcWhiteBlackListRequest) SetBWList(v string) *UploadCcWhiteBlackListRequest {
	s.BWList = &v
	return s
}

type UploadCcWhiteBlackListResponseBody struct {
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PromptInfo map[string]interface{} `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UploadCcWhiteBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCcWhiteBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCcWhiteBlackListResponseBody) SetRequestId(v string) *UploadCcWhiteBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCcWhiteBlackListResponseBody) SetPromptInfo(v map[string]interface{}) *UploadCcWhiteBlackListResponseBody {
	s.PromptInfo = v
	return s
}

type UploadCcWhiteBlackListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadCcWhiteBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadCcWhiteBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCcWhiteBlackListResponse) GoString() string {
	return s.String()
}

func (s *UploadCcWhiteBlackListResponse) SetHeaders(v map[string]*string) *UploadCcWhiteBlackListResponse {
	s.Headers = v
	return s
}

func (s *UploadCcWhiteBlackListResponse) SetBody(v *UploadCcWhiteBlackListResponseBody) *UploadCcWhiteBlackListResponse {
	s.Body = v
	return s
}

type UploadFlexAccRulesFileForParseRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
}

func (s UploadFlexAccRulesFileForParseRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexAccRulesFileForParseRequest) GoString() string {
	return s.String()
}

func (s *UploadFlexAccRulesFileForParseRequest) SetSourceIp(v string) *UploadFlexAccRulesFileForParseRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseRequest) SetFileKey(v string) *UploadFlexAccRulesFileForParseRequest {
	s.FileKey = &v
	return s
}

type UploadFlexAccRulesFileForParseResponseBody struct {
	FlexAccFwdRules []*UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules `json:"FlexAccFwdRules,omitempty" xml:"FlexAccFwdRules,omitempty" type:"Repeated"`
	RequestId       *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total           *int64                                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo      map[string]interface{}                                       `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
}

func (s UploadFlexAccRulesFileForParseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexAccRulesFileForParseResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFlexAccRulesFileForParseResponseBody) SetFlexAccFwdRules(v []*UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) *UploadFlexAccRulesFileForParseResponseBody {
	s.FlexAccFwdRules = v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBody) SetRequestId(v string) *UploadFlexAccRulesFileForParseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBody) SetTotal(v int64) *UploadFlexAccRulesFileForParseResponseBody {
	s.Total = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBody) SetPromptInfo(v map[string]interface{}) *UploadFlexAccRulesFileForParseResponseBody {
	s.PromptInfo = v
	return s
}

type UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules struct {
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Protocol     *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	AccType      *int32  `json:"AccType,omitempty" xml:"AccType,omitempty"`
	IpList       *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	DomainList   *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
}

func (s UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) GoString() string {
	return s.String()
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetMasterIpList(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.MasterIpList = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetRemark(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.Remark = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetSlaveIpList(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.SlaveIpList = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetIdentity(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.Identity = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetProtocol(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.Protocol = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetAccType(v int32) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.AccType = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetIpList(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.IpList = &v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules) SetDomainList(v string) *UploadFlexAccRulesFileForParseResponseBodyFlexAccFwdRules {
	s.DomainList = &v
	return s
}

type UploadFlexAccRulesFileForParseResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadFlexAccRulesFileForParseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadFlexAccRulesFileForParseResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexAccRulesFileForParseResponse) GoString() string {
	return s.String()
}

func (s *UploadFlexAccRulesFileForParseResponse) SetHeaders(v map[string]*string) *UploadFlexAccRulesFileForParseResponse {
	s.Headers = v
	return s
}

func (s *UploadFlexAccRulesFileForParseResponse) SetBody(v *UploadFlexAccRulesFileForParseResponseBody) *UploadFlexAccRulesFileForParseResponse {
	s.Body = v
	return s
}

type UploadFlexRulesFileForParseRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
}

func (s UploadFlexRulesFileForParseRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexRulesFileForParseRequest) GoString() string {
	return s.String()
}

func (s *UploadFlexRulesFileForParseRequest) SetSourceIp(v string) *UploadFlexRulesFileForParseRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadFlexRulesFileForParseRequest) SetFileKey(v string) *UploadFlexRulesFileForParseRequest {
	s.FileKey = &v
	return s
}

type UploadFlexRulesFileForParseResponseBody struct {
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	FlexFwdRules []*UploadFlexRulesFileForParseResponseBodyFlexFwdRules `json:"FlexFwdRules,omitempty" xml:"FlexFwdRules,omitempty" type:"Repeated"`
}

func (s UploadFlexRulesFileForParseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexRulesFileForParseResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFlexRulesFileForParseResponseBody) SetRequestId(v string) *UploadFlexRulesFileForParseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadFlexRulesFileForParseResponseBody) SetTotal(v int64) *UploadFlexRulesFileForParseResponseBody {
	s.Total = &v
	return s
}

func (s *UploadFlexRulesFileForParseResponseBody) SetFlexFwdRules(v []*UploadFlexRulesFileForParseResponseBodyFlexFwdRules) *UploadFlexRulesFileForParseResponseBody {
	s.FlexFwdRules = v
	return s
}

type UploadFlexRulesFileForParseResponseBodyFlexFwdRules struct {
	MasterIpList *string `json:"MasterIpList,omitempty" xml:"MasterIpList,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SlaveIpList  *string `json:"SlaveIpList,omitempty" xml:"SlaveIpList,omitempty"`
	Protocol     *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Identity     *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
}

func (s UploadFlexRulesFileForParseResponseBodyFlexFwdRules) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexRulesFileForParseResponseBodyFlexFwdRules) GoString() string {
	return s.String()
}

func (s *UploadFlexRulesFileForParseResponseBodyFlexFwdRules) SetMasterIpList(v string) *UploadFlexRulesFileForParseResponseBodyFlexFwdRules {
	s.MasterIpList = &v
	return s
}

func (s *UploadFlexRulesFileForParseResponseBodyFlexFwdRules) SetRemark(v string) *UploadFlexRulesFileForParseResponseBodyFlexFwdRules {
	s.Remark = &v
	return s
}

func (s *UploadFlexRulesFileForParseResponseBodyFlexFwdRules) SetSlaveIpList(v string) *UploadFlexRulesFileForParseResponseBodyFlexFwdRules {
	s.SlaveIpList = &v
	return s
}

func (s *UploadFlexRulesFileForParseResponseBodyFlexFwdRules) SetProtocol(v string) *UploadFlexRulesFileForParseResponseBodyFlexFwdRules {
	s.Protocol = &v
	return s
}

func (s *UploadFlexRulesFileForParseResponseBodyFlexFwdRules) SetIdentity(v string) *UploadFlexRulesFileForParseResponseBodyFlexFwdRules {
	s.Identity = &v
	return s
}

type UploadFlexRulesFileForParseResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadFlexRulesFileForParseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadFlexRulesFileForParseResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFlexRulesFileForParseResponse) GoString() string {
	return s.String()
}

func (s *UploadFlexRulesFileForParseResponse) SetHeaders(v map[string]*string) *UploadFlexRulesFileForParseResponse {
	s.Headers = v
	return s
}

func (s *UploadFlexRulesFileForParseResponse) SetBody(v *UploadFlexRulesFileForParseResponseBody) *UploadFlexRulesFileForParseResponse {
	s.Body = v
	return s
}

type UploadL4RulesFileForParseRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
}

func (s UploadL4RulesFileForParseRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadL4RulesFileForParseRequest) GoString() string {
	return s.String()
}

func (s *UploadL4RulesFileForParseRequest) SetSourceIp(v string) *UploadL4RulesFileForParseRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadL4RulesFileForParseRequest) SetLang(v string) *UploadL4RulesFileForParseRequest {
	s.Lang = &v
	return s
}

func (s *UploadL4RulesFileForParseRequest) SetFileKey(v string) *UploadL4RulesFileForParseRequest {
	s.FileKey = &v
	return s
}

type UploadL4RulesFileForParseResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	PromptInfo map[string]interface{}                             `json:"PromptInfo,omitempty" xml:"PromptInfo,omitempty"`
	L4RuleList []*UploadL4RulesFileForParseResponseBodyL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
}

func (s UploadL4RulesFileForParseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadL4RulesFileForParseResponseBody) GoString() string {
	return s.String()
}

func (s *UploadL4RulesFileForParseResponseBody) SetRequestId(v string) *UploadL4RulesFileForParseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBody) SetTotal(v int64) *UploadL4RulesFileForParseResponseBody {
	s.Total = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBody) SetPromptInfo(v map[string]interface{}) *UploadL4RulesFileForParseResponseBody {
	s.PromptInfo = v
	return s
}

func (s *UploadL4RulesFileForParseResponseBody) SetL4RuleList(v []*UploadL4RulesFileForParseResponseBodyL4RuleList) *UploadL4RulesFileForParseResponseBody {
	s.L4RuleList = v
	return s
}

type UploadL4RulesFileForParseResponseBodyL4RuleList struct {
	Protocol  *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	AppId     *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BackPort  *int32    `json:"BackPort,omitempty" xml:"BackPort,omitempty"`
	BizId     *int64    `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Rservers  []*string `json:"Rservers,omitempty" xml:"Rservers,omitempty" type:"Repeated"`
	LvsPolicy *string   `json:"LvsPolicy,omitempty" xml:"LvsPolicy,omitempty"`
	FrontPort *int32    `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	RuleId    *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UploadL4RulesFileForParseResponseBodyL4RuleList) String() string {
	return tea.Prettify(s)
}

func (s UploadL4RulesFileForParseResponseBodyL4RuleList) GoString() string {
	return s.String()
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetProtocol(v string) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.Protocol = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetAppId(v int64) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.AppId = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetBackPort(v int32) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.BackPort = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetBizId(v int64) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.BizId = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetRservers(v []*string) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.Rservers = v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetLvsPolicy(v string) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.LvsPolicy = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetFrontPort(v int32) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.FrontPort = &v
	return s
}

func (s *UploadL4RulesFileForParseResponseBodyL4RuleList) SetRuleId(v string) *UploadL4RulesFileForParseResponseBodyL4RuleList {
	s.RuleId = &v
	return s
}

type UploadL4RulesFileForParseResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadL4RulesFileForParseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadL4RulesFileForParseResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadL4RulesFileForParseResponse) GoString() string {
	return s.String()
}

func (s *UploadL4RulesFileForParseResponse) SetHeaders(v map[string]*string) *UploadL4RulesFileForParseResponse {
	s.Headers = v
	return s
}

func (s *UploadL4RulesFileForParseResponse) SetBody(v *UploadL4RulesFileForParseResponseBody) *UploadL4RulesFileForParseResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("gameshield"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CleanFlexAccFwdRulesWithOptions(request *CleanFlexAccFwdRulesRequest, runtime *util.RuntimeOptions) (_result *CleanFlexAccFwdRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CleanFlexAccFwdRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CleanFlexAccFwdRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CleanFlexAccFwdRules(request *CleanFlexAccFwdRulesRequest) (_result *CleanFlexAccFwdRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CleanFlexAccFwdRulesResponse{}
	_body, _err := client.CleanFlexAccFwdRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CleanFlexFwdRulesWithOptions(request *CleanFlexFwdRulesRequest, runtime *util.RuntimeOptions) (_result *CleanFlexFwdRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CleanFlexFwdRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CleanFlexFwdRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CleanFlexFwdRules(request *CleanFlexFwdRulesRequest) (_result *CleanFlexFwdRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CleanFlexFwdRulesResponse{}
	_body, _err := client.CleanFlexFwdRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearCcRouteRulesWithOptions(request *ClearCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *ClearCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClearCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClearCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearCcRouteRules(request *ClearCcRouteRulesRequest) (_result *ClearCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearCcRouteRulesResponse{}
	_body, _err := client.ClearCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApp"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppKeyWithOptions(request *CreateAppKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAppKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAppKey"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppKey(request *CreateAppKeyRequest) (_result *CreateAppKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.CreateAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBizWithOptions(request *CreateBizRequest, runtime *util.RuntimeOptions) (_result *CreateBizResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBizResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBiz"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBiz(request *CreateBizRequest) (_result *CreateBizResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBizResponse{}
	_body, _err := client.CreateBizWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCcRouteRulesWithOptions(request *CreateCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *CreateCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCcRouteRules(request *CreateCcRouteRulesRequest) (_result *CreateCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCcRouteRulesResponse{}
	_body, _err := client.CreateCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlexAccFwdRuleWithOptions(request *CreateFlexAccFwdRuleRequest, runtime *util.RuntimeOptions) (_result *CreateFlexAccFwdRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlexAccFwdRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlexAccFwdRule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlexAccFwdRule(request *CreateFlexAccFwdRuleRequest) (_result *CreateFlexAccFwdRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlexAccFwdRuleResponse{}
	_body, _err := client.CreateFlexAccFwdRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlexAccFwdRuleListWithOptions(request *CreateFlexAccFwdRuleListRequest, runtime *util.RuntimeOptions) (_result *CreateFlexAccFwdRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlexAccFwdRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlexAccFwdRuleList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlexAccFwdRuleList(request *CreateFlexAccFwdRuleListRequest) (_result *CreateFlexAccFwdRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlexAccFwdRuleListResponse{}
	_body, _err := client.CreateFlexAccFwdRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlexFwdRuleWithOptions(request *CreateFlexFwdRuleRequest, runtime *util.RuntimeOptions) (_result *CreateFlexFwdRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlexFwdRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlexFwdRule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlexFwdRule(request *CreateFlexFwdRuleRequest) (_result *CreateFlexFwdRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlexFwdRuleResponse{}
	_body, _err := client.CreateFlexFwdRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGroup"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayer4RuleWithOptions(request *CreateLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLayer4RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLayer4Rule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayer4Rule(request *CreateLayer4RuleRequest) (_result *CreateLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLayer4RuleResponse{}
	_body, _err := client.CreateLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayer4RulesWithOptions(request *CreateLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *CreateLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLayer4RulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLayer4Rules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayer4Rules(request *CreateLayer4RulesRequest) (_result *CreateLayer4RulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLayer4RulesResponse{}
	_body, _err := client.CreateLayer4RulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteApp"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppKeyWithOptions(request *DeleteAppKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteAppKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAppKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAppKey"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppKey(request *DeleteAppKeyRequest) (_result *DeleteAppKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppKeyResponse{}
	_body, _err := client.DeleteAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBizWithOptions(request *DeleteBizRequest, runtime *util.RuntimeOptions) (_result *DeleteBizResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBizResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBiz"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBiz(request *DeleteBizRequest) (_result *DeleteBizResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBizResponse{}
	_body, _err := client.DeleteBizWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCcRouteRulesWithOptions(request *DeleteCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCcRouteRules(request *DeleteCcRouteRulesRequest) (_result *DeleteCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCcRouteRulesResponse{}
	_body, _err := client.DeleteCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlexAccFwdRuleWithOptions(request *DeleteFlexAccFwdRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteFlexAccFwdRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlexAccFwdRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlexAccFwdRule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlexAccFwdRule(request *DeleteFlexAccFwdRuleRequest) (_result *DeleteFlexAccFwdRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlexAccFwdRuleResponse{}
	_body, _err := client.DeleteFlexAccFwdRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlexFwdRuleWithOptions(request *DeleteFlexFwdRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteFlexFwdRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlexFwdRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlexFwdRule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlexFwdRule(request *DeleteFlexFwdRuleRequest) (_result *DeleteFlexFwdRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlexFwdRuleResponse{}
	_body, _err := client.DeleteFlexFwdRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGroup"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer4RuleWithOptions(request *DeleteLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLayer4RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLayer4Rule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer4Rule(request *DeleteLayer4RuleRequest) (_result *DeleteLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer4RuleResponse{}
	_body, _err := client.DeleteLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccResSummaryWithOptions(request *DescribeAccResSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeAccResSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccResSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccResSummary"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccResSummary(request *DescribeAccResSummaryRequest) (_result *DescribeAccResSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccResSummaryResponse{}
	_body, _err := client.DescribeAccResSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllLocalIpsWithOptions(request *DescribeAllLocalIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeAllLocalIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllLocalIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllLocalIps"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllLocalIps(request *DescribeAllLocalIpsRequest) (_result *DescribeAllLocalIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllLocalIpsResponse{}
	_body, _err := client.DescribeAllLocalIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppDailyDetailsWithOptions(request *DescribeAppDailyDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppDailyDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppDailyDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAppDailyDetails"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppDailyDetails(request *DescribeAppDailyDetailsRequest) (_result *DescribeAppDailyDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppDailyDetailsResponse{}
	_body, _err := client.DescribeAppDailyDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppDailySdkVersionsWithOptions(request *DescribeAppDailySdkVersionsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppDailySdkVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppDailySdkVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAppDailySdkVersions"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppDailySdkVersions(request *DescribeAppDailySdkVersionsRequest) (_result *DescribeAppDailySdkVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppDailySdkVersionsResponse{}
	_body, _err := client.DescribeAppDailySdkVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppListWithOptions(request *DescribeAppListRequest, runtime *util.RuntimeOptions) (_result *DescribeAppListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAppList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppList(request *DescribeAppListRequest) (_result *DescribeAppListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppListResponse{}
	_body, _err := client.DescribeAppListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppSimpleListWithOptions(request *DescribeAppSimpleListRequest, runtime *util.RuntimeOptions) (_result *DescribeAppSimpleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppSimpleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAppSimpleList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppSimpleList(request *DescribeAppSimpleListRequest) (_result *DescribeAppSimpleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppSimpleListResponse{}
	_body, _err := client.DescribeAppSimpleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAsyncOperationWithOptions(request *DescribeAsyncOperationRequest, runtime *util.RuntimeOptions) (_result *DescribeAsyncOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAsyncOperationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAsyncOperation"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAsyncOperation(request *DescribeAsyncOperationRequest) (_result *DescribeAsyncOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAsyncOperationResponse{}
	_body, _err := client.DescribeAsyncOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutomaticTraceabilitySlsLogWithOptions(request *DescribeAutomaticTraceabilitySlsLogRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomaticTraceabilitySlsLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutomaticTraceabilitySlsLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutomaticTraceabilitySlsLog"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutomaticTraceabilitySlsLog(request *DescribeAutomaticTraceabilitySlsLogRequest) (_result *DescribeAutomaticTraceabilitySlsLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomaticTraceabilitySlsLogResponse{}
	_body, _err := client.DescribeAutomaticTraceabilitySlsLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBgpResSummaryWithOptions(request *DescribeBgpResSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeBgpResSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBgpResSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBgpResSummary"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBgpResSummary(request *DescribeBgpResSummaryRequest) (_result *DescribeBgpResSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBgpResSummaryResponse{}
	_body, _err := client.DescribeBgpResSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBizListWithOptions(request *DescribeBizListRequest, runtime *util.RuntimeOptions) (_result *DescribeBizListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBizListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBizList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizList(request *DescribeBizListRequest) (_result *DescribeBizListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizListResponse{}
	_body, _err := client.DescribeBizListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBizSimpleListWithOptions(request *DescribeBizSimpleListRequest, runtime *util.RuntimeOptions) (_result *DescribeBizSimpleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBizSimpleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBizSimpleList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizSimpleList(request *DescribeBizSimpleListRequest) (_result *DescribeBizSimpleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizSimpleListResponse{}
	_body, _err := client.DescribeBizSimpleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBpsFlowWithOptions(request *DescribeBpsFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeBpsFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBpsFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBpsFlow"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBpsFlow(request *DescribeBpsFlowRequest) (_result *DescribeBpsFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBpsFlowResponse{}
	_body, _err := client.DescribeBpsFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcBlackListWithOptions(request *DescribeCcBlackListRequest, runtime *util.RuntimeOptions) (_result *DescribeCcBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcBlackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcBlackList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcBlackList(request *DescribeCcBlackListRequest) (_result *DescribeCcBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcBlackListResponse{}
	_body, _err := client.DescribeCcBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcDevieIpListWithOptions(request *DescribeCcDevieIpListRequest, runtime *util.RuntimeOptions) (_result *DescribeCcDevieIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcDevieIpListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcDevieIpList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcDevieIpList(request *DescribeCcDevieIpListRequest) (_result *DescribeCcDevieIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcDevieIpListResponse{}
	_body, _err := client.DescribeCcDevieIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcFlowWithOptions(request *DescribeCcFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeCcFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcFlow"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcFlow(request *DescribeCcFlowRequest) (_result *DescribeCcFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcFlowResponse{}
	_body, _err := client.DescribeCcFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcIDCBlockSwitchWithOptions(request *DescribeCcIDCBlockSwitchRequest, runtime *util.RuntimeOptions) (_result *DescribeCcIDCBlockSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcIDCBlockSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcIDCBlockSwitch"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcIDCBlockSwitch(request *DescribeCcIDCBlockSwitchRequest) (_result *DescribeCcIDCBlockSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcIDCBlockSwitchResponse{}
	_body, _err := client.DescribeCcIDCBlockSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcMaxFwWithOptions(request *DescribeCcMaxFwRequest, runtime *util.RuntimeOptions) (_result *DescribeCcMaxFwResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcMaxFwResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcMaxFw"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcMaxFw(request *DescribeCcMaxFwRequest) (_result *DescribeCcMaxFwResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcMaxFwResponse{}
	_body, _err := client.DescribeCcMaxFwWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcResSummaryWithOptions(request *DescribeCcResSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeCcResSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcResSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcResSummary"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcResSummary(request *DescribeCcResSummaryRequest) (_result *DescribeCcResSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcResSummaryResponse{}
	_body, _err := client.DescribeCcResSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcRouteRulesWithOptions(request *DescribeCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcRouteRules(request *DescribeCcRouteRulesRequest) (_result *DescribeCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcRouteRulesResponse{}
	_body, _err := client.DescribeCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcRouteSwitchWithOptions(request *DescribeCcRouteSwitchRequest, runtime *util.RuntimeOptions) (_result *DescribeCcRouteSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcRouteSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcRouteSwitch"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcRouteSwitch(request *DescribeCcRouteSwitchRequest) (_result *DescribeCcRouteSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcRouteSwitchResponse{}
	_body, _err := client.DescribeCcRouteSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcSocketDetailWithOptions(request *DescribeCcSocketDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCcSocketDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcSocketDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcSocketDetail"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcSocketDetail(request *DescribeCcSocketDetailRequest) (_result *DescribeCcSocketDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcSocketDetailResponse{}
	_body, _err := client.DescribeCcSocketDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcTotalFwWithOptions(request *DescribeCcTotalFwRequest, runtime *util.RuntimeOptions) (_result *DescribeCcTotalFwResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcTotalFwResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcTotalFw"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcTotalFw(request *DescribeCcTotalFwRequest) (_result *DescribeCcTotalFwResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcTotalFwResponse{}
	_body, _err := client.DescribeCcTotalFwWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcTunnelWithOptions(request *DescribeCcTunnelRequest, runtime *util.RuntimeOptions) (_result *DescribeCcTunnelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcTunnelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcTunnel"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcTunnel(request *DescribeCcTunnelRequest) (_result *DescribeCcTunnelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcTunnelResponse{}
	_body, _err := client.DescribeCcTunnelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcWhiteListWithOptions(request *DescribeCcWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeCcWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcWhiteList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcWhiteList(request *DescribeCcWhiteListRequest) (_result *DescribeCcWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcWhiteListResponse{}
	_body, _err := client.DescribeCcWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcZoneBlockConfigWithOptions(request *DescribeCcZoneBlockConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeCcZoneBlockConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcZoneBlockConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcZoneBlockConfig"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcZoneBlockConfig(request *DescribeCcZoneBlockConfigRequest) (_result *DescribeCcZoneBlockConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcZoneBlockConfigResponse{}
	_body, _err := client.DescribeCcZoneBlockConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcZonesWithOptions(request *DescribeCcZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeCcZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcZones"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcZones(request *DescribeCcZonesRequest) (_result *DescribeCcZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcZonesResponse{}
	_body, _err := client.DescribeCcZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDailyDetailsWithOptions(request *DescribeDailyDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeDailyDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDailyDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDailyDetails"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDailyDetails(request *DescribeDailyDetailsRequest) (_result *DescribeDailyDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDailyDetailsResponse{}
	_body, _err := client.DescribeDailyDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadUrlsForAppKeyWithOptions(request *DescribeDownloadUrlsForAppKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadUrlsForAppKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDownloadUrlsForAppKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDownloadUrlsForAppKey"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadUrlsForAppKey(request *DescribeDownloadUrlsForAppKeyRequest) (_result *DescribeDownloadUrlsForAppKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadUrlsForAppKeyResponse{}
	_body, _err := client.DescribeDownloadUrlsForAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadUrlsForSDKWithOptions(request *DescribeDownloadUrlsForSDKRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadUrlsForSDKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDownloadUrlsForSDKResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDownloadUrlsForSDK"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadUrlsForSDK(request *DescribeDownloadUrlsForSDKRequest) (_result *DescribeDownloadUrlsForSDKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadUrlsForSDKResponse{}
	_body, _err := client.DescribeDownloadUrlsForSDKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexAccConfigWithOptions(request *DescribeFlexAccConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexAccConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexAccConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexAccConfig"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexAccConfig(request *DescribeFlexAccConfigRequest) (_result *DescribeFlexAccConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexAccConfigResponse{}
	_body, _err := client.DescribeFlexAccConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexAccFlowWithOptions(request *DescribeFlexAccFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexAccFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexAccFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexAccFlow"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexAccFlow(request *DescribeFlexAccFlowRequest) (_result *DescribeFlexAccFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexAccFlowResponse{}
	_body, _err := client.DescribeFlexAccFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexAccFwdRulesWithOptions(request *DescribeFlexAccFwdRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexAccFwdRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexAccFwdRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexAccFwdRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexAccFwdRules(request *DescribeFlexAccFwdRulesRequest) (_result *DescribeFlexAccFwdRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexAccFwdRulesResponse{}
	_body, _err := client.DescribeFlexAccFwdRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexBackupGroupsWithOptions(request *DescribeFlexBackupGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexBackupGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexBackupGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexBackupGroups"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexBackupGroups(request *DescribeFlexBackupGroupsRequest) (_result *DescribeFlexBackupGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexBackupGroupsResponse{}
	_body, _err := client.DescribeFlexBackupGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexConfigWithOptions(request *DescribeFlexConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexConfig"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexConfig(request *DescribeFlexConfigRequest) (_result *DescribeFlexConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexConfigResponse{}
	_body, _err := client.DescribeFlexConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexFwdRulesWithOptions(request *DescribeFlexFwdRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexFwdRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexFwdRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexFwdRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexFwdRules(request *DescribeFlexFwdRulesRequest) (_result *DescribeFlexFwdRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexFwdRulesResponse{}
	_body, _err := client.DescribeFlexFwdRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlexSechedPolicyWithOptions(request *DescribeFlexSechedPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeFlexSechedPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlexSechedPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlexSechedPolicy"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlexSechedPolicy(request *DescribeFlexSechedPolicyRequest) (_result *DescribeFlexSechedPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlexSechedPolicyResponse{}
	_body, _err := client.DescribeFlexSechedPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowReportWithOptions(request *DescribeFlowReportRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowReport"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowReport(request *DescribeFlowReportRequest) (_result *DescribeFlowReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowReportResponse{}
	_body, _err := client.DescribeFlowReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFullNodesWithOptions(request *DescribeFullNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeFullNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFullNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFullNodes"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFullNodes(request *DescribeFullNodesRequest) (_result *DescribeFullNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFullNodesResponse{}
	_body, _err := client.DescribeFullNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFullNodesSummaryWithOptions(request *DescribeFullNodesSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeFullNodesSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFullNodesSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFullNodesSummary"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFullNodesSummary(request *DescribeFullNodesSummaryRequest) (_result *DescribeFullNodesSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFullNodesSummaryResponse{}
	_body, _err := client.DescribeFullNodesSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGfResSummaryWithOptions(request *DescribeGfResSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeGfResSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGfResSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGfResSummary"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGfResSummary(request *DescribeGfResSummaryRequest) (_result *DescribeGfResSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGfResSummaryResponse{}
	_body, _err := client.DescribeGfResSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupListWithOptions(request *DescribeGroupListRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupList(request *DescribeGroupListRequest) (_result *DescribeGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupListResponse{}
	_body, _err := client.DescribeGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupSimpleListWithOptions(request *DescribeGroupSimpleListRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupSimpleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupSimpleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupSimpleList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupSimpleList(request *DescribeGroupSimpleListRequest) (_result *DescribeGroupSimpleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupSimpleListResponse{}
	_body, _err := client.DescribeGroupSimpleListWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstance"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeIpFlowReportWithOptions(request *DescribeIpFlowReportRequest, runtime *util.RuntimeOptions) (_result *DescribeIpFlowReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpFlowReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpFlowReport"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpFlowReport(request *DescribeIpFlowReportRequest) (_result *DescribeIpFlowReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpFlowReportResponse{}
	_body, _err := client.DescribeIpFlowReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJianYuTestGetWithOptions(request *DescribeJianYuTestGetRequest, runtime *util.RuntimeOptions) (_result *DescribeJianYuTestGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeJianYuTestGetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJianYuTestGet"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJianYuTestGet(request *DescribeJianYuTestGetRequest) (_result *DescribeJianYuTestGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJianYuTestGetResponse{}
	_body, _err := client.DescribeJianYuTestGetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJianYuTestListWithOptions(request *DescribeJianYuTestListRequest, runtime *util.RuntimeOptions) (_result *DescribeJianYuTestListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeJianYuTestListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJianYuTestList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJianYuTestList(request *DescribeJianYuTestListRequest) (_result *DescribeJianYuTestListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJianYuTestListResponse{}
	_body, _err := client.DescribeJianYuTestListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJianYuTestPaginationWithOptions(request *DescribeJianYuTestPaginationRequest, runtime *util.RuntimeOptions) (_result *DescribeJianYuTestPaginationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeJianYuTestPaginationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJianYuTestPagination"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJianYuTestPagination(request *DescribeJianYuTestPaginationRequest) (_result *DescribeJianYuTestPaginationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJianYuTestPaginationResponse{}
	_body, _err := client.DescribeJianYuTestPaginationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeL4EventsSelectiveWithOptions(request *DescribeL4EventsSelectiveRequest, runtime *util.RuntimeOptions) (_result *DescribeL4EventsSelectiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeL4EventsSelectiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeL4EventsSelective"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeL4EventsSelective(request *DescribeL4EventsSelectiveRequest) (_result *DescribeL4EventsSelectiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeL4EventsSelectiveResponse{}
	_body, _err := client.DescribeL4EventsSelectiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer4RulesWithOptions(request *DescribeLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLayer4RulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLayer4Rules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer4Rules(request *DescribeLayer4RulesRequest) (_result *DescribeLayer4RulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer4RulesResponse{}
	_body, _err := client.DescribeLayer4RulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMianLiuResSummaryWithOptions(request *DescribeMianLiuResSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeMianLiuResSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMianLiuResSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMianLiuResSummary"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMianLiuResSummary(request *DescribeMianLiuResSummaryRequest) (_result *DescribeMianLiuResSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMianLiuResSummaryResponse{}
	_body, _err := client.DescribeMianLiuResSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNgSourceDiagnosisLogWithOptions(request *DescribeNgSourceDiagnosisLogRequest, runtime *util.RuntimeOptions) (_result *DescribeNgSourceDiagnosisLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNgSourceDiagnosisLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNgSourceDiagnosisLog"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNgSourceDiagnosisLog(request *DescribeNgSourceDiagnosisLogRequest) (_result *DescribeNgSourceDiagnosisLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNgSourceDiagnosisLogResponse{}
	_body, _err := client.DescribeNgSourceDiagnosisLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRemainQpsWithOptions(request *DescribeRemainQpsRequest, runtime *util.RuntimeOptions) (_result *DescribeRemainQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRemainQpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRemainQps"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRemainQps(request *DescribeRemainQpsRequest) (_result *DescribeRemainQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRemainQpsResponse{}
	_body, _err := client.DescribeRemainQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticByEsnBizIntervalWithOptions(request *DescribeRequestStatisticByEsnBizIntervalRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticByEsnBizIntervalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticByEsnBizIntervalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticByEsnBizInterval"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticByEsnBizInterval(request *DescribeRequestStatisticByEsnBizIntervalRequest) (_result *DescribeRequestStatisticByEsnBizIntervalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticByEsnBizIntervalResponse{}
	_body, _err := client.DescribeRequestStatisticByEsnBizIntervalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz1DayWithOptions(request *DescribeRequestStatisticCountByEsnBiz1DayRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticCountByEsnBiz1DayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticCountByEsnBiz1DayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticCountByEsnBiz1Day"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz1Day(request *DescribeRequestStatisticCountByEsnBiz1DayRequest) (_result *DescribeRequestStatisticCountByEsnBiz1DayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticCountByEsnBiz1DayResponse{}
	_body, _err := client.DescribeRequestStatisticCountByEsnBiz1DayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz1DayProvinceWithOptions(request *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticCountByEsnBiz1DayProvince"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz1DayProvince(request *DescribeRequestStatisticCountByEsnBiz1DayProvinceRequest) (_result *DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticCountByEsnBiz1DayProvinceResponse{}
	_body, _err := client.DescribeRequestStatisticCountByEsnBiz1DayProvinceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz1M30MWithOptions(request *DescribeRequestStatisticCountByEsnBiz1M30MRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticCountByEsnBiz1M30MResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticCountByEsnBiz1M30MResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticCountByEsnBiz1M30M"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz1M30M(request *DescribeRequestStatisticCountByEsnBiz1M30MRequest) (_result *DescribeRequestStatisticCountByEsnBiz1M30MResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticCountByEsnBiz1M30MResponse{}
	_body, _err := client.DescribeRequestStatisticCountByEsnBiz1M30MWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz30MWithOptions(request *DescribeRequestStatisticCountByEsnBiz30MRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticCountByEsnBiz30MResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticCountByEsnBiz30MResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticCountByEsnBiz30M"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBiz30M(request *DescribeRequestStatisticCountByEsnBiz30MRequest) (_result *DescribeRequestStatisticCountByEsnBiz30MResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticCountByEsnBiz30MResponse{}
	_body, _err := client.DescribeRequestStatisticCountByEsnBiz30MWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBizGroup30MWithOptions(request *DescribeRequestStatisticCountByEsnBizGroup30MRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticCountByEsnBizGroup30MResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticCountByEsnBizGroup30MResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticCountByEsnBizGroup30M"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticCountByEsnBizGroup30M(request *DescribeRequestStatisticCountByEsnBizGroup30MRequest) (_result *DescribeRequestStatisticCountByEsnBizGroup30MResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticCountByEsnBizGroup30MResponse{}
	_body, _err := client.DescribeRequestStatisticCountByEsnBizGroup30MWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRequestStatisticLogWithOptions(request *DescribeRequestStatisticLogRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestStatisticLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRequestStatisticLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRequestStatisticLog"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestStatisticLog(request *DescribeRequestStatisticLogRequest) (_result *DescribeRequestStatisticLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestStatisticLogResponse{}
	_body, _err := client.DescribeRequestStatisticLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSDKStatisticLogWithOptions(request *DescribeSDKStatisticLogRequest, runtime *util.RuntimeOptions) (_result *DescribeSDKStatisticLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSDKStatisticLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSDKStatisticLog"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSDKStatisticLog(request *DescribeSDKStatisticLogRequest) (_result *DescribeSDKStatisticLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSDKStatisticLogResponse{}
	_body, _err := client.DescribeSDKStatisticLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizISP1M30MWithOptions(request *DescribeSDKStatisticResultByEsnBizISP1M30MRequest, runtime *util.RuntimeOptions) (_result *DescribeSDKStatisticResultByEsnBizISP1M30MResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSDKStatisticResultByEsnBizISP1M30MResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSDKStatisticResultByEsnBizISP1M30M"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizISP1M30M(request *DescribeSDKStatisticResultByEsnBizISP1M30MRequest) (_result *DescribeSDKStatisticResultByEsnBizISP1M30MResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSDKStatisticResultByEsnBizISP1M30MResponse{}
	_body, _err := client.DescribeSDKStatisticResultByEsnBizISP1M30MWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizISPIntervalWithOptions(request *DescribeSDKStatisticResultByEsnBizISPIntervalRequest, runtime *util.RuntimeOptions) (_result *DescribeSDKStatisticResultByEsnBizISPIntervalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSDKStatisticResultByEsnBizISPIntervalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSDKStatisticResultByEsnBizISPInterval"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizISPInterval(request *DescribeSDKStatisticResultByEsnBizISPIntervalRequest) (_result *DescribeSDKStatisticResultByEsnBizISPIntervalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSDKStatisticResultByEsnBizISPIntervalResponse{}
	_body, _err := client.DescribeSDKStatisticResultByEsnBizISPIntervalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizProvince1DayWithOptions(request *DescribeSDKStatisticResultByEsnBizProvince1DayRequest, runtime *util.RuntimeOptions) (_result *DescribeSDKStatisticResultByEsnBizProvince1DayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSDKStatisticResultByEsnBizProvince1DayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSDKStatisticResultByEsnBizProvince1Day"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizProvince1Day(request *DescribeSDKStatisticResultByEsnBizProvince1DayRequest) (_result *DescribeSDKStatisticResultByEsnBizProvince1DayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSDKStatisticResultByEsnBizProvince1DayResponse{}
	_body, _err := client.DescribeSDKStatisticResultByEsnBizProvince1DayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizProvince30MWithOptions(request *DescribeSDKStatisticResultByEsnBizProvince30MRequest, runtime *util.RuntimeOptions) (_result *DescribeSDKStatisticResultByEsnBizProvince30MResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSDKStatisticResultByEsnBizProvince30MResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSDKStatisticResultByEsnBizProvince30M"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSDKStatisticResultByEsnBizProvince30M(request *DescribeSDKStatisticResultByEsnBizProvince30MRequest) (_result *DescribeSDKStatisticResultByEsnBizProvince30MResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSDKStatisticResultByEsnBizProvince30MResponse{}
	_body, _err := client.DescribeSDKStatisticResultByEsnBizProvince30MWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSourceFailureTopIpWithOptions(request *DescribeSourceFailureTopIpRequest, runtime *util.RuntimeOptions) (_result *DescribeSourceFailureTopIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSourceFailureTopIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSourceFailureTopIp"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSourceFailureTopIp(request *DescribeSourceFailureTopIpRequest) (_result *DescribeSourceFailureTopIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSourceFailureTopIpResponse{}
	_body, _err := client.DescribeSourceFailureTopIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSourceFailureTrendGraphWithOptions(request *DescribeSourceFailureTrendGraphRequest, runtime *util.RuntimeOptions) (_result *DescribeSourceFailureTrendGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSourceFailureTrendGraphResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSourceFailureTrendGraph"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSourceFailureTrendGraph(request *DescribeSourceFailureTrendGraphRequest) (_result *DescribeSourceFailureTrendGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSourceFailureTrendGraphResponse{}
	_body, _err := client.DescribeSourceFailureTrendGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUploadPreSignWithOptions(request *DescribeUploadPreSignRequest, runtime *util.RuntimeOptions) (_result *DescribeUploadPreSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUploadPreSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUploadPreSign"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUploadPreSign(request *DescribeUploadPreSignRequest) (_result *DescribeUploadPreSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUploadPreSignResponse{}
	_body, _err := client.DescribeUploadPreSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserFlowInfoWithOptions(request *DescribeUserFlowInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUserFlowInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserFlowInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserFlowInfo"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserFlowInfo(request *DescribeUserFlowInfoRequest) (_result *DescribeUserFlowInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserFlowInfoResponse{}
	_body, _err := client.DescribeUserFlowInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserFlowReportWithOptions(request *DescribeUserFlowReportRequest, runtime *util.RuntimeOptions) (_result *DescribeUserFlowReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserFlowReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserFlowReport"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserFlowReport(request *DescribeUserFlowReportRequest) (_result *DescribeUserFlowReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserFlowReportResponse{}
	_body, _err := client.DescribeUserFlowReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserTotalFlowReportWithOptions(request *DescribeUserTotalFlowReportRequest, runtime *util.RuntimeOptions) (_result *DescribeUserTotalFlowReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserTotalFlowReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserTotalFlowReport"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserTotalFlowReport(request *DescribeUserTotalFlowReportRequest) (_result *DescribeUserTotalFlowReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserTotalFlowReportResponse{}
	_body, _err := client.DescribeUserTotalFlowReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadAppKeyFileWithOptions(request *DownloadAppKeyFileRequest, runtime *util.RuntimeOptions) (_result *DownloadAppKeyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadAppKeyFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadAppKeyFile"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadAppKeyFile(request *DownloadAppKeyFileRequest) (_result *DownloadAppKeyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadAppKeyFileResponse{}
	_body, _err := client.DownloadAppKeyFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadCcRouteRulesWithOptions(request *DownloadCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *DownloadCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadCcRouteRules(request *DownloadCcRouteRulesRequest) (_result *DownloadCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadCcRouteRulesResponse{}
	_body, _err := client.DownloadCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadFlexAccRulesFileWithOptions(request *DownloadFlexAccRulesFileRequest, runtime *util.RuntimeOptions) (_result *DownloadFlexAccRulesFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadFlexAccRulesFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadFlexAccRulesFile"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadFlexAccRulesFile(request *DownloadFlexAccRulesFileRequest) (_result *DownloadFlexAccRulesFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadFlexAccRulesFileResponse{}
	_body, _err := client.DownloadFlexAccRulesFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadLayer4RulesWithOptions(request *DownloadLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *DownloadLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadLayer4RulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadLayer4Rules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadLayer4Rules(request *DownloadLayer4RulesRequest) (_result *DownloadLayer4RulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadLayer4RulesResponse{}
	_body, _err := client.DownloadLayer4RulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadSdkFileWithOptions(request *DownloadSdkFileRequest, runtime *util.RuntimeOptions) (_result *DownloadSdkFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadSdkFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadSdkFile"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadSdkFile(request *DownloadSdkFileRequest) (_result *DownloadSdkFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadSdkFileResponse{}
	_body, _err := client.DownloadSdkFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadUserTotalFlowReportWithOptions(request *DownloadUserTotalFlowReportRequest, runtime *util.RuntimeOptions) (_result *DownloadUserTotalFlowReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadUserTotalFlowReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadUserTotalFlowReport"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadUserTotalFlowReport(request *DownloadUserTotalFlowReportRequest) (_result *DownloadUserTotalFlowReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadUserTotalFlowReportResponse{}
	_body, _err := client.DownloadUserTotalFlowReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FlushLayer4RulesWithOptions(request *FlushLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *FlushLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FlushLayer4RulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FlushLayer4Rules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FlushLayer4Rules(request *FlushLayer4RulesRequest) (_result *FlushLayer4RulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FlushLayer4RulesResponse{}
	_body, _err := client.FlushLayer4RulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReallocNgResourceWithOptions(request *ReallocNgResourceRequest, runtime *util.RuntimeOptions) (_result *ReallocNgResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReallocNgResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReallocNgResource"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReallocNgResource(request *ReallocNgResourceRequest) (_result *ReallocNgResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReallocNgResourceResponse{}
	_body, _err := client.ReallocNgResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceCcRouteRulesWithOptions(request *ReplaceCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *ReplaceCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReplaceCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReplaceCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceCcRouteRules(request *ReplaceCcRouteRulesRequest) (_result *ReplaceCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceCcRouteRulesResponse{}
	_body, _err := client.ReplaceCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFlexFwdRulesWithOptions(request *SearchFlexFwdRulesRequest, runtime *util.RuntimeOptions) (_result *SearchFlexFwdRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchFlexFwdRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchFlexFwdRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFlexFwdRules(request *SearchFlexFwdRulesRequest) (_result *SearchFlexFwdRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchFlexFwdRulesResponse{}
	_body, _err := client.SearchFlexFwdRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApp"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcBlackListWithOptions(request *UpdateCcBlackListRequest, runtime *util.RuntimeOptions) (_result *UpdateCcBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcBlackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcBlackList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcBlackList(request *UpdateCcBlackListRequest) (_result *UpdateCcBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcBlackListResponse{}
	_body, _err := client.UpdateCcBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcIDCBlockSwitchWithOptions(request *UpdateCcIDCBlockSwitchRequest, runtime *util.RuntimeOptions) (_result *UpdateCcIDCBlockSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcIDCBlockSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcIDCBlockSwitch"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcIDCBlockSwitch(request *UpdateCcIDCBlockSwitchRequest) (_result *UpdateCcIDCBlockSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcIDCBlockSwitchResponse{}
	_body, _err := client.UpdateCcIDCBlockSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcRouteRulesWithOptions(request *UpdateCcRouteRulesRequest, runtime *util.RuntimeOptions) (_result *UpdateCcRouteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcRouteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcRouteRules"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcRouteRules(request *UpdateCcRouteRulesRequest) (_result *UpdateCcRouteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcRouteRulesResponse{}
	_body, _err := client.UpdateCcRouteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcRouteSwitchWithOptions(request *UpdateCcRouteSwitchRequest, runtime *util.RuntimeOptions) (_result *UpdateCcRouteSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcRouteSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcRouteSwitch"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcRouteSwitch(request *UpdateCcRouteSwitchRequest) (_result *UpdateCcRouteSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcRouteSwitchResponse{}
	_body, _err := client.UpdateCcRouteSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcTunnelGrayAndOnlyAllowWithOptions(request *UpdateCcTunnelGrayAndOnlyAllowRequest, runtime *util.RuntimeOptions) (_result *UpdateCcTunnelGrayAndOnlyAllowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcTunnelGrayAndOnlyAllowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcTunnelGrayAndOnlyAllow"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcTunnelGrayAndOnlyAllow(request *UpdateCcTunnelGrayAndOnlyAllowRequest) (_result *UpdateCcTunnelGrayAndOnlyAllowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcTunnelGrayAndOnlyAllowResponse{}
	_body, _err := client.UpdateCcTunnelGrayAndOnlyAllowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcTunnelStatusWithOptions(request *UpdateCcTunnelStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateCcTunnelStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcTunnelStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcTunnelStatus"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcTunnelStatus(request *UpdateCcTunnelStatusRequest) (_result *UpdateCcTunnelStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcTunnelStatusResponse{}
	_body, _err := client.UpdateCcTunnelStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcUseWithOptions(request *UpdateCcUseRequest, runtime *util.RuntimeOptions) (_result *UpdateCcUseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcUseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcUse"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcUse(request *UpdateCcUseRequest) (_result *UpdateCcUseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcUseResponse{}
	_body, _err := client.UpdateCcUseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcWhiteListWithOptions(request *UpdateCcWhiteListRequest, runtime *util.RuntimeOptions) (_result *UpdateCcWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcWhiteList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcWhiteList(request *UpdateCcWhiteListRequest) (_result *UpdateCcWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcWhiteListResponse{}
	_body, _err := client.UpdateCcWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcZoneBlockConfigWithOptions(request *UpdateCcZoneBlockConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateCcZoneBlockConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcZoneBlockConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcZoneBlockConfig"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcZoneBlockConfig(request *UpdateCcZoneBlockConfigRequest) (_result *UpdateCcZoneBlockConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcZoneBlockConfigResponse{}
	_body, _err := client.UpdateCcZoneBlockConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcZoneBlockStatusWithOptions(request *UpdateCcZoneBlockStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateCcZoneBlockStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcZoneBlockStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcZoneBlockStatus"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcZoneBlockStatus(request *UpdateCcZoneBlockStatusRequest) (_result *UpdateCcZoneBlockStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcZoneBlockStatusResponse{}
	_body, _err := client.UpdateCcZoneBlockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexAccFwdRuleWithOptions(request *UpdateFlexAccFwdRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexAccFwdRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexAccFwdRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexAccFwdRule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexAccFwdRule(request *UpdateFlexAccFwdRuleRequest) (_result *UpdateFlexAccFwdRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexAccFwdRuleResponse{}
	_body, _err := client.UpdateFlexAccFwdRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexAccFwdRuleV2WithOptions(request *UpdateFlexAccFwdRuleV2Request, runtime *util.RuntimeOptions) (_result *UpdateFlexAccFwdRuleV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexAccFwdRuleV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexAccFwdRuleV2"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexAccFwdRuleV2(request *UpdateFlexAccFwdRuleV2Request) (_result *UpdateFlexAccFwdRuleV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexAccFwdRuleV2Response{}
	_body, _err := client.UpdateFlexAccFwdRuleV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexAccStatusWithOptions(request *UpdateFlexAccStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexAccStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexAccStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexAccStatus"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexAccStatus(request *UpdateFlexAccStatusRequest) (_result *UpdateFlexAccStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexAccStatusResponse{}
	_body, _err := client.UpdateFlexAccStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexAccTcpPortsWithOptions(request *UpdateFlexAccTcpPortsRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexAccTcpPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexAccTcpPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexAccTcpPorts"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexAccTcpPorts(request *UpdateFlexAccTcpPortsRequest) (_result *UpdateFlexAccTcpPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexAccTcpPortsResponse{}
	_body, _err := client.UpdateFlexAccTcpPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexAccUdpPortsWithOptions(request *UpdateFlexAccUdpPortsRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexAccUdpPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexAccUdpPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexAccUdpPorts"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexAccUdpPorts(request *UpdateFlexAccUdpPortsRequest) (_result *UpdateFlexAccUdpPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexAccUdpPortsResponse{}
	_body, _err := client.UpdateFlexAccUdpPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexBackupGroupsWithOptions(request *UpdateFlexBackupGroupsRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexBackupGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexBackupGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexBackupGroups"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexBackupGroups(request *UpdateFlexBackupGroupsRequest) (_result *UpdateFlexBackupGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexBackupGroupsResponse{}
	_body, _err := client.UpdateFlexBackupGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexFwdRuleWithOptions(request *UpdateFlexFwdRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexFwdRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexFwdRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexFwdRule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexFwdRule(request *UpdateFlexFwdRuleRequest) (_result *UpdateFlexFwdRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexFwdRuleResponse{}
	_body, _err := client.UpdateFlexFwdRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexInnerPolicyWithOptions(request *UpdateFlexInnerPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexInnerPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexInnerPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexInnerPolicy"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexInnerPolicy(request *UpdateFlexInnerPolicyRequest) (_result *UpdateFlexInnerPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexInnerPolicyResponse{}
	_body, _err := client.UpdateFlexInnerPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexLinkTypeWithOptions(request *UpdateFlexLinkTypeRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexLinkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexLinkTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexLinkType"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexLinkType(request *UpdateFlexLinkTypeRequest) (_result *UpdateFlexLinkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexLinkTypeResponse{}
	_body, _err := client.UpdateFlexLinkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexPortsWithOptions(request *UpdateFlexPortsRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexPorts"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexPorts(request *UpdateFlexPortsRequest) (_result *UpdateFlexPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexPortsResponse{}
	_body, _err := client.UpdateFlexPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlexStatusWithOptions(request *UpdateFlexStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateFlexStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlexStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlexStatus"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlexStatus(request *UpdateFlexStatusRequest) (_result *UpdateFlexStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlexStatusResponse{}
	_body, _err := client.UpdateFlexStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroup"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupDnsStatusWithOptions(request *UpdateGroupDnsStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupDnsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupDnsStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroupDnsStatus"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupDnsStatus(request *UpdateGroupDnsStatusRequest) (_result *UpdateGroupDnsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupDnsStatusResponse{}
	_body, _err := client.UpdateGroupDnsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupNodesWithOptions(request *UpdateGroupNodesRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroupNodes"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupNodes(request *UpdateGroupNodesRequest) (_result *UpdateGroupNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupNodesResponse{}
	_body, _err := client.UpdateGroupNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupStatusWithOptions(request *UpdateGroupStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateGroupStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateGroupStatus"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupStatus(request *UpdateGroupStatusRequest) (_result *UpdateGroupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupStatusResponse{}
	_body, _err := client.UpdateGroupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLayer4RuleWithOptions(request *UpdateLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *UpdateLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLayer4RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLayer4Rule"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLayer4Rule(request *UpdateLayer4RuleRequest) (_result *UpdateLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLayer4RuleResponse{}
	_body, _err := client.UpdateLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadCcRouteFileForParseWithOptions(request *UploadCcRouteFileForParseRequest, runtime *util.RuntimeOptions) (_result *UploadCcRouteFileForParseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadCcRouteFileForParseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadCcRouteFileForParse"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadCcRouteFileForParse(request *UploadCcRouteFileForParseRequest) (_result *UploadCcRouteFileForParseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadCcRouteFileForParseResponse{}
	_body, _err := client.UploadCcRouteFileForParseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadCcWhiteBlackListWithOptions(request *UploadCcWhiteBlackListRequest, runtime *util.RuntimeOptions) (_result *UploadCcWhiteBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadCcWhiteBlackListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadCcWhiteBlackList"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadCcWhiteBlackList(request *UploadCcWhiteBlackListRequest) (_result *UploadCcWhiteBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadCcWhiteBlackListResponse{}
	_body, _err := client.UploadCcWhiteBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFlexAccRulesFileForParseWithOptions(request *UploadFlexAccRulesFileForParseRequest, runtime *util.RuntimeOptions) (_result *UploadFlexAccRulesFileForParseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadFlexAccRulesFileForParseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadFlexAccRulesFileForParse"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadFlexAccRulesFileForParse(request *UploadFlexAccRulesFileForParseRequest) (_result *UploadFlexAccRulesFileForParseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadFlexAccRulesFileForParseResponse{}
	_body, _err := client.UploadFlexAccRulesFileForParseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFlexRulesFileForParseWithOptions(request *UploadFlexRulesFileForParseRequest, runtime *util.RuntimeOptions) (_result *UploadFlexRulesFileForParseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadFlexRulesFileForParseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadFlexRulesFileForParse"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadFlexRulesFileForParse(request *UploadFlexRulesFileForParseRequest) (_result *UploadFlexRulesFileForParseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadFlexRulesFileForParseResponse{}
	_body, _err := client.UploadFlexRulesFileForParseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadL4RulesFileForParseWithOptions(request *UploadL4RulesFileForParseRequest, runtime *util.RuntimeOptions) (_result *UploadL4RulesFileForParseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadL4RulesFileForParseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadL4RulesFileForParse"), tea.String("2018-03-05"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadL4RulesFileForParse(request *UploadL4RulesFileForParseRequest) (_result *UploadL4RulesFileForParseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadL4RulesFileForParseResponse{}
	_body, _err := client.UploadL4RulesFileForParseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
