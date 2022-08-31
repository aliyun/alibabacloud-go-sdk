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

type AddScdnDomainRequest struct {
	CheckUrl        *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s AddScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *AddScdnDomainRequest) SetCheckUrl(v string) *AddScdnDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddScdnDomainRequest) SetDomainName(v string) *AddScdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddScdnDomainRequest) SetOwnerAccount(v string) *AddScdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddScdnDomainRequest) SetOwnerId(v int64) *AddScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddScdnDomainRequest) SetResourceGroupId(v string) *AddScdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddScdnDomainRequest) SetScope(v string) *AddScdnDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddScdnDomainRequest) SetSecurityToken(v string) *AddScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddScdnDomainRequest) SetSources(v string) *AddScdnDomainRequest {
	s.Sources = &v
	return s
}

type AddScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddScdnDomainResponseBody) SetRequestId(v string) *AddScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddScdnDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *AddScdnDomainResponse) SetHeaders(v map[string]*string) *AddScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *AddScdnDomainResponse) SetStatusCode(v int32) *AddScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddScdnDomainResponse) SetBody(v *AddScdnDomainResponseBody) *AddScdnDomainResponse {
	s.Body = v
	return s
}

type BatchDeleteScdnDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchDeleteScdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteScdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteScdnDomainConfigsRequest) SetDomainNames(v string) *BatchDeleteScdnDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchDeleteScdnDomainConfigsRequest) SetFunctionNames(v string) *BatchDeleteScdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *BatchDeleteScdnDomainConfigsRequest) SetOwnerAccount(v string) *BatchDeleteScdnDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchDeleteScdnDomainConfigsRequest) SetOwnerId(v int64) *BatchDeleteScdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchDeleteScdnDomainConfigsRequest) SetSecurityToken(v string) *BatchDeleteScdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type BatchDeleteScdnDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteScdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteScdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteScdnDomainConfigsResponseBody) SetRequestId(v string) *BatchDeleteScdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteScdnDomainConfigsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteScdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteScdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteScdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteScdnDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchDeleteScdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteScdnDomainConfigsResponse) SetStatusCode(v int32) *BatchDeleteScdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteScdnDomainConfigsResponse) SetBody(v *BatchDeleteScdnDomainConfigsResponseBody) *BatchDeleteScdnDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchSetScdnDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetScdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetScdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetScdnDomainConfigsRequest) SetDomainNames(v string) *BatchSetScdnDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetScdnDomainConfigsRequest) SetFunctions(v string) *BatchSetScdnDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetScdnDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetScdnDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetScdnDomainConfigsRequest) SetOwnerId(v int64) *BatchSetScdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetScdnDomainConfigsRequest) SetSecurityToken(v string) *BatchSetScdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type BatchSetScdnDomainConfigsResponseBody struct {
	DomainConfigList *BatchSetScdnDomainConfigsResponseBodyDomainConfigList `json:"DomainConfigList,omitempty" xml:"DomainConfigList,omitempty" type:"Struct"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetScdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetScdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetScdnDomainConfigsResponseBody) SetDomainConfigList(v *BatchSetScdnDomainConfigsResponseBodyDomainConfigList) *BatchSetScdnDomainConfigsResponseBody {
	s.DomainConfigList = v
	return s
}

func (s *BatchSetScdnDomainConfigsResponseBody) SetRequestId(v string) *BatchSetScdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetScdnDomainConfigsResponseBodyDomainConfigList struct {
	DomainConfigModel []*BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel `json:"DomainConfigModel,omitempty" xml:"DomainConfigModel,omitempty" type:"Repeated"`
}

func (s BatchSetScdnDomainConfigsResponseBodyDomainConfigList) String() string {
	return tea.Prettify(s)
}

func (s BatchSetScdnDomainConfigsResponseBodyDomainConfigList) GoString() string {
	return s.String()
}

func (s *BatchSetScdnDomainConfigsResponseBodyDomainConfigList) SetDomainConfigModel(v []*BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) *BatchSetScdnDomainConfigsResponseBodyDomainConfigList {
	s.DomainConfigModel = v
	return s
}

type BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel struct {
	ConfigId     *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) String() string {
	return tea.Prettify(s)
}

func (s BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) GoString() string {
	return s.String()
}

func (s *BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) SetConfigId(v int64) *BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	s.ConfigId = &v
	return s
}

func (s *BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) SetDomainName(v string) *BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	s.DomainName = &v
	return s
}

func (s *BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel) SetFunctionName(v string) *BatchSetScdnDomainConfigsResponseBodyDomainConfigListDomainConfigModel {
	s.FunctionName = &v
	return s
}

type BatchSetScdnDomainConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchSetScdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetScdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetScdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetScdnDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetScdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetScdnDomainConfigsResponse) SetStatusCode(v int32) *BatchSetScdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetScdnDomainConfigsResponse) SetBody(v *BatchSetScdnDomainConfigsResponseBody) *BatchSetScdnDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchStartScdnDomainRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStartScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartScdnDomainRequest) SetDomainNames(v string) *BatchStartScdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStartScdnDomainRequest) SetOwnerId(v int64) *BatchStartScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartScdnDomainRequest) SetSecurityToken(v string) *BatchStartScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type BatchStartScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStartScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartScdnDomainResponseBody) SetRequestId(v string) *BatchStartScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStartScdnDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchStartScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartScdnDomainResponse) SetHeaders(v map[string]*string) *BatchStartScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartScdnDomainResponse) SetStatusCode(v int32) *BatchStartScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartScdnDomainResponse) SetBody(v *BatchStartScdnDomainResponseBody) *BatchStartScdnDomainResponse {
	s.Body = v
	return s
}

type BatchStopScdnDomainRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStopScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopScdnDomainRequest) SetDomainNames(v string) *BatchStopScdnDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStopScdnDomainRequest) SetOwnerId(v int64) *BatchStopScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopScdnDomainRequest) SetSecurityToken(v string) *BatchStopScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type BatchStopScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopScdnDomainResponseBody) SetRequestId(v string) *BatchStopScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStopScdnDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchStopScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopScdnDomainResponse) SetHeaders(v map[string]*string) *BatchStopScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopScdnDomainResponse) SetStatusCode(v int32) *BatchStopScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopScdnDomainResponse) SetBody(v *BatchStopScdnDomainResponseBody) *BatchStopScdnDomainResponse {
	s.Body = v
	return s
}

type BatchUpdateScdnDomainRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain  *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s BatchUpdateScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateScdnDomainRequest) SetDomainName(v string) *BatchUpdateScdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BatchUpdateScdnDomainRequest) SetOwnerId(v int64) *BatchUpdateScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUpdateScdnDomainRequest) SetResourceGroupId(v string) *BatchUpdateScdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchUpdateScdnDomainRequest) SetSecurityToken(v string) *BatchUpdateScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchUpdateScdnDomainRequest) SetSources(v string) *BatchUpdateScdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *BatchUpdateScdnDomainRequest) SetTopLevelDomain(v string) *BatchUpdateScdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type BatchUpdateScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchUpdateScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateScdnDomainResponseBody) SetRequestId(v string) *BatchUpdateScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchUpdateScdnDomainResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateScdnDomainResponse) SetHeaders(v map[string]*string) *BatchUpdateScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateScdnDomainResponse) SetStatusCode(v int32) *BatchUpdateScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateScdnDomainResponse) SetBody(v *BatchUpdateScdnDomainResponseBody) *BatchUpdateScdnDomainResponse {
	s.Body = v
	return s
}

type CheckScdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckScdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckScdnServiceRequest) GoString() string {
	return s.String()
}

func (s *CheckScdnServiceRequest) SetOwnerId(v int64) *CheckScdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckScdnServiceRequest) SetSecurityToken(v string) *CheckScdnServiceRequest {
	s.SecurityToken = &v
	return s
}

type CheckScdnServiceResponseBody struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InDebt        *bool   `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InDebtOverdue *bool   `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	OnService     *bool   `json:"OnService,omitempty" xml:"OnService,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckScdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckScdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckScdnServiceResponseBody) SetEnabled(v bool) *CheckScdnServiceResponseBody {
	s.Enabled = &v
	return s
}

func (s *CheckScdnServiceResponseBody) SetInDebt(v bool) *CheckScdnServiceResponseBody {
	s.InDebt = &v
	return s
}

func (s *CheckScdnServiceResponseBody) SetInDebtOverdue(v bool) *CheckScdnServiceResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *CheckScdnServiceResponseBody) SetOnService(v bool) *CheckScdnServiceResponseBody {
	s.OnService = &v
	return s
}

func (s *CheckScdnServiceResponseBody) SetRequestId(v string) *CheckScdnServiceResponseBody {
	s.RequestId = &v
	return s
}

type CheckScdnServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckScdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckScdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckScdnServiceResponse) GoString() string {
	return s.String()
}

func (s *CheckScdnServiceResponse) SetHeaders(v map[string]*string) *CheckScdnServiceResponse {
	s.Headers = v
	return s
}

func (s *CheckScdnServiceResponse) SetStatusCode(v int32) *CheckScdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckScdnServiceResponse) SetBody(v *CheckScdnServiceResponseBody) *CheckScdnServiceResponse {
	s.Body = v
	return s
}

type DeleteScdnDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteScdnDomainRequest) SetDomainName(v string) *DeleteScdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteScdnDomainRequest) SetOwnerAccount(v string) *DeleteScdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScdnDomainRequest) SetOwnerId(v int64) *DeleteScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScdnDomainRequest) SetSecurityToken(v string) *DeleteScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type DeleteScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScdnDomainResponseBody) SetRequestId(v string) *DeleteScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScdnDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteScdnDomainResponse) SetHeaders(v map[string]*string) *DeleteScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteScdnDomainResponse) SetStatusCode(v int32) *DeleteScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScdnDomainResponse) SetBody(v *DeleteScdnDomainResponseBody) *DeleteScdnDomainResponse {
	s.Body = v
	return s
}

type DeleteScdnSpecificConfigRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteScdnSpecificConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScdnSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteScdnSpecificConfigRequest) SetConfigId(v string) *DeleteScdnSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteScdnSpecificConfigRequest) SetDomainName(v string) *DeleteScdnSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteScdnSpecificConfigRequest) SetOwnerId(v int64) *DeleteScdnSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScdnSpecificConfigRequest) SetSecurityToken(v string) *DeleteScdnSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteScdnSpecificConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScdnSpecificConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScdnSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScdnSpecificConfigResponseBody) SetRequestId(v string) *DeleteScdnSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScdnSpecificConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScdnSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScdnSpecificConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScdnSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteScdnSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteScdnSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteScdnSpecificConfigResponse) SetStatusCode(v int32) *DeleteScdnSpecificConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScdnSpecificConfigResponse) SetBody(v *DeleteScdnSpecificConfigResponseBody) *DeleteScdnSpecificConfigResponse {
	s.Body = v
	return s
}

type DescribeScdnCcInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScdnCcInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcInfoResponseBody) SetRequestId(v string) *DescribeScdnCcInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnCcInfoResponseBody) SetStatus(v string) *DescribeScdnCcInfoResponseBody {
	s.Status = &v
	return s
}

type DescribeScdnCcInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnCcInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnCcInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcInfoResponse) SetHeaders(v map[string]*string) *DescribeScdnCcInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnCcInfoResponse) SetStatusCode(v int32) *DescribeScdnCcInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnCcInfoResponse) SetBody(v *DescribeScdnCcInfoResponseBody) *DescribeScdnCcInfoResponse {
	s.Body = v
	return s
}

type DescribeScdnCcQpsInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnCcQpsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoRequest) SetDomainName(v string) *DescribeScdnCcQpsInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnCcQpsInfoRequest) SetEndTime(v string) *DescribeScdnCcQpsInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnCcQpsInfoRequest) SetStartTime(v string) *DescribeScdnCcQpsInfoRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnCcQpsInfoResponseBody struct {
	Attacks    *DescribeScdnCcQpsInfoResponseBodyAttacks    `json:"Attacks,omitempty" xml:"Attacks,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeScopes *DescribeScdnCcQpsInfoResponseBodyTimeScopes `json:"TimeScopes,omitempty" xml:"TimeScopes,omitempty" type:"Struct"`
	Totals     *DescribeScdnCcQpsInfoResponseBodyTotals     `json:"Totals,omitempty" xml:"Totals,omitempty" type:"Struct"`
}

func (s DescribeScdnCcQpsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoResponseBody) SetAttacks(v *DescribeScdnCcQpsInfoResponseBodyAttacks) *DescribeScdnCcQpsInfoResponseBody {
	s.Attacks = v
	return s
}

func (s *DescribeScdnCcQpsInfoResponseBody) SetRequestId(v string) *DescribeScdnCcQpsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnCcQpsInfoResponseBody) SetTimeScopes(v *DescribeScdnCcQpsInfoResponseBodyTimeScopes) *DescribeScdnCcQpsInfoResponseBody {
	s.TimeScopes = v
	return s
}

func (s *DescribeScdnCcQpsInfoResponseBody) SetTotals(v *DescribeScdnCcQpsInfoResponseBodyTotals) *DescribeScdnCcQpsInfoResponseBody {
	s.Totals = v
	return s
}

type DescribeScdnCcQpsInfoResponseBodyAttacks struct {
	Attack []*string `json:"Attack,omitempty" xml:"Attack,omitempty" type:"Repeated"`
}

func (s DescribeScdnCcQpsInfoResponseBodyAttacks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoResponseBodyAttacks) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoResponseBodyAttacks) SetAttack(v []*string) *DescribeScdnCcQpsInfoResponseBodyAttacks {
	s.Attack = v
	return s
}

type DescribeScdnCcQpsInfoResponseBodyTimeScopes struct {
	TimeScope []*DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Repeated"`
}

func (s DescribeScdnCcQpsInfoResponseBodyTimeScopes) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoResponseBodyTimeScopes) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoResponseBodyTimeScopes) SetTimeScope(v []*DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope) *DescribeScdnCcQpsInfoResponseBodyTimeScopes {
	s.TimeScope = v
	return s
}

type DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope struct {
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope) SetInterval(v string) *DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope {
	s.Interval = &v
	return s
}

func (s *DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope) SetStart(v string) *DescribeScdnCcQpsInfoResponseBodyTimeScopesTimeScope {
	s.Start = &v
	return s
}

type DescribeScdnCcQpsInfoResponseBodyTotals struct {
	Total []*string `json:"Total,omitempty" xml:"Total,omitempty" type:"Repeated"`
}

func (s DescribeScdnCcQpsInfoResponseBodyTotals) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoResponseBodyTotals) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoResponseBodyTotals) SetTotal(v []*string) *DescribeScdnCcQpsInfoResponseBodyTotals {
	s.Total = v
	return s
}

type DescribeScdnCcQpsInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnCcQpsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnCcQpsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcQpsInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcQpsInfoResponse) SetHeaders(v map[string]*string) *DescribeScdnCcQpsInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnCcQpsInfoResponse) SetStatusCode(v int32) *DescribeScdnCcQpsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnCcQpsInfoResponse) SetBody(v *DescribeScdnCcQpsInfoResponseBody) *DescribeScdnCcQpsInfoResponse {
	s.Body = v
	return s
}

type DescribeScdnCcTopIpRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnCcTopIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopIpRequest) SetDomainName(v string) *DescribeScdnCcTopIpRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnCcTopIpRequest) SetEndTime(v string) *DescribeScdnCcTopIpRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnCcTopIpRequest) SetPageNumber(v string) *DescribeScdnCcTopIpRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnCcTopIpRequest) SetPageSize(v string) *DescribeScdnCcTopIpRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnCcTopIpRequest) SetStartTime(v string) *DescribeScdnCcTopIpRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnCcTopIpResponseBody struct {
	AttackIpDataList *DescribeScdnCcTopIpResponseBodyAttackIpDataList `json:"AttackIpDataList,omitempty" xml:"AttackIpDataList,omitempty" type:"Struct"`
	DomainName       *string                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total            *string                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeScdnCcTopIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopIpResponseBody) SetAttackIpDataList(v *DescribeScdnCcTopIpResponseBodyAttackIpDataList) *DescribeScdnCcTopIpResponseBody {
	s.AttackIpDataList = v
	return s
}

func (s *DescribeScdnCcTopIpResponseBody) SetDomainName(v string) *DescribeScdnCcTopIpResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnCcTopIpResponseBody) SetRequestId(v string) *DescribeScdnCcTopIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnCcTopIpResponseBody) SetTotal(v string) *DescribeScdnCcTopIpResponseBody {
	s.Total = &v
	return s
}

type DescribeScdnCcTopIpResponseBodyAttackIpDataList struct {
	AttackIpDatas []*DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas `json:"AttackIpDatas,omitempty" xml:"AttackIpDatas,omitempty" type:"Repeated"`
}

func (s DescribeScdnCcTopIpResponseBodyAttackIpDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopIpResponseBodyAttackIpDataList) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopIpResponseBodyAttackIpDataList) SetAttackIpDatas(v []*DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas) *DescribeScdnCcTopIpResponseBodyAttackIpDataList {
	s.AttackIpDatas = v
	return s
}

type DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas struct {
	AttackCount *string `json:"AttackCount,omitempty" xml:"AttackCount,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas) SetAttackCount(v string) *DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas {
	s.AttackCount = &v
	return s
}

func (s *DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas) SetIp(v string) *DescribeScdnCcTopIpResponseBodyAttackIpDataListAttackIpDatas {
	s.Ip = &v
	return s
}

type DescribeScdnCcTopIpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnCcTopIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnCcTopIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopIpResponse) SetHeaders(v map[string]*string) *DescribeScdnCcTopIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnCcTopIpResponse) SetStatusCode(v int32) *DescribeScdnCcTopIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnCcTopIpResponse) SetBody(v *DescribeScdnCcTopIpResponseBody) *DescribeScdnCcTopIpResponse {
	s.Body = v
	return s
}

type DescribeScdnCcTopUrlRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnCcTopUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopUrlRequest) SetDomainName(v string) *DescribeScdnCcTopUrlRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnCcTopUrlRequest) SetEndTime(v string) *DescribeScdnCcTopUrlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnCcTopUrlRequest) SetPageNumber(v string) *DescribeScdnCcTopUrlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnCcTopUrlRequest) SetPageSize(v string) *DescribeScdnCcTopUrlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnCcTopUrlRequest) SetStartTime(v string) *DescribeScdnCcTopUrlRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnCcTopUrlResponseBody struct {
	AttackUrlDataList *DescribeScdnCcTopUrlResponseBodyAttackUrlDataList `json:"AttackUrlDataList,omitempty" xml:"AttackUrlDataList,omitempty" type:"Struct"`
	DomainName        *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total             *string                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeScdnCcTopUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopUrlResponseBody) SetAttackUrlDataList(v *DescribeScdnCcTopUrlResponseBodyAttackUrlDataList) *DescribeScdnCcTopUrlResponseBody {
	s.AttackUrlDataList = v
	return s
}

func (s *DescribeScdnCcTopUrlResponseBody) SetDomainName(v string) *DescribeScdnCcTopUrlResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnCcTopUrlResponseBody) SetRequestId(v string) *DescribeScdnCcTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnCcTopUrlResponseBody) SetTotal(v string) *DescribeScdnCcTopUrlResponseBody {
	s.Total = &v
	return s
}

type DescribeScdnCcTopUrlResponseBodyAttackUrlDataList struct {
	AttackUrlDatas []*DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas `json:"AttackUrlDatas,omitempty" xml:"AttackUrlDatas,omitempty" type:"Repeated"`
}

func (s DescribeScdnCcTopUrlResponseBodyAttackUrlDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopUrlResponseBodyAttackUrlDataList) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopUrlResponseBodyAttackUrlDataList) SetAttackUrlDatas(v []*DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas) *DescribeScdnCcTopUrlResponseBodyAttackUrlDataList {
	s.AttackUrlDatas = v
	return s
}

type DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas struct {
	AttackCount *string `json:"AttackCount,omitempty" xml:"AttackCount,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas) SetAttackCount(v string) *DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas {
	s.AttackCount = &v
	return s
}

func (s *DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas) SetUrl(v string) *DescribeScdnCcTopUrlResponseBodyAttackUrlDataListAttackUrlDatas {
	s.Url = &v
	return s
}

type DescribeScdnCcTopUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnCcTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnCcTopUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCcTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnCcTopUrlResponse) SetHeaders(v map[string]*string) *DescribeScdnCcTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnCcTopUrlResponse) SetStatusCode(v int32) *DescribeScdnCcTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnCcTopUrlResponse) SetBody(v *DescribeScdnCcTopUrlResponseBody) *DescribeScdnCcTopUrlResponse {
	s.Body = v
	return s
}

type DescribeScdnCertificateDetailRequest struct {
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateDetailRequest) SetCertName(v string) *DescribeScdnCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeScdnCertificateDetailRequest) SetOwnerId(v int64) *DescribeScdnCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnCertificateDetailRequest) SetSecurityToken(v string) *DescribeScdnCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnCertificateDetailResponseBody struct {
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateDetailResponseBody) SetCert(v string) *DescribeScdnCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeScdnCertificateDetailResponseBody) SetCertId(v int64) *DescribeScdnCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeScdnCertificateDetailResponseBody) SetCertName(v string) *DescribeScdnCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeScdnCertificateDetailResponseBody) SetKey(v string) *DescribeScdnCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeScdnCertificateDetailResponseBody) SetRequestId(v string) *DescribeScdnCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnCertificateDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeScdnCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnCertificateDetailResponse) SetStatusCode(v int32) *DescribeScdnCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnCertificateDetailResponse) SetBody(v *DescribeScdnCertificateDetailResponseBody) *DescribeScdnCertificateDetailResponse {
	s.Body = v
	return s
}

type DescribeScdnCertificateListRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateListRequest) SetDomainName(v string) *DescribeScdnCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnCertificateListRequest) SetOwnerId(v int64) *DescribeScdnCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnCertificateListRequest) SetSecurityToken(v string) *DescribeScdnCertificateListRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnCertificateListResponseBody struct {
	CertificateListModel *DescribeScdnCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	RequestId            *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateListResponseBody) SetCertificateListModel(v *DescribeScdnCertificateListResponseBodyCertificateListModel) *DescribeScdnCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeScdnCertificateListResponseBody) SetRequestId(v string) *DescribeScdnCertificateListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnCertificateListResponseBodyCertificateListModel struct {
	CertList *DescribeScdnCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	Count    *int32                                                               `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeScdnCertificateListResponseBodyCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeScdnCertificateListResponseBodyCertificateListModelCertList) *DescribeScdnCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeScdnCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

type DescribeScdnCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeScdnCertificateListResponseBodyCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) *DescribeScdnCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

type DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert struct {
	CertId      *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	LastTime    *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeScdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

type DescribeScdnCertificateListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnCertificateListResponse) SetHeaders(v map[string]*string) *DescribeScdnCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnCertificateListResponse) SetStatusCode(v int32) *DescribeScdnCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnCertificateListResponse) SetBody(v *DescribeScdnCertificateListResponseBody) *DescribeScdnCertificateListResponse {
	s.Body = v
	return s
}

type DescribeScdnDDoSInfoResponseBody struct {
	ElasticBandwidth *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecBandwidth     *int32  `json:"SecBandwidth,omitempty" xml:"SecBandwidth,omitempty"`
}

func (s DescribeScdnDDoSInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSInfoResponseBody) SetElasticBandwidth(v int32) *DescribeScdnDDoSInfoResponseBody {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribeScdnDDoSInfoResponseBody) SetRequestId(v string) *DescribeScdnDDoSInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDDoSInfoResponseBody) SetSecBandwidth(v int32) *DescribeScdnDDoSInfoResponseBody {
	s.SecBandwidth = &v
	return s
}

type DescribeScdnDDoSInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDDoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDDoSInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSInfoResponse) SetHeaders(v map[string]*string) *DescribeScdnDDoSInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDDoSInfoResponse) SetStatusCode(v int32) *DescribeScdnDDoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDDoSInfoResponse) SetBody(v *DescribeScdnDDoSInfoResponseBody) *DescribeScdnDDoSInfoResponse {
	s.Body = v
	return s
}

type DescribeScdnDDoSTrafficInfoRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Line      *string `json:"Line,omitempty" xml:"Line,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDDoSTrafficInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoRequest) SetEndTime(v string) *DescribeScdnDDoSTrafficInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoRequest) SetLine(v string) *DescribeScdnDDoSTrafficInfoRequest {
	s.Line = &v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoRequest) SetStartTime(v string) *DescribeScdnDDoSTrafficInfoRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBody struct {
	BpsDrops   *DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops   `json:"BpsDrops,omitempty" xml:"BpsDrops,omitempty" type:"Struct"`
	BpsTotals  *DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals  `json:"BpsTotals,omitempty" xml:"BpsTotals,omitempty" type:"Struct"`
	PpsDrops   *DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops   `json:"PpsDrops,omitempty" xml:"PpsDrops,omitempty" type:"Struct"`
	PpsTotals  *DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals  `json:"PpsTotals,omitempty" xml:"PpsTotals,omitempty" type:"Struct"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeScopes *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes `json:"TimeScopes,omitempty" xml:"TimeScopes,omitempty" type:"Struct"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBody) SetBpsDrops(v *DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops) *DescribeScdnDDoSTrafficInfoResponseBody {
	s.BpsDrops = v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponseBody) SetBpsTotals(v *DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals) *DescribeScdnDDoSTrafficInfoResponseBody {
	s.BpsTotals = v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponseBody) SetPpsDrops(v *DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops) *DescribeScdnDDoSTrafficInfoResponseBody {
	s.PpsDrops = v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponseBody) SetPpsTotals(v *DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals) *DescribeScdnDDoSTrafficInfoResponseBody {
	s.PpsTotals = v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponseBody) SetRequestId(v string) *DescribeScdnDDoSTrafficInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponseBody) SetTimeScopes(v *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes) *DescribeScdnDDoSTrafficInfoResponseBody {
	s.TimeScopes = v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops struct {
	BpsDrop []*string `json:"BpsDrop,omitempty" xml:"BpsDrop,omitempty" type:"Repeated"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops) SetBpsDrop(v []*string) *DescribeScdnDDoSTrafficInfoResponseBodyBpsDrops {
	s.BpsDrop = v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals struct {
	BpsTotal []*string `json:"BpsTotal,omitempty" xml:"BpsTotal,omitempty" type:"Repeated"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals) SetBpsTotal(v []*string) *DescribeScdnDDoSTrafficInfoResponseBodyBpsTotals {
	s.BpsTotal = v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops struct {
	PpsDrop []*string `json:"PpsDrop,omitempty" xml:"PpsDrop,omitempty" type:"Repeated"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops) SetPpsDrop(v []*string) *DescribeScdnDDoSTrafficInfoResponseBodyPpsDrops {
	s.PpsDrop = v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals struct {
	PpsTotal []*string `json:"PpsTotal,omitempty" xml:"PpsTotal,omitempty" type:"Repeated"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals) SetPpsTotal(v []*string) *DescribeScdnDDoSTrafficInfoResponseBodyPpsTotals {
	s.PpsTotal = v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes struct {
	TimeScope []*DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Repeated"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes) SetTimeScope(v []*DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope) *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopes {
	s.TimeScope = v
	return s
}

type DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope struct {
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope) SetInterval(v string) *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope) SetStart(v string) *DescribeScdnDDoSTrafficInfoResponseBodyTimeScopesTimeScope {
	s.Start = &v
	return s
}

type DescribeScdnDDoSTrafficInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDDoSTrafficInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDDoSTrafficInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDDoSTrafficInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDDoSTrafficInfoResponse) SetHeaders(v map[string]*string) *DescribeScdnDDoSTrafficInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponse) SetStatusCode(v int32) *DescribeScdnDDoSTrafficInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDDoSTrafficInfoResponse) SetBody(v *DescribeScdnDDoSTrafficInfoResponseBody) *DescribeScdnDDoSTrafficInfoResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainBpsDataRequest) SetDomainName(v string) *DescribeScdnDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainBpsDataRequest) SetEndTime(v string) *DescribeScdnDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainBpsDataRequest) SetInterval(v string) *DescribeScdnDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainBpsDataRequest) SetIspNameEn(v string) *DescribeScdnDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainBpsDataRequest) SetStartTime(v string) *DescribeScdnDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeScdnDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBody) SetDomainName(v string) *DescribeScdnDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBody) SetEndTime(v string) *DescribeScdnDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBody) SetRequestId(v string) *DescribeScdnDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBody) SetStartTime(v string) *DescribeScdnDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeScdnDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	BpsValue      *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty"`
	HttpBpsValue  *string `json:"HttpBpsValue,omitempty" xml:"HttpBpsValue,omitempty"`
	HttpsBpsValue *string `json:"HttpsBpsValue,omitempty" xml:"HttpsBpsValue,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpBpsValue(v string) *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpBpsValue = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsBpsValue(v string) *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsBpsValue = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainBpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainBpsDataResponse) SetStatusCode(v int32) *DescribeScdnDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainBpsDataResponse) SetBody(v *DescribeScdnDomainBpsDataResponseBody) *DescribeScdnDomainBpsDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainCertificateInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeScdnDomainCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCertificateInfoRequest) SetDomainName(v string) *DescribeScdnDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

type DescribeScdnDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeScdnDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainCertificateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeScdnDomainCertificateInfoResponseBodyCertInfos) *DescribeScdnDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeScdnDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainCertificateInfoResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeScdnDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertLife       *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg        *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SSLProtocol    *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub         *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLProtocol(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeScdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

type DescribeScdnDomainCertificateInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponse) SetStatusCode(v int32) *DescribeScdnDomainCertificateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainCertificateInfoResponse) SetBody(v *DescribeScdnDomainCertificateInfoResponseBody) *DescribeScdnDomainCertificateInfoResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainCnameRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeScdnDomainCnameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCnameRequest) SetDomainName(v string) *DescribeScdnDomainCnameRequest {
	s.DomainName = &v
	return s
}

type DescribeScdnDomainCnameResponseBody struct {
	CnameDatas *DescribeScdnDomainCnameResponseBodyCnameDatas `json:"CnameDatas,omitempty" xml:"CnameDatas,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainCnameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCnameResponseBody) SetCnameDatas(v *DescribeScdnDomainCnameResponseBodyCnameDatas) *DescribeScdnDomainCnameResponseBody {
	s.CnameDatas = v
	return s
}

func (s *DescribeScdnDomainCnameResponseBody) SetRequestId(v string) *DescribeScdnDomainCnameResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainCnameResponseBodyCnameDatas struct {
	Data []*DescribeScdnDomainCnameResponseBodyCnameDatasData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainCnameResponseBodyCnameDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCnameResponseBodyCnameDatas) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCnameResponseBodyCnameDatas) SetData(v []*DescribeScdnDomainCnameResponseBodyCnameDatasData) *DescribeScdnDomainCnameResponseBodyCnameDatas {
	s.Data = v
	return s
}

