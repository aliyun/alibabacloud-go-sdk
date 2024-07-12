// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AuditPublicTemplateRegistrationRequest struct {
	AuditAction    *string `json:"AuditAction,omitempty" xml:"AuditAction,omitempty"`
	Comment        *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
}

func (s AuditPublicTemplateRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditPublicTemplateRegistrationRequest) GoString() string {
	return s.String()
}

func (s *AuditPublicTemplateRegistrationRequest) SetAuditAction(v string) *AuditPublicTemplateRegistrationRequest {
	s.AuditAction = &v
	return s
}

func (s *AuditPublicTemplateRegistrationRequest) SetComment(v string) *AuditPublicTemplateRegistrationRequest {
	s.Comment = &v
	return s
}

func (s *AuditPublicTemplateRegistrationRequest) SetRegionId(v string) *AuditPublicTemplateRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *AuditPublicTemplateRegistrationRequest) SetRegistrationId(v string) *AuditPublicTemplateRegistrationRequest {
	s.RegistrationId = &v
	return s
}

type AuditPublicTemplateRegistrationResponseBody struct {
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Detail          *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	RegistrationId  *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s AuditPublicTemplateRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuditPublicTemplateRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetComment(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.Comment = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetDetail(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.Detail = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetRegistrationId(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetRequestId(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetStatus(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.Status = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetTemplateId(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.TemplateId = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetTemplateName(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.TemplateName = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponseBody) SetTemplateVersion(v string) *AuditPublicTemplateRegistrationResponseBody {
	s.TemplateVersion = &v
	return s
}

type AuditPublicTemplateRegistrationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditPublicTemplateRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditPublicTemplateRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditPublicTemplateRegistrationResponse) GoString() string {
	return s.String()
}

func (s *AuditPublicTemplateRegistrationResponse) SetHeaders(v map[string]*string) *AuditPublicTemplateRegistrationResponse {
	s.Headers = v
	return s
}

func (s *AuditPublicTemplateRegistrationResponse) SetStatusCode(v int32) *AuditPublicTemplateRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditPublicTemplateRegistrationResponse) SetBody(v *AuditPublicTemplateRegistrationResponseBody) *AuditPublicTemplateRegistrationResponse {
	s.Body = v
	return s
}

type CreateActionRequest struct {
	// This parameter is required.
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// This parameter is required.
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// This parameter is required.
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Popularity *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequest) GoString() string {
	return s.String()
}

func (s *CreateActionRequest) SetActionName(v string) *CreateActionRequest {
	s.ActionName = &v
	return s
}

func (s *CreateActionRequest) SetActionType(v string) *CreateActionRequest {
	s.ActionType = &v
	return s
}

func (s *CreateActionRequest) SetContent(v string) *CreateActionRequest {
	s.Content = &v
	return s
}

func (s *CreateActionRequest) SetPopularity(v int32) *CreateActionRequest {
	s.Popularity = &v
	return s
}

func (s *CreateActionRequest) SetRegionId(v string) *CreateActionRequest {
	s.RegionId = &v
	return s
}

type CreateActionResponseBody struct {
	ActionName      *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ActionType      *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActionResponseBody) SetActionName(v string) *CreateActionResponseBody {
	s.ActionName = &v
	return s
}

func (s *CreateActionResponseBody) SetActionType(v string) *CreateActionResponseBody {
	s.ActionType = &v
	return s
}

func (s *CreateActionResponseBody) SetCreatedDate(v string) *CreateActionResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *CreateActionResponseBody) SetDescription(v string) *CreateActionResponseBody {
	s.Description = &v
	return s
}

func (s *CreateActionResponseBody) SetPopularity(v int32) *CreateActionResponseBody {
	s.Popularity = &v
	return s
}

func (s *CreateActionResponseBody) SetProperties(v string) *CreateActionResponseBody {
	s.Properties = &v
	return s
}

func (s *CreateActionResponseBody) SetRequestId(v string) *CreateActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateActionResponseBody) SetTemplateVersion(v string) *CreateActionResponseBody {
	s.TemplateVersion = &v
	return s
}

type CreateActionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateActionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponse) GoString() string {
	return s.String()
}

func (s *CreateActionResponse) SetHeaders(v map[string]*string) *CreateActionResponse {
	s.Headers = v
	return s
}

func (s *CreateActionResponse) SetStatusCode(v int32) *CreateActionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateActionResponse) SetBody(v *CreateActionResponseBody) *CreateActionResponse {
	s.Body = v
	return s
}

type CreatePublicParameterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ParameterType *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePublicParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicParameterRequest) GoString() string {
	return s.String()
}

func (s *CreatePublicParameterRequest) SetClientToken(v string) *CreatePublicParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePublicParameterRequest) SetConstraints(v string) *CreatePublicParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreatePublicParameterRequest) SetDescription(v string) *CreatePublicParameterRequest {
	s.Description = &v
	return s
}

func (s *CreatePublicParameterRequest) SetName(v string) *CreatePublicParameterRequest {
	s.Name = &v
	return s
}

func (s *CreatePublicParameterRequest) SetParameterType(v string) *CreatePublicParameterRequest {
	s.ParameterType = &v
	return s
}

func (s *CreatePublicParameterRequest) SetRegionId(v string) *CreatePublicParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePublicParameterRequest) SetValue(v string) *CreatePublicParameterRequest {
	s.Value = &v
	return s
}

type CreatePublicParameterResponseBody struct {
	Parameter *CreatePublicParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePublicParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublicParameterResponseBody) SetParameter(v *CreatePublicParameterResponseBodyParameter) *CreatePublicParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *CreatePublicParameterResponseBody) SetRequestId(v string) *CreatePublicParameterResponseBody {
	s.RequestId = &v
	return s
}

type CreatePublicParameterResponseBodyParameter struct {
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreatePublicParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreatePublicParameterResponseBodyParameter) SetConstraints(v string) *CreatePublicParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetCreatedBy(v string) *CreatePublicParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetCreatedDate(v string) *CreatePublicParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetDescription(v string) *CreatePublicParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetId(v string) *CreatePublicParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetName(v string) *CreatePublicParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetParameterVersion(v int32) *CreatePublicParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetRegionId(v string) *CreatePublicParameterResponseBodyParameter {
	s.RegionId = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetShareType(v string) *CreatePublicParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetType(v string) *CreatePublicParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetUpdatedBy(v string) *CreatePublicParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreatePublicParameterResponseBodyParameter) SetUpdatedDate(v string) *CreatePublicParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

type CreatePublicParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublicParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublicParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicParameterResponse) GoString() string {
	return s.String()
}

func (s *CreatePublicParameterResponse) SetHeaders(v map[string]*string) *CreatePublicParameterResponse {
	s.Headers = v
	return s
}

func (s *CreatePublicParameterResponse) SetStatusCode(v int32) *CreatePublicParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublicParameterResponse) SetBody(v *CreatePublicParameterResponseBody) *CreatePublicParameterResponse {
	s.Body = v
	return s
}

type CreatePublicPatchBaselineRequest struct {
	// This parameter is required.
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreatePublicPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreatePublicPatchBaselineRequest) SetApprovalRules(v string) *CreatePublicPatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePublicPatchBaselineRequest) SetClientToken(v string) *CreatePublicPatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePublicPatchBaselineRequest) SetDescription(v string) *CreatePublicPatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *CreatePublicPatchBaselineRequest) SetName(v string) *CreatePublicPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *CreatePublicPatchBaselineRequest) SetOperationSystem(v string) *CreatePublicPatchBaselineRequest {
	s.OperationSystem = &v
	return s
}

func (s *CreatePublicPatchBaselineRequest) SetRegionId(v string) *CreatePublicPatchBaselineRequest {
	s.RegionId = &v
	return s
}

type CreatePublicPatchBaselineResponseBody struct {
	PatchBaseline *CreatePublicPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePublicPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublicPatchBaselineResponseBody) SetPatchBaseline(v *CreatePublicPatchBaselineResponseBodyPatchBaseline) *CreatePublicPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *CreatePublicPatchBaselineResponseBody) SetRequestId(v string) *CreatePublicPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type CreatePublicPatchBaselineResponseBodyPatchBaseline struct {
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreatePublicPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetId(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetName(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *CreatePublicPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *CreatePublicPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type CreatePublicPatchBaselineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublicPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublicPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreatePublicPatchBaselineResponse) SetHeaders(v map[string]*string) *CreatePublicPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreatePublicPatchBaselineResponse) SetStatusCode(v int32) *CreatePublicPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublicPatchBaselineResponse) SetBody(v *CreatePublicPatchBaselineResponseBody) *CreatePublicPatchBaselineResponse {
	s.Body = v
	return s
}

type CreatePublicTemplateRequest struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	IsExample  *bool   `json:"IsExample,omitempty" xml:"IsExample,omitempty"`
	Popularity *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	Publisher  *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreatePublicTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreatePublicTemplateRequest) SetCategory(v string) *CreatePublicTemplateRequest {
	s.Category = &v
	return s
}

func (s *CreatePublicTemplateRequest) SetContent(v string) *CreatePublicTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreatePublicTemplateRequest) SetIsExample(v bool) *CreatePublicTemplateRequest {
	s.IsExample = &v
	return s
}

func (s *CreatePublicTemplateRequest) SetPopularity(v int32) *CreatePublicTemplateRequest {
	s.Popularity = &v
	return s
}

func (s *CreatePublicTemplateRequest) SetPublisher(v string) *CreatePublicTemplateRequest {
	s.Publisher = &v
	return s
}

func (s *CreatePublicTemplateRequest) SetRegionId(v string) *CreatePublicTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePublicTemplateRequest) SetTemplateName(v string) *CreatePublicTemplateRequest {
	s.TemplateName = &v
	return s
}

type CreatePublicTemplateResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *CreatePublicTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s CreatePublicTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublicTemplateResponseBody) SetRequestId(v string) *CreatePublicTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublicTemplateResponseBody) SetTemplate(v *CreatePublicTemplateResponseBodyTemplate) *CreatePublicTemplateResponseBody {
	s.Template = v
	return s
}

type CreatePublicTemplateResponseBodyTemplate struct {
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash            *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat  *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreatePublicTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetCategory(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.Category = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetCreatedBy(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetCreatedDate(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetDescription(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetHash(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetPopularity(v int32) *CreatePublicTemplateResponseBodyTemplate {
	s.Popularity = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetShareType(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetTemplateFormat(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetTemplateId(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetTemplateName(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetTemplateVersion(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetUpdatedBy(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *CreatePublicTemplateResponseBodyTemplate) SetUpdatedDate(v string) *CreatePublicTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type CreatePublicTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublicTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublicTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePublicTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreatePublicTemplateResponse) SetHeaders(v map[string]*string) *CreatePublicTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreatePublicTemplateResponse) SetStatusCode(v int32) *CreatePublicTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublicTemplateResponse) SetBody(v *CreatePublicTemplateResponseBody) *CreatePublicTemplateResponse {
	s.Body = v
	return s
}

type DeleteFailureMsgRequest struct {
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	RequestFingerprint *string `json:"RequestFingerprint,omitempty" xml:"RequestFingerprint,omitempty"`
}

func (s DeleteFailureMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFailureMsgRequest) GoString() string {
	return s.String()
}

func (s *DeleteFailureMsgRequest) SetOperation(v string) *DeleteFailureMsgRequest {
	s.Operation = &v
	return s
}

func (s *DeleteFailureMsgRequest) SetRequestFingerprint(v string) *DeleteFailureMsgRequest {
	s.RequestFingerprint = &v
	return s
}

type DeleteFailureMsgResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFailureMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFailureMsgResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFailureMsgResponseBody) SetRequestId(v string) *DeleteFailureMsgResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFailureMsgResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFailureMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFailureMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFailureMsgResponse) GoString() string {
	return s.String()
}

func (s *DeleteFailureMsgResponse) SetHeaders(v map[string]*string) *DeleteFailureMsgResponse {
	s.Headers = v
	return s
}

func (s *DeleteFailureMsgResponse) SetStatusCode(v int32) *DeleteFailureMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFailureMsgResponse) SetBody(v *DeleteFailureMsgResponseBody) *DeleteFailureMsgResponse {
	s.Body = v
	return s
}

type DeletePublicParameterRequest struct {
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePublicParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicParameterRequest) GoString() string {
	return s.String()
}

func (s *DeletePublicParameterRequest) SetName(v string) *DeletePublicParameterRequest {
	s.Name = &v
	return s
}

func (s *DeletePublicParameterRequest) SetRegionId(v string) *DeletePublicParameterRequest {
	s.RegionId = &v
	return s
}

type DeletePublicParameterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePublicParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicParameterResponseBody) SetRequestId(v string) *DeletePublicParameterResponseBody {
	s.RequestId = &v
	return s
}

type DeletePublicParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicParameterResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicParameterResponse) SetHeaders(v map[string]*string) *DeletePublicParameterResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicParameterResponse) SetStatusCode(v int32) *DeletePublicParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicParameterResponse) SetBody(v *DeletePublicParameterResponseBody) *DeletePublicParameterResponse {
	s.Body = v
	return s
}

type DeletePublicPatchBaselineRequest struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeletePublicPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeletePublicPatchBaselineRequest) SetName(v string) *DeletePublicPatchBaselineRequest {
	s.Name = &v
	return s
}

type DeletePublicPatchBaselineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePublicPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicPatchBaselineResponseBody) SetRequestId(v string) *DeletePublicPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type DeletePublicPatchBaselineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicPatchBaselineResponse) SetHeaders(v map[string]*string) *DeletePublicPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicPatchBaselineResponse) SetStatusCode(v int32) *DeletePublicPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicPatchBaselineResponse) SetBody(v *DeletePublicPatchBaselineResponseBody) *DeletePublicPatchBaselineResponse {
	s.Body = v
	return s
}

type DeletePublicTemplateRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DeletePublicTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeletePublicTemplateRequest) SetRegionId(v string) *DeletePublicTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePublicTemplateRequest) SetTemplateName(v string) *DeletePublicTemplateRequest {
	s.TemplateName = &v
	return s
}

type DeletePublicTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePublicTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePublicTemplateResponseBody) SetRequestId(v string) *DeletePublicTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeletePublicTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePublicTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicTemplateResponse) SetHeaders(v map[string]*string) *DeletePublicTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicTemplateResponse) SetStatusCode(v int32) *DeletePublicTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicTemplateResponse) SetBody(v *DeletePublicTemplateResponseBody) *DeletePublicTemplateResponse {
	s.Body = v
	return s
}

type DoCheckResourceRequest struct {
	Bid            *string `json:"bid,omitempty" xml:"bid,omitempty"`
	Country        *string `json:"country,omitempty" xml:"country,omitempty"`
	GmtWakeup      *string `json:"gmtWakeup,omitempty" xml:"gmtWakeup,omitempty"`
	Hid            *int32  `json:"hid,omitempty" xml:"hid,omitempty"`
	Interrupt      *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	Invoker        *string `json:"invoker,omitempty" xml:"invoker,omitempty"`
	Level          *int32  `json:"level,omitempty" xml:"level,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	Pk             *string `json:"pk,omitempty" xml:"pk,omitempty"`
	Prompt         *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TaskExtraData  *string `json:"taskExtraData,omitempty" xml:"taskExtraData,omitempty"`
	TaskIdentifier *string `json:"taskIdentifier,omitempty" xml:"taskIdentifier,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s DoCheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceRequest) GoString() string {
	return s.String()
}