type DescribeScdnDomainCnameResponseBodyCnameDatasData struct {
	Cname  *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScdnDomainCnameResponseBodyCnameDatasData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCnameResponseBodyCnameDatasData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCnameResponseBodyCnameDatasData) SetCname(v string) *DescribeScdnDomainCnameResponseBodyCnameDatasData {
	s.Cname = &v
	return s
}

func (s *DescribeScdnDomainCnameResponseBodyCnameDatasData) SetDomain(v string) *DescribeScdnDomainCnameResponseBodyCnameDatasData {
	s.Domain = &v
	return s
}

func (s *DescribeScdnDomainCnameResponseBodyCnameDatasData) SetStatus(v int32) *DescribeScdnDomainCnameResponseBodyCnameDatasData {
	s.Status = &v
	return s
}

type DescribeScdnDomainCnameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainCnameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainCnameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainCnameResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainCnameResponse) SetStatusCode(v int32) *DescribeScdnDomainCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainCnameResponse) SetBody(v *DescribeScdnDomainCnameResponseBody) *DescribeScdnDomainCnameResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainConfigsRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsRequest) SetConfigId(v string) *DescribeScdnDomainConfigsRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeScdnDomainConfigsRequest) SetDomainName(v string) *DescribeScdnDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainConfigsRequest) SetFunctionNames(v string) *DescribeScdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeScdnDomainConfigsRequest) SetOwnerId(v int64) *DescribeScdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnDomainConfigsRequest) SetSecurityToken(v string) *DescribeScdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnDomainConfigsResponseBody struct {
	DomainConfigs *DescribeScdnDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsResponseBody) SetDomainConfigs(v *DescribeScdnDomainConfigsResponseBodyDomainConfigs) *DescribeScdnDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeScdnDomainConfigsResponseBody) SetRequestId(v string) *DescribeScdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeScdnDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

type DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	ConfigId     *string                                                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	FunctionName *string                                                                     `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	ParentId     *string                                                                     `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Status       *string                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetParentId(v string) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ParentId = &v
	return s
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

type DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeScdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeScdnDomainConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainConfigsResponse) SetStatusCode(v int32) *DescribeScdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainConfigsResponse) SetBody(v *DescribeScdnDomainConfigsResponseBody) *DescribeScdnDomainConfigsResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainDetailRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainDetailRequest) SetDomainName(v string) *DescribeScdnDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainDetailRequest) SetOwnerId(v int64) *DescribeScdnDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnDomainDetailRequest) SetSecurityToken(v string) *DescribeScdnDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnDomainDetailResponseBody struct {
	DomainDetail *DescribeScdnDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainDetailResponseBody) SetDomainDetail(v *DescribeScdnDomainDetailResponseBodyDomainDetail) *DescribeScdnDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeScdnDomainDetailResponseBody) SetRequestId(v string) *DescribeScdnDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainDetailResponseBodyDomainDetail struct {
	CertName        *string                                                  `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cname           *string                                                  `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                  `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SSLProtocol     *string                                                  `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub          *string                                                  `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Scope           *string                                                  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Sources         *DescribeScdnDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainDetailResponseBodyDomainDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetResourceGroupId(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeScdnDomainDetailResponseBodyDomainDetailSources) *DescribeScdnDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

type DescribeScdnDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainDetailResponseBodyDomainDetailSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeScdnDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

type DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeScdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

type DescribeScdnDomainDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainDetailResponse) SetStatusCode(v int32) *DescribeScdnDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainDetailResponse) SetBody(v *DescribeScdnDomainDetailResponseBody) *DescribeScdnDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHitRateDataRequest) SetDomainName(v string) *DescribeScdnDomainHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataRequest) SetEndTime(v string) *DescribeScdnDomainHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataRequest) SetInterval(v string) *DescribeScdnDomainHitRateDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataRequest) SetStartTime(v string) *DescribeScdnDomainHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainHitRateDataResponseBody struct {
	DataInterval       *string                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HitRatePerInterval *DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval `json:"HitRatePerInterval,omitempty" xml:"HitRatePerInterval,omitempty" type:"Struct"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeScdnDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeScdnDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBody) SetHitRatePerInterval(v *DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval) *DescribeScdnDomainHitRateDataResponseBody {
	s.HitRatePerInterval = v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeScdnDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeScdnDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval struct {
	DataModule []*DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval) SetDataModule(v []*DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) *DescribeScdnDomainHitRateDataResponseBodyHitRatePerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule struct {
	ByteHitRate *string `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	ReqHitRate  *string `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TimeStamp   *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetByteHitRate(v string) *DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetReqHitRate(v string) *DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainHitRateDataResponseBodyHitRatePerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainHitRateDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponse) SetStatusCode(v int32) *DescribeScdnDomainHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainHitRateDataResponse) SetBody(v *DescribeScdnDomainHitRateDataResponseBody) *DescribeScdnDomainHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataRequest) SetDomainName(v string) *DescribeScdnDomainHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataRequest) SetEndTime(v string) *DescribeScdnDomainHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataRequest) SetInterval(v string) *DescribeScdnDomainHttpCodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataRequest) SetIspNameEn(v string) *DescribeScdnDomainHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataRequest) SetStartTime(v string) *DescribeScdnDomainHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainHttpCodeDataResponseBody struct {
	DataInterval    *string                                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DataPerInterval *DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval `json:"DataPerInterval,omitempty" xml:"DataPerInterval,omitempty" type:"Struct"`
	DomainName      *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBody) SetDataPerInterval(v *DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval) *DescribeScdnDomainHttpCodeDataResponseBody {
	s.DataPerInterval = v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeScdnDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeScdnDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeScdnDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeScdnDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval struct {
	DataModule []*DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval) SetDataModule(v []*DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule struct {
	HttpCodeDataPerInterval *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval `json:"HttpCodeDataPerInterval,omitempty" xml:"HttpCodeDataPerInterval,omitempty" type:"Struct"`
	TimeStamp               *string                                                                                     `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) SetHttpCodeDataPerInterval(v *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	s.HttpCodeDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval struct {
	HttpCodeDataModule []*DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule `json:"HttpCodeDataModule,omitempty" xml:"HttpCodeDataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval) SetHttpCodeDataModule(v []*DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerInterval {
	s.HttpCodeDataModule = v
	return s
}

type DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetCode(v string) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Code = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetCount(v string) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Count = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule) SetProportion(v string) *DescribeScdnDomainHttpCodeDataResponseBodyDataPerIntervalDataModuleHttpCodeDataPerIntervalHttpCodeDataModule {
	s.Proportion = &v
	return s
}

type DescribeScdnDomainHttpCodeDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponse) SetStatusCode(v int32) *DescribeScdnDomainHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainHttpCodeDataResponse) SetBody(v *DescribeScdnDomainHttpCodeDataResponseBody) *DescribeScdnDomainHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainIspDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainIspDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainIspDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainIspDataRequest) SetDomainName(v string) *DescribeScdnDomainIspDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainIspDataRequest) SetEndTime(v string) *DescribeScdnDomainIspDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainIspDataRequest) SetStartTime(v string) *DescribeScdnDomainIspDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainIspDataResponseBody struct {
	DataInterval *string                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeScdnDomainIspDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainIspDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainIspDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainIspDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainIspDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBody) SetDomainName(v string) *DescribeScdnDomainIspDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBody) SetEndTime(v string) *DescribeScdnDomainIspDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBody) SetRequestId(v string) *DescribeScdnDomainIspDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBody) SetStartTime(v string) *DescribeScdnDomainIspDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBody) SetValue(v *DescribeScdnDomainIspDataResponseBodyValue) *DescribeScdnDomainIspDataResponseBody {
	s.Value = v
	return s
}

type DescribeScdnDomainIspDataResponseBodyValue struct {
	ISPProportionData []*DescribeScdnDomainIspDataResponseBodyValueISPProportionData `json:"ISPProportionData,omitempty" xml:"ISPProportionData,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainIspDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainIspDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainIspDataResponseBodyValue) SetISPProportionData(v []*DescribeScdnDomainIspDataResponseBodyValueISPProportionData) *DescribeScdnDomainIspDataResponseBodyValue {
	s.ISPProportionData = v
	return s
}

type DescribeScdnDomainIspDataResponseBodyValueISPProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	ISP             *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IspEname        *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeScdnDomainIspDataResponseBodyValueISPProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainIspDataResponseBodyValueISPProportionData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetAvgObjectSize(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetAvgResponseRate(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetAvgResponseTime(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetBps(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetBytesProportion(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetISP(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.ISP = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetIspEname(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetProportion(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetQps(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetReqErrRate(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetTotalBytes(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetTotalQuery(v string) *DescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeScdnDomainIspDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainIspDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainIspDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainIspDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainIspDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainIspDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainIspDataResponse) SetStatusCode(v int32) *DescribeScdnDomainIspDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainIspDataResponse) SetBody(v *DescribeScdnDomainIspDataResponseBody) *DescribeScdnDomainIspDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainLogRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogRequest) SetDomainName(v string) *DescribeScdnDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainLogRequest) SetEndTime(v string) *DescribeScdnDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainLogRequest) SetPageNumber(v int64) *DescribeScdnDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnDomainLogRequest) SetPageSize(v int64) *DescribeScdnDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnDomainLogRequest) SetStartTime(v string) *DescribeScdnDomainLogRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainLogResponseBody struct {
	DomainLogDetails *DescribeScdnDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponseBody) SetDomainLogDetails(v *DescribeScdnDomainLogResponseBodyDomainLogDetails) *DescribeScdnDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeScdnDomainLogResponseBody) SetDomainName(v string) *DescribeScdnDomainLogResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBody) SetRequestId(v string) *DescribeScdnDomainLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeScdnDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

type DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount  *int64                                                                     `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

type DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

type DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageNumber(v int64) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeScdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

type DescribeScdnDomainLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainLogResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainLogResponse) SetStatusCode(v int32) *DescribeScdnDomainLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainLogResponse) SetBody(v *DescribeScdnDomainLogResponseBody) *DescribeScdnDomainLogResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainOriginBpsDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainOriginBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginBpsDataRequest) SetDomainName(v string) *DescribeScdnDomainOriginBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataRequest) SetEndTime(v string) *DescribeScdnDomainOriginBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataRequest) SetInterval(v string) *DescribeScdnDomainOriginBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataRequest) SetStartTime(v string) *DescribeScdnDomainOriginBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainOriginBpsDataResponseBody struct {
	DataInterval             *string                                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName               *string                                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                  *string                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OriginBpsDataPerInterval *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval `json:"OriginBpsDataPerInterval,omitempty" xml:"OriginBpsDataPerInterval,omitempty" type:"Struct"`
	RequestId                *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainOriginBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginBpsDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainOriginBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBody) SetDomainName(v string) *DescribeScdnDomainOriginBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBody) SetEndTime(v string) *DescribeScdnDomainOriginBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBody) SetOriginBpsDataPerInterval(v *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) *DescribeScdnDomainOriginBpsDataResponseBody {
	s.OriginBpsDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBody) SetRequestId(v string) *DescribeScdnDomainOriginBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBody) SetStartTime(v string) *DescribeScdnDomainOriginBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval struct {
	DataModule []*DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) SetDataModule(v []*DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule struct {
	HttpOriginBpsValue  *string `json:"HttpOriginBpsValue,omitempty" xml:"HttpOriginBpsValue,omitempty"`
	HttpsOriginBpsValue *string `json:"HttpsOriginBpsValue,omitempty" xml:"HttpsOriginBpsValue,omitempty"`
	OriginBpsValue      *string `json:"OriginBpsValue,omitempty" xml:"OriginBpsValue,omitempty"`
	TimeStamp           *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetHttpOriginBpsValue(v string) *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.HttpOriginBpsValue = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetHttpsOriginBpsValue(v string) *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.HttpsOriginBpsValue = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetOriginBpsValue(v string) *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.OriginBpsValue = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainOriginBpsDataResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainOriginBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainOriginBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginBpsDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainOriginBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponse) SetStatusCode(v int32) *DescribeScdnDomainOriginBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainOriginBpsDataResponse) SetBody(v *DescribeScdnDomainOriginBpsDataResponseBody) *DescribeScdnDomainOriginBpsDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainOriginTrafficDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainOriginTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginTrafficDataRequest) SetDomainName(v string) *DescribeScdnDomainOriginTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataRequest) SetEndTime(v string) *DescribeScdnDomainOriginTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataRequest) SetInterval(v string) *DescribeScdnDomainOriginTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataRequest) SetStartTime(v string) *DescribeScdnDomainOriginTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainOriginTrafficDataResponseBody struct {
	DataInterval                 *string                                                                      `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                   *string                                                                      `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                      *string                                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OriginTrafficDataPerInterval *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval `json:"OriginTrafficDataPerInterval,omitempty" xml:"OriginTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                    *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                    *string                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainOriginTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainOriginTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBody) SetDomainName(v string) *DescribeScdnDomainOriginTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBody) SetEndTime(v string) *DescribeScdnDomainOriginTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBody) SetOriginTrafficDataPerInterval(v *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) *DescribeScdnDomainOriginTrafficDataResponseBody {
	s.OriginTrafficDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBody) SetRequestId(v string) *DescribeScdnDomainOriginTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBody) SetStartTime(v string) *DescribeScdnDomainOriginTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval struct {
	DataModule []*DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) SetDataModule(v []*DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule struct {
	HttpTrafficValue  *string `json:"HttpTrafficValue,omitempty" xml:"HttpTrafficValue,omitempty"`
	HttpsTrafficValue *string `json:"HttpsTrafficValue,omitempty" xml:"HttpsTrafficValue,omitempty"`
	TimeStamp         *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrafficValue      *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetHttpTrafficValue(v string) *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.HttpTrafficValue = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetHttpsTrafficValue(v string) *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.HttpsTrafficValue = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeScdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

type DescribeScdnDomainOriginTrafficDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainOriginTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainOriginTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainOriginTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainOriginTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainOriginTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponse) SetStatusCode(v int32) *DescribeScdnDomainOriginTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainOriginTrafficDataResponse) SetBody(v *DescribeScdnDomainOriginTrafficDataResponseBody) *DescribeScdnDomainOriginTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainPvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainPvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainPvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainPvDataRequest) SetDomainName(v string) *DescribeScdnDomainPvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainPvDataRequest) SetEndTime(v string) *DescribeScdnDomainPvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainPvDataRequest) SetStartTime(v string) *DescribeScdnDomainPvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainPvDataResponseBody struct {
	DataInterval   *string                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvDataInterval *DescribeScdnDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainPvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainPvDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainPvDataResponseBody) SetDomainName(v string) *DescribeScdnDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainPvDataResponseBody) SetEndTime(v string) *DescribeScdnDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainPvDataResponseBody) SetPvDataInterval(v *DescribeScdnDomainPvDataResponseBodyPvDataInterval) *DescribeScdnDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeScdnDomainPvDataResponseBody) SetRequestId(v string) *DescribeScdnDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainPvDataResponseBody) SetStartTime(v string) *DescribeScdnDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainPvDataResponseBodyPvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeScdnDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

type DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeScdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeScdnDomainPvDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainPvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainPvDataResponse) SetStatusCode(v int32) *DescribeScdnDomainPvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainPvDataResponse) SetBody(v *DescribeScdnDomainPvDataResponseBody) *DescribeScdnDomainPvDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainQpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainQpsDataRequest) SetDomainName(v string) *DescribeScdnDomainQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainQpsDataRequest) SetEndTime(v string) *DescribeScdnDomainQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainQpsDataRequest) SetInterval(v string) *DescribeScdnDomainQpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainQpsDataRequest) SetIspNameEn(v string) *DescribeScdnDomainQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainQpsDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainQpsDataRequest) SetStartTime(v string) *DescribeScdnDomainQpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainQpsDataResponseBody struct {
	DataInterval       *string                                                  `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	QpsDataPerInterval *DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval `json:"QpsDataPerInterval,omitempty" xml:"QpsDataPerInterval,omitempty" type:"Struct"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBody) SetDomainName(v string) *DescribeScdnDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBody) SetEndTime(v string) *DescribeScdnDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBody) SetQpsDataPerInterval(v *DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval) *DescribeScdnDomainQpsDataResponseBody {
	s.QpsDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBody) SetRequestId(v string) *DescribeScdnDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBody) SetStartTime(v string) *DescribeScdnDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval struct {
	DataModule []*DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval) SetDataModule(v []*DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule struct {
	AccValue      *string `json:"AccValue,omitempty" xml:"AccValue,omitempty"`
	HttpAccValue  *string `json:"HttpAccValue,omitempty" xml:"HttpAccValue,omitempty"`
	HttpQpsValue  *string `json:"HttpQpsValue,omitempty" xml:"HttpQpsValue,omitempty"`
	HttpsAccValue *string `json:"HttpsAccValue,omitempty" xml:"HttpsAccValue,omitempty"`
	HttpsQpsValue *string `json:"HttpsQpsValue,omitempty" xml:"HttpsQpsValue,omitempty"`
	QpsValue      *string `json:"QpsValue,omitempty" xml:"QpsValue,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetAccValue(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.AccValue = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetHttpAccValue(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.HttpAccValue = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetHttpQpsValue(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.HttpQpsValue = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetHttpsAccValue(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.HttpsAccValue = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetHttpsQpsValue(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.HttpsQpsValue = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetQpsValue(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.QpsValue = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainQpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainQpsDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainQpsDataResponse) SetStatusCode(v int32) *DescribeScdnDomainQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainQpsDataResponse) SetBody(v *DescribeScdnDomainQpsDataResponseBody) *DescribeScdnDomainQpsDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeBpsDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataRequest) SetIspNameEn(v string) *DescribeScdnDomainRealTimeBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainRealTimeBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeBpsDataResponseBody struct {
	Data      *DescribeScdnDomainRealTimeBpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainRealTimeBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeBpsDataResponseBody) SetData(v *DescribeScdnDomainRealTimeBpsDataResponseBodyData) *DescribeScdnDomainRealTimeBpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainRealTimeBpsDataResponseBodyData struct {
	BpsModel []*DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel `json:"BpsModel,omitempty" xml:"BpsModel,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeBpsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeBpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeBpsDataResponseBodyData) SetBpsModel(v []*DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel) *DescribeScdnDomainRealTimeBpsDataResponseBodyData {
	s.BpsModel = v
	return s
}

type DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel struct {
	Bps       *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel) SetBps(v float32) *DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.Bps = &v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel) SetTimeStamp(v string) *DescribeScdnDomainRealTimeBpsDataResponseBodyDataBpsModel {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainRealTimeBpsDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeBpsDataResponse) SetBody(v *DescribeScdnDomainRealTimeBpsDataResponseBody) *DescribeScdnDomainRealTimeBpsDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeByteHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeByteHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeByteHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeByteHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeByteHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeByteHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeByteHitRateDataResponseBody struct {
	Data      *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponseBody) SetData(v *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData) *DescribeScdnDomainRealTimeByteHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeByteHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData struct {
	ByteHitRateDataModel []*DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel `json:"ByteHitRateDataModel,omitempty" xml:"ByteHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData) SetByteHitRateDataModel(v []*DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyData {
	s.ByteHitRateDataModel = v
	return s
}

type DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel struct {
	ByteHitRate *float32 `json:"ByteHitRate,omitempty" xml:"ByteHitRate,omitempty"`
	TimeStamp   *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetByteHitRate(v float32) *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.ByteHitRate = &v
	return s
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel) SetTimeStamp(v string) *DescribeScdnDomainRealTimeByteHitRateDataResponseBodyDataByteHitRateDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainRealTimeByteHitRateDataResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeByteHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeByteHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeByteHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeByteHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeByteHitRateDataResponse) SetBody(v *DescribeScdnDomainRealTimeByteHitRateDataResponseBody) *DescribeScdnDomainRealTimeByteHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeHttpCodeDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeHttpCodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataRequest) SetIspNameEn(v string) *DescribeScdnDomainRealTimeHttpCodeDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainRealTimeHttpCodeDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeHttpCodeDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataResponseBody struct {
	DataInterval         *string                                                                 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName           *string                                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime              *string                                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeHttpCodeData *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData `json:"RealTimeHttpCodeData,omitempty" xml:"RealTimeHttpCodeData,omitempty" type:"Struct"`
	RequestId            *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime            *string                                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) SetDomainName(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) SetEndTime(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) SetRealTimeHttpCodeData(v *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeScdnDomainRealTimeHttpCodeDataResponseBody {
	s.RealTimeHttpCodeData = v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) SetStartTime(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData struct {
	UsageData []*DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) SetUsageData(v []*DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	s.UsageData = v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData struct {
	TimeStamp *string                                                                               `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetValue(v *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.Value = v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue struct {
	RealTimeCodeProportionData []*DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData `json:"RealTimeCodeProportionData,omitempty" xml:"RealTimeCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) SetRealTimeCodeProportionData(v []*DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	s.RealTimeCodeProportionData = v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCode(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCount(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetProportion(v string) *DescribeScdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Proportion = &v
	return s
}

type DescribeScdnDomainRealTimeHttpCodeDataResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeScdnDomainRealTimeHttpCodeDataResponseBody) *DescribeScdnDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeQpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeQpsDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeQpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataRequest) SetIspNameEn(v string) *DescribeScdnDomainRealTimeQpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainRealTimeQpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeQpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeQpsDataResponseBody struct {
	Data      *DescribeScdnDomainRealTimeQpsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainRealTimeQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeQpsDataResponseBody) SetData(v *DescribeScdnDomainRealTimeQpsDataResponseBodyData) *DescribeScdnDomainRealTimeQpsDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeQpsDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainRealTimeQpsDataResponseBodyData struct {
	QpsModel []*DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel `json:"QpsModel,omitempty" xml:"QpsModel,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeQpsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeQpsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeQpsDataResponseBodyData) SetQpsModel(v []*DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel) *DescribeScdnDomainRealTimeQpsDataResponseBodyData {
	s.QpsModel = v
	return s
}

type DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel struct {
	Qps       *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	TimeStamp *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel) SetQps(v float32) *DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.Qps = &v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel) SetTimeStamp(v string) *DescribeScdnDomainRealTimeQpsDataResponseBodyDataQpsModel {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainRealTimeQpsDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeQpsDataResponse) SetBody(v *DescribeScdnDomainRealTimeQpsDataResponseBody) *DescribeScdnDomainRealTimeQpsDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeReqHitRateDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeReqHitRateDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeReqHitRateDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeReqHitRateDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeReqHitRateDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeReqHitRateDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeReqHitRateDataResponseBody struct {
	Data      *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponseBody) SetData(v *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData) *DescribeScdnDomainRealTimeReqHitRateDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeReqHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData struct {
	ReqHitRateDataModel []*DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel `json:"ReqHitRateDataModel,omitempty" xml:"ReqHitRateDataModel,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData) SetReqHitRateDataModel(v []*DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyData {
	s.ReqHitRateDataModel = v
	return s
}

type DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel struct {
	ReqHitRate *float32 `json:"ReqHitRate,omitempty" xml:"ReqHitRate,omitempty"`
	TimeStamp  *string  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetReqHitRate(v float32) *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.ReqHitRate = &v
	return s
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel) SetTimeStamp(v string) *DescribeScdnDomainRealTimeReqHitRateDataResponseBodyDataReqHitRateDataModel {
	s.TimeStamp = &v
	return s
}

type DescribeScdnDomainRealTimeReqHitRateDataResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeReqHitRateDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeReqHitRateDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeReqHitRateDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeReqHitRateDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeReqHitRateDataResponse) SetBody(v *DescribeScdnDomainRealTimeReqHitRateDataResponseBody) *DescribeScdnDomainRealTimeReqHitRateDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeSrcBpsDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeSrcBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeSrcBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeSrcBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeSrcBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeSrcBpsDataResponseBody struct {
	DataInterval                  *string                                                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                    *string                                                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                       *string                                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcBpsDataPerInterval *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval,omitempty" xml:"RealTimeSrcBpsDataPerInterval,omitempty" type:"Struct"`
	RequestId                     *string                                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                     *string                                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) SetDomainName(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) SetEndTime(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) SetRealTimeSrcBpsDataPerInterval(v *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeScdnDomainRealTimeSrcBpsDataResponseBody {
	s.RealTimeSrcBpsDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) SetStartTime(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval struct {
	DataModule []*DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) SetDataModule(v []*DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeScdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeScdnDomainRealTimeSrcBpsDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcBpsDataResponse) SetBody(v *DescribeScdnDomainRealTimeSrcBpsDataResponseBody) *DescribeScdnDomainRealTimeSrcBpsDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeSrcTrafficDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeSrcTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeSrcTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeSrcTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeSrcTrafficDataResponseBody struct {
	DataInterval                      *string                                                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                        *string                                                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                           *string                                                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeSrcTrafficDataPerInterval *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval,omitempty" xml:"RealTimeSrcTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                         *string                                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                         *string                                                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) SetRealTimeSrcTrafficDataPerInterval(v *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody {
	s.RealTimeSrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval struct {
	DataModule []*DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) SetDataModule(v []*DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeScdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeScdnDomainRealTimeSrcTrafficDataResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeSrcTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeSrcTrafficDataResponse) SetBody(v *DescribeScdnDomainRealTimeSrcTrafficDataResponseBody) *DescribeScdnDomainRealTimeSrcTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRealTimeTrafficDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeTrafficDataRequest) SetDomainName(v string) *DescribeScdnDomainRealTimeTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataRequest) SetEndTime(v string) *DescribeScdnDomainRealTimeTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataRequest) SetStartTime(v string) *DescribeScdnDomainRealTimeTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeTrafficDataResponseBody struct {
	DataInterval                   *string                                                                          `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName                     *string                                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                        *string                                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeTrafficDataPerInterval *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
	RequestId                      *string                                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime                      *string                                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRealTimeTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeScdnDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeScdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeScdnDomainRealTimeTrafficDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRealTimeTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRealTimeTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRealTimeTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRealTimeTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRealTimeTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRealTimeTrafficDataResponse) SetBody(v *DescribeScdnDomainRealTimeTrafficDataResponseBody) *DescribeScdnDomainRealTimeTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainRegionDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainRegionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRegionDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRegionDataRequest) SetDomainName(v string) *DescribeScdnDomainRegionDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRegionDataRequest) SetEndTime(v string) *DescribeScdnDomainRegionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRegionDataRequest) SetStartTime(v string) *DescribeScdnDomainRegionDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainRegionDataResponseBody struct {
	DataInterval *string                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeScdnDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainRegionDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBody) SetDomainName(v string) *DescribeScdnDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBody) SetEndTime(v string) *DescribeScdnDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBody) SetRequestId(v string) *DescribeScdnDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBody) SetStartTime(v string) *DescribeScdnDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBody) SetValue(v *DescribeScdnDomainRegionDataResponseBodyValue) *DescribeScdnDomainRegionDataResponseBody {
	s.Value = v
	return s
}

type DescribeScdnDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainRegionDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) *DescribeScdnDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

type DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionEname     *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeScdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

type DescribeScdnDomainRegionDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainRegionDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainRegionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainRegionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainRegionDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainRegionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainRegionDataResponse) SetStatusCode(v int32) *DescribeScdnDomainRegionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainRegionDataResponse) SetBody(v *DescribeScdnDomainRegionDataResponseBody) *DescribeScdnDomainRegionDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainTopReferVisitRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainTopReferVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopReferVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopReferVisitRequest) SetDomainName(v string) *DescribeScdnDomainTopReferVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitRequest) SetSortBy(v string) *DescribeScdnDomainTopReferVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitRequest) SetStartTime(v string) *DescribeScdnDomainTopReferVisitRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainTopReferVisitResponseBody struct {
	DomainName   *string                                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId    *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopReferList *DescribeScdnDomainTopReferVisitResponseBodyTopReferList `json:"TopReferList,omitempty" xml:"TopReferList,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainTopReferVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopReferVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopReferVisitResponseBody) SetDomainName(v string) *DescribeScdnDomainTopReferVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBody) SetRequestId(v string) *DescribeScdnDomainTopReferVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBody) SetStartTime(v string) *DescribeScdnDomainTopReferVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBody) SetTopReferList(v *DescribeScdnDomainTopReferVisitResponseBodyTopReferList) *DescribeScdnDomainTopReferVisitResponseBody {
	s.TopReferList = v
	return s
}

type DescribeScdnDomainTopReferVisitResponseBodyTopReferList struct {
	ReferList []*DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList `json:"ReferList,omitempty" xml:"ReferList,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTopReferVisitResponseBodyTopReferList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopReferVisitResponseBodyTopReferList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopReferVisitResponseBodyTopReferList) SetReferList(v []*DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) *DescribeScdnDomainTopReferVisitResponseBodyTopReferList {
	s.ReferList = v
	return s
}

type DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	ReferDetail     *string  `json:"ReferDetail,omitempty" xml:"ReferDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) SetFlow(v string) *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.Flow = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) SetFlowProportion(v float32) *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) SetReferDetail(v string) *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.ReferDetail = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitData(v string) *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitData = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList) SetVisitProportion(v float32) *DescribeScdnDomainTopReferVisitResponseBodyTopReferListReferList {
	s.VisitProportion = &v
	return s
}

type DescribeScdnDomainTopReferVisitResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainTopReferVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainTopReferVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopReferVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopReferVisitResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainTopReferVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponse) SetStatusCode(v int32) *DescribeScdnDomainTopReferVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainTopReferVisitResponse) SetBody(v *DescribeScdnDomainTopReferVisitResponseBody) *DescribeScdnDomainTopReferVisitResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainTopUrlVisitRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainTopUrlVisitRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitRequest) SetDomainName(v string) *DescribeScdnDomainTopUrlVisitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitRequest) SetSortBy(v string) *DescribeScdnDomainTopUrlVisitRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitRequest) SetStartTime(v string) *DescribeScdnDomainTopUrlVisitRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBody struct {
	AllUrlList *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList `json:"AllUrlList,omitempty" xml:"AllUrlList,omitempty" type:"Struct"`
	DomainName *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url200List *DescribeScdnDomainTopUrlVisitResponseBodyUrl200List `json:"Url200List,omitempty" xml:"Url200List,omitempty" type:"Struct"`
	Url300List *DescribeScdnDomainTopUrlVisitResponseBodyUrl300List `json:"Url300List,omitempty" xml:"Url300List,omitempty" type:"Struct"`
	Url400List *DescribeScdnDomainTopUrlVisitResponseBodyUrl400List `json:"Url400List,omitempty" xml:"Url400List,omitempty" type:"Struct"`
	Url500List *DescribeScdnDomainTopUrlVisitResponseBodyUrl500List `json:"Url500List,omitempty" xml:"Url500List,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetAllUrlList(v *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.AllUrlList = v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetDomainName(v string) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetRequestId(v string) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetStartTime(v string) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetUrl200List(v *DescribeScdnDomainTopUrlVisitResponseBodyUrl200List) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.Url200List = v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetUrl300List(v *DescribeScdnDomainTopUrlVisitResponseBodyUrl300List) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.Url300List = v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetUrl400List(v *DescribeScdnDomainTopUrlVisitResponseBodyUrl400List) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.Url400List = v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBody) SetUrl500List(v *DescribeScdnDomainTopUrlVisitResponseBodyUrl500List) *DescribeScdnDomainTopUrlVisitResponseBody {
	s.Url500List = v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList struct {
	UrlList []*DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList) SetUrlList(v []*DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlList {
	s.UrlList = v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlow(v string) *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetFlowProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetUrlDetail(v string) *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitData(v string) *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList) SetVisitProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyAllUrlListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl200List struct {
	UrlList []*DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl200List) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl200List) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl200List) SetUrlList(v []*DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) *DescribeScdnDomainTopUrlVisitResponseBodyUrl200List {
	s.UrlList = v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlow(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetFlowProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetUrlDetail(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitData(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList) SetVisitProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl200ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl300List struct {
	UrlList []*DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl300List) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl300List) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl300List) SetUrlList(v []*DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) *DescribeScdnDomainTopUrlVisitResponseBodyUrl300List {
	s.UrlList = v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlow(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetFlowProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetUrlDetail(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitData(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList) SetVisitProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl300ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl400List struct {
	UrlList []*DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl400List) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl400List) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl400List) SetUrlList(v []*DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) *DescribeScdnDomainTopUrlVisitResponseBodyUrl400List {
	s.UrlList = v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlow(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetFlowProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetUrlDetail(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitData(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList) SetVisitProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl400ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl500List struct {
	UrlList []*DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl500List) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl500List) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl500List) SetUrlList(v []*DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) *DescribeScdnDomainTopUrlVisitResponseBodyUrl500List {
	s.UrlList = v
	return s
}

type DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList struct {
	Flow            *string  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	FlowProportion  *float32 `json:"FlowProportion,omitempty" xml:"FlowProportion,omitempty"`
	UrlDetail       *string  `json:"UrlDetail,omitempty" xml:"UrlDetail,omitempty"`
	VisitData       *string  `json:"VisitData,omitempty" xml:"VisitData,omitempty"`
	VisitProportion *float32 `json:"VisitProportion,omitempty" xml:"VisitProportion,omitempty"`
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlow(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.Flow = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetFlowProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.FlowProportion = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetUrlDetail(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.UrlDetail = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitData(v string) *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitData = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList) SetVisitProportion(v float32) *DescribeScdnDomainTopUrlVisitResponseBodyUrl500ListUrlList {
	s.VisitProportion = &v
	return s
}

type DescribeScdnDomainTopUrlVisitResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainTopUrlVisitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainTopUrlVisitResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTopUrlVisitResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTopUrlVisitResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainTopUrlVisitResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponse) SetStatusCode(v int32) *DescribeScdnDomainTopUrlVisitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainTopUrlVisitResponse) SetBody(v *DescribeScdnDomainTopUrlVisitResponseBody) *DescribeScdnDomainTopUrlVisitResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTrafficDataRequest) SetDomainName(v string) *DescribeScdnDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataRequest) SetEndTime(v string) *DescribeScdnDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataRequest) SetInterval(v string) *DescribeScdnDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeScdnDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeScdnDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataRequest) SetStartTime(v string) *DescribeScdnDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainTrafficDataResponseBody struct {
	DataInterval           *string                                                          `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeScdnDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeScdnDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeScdnDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeScdnDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeScdnDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	HttpTrafficValue  *string `json:"HttpTrafficValue,omitempty" xml:"HttpTrafficValue,omitempty"`
	HttpsTrafficValue *string `json:"HttpsTrafficValue,omitempty" xml:"HttpsTrafficValue,omitempty"`
	TimeStamp         *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrafficValue      *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpTrafficValue(v string) *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpTrafficValue = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsTrafficValue(v string) *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsTrafficValue = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeScdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

type DescribeScdnDomainTrafficDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponse) SetStatusCode(v int32) *DescribeScdnDomainTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainTrafficDataResponse) SetBody(v *DescribeScdnDomainTrafficDataResponseBody) *DescribeScdnDomainTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeScdnDomainUvDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnDomainUvDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainUvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainUvDataRequest) SetDomainName(v string) *DescribeScdnDomainUvDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainUvDataRequest) SetEndTime(v string) *DescribeScdnDomainUvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainUvDataRequest) SetStartTime(v string) *DescribeScdnDomainUvDataRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnDomainUvDataResponseBody struct {
	DataInterval   *string                                             `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName     *string                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime      *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UvDataInterval *DescribeScdnDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeScdnDomainUvDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainUvDataResponseBody) SetDataInterval(v string) *DescribeScdnDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponseBody) SetDomainName(v string) *DescribeScdnDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponseBody) SetEndTime(v string) *DescribeScdnDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponseBody) SetRequestId(v string) *DescribeScdnDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponseBody) SetStartTime(v string) *DescribeScdnDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponseBody) SetUvDataInterval(v *DescribeScdnDomainUvDataResponseBodyUvDataInterval) *DescribeScdnDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

type DescribeScdnDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeScdnDomainUvDataResponseBodyUvDataInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeScdnDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

type DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeScdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

type DescribeScdnDomainUvDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnDomainUvDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeScdnDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnDomainUvDataResponse) SetStatusCode(v int32) *DescribeScdnDomainUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnDomainUvDataResponse) SetBody(v *DescribeScdnDomainUvDataResponseBody) *DescribeScdnDomainUvDataResponse {
	s.Body = v
	return s
}

type DescribeScdnRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnRefreshQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshQuotaRequest) SetOwnerId(v int64) *DescribeScdnRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnRefreshQuotaRequest) SetSecurityToken(v string) *DescribeScdnRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnRefreshQuotaResponseBody struct {
	BlockQuota    *string `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	DirQuota      *string `json:"DirQuota,omitempty" xml:"DirQuota,omitempty"`
	DirRemain     *string `json:"DirRemain,omitempty" xml:"DirRemain,omitempty"`
	PreloadQuota  *string `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	PreloadRemain *string `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlQuota      *string `json:"UrlQuota,omitempty" xml:"UrlQuota,omitempty"`
	UrlRemain     *string `json:"UrlRemain,omitempty" xml:"UrlRemain,omitempty"`
	BlockRemain   *string `json:"blockRemain,omitempty" xml:"blockRemain,omitempty"`
}