func (s *DoCheckResourceRequest) SetBid(v string) *DoCheckResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoCheckResourceRequest) SetCountry(v string) *DoCheckResourceRequest {
	s.Country = &v
	return s
}

func (s *DoCheckResourceRequest) SetGmtWakeup(v string) *DoCheckResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoCheckResourceRequest) SetHid(v int32) *DoCheckResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoCheckResourceRequest) SetInterrupt(v bool) *DoCheckResourceRequest {
	s.Interrupt = &v
	return s
}

func (s *DoCheckResourceRequest) SetInvoker(v string) *DoCheckResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoCheckResourceRequest) SetLevel(v int32) *DoCheckResourceRequest {
	s.Level = &v
	return s
}

func (s *DoCheckResourceRequest) SetMessage(v string) *DoCheckResourceRequest {
	s.Message = &v
	return s
}

func (s *DoCheckResourceRequest) SetPk(v string) *DoCheckResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoCheckResourceRequest) SetPrompt(v string) *DoCheckResourceRequest {
	s.Prompt = &v
	return s
}

func (s *DoCheckResourceRequest) SetSuccess(v bool) *DoCheckResourceRequest {
	s.Success = &v
	return s
}

func (s *DoCheckResourceRequest) SetTaskExtraData(v string) *DoCheckResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoCheckResourceRequest) SetTaskIdentifier(v string) *DoCheckResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoCheckResourceRequest) SetUrl(v string) *DoCheckResourceRequest {
	s.Url = &v
	return s
}

type DoCheckResourceResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Bid            *string `json:"bid,omitempty" xml:"bid,omitempty"`
	Country        *string `json:"country,omitempty" xml:"country,omitempty"`
	GmtWakeup      *string `json:"gmtWakeup,omitempty" xml:"gmtWakeup,omitempty"`
	Hid            *int32  `json:"hid,omitempty" xml:"hid,omitempty"`
	Interrupt      *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	Invoker        *string `json:"invoker,omitempty" xml:"invoker,omitempty"`
	Level          *int32  `json:"level,omitempty" xml:"level,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	Pk             *string `json:"pk,omitempty" xml:"pk,omitempty"`
	Prompt         *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TaskExtraData  *string `json:"taskExtraData,omitempty" xml:"taskExtraData,omitempty"`
	TaskIdentifier *string `json:"taskIdentifier,omitempty" xml:"taskIdentifier,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s DoCheckResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DoCheckResourceResponseBody) SetRequestId(v string) *DoCheckResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetBid(v string) *DoCheckResourceResponseBody {
	s.Bid = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetCountry(v string) *DoCheckResourceResponseBody {
	s.Country = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetGmtWakeup(v string) *DoCheckResourceResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetHid(v int32) *DoCheckResourceResponseBody {
	s.Hid = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetInterrupt(v bool) *DoCheckResourceResponseBody {
	s.Interrupt = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetInvoker(v string) *DoCheckResourceResponseBody {
	s.Invoker = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetLevel(v int32) *DoCheckResourceResponseBody {
	s.Level = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetMessage(v string) *DoCheckResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetPk(v string) *DoCheckResourceResponseBody {
	s.Pk = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetPrompt(v string) *DoCheckResourceResponseBody {
	s.Prompt = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetSuccess(v bool) *DoCheckResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetTaskExtraData(v string) *DoCheckResourceResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetTaskIdentifier(v string) *DoCheckResourceResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetUrl(v string) *DoCheckResourceResponseBody {
	s.Url = &v
	return s
}

type DoCheckResourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DoCheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DoCheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceResponse) GoString() string {
	return s.String()
}

func (s *DoCheckResourceResponse) SetHeaders(v map[string]*string) *DoCheckResourceResponse {
	s.Headers = v
	return s
}

func (s *DoCheckResourceResponse) SetStatusCode(v int32) *DoCheckResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DoCheckResourceResponse) SetBody(v *DoCheckResourceResponseBody) *DoCheckResourceResponse {
	s.Body = v
	return s
}

type GetActionRequest struct {
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetActionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActionRequest) GoString() string {
	return s.String()
}

func (s *GetActionRequest) SetActionName(v string) *GetActionRequest {
	s.ActionName = &v
	return s
}

func (s *GetActionRequest) SetRegionId(v string) *GetActionRequest {
	s.RegionId = &v
	return s
}

type GetActionResponseBody struct {
	ActionName   *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ActionType   *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	Content      []byte  `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Popularity   *string `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActionResponseBody) GoString() string {
	return s.String()
}

func (s *GetActionResponseBody) SetActionName(v string) *GetActionResponseBody {
	s.ActionName = &v
	return s
}

func (s *GetActionResponseBody) SetActionType(v string) *GetActionResponseBody {
	s.ActionType = &v
	return s
}

func (s *GetActionResponseBody) SetContent(v []byte) *GetActionResponseBody {
	s.Content = v
	return s
}

func (s *GetActionResponseBody) SetCreateTime(v string) *GetActionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetActionResponseBody) SetModifiedTime(v string) *GetActionResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetActionResponseBody) SetPopularity(v string) *GetActionResponseBody {
	s.Popularity = &v
	return s
}

func (s *GetActionResponseBody) SetRequestId(v string) *GetActionResponseBody {
	s.RequestId = &v
	return s
}

type GetActionResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetActionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActionResponse) GoString() string {
	return s.String()
}

func (s *GetActionResponse) SetHeaders(v map[string]*string) *GetActionResponse {
	s.Headers = v
	return s
}

func (s *GetActionResponse) SetStatusCode(v int32) *GetActionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActionResponse) SetBody(v *GetActionResponseBody) *GetActionResponse {
	s.Body = v
	return s
}

type GetFlowControlRequest struct {
	Api     *string `json:"Api,omitempty" xml:"Api,omitempty"`
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// This parameter is required.
	Type *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid  *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowControlRequest) GoString() string {
	return s.String()
}

func (s *GetFlowControlRequest) SetApi(v string) *GetFlowControlRequest {
	s.Api = &v
	return s
}

func (s *GetFlowControlRequest) SetService(v string) *GetFlowControlRequest {
	s.Service = &v
	return s
}

func (s *GetFlowControlRequest) SetType(v int32) *GetFlowControlRequest {
	s.Type = &v
	return s
}

func (s *GetFlowControlRequest) SetUid(v string) *GetFlowControlRequest {
	s.Uid = &v
	return s
}

type GetFlowControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Value     *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowControlResponseBody) SetRequestId(v string) *GetFlowControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowControlResponseBody) SetValue(v int32) *GetFlowControlResponseBody {
	s.Value = &v
	return s
}

type GetFlowControlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowControlResponse) GoString() string {
	return s.String()
}

func (s *GetFlowControlResponse) SetHeaders(v map[string]*string) *GetFlowControlResponse {
	s.Headers = v
	return s
}

func (s *GetFlowControlResponse) SetStatusCode(v int32) *GetFlowControlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowControlResponse) SetBody(v *GetFlowControlResponseBody) *GetFlowControlResponse {
	s.Body = v
	return s
}

type GetPublicParameterRequest struct {
	// This parameter is required.
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPublicParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicParameterRequest) GoString() string {
	return s.String()
}

func (s *GetPublicParameterRequest) SetName(v string) *GetPublicParameterRequest {
	s.Name = &v
	return s
}

func (s *GetPublicParameterRequest) SetParameterVersion(v int32) *GetPublicParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetPublicParameterRequest) SetRegionId(v string) *GetPublicParameterRequest {
	s.RegionId = &v
	return s
}

type GetPublicParameterResponseBody struct {
	Parameter *GetPublicParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicParameterResponseBody) SetParameter(v *GetPublicParameterResponseBodyParameter) *GetPublicParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *GetPublicParameterResponseBody) SetRequestId(v string) *GetPublicParameterResponseBody {
	s.RequestId = &v
	return s
}

type GetPublicParameterResponseBodyParameter struct {
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	Value            *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPublicParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s GetPublicParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetPublicParameterResponseBodyParameter) SetConstraints(v string) *GetPublicParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetCreatedBy(v string) *GetPublicParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetCreatedDate(v string) *GetPublicParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetDescription(v string) *GetPublicParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetId(v string) *GetPublicParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetName(v string) *GetPublicParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetParameterVersion(v int32) *GetPublicParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetRegionId(v string) *GetPublicParameterResponseBodyParameter {
	s.RegionId = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetShareType(v string) *GetPublicParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetType(v string) *GetPublicParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetUpdatedBy(v string) *GetPublicParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetUpdatedDate(v string) *GetPublicParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetPublicParameterResponseBodyParameter) SetValue(v string) *GetPublicParameterResponseBodyParameter {
	s.Value = &v
	return s
}

type GetPublicParameterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicParameterResponse) GoString() string {
	return s.String()
}

func (s *GetPublicParameterResponse) SetHeaders(v map[string]*string) *GetPublicParameterResponse {
	s.Headers = v
	return s
}

func (s *GetPublicParameterResponse) SetStatusCode(v int32) *GetPublicParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicParameterResponse) SetBody(v *GetPublicParameterResponseBody) *GetPublicParameterResponse {
	s.Body = v
	return s
}

type GetPublicPatchBaselineRequest struct {
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPublicPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetPublicPatchBaselineRequest) SetName(v string) *GetPublicPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *GetPublicPatchBaselineRequest) SetRegionId(v string) *GetPublicPatchBaselineRequest {
	s.RegionId = &v
	return s
}

type GetPublicPatchBaselineResponseBody struct {
	PatchBaseline *GetPublicPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicPatchBaselineResponseBody) SetPatchBaseline(v *GetPublicPatchBaselineResponseBodyPatchBaseline) *GetPublicPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *GetPublicPatchBaselineResponseBody) SetRequestId(v string) *GetPublicPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type GetPublicPatchBaselineResponseBodyPatchBaseline struct {
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetPublicPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s GetPublicPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetId(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetName(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *GetPublicPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *GetPublicPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type GetPublicPatchBaselineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetPublicPatchBaselineResponse) SetHeaders(v map[string]*string) *GetPublicPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetPublicPatchBaselineResponse) SetStatusCode(v int32) *GetPublicPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicPatchBaselineResponse) SetBody(v *GetPublicPatchBaselineResponseBody) *GetPublicPatchBaselineResponse {
	s.Body = v
	return s
}

type GetPublicTemplateRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetPublicTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetPublicTemplateRequest) SetRegionId(v string) *GetPublicTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetPublicTemplateRequest) SetTemplateName(v string) *GetPublicTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *GetPublicTemplateRequest) SetTemplateVersion(v string) *GetPublicTemplateRequest {
	s.TemplateVersion = &v
	return s
}

type GetPublicTemplateResponseBody struct {
	Content   *string                                `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *GetPublicTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetPublicTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicTemplateResponseBody) SetContent(v string) *GetPublicTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetPublicTemplateResponseBody) SetRequestId(v string) *GetPublicTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicTemplateResponseBody) SetTemplate(v *GetPublicTemplateResponseBodyTemplate) *GetPublicTemplateResponseBody {
	s.Template = v
	return s
}

type GetPublicTemplateResponseBodyTemplate struct {
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash            *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat  *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetPublicTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetPublicTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetPublicTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetPublicTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetPublicTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetDescription(v string) *GetPublicTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetHash(v string) *GetPublicTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetPopularity(v int32) *GetPublicTemplateResponseBodyTemplate {
	s.Popularity = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetShareType(v string) *GetPublicTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetPublicTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetTemplateId(v string) *GetPublicTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetTemplateName(v string) *GetPublicTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetPublicTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetPublicTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetPublicTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetPublicTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type GetPublicTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetPublicTemplateResponse) SetHeaders(v map[string]*string) *GetPublicTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetPublicTemplateResponse) SetStatusCode(v int32) *GetPublicTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicTemplateResponse) SetBody(v *GetPublicTemplateResponseBody) *GetPublicTemplateResponse {
	s.Body = v
	return s
}

type GetQuotaRequest struct {
	// This parameter is required.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) SetQuotaName(v string) *GetQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaRequest) SetRegionId(v string) *GetQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *GetQuotaRequest) SetUid(v string) *GetQuotaRequest {
	s.Uid = &v
	return s
}

type GetQuotaResponseBody struct {
	Quota     *GetQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uid       *string                    `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) SetQuota(v *GetQuotaResponseBodyQuota) *GetQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetUid(v string) *GetQuotaResponseBody {
	s.Uid = &v
	return s
}

type GetQuotaResponseBodyQuota struct {
	ConcurrentExecution *int32 `json:"ConcurrentExecution,omitempty" xml:"ConcurrentExecution,omitempty"`
	DailyTasks          *int32 `json:"DailyTasks,omitempty" xml:"DailyTasks,omitempty"`
	TotalTemplate       *int32 `json:"TotalTemplate,omitempty" xml:"TotalTemplate,omitempty"`
}

func (s GetQuotaResponseBodyQuota) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyQuota) SetConcurrentExecution(v int32) *GetQuotaResponseBodyQuota {
	s.ConcurrentExecution = &v
	return s
}

func (s *GetQuotaResponseBodyQuota) SetDailyTasks(v int32) *GetQuotaResponseBodyQuota {
	s.DailyTasks = &v
	return s
}

func (s *GetQuotaResponseBodyQuota) SetTotalTemplate(v int32) *GetQuotaResponseBodyQuota {
	s.TotalTemplate = &v
	return s
}

type GetQuotaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetStatusCode(v int32) *GetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

type GetUserExecutionTemplateRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUserExecutionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserExecutionTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetUserExecutionTemplateRequest) SetAliUid(v string) *GetUserExecutionTemplateRequest {
	s.AliUid = &v
	return s
}

func (s *GetUserExecutionTemplateRequest) SetExecutionId(v string) *GetUserExecutionTemplateRequest {
	s.ExecutionId = &v
	return s
}

func (s *GetUserExecutionTemplateRequest) SetRegionId(v string) *GetUserExecutionTemplateRequest {
	s.RegionId = &v
	return s
}

type GetUserExecutionTemplateResponseBody struct {
	Content   *string                                       `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *GetUserExecutionTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetUserExecutionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserExecutionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserExecutionTemplateResponseBody) SetContent(v string) *GetUserExecutionTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBody) SetRequestId(v string) *GetUserExecutionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBody) SetTemplate(v *GetUserExecutionTemplateResponseBodyTemplate) *GetUserExecutionTemplateResponseBody {
	s.Template = v
	return s
}

type GetUserExecutionTemplateResponseBodyTemplate struct {
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash            *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat  *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetUserExecutionTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetUserExecutionTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetDescription(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetHash(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetShareType(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetTemplateId(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetTemplateName(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetUserExecutionTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetUserExecutionTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type GetUserExecutionTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserExecutionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserExecutionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserExecutionTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetUserExecutionTemplateResponse) SetHeaders(v map[string]*string) *GetUserExecutionTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetUserExecutionTemplateResponse) SetStatusCode(v int32) *GetUserExecutionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserExecutionTemplateResponse) SetBody(v *GetUserExecutionTemplateResponseBody) *GetUserExecutionTemplateResponse {
	s.Body = v
	return s
}

type GetUserTemplateRequest struct {
	// This parameter is required.
	AliUid   *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetUserTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetUserTemplateRequest) SetAliUid(v string) *GetUserTemplateRequest {
	s.AliUid = &v
	return s
}

func (s *GetUserTemplateRequest) SetRegionId(v string) *GetUserTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserTemplateRequest) SetTemplateName(v string) *GetUserTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *GetUserTemplateRequest) SetTemplateVersion(v string) *GetUserTemplateRequest {
	s.TemplateVersion = &v
	return s
}

type GetUserTemplateResponseBody struct {
	Content   *string                              `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *GetUserTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetUserTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTemplateResponseBody) SetContent(v string) *GetUserTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetUserTemplateResponseBody) SetRequestId(v string) *GetUserTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserTemplateResponseBody) SetTemplate(v *GetUserTemplateResponseBodyTemplate) *GetUserTemplateResponseBody {
	s.Template = v
	return s
}