func (s DescribeScdnRefreshQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetRequestId(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeScdnRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

type DescribeScdnRefreshQuotaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnRefreshQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeScdnRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnRefreshQuotaResponse) SetStatusCode(v int32) *DescribeScdnRefreshQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnRefreshQuotaResponse) SetBody(v *DescribeScdnRefreshQuotaResponseBody) *DescribeScdnRefreshQuotaResponse {
	s.Body = v
	return s
}

type DescribeScdnRefreshTasksRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ObjectPath      *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType      *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScdnRefreshTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshTasksRequest) SetDomainName(v string) *DescribeScdnRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetEndTime(v string) *DescribeScdnRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetObjectPath(v string) *DescribeScdnRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetObjectType(v string) *DescribeScdnRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetOwnerId(v int64) *DescribeScdnRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetPageNumber(v int32) *DescribeScdnRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetPageSize(v int32) *DescribeScdnRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetResourceGroupId(v string) *DescribeScdnRefreshTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetSecurityToken(v string) *DescribeScdnRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetStartTime(v string) *DescribeScdnRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetStatus(v string) *DescribeScdnRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeScdnRefreshTasksRequest) SetTaskId(v string) *DescribeScdnRefreshTasksRequest {
	s.TaskId = &v
	return s
}