type GetUserTemplateResponseBodyTemplate struct {
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash            *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat  *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetUserTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetUserTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetUserTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetUserTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetUserTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetDescription(v string) *GetUserTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetHash(v string) *GetUserTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetShareType(v string) *GetUserTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetUserTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetTemplateId(v string) *GetUserTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetTemplateName(v string) *GetUserTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetUserTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetUserTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetUserTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetUserTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type GetUserTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetUserTemplateResponse) SetHeaders(v map[string]*string) *GetUserTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetUserTemplateResponse) SetStatusCode(v int32) *GetUserTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserTemplateResponse) SetBody(v *GetUserTemplateResponseBody) *GetUserTemplateResponse {
	s.Body = v
	return s
}

type ListActionsRequest struct {
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OOSActionName *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActionsRequest) GoString() string {
	return s.String()
}

func (s *ListActionsRequest) SetMaxResults(v int32) *ListActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionsRequest) SetNextToken(v string) *ListActionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionsRequest) SetOOSActionName(v string) *ListActionsRequest {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsRequest) SetRegionId(v string) *ListActionsRequest {
	s.RegionId = &v
	return s
}

type ListActionsResponseBody struct {
	Actions    []*ListActionsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	MaxResults *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBody) SetActions(v []*ListActionsResponseBodyActions) *ListActionsResponseBody {
	s.Actions = v
	return s
}

func (s *ListActionsResponseBody) SetMaxResults(v int32) *ListActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionsResponseBody) SetNextToken(v string) *ListActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionsResponseBody) SetRequestId(v string) *ListActionsResponseBody {
	s.RequestId = &v
	return s
}

type ListActionsResponseBodyActions struct {
	ActionType      *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OOSActionName   *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdateDate      *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListActionsResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponseBodyActions) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBodyActions) SetActionType(v string) *ListActionsResponseBodyActions {
	s.ActionType = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetCreatedDate(v string) *ListActionsResponseBodyActions {
	s.CreatedDate = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetDescription(v string) *ListActionsResponseBodyActions {
	s.Description = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetOOSActionName(v string) *ListActionsResponseBodyActions {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetPopularity(v int32) *ListActionsResponseBodyActions {
	s.Popularity = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetProperties(v string) *ListActionsResponseBodyActions {
	s.Properties = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetTemplateVersion(v string) *ListActionsResponseBodyActions {
	s.TemplateVersion = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetUpdateDate(v string) *ListActionsResponseBodyActions {
	s.UpdateDate = &v
	return s
}

type ListActionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponse) GoString() string {
	return s.String()
}

func (s *ListActionsResponse) SetHeaders(v map[string]*string) *ListActionsResponse {
	s.Headers = v
	return s
}

func (s *ListActionsResponse) SetStatusCode(v int32) *ListActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionsResponse) SetBody(v *ListActionsResponseBody) *ListActionsResponse {
	s.Body = v
	return s
}

type ListDefaultQuotaResponseBody struct {
	Quotas    []*ListDefaultQuotaResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDefaultQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ListDefaultQuotaResponseBody) SetQuotas(v []*ListDefaultQuotaResponseBodyQuotas) *ListDefaultQuotaResponseBody {
	s.Quotas = v
	return s
}

func (s *ListDefaultQuotaResponseBody) SetRequestId(v string) *ListDefaultQuotaResponseBody {
	s.RequestId = &v
	return s
}

type ListDefaultQuotaResponseBodyQuotas struct {
	ConcurrentExecution *int32 `json:"ConcurrentExecution,omitempty" xml:"ConcurrentExecution,omitempty"`
	DailyTasks          *int32 `json:"DailyTasks,omitempty" xml:"DailyTasks,omitempty"`
	TotalTemplate       *int32 `json:"TotalTemplate,omitempty" xml:"TotalTemplate,omitempty"`
}

func (s ListDefaultQuotaResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultQuotaResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListDefaultQuotaResponseBodyQuotas) SetConcurrentExecution(v int32) *ListDefaultQuotaResponseBodyQuotas {
	s.ConcurrentExecution = &v
	return s
}

func (s *ListDefaultQuotaResponseBodyQuotas) SetDailyTasks(v int32) *ListDefaultQuotaResponseBodyQuotas {
	s.DailyTasks = &v
	return s
}

func (s *ListDefaultQuotaResponseBodyQuotas) SetTotalTemplate(v int32) *ListDefaultQuotaResponseBodyQuotas {
	s.TotalTemplate = &v
	return s
}

type ListDefaultQuotaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDefaultQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDefaultQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDefaultQuotaResponse) GoString() string {
	return s.String()
}

func (s *ListDefaultQuotaResponse) SetHeaders(v map[string]*string) *ListDefaultQuotaResponse {
	s.Headers = v
	return s
}

func (s *ListDefaultQuotaResponse) SetStatusCode(v int32) *ListDefaultQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDefaultQuotaResponse) SetBody(v *ListDefaultQuotaResponseBody) *ListDefaultQuotaResponse {
	s.Body = v
	return s
}

type ListFailureMsgsRequest struct {
	MaxResults         *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestFingerprint *string `json:"RequestFingerprint,omitempty" xml:"RequestFingerprint,omitempty"`
}

func (s ListFailureMsgsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFailureMsgsRequest) GoString() string {
	return s.String()
}

func (s *ListFailureMsgsRequest) SetMaxResults(v int32) *ListFailureMsgsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFailureMsgsRequest) SetNextToken(v string) *ListFailureMsgsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFailureMsgsRequest) SetRequestFingerprint(v string) *ListFailureMsgsRequest {
	s.RequestFingerprint = &v
	return s
}

type ListFailureMsgsResponseBody struct {
	FailureMsgs []*ListFailureMsgsResponseBodyFailureMsgs `json:"FailureMsgs,omitempty" xml:"FailureMsgs,omitempty" type:"Repeated"`
	MaxResults  *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFailureMsgsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFailureMsgsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFailureMsgsResponseBody) SetFailureMsgs(v []*ListFailureMsgsResponseBodyFailureMsgs) *ListFailureMsgsResponseBody {
	s.FailureMsgs = v
	return s
}

func (s *ListFailureMsgsResponseBody) SetMaxResults(v int32) *ListFailureMsgsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFailureMsgsResponseBody) SetNextToken(v string) *ListFailureMsgsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFailureMsgsResponseBody) SetRequestId(v string) *ListFailureMsgsResponseBody {
	s.RequestId = &v
	return s
}

type ListFailureMsgsResponseBodyFailureMsgs struct {
	AliUid          *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ExecutionId     *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	MessageBody     *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
}

func (s ListFailureMsgsResponseBodyFailureMsgs) String() string {
	return tea.Prettify(s)
}

func (s ListFailureMsgsResponseBodyFailureMsgs) GoString() string {
	return s.String()
}

func (s *ListFailureMsgsResponseBodyFailureMsgs) SetAliUid(v string) *ListFailureMsgsResponseBodyFailureMsgs {
	s.AliUid = &v
	return s
}

func (s *ListFailureMsgsResponseBodyFailureMsgs) SetExecutionId(v string) *ListFailureMsgsResponseBodyFailureMsgs {
	s.ExecutionId = &v
	return s
}

func (s *ListFailureMsgsResponseBodyFailureMsgs) SetMessageBody(v string) *ListFailureMsgsResponseBodyFailureMsgs {
	s.MessageBody = &v
	return s
}

func (s *ListFailureMsgsResponseBodyFailureMsgs) SetReason(v string) *ListFailureMsgsResponseBodyFailureMsgs {
	s.Reason = &v
	return s
}

func (s *ListFailureMsgsResponseBodyFailureMsgs) SetTaskExecutionId(v string) *ListFailureMsgsResponseBodyFailureMsgs {
	s.TaskExecutionId = &v
	return s
}

type ListFailureMsgsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFailureMsgsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFailureMsgsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFailureMsgsResponse) GoString() string {
	return s.String()
}

func (s *ListFailureMsgsResponse) SetHeaders(v map[string]*string) *ListFailureMsgsResponse {
	s.Headers = v
	return s
}

func (s *ListFailureMsgsResponse) SetStatusCode(v int32) *ListFailureMsgsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFailureMsgsResponse) SetBody(v *ListFailureMsgsResponseBody) *ListFailureMsgsResponse {
	s.Body = v
	return s
}

type ListOOSLogsRequest struct {
	// This parameter is required.
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutionId        *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	MaxResults         *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestFingerprint *string `json:"RequestFingerprint,omitempty" xml:"RequestFingerprint,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListOOSLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOOSLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOOSLogsRequest) SetEndTime(v string) *ListOOSLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListOOSLogsRequest) SetExecutionId(v string) *ListOOSLogsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListOOSLogsRequest) SetMaxResults(v int32) *ListOOSLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOOSLogsRequest) SetNextToken(v string) *ListOOSLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOOSLogsRequest) SetRegionId(v string) *ListOOSLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOOSLogsRequest) SetRequestFingerprint(v string) *ListOOSLogsRequest {
	s.RequestFingerprint = &v
	return s
}

func (s *ListOOSLogsRequest) SetStartTime(v string) *ListOOSLogsRequest {
	s.StartTime = &v
	return s
}

type ListOOSLogsResponseBody struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OOSLogs    *string `json:"OOSLogs,omitempty" xml:"OOSLogs,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOOSLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOOSLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOOSLogsResponseBody) SetMaxResults(v int32) *ListOOSLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOOSLogsResponseBody) SetNextToken(v string) *ListOOSLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOOSLogsResponseBody) SetOOSLogs(v string) *ListOOSLogsResponseBody {
	s.OOSLogs = &v
	return s
}

func (s *ListOOSLogsResponseBody) SetRequestId(v string) *ListOOSLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListOOSLogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOOSLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOOSLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOOSLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOOSLogsResponse) SetHeaders(v map[string]*string) *ListOOSLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOOSLogsResponse) SetStatusCode(v int32) *ListOOSLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOOSLogsResponse) SetBody(v *ListOOSLogsResponseBody) *ListOOSLogsResponse {
	s.Body = v
	return s
}

type ListPublicParametersRequest struct {
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ParameterType *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	Path          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Recursive     *bool   `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortField     *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder     *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListPublicParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicParametersRequest) GoString() string {
	return s.String()
}

func (s *ListPublicParametersRequest) SetMaxResults(v int32) *ListPublicParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicParametersRequest) SetName(v string) *ListPublicParametersRequest {
	s.Name = &v
	return s
}

func (s *ListPublicParametersRequest) SetNextToken(v string) *ListPublicParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicParametersRequest) SetParameterType(v string) *ListPublicParametersRequest {
	s.ParameterType = &v
	return s
}

func (s *ListPublicParametersRequest) SetPath(v string) *ListPublicParametersRequest {
	s.Path = &v
	return s
}

func (s *ListPublicParametersRequest) SetRecursive(v bool) *ListPublicParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListPublicParametersRequest) SetRegionId(v string) *ListPublicParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListPublicParametersRequest) SetSortField(v string) *ListPublicParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListPublicParametersRequest) SetSortOrder(v string) *ListPublicParametersRequest {
	s.SortOrder = &v
	return s
}

type ListPublicParametersResponseBody struct {
	MaxResults *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Parameters []*ListPublicParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicParametersResponseBody) SetMaxResults(v int32) *ListPublicParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicParametersResponseBody) SetNextToken(v string) *ListPublicParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicParametersResponseBody) SetParameters(v []*ListPublicParametersResponseBodyParameters) *ListPublicParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *ListPublicParametersResponseBody) SetRequestId(v string) *ListPublicParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicParametersResponseBody) SetTotalCount(v int32) *ListPublicParametersResponseBody {
	s.TotalCount = &v
	return s
}

type ListPublicParametersResponseBodyParameters struct {
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *string `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListPublicParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListPublicParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListPublicParametersResponseBodyParameters) SetCreatedBy(v string) *ListPublicParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetCreatedDate(v string) *ListPublicParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetDescription(v string) *ListPublicParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetId(v string) *ListPublicParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetName(v string) *ListPublicParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetParameterVersion(v string) *ListPublicParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetRegionId(v string) *ListPublicParametersResponseBodyParameters {
	s.RegionId = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetShareType(v string) *ListPublicParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetType(v string) *ListPublicParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetUpdatedBy(v string) *ListPublicParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListPublicParametersResponseBodyParameters) SetUpdatedDate(v string) *ListPublicParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

type ListPublicParametersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicParametersResponse) GoString() string {
	return s.String()
}

func (s *ListPublicParametersResponse) SetHeaders(v map[string]*string) *ListPublicParametersResponse {
	s.Headers = v
	return s
}

func (s *ListPublicParametersResponse) SetStatusCode(v int32) *ListPublicParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicParametersResponse) SetBody(v *ListPublicParametersResponseBody) *ListPublicParametersResponse {
	s.Body = v
	return s
}

type ListPublicPatchBaselinesRequest struct {
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListPublicPatchBaselinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicPatchBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListPublicPatchBaselinesRequest) SetMaxResults(v int32) *ListPublicPatchBaselinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicPatchBaselinesRequest) SetName(v string) *ListPublicPatchBaselinesRequest {
	s.Name = &v
	return s
}

func (s *ListPublicPatchBaselinesRequest) SetNextToken(v string) *ListPublicPatchBaselinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicPatchBaselinesRequest) SetOperationSystem(v string) *ListPublicPatchBaselinesRequest {
	s.OperationSystem = &v
	return s
}

func (s *ListPublicPatchBaselinesRequest) SetRegionId(v string) *ListPublicPatchBaselinesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPublicPatchBaselinesRequest) SetShareType(v string) *ListPublicPatchBaselinesRequest {
	s.ShareType = &v
	return s
}

type ListPublicPatchBaselinesResponseBody struct {
	MaxResults     *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PatchBaselines []*ListPublicPatchBaselinesResponseBodyPatchBaselines `json:"PatchBaselines,omitempty" xml:"PatchBaselines,omitempty" type:"Repeated"`
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublicPatchBaselinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicPatchBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicPatchBaselinesResponseBody) SetMaxResults(v int32) *ListPublicPatchBaselinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBody) SetNextToken(v string) *ListPublicPatchBaselinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBody) SetPatchBaselines(v []*ListPublicPatchBaselinesResponseBodyPatchBaselines) *ListPublicPatchBaselinesResponseBody {
	s.PatchBaselines = v
	return s
}

func (s *ListPublicPatchBaselinesResponseBody) SetRequestId(v string) *ListPublicPatchBaselinesResponseBody {
	s.RequestId = &v
	return s
}

type ListPublicPatchBaselinesResponseBodyPatchBaselines struct {
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDefault       *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListPublicPatchBaselinesResponseBodyPatchBaselines) String() string {
	return tea.Prettify(s)
}

func (s ListPublicPatchBaselinesResponseBodyPatchBaselines) GoString() string {
	return s.String()
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetCreatedBy(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedBy = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetCreatedDate(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedDate = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetDescription(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.Description = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetId(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.Id = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetIsDefault(v bool) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.IsDefault = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetName(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.Name = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetOperationSystem(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.OperationSystem = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetShareType(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.ShareType = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetUpdatedBy(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedBy = &v
	return s
}

func (s *ListPublicPatchBaselinesResponseBodyPatchBaselines) SetUpdatedDate(v string) *ListPublicPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedDate = &v
	return s
}

type ListPublicPatchBaselinesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicPatchBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicPatchBaselinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicPatchBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListPublicPatchBaselinesResponse) SetHeaders(v map[string]*string) *ListPublicPatchBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListPublicPatchBaselinesResponse) SetStatusCode(v int32) *ListPublicPatchBaselinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicPatchBaselinesResponse) SetBody(v *ListPublicPatchBaselinesResponseBody) *ListPublicPatchBaselinesResponse {
	s.Body = v
	return s
}

type ListPublicTemplateRegistrationsRequest struct {
	MaxResults     *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListPublicTemplateRegistrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplateRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicTemplateRegistrationsRequest) SetMaxResults(v int64) *ListPublicTemplateRegistrationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicTemplateRegistrationsRequest) SetNextToken(v string) *ListPublicTemplateRegistrationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicTemplateRegistrationsRequest) SetRegionId(v string) *ListPublicTemplateRegistrationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListPublicTemplateRegistrationsRequest) SetRegistrationId(v string) *ListPublicTemplateRegistrationsRequest {
	s.RegistrationId = &v
	return s
}

func (s *ListPublicTemplateRegistrationsRequest) SetStatus(v string) *ListPublicTemplateRegistrationsRequest {
	s.Status = &v
	return s
}

func (s *ListPublicTemplateRegistrationsRequest) SetTemplateName(v string) *ListPublicTemplateRegistrationsRequest {
	s.TemplateName = &v
	return s
}

type ListPublicTemplateRegistrationsResponseBody struct {
	MaxResults    *int32                                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Registrations []*ListPublicTemplateRegistrationsResponseBodyRegistrations `json:"Registrations,omitempty" xml:"Registrations,omitempty" type:"Repeated"`
	RequestId     *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublicTemplateRegistrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplateRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicTemplateRegistrationsResponseBody) SetMaxResults(v int32) *ListPublicTemplateRegistrationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBody) SetNextToken(v string) *ListPublicTemplateRegistrationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBody) SetRegistrations(v []*ListPublicTemplateRegistrationsResponseBodyRegistrations) *ListPublicTemplateRegistrationsResponseBody {
	s.Registrations = v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBody) SetRequestId(v string) *ListPublicTemplateRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

type ListPublicTemplateRegistrationsResponseBodyRegistrations struct {
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Detail          *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	RegistrationId  *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	ShowPages       *string `json:"ShowPages,omitempty" xml:"ShowPages,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListPublicTemplateRegistrationsResponseBodyRegistrations) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplateRegistrationsResponseBodyRegistrations) GoString() string {
	return s.String()
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetComment(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.Comment = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetCreatedDate(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.CreatedDate = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetDetail(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.Detail = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetRegistrationId(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetShowPages(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.ShowPages = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetStatus(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.Status = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetTemplateId(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.TemplateId = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetTemplateName(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.TemplateName = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetTemplateVersion(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.TemplateVersion = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponseBodyRegistrations) SetUpdatedDate(v string) *ListPublicTemplateRegistrationsResponseBodyRegistrations {
	s.UpdatedDate = &v
	return s
}

type ListPublicTemplateRegistrationsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicTemplateRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicTemplateRegistrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplateRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicTemplateRegistrationsResponse) SetHeaders(v map[string]*string) *ListPublicTemplateRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicTemplateRegistrationsResponse) SetStatusCode(v int32) *ListPublicTemplateRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicTemplateRegistrationsResponse) SetBody(v *ListPublicTemplateRegistrationsResponseBody) *ListPublicTemplateRegistrationsResponse {
	s.Body = v
	return s
}

type ListPublicTemplatesRequest struct {
	CreatedBy         *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDateAfter  *string `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	CreatedDateBefore *string `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	IsExample         *bool   `json:"IsExample,omitempty" xml:"IsExample,omitempty"`
	MaxResults        *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Popularity        *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType         *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	SortField         *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder         *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	TemplateFormat    *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListPublicTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPublicTemplatesRequest) SetCreatedBy(v string) *ListPublicTemplatesRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetCreatedDateAfter(v string) *ListPublicTemplatesRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetCreatedDateBefore(v string) *ListPublicTemplatesRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetIsExample(v bool) *ListPublicTemplatesRequest {
	s.IsExample = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetMaxResults(v int32) *ListPublicTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetNextToken(v string) *ListPublicTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetPopularity(v int32) *ListPublicTemplatesRequest {
	s.Popularity = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetRegionId(v string) *ListPublicTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetShareType(v string) *ListPublicTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetSortField(v string) *ListPublicTemplatesRequest {
	s.SortField = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetSortOrder(v string) *ListPublicTemplatesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetTemplateFormat(v string) *ListPublicTemplatesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListPublicTemplatesRequest) SetTemplateName(v string) *ListPublicTemplatesRequest {
	s.TemplateName = &v
	return s
}

type ListPublicTemplatesResponseBody struct {
	MaxResults *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates  []*ListPublicTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListPublicTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicTemplatesResponseBody) SetMaxResults(v int32) *ListPublicTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicTemplatesResponseBody) SetNextToken(v string) *ListPublicTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicTemplatesResponseBody) SetRequestId(v string) *ListPublicTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicTemplatesResponseBody) SetTemplates(v []*ListPublicTemplatesResponseBodyTemplates) *ListPublicTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListPublicTemplatesResponseBodyTemplates struct {
	CreatedBy           *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate         *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash                *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Popularity          *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	ShareType           *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat      *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion     *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TotalExecutionCount *int32  `json:"TotalExecutionCount,omitempty" xml:"TotalExecutionCount,omitempty"`
	UpdatedBy           *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate         *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListPublicTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetCreatedBy(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.CreatedBy = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetCreatedDate(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.CreatedDate = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetDescription(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetHash(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.Hash = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetPopularity(v int32) *ListPublicTemplatesResponseBodyTemplates {
	s.Popularity = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetShareType(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetTemplateFormat(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.TemplateFormat = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetTotalExecutionCount(v int32) *ListPublicTemplatesResponseBodyTemplates {
	s.TotalExecutionCount = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetUpdatedBy(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.UpdatedBy = &v
	return s
}

func (s *ListPublicTemplatesResponseBodyTemplates) SetUpdatedDate(v string) *ListPublicTemplatesResponseBodyTemplates {
	s.UpdatedDate = &v
	return s
}

type ListPublicTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPublicTemplatesResponse) SetHeaders(v map[string]*string) *ListPublicTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPublicTemplatesResponse) SetStatusCode(v int32) *ListPublicTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicTemplatesResponse) SetBody(v *ListPublicTemplatesResponseBody) *ListPublicTemplatesResponse {
	s.Body = v
	return s
}

type ListUserExecutionLogsRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	ExecutionId     *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	LogType         *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
}

func (s ListUserExecutionLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionLogsRequest) GoString() string {
	return s.String()
}

func (s *ListUserExecutionLogsRequest) SetAliUid(v string) *ListUserExecutionLogsRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserExecutionLogsRequest) SetExecutionId(v string) *ListUserExecutionLogsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListUserExecutionLogsRequest) SetLogType(v string) *ListUserExecutionLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListUserExecutionLogsRequest) SetMaxResults(v int32) *ListUserExecutionLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserExecutionLogsRequest) SetNextToken(v string) *ListUserExecutionLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserExecutionLogsRequest) SetRegionId(v string) *ListUserExecutionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserExecutionLogsRequest) SetTaskExecutionId(v string) *ListUserExecutionLogsRequest {
	s.TaskExecutionId = &v
	return s
}

type ListUserExecutionLogsResponseBody struct {
	ExecutionLogs []*ListUserExecutionLogsResponseBodyExecutionLogs `json:"ExecutionLogs,omitempty" xml:"ExecutionLogs,omitempty" type:"Repeated"`
	MaxResults    *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserExecutionLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserExecutionLogsResponseBody) SetExecutionLogs(v []*ListUserExecutionLogsResponseBodyExecutionLogs) *ListUserExecutionLogsResponseBody {
	s.ExecutionLogs = v
	return s
}

func (s *ListUserExecutionLogsResponseBody) SetMaxResults(v int32) *ListUserExecutionLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserExecutionLogsResponseBody) SetNextToken(v string) *ListUserExecutionLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserExecutionLogsResponseBody) SetRequestId(v string) *ListUserExecutionLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListUserExecutionLogsResponseBodyExecutionLogs struct {
	LogType         *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	Timestamp       *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListUserExecutionLogsResponseBodyExecutionLogs) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionLogsResponseBodyExecutionLogs) GoString() string {
	return s.String()
}

func (s *ListUserExecutionLogsResponseBodyExecutionLogs) SetLogType(v string) *ListUserExecutionLogsResponseBodyExecutionLogs {
	s.LogType = &v
	return s
}

func (s *ListUserExecutionLogsResponseBodyExecutionLogs) SetMessage(v string) *ListUserExecutionLogsResponseBodyExecutionLogs {
	s.Message = &v
	return s
}

func (s *ListUserExecutionLogsResponseBodyExecutionLogs) SetTaskExecutionId(v string) *ListUserExecutionLogsResponseBodyExecutionLogs {
	s.TaskExecutionId = &v
	return s
}

func (s *ListUserExecutionLogsResponseBodyExecutionLogs) SetTimestamp(v string) *ListUserExecutionLogsResponseBodyExecutionLogs {
	s.Timestamp = &v
	return s
}

type ListUserExecutionLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserExecutionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserExecutionLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionLogsResponse) GoString() string {
	return s.String()
}

func (s *ListUserExecutionLogsResponse) SetHeaders(v map[string]*string) *ListUserExecutionLogsResponse {
	s.Headers = v
	return s
}

func (s *ListUserExecutionLogsResponse) SetStatusCode(v int32) *ListUserExecutionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserExecutionLogsResponse) SetBody(v *ListUserExecutionLogsResponseBody) *ListUserExecutionLogsResponse {
	s.Body = v
	return s
}

type ListUserExecutionsRequest struct {
	// This parameter is required.
	AliUid                *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndDateAfter          *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	EndDateBefore         *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	ExecutedBy            *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	ExecutionId           *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	IncludeChildExecution *bool   `json:"IncludeChildExecution,omitempty" xml:"IncludeChildExecution,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Mode                  *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ParentExecutionId     *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	RamRole               *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortField             *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder             *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartDateAfter        *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	StartDateBefore       *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName          *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListUserExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserExecutionsRequest) SetAliUid(v string) *ListUserExecutionsRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserExecutionsRequest) SetEndDateAfter(v string) *ListUserExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListUserExecutionsRequest) SetEndDateBefore(v string) *ListUserExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListUserExecutionsRequest) SetExecutedBy(v string) *ListUserExecutionsRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListUserExecutionsRequest) SetExecutionId(v string) *ListUserExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListUserExecutionsRequest) SetIncludeChildExecution(v bool) *ListUserExecutionsRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListUserExecutionsRequest) SetMaxResults(v int32) *ListUserExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserExecutionsRequest) SetMode(v string) *ListUserExecutionsRequest {
	s.Mode = &v
	return s
}

func (s *ListUserExecutionsRequest) SetNextToken(v string) *ListUserExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserExecutionsRequest) SetParentExecutionId(v string) *ListUserExecutionsRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListUserExecutionsRequest) SetRamRole(v string) *ListUserExecutionsRequest {
	s.RamRole = &v
	return s
}

func (s *ListUserExecutionsRequest) SetRegionId(v string) *ListUserExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserExecutionsRequest) SetSortField(v string) *ListUserExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListUserExecutionsRequest) SetSortOrder(v string) *ListUserExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserExecutionsRequest) SetStartDateAfter(v string) *ListUserExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListUserExecutionsRequest) SetStartDateBefore(v string) *ListUserExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListUserExecutionsRequest) SetStatus(v string) *ListUserExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListUserExecutionsRequest) SetTemplateName(v string) *ListUserExecutionsRequest {
	s.TemplateName = &v
	return s
}

type ListUserExecutionsResponseBody struct {
	Executions []*ListUserExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	MaxResults *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserExecutionsResponseBody) SetExecutions(v []*ListUserExecutionsResponseBodyExecutions) *ListUserExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListUserExecutionsResponseBody) SetMaxResults(v int32) *ListUserExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserExecutionsResponseBody) SetNextToken(v string) *ListUserExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserExecutionsResponseBody) SetRequestId(v string) *ListUserExecutionsResponseBody {
	s.RequestId = &v
	return s
}