type DescribeScdnRefreshTasksResponseBody struct {
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      *DescribeScdnRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScdnRefreshTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeScdnRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBody) SetPageSize(v int64) *DescribeScdnRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBody) SetRequestId(v string) *DescribeScdnRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBody) SetTasks(v *DescribeScdnRefreshTasksResponseBodyTasks) *DescribeScdnRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeScdnRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScdnRefreshTasksResponseBodyTasks struct {
	Task []*DescribeScdnRefreshTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeScdnRefreshTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshTasksResponseBodyTasks) SetTask(v []*DescribeScdnRefreshTasksResponseBodyTasksTask) *DescribeScdnRefreshTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeScdnRefreshTasksResponseBodyTasksTask struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScdnRefreshTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetCreationTime(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetDescription(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.Description = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetObjectPath(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetObjectType(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetProcess(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.Process = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetStatus(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeScdnRefreshTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

type DescribeScdnRefreshTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnRefreshTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeScdnRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnRefreshTasksResponse) SetStatusCode(v int32) *DescribeScdnRefreshTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnRefreshTasksResponse) SetBody(v *DescribeScdnRefreshTasksResponseBody) *DescribeScdnRefreshTasksResponse {
	s.Body = v
	return s
}

type DescribeScdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnServiceRequest) SetOwnerId(v int64) *DescribeScdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnServiceRequest) SetSecurityToken(v string) *DescribeScdnServiceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnServiceResponseBody struct {
	Bandwidth                     *string                                        `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthValue                *string                                        `json:"BandwidthValue,omitempty" xml:"BandwidthValue,omitempty"`
	CcProtection                  *string                                        `json:"CcProtection,omitempty" xml:"CcProtection,omitempty"`
	CcProtectionValue             *string                                        `json:"CcProtectionValue,omitempty" xml:"CcProtectionValue,omitempty"`
	ChangingAffectTime            *string                                        `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	ChangingChargeType            *string                                        `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	CurrentBandwidth              *string                                        `json:"CurrentBandwidth,omitempty" xml:"CurrentBandwidth,omitempty"`
	CurrentBandwidthValue         *string                                        `json:"CurrentBandwidthValue,omitempty" xml:"CurrentBandwidthValue,omitempty"`
	CurrentCcProtection           *string                                        `json:"CurrentCcProtection,omitempty" xml:"CurrentCcProtection,omitempty"`
	CurrentCcProtectionValue      *string                                        `json:"CurrentCcProtectionValue,omitempty" xml:"CurrentCcProtectionValue,omitempty"`
	CurrentDDoSBasic              *string                                        `json:"CurrentDDoSBasic,omitempty" xml:"CurrentDDoSBasic,omitempty"`
	CurrentDDoSBasicValue         *string                                        `json:"CurrentDDoSBasicValue,omitempty" xml:"CurrentDDoSBasicValue,omitempty"`
	CurrentDomainCount            *string                                        `json:"CurrentDomainCount,omitempty" xml:"CurrentDomainCount,omitempty"`
	CurrentDomainCountValue       *string                                        `json:"CurrentDomainCountValue,omitempty" xml:"CurrentDomainCountValue,omitempty"`
	CurrentElasticProtection      *string                                        `json:"CurrentElasticProtection,omitempty" xml:"CurrentElasticProtection,omitempty"`
	CurrentElasticProtectionValue *string                                        `json:"CurrentElasticProtectionValue,omitempty" xml:"CurrentElasticProtectionValue,omitempty"`
	CurrentProtectType            *string                                        `json:"CurrentProtectType,omitempty" xml:"CurrentProtectType,omitempty"`
	CurrentProtectTypeValue       *string                                        `json:"CurrentProtectTypeValue,omitempty" xml:"CurrentProtectTypeValue,omitempty"`
	DDoSBasic                     *string                                        `json:"DDoSBasic,omitempty" xml:"DDoSBasic,omitempty"`
	DDoSBasicValue                *string                                        `json:"DDoSBasicValue,omitempty" xml:"DDoSBasicValue,omitempty"`
	DomainCount                   *string                                        `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainCountValue              *string                                        `json:"DomainCountValue,omitempty" xml:"DomainCountValue,omitempty"`
	ElasticProtection             *string                                        `json:"ElasticProtection,omitempty" xml:"ElasticProtection,omitempty"`
	ElasticProtectionValue        *string                                        `json:"ElasticProtectionValue,omitempty" xml:"ElasticProtectionValue,omitempty"`
	EndTime                       *string                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId                    *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetChargeType            *string                                        `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OpenTime                      *string                                        `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	OperationLocks                *DescribeScdnServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	PriceType                     *string                                        `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	PricingCycle                  *string                                        `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ProtectType                   *string                                        `json:"ProtectType,omitempty" xml:"ProtectType,omitempty"`
	ProtectTypeValue              *string                                        `json:"ProtectTypeValue,omitempty" xml:"ProtectTypeValue,omitempty"`
	RequestId                     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnServiceResponseBody) SetBandwidth(v string) *DescribeScdnServiceResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetBandwidthValue(v string) *DescribeScdnServiceResponseBody {
	s.BandwidthValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCcProtection(v string) *DescribeScdnServiceResponseBody {
	s.CcProtection = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCcProtectionValue(v string) *DescribeScdnServiceResponseBody {
	s.CcProtectionValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetChangingAffectTime(v string) *DescribeScdnServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetChangingChargeType(v string) *DescribeScdnServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentBandwidth(v string) *DescribeScdnServiceResponseBody {
	s.CurrentBandwidth = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentBandwidthValue(v string) *DescribeScdnServiceResponseBody {
	s.CurrentBandwidthValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentCcProtection(v string) *DescribeScdnServiceResponseBody {
	s.CurrentCcProtection = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentCcProtectionValue(v string) *DescribeScdnServiceResponseBody {
	s.CurrentCcProtectionValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentDDoSBasic(v string) *DescribeScdnServiceResponseBody {
	s.CurrentDDoSBasic = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentDDoSBasicValue(v string) *DescribeScdnServiceResponseBody {
	s.CurrentDDoSBasicValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentDomainCount(v string) *DescribeScdnServiceResponseBody {
	s.CurrentDomainCount = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentDomainCountValue(v string) *DescribeScdnServiceResponseBody {
	s.CurrentDomainCountValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentElasticProtection(v string) *DescribeScdnServiceResponseBody {
	s.CurrentElasticProtection = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentElasticProtectionValue(v string) *DescribeScdnServiceResponseBody {
	s.CurrentElasticProtectionValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentProtectType(v string) *DescribeScdnServiceResponseBody {
	s.CurrentProtectType = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetCurrentProtectTypeValue(v string) *DescribeScdnServiceResponseBody {
	s.CurrentProtectTypeValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetDDoSBasic(v string) *DescribeScdnServiceResponseBody {
	s.DDoSBasic = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetDDoSBasicValue(v string) *DescribeScdnServiceResponseBody {
	s.DDoSBasicValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetDomainCount(v string) *DescribeScdnServiceResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetDomainCountValue(v string) *DescribeScdnServiceResponseBody {
	s.DomainCountValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetElasticProtection(v string) *DescribeScdnServiceResponseBody {
	s.ElasticProtection = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetElasticProtectionValue(v string) *DescribeScdnServiceResponseBody {
	s.ElasticProtectionValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetEndTime(v string) *DescribeScdnServiceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetInstanceId(v string) *DescribeScdnServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetInternetChargeType(v string) *DescribeScdnServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetOpenTime(v string) *DescribeScdnServiceResponseBody {
	s.OpenTime = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetOperationLocks(v *DescribeScdnServiceResponseBodyOperationLocks) *DescribeScdnServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetPriceType(v string) *DescribeScdnServiceResponseBody {
	s.PriceType = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetPricingCycle(v string) *DescribeScdnServiceResponseBody {
	s.PricingCycle = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetProtectType(v string) *DescribeScdnServiceResponseBody {
	s.ProtectType = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetProtectTypeValue(v string) *DescribeScdnServiceResponseBody {
	s.ProtectTypeValue = &v
	return s
}

func (s *DescribeScdnServiceResponseBody) SetRequestId(v string) *DescribeScdnServiceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeScdnServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeScdnServiceResponseBodyOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeScdnServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeScdnServiceResponseBodyOperationLocksLockReason) *DescribeScdnServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

type DescribeScdnServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeScdnServiceResponseBodyOperationLocksLockReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeScdnServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeScdnServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

type DescribeScdnServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnServiceResponse) SetHeaders(v map[string]*string) *DescribeScdnServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnServiceResponse) SetStatusCode(v int32) *DescribeScdnServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnServiceResponse) SetBody(v *DescribeScdnServiceResponseBody) *DescribeScdnServiceResponse {
	s.Body = v
	return s
}

type DescribeScdnTopDomainsByFlowRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScdnTopDomainsByFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnTopDomainsByFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnTopDomainsByFlowRequest) SetEndTime(v string) *DescribeScdnTopDomainsByFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowRequest) SetLimit(v int64) *DescribeScdnTopDomainsByFlowRequest {
	s.Limit = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowRequest) SetProduct(v string) *DescribeScdnTopDomainsByFlowRequest {
	s.Product = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowRequest) SetStartTime(v string) *DescribeScdnTopDomainsByFlowRequest {
	s.StartTime = &v
	return s
}

type DescribeScdnTopDomainsByFlowResponseBody struct {
	DomainCount       *int64                                              `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainOnlineCount *int64                                              `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	EndTime           *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime         *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains        *DescribeScdnTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeScdnTopDomainsByFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeScdnTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeScdnTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeScdnTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeScdnTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeScdnTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeScdnTopDomainsByFlowResponseBodyTopDomains) *DescribeScdnTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

type DescribeScdnTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeScdnTopDomainsByFlowResponseBodyTopDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeScdnTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

type DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MaxBps         *int64  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	Rank           *int64  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64  `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeScdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

type DescribeScdnTopDomainsByFlowResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnTopDomainsByFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnTopDomainsByFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnTopDomainsByFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnTopDomainsByFlowResponse) SetHeaders(v map[string]*string) *DescribeScdnTopDomainsByFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponse) SetStatusCode(v int32) *DescribeScdnTopDomainsByFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnTopDomainsByFlowResponse) SetBody(v *DescribeScdnTopDomainsByFlowResponseBody) *DescribeScdnTopDomainsByFlowResponse {
	s.Body = v
	return s
}

type DescribeScdnUserDomainsRequest struct {
	ChangeEndTime    *string `json:"ChangeEndTime,omitempty" xml:"ChangeEndTime,omitempty"`
	ChangeStartTime  *string `json:"ChangeStartTime,omitempty" xml:"ChangeStartTime,omitempty"`
	CheckDomainShow  *bool   `json:"CheckDomainShow,omitempty" xml:"CheckDomainShow,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainSearchType *string `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	DomainStatus     *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsRequest) SetChangeEndTime(v string) *DescribeScdnUserDomainsRequest {
	s.ChangeEndTime = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetChangeStartTime(v string) *DescribeScdnUserDomainsRequest {
	s.ChangeStartTime = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetCheckDomainShow(v bool) *DescribeScdnUserDomainsRequest {
	s.CheckDomainShow = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetDomainName(v string) *DescribeScdnUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetDomainSearchType(v string) *DescribeScdnUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetDomainStatus(v string) *DescribeScdnUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetOwnerId(v int64) *DescribeScdnUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetPageNumber(v int32) *DescribeScdnUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetPageSize(v int32) *DescribeScdnUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetResourceGroupId(v string) *DescribeScdnUserDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScdnUserDomainsRequest) SetSecurityToken(v string) *DescribeScdnUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnUserDomainsResponseBody struct {
	Domains    *DescribeScdnUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScdnUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsResponseBody) SetDomains(v *DescribeScdnUserDomainsResponseBodyDomains) *DescribeScdnUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeScdnUserDomainsResponseBody) SetPageNumber(v int64) *DescribeScdnUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBody) SetPageSize(v int64) *DescribeScdnUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBody) SetRequestId(v string) *DescribeScdnUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBody) SetTotalCount(v int64) *DescribeScdnUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScdnUserDomainsResponseBodyDomains struct {
	PageData []*DescribeScdnUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeScdnUserDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsResponseBodyDomains) SetPageData(v []*DescribeScdnUserDomainsResponseBodyDomainsPageData) *DescribeScdnUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeScdnUserDomainsResponseBodyDomainsPageData struct {
	Cname           *string                                                    `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description     *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName      *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus    *string                                                    `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated      *string                                                    `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string                                                    `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ResourceGroupId *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SSLProtocol     *string                                                    `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	Sandbox         *string                                                    `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Sources         *DescribeScdnUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeScdnUserDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetResourceGroupId(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetSSLProtocol(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeScdnUserDomainsResponseBodyDomainsPageDataSources) *DescribeScdnUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

type DescribeScdnUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeScdnUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeScdnUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeScdnUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

type DescribeScdnUserDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeScdnUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnUserDomainsResponse) SetStatusCode(v int32) *DescribeScdnUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnUserDomainsResponse) SetBody(v *DescribeScdnUserDomainsResponseBody) *DescribeScdnUserDomainsResponse {
	s.Body = v
	return s
}

type DescribeScdnUserProtectInfoResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceDDoS *int32  `json:"ServiceDDoS,omitempty" xml:"ServiceDDoS,omitempty"`
}

func (s DescribeScdnUserProtectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserProtectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserProtectInfoResponseBody) SetRequestId(v string) *DescribeScdnUserProtectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScdnUserProtectInfoResponseBody) SetServiceDDoS(v int32) *DescribeScdnUserProtectInfoResponseBody {
	s.ServiceDDoS = &v
	return s
}

type DescribeScdnUserProtectInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnUserProtectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnUserProtectInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserProtectInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserProtectInfoResponse) SetHeaders(v map[string]*string) *DescribeScdnUserProtectInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnUserProtectInfoResponse) SetStatusCode(v int32) *DescribeScdnUserProtectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnUserProtectInfoResponse) SetBody(v *DescribeScdnUserProtectInfoResponseBody) *DescribeScdnUserProtectInfoResponse {
	s.Body = v
	return s
}

type DescribeScdnUserQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeScdnUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserQuotaRequest) SetOwnerId(v int64) *DescribeScdnUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScdnUserQuotaRequest) SetSecurityToken(v string) *DescribeScdnUserQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeScdnUserQuotaResponseBody struct {
	BlockQuota       *int32  `json:"BlockQuota,omitempty" xml:"BlockQuota,omitempty"`
	BlockRemain      *int32  `json:"BlockRemain,omitempty" xml:"BlockRemain,omitempty"`
	DomainQuota      *int32  `json:"DomainQuota,omitempty" xml:"DomainQuota,omitempty"`
	PreloadQuota     *int32  `json:"PreloadQuota,omitempty" xml:"PreloadQuota,omitempty"`
	PreloadRemain    *int32  `json:"PreloadRemain,omitempty" xml:"PreloadRemain,omitempty"`
	RefreshDirQuota  *int32  `json:"RefreshDirQuota,omitempty" xml:"RefreshDirQuota,omitempty"`
	RefreshDirRemain *int32  `json:"RefreshDirRemain,omitempty" xml:"RefreshDirRemain,omitempty"`
	RefreshUrlQuota  *int32  `json:"RefreshUrlQuota,omitempty" xml:"RefreshUrlQuota,omitempty"`
	RefreshUrlRemain *int32  `json:"RefreshUrlRemain,omitempty" xml:"RefreshUrlRemain,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserQuotaResponseBody) SetBlockQuota(v int32) *DescribeScdnUserQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetBlockRemain(v int32) *DescribeScdnUserQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetDomainQuota(v int32) *DescribeScdnUserQuotaResponseBody {
	s.DomainQuota = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetPreloadQuota(v int32) *DescribeScdnUserQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetPreloadRemain(v int32) *DescribeScdnUserQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetRefreshDirQuota(v int32) *DescribeScdnUserQuotaResponseBody {
	s.RefreshDirQuota = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetRefreshDirRemain(v int32) *DescribeScdnUserQuotaResponseBody {
	s.RefreshDirRemain = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetRefreshUrlQuota(v int32) *DescribeScdnUserQuotaResponseBody {
	s.RefreshUrlQuota = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetRefreshUrlRemain(v int32) *DescribeScdnUserQuotaResponseBody {
	s.RefreshUrlRemain = &v
	return s
}

func (s *DescribeScdnUserQuotaResponseBody) SetRequestId(v string) *DescribeScdnUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnUserQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeScdnUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnUserQuotaResponse) SetStatusCode(v int32) *DescribeScdnUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnUserQuotaResponse) SetBody(v *DescribeScdnUserQuotaResponseBody) *DescribeScdnUserQuotaResponse {
	s.Body = v
	return s
}

type DescribeScdnVerifyContentRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeScdnVerifyContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeScdnVerifyContentRequest) SetDomainName(v string) *DescribeScdnVerifyContentRequest {
	s.DomainName = &v
	return s
}

type DescribeScdnVerifyContentResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScdnVerifyContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScdnVerifyContentResponseBody) SetContent(v string) *DescribeScdnVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeScdnVerifyContentResponseBody) SetRequestId(v string) *DescribeScdnVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScdnVerifyContentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScdnVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScdnVerifyContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScdnVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeScdnVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeScdnVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeScdnVerifyContentResponse) SetStatusCode(v int32) *DescribeScdnVerifyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScdnVerifyContentResponse) SetBody(v *DescribeScdnVerifyContentResponseBody) *DescribeScdnVerifyContentResponse {
	s.Body = v
	return s
}

type PreloadScdnObjectCachesRequest struct {
	Area          *string `json:"Area,omitempty" xml:"Area,omitempty"`
	L2Preload     *bool   `json:"L2Preload,omitempty" xml:"L2Preload,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PreloadScdnObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s PreloadScdnObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadScdnObjectCachesRequest) SetArea(v string) *PreloadScdnObjectCachesRequest {
	s.Area = &v
	return s
}

func (s *PreloadScdnObjectCachesRequest) SetL2Preload(v bool) *PreloadScdnObjectCachesRequest {
	s.L2Preload = &v
	return s
}

func (s *PreloadScdnObjectCachesRequest) SetObjectPath(v string) *PreloadScdnObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *PreloadScdnObjectCachesRequest) SetOwnerId(v int64) *PreloadScdnObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *PreloadScdnObjectCachesRequest) SetSecurityToken(v string) *PreloadScdnObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type PreloadScdnObjectCachesResponseBody struct {
	PreloadTaskId *string `json:"PreloadTaskId,omitempty" xml:"PreloadTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadScdnObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreloadScdnObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadScdnObjectCachesResponseBody) SetPreloadTaskId(v string) *PreloadScdnObjectCachesResponseBody {
	s.PreloadTaskId = &v
	return s
}

func (s *PreloadScdnObjectCachesResponseBody) SetRequestId(v string) *PreloadScdnObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type PreloadScdnObjectCachesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PreloadScdnObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreloadScdnObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s PreloadScdnObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadScdnObjectCachesResponse) SetHeaders(v map[string]*string) *PreloadScdnObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadScdnObjectCachesResponse) SetStatusCode(v int32) *PreloadScdnObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadScdnObjectCachesResponse) SetBody(v *PreloadScdnObjectCachesResponseBody) *PreloadScdnObjectCachesResponse {
	s.Body = v
	return s
}

type RefreshScdnObjectCachesRequest struct {
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshScdnObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshScdnObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshScdnObjectCachesRequest) SetObjectPath(v string) *RefreshScdnObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshScdnObjectCachesRequest) SetObjectType(v string) *RefreshScdnObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshScdnObjectCachesRequest) SetOwnerId(v int64) *RefreshScdnObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshScdnObjectCachesRequest) SetSecurityToken(v string) *RefreshScdnObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type RefreshScdnObjectCachesResponseBody struct {
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshScdnObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshScdnObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshScdnObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshScdnObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshScdnObjectCachesResponseBody) SetRequestId(v string) *RefreshScdnObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type RefreshScdnObjectCachesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshScdnObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshScdnObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshScdnObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshScdnObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshScdnObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshScdnObjectCachesResponse) SetStatusCode(v int32) *RefreshScdnObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshScdnObjectCachesResponse) SetBody(v *RefreshScdnObjectCachesResponseBody) *RefreshScdnObjectCachesResponse {
	s.Body = v
	return s
}

type SetScdnBotInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetScdnBotInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetScdnBotInfoRequest) GoString() string {
	return s.String()
}

func (s *SetScdnBotInfoRequest) SetDomainName(v string) *SetScdnBotInfoRequest {
	s.DomainName = &v
	return s
}

func (s *SetScdnBotInfoRequest) SetEnable(v string) *SetScdnBotInfoRequest {
	s.Enable = &v
	return s
}

func (s *SetScdnBotInfoRequest) SetStatus(v string) *SetScdnBotInfoRequest {
	s.Status = &v
	return s
}

type SetScdnBotInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetScdnBotInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetScdnBotInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetScdnBotInfoResponseBody) SetRequestId(v string) *SetScdnBotInfoResponseBody {
	s.RequestId = &v
	return s
}

type SetScdnBotInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetScdnBotInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetScdnBotInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetScdnBotInfoResponse) GoString() string {
	return s.String()
}

func (s *SetScdnBotInfoResponse) SetHeaders(v map[string]*string) *SetScdnBotInfoResponse {
	s.Headers = v
	return s
}

func (s *SetScdnBotInfoResponse) SetStatusCode(v int32) *SetScdnBotInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetScdnBotInfoResponse) SetBody(v *SetScdnBotInfoResponseBody) *SetScdnBotInfoResponse {
	s.Body = v
	return s
}

type SetScdnCcInfoRequest struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetScdnCcInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetScdnCcInfoRequest) GoString() string {
	return s.String()
}

func (s *SetScdnCcInfoRequest) SetStatus(v string) *SetScdnCcInfoRequest {
	s.Status = &v
	return s
}

type SetScdnCcInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetScdnCcInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetScdnCcInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetScdnCcInfoResponseBody) SetRequestId(v string) *SetScdnCcInfoResponseBody {
	s.RequestId = &v
	return s
}

type SetScdnCcInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetScdnCcInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetScdnCcInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetScdnCcInfoResponse) GoString() string {
	return s.String()
}

func (s *SetScdnCcInfoResponse) SetHeaders(v map[string]*string) *SetScdnCcInfoResponse {
	s.Headers = v
	return s
}

func (s *SetScdnCcInfoResponse) SetStatusCode(v int32) *SetScdnCcInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetScdnCcInfoResponse) SetBody(v *SetScdnCcInfoResponseBody) *SetScdnCcInfoResponse {
	s.Body = v
	return s
}

type SetScdnDDoSInfoRequest struct {
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
}

func (s SetScdnDDoSInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetScdnDDoSInfoRequest) GoString() string {
	return s.String()
}

func (s *SetScdnDDoSInfoRequest) SetElasticBandwidth(v int32) *SetScdnDDoSInfoRequest {
	s.ElasticBandwidth = &v
	return s
}

type SetScdnDDoSInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetScdnDDoSInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetScdnDDoSInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetScdnDDoSInfoResponseBody) SetRequestId(v string) *SetScdnDDoSInfoResponseBody {
	s.RequestId = &v
	return s
}

type SetScdnDDoSInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetScdnDDoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetScdnDDoSInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetScdnDDoSInfoResponse) GoString() string {
	return s.String()
}

func (s *SetScdnDDoSInfoResponse) SetHeaders(v map[string]*string) *SetScdnDDoSInfoResponse {
	s.Headers = v
	return s
}

func (s *SetScdnDDoSInfoResponse) SetStatusCode(v int32) *SetScdnDDoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetScdnDDoSInfoResponse) SetBody(v *SetScdnDDoSInfoResponseBody) *SetScdnDDoSInfoResponse {
	s.Body = v
	return s
}

type SetScdnDomainCertificateRequest struct {
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType      *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ForceSet      *string `json:"ForceSet,omitempty" xml:"ForceSet,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SSLPri        *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetScdnDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetScdnDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetScdnDomainCertificateRequest) SetCertName(v string) *SetScdnDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetCertType(v string) *SetScdnDomainCertificateRequest {
	s.CertType = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetDomainName(v string) *SetScdnDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetForceSet(v string) *SetScdnDomainCertificateRequest {
	s.ForceSet = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetOwnerId(v int64) *SetScdnDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetRegion(v string) *SetScdnDomainCertificateRequest {
	s.Region = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetSSLPri(v string) *SetScdnDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetSSLProtocol(v string) *SetScdnDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetSSLPub(v string) *SetScdnDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetScdnDomainCertificateRequest) SetSecurityToken(v string) *SetScdnDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

type SetScdnDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetScdnDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetScdnDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetScdnDomainCertificateResponseBody) SetRequestId(v string) *SetScdnDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetScdnDomainCertificateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetScdnDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetScdnDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetScdnDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetScdnDomainCertificateResponse) SetHeaders(v map[string]*string) *SetScdnDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetScdnDomainCertificateResponse) SetStatusCode(v int32) *SetScdnDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetScdnDomainCertificateResponse) SetBody(v *SetScdnDomainCertificateResponseBody) *SetScdnDomainCertificateResponse {
	s.Body = v
	return s
}

type StartScdnDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StartScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StartScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StartScdnDomainRequest) SetDomainName(v string) *StartScdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StartScdnDomainRequest) SetOwnerId(v int64) *StartScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StartScdnDomainRequest) SetSecurityToken(v string) *StartScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type StartScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartScdnDomainResponseBody) SetRequestId(v string) *StartScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type StartScdnDomainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StartScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StartScdnDomainResponse) SetHeaders(v map[string]*string) *StartScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StartScdnDomainResponse) SetStatusCode(v int32) *StartScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StartScdnDomainResponse) SetBody(v *StartScdnDomainResponseBody) *StartScdnDomainResponse {
	s.Body = v
	return s
}

type StopScdnDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StopScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s StopScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *StopScdnDomainRequest) SetDomainName(v string) *StopScdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *StopScdnDomainRequest) SetOwnerId(v int64) *StopScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *StopScdnDomainRequest) SetSecurityToken(v string) *StopScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

type StopScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopScdnDomainResponseBody) SetRequestId(v string) *StopScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type StopScdnDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s StopScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StopScdnDomainResponse) SetHeaders(v map[string]*string) *StopScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StopScdnDomainResponse) SetStatusCode(v int32) *StopScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StopScdnDomainResponse) SetBody(v *StopScdnDomainResponseBody) *StopScdnDomainResponse {
	s.Body = v
	return s
}

type TestAmpDescribeScdnDomainIspDataRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TestAmpDescribeScdnDomainIspDataRequest) String() string {
	return tea.Prettify(s)
}

func (s TestAmpDescribeScdnDomainIspDataRequest) GoString() string {
	return s.String()
}

func (s *TestAmpDescribeScdnDomainIspDataRequest) SetDomainName(v string) *TestAmpDescribeScdnDomainIspDataRequest {
	s.DomainName = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataRequest) SetEndTime(v string) *TestAmpDescribeScdnDomainIspDataRequest {
	s.EndTime = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataRequest) SetStartTime(v string) *TestAmpDescribeScdnDomainIspDataRequest {
	s.StartTime = &v
	return s
}