type ListUserExecutionsResponseBodyExecutions struct {
	Counters          *string                                                 `json:"Counters,omitempty" xml:"Counters,omitempty"`
	CreateDate        *string                                                 `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	CurrentTasks      []*ListUserExecutionsResponseBodyExecutionsCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
	EndDate           *string                                                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ExecutedBy        *string                                                 `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	ExecutionId       *string                                                 `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	IsParent          *bool                                                   `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	Mode              *string                                                 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Outputs           *string                                                 `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Parameters        *string                                                 `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ParentExecutionId *string                                                 `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	RamRole           *string                                                 `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	SafetyCheck       *string                                                 `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	StartDate         *string                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status            *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage     *string                                                 `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	TemplateId        *string                                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName      *string                                                 `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion   *string                                                 `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdateDate        *string                                                 `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	WaitingStatus     *string                                                 `json:"WaitingStatus,omitempty" xml:"WaitingStatus,omitempty"`
}

func (s ListUserExecutionsResponseBodyExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListUserExecutionsResponseBodyExecutions) SetCounters(v string) *ListUserExecutionsResponseBodyExecutions {
	s.Counters = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetCreateDate(v string) *ListUserExecutionsResponseBodyExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetCurrentTasks(v []*ListUserExecutionsResponseBodyExecutionsCurrentTasks) *ListUserExecutionsResponseBodyExecutions {
	s.CurrentTasks = v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetEndDate(v string) *ListUserExecutionsResponseBodyExecutions {
	s.EndDate = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetExecutedBy(v string) *ListUserExecutionsResponseBodyExecutions {
	s.ExecutedBy = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetExecutionId(v string) *ListUserExecutionsResponseBodyExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetIsParent(v bool) *ListUserExecutionsResponseBodyExecutions {
	s.IsParent = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetMode(v string) *ListUserExecutionsResponseBodyExecutions {
	s.Mode = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetOutputs(v string) *ListUserExecutionsResponseBodyExecutions {
	s.Outputs = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetParameters(v string) *ListUserExecutionsResponseBodyExecutions {
	s.Parameters = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetParentExecutionId(v string) *ListUserExecutionsResponseBodyExecutions {
	s.ParentExecutionId = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetRamRole(v string) *ListUserExecutionsResponseBodyExecutions {
	s.RamRole = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetSafetyCheck(v string) *ListUserExecutionsResponseBodyExecutions {
	s.SafetyCheck = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetStartDate(v string) *ListUserExecutionsResponseBodyExecutions {
	s.StartDate = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetStatus(v string) *ListUserExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetStatusMessage(v string) *ListUserExecutionsResponseBodyExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetTemplateId(v string) *ListUserExecutionsResponseBodyExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetTemplateName(v string) *ListUserExecutionsResponseBodyExecutions {
	s.TemplateName = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetTemplateVersion(v string) *ListUserExecutionsResponseBodyExecutions {
	s.TemplateVersion = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetUpdateDate(v string) *ListUserExecutionsResponseBodyExecutions {
	s.UpdateDate = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutions) SetWaitingStatus(v string) *ListUserExecutionsResponseBodyExecutions {
	s.WaitingStatus = &v
	return s
}

type ListUserExecutionsResponseBodyExecutionsCurrentTasks struct {
	TaskAction      *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListUserExecutionsResponseBodyExecutionsCurrentTasks) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionsResponseBodyExecutionsCurrentTasks) GoString() string {
	return s.String()
}

func (s *ListUserExecutionsResponseBodyExecutionsCurrentTasks) SetTaskAction(v string) *ListUserExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskAction = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutionsCurrentTasks) SetTaskExecutionId(v string) *ListUserExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *ListUserExecutionsResponseBodyExecutionsCurrentTasks) SetTaskName(v string) *ListUserExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskName = &v
	return s
}

type ListUserExecutionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserExecutionsResponse) SetHeaders(v map[string]*string) *ListUserExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserExecutionsResponse) SetStatusCode(v int32) *ListUserExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserExecutionsResponse) SetBody(v *ListUserExecutionsResponseBody) *ListUserExecutionsResponse {
	s.Body = v
	return s
}

type ListUserInstancePatchStatesRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserInstancePatchStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchStatesRequest) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchStatesRequest) SetAliUid(v string) *ListUserInstancePatchStatesRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserInstancePatchStatesRequest) SetInstanceIds(v string) *ListUserInstancePatchStatesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListUserInstancePatchStatesRequest) SetMaxResults(v int32) *ListUserInstancePatchStatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserInstancePatchStatesRequest) SetNextToken(v string) *ListUserInstancePatchStatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserInstancePatchStatesRequest) SetRegionId(v string) *ListUserInstancePatchStatesRequest {
	s.RegionId = &v
	return s
}

type ListUserInstancePatchStatesResponseBody struct {
	InstancePatchStates []*ListUserInstancePatchStatesResponseBodyInstancePatchStates `json:"InstancePatchStates,omitempty" xml:"InstancePatchStates,omitempty" type:"Repeated"`
	MaxResults          *int32                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserInstancePatchStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchStatesResponseBody) SetInstancePatchStates(v []*ListUserInstancePatchStatesResponseBodyInstancePatchStates) *ListUserInstancePatchStatesResponseBody {
	s.InstancePatchStates = v
	return s
}

func (s *ListUserInstancePatchStatesResponseBody) SetMaxResults(v int32) *ListUserInstancePatchStatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBody) SetNextToken(v string) *ListUserInstancePatchStatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBody) SetRequestId(v string) *ListUserInstancePatchStatesResponseBody {
	s.RequestId = &v
	return s
}

type ListUserInstancePatchStatesResponseBodyInstancePatchStates struct {
	BaselineId                  *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	FailedCount                 *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	InstalledCount              *string `json:"InstalledCount,omitempty" xml:"InstalledCount,omitempty"`
	InstalledOtherCount         *string `json:"InstalledOtherCount,omitempty" xml:"InstalledOtherCount,omitempty"`
	InstalledPendingRebootCount *string `json:"InstalledPendingRebootCount,omitempty" xml:"InstalledPendingRebootCount,omitempty"`
	InstalledRejectedCount      *string `json:"InstalledRejectedCount,omitempty" xml:"InstalledRejectedCount,omitempty"`
	InstanceId                  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MissingCount                *string `json:"MissingCount,omitempty" xml:"MissingCount,omitempty"`
	OperationEndTime            *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OperationStartTime          *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	OperationType               *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerInformation            *string `json:"OwnerInformation,omitempty" xml:"OwnerInformation,omitempty"`
	PatchGroup                  *string `json:"PatchGroup,omitempty" xml:"PatchGroup,omitempty"`
}

func (s ListUserInstancePatchStatesResponseBodyInstancePatchStates) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchStatesResponseBodyInstancePatchStates) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetBaselineId(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.BaselineId = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetFailedCount(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.FailedCount = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledCount(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledCount = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledOtherCount(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledOtherCount = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledPendingRebootCount(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledPendingRebootCount = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledRejectedCount(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledRejectedCount = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetInstanceId(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstanceId = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetMissingCount(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.MissingCount = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetOperationEndTime(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationEndTime = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetOperationStartTime(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationStartTime = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetOperationType(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationType = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetOwnerInformation(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.OwnerInformation = &v
	return s
}

func (s *ListUserInstancePatchStatesResponseBodyInstancePatchStates) SetPatchGroup(v string) *ListUserInstancePatchStatesResponseBodyInstancePatchStates {
	s.PatchGroup = &v
	return s
}

type ListUserInstancePatchStatesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserInstancePatchStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserInstancePatchStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchStatesResponse) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchStatesResponse) SetHeaders(v map[string]*string) *ListUserInstancePatchStatesResponse {
	s.Headers = v
	return s
}

func (s *ListUserInstancePatchStatesResponse) SetStatusCode(v int32) *ListUserInstancePatchStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserInstancePatchStatesResponse) SetBody(v *ListUserInstancePatchStatesResponseBody) *ListUserInstancePatchStatesResponse {
	s.Body = v
	return s
}

type ListUserInstancePatchesRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserInstancePatchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchesRequest) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchesRequest) SetAliUid(v string) *ListUserInstancePatchesRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserInstancePatchesRequest) SetInstanceId(v string) *ListUserInstancePatchesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserInstancePatchesRequest) SetMaxResults(v int32) *ListUserInstancePatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserInstancePatchesRequest) SetNextToken(v string) *ListUserInstancePatchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserInstancePatchesRequest) SetRegionId(v string) *ListUserInstancePatchesRequest {
	s.RegionId = &v
	return s
}

type ListUserInstancePatchesResponseBody struct {
	MaxResults *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Patches    []*ListUserInstancePatchesResponseBodyPatches `json:"Patches,omitempty" xml:"Patches,omitempty" type:"Repeated"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserInstancePatchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchesResponseBody) SetMaxResults(v int32) *ListUserInstancePatchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserInstancePatchesResponseBody) SetNextToken(v string) *ListUserInstancePatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserInstancePatchesResponseBody) SetPatches(v []*ListUserInstancePatchesResponseBodyPatches) *ListUserInstancePatchesResponseBody {
	s.Patches = v
	return s
}

func (s *ListUserInstancePatchesResponseBody) SetRequestId(v string) *ListUserInstancePatchesResponseBody {
	s.RequestId = &v
	return s
}

type ListUserInstancePatchesResponseBodyPatches struct {
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	InstalledTime  *string `json:"InstalledTime,omitempty" xml:"InstalledTime,omitempty"`
	KBId           *string `json:"KBId,omitempty" xml:"KBId,omitempty"`
	Severity       *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListUserInstancePatchesResponseBodyPatches) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchesResponseBodyPatches) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchesResponseBodyPatches) SetClassification(v string) *ListUserInstancePatchesResponseBodyPatches {
	s.Classification = &v
	return s
}

func (s *ListUserInstancePatchesResponseBodyPatches) SetInstalledTime(v string) *ListUserInstancePatchesResponseBodyPatches {
	s.InstalledTime = &v
	return s
}

func (s *ListUserInstancePatchesResponseBodyPatches) SetKBId(v string) *ListUserInstancePatchesResponseBodyPatches {
	s.KBId = &v
	return s
}

func (s *ListUserInstancePatchesResponseBodyPatches) SetSeverity(v string) *ListUserInstancePatchesResponseBodyPatches {
	s.Severity = &v
	return s
}

func (s *ListUserInstancePatchesResponseBodyPatches) SetStatus(v string) *ListUserInstancePatchesResponseBodyPatches {
	s.Status = &v
	return s
}

func (s *ListUserInstancePatchesResponseBodyPatches) SetTitle(v string) *ListUserInstancePatchesResponseBodyPatches {
	s.Title = &v
	return s
}

type ListUserInstancePatchesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserInstancePatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserInstancePatchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserInstancePatchesResponse) GoString() string {
	return s.String()
}

func (s *ListUserInstancePatchesResponse) SetHeaders(v map[string]*string) *ListUserInstancePatchesResponse {
	s.Headers = v
	return s
}

func (s *ListUserInstancePatchesResponse) SetStatusCode(v int32) *ListUserInstancePatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserInstancePatchesResponse) SetBody(v *ListUserInstancePatchesResponseBody) *ListUserInstancePatchesResponse {
	s.Body = v
	return s
}

type ListUserInventoryEntriesRequest struct {
	// This parameter is required.
	AliUid *string                                  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Filter []*ListUserInventoryEntriesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListUserInventoryEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserInventoryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListUserInventoryEntriesRequest) SetAliUid(v string) *ListUserInventoryEntriesRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserInventoryEntriesRequest) SetFilter(v []*ListUserInventoryEntriesRequestFilter) *ListUserInventoryEntriesRequest {
	s.Filter = v
	return s
}

func (s *ListUserInventoryEntriesRequest) SetInstanceId(v string) *ListUserInventoryEntriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserInventoryEntriesRequest) SetMaxResults(v int32) *ListUserInventoryEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserInventoryEntriesRequest) SetNextToken(v string) *ListUserInventoryEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserInventoryEntriesRequest) SetRegionId(v string) *ListUserInventoryEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserInventoryEntriesRequest) SetTypeName(v string) *ListUserInventoryEntriesRequest {
	s.TypeName = &v
	return s
}

type ListUserInventoryEntriesRequestFilter struct {
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListUserInventoryEntriesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListUserInventoryEntriesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListUserInventoryEntriesRequestFilter) SetName(v string) *ListUserInventoryEntriesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListUserInventoryEntriesRequestFilter) SetOperator(v string) *ListUserInventoryEntriesRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListUserInventoryEntriesRequestFilter) SetValue(v []*string) *ListUserInventoryEntriesRequestFilter {
	s.Value = v
	return s
}

type ListUserInventoryEntriesResponseBody struct {
	CaptureTime   *string                  `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	Entries       []map[string]interface{} `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	InstanceId    *string                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults    *int32                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchemaVersion *string                  `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	TypeName      *string                  `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListUserInventoryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserInventoryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserInventoryEntriesResponseBody) SetCaptureTime(v string) *ListUserInventoryEntriesResponseBody {
	s.CaptureTime = &v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetEntries(v []map[string]interface{}) *ListUserInventoryEntriesResponseBody {
	s.Entries = v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetInstanceId(v string) *ListUserInventoryEntriesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetMaxResults(v int32) *ListUserInventoryEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetNextToken(v string) *ListUserInventoryEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetRequestId(v string) *ListUserInventoryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetSchemaVersion(v string) *ListUserInventoryEntriesResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *ListUserInventoryEntriesResponseBody) SetTypeName(v string) *ListUserInventoryEntriesResponseBody {
	s.TypeName = &v
	return s
}

type ListUserInventoryEntriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserInventoryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserInventoryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserInventoryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListUserInventoryEntriesResponse) SetHeaders(v map[string]*string) *ListUserInventoryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListUserInventoryEntriesResponse) SetStatusCode(v int32) *ListUserInventoryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserInventoryEntriesResponse) SetBody(v *ListUserInventoryEntriesResponseBody) *ListUserInventoryEntriesResponse {
	s.Body = v
	return s
}

type ListUserTaskExecutionsRequest struct {
	// This parameter is required.
	AliUid                    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndDateAfter              *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	EndDateBefore             *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	ExecutionId               *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	IncludeChildTaskExecution *bool   `json:"IncludeChildTaskExecution,omitempty" xml:"IncludeChildTaskExecution,omitempty"`
	MaxResults                *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ParentTaskExecutionId     *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortField                 *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder                 *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartDateAfter            *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	StartDateBefore           *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskAction                *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskExecutionId           *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	TaskName                  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListUserTaskExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserTaskExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserTaskExecutionsRequest) SetAliUid(v string) *ListUserTaskExecutionsRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetEndDateAfter(v string) *ListUserTaskExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetEndDateBefore(v string) *ListUserTaskExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetExecutionId(v string) *ListUserTaskExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetIncludeChildTaskExecution(v bool) *ListUserTaskExecutionsRequest {
	s.IncludeChildTaskExecution = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetMaxResults(v int32) *ListUserTaskExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetNextToken(v string) *ListUserTaskExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetParentTaskExecutionId(v string) *ListUserTaskExecutionsRequest {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetRegionId(v string) *ListUserTaskExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetSortField(v string) *ListUserTaskExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetSortOrder(v string) *ListUserTaskExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetStartDateAfter(v string) *ListUserTaskExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetStartDateBefore(v string) *ListUserTaskExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetStatus(v string) *ListUserTaskExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetTaskAction(v string) *ListUserTaskExecutionsRequest {
	s.TaskAction = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetTaskExecutionId(v string) *ListUserTaskExecutionsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsRequest) SetTaskName(v string) *ListUserTaskExecutionsRequest {
	s.TaskName = &v
	return s
}

type ListUserTaskExecutionsResponseBody struct {
	MaxResults     *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskExecutions []*ListUserTaskExecutionsResponseBodyTaskExecutions `json:"TaskExecutions,omitempty" xml:"TaskExecutions,omitempty" type:"Repeated"`
}

func (s ListUserTaskExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserTaskExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserTaskExecutionsResponseBody) SetMaxResults(v int32) *ListUserTaskExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBody) SetNextToken(v string) *ListUserTaskExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBody) SetRequestId(v string) *ListUserTaskExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBody) SetTaskExecutions(v []*ListUserTaskExecutionsResponseBodyTaskExecutions) *ListUserTaskExecutionsResponseBody {
	s.TaskExecutions = v
	return s
}

type ListUserTaskExecutionsResponseBodyTaskExecutions struct {
	ChildExecutionId      *string `json:"ChildExecutionId,omitempty" xml:"ChildExecutionId,omitempty"`
	CreateDate            *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	EndDate               *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ExecutionId           *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	ExtraData             *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Loop                  *string `json:"Loop,omitempty" xml:"Loop,omitempty"`
	LoopBatchNumber       *int32  `json:"LoopBatchNumber,omitempty" xml:"LoopBatchNumber,omitempty"`
	LoopItem              *string `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	Outputs               *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	ParentTaskExecutionId *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	Properties            *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	StartDate             *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage         *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	TaskAction            *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskExecutionId       *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	TaskName              *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateId            *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UpdateDate            *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListUserTaskExecutionsResponseBodyTaskExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListUserTaskExecutionsResponseBodyTaskExecutions) GoString() string {
	return s.String()
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetChildExecutionId(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.ChildExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetCreateDate(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetEndDate(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.EndDate = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetExecutionId(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetExtraData(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.ExtraData = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetLoop(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.Loop = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetLoopBatchNumber(v int32) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.LoopBatchNumber = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetLoopItem(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.LoopItem = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetOutputs(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.Outputs = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetParentTaskExecutionId(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetProperties(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.Properties = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetStartDate(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.StartDate = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetStatus(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.Status = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetStatusMessage(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetTaskAction(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.TaskAction = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetTaskExecutionId(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.TaskExecutionId = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetTaskName(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.TaskName = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetTemplateId(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListUserTaskExecutionsResponseBodyTaskExecutions) SetUpdateDate(v string) *ListUserTaskExecutionsResponseBodyTaskExecutions {
	s.UpdateDate = &v
	return s
}

type ListUserTaskExecutionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserTaskExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserTaskExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserTaskExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserTaskExecutionsResponse) SetHeaders(v map[string]*string) *ListUserTaskExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserTaskExecutionsResponse) SetStatusCode(v int32) *ListUserTaskExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserTaskExecutionsResponse) SetBody(v *ListUserTaskExecutionsResponseBody) *ListUserTaskExecutionsResponse {
	s.Body = v
	return s
}

type ListUserTemplatesRequest struct {
	// This parameter is required.
	AliUid            *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Category          *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedBy         *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDateAfter  *string `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	CreatedDateBefore *string `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	MaxResults        *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Popularity        *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType         *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	SortField         *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SortOrder         *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	TemplateFormat    *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType      *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListUserTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListUserTemplatesRequest) SetAliUid(v string) *ListUserTemplatesRequest {
	s.AliUid = &v
	return s
}

func (s *ListUserTemplatesRequest) SetCategory(v string) *ListUserTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListUserTemplatesRequest) SetCreatedBy(v string) *ListUserTemplatesRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListUserTemplatesRequest) SetCreatedDateAfter(v string) *ListUserTemplatesRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListUserTemplatesRequest) SetCreatedDateBefore(v string) *ListUserTemplatesRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListUserTemplatesRequest) SetMaxResults(v int32) *ListUserTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserTemplatesRequest) SetNextToken(v string) *ListUserTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserTemplatesRequest) SetPopularity(v int32) *ListUserTemplatesRequest {
	s.Popularity = &v
	return s
}

func (s *ListUserTemplatesRequest) SetRegionId(v string) *ListUserTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserTemplatesRequest) SetShareType(v string) *ListUserTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListUserTemplatesRequest) SetSortField(v string) *ListUserTemplatesRequest {
	s.SortField = &v
	return s
}

func (s *ListUserTemplatesRequest) SetSortOrder(v string) *ListUserTemplatesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserTemplatesRequest) SetTemplateFormat(v string) *ListUserTemplatesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListUserTemplatesRequest) SetTemplateName(v string) *ListUserTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListUserTemplatesRequest) SetTemplateType(v string) *ListUserTemplatesRequest {
	s.TemplateType = &v
	return s
}

type ListUserTemplatesResponseBody struct {
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates  []*ListUserTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListUserTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserTemplatesResponseBody) SetMaxResults(v int32) *ListUserTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserTemplatesResponseBody) SetNextToken(v string) *ListUserTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserTemplatesResponseBody) SetRequestId(v string) *ListUserTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserTemplatesResponseBody) SetTemplates(v []*ListUserTemplatesResponseBodyTemplates) *ListUserTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListUserTemplatesResponseBodyTemplates struct {
	CreatedBy           *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate         *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash                *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Popularity          *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	ShareType           *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat      *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion     *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	TotalExecutionCount *int32  `json:"TotalExecutionCount,omitempty" xml:"TotalExecutionCount,omitempty"`
	UpdatedBy           *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate         *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListUserTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListUserTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListUserTemplatesResponseBodyTemplates) SetCreatedBy(v string) *ListUserTemplatesResponseBodyTemplates {
	s.CreatedBy = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetCreatedDate(v string) *ListUserTemplatesResponseBodyTemplates {
	s.CreatedDate = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetDescription(v string) *ListUserTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetHash(v string) *ListUserTemplatesResponseBodyTemplates {
	s.Hash = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetPopularity(v int32) *ListUserTemplatesResponseBodyTemplates {
	s.Popularity = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetShareType(v string) *ListUserTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetTemplateFormat(v string) *ListUserTemplatesResponseBodyTemplates {
	s.TemplateFormat = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListUserTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListUserTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListUserTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetTotalExecutionCount(v int32) *ListUserTemplatesResponseBodyTemplates {
	s.TotalExecutionCount = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetUpdatedBy(v string) *ListUserTemplatesResponseBodyTemplates {
	s.UpdatedBy = &v
	return s
}

func (s *ListUserTemplatesResponseBodyTemplates) SetUpdatedDate(v string) *ListUserTemplatesResponseBodyTemplates {
	s.UpdatedDate = &v
	return s
}

type ListUserTemplatesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListUserTemplatesResponse) SetHeaders(v map[string]*string) *ListUserTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListUserTemplatesResponse) SetStatusCode(v int32) *ListUserTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserTemplatesResponse) SetBody(v *ListUserTemplatesResponseBody) *ListUserTemplatesResponse {
	s.Body = v
	return s
}

type ResetTimerTriggerExecutionRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetTimerTriggerExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetTimerTriggerExecutionRequest) GoString() string {
	return s.String()
}

func (s *ResetTimerTriggerExecutionRequest) SetAliUid(v string) *ResetTimerTriggerExecutionRequest {
	s.AliUid = &v
	return s
}

func (s *ResetTimerTriggerExecutionRequest) SetExecutionId(v string) *ResetTimerTriggerExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *ResetTimerTriggerExecutionRequest) SetRegionId(v string) *ResetTimerTriggerExecutionRequest {
	s.RegionId = &v
	return s
}

type ResetTimerTriggerExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetTimerTriggerExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetTimerTriggerExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *ResetTimerTriggerExecutionResponseBody) SetRequestId(v string) *ResetTimerTriggerExecutionResponseBody {
	s.RequestId = &v
	return s
}

type ResetTimerTriggerExecutionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetTimerTriggerExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetTimerTriggerExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetTimerTriggerExecutionResponse) GoString() string {
	return s.String()
}

func (s *ResetTimerTriggerExecutionResponse) SetHeaders(v map[string]*string) *ResetTimerTriggerExecutionResponse {
	s.Headers = v
	return s
}

func (s *ResetTimerTriggerExecutionResponse) SetStatusCode(v int32) *ResetTimerTriggerExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetTimerTriggerExecutionResponse) SetBody(v *ResetTimerTriggerExecutionResponseBody) *ResetTimerTriggerExecutionResponse {
	s.Body = v
	return s
}

type ResetUserExecutionRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResetUserExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetUserExecutionRequest) GoString() string {
	return s.String()
}

func (s *ResetUserExecutionRequest) SetAliUid(v string) *ResetUserExecutionRequest {
	s.AliUid = &v
	return s
}

func (s *ResetUserExecutionRequest) SetExecutionId(v string) *ResetUserExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *ResetUserExecutionRequest) SetRegionId(v string) *ResetUserExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *ResetUserExecutionRequest) SetStatus(v string) *ResetUserExecutionRequest {
	s.Status = &v
	return s
}

type ResetUserExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetUserExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetUserExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserExecutionResponseBody) SetRequestId(v string) *ResetUserExecutionResponseBody {
	s.RequestId = &v
	return s
}

type ResetUserExecutionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetUserExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetUserExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetUserExecutionResponse) GoString() string {
	return s.String()
}

func (s *ResetUserExecutionResponse) SetHeaders(v map[string]*string) *ResetUserExecutionResponse {
	s.Headers = v
	return s
}

func (s *ResetUserExecutionResponse) SetStatusCode(v int32) *ResetUserExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetUserExecutionResponse) SetBody(v *ResetUserExecutionResponseBody) *ResetUserExecutionResponse {
	s.Body = v
	return s
}

type SetFlowControlRequest struct {
	Api     *string `json:"Api,omitempty" xml:"Api,omitempty"`
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// This parameter is required.
	Type *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid  *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// This parameter is required.
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFlowControlRequest) GoString() string {
	return s.String()
}

func (s *SetFlowControlRequest) SetApi(v string) *SetFlowControlRequest {
	s.Api = &v
	return s
}

func (s *SetFlowControlRequest) SetService(v string) *SetFlowControlRequest {
	s.Service = &v
	return s
}

func (s *SetFlowControlRequest) SetType(v int32) *SetFlowControlRequest {
	s.Type = &v
	return s
}

func (s *SetFlowControlRequest) SetUid(v string) *SetFlowControlRequest {
	s.Uid = &v
	return s
}

func (s *SetFlowControlRequest) SetValue(v int32) *SetFlowControlRequest {
	s.Value = &v
	return s
}

type SetFlowControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *SetFlowControlResponseBody) SetRequestId(v string) *SetFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type SetFlowControlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFlowControlResponse) GoString() string {
	return s.String()
}

func (s *SetFlowControlResponse) SetHeaders(v map[string]*string) *SetFlowControlResponse {
	s.Headers = v
	return s
}

func (s *SetFlowControlResponse) SetStatusCode(v int32) *SetFlowControlResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFlowControlResponse) SetBody(v *SetFlowControlResponseBody) *SetFlowControlResponse {
	s.Body = v
	return s
}

type SetQuotaRequest struct {
	// This parameter is required.
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetQuotaRequest) SetQuotaName(v string) *SetQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *SetQuotaRequest) SetRegionId(v string) *SetQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *SetQuotaRequest) SetUid(v string) *SetQuotaRequest {
	s.Uid = &v
	return s
}

func (s *SetQuotaRequest) SetValue(v string) *SetQuotaRequest {
	s.Value = &v
	return s
}

type SetQuotaResponseBody struct {
	Quota     *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetQuotaResponseBody) SetQuota(v int32) *SetQuotaResponseBody {
	s.Quota = &v
	return s
}

func (s *SetQuotaResponseBody) SetRequestId(v string) *SetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetQuotaResponseBody) SetUid(v string) *SetQuotaResponseBody {
	s.Uid = &v
	return s
}

type SetQuotaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetQuotaResponse) SetHeaders(v map[string]*string) *SetQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetQuotaResponse) SetStatusCode(v int32) *SetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetQuotaResponse) SetBody(v *SetQuotaResponseBody) *SetQuotaResponse {
	s.Body = v
	return s
}

type TerminateUserExecutionRequest struct {
	// This parameter is required.
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s TerminateUserExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateUserExecutionRequest) GoString() string {
	return s.String()
}

func (s *TerminateUserExecutionRequest) SetAliUid(v string) *TerminateUserExecutionRequest {
	s.AliUid = &v
	return s
}

func (s *TerminateUserExecutionRequest) SetExecutionId(v string) *TerminateUserExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *TerminateUserExecutionRequest) SetRegionId(v string) *TerminateUserExecutionRequest {
	s.RegionId = &v
	return s
}

type TerminateUserExecutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateUserExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateUserExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateUserExecutionResponseBody) SetRequestId(v string) *TerminateUserExecutionResponseBody {
	s.RequestId = &v
	return s
}

type TerminateUserExecutionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateUserExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateUserExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateUserExecutionResponse) GoString() string {
	return s.String()
}

func (s *TerminateUserExecutionResponse) SetHeaders(v map[string]*string) *TerminateUserExecutionResponse {
	s.Headers = v
	return s
}

func (s *TerminateUserExecutionResponse) SetStatusCode(v int32) *TerminateUserExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateUserExecutionResponse) SetBody(v *TerminateUserExecutionResponseBody) *TerminateUserExecutionResponse {
	s.Body = v
	return s
}

type UpdateActionRequest struct {
	// This parameter is required.
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// This parameter is required.
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// This parameter is required.
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Popularity *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateActionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionRequest) GoString() string {
	return s.String()
}

func (s *UpdateActionRequest) SetActionName(v string) *UpdateActionRequest {
	s.ActionName = &v
	return s
}

func (s *UpdateActionRequest) SetActionType(v string) *UpdateActionRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateActionRequest) SetContent(v string) *UpdateActionRequest {
	s.Content = &v
	return s
}

func (s *UpdateActionRequest) SetPopularity(v int32) *UpdateActionRequest {
	s.Popularity = &v
	return s
}

func (s *UpdateActionRequest) SetRegionId(v string) *UpdateActionRequest {
	s.RegionId = &v
	return s
}

type UpdateActionResponseBody struct {
	ActionName      *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ActionType      *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateActionResponseBody) SetActionName(v string) *UpdateActionResponseBody {
	s.ActionName = &v
	return s
}

func (s *UpdateActionResponseBody) SetActionType(v string) *UpdateActionResponseBody {
	s.ActionType = &v
	return s
}

func (s *UpdateActionResponseBody) SetCreatedDate(v string) *UpdateActionResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *UpdateActionResponseBody) SetDescription(v string) *UpdateActionResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateActionResponseBody) SetPopularity(v int32) *UpdateActionResponseBody {
	s.Popularity = &v
	return s
}

func (s *UpdateActionResponseBody) SetProperties(v string) *UpdateActionResponseBody {
	s.Properties = &v
	return s
}

func (s *UpdateActionResponseBody) SetRequestId(v string) *UpdateActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateActionResponseBody) SetTemplateVersion(v string) *UpdateActionResponseBody {
	s.TemplateVersion = &v
	return s
}

type UpdateActionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateActionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateActionResponse) GoString() string {
	return s.String()
}

func (s *UpdateActionResponse) SetHeaders(v map[string]*string) *UpdateActionResponse {
	s.Headers = v
	return s
}

func (s *UpdateActionResponse) SetStatusCode(v int32) *UpdateActionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateActionResponse) SetBody(v *UpdateActionResponseBody) *UpdateActionResponse {
	s.Body = v
	return s
}

type UpdatePublicParameterRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePublicParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicParameterRequest) SetDescription(v string) *UpdatePublicParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdatePublicParameterRequest) SetName(v string) *UpdatePublicParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdatePublicParameterRequest) SetRegionId(v string) *UpdatePublicParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePublicParameterRequest) SetValue(v string) *UpdatePublicParameterRequest {
	s.Value = &v
	return s
}

type UpdatePublicParameterResponseBody struct {
	Parameter *UpdatePublicParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublicParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicParameterResponseBody) SetParameter(v *UpdatePublicParameterResponseBodyParameter) *UpdatePublicParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *UpdatePublicParameterResponseBody) SetRequestId(v string) *UpdatePublicParameterResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePublicParameterResponseBodyParameter struct {
	Constraints      *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedBy        *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate      *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterVersion *int32  `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShareType        *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdatedBy        *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate      *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdatePublicParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdatePublicParameterResponseBodyParameter) SetConstraints(v string) *UpdatePublicParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetCreatedBy(v string) *UpdatePublicParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetCreatedDate(v string) *UpdatePublicParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetDescription(v string) *UpdatePublicParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetId(v string) *UpdatePublicParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetName(v string) *UpdatePublicParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdatePublicParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetRegionId(v string) *UpdatePublicParameterResponseBodyParameter {
	s.RegionId = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetShareType(v string) *UpdatePublicParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetType(v string) *UpdatePublicParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdatePublicParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdatePublicParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdatePublicParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

type UpdatePublicParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicParameterResponse) SetHeaders(v map[string]*string) *UpdatePublicParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicParameterResponse) SetStatusCode(v int32) *UpdatePublicParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicParameterResponse) SetBody(v *UpdatePublicParameterResponseBody) *UpdatePublicParameterResponse {
	s.Body = v
	return s
}

type UpdatePublicPatchBaselineRequest struct {
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdatePublicPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicPatchBaselineRequest) SetApprovalRules(v string) *UpdatePublicPatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePublicPatchBaselineRequest) SetClientToken(v string) *UpdatePublicPatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePublicPatchBaselineRequest) SetDescription(v string) *UpdatePublicPatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *UpdatePublicPatchBaselineRequest) SetName(v string) *UpdatePublicPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePublicPatchBaselineRequest) SetRegionId(v string) *UpdatePublicPatchBaselineRequest {
	s.RegionId = &v
	return s
}

type UpdatePublicPatchBaselineResponseBody struct {
	PatchBaseline *UpdatePublicPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublicPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicPatchBaselineResponseBody) SetPatchBaseline(v *UpdatePublicPatchBaselineResponseBodyPatchBaseline) *UpdatePublicPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBody) SetRequestId(v string) *UpdatePublicPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePublicPatchBaselineResponseBodyPatchBaseline struct {
	ApprovalRules   *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdatePublicPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetId(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetName(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *UpdatePublicPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type UpdatePublicPatchBaselineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicPatchBaselineResponse) SetHeaders(v map[string]*string) *UpdatePublicPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicPatchBaselineResponse) SetStatusCode(v int32) *UpdatePublicPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicPatchBaselineResponse) SetBody(v *UpdatePublicPatchBaselineResponseBody) *UpdatePublicPatchBaselineResponse {
	s.Body = v
	return s
}