type TestAmpDescribeScdnDomainIspDataResponseBody struct {
	DataInterval *string                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                            `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *TestAmpDescribeScdnDomainIspDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s TestAmpDescribeScdnDomainIspDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestAmpDescribeScdnDomainIspDataResponseBody) GoString() string {
	return s.String()
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBody) SetDataInterval(v string) *TestAmpDescribeScdnDomainIspDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBody) SetDomainName(v string) *TestAmpDescribeScdnDomainIspDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBody) SetEndTime(v string) *TestAmpDescribeScdnDomainIspDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBody) SetRequestId(v string) *TestAmpDescribeScdnDomainIspDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBody) SetStartTime(v string) *TestAmpDescribeScdnDomainIspDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBody) SetValue(v *TestAmpDescribeScdnDomainIspDataResponseBodyValue) *TestAmpDescribeScdnDomainIspDataResponseBody {
	s.Value = v
	return s
}

type TestAmpDescribeScdnDomainIspDataResponseBodyValue struct {
	ISPProportionData []*TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData `json:"ISPProportionData,omitempty" xml:"ISPProportionData,omitempty" type:"Repeated"`
}

func (s TestAmpDescribeScdnDomainIspDataResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s TestAmpDescribeScdnDomainIspDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValue) SetISPProportionData(v []*TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) *TestAmpDescribeScdnDomainIspDataResponseBodyValue {
	s.ISPProportionData = v
	return s
}

type TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	ISP             *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IspEname        *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) String() string {
	return tea.Prettify(s)
}

func (s TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) GoString() string {
	return s.String()
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetAvgObjectSize(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetAvgResponseRate(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetAvgResponseTime(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetBps(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.Bps = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetBytesProportion(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.BytesProportion = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetISP(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.ISP = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetIspEname(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.IspEname = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetProportion(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.Proportion = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetQps(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.Qps = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetReqErrRate(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetTotalBytes(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.TotalBytes = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData) SetTotalQuery(v string) *TestAmpDescribeScdnDomainIspDataResponseBodyValueISPProportionData {
	s.TotalQuery = &v
	return s
}

type TestAmpDescribeScdnDomainIspDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TestAmpDescribeScdnDomainIspDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestAmpDescribeScdnDomainIspDataResponse) String() string {
	return tea.Prettify(s)
}

func (s TestAmpDescribeScdnDomainIspDataResponse) GoString() string {
	return s.String()
}

func (s *TestAmpDescribeScdnDomainIspDataResponse) SetHeaders(v map[string]*string) *TestAmpDescribeScdnDomainIspDataResponse {
	s.Headers = v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponse) SetStatusCode(v int32) *TestAmpDescribeScdnDomainIspDataResponse {
	s.StatusCode = &v
	return s
}

func (s *TestAmpDescribeScdnDomainIspDataResponse) SetBody(v *TestAmpDescribeScdnDomainIspDataResponseBody) *TestAmpDescribeScdnDomainIspDataResponse {
	s.Body = v
	return s
}

type UpdateScdnDomainRequest struct {
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources         *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s UpdateScdnDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScdnDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateScdnDomainRequest) SetDomainName(v string) *UpdateScdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateScdnDomainRequest) SetOwnerId(v int64) *UpdateScdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateScdnDomainRequest) SetResourceGroupId(v string) *UpdateScdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateScdnDomainRequest) SetSecurityToken(v string) *UpdateScdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateScdnDomainRequest) SetSources(v string) *UpdateScdnDomainRequest {
	s.Sources = &v
	return s
}

type UpdateScdnDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScdnDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScdnDomainResponseBody) SetRequestId(v string) *UpdateScdnDomainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateScdnDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateScdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateScdnDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScdnDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateScdnDomainResponse) SetHeaders(v map[string]*string) *UpdateScdnDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateScdnDomainResponse) SetStatusCode(v int32) *UpdateScdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScdnDomainResponse) SetBody(v *UpdateScdnDomainResponseBody) *UpdateScdnDomainResponse {
	s.Body = v
	return s
}

type VerifyScdnDomainOwnerRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyScdnDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyScdnDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyScdnDomainOwnerRequest) SetDomainName(v string) *VerifyScdnDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyScdnDomainOwnerRequest) SetVerifyType(v string) *VerifyScdnDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

type VerifyScdnDomainOwnerResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyScdnDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyScdnDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyScdnDomainOwnerResponseBody) SetContent(v string) *VerifyScdnDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyScdnDomainOwnerResponseBody) SetRequestId(v string) *VerifyScdnDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyScdnDomainOwnerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyScdnDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyScdnDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyScdnDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyScdnDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyScdnDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyScdnDomainOwnerResponse) SetStatusCode(v int32) *VerifyScdnDomainOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyScdnDomainOwnerResponse) SetBody(v *VerifyScdnDomainOwnerResponseBody) *VerifyScdnDomainOwnerResponse {
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
		"ap-northeast-1":              tea.String("scdn.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("scdn.aliyuncs.com"),
		"ap-south-1":                  tea.String("scdn.aliyuncs.com"),
		"ap-southeast-1":              tea.String("scdn.aliyuncs.com"),
		"ap-southeast-2":              tea.String("scdn.aliyuncs.com"),
		"ap-southeast-3":              tea.String("scdn.aliyuncs.com"),
		"ap-southeast-5":              tea.String("scdn.aliyuncs.com"),
		"cn-beijing":                  tea.String("scdn.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("scdn.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("scdn.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("scdn.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("scdn.aliyuncs.com"),
		"cn-chengdu":                  tea.String("scdn.aliyuncs.com"),
		"cn-edge-1":                   tea.String("scdn.aliyuncs.com"),
		"cn-fujian":                   tea.String("scdn.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("scdn.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("scdn.aliyuncs.com"),
		"cn-hongkong":                 tea.String("scdn.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("scdn.aliyuncs.com"),
		"cn-huhehaote":                tea.String("scdn.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("scdn.aliyuncs.com"),
		"cn-qingdao":                  tea.String("scdn.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("scdn.aliyuncs.com"),
		"cn-shanghai":                 tea.String("scdn.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("scdn.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("scdn.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("scdn.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("scdn.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("scdn.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("scdn.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("scdn.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("scdn.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("scdn.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("scdn.aliyuncs.com"),
		"cn-wuhan":                    tea.String("scdn.aliyuncs.com"),
		"cn-yushanfang":               tea.String("scdn.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("scdn.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("scdn.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("scdn.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("scdn.aliyuncs.com"),
		"eu-central-1":                tea.String("scdn.aliyuncs.com"),
		"eu-west-1":                   tea.String("scdn.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("scdn.aliyuncs.com"),
		"me-east-1":                   tea.String("scdn.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("scdn.aliyuncs.com"),
		"us-east-1":                   tea.String("scdn.aliyuncs.com"),
		"us-west-1":                   tea.String("scdn.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("scdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddScdnDomainWithOptions(request *AddScdnDomainRequest, runtime *util.RuntimeOptions) (_result *AddScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckUrl)) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddScdnDomain(request *AddScdnDomainRequest) (_result *AddScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddScdnDomainResponse{}
	_body, _err := client.AddScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteScdnDomainConfigsWithOptions(request *BatchDeleteScdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteScdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteScdnDomainConfigs"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteScdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteScdnDomainConfigs(request *BatchDeleteScdnDomainConfigsRequest) (_result *BatchDeleteScdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteScdnDomainConfigsResponse{}
	_body, _err := client.BatchDeleteScdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetScdnDomainConfigsWithOptions(request *BatchSetScdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchSetScdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.Functions)) {
		query["Functions"] = request.Functions
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSetScdnDomainConfigs"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSetScdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetScdnDomainConfigs(request *BatchSetScdnDomainConfigsRequest) (_result *BatchSetScdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetScdnDomainConfigsResponse{}
	_body, _err := client.BatchSetScdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartScdnDomainWithOptions(request *BatchStartScdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStartScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStartScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStartScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartScdnDomain(request *BatchStartScdnDomainRequest) (_result *BatchStartScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStartScdnDomainResponse{}
	_body, _err := client.BatchStartScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopScdnDomainWithOptions(request *BatchStopScdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStopScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStopScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStopScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopScdnDomain(request *BatchStopScdnDomainRequest) (_result *BatchStopScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopScdnDomainResponse{}
	_body, _err := client.BatchStopScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateScdnDomainWithOptions(request *BatchUpdateScdnDomainRequest, runtime *util.RuntimeOptions) (_result *BatchUpdateScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	if !tea.BoolValue(util.IsUnset(request.TopLevelDomain)) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateScdnDomain(request *BatchUpdateScdnDomainRequest) (_result *BatchUpdateScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchUpdateScdnDomainResponse{}
	_body, _err := client.BatchUpdateScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckScdnServiceWithOptions(request *CheckScdnServiceRequest, runtime *util.RuntimeOptions) (_result *CheckScdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckScdnService"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckScdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckScdnService(request *CheckScdnServiceRequest) (_result *CheckScdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckScdnServiceResponse{}
	_body, _err := client.CheckScdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScdnDomainWithOptions(request *DeleteScdnDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScdnDomain(request *DeleteScdnDomainRequest) (_result *DeleteScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScdnDomainResponse{}
	_body, _err := client.DeleteScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScdnSpecificConfigWithOptions(request *DeleteScdnSpecificConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteScdnSpecificConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScdnSpecificConfig"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScdnSpecificConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScdnSpecificConfig(request *DeleteScdnSpecificConfigRequest) (_result *DeleteScdnSpecificConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScdnSpecificConfigResponse{}
	_body, _err := client.DeleteScdnSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnCcInfoWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScdnCcInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnCcInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnCcInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnCcInfo() (_result *DescribeScdnCcInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnCcInfoResponse{}
	_body, _err := client.DescribeScdnCcInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnCcQpsInfoWithOptions(request *DescribeScdnCcQpsInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnCcQpsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnCcQpsInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnCcQpsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnCcQpsInfo(request *DescribeScdnCcQpsInfoRequest) (_result *DescribeScdnCcQpsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnCcQpsInfoResponse{}
	_body, _err := client.DescribeScdnCcQpsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnCcTopIpWithOptions(request *DescribeScdnCcTopIpRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnCcTopIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnCcTopIp"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnCcTopIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnCcTopIp(request *DescribeScdnCcTopIpRequest) (_result *DescribeScdnCcTopIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnCcTopIpResponse{}
	_body, _err := client.DescribeScdnCcTopIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnCcTopUrlWithOptions(request *DescribeScdnCcTopUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnCcTopUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnCcTopUrl"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnCcTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnCcTopUrl(request *DescribeScdnCcTopUrlRequest) (_result *DescribeScdnCcTopUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnCcTopUrlResponse{}
	_body, _err := client.DescribeScdnCcTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnCertificateDetailWithOptions(request *DescribeScdnCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnCertificateDetail"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnCertificateDetail(request *DescribeScdnCertificateDetailRequest) (_result *DescribeScdnCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnCertificateDetailResponse{}
	_body, _err := client.DescribeScdnCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnCertificateListWithOptions(request *DescribeScdnCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnCertificateList"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnCertificateList(request *DescribeScdnCertificateListRequest) (_result *DescribeScdnCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnCertificateListResponse{}
	_body, _err := client.DescribeScdnCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDDoSInfoWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScdnDDoSInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDDoSInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDDoSInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDDoSInfo() (_result *DescribeScdnDDoSInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDDoSInfoResponse{}
	_body, _err := client.DescribeScdnDDoSInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDDoSTrafficInfoWithOptions(request *DescribeScdnDDoSTrafficInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDDoSTrafficInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDDoSTrafficInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDDoSTrafficInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDDoSTrafficInfo(request *DescribeScdnDDoSTrafficInfoRequest) (_result *DescribeScdnDDoSTrafficInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDDoSTrafficInfoResponse{}
	_body, _err := client.DescribeScdnDDoSTrafficInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainBpsDataWithOptions(request *DescribeScdnDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainBpsData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainBpsData(request *DescribeScdnDomainBpsDataRequest) (_result *DescribeScdnDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainBpsDataResponse{}
	_body, _err := client.DescribeScdnDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainCertificateInfoWithOptions(request *DescribeScdnDomainCertificateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainCertificateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainCertificateInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainCertificateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainCertificateInfo(request *DescribeScdnDomainCertificateInfoRequest) (_result *DescribeScdnDomainCertificateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainCertificateInfoResponse{}
	_body, _err := client.DescribeScdnDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainCnameWithOptions(request *DescribeScdnDomainCnameRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainCnameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainCname"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainCname(request *DescribeScdnDomainCnameRequest) (_result *DescribeScdnDomainCnameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainCnameResponse{}
	_body, _err := client.DescribeScdnDomainCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainConfigsWithOptions(request *DescribeScdnDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNames)) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainConfigs"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainConfigs(request *DescribeScdnDomainConfigsRequest) (_result *DescribeScdnDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainConfigsResponse{}
	_body, _err := client.DescribeScdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainDetailWithOptions(request *DescribeScdnDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainDetail"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainDetail(request *DescribeScdnDomainDetailRequest) (_result *DescribeScdnDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainDetailResponse{}
	_body, _err := client.DescribeScdnDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainHitRateDataWithOptions(request *DescribeScdnDomainHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainHitRateData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainHitRateData(request *DescribeScdnDomainHitRateDataRequest) (_result *DescribeScdnDomainHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainHitRateDataResponse{}
	_body, _err := client.DescribeScdnDomainHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainHttpCodeDataWithOptions(request *DescribeScdnDomainHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainHttpCodeData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainHttpCodeData(request *DescribeScdnDomainHttpCodeDataRequest) (_result *DescribeScdnDomainHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainHttpCodeDataResponse{}
	_body, _err := client.DescribeScdnDomainHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainIspDataWithOptions(request *DescribeScdnDomainIspDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainIspDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainIspData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainIspDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainIspData(request *DescribeScdnDomainIspDataRequest) (_result *DescribeScdnDomainIspDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainIspDataResponse{}
	_body, _err := client.DescribeScdnDomainIspDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainLogWithOptions(request *DescribeScdnDomainLogRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainLog"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainLog(request *DescribeScdnDomainLogRequest) (_result *DescribeScdnDomainLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainLogResponse{}
	_body, _err := client.DescribeScdnDomainLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainOriginBpsDataWithOptions(request *DescribeScdnDomainOriginBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainOriginBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainOriginBpsData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainOriginBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainOriginBpsData(request *DescribeScdnDomainOriginBpsDataRequest) (_result *DescribeScdnDomainOriginBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainOriginBpsDataResponse{}
	_body, _err := client.DescribeScdnDomainOriginBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainOriginTrafficDataWithOptions(request *DescribeScdnDomainOriginTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainOriginTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainOriginTrafficData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainOriginTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainOriginTrafficData(request *DescribeScdnDomainOriginTrafficDataRequest) (_result *DescribeScdnDomainOriginTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainOriginTrafficDataResponse{}
	_body, _err := client.DescribeScdnDomainOriginTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainPvDataWithOptions(request *DescribeScdnDomainPvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainPvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainPvData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainPvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainPvData(request *DescribeScdnDomainPvDataRequest) (_result *DescribeScdnDomainPvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainPvDataResponse{}
	_body, _err := client.DescribeScdnDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainQpsDataWithOptions(request *DescribeScdnDomainQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainQpsData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainQpsData(request *DescribeScdnDomainQpsDataRequest) (_result *DescribeScdnDomainQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainQpsDataResponse{}
	_body, _err := client.DescribeScdnDomainQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeBpsDataWithOptions(request *DescribeScdnDomainRealTimeBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeBpsData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeBpsData(request *DescribeScdnDomainRealTimeBpsDataRequest) (_result *DescribeScdnDomainRealTimeBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeByteHitRateDataWithOptions(request *DescribeScdnDomainRealTimeByteHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeByteHitRateData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeByteHitRateData(request *DescribeScdnDomainRealTimeByteHitRateDataRequest) (_result *DescribeScdnDomainRealTimeByteHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeByteHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeHttpCodeDataWithOptions(request *DescribeScdnDomainRealTimeHttpCodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeHttpCodeData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeHttpCodeData(request *DescribeScdnDomainRealTimeHttpCodeDataRequest) (_result *DescribeScdnDomainRealTimeHttpCodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeQpsDataWithOptions(request *DescribeScdnDomainRealTimeQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeQpsData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeQpsData(request *DescribeScdnDomainRealTimeQpsDataRequest) (_result *DescribeScdnDomainRealTimeQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeReqHitRateDataWithOptions(request *DescribeScdnDomainRealTimeReqHitRateDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeReqHitRateData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeReqHitRateData(request *DescribeScdnDomainRealTimeReqHitRateDataRequest) (_result *DescribeScdnDomainRealTimeReqHitRateDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeSrcBpsDataWithOptions(request *DescribeScdnDomainRealTimeSrcBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeSrcBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeSrcBpsData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeSrcBpsData(request *DescribeScdnDomainRealTimeSrcBpsDataRequest) (_result *DescribeScdnDomainRealTimeSrcBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeSrcTrafficDataWithOptions(request *DescribeScdnDomainRealTimeSrcTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeSrcTrafficData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeSrcTrafficData(request *DescribeScdnDomainRealTimeSrcTrafficDataRequest) (_result *DescribeScdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeTrafficDataWithOptions(request *DescribeScdnDomainRealTimeTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRealTimeTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRealTimeTrafficData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRealTimeTrafficData(request *DescribeScdnDomainRealTimeTrafficDataRequest) (_result *DescribeScdnDomainRealTimeTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DescribeScdnDomainRealTimeTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainRegionDataWithOptions(request *DescribeScdnDomainRegionDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainRegionDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainRegionData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainRegionData(request *DescribeScdnDomainRegionDataRequest) (_result *DescribeScdnDomainRegionDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainRegionDataResponse{}
	_body, _err := client.DescribeScdnDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainTopReferVisitWithOptions(request *DescribeScdnDomainTopReferVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainTopReferVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainTopReferVisit"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainTopReferVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainTopReferVisit(request *DescribeScdnDomainTopReferVisitRequest) (_result *DescribeScdnDomainTopReferVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainTopReferVisitResponse{}
	_body, _err := client.DescribeScdnDomainTopReferVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainTopUrlVisitWithOptions(request *DescribeScdnDomainTopUrlVisitRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainTopUrlVisitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainTopUrlVisit"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainTopUrlVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainTopUrlVisit(request *DescribeScdnDomainTopUrlVisitRequest) (_result *DescribeScdnDomainTopUrlVisitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainTopUrlVisitResponse{}
	_body, _err := client.DescribeScdnDomainTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainTrafficDataWithOptions(request *DescribeScdnDomainTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.IspNameEn)) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.LocationNameEn)) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainTrafficData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainTrafficData(request *DescribeScdnDomainTrafficDataRequest) (_result *DescribeScdnDomainTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainTrafficDataResponse{}
	_body, _err := client.DescribeScdnDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnDomainUvDataWithOptions(request *DescribeScdnDomainUvDataRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnDomainUvDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnDomainUvData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnDomainUvData(request *DescribeScdnDomainUvDataRequest) (_result *DescribeScdnDomainUvDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnDomainUvDataResponse{}
	_body, _err := client.DescribeScdnDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnRefreshQuotaWithOptions(request *DescribeScdnRefreshQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnRefreshQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnRefreshQuota"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnRefreshQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnRefreshQuota(request *DescribeScdnRefreshQuotaRequest) (_result *DescribeScdnRefreshQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnRefreshQuotaResponse{}
	_body, _err := client.DescribeScdnRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnRefreshTasksWithOptions(request *DescribeScdnRefreshTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnRefreshTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnRefreshTasks"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnRefreshTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnRefreshTasks(request *DescribeScdnRefreshTasksRequest) (_result *DescribeScdnRefreshTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnRefreshTasksResponse{}
	_body, _err := client.DescribeScdnRefreshTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnServiceWithOptions(request *DescribeScdnServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnService"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnService(request *DescribeScdnServiceRequest) (_result *DescribeScdnServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnServiceResponse{}
	_body, _err := client.DescribeScdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnTopDomainsByFlowWithOptions(request *DescribeScdnTopDomainsByFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnTopDomainsByFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnTopDomainsByFlow"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnTopDomainsByFlow(request *DescribeScdnTopDomainsByFlowRequest) (_result *DescribeScdnTopDomainsByFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnTopDomainsByFlowResponse{}
	_body, _err := client.DescribeScdnTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnUserDomainsWithOptions(request *DescribeScdnUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeEndTime)) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeStartTime)) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CheckDomainShow)) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainSearchType)) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainStatus)) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnUserDomains"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnUserDomains(request *DescribeScdnUserDomainsRequest) (_result *DescribeScdnUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnUserDomainsResponse{}
	_body, _err := client.DescribeScdnUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnUserProtectInfoWithOptions(runtime *util.RuntimeOptions) (_result *DescribeScdnUserProtectInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnUserProtectInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnUserProtectInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnUserProtectInfo() (_result *DescribeScdnUserProtectInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnUserProtectInfoResponse{}
	_body, _err := client.DescribeScdnUserProtectInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnUserQuotaWithOptions(request *DescribeScdnUserQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnUserQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnUserQuota"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnUserQuota(request *DescribeScdnUserQuotaRequest) (_result *DescribeScdnUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnUserQuotaResponse{}
	_body, _err := client.DescribeScdnUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScdnVerifyContentWithOptions(request *DescribeScdnVerifyContentRequest, runtime *util.RuntimeOptions) (_result *DescribeScdnVerifyContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScdnVerifyContent"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScdnVerifyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScdnVerifyContent(request *DescribeScdnVerifyContentRequest) (_result *DescribeScdnVerifyContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScdnVerifyContentResponse{}
	_body, _err := client.DescribeScdnVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreloadScdnObjectCachesWithOptions(request *PreloadScdnObjectCachesRequest, runtime *util.RuntimeOptions) (_result *PreloadScdnObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Area)) {
		query["Area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.L2Preload)) {
		query["L2Preload"] = request.L2Preload
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreloadScdnObjectCaches"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreloadScdnObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreloadScdnObjectCaches(request *PreloadScdnObjectCachesRequest) (_result *PreloadScdnObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreloadScdnObjectCachesResponse{}
	_body, _err := client.PreloadScdnObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshScdnObjectCachesWithOptions(request *RefreshScdnObjectCachesRequest, runtime *util.RuntimeOptions) (_result *RefreshScdnObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectPath)) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshScdnObjectCaches"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshScdnObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshScdnObjectCaches(request *RefreshScdnObjectCachesRequest) (_result *RefreshScdnObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshScdnObjectCachesResponse{}
	_body, _err := client.RefreshScdnObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetScdnBotInfoWithOptions(request *SetScdnBotInfoRequest, runtime *util.RuntimeOptions) (_result *SetScdnBotInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetScdnBotInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetScdnBotInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetScdnBotInfo(request *SetScdnBotInfoRequest) (_result *SetScdnBotInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetScdnBotInfoResponse{}
	_body, _err := client.SetScdnBotInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetScdnCcInfoWithOptions(request *SetScdnCcInfoRequest, runtime *util.RuntimeOptions) (_result *SetScdnCcInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetScdnCcInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetScdnCcInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetScdnCcInfo(request *SetScdnCcInfoRequest) (_result *SetScdnCcInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetScdnCcInfoResponse{}
	_body, _err := client.SetScdnCcInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetScdnDDoSInfoWithOptions(request *SetScdnDDoSInfoRequest, runtime *util.RuntimeOptions) (_result *SetScdnDDoSInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetScdnDDoSInfo"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetScdnDDoSInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetScdnDDoSInfo(request *SetScdnDDoSInfoRequest) (_result *SetScdnDDoSInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetScdnDDoSInfoResponse{}
	_body, _err := client.SetScdnDDoSInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetScdnDomainCertificateWithOptions(request *SetScdnDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *SetScdnDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertType)) {
		query["CertType"] = request.CertType
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ForceSet)) {
		query["ForceSet"] = request.ForceSet
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPri)) {
		query["SSLPri"] = request.SSLPri
	}

	if !tea.BoolValue(util.IsUnset(request.SSLProtocol)) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SSLPub)) {
		query["SSLPub"] = request.SSLPub
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetScdnDomainCertificate"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetScdnDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetScdnDomainCertificate(request *SetScdnDomainCertificateRequest) (_result *SetScdnDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetScdnDomainCertificateResponse{}
	_body, _err := client.SetScdnDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartScdnDomainWithOptions(request *StartScdnDomainRequest, runtime *util.RuntimeOptions) (_result *StartScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartScdnDomain(request *StartScdnDomainRequest) (_result *StartScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartScdnDomainResponse{}
	_body, _err := client.StartScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopScdnDomainWithOptions(request *StopScdnDomainRequest, runtime *util.RuntimeOptions) (_result *StopScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopScdnDomain(request *StopScdnDomainRequest) (_result *StopScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopScdnDomainResponse{}
	_body, _err := client.StopScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestAmpDescribeScdnDomainIspDataWithOptions(request *TestAmpDescribeScdnDomainIspDataRequest, runtime *util.RuntimeOptions) (_result *TestAmpDescribeScdnDomainIspDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TestAmpDescribeScdnDomainIspData"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestAmpDescribeScdnDomainIspDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestAmpDescribeScdnDomainIspData(request *TestAmpDescribeScdnDomainIspDataRequest) (_result *TestAmpDescribeScdnDomainIspDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestAmpDescribeScdnDomainIspDataResponse{}
	_body, _err := client.TestAmpDescribeScdnDomainIspDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateScdnDomainWithOptions(request *UpdateScdnDomainRequest, runtime *util.RuntimeOptions) (_result *UpdateScdnDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Sources)) {
		query["Sources"] = request.Sources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScdnDomain"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateScdnDomain(request *UpdateScdnDomainRequest) (_result *UpdateScdnDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScdnDomainResponse{}
	_body, _err := client.UpdateScdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyScdnDomainOwnerWithOptions(request *VerifyScdnDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyScdnDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyScdnDomainOwner"),
		Version:     tea.String("2017-11-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyScdnDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyScdnDomainOwner(request *VerifyScdnDomainOwnerRequest) (_result *VerifyScdnDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyScdnDomainOwnerResponse{}
	_body, _err := client.VerifyScdnDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