type UpdatePublicTemplateRequest struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Popularity *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	Publisher  *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdatePublicTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicTemplateRequest) SetCategory(v string) *UpdatePublicTemplateRequest {
	s.Category = &v
	return s
}

func (s *UpdatePublicTemplateRequest) SetContent(v string) *UpdatePublicTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdatePublicTemplateRequest) SetPopularity(v int32) *UpdatePublicTemplateRequest {
	s.Popularity = &v
	return s
}

func (s *UpdatePublicTemplateRequest) SetPublisher(v string) *UpdatePublicTemplateRequest {
	s.Publisher = &v
	return s
}

func (s *UpdatePublicTemplateRequest) SetRegionId(v string) *UpdatePublicTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePublicTemplateRequest) SetTemplateName(v string) *UpdatePublicTemplateRequest {
	s.TemplateName = &v
	return s
}

type UpdatePublicTemplateResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *UpdatePublicTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s UpdatePublicTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicTemplateResponseBody) SetRequestId(v string) *UpdatePublicTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicTemplateResponseBody) SetTemplate(v *UpdatePublicTemplateResponseBodyTemplate) *UpdatePublicTemplateResponseBody {
	s.Template = v
	return s
}

type UpdatePublicTemplateResponseBodyTemplate struct {
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreatedBy       *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	CreatedDate     *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Hash            *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Popularity      *int32  `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	ShareType       *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	TemplateFormat  *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdatedBy       *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	UpdatedDate     *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdatePublicTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetCategory(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.Category = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetCreatedBy(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetCreatedDate(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetDescription(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetHash(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetPopularity(v int32) *UpdatePublicTemplateResponseBodyTemplate {
	s.Popularity = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetShareType(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetTemplateFormat(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetTemplateId(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetTemplateName(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetTemplateVersion(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetUpdatedBy(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *UpdatePublicTemplateResponseBodyTemplate) SetUpdatedDate(v string) *UpdatePublicTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type UpdatePublicTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePublicTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePublicTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePublicTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdatePublicTemplateResponse) SetHeaders(v map[string]*string) *UpdatePublicTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdatePublicTemplateResponse) SetStatusCode(v int32) *UpdatePublicTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePublicTemplateResponse) SetBody(v *UpdatePublicTemplateResponseBody) *UpdatePublicTemplateResponse {
	s.Body = v
	return s
}

type ValidatePublicTemplateContentRequest struct {
	// This parameter is required.
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValidatePublicTemplateContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidatePublicTemplateContentRequest) GoString() string {
	return s.String()
}

func (s *ValidatePublicTemplateContentRequest) SetContent(v string) *ValidatePublicTemplateContentRequest {
	s.Content = &v
	return s
}

func (s *ValidatePublicTemplateContentRequest) SetRegionId(v string) *ValidatePublicTemplateContentRequest {
	s.RegionId = &v
	return s
}

func (s *ValidatePublicTemplateContentRequest) SetTemplateName(v string) *ValidatePublicTemplateContentRequest {
	s.TemplateName = &v
	return s
}

func (s *ValidatePublicTemplateContentRequest) SetType(v string) *ValidatePublicTemplateContentRequest {
	s.Type = &v
	return s
}

type ValidatePublicTemplateContentResponseBody struct {
	Description *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Outputs     *string                                           `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Parameters  *string                                           `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RamRole     *string                                           `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks       []*ValidatePublicTemplateContentResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ValidatePublicTemplateContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidatePublicTemplateContentResponseBody) GoString() string {
	return s.String()
}

func (s *ValidatePublicTemplateContentResponseBody) SetDescription(v string) *ValidatePublicTemplateContentResponseBody {
	s.Description = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBody) SetOutputs(v string) *ValidatePublicTemplateContentResponseBody {
	s.Outputs = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBody) SetParameters(v string) *ValidatePublicTemplateContentResponseBody {
	s.Parameters = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBody) SetRamRole(v string) *ValidatePublicTemplateContentResponseBody {
	s.RamRole = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBody) SetRequestId(v string) *ValidatePublicTemplateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBody) SetTasks(v []*ValidatePublicTemplateContentResponseBodyTasks) *ValidatePublicTemplateContentResponseBody {
	s.Tasks = v
	return s
}

type ValidatePublicTemplateContentResponseBodyTasks struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Outputs     *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Properties  *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValidatePublicTemplateContentResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ValidatePublicTemplateContentResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ValidatePublicTemplateContentResponseBodyTasks) SetDescription(v string) *ValidatePublicTemplateContentResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBodyTasks) SetName(v string) *ValidatePublicTemplateContentResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBodyTasks) SetOutputs(v string) *ValidatePublicTemplateContentResponseBodyTasks {
	s.Outputs = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBodyTasks) SetProperties(v string) *ValidatePublicTemplateContentResponseBodyTasks {
	s.Properties = &v
	return s
}

func (s *ValidatePublicTemplateContentResponseBodyTasks) SetType(v string) *ValidatePublicTemplateContentResponseBodyTasks {
	s.Type = &v
	return s
}

type ValidatePublicTemplateContentResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidatePublicTemplateContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidatePublicTemplateContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidatePublicTemplateContentResponse) GoString() string {
	return s.String()
}

func (s *ValidatePublicTemplateContentResponse) SetHeaders(v map[string]*string) *ValidatePublicTemplateContentResponse {
	s.Headers = v
	return s
}

func (s *ValidatePublicTemplateContentResponse) SetStatusCode(v int32) *ValidatePublicTemplateContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidatePublicTemplateContentResponse) SetBody(v *ValidatePublicTemplateContentResponseBody) *ValidatePublicTemplateContentResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("oosops"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AuditPublicTemplateRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditPublicTemplateRegistrationResponse
func (client *Client) AuditPublicTemplateRegistrationWithOptions(request *AuditPublicTemplateRegistrationRequest, runtime *util.RuntimeOptions) (_result *AuditPublicTemplateRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditAction)) {
		query["AuditAction"] = request.AuditAction
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationId)) {
		query["RegistrationId"] = request.RegistrationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuditPublicTemplateRegistration"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuditPublicTemplateRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuditPublicTemplateRegistrationRequest
//
// @return AuditPublicTemplateRegistrationResponse
func (client *Client) AuditPublicTemplateRegistration(request *AuditPublicTemplateRegistrationRequest) (_result *AuditPublicTemplateRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuditPublicTemplateRegistrationResponse{}
	_body, _err := client.AuditPublicTemplateRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateActionResponse
func (client *Client) CreateActionWithOptions(request *CreateActionRequest, runtime *util.RuntimeOptions) (_result *CreateActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Popularity)) {
		query["Popularity"] = request.Popularity
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAction"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateActionRequest
//
// @return CreateActionResponse
func (client *Client) CreateAction(request *CreateActionRequest) (_result *CreateActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateActionResponse{}
	_body, _err := client.CreateActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个公共参数。
//
// @param request - CreatePublicParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublicParameterResponse
func (client *Client) CreatePublicParameterWithOptions(request *CreatePublicParameterRequest, runtime *util.RuntimeOptions) (_result *CreatePublicParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Constraints)) {
		query["Constraints"] = request.Constraints
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterType)) {
		query["ParameterType"] = request.ParameterType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePublicParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePublicParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个公共参数。
//
// @param request - CreatePublicParameterRequest
//
// @return CreatePublicParameterResponse
func (client *Client) CreatePublicParameter(request *CreatePublicParameterRequest) (_result *CreatePublicParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePublicParameterResponse{}
	_body, _err := client.CreatePublicParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreatePublicPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublicPatchBaselineResponse
func (client *Client) CreatePublicPatchBaselineWithOptions(request *CreatePublicPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *CreatePublicPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalRules)) {
		query["ApprovalRules"] = request.ApprovalRules
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperationSystem)) {
		query["OperationSystem"] = request.OperationSystem
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePublicPatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePublicPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreatePublicPatchBaselineRequest
//
// @return CreatePublicPatchBaselineResponse
func (client *Client) CreatePublicPatchBaseline(request *CreatePublicPatchBaselineRequest) (_result *CreatePublicPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePublicPatchBaselineResponse{}
	_body, _err := client.CreatePublicPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreatePublicTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublicTemplateResponse
func (client *Client) CreatePublicTemplateWithOptions(request *CreatePublicTemplateRequest, runtime *util.RuntimeOptions) (_result *CreatePublicTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.IsExample)) {
		query["IsExample"] = request.IsExample
	}

	if !tea.BoolValue(util.IsUnset(request.Popularity)) {
		query["Popularity"] = request.Popularity
	}

	if !tea.BoolValue(util.IsUnset(request.Publisher)) {
		query["Publisher"] = request.Publisher
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePublicTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePublicTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreatePublicTemplateRequest
//
// @return CreatePublicTemplateResponse
func (client *Client) CreatePublicTemplate(request *CreatePublicTemplateRequest) (_result *CreatePublicTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePublicTemplateResponse{}
	_body, _err := client.CreatePublicTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFailureMsgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFailureMsgResponse
func (client *Client) DeleteFailureMsgWithOptions(request *DeleteFailureMsgRequest, runtime *util.RuntimeOptions) (_result *DeleteFailureMsgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.RequestFingerprint)) {
		query["RequestFingerprint"] = request.RequestFingerprint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFailureMsg"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFailureMsgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFailureMsgRequest
//
// @return DeleteFailureMsgResponse
func (client *Client) DeleteFailureMsg(request *DeleteFailureMsgRequest) (_result *DeleteFailureMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFailureMsgResponse{}
	_body, _err := client.DeleteFailureMsgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除公共参数。
//
// @param request - DeletePublicParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicParameterResponse
func (client *Client) DeletePublicParameterWithOptions(request *DeletePublicParameterRequest, runtime *util.RuntimeOptions) (_result *DeletePublicParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePublicParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePublicParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除公共参数。
//
// @param request - DeletePublicParameterRequest
//
// @return DeletePublicParameterResponse
func (client *Client) DeletePublicParameter(request *DeletePublicParameterRequest) (_result *DeletePublicParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePublicParameterResponse{}
	_body, _err := client.DeletePublicParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeletePublicPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicPatchBaselineResponse
func (client *Client) DeletePublicPatchBaselineWithOptions(request *DeletePublicPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *DeletePublicPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePublicPatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePublicPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePublicPatchBaselineRequest
//
// @return DeletePublicPatchBaselineResponse
func (client *Client) DeletePublicPatchBaseline(request *DeletePublicPatchBaselineRequest) (_result *DeletePublicPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePublicPatchBaselineResponse{}
	_body, _err := client.DeletePublicPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeletePublicTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePublicTemplateResponse
func (client *Client) DeletePublicTemplateWithOptions(request *DeletePublicTemplateRequest, runtime *util.RuntimeOptions) (_result *DeletePublicTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePublicTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePublicTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePublicTemplateRequest
//
// @return DeletePublicTemplateResponse
func (client *Client) DeletePublicTemplate(request *DeletePublicTemplateRequest) (_result *DeletePublicTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePublicTemplateResponse{}
	_body, _err := client.DeletePublicTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DoCheckResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DoCheckResourceResponse
func (client *Client) DoCheckResourceWithOptions(request *DoCheckResourceRequest, runtime *util.RuntimeOptions) (_result *DoCheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.GmtWakeup)) {
		query["gmtWakeup"] = request.GmtWakeup
	}

	if !tea.BoolValue(util.IsUnset(request.Hid)) {
		query["hid"] = request.Hid
	}

	if !tea.BoolValue(util.IsUnset(request.Interrupt)) {
		query["interrupt"] = request.Interrupt
	}

	if !tea.BoolValue(util.IsUnset(request.Invoker)) {
		query["invoker"] = request.Invoker
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Success)) {
		query["success"] = request.Success
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExtraData)) {
		query["taskExtraData"] = request.TaskExtraData
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdentifier)) {
		query["taskIdentifier"] = request.TaskIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DoCheckResource"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoCheckResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DoCheckResourceRequest
//
// @return DoCheckResourceResponse
func (client *Client) DoCheckResource(request *DoCheckResourceRequest) (_result *DoCheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoCheckResourceResponse{}
	_body, _err := client.DoCheckResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取action的详细信息
//
// @param request - GetActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActionResponse
func (client *Client) GetActionWithOptions(request *GetActionRequest, runtime *util.RuntimeOptions) (_result *GetActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAction"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取action的详细信息
//
// @param request - GetActionRequest
//
// @return GetActionResponse
func (client *Client) GetAction(request *GetActionRequest) (_result *GetActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetActionResponse{}
	_body, _err := client.GetActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetFlowControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowControlResponse
func (client *Client) GetFlowControlWithOptions(request *GetFlowControlRequest, runtime *util.RuntimeOptions) (_result *GetFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Api)) {
		query["Api"] = request.Api
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFlowControl"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFlowControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetFlowControlRequest
//
// @return GetFlowControlResponse
func (client *Client) GetFlowControl(request *GetFlowControlRequest) (_result *GetFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowControlResponse{}
	_body, _err := client.GetFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一个公共参数，包括参数值。
//
// @param request - GetPublicParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicParameterResponse
func (client *Client) GetPublicParameterWithOptions(request *GetPublicParameterRequest, runtime *util.RuntimeOptions) (_result *GetPublicParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterVersion)) {
		query["ParameterVersion"] = request.ParameterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublicParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublicParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一个公共参数，包括参数值。
//
// @param request - GetPublicParameterRequest
//
// @return GetPublicParameterResponse
func (client *Client) GetPublicParameter(request *GetPublicParameterRequest) (_result *GetPublicParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicParameterResponse{}
	_body, _err := client.GetPublicParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetPublicPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicPatchBaselineResponse
func (client *Client) GetPublicPatchBaselineWithOptions(request *GetPublicPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *GetPublicPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublicPatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublicPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPublicPatchBaselineRequest
//
// @return GetPublicPatchBaselineResponse
func (client *Client) GetPublicPatchBaseline(request *GetPublicPatchBaselineRequest) (_result *GetPublicPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicPatchBaselineResponse{}
	_body, _err := client.GetPublicPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetPublicTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicTemplateResponse
func (client *Client) GetPublicTemplateWithOptions(request *GetPublicTemplateRequest, runtime *util.RuntimeOptions) (_result *GetPublicTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublicTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublicTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPublicTemplateRequest
//
// @return GetPublicTemplateResponse
func (client *Client) GetPublicTemplate(request *GetPublicTemplateRequest) (_result *GetPublicTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicTemplateResponse{}
	_body, _err := client.GetPublicTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaResponse
func (client *Client) GetQuotaWithOptions(request *GetQuotaRequest, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		query["QuotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuota"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQuotaRequest
//
// @return GetQuotaResponse
func (client *Client) GetQuota(request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUserExecutionTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserExecutionTemplateResponse
func (client *Client) GetUserExecutionTemplateWithOptions(request *GetUserExecutionTemplateRequest, runtime *util.RuntimeOptions) (_result *GetUserExecutionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserExecutionTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserExecutionTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserExecutionTemplateRequest
//
// @return GetUserExecutionTemplateResponse
func (client *Client) GetUserExecutionTemplate(request *GetUserExecutionTemplateRequest) (_result *GetUserExecutionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserExecutionTemplateResponse{}
	_body, _err := client.GetUserExecutionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUserTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserTemplateResponse
func (client *Client) GetUserTemplateWithOptions(request *GetUserTemplateRequest, runtime *util.RuntimeOptions) (_result *GetUserTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserTemplateRequest
//
// @return GetUserTemplateResponse
func (client *Client) GetUserTemplate(request *GetUserTemplateRequest) (_result *GetUserTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserTemplateResponse{}
	_body, _err := client.GetUserTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionsResponse
func (client *Client) ListActionsWithOptions(request *ListActionsRequest, runtime *util.RuntimeOptions) (_result *ListActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OOSActionName)) {
		query["OOSActionName"] = request.OOSActionName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListActions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListActionsRequest
//
// @return ListActionsResponse
func (client *Client) ListActions(request *ListActionsRequest) (_result *ListActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActionsResponse{}
	_body, _err := client.ListActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDefaultQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDefaultQuotaResponse
func (client *Client) ListDefaultQuotaWithOptions(runtime *util.RuntimeOptions) (_result *ListDefaultQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListDefaultQuota"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDefaultQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return ListDefaultQuotaResponse
func (client *Client) ListDefaultQuota() (_result *ListDefaultQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDefaultQuotaResponse{}
	_body, _err := client.ListDefaultQuotaWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListFailureMsgsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFailureMsgsResponse
func (client *Client) ListFailureMsgsWithOptions(request *ListFailureMsgsRequest, runtime *util.RuntimeOptions) (_result *ListFailureMsgsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RequestFingerprint)) {
		query["RequestFingerprint"] = request.RequestFingerprint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFailureMsgs"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFailureMsgsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFailureMsgsRequest
//
// @return ListFailureMsgsResponse
func (client *Client) ListFailureMsgs(request *ListFailureMsgsRequest) (_result *ListFailureMsgsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFailureMsgsResponse{}
	_body, _err := client.ListFailureMsgsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListOOSLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOOSLogsResponse
func (client *Client) ListOOSLogsWithOptions(request *ListOOSLogsRequest, runtime *util.RuntimeOptions) (_result *ListOOSLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestFingerprint)) {
		query["RequestFingerprint"] = request.RequestFingerprint
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOOSLogs"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOOSLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListOOSLogsRequest
//
// @return ListOOSLogsResponse
func (client *Client) ListOOSLogs(request *ListOOSLogsRequest) (_result *ListOOSLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOOSLogsResponse{}
	_body, _err := client.ListOOSLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询公共参数。支持多种查询
//
// @param request - ListPublicParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicParametersResponse
func (client *Client) ListPublicParametersWithOptions(request *ListPublicParametersRequest, runtime *util.RuntimeOptions) (_result *ListPublicParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterType)) {
		query["ParameterType"] = request.ParameterType
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicParameters"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询公共参数。支持多种查询
//
// @param request - ListPublicParametersRequest
//
// @return ListPublicParametersResponse
func (client *Client) ListPublicParameters(request *ListPublicParametersRequest) (_result *ListPublicParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicParametersResponse{}
	_body, _err := client.ListPublicParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPublicPatchBaselinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicPatchBaselinesResponse
func (client *Client) ListPublicPatchBaselinesWithOptions(request *ListPublicPatchBaselinesRequest, runtime *util.RuntimeOptions) (_result *ListPublicPatchBaselinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperationSystem)) {
		query["OperationSystem"] = request.OperationSystem
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicPatchBaselines"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicPatchBaselinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPublicPatchBaselinesRequest
//
// @return ListPublicPatchBaselinesResponse
func (client *Client) ListPublicPatchBaselines(request *ListPublicPatchBaselinesRequest) (_result *ListPublicPatchBaselinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicPatchBaselinesResponse{}
	_body, _err := client.ListPublicPatchBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPublicTemplateRegistrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicTemplateRegistrationsResponse
func (client *Client) ListPublicTemplateRegistrationsWithOptions(request *ListPublicTemplateRegistrationsRequest, runtime *util.RuntimeOptions) (_result *ListPublicTemplateRegistrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationId)) {
		query["RegistrationId"] = request.RegistrationId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicTemplateRegistrations"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicTemplateRegistrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPublicTemplateRegistrationsRequest
//
// @return ListPublicTemplateRegistrationsResponse
func (client *Client) ListPublicTemplateRegistrations(request *ListPublicTemplateRegistrationsRequest) (_result *ListPublicTemplateRegistrationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicTemplateRegistrationsResponse{}
	_body, _err := client.ListPublicTemplateRegistrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPublicTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicTemplatesResponse
func (client *Client) ListPublicTemplatesWithOptions(request *ListPublicTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListPublicTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatedBy)) {
		query["CreatedBy"] = request.CreatedBy
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedDateAfter)) {
		query["CreatedDateAfter"] = request.CreatedDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedDateBefore)) {
		query["CreatedDateBefore"] = request.CreatedDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.IsExample)) {
		query["IsExample"] = request.IsExample
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Popularity)) {
		query["Popularity"] = request.Popularity
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateFormat)) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublicTemplates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublicTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPublicTemplatesRequest
//
// @return ListPublicTemplatesResponse
func (client *Client) ListPublicTemplates(request *ListPublicTemplatesRequest) (_result *ListPublicTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicTemplatesResponse{}
	_body, _err := client.ListPublicTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserExecutionLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserExecutionLogsResponse
func (client *Client) ListUserExecutionLogsWithOptions(request *ListUserExecutionLogsRequest, runtime *util.RuntimeOptions) (_result *ListUserExecutionLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutionId)) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserExecutionLogs"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserExecutionLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserExecutionLogsRequest
//
// @return ListUserExecutionLogsResponse
func (client *Client) ListUserExecutionLogs(request *ListUserExecutionLogsRequest) (_result *ListUserExecutionLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserExecutionLogsResponse{}
	_body, _err := client.ListUserExecutionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserExecutionsResponse
func (client *Client) ListUserExecutionsWithOptions(request *ListUserExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListUserExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateAfter)) {
		query["EndDateAfter"] = request.EndDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateBefore)) {
		query["EndDateBefore"] = request.EndDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutedBy)) {
		query["ExecutedBy"] = request.ExecutedBy
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeChildExecution)) {
		query["IncludeChildExecution"] = request.IncludeChildExecution
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ParentExecutionId)) {
		query["ParentExecutionId"] = request.ParentExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		query["RamRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateAfter)) {
		query["StartDateAfter"] = request.StartDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateBefore)) {
		query["StartDateBefore"] = request.StartDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserExecutions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserExecutionsRequest
//
// @return ListUserExecutionsResponse
func (client *Client) ListUserExecutions(request *ListUserExecutionsRequest) (_result *ListUserExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserExecutionsResponse{}
	_body, _err := client.ListUserExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserInstancePatchStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserInstancePatchStatesResponse
func (client *Client) ListUserInstancePatchStatesWithOptions(request *ListUserInstancePatchStatesRequest, runtime *util.RuntimeOptions) (_result *ListUserInstancePatchStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserInstancePatchStates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserInstancePatchStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserInstancePatchStatesRequest
//
// @return ListUserInstancePatchStatesResponse
func (client *Client) ListUserInstancePatchStates(request *ListUserInstancePatchStatesRequest) (_result *ListUserInstancePatchStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserInstancePatchStatesResponse{}
	_body, _err := client.ListUserInstancePatchStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserInstancePatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserInstancePatchesResponse
func (client *Client) ListUserInstancePatchesWithOptions(request *ListUserInstancePatchesRequest, runtime *util.RuntimeOptions) (_result *ListUserInstancePatchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserInstancePatches"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserInstancePatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserInstancePatchesRequest
//
// @return ListUserInstancePatchesResponse
func (client *Client) ListUserInstancePatches(request *ListUserInstancePatchesRequest) (_result *ListUserInstancePatchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserInstancePatchesResponse{}
	_body, _err := client.ListUserInstancePatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserInventoryEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserInventoryEntriesResponse
func (client *Client) ListUserInventoryEntriesWithOptions(request *ListUserInventoryEntriesRequest, runtime *util.RuntimeOptions) (_result *ListUserInventoryEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TypeName)) {
		query["TypeName"] = request.TypeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserInventoryEntries"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserInventoryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserInventoryEntriesRequest
//
// @return ListUserInventoryEntriesResponse
func (client *Client) ListUserInventoryEntries(request *ListUserInventoryEntriesRequest) (_result *ListUserInventoryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserInventoryEntriesResponse{}
	_body, _err := client.ListUserInventoryEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserTaskExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserTaskExecutionsResponse
func (client *Client) ListUserTaskExecutionsWithOptions(request *ListUserTaskExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListUserTaskExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateAfter)) {
		query["EndDateAfter"] = request.EndDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateBefore)) {
		query["EndDateBefore"] = request.EndDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeChildTaskExecution)) {
		query["IncludeChildTaskExecution"] = request.IncludeChildTaskExecution
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskExecutionId)) {
		query["ParentTaskExecutionId"] = request.ParentTaskExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateAfter)) {
		query["StartDateAfter"] = request.StartDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateBefore)) {
		query["StartDateBefore"] = request.StartDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAction)) {
		query["TaskAction"] = request.TaskAction
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutionId)) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserTaskExecutions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserTaskExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserTaskExecutionsRequest
//
// @return ListUserTaskExecutionsResponse
func (client *Client) ListUserTaskExecutions(request *ListUserTaskExecutionsRequest) (_result *ListUserTaskExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserTaskExecutionsResponse{}
	_body, _err := client.ListUserTaskExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserTemplatesResponse
func (client *Client) ListUserTemplatesWithOptions(request *ListUserTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListUserTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedBy)) {
		query["CreatedBy"] = request.CreatedBy
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedDateAfter)) {
		query["CreatedDateAfter"] = request.CreatedDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedDateBefore)) {
		query["CreatedDateBefore"] = request.CreatedDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Popularity)) {
		query["Popularity"] = request.Popularity
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateFormat)) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserTemplates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserTemplatesRequest
//
// @return ListUserTemplatesResponse
func (client *Client) ListUserTemplates(request *ListUserTemplatesRequest) (_result *ListUserTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserTemplatesResponse{}
	_body, _err := client.ListUserTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetTimerTriggerExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetTimerTriggerExecutionResponse
func (client *Client) ResetTimerTriggerExecutionWithOptions(request *ResetTimerTriggerExecutionRequest, runtime *util.RuntimeOptions) (_result *ResetTimerTriggerExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetTimerTriggerExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetTimerTriggerExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetTimerTriggerExecutionRequest
//
// @return ResetTimerTriggerExecutionResponse
func (client *Client) ResetTimerTriggerExecution(request *ResetTimerTriggerExecutionRequest) (_result *ResetTimerTriggerExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetTimerTriggerExecutionResponse{}
	_body, _err := client.ResetTimerTriggerExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetUserExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserExecutionResponse
func (client *Client) ResetUserExecutionWithOptions(request *ResetUserExecutionRequest, runtime *util.RuntimeOptions) (_result *ResetUserExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetUserExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetUserExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetUserExecutionRequest
//
// @return ResetUserExecutionResponse
func (client *Client) ResetUserExecution(request *ResetUserExecutionRequest) (_result *ResetUserExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetUserExecutionResponse{}
	_body, _err := client.ResetUserExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetFlowControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFlowControlResponse
func (client *Client) SetFlowControlWithOptions(request *SetFlowControlRequest, runtime *util.RuntimeOptions) (_result *SetFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Api)) {
		query["Api"] = request.Api
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFlowControl"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFlowControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetFlowControlRequest
//
// @return SetFlowControlResponse
func (client *Client) SetFlowControl(request *SetFlowControlRequest) (_result *SetFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFlowControlResponse{}
	_body, _err := client.SetFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetQuotaResponse
func (client *Client) SetQuotaWithOptions(request *SetQuotaRequest, runtime *util.RuntimeOptions) (_result *SetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		query["QuotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetQuota"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetQuotaRequest
//
// @return SetQuotaResponse
func (client *Client) SetQuota(request *SetQuotaRequest) (_result *SetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQuotaResponse{}
	_body, _err := client.SetQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - TerminateUserExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateUserExecutionResponse
func (client *Client) TerminateUserExecutionWithOptions(request *TerminateUserExecutionRequest, runtime *util.RuntimeOptions) (_result *TerminateUserExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TerminateUserExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TerminateUserExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - TerminateUserExecutionRequest
//
// @return TerminateUserExecutionResponse
func (client *Client) TerminateUserExecution(request *TerminateUserExecutionRequest) (_result *TerminateUserExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TerminateUserExecutionResponse{}
	_body, _err := client.TerminateUserExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateActionResponse
func (client *Client) UpdateActionWithOptions(request *UpdateActionRequest, runtime *util.RuntimeOptions) (_result *UpdateActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		query["ActionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Popularity)) {
		query["Popularity"] = request.Popularity
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAction"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateActionRequest
//
// @return UpdateActionResponse
func (client *Client) UpdateAction(request *UpdateActionRequest) (_result *UpdateActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateActionResponse{}
	_body, _err := client.UpdateActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个已存在的公共参数。
//
// @param request - UpdatePublicParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicParameterResponse
func (client *Client) UpdatePublicParameterWithOptions(request *UpdatePublicParameterRequest, runtime *util.RuntimeOptions) (_result *UpdatePublicParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePublicParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePublicParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个已存在的公共参数。
//
// @param request - UpdatePublicParameterRequest
//
// @return UpdatePublicParameterResponse
func (client *Client) UpdatePublicParameter(request *UpdatePublicParameterRequest) (_result *UpdatePublicParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePublicParameterResponse{}
	_body, _err := client.UpdatePublicParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdatePublicPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicPatchBaselineResponse
func (client *Client) UpdatePublicPatchBaselineWithOptions(request *UpdatePublicPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *UpdatePublicPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalRules)) {
		query["ApprovalRules"] = request.ApprovalRules
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePublicPatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePublicPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdatePublicPatchBaselineRequest
//
// @return UpdatePublicPatchBaselineResponse
func (client *Client) UpdatePublicPatchBaseline(request *UpdatePublicPatchBaselineRequest) (_result *UpdatePublicPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePublicPatchBaselineResponse{}
	_body, _err := client.UpdatePublicPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdatePublicTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicTemplateResponse
func (client *Client) UpdatePublicTemplateWithOptions(request *UpdatePublicTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdatePublicTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Popularity)) {
		query["Popularity"] = request.Popularity
	}

	if !tea.BoolValue(util.IsUnset(request.Publisher)) {
		query["Publisher"] = request.Publisher
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePublicTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePublicTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdatePublicTemplateRequest
//
// @return UpdatePublicTemplateResponse
func (client *Client) UpdatePublicTemplate(request *UpdatePublicTemplateRequest) (_result *UpdatePublicTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePublicTemplateResponse{}
	_body, _err := client.UpdatePublicTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ValidatePublicTemplateContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidatePublicTemplateContentResponse
func (client *Client) ValidatePublicTemplateContentWithOptions(request *ValidatePublicTemplateContentRequest, runtime *util.RuntimeOptions) (_result *ValidatePublicTemplateContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidatePublicTemplateContent"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidatePublicTemplateContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ValidatePublicTemplateContentRequest
//
// @return ValidatePublicTemplateContentResponse
func (client *Client) ValidatePublicTemplateContent(request *ValidatePublicTemplateContentRequest) (_result *ValidatePublicTemplateContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidatePublicTemplateContentResponse{}
	_body, _err := client.ValidatePublicTemplateContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
